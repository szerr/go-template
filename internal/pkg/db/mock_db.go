// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/pkg/db/db.go
//
// Generated by this command:
//
//	mockgen -source=./internal/pkg/db/db.go -destination=./internal/pkg/db/mock_db.go -package=db
//

// Package db is a generated GoMock package.
package db

import (
	sql "database/sql"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockIDB is a mock of IDB interface.
type MockIDB struct {
	ctrl     *gomock.Controller
	recorder *MockIDBMockRecorder
	isgomock struct{}
}

// MockIDBMockRecorder is the mock recorder for MockIDB.
type MockIDBMockRecorder struct {
	mock *MockIDB
}

// NewMockIDB creates a new mock instance.
func NewMockIDB(ctrl *gomock.Controller) *MockIDB {
	mock := &MockIDB{ctrl: ctrl}
	mock.recorder = &MockIDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIDB) EXPECT() *MockIDBMockRecorder {
	return m.recorder
}

// Begin mocks base method.
func (m *MockIDB) Begin(opts ...*sql.TxOptions) *DB {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Begin", varargs...)
	ret0, _ := ret[0].(*DB)
	return ret0
}

// Begin indicates an expected call of Begin.
func (mr *MockIDBMockRecorder) Begin(opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Begin", reflect.TypeOf((*MockIDB)(nil).Begin), opts...)
}

// Commit mocks base method.
func (m *MockIDB) Commit() *DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit")
	ret0, _ := ret[0].(*DB)
	return ret0
}

// Commit indicates an expected call of Commit.
func (mr *MockIDBMockRecorder) Commit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockIDB)(nil).Commit))
}

// Count mocks base method.
func (m *MockIDB) Count(count *int64) *DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", count)
	ret0, _ := ret[0].(*DB)
	return ret0
}

// Count indicates an expected call of Count.
func (mr *MockIDBMockRecorder) Count(count any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockIDB)(nil).Count), count)
}

// Create mocks base method.
func (m *MockIDB) Create(value any) *DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", value)
	ret0, _ := ret[0].(*DB)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockIDBMockRecorder) Create(value any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIDB)(nil).Create), value)
}

// Delete mocks base method.
func (m *MockIDB) Delete(value any, conds ...any) *DB {
	m.ctrl.T.Helper()
	varargs := []any{value}
	for _, a := range conds {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(*DB)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockIDBMockRecorder) Delete(value any, conds ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{value}, conds...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIDB)(nil).Delete), varargs...)
}

// Error mocks base method.
func (m *MockIDB) Error() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Error")
	ret0, _ := ret[0].(error)
	return ret0
}

// Error indicates an expected call of Error.
func (mr *MockIDBMockRecorder) Error() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockIDB)(nil).Error))
}

// Find mocks base method.
func (m *MockIDB) Find(dest any, conds ...any) *DB {
	m.ctrl.T.Helper()
	varargs := []any{dest}
	for _, a := range conds {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Find", varargs...)
	ret0, _ := ret[0].(*DB)
	return ret0
}

// Find indicates an expected call of Find.
func (mr *MockIDBMockRecorder) Find(dest any, conds ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{dest}, conds...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockIDB)(nil).Find), varargs...)
}

// First mocks base method.
func (m *MockIDB) First(dest any, conds ...any) *DB {
	m.ctrl.T.Helper()
	varargs := []any{dest}
	for _, a := range conds {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "First", varargs...)
	ret0, _ := ret[0].(*DB)
	return ret0
}

// First indicates an expected call of First.
func (mr *MockIDBMockRecorder) First(dest any, conds ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{dest}, conds...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "First", reflect.TypeOf((*MockIDB)(nil).First), varargs...)
}

