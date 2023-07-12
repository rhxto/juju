// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/cmd/juju/application/deployer (interfaces: ModelCommand,ConsumeDetails,ModelConfigGetter)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	httpbakery "github.com/go-macaroon-bakery/macaroon-bakery/v3/httpbakery"
	modelcmd "github.com/juju/juju/cmd/modelcmd"
	model "github.com/juju/juju/core/model"
	jujuclient "github.com/juju/juju/jujuclient"
	params "github.com/juju/juju/rpc/params"
	gomock "go.uber.org/mock/gomock"
)

// MockModelCommand is a mock of ModelCommand interface.
type MockModelCommand struct {
	ctrl     *gomock.Controller
	recorder *MockModelCommandMockRecorder
}

// MockModelCommandMockRecorder is the mock recorder for MockModelCommand.
type MockModelCommandMockRecorder struct {
	mock *MockModelCommand
}

// NewMockModelCommand creates a new mock instance.
func NewMockModelCommand(ctrl *gomock.Controller) *MockModelCommand {
	mock := &MockModelCommand{ctrl: ctrl}
	mock.recorder = &MockModelCommandMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockModelCommand) EXPECT() *MockModelCommandMockRecorder {
	return m.recorder
}

// BakeryClient mocks base method.
func (m *MockModelCommand) BakeryClient() (*httpbakery.Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BakeryClient")
	ret0, _ := ret[0].(*httpbakery.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BakeryClient indicates an expected call of BakeryClient.
func (mr *MockModelCommandMockRecorder) BakeryClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BakeryClient", reflect.TypeOf((*MockModelCommand)(nil).BakeryClient))
}

// ControllerName mocks base method.
func (m *MockModelCommand) ControllerName() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerName")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ControllerName indicates an expected call of ControllerName.
func (mr *MockModelCommandMockRecorder) ControllerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerName", reflect.TypeOf((*MockModelCommand)(nil).ControllerName))
}

// CurrentAccountDetails mocks base method.
func (m *MockModelCommand) CurrentAccountDetails() (*jujuclient.AccountDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CurrentAccountDetails")
	ret0, _ := ret[0].(*jujuclient.AccountDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CurrentAccountDetails indicates an expected call of CurrentAccountDetails.
func (mr *MockModelCommandMockRecorder) CurrentAccountDetails() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurrentAccountDetails", reflect.TypeOf((*MockModelCommand)(nil).CurrentAccountDetails))
}

// Filesystem mocks base method.
func (m *MockModelCommand) Filesystem() modelcmd.Filesystem {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Filesystem")
	ret0, _ := ret[0].(modelcmd.Filesystem)
	return ret0
}

// Filesystem indicates an expected call of Filesystem.
func (mr *MockModelCommandMockRecorder) Filesystem() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Filesystem", reflect.TypeOf((*MockModelCommand)(nil).Filesystem))
}

// ModelDetails mocks base method.
func (m *MockModelCommand) ModelDetails() (string, *jujuclient.ModelDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelDetails")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(*jujuclient.ModelDetails)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ModelDetails indicates an expected call of ModelDetails.
func (mr *MockModelCommandMockRecorder) ModelDetails() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelDetails", reflect.TypeOf((*MockModelCommand)(nil).ModelDetails))
}

// ModelType mocks base method.
func (m *MockModelCommand) ModelType() (model.ModelType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelType")
	ret0, _ := ret[0].(model.ModelType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModelType indicates an expected call of ModelType.
func (mr *MockModelCommandMockRecorder) ModelType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelType", reflect.TypeOf((*MockModelCommand)(nil).ModelType))
}

// MockConsumeDetails is a mock of ConsumeDetails interface.
type MockConsumeDetails struct {
	ctrl     *gomock.Controller
	recorder *MockConsumeDetailsMockRecorder
}

// MockConsumeDetailsMockRecorder is the mock recorder for MockConsumeDetails.
type MockConsumeDetailsMockRecorder struct {
	mock *MockConsumeDetails
}

// NewMockConsumeDetails creates a new mock instance.
func NewMockConsumeDetails(ctrl *gomock.Controller) *MockConsumeDetails {
	mock := &MockConsumeDetails{ctrl: ctrl}
	mock.recorder = &MockConsumeDetailsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConsumeDetails) EXPECT() *MockConsumeDetailsMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockConsumeDetails) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockConsumeDetailsMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockConsumeDetails)(nil).Close))
}

// GetConsumeDetails mocks base method.
func (m *MockConsumeDetails) GetConsumeDetails(arg0 string) (params.ConsumeOfferDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConsumeDetails", arg0)
	ret0, _ := ret[0].(params.ConsumeOfferDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConsumeDetails indicates an expected call of GetConsumeDetails.
func (mr *MockConsumeDetailsMockRecorder) GetConsumeDetails(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConsumeDetails", reflect.TypeOf((*MockConsumeDetails)(nil).GetConsumeDetails), arg0)
}

// MockModelConfigGetter is a mock of ModelConfigGetter interface.
type MockModelConfigGetter struct {
	ctrl     *gomock.Controller
	recorder *MockModelConfigGetterMockRecorder
}

// MockModelConfigGetterMockRecorder is the mock recorder for MockModelConfigGetter.
type MockModelConfigGetterMockRecorder struct {
	mock *MockModelConfigGetter
}

// NewMockModelConfigGetter creates a new mock instance.
func NewMockModelConfigGetter(ctrl *gomock.Controller) *MockModelConfigGetter {
	mock := &MockModelConfigGetter{ctrl: ctrl}
	mock.recorder = &MockModelConfigGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockModelConfigGetter) EXPECT() *MockModelConfigGetterMockRecorder {
	return m.recorder
}

// ModelGet mocks base method.
func (m *MockModelConfigGetter) ModelGet() (map[string]interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelGet")
	ret0, _ := ret[0].(map[string]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModelGet indicates an expected call of ModelGet.
func (mr *MockModelConfigGetterMockRecorder) ModelGet() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelGet", reflect.TypeOf((*MockModelConfigGetter)(nil).ModelGet))
}
