// Code generated by counterfeiter. DO NOT EDIT.
package apifakes

import (
	"context"
	"sync"

	"github.com/dembygenesis/local.tools/internal/model"
)

type FakeClickTrackerService struct {
	CreateClickTrackerStub        func(context.Context, *model.CreateClickTracker) (*model.ClickTracker, error)
	createClickTrackerMutex       sync.RWMutex
	createClickTrackerArgsForCall []struct {
		arg1 context.Context
		arg2 *model.CreateClickTracker
	}
	createClickTrackerReturns struct {
		result1 *model.ClickTracker
		result2 error
	}
	createClickTrackerReturnsOnCall map[int]struct {
		result1 *model.ClickTracker
		result2 error
	}
	ListClickTrackersStub        func(context.Context, *model.ClickTrackerFilters) (*model.PaginatedClickTrackers, error)
	listClickTrackersMutex       sync.RWMutex
	listClickTrackersArgsForCall []struct {
		arg1 context.Context
		arg2 *model.ClickTrackerFilters
	}
	listClickTrackersReturns struct {
		result1 *model.PaginatedClickTrackers
		result2 error
	}
	listClickTrackersReturnsOnCall map[int]struct {
		result1 *model.PaginatedClickTrackers
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeClickTrackerService) CreateClickTracker(arg1 context.Context, arg2 *model.CreateClickTracker) (*model.ClickTracker, error) {
	fake.createClickTrackerMutex.Lock()
	ret, specificReturn := fake.createClickTrackerReturnsOnCall[len(fake.createClickTrackerArgsForCall)]
	fake.createClickTrackerArgsForCall = append(fake.createClickTrackerArgsForCall, struct {
		arg1 context.Context
		arg2 *model.CreateClickTracker
	}{arg1, arg2})
	stub := fake.CreateClickTrackerStub
	fakeReturns := fake.createClickTrackerReturns
	fake.recordInvocation("CreateClickTracker", []interface{}{arg1, arg2})
	fake.createClickTrackerMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeClickTrackerService) CreateClickTrackerCallCount() int {
	fake.createClickTrackerMutex.RLock()
	defer fake.createClickTrackerMutex.RUnlock()
	return len(fake.createClickTrackerArgsForCall)
}

func (fake *FakeClickTrackerService) CreateClickTrackerCalls(stub func(context.Context, *model.CreateClickTracker) (*model.ClickTracker, error)) {
	fake.createClickTrackerMutex.Lock()
	defer fake.createClickTrackerMutex.Unlock()
	fake.CreateClickTrackerStub = stub
}

func (fake *FakeClickTrackerService) CreateClickTrackerArgsForCall(i int) (context.Context, *model.CreateClickTracker) {
	fake.createClickTrackerMutex.RLock()
	defer fake.createClickTrackerMutex.RUnlock()
	argsForCall := fake.createClickTrackerArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeClickTrackerService) CreateClickTrackerReturns(result1 *model.ClickTracker, result2 error) {
	fake.createClickTrackerMutex.Lock()
	defer fake.createClickTrackerMutex.Unlock()
	fake.CreateClickTrackerStub = nil
	fake.createClickTrackerReturns = struct {
		result1 *model.ClickTracker
		result2 error
	}{result1, result2}
}

func (fake *FakeClickTrackerService) CreateClickTrackerReturnsOnCall(i int, result1 *model.ClickTracker, result2 error) {
	fake.createClickTrackerMutex.Lock()
	defer fake.createClickTrackerMutex.Unlock()
	fake.CreateClickTrackerStub = nil
	if fake.createClickTrackerReturnsOnCall == nil {
		fake.createClickTrackerReturnsOnCall = make(map[int]struct {
			result1 *model.ClickTracker
			result2 error
		})
	}
	fake.createClickTrackerReturnsOnCall[i] = struct {
		result1 *model.ClickTracker
		result2 error
	}{result1, result2}
}

func (fake *FakeClickTrackerService) ListClickTrackers(arg1 context.Context, arg2 *model.ClickTrackerFilters) (*model.PaginatedClickTrackers, error) {
	fake.listClickTrackersMutex.Lock()
	ret, specificReturn := fake.listClickTrackersReturnsOnCall[len(fake.listClickTrackersArgsForCall)]
	fake.listClickTrackersArgsForCall = append(fake.listClickTrackersArgsForCall, struct {
		arg1 context.Context
		arg2 *model.ClickTrackerFilters
	}{arg1, arg2})
	stub := fake.ListClickTrackersStub
	fakeReturns := fake.listClickTrackersReturns
	fake.recordInvocation("ListClickTrackers", []interface{}{arg1, arg2})
	fake.listClickTrackersMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeClickTrackerService) ListClickTrackersCallCount() int {
	fake.listClickTrackersMutex.RLock()
	defer fake.listClickTrackersMutex.RUnlock()
	return len(fake.listClickTrackersArgsForCall)
}

func (fake *FakeClickTrackerService) ListClickTrackersCalls(stub func(context.Context, *model.ClickTrackerFilters) (*model.PaginatedClickTrackers, error)) {
	fake.listClickTrackersMutex.Lock()
	defer fake.listClickTrackersMutex.Unlock()
	fake.ListClickTrackersStub = stub
}

func (fake *FakeClickTrackerService) ListClickTrackersArgsForCall(i int) (context.Context, *model.ClickTrackerFilters) {
	fake.listClickTrackersMutex.RLock()
	defer fake.listClickTrackersMutex.RUnlock()
	argsForCall := fake.listClickTrackersArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeClickTrackerService) ListClickTrackersReturns(result1 *model.PaginatedClickTrackers, result2 error) {
	fake.listClickTrackersMutex.Lock()
	defer fake.listClickTrackersMutex.Unlock()
	fake.ListClickTrackersStub = nil
	fake.listClickTrackersReturns = struct {
		result1 *model.PaginatedClickTrackers
		result2 error
	}{result1, result2}
}

func (fake *FakeClickTrackerService) ListClickTrackersReturnsOnCall(i int, result1 *model.PaginatedClickTrackers, result2 error) {
	fake.listClickTrackersMutex.Lock()
	defer fake.listClickTrackersMutex.Unlock()
	fake.ListClickTrackersStub = nil
	if fake.listClickTrackersReturnsOnCall == nil {
		fake.listClickTrackersReturnsOnCall = make(map[int]struct {
			result1 *model.PaginatedClickTrackers
			result2 error
		})
	}
	fake.listClickTrackersReturnsOnCall[i] = struct {
		result1 *model.PaginatedClickTrackers
		result2 error
	}{result1, result2}
}

func (fake *FakeClickTrackerService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createClickTrackerMutex.RLock()
	defer fake.createClickTrackerMutex.RUnlock()
	fake.listClickTrackersMutex.RLock()
	defer fake.listClickTrackersMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeClickTrackerService) recordInvocation(key string, args []interface{}) {
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
