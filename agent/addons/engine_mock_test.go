// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/cmd/jujud/agent/engine (interfaces: MetricSink)

// Package addons_test is a generated GoMock package.
package addons_test

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockMetricSink is a mock of MetricSink interface.
type MockMetricSink struct {
	ctrl     *gomock.Controller
	recorder *MockMetricSinkMockRecorder
}

// MockMetricSinkMockRecorder is the mock recorder for MockMetricSink.
type MockMetricSinkMockRecorder struct {
	mock *MockMetricSink
}

// NewMockMetricSink creates a new mock instance.
func NewMockMetricSink(ctrl *gomock.Controller) *MockMetricSink {
	mock := &MockMetricSink{ctrl: ctrl}
	mock.recorder = &MockMetricSinkMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMetricSink) EXPECT() *MockMetricSinkMockRecorder {
	return m.recorder
}

// RecordStart mocks base method.
func (m *MockMetricSink) RecordStart(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RecordStart", arg0)
}

// RecordStart indicates an expected call of RecordStart.
func (mr *MockMetricSinkMockRecorder) RecordStart(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecordStart", reflect.TypeOf((*MockMetricSink)(nil).RecordStart), arg0)
}

// Unregister mocks base method.
func (m *MockMetricSink) Unregister() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unregister")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Unregister indicates an expected call of Unregister.
func (mr *MockMetricSinkMockRecorder) Unregister() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unregister", reflect.TypeOf((*MockMetricSink)(nil).Unregister))
}
