package exceptions

import "errors"

var (
	ErrServiceUnavailable  = errors.New("servidor temporariamente indisponível")
	ErrInternalServerError = errors.New("erro interno do servidor")
)
