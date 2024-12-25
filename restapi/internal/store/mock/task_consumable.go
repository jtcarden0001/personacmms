package mock

import (
	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
)

// task consumable
func (m *MockStore) CreateTaskConsumable(tc tp.TaskConsumable) (tp.TaskConsumable, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.data["taskConsumables"] == nil {
		m.data["taskConsumables"] = make(map[string]interface{})
	}
	id := tc.TaskId.String() + tc.ConsumableId.String()
	m.data["taskConsumables"][id] = tc
	return tc, nil
}

func (m *MockStore) DeleteTaskConsumable(taskId, consumableId uuid.UUID) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	id := taskId.String() + consumableId.String()
	if _, ok := m.data["taskConsumables"][id]; !ok {
		return ae.New(ae.CodeNotFound, "task consumable not found")
	}
	delete(m.data["taskConsumables"], id)
	return nil
}

func (m *MockStore) GetTaskConsumable(taskId, consumableId uuid.UUID) (tp.TaskConsumable, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	id := taskId.String() + consumableId.String()
	if tc, ok := m.data["taskConsumables"][id]; ok {
		return tc.(tp.TaskConsumable), nil
	}
	return tp.TaskConsumable{}, nil
}

func (m *MockStore) ListTaskConsumables() ([]tp.TaskConsumable, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var taskConsumables []tp.TaskConsumable
	for _, tc := range m.data["taskConsumables"] {
		taskConsumables = append(taskConsumables, tc.(tp.TaskConsumable))
	}
	return taskConsumables, nil
}

func (m *MockStore) ListTaskConsumablesByTaskId(taskId uuid.UUID) ([]tp.TaskConsumable, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var taskConsumables []tp.TaskConsumable
	for _, tc := range m.data["taskConsumables"] {
		if tc.(tp.TaskConsumable).TaskId == taskId {
			taskConsumables = append(taskConsumables, tc.(tp.TaskConsumable))
		}
	}
	return taskConsumables, nil
}

func (m *MockStore) UpdateTaskConsumable(tc tp.TaskConsumable) (tp.TaskConsumable, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	id := tc.TaskId.String() + tc.ConsumableId.String()
	if _, ok := m.data["taskConsumables"][id]; !ok {
		return tp.TaskConsumable{}, ae.New(ae.CodeNotFound, "task consumable not found")
	}
	m.data["taskConsumables"][id] = tc
	return tc, nil
}
