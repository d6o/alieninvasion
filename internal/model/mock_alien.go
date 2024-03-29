// Code generated by MockGen. DO NOT EDIT.
// Source: alien.go

// Package model is a generated GoMock package.
package model

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAlienRepository is a mock of AlienRepository interface.
type MockAlienRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAlienRepositoryMockRecorder
}

// MockAlienRepositoryMockRecorder is the mock recorder for MockAlienRepository.
type MockAlienRepositoryMockRecorder struct {
	mock *MockAlienRepository
}

// NewMockAlienRepository creates a new mock instance.
func NewMockAlienRepository(ctrl *gomock.Controller) *MockAlienRepository {
	mock := &MockAlienRepository{ctrl: ctrl}
	mock.recorder = &MockAlienRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAlienRepository) EXPECT() *MockAlienRepositoryMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockAlienRepository) Add(alien *Alien) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Add", alien)
}

// Add indicates an expected call of Add.
func (mr *MockAlienRepositoryMockRecorder) Add(alien interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockAlienRepository)(nil).Add), alien)
}

// All mocks base method.
func (m *MockAlienRepository) All() map[int]*Alien {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All")
	ret0, _ := ret[0].(map[int]*Alien)
	return ret0
}

// All indicates an expected call of All.
func (mr *MockAlienRepositoryMockRecorder) All() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockAlienRepository)(nil).All))
}

// NextID mocks base method.
func (m *MockAlienRepository) NextID() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NextID")
	ret0, _ := ret[0].(int)
	return ret0
}

// NextID indicates an expected call of NextID.
func (mr *MockAlienRepositoryMockRecorder) NextID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NextID", reflect.TypeOf((*MockAlienRepository)(nil).NextID))
}

// Remove mocks base method.
func (m *MockAlienRepository) Remove(aliens map[int]*Alien) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Remove", aliens)
}

// Remove indicates an expected call of Remove.
func (mr *MockAlienRepositoryMockRecorder) Remove(aliens interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockAlienRepository)(nil).Remove), aliens)
}
