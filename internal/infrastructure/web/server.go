package web

import (
	"net/http"

	"github.com/pos-graduacao-go-expert-desafio1/internal/infrastructure/handlers"
)

func StartServer(handler *handlers.CotationHandler) {
	http.HandleFunc("/cotacao", handler.CreateCotation)
	http.ListenAndServe(":8080", nil)
}
