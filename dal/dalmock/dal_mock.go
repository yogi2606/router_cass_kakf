// Code generated by MockGen. DO NOT EDIT.
// Source: Idal.go

package dalmock

import (
	gomock "github.com/golang/mock/gomock"
	model "practise/router_cass_kakf/model"
	reflect "reflect"
)

// MockCRMInterface is a mock of CRMInterface interface
type MockCRMInterface struct {
	ctrl     *gomock.Controller
	recorder *MockCRMInterfaceMockRecorder
}

// MockCRMInterfaceMockRecorder is the mock recorder for MockCRMInterface
type MockCRMInterfaceMockRecorder struct {
	mock *MockCRMInterface
}

// NewMockCRMInterface creates a new mock instance
func NewMockCRMInterface(ctrl *gomock.Controller) *MockCRMInterface {
	mock := &MockCRMInterface{ctrl: ctrl}
	mock.recorder = &MockCRMInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockCRMInterface) EXPECT() *MockCRMInterfaceMockRecorder {
	return _m.recorder
}

// InsertCRM mocks base method
func (_m *MockCRMInterface) InsertCRM(customer *model.Customer) error {
	ret := _m.ctrl.Call(_m, "InsertCRM", customer)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertCRM indicates an expected call of InsertCRM
func (_mr *MockCRMInterfaceMockRecorder) InsertCRM(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "InsertCRM", reflect.TypeOf((*MockCRMInterface)(nil).InsertCRM), arg0)
}

// GetAll mocks base method
func (_m *MockCRMInterface) GetAll() []model.Customer {
	ret := _m.ctrl.Call(_m, "GetAll")
	ret0, _ := ret[0].([]model.Customer)
	return ret0
}

// GetAll indicates an expected call of GetAll
func (_mr *MockCRMInterfaceMockRecorder) GetAll() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetAll", reflect.TypeOf((*MockCRMInterface)(nil).GetAll))
}

// Get mocks base method
func (_m *MockCRMInterface) Get(customerID string) []model.Customer {
	ret := _m.ctrl.Call(_m, "Get", customerID)
	ret0, _ := ret[0].([]model.Customer)
	return ret0
}

// Get indicates an expected call of Get
func (_mr *MockCRMInterfaceMockRecorder) Get(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Get", reflect.TypeOf((*MockCRMInterface)(nil).Get), arg0)
}

// Delete mocks base method
func (_m *MockCRMInterface) Delete(customerID string) error {
	ret := _m.ctrl.Call(_m, "Delete", customerID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (_mr *MockCRMInterfaceMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Delete", reflect.TypeOf((*MockCRMInterface)(nil).Delete), arg0)
}
