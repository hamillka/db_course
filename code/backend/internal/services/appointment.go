//go:generate mockgen -package=mocks -source=$GOFILE -destination=mocks/$GOFILE
package services

import (
	"time"

	"github.com/hamillka/ppo/backend/internal/models"
)

type AppointmentRepository interface {
	CreateAppointment(patientID, doctorID int64, dateTime time.Time) (int64, error)
	CancelAppointment(id int64) error
	GetAppointmentsByPatient(id int64) ([]*models.Appointment, error)
	GetAppointmentsByDoctor(id int64) ([]*models.Appointment, error)
	EditAppointment(id, doctorID, patientID int64, dateTime time.Time) (int64, error)
	GetAppointmentByID(id int64) (*models.Appointment, error)
}

type AppointmentService struct {
	repo AppointmentRepository
}

func NewAppointmentService(repository AppointmentRepository) *AppointmentService {
	return &AppointmentService{repo: repository}
}

func (as *AppointmentService) CreateAppointment(
	patientID,
	doctorID int64,
	dateTime time.Time,
) (int64, error) {
	id, err := as.repo.CreateAppointment(patientID, doctorID, dateTime)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (as *AppointmentService) CancelAppointment(id int64) error {
	err := as.repo.CancelAppointment(id)
	if err != nil {
		return err
	}

	return nil
}

func (as *AppointmentService) EditAppointment(
	id,
	doctorID,
	patientID int64,
	dateTime time.Time,
) (int64, error) {
	id, err := as.repo.EditAppointment(id, doctorID, patientID, dateTime)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (as *AppointmentService) GetAppointmentsByPatient(id int64) ([]*models.Appointment, error) {
	appointments, err := as.repo.GetAppointmentsByPatient(id)
	if err != nil {
		return nil, err
	}

	return appointments, nil
}

func (as *AppointmentService) GetAppointmentsByDoctor(id int64) ([]*models.Appointment, error) {
	appointments, err := as.repo.GetAppointmentsByDoctor(id)
	if err != nil {
		return nil, err
	}

	return appointments, nil
}

func (as *AppointmentService) GetAppointmentByID(id int64) (*models.Appointment, error) {
	appointment, err := as.repo.GetAppointmentByID(id)
	if err != nil {
		return nil, err
	}

	return appointment, nil
}
