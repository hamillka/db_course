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

type OfficeService interface {
	EditOffice(id, number, floor, branchID int64) (int64, error)
	AddOffice(number, floor, branchID int64) (int64, error)
	GetAllOffices() ([]models.Office, error)
	GetOfficeByID(id int64) (models.Office, error)
}

type OfficeHandler struct {
	service OfficeService
	logger  *zap.SugaredLogger
}

func NewOfficeHandler(s OfficeService, logger *zap.SugaredLogger) *OfficeHandler {
	return &OfficeHandler{
		service: s,
		logger:  logger,
	}
}

func (oh *OfficeHandler) AddOffice(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	key := middlewares.Key("props")
	claims := ctx.Value(key).(jwt.MapClaims) //nolint: errcheck // always props
	role := int64(claims["role"].(float64))
	if role != dto.ADMIN {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	var office dto.CreateOrEditOfficeRequestDto

	w.Header().Add("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&office)
	if err != nil {
		oh.logger.Errorf("office handler: json decode %s", err)
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

	id, err := oh.service.AddOffice(
		office.Number,
		office.Floor,
		office.BranchID,
	)
	if err != nil {
		oh.logger.Errorf("office handler: add office service method: %s", err)
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

	createOrEditOfficeResponseDto := dto.CreateOrEditOfficeResponseDto{
		ID: id,
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(createOrEditOfficeResponseDto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

//nolint:funlen // it's ok
func (oh *OfficeHandler) EditOffice(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	key := middlewares.Key("props")
	claims := ctx.Value(key).(jwt.MapClaims) //nolint: errcheck // always props
	role := int64(claims["role"].(float64))
	if role != dto.ADMIN {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	var office dto.CreateOrEditOfficeRequestDto

	param, ok := mux.Vars(r)["id"]
	officeID, err := strToInt64(param)
	if !ok || err != nil {
		oh.logger.Errorf("office handler: %s", err)
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
	err = json.NewDecoder(r.Body).Decode(&office)
	if err != nil {
		oh.logger.Errorf("office handler: json decode %s", err)
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

	id, err := oh.service.EditOffice(
		officeID,
		office.Number,
		office.Floor,
		office.BranchID,
	)
	if err != nil {
		oh.logger.Errorf("office handler: edit office service method: %s", err)
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

	createOrUpdateOfficeResponseDto := dto.CreateOrEditOfficeResponseDto{
		ID: id,
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(createOrUpdateOfficeResponseDto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (oh *OfficeHandler) GetAllOffices(w http.ResponseWriter, r *http.Request) {
	allOffices, err := oh.service.GetAllOffices()
	if err != nil {
		oh.logger.Errorf("office handler: get all offices service method: %s", err)
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

	getAllOfficesResponseDto := make([]*dto.GetOfficeResponseDto, 0)

	for idx := range allOffices {
		getAllOfficesResponseDto = append(
			getAllOfficesResponseDto,
			dto.ConvertToOfficeDto(&allOffices[idx]),
		)
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(getAllOfficesResponseDto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (oh *OfficeHandler) GetOfficeByID(w http.ResponseWriter, r *http.Request) {
	param, ok := mux.Vars(r)["id"]
	officeID, err := strToInt64(param)
	if !ok || err != nil {
		oh.logger.Errorf("office handler: %s", err)
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

	office, err := oh.service.GetOfficeByID(officeID)
	if err != nil {
		oh.logger.Errorf("office handler: get office by id service method: %s", err)
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

	getOfficeResponseDto := dto.ConvertToOfficeDto(&office)

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(getOfficeResponseDto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
