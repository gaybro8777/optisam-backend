// Copyright (C) 2019 Orange
// 
// This software is distributed under the terms and conditions of the 'Apache License 2.0'
// license which can be found in the file 'License.txt' in this package distribution 
// or at 'http://www.apache.org/licenses/LICENSE-2.0'. 

// Code generated by MockGen. DO NOT EDIT.
// Source: optisam-backend/license-service/pkg/api/v1 (interfaces: LicenseServiceClient)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	v1 "optisam-backend/license-service/pkg/api/v1"
	reflect "reflect"
)

// MockLicenseServiceClient is a mock of LicenseServiceClient interface
type MockLicenseServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockLicenseServiceClientMockRecorder
}

// MockLicenseServiceClientMockRecorder is the mock recorder for MockLicenseServiceClient
type MockLicenseServiceClientMockRecorder struct {
	mock *MockLicenseServiceClient
}

// NewMockLicenseServiceClient creates a new mock instance
func NewMockLicenseServiceClient(ctrl *gomock.Controller) *MockLicenseServiceClient {
	mock := &MockLicenseServiceClient{ctrl: ctrl}
	mock.recorder = &MockLicenseServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLicenseServiceClient) EXPECT() *MockLicenseServiceClientMockRecorder {
	return m.recorder
}

// CreateProductAggregation mocks base method
func (m *MockLicenseServiceClient) CreateProductAggregation(arg0 context.Context, arg1 *v1.ProductAggregation, arg2 ...grpc.CallOption) (*v1.ProductAggregation, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateProductAggregation", varargs...)
	ret0, _ := ret[0].(*v1.ProductAggregation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProductAggregation indicates an expected call of CreateProductAggregation
func (mr *MockLicenseServiceClientMockRecorder) CreateProductAggregation(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProductAggregation", reflect.TypeOf((*MockLicenseServiceClient)(nil).CreateProductAggregation), varargs...)
}

// DeleteProductAggregation mocks base method
func (m *MockLicenseServiceClient) DeleteProductAggregation(arg0 context.Context, arg1 *v1.DeleteProductAggregationRequest, arg2 ...grpc.CallOption) (*v1.ListProductAggregationResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteProductAggregation", varargs...)
	ret0, _ := ret[0].(*v1.ListProductAggregationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteProductAggregation indicates an expected call of DeleteProductAggregation
func (mr *MockLicenseServiceClientMockRecorder) DeleteProductAggregation(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProductAggregation", reflect.TypeOf((*MockLicenseServiceClient)(nil).DeleteProductAggregation), varargs...)
}

// LicensesForEquipAndMetric mocks base method
func (m *MockLicenseServiceClient) LicensesForEquipAndMetric(arg0 context.Context, arg1 *v1.LicensesForEquipAndMetricRequest, arg2 ...grpc.CallOption) (*v1.LicensesForEquipAndMetricResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "LicensesForEquipAndMetric", varargs...)
	ret0, _ := ret[0].(*v1.LicensesForEquipAndMetricResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LicensesForEquipAndMetric indicates an expected call of LicensesForEquipAndMetric
func (mr *MockLicenseServiceClientMockRecorder) LicensesForEquipAndMetric(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LicensesForEquipAndMetric", reflect.TypeOf((*MockLicenseServiceClient)(nil).LicensesForEquipAndMetric), varargs...)
}

// ListAcqRightsForProduct mocks base method
func (m *MockLicenseServiceClient) ListAcqRightsForProduct(arg0 context.Context, arg1 *v1.ListAcquiredRightsForProductRequest, arg2 ...grpc.CallOption) (*v1.ListAcquiredRightsForProductResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAcqRightsForProduct", varargs...)
	ret0, _ := ret[0].(*v1.ListAcquiredRightsForProductResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAcqRightsForProduct indicates an expected call of ListAcqRightsForProduct
func (mr *MockLicenseServiceClientMockRecorder) ListAcqRightsForProduct(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAcqRightsForProduct", reflect.TypeOf((*MockLicenseServiceClient)(nil).ListAcqRightsForProduct), varargs...)
}

// ListAcqRightsForProductAggregation mocks base method
func (m *MockLicenseServiceClient) ListAcqRightsForProductAggregation(arg0 context.Context, arg1 *v1.ListAcqRightsForProductAggregationRequest, arg2 ...grpc.CallOption) (*v1.ListAcqRightsForProductAggregationResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAcqRightsForProductAggregation", varargs...)
	ret0, _ := ret[0].(*v1.ListAcqRightsForProductAggregationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAcqRightsForProductAggregation indicates an expected call of ListAcqRightsForProductAggregation
func (mr *MockLicenseServiceClientMockRecorder) ListAcqRightsForProductAggregation(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAcqRightsForProductAggregation", reflect.TypeOf((*MockLicenseServiceClient)(nil).ListAcqRightsForProductAggregation), varargs...)
}

// MetricesForEqType mocks base method
func (m *MockLicenseServiceClient) MetricesForEqType(arg0 context.Context, arg1 *v1.MetricesForEqTypeRequest, arg2 ...grpc.CallOption) (*v1.ListMetricResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MetricesForEqType", varargs...)
	ret0, _ := ret[0].(*v1.ListMetricResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MetricesForEqType indicates an expected call of MetricesForEqType
func (mr *MockLicenseServiceClientMockRecorder) MetricesForEqType(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MetricesForEqType", reflect.TypeOf((*MockLicenseServiceClient)(nil).MetricesForEqType), varargs...)
}

// ProductLicensesForMetric mocks base method
func (m *MockLicenseServiceClient) ProductLicensesForMetric(arg0 context.Context, arg1 *v1.ProductLicensesForMetricRequest, arg2 ...grpc.CallOption) (*v1.ProductLicensesForMetricResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ProductLicensesForMetric", varargs...)
	ret0, _ := ret[0].(*v1.ProductLicensesForMetricResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProductLicensesForMetric indicates an expected call of ProductLicensesForMetric
func (mr *MockLicenseServiceClientMockRecorder) ProductLicensesForMetric(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProductLicensesForMetric", reflect.TypeOf((*MockLicenseServiceClient)(nil).ProductLicensesForMetric), varargs...)
}

// UpdateProductAggregation mocks base method
func (m *MockLicenseServiceClient) UpdateProductAggregation(arg0 context.Context, arg1 *v1.UpdateProductAggregationRequest, arg2 ...grpc.CallOption) (*v1.ProductAggregation, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateProductAggregation", varargs...)
	ret0, _ := ret[0].(*v1.ProductAggregation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProductAggregation indicates an expected call of UpdateProductAggregation
func (mr *MockLicenseServiceClientMockRecorder) UpdateProductAggregation(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProductAggregation", reflect.TypeOf((*MockLicenseServiceClient)(nil).UpdateProductAggregation), varargs...)
}
