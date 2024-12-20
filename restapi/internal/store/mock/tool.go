package mock

import (
	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
)

// tool
func (m *MockStore) CreateTool(tool tp.Tool) (tp.Tool, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.data["tools"] == nil {
		m.data["tools"] = make(map[string]interface{})
	}
	tool.Id = uuid.New()
	m.data["tools"][tool.Title] = tool
	return tool, nil
}

func (m *MockStore) DeleteTool(title string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data["tools"][title]; !ok {
		return ae.New(ae.CodeNotFound, "tool not found")
	}
	delete(m.data["tools"], title)
	return nil
}

func (m *MockStore) GetTool(title string) (tp.Tool, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if tool, ok := m.data["tools"][title]; ok {
		return tool.(tp.Tool), nil
	}
	return tp.Tool{}, nil
}

func (m *MockStore) GetToolById(id uuid.UUID) (tp.Tool, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	for _, tool := range m.data["tools"] {
		if tool.(tp.Tool).Id == id {
			return tool.(tp.Tool), nil
		}
	}

	return tp.Tool{}, nil
}

func (m *MockStore) ListTools() ([]tp.Tool, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var tools []tp.Tool
	for _, tool := range m.data["tools"] {
		tools = append(tools, tool.(tp.Tool))
	}
	return tools, nil
}

func (m *MockStore) UpdateTool(title string, tool tp.Tool) (tp.Tool, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data["tools"][title]; !ok {
		return tp.Tool{}, ae.New(ae.CodeNotFound, "tool not found")
	}
	m.data["tools"][title] = tool
	return tool, nil
}
