package repositories

import (
	"database/sql"
	"errors"

	"github.com/AlekSi/pointer"
	"github.com/hamillka/ppo/backend/internal/models"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

const (
	createUser = "INSERT INTO users (login, password, role) VALUES " +
		"($1, $2, $3) RETURNING id"
	createUserDoctor = "INSERT INTO users (login, password, role, doctorID) " +
		"VALUES ($1, $2, $3, $4) RETURNING id"
	createUserPatient = "INSERT INTO users (login, password, role, patientID) " +
		"VALUES ($1, $2, $3, $4) RETURNING id"
	selectRole       = "SELECT role FROM users WHERE id = $1"
	selectByLogoPass = "SELECT * FROM users WHERE login LIKE $1 AND password LIKE $2" // #nosec G101
	selectByLogin    = "SELECT * FROM users WHERE login LIKE $1"                      // #nosec G101

)

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) CreateUser(login, password string, role int64) (int64, error) {
	var newID int64
	err := ur.db.QueryRow(createUser, login, password, role).Scan(&newID) //nolint:execinquery,lll //exec doesn't work
	if err != nil {
		return 0, ErrRecordAlreadyExists
	}

	return newID, nil
}

func (ur *UserRepository) CreateUserDoctor(
	login,
	password string,
	role,
	docID int64,
) (int64, error) {
	var newID int64
	err := ur.db.QueryRow(createUserDoctor, login, password, role, docID).Scan(&newID) //nolint:execinquery,lll //exec doesn't work
	if err != nil {
		return 0, ErrRecordAlreadyExists
	}

	return newID, nil
}

func (ur *UserRepository) CreateUserPatient(
	login,
	password string,
	role,
	patID int64,
) (int64, error) {
	var newID int64
	err := ur.db.QueryRow(createUserPatient, login, password, role, patID).Scan(&newID) //nolint:execinquery,lll //exec doesn't work
	if err != nil {
		return 0, ErrRecordAlreadyExists
	}

	return newID, nil
}

func (ur *UserRepository) CheckUserRole(id int64) (int64, error) {
	var role int64

	err := ur.db.QueryRow(selectRole, id).Scan(
		&role,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return -1, ErrRecordNotFound
		}
	}

	return role, nil
}

func (ur *UserRepository) GetUserByLoginAndPassword(login, password string) (models.User, error) {
	var user models.User

	err := ur.db.QueryRow(selectByLogoPass, login, password).Scan(
		&user.ID,
		&user.Login,
		&user.Password,
		&user.Role,
		&user.PatientID,
		&user.DoctorID,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.User{}, ErrRecordNotFound
		}
		return models.User{}, err
	}

	if user.DoctorID == nil {
		user.DoctorID = pointer.ToInt64(0)
	}
	if user.PatientID == nil {
		user.PatientID = pointer.ToInt64(0)
	}

	return user, nil
}

func (ur *UserRepository) GetUserByLogin(login string) (*models.User, error) {
	var user models.User

	err := ur.db.QueryRow(selectByLogin, login).Scan(
		&user.ID,
		&user.Login,
		&user.Password,
		&user.Role,
		&user.PatientID,
		&user.DoctorID,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrRecordNotFound
		}
		return nil, err
	}

	if user.DoctorID == nil {
		user.DoctorID = pointer.ToInt64(0)
	}
	if user.PatientID == nil {
		user.PatientID = pointer.ToInt64(0)
	}

	return &user, nil
}
