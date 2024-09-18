package models

type User struct {
	PatientID *int64
	DoctorID  *int64
	Login     string
	Password  string
	Role      int64
	ID        int64
}
