# \GroupsAPI

All URIs are relative to *http://127.0.0.1:8080/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssetsAssetIdGroupsGet**](GroupsAPI.md#AssetsAssetIdGroupsGet) | **Get** /assets/{assetId}/groups | List groups by asset
[**GroupsGet**](GroupsAPI.md#GroupsGet) | **Get** /groups | List all groups
[**GroupsGroupIdDelete**](GroupsAPI.md#GroupsGroupIdDelete) | **Delete** /groups/{groupId} | Delete a group
[**GroupsGroupIdGet**](GroupsAPI.md#GroupsGroupIdGet) | **Get** /groups/{groupId} | Get a group
[**GroupsGroupIdPut**](GroupsAPI.md#GroupsGroupIdPut) | **Put** /groups/{groupId} | Update a group
[**GroupsPost**](GroupsAPI.md#GroupsPost) | **Post** /groups | Create a group



## AssetsAssetIdGroupsGet

> []TypesGroup AssetsAssetIdGroupsGet(ctx, assetId).Execute()

List groups by asset



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
	resp, r, err := apiClient.GroupsAPI.AssetsAssetIdGroupsGet(context.Background(), assetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.AssetsAssetIdGroupsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssetsAssetIdGroupsGet`: []TypesGroup
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.AssetsAssetIdGroupsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetId** | **string** | Asset Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetsAssetIdGroupsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]TypesGroup**](TypesGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsGet

> []TypesGroup GroupsGet(ctx).Execute()

List all groups



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
	resp, r, err := apiClient.GroupsAPI.GroupsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GroupsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GroupsGet`: []TypesGroup
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GroupsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsGetRequest struct via the builder pattern


### Return type

[**[]TypesGroup**](TypesGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsGroupIdDelete

> GroupsGroupIdDelete(ctx, groupId).Execute()

Delete a group



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
	groupId := "groupId_example" // string | Group Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GroupsAPI.GroupsGroupIdDelete(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GroupsGroupIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Group Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsGroupIdDeleteRequest struct via the builder pattern


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


## GroupsGroupIdGet

> TypesGroup GroupsGroupIdGet(ctx, groupId).Execute()

Get a group



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
	groupId := "groupId_example" // string | Group Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.GroupsGroupIdGet(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GroupsGroupIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GroupsGroupIdGet`: TypesGroup
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GroupsGroupIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Group Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsGroupIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TypesGroup**](TypesGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsGroupIdPut

> TypesGroup GroupsGroupIdPut(ctx, groupId).TypesGroup(typesGroup).Execute()

Update a group



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
	groupId := "groupId_example" // string | Group Id
	typesGroup := *openapiclient.NewTypesGroup("Title_example") // TypesGroup | Group object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.GroupsGroupIdPut(context.Background(), groupId).TypesGroup(typesGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GroupsGroupIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GroupsGroupIdPut`: TypesGroup
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GroupsGroupIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Group Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsGroupIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **typesGroup** | [**TypesGroup**](TypesGroup.md) | Group object | 

### Return type

[**TypesGroup**](TypesGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsPost

> TypesGroup GroupsPost(ctx).TypesGroup(typesGroup).Execute()

Create a group



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
	typesGroup := *openapiclient.NewTypesGroup("Title_example") // TypesGroup | Group object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.GroupsPost(context.Background()).TypesGroup(typesGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GroupsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GroupsPost`: TypesGroup
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GroupsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGroupsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **typesGroup** | [**TypesGroup**](TypesGroup.md) | Group object | 

### Return type

[**TypesGroup**](TypesGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

