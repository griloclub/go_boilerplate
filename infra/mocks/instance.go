// Code generated by MockGen. DO NOT EDIT.
// Source: infra/contract/instance.go
//
// Generated by this command:
//
//	mockgen -package mocks -source=infra/contract/instance.go -destination=infra/mocks/instance.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	contract "github.com/diegoclair/go_boilerplate/domain/contract"
	contract0 "github.com/diegoclair/go_boilerplate/infra/contract"
	logger "github.com/diegoclair/go_utils/logger"
	validator "github.com/diegoclair/go_utils/validator"
	gomock "go.uber.org/mock/gomock"
)

// MockInfrastructure is a mock of Infrastructure interface.
type MockInfrastructure struct {
	ctrl     *gomock.Controller
	recorder *MockInfrastructureMockRecorder
}

// MockInfrastructureMockRecorder is the mock recorder for MockInfrastructure.
type MockInfrastructureMockRecorder struct {
	mock *MockInfrastructure
}

// NewMockInfrastructure creates a new mock instance.
func NewMockInfrastructure(ctrl *gomock.Controller) *MockInfrastructure {
	mock := &MockInfrastructure{ctrl: ctrl}
	mock.recorder = &MockInfrastructureMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInfrastructure) EXPECT() *MockInfrastructureMockRecorder {
	return m.recorder
}

// AuthToken mocks base method.
func (m *MockInfrastructure) AuthToken() contract0.AuthToken {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthToken")
	ret0, _ := ret[0].(contract0.AuthToken)
	return ret0
}

// AuthToken indicates an expected call of AuthToken.
func (mr *MockInfrastructureMockRecorder) AuthToken() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthToken", reflect.TypeOf((*MockInfrastructure)(nil).AuthToken))
}

// CacheManager mocks base method.
func (m *MockInfrastructure) CacheManager() contract0.CacheManager {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CacheManager")
	ret0, _ := ret[0].(contract0.CacheManager)
	return ret0
}

// CacheManager indicates an expected call of CacheManager.
func (mr *MockInfrastructureMockRecorder) CacheManager() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CacheManager", reflect.TypeOf((*MockInfrastructure)(nil).CacheManager))
}

// Crypto mocks base method.
func (m *MockInfrastructure) Crypto() contract0.Crypto {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Crypto")
	ret0, _ := ret[0].(contract0.Crypto)
	return ret0
}

// Crypto indicates an expected call of Crypto.
func (mr *MockInfrastructureMockRecorder) Crypto() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Crypto", reflect.TypeOf((*MockInfrastructure)(nil).Crypto))
}

// DataManager mocks base method.
func (m *MockInfrastructure) DataManager() contract.DataManager {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DataManager")
	ret0, _ := ret[0].(contract.DataManager)
	return ret0
}

// DataManager indicates an expected call of DataManager.
func (mr *MockInfrastructureMockRecorder) DataManager() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DataManager", reflect.TypeOf((*MockInfrastructure)(nil).DataManager))
}

// Logger mocks base method.
func (m *MockInfrastructure) Logger() logger.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Logger")
	ret0, _ := ret[0].(logger.Logger)
	return ret0
}

// Logger indicates an expected call of Logger.
func (mr *MockInfrastructureMockRecorder) Logger() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logger", reflect.TypeOf((*MockInfrastructure)(nil).Logger))
}

// Validator mocks base method.
func (m *MockInfrastructure) Validator() validator.Validator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validator")
	ret0, _ := ret[0].(validator.Validator)
	return ret0
}

// Validator indicates an expected call of Validator.
func (mr *MockInfrastructureMockRecorder) Validator() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validator", reflect.TypeOf((*MockInfrastructure)(nil).Validator))
}
