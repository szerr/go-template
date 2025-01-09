// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/biz/user.go

// Package mock_biz is a generated GoMock package.
package mock_biz

import (
	context "context"
	request "go-template/internal/domain"
	model "go-template/internal/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIUserRepo is a mock of IUserRepo interface.
type MockIUserRepo struct {
	ctrl     *gomock.Controller
	recorder *MockIUserRepoMockRecorder
}

// MockIUserRepoMockRecorder is the mock recorder for MockIUserRepo.
type MockIUserRepoMockRecorder struct {
	mock *MockIUserRepo
}

// NewMockIUserRepo creates a new mock instance.
func NewMockIUserRepo(ctrl *gomock.Controller) *MockIUserRepo {
	mock := &MockIUserRepo{ctrl: ctrl}
	mock.recorder = &MockIUserRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUserRepo) EXPECT() *MockIUserRepoMockRecorder {
	return m.recorder
}

// Count mocks base method.
func (m *MockIUserRepo) Count(ctx context.Context, req *request.UserListRequest) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, req)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockIUserRepoMockRecorder) Count(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockIUserRepo)(nil).Count), ctx, req)
}

// Create mocks base method.
func (m *MockIUserRepo) Create(ctx context.Context, cp *model.SysUser) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, cp)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockIUserRepoMockRecorder) Create(ctx, cp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIUserRepo)(nil).Create), ctx, cp)
}

// Delete mocks base method.
func (m *MockIUserRepo) Delete(ctx context.Context, id uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockIUserRepoMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIUserRepo)(nil).Delete), ctx, id)
}

// GetByName mocks base method.
func (m *MockIUserRepo) GetByName(ctx context.Context, userName string) (*model.SysUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByUserName", ctx, userName)
	ret0, _ := ret[0].(*model.SysUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByName indicates an expected call of GetByName.
func (mr *MockIUserRepoMockRecorder) GetByName(ctx, userName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUserName", reflect.TypeOf((*MockIUserRepo)(nil).GetByName), ctx, userName)
}

// List mocks base method.
func (m *MockIUserRepo) List(ctx context.Context, req *request.UserListRequest) ([]*model.SysUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, req)
	ret0, _ := ret[0].([]*model.SysUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockIUserRepoMockRecorder) List(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockIUserRepo)(nil).List), ctx, req)
}

// Retrieve mocks base method.
func (m *MockIUserRepo) Retrieve(ctx context.Context, id uint64) (*model.SysUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Retrieve", ctx, id)
	ret0, _ := ret[0].(*model.SysUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Retrieve indicates an expected call of Retrieve.
func (mr *MockIUserRepoMockRecorder) Retrieve(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Retrieve", reflect.TypeOf((*MockIUserRepo)(nil).Retrieve), ctx, id)
}

// Update mocks base method.
func (m *MockIUserRepo) Update(ctx context.Context, cp *model.SysUser) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, cp)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockIUserRepoMockRecorder) Update(ctx, cp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIUserRepo)(nil).Update), ctx, cp)
}
