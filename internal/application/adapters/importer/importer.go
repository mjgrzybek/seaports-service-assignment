package jsonImporter

import (
	"context"
	"seaports-service-assignment/internal/ports/importer"
	"seaports-service-assignment/internal/ports/store"
)

type JsonImporter struct {
}

func (i JsonImporter) Import(ctx context.Context, sourcePath string, destination store.Store) error {
	//TODO implement me
	panic("implement me")
}

var _ importer.Importer = (*JsonImporter)(nil)
