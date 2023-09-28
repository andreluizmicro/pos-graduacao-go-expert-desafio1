package main

import (
	"github.com/pos-graduacao-go-expert-desafio1/client/internal/application"
	"github.com/pos-graduacao-go-expert-desafio1/client/internal/infrastructure/repository"
)

func main() {

	repository := repository.NewCotationRepository()
	service := application.NewCotationService(repository)

	err := service.Create()
	if err != nil {
		panic(err)
	}
}
