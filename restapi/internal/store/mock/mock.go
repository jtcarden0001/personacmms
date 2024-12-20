package mock

import (
	"sync"
)

type MockStore struct {
	data map[string]map[string]interface{}
	mu   sync.RWMutex
}

func New() *MockStore {
	return &MockStore{
		data: make(map[string]map[string]interface{}),
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
