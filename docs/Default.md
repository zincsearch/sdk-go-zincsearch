# \Default

All URIs are relative to *http://localhost:4080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Healthz**](Default.md#Healthz) | **Get** /healthz | Get healthz
[**Version**](Default.md#Version) | **Get** /version | Get version



## Healthz

> MetaHealthzResponse Healthz(ctx).Execute()

Get healthz

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    client "github.com/zinclabs/sdk-go-zincsearch"
)

func main() {

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.Default.Healthz(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Default.Healthz``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Healthz`: MetaHealthzResponse
    fmt.Fprintf(os.Stdout, "Response from `Default.Healthz`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiHealthzRequest struct via the builder pattern


### Return type

[**MetaHealthzResponse**](MetaHealthzResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Version

> MetaVersionResponse Version(ctx).Execute()

Get version

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    client "github.com/zinclabs/sdk-go-zincsearch"
)

func main() {

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.Default.Version(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Default.Version``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Version`: MetaVersionResponse
    fmt.Fprintf(os.Stdout, "Response from `Default.Version`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiVersionRequest struct via the builder pattern


### Return type

[**MetaVersionResponse**](MetaVersionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

