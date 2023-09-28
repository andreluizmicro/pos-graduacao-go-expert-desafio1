package exceptions

import "errors"

var (
	ErrConvertFailed = errors.New("erro ao tentar converter a cotação")
)
