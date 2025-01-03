# \CategoriesAPI

All URIs are relative to *http://127.0.0.1:8080/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssetsAssetIdCategoriesGet**](CategoriesAPI.md#AssetsAssetIdCategoriesGet) | **Get** /assets/{assetId}/categories | List asset categories
[**CategoriesCategoryIdDelete**](CategoriesAPI.md#CategoriesCategoryIdDelete) | **Delete** /categories/{categoryId} | Delete an asset category
[**CategoriesCategoryIdGet**](CategoriesAPI.md#CategoriesCategoryIdGet) | **Get** /categories/{categoryId} | Get an asset category
[**CategoriesCategoryIdPut**](CategoriesAPI.md#CategoriesCategoryIdPut) | **Put** /categories/{categoryId} | Update an asset category
[**CategoriesGet**](CategoriesAPI.md#CategoriesGet) | **Get** /categories | List asset categories
[**CategoriesPost**](CategoriesAPI.md#CategoriesPost) | **Post** /categories | Create an asset category



## AssetsAssetIdCategoriesGet

> []TypesCategory AssetsAssetIdCategoriesGet(ctx, assetId).Execute()

List asset categories



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
	resp, r, err := apiClient.CategoriesAPI.AssetsAssetIdCategoriesGet(context.Background(), assetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CategoriesAPI.AssetsAssetIdCategoriesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssetsAssetIdCategoriesGet`: []TypesCategory
	fmt.Fprintf(os.Stdout, "Response from `CategoriesAPI.AssetsAssetIdCategoriesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetId** | **string** | Asset Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetsAssetIdCategoriesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]TypesCategory**](TypesCategory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CategoriesCategoryIdDelete

> CategoriesCategoryIdDelete(ctx, categoryId).Execute()

Delete an asset category



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
	r, err := apiClient.CategoriesAPI.CategoriesCategoryIdDelete(context.Background(), categoryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CategoriesAPI.CategoriesCategoryIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categoryId** | **string** | Category Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCategoriesCategoryIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CategoriesCategoryIdGet

> TypesCategory CategoriesCategoryIdGet(ctx, categoryId).Execute()

Get an asset category



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
	resp, r, err := apiClient.CategoriesAPI.CategoriesCategoryIdGet(context.Background(), categoryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CategoriesAPI.CategoriesCategoryIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CategoriesCategoryIdGet`: TypesCategory
	fmt.Fprintf(os.Stdout, "Response from `CategoriesAPI.CategoriesCategoryIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categoryId** | **string** | Category Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCategoriesCategoryIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TypesCategory**](TypesCategory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CategoriesCategoryIdPut

> TypesCategory CategoriesCategoryIdPut(ctx, categoryId).TypesCategory(typesCategory).Execute()

Update an asset category



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
	typesCategory := *openapiclient.NewTypesCategory("Title_example") // TypesCategory | Category object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CategoriesAPI.CategoriesCategoryIdPut(context.Background(), categoryId).TypesCategory(typesCategory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CategoriesAPI.CategoriesCategoryIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CategoriesCategoryIdPut`: TypesCategory
	fmt.Fprintf(os.Stdout, "Response from `CategoriesAPI.CategoriesCategoryIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categoryId** | **string** | Category Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCategoriesCategoryIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **typesCategory** | [**TypesCategory**](TypesCategory.md) | Category object | 

### Return type

[**TypesCategory**](TypesCategory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CategoriesGet

> []TypesCategory CategoriesGet(ctx).Execute()

List asset categories



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
	resp, r, err := apiClient.CategoriesAPI.CategoriesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CategoriesAPI.CategoriesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CategoriesGet`: []TypesCategory
	fmt.Fprintf(os.Stdout, "Response from `CategoriesAPI.CategoriesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCategoriesGetRequest struct via the builder pattern


### Return type

[**[]TypesCategory**](TypesCategory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CategoriesPost

> TypesCategory CategoriesPost(ctx).TypesCategory(typesCategory).Execute()

Create an asset category



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
	typesCategory := *openapiclient.NewTypesCategory("Title_example") // TypesCategory | Category object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CategoriesAPI.CategoriesPost(context.Background()).TypesCategory(typesCategory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CategoriesAPI.CategoriesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CategoriesPost`: TypesCategory
	fmt.Fprintf(os.Stdout, "Response from `CategoriesAPI.CategoriesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCategoriesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **typesCategory** | [**TypesCategory**](TypesCategory.md) | Category object | 

### Return type

[**TypesCategory**](TypesCategory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

