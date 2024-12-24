package cmmsapp

import (
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
)

func (a *App) CreateCategory(cat tp.Category) (tp.Category, error) {
	return tp.Category{}, ae.New(ae.CodeNotImplemented, "CreateCategory not implemented")
}

func (a *App) DeleteCategory(title string) error {
	return ae.New(ae.CodeNotImplemented, "DeleteCategory not implemented")
}

func (a *App) ListCategories() ([]tp.Category, error) {
	return nil, ae.New(ae.CodeNotImplemented, "ListCategories not implemented")
}

func (a *App) GetCategory(title string) (tp.Category, error) {
	return tp.Category{}, ae.New(ae.CodeNotImplemented, "GetCategory not implemented")
}

func (a *App) UpdateCategory(oldTitle string, cat tp.Category) (tp.Category, error) {
	return tp.Category{}, ae.New(ae.CodeNotImplemented, "UpdateCategory not implemented")
}

func (a *App) ListCategoriesByAsset(assetId string) ([]tp.Category, error) {
	return nil, ae.New(ae.CodeNotImplemented, "ListCategoriesByAsset not implemented")
}
