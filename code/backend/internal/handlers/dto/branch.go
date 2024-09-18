package dto

import "github.com/hamillka/ppo/backend/internal/models"

type CreateOrEditBranchRequestDto struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phoneNumber"`
	ID          int64  `json:"id"`
}

type CreateOrEditBranchResponseDto struct {
	ID int64 `json:"id"`
}

type GetBranchResponseDto struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phoneNumber"`
	ID          int64  `json:"id"`
}

func ConvertToBranchDto(branch *models.Branch) *GetBranchResponseDto {
	return &GetBranchResponseDto{
		Name:        branch.Name,
		Address:     branch.Address,
		PhoneNumber: branch.PhoneNumber,
		ID:          branch.ID,
	}
}
