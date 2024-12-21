package mock

import (
	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
)

// group
func (m *MockStore) CreateGroup(group tp.Group) (tp.Group, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.data["groups"] == nil {
		m.data["groups"] = make(map[string]interface{})
	}
	group.Id = uuid.New()
	m.data["groups"][group.Title] = group
	return group, nil
}

func (m *MockStore) DeleteGroup(title string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data["groups"][title]; !ok {
		return ae.New(ae.CodeNotFound, "group not found")
	}
	delete(m.data["groups"], title)
	return nil
}

func (m *MockStore) GetGroup(title string) (tp.Group, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if group, ok := m.data["groups"][title]; ok {
		return group.(tp.Group), nil
	}

	return tp.Group{}, ae.ErrNotFound
}

func (m *MockStore) ListGroups() ([]tp.Group, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var groups []tp.Group
	for _, group := range m.data["groups"] {
		groups = append(groups, group.(tp.Group))
	}
	return groups, nil
}

func (m *MockStore) UpdateGroup(title string, group tp.Group) (tp.Group, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data["groups"][title]; !ok {
		return tp.Group{}, ae.New(ae.CodeNotFound, "group not found")
	}
	m.data["groups"][title] = group
	return group, nil
}
