package inMemoryStore

import (
	log "github.com/sirupsen/logrus"
	"seaports-service-assignment/internal/domain/model"
	"seaports-service-assignment/internal/ports/store"
	"sync"
)

type InMemoryStore struct {
	mutext sync.Mutex
	m      map[string]model.Seaport
}

func NewInMemoryStore() *InMemoryStore {
	log.Info("Store initialized")
	return &InMemoryStore{
		m: map[string]model.Seaport{},
	}
}

func (s *InMemoryStore) Create(port model.Seaport) error {
	s.mutext.Lock()
	defer s.mutext.Unlock()

	if _, ok := s.m[port.Id]; ok {
		return store.ErrAlreadyExists
	}

	s.m[port.Id] = port
	return nil
}

func (s *InMemoryStore) Update(port model.Seaport) error {
	s.mutext.Lock()
	defer s.mutext.Unlock()

	if _, ok := s.m[port.Id]; !ok {
		return store.ErrNotFound
	}

	s.m[port.Id] = port
	return nil
}

var _ store.Store = (*InMemoryStore)(nil)
