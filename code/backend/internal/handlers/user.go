package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/hamillka/ppo/backend/internal/handlers/dto"
	"github.com/hamillka/ppo/backend/internal/handlers/middlewares"
	"github.com/hamillka/ppo/backend/internal/models"
	"go.uber.org/zap"
)

type UserService interface {
	Login() error
	Register(fio, phoneNumber, email, insurance, specialization, login, password string, role int64) (int64, error)
	CheckUserRole(id int64) (int64, error)
	GetUserByLoginAndPassword(login, password string) (models.User, error)
}

type UserHandler struct {
	service UserService
	logger  *zap.SugaredLogger
}

func NewUserHandler(s UserService, logger *zap.SugaredLogger) *UserHandler {
	return &UserHandler{
		service: s,
		logger:  logger,
	}
}

func createToken(role int64) (string, error) {
	payload := jwt.MapClaims{
		"role": role,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	t, err := token.SignedString(middlewares.Secret)
	if err != nil {
		return "", err
	}

	return t, nil
}

func (uh *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var userLoginDto dto.UserLoginRequestDto

	w.Header().Add("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&userLoginDto)
	if err != nil {
		uh.logger.Errorf("user handler: json decode %s", err)
		w.WriteHeader(http.StatusBadRequest)
		errorDto := &dto.ErrorDto{
			Error: "Некорректные данные",
		}
		err = json.NewEncoder(w).Encode(errorDto)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	user, err := uh.service.GetUserByLoginAndPassword(userLoginDto.Login, userLoginDto.Password)
	if err != nil {
		uh.logger.Errorf("user handler: get user by login and password service method: %s", err)
		w.WriteHeader(http.StatusUnauthorized)
		errorDto := &dto.ErrorDto{
			Error: "Неверное имя пользователя или пароль",
		}
		err = json.NewEncoder(w).Encode(errorDto)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	t, err := createToken(user.Role)
	if err != nil {
		uh.logger.Errorf("user handler: create token method: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		errorDto := &dto.ErrorDto{
			Error: "Возникла внутренняя ошибка при авторизации",
		}
		err = json.NewEncoder(w).Encode(errorDto)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	userReqDto := dto.ConvertToUserDto(&user)

	userResponseDto := dto.UserLoginResponseDto{
		JWTToken: t,
		User:     *userReqDto,
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(userResponseDto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (uh *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var userDto dto.UserRegisterRequestDto

	w.Header().Add("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&userDto)
	if err != nil {
		uh.logger.Errorf("user handler: json decode %s", err)

		w.WriteHeader(http.StatusBadRequest)
		errorDto := &dto.ErrorDto{
			Error: "Некорректные данные",
		}
		err = json.NewEncoder(w).Encode(errorDto)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	_, err = uh.service.Register(
		userDto.FIO,
		userDto.PhoneNumber,
		userDto.Email,
		userDto.Insurance,
		userDto.Specialization,
		userDto.Login,
		userDto.Password,
		userDto.Role,
	)
	if err != nil {
		uh.logger.Errorf("user handler: register service method: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		errorDto := &dto.ErrorDto{
			Error: "Внутренняя ошибка сервера",
		}
		err = json.NewEncoder(w).Encode(errorDto)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	w.WriteHeader(http.StatusOK)
}
