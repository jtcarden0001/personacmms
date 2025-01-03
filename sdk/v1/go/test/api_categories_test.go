/*
PersonaCMMS API

Testing CategoriesAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/jtcarden0001/personacmms/sdk/v1/go"
)

func Test_openapi_CategoriesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CategoriesAPIService AssetsAssetIdCategoriesGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var assetId string

		resp, httpRes, err := apiClient.CategoriesAPI.AssetsAssetIdCategoriesGet(context.Background(), assetId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CategoriesAPIService CategoriesCategoryIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var categoryId string

		httpRes, err := apiClient.CategoriesAPI.CategoriesCategoryIdDelete(context.Background(), categoryId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CategoriesAPIService CategoriesCategoryIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var categoryId string

		resp, httpRes, err := apiClient.CategoriesAPI.CategoriesCategoryIdGet(context.Background(), categoryId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CategoriesAPIService CategoriesCategoryIdPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var categoryId string

		resp, httpRes, err := apiClient.CategoriesAPI.CategoriesCategoryIdPut(context.Background(), categoryId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CategoriesAPIService CategoriesGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CategoriesAPI.CategoriesGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CategoriesAPIService CategoriesPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CategoriesAPI.CategoriesPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}