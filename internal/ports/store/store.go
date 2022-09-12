package store

import (
	"errors"
	"seaports-service-assignment/internal/domain/model"
)

var ErrAlreadyExists = errors.New("Already exists in database")
var ErrNotFound = errors.New("Not found")

type Store interface {
	Create(port model.Seaport) error
	Update(port model.Seaport) error
}
