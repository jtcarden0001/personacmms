package cmmsapp

import (
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

type Category interface {
	CreateCategory(string, string) (tp.Category, error)
	DeleteCategory(string) error
	ListCategories() ([]tp.Category, error)
	GetCategory(string) (tp.Category, error)
	UpdateCategory(string, string, string) (tp.Category, error)
}

func (a *App) CreateCategory(title, description string) (tp.Category, error) {
	return a.db.CreateCategory(title, description)
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

func (a *App) UpdateCategory(oldTitle, newTitle, newDescription string) (tp.Category, error) {
	return a.db.UpdateCategory(oldTitle, newTitle, newDescription)
}
