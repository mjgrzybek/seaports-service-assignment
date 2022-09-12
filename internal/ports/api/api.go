package api

import "seaports-service-assignment/internal/domain/model"

type Api interface {
	Create(port model.Seaport) error
	Update(port model.Seaport) error
}
