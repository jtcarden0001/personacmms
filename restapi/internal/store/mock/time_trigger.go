package mock

import (
	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
)

var timeTriggerTable = "timeTriggers"

func (m *MockStore) CreateTimeTrigger(tt tp.TimeTrigger) (tp.TimeTrigger, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.data[timeTriggerTable] == nil {
		m.data[timeTriggerTable] = make(map[uuid.UUID]interface{})
	}

	m.data[timeTriggerTable][tt.Id] = tt
	return tt, nil
}

func (m *MockStore) DeleteTimeTrigger(id uuid.UUID) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data[timeTriggerTable][id]; !ok {
		return ae.New(ae.CodeNotFound, "deletetimetrigger - time trigger not found")
	}

	delete(m.data[timeTriggerTable], id)
	return nil
}

func (m *MockStore) GetTimeTrigger(id uuid.UUID) (tp.TimeTrigger, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if tt, ok := m.data[timeTriggerTable][id]; ok {
		return tt.(tp.TimeTrigger), nil
	}

	return tp.TimeTrigger{}, ae.New(ae.CodeNotFound, "gettimetrigger - time trigger not found")
}

func (m *MockStore) ListTimeTriggers() ([]tp.TimeTrigger, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var timeTriggers []tp.TimeTrigger
	for _, tt := range m.data[timeTriggerTable] {
		timeTriggers = append(timeTriggers, tt.(tp.TimeTrigger))
	}

	return timeTriggers, nil
}

func (m *MockStore) ListTimeTriggersByTaskId(id uuid.UUID) ([]tp.TimeTrigger, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var timeTriggers []tp.TimeTrigger
	for _, tt := range m.data[timeTriggerTable] {
		if tt.(tp.TimeTrigger).TaskId == id {
			timeTriggers = append(timeTriggers, tt.(tp.TimeTrigger))
		}
	}

	return timeTriggers, nil
}

func (m *MockStore) UpdateTimeTrigger(tt tp.TimeTrigger) (tp.TimeTrigger, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data[timeTriggerTable][tt.Id]; !ok {
		return tp.TimeTrigger{}, ae.New(ae.CodeNotFound, "updatetimetrigger - time trigger not found")
	}

	m.data[timeTriggerTable][tt.Id] = tt
	return tt, nil
}
