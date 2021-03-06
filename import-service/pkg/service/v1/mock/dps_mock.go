// Copyright (C) 2019 Orange
// 
// This software is distributed under the terms and conditions of the 'Apache License 2.0'
// license which can be found in the file 'License.txt' in this package distribution 
// or at 'http://www.apache.org/licenses/LICENSE-2.0'. 

// Code generated by MockGen. DO NOT EDIT.
// Source: optisam-backend/dps-service/pkg/api/v1 (interfaces: DpsServiceClient)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	v1 "optisam-backend/dps-service/pkg/api/v1"
	reflect "reflect"
)

// MockDpsServiceClient is a mock of DpsServiceClient interface
type MockDpsServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockDpsServiceClientMockRecorder
}

// MockDpsServiceClientMockRecorder is the mock recorder for MockDpsServiceClient
type MockDpsServiceClientMockRecorder struct {
	mock *MockDpsServiceClient
}

// NewMockDpsServiceClient creates a new mock instance
func NewMockDpsServiceClient(ctrl *gomock.Controller) *MockDpsServiceClient {
	mock := &MockDpsServiceClient{ctrl: ctrl}
	mock.recorder = &MockDpsServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDpsServiceClient) EXPECT() *MockDpsServiceClientMockRecorder {
	return m.recorder
}

// ListUpload mocks base method
func (m *MockDpsServiceClient) ListUpload(arg0 context.Context, arg1 *v1.ListUploadRequest, arg2 ...grpc.CallOption) (*v1.ListUploadResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListUpload", varargs...)
	ret0, _ := ret[0].(*v1.ListUploadResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUpload indicates an expected call of ListUpload
func (mr *MockDpsServiceClientMockRecorder) ListUpload(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUpload", reflect.TypeOf((*MockDpsServiceClient)(nil).ListUpload), varargs...)
}

// NotifyUpload mocks base method
func (m *MockDpsServiceClient) NotifyUpload(arg0 context.Context, arg1 *v1.NotifyUploadRequest, arg2 ...grpc.CallOption) (*v1.NotifyUploadResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "NotifyUpload", varargs...)
	ret0, _ := ret[0].(*v1.NotifyUploadResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NotifyUpload indicates an expected call of NotifyUpload
func (mr *MockDpsServiceClientMockRecorder) NotifyUpload(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyUpload", reflect.TypeOf((*MockDpsServiceClient)(nil).NotifyUpload), varargs...)
}
