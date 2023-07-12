// Code generated by MockGen. DO NOT EDIT.
// Source: k8s.io/client-go/kubernetes/typed/admissionregistration/v1beta1 (interfaces: AdmissionregistrationV1beta1Interface,MutatingWebhookConfigurationInterface,ValidatingWebhookConfigurationInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	v1beta1 "k8s.io/api/admissionregistration/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	v1beta10 "k8s.io/client-go/applyconfigurations/admissionregistration/v1beta1"
	v1beta11 "k8s.io/client-go/kubernetes/typed/admissionregistration/v1beta1"
	rest "k8s.io/client-go/rest"
)

// MockAdmissionregistrationV1beta1Interface is a mock of AdmissionregistrationV1beta1Interface interface.
type MockAdmissionregistrationV1beta1Interface struct {
	ctrl     *gomock.Controller
	recorder *MockAdmissionregistrationV1beta1InterfaceMockRecorder
}

// MockAdmissionregistrationV1beta1InterfaceMockRecorder is the mock recorder for MockAdmissionregistrationV1beta1Interface.
type MockAdmissionregistrationV1beta1InterfaceMockRecorder struct {
	mock *MockAdmissionregistrationV1beta1Interface
}

// NewMockAdmissionregistrationV1beta1Interface creates a new mock instance.
func NewMockAdmissionregistrationV1beta1Interface(ctrl *gomock.Controller) *MockAdmissionregistrationV1beta1Interface {
	mock := &MockAdmissionregistrationV1beta1Interface{ctrl: ctrl}
	mock.recorder = &MockAdmissionregistrationV1beta1InterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAdmissionregistrationV1beta1Interface) EXPECT() *MockAdmissionregistrationV1beta1InterfaceMockRecorder {
	return m.recorder
}

// MutatingWebhookConfigurations mocks base method.
func (m *MockAdmissionregistrationV1beta1Interface) MutatingWebhookConfigurations() v1beta11.MutatingWebhookConfigurationInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MutatingWebhookConfigurations")
	ret0, _ := ret[0].(v1beta11.MutatingWebhookConfigurationInterface)
	return ret0
}

// MutatingWebhookConfigurations indicates an expected call of MutatingWebhookConfigurations.
func (mr *MockAdmissionregistrationV1beta1InterfaceMockRecorder) MutatingWebhookConfigurations() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MutatingWebhookConfigurations", reflect.TypeOf((*MockAdmissionregistrationV1beta1Interface)(nil).MutatingWebhookConfigurations))
}

// RESTClient mocks base method.
func (m *MockAdmissionregistrationV1beta1Interface) RESTClient() rest.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RESTClient")
	ret0, _ := ret[0].(rest.Interface)
	return ret0
}

// RESTClient indicates an expected call of RESTClient.
func (mr *MockAdmissionregistrationV1beta1InterfaceMockRecorder) RESTClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RESTClient", reflect.TypeOf((*MockAdmissionregistrationV1beta1Interface)(nil).RESTClient))
}

// ValidatingWebhookConfigurations mocks base method.
func (m *MockAdmissionregistrationV1beta1Interface) ValidatingWebhookConfigurations() v1beta11.ValidatingWebhookConfigurationInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidatingWebhookConfigurations")
	ret0, _ := ret[0].(v1beta11.ValidatingWebhookConfigurationInterface)
	return ret0
}

// ValidatingWebhookConfigurations indicates an expected call of ValidatingWebhookConfigurations.
func (mr *MockAdmissionregistrationV1beta1InterfaceMockRecorder) ValidatingWebhookConfigurations() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidatingWebhookConfigurations", reflect.TypeOf((*MockAdmissionregistrationV1beta1Interface)(nil).ValidatingWebhookConfigurations))
}

// MockMutatingWebhookConfigurationV1Beta1Interface is a mock of MutatingWebhookConfigurationInterface interface.
type MockMutatingWebhookConfigurationV1Beta1Interface struct {
	ctrl     *gomock.Controller
	recorder *MockMutatingWebhookConfigurationV1Beta1InterfaceMockRecorder
}

// MockMutatingWebhookConfigurationV1Beta1InterfaceMockRecorder is the mock recorder for MockMutatingWebhookConfigurationV1Beta1Interface.
type MockMutatingWebhookConfigurationV1Beta1InterfaceMockRecorder struct {
	mock *MockMutatingWebhookConfigurationV1Beta1Interface
}

// NewMockMutatingWebhookConfigurationV1Beta1Interface creates a new mock instance.
func NewMockMutatingWebhookConfigurationV1Beta1Interface(ctrl *gomock.Controller) *MockMutatingWebhookConfigurationV1Beta1Interface {
	mock := &MockMutatingWebhookConfigurationV1Beta1Interface{ctrl: ctrl}
	mock.recorder = &MockMutatingWebhookConfigurationV1Beta1InterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMutatingWebhookConfigurationV1Beta1Interface) EXPECT() *MockMutatingWebhookConfigurationV1Beta1InterfaceMockRecorder {
	return m.recorder
}

