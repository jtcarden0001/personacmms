package mock

import (
	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
)

// category
func (m *MockStore) CreateCategory(category tp.Category) (tp.Category, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.data["categories"] == nil {
		m.data["categories"] = make(map[string]interface{})
	}
	category.Id = uuid.New()
	m.data["categories"][category.Title] = category
	return category, nil
}

func (m *MockStore) DeleteCategory(title string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data["categories"][title]; !ok {
		return ae.New(ae.CodeNotFound, "category not found")
	}
	delete(m.data["categories"], title)
	return nil
}

func (m *MockStore) GetCategory(title string) (tp.Category, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if category, ok := m.data["categories"][title]; ok {
		return category.(tp.Category), nil
	}
	return tp.Category{}, ae.ErrNotFound
}

func (m *MockStore) ListCategories() ([]tp.Category, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var categories []tp.Category
	for _, category := range m.data["categories"] {
		categories = append(categories, category.(tp.Category))
	}
	return categories, nil
}

func (m *MockStore) UpdateCategory(title string, category tp.Category) (tp.Category, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data["categories"][title]; !ok {
		return tp.Category{}, ae.New(ae.CodeNotFound, "category not found")
	}
	delete(m.data["categories"], title)
	m.data["categories"][category.Title] = category
	return category, nil
}
