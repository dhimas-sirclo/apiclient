# \ProductAPI

All URIs are relative to *https://fs.tokopedia.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProductInfo**](ProductAPI.md#GetProductInfo) | **Get** /inventory/v1/fs/{fs_id}/product/info | 
[**GetVariantsByCategoryId**](ProductAPI.md#GetVariantsByCategoryId) | **Get** /inventory/v2/fs/{fs_id}/category/get_variant | 
[**GetVariantsByProductId**](ProductAPI.md#GetVariantsByProductId) | **Get** /inventory/v1/fs/{fs_id}/product/variant/{product_id} | 



## GetProductInfo

> GetProductInfoResponse GetProductInfo(ctx, fsId).Product(product).ProductUrl(productUrl).Sku(sku).ShopId(shopId).Page(page).PerPage(perPage).LastSort(lastSort).Execute()





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
    fsId := int32(56) // int32 | Fulfillment service unique identifier
    product := "product_example" // string | Can input more than one product_id (optional)
    productUrl := "productUrl_example" // string | Can input more than one product_url (optional)
    sku := "sku_example" // string | Product’s SKU (optional)
    shopId := int32(56) // int32 | Shop Identifier (optional)
    page := int32(56) // int32 | Determine which page the order list should start. The minimal value is 1. Page (required if shop_id is filled) (optional)
    perPage := int32(56) // int32 | Page per item (required if shop_id is filled). Maximun items are 50 for 1 page (optional)
    lastSort := "lastSort_example" // string | This parameter is required when the product exceeds 10.000 products (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductAPI.GetProductInfo(context.Background(), fsId).Product(product).ProductUrl(productUrl).Sku(sku).ShopId(shopId).Page(page).PerPage(perPage).LastSort(lastSort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.GetProductInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProductInfo`: GetProductInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductAPI.GetProductInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsId** | **int32** | Fulfillment service unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProductInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **product** | **string** | Can input more than one product_id | 
 **productUrl** | **string** | Can input more than one product_url | 
 **sku** | **string** | Product’s SKU | 
 **shopId** | **int32** | Shop Identifier | 
 **page** | **int32** | Determine which page the order list should start. The minimal value is 1. Page (required if shop_id is filled) | 
 **perPage** | **int32** | Page per item (required if shop_id is filled). Maximun items are 50 for 1 page | 
 **lastSort** | **string** | This parameter is required when the product exceeds 10.000 products | 

### Return type

[**GetProductInfoResponse**](GetProductInfoResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVariantsByCategoryId

> CategoryVariants GetVariantsByCategoryId(ctx, fsId).CatId(catId).Execute()





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
    fsId := int32(56) // int32 | Fulfillment service unique identifier
    catId := int32(56) // int32 | Category unique identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductAPI.GetVariantsByCategoryId(context.Background(), fsId).CatId(catId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.GetVariantsByCategoryId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVariantsByCategoryId`: CategoryVariants
    fmt.Fprintf(os.Stdout, "Response from `ProductAPI.GetVariantsByCategoryId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsId** | **int32** | Fulfillment service unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVariantsByCategoryIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **catId** | **int32** | Category unique identifier | 

### Return type

[**CategoryVariants**](CategoryVariants.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVariantsByProductId

> GetProductVariantResponse GetVariantsByProductId(ctx, fsId, productId).Execute()





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
    fsId := int32(56) // int32 | Fulfillment service unique identifier
    productId := int32(56) // int32 | Product unique identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductAPI.GetVariantsByProductId(context.Background(), fsId, productId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.GetVariantsByProductId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVariantsByProductId`: GetProductVariantResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductAPI.GetVariantsByProductId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsId** | **int32** | Fulfillment service unique identifier | 
**productId** | **int32** | Product unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVariantsByProductIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetProductVariantResponse**](GetProductVariantResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

