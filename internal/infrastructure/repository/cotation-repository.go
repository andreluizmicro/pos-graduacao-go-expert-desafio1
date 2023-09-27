package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

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
	req, err := http.Get(fmt.Sprintf("%s/json/last/USD-BRL", ECONOMIA_AWESOME_API))
	if err != nil {
		return nil, exceptions.ErrServiceUnavailable
	}
	defer req.Body.Close()

	res, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	var cotation *entity.Cotation
	json.Unmarshal(res, &cotation)

	cotation.ID = id.New().ID()

	stmt, err := repository.db.Prepare("INSERT INTO cotation (id, code, code_in, name, high, low, var_bid, pct_change, bid, ask, created_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(
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
		fmt.Println("dsadsadsada")
		return nil, exceptions.ErrInternalServerError
	}

	return cotation, nil
}
