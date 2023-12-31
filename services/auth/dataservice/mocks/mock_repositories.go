// Code generated by MockGen. DO NOT EDIT.
// Source: services/auth/dataservice/dataservice.go
//
// Generated by this command:
//
//	mockgen -source=services/auth/dataservice/dataservice.go -destination=services/auth/dataservice/mocks/mock_repositories.go
//
// Package mock_dataservice is a generated GoMock package.
package mock_dataservice

import (
	context "context"
	multipart "mime/multipart"
	reflect "reflect"
	errors "warehouseai/auth/errors"
	model "warehouseai/auth/model"

	gomock "go.uber.org/mock/gomock"
)

// MockResetTokenInterface is a mock of ResetTokenInterface interface.
type MockResetTokenInterface struct {
	ctrl     *gomock.Controller
	recorder *MockResetTokenInterfaceMockRecorder
}

// MockResetTokenInterfaceMockRecorder is the mock recorder for MockResetTokenInterface.
type MockResetTokenInterfaceMockRecorder struct {
	mock *MockResetTokenInterface
}

// NewMockResetTokenInterface creates a new mock instance.
func NewMockResetTokenInterface(ctrl *gomock.Controller) *MockResetTokenInterface {
	mock := &MockResetTokenInterface{ctrl: ctrl}
	mock.recorder = &MockResetTokenInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResetTokenInterface) EXPECT() *MockResetTokenInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockResetTokenInterface) Create(newResetToken *model.ResetToken) *errors.DBError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", newResetToken)
	ret0, _ := ret[0].(*errors.DBError)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockResetTokenInterfaceMockRecorder) Create(newResetToken any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockResetTokenInterface)(nil).Create), newResetToken)
}

// Delete mocks base method.
func (m *MockResetTokenInterface) Delete(condition map[string]any) *errors.DBError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", condition)
	ret0, _ := ret[0].(*errors.DBError)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockResetTokenInterfaceMockRecorder) Delete(condition any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockResetTokenInterface)(nil).Delete), condition)
}

// Get mocks base method.
func (m *MockResetTokenInterface) Get(condition map[string]any) (*model.ResetToken, *errors.DBError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", condition)
	ret0, _ := ret[0].(*model.ResetToken)
	ret1, _ := ret[1].(*errors.DBError)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockResetTokenInterfaceMockRecorder) Get(condition any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockResetTokenInterface)(nil).Get), condition)
}

// MockVerificationTokenInterface is a mock of VerificationTokenInterface interface.
type MockVerificationTokenInterface struct {
	ctrl     *gomock.Controller
	recorder *MockVerificationTokenInterfaceMockRecorder
}

// MockVerificationTokenInterfaceMockRecorder is the mock recorder for MockVerificationTokenInterface.
type MockVerificationTokenInterfaceMockRecorder struct {
	mock *MockVerificationTokenInterface
}

// NewMockVerificationTokenInterface creates a new mock instance.
func NewMockVerificationTokenInterface(ctrl *gomock.Controller) *MockVerificationTokenInterface {
	mock := &MockVerificationTokenInterface{ctrl: ctrl}
	mock.recorder = &MockVerificationTokenInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVerificationTokenInterface) EXPECT() *MockVerificationTokenInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockVerificationTokenInterface) Create(newVerificationToken *model.VerificationToken) *errors.DBError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", newVerificationToken)
	ret0, _ := ret[0].(*errors.DBError)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockVerificationTokenInterfaceMockRecorder) Create(newVerificationToken any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockVerificationTokenInterface)(nil).Create), newVerificationToken)
}

// Delete mocks base method.
func (m *MockVerificationTokenInterface) Delete(condition map[string]any) *errors.DBError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", condition)
	ret0, _ := ret[0].(*errors.DBError)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockVerificationTokenInterfaceMockRecorder) Delete(condition any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockVerificationTokenInterface)(nil).Delete), condition)
}

