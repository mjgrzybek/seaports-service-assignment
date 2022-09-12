package jsonImporter

import (
	"bufio"
	"context"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"os"
	"seaports-service-assignment/internal/domain/model"
	"seaports-service-assignment/internal/ports/importer"
	"seaports-service-assignment/internal/ports/store"
)

type JsonImporter struct {
	store store.Store
}

func NewJsonImporter(store store.Store) *JsonImporter {
	log.Info("JsonImporter initialized")
	return &JsonImporter{store: store}
}

// TODO: check and handle file and decoder errors
func (i JsonImporter) Import(ctx context.Context, sourcePath string) error {
	f, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer f.Close() // TODO: handle error

	r := bufio.NewReader(f)
	d := json.NewDecoder(r)

	_, err = d.Token()
	PanicOnError(err)
	for d.More() {
		select {
		case <-ctx.Done():
			return importer.ErrImportCancelled
		default:
		}

		key, _ := d.Token()

		port := model.Seaport{}
		d.Decode(&port)

		port.Id = key.(string)

		// load to store
		i.store.Create(port)
	}
	d.Token()

	return err
}

func PanicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

var _ importer.Importer = (*JsonImporter)(nil)
