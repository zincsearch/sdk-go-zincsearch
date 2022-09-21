# \Index

All URIs are relative to *http://localhost:4080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddOrRemoveESAlias**](Index.md#AddOrRemoveESAlias) | **Post** /es/_aliases | Add or remove index alias for compatible ES
[**Analyze**](Index.md#Analyze) | **Post** /api/_analyze | Analyze
[**AnalyzeIndex**](Index.md#AnalyzeIndex) | **Post** /api/{index}/_analyze | Analyze
[**Create**](Index.md#Create) | **Post** /api/index | Create index
[**CreateTemplate**](Index.md#CreateTemplate) | **Post** /es/_index_template | Create update index template
[**Delete**](Index.md#Delete) | **Delete** /api/index/{index} | Delete index
[**DeleteTemplate**](Index.md#DeleteTemplate) | **Delete** /es/_index_template/{name} | Delete template
[**ESCreateIndex**](Index.md#ESCreateIndex) | **Put** /es/{index} | Create index for compatible ES
[**ESGetMapping**](Index.md#ESGetMapping) | **Get** /es/{index}/_mapping | Get index mappings for compatible ES
[**EsExists**](Index.md#EsExists) | **Head** /es/{index} | Checks if the index exists for compatible ES
[**Exists**](Index.md#Exists) | **Head** /api/index/{index} | Checks if the index exists
[**GetESAliases**](Index.md#GetESAliases) | **Get** /es/{target}/_alias/{target_alias} | Get index alias for compatible ES
[**GetIndex**](Index.md#GetIndex) | **Get** /api/index/{index} | Get index metadata
[**GetMapping**](Index.md#GetMapping) | **Get** /api/{index}/_mapping | Get index mappings
[**GetSettings**](Index.md#GetSettings) | **Get** /api/{index}/_settings | Get index settings
[**GetTemplate**](Index.md#GetTemplate) | **Get** /es/_index_template/{name} | Get index template
[**IndexNameList**](Index.md#IndexNameList) | **Get** /api/index_name | List index Name
[**List**](Index.md#List) | **Get** /api/index | List indexes
[**ListTemplates**](Index.md#ListTemplates) | **Get** /es/_index_template | List index teplates
[**Refresh**](Index.md#Refresh) | **Post** /api/index/{index}/refresh | Resfresh index
[**SetMapping**](Index.md#SetMapping) | **Put** /api/{index}/_mapping | Set index mappings
[**SetSettings**](Index.md#SetSettings) | **Put** /api/{index}/_settings | Set index Settings
[**UpdateTemplate**](Index.md#UpdateTemplate) | **Put** /es/_index_template/{name} | Create update index template



## AddOrRemoveESAlias

> map[string]interface{} AddOrRemoveESAlias(ctx).Execute()

Add or remove index alias for compatible ES

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
    resp, r, err := apiClient.Index.AddOrRemoveESAlias(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Index.AddOrRemoveESAlias``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddOrRemoveESAlias`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `Index.AddOrRemoveESAlias`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAddOrRemoveESAliasRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Analyze

> IndexAnalyzeResponse Analyze(ctx).Query(query).Execute()

Analyze

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
    query := map[string]interface{}{ ... } // map[string]interface{} | Query

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.Index.Analyze(context.Background()).Query(query).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Index.Analyze``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Analyze`: IndexAnalyzeResponse
    fmt.Fprintf(os.Stdout, "Response from `Index.Analyze`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAnalyzeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **map[string]interface{}** | Query | 

### Return type

[**IndexAnalyzeResponse**](IndexAnalyzeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyzeIndex

> IndexAnalyzeResponse AnalyzeIndex(ctx, index).Query(query).Execute()

Analyze

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
    query := map[string]interface{}{ ... } // map[string]interface{} | Query

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.Index.AnalyzeIndex(context.Background(), index).Query(query).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Index.AnalyzeIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AnalyzeIndex`: IndexAnalyzeResponse
    fmt.Fprintf(os.Stdout, "Response from `Index.AnalyzeIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | Index | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyzeIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **query** | **map[string]interface{}** | Query | 

### Return type

[**IndexAnalyzeResponse**](IndexAnalyzeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Create

> MetaHTTPResponseIndex Create(ctx).Data(data).Execute()

Create index

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
    data := *client.NewMetaIndexSimple() // MetaIndexSimple | Index data

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.Index.Create(context.Background()).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Index.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: MetaHTTPResponseIndex
    fmt.Fprintf(os.Stdout, "Response from `Index.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**MetaIndexSimple**](MetaIndexSimple.md) | Index data | 

### Return type

[**MetaHTTPResponseIndex**](MetaHTTPResponseIndex.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTemplate

> MetaHTTPResponseTemplate CreateTemplate(ctx).Template(template).Execute()

Create update index template

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
    template := *client.NewMetaIndexTemplate() // MetaIndexTemplate | Template data

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.Index.CreateTemplate(context.Background()).Template(template).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Index.CreateTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTemplate`: MetaHTTPResponseTemplate
    fmt.Fprintf(os.Stdout, "Response from `Index.CreateTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **template** | [**MetaIndexTemplate**](MetaIndexTemplate.md) | Template data | 

### Return type

[**MetaHTTPResponseTemplate**](MetaHTTPResponseTemplate.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> MetaHTTPResponseIndex Delete(ctx, index).Execute()

Delete index

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

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.Index.Delete(context.Background(), index).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Index.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Delete`: MetaHTTPResponseIndex
    fmt.Fprintf(os.Stdout, "Response from `Index.Delete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | Index | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MetaHTTPResponseIndex**](MetaHTTPResponseIndex.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTemplate

> MetaHTTPResponse DeleteTemplate(ctx, name).Execute()

Delete template

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
    name := "name_example" // string | Template

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.Index.DeleteTemplate(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Index.DeleteTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteTemplate`: MetaHTTPResponse
    fmt.Fprintf(os.Stdout, "Response from `Index.DeleteTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Template | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MetaHTTPResponse**](MetaHTTPResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ESCreateIndex

> map[string]interface{} ESCreateIndex(ctx, index).Data(data).Execute()

Create index for compatible ES

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
    data := *client.NewMetaIndexSimple() // MetaIndexSimple | Index data

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.Index.ESCreateIndex(context.Background(), index).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Index.ESCreateIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ESCreateIndex`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `Index.ESCreateIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | Index | 

### Other Parameters

Other parameters are passed through a pointer to a apiESCreateIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**MetaIndexSimple**](MetaIndexSimple.md) | Index data | 

### Return type

**map[string]interface{}**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ESGetMapping

> map[string]interface{} ESGetMapping(ctx, index).Execute()

Get index mappings for compatible ES

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

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.Index.ESGetMapping(context.Background(), index).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Index.ESGetMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ESGetMapping`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `Index.ESGetMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | Index | 

### Other Parameters

Other parameters are passed through a pointer to a apiESGetMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EsExists

> MetaHTTPResponse EsExists(ctx, index).Execute()

Checks if the index exists for compatible ES

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

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.Index.EsExists(context.Background(), index).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Index.EsExists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EsExists`: MetaHTTPResponse
    fmt.Fprintf(os.Stdout, "Response from `Index.EsExists`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | Index | 

### Other Parameters

Other parameters are passed through a pointer to a apiEsExistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MetaHTTPResponse**](MetaHTTPResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Exists

> MetaHTTPResponse Exists(ctx, index).Execute()

Checks if the index exists

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

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.Index.Exists(context.Background(), index).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Index.Exists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Exists`: MetaHTTPResponse
    fmt.Fprintf(os.Stdout, "Response from `Index.Exists`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | Index | 

### Other Parameters

Other parameters are passed through a pointer to a apiExistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MetaHTTPResponse**](MetaHTTPResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetESAliases

> map[string]interface{} GetESAliases(ctx, target, targetAlias).Execute()

Get index alias for compatible ES

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
    target := "target_example" // string | Target Index
    targetAlias := "targetAlias_example" // string | Target Alias

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.Index.GetESAliases(context.Background(), target, targetAlias).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Index.GetESAliases``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetESAliases`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `Index.GetESAliases`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**target** | **string** | Target Index | 
**targetAlias** | **string** | Target Alias | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetESAliasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIndex

> map[string]interface{} GetIndex(ctx, index).Execute()

Get index metadata

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

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.Index.GetIndex(context.Background(), index).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Index.GetIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIndex`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `Index.GetIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | Index | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMapping

> map[string]interface{} GetMapping(ctx, index).Execute()

Get index mappings

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

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.Index.GetMapping(context.Background(), index).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Index.GetMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMapping`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `Index.GetMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | Index | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSettings

> map[string]interface{} GetSettings(ctx, index).Execute()

Get index settings

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

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.Index.GetSettings(context.Background(), index).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Index.GetSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSettings`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `Index.GetSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | Index | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTemplate

> MetaIndexTemplate GetTemplate(ctx, name).Execute()

Get index template

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
    name := "name_example" // string | Template

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.Index.GetTemplate(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Index.GetTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTemplate`: MetaIndexTemplate
    fmt.Fprintf(os.Stdout, "Response from `Index.GetTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Template | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MetaIndexTemplate**](MetaIndexTemplate.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IndexNameList

> []string IndexNameList(ctx).Name(name).Execute()

List index Name

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
    name := "name_example" // string | IndexName (optional)

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.Index.IndexNameList(context.Background()).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Index.IndexNameList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IndexNameList`: []string
    fmt.Fprintf(os.Stdout, "Response from `Index.IndexNameList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIndexNameListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | IndexName | 

### Return type

**[]string**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> IndexIndexListResponse List(ctx).PageNum(pageNum).PageSize(pageSize).SortBy(sortBy).Desc(desc).Name(name).Execute()

List indexes

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
    pageNum := int32(56) // int32 | page num (optional)
    pageSize := int32(56) // int32 | page size (optional)
    sortBy := "sortBy_example" // string | sort by (optional)
    desc := true // bool | desc (optional)
    name := "name_example" // string | name (optional)

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.Index.List(context.Background()).PageNum(pageNum).PageSize(pageSize).SortBy(sortBy).Desc(desc).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Index.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: IndexIndexListResponse
    fmt.Fprintf(os.Stdout, "Response from `Index.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNum** | **int32** | page num | 
 **pageSize** | **int32** | page size | 
 **sortBy** | **string** | sort by | 
 **desc** | **bool** | desc | 
 **name** | **string** | name | 

### Return type

[**IndexIndexListResponse**](IndexIndexListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTemplates

> []MetaTemplate ListTemplates(ctx).Execute()

List index teplates

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
    resp, r, err := apiClient.Index.ListTemplates(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Index.ListTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTemplates`: []MetaTemplate
    fmt.Fprintf(os.Stdout, "Response from `Index.ListTemplates`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListTemplatesRequest struct via the builder pattern


### Return type

[**[]MetaTemplate**](MetaTemplate.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Refresh

> MetaHTTPResponse Refresh(ctx, index).Execute()

Resfresh index

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

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.Index.Refresh(context.Background(), index).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Index.Refresh``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Refresh`: MetaHTTPResponse
    fmt.Fprintf(os.Stdout, "Response from `Index.Refresh`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | Index | 

### Other Parameters

Other parameters are passed through a pointer to a apiRefreshRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MetaHTTPResponse**](MetaHTTPResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetMapping

> MetaHTTPResponse SetMapping(ctx, index).Mapping(mapping).Execute()

Set index mappings

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
    mapping := *client.NewMetaMappings() // MetaMappings | Mapping

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.Index.SetMapping(context.Background(), index).Mapping(mapping).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Index.SetMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetMapping`: MetaHTTPResponse
    fmt.Fprintf(os.Stdout, "Response from `Index.SetMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | Index | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mapping** | [**MetaMappings**](MetaMappings.md) | Mapping | 

### Return type

[**MetaHTTPResponse**](MetaHTTPResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetSettings

> MetaHTTPResponse SetSettings(ctx, index).Settings(settings).Execute()

Set index Settings

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
    settings := *client.NewMetaIndexSettings() // MetaIndexSettings | Settings

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.Index.SetSettings(context.Background(), index).Settings(settings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Index.SetSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetSettings`: MetaHTTPResponse
    fmt.Fprintf(os.Stdout, "Response from `Index.SetSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | Index | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **settings** | [**MetaIndexSettings**](MetaIndexSettings.md) | Settings | 

### Return type

[**MetaHTTPResponse**](MetaHTTPResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTemplate

> MetaHTTPResponseTemplate UpdateTemplate(ctx, name).Template(template).Execute()

Create update index template

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
    name := "name_example" // string | Template
    template := *client.NewMetaIndexTemplate() // MetaIndexTemplate | Template data

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.Index.UpdateTemplate(context.Background(), name).Template(template).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Index.UpdateTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTemplate`: MetaHTTPResponseTemplate
    fmt.Fprintf(os.Stdout, "Response from `Index.UpdateTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Template | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **template** | [**MetaIndexTemplate**](MetaIndexTemplate.md) | Template data | 

### Return type

[**MetaHTTPResponseTemplate**](MetaHTTPResponseTemplate.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

