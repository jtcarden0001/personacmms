package mock

import (
	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
)

// usage trigger
func (m *MockStore) CreateUsageTrigger(ut tp.UsageTrigger) (tp.UsageTrigger, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.data["usageTriggers"] == nil {
		m.data["usageTriggers"] = make(map[string]interface{})
	}
	ut.Id = uuid.New()
	m.data["usageTriggers"][ut.Id.String()] = ut
	return ut, nil
}

func (m *MockStore) DeleteUsageTrigger(id tp.UUID) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data["usageTriggers"][id.String()]; !ok {
		return ae.New(ae.CodeNotFound, "usage trigger not found")
	}
	delete(m.data["usageTriggers"], id.String())
	return nil
}

func (m *MockStore) GetUsageTrigger(id tp.UUID) (tp.UsageTrigger, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if ut, ok := m.data["usageTriggers"][id.String()]; ok {
		return ut.(tp.UsageTrigger), nil
	}
	return tp.UsageTrigger{}, nil
}

func (m *MockStore) ListUsageTriggers() ([]tp.UsageTrigger, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var usageTriggers []tp.UsageTrigger
	for _, ut := range m.data["usageTriggers"] {
		usageTriggers = append(usageTriggers, ut.(tp.UsageTrigger))
	}
	return usageTriggers, nil
}

func (m *MockStore) ListUsageTriggersByTaskId(id tp.UUID) ([]tp.UsageTrigger, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var usageTriggers []tp.UsageTrigger
	for _, ut := range m.data["usageTriggers"] {
		if ut.(tp.UsageTrigger).TaskId == id {
			usageTriggers = append(usageTriggers, ut.(tp.UsageTrigger))
		}
	}
	return usageTriggers, nil
}

func (m *MockStore) UpdateUsageTrigger(id tp.UUID, ut tp.UsageTrigger) (tp.UsageTrigger, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data["usageTriggers"][id.String()]; !ok {
		return tp.UsageTrigger{}, ae.New(ae.CodeNotFound, "usage trigger not found")
	}
	m.data["usageTriggers"][id.String()] = ut
	return ut, nil
}
