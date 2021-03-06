// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/replicatedhq/ship/pkg/lifecycle (interfaces: Renderer)

// Package lifecycle is a generated GoMock package.
package lifecycle

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	api "github.com/replicatedhq/ship/pkg/api"
	lifecycle "github.com/replicatedhq/ship/pkg/lifecycle"
	daemontypes "github.com/replicatedhq/ship/pkg/lifecycle/daemon/daemontypes"
	planner "github.com/replicatedhq/ship/pkg/lifecycle/render/planner"
)

// MockRenderer is a mock of Renderer interface
type MockRenderer struct {
	ctrl     *gomock.Controller
	recorder *MockRendererMockRecorder
}

// MockRendererMockRecorder is the mock recorder for MockRenderer
type MockRendererMockRecorder struct {
	mock *MockRenderer
}

// NewMockRenderer creates a new mock instance
func NewMockRenderer(ctrl *gomock.Controller) *MockRenderer {
	mock := &MockRenderer{ctrl: ctrl}
	mock.recorder = &MockRendererMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRenderer) EXPECT() *MockRendererMockRecorder {
	return m.recorder
}

// Execute mocks base method
func (m *MockRenderer) Execute(arg0 context.Context, arg1 *api.Release, arg2 *api.Render) error {
	ret := m.ctrl.Call(m, "Execute", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Execute indicates an expected call of Execute
func (mr *MockRendererMockRecorder) Execute(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockRenderer)(nil).Execute), arg0, arg1, arg2)
}

// WithPlanner mocks base method
func (m *MockRenderer) WithPlanner(arg0 planner.Planner) lifecycle.Renderer {
	ret := m.ctrl.Call(m, "WithPlanner", arg0)
	ret0, _ := ret[0].(lifecycle.Renderer)
	return ret0
}

// WithPlanner indicates an expected call of WithPlanner
func (mr *MockRendererMockRecorder) WithPlanner(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithPlanner", reflect.TypeOf((*MockRenderer)(nil).WithPlanner), arg0)
}

// WithStatusReceiver mocks base method
func (m *MockRenderer) WithStatusReceiver(arg0 daemontypes.StatusReceiver) lifecycle.Renderer {
	ret := m.ctrl.Call(m, "WithStatusReceiver", arg0)
	ret0, _ := ret[0].(lifecycle.Renderer)
	return ret0
}

// WithStatusReceiver indicates an expected call of WithStatusReceiver
func (mr *MockRendererMockRecorder) WithStatusReceiver(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithStatusReceiver", reflect.TypeOf((*MockRenderer)(nil).WithStatusReceiver), arg0)
}
