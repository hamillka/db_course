package dto

import (
	"time"
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
