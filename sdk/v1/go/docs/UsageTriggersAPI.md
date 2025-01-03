# \UsageTriggersAPI

All URIs are relative to *http://127.0.0.1:8080/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssetsAssetIdTasksTaskIdUsageTriggersGet**](UsageTriggersAPI.md#AssetsAssetIdTasksTaskIdUsageTriggersGet) | **Get** /assets/{assetId}/tasks/{taskId}/usage-triggers | List usage triggers
[**AssetsAssetIdTasksTaskIdUsageTriggersPost**](UsageTriggersAPI.md#AssetsAssetIdTasksTaskIdUsageTriggersPost) | **Post** /assets/{assetId}/tasks/{taskId}/usage-triggers | Create a usage trigger
[**AssetsAssetIdTasksTaskIdUsageTriggersUsageTriggerIdDelete**](UsageTriggersAPI.md#AssetsAssetIdTasksTaskIdUsageTriggersUsageTriggerIdDelete) | **Delete** /assets/{assetId}/tasks/{taskId}/usage-triggers/{usageTriggerId} | Delete a usage trigger
[**AssetsAssetIdTasksTaskIdUsageTriggersUsageTriggerIdGet**](UsageTriggersAPI.md#AssetsAssetIdTasksTaskIdUsageTriggersUsageTriggerIdGet) | **Get** /assets/{assetId}/tasks/{taskId}/usage-triggers/{usageTriggerId} | Get a usage trigger
[**AssetsAssetIdTasksTaskIdUsageTriggersUsageTriggerIdPut**](UsageTriggersAPI.md#AssetsAssetIdTasksTaskIdUsageTriggersUsageTriggerIdPut) | **Put** /assets/{assetId}/tasks/{taskId}/usage-triggers/{usageTriggerId} | Update a usage trigger



## AssetsAssetIdTasksTaskIdUsageTriggersGet

> []TypesUsageTrigger AssetsAssetIdTasksTaskIdUsageTriggersGet(ctx, assetId, taskId).Execute()

List usage triggers



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/jtcarden0001/personacmms/sdk/v1/go"
)

func main() {
	assetId := "assetId_example" // string | Asset Id
	taskId := "taskId_example" // string | Asset Task Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsageTriggersAPI.AssetsAssetIdTasksTaskIdUsageTriggersGet(context.Background(), assetId, taskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageTriggersAPI.AssetsAssetIdTasksTaskIdUsageTriggersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssetsAssetIdTasksTaskIdUsageTriggersGet`: []TypesUsageTrigger
	fmt.Fprintf(os.Stdout, "Response from `UsageTriggersAPI.AssetsAssetIdTasksTaskIdUsageTriggersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetId** | **string** | Asset Id | 
**taskId** | **string** | Asset Task Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetsAssetIdTasksTaskIdUsageTriggersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]TypesUsageTrigger**](TypesUsageTrigger.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetsAssetIdTasksTaskIdUsageTriggersPost

> TypesUsageTrigger AssetsAssetIdTasksTaskIdUsageTriggersPost(ctx, assetId, taskId).TypesUsageTrigger(typesUsageTrigger).Execute()

Create a usage trigger



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/jtcarden0001/personacmms/sdk/v1/go"
)

func main() {
	assetId := "assetId_example" // string | Asset Id
	taskId := "taskId_example" // string | Asset Task Id
	typesUsageTrigger := *openapiclient.NewTypesUsageTrigger(int32(123), "UsageUnit_example") // TypesUsageTrigger | Usage Trigger object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsageTriggersAPI.AssetsAssetIdTasksTaskIdUsageTriggersPost(context.Background(), assetId, taskId).TypesUsageTrigger(typesUsageTrigger).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageTriggersAPI.AssetsAssetIdTasksTaskIdUsageTriggersPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssetsAssetIdTasksTaskIdUsageTriggersPost`: TypesUsageTrigger
	fmt.Fprintf(os.Stdout, "Response from `UsageTriggersAPI.AssetsAssetIdTasksTaskIdUsageTriggersPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetId** | **string** | Asset Id | 
**taskId** | **string** | Asset Task Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetsAssetIdTasksTaskIdUsageTriggersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **typesUsageTrigger** | [**TypesUsageTrigger**](TypesUsageTrigger.md) | Usage Trigger object | 

### Return type

[**TypesUsageTrigger**](TypesUsageTrigger.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetsAssetIdTasksTaskIdUsageTriggersUsageTriggerIdDelete

> AssetsAssetIdTasksTaskIdUsageTriggersUsageTriggerIdDelete(ctx, assetId, taskId, usageTriggerId).Execute()

Delete a usage trigger



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/jtcarden0001/personacmms/sdk/v1/go"
)

func main() {
	assetId := "assetId_example" // string | Asset Id
	taskId := "taskId_example" // string | Asset Task Id
	usageTriggerId := "usageTriggerId_example" // string | Usage Trigger Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsageTriggersAPI.AssetsAssetIdTasksTaskIdUsageTriggersUsageTriggerIdDelete(context.Background(), assetId, taskId, usageTriggerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageTriggersAPI.AssetsAssetIdTasksTaskIdUsageTriggersUsageTriggerIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetId** | **string** | Asset Id | 
**taskId** | **string** | Asset Task Id | 
**usageTriggerId** | **string** | Usage Trigger Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetsAssetIdTasksTaskIdUsageTriggersUsageTriggerIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetsAssetIdTasksTaskIdUsageTriggersUsageTriggerIdGet

> TypesUsageTrigger AssetsAssetIdTasksTaskIdUsageTriggersUsageTriggerIdGet(ctx, assetId, taskId, usageTriggerId).Execute()

Get a usage trigger



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/jtcarden0001/personacmms/sdk/v1/go"
)

func main() {
	assetId := "assetId_example" // string | Asset Id
	taskId := "taskId_example" // string | Asset Task Id
	usageTriggerId := "usageTriggerId_example" // string | Usage Trigger Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsageTriggersAPI.AssetsAssetIdTasksTaskIdUsageTriggersUsageTriggerIdGet(context.Background(), assetId, taskId, usageTriggerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageTriggersAPI.AssetsAssetIdTasksTaskIdUsageTriggersUsageTriggerIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssetsAssetIdTasksTaskIdUsageTriggersUsageTriggerIdGet`: TypesUsageTrigger
	fmt.Fprintf(os.Stdout, "Response from `UsageTriggersAPI.AssetsAssetIdTasksTaskIdUsageTriggersUsageTriggerIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetId** | **string** | Asset Id | 
**taskId** | **string** | Asset Task Id | 
**usageTriggerId** | **string** | Usage Trigger Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetsAssetIdTasksTaskIdUsageTriggersUsageTriggerIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**TypesUsageTrigger**](TypesUsageTrigger.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetsAssetIdTasksTaskIdUsageTriggersUsageTriggerIdPut

> TypesUsageTrigger AssetsAssetIdTasksTaskIdUsageTriggersUsageTriggerIdPut(ctx, assetId, taskId, usageTriggerId).TypesUsageTrigger(typesUsageTrigger).Execute()

Update a usage trigger



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/jtcarden0001/personacmms/sdk/v1/go"
)

func main() {
	assetId := "assetId_example" // string | Asset Id
	taskId := "taskId_example" // string | Asset Task Id
	usageTriggerId := "usageTriggerId_example" // string | Usage Trigger Id
	typesUsageTrigger := *openapiclient.NewTypesUsageTrigger(int32(123), "UsageUnit_example") // TypesUsageTrigger | Usage Trigger object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsageTriggersAPI.AssetsAssetIdTasksTaskIdUsageTriggersUsageTriggerIdPut(context.Background(), assetId, taskId, usageTriggerId).TypesUsageTrigger(typesUsageTrigger).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageTriggersAPI.AssetsAssetIdTasksTaskIdUsageTriggersUsageTriggerIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssetsAssetIdTasksTaskIdUsageTriggersUsageTriggerIdPut`: TypesUsageTrigger
	fmt.Fprintf(os.Stdout, "Response from `UsageTriggersAPI.AssetsAssetIdTasksTaskIdUsageTriggersUsageTriggerIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetId** | **string** | Asset Id | 
**taskId** | **string** | Asset Task Id | 
**usageTriggerId** | **string** | Usage Trigger Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetsAssetIdTasksTaskIdUsageTriggersUsageTriggerIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **typesUsageTrigger** | [**TypesUsageTrigger**](TypesUsageTrigger.md) | Usage Trigger object | 

### Return type

[**TypesUsageTrigger**](TypesUsageTrigger.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

