# \TimeUnitsAPI

All URIs are relative to *http://127.0.0.1:8080/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TimeUnitsGet**](TimeUnitsAPI.md#TimeUnitsGet) | **Get** /time-units | List time units



## TimeUnitsGet

> []string TimeUnitsGet(ctx).Execute()

List time units



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
	resp, r, err := apiClient.TimeUnitsAPI.TimeUnitsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimeUnitsAPI.TimeUnitsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TimeUnitsGet`: []string
	fmt.Fprintf(os.Stdout, "Response from `TimeUnitsAPI.TimeUnitsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTimeUnitsGetRequest struct via the builder pattern


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

