// Code generated by MockGen. DO NOT EDIT.
// Source: dep.go

// Package dep_test is a generated GoMock package.
package dep_test

import (
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRunner is a mock of Runner interface
type MockRunner struct {
	ctrl     *gomock.Controller
	recorder *MockRunnerMockRecorder
}

// MockRunnerMockRecorder is the mock recorder for MockRunner
type MockRunnerMockRecorder struct {
	mock *MockRunner
}

// NewMockRunner creates a new mock instance
func NewMockRunner(ctrl *gomock.Controller) *MockRunner {
	mock := &MockRunner{ctrl: ctrl}
	mock.recorder = &MockRunnerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRunner) EXPECT() *MockRunnerMockRecorder {
	return m.recorder
}

// Run mocks base method
func (m *MockRunner) Run(dir, bin string, args ...string) (string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{dir, bin}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Run", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Run indicates an expected call of Run
func (mr *MockRunnerMockRecorder) Run(dir, bin interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{dir, bin}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockRunner)(nil).Run), varargs...)
}

// RunSilent mocks base method
func (m *MockRunner) RunSilent(dir, bin string, args ...string) (string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{dir, bin}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunSilent", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RunSilent indicates an expected call of RunSilent
func (mr *MockRunnerMockRecorder) RunSilent(dir, bin interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{dir, bin}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunSilent", reflect.TypeOf((*MockRunner)(nil).RunSilent), varargs...)
}

// CustomRun mocks base method
func (m *MockRunner) CustomRun(dir string, env []string, out, err io.Writer, bin string, args ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{dir, env, out, err, bin}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CustomRun", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CustomRun indicates an expected call of CustomRun
func (mr *MockRunnerMockRecorder) CustomRun(dir, env, out, err, bin interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{dir, env, out, err, bin}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CustomRun", reflect.TypeOf((*MockRunner)(nil).CustomRun), varargs...)
}
