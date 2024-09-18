package dto

import "github.com/hamillka/ppo/backend/internal/models"

const (
	ADMIN = iota
	USER
	DOCTOR
)

type UserDto struct {
	PatientID *int64 `json:"patientId"`
	DoctorID  *int64 `json:"doctorId"`
	Login     string `json:"login"`
	Password  string `json:"password"`
	Role      int64  `json:"role"`
	ID        int64  `json:"id"`
}

type UserLoginRequestDto struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UserLoginResponseDto struct {
	JWTToken string  `json:"jwtToken"`
	User     UserDto `json:"user"`
}

type UserRegisterRequestDto struct {
	FIO            string `json:"fio"`
	PhoneNumber    string `json:"phoneNumber"`
	Email          string `json:"email"`
	Insurance      string `json:"insurance"`
	Specialization string `json:"specialization"`
	Login          string `json:"login"`
	Password       string `json:"password"`
	Role           int64  `json:"role"`
}

func ConvertToUserDto(user *models.User) *UserDto {
	return &UserDto{
		PatientID: user.PatientID,
		DoctorID:  user.DoctorID,
		Login:     user.Login,
		Password:  user.Password,
		Role:      user.Role,
		ID:        user.ID,
	}
}
