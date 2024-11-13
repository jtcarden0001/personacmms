package cmmsapp

import (
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

type Category interface {
	CreateCategory(string) (tp.Category, error)
	DeleteCategory(string) error
	ListCategory() ([]tp.Category, error)
	GetCategory(string) (tp.Category, error)
	UpdateCategory(string, string) (tp.Category, error)
}

func (a *App) CreateCategory(title string) (tp.Category, error) {
	return a.db.CreateCategory(title)
}

func (a *App) DeleteCategory(title string) error {
	return a.db.DeleteCategory(title)
}

func (a *App) ListCategory() ([]tp.Category, error) {
	return a.db.ListCategory()
}

func (a *App) GetCategory(title string) (tp.Category, error) {
	return a.db.GetCategory(title)
}

func (a *App) UpdateCategory(oldTitle, newTitle string) (tp.Category, error) {
	return a.db.UpdateCategory(oldTitle, newTitle)
}
