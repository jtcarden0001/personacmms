# \AssetsAPI

All URIs are relative to *http://127.0.0.1:8080/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssetsAssetIdDelete**](AssetsAPI.md#AssetsAssetIdDelete) | **Delete** /assets/{assetId} | Delete an asset
[**AssetsAssetIdGet**](AssetsAPI.md#AssetsAssetIdGet) | **Get** /assets/{assetId} | Get an asset
[**AssetsAssetIdPut**](AssetsAPI.md#AssetsAssetIdPut) | **Put** /assets/{assetId} | Update an asset
[**AssetsGet**](AssetsAPI.md#AssetsGet) | **Get** /assets | List assets
[**AssetsPost**](AssetsAPI.md#AssetsPost) | **Post** /assets | Create an asset
[**CategoriesCategoryIdAssetsAssetIdDelete**](AssetsAPI.md#CategoriesCategoryIdAssetsAssetIdDelete) | **Delete** /categories/{categoryId}/assets/{assetId} | Disassociate an asset with a category
[**CategoriesCategoryIdAssetsAssetIdPut**](AssetsAPI.md#CategoriesCategoryIdAssetsAssetIdPut) | **Put** /categories/{categoryId}/assets/{assetId} | Associate an asset with a category
[**CategoriesCategoryIdAssetsGet**](AssetsAPI.md#CategoriesCategoryIdAssetsGet) | **Get** /categories/{categoryId}/assets | List assets by category
[**CategoriesCategoryIdGroupsGroupIdAssetsGet**](AssetsAPI.md#CategoriesCategoryIdGroupsGroupIdAssetsGet) | **Get** /categories/{categoryId}/groups/{groupId}/assets | List assets by category and group
[**GroupsGroupIdAssetsAssetIdDelete**](AssetsAPI.md#GroupsGroupIdAssetsAssetIdDelete) | **Delete** /groups/{groupId}/assets/{assetId} | Disassociate an asset with a group
[**GroupsGroupIdAssetsAssetIdPut**](AssetsAPI.md#GroupsGroupIdAssetsAssetIdPut) | **Put** /groups/{groupId}/assets/{assetId} | Associate an asset with a group
[**GroupsGroupIdAssetsGet**](AssetsAPI.md#GroupsGroupIdAssetsGet) | **Get** /groups/{groupId}/assets | List assets by group



## AssetsAssetIdDelete

> AssetsAssetIdDelete(ctx, assetId).Execute()

Delete an asset



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
	assetId := "assetId_example" // string | Asset Title

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AssetsAPI.AssetsAssetIdDelete(context.Background(), assetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.AssetsAssetIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetId** | **string** | Asset Title | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetsAssetIdDeleteRequest struct via the builder pattern


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


## AssetsAssetIdGet

> TypesAsset AssetsAssetIdGet(ctx, assetId).Execute()

Get an asset



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
	assetId := "assetId_example" // string | Asset Title

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetsAPI.AssetsAssetIdGet(context.Background(), assetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.AssetsAssetIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssetsAssetIdGet`: TypesAsset
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.AssetsAssetIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetId** | **string** | Asset Title | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetsAssetIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TypesAsset**](TypesAsset.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetsAssetIdPut

> TypesAsset AssetsAssetIdPut(ctx, assetId).TypesAsset(typesAsset).Execute()

Update an asset



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
	assetId := "assetId_example" // string | Asset Title
	typesAsset := *openapiclient.NewTypesAsset("Title_example") // TypesAsset | Asset object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetsAPI.AssetsAssetIdPut(context.Background(), assetId).TypesAsset(typesAsset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.AssetsAssetIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssetsAssetIdPut`: TypesAsset
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.AssetsAssetIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetId** | **string** | Asset Title | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetsAssetIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **typesAsset** | [**TypesAsset**](TypesAsset.md) | Asset object | 

### Return type

[**TypesAsset**](TypesAsset.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetsGet

> []TypesAsset AssetsGet(ctx).Execute()

List assets



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
	resp, r, err := apiClient.AssetsAPI.AssetsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.AssetsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssetsGet`: []TypesAsset
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.AssetsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAssetsGetRequest struct via the builder pattern


### Return type

[**[]TypesAsset**](TypesAsset.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetsPost

> TypesAsset AssetsPost(ctx).TypesAsset(typesAsset).Execute()

Create an asset



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
	typesAsset := *openapiclient.NewTypesAsset("Title_example") // TypesAsset | Asset object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetsAPI.AssetsPost(context.Background()).TypesAsset(typesAsset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.AssetsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssetsPost`: TypesAsset
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.AssetsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAssetsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **typesAsset** | [**TypesAsset**](TypesAsset.md) | Asset object | 

### Return type

[**TypesAsset**](TypesAsset.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CategoriesCategoryIdAssetsAssetIdDelete

> CategoriesCategoryIdAssetsAssetIdDelete(ctx, categoryId, assetId).Execute()

Disassociate an asset with a category



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
	categoryId := "categoryId_example" // string | Category Id
	assetId := "assetId_example" // string | Asset Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AssetsAPI.CategoriesCategoryIdAssetsAssetIdDelete(context.Background(), categoryId, assetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.CategoriesCategoryIdAssetsAssetIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categoryId** | **string** | Category Id | 
**assetId** | **string** | Asset Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCategoriesCategoryIdAssetsAssetIdDeleteRequest struct via the builder pattern


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


## CategoriesCategoryIdAssetsAssetIdPut

> TypesAsset CategoriesCategoryIdAssetsAssetIdPut(ctx, categoryId, assetId).Execute()

Associate an asset with a category



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
	categoryId := "categoryId_example" // string | Category Id
	assetId := "assetId_example" // string | Asset Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetsAPI.CategoriesCategoryIdAssetsAssetIdPut(context.Background(), categoryId, assetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.CategoriesCategoryIdAssetsAssetIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CategoriesCategoryIdAssetsAssetIdPut`: TypesAsset
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.CategoriesCategoryIdAssetsAssetIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categoryId** | **string** | Category Id | 
**assetId** | **string** | Asset Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCategoriesCategoryIdAssetsAssetIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TypesAsset**](TypesAsset.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CategoriesCategoryIdAssetsGet

> []TypesAsset CategoriesCategoryIdAssetsGet(ctx, categoryId).Execute()

List assets by category



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
	categoryId := "categoryId_example" // string | Category Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetsAPI.CategoriesCategoryIdAssetsGet(context.Background(), categoryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.CategoriesCategoryIdAssetsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CategoriesCategoryIdAssetsGet`: []TypesAsset
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.CategoriesCategoryIdAssetsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categoryId** | **string** | Category Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCategoriesCategoryIdAssetsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]TypesAsset**](TypesAsset.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CategoriesCategoryIdGroupsGroupIdAssetsGet

> []TypesAsset CategoriesCategoryIdGroupsGroupIdAssetsGet(ctx, categoryId, groupId).Execute()

List assets by category and group



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
	categoryId := "categoryId_example" // string | Category Id
	groupId := "groupId_example" // string | Group Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetsAPI.CategoriesCategoryIdGroupsGroupIdAssetsGet(context.Background(), categoryId, groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.CategoriesCategoryIdGroupsGroupIdAssetsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CategoriesCategoryIdGroupsGroupIdAssetsGet`: []TypesAsset
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.CategoriesCategoryIdGroupsGroupIdAssetsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categoryId** | **string** | Category Id | 
**groupId** | **string** | Group Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCategoriesCategoryIdGroupsGroupIdAssetsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]TypesAsset**](TypesAsset.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsGroupIdAssetsAssetIdDelete

> GroupsGroupIdAssetsAssetIdDelete(ctx, groupId, assetId).Execute()

Disassociate an asset with a group



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
	assetId := "assetId_example" // string | Asset Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AssetsAPI.GroupsGroupIdAssetsAssetIdDelete(context.Background(), groupId, assetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.GroupsGroupIdAssetsAssetIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Group Id | 
**assetId** | **string** | Asset Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsGroupIdAssetsAssetIdDeleteRequest struct via the builder pattern


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


## GroupsGroupIdAssetsAssetIdPut

> TypesAsset GroupsGroupIdAssetsAssetIdPut(ctx, groupId, assetId).Execute()

Associate an asset with a group



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
	assetId := "assetId_example" // string | Asset Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetsAPI.GroupsGroupIdAssetsAssetIdPut(context.Background(), groupId, assetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.GroupsGroupIdAssetsAssetIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GroupsGroupIdAssetsAssetIdPut`: TypesAsset
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.GroupsGroupIdAssetsAssetIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Group Id | 
**assetId** | **string** | Asset Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsGroupIdAssetsAssetIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TypesAsset**](TypesAsset.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsGroupIdAssetsGet

> []TypesAsset GroupsGroupIdAssetsGet(ctx, groupId).Execute()

List assets by group



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
	resp, r, err := apiClient.AssetsAPI.GroupsGroupIdAssetsGet(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.GroupsGroupIdAssetsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GroupsGroupIdAssetsGet`: []TypesAsset
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.GroupsGroupIdAssetsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Group Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsGroupIdAssetsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]TypesAsset**](TypesAsset.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

