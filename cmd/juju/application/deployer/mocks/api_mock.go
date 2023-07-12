// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/api (interfaces: AllWatch)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	params "github.com/juju/juju/rpc/params"
	gomock "go.uber.org/mock/gomock"
)

// MockAllWatch is a mock of AllWatch interface.
type MockAllWatch struct {
	ctrl     *gomock.Controller
	recorder *MockAllWatchMockRecorder
}

// MockAllWatchMockRecorder is the mock recorder for MockAllWatch.
type MockAllWatchMockRecorder struct {
	mock *MockAllWatch
}

// NewMockAllWatch creates a new mock instance.
func NewMockAllWatch(ctrl *gomock.Controller) *MockAllWatch {
	mock := &MockAllWatch{ctrl: ctrl}
	mock.recorder = &MockAllWatchMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAllWatch) EXPECT() *MockAllWatchMockRecorder {
	return m.recorder
}

// Next mocks base method.
func (m *MockAllWatch) Next() ([]params.Delta, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next")
	ret0, _ := ret[0].([]params.Delta)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Next indicates an expected call of Next.
func (mr *MockAllWatchMockRecorder) Next() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockAllWatch)(nil).Next))
}

// Stop mocks base method.
func (m *MockAllWatch) Stop() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop.
func (mr *MockAllWatchMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockAllWatch)(nil).Stop))
}
