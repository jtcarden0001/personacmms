# \TimeTriggersAPI

All URIs are relative to *http://127.0.0.1:8080/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssetsAssetIdTasksTaskIdTimeTriggersGet**](TimeTriggersAPI.md#AssetsAssetIdTasksTaskIdTimeTriggersGet) | **Get** /assets/{assetId}/tasks/{taskId}/time-triggers | List time triggers
[**AssetsAssetIdTasksTaskIdTimeTriggersPost**](TimeTriggersAPI.md#AssetsAssetIdTasksTaskIdTimeTriggersPost) | **Post** /assets/{assetId}/tasks/{taskId}/time-triggers | Create a time trigger
[**AssetsAssetIdTasksTaskIdTimeTriggersTimeTriggerIdDelete**](TimeTriggersAPI.md#AssetsAssetIdTasksTaskIdTimeTriggersTimeTriggerIdDelete) | **Delete** /assets/{assetId}/tasks/{taskId}/time-triggers/{timeTriggerId} | Delete a time trigger
[**AssetsAssetIdTasksTaskIdTimeTriggersTimeTriggerIdGet**](TimeTriggersAPI.md#AssetsAssetIdTasksTaskIdTimeTriggersTimeTriggerIdGet) | **Get** /assets/{assetId}/tasks/{taskId}/time-triggers/{timeTriggerId} | Get a time trigger
[**AssetsAssetIdTasksTaskIdTimeTriggersTimeTriggerIdPut**](TimeTriggersAPI.md#AssetsAssetIdTasksTaskIdTimeTriggersTimeTriggerIdPut) | **Put** /assets/{assetId}/tasks/{taskId}/time-triggers/{timeTriggerId} | Update a time trigger



## AssetsAssetIdTasksTaskIdTimeTriggersGet

> []TypesTimeTrigger AssetsAssetIdTasksTaskIdTimeTriggersGet(ctx, assetId, taskId).Execute()

List time triggers



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
	resp, r, err := apiClient.TimeTriggersAPI.AssetsAssetIdTasksTaskIdTimeTriggersGet(context.Background(), assetId, taskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimeTriggersAPI.AssetsAssetIdTasksTaskIdTimeTriggersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssetsAssetIdTasksTaskIdTimeTriggersGet`: []TypesTimeTrigger
	fmt.Fprintf(os.Stdout, "Response from `TimeTriggersAPI.AssetsAssetIdTasksTaskIdTimeTriggersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetId** | **string** | Asset Id | 
**taskId** | **string** | Asset Task Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetsAssetIdTasksTaskIdTimeTriggersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]TypesTimeTrigger**](TypesTimeTrigger.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetsAssetIdTasksTaskIdTimeTriggersPost

> TypesTimeTrigger AssetsAssetIdTasksTaskIdTimeTriggersPost(ctx, assetId, taskId).TypesTimeTrigger(typesTimeTrigger).Execute()

Create a time trigger



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
	typesTimeTrigger := *openapiclient.NewTypesTimeTrigger(int32(123), "TimeUnit_example") // TypesTimeTrigger | Time Trigger object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TimeTriggersAPI.AssetsAssetIdTasksTaskIdTimeTriggersPost(context.Background(), assetId, taskId).TypesTimeTrigger(typesTimeTrigger).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimeTriggersAPI.AssetsAssetIdTasksTaskIdTimeTriggersPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssetsAssetIdTasksTaskIdTimeTriggersPost`: TypesTimeTrigger
	fmt.Fprintf(os.Stdout, "Response from `TimeTriggersAPI.AssetsAssetIdTasksTaskIdTimeTriggersPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetId** | **string** | Asset Id | 
**taskId** | **string** | Asset Task Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetsAssetIdTasksTaskIdTimeTriggersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **typesTimeTrigger** | [**TypesTimeTrigger**](TypesTimeTrigger.md) | Time Trigger object | 

### Return type

[**TypesTimeTrigger**](TypesTimeTrigger.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetsAssetIdTasksTaskIdTimeTriggersTimeTriggerIdDelete

> AssetsAssetIdTasksTaskIdTimeTriggersTimeTriggerIdDelete(ctx, assetId, taskId, timeTriggerId).Execute()

Delete a time trigger



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
	timeTriggerId := "timeTriggerId_example" // string | Time Trigger Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TimeTriggersAPI.AssetsAssetIdTasksTaskIdTimeTriggersTimeTriggerIdDelete(context.Background(), assetId, taskId, timeTriggerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimeTriggersAPI.AssetsAssetIdTasksTaskIdTimeTriggersTimeTriggerIdDelete``: %v\n", err)
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
**timeTriggerId** | **string** | Time Trigger Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetsAssetIdTasksTaskIdTimeTriggersTimeTriggerIdDeleteRequest struct via the builder pattern


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


## AssetsAssetIdTasksTaskIdTimeTriggersTimeTriggerIdGet

> TypesTimeTrigger AssetsAssetIdTasksTaskIdTimeTriggersTimeTriggerIdGet(ctx, assetId, taskId, timeTriggerId).Execute()

Get a time trigger



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
	timeTriggerId := "timeTriggerId_example" // string | Time Trigger Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TimeTriggersAPI.AssetsAssetIdTasksTaskIdTimeTriggersTimeTriggerIdGet(context.Background(), assetId, taskId, timeTriggerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimeTriggersAPI.AssetsAssetIdTasksTaskIdTimeTriggersTimeTriggerIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssetsAssetIdTasksTaskIdTimeTriggersTimeTriggerIdGet`: TypesTimeTrigger
	fmt.Fprintf(os.Stdout, "Response from `TimeTriggersAPI.AssetsAssetIdTasksTaskIdTimeTriggersTimeTriggerIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetId** | **string** | Asset Id | 
**taskId** | **string** | Asset Task Id | 
**timeTriggerId** | **string** | Time Trigger Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetsAssetIdTasksTaskIdTimeTriggersTimeTriggerIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**TypesTimeTrigger**](TypesTimeTrigger.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetsAssetIdTasksTaskIdTimeTriggersTimeTriggerIdPut

> TypesTimeTrigger AssetsAssetIdTasksTaskIdTimeTriggersTimeTriggerIdPut(ctx, assetId, taskId, timeTriggerId).TypesTimeTrigger(typesTimeTrigger).Execute()

Update a time trigger



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
	timeTriggerId := "timeTriggerId_example" // string | Time Trigger Id
	typesTimeTrigger := *openapiclient.NewTypesTimeTrigger(int32(123), "TimeUnit_example") // TypesTimeTrigger | Time Trigger object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TimeTriggersAPI.AssetsAssetIdTasksTaskIdTimeTriggersTimeTriggerIdPut(context.Background(), assetId, taskId, timeTriggerId).TypesTimeTrigger(typesTimeTrigger).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimeTriggersAPI.AssetsAssetIdTasksTaskIdTimeTriggersTimeTriggerIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssetsAssetIdTasksTaskIdTimeTriggersTimeTriggerIdPut`: TypesTimeTrigger
	fmt.Fprintf(os.Stdout, "Response from `TimeTriggersAPI.AssetsAssetIdTasksTaskIdTimeTriggersTimeTriggerIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetId** | **string** | Asset Id | 
**taskId** | **string** | Task Id | 
**timeTriggerId** | **string** | Time Trigger Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetsAssetIdTasksTaskIdTimeTriggersTimeTriggerIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **typesTimeTrigger** | [**TypesTimeTrigger**](TypesTimeTrigger.md) | Time Trigger object | 

### Return type

[**TypesTimeTrigger**](TypesTimeTrigger.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

