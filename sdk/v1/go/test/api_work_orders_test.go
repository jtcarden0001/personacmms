/*
PersonaCMMS API

Testing WorkOrdersAPIService

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

func Test_openapi_WorkOrdersAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test WorkOrdersAPIService AssetsAssetIdTasksTaskIdWorkOrdersWorkOrderIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var assetId string
		var taskId string
		var workOrderId string

		httpRes, err := apiClient.WorkOrdersAPI.AssetsAssetIdTasksTaskIdWorkOrdersWorkOrderIdDelete(context.Background(), assetId, taskId, workOrderId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WorkOrdersAPIService AssetsAssetIdTasksTaskIdWorkOrdersWorkOrderIdPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var assetId string
		var taskId string
		var workOrderId string

		resp, httpRes, err := apiClient.WorkOrdersAPI.AssetsAssetIdTasksTaskIdWorkOrdersWorkOrderIdPut(context.Background(), assetId, taskId, workOrderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WorkOrdersAPIService AssetsAssetIdWorkOrdersGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var assetId string
		var taskId string

		resp, httpRes, err := apiClient.WorkOrdersAPI.AssetsAssetIdWorkOrdersGet(context.Background(), assetId, taskId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WorkOrdersAPIService AssetsAssetIdWorkOrdersPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var assetId string
		var taskId string

		resp, httpRes, err := apiClient.WorkOrdersAPI.AssetsAssetIdWorkOrdersPost(context.Background(), assetId, taskId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WorkOrdersAPIService AssetsAssetIdWorkOrdersWorkOrderIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var assetId string
		var taskId string
		var workOrderId string

		httpRes, err := apiClient.WorkOrdersAPI.AssetsAssetIdWorkOrdersWorkOrderIdDelete(context.Background(), assetId, taskId, workOrderId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WorkOrdersAPIService AssetsAssetIdWorkOrdersWorkOrderIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var assetId string
		var taskId string
		var workOrderId string

		resp, httpRes, err := apiClient.WorkOrdersAPI.AssetsAssetIdWorkOrdersWorkOrderIdGet(context.Background(), assetId, taskId, workOrderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WorkOrdersAPIService AssetsAssetIdWorkOrdersWorkOrderIdPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var assetId string
		var taskId string
		var workOrderId string

		resp, httpRes, err := apiClient.WorkOrdersAPI.AssetsAssetIdWorkOrdersWorkOrderIdPut(context.Background(), assetId, taskId, workOrderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
