// This file was generated by counterfeiter
package serverfakes

import (
	"sync"

	"github.com/cloudfoundry/dns-release/src/dns/server"
)

type FakeListenAndServer struct {
	ListenAndServeStub        func() error
	listenAndServeMutex       sync.RWMutex
	listenAndServeArgsForCall []struct{}
	listenAndServeReturns     struct {
		result1 error
	}
	listenAndServeReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeListenAndServer) ListenAndServe() error {
	fake.listenAndServeMutex.Lock()
	ret, specificReturn := fake.listenAndServeReturnsOnCall[len(fake.listenAndServeArgsForCall)]
	fake.listenAndServeArgsForCall = append(fake.listenAndServeArgsForCall, struct{}{})
	fake.recordInvocation("ListenAndServe", []interface{}{})
	fake.listenAndServeMutex.Unlock()
	if fake.ListenAndServeStub != nil {
		return fake.ListenAndServeStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.listenAndServeReturns.result1
}

func (fake *FakeListenAndServer) ListenAndServeCallCount() int {
	fake.listenAndServeMutex.RLock()
	defer fake.listenAndServeMutex.RUnlock()
	return len(fake.listenAndServeArgsForCall)
}

func (fake *FakeListenAndServer) ListenAndServeReturns(result1 error) {
	fake.ListenAndServeStub = nil
	fake.listenAndServeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeListenAndServer) ListenAndServeReturnsOnCall(i int, result1 error) {
	fake.ListenAndServeStub = nil
	if fake.listenAndServeReturnsOnCall == nil {
		fake.listenAndServeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.listenAndServeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeListenAndServer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.listenAndServeMutex.RLock()
	defer fake.listenAndServeMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeListenAndServer) recordInvocation(key string, args []interface{}) {
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

var _ server.ListenAndServer = new(FakeListenAndServer)