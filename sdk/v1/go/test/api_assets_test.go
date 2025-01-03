/*
PersonaCMMS API

Testing AssetsAPIService

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

func Test_openapi_AssetsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AssetsAPIService AssetsAssetIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var assetId string

		httpRes, err := apiClient.AssetsAPI.AssetsAssetIdDelete(context.Background(), assetId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AssetsAPIService AssetsAssetIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var assetId string

		resp, httpRes, err := apiClient.AssetsAPI.AssetsAssetIdGet(context.Background(), assetId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AssetsAPIService AssetsAssetIdPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var assetId string

		resp, httpRes, err := apiClient.AssetsAPI.AssetsAssetIdPut(context.Background(), assetId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AssetsAPIService AssetsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AssetsAPI.AssetsGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AssetsAPIService AssetsPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AssetsAPI.AssetsPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AssetsAPIService CategoriesCategoryIdAssetsAssetIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var categoryId string
		var assetId string

		httpRes, err := apiClient.AssetsAPI.CategoriesCategoryIdAssetsAssetIdDelete(context.Background(), categoryId, assetId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AssetsAPIService CategoriesCategoryIdAssetsAssetIdPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var categoryId string
		var assetId string

		resp, httpRes, err := apiClient.AssetsAPI.CategoriesCategoryIdAssetsAssetIdPut(context.Background(), categoryId, assetId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AssetsAPIService CategoriesCategoryIdAssetsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var categoryId string

		resp, httpRes, err := apiClient.AssetsAPI.CategoriesCategoryIdAssetsGet(context.Background(), categoryId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AssetsAPIService CategoriesCategoryIdGroupsGroupIdAssetsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var categoryId string
		var groupId string

		resp, httpRes, err := apiClient.AssetsAPI.CategoriesCategoryIdGroupsGroupIdAssetsGet(context.Background(), categoryId, groupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AssetsAPIService GroupsGroupIdAssetsAssetIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var groupId string
		var assetId string

		httpRes, err := apiClient.AssetsAPI.GroupsGroupIdAssetsAssetIdDelete(context.Background(), groupId, assetId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AssetsAPIService GroupsGroupIdAssetsAssetIdPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var groupId string
		var assetId string

		resp, httpRes, err := apiClient.AssetsAPI.GroupsGroupIdAssetsAssetIdPut(context.Background(), groupId, assetId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AssetsAPIService GroupsGroupIdAssetsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var groupId string

		resp, httpRes, err := apiClient.AssetsAPI.GroupsGroupIdAssetsGet(context.Background(), groupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
