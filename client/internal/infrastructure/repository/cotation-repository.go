package repository

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/pos-graduacao-go-expert-desafio1/client/internal/domain/entity"
)

type CotationRepository struct {
}

func NewCotationRepository() *CotationRepository {
	return &CotationRepository{}
}

func (repository *CotationRepository) CreateCotation() (*entity.Cotation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3000*time.Millisecond)
	defer cancel()

	select {
	case <-time.After(200 * time.Millisecond):
		log.Println("Cotação realizada com sucesso")
	case <-ctx.Done():
		log.Println("timeout")
	}

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var cotation entity.Cotation
	json.Unmarshal(data, &cotation)

	return &cotation, nil
}
