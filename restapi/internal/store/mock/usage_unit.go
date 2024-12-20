package mock

import (
	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
)

// usage unit
func (m *MockStore) CreateUsageUnit(uu tp.UsageUnit) (tp.UsageUnit, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.data["usageUnits"] == nil {
		m.data["usageUnits"] = make(map[string]interface{})
	}
	uu.Id = uuid.New()
	m.data["usageUnits"][uu.Title] = uu
	return uu, nil
}

func (m *MockStore) DeleteUsageUnit(title string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data["usageUnits"][title]; !ok {
		return ae.New(ae.CodeNotFound, "usage unit not found")
	}
	delete(m.data["usageUnits"], title)
	return nil
}

func (m *MockStore) GetUsageUnit(title string) (tp.UsageUnit, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if uu, ok := m.data["usageUnits"][title]; ok {
		return uu.(tp.UsageUnit), nil
	}
	return tp.UsageUnit{}, nil
}

func (m *MockStore) ListUsageUnits() ([]tp.UsageUnit, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var usageUnits []tp.UsageUnit
	for _, uu := range m.data["usageUnits"] {
		usageUnits = append(usageUnits, uu.(tp.UsageUnit))
	}
	return usageUnits, nil
}

func (m *MockStore) UpdateUsageUnit(title string, uu tp.UsageUnit) (tp.UsageUnit, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data["usageUnits"][title]; !ok {
		return tp.UsageUnit{}, ae.New(ae.CodeNotFound, "usage unit not found")
	}
	m.data["usageUnits"][title] = uu
	return uu, nil
}