// Get mocks base method.
func (m *MockVerificationTokenInterface) Get(condition map[string]any) (*model.VerificationToken, *errors.DBError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", condition)
	ret0, _ := ret[0].(*model.VerificationToken)
	ret1, _ := ret[1].(*errors.DBError)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockVerificationTokenInterfaceMockRecorder) Get(condition any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockVerificationTokenInterface)(nil).Get), condition)
}

// MockSessionInterface is a mock of SessionInterface interface.
type MockSessionInterface struct {
	ctrl     *gomock.Controller
	recorder *MockSessionInterfaceMockRecorder
}

// MockSessionInterfaceMockRecorder is the mock recorder for MockSessionInterface.
type MockSessionInterfaceMockRecorder struct {
	mock *MockSessionInterface
}

// NewMockSessionInterface creates a new mock instance.
func NewMockSessionInterface(ctrl *gomock.Controller) *MockSessionInterface {
	mock := &MockSessionInterface{ctrl: ctrl}
	mock.recorder = &MockSessionInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSessionInterface) EXPECT() *MockSessionInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockSessionInterface) Create(ctx context.Context, userId string) (*model.Session, *errors.DBError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, userId)
	ret0, _ := ret[0].(*model.Session)
	ret1, _ := ret[1].(*errors.DBError)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockSessionInterfaceMockRecorder) Create(ctx, userId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockSessionInterface)(nil).Create), ctx, userId)
}

// Delete mocks base method.
func (m *MockSessionInterface) Delete(ctx context.Context, sessionId string) *errors.DBError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, sessionId)
	ret0, _ := ret[0].(*errors.DBError)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockSessionInterfaceMockRecorder) Delete(ctx, sessionId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockSessionInterface)(nil).Delete), ctx, sessionId)
}

// Get mocks base method.
func (m *MockSessionInterface) Get(ctx context.Context, sessionId string) (*model.Session, *errors.DBError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, sessionId)
	ret0, _ := ret[0].(*model.Session)
	ret1, _ := ret[1].(*errors.DBError)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockSessionInterfaceMockRecorder) Get(ctx, sessionId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockSessionInterface)(nil).Get), ctx, sessionId)
}

// Update mocks base method.
func (m *MockSessionInterface) Update(ctx context.Context, sessionId string) (*string, *model.Session, *errors.DBError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, sessionId)
	ret0, _ := ret[0].(*string)
	ret1, _ := ret[1].(*model.Session)
	ret2, _ := ret[2].(*errors.DBError)
	return ret0, ret1, ret2
}

// Update indicates an expected call of Update.
func (mr *MockSessionInterfaceMockRecorder) Update(ctx, sessionId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockSessionInterface)(nil).Update), ctx, sessionId)
}

// MockPictureInterface is a mock of PictureInterface interface.
type MockPictureInterface struct {
	ctrl     *gomock.Controller
	recorder *MockPictureInterfaceMockRecorder
}

// MockPictureInterfaceMockRecorder is the mock recorder for MockPictureInterface.
type MockPictureInterfaceMockRecorder struct {
	mock *MockPictureInterface
}

// NewMockPictureInterface creates a new mock instance.
func NewMockPictureInterface(ctrl *gomock.Controller) *MockPictureInterface {
	mock := &MockPictureInterface{ctrl: ctrl}
	mock.recorder = &MockPictureInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPictureInterface) EXPECT() *MockPictureInterfaceMockRecorder {
	return m.recorder
}

// DeleteImage mocks base method.
func (m *MockPictureInterface) DeleteImage(fileName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteImage", fileName)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteImage indicates an expected call of DeleteImage.
func (mr *MockPictureInterfaceMockRecorder) DeleteImage(fileName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteImage", reflect.TypeOf((*MockPictureInterface)(nil).DeleteImage), fileName)
}

// UploadFile mocks base method.
func (m *MockPictureInterface) UploadFile(file multipart.File, fileName string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadFile", file, fileName)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UploadFile indicates an expected call of UploadFile.
func (mr *MockPictureInterfaceMockRecorder) UploadFile(file, fileName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadFile", reflect.TypeOf((*MockPictureInterface)(nil).UploadFile), file, fileName)
}
