// Code generated by MockGen. DO NOT EDIT.
// Source: tcp-pow/internal/handlers (interfaces: SolutionCalculator)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockSolutionCalculator is a mock of SolutionCalculator interface.
type MockSolutionCalculator struct {
	ctrl     *gomock.Controller
	recorder *MockSolutionCalculatorMockRecorder
}

// MockSolutionCalculatorMockRecorder is the mock recorder for MockSolutionCalculator.
type MockSolutionCalculatorMockRecorder struct {
	mock *MockSolutionCalculator
}

// NewMockSolutionCalculator creates a new mock instance.
func NewMockSolutionCalculator(ctrl *gomock.Controller) *MockSolutionCalculator {
	mock := &MockSolutionCalculator{ctrl: ctrl}
	mock.recorder = &MockSolutionCalculatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSolutionCalculator) EXPECT() *MockSolutionCalculatorMockRecorder {
	return m.recorder
}

// Calculate mocks base method.
func (m *MockSolutionCalculator) Calculate(arg0 string, arg1 int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Calculate", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Calculate indicates an expected call of Calculate.
func (mr *MockSolutionCalculatorMockRecorder) Calculate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Calculate", reflect.TypeOf((*MockSolutionCalculator)(nil).Calculate), arg0, arg1)
}
