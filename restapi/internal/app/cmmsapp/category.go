package cmmsapp

import (
	"fmt"

	"github.com/google/uuid"
	apitp "github.com/jtcarden0001/personacmms/restapi/internal/types/api"
	storetp "github.com/jtcarden0001/personacmms/restapi/internal/types/store"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
	"github.com/pkg/errors"
)

// TODO: ensure the returned category has a list of asset referenes associated with it

func (a *App) CreateCategory(cat apitp.CategoryRequest) (apitp.CategoryResponse, error) {
	if cat.Id != uuid.Nil {
		return apitp.CategoryResponse{}, ae.New(ae.CodeInvalid, "category id must be nil on create, we will create an id for you")
	}
	cat.Id = uuid.New()

	err := a.validateCategory(cat)
	if err != nil {
		return apitp.CategoryResponse{}, errors.Wrapf(err, "CreateCategory validation failed")
	}

	stCategory, err := convertApiCategoryRequestToStoreCategory(cat)
	if err != nil {
		return apitp.CategoryResponse{}, errors.Wrapf(err, "CreateCategory ConvertApiCategoryRequestToStoreCategory failed")
	}

	stCategory, err = a.db.CreateCategory(stCategory)
	if err != nil {
		return apitp.CategoryResponse{}, errors.Wrapf(err, "CreateCategory CreateCategory failed")
	}

	return convertStoreCategoryToApiCategoryResponse(stCategory)
}

func (a *App) DeleteCategory(id string) error {
	catUuid, err := uuid.Parse(id)
	if err != nil {
		return ae.New(ae.CodeInvalid, "category id must be a valid uuid")
	}

	// TODO: block deletion if the category is in use

	return a.db.DeleteCategory(catUuid)
}

func (a *App) ListCategories() ([]apitp.CategoryResponse, error) {
	stCategories, err := a.db.ListCategories()
	if err != nil {
		return nil, errors.Wrapf(err, "ListCategories failed")
	}

	return convertStoreCategoryListToApiCategoryResponseList(stCategories)
}

func (a *App) GetCategory(id string) (apitp.CategoryResponse, error) {
	catUuid, err := uuid.Parse(id)
	if err != nil {
		return apitp.CategoryResponse{}, ae.New(ae.CodeInvalid, "category id must be a valid uuid")
	}

	stCategory, err := a.db.GetCategory(catUuid)
	if err != nil {
		return apitp.CategoryResponse{}, errors.Wrapf(err, "GetCategory failed")
	}

	return convertStoreCategoryToApiCategoryResponse(stCategory)
}

func (a *App) UpdateCategory(id string, cat apitp.CategoryRequest) (apitp.CategoryResponse, error) {
	catUuid, err := uuid.Parse(id)
	if err != nil {
		return apitp.CategoryResponse{}, ae.New(ae.CodeInvalid, "category id must be a valid uuid")
	}

	if cat.Id != uuid.Nil && cat.Id != catUuid {
		return apitp.CategoryResponse{}, ae.New(ae.CodeInvalid, fmt.Sprintf("category id mismatch between [%s] and [%s]", id, cat.Id.String()))
	}

	cat.Id = catUuid
	err = a.validateCategory(cat)
	if err != nil {
		return apitp.CategoryResponse{}, errors.Wrapf(err, "UpdateCategory validation failed")
	}

	stCatRequest, err := convertApiCategoryRequestToStoreCategory(cat)
	if err != nil {
		return apitp.CategoryResponse{}, errors.Wrapf(err, "UpdateCategory ConvertApiCategoryRequestToStoreCategory failed")
	}

	stCatResponse, err := a.db.UpdateCategory(stCatRequest)
	if err != nil {
		return apitp.CategoryResponse{}, errors.Wrapf(err, "UpdateCategory UpdateCategory failed")
	}

	return convertStoreCategoryToApiCategoryResponse(stCatResponse)
}

// candidate to offload to store layer
func (a *App) ListCategoriesByAsset(assetId string) ([]apitp.CategoryResponse, error) {
	auid, aex, err := a.assetExists(assetId)
	if err != nil {
		return nil, errors.Wrapf(err, "ListCategoriesByAsset - GetAsset failed")
	}

	if !aex {
		return nil, ae.New(ae.CodeNotFound, fmt.Sprintf("asset with id [%s] not found", assetId))
	}

	stCatResponses, err := a.db.ListCategoriesByAsset(auid)
	if err != nil {
		return nil, errors.Wrapf(err, "ListCategoriesByAsset failed")
	}

	return convertStoreCategoryListToApiCategoryResponseList(stCatResponses)
}

func (a *App) validateCategory(cat apitp.CategoryRequest) error {
	if cat.Id == uuid.Nil {
		return ae.New(ae.CodeInvalid, "category id is required")
	}

	if len(cat.Title) < apitp.MinEntityTitleLength || len(cat.Title) > apitp.MaxEntityTitleLength {
		return ae.New(ae.CodeInvalid,
			fmt.Sprintf("category title length must be between [%d] and [%d] characters",
				apitp.MinEntityTitleLength,
				apitp.MaxEntityTitleLength))
	}

	return nil
}

func (a *App) categoryExists(id string) (uuid.UUID, bool, error) {
	cUuid, err := uuid.Parse(id)
	if err != nil || cUuid == uuid.Nil {
		return uuid.Nil, false, ae.New(ae.CodeInvalid, "category id must be a valid uuid")
	}

	_, err = a.db.GetCategory(cUuid)
	if err != nil {
		var appErr ae.AppError
		if errors.As(err, &appErr); appErr.Code == ae.CodeNotFound {
			return cUuid, false, nil
		}
		return cUuid, false, err
	}

	return cUuid, true, nil
}

func convertApiCategoryRequestToStoreCategory(catRequest apitp.CategoryRequest) (storetp.Category, error) {
	return storetp.Category{}, ae.New(ae.CodeNotImplemented, "ConvertApiCategoryRequestToStoreCategory not implemented")
}

func convertStoreCategoryListToApiCategoryResponseList(cats []storetp.Category) ([]apitp.CategoryResponse, error) {
	return []apitp.CategoryResponse{}, ae.New(ae.CodeNotImplemented, "ConvertStoreCategoryListToApiCategoryResponseList not implemented")
}

func convertStoreCategoryToApiCategoryResponse(cat storetp.Category) (apitp.CategoryResponse, error) {
	return apitp.CategoryResponse{}, ae.New(ae.CodeNotImplemented, "ConvertStoreCategoryToApiCategoryResponse not implemented")
}
