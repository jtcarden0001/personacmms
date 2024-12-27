package mock

import (
	"sync"

	"github.com/google/uuid"
)

type MockStore struct {
	data map[string]map[uuid.UUID]interface{}
	mu   sync.RWMutex
}

func New() *MockStore {
	return &MockStore{
		data: make(map[string]map[uuid.UUID]interface{}),
	}
}

// utilities
func (m *MockStore) Exec(query string) error {
	// Mock implementation
	return nil
}

func (m *MockStore) Close() error {
	// Mock implementation
	return nil
}
