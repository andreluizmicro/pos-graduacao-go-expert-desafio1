package application

import (
	"github.com/pos-graduacao-go-expert-desafio1/internal/domain/entity"
	"github.com/pos-graduacao-go-expert-desafio1/internal/infrastructure/repository"
)

type CotationService struct {
	repository *repository.CotationRepository
}

func NewCotationService(repository *repository.CotationRepository) *CotationService {
	return &CotationService{
		repository: repository,
	}
}

func (service *CotationService) Create() (*entity.Cotation, error) {
	return service.repository.Create()
}
