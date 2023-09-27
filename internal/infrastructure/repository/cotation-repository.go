package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/pos-graduacao-go-expert-desafio1/internal/domain/entity"
	"github.com/pos-graduacao-go-expert-desafio1/internal/domain/exceptions"
	id "github.com/pos-graduacao-go-expert-desafio1/pkg/entity"
)

const ECONOMIA_AWESOME_API = "https://economia.awesomeapi.com.br"

type CotationRepository struct {
	db *sql.DB
}

func NewContationRepository(db *sql.DB) *CotationRepository {
	return &CotationRepository{
		db: db,
	}
}

func (repository *CotationRepository) Create() (*entity.Cotation, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s/json/last/USD-BRL", ECONOMIA_AWESOME_API), nil)
	if err != nil {
		return nil, exceptions.ErrServiceUnavailable
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var cotation *entity.Cotation
	json.Unmarshal(data, &cotation)

	cotation.ID = id.New().ID()

	stmt, err := repository.db.Prepare("INSERT INTO cotation (id, code, code_in, name, high, low, var_bid, pct_change, bid, ask, created_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	_, err = stmt.ExecContext(
		ctx,
		cotation.ID,
		cotation.USDBRL.Code,
		cotation.USDBRL.Codein,
		cotation.USDBRL.Name,
		cotation.USDBRL.High,
		cotation.USDBRL.Low,
		cotation.USDBRL.VarBid,
		cotation.USDBRL.PctChange,
		cotation.USDBRL.Bid,
		cotation.USDBRL.Ask,
		cotation.USDBRL.CreateDate,
	)
	if err != nil {
		return nil, exceptions.ErrInternalServerError
	}

	return cotation, nil
}
