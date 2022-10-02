// Code generated by MockGen. DO NOT EDIT.
// Source: contract.go

// Package usecase is a generated GoMock package.
package usecase

import (
	context "context"
	reflect "reflect"

	models "github.com/Moxxx1e/rsoi-2022-lab1-ci-cd-Moxxx1e/internal/models"
	gomock "github.com/golang/mock/gomock"
)

// Mockrepo is a mock of repo interface.
type Mockrepo struct {
	ctrl     *gomock.Controller
	recorder *MockrepoMockRecorder
}

// MockrepoMockRecorder is the mock recorder for Mockrepo.
type MockrepoMockRecorder struct {
	mock *Mockrepo
}

// NewMockrepo creates a new mock instance.
func NewMockrepo(ctrl *gomock.Controller) *Mockrepo {
	mock := &Mockrepo{ctrl: ctrl}
	mock.recorder = &MockrepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockrepo) EXPECT() *MockrepoMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *Mockrepo) Create(ctx context.Context, person models.Person) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, person)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockrepoMockRecorder) Create(ctx, person interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*Mockrepo)(nil).Create), ctx, person)
}

// Delete mocks base method.
func (m *Mockrepo) Delete(ctx context.Context, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockrepoMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*Mockrepo)(nil).Delete), ctx, id)
}

// GetAll mocks base method.
func (m *Mockrepo) GetAll(ctx context.Context) (*[]models.Person, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx)
	ret0, _ := ret[0].(*[]models.Person)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockrepoMockRecorder) GetAll(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*Mockrepo)(nil).GetAll), ctx)
}

// GetByID mocks base method.
func (m *Mockrepo) GetByID(ctx context.Context, id int64) (models.Person, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, id)
	ret0, _ := ret[0].(models.Person)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockrepoMockRecorder) GetByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*Mockrepo)(nil).GetByID), ctx, id)
}

// Update mocks base method.
func (m *Mockrepo) Update(ctx context.Context, person models.Person) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, person)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockrepoMockRecorder) Update(ctx, person interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*Mockrepo)(nil).Update), ctx, person)
}