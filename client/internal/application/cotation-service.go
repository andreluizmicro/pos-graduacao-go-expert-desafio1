package application

import (
	"fmt"
	"os"
	"strconv"

	"github.com/pos-graduacao-go-expert-desafio1/client/internal/domain/entity"
	"github.com/pos-graduacao-go-expert-desafio1/client/internal/infrastructure/repository"
)

type CotationService struct {
	repository *repository.CotationRepository
}

func NewCotationService(repository *repository.CotationRepository) *CotationService {
	return &CotationService{
		repository: repository,
	}
}

func (service *CotationService) Create() error {
	cotation, err := service.repository.CreateCotation()
	if err != nil {
		return err
	}

	err = saveCotation(cotation)
	if err != nil {
		return err
	}

	return nil
}

func saveCotation(cotation *entity.Cotation) error {
	file, err := os.Create("tmp/cotacao.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	dolar, err := strconv.ParseFloat(cotation.Dolar, 64)
	if err != nil {
		return err
	}

	_, err = file.Write([]byte(fmt.Sprintf("Dólar:{%.4f}", dolar)))
	if err != nil {
		return err
	}

	return nil
}
