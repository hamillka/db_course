package repositories

import (
	"database/sql"
	"errors"

	"github.com/hamillka/ppo/backend/internal/models"
	"github.com/jmoiron/sqlx"
)

type PatientRepository struct {
	db *sqlx.DB
}

const (
	createPatient = "INSERT INTO patients (fio, phoneNumber, email, insurance) VALUES " +
		"($1, $2, $3, $4) RETURNING id"
	selectPatient = "SELECT * FROM patients WHERE id = $1"
	updatePatient = "UPDATE patients " +
		"SET fio = $1, phoneNumber = $2, email = $3, insurance = $4 WHERE id = $5"
	selectPatients = "SELECT * FROM patients"
)

func NewPatientRepository(db *sqlx.DB) *PatientRepository {
	return &PatientRepository{db: db}
}

func (pr *PatientRepository) EditPatient(
	id int64,
	fio,
	phoneNumber,
	email,
	insurance string,
) (int64, error) {
	patient := new(models.Patient)

	err := pr.db.QueryRow(selectPatient, id).Scan(
		&patient.ID,
		&patient.Fio,
		&patient.PhoneNumber,
		&patient.Email,
		&patient.Insurance,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, ErrRecordNotFound
		}
	}

	if fio != "" {
		patient.Fio = fio
	}
	if phoneNumber != "" {
		patient.PhoneNumber = phoneNumber
	}
	if email != "" {
		patient.Email = email
	}
	if insurance != "" {
		patient.Insurance = insurance
	}

	_, err = pr.db.Exec(updatePatient,
		patient.Fio,
		patient.PhoneNumber,
		patient.Email,
		patient.Insurance,
		patient.ID,
	)
	if err != nil {
		return 0, ErrDatabaseUpdatingError
	}
	return id, nil
}

func (pr *PatientRepository) AddPatient(
	fio,
	phoneNumber,
	email,
	insurance string,
) (int64, error) {
	var newID int64
	err := pr.db.QueryRow(createPatient, fio, phoneNumber, email, insurance).Scan(&newID) //nolint:execinquery,lll //exec doesn't work
	if err != nil {
		return 0, ErrRecordAlreadyExists
	}

	return newID, nil
}

func (pr *PatientRepository) GetAllPatients() ([]models.Patient, error) {
	var patients []models.Patient

	rows, err := pr.db.Query(selectPatients)
	if err != nil {
		return nil, ErrDatabaseReadingError
	}
	if err := rows.Err(); err != nil {
		return nil, ErrDatabaseReadingError
	}

	for rows.Next() {
		patient := new(models.Patient)
		if err := rows.Scan(
			&patient.ID,
			&patient.Fio,
			&patient.PhoneNumber,
			&patient.Email,
			&patient.Insurance,
		); err != nil {
			return nil, ErrDatabaseReadingError
		}
		patients = append(patients, *patient)
	}
	defer rows.Close()

	return patients, nil
}

func (pr *PatientRepository) GetPatientByID(id int64) (models.Patient, error) {
	var patient models.Patient

	err := pr.db.QueryRow(selectPatient, id).Scan(
		&patient.ID,
		&patient.Fio,
		&patient.PhoneNumber,
		&patient.Email,
		&patient.Insurance,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Patient{}, ErrRecordNotFound
		}
	}

	return patient, nil
}
