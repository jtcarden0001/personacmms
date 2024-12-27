package mock

import (
	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
)

var categoryTable = "categories"

func (m *MockStore) CreateCategory(category tp.Category) (tp.Category, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.data[categoryTable] == nil {
		m.data[categoryTable] = make(map[uuid.UUID]interface{})
	}

	m.data[categoryTable][category.Id] = category
	return category, nil
}

func (m *MockStore) DeleteCategory(id uuid.UUID) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data[categoryTable][id]; !ok {
		return ae.New(ae.CodeNotFound, "deletecategory - category not found")
	}

	delete(m.data[categoryTable], id)
	return nil
}

func (m *MockStore) GetCategory(id uuid.UUID) (tp.Category, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if category, ok := m.data[categoryTable][id]; ok {
		return category.(tp.Category), nil
	}

	return tp.Category{}, ae.New(ae.CodeNotFound, "getcategory - category not found")
}

func (m *MockStore) ListCategories() ([]tp.Category, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var categories []tp.Category
	for _, category := range m.data[categoryTable] {
		categories = append(categories, category.(tp.Category))
	}

	return categories, nil
}

func (m *MockStore) UpdateCategory(category tp.Category) (tp.Category, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data[categoryTable][category.Id]; !ok {
		return tp.Category{}, ae.New(ae.CodeNotFound, "updatecategory - category not found")
	}

	m.data[categoryTable][category.Id] = category
	return category, nil
}
