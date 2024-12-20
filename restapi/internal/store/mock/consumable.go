package mock

import (
	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
)

// consumable
func (m *MockStore) CreateConsumable(consumable tp.Consumable) (tp.Consumable, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.data["consumables"] == nil {
		m.data["consumables"] = make(map[string]interface{})
	}
	consumable.Id = uuid.New()
	m.data["consumables"][consumable.Title] = consumable
	return consumable, nil
}

func (m *MockStore) DeleteConsumable(title string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data["consumables"][title]; !ok {
		return ae.New(ae.CodeNotFound, "consumable not found")
	}
	delete(m.data["consumables"], title)
	return nil
}

func (m *MockStore) GetConsumableByTitle(title string) (tp.Consumable, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if consumable, ok := m.data["consumables"][title]; ok {
		return consumable.(tp.Consumable), nil
	}
	return tp.Consumable{}, nil
}

func (m *MockStore) GetConsumableById(id uuid.UUID) (tp.Consumable, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	for _, consumable := range m.data["consumables"] {
		if consumable.(tp.Consumable).Id == id {
			return consumable.(tp.Consumable), nil
		}
	}
	return tp.Consumable{}, nil
}

func (m *MockStore) ListConsumables() ([]tp.Consumable, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var consumables []tp.Consumable
	for _, consumable := range m.data["consumables"] {
		consumables = append(consumables, consumable.(tp.Consumable))
	}
	return consumables, nil
}

func (m *MockStore) UpdateConsumable(title string, consumable tp.Consumable) (tp.Consumable, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data["consumables"][title]; !ok {
		return tp.Consumable{}, ae.New(ae.CodeNotFound, "consumable not found")
	}
	m.data["consumables"][title] = consumable
	return consumable, nil
}
