package repositories

import (
	"database/sql"
	"errors"

	"github.com/hamillka/ppo/backend/internal/models"
	"github.com/jmoiron/sqlx"
)

type OfficeRepository struct {
	db *sqlx.DB
}

const (
	createOffice  = "INSERT INTO offices (number, floor, branchID) VALUES ($1, $2, $3) RETURNING id"
	selectOffice  = "SELECT * FROM offices WHERE id = $1"
	updateOffice  = "UPDATE offices SET number = $1, floor = $2, branchID = $3 WHERE id = $4"
	selectOffices = "SELECT * FROM offices"
)

func NewOfficeRepository(db *sqlx.DB) *OfficeRepository {
	return &OfficeRepository{db: db}
}

func (or *OfficeRepository) EditOffice(id, number, floor, branchID int64) (int64, error) {
	office := new(models.Office)

	err := or.db.QueryRow(selectOffice, id).Scan(
		&office.ID,
		&office.Number,
		&office.Floor,
		&office.BranchID,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, ErrRecordNotFound
		}
	}

	if number != 0 {
		office.Number = number
	}
	if floor != 0 {
		office.Floor = floor
	}
	if branchID != 0 {
		office.BranchID = branchID
	}

	_, err = or.db.Exec(updateOffice, office.Number, office.Floor, office.BranchID, office.ID)
	if err != nil {
		return 0, ErrDatabaseUpdatingError
	}
	return id, nil
}

func (or *OfficeRepository) AddOffice(number, floor, branchID int64) (int64, error) {
	var newID int64
	err := or.db.QueryRow(createOffice, number, floor, branchID).Scan(&newID) //nolint:execinquery,lll //exec doesn't work
	if err != nil {
		return 0, ErrRecordAlreadyExists
	}

	return newID, nil
}

func (or *OfficeRepository) GetAllOffices() ([]models.Office, error) {
	var offices []models.Office

	rows, err := or.db.Query(selectOffices)
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

func (or *OfficeRepository) GetOfficeByID(id int64) (models.Office, error) {
	var office models.Office

	err := or.db.QueryRow(selectOffice, id).Scan(
		&office.ID,
		&office.Number,
		&office.Floor,
		&office.BranchID,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Office{}, ErrRecordNotFound
		}
	}

	return office, nil
}
