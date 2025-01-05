package convert

import (
	apitp "github.com/jtcarden0001/personacmms/restapi/internal/types/api"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"

	storetp "github.com/jtcarden0001/personacmms/restapi/internal/types/store"
)

func ConvertApiCategoryRequestToStoreCategory(catRequest apitp.CategoryRequest) (storetp.Category, error) {
	return storetp.Category{}, ae.New(ae.CodeNotImplemented, "ConvertApiCategoryRequestToStoreCategory not implemented")
}

func ConvertStoreCategoryListToApiCategoryResponseList(cats []storetp.Category) ([]apitp.CategoryResponse, error) {
	return []apitp.CategoryResponse{}, ae.New(ae.CodeNotImplemented, "ConvertStoreCategoryListToApiCategoryResponseList not implemented")
}

func ConvertStoreCategoryToApiCategoryResponse(cat storetp.Category) (apitp.CategoryResponse, error) {
	return apitp.CategoryResponse{}, ae.New(ae.CodeNotImplemented, "ConvertStoreCategoryToApiCategoryResponse not implemented")
}
