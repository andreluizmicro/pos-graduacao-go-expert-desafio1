package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/pos-graduacao-go-expert-desafio1/internal/application"
)

type CotationHandler struct {
	service *application.CotationService
}

func NewCotationHandler(service *application.CotationService) *CotationHandler {
	return &CotationHandler{
		service: service,
	}
}

func (handler *CotationHandler) CreateCotation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cotation, err := handler.service.Create()

	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cotation)
}
