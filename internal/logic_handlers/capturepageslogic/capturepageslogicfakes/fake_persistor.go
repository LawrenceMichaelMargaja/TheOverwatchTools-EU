// Code generated by counterfeiter. DO NOT EDIT.
package capturepageslogicfakes

import (
	"context"
	"sync"

	"github.com/dembygenesis/local.tools/internal/model"
	"github.com/dembygenesis/local.tools/internal/persistence"
)

type FakePersistor struct {
	CreateCapturePagesStub        func(context.Context, persistence.TransactionHandler, *model.CapturePages) (*model.CapturePages, error)
	createCapturePagesMutex       sync.RWMutex
	createCapturePagesArgsForCall []struct {
		arg1 context.Context
		arg2 persistence.TransactionHandler
		arg3 *model.CapturePages
	}
	createCapturePagesReturns struct {
		result1 *model.CapturePages
		result2 error
	}
	createCapturePagesReturnsOnCall map[int]struct {
		result1 *model.CapturePages
		result2 error
	}
	DeleteCapturePagesStub        func(context.Context, persistence.TransactionHandler, int) error
	deleteCapturePagesMutex       sync.RWMutex
	deleteCapturePagesArgsForCall []struct {
		arg1 context.Context
		arg2 persistence.TransactionHandler
		arg3 int
	}
	deleteCapturePagesReturns struct {
		result1 error
	}
	deleteCapturePagesReturnsOnCall map[int]struct {
		result1 error
	}
	GetCapturePageTypeByIdStub        func(context.Context, persistence.TransactionHandler, int) (*model.CapturePageType, error)
	getCapturePageTypeByIdMutex       sync.RWMutex
	getCapturePageTypeByIdArgsForCall []struct {
		arg1 context.Context
		arg2 persistence.TransactionHandler
		arg3 int
	}
	getCapturePageTypeByIdReturns struct {
		result1 *model.CapturePageType
		result2 error
	}
	getCapturePageTypeByIdReturnsOnCall map[int]struct {
		result1 *model.CapturePageType
		result2 error
	}
	GetCapturePagesStub        func(context.Context, persistence.TransactionHandler, *model.CapturePagesFilters) (*model.PaginatedCapturePages, error)
	getCapturePagesMutex       sync.RWMutex
	getCapturePagesArgsForCall []struct {
		arg1 context.Context
		arg2 persistence.TransactionHandler
		arg3 *model.CapturePagesFilters
	}
	getCapturePagesReturns struct {
		result1 *model.PaginatedCapturePages
		result2 error
	}
	getCapturePagesReturnsOnCall map[int]struct {
		result1 *model.PaginatedCapturePages
		result2 error
	}
	GetCapturePagesByIdStub        func(context.Context, persistence.TransactionHandler, int) (*model.CapturePages, error)
	getCapturePagesByIdMutex       sync.RWMutex
	getCapturePagesByIdArgsForCall []struct {
		arg1 context.Context
		arg2 persistence.TransactionHandler
		arg3 int
	}
	getCapturePagesByIdReturns struct {
		result1 *model.CapturePages
		result2 error
	}
	getCapturePagesByIdReturnsOnCall map[int]struct {
		result1 *model.CapturePages
		result2 error
	}
	RestoreCapturePagesStub        func(context.Context, persistence.TransactionHandler, int) error
	restoreCapturePagesMutex       sync.RWMutex
	restoreCapturePagesArgsForCall []struct {
		arg1 context.Context
		arg2 persistence.TransactionHandler
		arg3 int
	}
	restoreCapturePagesReturns struct {
		result1 error
	}
	restoreCapturePagesReturnsOnCall map[int]struct {
		result1 error
	}
	UpdateCapturePagesStub        func(context.Context, persistence.TransactionHandler, *model.UpdateCapturePages) (*model.CapturePages, error)
	updateCapturePagesMutex       sync.RWMutex
	updateCapturePagesArgsForCall []struct {
		arg1 context.Context
		arg2 persistence.TransactionHandler
		arg3 *model.UpdateCapturePages
	}
	updateCapturePagesReturns struct {
		result1 *model.CapturePages
		result2 error
	}
	updateCapturePagesReturnsOnCall map[int]struct {
		result1 *model.CapturePages
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePersistor) CreateCapturePages(arg1 context.Context, arg2 persistence.TransactionHandler, arg3 *model.CapturePages) (*model.CapturePages, error) {
	fake.createCapturePagesMutex.Lock()
	ret, specificReturn := fake.createCapturePagesReturnsOnCall[len(fake.createCapturePagesArgsForCall)]
	fake.createCapturePagesArgsForCall = append(fake.createCapturePagesArgsForCall, struct {
		arg1 context.Context
		arg2 persistence.TransactionHandler
		arg3 *model.CapturePages
	}{arg1, arg2, arg3})
	stub := fake.CreateCapturePagesStub
	fakeReturns := fake.createCapturePagesReturns
	fake.recordInvocation("CreateCapturePages", []interface{}{arg1, arg2, arg3})
	fake.createCapturePagesMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePersistor) CreateCapturePagesCallCount() int {
	fake.createCapturePagesMutex.RLock()
	defer fake.createCapturePagesMutex.RUnlock()
	return len(fake.createCapturePagesArgsForCall)
}

func (fake *FakePersistor) CreateCapturePagesCalls(stub func(context.Context, persistence.TransactionHandler, *model.CapturePages) (*model.CapturePages, error)) {
	fake.createCapturePagesMutex.Lock()
	defer fake.createCapturePagesMutex.Unlock()
	fake.CreateCapturePagesStub = stub
}

func (fake *FakePersistor) CreateCapturePagesArgsForCall(i int) (context.Context, persistence.TransactionHandler, *model.CapturePages) {
	fake.createCapturePagesMutex.RLock()
	defer fake.createCapturePagesMutex.RUnlock()
	argsForCall := fake.createCapturePagesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakePersistor) CreateCapturePagesReturns(result1 *model.CapturePages, result2 error) {
	fake.createCapturePagesMutex.Lock()
	defer fake.createCapturePagesMutex.Unlock()
	fake.CreateCapturePagesStub = nil
	fake.createCapturePagesReturns = struct {
		result1 *model.CapturePages
		result2 error
	}{result1, result2}
}

func (fake *FakePersistor) CreateCapturePagesReturnsOnCall(i int, result1 *model.CapturePages, result2 error) {
	fake.createCapturePagesMutex.Lock()
	defer fake.createCapturePagesMutex.Unlock()
	fake.CreateCapturePagesStub = nil
	if fake.createCapturePagesReturnsOnCall == nil {
		fake.createCapturePagesReturnsOnCall = make(map[int]struct {
			result1 *model.CapturePages
			result2 error
		})
	}
	fake.createCapturePagesReturnsOnCall[i] = struct {
		result1 *model.CapturePages
		result2 error
	}{result1, result2}
}

func (fake *FakePersistor) DeleteCapturePages(arg1 context.Context, arg2 persistence.TransactionHandler, arg3 int) error {
	fake.deleteCapturePagesMutex.Lock()
	ret, specificReturn := fake.deleteCapturePagesReturnsOnCall[len(fake.deleteCapturePagesArgsForCall)]
	fake.deleteCapturePagesArgsForCall = append(fake.deleteCapturePagesArgsForCall, struct {
		arg1 context.Context
		arg2 persistence.TransactionHandler
		arg3 int
	}{arg1, arg2, arg3})
	stub := fake.DeleteCapturePagesStub
	fakeReturns := fake.deleteCapturePagesReturns
	fake.recordInvocation("DeleteCapturePages", []interface{}{arg1, arg2, arg3})
	fake.deleteCapturePagesMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakePersistor) DeleteCapturePagesCallCount() int {
	fake.deleteCapturePagesMutex.RLock()
	defer fake.deleteCapturePagesMutex.RUnlock()
	return len(fake.deleteCapturePagesArgsForCall)
}

func (fake *FakePersistor) DeleteCapturePagesCalls(stub func(context.Context, persistence.TransactionHandler, int) error) {
	fake.deleteCapturePagesMutex.Lock()
	defer fake.deleteCapturePagesMutex.Unlock()
	fake.DeleteCapturePagesStub = stub
}

func (fake *FakePersistor) DeleteCapturePagesArgsForCall(i int) (context.Context, persistence.TransactionHandler, int) {
	fake.deleteCapturePagesMutex.RLock()
	defer fake.deleteCapturePagesMutex.RUnlock()
	argsForCall := fake.deleteCapturePagesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakePersistor) DeleteCapturePagesReturns(result1 error) {
	fake.deleteCapturePagesMutex.Lock()
	defer fake.deleteCapturePagesMutex.Unlock()
	fake.DeleteCapturePagesStub = nil
	fake.deleteCapturePagesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePersistor) DeleteCapturePagesReturnsOnCall(i int, result1 error) {
	fake.deleteCapturePagesMutex.Lock()
	defer fake.deleteCapturePagesMutex.Unlock()
	fake.DeleteCapturePagesStub = nil
	if fake.deleteCapturePagesReturnsOnCall == nil {
		fake.deleteCapturePagesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteCapturePagesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakePersistor) GetCapturePageTypeById(arg1 context.Context, arg2 persistence.TransactionHandler, arg3 int) (*model.CapturePageType, error) {
	fake.getCapturePageTypeByIdMutex.Lock()
	ret, specificReturn := fake.getCapturePageTypeByIdReturnsOnCall[len(fake.getCapturePageTypeByIdArgsForCall)]
	fake.getCapturePageTypeByIdArgsForCall = append(fake.getCapturePageTypeByIdArgsForCall, struct {
		arg1 context.Context
		arg2 persistence.TransactionHandler
		arg3 int
	}{arg1, arg2, arg3})
	stub := fake.GetCapturePageTypeByIdStub
	fakeReturns := fake.getCapturePageTypeByIdReturns
	fake.recordInvocation("GetCapturePageTypeById", []interface{}{arg1, arg2, arg3})
	fake.getCapturePageTypeByIdMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePersistor) GetCapturePageTypeByIdCallCount() int {
	fake.getCapturePageTypeByIdMutex.RLock()
	defer fake.getCapturePageTypeByIdMutex.RUnlock()
	return len(fake.getCapturePageTypeByIdArgsForCall)
}

func (fake *FakePersistor) GetCapturePageTypeByIdCalls(stub func(context.Context, persistence.TransactionHandler, int) (*model.CapturePageType, error)) {
	fake.getCapturePageTypeByIdMutex.Lock()
	defer fake.getCapturePageTypeByIdMutex.Unlock()
	fake.GetCapturePageTypeByIdStub = stub
}

func (fake *FakePersistor) GetCapturePageTypeByIdArgsForCall(i int) (context.Context, persistence.TransactionHandler, int) {
	fake.getCapturePageTypeByIdMutex.RLock()
	defer fake.getCapturePageTypeByIdMutex.RUnlock()
	argsForCall := fake.getCapturePageTypeByIdArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakePersistor) GetCapturePageTypeByIdReturns(result1 *model.CapturePageType, result2 error) {
	fake.getCapturePageTypeByIdMutex.Lock()
	defer fake.getCapturePageTypeByIdMutex.Unlock()
	fake.GetCapturePageTypeByIdStub = nil
	fake.getCapturePageTypeByIdReturns = struct {
		result1 *model.CapturePageType
		result2 error
	}{result1, result2}
}

func (fake *FakePersistor) GetCapturePageTypeByIdReturnsOnCall(i int, result1 *model.CapturePageType, result2 error) {
	fake.getCapturePageTypeByIdMutex.Lock()
	defer fake.getCapturePageTypeByIdMutex.Unlock()
	fake.GetCapturePageTypeByIdStub = nil
	if fake.getCapturePageTypeByIdReturnsOnCall == nil {
		fake.getCapturePageTypeByIdReturnsOnCall = make(map[int]struct {
			result1 *model.CapturePageType
			result2 error
		})
	}
	fake.getCapturePageTypeByIdReturnsOnCall[i] = struct {
		result1 *model.CapturePageType
		result2 error
	}{result1, result2}
}

func (fake *FakePersistor) GetCapturePages(arg1 context.Context, arg2 persistence.TransactionHandler, arg3 *model.CapturePagesFilters) (*model.PaginatedCapturePages, error) {
	fake.getCapturePagesMutex.Lock()
	ret, specificReturn := fake.getCapturePagesReturnsOnCall[len(fake.getCapturePagesArgsForCall)]
	fake.getCapturePagesArgsForCall = append(fake.getCapturePagesArgsForCall, struct {
		arg1 context.Context
		arg2 persistence.TransactionHandler
		arg3 *model.CapturePagesFilters
	}{arg1, arg2, arg3})
	stub := fake.GetCapturePagesStub
	fakeReturns := fake.getCapturePagesReturns
	fake.recordInvocation("GetCapturePages", []interface{}{arg1, arg2, arg3})
	fake.getCapturePagesMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePersistor) GetCapturePagesCallCount() int {
	fake.getCapturePagesMutex.RLock()
	defer fake.getCapturePagesMutex.RUnlock()
	return len(fake.getCapturePagesArgsForCall)
}

func (fake *FakePersistor) GetCapturePagesCalls(stub func(context.Context, persistence.TransactionHandler, *model.CapturePagesFilters) (*model.PaginatedCapturePages, error)) {
	fake.getCapturePagesMutex.Lock()
	defer fake.getCapturePagesMutex.Unlock()
	fake.GetCapturePagesStub = stub
}

func (fake *FakePersistor) GetCapturePagesArgsForCall(i int) (context.Context, persistence.TransactionHandler, *model.CapturePagesFilters) {
	fake.getCapturePagesMutex.RLock()
	defer fake.getCapturePagesMutex.RUnlock()
	argsForCall := fake.getCapturePagesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakePersistor) GetCapturePagesReturns(result1 *model.PaginatedCapturePages, result2 error) {
	fake.getCapturePagesMutex.Lock()
	defer fake.getCapturePagesMutex.Unlock()
	fake.GetCapturePagesStub = nil
	fake.getCapturePagesReturns = struct {
		result1 *model.PaginatedCapturePages
		result2 error
	}{result1, result2}
}

func (fake *FakePersistor) GetCapturePagesReturnsOnCall(i int, result1 *model.PaginatedCapturePages, result2 error) {
	fake.getCapturePagesMutex.Lock()
	defer fake.getCapturePagesMutex.Unlock()
	fake.GetCapturePagesStub = nil
	if fake.getCapturePagesReturnsOnCall == nil {
		fake.getCapturePagesReturnsOnCall = make(map[int]struct {
			result1 *model.PaginatedCapturePages
			result2 error
		})
	}
	fake.getCapturePagesReturnsOnCall[i] = struct {
		result1 *model.PaginatedCapturePages
		result2 error
	}{result1, result2}
}

func (fake *FakePersistor) GetCapturePagesById(arg1 context.Context, arg2 persistence.TransactionHandler, arg3 int) (*model.CapturePages, error) {
	fake.getCapturePagesByIdMutex.Lock()
	ret, specificReturn := fake.getCapturePagesByIdReturnsOnCall[len(fake.getCapturePagesByIdArgsForCall)]
	fake.getCapturePagesByIdArgsForCall = append(fake.getCapturePagesByIdArgsForCall, struct {
		arg1 context.Context
		arg2 persistence.TransactionHandler
		arg3 int
	}{arg1, arg2, arg3})
	stub := fake.GetCapturePagesByIdStub
	fakeReturns := fake.getCapturePagesByIdReturns
	fake.recordInvocation("GetCapturePagesById", []interface{}{arg1, arg2, arg3})
	fake.getCapturePagesByIdMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePersistor) GetCapturePagesByIdCallCount() int {
	fake.getCapturePagesByIdMutex.RLock()
	defer fake.getCapturePagesByIdMutex.RUnlock()
	return len(fake.getCapturePagesByIdArgsForCall)
}

func (fake *FakePersistor) GetCapturePagesByIdCalls(stub func(context.Context, persistence.TransactionHandler, int) (*model.CapturePages, error)) {
	fake.getCapturePagesByIdMutex.Lock()
	defer fake.getCapturePagesByIdMutex.Unlock()
	fake.GetCapturePagesByIdStub = stub
}

func (fake *FakePersistor) GetCapturePagesByIdArgsForCall(i int) (context.Context, persistence.TransactionHandler, int) {
	fake.getCapturePagesByIdMutex.RLock()
	defer fake.getCapturePagesByIdMutex.RUnlock()
	argsForCall := fake.getCapturePagesByIdArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakePersistor) GetCapturePagesByIdReturns(result1 *model.CapturePages, result2 error) {
	fake.getCapturePagesByIdMutex.Lock()
	defer fake.getCapturePagesByIdMutex.Unlock()
	fake.GetCapturePagesByIdStub = nil
	fake.getCapturePagesByIdReturns = struct {
		result1 *model.CapturePages
		result2 error
	}{result1, result2}
}

func (fake *FakePersistor) GetCapturePagesByIdReturnsOnCall(i int, result1 *model.CapturePages, result2 error) {
	fake.getCapturePagesByIdMutex.Lock()
	defer fake.getCapturePagesByIdMutex.Unlock()
	fake.GetCapturePagesByIdStub = nil
	if fake.getCapturePagesByIdReturnsOnCall == nil {
		fake.getCapturePagesByIdReturnsOnCall = make(map[int]struct {
			result1 *model.CapturePages
			result2 error
		})
	}
	fake.getCapturePagesByIdReturnsOnCall[i] = struct {
		result1 *model.CapturePages
		result2 error
	}{result1, result2}
}

func (fake *FakePersistor) RestoreCapturePages(arg1 context.Context, arg2 persistence.TransactionHandler, arg3 int) error {
	fake.restoreCapturePagesMutex.Lock()
	ret, specificReturn := fake.restoreCapturePagesReturnsOnCall[len(fake.restoreCapturePagesArgsForCall)]
	fake.restoreCapturePagesArgsForCall = append(fake.restoreCapturePagesArgsForCall, struct {
		arg1 context.Context
		arg2 persistence.TransactionHandler
		arg3 int
	}{arg1, arg2, arg3})
	stub := fake.RestoreCapturePagesStub
	fakeReturns := fake.restoreCapturePagesReturns
	fake.recordInvocation("RestoreCapturePages", []interface{}{arg1, arg2, arg3})
	fake.restoreCapturePagesMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakePersistor) RestoreCapturePagesCallCount() int {
	fake.restoreCapturePagesMutex.RLock()
	defer fake.restoreCapturePagesMutex.RUnlock()
	return len(fake.restoreCapturePagesArgsForCall)
}

func (fake *FakePersistor) RestoreCapturePagesCalls(stub func(context.Context, persistence.TransactionHandler, int) error) {
	fake.restoreCapturePagesMutex.Lock()
	defer fake.restoreCapturePagesMutex.Unlock()
	fake.RestoreCapturePagesStub = stub
}

func (fake *FakePersistor) RestoreCapturePagesArgsForCall(i int) (context.Context, persistence.TransactionHandler, int) {
	fake.restoreCapturePagesMutex.RLock()
	defer fake.restoreCapturePagesMutex.RUnlock()
	argsForCall := fake.restoreCapturePagesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakePersistor) RestoreCapturePagesReturns(result1 error) {
	fake.restoreCapturePagesMutex.Lock()
	defer fake.restoreCapturePagesMutex.Unlock()
	fake.RestoreCapturePagesStub = nil
	fake.restoreCapturePagesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePersistor) RestoreCapturePagesReturnsOnCall(i int, result1 error) {
	fake.restoreCapturePagesMutex.Lock()
	defer fake.restoreCapturePagesMutex.Unlock()
	fake.RestoreCapturePagesStub = nil
	if fake.restoreCapturePagesReturnsOnCall == nil {
		fake.restoreCapturePagesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.restoreCapturePagesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakePersistor) UpdateCapturePages(arg1 context.Context, arg2 persistence.TransactionHandler, arg3 *model.UpdateCapturePages) (*model.CapturePages, error) {
	fake.updateCapturePagesMutex.Lock()
	ret, specificReturn := fake.updateCapturePagesReturnsOnCall[len(fake.updateCapturePagesArgsForCall)]
	fake.updateCapturePagesArgsForCall = append(fake.updateCapturePagesArgsForCall, struct {
		arg1 context.Context
		arg2 persistence.TransactionHandler
		arg3 *model.UpdateCapturePages
	}{arg1, arg2, arg3})
	stub := fake.UpdateCapturePagesStub
	fakeReturns := fake.updateCapturePagesReturns
	fake.recordInvocation("UpdateCapturePages", []interface{}{arg1, arg2, arg3})
	fake.updateCapturePagesMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePersistor) UpdateCapturePagesCallCount() int {
	fake.updateCapturePagesMutex.RLock()
	defer fake.updateCapturePagesMutex.RUnlock()
	return len(fake.updateCapturePagesArgsForCall)
}

func (fake *FakePersistor) UpdateCapturePagesCalls(stub func(context.Context, persistence.TransactionHandler, *model.UpdateCapturePages) (*model.CapturePages, error)) {
	fake.updateCapturePagesMutex.Lock()
	defer fake.updateCapturePagesMutex.Unlock()
	fake.UpdateCapturePagesStub = stub
}

func (fake *FakePersistor) UpdateCapturePagesArgsForCall(i int) (context.Context, persistence.TransactionHandler, *model.UpdateCapturePages) {
	fake.updateCapturePagesMutex.RLock()
	defer fake.updateCapturePagesMutex.RUnlock()
	argsForCall := fake.updateCapturePagesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakePersistor) UpdateCapturePagesReturns(result1 *model.CapturePages, result2 error) {
	fake.updateCapturePagesMutex.Lock()
	defer fake.updateCapturePagesMutex.Unlock()
	fake.UpdateCapturePagesStub = nil
	fake.updateCapturePagesReturns = struct {
		result1 *model.CapturePages
		result2 error
	}{result1, result2}
}

func (fake *FakePersistor) UpdateCapturePagesReturnsOnCall(i int, result1 *model.CapturePages, result2 error) {
	fake.updateCapturePagesMutex.Lock()
	defer fake.updateCapturePagesMutex.Unlock()
	fake.UpdateCapturePagesStub = nil
	if fake.updateCapturePagesReturnsOnCall == nil {
		fake.updateCapturePagesReturnsOnCall = make(map[int]struct {
			result1 *model.CapturePages
			result2 error
		})
	}
	fake.updateCapturePagesReturnsOnCall[i] = struct {
		result1 *model.CapturePages
		result2 error
	}{result1, result2}
}

func (fake *FakePersistor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createCapturePagesMutex.RLock()
	defer fake.createCapturePagesMutex.RUnlock()
	fake.deleteCapturePagesMutex.RLock()
	defer fake.deleteCapturePagesMutex.RUnlock()
	fake.getCapturePageTypeByIdMutex.RLock()
	defer fake.getCapturePageTypeByIdMutex.RUnlock()
	fake.getCapturePagesMutex.RLock()
	defer fake.getCapturePagesMutex.RUnlock()
	fake.getCapturePagesByIdMutex.RLock()
	defer fake.getCapturePagesByIdMutex.RUnlock()
	fake.restoreCapturePagesMutex.RLock()
	defer fake.restoreCapturePagesMutex.RUnlock()
	fake.updateCapturePagesMutex.RLock()
	defer fake.updateCapturePagesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakePersistor) recordInvocation(key string, args []interface{}) {
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
