package dto

import "github.com/hamillka/ppo/backend/internal/models"

type CreateOrEditDoctorRequestDto struct {
	Fio            string `json:"fio"`
	PhoneNumber    string `json:"phoneNumber"`
	Email          string `json:"email"`
	ID             int64  `json:"id"`
	Specialization string `json:"specialization"`
}

type CreateOrEditDoctorResponseDto struct {
	ID int64 `json:"id"`
}

type GetDoctorResponseDto struct {
	Fio            string `json:"fio"`
	PhoneNumber    string `json:"phoneNumber"`
	Email          string `json:"email"`
	ID             int64  `json:"id"`
	Specialization string `json:"specialization"`
}

func ConvertToDoctorDto(doctor *models.Doctor) *GetDoctorResponseDto {
	return &GetDoctorResponseDto{
		Fio:            doctor.Fio,
		PhoneNumber:    doctor.PhoneNumber,
		Email:          doctor.Email,
		ID:             doctor.ID,
		Specialization: doctor.Specialization,
	}
}
