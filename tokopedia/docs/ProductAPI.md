# \ProductAPI

All URIs are relative to *https://fs.tokopedia.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllCategories**](ProductAPI.md#GetAllCategories) | **Get** /inventory/v1/fs/{fs_id}/product/category | 
[**GetProductAnnotationByCategoryId**](ProductAPI.md#GetProductAnnotationByCategoryId) | **Get** /v1/fs/{fs_id}/product/annotation | 
[**GetProductInfo**](ProductAPI.md#GetProductInfo) | **Get** /inventory/v1/fs/{fs_id}/product/info | 
[**GetVariantsByCategoryId**](ProductAPI.md#GetVariantsByCategoryId) | **Get** /inventory/v2/fs/{fs_id}/category/get_variant | 
[**GetVariantsByProductId**](ProductAPI.md#GetVariantsByProductId) | **Get** /inventory/v1/fs/{fs_id}/product/variant/{product_id} | 
[**UpdateStock**](ProductAPI.md#UpdateStock) | **Post** /inventory/v1/fs/{fs_id}/stock/update | 



## GetAllCategories

> GetAllCategories200Response GetAllCategories(ctx, fsId).Keyword(keyword).Execute()





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
    keyword := "keyword_example" // string | Keyword or Product Name to get recommended category. Leave it blank to get full categories (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductAPI.GetAllCategories(context.Background(), fsId).Keyword(keyword).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.GetAllCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllCategories`: GetAllCategories200Response
    fmt.Fprintf(os.Stdout, "Response from `ProductAPI.GetAllCategories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsId** | **int64** | Fulfillment service unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **keyword** | **string** | Keyword or Product Name to get recommended category. Leave it blank to get full categories | 

### Return type

[**GetAllCategories200Response**](GetAllCategories200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/html, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProductAnnotationByCategoryId

> GetProductAnnotationByCategoryId200Response GetProductAnnotationByCategoryId(ctx, fsId).CatId(catId).Execute()





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
    resp, r, err := apiClient.ProductAPI.GetProductAnnotationByCategoryId(context.Background(), fsId).CatId(catId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.GetProductAnnotationByCategoryId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProductAnnotationByCategoryId`: GetProductAnnotationByCategoryId200Response
    fmt.Fprintf(os.Stdout, "Response from `ProductAPI.GetProductAnnotationByCategoryId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsId** | **int64** | Fulfillment service unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProductAnnotationByCategoryIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **catId** | **int64** | Category unique identifier | 

### Return type

[**GetProductAnnotationByCategoryId200Response**](GetProductAnnotationByCategoryId200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
**fsId** | **int64** | Fulfillment service unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProductInfoRequest struct via the builder pattern


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
    fsId := int64(789) // int64 | Fulfillment service unique identifier
    catId := int64(789) // int64 | Category unique identifier

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
**fsId** | **int64** | Fulfillment service unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVariantsByCategoryIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **catId** | **int64** | Category unique identifier | 

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
    fsId := int64(789) // int64 | Fulfillment service unique identifier
    productId := int64(789) // int64 | Product unique identifier

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
**fsId** | **int64** | Fulfillment service unique identifier | 
**productId** | **int64** | Product unique identifier | 

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


## UpdateStock

> UpdateStockResponse UpdateStock(ctx, fsId).ShopId(shopId).UpdateStockInput(updateStockInput).BypassUpdateProductStatus(bypassUpdateProductStatus).WarehouseId(warehouseId).Execute()





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
    shopId := int64(789) // int64 | 
    updateStockInput := []openapiclient.UpdateStockInput{*openapiclient.NewUpdateStockInput(int64(123))} // []UpdateStockInput | 
    bypassUpdateProductStatus := true // bool |  (optional)
    warehouseId := int64(789) // int64 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductAPI.UpdateStock(context.Background(), fsId).ShopId(shopId).UpdateStockInput(updateStockInput).BypassUpdateProductStatus(bypassUpdateProductStatus).WarehouseId(warehouseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.UpdateStock``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateStock`: UpdateStockResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductAPI.UpdateStock`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsId** | **int64** | Fulfillment service unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shopId** | **int64** |  | 
 **updateStockInput** | [**[]UpdateStockInput**](UpdateStockInput.md) |  | 
 **bypassUpdateProductStatus** | **bool** |  | 
 **warehouseId** | **int64** |  | 

### Return type

[**UpdateStockResponse**](UpdateStockResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

