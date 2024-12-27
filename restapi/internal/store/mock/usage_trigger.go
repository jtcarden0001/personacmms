package mock

import (
	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
)

var usageTriggerTable = "usageTriggers"

func (m *MockStore) CreateUsageTrigger(ut tp.UsageTrigger) (tp.UsageTrigger, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.data[usageTriggerTable] == nil {
		m.data[usageTriggerTable] = make(map[uuid.UUID]interface{})
	}

	m.data[usageTriggerTable][ut.Id] = ut
	return ut, nil
}

func (m *MockStore) DeleteUsageTrigger(id uuid.UUID) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data[usageTriggerTable][id]; !ok {
		return ae.New(ae.CodeNotFound, "deleteusagetrigger - usage trigger not found")
	}

	delete(m.data[usageTriggerTable], id)
	return nil
}

func (m *MockStore) GetUsageTrigger(id uuid.UUID) (tp.UsageTrigger, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if ut, ok := m.data[usageTriggerTable][id]; ok {
		return ut.(tp.UsageTrigger), nil
	}

	return tp.UsageTrigger{}, ae.New(ae.CodeNotFound, "getusagetrigger - usage trigger not found")
}

func (m *MockStore) ListUsageTriggers() ([]tp.UsageTrigger, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var usageTriggers []tp.UsageTrigger
	for _, ut := range m.data[usageTriggerTable] {
		usageTriggers = append(usageTriggers, ut.(tp.UsageTrigger))
	}

	return usageTriggers, nil
}

func (m *MockStore) UpdateUsageTrigger(ut tp.UsageTrigger) (tp.UsageTrigger, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data[usageTriggerTable][ut.Id]; !ok {
		return tp.UsageTrigger{}, ae.New(ae.CodeNotFound, "updateusagetrigger - usage trigger not found")
	}

	m.data[usageTriggerTable][ut.Id] = ut
	return ut, nil
}
