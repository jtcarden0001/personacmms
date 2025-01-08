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
	err := a.validateCategory(cat)
	if err != nil {
		return apitp.CategoryResponse{}, errors.Wrapf(err, "CreateCategory validation failed")
	}

	catId := uuid.New()
	stCategory, err := a.convertApiCategoryRequestToStoreCategory(catId, cat)
	if err != nil {
		return apitp.CategoryResponse{}, errors.Wrapf(err, "CreateCategory ConvertApiCategoryRequestToStoreCategory failed")
	}

	stCategory, err = a.db.CreateCategory(stCategory)
	if err != nil {
		return apitp.CategoryResponse{}, errors.Wrapf(err, "CreateCategory CreateCategory failed")
	}

	return a.convertStoreCategoryToApiCategoryResponse(stCategory)
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

	return a.convertStoreCategoryListToApiCategoryResponseList(stCategories)
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

	return a.convertStoreCategoryToApiCategoryResponse(stCategory)
}

func (a *App) UpdateCategory(id string, cat apitp.CategoryRequest) (apitp.CategoryResponse, error) {
	catUuid, err := uuid.Parse(id)
	if err != nil {
		return apitp.CategoryResponse{}, ae.New(ae.CodeInvalid, "category id must be a valid uuid")
	}

	err = a.validateCategory(cat)
	if err != nil {
		return apitp.CategoryResponse{}, errors.Wrapf(err, "UpdateCategory validation failed")
	}

	stCatRequest, err := a.convertApiCategoryRequestToStoreCategory(catUuid, cat)
	if err != nil {
		return apitp.CategoryResponse{}, errors.Wrapf(err, "UpdateCategory ConvertApiCategoryRequestToStoreCategory failed")
	}

	stCatResponse, err := a.db.UpdateCategory(stCatRequest)
	if err != nil {
		return apitp.CategoryResponse{}, errors.Wrapf(err, "UpdateCategory UpdateCategory failed")
	}

	return a.convertStoreCategoryToApiCategoryResponse(stCatResponse)
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

	return a.convertStoreCategoryListToApiCategoryResponseList(stCatResponses)
}

func (a *App) validateCategory(cat apitp.CategoryRequest) error {
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

func (a *App) convertApiCategoryRequestToStoreCategory(id uuid.UUID, catRequest apitp.CategoryRequest) (storetp.Category, error) {
	storeCat := storetp.Category{
		Id:          id,
		Title:       catRequest.Title,
		Description: catRequest.Description,
	}

	return storeCat, nil
}

func (a *App) convertStoreCategoryListToApiCategoryResponseList(cats []storetp.Category) ([]apitp.CategoryResponse, error) {
	apiCats := make([]apitp.CategoryResponse, 0, len(cats))
	for _, c := range cats {
		apiCat, err := a.convertStoreCategoryToApiCategoryResponse(c)
		if err != nil {
			return nil, errors.Wrapf(err, "convertStoreCategoryListToApiCategoryResponseList failed")
		}
		apiCats = append(apiCats, apiCat)
	}

	return apiCats, nil
}

func (a *App) convertStoreCategoryToApiCategoryResponse(cat storetp.Category) (apitp.CategoryResponse, error) {
	apiCat := apitp.CategoryResponse{
		Id:          cat.Id,
		Title:       cat.Title,
		Description: cat.Description,
	}

	return apiCat, nil
}
