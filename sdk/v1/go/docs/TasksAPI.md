# \TasksAPI

All URIs are relative to *http://127.0.0.1:8080/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssetsAssetIdTasksGet**](TasksAPI.md#AssetsAssetIdTasksGet) | **Get** /assets/{assetId}/tasks | List tasks by asset
[**AssetsAssetIdTasksPost**](TasksAPI.md#AssetsAssetIdTasksPost) | **Post** /assets/{assetId}/tasks | Create a task
[**AssetsAssetIdTasksTaskIdDelete**](TasksAPI.md#AssetsAssetIdTasksTaskIdDelete) | **Delete** /assets/{assetId}/tasks/{taskId} | Delete a task
[**AssetsAssetIdTasksTaskIdGet**](TasksAPI.md#AssetsAssetIdTasksTaskIdGet) | **Get** /assets/{assetId}/tasks/{taskId} | Get a task
[**AssetsAssetIdTasksTaskIdPut**](TasksAPI.md#AssetsAssetIdTasksTaskIdPut) | **Put** /assets/{assetId}/tasks/{taskId} | Update a task
[**AssetsAssetIdWorkOrdersWorkOrderIdTasksDelete**](TasksAPI.md#AssetsAssetIdWorkOrdersWorkOrderIdTasksDelete) | **Delete** /assets/{assetId}/work-orders/{workOrderId}/tasks | Disassociate a task with a work order



## AssetsAssetIdTasksGet

> []TypesTask AssetsAssetIdTasksGet(ctx, assetId).Execute()

List tasks by asset



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksAPI.AssetsAssetIdTasksGet(context.Background(), assetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.AssetsAssetIdTasksGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssetsAssetIdTasksGet`: []TypesTask
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.AssetsAssetIdTasksGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetId** | **string** | Asset Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetsAssetIdTasksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]TypesTask**](TypesTask.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetsAssetIdTasksPost

> TypesTask AssetsAssetIdTasksPost(ctx, assetId).TypesTask(typesTask).Execute()

Create a task



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
	typesTask := *openapiclient.NewTypesTask("Title_example") // TypesTask | Task object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksAPI.AssetsAssetIdTasksPost(context.Background(), assetId).TypesTask(typesTask).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.AssetsAssetIdTasksPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssetsAssetIdTasksPost`: TypesTask
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.AssetsAssetIdTasksPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetId** | **string** | Asset Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetsAssetIdTasksPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **typesTask** | [**TypesTask**](TypesTask.md) | Task object | 

### Return type

[**TypesTask**](TypesTask.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetsAssetIdTasksTaskIdDelete

> AssetsAssetIdTasksTaskIdDelete(ctx, assetId, taskId).Execute()

Delete a task



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
	r, err := apiClient.TasksAPI.AssetsAssetIdTasksTaskIdDelete(context.Background(), assetId, taskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.AssetsAssetIdTasksTaskIdDelete``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiAssetsAssetIdTasksTaskIdDeleteRequest struct via the builder pattern


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


## AssetsAssetIdTasksTaskIdGet

> TypesTask AssetsAssetIdTasksTaskIdGet(ctx, assetId, taskId).Execute()

Get a task



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
	resp, r, err := apiClient.TasksAPI.AssetsAssetIdTasksTaskIdGet(context.Background(), assetId, taskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.AssetsAssetIdTasksTaskIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssetsAssetIdTasksTaskIdGet`: TypesTask
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.AssetsAssetIdTasksTaskIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetId** | **string** | Asset Id | 
**taskId** | **string** | Task Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetsAssetIdTasksTaskIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TypesTask**](TypesTask.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetsAssetIdTasksTaskIdPut

> TypesTask AssetsAssetIdTasksTaskIdPut(ctx, assetId, taskId).TypesTask(typesTask).Execute()

Update a task



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
	typesTask := *openapiclient.NewTypesTask("Title_example") // TypesTask | Task object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksAPI.AssetsAssetIdTasksTaskIdPut(context.Background(), assetId, taskId).TypesTask(typesTask).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.AssetsAssetIdTasksTaskIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssetsAssetIdTasksTaskIdPut`: TypesTask
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.AssetsAssetIdTasksTaskIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetId** | **string** | Asset Id | 
**taskId** | **string** | Task Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetsAssetIdTasksTaskIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **typesTask** | [**TypesTask**](TypesTask.md) | Task object | 

### Return type

[**TypesTask**](TypesTask.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetsAssetIdWorkOrdersWorkOrderIdTasksDelete

> AssetsAssetIdWorkOrdersWorkOrderIdTasksDelete(ctx, assetId, workOrderId).Execute()

Disassociate a task with a work order



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
	workOrderId := "workOrderId_example" // string | Work Order Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TasksAPI.AssetsAssetIdWorkOrdersWorkOrderIdTasksDelete(context.Background(), assetId, workOrderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.AssetsAssetIdWorkOrdersWorkOrderIdTasksDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetId** | **string** | Asset Id | 
**workOrderId** | **string** | Work Order Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetsAssetIdWorkOrdersWorkOrderIdTasksDeleteRequest struct via the builder pattern


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

