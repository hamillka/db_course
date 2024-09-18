package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
	"github.com/hamillka/ppo/backend/internal/handlers/dto"
	"github.com/hamillka/ppo/backend/internal/handlers/middlewares"
	"github.com/hamillka/ppo/backend/internal/models"
	"github.com/hamillka/ppo/backend/internal/repositories"
	"go.uber.org/zap"
)

type DoctorService interface {
	EditDoctor(id int64, fio, phoneNumber, email, specialization string) (int64, error)
	AddDoctor(fio, phoneNumber, email, specialization string) (int64, error)
	GetAllDoctors() ([]models.Doctor, error)
	GetDoctorByID(id int64) (models.Doctor, error)
}

type DoctorHandler struct {
	service DoctorService
	logger  *zap.SugaredLogger
}

func NewDoctorHandler(s DoctorService, logger *zap.SugaredLogger) *DoctorHandler {
	return &DoctorHandler{
		service: s,
		logger:  logger,
	}
}

func (dh *DoctorHandler) AddDoctor(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	key := middlewares.Key("props")
	claims := ctx.Value(key).(jwt.MapClaims) //nolint: errcheck // always props
	role := int64(claims["role"].(float64))
	if role != dto.ADMIN {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	var doctor dto.CreateOrEditDoctorRequestDto

	w.Header().Add("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&doctor)
	if err != nil {
		dh.logger.Errorf("doctor handler: json decode %s", err)
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

	id, err := dh.service.AddDoctor(
		doctor.Fio,
		doctor.PhoneNumber,
		doctor.Email,
		doctor.Specialization,
	)
	if err != nil {
		dh.logger.Errorf("doctor handler: add doctor service method: %s", err)
		var errorDto *dto.ErrorDto
		w.WriteHeader(http.StatusInternalServerError)
		errorDto = &dto.ErrorDto{
			Error: "Внутренняя ошибка сервера",
		}
		err = json.NewEncoder(w).Encode(errorDto)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	createOrEditDoctorResponseDto := dto.CreateOrEditDoctorResponseDto{
		ID: id,
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(createOrEditDoctorResponseDto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

//nolint:funlen // it's ok
func (dh *DoctorHandler) EditDoctor(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	key := middlewares.Key("props")
	claims := ctx.Value(key).(jwt.MapClaims) //nolint: errcheck // always props
	role := int64(claims["role"].(float64))
	if role == dto.USER {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	var doctor dto.CreateOrEditDoctorRequestDto

	param, ok := mux.Vars(r)["id"]
	doctorID, err := strToInt64(param)
	if !ok || err != nil {
		dh.logger.Errorf("doctor handler: %s", err)
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

	w.Header().Add("Content-Type", "application/json")
	err = json.NewDecoder(r.Body).Decode(&doctor)
	if err != nil {
		dh.logger.Errorf("doctor handler: json decode %s", err)

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

	id, err := dh.service.EditDoctor(
		doctorID,
		doctor.Fio,
		doctor.PhoneNumber,
		doctor.Email,
		doctor.Specialization,
	)
	if err != nil {
		dh.logger.Errorf("doctor handler: edit doctor service method: %s", err)
		var errorDto *dto.ErrorDto
		w.WriteHeader(http.StatusInternalServerError)
		errorDto = &dto.ErrorDto{
			Error: "Внутренняя ошибка сервера",
		}
		err = json.NewEncoder(w).Encode(errorDto)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	createOrUpdateDoctorResponseDto := dto.CreateOrEditDoctorResponseDto{
		ID: id,
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(createOrUpdateDoctorResponseDto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (dh *DoctorHandler) GetAllDoctors(w http.ResponseWriter, r *http.Request) {
	allDoctors, err := dh.service.GetAllDoctors()
	if err != nil {
		dh.logger.Errorf("doctor handler: get all doctors service method: %s", err)
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

	getAllDoctorsResponseDto := make([]*dto.GetDoctorResponseDto, 0)

	for idx := range allDoctors {
		getAllDoctorsResponseDto = append(
			getAllDoctorsResponseDto,
			dto.ConvertToDoctorDto(&allDoctors[idx]),
		)
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(getAllDoctorsResponseDto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (dh *DoctorHandler) GetDoctorByID(w http.ResponseWriter, r *http.Request) {
	param, ok := mux.Vars(r)["id"]
	doctorID, err := strToInt64(param)
	if !ok || err != nil {
		dh.logger.Errorf("doctor handler: %s", err)
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

	doctor, err := dh.service.GetDoctorByID(doctorID)
	if err != nil {
		dh.logger.Errorf("doctor handler: get doctor by id service method: %s", err)
		var errorDto *dto.ErrorDto
		if errors.Is(err, repositories.ErrRecordNotFound) {
			w.WriteHeader(http.StatusNotFound)
			errorDto = &dto.ErrorDto{
				Error: "Запись не найдена",
			}
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			errorDto = &dto.ErrorDto{
				Error: "Внутренняя ошибка сервера",
			}
		}
		err = json.NewEncoder(w).Encode(errorDto)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	getDoctorResponseDto := dto.ConvertToDoctorDto(&doctor)

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(getDoctorResponseDto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