// Joins mocks base method.
func (m *MockIDB) Joins(query string, args ...any) *DB {
	m.ctrl.T.Helper()
	varargs := []any{query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Joins", varargs...)
	ret0, _ := ret[0].(*DB)
	return ret0
}

// Joins indicates an expected call of Joins.
func (mr *MockIDBMockRecorder) Joins(query any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Joins", reflect.TypeOf((*MockIDB)(nil).Joins), varargs...)
}

// Last mocks base method.
func (m *MockIDB) Last(dest any, conds ...any) *DB {
	m.ctrl.T.Helper()
	varargs := []any{dest}
	for _, a := range conds {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Last", varargs...)
	ret0, _ := ret[0].(*DB)
	return ret0
}

// Last indicates an expected call of Last.
func (mr *MockIDBMockRecorder) Last(dest any, conds ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{dest}, conds...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Last", reflect.TypeOf((*MockIDB)(nil).Last), varargs...)
}

// Limit mocks base method.
func (m *MockIDB) Limit(limit int) *DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Limit", limit)
	ret0, _ := ret[0].(*DB)
	return ret0
}

// Limit indicates an expected call of Limit.
func (mr *MockIDBMockRecorder) Limit(limit any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Limit", reflect.TypeOf((*MockIDB)(nil).Limit), limit)
}

// Model mocks base method.
func (m *MockIDB) Model(value any) *DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Model", value)
	ret0, _ := ret[0].(*DB)
	return ret0
}

// Model indicates an expected call of Model.
func (mr *MockIDBMockRecorder) Model(value any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Model", reflect.TypeOf((*MockIDB)(nil).Model), value)
}

// Not mocks base method.
func (m *MockIDB) Not(query any, args ...any) *DB {
	m.ctrl.T.Helper()
	varargs := []any{query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Not", varargs...)
	ret0, _ := ret[0].(*DB)
	return ret0
}

// Not indicates an expected call of Not.
func (mr *MockIDBMockRecorder) Not(query any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Not", reflect.TypeOf((*MockIDB)(nil).Not), varargs...)
}

// Offset mocks base method.
func (m *MockIDB) Offset(offset int) *DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Offset", offset)
	ret0, _ := ret[0].(*DB)
	return ret0
}

// Offset indicates an expected call of Offset.
func (mr *MockIDBMockRecorder) Offset(offset any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Offset", reflect.TypeOf((*MockIDB)(nil).Offset), offset)
}

// Omit mocks base method.
func (m *MockIDB) Omit(columns ...string) *DB {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range columns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Omit", varargs...)
	ret0, _ := ret[0].(*DB)
	return ret0
}

// Omit indicates an expected call of Omit.
func (mr *MockIDBMockRecorder) Omit(columns ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Omit", reflect.TypeOf((*MockIDB)(nil).Omit), columns...)
}

// Or mocks base method.
func (m *MockIDB) Or(query any, args ...any) *DB {
	m.ctrl.T.Helper()
	varargs := []any{query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Or", varargs...)
	ret0, _ := ret[0].(*DB)
	return ret0
}

// Or indicates an expected call of Or.
func (mr *MockIDBMockRecorder) Or(query any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Or", reflect.TypeOf((*MockIDB)(nil).Or), varargs...)
}

// Order mocks base method.
func (m *MockIDB) Order(value any) *DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Order", value)
	ret0, _ := ret[0].(*DB)
	return ret0
}

// Order indicates an expected call of Order.
func (mr *MockIDBMockRecorder) Order(value any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Order", reflect.TypeOf((*MockIDB)(nil).Order), value)
}

// Rollback mocks base method.
func (m *MockIDB) Rollback() *DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rollback")
	ret0, _ := ret[0].(*DB)
	return ret0
}

// Rollback indicates an expected call of Rollback.
func (mr *MockIDBMockRecorder) Rollback() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rollback", reflect.TypeOf((*MockIDB)(nil).Rollback))
}

// RollbackTo mocks base method.
func (m *MockIDB) RollbackTo(name string) *DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RollbackTo", name)
	ret0, _ := ret[0].(*DB)
	return ret0
}

