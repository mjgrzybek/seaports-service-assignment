package store

import "seaports-service-assignment/internal/domain/model"

type Store interface {
	Create(port model.Seaports) error
	Update(port model.Seaports) error
}
