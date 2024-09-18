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

type MedicalHistoryService interface {
	GetHistoryByPatient(id int64) (models.MedicalHistory, error)
	CreateMedicalHistory(chronicDiseases, allergies, bloodType, vaccination string, patientID int64) (int64, error)
	UpdateMedicalHistory(id int64, chronicDiseases, allergies, bloodType, vaccination string) (int64, error)
}

type MedicalHistoryHandler struct {
	service MedicalHistoryService
	logger  *zap.SugaredLogger
}

func NewMedicalHistoryHandler(
	s MedicalHistoryService,
	logger *zap.SugaredLogger,
) *MedicalHistoryHandler {
	return &MedicalHistoryHandler{
		service: s,
		logger:  logger,
	}
}

func (mhh *MedicalHistoryHandler) GetHistoryByPatient(
	w http.ResponseWriter,
	r *http.Request,
) {
	ctx := r.Context()
	key := middlewares.Key("props")
	claims := ctx.Value(key).(jwt.MapClaims) //nolint: errcheck // always props
	role := int64(claims["role"].(float64))
	if role != dto.DOCTOR && role != dto.USER {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	param, ok := mux.Vars(r)["id"]
	patientID, err := strToInt64(param)
	if !ok || err != nil {
		mhh.logger.Errorf("doctor handler: %s", err)
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

	medicalHistory, err := mhh.service.GetHistoryByPatient(patientID)
	if err != nil {
		mhh.logger.Errorf("medical history handler: get history by patient id service method: %s", err)
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

	getMedicalHistoryResponseDto := dto.ConvertToMedicalHistoryDto(medicalHistory)

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(getMedicalHistoryResponseDto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (mhh *MedicalHistoryHandler) CreateMedicalHistory(
	w http.ResponseWriter,
	r *http.Request,
) {
	var history dto.CreateOrEditMedicalHistoryRequestDto

	w.Header().Add("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&history)
	if err != nil {
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

	id, err := mhh.service.CreateMedicalHistory(
		history.ChronicDiseases,
		history.Allergies,
		history.BloodType,
		history.Vaccination,
		history.PatientID,
	)
	if err != nil {
		mhh.logger.Errorf("medical history handler: create history service method: %s", err)
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

	createOrEditMedicalHistoryResponseDto := dto.CreateOrEditMedicalHistoryResponseDto{
		ID: id,
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(createOrEditMedicalHistoryResponseDto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

//nolint:funlen // it's ok
func (mhh *MedicalHistoryHandler) UpdateMedicalHistory(
	w http.ResponseWriter,
	r *http.Request,
) {
	ctx := r.Context()
	key := middlewares.Key("props")
	claims := ctx.Value(key).(jwt.MapClaims) //nolint: errcheck // always props
	role := int64(claims["role"].(float64))
	if role != dto.DOCTOR {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	var history dto.CreateOrEditMedicalHistoryRequestDto
	param, ok := mux.Vars(r)["id"]
	patientID, err := strToInt64(param)
	if !ok || err != nil {
		mhh.logger.Errorf("medical history handler: %s", err)
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
	err = json.NewDecoder(r.Body).Decode(&history)
	if err != nil {
		mhh.logger.Errorf("appointment handler: decode json %s", err)
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

	id, err := mhh.service.UpdateMedicalHistory(
		patientID,
		history.ChronicDiseases,
		history.Allergies,
		history.BloodType,
		history.Vaccination,
	)
	if err != nil {
		mhh.logger.Errorf("medical history handler: update medical history service method: %s", err)
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

	createOrUpdateMedicalHistoryResponseDto := dto.CreateOrEditMedicalHistoryResponseDto{
		ID: id,
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(createOrUpdateMedicalHistoryResponseDto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
