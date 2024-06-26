// Code generated by MockGen. DO NOT EDIT.
// Source: internal/interface/repository/user.go

// Package mocks is a generated GoMock package.
package mocks

import (
	domain "auth-app/internal/domain"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// GetUserByUsernamePassword mocks base method.
func (m *MockUserRepository) GetUserByUsernamePassword(ctx context.Context, username, password string) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByUsernamePassword", ctx, username, password)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByUsernamePassword indicates an expected call of GetUserByUsernamePassword.
func (mr *MockUserRepositoryMockRecorder) GetUserByUsernamePassword(ctx, username, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByUsernamePassword", reflect.TypeOf((*MockUserRepository)(nil).GetUserByUsernamePassword), ctx, username, password)
}

// SaveUser mocks base method.
func (m *MockUserRepository) SaveUser(ctx context.Context, user *domain.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveUser", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveUser indicates an expected call of SaveUser.
func (mr *MockUserRepositoryMockRecorder) SaveUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveUser", reflect.TypeOf((*MockUserRepository)(nil).SaveUser), ctx, user)
}
