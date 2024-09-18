package repositories

import (
	"database/sql"
	"errors"

	"github.com/hamillka/ppo/backend/internal/models"
	"github.com/jmoiron/sqlx"
)

type BranchRepository struct {
	db *sqlx.DB
}

const (
	createBranch = "INSERT INTO branches (name, address, phoneNumber) VALUES " +
		"($1, $2, $3) RETURNING id"
	selectBranch = "SELECT * FROM branches WHERE id = $1"
	updateBranch = "UPDATE branches " +
		"SET name = $1, address = $2, phoneNumber = $3 WHERE id = $4"
	selectBranches = "SELECT * FROM branches"
)

func NewBranchRepository(db *sqlx.DB) *BranchRepository {
	return &BranchRepository{db: db}
}

func (br *BranchRepository) EditBranch(id int64, name, address, phoneNumber string) (int64, error) {
	branch := new(models.Branch)

	err := br.db.QueryRow(selectBranch, id).Scan(
		&branch.ID,
		&branch.Name,
		&branch.Address,
		&branch.PhoneNumber,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, ErrRecordNotFound
		}
		return 0, err
	}

	if name != "" {
		branch.Name = name
	}
	if address != "" {
		branch.Address = address
	}
	if phoneNumber != "" {
		branch.PhoneNumber = phoneNumber
	}

	_, err = br.db.Exec(updateBranch,
		branch.Name,
		branch.Address,
		branch.PhoneNumber,
		branch.ID,
	)
	if err != nil {
		return 0, ErrDatabaseUpdatingError
	}
	return id, nil
}

func (br *BranchRepository) AddBranch(name, address, phoneNumber string) (int64, error) {
	var newID int64
	err := br.db.QueryRow(createBranch, name, address, phoneNumber).Scan(&newID) //nolint:execinquery,lll //exec doesn't work
	if err != nil {
		return 0, ErrRecordAlreadyExists
	}

	return newID, nil
}

func (br *BranchRepository) GetAllBranches() ([]*models.Branch, error) {
	var branches []*models.Branch

	rows, err := br.db.Query(selectBranches)
	if err != nil {
		return nil, ErrDatabaseReadingError
	}
	if err := rows.Err(); err != nil {
		return nil, ErrDatabaseReadingError
	}

	for rows.Next() {
		branch := new(models.Branch)
		if err := rows.Scan(
			&branch.ID,
			&branch.Name,
			&branch.Address,
			&branch.PhoneNumber,
		); err != nil {
			return nil, ErrDatabaseReadingError
		}
		branches = append(branches, branch)
	}
	defer rows.Close()

	return branches, nil
}

func (br *BranchRepository) GetBranchByID(id int64) (models.Branch, error) {
	var branch models.Branch

	err := br.db.QueryRow(selectBranch, id).Scan(
		&branch.ID,
		&branch.Name,
		&branch.Address,
		&branch.PhoneNumber,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Branch{}, ErrRecordNotFound
		}
	}

	return branch, nil
}
