package entity

type Cotation struct {
	Dolar string `json:"bid"`
}

func NewCotation(value string) (*Cotation, error) {

	return &Cotation{
		Dolar: value,
	}, nil
}
