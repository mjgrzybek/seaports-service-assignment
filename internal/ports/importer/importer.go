package importer

import (
	"context"
	"errors"
)

var ErrImportCancelled = errors.New("Import cancelled by context")

type Importer interface {
	Import(ctx context.Context, sourcePath string) error
}
