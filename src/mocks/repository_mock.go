// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package mock_repository is a generated GoMock package.
package mocks
import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIRepository is a mock of IRepository interface
type MockIRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIRepositoryMockRecorder
}

// MockIRepositoryMockRecorder is the mock recorder for MockIRepository
type MockIRepositoryMockRecorder struct {
	mock *MockIRepository
}

// NewMockIRepository creates a new mock instance
func NewMockIRepository(ctrl *gomock.Controller) *MockIRepository {
	mock := &MockIRepository{ctrl: ctrl}
	mock.recorder = &MockIRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIRepository) EXPECT() *MockIRepositoryMockRecorder {
	return m.recorder
}

// SearchWord mocks base method
func (m *MockIRepository) SearchWord(word string) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchWord", word)
	ret0, _ := ret[0].(int)
	return ret0
}

// SearchWord indicates an expected call of SearchWord
func (mr *MockIRepositoryMockRecorder) SearchWord(word interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchWord", reflect.TypeOf((*MockIRepository)(nil).SearchWord), word)
}
