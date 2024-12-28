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

	err := a.validateCategory(cat)
	if err != nil {
		return tp.Category{}, errors.Wrapf(err, "CreateCategory validation failed")
	}

	return a.db.CreateCategory(cat)
}

func (a *App) DeleteCategory(id string) error {
	catUuid, err := uuid.Parse(id)
	if err != nil {
		return ae.New(ae.CodeInvalid, "category id must be a valid uuid")
	}

	return a.db.DeleteCategory(catUuid)
}

func (a *App) ListCategories() ([]tp.Category, error) {
	return a.db.ListCategories()
}

func (a *App) GetCategory(id string) (tp.Category, error) {
	catUuid, err := uuid.Parse(id)
	if err != nil {
		return tp.Category{}, ae.New(ae.CodeInvalid, "category id must be a valid uuid")
	}

	return a.db.GetCategory(catUuid)
}

func (a *App) UpdateCategory(id string, cat tp.Category) (tp.Category, error) {
	catUuid, err := uuid.Parse(id)
	if err != nil {
		return tp.Category{}, ae.New(ae.CodeInvalid, "category id must be a valid uuid")
	}

	if cat.Id != uuid.Nil && cat.Id != catUuid {
		return tp.Category{}, ae.New(ae.CodeInvalid, fmt.Sprintf("category id mismatch between [%s] and [%s]", id, cat.Id.String()))
	}

	cat.Id = catUuid
	err = a.validateCategory(cat)
	if err != nil {
		return tp.Category{}, errors.Wrapf(err, "UpdateCategory validation failed")
	}

	return a.db.UpdateCategory(cat)
}

// candidate to offload to store layer
func (a *App) ListCategoriesByAsset(assetId string) ([]tp.Category, error) {
	assetUuid, err := uuid.Parse(assetId)
	if err != nil {
		return nil, ae.New(ae.CodeInvalid, "asset id must be a valid uuid")
	}

	aex, err := a.assetExists(assetUuid)
	if err != nil {
		return nil, errors.Wrapf(err, "ListCategoriesByAsset - GetAsset failed")
	}

	if !aex {
		return nil, ae.New(ae.CodeNotFound, fmt.Sprintf("asset with id [%s] not found", assetUuid.String()))
	}

	// TODO: implement supporting functionality in the store layer
	return []tp.Category{}, ae.New(ae.CodeNotImplemented, "ListCategoriesByAsset not implemented")
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
