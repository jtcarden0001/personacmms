/*
PersonaCMMS API

Testing TasksAPIService

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

func Test_openapi_TasksAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TasksAPIService AssetsAssetIdTasksGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var assetId string

		resp, httpRes, err := apiClient.TasksAPI.AssetsAssetIdTasksGet(context.Background(), assetId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TasksAPIService AssetsAssetIdTasksPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var assetId string

		resp, httpRes, err := apiClient.TasksAPI.AssetsAssetIdTasksPost(context.Background(), assetId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TasksAPIService AssetsAssetIdTasksTaskIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var assetId string
		var taskId string

		httpRes, err := apiClient.TasksAPI.AssetsAssetIdTasksTaskIdDelete(context.Background(), assetId, taskId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TasksAPIService AssetsAssetIdTasksTaskIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var assetId string
		var taskId string

		resp, httpRes, err := apiClient.TasksAPI.AssetsAssetIdTasksTaskIdGet(context.Background(), assetId, taskId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TasksAPIService AssetsAssetIdTasksTaskIdPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var assetId string
		var taskId string

		resp, httpRes, err := apiClient.TasksAPI.AssetsAssetIdTasksTaskIdPut(context.Background(), assetId, taskId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TasksAPIService AssetsAssetIdWorkOrdersWorkOrderIdTasksDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var assetId string
		var workOrderId string

		httpRes, err := apiClient.TasksAPI.AssetsAssetIdWorkOrdersWorkOrderIdTasksDelete(context.Background(), assetId, workOrderId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
