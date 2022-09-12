package httpServer

import (
	"context"
	"net/http"
	"seaports-service-assignment/internal/ports/api"
)

type HttpSeaportsServer struct {
	srv http.Server
}

func NewHttpSeaportsServer() *HttpSeaportsServer {
	return &HttpSeaportsServer{}
}

func (s *HttpSeaportsServer) Start(context.Context) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/create", s.create)
	mux.HandleFunc("/update", s.update)
	return http.ListenAndServe(":3333", mux)
}

func (s *HttpSeaportsServer) create(writer http.ResponseWriter, request *http.Request) {

}

func (s *HttpSeaportsServer) update(writer http.ResponseWriter, request *http.Request) {

}

var _ api.Api = (*HttpSeaportsServer)(nil)
