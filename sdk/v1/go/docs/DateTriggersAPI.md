# \DateTriggersAPI

All URIs are relative to *http://127.0.0.1:8080/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssetsAssetIdTasksTaskIdDateTriggersDateTriggerIdDelete**](DateTriggersAPI.md#AssetsAssetIdTasksTaskIdDateTriggersDateTriggerIdDelete) | **Delete** /assets/{assetId}/tasks/{taskId}/date-triggers/{dateTriggerId} | Delete a date trigger
[**AssetsAssetIdTasksTaskIdDateTriggersDateTriggerIdGet**](DateTriggersAPI.md#AssetsAssetIdTasksTaskIdDateTriggersDateTriggerIdGet) | **Get** /assets/{assetId}/tasks/{taskId}/date-triggers/{dateTriggerId} | Get a date trigger
[**AssetsAssetIdTasksTaskIdDateTriggersDateTriggerIdPut**](DateTriggersAPI.md#AssetsAssetIdTasksTaskIdDateTriggersDateTriggerIdPut) | **Put** /assets/{assetId}/tasks/{taskId}/date-triggers/{dateTriggerId} | Update a date trigger
[**AssetsAssetIdTasksTaskIdDateTriggersGet**](DateTriggersAPI.md#AssetsAssetIdTasksTaskIdDateTriggersGet) | **Get** /assets/{assetId}/tasks/{taskId}/date-triggers | List date triggers
[**AssetsAssetIdTasksTaskIdDateTriggersPost**](DateTriggersAPI.md#AssetsAssetIdTasksTaskIdDateTriggersPost) | **Post** /assets/{assetId}/tasks/{taskId}/date-triggers | Create a date trigger



## AssetsAssetIdTasksTaskIdDateTriggersDateTriggerIdDelete

> AssetsAssetIdTasksTaskIdDateTriggersDateTriggerIdDelete(ctx, assetId, taskId, dateTriggerId).Execute()

Delete a date trigger



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
	taskId := "taskId_example" // string | Task Id
	dateTriggerId := "dateTriggerId_example" // string | Date Trigger Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DateTriggersAPI.AssetsAssetIdTasksTaskIdDateTriggersDateTriggerIdDelete(context.Background(), assetId, taskId, dateTriggerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DateTriggersAPI.AssetsAssetIdTasksTaskIdDateTriggersDateTriggerIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetId** | **string** | Asset Id | 
**taskId** | **string** | Task Id | 
**dateTriggerId** | **string** | Date Trigger Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetsAssetIdTasksTaskIdDateTriggersDateTriggerIdDeleteRequest struct via the builder pattern


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


## AssetsAssetIdTasksTaskIdDateTriggersDateTriggerIdGet

> TypesDateTrigger AssetsAssetIdTasksTaskIdDateTriggersDateTriggerIdGet(ctx, assetId, taskId, dateTriggerId).Execute()

Get a date trigger



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
	taskId := "taskId_example" // string | Task Id
	dateTriggerId := "dateTriggerId_example" // string | Date Trigger Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DateTriggersAPI.AssetsAssetIdTasksTaskIdDateTriggersDateTriggerIdGet(context.Background(), assetId, taskId, dateTriggerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DateTriggersAPI.AssetsAssetIdTasksTaskIdDateTriggersDateTriggerIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssetsAssetIdTasksTaskIdDateTriggersDateTriggerIdGet`: TypesDateTrigger
	fmt.Fprintf(os.Stdout, "Response from `DateTriggersAPI.AssetsAssetIdTasksTaskIdDateTriggersDateTriggerIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetId** | **string** | Asset Id | 
**taskId** | **string** | Task Id | 
**dateTriggerId** | **string** | Date Trigger Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetsAssetIdTasksTaskIdDateTriggersDateTriggerIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**TypesDateTrigger**](TypesDateTrigger.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetsAssetIdTasksTaskIdDateTriggersDateTriggerIdPut

> TypesDateTrigger AssetsAssetIdTasksTaskIdDateTriggersDateTriggerIdPut(ctx, assetId, taskId, dateTriggerId).TypesDateTrigger(typesDateTrigger).Execute()

Update a date trigger



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
	taskId := "taskId_example" // string | Task Id
	dateTriggerId := "dateTriggerId_example" // string | Date Trigger Id
	typesDateTrigger := *openapiclient.NewTypesDateTrigger() // TypesDateTrigger | Date Trigger object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DateTriggersAPI.AssetsAssetIdTasksTaskIdDateTriggersDateTriggerIdPut(context.Background(), assetId, taskId, dateTriggerId).TypesDateTrigger(typesDateTrigger).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DateTriggersAPI.AssetsAssetIdTasksTaskIdDateTriggersDateTriggerIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssetsAssetIdTasksTaskIdDateTriggersDateTriggerIdPut`: TypesDateTrigger
	fmt.Fprintf(os.Stdout, "Response from `DateTriggersAPI.AssetsAssetIdTasksTaskIdDateTriggersDateTriggerIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetId** | **string** | Asset Id | 
**taskId** | **string** | Task Id | 
**dateTriggerId** | **string** | Date Trigger Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetsAssetIdTasksTaskIdDateTriggersDateTriggerIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **typesDateTrigger** | [**TypesDateTrigger**](TypesDateTrigger.md) | Date Trigger object | 

### Return type

[**TypesDateTrigger**](TypesDateTrigger.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetsAssetIdTasksTaskIdDateTriggersGet

> []TypesDateTrigger AssetsAssetIdTasksTaskIdDateTriggersGet(ctx, assetId, taskId).Execute()

List date triggers



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
	taskId := "taskId_example" // string | Task Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DateTriggersAPI.AssetsAssetIdTasksTaskIdDateTriggersGet(context.Background(), assetId, taskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DateTriggersAPI.AssetsAssetIdTasksTaskIdDateTriggersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssetsAssetIdTasksTaskIdDateTriggersGet`: []TypesDateTrigger
	fmt.Fprintf(os.Stdout, "Response from `DateTriggersAPI.AssetsAssetIdTasksTaskIdDateTriggersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetId** | **string** | Asset Id | 
**taskId** | **string** | Task Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetsAssetIdTasksTaskIdDateTriggersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]TypesDateTrigger**](TypesDateTrigger.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetsAssetIdTasksTaskIdDateTriggersPost

> TypesDateTrigger AssetsAssetIdTasksTaskIdDateTriggersPost(ctx, assetId, taskId).TypesDateTrigger(typesDateTrigger).Execute()

Create a date trigger



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
	taskId := "taskId_example" // string | Task Id
	typesDateTrigger := *openapiclient.NewTypesDateTrigger() // TypesDateTrigger | Date Trigger object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DateTriggersAPI.AssetsAssetIdTasksTaskIdDateTriggersPost(context.Background(), assetId, taskId).TypesDateTrigger(typesDateTrigger).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DateTriggersAPI.AssetsAssetIdTasksTaskIdDateTriggersPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssetsAssetIdTasksTaskIdDateTriggersPost`: TypesDateTrigger
	fmt.Fprintf(os.Stdout, "Response from `DateTriggersAPI.AssetsAssetIdTasksTaskIdDateTriggersPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetId** | **string** | Asset Id | 
**taskId** | **string** | Task Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetsAssetIdTasksTaskIdDateTriggersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **typesDateTrigger** | [**TypesDateTrigger**](TypesDateTrigger.md) | Date Trigger object | 

### Return type

[**TypesDateTrigger**](TypesDateTrigger.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

