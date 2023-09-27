package entity

import "github.com/google/uuid"

type ID struct {
	value string
}

func New() *ID {
	return &ID{
		value: uuid.New().String(),
	}
}

func (id *ID) ID() string {
	return id.value
}
