// Code generated by counterfeiter. DO NOT EDIT.
package commandfakes

import (
	"sync"

	"github.com/ablease/zinn/command"
)

type FakeMasteriesClient struct {
	MasteriesStub        func() ([]string, error)
	masteriesMutex       sync.RWMutex
	masteriesArgsForCall []struct{}
	masteriesReturns     struct {
		result1 []string
		result2 error
	}
	masteriesReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeMasteriesClient) Masteries() ([]string, error) {
	fake.masteriesMutex.Lock()
	ret, specificReturn := fake.masteriesReturnsOnCall[len(fake.masteriesArgsForCall)]
	fake.masteriesArgsForCall = append(fake.masteriesArgsForCall, struct{}{})
	fake.recordInvocation("Masteries", []interface{}{})
	fake.masteriesMutex.Unlock()
	if fake.MasteriesStub != nil {
		return fake.MasteriesStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.masteriesReturns.result1, fake.masteriesReturns.result2
}

func (fake *FakeMasteriesClient) MasteriesCallCount() int {
	fake.masteriesMutex.RLock()
	defer fake.masteriesMutex.RUnlock()
	return len(fake.masteriesArgsForCall)
}

func (fake *FakeMasteriesClient) MasteriesReturns(result1 []string, result2 error) {
	fake.MasteriesStub = nil
	fake.masteriesReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeMasteriesClient) MasteriesReturnsOnCall(i int, result1 []string, result2 error) {
	fake.MasteriesStub = nil
	if fake.masteriesReturnsOnCall == nil {
		fake.masteriesReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.masteriesReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeMasteriesClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.masteriesMutex.RLock()
	defer fake.masteriesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeMasteriesClient) recordInvocation(key string, args []interface{}) {
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

var _ command.MasteriesClient = new(FakeMasteriesClient)