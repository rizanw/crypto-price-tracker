// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package mock is a generated GoMock package.
package mock

import (
	coincap "crypto-tracker/internal/model/coincap"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRepo is a mock of Repo interface.
type MockRepo struct {
	ctrl     *gomock.Controller
	recorder *MockRepoMockRecorder
}

// MockRepoMockRecorder is the mock recorder for MockRepo.
type MockRepoMockRecorder struct {
	mock *MockRepo
}

// NewMockRepo creates a new mock instance.
func NewMockRepo(ctrl *gomock.Controller) *MockRepo {
	mock := &MockRepo{ctrl: ctrl}
	mock.recorder = &MockRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepo) EXPECT() *MockRepoMockRecorder {
	return m.recorder
}

// FindRate mocks base method.
func (m *MockRepo) FindRate(id string) (coincap.Rate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindRate", id)
	ret0, _ := ret[0].(coincap.Rate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindRate indicates an expected call of FindRate.
func (mr *MockRepoMockRecorder) FindRate(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindRate", reflect.TypeOf((*MockRepo)(nil).FindRate), id)
}