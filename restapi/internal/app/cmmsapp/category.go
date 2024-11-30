package cmmsapp

import (
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

func (a *App) CreateCategory(cat tp.Category) (tp.Category, error) {
	return a.db.CreateCategory(cat)
}

func (a *App) DeleteCategory(title string) error {
	if _, err := a.GetCategory(title); err != nil {
		return err
	}

	return a.db.DeleteCategory(title)
}

func (a *App) ListCategories() ([]tp.Category, error) {
	return a.db.ListCategories()
}

func (a *App) GetCategory(title string) (tp.Category, error) {
	return a.db.GetCategory(title)
}

func (a *App) UpdateCategory(oldTitle string, cat tp.Category) (tp.Category, error) {
	return a.db.UpdateCategory(oldTitle, cat)
}
