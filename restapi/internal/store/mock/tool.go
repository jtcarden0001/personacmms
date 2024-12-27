package mock

import (
	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
)

var toolTable = "tools"

func (m *MockStore) CreateTool(tool tp.Tool) (tp.Tool, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.data[toolTable] == nil {
		m.data[toolTable] = make(map[uuid.UUID]interface{})
	}

	m.data[toolTable][tool.Id] = tool
	return tool, nil
}

func (m *MockStore) DeleteTool(id uuid.UUID) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data[toolTable][id]; !ok {
		return ae.New(ae.CodeNotFound, "deletetool - tool not found")
	}

	delete(m.data[toolTable], id)
	return nil
}

func (m *MockStore) GetTool(id uuid.UUID) (tp.Tool, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if tool, ok := m.data[toolTable][id]; ok {
		return tool.(tp.Tool), nil
	}

	return tp.Tool{}, ae.New(ae.CodeNotFound, "gettool - tool not found")
}

func (m *MockStore) ListTools() ([]tp.Tool, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var tools []tp.Tool
	for _, tool := range m.data[toolTable] {
		tools = append(tools, tool.(tp.Tool))
	}

	return tools, nil
}

func (m *MockStore) UpdateTool(tool tp.Tool) (tp.Tool, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data[toolTable][tool.Id]; !ok {
		return tp.Tool{}, ae.New(ae.CodeNotFound, "updatetool - tool not found")
	}

	m.data[toolTable][tool.Id] = tool
	return tool, nil
}
