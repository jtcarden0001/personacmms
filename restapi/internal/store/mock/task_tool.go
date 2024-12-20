package mock

import (
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
)

// task tool
func (m *MockStore) CreateTaskTool(tt tp.TaskTool) (tp.TaskTool, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.data["taskTools"] == nil {
		m.data["taskTools"] = make(map[string]interface{})
	}
	id := tt.TaskId.String() + tt.ToolId.String()
	m.data["taskTools"][id] = tt
	return tt, nil
}

func (m *MockStore) DeleteTaskTool(taskId, toolId tp.UUID) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	id := taskId.String() + toolId.String()
	if _, ok := m.data["taskTools"][id]; !ok {
		return ae.New(ae.CodeNotFound, "task tool not found")
	}
	delete(m.data["taskTools"], id)
	return nil
}

func (m *MockStore) GetTaskTool(taskId, toolId tp.UUID) (tp.TaskTool, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	id := taskId.String() + toolId.String()
	if tt, ok := m.data["taskTools"][id]; ok {
		return tt.(tp.TaskTool), nil
	}
	return tp.TaskTool{}, nil
}

func (m *MockStore) ListTaskTools() ([]tp.TaskTool, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var taskTools []tp.TaskTool
	for _, tt := range m.data["taskTools"] {
		taskTools = append(taskTools, tt.(tp.TaskTool))
	}
	return taskTools, nil
}

func (m *MockStore) ListTaskToolsByTaskId(taskId tp.UUID) ([]tp.TaskTool, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var taskTools []tp.TaskTool
	for _, tt := range m.data["taskTools"] {
		if tt.(tp.TaskTool).TaskId == taskId {
			taskTools = append(taskTools, tt.(tp.TaskTool))
		}
	}
	return taskTools, nil
}
