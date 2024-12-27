package mock

import (
	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
)

var taskTable = "tasks"

func (m *MockStore) CreateTask(task tp.Task) (tp.Task, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.data[taskTable] == nil {
		m.data[taskTable] = make(map[uuid.UUID]interface{})
	}

	m.data[taskTable][task.Id] = task
	return task, nil
}

func (m *MockStore) DeleteTask(id uuid.UUID) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data[taskTable][id]; !ok {
		return ae.New(ae.CodeNotFound, "deletetask - task not found")
	}

	delete(m.data[taskTable], id)
	return nil
}

func (m *MockStore) GetTask(id uuid.UUID) (tp.Task, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if task, ok := m.data[taskTable][id]; ok {
		return task.(tp.Task), nil
	}

	return tp.Task{}, ae.New(ae.CodeNotFound, "gettask - task not found")
}

func (m *MockStore) ListTasks() ([]tp.Task, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var tasks []tp.Task
	for _, task := range m.data[taskTable] {
		tasks = append(tasks, task.(tp.Task))
	}

	return tasks, nil
}

func (m *MockStore) UpdateTask(task tp.Task) (tp.Task, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data[taskTable][task.Id]; !ok {
		return tp.Task{}, ae.New(ae.CodeNotFound, "updatetask - task not found")
	}

	m.data[taskTable][task.Id] = task
	return task, nil
}
