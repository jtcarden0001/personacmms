# \ConsumablesAPI

All URIs are relative to *http://127.0.0.1:8080/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssetsAssetIdTasksTaskIdConsumablesConsumableIdDelete**](ConsumablesAPI.md#AssetsAssetIdTasksTaskIdConsumablesConsumableIdDelete) | **Delete** /assets/{assetId}/tasks/{taskId}/consumables/{consumableId} | Disassociate a consumable with a task
[**AssetsAssetIdTasksTaskIdConsumablesConsumableIdPut**](ConsumablesAPI.md#AssetsAssetIdTasksTaskIdConsumablesConsumableIdPut) | **Put** /assets/{assetId}/tasks/{taskId}/consumables/{consumableId} | Associate a consumable with a task
[**AssetsAssetIdWorkOrdersWorkOrderIdConsumablesConsumableIdDelete**](ConsumablesAPI.md#AssetsAssetIdWorkOrdersWorkOrderIdConsumablesConsumableIdDelete) | **Delete** /assets/{assetId}/work-orders/{workOrderId}/consumables/{consumableId} | Disassociate a consumable with a work order
[**AssetsAssetIdWorkOrdersWorkOrderIdConsumablesConsumableIdPut**](ConsumablesAPI.md#AssetsAssetIdWorkOrdersWorkOrderIdConsumablesConsumableIdPut) | **Put** /assets/{assetId}/work-orders/{workOrderId}/consumables/{consumableId} | Associate a consumable with a work order
[**ConsumablesConsumableIdDelete**](ConsumablesAPI.md#ConsumablesConsumableIdDelete) | **Delete** /consumables/{consumableId} | Delete a consumable
[**ConsumablesConsumableIdGet**](ConsumablesAPI.md#ConsumablesConsumableIdGet) | **Get** /consumables/{consumableId} | Get a consumable
[**ConsumablesConsumableIdPut**](ConsumablesAPI.md#ConsumablesConsumableIdPut) | **Put** /consumables/{consumableId} | Update a consumable
[**ConsumablesGet**](ConsumablesAPI.md#ConsumablesGet) | **Get** /consumables | List consumables
[**ConsumablesPost**](ConsumablesAPI.md#ConsumablesPost) | **Post** /consumables | Create a consumable



## AssetsAssetIdTasksTaskIdConsumablesConsumableIdDelete

> AssetsAssetIdTasksTaskIdConsumablesConsumableIdDelete(ctx, assetId, taskId, consumableId).Execute()

Disassociate a consumable with a task



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
	consumableId := "consumableId_example" // string | Consumable Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConsumablesAPI.AssetsAssetIdTasksTaskIdConsumablesConsumableIdDelete(context.Background(), assetId, taskId, consumableId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumablesAPI.AssetsAssetIdTasksTaskIdConsumablesConsumableIdDelete``: %v\n", err)
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
**consumableId** | **string** | Consumable Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetsAssetIdTasksTaskIdConsumablesConsumableIdDeleteRequest struct via the builder pattern


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


## AssetsAssetIdTasksTaskIdConsumablesConsumableIdPut

> TypesConsumable AssetsAssetIdTasksTaskIdConsumablesConsumableIdPut(ctx, assetId, taskId, consumableId).TypesConsumableQuantity(typesConsumableQuantity).Execute()

Associate a consumable with a task



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
	consumableId := "consumableId_example" // string | Consumable ID
	typesConsumableQuantity := *openapiclient.NewTypesConsumableQuantity("Quantity_example") // TypesConsumableQuantity | Consumable object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsumablesAPI.AssetsAssetIdTasksTaskIdConsumablesConsumableIdPut(context.Background(), assetId, taskId, consumableId).TypesConsumableQuantity(typesConsumableQuantity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumablesAPI.AssetsAssetIdTasksTaskIdConsumablesConsumableIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssetsAssetIdTasksTaskIdConsumablesConsumableIdPut`: TypesConsumable
	fmt.Fprintf(os.Stdout, "Response from `ConsumablesAPI.AssetsAssetIdTasksTaskIdConsumablesConsumableIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetId** | **string** | Asset ID | 
**taskId** | **string** | Task ID | 
**consumableId** | **string** | Consumable ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetsAssetIdTasksTaskIdConsumablesConsumableIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **typesConsumableQuantity** | [**TypesConsumableQuantity**](TypesConsumableQuantity.md) | Consumable object | 

### Return type

[**TypesConsumable**](TypesConsumable.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetsAssetIdWorkOrdersWorkOrderIdConsumablesConsumableIdDelete

> AssetsAssetIdWorkOrdersWorkOrderIdConsumablesConsumableIdDelete(ctx, assetId, workOrderId, consumableId).Execute()

Disassociate a consumable with a work order



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
	consumableId := "consumableId_example" // string | Consumable Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConsumablesAPI.AssetsAssetIdWorkOrdersWorkOrderIdConsumablesConsumableIdDelete(context.Background(), assetId, workOrderId, consumableId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumablesAPI.AssetsAssetIdWorkOrdersWorkOrderIdConsumablesConsumableIdDelete``: %v\n", err)
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
**consumableId** | **string** | Consumable Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetsAssetIdWorkOrdersWorkOrderIdConsumablesConsumableIdDeleteRequest struct via the builder pattern


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


## AssetsAssetIdWorkOrdersWorkOrderIdConsumablesConsumableIdPut

> TypesConsumable AssetsAssetIdWorkOrdersWorkOrderIdConsumablesConsumableIdPut(ctx, assetId, workOrderId, consumableId).TypesConsumableQuantity(typesConsumableQuantity).Execute()

Associate a consumable with a work order



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
	consumableId := "consumableId_example" // string | Consumable Id
	typesConsumableQuantity := *openapiclient.NewTypesConsumableQuantity("Quantity_example") // TypesConsumableQuantity | Consumable object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsumablesAPI.AssetsAssetIdWorkOrdersWorkOrderIdConsumablesConsumableIdPut(context.Background(), assetId, workOrderId, consumableId).TypesConsumableQuantity(typesConsumableQuantity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumablesAPI.AssetsAssetIdWorkOrdersWorkOrderIdConsumablesConsumableIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssetsAssetIdWorkOrdersWorkOrderIdConsumablesConsumableIdPut`: TypesConsumable
	fmt.Fprintf(os.Stdout, "Response from `ConsumablesAPI.AssetsAssetIdWorkOrdersWorkOrderIdConsumablesConsumableIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetId** | **string** | Asset Id | 
**workOrderId** | **string** | Work Order Id | 
**consumableId** | **string** | Consumable Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetsAssetIdWorkOrdersWorkOrderIdConsumablesConsumableIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **typesConsumableQuantity** | [**TypesConsumableQuantity**](TypesConsumableQuantity.md) | Consumable object | 

### Return type

[**TypesConsumable**](TypesConsumable.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsumablesConsumableIdDelete

> ConsumablesConsumableIdDelete(ctx, consumableId).Execute()

Delete a consumable



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
	consumableId := "consumableId_example" // string | Consumable Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConsumablesAPI.ConsumablesConsumableIdDelete(context.Background(), consumableId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumablesAPI.ConsumablesConsumableIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consumableId** | **string** | Consumable Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsumablesConsumableIdDeleteRequest struct via the builder pattern


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


## ConsumablesConsumableIdGet

> TypesConsumable ConsumablesConsumableIdGet(ctx, consumableId).Execute()

Get a consumable



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
	consumableId := "consumableId_example" // string | Consumable Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsumablesAPI.ConsumablesConsumableIdGet(context.Background(), consumableId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumablesAPI.ConsumablesConsumableIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsumablesConsumableIdGet`: TypesConsumable
	fmt.Fprintf(os.Stdout, "Response from `ConsumablesAPI.ConsumablesConsumableIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consumableId** | **string** | Consumable Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsumablesConsumableIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TypesConsumable**](TypesConsumable.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsumablesConsumableIdPut

> TypesConsumable ConsumablesConsumableIdPut(ctx, consumableId).TypesConsumable(typesConsumable).Execute()

Update a consumable



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
	consumableId := "consumableId_example" // string | Consumable Id
	typesConsumable := *openapiclient.NewTypesConsumable("Title_example") // TypesConsumable | Consumable object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsumablesAPI.ConsumablesConsumableIdPut(context.Background(), consumableId).TypesConsumable(typesConsumable).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumablesAPI.ConsumablesConsumableIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsumablesConsumableIdPut`: TypesConsumable
	fmt.Fprintf(os.Stdout, "Response from `ConsumablesAPI.ConsumablesConsumableIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consumableId** | **string** | Consumable Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsumablesConsumableIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **typesConsumable** | [**TypesConsumable**](TypesConsumable.md) | Consumable object | 

### Return type

[**TypesConsumable**](TypesConsumable.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsumablesGet

> []TypesConsumable ConsumablesGet(ctx).Execute()

List consumables



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsumablesAPI.ConsumablesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumablesAPI.ConsumablesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsumablesGet`: []TypesConsumable
	fmt.Fprintf(os.Stdout, "Response from `ConsumablesAPI.ConsumablesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiConsumablesGetRequest struct via the builder pattern


### Return type

[**[]TypesConsumable**](TypesConsumable.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsumablesPost

> TypesConsumable ConsumablesPost(ctx).TypesConsumable(typesConsumable).Execute()

Create a consumable



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
	typesConsumable := *openapiclient.NewTypesConsumable("Title_example") // TypesConsumable | Consumable object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsumablesAPI.ConsumablesPost(context.Background()).TypesConsumable(typesConsumable).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumablesAPI.ConsumablesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsumablesPost`: TypesConsumable
	fmt.Fprintf(os.Stdout, "Response from `ConsumablesAPI.ConsumablesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsumablesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **typesConsumable** | [**TypesConsumable**](TypesConsumable.md) | Consumable object | 

### Return type

[**TypesConsumable**](TypesConsumable.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

