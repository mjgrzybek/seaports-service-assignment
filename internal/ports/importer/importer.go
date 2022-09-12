package importer

import "seaports-service-assignment/internal/ports/store"

type Importer interface {
	Import(sourcePath string, destination store.Store) error
}
