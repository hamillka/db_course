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

type BranchService interface {
	EditBranch(id int64, name, address, phoneNumber string) (int64, error)
	AddBranch(name, address, phoneNumber string) (int64, error)
	GetAllBranches() ([]*models.Branch, error)
	GetBranchByID(id int64) (models.Branch, error)
}

type BranchHandler struct {
	service BranchService
	logger  *zap.SugaredLogger
}

func NewBranchHandler(s BranchService, logger *zap.SugaredLogger) *BranchHandler {
	return &BranchHandler{
		service: s,
		logger:  logger,
	}
}

func (bh *BranchHandler) AddBranch(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	key := middlewares.Key("props")
	claims := ctx.Value(key).(jwt.MapClaims) //nolint: errcheck // always props
	role := int64(claims["role"].(float64))
	if role != dto.ADMIN {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	var branch dto.CreateOrEditBranchRequestDto

	w.Header().Add("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&branch)
	if err != nil {
		bh.logger.Errorf("branch handler: decode json: %s", err)
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

	id, err := bh.service.AddBranch(
		branch.Name,
		branch.Address,
		branch.PhoneNumber,
	)
	if err != nil {
		bh.logger.Errorf("branch handler: add branch service method: %s", err)
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

	createOrEditBranchResponseDto := dto.CreateOrEditBranchResponseDto{
		ID: id,
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(createOrEditBranchResponseDto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

//nolint:funlen // it's ok
func (bh *BranchHandler) EditBranch(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	key := middlewares.Key("props")
	claims := ctx.Value(key).(jwt.MapClaims) //nolint: errcheck // always props
	role := int64(claims["role"].(float64))
	if role != dto.ADMIN {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	var branch dto.CreateOrEditBranchRequestDto

	param, ok := mux.Vars(r)["id"]
	branchID, err := strToInt64(param)
	if !ok || err != nil {
		bh.logger.Errorf("branch handler: %s", err)
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
	err = json.NewDecoder(r.Body).Decode(&branch)
	if err != nil {
		bh.logger.Errorf("branch handler: json decode %s", err)
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

	id, err := bh.service.EditBranch(
		branchID,
		branch.Name,
		branch.Address,
		branch.PhoneNumber,
	)
	if err != nil {
		bh.logger.Errorf("branch handler: edit branch service method: %s", err)
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

	createOrUpdateBranchResponseDto := dto.CreateOrEditBranchResponseDto{
		ID: id,
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(createOrUpdateBranchResponseDto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (bh *BranchHandler) GetAllBranches(w http.ResponseWriter, r *http.Request) {
	allBranches, err := bh.service.GetAllBranches()
	if err != nil {
		bh.logger.Errorf("branch handler: get all branches service method: %s", err)
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

	getAllBranchesResponseDto := make([]*dto.GetBranchResponseDto, 0)

	for _, branch := range allBranches {
		getAllBranchesResponseDto = append(getAllBranchesResponseDto, dto.ConvertToBranchDto(branch))
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(getAllBranchesResponseDto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (bh *BranchHandler) GetBranchByID(w http.ResponseWriter, r *http.Request) {
	param, ok := mux.Vars(r)["id"]
	branchID, err := strToInt64(param)
	if !ok || err != nil {
		bh.logger.Errorf("branch handler: %s", err)
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

	branch, err := bh.service.GetBranchByID(branchID)
	if err != nil {
		bh.logger.Errorf("branch handler: get branch by id service method: %s", err)
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

	getBranchResponseDto := dto.ConvertToBranchDto(&branch)

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(getBranchResponseDto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
