package mock

import (
	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
)

// time unit
func (m *MockStore) CreateTimeUnit(tu tp.TimeUnit) (tp.TimeUnit, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.data["timeUnits"] == nil {
		m.data["timeUnits"] = make(map[string]interface{})
	}
	tu.Id = uuid.New()
	m.data["timeUnits"][tu.Id.String()] = tu
	return tu, nil
}

func (m *MockStore) DeleteTimeUnit(id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data["timeUnits"][id]; !ok {
		return ae.New(ae.CodeNotFound, "time unit not found")
	}
	delete(m.data["timeUnits"], id)
	return nil
}

func (m *MockStore) GetTimeUnit(id string) (tp.TimeUnit, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if tu, ok := m.data["timeUnits"][id]; ok {
		return tu.(tp.TimeUnit), nil
	}
	return tp.TimeUnit{}, nil
}

func (m *MockStore) ListTimeUnits() ([]tp.TimeUnit, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var timeUnits []tp.TimeUnit
	for _, tu := range m.data["timeUnits"] {
		timeUnits = append(timeUnits, tu.(tp.TimeUnit))
	}
	return timeUnits, nil
}

func (m *MockStore) UpdateTimeUnit(id string, tu tp.TimeUnit) (tp.TimeUnit, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data["timeUnits"][id]; !ok {
		return tp.TimeUnit{}, ae.New(ae.CodeNotFound, "time unit not found")
	}
	m.data["timeUnits"][id] = tu
	return tu, nil
}
