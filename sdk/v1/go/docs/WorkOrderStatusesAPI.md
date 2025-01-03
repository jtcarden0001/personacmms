# \WorkOrderStatusesAPI

All URIs are relative to *http://127.0.0.1:8080/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorkOrderStatusesGet**](WorkOrderStatusesAPI.md#WorkOrderStatusesGet) | **Get** /work-order-statuses | List work order statuses



## WorkOrderStatusesGet

> []string WorkOrderStatusesGet(ctx).Execute()

List work order statuses



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
	resp, r, err := apiClient.WorkOrderStatusesAPI.WorkOrderStatusesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkOrderStatusesAPI.WorkOrderStatusesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorkOrderStatusesGet`: []string
	fmt.Fprintf(os.Stdout, "Response from `WorkOrderStatusesAPI.WorkOrderStatusesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWorkOrderStatusesGetRequest struct via the builder pattern


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

