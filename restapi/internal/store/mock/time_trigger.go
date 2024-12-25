package mock

import (
	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
)

// time trigger
func (m *MockStore) CreateTimeTrigger(tt tp.TimeTrigger) (tp.TimeTrigger, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.data["timeTriggers"] == nil {
		m.data["timeTriggers"] = make(map[string]interface{})
	}
	tt.Id = uuid.New()
	m.data["timeTriggers"][tt.Id.String()] = tt
	return tt, nil
}

func (m *MockStore) DeleteTimeTrigger(id uuid.UUID) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data["timeTriggers"][id.String()]; !ok {
		return ae.New(ae.CodeNotFound, "time trigger not found")
	}
	delete(m.data["timeTriggers"], id.String())
	return nil
}

func (m *MockStore) GetTimeTrigger(id uuid.UUID) (tp.TimeTrigger, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if tt, ok := m.data["timeTriggers"][id.String()]; ok {
		return tt.(tp.TimeTrigger), nil
	}
	return tp.TimeTrigger{}, nil
}

func (m *MockStore) ListTimeTriggers() ([]tp.TimeTrigger, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var timeTriggers []tp.TimeTrigger
	for _, tt := range m.data["timeTriggers"] {
		timeTriggers = append(timeTriggers, tt.(tp.TimeTrigger))
	}
	return timeTriggers, nil
}

func (m *MockStore) ListTimeTriggersByTaskId(id uuid.UUID) ([]tp.TimeTrigger, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var timeTriggers []tp.TimeTrigger
	for _, tt := range m.data["timeTriggers"] {
		if tt.(tp.TimeTrigger).TaskId == id {
			timeTriggers = append(timeTriggers, tt.(tp.TimeTrigger))
		}
	}
	return timeTriggers, nil
}

func (m *MockStore) UpdateTimeTrigger(id uuid.UUID, tt tp.TimeTrigger) (tp.TimeTrigger, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data["timeTriggers"][id.String()]; !ok {
		return tp.TimeTrigger{}, ae.New(ae.CodeNotFound, "time trigger not found")
	}
	m.data["timeTriggers"][id.String()] = tt
	return tt, nil
}
