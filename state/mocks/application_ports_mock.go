// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/state (interfaces: ApplicationPortRanges)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	network "github.com/juju/juju/core/network"
	state "github.com/juju/juju/state"
	gomock "go.uber.org/mock/gomock"
)

// MockApplicationPortRanges is a mock of ApplicationPortRanges interface.
type MockApplicationPortRanges struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationPortRangesMockRecorder
}

// MockApplicationPortRangesMockRecorder is the mock recorder for MockApplicationPortRanges.
type MockApplicationPortRangesMockRecorder struct {
	mock *MockApplicationPortRanges
}

// NewMockApplicationPortRanges creates a new mock instance.
func NewMockApplicationPortRanges(ctrl *gomock.Controller) *MockApplicationPortRanges {
	mock := &MockApplicationPortRanges{ctrl: ctrl}
	mock.recorder = &MockApplicationPortRangesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApplicationPortRanges) EXPECT() *MockApplicationPortRangesMockRecorder {
	return m.recorder
}

// ApplicationName mocks base method.
func (m *MockApplicationPortRanges) ApplicationName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplicationName")
	ret0, _ := ret[0].(string)
	return ret0
}

// ApplicationName indicates an expected call of ApplicationName.
func (mr *MockApplicationPortRangesMockRecorder) ApplicationName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplicationName", reflect.TypeOf((*MockApplicationPortRanges)(nil).ApplicationName))
}

// ByEndpoint mocks base method.
func (m *MockApplicationPortRanges) ByEndpoint() network.GroupedPortRanges {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ByEndpoint")
	ret0, _ := ret[0].(network.GroupedPortRanges)
	return ret0
}

// ByEndpoint indicates an expected call of ByEndpoint.
func (mr *MockApplicationPortRangesMockRecorder) ByEndpoint() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ByEndpoint", reflect.TypeOf((*MockApplicationPortRanges)(nil).ByEndpoint))
}

// ByUnit mocks base method.
func (m *MockApplicationPortRanges) ByUnit() map[string]state.UnitPortRanges {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ByUnit")
	ret0, _ := ret[0].(map[string]state.UnitPortRanges)
	return ret0
}

// ByUnit indicates an expected call of ByUnit.
func (mr *MockApplicationPortRangesMockRecorder) ByUnit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ByUnit", reflect.TypeOf((*MockApplicationPortRanges)(nil).ByUnit))
}

// Changes mocks base method.
func (m *MockApplicationPortRanges) Changes() state.ModelOperation {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Changes")
	ret0, _ := ret[0].(state.ModelOperation)
	return ret0
}

// Changes indicates an expected call of Changes.
func (mr *MockApplicationPortRangesMockRecorder) Changes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Changes", reflect.TypeOf((*MockApplicationPortRanges)(nil).Changes))
}

// ForUnit mocks base method.
func (m *MockApplicationPortRanges) ForUnit(arg0 string) state.UnitPortRanges {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForUnit", arg0)
	ret0, _ := ret[0].(state.UnitPortRanges)
	return ret0
}

// ForUnit indicates an expected call of ForUnit.
func (mr *MockApplicationPortRangesMockRecorder) ForUnit(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForUnit", reflect.TypeOf((*MockApplicationPortRanges)(nil).ForUnit), arg0)
}

// UniquePortRanges mocks base method.
func (m *MockApplicationPortRanges) UniquePortRanges() []network.PortRange {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UniquePortRanges")
	ret0, _ := ret[0].([]network.PortRange)
	return ret0
}

// UniquePortRanges indicates an expected call of UniquePortRanges.
func (mr *MockApplicationPortRangesMockRecorder) UniquePortRanges() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UniquePortRanges", reflect.TypeOf((*MockApplicationPortRanges)(nil).UniquePortRanges))
}
