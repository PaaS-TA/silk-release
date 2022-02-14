// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"code.cloudfoundry.org/vxlan-policy-agent/enforcer"
)

type RuleEnforcer struct {
	DeleteChainStub        func(string, string) error
	deleteChainMutex       sync.RWMutex
	deleteChainArgsForCall []struct {
		arg1 string
		arg2 string
	}
	deleteChainReturns struct {
		result1 error
	}
	deleteChainReturnsOnCall map[int]struct {
		result1 error
	}
	EnforceRulesAndChainStub        func(enforcer.RulesWithChain) (string, error)
	enforceRulesAndChainMutex       sync.RWMutex
	enforceRulesAndChainArgsForCall []struct {
		arg1 enforcer.RulesWithChain
	}
	enforceRulesAndChainReturns struct {
		result1 string
		result2 error
	}
	enforceRulesAndChainReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *RuleEnforcer) DeleteChain(arg1 string, arg2 string) error {
	fake.deleteChainMutex.Lock()
	ret, specificReturn := fake.deleteChainReturnsOnCall[len(fake.deleteChainArgsForCall)]
	fake.deleteChainArgsForCall = append(fake.deleteChainArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.DeleteChainStub
	fakeReturns := fake.deleteChainReturns
	fake.recordInvocation("DeleteChain", []interface{}{arg1, arg2})
	fake.deleteChainMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *RuleEnforcer) DeleteChainCallCount() int {
	fake.deleteChainMutex.RLock()
	defer fake.deleteChainMutex.RUnlock()
	return len(fake.deleteChainArgsForCall)
}

func (fake *RuleEnforcer) DeleteChainCalls(stub func(string, string) error) {
	fake.deleteChainMutex.Lock()
	defer fake.deleteChainMutex.Unlock()
	fake.DeleteChainStub = stub
}

func (fake *RuleEnforcer) DeleteChainArgsForCall(i int) (string, string) {
	fake.deleteChainMutex.RLock()
	defer fake.deleteChainMutex.RUnlock()
	argsForCall := fake.deleteChainArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *RuleEnforcer) DeleteChainReturns(result1 error) {
	fake.deleteChainMutex.Lock()
	defer fake.deleteChainMutex.Unlock()
	fake.DeleteChainStub = nil
	fake.deleteChainReturns = struct {
		result1 error
	}{result1}
}

func (fake *RuleEnforcer) DeleteChainReturnsOnCall(i int, result1 error) {
	fake.deleteChainMutex.Lock()
	defer fake.deleteChainMutex.Unlock()
	fake.DeleteChainStub = nil
	if fake.deleteChainReturnsOnCall == nil {
		fake.deleteChainReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteChainReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *RuleEnforcer) EnforceRulesAndChain(arg1 enforcer.RulesWithChain) (string, error) {
	fake.enforceRulesAndChainMutex.Lock()
	ret, specificReturn := fake.enforceRulesAndChainReturnsOnCall[len(fake.enforceRulesAndChainArgsForCall)]
	fake.enforceRulesAndChainArgsForCall = append(fake.enforceRulesAndChainArgsForCall, struct {
		arg1 enforcer.RulesWithChain
	}{arg1})
	stub := fake.EnforceRulesAndChainStub
	fakeReturns := fake.enforceRulesAndChainReturns
	fake.recordInvocation("EnforceRulesAndChain", []interface{}{arg1})
	fake.enforceRulesAndChainMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *RuleEnforcer) EnforceRulesAndChainCallCount() int {
	fake.enforceRulesAndChainMutex.RLock()
	defer fake.enforceRulesAndChainMutex.RUnlock()
	return len(fake.enforceRulesAndChainArgsForCall)
}

func (fake *RuleEnforcer) EnforceRulesAndChainCalls(stub func(enforcer.RulesWithChain) (string, error)) {
	fake.enforceRulesAndChainMutex.Lock()
	defer fake.enforceRulesAndChainMutex.Unlock()
	fake.EnforceRulesAndChainStub = stub
}

func (fake *RuleEnforcer) EnforceRulesAndChainArgsForCall(i int) enforcer.RulesWithChain {
	fake.enforceRulesAndChainMutex.RLock()
	defer fake.enforceRulesAndChainMutex.RUnlock()
	argsForCall := fake.enforceRulesAndChainArgsForCall[i]
	return argsForCall.arg1
}

func (fake *RuleEnforcer) EnforceRulesAndChainReturns(result1 string, result2 error) {
	fake.enforceRulesAndChainMutex.Lock()
	defer fake.enforceRulesAndChainMutex.Unlock()
	fake.EnforceRulesAndChainStub = nil
	fake.enforceRulesAndChainReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *RuleEnforcer) EnforceRulesAndChainReturnsOnCall(i int, result1 string, result2 error) {
	fake.enforceRulesAndChainMutex.Lock()
	defer fake.enforceRulesAndChainMutex.Unlock()
	fake.EnforceRulesAndChainStub = nil
	if fake.enforceRulesAndChainReturnsOnCall == nil {
		fake.enforceRulesAndChainReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.enforceRulesAndChainReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *RuleEnforcer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.deleteChainMutex.RLock()
	defer fake.deleteChainMutex.RUnlock()
	fake.enforceRulesAndChainMutex.RLock()
	defer fake.enforceRulesAndChainMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *RuleEnforcer) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}
