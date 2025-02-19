package http

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"weatherapp/internal/domain"

	"github.com/go-playground/validator/v10"
)

type (
	getCityCommand interface {
		GetCity(ctx context.Context, id int64) (*domain.City, error)
	}

	GetCityHandler struct {
		name           string
		getCityCommand getCityCommand
	}

	getCityRequest struct {
		Id int64 `validate:"required"`
	}
)

func NewCityGetHandler(command getCityCommand, name string) *GetCityHandler {
	return &GetCityHandler{
		name:           name,
		getCityCommand: command,
	}
}

func (h *GetCityHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		ctx         = r.Context()
		err         error
		requestData *getCityRequest
	)

	if requestData, err = h.getRequestData(r); err != nil {
		GetErrorResponse(w, h.name, err, http.StatusBadRequest)
		return
	}

	responseRawBody, err := h.getCityCommand.GetCity(ctx, requestData.Id)

	if err != nil {
		if errors.Is(err, domain.ErrCityNotFound) {
			GetErrorResponse(w, h.name, err, http.StatusBadRequest)
			return
		}

		GetErrorResponse(w, h.name, fmt.Errorf("command handler failed: %w", err), http.StatusInternalServerError)
		return
	}

	responseBody, err := json.Marshal(responseRawBody)
	if err != nil {
		GetErrorResponse(w, h.name, fmt.Errorf("json marshalling failed: %w", err), http.StatusInternalServerError)
		return
	}

	GetSuccessResponseWithJSON(w, responseBody)
}

func (h *GetCityHandler) getRequestData(r *http.Request) (requestData *getCityRequest, err error) {
	requestData = &getCityRequest{}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		return
	}

	requestData.Id = int64(id)

	return
}

func (h *GetCityHandler) validateRequestData(requestData *getCityRequest) error {
	return validator.New().Struct(requestData)
}
