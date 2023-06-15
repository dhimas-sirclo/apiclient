# \DefaultApi

All URIs are relative to *https://fs.tokopedia.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InventoryV1FsFsIdProductInfoGet**](DefaultApi.md#InventoryV1FsFsIdProductInfoGet) | **Get** /inventory/v1/fs/{fs_id}/product/info | 
[**InventoryV2FsFsIdCategoryGetVariantGet**](DefaultApi.md#InventoryV2FsFsIdCategoryGetVariantGet) | **Get** /inventory/v2/fs/{fs_id}/category/get_variant | 



## InventoryV1FsFsIdProductInfoGet

> GetProductResponse InventoryV1FsFsIdProductInfoGet(ctx, fsId).Product(product).ProductUrl(productUrl).Sku(sku).ShopId(shopId).Page(page).PerPage(perPage).LastSort(lastSort).Execute()





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
    product := "product_example" // string | Can input more than one product_id (optional)
    productUrl := "productUrl_example" // string | Can input more than one product_url (optional)
    sku := "sku_example" // string | Product’s SKU (optional)
    shopId := int64(789) // int64 | Shop Identifier (optional)
    page := int64(789) // int64 | Determine which page the order list should start. The minimal value is 1. Page (required if shop_id is filled) (optional)
    perPage := int64(789) // int64 | Page per item (required if shop_id is filled). Maximun items are 50 for 1 page (optional)
    lastSort := "lastSort_example" // string | This parameter is required when the product exceeds 10.000 products (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.InventoryV1FsFsIdProductInfoGet(context.Background(), fsId).Product(product).ProductUrl(productUrl).Sku(sku).ShopId(shopId).Page(page).PerPage(perPage).LastSort(lastSort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.InventoryV1FsFsIdProductInfoGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InventoryV1FsFsIdProductInfoGet`: GetProductResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.InventoryV1FsFsIdProductInfoGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsId** | **int64** | Fulfillment service unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryV1FsFsIdProductInfoGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **product** | **string** | Can input more than one product_id | 
 **productUrl** | **string** | Can input more than one product_url | 
 **sku** | **string** | Product’s SKU | 
 **shopId** | **int64** | Shop Identifier | 
 **page** | **int64** | Determine which page the order list should start. The minimal value is 1. Page (required if shop_id is filled) | 
 **perPage** | **int64** | Page per item (required if shop_id is filled). Maximun items are 50 for 1 page | 
 **lastSort** | **string** | This parameter is required when the product exceeds 10.000 products | 

### Return type

[**GetProductResponse**](GetProductResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryV2FsFsIdCategoryGetVariantGet

> GetVariantResponse InventoryV2FsFsIdCategoryGetVariantGet(ctx, fsId).CatId(catId).Execute()





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
    catId := int64(789) // int64 | Category unique identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.InventoryV2FsFsIdCategoryGetVariantGet(context.Background(), fsId).CatId(catId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.InventoryV2FsFsIdCategoryGetVariantGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InventoryV2FsFsIdCategoryGetVariantGet`: GetVariantResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.InventoryV2FsFsIdCategoryGetVariantGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsId** | **int64** | Fulfillment service unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryV2FsFsIdCategoryGetVariantGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **catId** | **int64** | Category unique identifier | 

### Return type

[**GetVariantResponse**](GetVariantResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

