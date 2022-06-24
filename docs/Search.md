# \Search

All URIs are relative to *http://localhost:4080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MSearch**](Search.md#MSearch) | **Post** /es/_msearch | Search V2 MultipleSearch for compatible ES
[**Search**](Search.md#Search) | **Post** /es/{index}/_search | Search V2 DSL for compatible ES
[**SearchV1**](Search.md#SearchV1) | **Post** /api/{index}/_search | Search V1



## MSearch

> MetaSearchResponse MSearch(ctx).Query(query).Execute()

Search V2 MultipleSearch for compatible ES

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
    query := "query_example" // string | Query

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.Search.MSearch(context.Background()).Query(query).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Search.MSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MSearch`: MetaSearchResponse
    fmt.Fprintf(os.Stdout, "Response from `Search.MSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | Query | 

### Return type

[**MetaSearchResponse**](MetaSearchResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Search

> MetaSearchResponse Search(ctx, index).Query(query).Execute()

Search V2 DSL for compatible ES

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
    index := "index_example" // string | Index
    query := *client.NewMetaZincQuery() // MetaZincQuery | Query

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.Search.Search(context.Background(), index).Query(query).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Search.Search``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Search`: MetaSearchResponse
    fmt.Fprintf(os.Stdout, "Response from `Search.Search`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | Index | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **query** | [**MetaZincQuery**](MetaZincQuery.md) | Query | 

### Return type

[**MetaSearchResponse**](MetaSearchResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchV1

> V1SearchResponse SearchV1(ctx, index).Query(query).Execute()

Search V1

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
    index := "index_example" // string | Index
    query := *client.NewV1ZincQuery() // V1ZincQuery | Query

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.Search.SearchV1(context.Background(), index).Query(query).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Search.SearchV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchV1`: V1SearchResponse
    fmt.Fprintf(os.Stdout, "Response from `Search.SearchV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | Index | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **query** | [**V1ZincQuery**](V1ZincQuery.md) | Query | 

### Return type

[**V1SearchResponse**](V1SearchResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