// RollbackTo indicates an expected call of RollbackTo.
func (mr *MockIDBMockRecorder) RollbackTo(name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RollbackTo", reflect.TypeOf((*MockIDB)(nil).RollbackTo), name)
}

// RowsAffected mocks base method.
func (m *MockIDB) RowsAffected() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RowsAffected")
	ret0, _ := ret[0].(int64)
	return ret0
}

// RowsAffected indicates an expected call of RowsAffected.
func (mr *MockIDBMockRecorder) RowsAffected() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RowsAffected", reflect.TypeOf((*MockIDB)(nil).RowsAffected))
}

// Save mocks base method.
func (m *MockIDB) Save(value any) *DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", value)
	ret0, _ := ret[0].(*DB)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockIDBMockRecorder) Save(value any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockIDB)(nil).Save), value)
}

// SavePoint mocks base method.
func (m *MockIDB) SavePoint(name string) *DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SavePoint", name)
	ret0, _ := ret[0].(*DB)
	return ret0
}

// SavePoint indicates an expected call of SavePoint.
func (mr *MockIDBMockRecorder) SavePoint(name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SavePoint", reflect.TypeOf((*MockIDB)(nil).SavePoint), name)
}

// Scan mocks base method.
func (m *MockIDB) Scan(dest any) *DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Scan", dest)
	ret0, _ := ret[0].(*DB)
	return ret0
}

// Scan indicates an expected call of Scan.
func (mr *MockIDBMockRecorder) Scan(dest any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scan", reflect.TypeOf((*MockIDB)(nil).Scan), dest)
}

// Select mocks base method.
func (m *MockIDB) Select(query any, args ...any) *DB {
	m.ctrl.T.Helper()
	varargs := []any{query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Select", varargs...)
	ret0, _ := ret[0].(*DB)
	return ret0
}

// Select indicates an expected call of Select.
func (mr *MockIDBMockRecorder) Select(query any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Select", reflect.TypeOf((*MockIDB)(nil).Select), varargs...)
}

// Take mocks base method.
func (m *MockIDB) Take(dest any, conds ...any) *DB {
	m.ctrl.T.Helper()
	varargs := []any{dest}
	for _, a := range conds {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Take", varargs...)
	ret0, _ := ret[0].(*DB)
	return ret0
}

// Take indicates an expected call of Take.
func (mr *MockIDBMockRecorder) Take(dest any, conds ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{dest}, conds...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Take", reflect.TypeOf((*MockIDB)(nil).Take), varargs...)
}

// Transaction mocks base method.
func (m *MockIDB) Transaction(fc func(*DB) error, opts ...*sql.TxOptions) error {
	m.ctrl.T.Helper()
	varargs := []any{fc}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Transaction", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Transaction indicates an expected call of Transaction.
func (mr *MockIDBMockRecorder) Transaction(fc any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{fc}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Transaction", reflect.TypeOf((*MockIDB)(nil).Transaction), varargs...)
}

// Update mocks base method.
func (m *MockIDB) Update(column string, value any) *DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", column, value)
	ret0, _ := ret[0].(*DB)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockIDBMockRecorder) Update(column, value any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIDB)(nil).Update), column, value)
}

// Updates mocks base method.
func (m *MockIDB) Updates(value any) *DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Updates", value)
	ret0, _ := ret[0].(*DB)
	return ret0
}

// Updates indicates an expected call of Updates.
func (mr *MockIDBMockRecorder) Updates(value any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Updates", reflect.TypeOf((*MockIDB)(nil).Updates), value)
}

// Where mocks base method.
func (m *MockIDB) Where(query any, args ...any) *DB {
	m.ctrl.T.Helper()
	varargs := []any{query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Where", varargs...)
	ret0, _ := ret[0].(*DB)
	return ret0
}

// Where indicates an expected call of Where.
func (mr *MockIDBMockRecorder) Where(query any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Where", reflect.TypeOf((*MockIDB)(nil).Where), varargs...)
}
