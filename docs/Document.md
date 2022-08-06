# \Document

All URIs are relative to *http://localhost:4080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Bulk**](Document.md#Bulk) | **Post** /api/_bulk | Bulk documents
[**Bulkv2**](Document.md#Bulkv2) | **Post** /api/_bulkv2 | Bulkv2 documents
[**Delete**](Document.md#Delete) | **Delete** /api/{index}/_doc/{id} | Delete document
[**ESBulk**](Document.md#ESBulk) | **Post** /es/_bulk | ES bulk documents
[**Index**](Document.md#Index) | **Post** /api/{index}/_doc | Create or update document
[**IndexWithID**](Document.md#IndexWithID) | **Put** /api/{index}/_doc/{id} | Create or update document with id
[**Multi**](Document.md#Multi) | **Post** /api/{index}/_multi | Multi documents
[**Update**](Document.md#Update) | **Post** /api/{index}/_update/{id} | Update document with id



## Bulk

> MetaHTTPResponseRecordCount Bulk(ctx).Query(query).Execute()

Bulk documents

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
    resp, r, err := apiClient.Document.Bulk(context.Background()).Query(query).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Document.Bulk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Bulk`: MetaHTTPResponseRecordCount
    fmt.Fprintf(os.Stdout, "Response from `Document.Bulk`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | Query | 

### Return type

[**MetaHTTPResponseRecordCount**](MetaHTTPResponseRecordCount.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Bulkv2

> MetaHTTPResponseRecordCount Bulkv2(ctx).Query(query).Execute()

Bulkv2 documents

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
    query := *client.NewMetaJSONIngest() // MetaJSONIngest | Query

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.Document.Bulkv2(context.Background()).Query(query).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Document.Bulkv2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Bulkv2`: MetaHTTPResponseRecordCount
    fmt.Fprintf(os.Stdout, "Response from `Document.Bulkv2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkv2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | [**MetaJSONIngest**](MetaJSONIngest.md) | Query | 

### Return type

[**MetaHTTPResponseRecordCount**](MetaHTTPResponseRecordCount.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> MetaHTTPResponseDocument Delete(ctx, index, id).Execute()

Delete document

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
    id := "id_example" // string | ID

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.Document.Delete(context.Background(), index, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Document.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Delete`: MetaHTTPResponseDocument
    fmt.Fprintf(os.Stdout, "Response from `Document.Delete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | Index | 
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MetaHTTPResponseDocument**](MetaHTTPResponseDocument.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ESBulk

> map[string]interface{} ESBulk(ctx).Query(query).Execute()

ES bulk documents

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
    resp, r, err := apiClient.Document.ESBulk(context.Background()).Query(query).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Document.ESBulk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ESBulk`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `Document.ESBulk`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiESBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | Query | 

### Return type

**map[string]interface{}**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Index

> MetaHTTPResponseID Index(ctx, index).Document(document).Execute()

Create or update document

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
    document := map[string]interface{}{ ... } // map[string]interface{} | Document

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.Document.Index(context.Background(), index).Document(document).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Document.Index``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Index`: MetaHTTPResponseID
    fmt.Fprintf(os.Stdout, "Response from `Document.Index`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | Index | 

### Other Parameters

Other parameters are passed through a pointer to a apiIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **document** | **map[string]interface{}** | Document | 

### Return type

[**MetaHTTPResponseID**](MetaHTTPResponseID.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IndexWithID

> MetaHTTPResponseID IndexWithID(ctx, index, id).Document(document).Execute()

Create or update document with id

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
    id := "id_example" // string | ID
    document := map[string]interface{}{ ... } // map[string]interface{} | Document

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.Document.IndexWithID(context.Background(), index, id).Document(document).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Document.IndexWithID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IndexWithID`: MetaHTTPResponseID
    fmt.Fprintf(os.Stdout, "Response from `Document.IndexWithID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | Index | 
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiIndexWithIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **document** | **map[string]interface{}** | Document | 

### Return type

[**MetaHTTPResponseID**](MetaHTTPResponseID.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Multi

> MetaHTTPResponseRecordCount Multi(ctx, index).Query(query).Execute()

Multi documents

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
    query := "query_example" // string | Query

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.Document.Multi(context.Background(), index).Query(query).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Document.Multi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Multi`: MetaHTTPResponseRecordCount
    fmt.Fprintf(os.Stdout, "Response from `Document.Multi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | Index | 

### Other Parameters

Other parameters are passed through a pointer to a apiMultiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **query** | **string** | Query | 

### Return type

[**MetaHTTPResponseRecordCount**](MetaHTTPResponseRecordCount.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> MetaHTTPResponseID Update(ctx, index, id).Document(document).Execute()

Update document with id

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
    id := "id_example" // string | ID
    document := map[string]interface{}{ ... } // map[string]interface{} | Document

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.Document.Update(context.Background(), index, id).Document(document).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Document.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: MetaHTTPResponseID
    fmt.Fprintf(os.Stdout, "Response from `Document.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | Index | 
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **document** | **map[string]interface{}** | Document | 

### Return type

[**MetaHTTPResponseID**](MetaHTTPResponseID.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

