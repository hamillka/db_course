//go:generate mockgen -package=mocks -source=$GOFILE -destination=mocks/$GOFILE
package services

import "github.com/hamillka/ppo/backend/internal/models"

type DoctorRepository interface {
	EditDoctor(id int64, fio, phoneNumber, email, specialization string) (int64, error)
	AddDoctor(fio, phoneNumber, email, specialization string) (int64, error)
	GetAllDoctors() ([]models.Doctor, error)
	GetDoctorByID(id int64) (models.Doctor, error)
}

type DoctorService struct {
	doctorRepo    DoctorRepository
	timetableRepo TimetableRepository
}

func NewDoctorService(dr DoctorRepository, tr TimetableRepository) *DoctorService {
	return &DoctorService{doctorRepo: dr, timetableRepo: tr}
}

func (ds *DoctorService) EditDoctor(id int64, fio, phoneNumber, email, specialization string) (int64, error) {
	id, err := ds.doctorRepo.EditDoctor(id, fio, phoneNumber, email, specialization)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (ds *DoctorService) AddDoctor(fio, phoneNumber, email, specialization string) (int64, error) {
	id, err := ds.doctorRepo.AddDoctor(fio, phoneNumber, email, specialization)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (ds *DoctorService) GetAllDoctors() ([]models.Doctor, error) {
	doctors, err := ds.doctorRepo.GetAllDoctors()
	if err != nil {
		return []models.Doctor{}, err
	}

	return doctors, nil
}

func (ds *DoctorService) GetDoctorByID(id int64) (models.Doctor, error) {
	doctor, err := ds.doctorRepo.GetDoctorByID(id)
	if err != nil {
		return models.Doctor{}, err
	}

	return doctor, nil
}
