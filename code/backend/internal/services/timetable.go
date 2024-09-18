//go:generate mockgen -package=mocks -source=$GOFILE -destination=mocks/$GOFILE
package services

import "github.com/hamillka/ppo/backend/internal/models"

type TimetableRepository interface {
	GetLocationsByDoctor(id int64) ([]models.Office, error)
	GetDoctorsByLocation(id int64) ([]models.Doctor, error)
	GetWorkdaysByDoctor(id int64) ([]*models.Timetable, error)
}

type TimetableService struct {
	repo TimetableRepository
}

func NewTimetableService(repository TimetableRepository) *TimetableService {
	return &TimetableService{repo: repository}
}

func (ts *TimetableService) GetLocationsByDoctor(id int64) ([]models.Office, error) {
	offices, err := ts.repo.GetLocationsByDoctor(id)
	if err != nil {
		return []models.Office{}, err
	}

	return offices, nil
}

func (ts *TimetableService) GetDoctorsByLocation(id int64) ([]models.Doctor, error) {
	doctors, err := ts.repo.GetDoctorsByLocation(id)
	if err != nil {
		return []models.Doctor{}, err
	}

	return doctors, nil
}

func (ts *TimetableService) GetWorkdaysByDoctor(id int64) ([]*models.Timetable, error) {
	timetable, err := ts.repo.GetWorkdaysByDoctor(id)
	if err != nil {
		return nil, err
	}

	return timetable, nil
}
