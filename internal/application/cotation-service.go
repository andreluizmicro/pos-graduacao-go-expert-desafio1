package application

import (
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

func (service *CotationService) Create() (*CotationOutputDTO, error) {
	cotation, err := service.repository.Create()
	if err != nil {
		return &CotationOutputDTO{}, nil
	}

	return &CotationOutputDTO{
		Bid: cotation.Bid,
	}, nil
}
