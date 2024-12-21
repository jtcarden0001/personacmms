package mock

import (
	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
)

// task template
func (m *MockStore) CreateTaskTemplate(tt tp.TaskTemplate) (tp.TaskTemplate, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.data["taskTemplates"] == nil {
		m.data["taskTemplates"] = make(map[string]interface{})
	}
	tt.Id = uuid.New()
	m.data["taskTemplates"][tt.Title] = tt
	return tt, nil
}

func (m *MockStore) DeleteTaskTemplate(id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data["taskTemplates"][id]; !ok {
		return ae.New(ae.CodeNotFound, "task template not found")
	}
	delete(m.data["taskTemplates"], id)
	return nil
}

func (m *MockStore) GetTaskTemplate(id string) (tp.TaskTemplate, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if tt, ok := m.data["taskTemplates"][id]; ok {
		return tt.(tp.TaskTemplate), nil
	}
	return tp.TaskTemplate{}, nil
}

func (m *MockStore) ListTaskTemplates() ([]tp.TaskTemplate, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var taskTemplates []tp.TaskTemplate
	for _, tt := range m.data["taskTemplates"] {
		taskTemplates = append(taskTemplates, tt.(tp.TaskTemplate))
	}
	return taskTemplates, nil
}

func (m *MockStore) UpdateTaskTemplate(title string, tt tp.TaskTemplate) (tp.TaskTemplate, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data["taskTemplates"][title]; !ok {
		return tp.TaskTemplate{}, ae.New(ae.CodeNotFound, "task template not found")
	}
	delete(m.data["taskTemplates"], title)
	m.data["taskTemplates"][tt.Title] = tt
	return tt, nil
}
