// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/enbility/eebus-go/service (interfaces: ServiceProvider,ConnectionsHub)

// Package service is a generated GoMock package.
package service

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockServiceProvider is a mock of ServiceProvider interface.
type MockServiceProvider struct {
	ctrl     *gomock.Controller
	recorder *MockServiceProviderMockRecorder
}

// MockServiceProviderMockRecorder is the mock recorder for MockServiceProvider.
type MockServiceProviderMockRecorder struct {
	mock *MockServiceProvider
}

// NewMockServiceProvider creates a new mock instance.
func NewMockServiceProvider(ctrl *gomock.Controller) *MockServiceProvider {
	mock := &MockServiceProvider{ctrl: ctrl}
	mock.recorder = &MockServiceProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceProvider) EXPECT() *MockServiceProviderMockRecorder {
	return m.recorder
}

// AllowWaitingForTrust mocks base method.
func (m *MockServiceProvider) AllowWaitingForTrust(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllowWaitingForTrust", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// AllowWaitingForTrust indicates an expected call of AllowWaitingForTrust.
func (mr *MockServiceProviderMockRecorder) AllowWaitingForTrust(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllowWaitingForTrust", reflect.TypeOf((*MockServiceProvider)(nil).AllowWaitingForTrust), arg0)
}

// RemoteSKIConnected mocks base method.
func (m *MockServiceProvider) RemoteSKIConnected(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoteSKIConnected", arg0)
}

// RemoteSKIConnected indicates an expected call of RemoteSKIConnected.
func (mr *MockServiceProviderMockRecorder) RemoteSKIConnected(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoteSKIConnected", reflect.TypeOf((*MockServiceProvider)(nil).RemoteSKIConnected), arg0)
}

// RemoteSKIDisconnected mocks base method.
func (m *MockServiceProvider) RemoteSKIDisconnected(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoteSKIDisconnected", arg0)
}

// RemoteSKIDisconnected indicates an expected call of RemoteSKIDisconnected.
func (mr *MockServiceProviderMockRecorder) RemoteSKIDisconnected(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoteSKIDisconnected", reflect.TypeOf((*MockServiceProvider)(nil).RemoteSKIDisconnected), arg0)
}

// ServicePairingDetailUpdate mocks base method.
func (m *MockServiceProvider) ServicePairingDetailUpdate(arg0 string, arg1 *ConnectionStateDetail) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ServicePairingDetailUpdate", arg0, arg1)
}

// ServicePairingDetailUpdate indicates an expected call of ServicePairingDetailUpdate.
func (mr *MockServiceProviderMockRecorder) ServicePairingDetailUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServicePairingDetailUpdate", reflect.TypeOf((*MockServiceProvider)(nil).ServicePairingDetailUpdate), arg0, arg1)
}

// ServiceShipIDUpdate mocks base method.
func (m *MockServiceProvider) ServiceShipIDUpdate(arg0, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ServiceShipIDUpdate", arg0, arg1)
}

// ServiceShipIDUpdate indicates an expected call of ServiceShipIDUpdate.
func (mr *MockServiceProviderMockRecorder) ServiceShipIDUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServiceShipIDUpdate", reflect.TypeOf((*MockServiceProvider)(nil).ServiceShipIDUpdate), arg0, arg1)
}

// VisibleMDNSRecordsUpdated mocks base method.
func (m *MockServiceProvider) VisibleMDNSRecordsUpdated(arg0 []*MdnsEntry) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "VisibleMDNSRecordsUpdated", arg0)
}

// VisibleMDNSRecordsUpdated indicates an expected call of VisibleMDNSRecordsUpdated.
func (mr *MockServiceProviderMockRecorder) VisibleMDNSRecordsUpdated(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VisibleMDNSRecordsUpdated", reflect.TypeOf((*MockServiceProvider)(nil).VisibleMDNSRecordsUpdated), arg0)
}

// MockConnectionsHub is a mock of ConnectionsHub interface.
type MockConnectionsHub struct {
	ctrl     *gomock.Controller
	recorder *MockConnectionsHubMockRecorder
}

// MockConnectionsHubMockRecorder is the mock recorder for MockConnectionsHub.
type MockConnectionsHubMockRecorder struct {
	mock *MockConnectionsHub
}

// NewMockConnectionsHub creates a new mock instance.
func NewMockConnectionsHub(ctrl *gomock.Controller) *MockConnectionsHub {
	mock := &MockConnectionsHub{ctrl: ctrl}
	mock.recorder = &MockConnectionsHubMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConnectionsHub) EXPECT() *MockConnectionsHubMockRecorder {
	return m.recorder
}

