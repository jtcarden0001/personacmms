# \WorkOrdersAPI

All URIs are relative to *http://127.0.0.1:8080/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssetsAssetIdTasksTaskIdWorkOrdersWorkOrderIdDelete**](WorkOrdersAPI.md#AssetsAssetIdTasksTaskIdWorkOrdersWorkOrderIdDelete) | **Delete** /assets/{assetId}/tasks/{taskId}/work-orders/{workOrderId} | Disassociate a work order with a task
[**AssetsAssetIdTasksTaskIdWorkOrdersWorkOrderIdPut**](WorkOrdersAPI.md#AssetsAssetIdTasksTaskIdWorkOrdersWorkOrderIdPut) | **Put** /assets/{assetId}/tasks/{taskId}/work-orders/{workOrderId} | Associate a work order with a task
[**AssetsAssetIdWorkOrdersGet**](WorkOrdersAPI.md#AssetsAssetIdWorkOrdersGet) | **Get** /assets/{assetId}/work-orders | List asset task work orders
[**AssetsAssetIdWorkOrdersPost**](WorkOrdersAPI.md#AssetsAssetIdWorkOrdersPost) | **Post** /assets/{assetId}/work-orders | Create a work order
[**AssetsAssetIdWorkOrdersWorkOrderIdDelete**](WorkOrdersAPI.md#AssetsAssetIdWorkOrdersWorkOrderIdDelete) | **Delete** /assets/{assetId}/work-orders/{workOrderId} | Delete an asset task work order
[**AssetsAssetIdWorkOrdersWorkOrderIdGet**](WorkOrdersAPI.md#AssetsAssetIdWorkOrdersWorkOrderIdGet) | **Get** /assets/{assetId}/work-orders/{workOrderId} | Get an asset task work order
[**AssetsAssetIdWorkOrdersWorkOrderIdPut**](WorkOrdersAPI.md#AssetsAssetIdWorkOrdersWorkOrderIdPut) | **Put** /assets/{assetId}/work-orders/{workOrderId} | Update an asset task work order



## AssetsAssetIdTasksTaskIdWorkOrdersWorkOrderIdDelete

> AssetsAssetIdTasksTaskIdWorkOrdersWorkOrderIdDelete(ctx, assetId, taskId, workOrderId).Execute()

Disassociate a work order with a task



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
	workOrderId := "workOrderId_example" // string | Work Order Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkOrdersAPI.AssetsAssetIdTasksTaskIdWorkOrdersWorkOrderIdDelete(context.Background(), assetId, taskId, workOrderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkOrdersAPI.AssetsAssetIdTasksTaskIdWorkOrdersWorkOrderIdDelete``: %v\n", err)
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
**workOrderId** | **string** | Work Order Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetsAssetIdTasksTaskIdWorkOrdersWorkOrderIdDeleteRequest struct via the builder pattern


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


## AssetsAssetIdTasksTaskIdWorkOrdersWorkOrderIdPut

> TypesWorkOrder AssetsAssetIdTasksTaskIdWorkOrdersWorkOrderIdPut(ctx, assetId, taskId, workOrderId).Execute()

Associate a work order with a task



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
	assetId := "assetId_example" // string | Asset ID
	taskId := "taskId_example" // string | Task ID
	workOrderId := "workOrderId_example" // string | Work Order ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkOrdersAPI.AssetsAssetIdTasksTaskIdWorkOrdersWorkOrderIdPut(context.Background(), assetId, taskId, workOrderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkOrdersAPI.AssetsAssetIdTasksTaskIdWorkOrdersWorkOrderIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssetsAssetIdTasksTaskIdWorkOrdersWorkOrderIdPut`: TypesWorkOrder
	fmt.Fprintf(os.Stdout, "Response from `WorkOrdersAPI.AssetsAssetIdTasksTaskIdWorkOrdersWorkOrderIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetId** | **string** | Asset ID | 
**taskId** | **string** | Task ID | 
**workOrderId** | **string** | Work Order ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetsAssetIdTasksTaskIdWorkOrdersWorkOrderIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**TypesWorkOrder**](TypesWorkOrder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetsAssetIdWorkOrdersGet

> []TypesWorkOrder AssetsAssetIdWorkOrdersGet(ctx, assetId, taskId).Execute()

List asset task work orders



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
	resp, r, err := apiClient.WorkOrdersAPI.AssetsAssetIdWorkOrdersGet(context.Background(), assetId, taskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkOrdersAPI.AssetsAssetIdWorkOrdersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssetsAssetIdWorkOrdersGet`: []TypesWorkOrder
	fmt.Fprintf(os.Stdout, "Response from `WorkOrdersAPI.AssetsAssetIdWorkOrdersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetId** | **string** | Asset Id | 
**taskId** | **string** | Asset Task Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetsAssetIdWorkOrdersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]TypesWorkOrder**](TypesWorkOrder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetsAssetIdWorkOrdersPost

> TypesWorkOrder AssetsAssetIdWorkOrdersPost(ctx, assetId, taskId).TypesWorkOrder(typesWorkOrder).Execute()

Create a work order



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
	typesWorkOrder := *openapiclient.NewTypesWorkOrder("Status_example", "Title_example") // TypesWorkOrder | Work Order object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkOrdersAPI.AssetsAssetIdWorkOrdersPost(context.Background(), assetId, taskId).TypesWorkOrder(typesWorkOrder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkOrdersAPI.AssetsAssetIdWorkOrdersPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssetsAssetIdWorkOrdersPost`: TypesWorkOrder
	fmt.Fprintf(os.Stdout, "Response from `WorkOrdersAPI.AssetsAssetIdWorkOrdersPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetId** | **string** | Asset Id | 
**taskId** | **string** | Asset Task Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetsAssetIdWorkOrdersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **typesWorkOrder** | [**TypesWorkOrder**](TypesWorkOrder.md) | Work Order object | 

### Return type

[**TypesWorkOrder**](TypesWorkOrder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetsAssetIdWorkOrdersWorkOrderIdDelete

> AssetsAssetIdWorkOrdersWorkOrderIdDelete(ctx, assetId, taskId, workOrderId).Execute()

Delete an asset task work order



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
	workOrderId := "workOrderId_example" // string | Work Order Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkOrdersAPI.AssetsAssetIdWorkOrdersWorkOrderIdDelete(context.Background(), assetId, taskId, workOrderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkOrdersAPI.AssetsAssetIdWorkOrdersWorkOrderIdDelete``: %v\n", err)
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
**workOrderId** | **string** | Work Order Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetsAssetIdWorkOrdersWorkOrderIdDeleteRequest struct via the builder pattern


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


## AssetsAssetIdWorkOrdersWorkOrderIdGet

> TypesWorkOrder AssetsAssetIdWorkOrdersWorkOrderIdGet(ctx, assetId, taskId, workOrderId).Execute()

Get an asset task work order



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
	workOrderId := "workOrderId_example" // string | Work Order Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkOrdersAPI.AssetsAssetIdWorkOrdersWorkOrderIdGet(context.Background(), assetId, taskId, workOrderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkOrdersAPI.AssetsAssetIdWorkOrdersWorkOrderIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssetsAssetIdWorkOrdersWorkOrderIdGet`: TypesWorkOrder
	fmt.Fprintf(os.Stdout, "Response from `WorkOrdersAPI.AssetsAssetIdWorkOrdersWorkOrderIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetId** | **string** | Asset Id | 
**taskId** | **string** | Asset Task Id | 
**workOrderId** | **string** | Work Order Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetsAssetIdWorkOrdersWorkOrderIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**TypesWorkOrder**](TypesWorkOrder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetsAssetIdWorkOrdersWorkOrderIdPut

> TypesWorkOrder AssetsAssetIdWorkOrdersWorkOrderIdPut(ctx, assetId, taskId, workOrderId).TypesWorkOrder(typesWorkOrder).Execute()

Update an asset task work order



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
	workOrderId := "workOrderId_example" // string | Work Order Id
	typesWorkOrder := *openapiclient.NewTypesWorkOrder("Status_example", "Title_example") // TypesWorkOrder | Work Order object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkOrdersAPI.AssetsAssetIdWorkOrdersWorkOrderIdPut(context.Background(), assetId, taskId, workOrderId).TypesWorkOrder(typesWorkOrder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkOrdersAPI.AssetsAssetIdWorkOrdersWorkOrderIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssetsAssetIdWorkOrdersWorkOrderIdPut`: TypesWorkOrder
	fmt.Fprintf(os.Stdout, "Response from `WorkOrdersAPI.AssetsAssetIdWorkOrdersWorkOrderIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetId** | **string** | Asset Id | 
**taskId** | **string** | Asset Task Id | 
**workOrderId** | **string** | Work Order Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetsAssetIdWorkOrdersWorkOrderIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **typesWorkOrder** | [**TypesWorkOrder**](TypesWorkOrder.md) | Work Order object | 

### Return type

[**TypesWorkOrder**](TypesWorkOrder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

