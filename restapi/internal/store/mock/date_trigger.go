package mock

import (
	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
)

var dtTable = "dateTriggers"

func (m *MockStore) CreateDateTrigger(dt tp.DateTrigger) (tp.DateTrigger, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.data[dtTable] == nil {
		m.data[dtTable] = make(map[uuid.UUID]interface{})
	}

	m.data[dtTable][dt.Id] = dt
	return dt, nil
}

func (m *MockStore) DeleteDateTrigger(id uuid.UUID) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data[dtTable][id]; !ok {
		return ae.New(ae.CodeNotFound, "deletedatetrigger - date trigger not found")
	}

	delete(m.data[dtTable], id)
	return nil
}

func (m *MockStore) GetDateTrigger(id uuid.UUID) (tp.DateTrigger, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if dt, ok := m.data[dtTable][id]; ok {
		return dt.(tp.DateTrigger), nil
	}

	return tp.DateTrigger{}, ae.New(ae.CodeNotFound, "getdatetrigger - date trigger not found")
}

func (m *MockStore) ListDateTriggers() ([]tp.DateTrigger, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var dateTriggers []tp.DateTrigger
	for _, dt := range m.data[dtTable] {
		dateTriggers = append(dateTriggers, dt.(tp.DateTrigger))
	}

	return dateTriggers, nil
}

func (m *MockStore) UpdateDateTrigger(dt tp.DateTrigger) (tp.DateTrigger, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data[dtTable][dt.Id]; !ok {
		return tp.DateTrigger{}, ae.New(ae.CodeNotFound, "updatedatetrigger - date trigger not found")
	}

	m.data[dtTable][dt.Id] = dt
	return dt, nil
}
