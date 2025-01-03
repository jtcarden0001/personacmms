# \UsageUnitsAPI

All URIs are relative to *http://127.0.0.1:8080/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsageUnitsGet**](UsageUnitsAPI.md#UsageUnitsGet) | **Get** /usage-units | List usage units



## UsageUnitsGet

> []string UsageUnitsGet(ctx).Execute()

List usage units



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
	resp, r, err := apiClient.UsageUnitsAPI.UsageUnitsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageUnitsAPI.UsageUnitsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsageUnitsGet`: []string
	fmt.Fprintf(os.Stdout, "Response from `UsageUnitsAPI.UsageUnitsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUsageUnitsGetRequest struct via the builder pattern


### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

