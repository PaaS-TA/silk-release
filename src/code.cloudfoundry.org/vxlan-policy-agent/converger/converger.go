package converger

import (
	"fmt"
	"time"

	"code.cloudfoundry.org/vxlan-policy-agent/enforcer"

	"sync"

	"code.cloudfoundry.org/lager"
)

//go:generate counterfeiter -o fakes/planner.go --fake-name Planner . Planner
type Planner interface {
	GetPolicyRulesAndChain() (enforcer.RulesWithChain, error)
	GetASGRulesAndChains() ([]enforcer.RulesWithChain, error)
}

//go:generate counterfeiter -o fakes/rule_enforcer.go --fake-name RuleEnforcer . ruleEnforcer
type ruleEnforcer interface {
	EnforceRulesAndChain(enforcer.RulesWithChain) (string, error)
	DeleteChain(chainName string, parentChain string) error
}

//go:generate counterfeiter -o fakes/metrics_sender.go --fake-name MetricsSender . metricsSender
type metricsSender interface {
	SendDuration(string, time.Duration)
}

type SinglePollCycle struct {
	planners            []Planner
	enforcer            ruleEnforcer
	metricsSender       metricsSender
	logger              lager.Logger
	policyRuleSets      map[enforcer.Chain]enforcer.RulesWithChain
	asgRuleSets         map[enforcer.Chain]enforcer.RulesWithChain
	containerToASGChain map[enforcer.Chain]string
	policyMutex         sync.Locker
	asgMutex            sync.Locker
}

func NewSinglePollCycle(planners []Planner, re ruleEnforcer, ms metricsSender, logger lager.Logger) *SinglePollCycle {
	return &SinglePollCycle{
		planners:      planners,
		enforcer:      re,
		metricsSender: ms,
		logger:        logger,
		policyMutex:   new(sync.Mutex),
		asgMutex:      new(sync.Mutex),
	}
}

const metricEnforceDuration = "iptablesEnforceTime"
const metricPollDuration = "totalPollTime"

const metricASGEnforceDuration = "asgIptablesEnforceTime"
const metricASGPollDuration = "asgTotalPollTime"

func (m *SinglePollCycle) DoPolicyCycle() error {
	m.policyMutex.Lock()

	if m.policyRuleSets == nil {
		m.policyRuleSets = make(map[enforcer.Chain]enforcer.RulesWithChain)
	}

	pollStartTime := time.Now()
	var enforceDuration time.Duration
	for _, p := range m.planners {
		ruleSet, err := p.GetPolicyRulesAndChain()
		if err != nil {
			m.policyMutex.Unlock()
			return fmt.Errorf("get-rules: %s", err)
		}
		enforceStartTime := time.Now()

		oldRuleSet := m.policyRuleSets[ruleSet.Chain]
		if !ruleSet.Equals(oldRuleSet) {
			m.logger.Debug("poll-cycle", lager.Data{
				"message":       "updating iptables rules",
				"num old rules": len(oldRuleSet.Rules),
				"num new rules": len(ruleSet.Rules),
				"old rules":     oldRuleSet,
				"new rules":     ruleSet,
			})
			_, err = m.enforcer.EnforceRulesAndChain(ruleSet)
			if err != nil {
				m.policyMutex.Unlock()
				return fmt.Errorf("enforce: %s", err)
			}
			m.policyRuleSets[ruleSet.Chain] = ruleSet
		}

		enforceDuration += time.Now().Sub(enforceStartTime)
	}

	m.policyMutex.Unlock()

	pollDuration := time.Now().Sub(pollStartTime)
	m.metricsSender.SendDuration(metricEnforceDuration, enforceDuration)
	m.metricsSender.SendDuration(metricPollDuration, pollDuration)

	return nil
}

func (m *SinglePollCycle) DoASGCycle() error {
	m.asgMutex.Lock()

	if m.asgRuleSets == nil {
		m.asgRuleSets = make(map[enforcer.Chain]enforcer.RulesWithChain)
	}
	if m.containerToASGChain == nil {
		m.containerToASGChain = make(map[enforcer.Chain]string)
	}

	pollStartTime := time.Now()
	var enforceDuration time.Duration

	for _, p := range m.planners {
		asgrulesets, err := p.GetASGRulesAndChains()
		if err != nil {
			m.asgMutex.Unlock()
			return fmt.Errorf("get-asg-rules: %s", err)
		}
		enforceStartTime := time.Now()

		for _, ruleset := range asgrulesets {
			oldRuleSet := m.asgRuleSets[ruleset.Chain]
			if !ruleset.Equals(oldRuleSet) {
				m.logger.Debug("poll-cycle-asg", lager.Data{
					"message":       "updating iptables rules",
					"num old rules": len(oldRuleSet.Rules),
					"num new rules": len(ruleset.Rules),
					"old rules":     oldRuleSet,
					"new rules":     ruleset,
				})
				chain, err := m.enforcer.EnforceRulesAndChain(ruleset)
				m.containerToASGChain[ruleset.Chain] = chain

				if err != nil {
					m.asgMutex.Unlock()
					return fmt.Errorf("enforce-asg: %s", err)
				}
				m.asgRuleSets[ruleset.Chain] = ruleset
			}
		}
		for parentChain, _ := range m.containerToASGChain {
			if _, ok := m.asgRuleSets[parentChain]; !ok {
				err := m.enforcer.DeleteChain(parentChain.Table, m.containerToASGChain[parentChain])
				if err != nil {
					m.asgMutex.Unlock()
					return fmt.Errorf("clean-up-orphaned-asg-chains: %s", err)
				}
				delete(m.containerToASGChain, parentChain)
				delete(m.asgRuleSets, parentChain)
			}
		}

		enforceDuration += time.Now().Sub(enforceStartTime)
	}

	m.asgMutex.Unlock()

	pollDuration := time.Now().Sub(pollStartTime)
	m.metricsSender.SendDuration(metricASGEnforceDuration, enforceDuration)
	m.metricsSender.SendDuration(metricASGPollDuration, pollDuration)

	return nil
}
