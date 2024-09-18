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

type PatientService interface {
	EditPatient(id int64, fio, phoneNumber, email, insurance string) (int64, error)
	AddPatient(fio, phoneNumber, email, insurance string) (int64, error)
	GetAllPatients() ([]models.Patient, error)
	GetPatientByID(id int64) (models.Patient, error)
}

type PatientHandler struct {
	service PatientService
	logger  *zap.SugaredLogger
}

func NewPatientHandler(s PatientService, logger *zap.SugaredLogger) *PatientHandler {
	return &PatientHandler{
		service: s,
		logger:  logger,
	}
}

func (ph *PatientHandler) AddPatient(w http.ResponseWriter, r *http.Request) {
	var patient dto.CreateOrEditPatientRequestDto

	w.Header().Add("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&patient)
	if err != nil {
		ph.logger.Errorf("patient handler: json decode %s", err)
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

	id, err := ph.service.AddPatient(
		patient.Fio,
		patient.PhoneNumber,
		patient.Email,
		patient.Insurance,
	)
	if err != nil {
		ph.logger.Errorf("patient handler: add patient service method: %s", err)
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

	createOrEditPatientResponseDto := dto.CreateOrEditPatientResponseDto{
		ID: id,
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(createOrEditPatientResponseDto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

//nolint:funlen // it's ok
func (ph *PatientHandler) EditPatient(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	key := middlewares.Key("props")
	claims := ctx.Value(key).(jwt.MapClaims) //nolint: errcheck // always props
	role := int64(claims["role"].(float64))
	if role == dto.DOCTOR {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	var patient dto.CreateOrEditPatientRequestDto

	param, ok := mux.Vars(r)["id"]
	patientID, err := strToInt64(param)
	if !ok || err != nil {
		ph.logger.Errorf("patient handler: %s", err)
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
	err = json.NewDecoder(r.Body).Decode(&patient)
	if err != nil {
		ph.logger.Errorf("patient handler: json decode %s", err)
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

	id, err := ph.service.EditPatient(
		patientID,
		patient.Fio,
		patient.PhoneNumber,
		patient.Email,
		patient.Insurance,
	)
	if err != nil {
		ph.logger.Errorf("patient handler: edit patient service method: %s", err)
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

	createOrUpdatePatientResponseDto := dto.CreateOrEditPatientResponseDto{
		ID: id,
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(createOrUpdatePatientResponseDto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (ph *PatientHandler) GetAllPatients(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	key := middlewares.Key("props")
	claims := ctx.Value(key).(jwt.MapClaims) //nolint: errcheck // always props
	role := int64(claims["role"].(float64))
	if role != dto.ADMIN {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	allPatients, err := ph.service.GetAllPatients()
	if err != nil {
		ph.logger.Errorf("patient handler: get all patients service method: %s", err)
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

	getAllPatientsResponseDto := make([]*dto.GetPatientResponseDto, 0)

	for idx := range allPatients {
		getAllPatientsResponseDto = append(
			getAllPatientsResponseDto,
			dto.ConvertToPatientDto(&allPatients[idx]),
		)
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(getAllPatientsResponseDto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (ph *PatientHandler) GetPatientByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	key := middlewares.Key("props")
	claims := ctx.Value(key).(jwt.MapClaims) //nolint: errcheck // always props
	role := int64(claims["role"].(float64))
	if role == dto.USER {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	param, ok := mux.Vars(r)["id"]
	patientID, err := strToInt64(param)
	if !ok || err != nil {
		ph.logger.Errorf("patient handler: %s", err)
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

	patient, err := ph.service.GetPatientByID(patientID)
	if err != nil {
		ph.logger.Errorf("patient handler: get patient by id service method: %s", err)
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

	getPatientResponseDto := dto.ConvertToPatientDto(&patient)

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(getPatientResponseDto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
