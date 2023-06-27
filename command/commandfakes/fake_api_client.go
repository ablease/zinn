// Code generated by counterfeiter. DO NOT EDIT.
package commandfakes

import (
	"sync"

	"github.com/ablease/zinn/api"
	"github.com/ablease/zinn/command"
)

type FakeApiClient struct {
	AchievementIDsStub        func() ([]int, error)
	achievementIDsMutex       sync.RWMutex
	achievementIDsArgsForCall []struct {
	}
	achievementIDsReturns struct {
		result1 []int
		result2 error
	}
	achievementIDsReturnsOnCall map[int]struct {
		result1 []int
		result2 error
	}
	GetMasteryIDsStub        func() ([]int, error)
	getMasteryIDsMutex       sync.RWMutex
	getMasteryIDsArgsForCall []struct {
	}
	getMasteryIDsReturns struct {
		result1 []int
		result2 error
	}
	getMasteryIDsReturnsOnCall map[int]struct {
		result1 []int
		result2 error
	}
	MasteriesStub        func([]int) ([]api.Mastery, error)
	masteriesMutex       sync.RWMutex
	masteriesArgsForCall []struct {
		arg1 []int
	}
	masteriesReturns struct {
		result1 []api.Mastery
		result2 error
	}
	masteriesReturnsOnCall map[int]struct {
		result1 []api.Mastery
		result2 error
	}
	ProfessionsStub        func() ([]string, error)
	professionsMutex       sync.RWMutex
	professionsArgsForCall []struct {
	}
	professionsReturns struct {
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

func (fake *FakeApiClient) AchievementIDs() ([]int, error) {
	fake.achievementIDsMutex.Lock()
	ret, specificReturn := fake.achievementIDsReturnsOnCall[len(fake.achievementIDsArgsForCall)]
	fake.achievementIDsArgsForCall = append(fake.achievementIDsArgsForCall, struct {
	}{})
	stub := fake.AchievementIDsStub
	fakeReturns := fake.achievementIDsReturns
	fake.recordInvocation("AchievementIDs", []interface{}{})
	fake.achievementIDsMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeApiClient) AchievementIDsCallCount() int {
	fake.achievementIDsMutex.RLock()
	defer fake.achievementIDsMutex.RUnlock()
	return len(fake.achievementIDsArgsForCall)
}

func (fake *FakeApiClient) AchievementIDsCalls(stub func() ([]int, error)) {
	fake.achievementIDsMutex.Lock()
	defer fake.achievementIDsMutex.Unlock()
	fake.AchievementIDsStub = stub
}

func (fake *FakeApiClient) AchievementIDsReturns(result1 []int, result2 error) {
	fake.achievementIDsMutex.Lock()
	defer fake.achievementIDsMutex.Unlock()
	fake.AchievementIDsStub = nil
	fake.achievementIDsReturns = struct {
		result1 []int
		result2 error
	}{result1, result2}
}

func (fake *FakeApiClient) AchievementIDsReturnsOnCall(i int, result1 []int, result2 error) {
	fake.achievementIDsMutex.Lock()
	defer fake.achievementIDsMutex.Unlock()
	fake.AchievementIDsStub = nil
	if fake.achievementIDsReturnsOnCall == nil {
		fake.achievementIDsReturnsOnCall = make(map[int]struct {
			result1 []int
			result2 error
		})
	}
	fake.achievementIDsReturnsOnCall[i] = struct {
		result1 []int
		result2 error
	}{result1, result2}
}

func (fake *FakeApiClient) GetMasteryIDs() ([]int, error) {
	fake.getMasteryIDsMutex.Lock()
	ret, specificReturn := fake.getMasteryIDsReturnsOnCall[len(fake.getMasteryIDsArgsForCall)]
	fake.getMasteryIDsArgsForCall = append(fake.getMasteryIDsArgsForCall, struct {
	}{})
	stub := fake.GetMasteryIDsStub
	fakeReturns := fake.getMasteryIDsReturns
	fake.recordInvocation("GetMasteryIDs", []interface{}{})
	fake.getMasteryIDsMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeApiClient) GetMasteryIDsCallCount() int {
	fake.getMasteryIDsMutex.RLock()
	defer fake.getMasteryIDsMutex.RUnlock()
	return len(fake.getMasteryIDsArgsForCall)
}

func (fake *FakeApiClient) GetMasteryIDsCalls(stub func() ([]int, error)) {
	fake.getMasteryIDsMutex.Lock()
	defer fake.getMasteryIDsMutex.Unlock()
	fake.GetMasteryIDsStub = stub
}

func (fake *FakeApiClient) GetMasteryIDsReturns(result1 []int, result2 error) {
	fake.getMasteryIDsMutex.Lock()
	defer fake.getMasteryIDsMutex.Unlock()
	fake.GetMasteryIDsStub = nil
	fake.getMasteryIDsReturns = struct {
		result1 []int
		result2 error
	}{result1, result2}
}

func (fake *FakeApiClient) GetMasteryIDsReturnsOnCall(i int, result1 []int, result2 error) {
	fake.getMasteryIDsMutex.Lock()
	defer fake.getMasteryIDsMutex.Unlock()
	fake.GetMasteryIDsStub = nil
	if fake.getMasteryIDsReturnsOnCall == nil {
		fake.getMasteryIDsReturnsOnCall = make(map[int]struct {
			result1 []int
			result2 error
		})
	}
	fake.getMasteryIDsReturnsOnCall[i] = struct {
		result1 []int
		result2 error
	}{result1, result2}
}

func (fake *FakeApiClient) Masteries(arg1 []int) ([]api.Mastery, error) {
	var arg1Copy []int
	if arg1 != nil {
		arg1Copy = make([]int, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.masteriesMutex.Lock()
	ret, specificReturn := fake.masteriesReturnsOnCall[len(fake.masteriesArgsForCall)]
	fake.masteriesArgsForCall = append(fake.masteriesArgsForCall, struct {
		arg1 []int
	}{arg1Copy})
	stub := fake.MasteriesStub
	fakeReturns := fake.masteriesReturns
	fake.recordInvocation("Masteries", []interface{}{arg1Copy})
	fake.masteriesMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeApiClient) MasteriesCallCount() int {
	fake.masteriesMutex.RLock()
	defer fake.masteriesMutex.RUnlock()
	return len(fake.masteriesArgsForCall)
}

func (fake *FakeApiClient) MasteriesCalls(stub func([]int) ([]api.Mastery, error)) {
	fake.masteriesMutex.Lock()
	defer fake.masteriesMutex.Unlock()
	fake.MasteriesStub = stub
}

func (fake *FakeApiClient) MasteriesArgsForCall(i int) []int {
	fake.masteriesMutex.RLock()
	defer fake.masteriesMutex.RUnlock()
	argsForCall := fake.masteriesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeApiClient) MasteriesReturns(result1 []api.Mastery, result2 error) {
	fake.masteriesMutex.Lock()
	defer fake.masteriesMutex.Unlock()
	fake.MasteriesStub = nil
	fake.masteriesReturns = struct {
		result1 []api.Mastery
		result2 error
	}{result1, result2}
}

func (fake *FakeApiClient) MasteriesReturnsOnCall(i int, result1 []api.Mastery, result2 error) {
	fake.masteriesMutex.Lock()
	defer fake.masteriesMutex.Unlock()
	fake.MasteriesStub = nil
	if fake.masteriesReturnsOnCall == nil {
		fake.masteriesReturnsOnCall = make(map[int]struct {
			result1 []api.Mastery
			result2 error
		})
	}
	fake.masteriesReturnsOnCall[i] = struct {
		result1 []api.Mastery
		result2 error
	}{result1, result2}
}

func (fake *FakeApiClient) Professions() ([]string, error) {
	fake.professionsMutex.Lock()
	ret, specificReturn := fake.professionsReturnsOnCall[len(fake.professionsArgsForCall)]
	fake.professionsArgsForCall = append(fake.professionsArgsForCall, struct {
	}{})
	stub := fake.ProfessionsStub
	fakeReturns := fake.professionsReturns
	fake.recordInvocation("Professions", []interface{}{})
	fake.professionsMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeApiClient) ProfessionsCallCount() int {
	fake.professionsMutex.RLock()
	defer fake.professionsMutex.RUnlock()
	return len(fake.professionsArgsForCall)
}

func (fake *FakeApiClient) ProfessionsCalls(stub func() ([]string, error)) {
	fake.professionsMutex.Lock()
	defer fake.professionsMutex.Unlock()
	fake.ProfessionsStub = stub
}

func (fake *FakeApiClient) ProfessionsReturns(result1 []string, result2 error) {
	fake.professionsMutex.Lock()
	defer fake.professionsMutex.Unlock()
	fake.ProfessionsStub = nil
	fake.professionsReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeApiClient) ProfessionsReturnsOnCall(i int, result1 []string, result2 error) {
	fake.professionsMutex.Lock()
	defer fake.professionsMutex.Unlock()
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

func (fake *FakeApiClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.achievementIDsMutex.RLock()
	defer fake.achievementIDsMutex.RUnlock()
	fake.getMasteryIDsMutex.RLock()
	defer fake.getMasteryIDsMutex.RUnlock()
	fake.masteriesMutex.RLock()
	defer fake.masteriesMutex.RUnlock()
	fake.professionsMutex.RLock()
	defer fake.professionsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeApiClient) recordInvocation(key string, args []interface{}) {
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

var _ command.ApiClient = new(FakeApiClient)
