// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package nvml

import (
	"sync"
)

// Ensure, that ComputeInstanceMock does implement ComputeInstance.
// If this is not the case, regenerate this file with moq.
var _ ComputeInstance = &ComputeInstanceMock{}

// ComputeInstanceMock is a mock implementation of ComputeInstance.
//
//	func TestSomethingThatUsesComputeInstance(t *testing.T) {
//
//		// make and configure a mocked ComputeInstance
//		mockedComputeInstance := &ComputeInstanceMock{
//			DestroyFunc: func() Return {
//				panic("mock out the Destroy method")
//			},
//			GetInfoFunc: func() (ComputeInstanceInfo, Return) {
//				panic("mock out the GetInfo method")
//			},
//		}
//
//		// use mockedComputeInstance in code that requires ComputeInstance
//		// and then make assertions.
//
//	}
type ComputeInstanceMock struct {
	// DestroyFunc mocks the Destroy method.
	DestroyFunc func() Return

	// GetInfoFunc mocks the GetInfo method.
	GetInfoFunc func() (ComputeInstanceInfo, Return)

	// calls tracks calls to the methods.
	calls struct {
		// Destroy holds details about calls to the Destroy method.
		Destroy []struct {
		}
		// GetInfo holds details about calls to the GetInfo method.
		GetInfo []struct {
		}
	}
	lockDestroy sync.RWMutex
	lockGetInfo sync.RWMutex
}

// Destroy calls DestroyFunc.
func (mock *ComputeInstanceMock) Destroy() Return {
	if mock.DestroyFunc == nil {
		panic("ComputeInstanceMock.DestroyFunc: method is nil but ComputeInstance.Destroy was just called")
	}
	callInfo := struct {
	}{}
	mock.lockDestroy.Lock()
	mock.calls.Destroy = append(mock.calls.Destroy, callInfo)
	mock.lockDestroy.Unlock()
	return mock.DestroyFunc()
}

// DestroyCalls gets all the calls that were made to Destroy.
// Check the length with:
//
//	len(mockedComputeInstance.DestroyCalls())
func (mock *ComputeInstanceMock) DestroyCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockDestroy.RLock()
	calls = mock.calls.Destroy
	mock.lockDestroy.RUnlock()
	return calls
}

// GetInfo calls GetInfoFunc.
func (mock *ComputeInstanceMock) GetInfo() (ComputeInstanceInfo, Return) {
	if mock.GetInfoFunc == nil {
		panic("ComputeInstanceMock.GetInfoFunc: method is nil but ComputeInstance.GetInfo was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetInfo.Lock()
	mock.calls.GetInfo = append(mock.calls.GetInfo, callInfo)
	mock.lockGetInfo.Unlock()
	return mock.GetInfoFunc()
}

// GetInfoCalls gets all the calls that were made to GetInfo.
// Check the length with:
//
//	len(mockedComputeInstance.GetInfoCalls())
func (mock *ComputeInstanceMock) GetInfoCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetInfo.RLock()
	calls = mock.calls.GetInfo
	mock.lockGetInfo.RUnlock()
	return calls
}
