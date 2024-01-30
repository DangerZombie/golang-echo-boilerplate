// Code generated by MockGen. DO NOT EDIT.
// Source: repository/repository_user/repository.go
//
// Generated by this command:
//
//	mockgen -source=repository/repository_user/repository.go -destination=repository/repository_user/repository.mock.gen.go -package=repository_user
//

// Package repository_user is a generated GoMock package.
package repository_user

import (
	entity "go-echo/model/entity"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	gorm "gorm.io/gorm"
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

// FindUserByUsernameAndPassword mocks base method.
func (m *MockUserRepository) FindUserByUsernameAndPassword(db *gorm.DB, username, password string) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByUsernameAndPassword", db, username, password)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByUsernameAndPassword indicates an expected call of FindUserByUsernameAndPassword.
func (mr *MockUserRepositoryMockRecorder) FindUserByUsernameAndPassword(db, username, password any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByUsernameAndPassword", reflect.TypeOf((*MockUserRepository)(nil).FindUserByUsernameAndPassword), db, username, password)
}
