package jsonImporter

import (
	"bufio"
	"context"
	"encoding/json"
	"os"
	"seaports-service-assignment/internal/domain/model"
	"seaports-service-assignment/internal/ports/importer"
	"seaports-service-assignment/internal/ports/store"
)

type JsonImporter struct {
	store store.Store
}

func NewJsonImporter(store store.Store) *JsonImporter {
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

	d.Token()
	for d.More() {
		var key string
		d.Decode(&key)

		port := model.Seaport{}
		d.Decode(&port)

		// load to store
		i.store.Create(port)
	}
	d.Token()

	return err
}

var _ importer.Importer = (*JsonImporter)(nil)
