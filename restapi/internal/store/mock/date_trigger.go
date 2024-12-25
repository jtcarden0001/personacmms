package mock

import (
	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
)

// date trigger
func (m *MockStore) CreateDateTrigger(dt tp.DateTrigger) (tp.DateTrigger, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.data["dateTriggers"] == nil {
		m.data["dateTriggers"] = make(map[string]interface{})
	}
	dt.Id = uuid.New()
	m.data["dateTriggers"][dt.Id.String()] = dt
	return dt, nil
}

func (m *MockStore) DeleteDateTrigger(id uuid.UUID) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data["dateTriggers"][id.String()]; !ok {
		return ae.New(ae.CodeNotFound, "date trigger not found")
	}
	delete(m.data["dateTriggers"], id.String())
	return nil
}

func (m *MockStore) GetDateTrigger(id uuid.UUID) (tp.DateTrigger, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if dt, ok := m.data["dateTriggers"][id.String()]; ok {
		return dt.(tp.DateTrigger), nil
	}
	return tp.DateTrigger{}, nil
}

func (m *MockStore) ListDateTriggers() ([]tp.DateTrigger, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var dateTriggers []tp.DateTrigger
	for _, dt := range m.data["dateTriggers"] {
		dateTriggers = append(dateTriggers, dt.(tp.DateTrigger))
	}
	return dateTriggers, nil
}

func (m *MockStore) ListDateTriggersByTaskId(id uuid.UUID) ([]tp.DateTrigger, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var dateTriggers []tp.DateTrigger
	for _, dt := range m.data["dateTriggers"] {
		if dt.(tp.DateTrigger).TaskId == id {
			dateTriggers = append(dateTriggers, dt.(tp.DateTrigger))
		}
	}
	return dateTriggers, nil
}

func (m *MockStore) UpdateDateTrigger(id uuid.UUID, dt tp.DateTrigger) (tp.DateTrigger, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data["dateTriggers"][id.String()]; !ok {
		return tp.DateTrigger{}, ae.New(ae.CodeNotFound, "date trigger not found")
	}
	m.data["dateTriggers"][id.String()] = dt
	return dt, nil
}
