package dto

import "github.com/hamillka/ppo/backend/internal/models"

type CreateOrEditOfficeRequestDto struct {
	ID       int64 `json:"id"`
	Number   int64 `json:"number"`
	Floor    int64 `json:"floor"`
	BranchID int64 `json:"branchId"`
}

type CreateOrEditOfficeResponseDto struct {
	ID int64 `json:"id"`
}

type GetOfficeResponseDto struct {
	ID       int64 `json:"id"`
	Number   int64 `json:"number"`
	Floor    int64 `json:"floor"`
	BranchID int64 `json:"branchId"`
}

func ConvertToOfficeDto(office *models.Office) *GetOfficeResponseDto {
	return &GetOfficeResponseDto{
		ID:       office.ID,
		Number:   office.Number,
		Floor:    office.Floor,
		BranchID: office.BranchID,
	}
}
