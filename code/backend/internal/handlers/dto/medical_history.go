package dto

import (
	"github.com/hamillka/ppo/backend/internal/models"
)

type GetMedicalHistoryResponseDto struct {
	ChronicDiseases string `json:"chronicDiseases"`
	Allergies       string `json:"allergies"`
	BloodType       string `json:"bloodType"`
	Vaccination     string `json:"vaccination"`
	ID              int64  `json:"id"`
	PatientID       int64  `json:"patientId"`
}

type CreateOrEditMedicalHistoryRequestDto struct {
	ChronicDiseases string `json:"chronicDiseases"`
	Allergies       string `json:"allergies"`
	BloodType       string `json:"bloodType"`
	Vaccination     string `json:"vaccination"`
	ID              int64  `json:"id"`
	PatientID       int64  `json:"patientId"`
}

type CreateOrEditMedicalHistoryResponseDto struct {
	ID int64 `json:"id"`
}

func ConvertToMedicalHistoryDto(history models.MedicalHistory) *GetMedicalHistoryResponseDto {
	resp := &GetMedicalHistoryResponseDto{
		ID:        history.ID,
		PatientID: history.PatientID,
	}
	if history.ChronicDiseases == nil {
		resp.ChronicDiseases = ""
	} else {
		resp.ChronicDiseases = *history.ChronicDiseases
	}

	if history.Allergies == nil {
		resp.Allergies = ""
	} else {
		resp.Allergies = *history.Allergies
	}

	if history.BloodType == nil {
		resp.BloodType = ""
	} else {
		resp.BloodType = *history.BloodType
	}

	if history.Vaccination == nil {
		resp.Vaccination = ""
	} else {
		resp.Vaccination = *history.Vaccination
	}

	return resp
}