// CancelPairingWithSKI mocks base method.
func (m *MockConnectionsHub) CancelPairingWithSKI(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CancelPairingWithSKI", arg0)
}

// CancelPairingWithSKI indicates an expected call of CancelPairingWithSKI.
func (mr *MockConnectionsHubMockRecorder) CancelPairingWithSKI(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelPairingWithSKI", reflect.TypeOf((*MockConnectionsHub)(nil).CancelPairingWithSKI), arg0)
}

// DisconnectSKI mocks base method.
func (m *MockConnectionsHub) DisconnectSKI(arg0, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DisconnectSKI", arg0, arg1)
}

// DisconnectSKI indicates an expected call of DisconnectSKI.
func (mr *MockConnectionsHubMockRecorder) DisconnectSKI(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisconnectSKI", reflect.TypeOf((*MockConnectionsHub)(nil).DisconnectSKI), arg0, arg1)
}

// InitiatePairingWithSKI mocks base method.
func (m *MockConnectionsHub) InitiatePairingWithSKI(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "InitiatePairingWithSKI", arg0)
}

// InitiatePairingWithSKI indicates an expected call of InitiatePairingWithSKI.
func (mr *MockConnectionsHubMockRecorder) InitiatePairingWithSKI(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitiatePairingWithSKI", reflect.TypeOf((*MockConnectionsHub)(nil).InitiatePairingWithSKI), arg0)
}

// PairingDetailForSki mocks base method.
func (m *MockConnectionsHub) PairingDetailForSki(arg0 string) *ConnectionStateDetail {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PairingDetailForSki", arg0)
	ret0, _ := ret[0].(*ConnectionStateDetail)
	return ret0
}

// PairingDetailForSki indicates an expected call of PairingDetailForSki.
func (mr *MockConnectionsHubMockRecorder) PairingDetailForSki(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PairingDetailForSki", reflect.TypeOf((*MockConnectionsHub)(nil).PairingDetailForSki), arg0)
}

// RegisterRemoteSKI mocks base method.
func (m *MockConnectionsHub) RegisterRemoteSKI(arg0 string, arg1 bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterRemoteSKI", arg0, arg1)
}

// RegisterRemoteSKI indicates an expected call of RegisterRemoteSKI.
func (mr *MockConnectionsHubMockRecorder) RegisterRemoteSKI(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterRemoteSKI", reflect.TypeOf((*MockConnectionsHub)(nil).RegisterRemoteSKI), arg0, arg1)
}

// ServiceForSKI mocks base method.
func (m *MockConnectionsHub) ServiceForSKI(arg0 string) *ServiceDetails {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServiceForSKI", arg0)
	ret0, _ := ret[0].(*ServiceDetails)
	return ret0
}

// ServiceForSKI indicates an expected call of ServiceForSKI.
func (mr *MockConnectionsHubMockRecorder) ServiceForSKI(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServiceForSKI", reflect.TypeOf((*MockConnectionsHub)(nil).ServiceForSKI), arg0)
}

// Shutdown mocks base method.
func (m *MockConnectionsHub) Shutdown() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Shutdown")
}

// Shutdown indicates an expected call of Shutdown.
func (mr *MockConnectionsHubMockRecorder) Shutdown() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*MockConnectionsHub)(nil).Shutdown))
}

// Start mocks base method.
func (m *MockConnectionsHub) Start() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start")
}

// Start indicates an expected call of Start.
func (mr *MockConnectionsHubMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockConnectionsHub)(nil).Start))
}

// StartBrowseMdnsSearch mocks base method.
func (m *MockConnectionsHub) StartBrowseMdnsSearch() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StartBrowseMdnsSearch")
}

// StartBrowseMdnsSearch indicates an expected call of StartBrowseMdnsSearch.
func (mr *MockConnectionsHubMockRecorder) StartBrowseMdnsSearch() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartBrowseMdnsSearch", reflect.TypeOf((*MockConnectionsHub)(nil).StartBrowseMdnsSearch))
}

// StopBrowseMdnsSearch mocks base method.
func (m *MockConnectionsHub) StopBrowseMdnsSearch() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StopBrowseMdnsSearch")
}

// StopBrowseMdnsSearch indicates an expected call of StopBrowseMdnsSearch.
func (mr *MockConnectionsHubMockRecorder) StopBrowseMdnsSearch() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopBrowseMdnsSearch", reflect.TypeOf((*MockConnectionsHub)(nil).StopBrowseMdnsSearch))
}