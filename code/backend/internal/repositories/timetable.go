package repositories

import (
	"github.com/hamillka/ppo/backend/internal/models"
	"github.com/jmoiron/sqlx"
)

type TimetableRepository struct {
	db *sqlx.DB
}

const (
	selectLocationsByDoctor = "SELECT o.id, o.number, o.floor, o.branchID " +
		"FROM timetable t JOIN offices o ON t.officeID = o.id WHERE t.doctorID = $1"
	selectDoctorsByLocation = "SELECT d.id, d.fio, d.phoneNumber, d.email " +
		"FROM timetable t JOIN doctors d ON t.doctorID = d.id WHERE t.officeID = $1"
	selectWorkDaysByDoctor = "SELECT * FROM timetable where doctorid = $1"
)

func NewTimetableRepository(db *sqlx.DB) *TimetableRepository {
	return &TimetableRepository{db: db}
}

func (tr *TimetableRepository) GetLocationsByDoctor(id int64) ([]models.Office, error) {
	var offices []models.Office

	rows, err := tr.db.Query(selectLocationsByDoctor, id)
	if err != nil {
		return nil, ErrDatabaseReadingError
	}
	if err := rows.Err(); err != nil {
		return nil, ErrDatabaseReadingError
	}

	for rows.Next() {
		office := new(models.Office)
		if err := rows.Scan(
			&office.ID,
			&office.Number,
			&office.Floor,
			&office.BranchID,
		); err != nil {
			return nil, ErrDatabaseReadingError
		}
		offices = append(offices, *office)
	}
	defer rows.Close()

	return offices, nil
}

func (tr *TimetableRepository) GetDoctorsByLocation(id int64) ([]models.Doctor, error) {
	var doctors []models.Doctor

	rows, err := tr.db.Query(selectDoctorsByLocation, id)
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
		); err != nil {
			return nil, ErrDatabaseReadingError
		}
		doctors = append(doctors, *doctor)
	}
	defer rows.Close()

	return doctors, nil
}

func (tr *TimetableRepository) GetWorkdaysByDoctor(id int64) ([]*models.Timetable, error) {
	var timetables []*models.Timetable

	rows, err := tr.db.Query(selectWorkDaysByDoctor, id)
	if err != nil {
		return nil, ErrDatabaseDeletingError
	}
	if err = rows.Err(); err != nil {
		return nil, ErrDatabaseReadingError
	}

	for rows.Next() {
		tt := new(models.Timetable)
		if err = rows.Scan(
			&tt.ID,
			&tt.DoctorID,
			&tt.OfficeID,
			&tt.WorkDay,
		); err != nil {
			return nil, ErrDatabaseReadingError
		}
		timetables = append(timetables, tt)
	}
	defer rows.Close()

	return timetables, nil
}
