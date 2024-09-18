package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
	"github.com/hamillka/ppo/backend/internal/handlers/dto"
	"github.com/hamillka/ppo/backend/internal/handlers/middlewares"
	"github.com/hamillka/ppo/backend/internal/models"
	"github.com/hamillka/ppo/backend/internal/repositories"
	"go.uber.org/zap"
)

var ErrValidate = errors.New("validation error")

type AppointmentService interface {
	CreateAppointment(patientID, doctorID int64, dateTime time.Time) (int64, error)
	CancelAppointment(id int64) error
	GetAppointmentsByPatient(id int64) ([]*models.Appointment, error)
	GetAppointmentsByDoctor(id int64) ([]*models.Appointment, error)
	EditAppointment(id, doctorID, patientID int64, dateTime time.Time) (int64, error)
	GetAppointmentByID(id int64) (*models.Appointment, error)
}

type AppointmentHandler struct {
	service AppointmentService
	logger  *zap.SugaredLogger
}

func NewAppointmentHandler(s AppointmentService, logger *zap.SugaredLogger) *AppointmentHandler {
	return &AppointmentHandler{
		service: s,
		logger:  logger,
	}
}

func (ah *AppointmentHandler) CreateAppointment(w http.ResponseWriter, r *http.Request) {
	var appointment dto.CreateOrEditAppointmentRequestDto

	w.Header().Add("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&appointment)
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

	id, err := ah.service.CreateAppointment(
		appointment.PatientID,
		appointment.DoctorID,
		appointment.DateTime,
	)
	if err != nil {
		ah.logger.Errorf("appointment handler: create appointment service method: %s", err)
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

	createOrEditAppointmentResponseDto := dto.CreateOrEditAppointmentResponseDto{
		ID: id,
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(createOrEditAppointmentResponseDto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

//nolint:funlen // it's ok
func (ah *AppointmentHandler) EditAppointment(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	key := middlewares.Key("props")
	claims := ctx.Value(key).(jwt.MapClaims) //nolint: errcheck // always props
	role := int64(claims["role"].(float64))
	if role == dto.DOCTOR {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	var appointment dto.CreateOrEditAppointmentRequestDto
	param, ok := mux.Vars(r)["id"]
	appointmentID, err := strToInt64(param)
	if !ok || err != nil {
		ah.logger.Errorf("appointment handler: %s", err)
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
	err = json.NewDecoder(r.Body).Decode(&appointment)
	if err != nil {
		ah.logger.Errorf("appointment handler: decode json %s", err)
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

	id, err := ah.service.EditAppointment(
		appointmentID,
		appointment.DoctorID,
		appointment.PatientID,
		appointment.DateTime,
	)
	if err != nil {
		ah.logger.Errorf("appointment handler: edit appointment service method: %s", err)
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

	createOrUpdateAppointmentResponseDto := dto.CreateOrEditAppointmentResponseDto{
		ID: id,
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(createOrUpdateAppointmentResponseDto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (ah *AppointmentHandler) CancelAppointment(w http.ResponseWriter, r *http.Request) {
	param, ok := mux.Vars(r)["id"]
	appointmentID, err := strToInt64(param)
	if !ok || err != nil {
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

	err = ah.service.CancelAppointment(appointmentID)
	if err != nil {
		ah.logger.Errorf("appointment handler: cancel appointment service method: %s", err)
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

	w.WriteHeader(http.StatusNoContent)
}

func (ah *AppointmentHandler) GetAppointmentByID(w http.ResponseWriter, r *http.Request) {
	param, ok := mux.Vars(r)["id"]
	appointmentID, err := strToInt64(param)
	if !ok || err != nil {
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

	appointment, err := ah.service.GetAppointmentByID(appointmentID)
	if err != nil {
		ah.logger.Errorf("appointment handler: get appointment by id service method: %s", err)
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

	getAppointmentResponseDto := dto.ConvertToDto(appointment)

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(getAppointmentResponseDto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (ah *AppointmentHandler) GetAppointments( //nolint: cyclop // no other ways
	w http.ResponseWriter,
	r *http.Request,
) {
	doctorID, errDoc := getQueryParam(r, "doctor_id")
	patientID, errPat := getQueryParam(r, "patient_id")
	getAppointmentsResponseDto := make([]*dto.GetAppointmentResponseDto, 0)
	var (
		appointments []*models.Appointment
		err          error
	)

	switch {
	case errDoc == nil && doctorID != 0:
		appointments, err = ah.service.GetAppointmentsByDoctor(doctorID)
		if err != nil {
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
	case errPat == nil && patientID != 0:
		appointments, err = ah.service.GetAppointmentsByPatient(patientID)
		if err != nil {
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
	default:
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

	for _, appointment := range appointments {
		getAppointmentsResponseDto = append(getAppointmentsResponseDto, dto.ConvertToDto(appointment))
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(getAppointmentsResponseDto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func strToInt64(val string) (int64, error) {
	result, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return 0, ErrValidate
	}

	return result, nil
}

func getQueryParam(r *http.Request, key string) (int64, error) {
	var val string
	if val = r.URL.Query().Get(key); val == "" {
		return 0, nil
	}

	result, err := strToInt64(val)
	if err != nil {
		return 0, err
	}

	return result, nil
}
