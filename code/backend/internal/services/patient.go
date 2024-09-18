//go:generate mockgen -package=mocks -source=$GOFILE -destination=mocks/$GOFILE
package services

import "github.com/hamillka/ppo/backend/internal/models"

type PatientRepository interface {
	EditPatient(id int64, fio, phoneNumber, email, insurance string) (int64, error)
	AddPatient(fio, phoneNumber, email, insurance string) (int64, error)
	GetAllPatients() ([]models.Patient, error)
	GetPatientByID(id int64) (models.Patient, error)
}

type PatientService struct {
	repo PatientRepository
}

func NewPatientService(repository PatientRepository) *PatientService {
	return &PatientService{repo: repository}
}

func (ps *PatientService) EditPatient(
	id int64,
	fio,
	phoneNumber,
	email,
	insurance string,
) (int64, error) {
	id, err := ps.repo.EditPatient(id, fio, phoneNumber, email, insurance)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (ps *PatientService) AddPatient(
	fio,
	phoneNumber,
	email,
	insurance string,
) (int64, error) {
	id, err := ps.repo.AddPatient(fio, phoneNumber, email, insurance)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (ps *PatientService) GetAllPatients() ([]models.Patient, error) {
	patients, err := ps.repo.GetAllPatients()
	if err != nil {
		return []models.Patient{}, err
	}

	return patients, nil
}

func (ps *PatientService) GetPatientByID(id int64) (models.Patient, error) {
	patient, err := ps.repo.GetPatientByID(id)
	if err != nil {
		return models.Patient{}, err
	}

	return patient, nil
}
