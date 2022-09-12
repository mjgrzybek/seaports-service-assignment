package importer

import (
	"context"
)

type Importer interface {
	Import(ctx context.Context, sourcePath string) error
}
