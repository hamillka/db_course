package dto

import (
	"time"

	"github.com/hamillka/ppo/backend/internal/models"
)

type CreateOrEditAppointmentRequestDto struct {
	DateTime  time.Time `json:"dateTime"`
	ID        int64     `json:"id"`
	PatientID int64     `json:"patientId"`
	DoctorID  int64     `json:"doctorId"`
}

type CreateOrEditAppointmentResponseDto struct {
	ID int64 `json:"id"`
}

type CancelAppointmentRequestDto struct {
	ID int64 `json:"id"`
}

type GetAppointmentResponseDto struct {
	DateTime  time.Time `json:"dateTime"`
	ID        int64     `json:"id"`
	PatientID int64     `json:"patientId"`
	DoctorID  int64     `json:"doctorId"`
}

func ConvertToDto(appointment *models.Appointment) *GetAppointmentResponseDto {
	return &GetAppointmentResponseDto{
		DateTime:  appointment.DateTime,
		ID:        appointment.ID,
		PatientID: appointment.PatientID,
		DoctorID:  appointment.DoctorID,
	}
}
