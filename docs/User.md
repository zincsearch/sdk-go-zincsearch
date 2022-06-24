# \User

All URIs are relative to *http://localhost:4080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](User.md#Create) | **Post** /api/user | Create user
[**Delete**](User.md#Delete) | **Delete** /api/user/{id} | Delete user
[**List**](User.md#List) | **Get** /api/user | List user
[**Login**](User.md#Login) | **Post** /api/login | Login
[**Update**](User.md#Update) | **Put** /api/user | Update user



## Create

> MetaHTTPResponseID Create(ctx).User(user).Execute()

Create user

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
    user := *client.NewMetaUser() // MetaUser | User data

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.User.Create(context.Background()).User(user).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `User.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: MetaHTTPResponseID
    fmt.Fprintf(os.Stdout, "Response from `User.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user** | [**MetaUser**](MetaUser.md) | User data | 

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


## Delete

> MetaHTTPResponseID Delete(ctx, id).Execute()

Delete user

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
    id := "id_example" // string | User id

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.User.Delete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `User.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Delete`: MetaHTTPResponseID
    fmt.Fprintf(os.Stdout, "Response from `User.Delete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | User id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MetaHTTPResponseID**](MetaHTTPResponseID.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> []MetaUser List(ctx).Execute()

List user

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
    resp, r, err := apiClient.User.List(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `User.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: []MetaUser
    fmt.Fprintf(os.Stdout, "Response from `User.List`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


### Return type

[**[]MetaUser**](MetaUser.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Login

> AuthLoginResponse Login(ctx).Login(login).Execute()

Login

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
    login := *client.NewAuthLoginRequest() // AuthLoginRequest | Login credentials

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.User.Login(context.Background()).Login(login).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `User.Login``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Login`: AuthLoginResponse
    fmt.Fprintf(os.Stdout, "Response from `User.Login`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **login** | [**AuthLoginRequest**](AuthLoginRequest.md) | Login credentials | 

### Return type

[**AuthLoginResponse**](AuthLoginResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> MetaHTTPResponseID Update(ctx).User(user).Execute()

Update user

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
    user := *client.NewMetaUser() // MetaUser | User data

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.User.Update(context.Background()).User(user).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `User.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: MetaHTTPResponseID
    fmt.Fprintf(os.Stdout, "Response from `User.Update`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user** | [**MetaUser**](MetaUser.md) | User data | 

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