// Apply mocks base method.
func (m *MockMutatingWebhookConfigurationV1Beta1Interface) Apply(arg0 context.Context, arg1 *v1beta10.MutatingWebhookConfigurationApplyConfiguration, arg2 v1.ApplyOptions) (*v1beta1.MutatingWebhookConfiguration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Apply", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1beta1.MutatingWebhookConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Apply indicates an expected call of Apply.
func (mr *MockMutatingWebhookConfigurationV1Beta1InterfaceMockRecorder) Apply(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Apply", reflect.TypeOf((*MockMutatingWebhookConfigurationV1Beta1Interface)(nil).Apply), arg0, arg1, arg2)
}

// Create mocks base method.
func (m *MockMutatingWebhookConfigurationV1Beta1Interface) Create(arg0 context.Context, arg1 *v1beta1.MutatingWebhookConfiguration, arg2 v1.CreateOptions) (*v1beta1.MutatingWebhookConfiguration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1beta1.MutatingWebhookConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockMutatingWebhookConfigurationV1Beta1InterfaceMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockMutatingWebhookConfigurationV1Beta1Interface)(nil).Create), arg0, arg1, arg2)
}

// Delete mocks base method.
func (m *MockMutatingWebhookConfigurationV1Beta1Interface) Delete(arg0 context.Context, arg1 string, arg2 v1.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockMutatingWebhookConfigurationV1Beta1InterfaceMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockMutatingWebhookConfigurationV1Beta1Interface)(nil).Delete), arg0, arg1, arg2)
}

// DeleteCollection mocks base method.
func (m *MockMutatingWebhookConfigurationV1Beta1Interface) DeleteCollection(arg0 context.Context, arg1 v1.DeleteOptions, arg2 v1.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCollection", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCollection indicates an expected call of DeleteCollection.
func (mr *MockMutatingWebhookConfigurationV1Beta1InterfaceMockRecorder) DeleteCollection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCollection", reflect.TypeOf((*MockMutatingWebhookConfigurationV1Beta1Interface)(nil).DeleteCollection), arg0, arg1, arg2)
}

// Get mocks base method.
func (m *MockMutatingWebhookConfigurationV1Beta1Interface) Get(arg0 context.Context, arg1 string, arg2 v1.GetOptions) (*v1beta1.MutatingWebhookConfiguration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1beta1.MutatingWebhookConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockMutatingWebhookConfigurationV1Beta1InterfaceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockMutatingWebhookConfigurationV1Beta1Interface)(nil).Get), arg0, arg1, arg2)
}

// List mocks base method.
func (m *MockMutatingWebhookConfigurationV1Beta1Interface) List(arg0 context.Context, arg1 v1.ListOptions) (*v1beta1.MutatingWebhookConfigurationList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*v1beta1.MutatingWebhookConfigurationList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockMutatingWebhookConfigurationV1Beta1InterfaceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockMutatingWebhookConfigurationV1Beta1Interface)(nil).List), arg0, arg1)
}

// Patch mocks base method.
func (m *MockMutatingWebhookConfigurationV1Beta1Interface) Patch(arg0 context.Context, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 v1.PatchOptions, arg5 ...string) (*v1beta1.MutatingWebhookConfiguration, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Patch", varargs...)
	ret0, _ := ret[0].(*v1beta1.MutatingWebhookConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Patch indicates an expected call of Patch.
func (mr *MockMutatingWebhookConfigurationV1Beta1InterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockMutatingWebhookConfigurationV1Beta1Interface)(nil).Patch), varargs...)
}

// Update mocks base method.
func (m *MockMutatingWebhookConfigurationV1Beta1Interface) Update(arg0 context.Context, arg1 *v1beta1.MutatingWebhookConfiguration, arg2 v1.UpdateOptions) (*v1beta1.MutatingWebhookConfiguration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1beta1.MutatingWebhookConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockMutatingWebhookConfigurationV1Beta1InterfaceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockMutatingWebhookConfigurationV1Beta1Interface)(nil).Update), arg0, arg1, arg2)
}

// Watch mocks base method.
func (m *MockMutatingWebhookConfigurationV1Beta1Interface) Watch(arg0 context.Context, arg1 v1.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch.
func (mr *MockMutatingWebhookConfigurationV1Beta1InterfaceMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockMutatingWebhookConfigurationV1Beta1Interface)(nil).Watch), arg0, arg1)
}

// MockValidatingWebhookConfigurationV1Beta1Interface is a mock of ValidatingWebhookConfigurationInterface interface.
type MockValidatingWebhookConfigurationV1Beta1Interface struct {
	ctrl     *gomock.Controller
	recorder *MockValidatingWebhookConfigurationV1Beta1InterfaceMockRecorder
}

// MockValidatingWebhookConfigurationV1Beta1InterfaceMockRecorder is the mock recorder for MockValidatingWebhookConfigurationV1Beta1Interface.
type MockValidatingWebhookConfigurationV1Beta1InterfaceMockRecorder struct {
	mock *MockValidatingWebhookConfigurationV1Beta1Interface
}

