//go:generate mockgen -package=mocks -source=$GOFILE -destination=mocks/$GOFILE
package services

import "github.com/hamillka/ppo/backend/internal/models"

type OfficeRepository interface {
	EditOffice(id, number, floor, branchID int64) (int64, error)
	AddOffice(number, floor, branchID int64) (int64, error)
	GetAllOffices() ([]models.Office, error)
	GetOfficeByID(id int64) (models.Office, error)
}

type OfficeService struct {
	repo OfficeRepository
}

func NewOfficeService(repository OfficeRepository) *OfficeService {
	return &OfficeService{repo: repository}
}

func (os *OfficeService) EditOffice(id, number, floor, branchID int64) (int64, error) {
	id, err := os.repo.EditOffice(id, number, floor, branchID)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (os *OfficeService) AddOffice(number, floor, branchID int64) (int64, error) {
	id, err := os.repo.AddOffice(number, floor, branchID)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (os *OfficeService) GetAllOffices() ([]models.Office, error) {
	offices, err := os.repo.GetAllOffices()
	if err != nil {
		return []models.Office{}, err
	}

	return offices, nil
}

func (os *OfficeService) GetOfficeByID(id int64) (models.Office, error) {
	office, err := os.repo.GetOfficeByID(id)
	if err != nil {
		return models.Office{}, err
	}

	return office, nil
}
