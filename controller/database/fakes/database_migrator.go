// This file was generated by counterfeiter
package fakes

import "sync"

type DatabaseMigrator struct {
	MigrateStub        func() (int, error)
	migrateMutex       sync.RWMutex
	migrateArgsForCall []struct{}
	migrateReturns     struct {
		result1 int
		result2 error
	}
	migrateReturnsOnCall map[int]struct {
		result1 int
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *DatabaseMigrator) Migrate() (int, error) {
	fake.migrateMutex.Lock()
	ret, specificReturn := fake.migrateReturnsOnCall[len(fake.migrateArgsForCall)]
	fake.migrateArgsForCall = append(fake.migrateArgsForCall, struct{}{})
	fake.recordInvocation("Migrate", []interface{}{})
	fake.migrateMutex.Unlock()
	if fake.MigrateStub != nil {
		return fake.MigrateStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.migrateReturns.result1, fake.migrateReturns.result2
}

func (fake *DatabaseMigrator) MigrateCallCount() int {
	fake.migrateMutex.RLock()
	defer fake.migrateMutex.RUnlock()
	return len(fake.migrateArgsForCall)
}

func (fake *DatabaseMigrator) MigrateReturns(result1 int, result2 error) {
	fake.MigrateStub = nil
	fake.migrateReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *DatabaseMigrator) MigrateReturnsOnCall(i int, result1 int, result2 error) {
	fake.MigrateStub = nil
	if fake.migrateReturnsOnCall == nil {
		fake.migrateReturnsOnCall = make(map[int]struct {
			result1 int
			result2 error
		})
	}
	fake.migrateReturnsOnCall[i] = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *DatabaseMigrator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.migrateMutex.RLock()
	defer fake.migrateMutex.RUnlock()
	return fake.invocations
}

func (fake *DatabaseMigrator) recordInvocation(key string, args []interface{}) {
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
