package container

import (
	"database/sql"

	"github.com/pos-graduacao-go-expert-desafio1/internal/application"
	"github.com/pos-graduacao-go-expert-desafio1/internal/infrastructure/repository"
)

type Container struct {
	CotationService application.CotationService
}

func New(db *sql.DB) *Container {
	return &Container{
		CotationService: *application.NewCotationService(
			repository.NewContationRepository(db),
		),
	}
}
