package mock

import (
	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
)

var consumableTable = "consumables"

func (m *MockStore) CreateConsumable(consumable tp.Consumable) (tp.Consumable, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.data[consumableTable] == nil {
		m.data[consumableTable] = make(map[uuid.UUID]interface{})
	}

	m.data[consumableTable][consumable.Id] = consumable
	return consumable, nil
}

func (m *MockStore) DeleteConsumable(id uuid.UUID) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data[consumableTable][id]; !ok {
		return ae.New(ae.CodeNotFound, "deleteconsumable - consumable not found")
	}

	delete(m.data[consumableTable], id)
	return nil
}

func (m *MockStore) GetConsumable(id uuid.UUID) (tp.Consumable, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if consumable, ok := m.data[consumableTable][id]; ok {
		return consumable.(tp.Consumable), nil
	}

	return tp.Consumable{}, ae.New(ae.CodeNotFound, "getconsumable - consumable not found")
}

func (m *MockStore) ListConsumables() ([]tp.Consumable, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var consumables []tp.Consumable
	for _, consumable := range m.data[consumableTable] {
		consumables = append(consumables, consumable.(tp.Consumable))
	}

	return consumables, nil
}

func (m *MockStore) UpdateConsumable(consumable tp.Consumable) (tp.Consumable, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data[consumableTable][consumable.Id]; !ok {
		return tp.Consumable{}, ae.New(ae.CodeNotFound, "updateconsumable - consumable not found")
	}

	m.data[consumableTable][consumable.Id] = consumable
	return consumable, nil
}
