package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hamillka/ppo/backend/internal/handlers/dto"
	"github.com/hamillka/ppo/backend/internal/models"
	"go.uber.org/zap"
)

type TimetableService interface {
	GetLocationsByDoctor(id int64) ([]models.Office, error)
	GetDoctorsByLocation(id int64) ([]models.Doctor, error)
	GetWorkdaysByDoctor(id int64) ([]*models.Timetable, error)
}

type TimetableHandler struct {
	service TimetableService
	logger  *zap.SugaredLogger
}

func NewTimetableHandler(s TimetableService, logger *zap.SugaredLogger) *TimetableHandler {
	return &TimetableHandler{
		service: s,
		logger:  logger,
	}
}

func (th *TimetableHandler) GetLocationsByDoctor(w http.ResponseWriter, r *http.Request) {
	param, ok := mux.Vars(r)["id"]
	doctorID, err := strToInt64(param)
	if !ok || err != nil {
		th.logger.Errorf("timetable handler: %s", err)
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

	offices, err := th.service.GetLocationsByDoctor(doctorID)
	if err != nil {
		th.logger.Errorf("timetable handler: get location by doctor service method: %s", err)
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

	getLocationsByDoctor := make([]*dto.GetOfficeResponseDto, 0)

	for idx := range offices {
		getLocationsByDoctor = append(getLocationsByDoctor, dto.ConvertToOfficeDto(&offices[idx]))
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(getLocationsByDoctor)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (th *TimetableHandler) GetDoctorsByLocation(w http.ResponseWriter, r *http.Request) {
	param, ok := mux.Vars(r)["id"]
	officeID, err := strToInt64(param)
	if !ok || err != nil {
		th.logger.Errorf("timetable handler: %s", err)
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

	doctors, err := th.service.GetDoctorsByLocation(officeID)
	if err != nil {
		th.logger.Errorf("timetable handler: get doctors by location service method: %s", err)
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

	getDoctorsByLocation := make([]*dto.GetDoctorResponseDto, 0)

	for idx := range doctors {
		getDoctorsByLocation = append(getDoctorsByLocation, dto.ConvertToDoctorDto(&doctors[idx]))
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(getDoctorsByLocation)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (th *TimetableHandler) GetTimetableByDoctor(w http.ResponseWriter, r *http.Request) {
	param, ok := mux.Vars(r)["id"]
	docID, err := strToInt64(param)
	if !ok || err != nil {
		th.logger.Errorf("timetable handler: %s", err)
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

	timetable, err := th.service.GetWorkdaysByDoctor(docID)
	if err != nil {
		th.logger.Errorf("timetable handler: get timetable by doctor service method: %s", err)
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

	getTimetableByDoctor := make([]*dto.GetTimetableResponseDto, 0)

	for idx := range timetable {
		getTimetableByDoctor = append(getTimetableByDoctor, dto.ConvertToTimetableDto(timetable[idx]))
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(getTimetableByDoctor)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
