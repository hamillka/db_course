package models

type MedicalHistory struct {
	ChronicDiseases *string
	Allergies       *string
	BloodType       *string
	Vaccination     *string
	ID              int64
	PatientID       int64
}
