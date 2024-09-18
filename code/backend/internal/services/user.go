//go:generate mockgen -package=mocks -source=$GOFILE -destination=mocks/$GOFILE
package services

import (
	goErrors "errors"

	"github.com/hamillka/ppo/backend/internal/handlers/dto"
	"github.com/hamillka/ppo/backend/internal/models"
)

type UserRepository interface {
	CreateUser(login, password string, role int64) (int64, error)
	CreateUserDoctor(login, password string, role, docID int64) (int64, error)
	CreateUserPatient(login, password string, role, patID int64) (int64, error)
	CheckUserRole(id int64) (int64, error)
	GetUserByLoginAndPassword(login, password string) (models.User, error)
	GetUserByLogin(login string) (*models.User, error)
}

type UserService struct {
	userRepo    UserRepository
	doctorRepo  DoctorRepository
	patientRepo PatientRepository
}

func NewUserService(ur UserRepository, dr DoctorRepository, pr PatientRepository) *UserService {
	return &UserService{
		userRepo:    ur,
		doctorRepo:  dr,
		patientRepo: pr,
	}
}

func (us *UserService) Login() error {
	// login logic

	return nil
}

func (us *UserService) Register(
	fio,
	phoneNumber,
	email,
	insurance,
	specialization,
	login,
	password string,
	role int64,
) (int64, error) {
	var (
		id, docID, patID int64
		err              error
	)

	user, err := us.userRepo.GetUserByLogin(login)
	if user != nil && user.Login == login {
		return 0, goErrors.New("username is busy")
	}

	switch role {
	case dto.DOCTOR:
		docID, err = us.doctorRepo.AddDoctor(fio, phoneNumber, email, specialization)
		if err != nil {
			return 0, err
		}
		id, err = us.userRepo.CreateUserDoctor(login, password, role, docID)
		if err != nil {
			return 0, err
		}
	case dto.USER:
		patID, err = us.patientRepo.AddPatient(fio, phoneNumber, email, insurance)
		if err != nil {
			return 0, err
		}
		id, err = us.userRepo.CreateUserPatient(login, password, role, patID)
		if err != nil {
			return 0, err
		}
	default:
		id, err = us.userRepo.CreateUser(login, password, role)
		if err != nil {
			return 0, err
		}
	}

	return id, nil
}

func (us *UserService) CheckUserRole(id int64) (int64, error) {
	role, err := us.userRepo.CheckUserRole(id)
	if err != nil {
		return -1, err
	}

	return role, nil
}

func (us *UserService) GetUserByLoginAndPassword(login, password string) (models.User, error) {
	user, err := us.userRepo.GetUserByLoginAndPassword(login, password)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
