package mock

import (
	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
)

var groupTable = "groups"

func (m *MockStore) CreateGroup(group tp.Group) (tp.Group, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.data[groupTable] == nil {
		m.data[groupTable] = make(map[uuid.UUID]interface{})
	}

	m.data[groupTable][group.Id] = group
	return group, nil
}

func (m *MockStore) DeleteGroup(id uuid.UUID) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data[groupTable][id]; !ok {
		return ae.New(ae.CodeNotFound, "deletegroup - group not found")
	}

	delete(m.data[groupTable], id)
	return nil
}

func (m *MockStore) GetGroup(id uuid.UUID) (tp.Group, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if group, ok := m.data[groupTable][id]; ok {
		return group.(tp.Group), nil
	}

	return tp.Group{}, ae.New(ae.CodeNotFound, "getgroup - group not found")
}

func (m *MockStore) ListGroups() ([]tp.Group, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var groups []tp.Group
	for _, group := range m.data[groupTable] {
		groups = append(groups, group.(tp.Group))
	}

	return groups, nil
}

func (m *MockStore) UpdateGroup(group tp.Group) (tp.Group, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data[groupTable][group.Id]; !ok {
		return tp.Group{}, ae.New(ae.CodeNotFound, "updategroup - group not found")
	}

	m.data[groupTable][group.Id] = group
	return group, nil
}
