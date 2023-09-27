# \BundleAPI

All URIs are relative to *https://fs.tokopedia.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelBundle**](BundleAPI.md#CancelBundle) | **Patch** /v1/products/bundle/fs/{fs_id}/edit | 
[**CreateBundle**](BundleAPI.md#CreateBundle) | **Post** /v1/products/bundle/fs/{fs_id}/create | 
[**GetBundleInfo**](BundleAPI.md#GetBundleInfo) | **Get** /v1/products/bundle/fs/{fs_id}/info | 
[**GetBundleList**](BundleAPI.md#GetBundleList) | **Get** /v1/products/bundle/fs/{fs_id}/list | 



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
    resp, r, err := apiClient.BundleAPI.CancelBundle(context.Background(), fsId).ShopId(shopId).CancelBundleRequest(cancelBundleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BundleAPI.CancelBundle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelBundle`: CancelBundle200Response
    fmt.Fprintf(os.Stdout, "Response from `BundleAPI.CancelBundle`: %v\n", resp)
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
    createBundleRequest := *openapiclient.NewCreateBundleRequest(*openapiclient.NewCreateBundleRequestBundle("BundleName_example", int64(123), int64(123), int64(123), []int64{int64(123)}, []openapiclient.CreateBundleRequestBundleBundleItemsInner{*openapiclient.NewCreateBundleRequestBundleBundleItemsInner(int64(123), int64(123), int64(123))})) // CreateBundleRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BundleAPI.CreateBundle(context.Background(), fsId).ShopId(shopId).CreateBundleRequest(createBundleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BundleAPI.CreateBundle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBundle`: CreateBundle200Response
    fmt.Fprintf(os.Stdout, "Response from `BundleAPI.CreateBundle`: %v\n", resp)
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
    resp, r, err := apiClient.BundleAPI.GetBundleInfo(context.Background(), fsId).ShopId(shopId).BundleId(bundleId).ProductId(productId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BundleAPI.GetBundleInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBundleInfo`: GetBundleInfo200Response
    fmt.Fprintf(os.Stdout, "Response from `BundleAPI.GetBundleInfo`: %v\n", resp)
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
    resp, r, err := apiClient.BundleAPI.GetBundleList(context.Background(), fsId).ShopId(shopId).Type_(type_).Status(status).LastGroupId(lastGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BundleAPI.GetBundleList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBundleList`: GetBundleList200Response
    fmt.Fprintf(os.Stdout, "Response from `BundleAPI.GetBundleList`: %v\n", resp)
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

