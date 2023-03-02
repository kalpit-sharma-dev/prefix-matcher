// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package mock_service is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIService is a mock of IService interface
type MockIService struct {
	ctrl     *gomock.Controller
	recorder *MockIServiceMockRecorder
}

// MockIServiceMockRecorder is the mock recorder for MockIService
type MockIServiceMockRecorder struct {
	mock *MockIService
}

// NewMockIService creates a new mock instance
func NewMockIService(ctrl *gomock.Controller) *MockIService {
	mock := &MockIService{ctrl: ctrl}
	mock.recorder = &MockIServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIService) EXPECT() *MockIServiceMockRecorder {
	return m.recorder
}

// SearchWord mocks base method
func (m *MockIService) SearchWord(word string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchWord", word)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchWord indicates an expected call of SearchWord
func (mr *MockIServiceMockRecorder) SearchWord(word interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchWord", reflect.TypeOf((*MockIService)(nil).SearchWord), word)
}