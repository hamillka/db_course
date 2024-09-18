package repositories

import (
	"github.com/hamillka/ppo/backend/internal/models"
	"github.com/jmoiron/sqlx"
)

type MedicalHistoryRepository struct {
	db *sqlx.DB
}

const (
	selectHistoryByPatient = "SELECT * FROM medical_histories WHERE patientid = $1"
	createMedicalHistory   = "INSERT INTO medical_histories " +
		"(chronic_diseases, allergies, blood_type, vaccination, patientID) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	updateMedicalHistory = "UPDATE medical_histories " +
		"SET chronic_diseases = $2, allergies = $3, blood_type = $4, vaccination = $5 WHERE patientid = $1 RETURNING id"
)

func NewMedicalHistoryRepository(db *sqlx.DB) *MedicalHistoryRepository {
	return &MedicalHistoryRepository{db: db}
}

func (mhr *MedicalHistoryRepository) GetHistoryByPatient(id int64) (models.MedicalHistory, error) {
	var history models.MedicalHistory
	err := mhr.db.QueryRow(selectHistoryByPatient, id).Scan(
		&history.ID,
		&history.ChronicDiseases,
		&history.Allergies,
		&history.BloodType,
		&history.Vaccination,
		&history.PatientID,
	)
	if err != nil {
		return models.MedicalHistory{}, ErrRecordNotFound
	}

	return history, nil
}

func (mhr *MedicalHistoryRepository) CreateMedicalHistory(
	chronicDiseases,
	allergies,
	bloodType,
	vaccination string,
	patientID int64,
) (int64, error) {
	var id int64
	err := mhr.db.QueryRow(createMedicalHistory, chronicDiseases, allergies, bloodType, vaccination, patientID).Scan(&id) //nolint:execinquery,lll //exec doesn't work
	if err != nil {
		return 0, ErrRecordAlreadyExists
	}

	return id, nil
}

func (mhr *MedicalHistoryRepository) UpdateMedicalHistory(
	id int64,
	chronicDiseases,
	allergies,
	bloodType,
	vaccination string,
) (int64, error) {
	var newID int64
	err := mhr.db.QueryRow(updateMedicalHistory, //nolint:execinquery //exec doesn't work
		id,
		chronicDiseases,
		allergies,
		bloodType,
		vaccination,
	).Scan(&newID)
	if err != nil {
		return 0, ErrDatabaseUpdatingError
	}

	return newID, nil
}
