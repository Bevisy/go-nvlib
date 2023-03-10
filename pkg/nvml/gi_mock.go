// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package nvml

import (
	"sync"
)

// Ensure, that GpuInstanceMock does implement GpuInstance.
// If this is not the case, regenerate this file with moq.
var _ GpuInstance = &GpuInstanceMock{}

// GpuInstanceMock is a mock implementation of GpuInstance.
//
//	func TestSomethingThatUsesGpuInstance(t *testing.T) {
//
//		// make and configure a mocked GpuInstance
//		mockedGpuInstance := &GpuInstanceMock{
//			CreateComputeInstanceFunc: func(Info *ComputeInstanceProfileInfo) (ComputeInstance, Return) {
//				panic("mock out the CreateComputeInstance method")
//			},
//			DestroyFunc: func() Return {
//				panic("mock out the Destroy method")
//			},
//			GetComputeInstanceByIdFunc: func(ID int) (ComputeInstance, Return) {
//				panic("mock out the GetComputeInstanceById method")
//			},
//			GetComputeInstanceProfileInfoFunc: func(Profile int, EngProfile int) (ComputeInstanceProfileInfo, Return) {
//				panic("mock out the GetComputeInstanceProfileInfo method")
//			},
//			GetComputeInstancesFunc: func(Info *ComputeInstanceProfileInfo) ([]ComputeInstance, Return) {
//				panic("mock out the GetComputeInstances method")
//			},
//			GetInfoFunc: func() (GpuInstanceInfo, Return) {
//				panic("mock out the GetInfo method")
//			},
//		}
//
//		// use mockedGpuInstance in code that requires GpuInstance
//		// and then make assertions.
//
//	}
type GpuInstanceMock struct {
	// CreateComputeInstanceFunc mocks the CreateComputeInstance method.
	CreateComputeInstanceFunc func(Info *ComputeInstanceProfileInfo) (ComputeInstance, Return)

	// DestroyFunc mocks the Destroy method.
	DestroyFunc func() Return

	// GetComputeInstanceByIdFunc mocks the GetComputeInstanceById method.
	GetComputeInstanceByIdFunc func(ID int) (ComputeInstance, Return)

	// GetComputeInstanceProfileInfoFunc mocks the GetComputeInstanceProfileInfo method.
	GetComputeInstanceProfileInfoFunc func(Profile int, EngProfile int) (ComputeInstanceProfileInfo, Return)

	// GetComputeInstancesFunc mocks the GetComputeInstances method.
	GetComputeInstancesFunc func(Info *ComputeInstanceProfileInfo) ([]ComputeInstance, Return)

	// GetInfoFunc mocks the GetInfo method.
	GetInfoFunc func() (GpuInstanceInfo, Return)

	// calls tracks calls to the methods.
	calls struct {
		// CreateComputeInstance holds details about calls to the CreateComputeInstance method.
		CreateComputeInstance []struct {
			// Info is the Info argument value.
			Info *ComputeInstanceProfileInfo
		}
		// Destroy holds details about calls to the Destroy method.
		Destroy []struct {
		}
		// GetComputeInstanceById holds details about calls to the GetComputeInstanceById method.
		GetComputeInstanceById []struct {
			// ID is the ID argument value.
			ID int
		}
		// GetComputeInstanceProfileInfo holds details about calls to the GetComputeInstanceProfileInfo method.
		GetComputeInstanceProfileInfo []struct {
			// Profile is the Profile argument value.
			Profile int
			// EngProfile is the EngProfile argument value.
			EngProfile int
		}
		// GetComputeInstances holds details about calls to the GetComputeInstances method.
		GetComputeInstances []struct {
			// Info is the Info argument value.
			Info *ComputeInstanceProfileInfo
		}
		// GetInfo holds details about calls to the GetInfo method.
		GetInfo []struct {
		}
	}
	lockCreateComputeInstance         sync.RWMutex
	lockDestroy                       sync.RWMutex
	lockGetComputeInstanceById        sync.RWMutex
	lockGetComputeInstanceProfileInfo sync.RWMutex
	lockGetComputeInstances           sync.RWMutex
	lockGetInfo                       sync.RWMutex
}

// CreateComputeInstance calls CreateComputeInstanceFunc.
func (mock *GpuInstanceMock) CreateComputeInstance(Info *ComputeInstanceProfileInfo) (ComputeInstance, Return) {
	if mock.CreateComputeInstanceFunc == nil {
		panic("GpuInstanceMock.CreateComputeInstanceFunc: method is nil but GpuInstance.CreateComputeInstance was just called")
	}
	callInfo := struct {
		Info *ComputeInstanceProfileInfo
	}{
		Info: Info,
	}
	mock.lockCreateComputeInstance.Lock()
	mock.calls.CreateComputeInstance = append(mock.calls.CreateComputeInstance, callInfo)
	mock.lockCreateComputeInstance.Unlock()
	return mock.CreateComputeInstanceFunc(Info)
}

// CreateComputeInstanceCalls gets all the calls that were made to CreateComputeInstance.
// Check the length with:
//
//	len(mockedGpuInstance.CreateComputeInstanceCalls())
func (mock *GpuInstanceMock) CreateComputeInstanceCalls() []struct {
	Info *ComputeInstanceProfileInfo
} {
	var calls []struct {
		Info *ComputeInstanceProfileInfo
	}
	mock.lockCreateComputeInstance.RLock()
	calls = mock.calls.CreateComputeInstance
	mock.lockCreateComputeInstance.RUnlock()
	return calls
}

