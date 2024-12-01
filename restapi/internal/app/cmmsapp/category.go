package cmmsapp

import (
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

// A Category is a logical grouping of asset types

// Create a Category
func (a *App) CreateCategory(cat tp.Category) (tp.Category, error) {
	return a.db.CreateCategory(cat)
}

// Delete a Category
func (a *App) DeleteCategory(title string) error {
	// Get before delete so we can return a not found error.
	if _, err := a.GetCategory(title); err != nil {
		return err
	}

	return a.db.DeleteCategory(title)
}

// List all Categories
func (a *App) ListCategories() ([]tp.Category, error) {
	return a.db.ListCategories()
}

// Get a Category
func (a *App) GetCategory(title string) (tp.Category, error) {
	return a.db.GetCategory(title)
}

// Update a Category
func (a *App) UpdateCategory(oldTitle string, cat tp.Category) (tp.Category, error) {
	if cat.Title == "" {
		cat.Title = oldTitle
	}

	return a.db.UpdateCategory(oldTitle, cat)
}
