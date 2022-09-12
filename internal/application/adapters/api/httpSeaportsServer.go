package httpServer

import (
	"context"
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
	"net/http"
	"seaports-service-assignment/internal/domain/model"
	"seaports-service-assignment/internal/ports/api"
	"seaports-service-assignment/internal/ports/store"
)

type HttpSeaportsServer struct {
	store store.Store
}

func NewHttpSeaportsServer(store store.Store) *HttpSeaportsServer {
	return &HttpSeaportsServer{
		store: store,
	}
}

func (s *HttpSeaportsServer) Start(context.Context) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/create", s.create)
	mux.HandleFunc("/update", s.update)

	log.Info("Staring HttpSeaportsServer")
	return http.ListenAndServe(":3333", mux)
}

func (s *HttpSeaportsServer) create(writer http.ResponseWriter, request *http.Request) {
	p, err := decodeSeaport(request)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	err = s.store.Create(p)
	if errors.Is(err, store.ErrAlreadyExists) {
		http.Error(writer, err.Error(), http.StatusConflict)
		return
	}

	writer.WriteHeader(http.StatusCreated)
}

func (s *HttpSeaportsServer) update(writer http.ResponseWriter, request *http.Request) {
	p, err := decodeSeaport(request)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	err = s.store.Update(p)
	if errors.Is(err, store.ErrAlreadyExists) {
		http.Error(writer, err.Error(), http.StatusNotFound)
		return
	}

	writer.WriteHeader(http.StatusCreated)
}

func decodeSeaport(request *http.Request) (model.Seaport, error) {
	var p model.Seaport
	err := json.NewDecoder(request.Body).Decode(&p)
	return p, err
}

var _ api.Api = (*HttpSeaportsServer)(nil)
