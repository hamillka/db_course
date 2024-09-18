//go:generate mockgen -package=mocks -source=$GOFILE -destination=mocks/$GOFILE
package services

import "github.com/hamillka/ppo/backend/internal/models"

type BranchRepository interface {
	EditBranch(id int64, name, address, phoneNumber string) (int64, error)
	AddBranch(name, address, phoneNumber string) (int64, error)
	GetAllBranches() ([]*models.Branch, error)
	GetBranchByID(id int64) (models.Branch, error)
}

type BranchService struct {
	repo BranchRepository
}

func NewBranchService(repository BranchRepository) *BranchService {
	return &BranchService{repo: repository}
}

func (bs *BranchService) EditBranch(id int64, name, address, phoneNumber string) (int64, error) {
	id, err := bs.repo.EditBranch(id, name, address, phoneNumber)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (bs *BranchService) AddBranch(name, address, phoneNumber string) (int64, error) {
	id, err := bs.repo.AddBranch(name, address, phoneNumber)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (bs *BranchService) GetAllBranches() ([]*models.Branch, error) {
	branches, err := bs.repo.GetAllBranches()
	if err != nil {
		return []*models.Branch{}, err
	}

	return branches, nil
}

func (bs *BranchService) GetBranchByID(id int64) (models.Branch, error) {
	branch, err := bs.repo.GetBranchByID(id)
	if err != nil {
		return models.Branch{}, err
	}

	return branch, nil
}
