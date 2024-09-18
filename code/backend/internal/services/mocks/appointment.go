// Code generated by MockGen. DO NOT EDIT.
// Source: appointment.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	models "github.com/hamillka/ppo/backend/internal/models"
)

// MockAppointmentRepository is a mock of AppointmentRepository interface.
type MockAppointmentRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAppointmentRepositoryMockRecorder
}

// MockAppointmentRepositoryMockRecorder is the mock recorder for MockAppointmentRepository.
type MockAppointmentRepositoryMockRecorder struct {
	mock *MockAppointmentRepository
}

// NewMockAppointmentRepository creates a new mock instance.
func NewMockAppointmentRepository(ctrl *gomock.Controller) *MockAppointmentRepository {
	mock := &MockAppointmentRepository{ctrl: ctrl}
	mock.recorder = &MockAppointmentRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppointmentRepository) EXPECT() *MockAppointmentRepositoryMockRecorder {
	return m.recorder
}

// CancelAppointment mocks base method.
func (m *MockAppointmentRepository) CancelAppointment(id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelAppointment", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// CancelAppointment indicates an expected call of CancelAppointment.
func (mr *MockAppointmentRepositoryMockRecorder) CancelAppointment(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelAppointment", reflect.TypeOf((*MockAppointmentRepository)(nil).CancelAppointment), id)
}

// CreateAppointment mocks base method.
func (m *MockAppointmentRepository) CreateAppointment(patientID, doctorID int64, dateTime time.Time) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAppointment", patientID, doctorID, dateTime)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAppointment indicates an expected call of CreateAppointment.
func (mr *MockAppointmentRepositoryMockRecorder) CreateAppointment(patientID, doctorID, dateTime interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAppointment", reflect.TypeOf((*MockAppointmentRepository)(nil).CreateAppointment), patientID, doctorID, dateTime)
}

// EditAppointment mocks base method.
func (m *MockAppointmentRepository) EditAppointment(id, doctorID, patientID int64, dateTime time.Time) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditAppointment", id, doctorID, patientID, dateTime)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EditAppointment indicates an expected call of EditAppointment.
func (mr *MockAppointmentRepositoryMockRecorder) EditAppointment(id, doctorID, patientID, dateTime interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditAppointment", reflect.TypeOf((*MockAppointmentRepository)(nil).EditAppointment), id, doctorID, patientID, dateTime)
}

// GetAppointmentByID mocks base method.
func (m *MockAppointmentRepository) GetAppointmentByID(id int64) (*models.Appointment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAppointmentByID", id)
	ret0, _ := ret[0].(*models.Appointment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAppointmentByID indicates an expected call of GetAppointmentByID.
func (mr *MockAppointmentRepositoryMockRecorder) GetAppointmentByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAppointmentByID", reflect.TypeOf((*MockAppointmentRepository)(nil).GetAppointmentByID), id)
}

// GetAppointmentsByDoctor mocks base method.
func (m *MockAppointmentRepository) GetAppointmentsByDoctor(id int64) ([]*models.Appointment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAppointmentsByDoctor", id)
	ret0, _ := ret[0].([]*models.Appointment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAppointmentsByDoctor indicates an expected call of GetAppointmentsByDoctor.
func (mr *MockAppointmentRepositoryMockRecorder) GetAppointmentsByDoctor(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAppointmentsByDoctor", reflect.TypeOf((*MockAppointmentRepository)(nil).GetAppointmentsByDoctor), id)
}

// GetAppointmentsByPatient mocks base method.
func (m *MockAppointmentRepository) GetAppointmentsByPatient(id int64) ([]*models.Appointment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAppointmentsByPatient", id)
	ret0, _ := ret[0].([]*models.Appointment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAppointmentsByPatient indicates an expected call of GetAppointmentsByPatient.
func (mr *MockAppointmentRepositoryMockRecorder) GetAppointmentsByPatient(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAppointmentsByPatient", reflect.TypeOf((*MockAppointmentRepository)(nil).GetAppointmentsByPatient), id)
}
