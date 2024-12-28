package cmmsapp

import (
	"fmt"

	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
	"github.com/pkg/errors"
)

func (a *App) CreateCategory(cat tp.Category) (tp.Category, error) {
	if cat.Id != uuid.Nil {
		return tp.Category{}, ae.New(ae.CodeInvalid, "category id must be nil on create, we will create an id for you")
	}
	cat.Id = uuid.New()

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
	if cat.Id == uuid.Nil {
		return ae.New(ae.CodeInvalid, "category id is required")
	}

	if len(cat.Title) < tp.MinEntityTitleLength || len(cat.Title) > tp.MaxEntityTitleLength {
		return ae.New(ae.CodeInvalid,
			fmt.Sprintf("category title length must be between [%d] and [%d] characters",
				tp.MinEntityTitleLength,
				tp.MaxEntityTitleLength))
	}

	return nil
}

func (a *App) categoryExists(id uuid.UUID) (bool, error) {
	_, err := a.db.GetCategory(id)
	if err != nil {
		var appErr ae.AppError
		if errors.As(err, &appErr); appErr.Code == ae.CodeNotFound {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
