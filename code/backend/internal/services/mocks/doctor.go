// Code generated by MockGen. DO NOT EDIT.
// Source: doctor.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/hamillka/ppo/backend/internal/models"
)

// MockDoctorRepository is a mock of DoctorRepository interface.
type MockDoctorRepository struct {
	ctrl     *gomock.Controller
	recorder *MockDoctorRepositoryMockRecorder
}

// MockDoctorRepositoryMockRecorder is the mock recorder for MockDoctorRepository.
type MockDoctorRepositoryMockRecorder struct {
	mock *MockDoctorRepository
}

// NewMockDoctorRepository creates a new mock instance.
func NewMockDoctorRepository(ctrl *gomock.Controller) *MockDoctorRepository {
	mock := &MockDoctorRepository{ctrl: ctrl}
	mock.recorder = &MockDoctorRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDoctorRepository) EXPECT() *MockDoctorRepositoryMockRecorder {
	return m.recorder
}

// AddDoctor mocks base method.
func (m *MockDoctorRepository) AddDoctor(fio, phoneNumber, email string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddDoctor", fio, phoneNumber, email)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddDoctor indicates an expected call of AddDoctor.
func (mr *MockDoctorRepositoryMockRecorder) AddDoctor(fio, phoneNumber, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddDoctor", reflect.TypeOf((*MockDoctorRepository)(nil).AddDoctor), fio, phoneNumber, email)
}

// EditDoctor mocks base method.
func (m *MockDoctorRepository) EditDoctor(id int64, fio, phoneNumber, email string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditDoctor", id, fio, phoneNumber, email)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EditDoctor indicates an expected call of EditDoctor.
func (mr *MockDoctorRepositoryMockRecorder) EditDoctor(id, fio, phoneNumber, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditDoctor", reflect.TypeOf((*MockDoctorRepository)(nil).EditDoctor), id, fio, phoneNumber, email)
}

// GetAllDoctors mocks base method.
func (m *MockDoctorRepository) GetAllDoctors() ([]models.Doctor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllDoctors")
	ret0, _ := ret[0].([]models.Doctor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllDoctors indicates an expected call of GetAllDoctors.
func (mr *MockDoctorRepositoryMockRecorder) GetAllDoctors() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllDoctors", reflect.TypeOf((*MockDoctorRepository)(nil).GetAllDoctors))
}

// GetDoctorByID mocks base method.
func (m *MockDoctorRepository) GetDoctorByID(id int64) (models.Doctor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDoctorByID", id)
	ret0, _ := ret[0].(models.Doctor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDoctorByID indicates an expected call of GetDoctorByID.
func (mr *MockDoctorRepositoryMockRecorder) GetDoctorByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDoctorByID", reflect.TypeOf((*MockDoctorRepository)(nil).GetDoctorByID), id)
}
