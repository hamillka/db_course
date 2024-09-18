package dto

import "github.com/hamillka/ppo/backend/internal/models"

type CreateOrEditPatientRequestDto struct {
	Fio         string `json:"fio"`
	PhoneNumber string `json:"phoneNumber"`
	Email       string `json:"email"`
	Insurance   string `json:"insurance"`
	ID          int64  `json:"id"`
}

type CreateOrEditPatientResponseDto struct {
	ID int64 `json:"id"`
}

type GetPatientResponseDto struct {
	Fio         string `json:"fio"`
	PhoneNumber string `json:"phoneNumber"`
	Email       string `json:"email"`
	Insurance   string `json:"insurance"`
	ID          int64  `json:"id"`
}

func ConvertToPatientDto(patient *models.Patient) *GetPatientResponseDto {
	return &GetPatientResponseDto{
		Fio:         patient.Fio,
		PhoneNumber: patient.PhoneNumber,
		Email:       patient.Email,
		Insurance:   patient.Insurance,
		ID:          patient.ID,
	}
}
