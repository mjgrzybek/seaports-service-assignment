package importer

import (
	"context"
	"seaports-service-assignment/internal/ports/store"
)

type Importer interface {
	Import(ctx context.Context, sourcePath string, destination store.Store) error
}
