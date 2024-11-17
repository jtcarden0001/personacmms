package cmmsapp

import (
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

type Category interface {
	CreateCategory(tp.Category) (tp.Category, error)
	DeleteCategory(string) error
	ListCategories() ([]tp.Category, error)
	GetCategory(string) (tp.Category, error)
	UpdateCategory(string, tp.Category) (tp.Category, error)
}

func (a *App) CreateCategory(cat tp.Category) (tp.Category, error) {
	return a.db.CreateCategory(cat)
}

func (a *App) DeleteCategory(title string) error {
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
