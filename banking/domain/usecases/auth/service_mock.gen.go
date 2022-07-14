// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package auth

import (
	"sync"
)

// Ensure, that ServiceMock does implement Service.
// If this is not the case, regenerate this file with moq.
var _ Service = &ServiceMock{}

// ServiceMock is a mock implementation of Service.
//
// 	func TestSomethingThatUsesService(t *testing.T) {
//
// 		// make and configure a mocked Service
// 		mockedService := &ServiceMock{
// 			NewTokenFunc: func(accountID string) (string, error) {
// 				panic("mock out the NewToken method")
// 			},
// 		}
//
// 		// use mockedService in code that requires Service
// 		// and then make assertions.
//
// 	}
type ServiceMock struct {
	// NewTokenFunc mocks the NewToken method.
	NewTokenFunc func(accountID string) (string, error)

	// calls tracks calls to the methods.
	calls struct {
		// NewToken holds details about calls to the NewToken method.
		NewToken []struct {
			// AccountID is the accountID argument value.
			AccountID string
		}
	}
	lockNewToken sync.RWMutex
}

// NewToken calls NewTokenFunc.
func (mock *ServiceMock) NewToken(accountID string) (string, error) {
	if mock.NewTokenFunc == nil {
		panic("ServiceMock.NewTokenFunc: method is nil but Service.NewToken was just called")
	}
	callInfo := struct {
		AccountID string
	}{
		AccountID: accountID,
	}
	mock.lockNewToken.Lock()
	mock.calls.NewToken = append(mock.calls.NewToken, callInfo)
	mock.lockNewToken.Unlock()
	return mock.NewTokenFunc(accountID)
}

// NewTokenCalls gets all the calls that were made to NewToken.
// Check the length with:
//     len(mockedService.NewTokenCalls())
func (mock *ServiceMock) NewTokenCalls() []struct {
	AccountID string
} {
	var calls []struct {
		AccountID string
	}
	mock.lockNewToken.RLock()
	calls = mock.calls.NewToken
	mock.lockNewToken.RUnlock()
	return calls
}