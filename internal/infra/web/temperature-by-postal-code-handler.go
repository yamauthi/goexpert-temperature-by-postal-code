package web

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/yamauthi/goexpert-temperature-by-postalcode/internal/application/service"
)

type TemperatureByPostalCodeHandler struct {
	Service *service.TemperatureByPostalCodeService
}

func NewTemperatureByPostalCodeHandler(service *service.TemperatureByPostalCodeService) *TemperatureByPostalCodeHandler {
	return &TemperatureByPostalCodeHandler{
		Service: service,
	}
}

func (h *TemperatureByPostalCodeHandler) TemperatureByPostalCode(w http.ResponseWriter, r *http.Request) {
	input := service.PostalCodeInput{
		PostalCode: chi.URLParam(r, "postalCode"),
	}

	output, err := h.Service.Execute(input)
	if err != nil {
		statusCodeError := http.StatusInternalServerError
		if errors.Is(err, service.ErrInvalidPostalCode) {
			statusCodeError = http.StatusUnprocessableEntity
		}

		if errors.Is(err, service.ErrPostalCodeNotFound) {
			statusCodeError = http.StatusNotFound
		}

		http.Error(w, err.Error(), statusCodeError)
		return
	}

	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
