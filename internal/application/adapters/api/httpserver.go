package httpServer

import (
	"context"
	"seaports-service-assignment/internal/ports/api"
)

type HttpServer struct{}

func (s HttpServer) Start(context.Context) {
	//TODO implement me
	panic("implement me")
}

var _ api.Api = (*HttpServer)(nil)
