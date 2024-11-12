package cmmsapp

import (
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

type Category interface {
	CreateCategory(string) (tp.UUID, error)
	DeleteCategory(tp.UUID) error
	ListCategory() ([]tp.Category, error)
	GetCategory(tp.UUID) (tp.Category, error)
	UpdateCategory(tp.UUID, string) error
}

func (a *App) CreateCategory(title string) (tp.UUID, error) {
	return a.db.CreateCategory(title)
}

func (a *App) DeleteCategory(id tp.UUID) error {
	return a.db.DeleteCategory(id)
}

func (a *App) ListCategory() ([]tp.Category, error) {
	return a.db.ListCategory()
}

func (a *App) GetCategory(id tp.UUID) (tp.Category, error) {
	return a.db.GetCategory(id)
}

func (a *App) UpdateCategory(id tp.UUID, title string) error {
	return a.db.UpdateCategory(id, title)
}
