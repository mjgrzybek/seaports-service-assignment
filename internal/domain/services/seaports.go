package services

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"seaports-service-assignment/internal/application/config"
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
	// import data
	if viper.IsSet(config.PortsFile) {
		if err := s.Importer.Import(ctx, viper.GetString(config.PortsFile), s.Store); err != nil {
			log.WithError(err).Errorln("Failed to import ports from file")
			return
		}
	}

	// start server
	_ = s.Api.Start(ctx)

	safeExit.Done <- struct{}{}
}
