# \ShopAPI

All URIs are relative to *https://fs.tokopedia.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateShowcase**](ShopAPI.md#CreateShowcase) | **Post** /v1/showcase/fs/{fs_id}/create | 
[**DeleteShowcase**](ShopAPI.md#DeleteShowcase) | **Post** /v1/showcase/fs/{fs_id}/delete | 
[**GetShopInfo**](ShopAPI.md#GetShopInfo) | **Get** /v1/shop/fs/{fs_id}/shop-info | 
[**GetShowcase**](ShopAPI.md#GetShowcase) | **Get** /v1/showcase/fs/{fs_id}/get | 
[**GetShowcaseAllEtalase**](ShopAPI.md#GetShowcaseAllEtalase) | **Get** /inventory/v1/fs/{fs_id}/product/etalase | 
[**UpdateShopStatus**](ShopAPI.md#UpdateShopStatus) | **Post** /v2/shop/fs/{fs_id}/shop-status | 
[**UpdateShowcase**](ShopAPI.md#UpdateShowcase) | **Post** /v1/showcase/fs/{fs_id}/update | 



## CreateShowcase

> CreateShowcase200Response CreateShowcase(ctx, fsId).ShopId(shopId).CreateShowcaseRequest(createShowcaseRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/dhimas-sirclo/apiclient/tokopedia"
)

func main() {
    fsId := int64(789) // int64 | Fulfillment service unique identifier
    shopId := int64(789) // int64 | Shop unique identifier
    createShowcaseRequest := *openapiclient.NewCreateShowcaseRequest("Name_example") // CreateShowcaseRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShopAPI.CreateShowcase(context.Background(), fsId).ShopId(shopId).CreateShowcaseRequest(createShowcaseRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShopAPI.CreateShowcase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateShowcase`: CreateShowcase200Response
    fmt.Fprintf(os.Stdout, "Response from `ShopAPI.CreateShowcase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsId** | **int64** | Fulfillment service unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateShowcaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shopId** | **int64** | Shop unique identifier | 
 **createShowcaseRequest** | [**CreateShowcaseRequest**](CreateShowcaseRequest.md) |  | 

### Return type

[**CreateShowcase200Response**](CreateShowcase200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteShowcase

> DeleteShowcase200Response DeleteShowcase(ctx, fsId).ShopId(shopId).DeleteShowcaseRequest(deleteShowcaseRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/dhimas-sirclo/apiclient/tokopedia"
)

func main() {
    fsId := int64(789) // int64 | Fulfillment service unique identifier
    shopId := int64(789) // int64 | Shop unique identifier
    deleteShowcaseRequest := *openapiclient.NewDeleteShowcaseRequest(int64(123)) // DeleteShowcaseRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShopAPI.DeleteShowcase(context.Background(), fsId).ShopId(shopId).DeleteShowcaseRequest(deleteShowcaseRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShopAPI.DeleteShowcase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteShowcase`: DeleteShowcase200Response
    fmt.Fprintf(os.Stdout, "Response from `ShopAPI.DeleteShowcase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsId** | **int64** | Fulfillment service unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteShowcaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shopId** | **int64** | Shop unique identifier | 
 **deleteShowcaseRequest** | [**DeleteShowcaseRequest**](DeleteShowcaseRequest.md) |  | 

### Return type

[**DeleteShowcase200Response**](DeleteShowcase200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetShopInfo

> GetShopInfo200Response GetShopInfo(ctx, fsId).ShopId(shopId).Page(page).PerPage(perPage).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/dhimas-sirclo/apiclient/tokopedia"
)

func main() {
    fsId := int64(789) // int64 | Fulfillment service unique identifier
    shopId := int64(789) // int64 | Shop unique identifier (optional)
    page := int64(789) // int64 | Determine which page the order list should start. The minimal value is 1 (optional)
    perPage := int64(789) // int64 | Determine how many orders will be shown per page (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShopAPI.GetShopInfo(context.Background(), fsId).ShopId(shopId).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShopAPI.GetShopInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetShopInfo`: GetShopInfo200Response
    fmt.Fprintf(os.Stdout, "Response from `ShopAPI.GetShopInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsId** | **int64** | Fulfillment service unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetShopInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shopId** | **int64** | Shop unique identifier | 
 **page** | **int64** | Determine which page the order list should start. The minimal value is 1 | 
 **perPage** | **int64** | Determine how many orders will be shown per page | 

### Return type

[**GetShopInfo200Response**](GetShopInfo200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetShowcase

> GetShowcase200Response GetShowcase(ctx, fsId).ShopId(shopId).Page(page).PageCount(pageCount).HideZero(hideZero).Display(display).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/dhimas-sirclo/apiclient/tokopedia"
)

func main() {
    fsId := int64(789) // int64 | Fulfillment service unique identifier
    shopId := int64(789) // int64 | Shop unique identifier
    page := int64(789) // int64 | USelect which page to show, default value: 1  (optional)
    pageCount := int64(789) // int64 | Showcase contained in each page, default value: 10  (optional)
    hideZero := true // bool | Hide showcase if product equal to zero, default value: false  (optional)
    display := "display_example" // string | To determine which showcase that you want to see. Accepted value: showcase to get user showcases only, group to get group showcases only, all to get both, default value: all  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShopAPI.GetShowcase(context.Background(), fsId).ShopId(shopId).Page(page).PageCount(pageCount).HideZero(hideZero).Display(display).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShopAPI.GetShowcase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetShowcase`: GetShowcase200Response
    fmt.Fprintf(os.Stdout, "Response from `ShopAPI.GetShowcase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsId** | **int64** | Fulfillment service unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetShowcaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shopId** | **int64** | Shop unique identifier | 
 **page** | **int64** | USelect which page to show, default value: 1  | 
 **pageCount** | **int64** | Showcase contained in each page, default value: 10  | 
 **hideZero** | **bool** | Hide showcase if product equal to zero, default value: false  | 
 **display** | **string** | To determine which showcase that you want to see. Accepted value: showcase to get user showcases only, group to get group showcases only, all to get both, default value: all  | 

### Return type

[**GetShowcase200Response**](GetShowcase200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetShowcaseAllEtalase

> GetShowcaseAllEtalase200Response GetShowcaseAllEtalase(ctx, fsId).ShopId(shopId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/dhimas-sirclo/apiclient/tokopedia"
)

func main() {
    fsId := int64(789) // int64 | Fulfillment service unique identifier
    shopId := int64(789) // int64 | Shop unique identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShopAPI.GetShowcaseAllEtalase(context.Background(), fsId).ShopId(shopId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShopAPI.GetShowcaseAllEtalase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetShowcaseAllEtalase`: GetShowcaseAllEtalase200Response
    fmt.Fprintf(os.Stdout, "Response from `ShopAPI.GetShowcaseAllEtalase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsId** | **int64** | Fulfillment service unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetShowcaseAllEtalaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shopId** | **int64** | Shop unique identifier | 

### Return type

[**GetShowcaseAllEtalase200Response**](GetShowcaseAllEtalase200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateShopStatus

> UpdateShopStatusDefaultResponse UpdateShopStatus(ctx, fsId).UpdateShopStatusRequest(updateShopStatusRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/dhimas-sirclo/apiclient/tokopedia"
)

func main() {
    fsId := int64(789) // int64 | Fulfillment service unique identifier
    updateShopStatusRequest := *openapiclient.NewUpdateShopStatusRequest(int64(123), "Action_example") // UpdateShopStatusRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShopAPI.UpdateShopStatus(context.Background(), fsId).UpdateShopStatusRequest(updateShopStatusRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShopAPI.UpdateShopStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateShopStatus`: UpdateShopStatusDefaultResponse
    fmt.Fprintf(os.Stdout, "Response from `ShopAPI.UpdateShopStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsId** | **int64** | Fulfillment service unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateShopStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateShopStatusRequest** | [**UpdateShopStatusRequest**](UpdateShopStatusRequest.md) |  | 

### Return type

[**UpdateShopStatusDefaultResponse**](UpdateShopStatusDefaultResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateShowcase

> DeleteShowcase200Response UpdateShowcase(ctx, fsId).ShopId(shopId).UpdateShowcaseRequest(updateShowcaseRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/dhimas-sirclo/apiclient/tokopedia"
)

func main() {
    fsId := int64(789) // int64 | Fulfillment service unique identifier
    shopId := int64(789) // int64 | Shop unique identifier
    updateShowcaseRequest := *openapiclient.NewUpdateShowcaseRequest(int64(123), "Name_example") // UpdateShowcaseRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShopAPI.UpdateShowcase(context.Background(), fsId).ShopId(shopId).UpdateShowcaseRequest(updateShowcaseRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShopAPI.UpdateShowcase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateShowcase`: DeleteShowcase200Response
    fmt.Fprintf(os.Stdout, "Response from `ShopAPI.UpdateShowcase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsId** | **int64** | Fulfillment service unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateShowcaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shopId** | **int64** | Shop unique identifier | 
 **updateShowcaseRequest** | [**UpdateShowcaseRequest**](UpdateShowcaseRequest.md) |  | 

### Return type

[**DeleteShowcase200Response**](DeleteShowcase200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

