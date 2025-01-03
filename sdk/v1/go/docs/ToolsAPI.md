# \ToolsAPI

All URIs are relative to *http://127.0.0.1:8080/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssetsAssetIdTasksTaskIdToolsToolIdDelete**](ToolsAPI.md#AssetsAssetIdTasksTaskIdToolsToolIdDelete) | **Delete** /assets/{assetId}/tasks/{taskId}/tools/{toolId} | Disassociate a tool with a task
[**AssetsAssetIdTasksTaskIdToolsToolIdPut**](ToolsAPI.md#AssetsAssetIdTasksTaskIdToolsToolIdPut) | **Put** /assets/{assetId}/tasks/{taskId}/tools/{toolId} | Associate a tool with a task
[**AssetsAssetIdWorkOrdersWorkOrderIdToolsToolIdDelete**](ToolsAPI.md#AssetsAssetIdWorkOrdersWorkOrderIdToolsToolIdDelete) | **Delete** /assets/{assetId}/work-orders/{workOrderId}/tools/{toolId} | Disassociate a tool with a work order
[**AssetsAssetIdWorkOrdersWorkOrderIdToolsToolIdPut**](ToolsAPI.md#AssetsAssetIdWorkOrdersWorkOrderIdToolsToolIdPut) | **Put** /assets/{assetId}/work-orders/{workOrderId}/tools/{toolId} | Associate a tool with a work order
[**ToolsGet**](ToolsAPI.md#ToolsGet) | **Get** /tools | List tools
[**ToolsPost**](ToolsAPI.md#ToolsPost) | **Post** /tools | Create a tool
[**ToolsToolIdDelete**](ToolsAPI.md#ToolsToolIdDelete) | **Delete** /tools/{toolId} | Delete a tool
[**ToolsToolIdGet**](ToolsAPI.md#ToolsToolIdGet) | **Get** /tools/{toolId} | Get a tool
[**ToolsToolIdPut**](ToolsAPI.md#ToolsToolIdPut) | **Put** /tools/{toolId} | Update a tool



## AssetsAssetIdTasksTaskIdToolsToolIdDelete

> AssetsAssetIdTasksTaskIdToolsToolIdDelete(ctx, assetId, taskId, toolId).Execute()

Disassociate a tool with a task



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
	toolId := "toolId_example" // string | Tool ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ToolsAPI.AssetsAssetIdTasksTaskIdToolsToolIdDelete(context.Background(), assetId, taskId, toolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.AssetsAssetIdTasksTaskIdToolsToolIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetId** | **string** | Asset ID | 
**taskId** | **string** | Task ID | 
**toolId** | **string** | Tool ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetsAssetIdTasksTaskIdToolsToolIdDeleteRequest struct via the builder pattern


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


## AssetsAssetIdTasksTaskIdToolsToolIdPut

> TypesTool AssetsAssetIdTasksTaskIdToolsToolIdPut(ctx, assetId, taskId, toolId).TypesToolSize(typesToolSize).Execute()

Associate a tool with a task



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
	toolId := "toolId_example" // string | Tool ID
	typesToolSize := *openapiclient.NewTypesToolSize() // TypesToolSize | Tool object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToolsAPI.AssetsAssetIdTasksTaskIdToolsToolIdPut(context.Background(), assetId, taskId, toolId).TypesToolSize(typesToolSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.AssetsAssetIdTasksTaskIdToolsToolIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssetsAssetIdTasksTaskIdToolsToolIdPut`: TypesTool
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.AssetsAssetIdTasksTaskIdToolsToolIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetId** | **string** | Asset ID | 
**taskId** | **string** | Task ID | 
**toolId** | **string** | Tool ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetsAssetIdTasksTaskIdToolsToolIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **typesToolSize** | [**TypesToolSize**](TypesToolSize.md) | Tool object | 

### Return type

[**TypesTool**](TypesTool.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetsAssetIdWorkOrdersWorkOrderIdToolsToolIdDelete

> AssetsAssetIdWorkOrdersWorkOrderIdToolsToolIdDelete(ctx, assetId, workOrderId, toolId).Execute()

Disassociate a tool with a work order



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
	workOrderId := "workOrderId_example" // string | Work Order ID
	toolId := "toolId_example" // string | Tool ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ToolsAPI.AssetsAssetIdWorkOrdersWorkOrderIdToolsToolIdDelete(context.Background(), assetId, workOrderId, toolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.AssetsAssetIdWorkOrdersWorkOrderIdToolsToolIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetId** | **string** | Asset ID | 
**workOrderId** | **string** | Work Order ID | 
**toolId** | **string** | Tool ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetsAssetIdWorkOrdersWorkOrderIdToolsToolIdDeleteRequest struct via the builder pattern


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


## AssetsAssetIdWorkOrdersWorkOrderIdToolsToolIdPut

> TypesTool AssetsAssetIdWorkOrdersWorkOrderIdToolsToolIdPut(ctx, assetId, workOrderId, toolId).TypesToolSize(typesToolSize).Execute()

Associate a tool with a work order



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
	workOrderId := "workOrderId_example" // string | Work Order ID
	toolId := "toolId_example" // string | Tool ID
	typesToolSize := *openapiclient.NewTypesToolSize() // TypesToolSize | Tool object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToolsAPI.AssetsAssetIdWorkOrdersWorkOrderIdToolsToolIdPut(context.Background(), assetId, workOrderId, toolId).TypesToolSize(typesToolSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.AssetsAssetIdWorkOrdersWorkOrderIdToolsToolIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssetsAssetIdWorkOrdersWorkOrderIdToolsToolIdPut`: TypesTool
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.AssetsAssetIdWorkOrdersWorkOrderIdToolsToolIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetId** | **string** | Asset ID | 
**workOrderId** | **string** | Work Order ID | 
**toolId** | **string** | Tool ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetsAssetIdWorkOrdersWorkOrderIdToolsToolIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **typesToolSize** | [**TypesToolSize**](TypesToolSize.md) | Tool object | 

### Return type

[**TypesTool**](TypesTool.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ToolsGet

> []TypesTool ToolsGet(ctx).Execute()

List tools



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
	resp, r, err := apiClient.ToolsAPI.ToolsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.ToolsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ToolsGet`: []TypesTool
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.ToolsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiToolsGetRequest struct via the builder pattern


### Return type

[**[]TypesTool**](TypesTool.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ToolsPost

> TypesTool ToolsPost(ctx).TypesTool(typesTool).Execute()

Create a tool



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
	typesTool := *openapiclient.NewTypesTool("Title_example") // TypesTool | Tool object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToolsAPI.ToolsPost(context.Background()).TypesTool(typesTool).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.ToolsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ToolsPost`: TypesTool
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.ToolsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiToolsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **typesTool** | [**TypesTool**](TypesTool.md) | Tool object | 

### Return type

[**TypesTool**](TypesTool.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ToolsToolIdDelete

> ToolsToolIdDelete(ctx, toolId).Execute()

Delete a tool



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
	toolId := "toolId_example" // string | Tool Title

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ToolsAPI.ToolsToolIdDelete(context.Background(), toolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.ToolsToolIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**toolId** | **string** | Tool Title | 

### Other Parameters

Other parameters are passed through a pointer to a apiToolsToolIdDeleteRequest struct via the builder pattern


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


## ToolsToolIdGet

> TypesTool ToolsToolIdGet(ctx, toolId).Execute()

Get a tool



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
	toolId := "toolId_example" // string | Tool Title

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToolsAPI.ToolsToolIdGet(context.Background(), toolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.ToolsToolIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ToolsToolIdGet`: TypesTool
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.ToolsToolIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**toolId** | **string** | Tool Title | 

### Other Parameters

Other parameters are passed through a pointer to a apiToolsToolIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TypesTool**](TypesTool.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ToolsToolIdPut

> TypesTool ToolsToolIdPut(ctx, toolId).TypesTool(typesTool).Execute()

Update a tool



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
	toolId := "toolId_example" // string | Tool Title
	typesTool := *openapiclient.NewTypesTool("Title_example") // TypesTool | Tool object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToolsAPI.ToolsToolIdPut(context.Background(), toolId).TypesTool(typesTool).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.ToolsToolIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ToolsToolIdPut`: TypesTool
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.ToolsToolIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**toolId** | **string** | Tool Title | 

### Other Parameters

Other parameters are passed through a pointer to a apiToolsToolIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **typesTool** | [**TypesTool**](TypesTool.md) | Tool object | 

### Return type

[**TypesTool**](TypesTool.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

