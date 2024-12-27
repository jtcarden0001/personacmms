package mock

import (
	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
)

var workOrderTable = "workOrders"

func (m *MockStore) CreateWorkOrder(wo tp.WorkOrder) (tp.WorkOrder, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.data[workOrderTable] == nil {
		m.data[workOrderTable] = make(map[uuid.UUID]interface{})
	}

	m.data[workOrderTable][wo.Id] = wo
	return wo, nil
}

func (m *MockStore) DeleteWorkOrder(id uuid.UUID) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data[workOrderTable][id]; !ok {
		return ae.New(ae.CodeNotFound, "deleteworkorder - work order not found")
	}

	delete(m.data[workOrderTable], id)
	return nil
}

func (m *MockStore) GetWorkOrder(id uuid.UUID) (tp.WorkOrder, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if wo, ok := m.data[workOrderTable][id]; ok {
		return wo.(tp.WorkOrder), nil
	}

	return tp.WorkOrder{}, ae.New(ae.CodeNotFound, "getworkorder - work order not found")
}

func (m *MockStore) ListWorkOrders() ([]tp.WorkOrder, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var workOrders []tp.WorkOrder
	for _, wo := range m.data[workOrderTable] {
		workOrders = append(workOrders, wo.(tp.WorkOrder))
	}

	return workOrders, nil
}

func (m *MockStore) UpdateWorkOrder(wo tp.WorkOrder) (tp.WorkOrder, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data[workOrderTable][wo.Id]; !ok {
		return tp.WorkOrder{}, ae.New(ae.CodeNotFound, "updateworkorder - work order not found")
	}

	m.data[workOrderTable][wo.Id] = wo
	return wo, nil
}
