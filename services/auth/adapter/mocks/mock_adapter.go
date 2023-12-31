// Code generated by MockGen. DO NOT EDIT.
// Source: services/auth/adapter/adapter.go
//
// Generated by this command:
//
//	mockgen -source=services/auth/adapter/adapter.go -destination=services/auth/adapter/mocks/mock_adapter.go
//
// Package mock_adapter is a generated GoMock package.
package mock_adapter

import (
	context "context"
	reflect "reflect"
	gen "warehouseai/auth/adapter/grpc/gen"
	errors "warehouseai/auth/errors"
	model "warehouseai/auth/model"

	gomock "go.uber.org/mock/gomock"
)

// MockUserGrpcInterface is a mock of UserGrpcInterface interface.
type MockUserGrpcInterface struct {
	ctrl     *gomock.Controller
	recorder *MockUserGrpcInterfaceMockRecorder
}

// MockUserGrpcInterfaceMockRecorder is the mock recorder for MockUserGrpcInterface.
type MockUserGrpcInterfaceMockRecorder struct {
	mock *MockUserGrpcInterface
}

// NewMockUserGrpcInterface creates a new mock instance.
func NewMockUserGrpcInterface(ctrl *gomock.Controller) *MockUserGrpcInterface {
	mock := &MockUserGrpcInterface{ctrl: ctrl}
	mock.recorder = &MockUserGrpcInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserGrpcInterface) EXPECT() *MockUserGrpcInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUserGrpcInterface) Create(ctx context.Context, userInfo *gen.CreateUserMsg) (string, *errors.ErrorResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, userInfo)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(*errors.ErrorResponse)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockUserGrpcInterfaceMockRecorder) Create(ctx, userInfo any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserGrpcInterface)(nil).Create), ctx, userInfo)
}

// GetByEmail mocks base method.
func (m *MockUserGrpcInterface) GetByEmail(ctx context.Context, email string) (*gen.User, *errors.ErrorResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByEmail", ctx, email)
	ret0, _ := ret[0].(*gen.User)
	ret1, _ := ret[1].(*errors.ErrorResponse)
	return ret0, ret1
}

// GetByEmail indicates an expected call of GetByEmail.
func (mr *MockUserGrpcInterfaceMockRecorder) GetByEmail(ctx, email any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByEmail", reflect.TypeOf((*MockUserGrpcInterface)(nil).GetByEmail), ctx, email)
}

// GetById mocks base method.
func (m *MockUserGrpcInterface) GetById(ctx context.Context, userId string) (*gen.User, *errors.ErrorResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", ctx, userId)
	ret0, _ := ret[0].(*gen.User)
	ret1, _ := ret[1].(*errors.ErrorResponse)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockUserGrpcInterfaceMockRecorder) GetById(ctx, userId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockUserGrpcInterface)(nil).GetById), ctx, userId)
}

// ResetPassword mocks base method.
func (m *MockUserGrpcInterface) ResetPassword(ctx context.Context, resetPasswordRequest *gen.ResetPasswordRequest) (string, *errors.ErrorResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResetPassword", ctx, resetPasswordRequest)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(*errors.ErrorResponse)
	return ret0, ret1
}

// ResetPassword indicates an expected call of ResetPassword.
func (mr *MockUserGrpcInterfaceMockRecorder) ResetPassword(ctx, resetPasswordRequest any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetPassword", reflect.TypeOf((*MockUserGrpcInterface)(nil).ResetPassword), ctx, resetPasswordRequest)
}

// UpdateVerificationStatus mocks base method.
func (m *MockUserGrpcInterface) UpdateVerificationStatus(ctx context.Context, userId string) (bool, *errors.ErrorResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateVerificationStatus", ctx, userId)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(*errors.ErrorResponse)
	return ret0, ret1
}

// UpdateVerificationStatus indicates an expected call of UpdateVerificationStatus.
func (mr *MockUserGrpcInterfaceMockRecorder) UpdateVerificationStatus(ctx, userId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateVerificationStatus", reflect.TypeOf((*MockUserGrpcInterface)(nil).UpdateVerificationStatus), ctx, userId)
}

// MockBrokerInterface is a mock of BrokerInterface interface.
type MockBrokerInterface struct {
	ctrl     *gomock.Controller
	recorder *MockBrokerInterfaceMockRecorder
}

// MockBrokerInterfaceMockRecorder is the mock recorder for MockBrokerInterface.
type MockBrokerInterfaceMockRecorder struct {
	mock *MockBrokerInterface
}

// NewMockBrokerInterface creates a new mock instance.
func NewMockBrokerInterface(ctrl *gomock.Controller) *MockBrokerInterface {
	mock := &MockBrokerInterface{ctrl: ctrl}
	mock.recorder = &MockBrokerInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBrokerInterface) EXPECT() *MockBrokerInterfaceMockRecorder {
	return m.recorder
}

// SendEmail mocks base method.
func (m *MockBrokerInterface) SendEmail(email model.Email) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendEmail", email)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendEmail indicates an expected call of SendEmail.
func (mr *MockBrokerInterfaceMockRecorder) SendEmail(email any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendEmail", reflect.TypeOf((*MockBrokerInterface)(nil).SendEmail), email)
}

// SendUserReject mocks base method.
func (m *MockBrokerInterface) SendUserReject(userId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendUserReject", userId)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendUserReject indicates an expected call of SendUserReject.
func (mr *MockBrokerInterfaceMockRecorder) SendUserReject(userId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendUserReject", reflect.TypeOf((*MockBrokerInterface)(nil).SendUserReject), userId)
}
