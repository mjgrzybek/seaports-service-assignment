package services

import (
	"context"
	log "github.com/sirupsen/logrus"
	"seaports-service-assignment/internal/application/safeExit"
	"seaports-service-assignment/internal/ports/api"
	"seaports-service-assignment/internal/ports/importer"
	"seaports-service-assignment/internal/ports/store"
)

type Seaports struct {
	Store    store.Store
	Api      api.Api
	Importer importer.Importer
}

func (s Seaports) Run(ctx context.Context) {
	log.Println("Domain logic execution here")
	safeExit.Done <- struct{}{}
}
