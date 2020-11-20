// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/cmd/juju/charmhub (interfaces: DownloadCommandAPI,InfoCommandAPI,FindCommandAPI,ModelConfigClient,ModelConfigGetter,CharmHubClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	charmhub "github.com/juju/juju/api/charmhub"
	transport "github.com/juju/juju/charmhub/transport"
	url "net/url"
	reflect "reflect"
)

// MockDownloadCommandAPI is a mock of DownloadCommandAPI interface
type MockDownloadCommandAPI struct {
	ctrl     *gomock.Controller
	recorder *MockDownloadCommandAPIMockRecorder
}

// MockDownloadCommandAPIMockRecorder is the mock recorder for MockDownloadCommandAPI
type MockDownloadCommandAPIMockRecorder struct {
	mock *MockDownloadCommandAPI
}

// NewMockDownloadCommandAPI creates a new mock instance
func NewMockDownloadCommandAPI(ctrl *gomock.Controller) *MockDownloadCommandAPI {
	mock := &MockDownloadCommandAPI{ctrl: ctrl}
	mock.recorder = &MockDownloadCommandAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDownloadCommandAPI) EXPECT() *MockDownloadCommandAPIMockRecorder {
	return m.recorder
}

// Download mocks base method
func (m *MockDownloadCommandAPI) Download(arg0 context.Context, arg1 *url.URL, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Download", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Download indicates an expected call of Download
func (mr *MockDownloadCommandAPIMockRecorder) Download(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Download", reflect.TypeOf((*MockDownloadCommandAPI)(nil).Download), arg0, arg1, arg2)
}

// Info mocks base method
func (m *MockDownloadCommandAPI) Info(arg0 context.Context, arg1 string) (transport.InfoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Info", arg0, arg1)
	ret0, _ := ret[0].(transport.InfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Info indicates an expected call of Info
func (mr *MockDownloadCommandAPIMockRecorder) Info(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockDownloadCommandAPI)(nil).Info), arg0, arg1)
}

// MockInfoCommandAPI is a mock of InfoCommandAPI interface
type MockInfoCommandAPI struct {
	ctrl     *gomock.Controller
	recorder *MockInfoCommandAPIMockRecorder
}

// MockInfoCommandAPIMockRecorder is the mock recorder for MockInfoCommandAPI
type MockInfoCommandAPIMockRecorder struct {
	mock *MockInfoCommandAPI
}

// NewMockInfoCommandAPI creates a new mock instance
func NewMockInfoCommandAPI(ctrl *gomock.Controller) *MockInfoCommandAPI {
	mock := &MockInfoCommandAPI{ctrl: ctrl}
	mock.recorder = &MockInfoCommandAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInfoCommandAPI) EXPECT() *MockInfoCommandAPIMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockInfoCommandAPI) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockInfoCommandAPIMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockInfoCommandAPI)(nil).Close))
}

// Info mocks base method
func (m *MockInfoCommandAPI) Info(arg0 string) (charmhub.InfoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Info", arg0)
	ret0, _ := ret[0].(charmhub.InfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Info indicates an expected call of Info
func (mr *MockInfoCommandAPIMockRecorder) Info(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockInfoCommandAPI)(nil).Info), arg0)
}

// MockFindCommandAPI is a mock of FindCommandAPI interface
type MockFindCommandAPI struct {
	ctrl     *gomock.Controller
	recorder *MockFindCommandAPIMockRecorder
}

// MockFindCommandAPIMockRecorder is the mock recorder for MockFindCommandAPI
type MockFindCommandAPIMockRecorder struct {
	mock *MockFindCommandAPI
}

// NewMockFindCommandAPI creates a new mock instance
func NewMockFindCommandAPI(ctrl *gomock.Controller) *MockFindCommandAPI {
	mock := &MockFindCommandAPI{ctrl: ctrl}
	mock.recorder = &MockFindCommandAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFindCommandAPI) EXPECT() *MockFindCommandAPIMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockFindCommandAPI) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockFindCommandAPIMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockFindCommandAPI)(nil).Close))
}

