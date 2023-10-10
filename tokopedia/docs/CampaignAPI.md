# \CampaignAPI

All URIs are relative to *https://fs.tokopedia.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddSlashPrice**](CampaignAPI.md#AddSlashPrice) | **Post** /v1/slash-price/fs/{fs_id}/add | 
[**CancelBundle**](CampaignAPI.md#CancelBundle) | **Patch** /v1/products/bundle/fs/{fs_id}/edit | 
[**CancelSlashPrice**](CampaignAPI.md#CancelSlashPrice) | **Post** /v1/slash-price/fs/{fs_id}/cancel | 
[**CreateBundle**](CampaignAPI.md#CreateBundle) | **Post** /v1/products/bundle/fs/{fs_id}/create | 
[**GetBundleInfo**](CampaignAPI.md#GetBundleInfo) | **Get** /v1/products/bundle/fs/{fs_id}/info | 
[**GetBundleList**](CampaignAPI.md#GetBundleList) | **Get** /v1/products/bundle/fs/{fs_id}/list | 
[**UpdateSlashPrice**](CampaignAPI.md#UpdateSlashPrice) | **Post** /v1/slash-price/fs/{fs_id}/update | 
[**ViewCampaignProducts**](CampaignAPI.md#ViewCampaignProducts) | **Get** /v1/campaign/fs/{fs_id}/view | 
[**ViewSlashPrice**](CampaignAPI.md#ViewSlashPrice) | **Get** /v2/slash-price/fs/{fs_id}/view | 



## AddSlashPrice

> CancelSlashPrice200Response AddSlashPrice(ctx, fsId).ShopId(shopId).UpdateSlashPriceRequestInner(updateSlashPriceRequestInner).Execute()





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
    updateSlashPriceRequestInner := []openapiclient.UpdateSlashPriceRequestInner{*openapiclient.NewUpdateSlashPriceRequestInner(int64(123), int64(123), int64(123), int64(123))} // []UpdateSlashPriceRequestInner | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CampaignAPI.AddSlashPrice(context.Background(), fsId).ShopId(shopId).UpdateSlashPriceRequestInner(updateSlashPriceRequestInner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignAPI.AddSlashPrice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddSlashPrice`: CancelSlashPrice200Response
    fmt.Fprintf(os.Stdout, "Response from `CampaignAPI.AddSlashPrice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsId** | **int64** | Fulfillment service unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddSlashPriceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shopId** | **int64** | Shop unique identifier | 
 **updateSlashPriceRequestInner** | [**[]UpdateSlashPriceRequestInner**](UpdateSlashPriceRequestInner.md) |  | 

### Return type

[**CancelSlashPrice200Response**](CancelSlashPrice200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelBundle

> CancelBundle200Response CancelBundle(ctx, fsId).ShopId(shopId).CancelBundleRequest(cancelBundleRequest).Execute()





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
    cancelBundleRequest := *openapiclient.NewCancelBundleRequest(int64(123)) // CancelBundleRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CampaignAPI.CancelBundle(context.Background(), fsId).ShopId(shopId).CancelBundleRequest(cancelBundleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignAPI.CancelBundle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelBundle`: CancelBundle200Response
    fmt.Fprintf(os.Stdout, "Response from `CampaignAPI.CancelBundle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsId** | **int64** | Fulfillment service unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelBundleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shopId** | **int64** | Shop unique identifier | 
 **cancelBundleRequest** | [**CancelBundleRequest**](CancelBundleRequest.md) |  | 

### Return type

[**CancelBundle200Response**](CancelBundle200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelSlashPrice

> CancelSlashPrice200Response CancelSlashPrice(ctx, fsId).ShopId(shopId).CancelSlashPriceRequestInner(cancelSlashPriceRequestInner).Execute()





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
    cancelSlashPriceRequestInner := []openapiclient.CancelSlashPriceRequestInner{*openapiclient.NewCancelSlashPriceRequestInner(int64(123), int64(123))} // []CancelSlashPriceRequestInner | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CampaignAPI.CancelSlashPrice(context.Background(), fsId).ShopId(shopId).CancelSlashPriceRequestInner(cancelSlashPriceRequestInner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignAPI.CancelSlashPrice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelSlashPrice`: CancelSlashPrice200Response
    fmt.Fprintf(os.Stdout, "Response from `CampaignAPI.CancelSlashPrice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsId** | **int64** | Fulfillment service unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelSlashPriceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shopId** | **int64** | Shop unique identifier | 
 **cancelSlashPriceRequestInner** | [**[]CancelSlashPriceRequestInner**](CancelSlashPriceRequestInner.md) |  | 

### Return type

[**CancelSlashPrice200Response**](CancelSlashPrice200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBundle

> CreateBundle200Response CreateBundle(ctx, fsId).ShopId(shopId).CreateBundleRequest(createBundleRequest).Execute()





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
    createBundleRequest := *openapiclient.NewCreateBundleRequest(*openapiclient.NewCreateBundleRequestBundle("BundleName_example", int64(123), int64(123), int64(123), []int64{int64(123)}, []openapiclient.CreateBundleRequestBundleBundleItemsInner{*openapiclient.NewCreateBundleRequestBundleBundleItemsInner(int64(123), int64(123), float64(123))})) // CreateBundleRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CampaignAPI.CreateBundle(context.Background(), fsId).ShopId(shopId).CreateBundleRequest(createBundleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignAPI.CreateBundle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBundle`: CreateBundle200Response
    fmt.Fprintf(os.Stdout, "Response from `CampaignAPI.CreateBundle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsId** | **int64** | Fulfillment service unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBundleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shopId** | **int64** | Shop unique identifier | 
 **createBundleRequest** | [**CreateBundleRequest**](CreateBundleRequest.md) |  | 

### Return type

[**CreateBundle200Response**](CreateBundle200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBundleInfo

> GetBundleInfo200Response GetBundleInfo(ctx, fsId).ShopId(shopId).BundleId(bundleId).ProductId(productId).Execute()





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
    bundleId := int64(789) // int64 | Bundle IDs, could add more than one. Example: 4360,4361  (optional)
    productId := int64(789) // int64 | The product unique identifier that associated with the bundle (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CampaignAPI.GetBundleInfo(context.Background(), fsId).ShopId(shopId).BundleId(bundleId).ProductId(productId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignAPI.GetBundleInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBundleInfo`: GetBundleInfo200Response
    fmt.Fprintf(os.Stdout, "Response from `CampaignAPI.GetBundleInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsId** | **int64** | Fulfillment service unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBundleInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shopId** | **int64** | Shop unique identifier | 
 **bundleId** | **int64** | Bundle IDs, could add more than one. Example: 4360,4361  | 
 **productId** | **int64** | The product unique identifier that associated with the bundle | 

### Return type

[**GetBundleInfo200Response**](GetBundleInfo200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBundleList

> GetBundleList200Response GetBundleList(ctx, fsId).ShopId(shopId).Type_(type_).Status(status).LastGroupId(lastGroupId).Execute()





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
    type_ := int64(789) // int64 | Filter type of bundle list: * 1 - SINGLE * 2 - MULTIPLE 
    status := int64(789) // int64 | Filter status of bundle list: * 1 = ACTIVE * 2 = UPCOMING * -5 = CANCELED fill with 0 or keep it empty to show all bundle list status  (optional)
    lastGroupId := int64(789) // int64 | For pagination purpose. Fill with 0 or keep it empty to show all bundle list (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CampaignAPI.GetBundleList(context.Background(), fsId).ShopId(shopId).Type_(type_).Status(status).LastGroupId(lastGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignAPI.GetBundleList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBundleList`: GetBundleList200Response
    fmt.Fprintf(os.Stdout, "Response from `CampaignAPI.GetBundleList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsId** | **int64** | Fulfillment service unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBundleListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shopId** | **int64** | Shop unique identifier | 
 **type_** | **int64** | Filter type of bundle list: * 1 - SINGLE * 2 - MULTIPLE  | 
 **status** | **int64** | Filter status of bundle list: * 1 &#x3D; ACTIVE * 2 &#x3D; UPCOMING * -5 &#x3D; CANCELED fill with 0 or keep it empty to show all bundle list status  | 
 **lastGroupId** | **int64** | For pagination purpose. Fill with 0 or keep it empty to show all bundle list | 

### Return type

[**GetBundleList200Response**](GetBundleList200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSlashPrice

> CancelSlashPrice200Response UpdateSlashPrice(ctx, fsId).ShopId(shopId).UpdateSlashPriceRequestInner(updateSlashPriceRequestInner).Execute()





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
    updateSlashPriceRequestInner := []openapiclient.UpdateSlashPriceRequestInner{*openapiclient.NewUpdateSlashPriceRequestInner(int64(123), int64(123), int64(123), int64(123))} // []UpdateSlashPriceRequestInner | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CampaignAPI.UpdateSlashPrice(context.Background(), fsId).ShopId(shopId).UpdateSlashPriceRequestInner(updateSlashPriceRequestInner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignAPI.UpdateSlashPrice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSlashPrice`: CancelSlashPrice200Response
    fmt.Fprintf(os.Stdout, "Response from `CampaignAPI.UpdateSlashPrice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsId** | **int64** | Fulfillment service unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSlashPriceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shopId** | **int64** | Shop unique identifier | 
 **updateSlashPriceRequestInner** | [**[]UpdateSlashPriceRequestInner**](UpdateSlashPriceRequestInner.md) |  | 

### Return type

[**CancelSlashPrice200Response**](CancelSlashPrice200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ViewCampaignProducts

> ViewCampaignProducts200Response ViewCampaignProducts(ctx, fsId).ShopId(shopId).ProductId(productId).Execute()





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
    productId := int64(789) // int64 | Product unique identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CampaignAPI.ViewCampaignProducts(context.Background(), fsId).ShopId(shopId).ProductId(productId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignAPI.ViewCampaignProducts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ViewCampaignProducts`: ViewCampaignProducts200Response
    fmt.Fprintf(os.Stdout, "Response from `CampaignAPI.ViewCampaignProducts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsId** | **int64** | Fulfillment service unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiViewCampaignProductsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shopId** | **int64** | Shop unique identifier | 
 **productId** | **int64** | Product unique identifier | 

### Return type

[**ViewCampaignProducts200Response**](ViewCampaignProducts200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ViewSlashPrice

> ViewSlashPrice200Response ViewSlashPrice(ctx, fsId).ShopId(shopId).WarehouseId(warehouseId).ProductId(productId).Page(page).PerPage(perPage).Status(status).Execute()





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
    warehouseId := int64(789) // int64 | Warehouse unique identifier
    productId := int64(789) // int64 | Product unique identifier
    page := int64(789) // int64 | Determine which page the order list should start. The minimal value is 1
    perPage := int64(789) // int64 | Determine how many orders will be shown per page
    status := "status_example" // string | Filter data by slash price status. The possible values are SCHEDULED,ONGOING, and PAUSED. By default, it will show all status if we don’t include this into query parameter (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CampaignAPI.ViewSlashPrice(context.Background(), fsId).ShopId(shopId).WarehouseId(warehouseId).ProductId(productId).Page(page).PerPage(perPage).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignAPI.ViewSlashPrice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ViewSlashPrice`: ViewSlashPrice200Response
    fmt.Fprintf(os.Stdout, "Response from `CampaignAPI.ViewSlashPrice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsId** | **int64** | Fulfillment service unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiViewSlashPriceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shopId** | **int64** | Shop unique identifier | 
 **warehouseId** | **int64** | Warehouse unique identifier | 
 **productId** | **int64** | Product unique identifier | 
 **page** | **int64** | Determine which page the order list should start. The minimal value is 1 | 
 **perPage** | **int64** | Determine how many orders will be shown per page | 
 **status** | **string** | Filter data by slash price status. The possible values are SCHEDULED,ONGOING, and PAUSED. By default, it will show all status if we don’t include this into query parameter | 

### Return type

[**ViewSlashPrice200Response**](ViewSlashPrice200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

