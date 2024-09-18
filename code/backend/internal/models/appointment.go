package models

import (
	"time"
)

type Appointment struct {
	DateTime  time.Time
	ID        int64
	PatientID int64
	DoctorID  int64
}
