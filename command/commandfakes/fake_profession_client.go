// Code generated by counterfeiter. DO NOT EDIT.
package commandfakes

import (
	"sync"

	"github.com/ablease/zinn/command"
)

type FakeProfessionClient struct {
	ProfessionsStub        func() ([]string, error)
	professionsMutex       sync.RWMutex
	professionsArgsForCall []struct{}
	professionsReturns     struct {
		result1 []string
		result2 error
	}
	professionsReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeProfessionClient) Professions() ([]string, error) {
	fake.professionsMutex.Lock()
	ret, specificReturn := fake.professionsReturnsOnCall[len(fake.professionsArgsForCall)]
	fake.professionsArgsForCall = append(fake.professionsArgsForCall, struct{}{})
	fake.recordInvocation("Professions", []interface{}{})
	fake.professionsMutex.Unlock()
	if fake.ProfessionsStub != nil {
		return fake.ProfessionsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.professionsReturns.result1, fake.professionsReturns.result2
}

func (fake *FakeProfessionClient) ProfessionsCallCount() int {
	fake.professionsMutex.RLock()
	defer fake.professionsMutex.RUnlock()
	return len(fake.professionsArgsForCall)
}

func (fake *FakeProfessionClient) ProfessionsReturns(result1 []string, result2 error) {
	fake.ProfessionsStub = nil
	fake.professionsReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeProfessionClient) ProfessionsReturnsOnCall(i int, result1 []string, result2 error) {
	fake.ProfessionsStub = nil
	if fake.professionsReturnsOnCall == nil {
		fake.professionsReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.professionsReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeProfessionClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.professionsMutex.RLock()
	defer fake.professionsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeProfessionClient) recordInvocation(key string, args []interface{}) {
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

var _ command.ProfessionClient = new(FakeProfessionClient)
