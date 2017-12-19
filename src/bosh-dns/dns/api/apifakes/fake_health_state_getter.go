// Code generated by counterfeiter. DO NOT EDIT.
package apifakes

import (
	"bosh-dns/dns/api"
	"sync"
)

type FakeHealthStateGetter struct {
	HealthStateStub        func(ip string) string
	healthStateMutex       sync.RWMutex
	healthStateArgsForCall []struct {
		ip string
	}
	healthStateReturns struct {
		result1 string
	}
	healthStateReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeHealthStateGetter) HealthState(ip string) string {
	fake.healthStateMutex.Lock()
	ret, specificReturn := fake.healthStateReturnsOnCall[len(fake.healthStateArgsForCall)]
	fake.healthStateArgsForCall = append(fake.healthStateArgsForCall, struct {
		ip string
	}{ip})
	fake.recordInvocation("HealthState", []interface{}{ip})
	fake.healthStateMutex.Unlock()
	if fake.HealthStateStub != nil {
		return fake.HealthStateStub(ip)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.healthStateReturns.result1
}

func (fake *FakeHealthStateGetter) HealthStateCallCount() int {
	fake.healthStateMutex.RLock()
	defer fake.healthStateMutex.RUnlock()
	return len(fake.healthStateArgsForCall)
}

func (fake *FakeHealthStateGetter) HealthStateArgsForCall(i int) string {
	fake.healthStateMutex.RLock()
	defer fake.healthStateMutex.RUnlock()
	return fake.healthStateArgsForCall[i].ip
}

func (fake *FakeHealthStateGetter) HealthStateReturns(result1 string) {
	fake.HealthStateStub = nil
	fake.healthStateReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeHealthStateGetter) HealthStateReturnsOnCall(i int, result1 string) {
	fake.HealthStateStub = nil
	if fake.healthStateReturnsOnCall == nil {
		fake.healthStateReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.healthStateReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeHealthStateGetter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.healthStateMutex.RLock()
	defer fake.healthStateMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeHealthStateGetter) recordInvocation(key string, args []interface{}) {
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

var _ api.HealthStateGetter = new(FakeHealthStateGetter)