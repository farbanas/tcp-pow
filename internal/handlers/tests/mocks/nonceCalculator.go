// Code generated by MockGen. DO NOT EDIT.
// Source: tcp-pow/internal/handlers (interfaces: NonceCalculator)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockNonceCalculator is a mock of NonceCalculator interface.
type MockNonceCalculator struct {
	ctrl     *gomock.Controller
	recorder *MockNonceCalculatorMockRecorder
}

// MockNonceCalculatorMockRecorder is the mock recorder for MockNonceCalculator.
type MockNonceCalculatorMockRecorder struct {
	mock *MockNonceCalculator
}

// NewMockNonceCalculator creates a new mock instance.
func NewMockNonceCalculator(ctrl *gomock.Controller) *MockNonceCalculator {
	mock := &MockNonceCalculator{ctrl: ctrl}
	mock.recorder = &MockNonceCalculatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNonceCalculator) EXPECT() *MockNonceCalculatorMockRecorder {
	return m.recorder
}

// GetNewValue mocks base method.
func (m *MockNonceCalculator) GetNewValue() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNewValue")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetNewValue indicates an expected call of GetNewValue.
func (mr *MockNonceCalculatorMockRecorder) GetNewValue() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNewValue", reflect.TypeOf((*MockNonceCalculator)(nil).GetNewValue))
}