// Find mocks base method
func (m *MockFindCommandAPI) Find(arg0 string) ([]charmhub.FindResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0)
	ret0, _ := ret[0].([]charmhub.FindResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockFindCommandAPIMockRecorder) Find(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockFindCommandAPI)(nil).Find), arg0)
}

// MockModelConfigClient is a mock of ModelConfigClient interface
type MockModelConfigClient struct {
	ctrl     *gomock.Controller
	recorder *MockModelConfigClientMockRecorder
}

// MockModelConfigClientMockRecorder is the mock recorder for MockModelConfigClient
type MockModelConfigClientMockRecorder struct {
	mock *MockModelConfigClient
}

// NewMockModelConfigClient creates a new mock instance
func NewMockModelConfigClient(ctrl *gomock.Controller) *MockModelConfigClient {
	mock := &MockModelConfigClient{ctrl: ctrl}
	mock.recorder = &MockModelConfigClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockModelConfigClient) EXPECT() *MockModelConfigClientMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockModelConfigClient) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockModelConfigClientMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockModelConfigClient)(nil).Close))
}

// ModelGet mocks base method
func (m *MockModelConfigClient) ModelGet() (map[string]interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelGet")
	ret0, _ := ret[0].(map[string]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModelGet indicates an expected call of ModelGet
func (mr *MockModelConfigClientMockRecorder) ModelGet() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelGet", reflect.TypeOf((*MockModelConfigClient)(nil).ModelGet))
}

// MockModelConfigGetter is a mock of ModelConfigGetter interface
type MockModelConfigGetter struct {
	ctrl     *gomock.Controller
	recorder *MockModelConfigGetterMockRecorder
}

// MockModelConfigGetterMockRecorder is the mock recorder for MockModelConfigGetter
type MockModelConfigGetterMockRecorder struct {
	mock *MockModelConfigGetter
}

// NewMockModelConfigGetter creates a new mock instance
func NewMockModelConfigGetter(ctrl *gomock.Controller) *MockModelConfigGetter {
	mock := &MockModelConfigGetter{ctrl: ctrl}
	mock.recorder = &MockModelConfigGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockModelConfigGetter) EXPECT() *MockModelConfigGetterMockRecorder {
	return m.recorder
}

// ModelGet mocks base method
func (m *MockModelConfigGetter) ModelGet() (map[string]interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelGet")
	ret0, _ := ret[0].(map[string]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModelGet indicates an expected call of ModelGet
func (mr *MockModelConfigGetterMockRecorder) ModelGet() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelGet", reflect.TypeOf((*MockModelConfigGetter)(nil).ModelGet))
}

// MockCharmHubClient is a mock of CharmHubClient interface
type MockCharmHubClient struct {
	ctrl     *gomock.Controller
	recorder *MockCharmHubClientMockRecorder
}

// MockCharmHubClientMockRecorder is the mock recorder for MockCharmHubClient
type MockCharmHubClientMockRecorder struct {
	mock *MockCharmHubClient
}

// NewMockCharmHubClient creates a new mock instance
func NewMockCharmHubClient(ctrl *gomock.Controller) *MockCharmHubClient {
	mock := &MockCharmHubClient{ctrl: ctrl}
	mock.recorder = &MockCharmHubClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCharmHubClient) EXPECT() *MockCharmHubClientMockRecorder {
	return m.recorder
}

// Download mocks base method
func (m *MockCharmHubClient) Download(arg0 context.Context, arg1 *url.URL, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Download", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Download indicates an expected call of Download
func (mr *MockCharmHubClientMockRecorder) Download(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Download", reflect.TypeOf((*MockCharmHubClient)(nil).Download), arg0, arg1, arg2)
}

// Info mocks base method
func (m *MockCharmHubClient) Info(arg0 context.Context, arg1 string) (transport.InfoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Info", arg0, arg1)
	ret0, _ := ret[0].(transport.InfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Info indicates an expected call of Info
func (mr *MockCharmHubClientMockRecorder) Info(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockCharmHubClient)(nil).Info), arg0, arg1)
}
