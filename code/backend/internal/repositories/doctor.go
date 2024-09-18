package repositories

import (
	"database/sql"
	"errors"

	"github.com/hamillka/ppo/backend/internal/models"
	"github.com/jmoiron/sqlx"
)

type DoctorRepository struct {
	db *sqlx.DB
}

const (
	createDoctor  = "INSERT INTO doctors (fio, phoneNumber, email, specialization) VALUES ($1, $2, $3, $4) RETURNING id"
	selectDoctor  = "SELECT * FROM doctors WHERE id = $1"
	updateDoctor  = "UPDATE doctors SET fio = $1, phoneNumber = $2, email = $3, specialization = $4 WHERE id = $5"
	selectDoctors = "SELECT * FROM doctors"
)

func NewDoctorRepository(db *sqlx.DB) *DoctorRepository {
	return &DoctorRepository{db: db}
}

func (dr *DoctorRepository) EditDoctor(id int64, fio, phoneNumber, email, specialization string) (int64, error) {
	doctor := new(models.Doctor)

	err := dr.db.QueryRow(selectDoctor, id).Scan(
		&doctor.ID,
		&doctor.Fio,
		&doctor.PhoneNumber,
		&doctor.Email,
		&doctor.Specialization,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, ErrRecordNotFound
		}
	}

	if fio != "" {
		doctor.Fio = fio
	}
	if phoneNumber != "" {
		doctor.PhoneNumber = phoneNumber
	}
	if email != "" {
		doctor.Email = email
	}
	if specialization != "" {
		doctor.Specialization = specialization
	}

	_, err = dr.db.Exec(updateDoctor, doctor.Fio, doctor.PhoneNumber, doctor.Email, doctor.Specialization, doctor.ID)
	if err != nil {
		return 0, ErrDatabaseUpdatingError
	}
	return id, nil
}

func (dr *DoctorRepository) AddDoctor(fio, phoneNumber, email, specialization string) (int64, error) {
	var id int64
	err := dr.db.QueryRow(createDoctor, fio, phoneNumber, email, specialization).Scan(&id) //nolint:execinquery,lll //exec doesn't work
	if err != nil {
		return 0, ErrRecordAlreadyExists
	}

	for i := 1; i < 7; i++ {
		_, err := dr.db.Exec("INSERT INTO timetable (doctorid, officeid, workdays) VALUES ($1, $2, $3)",
			id, 100, i)
		if err != nil {
			return 0, err
		}
	}
	return id, nil
}

func (dr *DoctorRepository) GetAllDoctors() ([]models.Doctor, error) {
	var doctors []models.Doctor

	rows, err := dr.db.Query(selectDoctors)
	if err != nil {
		return nil, ErrDatabaseReadingError
	}
	if err := rows.Err(); err != nil {
		return nil, ErrDatabaseReadingError
	}

	for rows.Next() {
		doctor := new(models.Doctor)
		if err := rows.Scan(
			&doctor.ID,
			&doctor.Fio,
			&doctor.PhoneNumber,
			&doctor.Email,
			&doctor.Specialization,
		); err != nil {
			return nil, ErrDatabaseReadingError
		}
		doctors = append(doctors, *doctor)
	}
	defer rows.Close()

	return doctors, nil
}

func (dr *DoctorRepository) GetDoctorByID(id int64) (models.Doctor, error) {
	var doctor models.Doctor

	err := dr.db.QueryRow(selectDoctor, id).Scan(
		&doctor.ID,
		&doctor.Fio,
		&doctor.PhoneNumber,
		&doctor.Email,
		&doctor.Specialization,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Doctor{}, ErrRecordNotFound
		}
	}

	return doctor, nil
}