// NewMockValidatingWebhookConfigurationV1Beta1Interface creates a new mock instance.
func NewMockValidatingWebhookConfigurationV1Beta1Interface(ctrl *gomock.Controller) *MockValidatingWebhookConfigurationV1Beta1Interface {
	mock := &MockValidatingWebhookConfigurationV1Beta1Interface{ctrl: ctrl}
	mock.recorder = &MockValidatingWebhookConfigurationV1Beta1InterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockValidatingWebhookConfigurationV1Beta1Interface) EXPECT() *MockValidatingWebhookConfigurationV1Beta1InterfaceMockRecorder {
	return m.recorder
}

// Apply mocks base method.
func (m *MockValidatingWebhookConfigurationV1Beta1Interface) Apply(arg0 context.Context, arg1 *v1beta10.ValidatingWebhookConfigurationApplyConfiguration, arg2 v1.ApplyOptions) (*v1beta1.ValidatingWebhookConfiguration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Apply", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1beta1.ValidatingWebhookConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Apply indicates an expected call of Apply.
func (mr *MockValidatingWebhookConfigurationV1Beta1InterfaceMockRecorder) Apply(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Apply", reflect.TypeOf((*MockValidatingWebhookConfigurationV1Beta1Interface)(nil).Apply), arg0, arg1, arg2)
}

// Create mocks base method.
func (m *MockValidatingWebhookConfigurationV1Beta1Interface) Create(arg0 context.Context, arg1 *v1beta1.ValidatingWebhookConfiguration, arg2 v1.CreateOptions) (*v1beta1.ValidatingWebhookConfiguration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1beta1.ValidatingWebhookConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockValidatingWebhookConfigurationV1Beta1InterfaceMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockValidatingWebhookConfigurationV1Beta1Interface)(nil).Create), arg0, arg1, arg2)
}

// Delete mocks base method.
func (m *MockValidatingWebhookConfigurationV1Beta1Interface) Delete(arg0 context.Context, arg1 string, arg2 v1.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockValidatingWebhookConfigurationV1Beta1InterfaceMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockValidatingWebhookConfigurationV1Beta1Interface)(nil).Delete), arg0, arg1, arg2)
}

// DeleteCollection mocks base method.
func (m *MockValidatingWebhookConfigurationV1Beta1Interface) DeleteCollection(arg0 context.Context, arg1 v1.DeleteOptions, arg2 v1.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCollection", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCollection indicates an expected call of DeleteCollection.
func (mr *MockValidatingWebhookConfigurationV1Beta1InterfaceMockRecorder) DeleteCollection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCollection", reflect.TypeOf((*MockValidatingWebhookConfigurationV1Beta1Interface)(nil).DeleteCollection), arg0, arg1, arg2)
}

// Get mocks base method.
func (m *MockValidatingWebhookConfigurationV1Beta1Interface) Get(arg0 context.Context, arg1 string, arg2 v1.GetOptions) (*v1beta1.ValidatingWebhookConfiguration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1beta1.ValidatingWebhookConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockValidatingWebhookConfigurationV1Beta1InterfaceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockValidatingWebhookConfigurationV1Beta1Interface)(nil).Get), arg0, arg1, arg2)
}

// List mocks base method.
func (m *MockValidatingWebhookConfigurationV1Beta1Interface) List(arg0 context.Context, arg1 v1.ListOptions) (*v1beta1.ValidatingWebhookConfigurationList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*v1beta1.ValidatingWebhookConfigurationList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockValidatingWebhookConfigurationV1Beta1InterfaceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockValidatingWebhookConfigurationV1Beta1Interface)(nil).List), arg0, arg1)
}

// Patch mocks base method.
func (m *MockValidatingWebhookConfigurationV1Beta1Interface) Patch(arg0 context.Context, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 v1.PatchOptions, arg5 ...string) (*v1beta1.ValidatingWebhookConfiguration, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Patch", varargs...)
	ret0, _ := ret[0].(*v1beta1.ValidatingWebhookConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Patch indicates an expected call of Patch.
func (mr *MockValidatingWebhookConfigurationV1Beta1InterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockValidatingWebhookConfigurationV1Beta1Interface)(nil).Patch), varargs...)
}

// Update mocks base method.
func (m *MockValidatingWebhookConfigurationV1Beta1Interface) Update(arg0 context.Context, arg1 *v1beta1.ValidatingWebhookConfiguration, arg2 v1.UpdateOptions) (*v1beta1.ValidatingWebhookConfiguration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1beta1.ValidatingWebhookConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockValidatingWebhookConfigurationV1Beta1InterfaceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockValidatingWebhookConfigurationV1Beta1Interface)(nil).Update), arg0, arg1, arg2)
}

// Watch mocks base method.
func (m *MockValidatingWebhookConfigurationV1Beta1Interface) Watch(arg0 context.Context, arg1 v1.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch.
func (mr *MockValidatingWebhookConfigurationV1Beta1InterfaceMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockValidatingWebhookConfigurationV1Beta1Interface)(nil).Watch), arg0, arg1)
}
