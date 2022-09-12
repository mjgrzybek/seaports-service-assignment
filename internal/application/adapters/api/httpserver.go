package httpServer

import (
	"seaports-service-assignment/internal/domain/model"
	"seaports-service-assignment/internal/ports/api"
)

type HttpServer struct{}

func (s HttpServer) Create(port model.Seaport) error {
	//TODO implement me
	panic("implement me")
}

func (s HttpServer) Update(port model.Seaport) error {
	//TODO implement me
	panic("implement me")
}

var _ api.Api = (*HttpServer)(nil)
