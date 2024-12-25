package mock

import (
	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
)

// work order
func (m *MockStore) CreateWorkOrder(wo tp.WorkOrder) (tp.WorkOrder, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.data["workOrders"] == nil {
		m.data["workOrders"] = make(map[string]interface{})
	}
	wo.Id = uuid.New()
	m.data["workOrders"][wo.Id.String()] = wo
	return wo, nil
}

func (m *MockStore) DeleteWorkOrder(id uuid.UUID) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data["workOrders"][id.String()]; !ok {
		return ae.New(ae.CodeNotFound, "work order not found")
	}
	delete(m.data["workOrders"], id.String())
	return nil
}

func (m *MockStore) GetWorkOrder(id uuid.UUID) (tp.WorkOrder, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if wo, ok := m.data["workOrders"][id.String()]; ok {
		return wo.(tp.WorkOrder), nil
	}
	return tp.WorkOrder{}, nil
}

func (m *MockStore) GetWorkOrderForTask(taskId uuid.UUID, woId uuid.UUID) (tp.WorkOrder, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if wo, ok := m.data["workOrders"][woId.String()]; ok {
		if wo.(tp.WorkOrder).TaskId == taskId {
			return wo.(tp.WorkOrder), nil
		}
	}
	return tp.WorkOrder{}, ae.New(ae.CodeNotFound, "work order not found")
}

func (m *MockStore) ListWorkOrders() ([]tp.WorkOrder, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var workOrders []tp.WorkOrder
	for _, wo := range m.data["workOrders"] {
		workOrders = append(workOrders, wo.(tp.WorkOrder))
	}
	return workOrders, nil
}

func (m *MockStore) ListWorkOrdersByTaskId(taskId uuid.UUID) ([]tp.WorkOrder, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var workOrders []tp.WorkOrder
	for _, wo := range m.data["workOrders"] {
		if wo.(tp.WorkOrder).TaskId == taskId {
			workOrders = append(workOrders, wo.(tp.WorkOrder))
		}
	}
	return workOrders, nil
}

func (m *MockStore) UpdateWorkOrder(id uuid.UUID, wo tp.WorkOrder) (tp.WorkOrder, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data["workOrders"][id.String()]; !ok {
		return tp.WorkOrder{}, ae.New(ae.CodeNotFound, "work order not found")
	}
	m.data["workOrders"][id.String()] = wo
	return wo, nil
}
