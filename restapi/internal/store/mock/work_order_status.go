package mock

import (
	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
)

// work order status
func (m *MockStore) CreateWorkOrderStatus(wos tp.WorkOrderStatus) (tp.WorkOrderStatus, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.data["workOrderStatuses"] == nil {
		m.data["workOrderStatuses"] = make(map[string]interface{})
	}
	wos.Id = uuid.New()
	m.data["workOrderStatuses"][wos.Title] = wos
	return wos, nil
}

func (m *MockStore) DeleteWorkOrderStatus(title string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data["workOrderStatuses"][title]; !ok {
		return ae.New(ae.CodeNotFound, "work order status not found")
	}
	delete(m.data["workOrderStatuses"], title)
	return nil
}

func (m *MockStore) GetWorkOrderStatus(title string) (tp.WorkOrderStatus, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if wos, ok := m.data["workOrderStatuses"][title]; ok {
		return wos.(tp.WorkOrderStatus), nil
	}
	return tp.WorkOrderStatus{}, nil
}

func (m *MockStore) ListWorkOrderStatuses() ([]tp.WorkOrderStatus, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var statuses []tp.WorkOrderStatus
	for _, wos := range m.data["workOrderStatuses"] {
		statuses = append(statuses, wos.(tp.WorkOrderStatus))
	}
	return statuses, nil
}

func (m *MockStore) UpdateWorkOrderStatus(title string, wos tp.WorkOrderStatus) (tp.WorkOrderStatus, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data["workOrderStatuses"][title]; !ok {
		return tp.WorkOrderStatus{}, ae.New(ae.CodeNotFound, "work order status not found")
	}
	m.data["workOrderStatuses"][wos.Title] = wos
	return wos, nil
}
