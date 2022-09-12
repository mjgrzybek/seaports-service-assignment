package api

import "seaports-service-assignment/internal/domain/model"

type Api interface {
	Create(port model.Seaports) error
	Update(port model.Seaports) error
}
