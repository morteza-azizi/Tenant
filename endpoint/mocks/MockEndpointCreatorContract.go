// Code generated by MockGen. DO NOT EDIT.
// Source: ../contracts/EndpointCreatorContract.go

// Package mock_contracts is a generated GoMock package.
package mock_contracts

import (
	endpoint "github.com/go-kit/kit/endpoint"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockEndpointCreatorContract is a mock of EndpointCreatorContract interface
type MockEndpointCreatorContract struct {
	ctrl     *gomock.Controller
	recorder *MockEndpointCreatorContractMockRecorder
}

// MockEndpointCreatorContractMockRecorder is the mock recorder for MockEndpointCreatorContract
type MockEndpointCreatorContractMockRecorder struct {
	mock *MockEndpointCreatorContract
}

// NewMockEndpointCreatorContract creates a new mock instance
func NewMockEndpointCreatorContract(ctrl *gomock.Controller) *MockEndpointCreatorContract {
	mock := &MockEndpointCreatorContract{ctrl: ctrl}
	mock.recorder = &MockEndpointCreatorContractMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEndpointCreatorContract) EXPECT() *MockEndpointCreatorContractMockRecorder {
	return m.recorder
}

// CreateTenantEndpoint mocks base method
func (m *MockEndpointCreatorContract) CreateTenantEndpoint() endpoint.Endpoint {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTenantEndpoint")
	ret0, _ := ret[0].(endpoint.Endpoint)
	return ret0
}

// CreateTenantEndpoint indicates an expected call of CreateTenantEndpoint
func (mr *MockEndpointCreatorContractMockRecorder) CreateTenantEndpoint() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTenantEndpoint", reflect.TypeOf((*MockEndpointCreatorContract)(nil).CreateTenantEndpoint))
}

// ReadTenantEndpoint mocks base method
func (m *MockEndpointCreatorContract) ReadTenantEndpoint() endpoint.Endpoint {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadTenantEndpoint")
	ret0, _ := ret[0].(endpoint.Endpoint)
	return ret0
}

// ReadTenantEndpoint indicates an expected call of ReadTenantEndpoint
func (mr *MockEndpointCreatorContractMockRecorder) ReadTenantEndpoint() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadTenantEndpoint", reflect.TypeOf((*MockEndpointCreatorContract)(nil).ReadTenantEndpoint))
}

// UpdateTenantEndpoint mocks base method
func (m *MockEndpointCreatorContract) UpdateTenantEndpoint() endpoint.Endpoint {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTenantEndpoint")
	ret0, _ := ret[0].(endpoint.Endpoint)
	return ret0
}

// UpdateTenantEndpoint indicates an expected call of UpdateTenantEndpoint
func (mr *MockEndpointCreatorContractMockRecorder) UpdateTenantEndpoint() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTenantEndpoint", reflect.TypeOf((*MockEndpointCreatorContract)(nil).UpdateTenantEndpoint))
}

// DeleteTenantEndpoint mocks base method
func (m *MockEndpointCreatorContract) DeleteTenantEndpoint() endpoint.Endpoint {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTenantEndpoint")
	ret0, _ := ret[0].(endpoint.Endpoint)
	return ret0
}

// DeleteTenantEndpoint indicates an expected call of DeleteTenantEndpoint
func (mr *MockEndpointCreatorContractMockRecorder) DeleteTenantEndpoint() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTenantEndpoint", reflect.TypeOf((*MockEndpointCreatorContract)(nil).DeleteTenantEndpoint))
}