// Destroy calls DestroyFunc.
func (mock *GpuInstanceMock) Destroy() Return {
	if mock.DestroyFunc == nil {
		panic("GpuInstanceMock.DestroyFunc: method is nil but GpuInstance.Destroy was just called")
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
//	len(mockedGpuInstance.DestroyCalls())
func (mock *GpuInstanceMock) DestroyCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockDestroy.RLock()
	calls = mock.calls.Destroy
	mock.lockDestroy.RUnlock()
	return calls
}

// GetComputeInstanceById calls GetComputeInstanceByIdFunc.
func (mock *GpuInstanceMock) GetComputeInstanceById(ID int) (ComputeInstance, Return) {
	if mock.GetComputeInstanceByIdFunc == nil {
		panic("GpuInstanceMock.GetComputeInstanceByIdFunc: method is nil but GpuInstance.GetComputeInstanceById was just called")
	}
	callInfo := struct {
		ID int
	}{
		ID: ID,
	}
	mock.lockGetComputeInstanceById.Lock()
	mock.calls.GetComputeInstanceById = append(mock.calls.GetComputeInstanceById, callInfo)
	mock.lockGetComputeInstanceById.Unlock()
	return mock.GetComputeInstanceByIdFunc(ID)
}

// GetComputeInstanceByIdCalls gets all the calls that were made to GetComputeInstanceById.
// Check the length with:
//
//	len(mockedGpuInstance.GetComputeInstanceByIdCalls())
func (mock *GpuInstanceMock) GetComputeInstanceByIdCalls() []struct {
	ID int
} {
	var calls []struct {
		ID int
	}
	mock.lockGetComputeInstanceById.RLock()
	calls = mock.calls.GetComputeInstanceById
	mock.lockGetComputeInstanceById.RUnlock()
	return calls
}

// GetComputeInstanceProfileInfo calls GetComputeInstanceProfileInfoFunc.
func (mock *GpuInstanceMock) GetComputeInstanceProfileInfo(Profile int, EngProfile int) (ComputeInstanceProfileInfo, Return) {
	if mock.GetComputeInstanceProfileInfoFunc == nil {
		panic("GpuInstanceMock.GetComputeInstanceProfileInfoFunc: method is nil but GpuInstance.GetComputeInstanceProfileInfo was just called")
	}
	callInfo := struct {
		Profile    int
		EngProfile int
	}{
		Profile:    Profile,
		EngProfile: EngProfile,
	}
	mock.lockGetComputeInstanceProfileInfo.Lock()
	mock.calls.GetComputeInstanceProfileInfo = append(mock.calls.GetComputeInstanceProfileInfo, callInfo)
	mock.lockGetComputeInstanceProfileInfo.Unlock()
	return mock.GetComputeInstanceProfileInfoFunc(Profile, EngProfile)
}

// GetComputeInstanceProfileInfoCalls gets all the calls that were made to GetComputeInstanceProfileInfo.
// Check the length with:
//
//	len(mockedGpuInstance.GetComputeInstanceProfileInfoCalls())
func (mock *GpuInstanceMock) GetComputeInstanceProfileInfoCalls() []struct {
	Profile    int
	EngProfile int
} {
	var calls []struct {
		Profile    int
		EngProfile int
	}
	mock.lockGetComputeInstanceProfileInfo.RLock()
	calls = mock.calls.GetComputeInstanceProfileInfo
	mock.lockGetComputeInstanceProfileInfo.RUnlock()
	return calls
}

// GetComputeInstances calls GetComputeInstancesFunc.
func (mock *GpuInstanceMock) GetComputeInstances(Info *ComputeInstanceProfileInfo) ([]ComputeInstance, Return) {
	if mock.GetComputeInstancesFunc == nil {
		panic("GpuInstanceMock.GetComputeInstancesFunc: method is nil but GpuInstance.GetComputeInstances was just called")
	}
	callInfo := struct {
		Info *ComputeInstanceProfileInfo
	}{
		Info: Info,
	}
	mock.lockGetComputeInstances.Lock()
	mock.calls.GetComputeInstances = append(mock.calls.GetComputeInstances, callInfo)
	mock.lockGetComputeInstances.Unlock()
	return mock.GetComputeInstancesFunc(Info)
}

// GetComputeInstancesCalls gets all the calls that were made to GetComputeInstances.
// Check the length with:
//
//	len(mockedGpuInstance.GetComputeInstancesCalls())
func (mock *GpuInstanceMock) GetComputeInstancesCalls() []struct {
	Info *ComputeInstanceProfileInfo
} {
	var calls []struct {
		Info *ComputeInstanceProfileInfo
	}
	mock.lockGetComputeInstances.RLock()
	calls = mock.calls.GetComputeInstances
	mock.lockGetComputeInstances.RUnlock()
	return calls
}

// GetInfo calls GetInfoFunc.
func (mock *GpuInstanceMock) GetInfo() (GpuInstanceInfo, Return) {
	if mock.GetInfoFunc == nil {
		panic("GpuInstanceMock.GetInfoFunc: method is nil but GpuInstance.GetInfo was just called")
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
//	len(mockedGpuInstance.GetInfoCalls())
func (mock *GpuInstanceMock) GetInfoCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetInfo.RLock()
	calls = mock.calls.GetInfo
	mock.lockGetInfo.RUnlock()
	return calls
}
