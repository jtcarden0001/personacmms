package mock

import (
	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
)

// task
func (m *MockStore) CreateTask(task tp.Task) (tp.Task, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.data["tasks"] == nil {
		m.data["tasks"] = make(map[string]interface{})
	}
	task.Id = uuid.New()
	m.data["tasks"][task.Id.String()] = task
	return task, nil
}

func (m *MockStore) DeleteTask(id uuid.UUID) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data["tasks"][id.String()]; !ok {
		return ae.New(ae.CodeNotFound, "task not found")
	}
	delete(m.data["tasks"], id.String())
	return nil
}

func (m *MockStore) GetTask(id uuid.UUID) (tp.Task, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if task, ok := m.data["tasks"][id.String()]; ok {
		return task.(tp.Task), nil
	}
	return tp.Task{}, nil
}

func (m *MockStore) GetTaskByAssetId(assetId uuid.UUID, taskId uuid.UUID) (tp.Task, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if task, ok := m.data["tasks"][taskId.String()]; ok {
		if task.(tp.Task).AssetId == assetId {
			return task.(tp.Task), nil
		}
	}

	return tp.Task{}, ae.New(ae.CodeNotFound, "task not found")
}

func (m *MockStore) ListTasks() ([]tp.Task, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var tasks []tp.Task
	for _, task := range m.data["tasks"] {
		tasks = append(tasks, task.(tp.Task))
	}
	return tasks, nil
}

func (m *MockStore) ListTasksByAssetId(id uuid.UUID) ([]tp.Task, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var tasks []tp.Task
	for _, task := range m.data["tasks"] {
		if task.(tp.Task).AssetId == id {
			tasks = append(tasks, task.(tp.Task))
		}
	}
	return tasks, nil
}

func (m *MockStore) UpdateTask(id uuid.UUID, task tp.Task) (tp.Task, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data["tasks"][id.String()]; !ok {
		return tp.Task{}, ae.New(ae.CodeNotFound, "task not found")
	}
	m.data["tasks"][id.String()] = task
	return task, nil
}
