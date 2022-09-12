package inMemoryStore

import (
	"seaports-service-assignment/internal/domain/model"
	"seaports-service-assignment/internal/ports/store"
)

type InMemoryStore struct {
}

func (i InMemoryStore) Create(port model.Seaports) error {
	//TODO implement me
	panic("implement me")
}

func (i InMemoryStore) Update(port model.Seaports) error {
	//TODO implement me
	panic("implement me")
}

var _ store.Store = (*InMemoryStore)(nil)
