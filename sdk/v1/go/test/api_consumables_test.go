/*
PersonaCMMS API

Testing ConsumablesAPIService

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

func Test_openapi_ConsumablesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ConsumablesAPIService AssetsAssetIdTasksTaskIdConsumablesConsumableIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var assetId string
		var taskId string
		var consumableId string

		httpRes, err := apiClient.ConsumablesAPI.AssetsAssetIdTasksTaskIdConsumablesConsumableIdDelete(context.Background(), assetId, taskId, consumableId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConsumablesAPIService AssetsAssetIdTasksTaskIdConsumablesConsumableIdPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var assetId string
		var taskId string
		var consumableId string

		resp, httpRes, err := apiClient.ConsumablesAPI.AssetsAssetIdTasksTaskIdConsumablesConsumableIdPut(context.Background(), assetId, taskId, consumableId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConsumablesAPIService AssetsAssetIdWorkOrdersWorkOrderIdConsumablesConsumableIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var assetId string
		var workOrderId string
		var consumableId string

		httpRes, err := apiClient.ConsumablesAPI.AssetsAssetIdWorkOrdersWorkOrderIdConsumablesConsumableIdDelete(context.Background(), assetId, workOrderId, consumableId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConsumablesAPIService AssetsAssetIdWorkOrdersWorkOrderIdConsumablesConsumableIdPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var assetId string
		var workOrderId string
		var consumableId string

		resp, httpRes, err := apiClient.ConsumablesAPI.AssetsAssetIdWorkOrdersWorkOrderIdConsumablesConsumableIdPut(context.Background(), assetId, workOrderId, consumableId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConsumablesAPIService ConsumablesConsumableIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var consumableId string

		httpRes, err := apiClient.ConsumablesAPI.ConsumablesConsumableIdDelete(context.Background(), consumableId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConsumablesAPIService ConsumablesConsumableIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var consumableId string

		resp, httpRes, err := apiClient.ConsumablesAPI.ConsumablesConsumableIdGet(context.Background(), consumableId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConsumablesAPIService ConsumablesConsumableIdPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var consumableId string

		resp, httpRes, err := apiClient.ConsumablesAPI.ConsumablesConsumableIdPut(context.Background(), consumableId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConsumablesAPIService ConsumablesGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ConsumablesAPI.ConsumablesGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConsumablesAPIService ConsumablesPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ConsumablesAPI.ConsumablesPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}