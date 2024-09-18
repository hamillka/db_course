package repositories

import (
	"database/sql"
	"errors"
	"time"

	"github.com/hamillka/ppo/backend/internal/models"
	"github.com/jmoiron/sqlx"
)

type AppointmentRepository struct {
	db *sqlx.DB
}

const (
	//createAppointment = "INSERT INTO " +
	//	"appointments (patientID, doctorID, dateTime) " +
	//	"VALUES ($1, $2, $3) " +
	//	"RETURNING id"
	createAppointment = "SELECT * from add_appointment($1, $2, $3)"
	updateAppointment = "UPDATE appointments SET doctorid = $2, patientid = $3, datetime = $4 " +
		"WHERE id = $1 RETURNING id"
	cancelAppointment        = "CALL delete_appointment($1)"
	getAppointmentsByPatient = "SELECT * FROM appointments WHERE patientID = $1"
	getAppointmentsByDoctor  = "SELECT * FROM appointments WHERE doctorID = $1"
	selectAppointmentByID    = "SELECT id, doctorid, patientid, datetime " +
		"FROM appointments WHERE id = $1"
)

func NewAppointmentRepository(db *sqlx.DB) *AppointmentRepository {
	return &AppointmentRepository{db: db}
}

func (ar *AppointmentRepository) CreateAppointment(
	patientID,
	doctorID int64,
	dateTime time.Time,
) (int64, error) {
	var id int64
	err := ar.db.QueryRow(createAppointment, doctorID, patientID, dateTime).Scan(&id) //nolint:execinquery,lll //exec doesn't work
	if err != nil {
		return 0, ErrRecordAlreadyExists
	}
	return id, nil
}

func (ar *AppointmentRepository) CancelAppointment(id int64) error {
	_, err := ar.db.Exec(cancelAppointment, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrRecordNotFound
		}

		return ErrDatabaseDeletingError
	}

	return nil
}

func (ar *AppointmentRepository) EditAppointment(
	id,
	doctorID,
	patientID int64,
	dateTime time.Time,
) (int64, error) {
	var newID int64
	err := ar.db.QueryRow(updateAppointment, //nolint:execinquery //exec doesn't work
		id,
		doctorID,
		patientID,
		dateTime.Format(time.RFC3339),
	).Scan(&newID)
	if err != nil {
		return 0, ErrDatabaseUpdatingError
	}

	return newID, nil
}

func (ar *AppointmentRepository) GetAppointmentsByPatient(id int64) ([]*models.Appointment, error) {
	var appointments []*models.Appointment

	rows, err := ar.db.Query(getAppointmentsByPatient, id)
	if err != nil {
		return nil, ErrRecordNotFound
	}
	if err := rows.Err(); err != nil {
		return nil, ErrDatabaseReadingError
	}

	for rows.Next() {
		appointment := new(models.Appointment)
		if err := rows.Scan(
			&appointment.ID,
			&appointment.DoctorID,
			&appointment.PatientID,
			&appointment.DateTime,
		); err != nil {
			return nil, ErrDatabaseReadingError
		}
		appointments = append(appointments, appointment)
	}
	defer rows.Close()

	return appointments, nil
}

func (ar *AppointmentRepository) GetAppointmentsByDoctor(id int64) ([]*models.Appointment, error) {
	var appointments []*models.Appointment

	rows, err := ar.db.Query(getAppointmentsByDoctor, id)
	if err != nil {
		return nil, ErrRecordNotFound
	}
	if err := rows.Err(); err != nil {
		return nil, ErrDatabaseReadingError
	}

	for rows.Next() {
		appointment := new(models.Appointment)
		if err := rows.Scan(
			&appointment.ID,
			&appointment.DoctorID,
			&appointment.PatientID,
			&appointment.DateTime,
		); err != nil {
			return nil, ErrDatabaseReadingError
		}
		appointments = append(appointments, appointment)
	}
	defer rows.Close()

	return appointments, nil
}

func (ar *AppointmentRepository) GetAppointmentByID(id int64) (*models.Appointment, error) {
	appointment := new(models.Appointment)

	err := ar.db.QueryRow(selectAppointmentByID, id).Scan(
		&appointment.ID,
		&appointment.DoctorID,
		&appointment.PatientID,
		&appointment.DateTime,
	)
	if err != nil {
		return nil, ErrRecordNotFound
	}

	return appointment, nil
}
