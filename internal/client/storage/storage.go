// Operational storage
package storage

import (
	"context"
	"errors"
	"sync"

	"github.com/srg-bnd/keeper/internal/client/models"
)

// Type for Storage
type Storage struct {
	items map[uint]models.Secret
	mtx   sync.RWMutex
}

// Returns new storage
func NewStorage() *Storage {
	return &Storage{
		mtx:   sync.RWMutex{},
		items: make(map[uint]models.Secret),
	}
}

// Add secrets to storage
func (s *Storage) AddList(ctx context.Context, secrets []models.Secret) error {
	for _, secret := range secrets {
		s.Add(ctx, secret)
	}

	return nil
}

// Adds secret to storage
func (s *Storage) Add(ctx context.Context, secret models.Secret) error {
	s.mtx.Lock()
	s.items[secret.ID] = secret
	s.mtx.Unlock()

	return nil
}

// Deletes secret from storage
func (s *Storage) Delete(ctx context.Context, secretID uint) error {
	s.mtx.Lock()
	if _, ok := s.items[secretID]; ok {
		delete(s.items, secretID)
	}
	s.mtx.Unlock()

	return nil
}

// Gets secret by ID
func (s *Storage) Get(ctx context.Context, secretID uint) (*models.Secret, error) {
	s.mtx.RLock()
	secret, ok := s.items[secretID]
	s.mtx.RLock()

	if ok {
		return &secret, nil
	} else {
		return nil, errors.New("not found")
	}
}

// Gets secret list
func (s *Storage) List(ctx context.Context) (map[uint]models.Secret, error) {
	return s.items, nil
}
