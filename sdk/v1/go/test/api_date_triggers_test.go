/*
PersonaCMMS API

Testing DateTriggersAPIService

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

func Test_openapi_DateTriggersAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DateTriggersAPIService AssetsAssetIdTasksTaskIdDateTriggersDateTriggerIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var assetId string
		var taskId string
		var dateTriggerId string

		httpRes, err := apiClient.DateTriggersAPI.AssetsAssetIdTasksTaskIdDateTriggersDateTriggerIdDelete(context.Background(), assetId, taskId, dateTriggerId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DateTriggersAPIService AssetsAssetIdTasksTaskIdDateTriggersDateTriggerIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var assetId string
		var taskId string
		var dateTriggerId string

		resp, httpRes, err := apiClient.DateTriggersAPI.AssetsAssetIdTasksTaskIdDateTriggersDateTriggerIdGet(context.Background(), assetId, taskId, dateTriggerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DateTriggersAPIService AssetsAssetIdTasksTaskIdDateTriggersDateTriggerIdPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var assetId string
		var taskId string
		var dateTriggerId string

		resp, httpRes, err := apiClient.DateTriggersAPI.AssetsAssetIdTasksTaskIdDateTriggersDateTriggerIdPut(context.Background(), assetId, taskId, dateTriggerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DateTriggersAPIService AssetsAssetIdTasksTaskIdDateTriggersGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var assetId string
		var taskId string

		resp, httpRes, err := apiClient.DateTriggersAPI.AssetsAssetIdTasksTaskIdDateTriggersGet(context.Background(), assetId, taskId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DateTriggersAPIService AssetsAssetIdTasksTaskIdDateTriggersPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var assetId string
		var taskId string

		resp, httpRes, err := apiClient.DateTriggersAPI.AssetsAssetIdTasksTaskIdDateTriggersPost(context.Background(), assetId, taskId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
