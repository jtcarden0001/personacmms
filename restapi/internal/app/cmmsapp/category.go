package cmmsapp

import (
	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
)

func (a *App) CreateCategory(cat tp.Category) (tp.Category, error) {
	return tp.Category{}, ae.New(ae.CodeNotImplemented, "CreateCategory not implemented")
}

func (a *App) DeleteCategory(id string) error {
	return ae.New(ae.CodeNotImplemented, "DeleteCategory not implemented")
}

func (a *App) ListCategories() ([]tp.Category, error) {
	return nil, ae.New(ae.CodeNotImplemented, "ListCategories not implemented")
}

func (a *App) GetCategory(id string) (tp.Category, error) {
	return tp.Category{}, ae.New(ae.CodeNotImplemented, "GetCategory not implemented")
}

func (a *App) UpdateCategory(id string, cat tp.Category) (tp.Category, error) {
	return tp.Category{}, ae.New(ae.CodeNotImplemented, "UpdateCategory not implemented")
}

func (a *App) ListCategoriesByAsset(assetId string) ([]tp.Category, error) {
	return nil, ae.New(ae.CodeNotImplemented, "ListCategoriesByAsset not implemented")
}

func (a *App) validateCategory(cat tp.Category) error {
	return ae.New(ae.CodeNotImplemented, "validateCategory not implemented")
}

func (a *App) categoryExists(id uuid.UUID) (bool, error) {
	return false, ae.New(ae.CodeNotImplemented, "categoryExists not implemented")
}
