# \ProductAPI

All URIs are relative to *https://fs.tokopedia.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckUploadStatus**](ProductAPI.md#CheckUploadStatus) | **Get** /v2/products/fs/{fs_id}/status/{upload_id} | 
[**CreateProductV3**](ProductAPI.md#CreateProductV3) | **Post** /v3/products/fs/{fs_id}/create | 
[**DeleteProduct**](ProductAPI.md#DeleteProduct) | **Delete** /v3/products/fs/{fs_id}/delete | 
[**EditProductV2**](ProductAPI.md#EditProductV2) | **Patch** /v2/products/fs/{fs_id}/edit | 
[**EditProductV3**](ProductAPI.md#EditProductV3) | **Patch** /v3/products/fs/{fs_id}/edit | 
[**GetAllCategories**](ProductAPI.md#GetAllCategories) | **Get** /inventory/v1/fs/{fs_id}/product/category | 
[**GetProductAnnotationByCategoryId**](ProductAPI.md#GetProductAnnotationByCategoryId) | **Get** /v1/fs/{fs_id}/product/annotation | 
[**GetProductDiscussion**](ProductAPI.md#GetProductDiscussion) | **Get** /v1/discussion/fs/{fs_id}/list | 
[**GetProductInfo**](ProductAPI.md#GetProductInfo) | **Get** /inventory/v1/fs/{fs_id}/product/info | 
[**GetVariantsByCategoryId**](ProductAPI.md#GetVariantsByCategoryId) | **Get** /inventory/v2/fs/{fs_id}/category/get_variant | 
[**GetVariantsByProductId**](ProductAPI.md#GetVariantsByProductId) | **Get** /inventory/v1/fs/{fs_id}/product/variant/{product_id} | 
[**SetActiveProduct**](ProductAPI.md#SetActiveProduct) | **Post** /v1/products/fs/{fs_id}/active | 
[**SetInactiveProduct**](ProductAPI.md#SetInactiveProduct) | **Post** /v1/products/fs/{fs_id}/inactive | 
[**UpdatePrice**](ProductAPI.md#UpdatePrice) | **Post** /inventory/v1/fs/{fs_id}/price/update | 
[**UpdateStockDecrement**](ProductAPI.md#UpdateStockDecrement) | **Post** /inventory/v1/fs/{fs_id}/stock/decrement | 
[**UpdateStockIncrement**](ProductAPI.md#UpdateStockIncrement) | **Post** /inventory/v1/fs/{fs_id}/stock/increment | 
[**UpdateStockOverwrite**](ProductAPI.md#UpdateStockOverwrite) | **Post** /inventory/v1/fs/{fs_id}/stock/update | 



## CheckUploadStatus

> CheckUploadStatus200Response CheckUploadStatus(ctx, fsId, uploadId).ShopId(shopId).Execute()





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
    uploadId := int64(789) // int64 | Upload id of the product to check
    shopId := int64(789) // int64 | Shop unique identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductAPI.CheckUploadStatus(context.Background(), fsId, uploadId).ShopId(shopId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.CheckUploadStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckUploadStatus`: CheckUploadStatus200Response
    fmt.Fprintf(os.Stdout, "Response from `ProductAPI.CheckUploadStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsId** | **int64** | Fulfillment service unique identifier | 
**uploadId** | **int64** | Upload id of the product to check | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckUploadStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **shopId** | **int64** | Shop unique identifier | 

### Return type

[**CheckUploadStatus200Response**](CheckUploadStatus200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateProductV3

> CreateProductV3200Response CreateProductV3(ctx, fsId).ShopId(shopId).CreateProductV3RequestInner(createProductV3RequestInner).Execute()





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
    createProductV3RequestInner := []openapiclient.CreateProductV3RequestInner{*openapiclient.NewCreateProductV3RequestInner()} // []CreateProductV3RequestInner | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductAPI.CreateProductV3(context.Background(), fsId).ShopId(shopId).CreateProductV3RequestInner(createProductV3RequestInner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.CreateProductV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProductV3`: CreateProductV3200Response
    fmt.Fprintf(os.Stdout, "Response from `ProductAPI.CreateProductV3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsId** | **int64** | Fulfillment service unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProductV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shopId** | **int64** | Shop unique identifier | 
 **createProductV3RequestInner** | [**[]CreateProductV3RequestInner**](CreateProductV3RequestInner.md) |  | 

### Return type

[**CreateProductV3200Response**](CreateProductV3200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProduct

> DeleteProductDefaultResponse DeleteProduct(ctx, fsId).ShopId(shopId).DeleteProductRequest(deleteProductRequest).Execute()





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
    deleteProductRequest := *openapiclient.NewDeleteProductRequest([]int64{int64(123)}) // DeleteProductRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductAPI.DeleteProduct(context.Background(), fsId).ShopId(shopId).DeleteProductRequest(deleteProductRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.DeleteProduct``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteProduct`: DeleteProductDefaultResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductAPI.DeleteProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsId** | **int64** | Fulfillment service unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shopId** | **int64** | Shop unique identifier | 
 **deleteProductRequest** | [**DeleteProductRequest**](DeleteProductRequest.md) |  | 

### Return type

[**DeleteProductDefaultResponse**](DeleteProductDefaultResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditProductV2

> EditProductV2200Response EditProductV2(ctx, fsId).ShopId(shopId).EditProductV2Request(editProductV2Request).Execute()





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
    editProductV2Request := *openapiclient.NewEditProductV2Request([]openapiclient.EditProductV2RequestProductsInner{*openapiclient.NewEditProductV2RequestProductsInner()}) // EditProductV2Request | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductAPI.EditProductV2(context.Background(), fsId).ShopId(shopId).EditProductV2Request(editProductV2Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.EditProductV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditProductV2`: EditProductV2200Response
    fmt.Fprintf(os.Stdout, "Response from `ProductAPI.EditProductV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsId** | **int64** | Fulfillment service unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditProductV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shopId** | **int64** | Shop unique identifier | 
 **editProductV2Request** | [**EditProductV2Request**](EditProductV2Request.md) |  | 

### Return type

[**EditProductV2200Response**](EditProductV2200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditProductV3

> EditProductV3200Response EditProductV3(ctx, fsId).ShopId(shopId).EditProductV3Request(editProductV3Request).Execute()





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
    editProductV3Request := *openapiclient.NewEditProductV3Request([]openapiclient.EditProductV3RequestProductsInner{*openapiclient.NewEditProductV3RequestProductsInner(int64(123), "Status_example", false)}) // EditProductV3Request | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductAPI.EditProductV3(context.Background(), fsId).ShopId(shopId).EditProductV3Request(editProductV3Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.EditProductV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditProductV3`: EditProductV3200Response
    fmt.Fprintf(os.Stdout, "Response from `ProductAPI.EditProductV3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsId** | **int64** | Fulfillment service unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditProductV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shopId** | **int64** | Shop unique identifier | 
 **editProductV3Request** | [**EditProductV3Request**](EditProductV3Request.md) |  | 

### Return type

[**EditProductV3200Response**](EditProductV3200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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


## GetProductDiscussion

> GetProductDiscussion200Response GetProductDiscussion(ctx, fsId).ShopId(shopId).ProductId(productId).Page(page).PerPage(perPage).Execute()





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
    productId := int64(789) // int64 | List of Product ID
    page := int64(789) // int64 | Determine which page the order list should start. The minimal value is 1 (default to 1)
    perPage := int64(789) // int64 | Determine how many orders will be shown per page. The maximal value is 10 (default to 10)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductAPI.GetProductDiscussion(context.Background(), fsId).ShopId(shopId).ProductId(productId).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.GetProductDiscussion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProductDiscussion`: GetProductDiscussion200Response
    fmt.Fprintf(os.Stdout, "Response from `ProductAPI.GetProductDiscussion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsId** | **int64** | Fulfillment service unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProductDiscussionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shopId** | **int64** | Shop unique identifier | 
 **productId** | **int64** | List of Product ID | 
 **page** | **int64** | Determine which page the order list should start. The minimal value is 1 | [default to 1]
 **perPage** | **int64** | Determine how many orders will be shown per page. The maximal value is 10 | [default to 10]

### Return type

[**GetProductDiscussion200Response**](GetProductDiscussion200Response.md)

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


## SetActiveProduct

> CancelSlashPrice200Response SetActiveProduct(ctx, fsId).ShopId(shopId).SetInactiveProductRequest(setInactiveProductRequest).Execute()





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
    setInactiveProductRequest := *openapiclient.NewSetInactiveProductRequest() // SetInactiveProductRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductAPI.SetActiveProduct(context.Background(), fsId).ShopId(shopId).SetInactiveProductRequest(setInactiveProductRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.SetActiveProduct``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetActiveProduct`: CancelSlashPrice200Response
    fmt.Fprintf(os.Stdout, "Response from `ProductAPI.SetActiveProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsId** | **int64** | Fulfillment service unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetActiveProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shopId** | **int64** | Shop unique identifier | 
 **setInactiveProductRequest** | [**SetInactiveProductRequest**](SetInactiveProductRequest.md) |  | 

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


## SetInactiveProduct

> CancelSlashPrice200Response SetInactiveProduct(ctx, fsId).ShopId(shopId).SetInactiveProductRequest(setInactiveProductRequest).Execute()





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
    setInactiveProductRequest := *openapiclient.NewSetInactiveProductRequest() // SetInactiveProductRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductAPI.SetInactiveProduct(context.Background(), fsId).ShopId(shopId).SetInactiveProductRequest(setInactiveProductRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.SetInactiveProduct``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetInactiveProduct`: CancelSlashPrice200Response
    fmt.Fprintf(os.Stdout, "Response from `ProductAPI.SetInactiveProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsId** | **int64** | Fulfillment service unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetInactiveProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shopId** | **int64** | Shop unique identifier | 
 **setInactiveProductRequest** | [**SetInactiveProductRequest**](SetInactiveProductRequest.md) |  | 

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


## UpdatePrice

> UpdatePriceDefaultResponse UpdatePrice(ctx, fsId).ShopId(shopId).UpdatePriceRequestInner(updatePriceRequestInner).WarehouseId(warehouseId).Execute()





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
    updatePriceRequestInner := []openapiclient.UpdatePriceRequestInner{*openapiclient.NewUpdatePriceRequestInner()} // []UpdatePriceRequestInner | 
    warehouseId := int64(789) // int64 | Warehouse unique identifer (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductAPI.UpdatePrice(context.Background(), fsId).ShopId(shopId).UpdatePriceRequestInner(updatePriceRequestInner).WarehouseId(warehouseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.UpdatePrice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePrice`: UpdatePriceDefaultResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductAPI.UpdatePrice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsId** | **int64** | Fulfillment service unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePriceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shopId** | **int64** | Shop unique identifier | 
 **updatePriceRequestInner** | [**[]UpdatePriceRequestInner**](UpdatePriceRequestInner.md) |  | 
 **warehouseId** | **int64** | Warehouse unique identifer | 

### Return type

[**UpdatePriceDefaultResponse**](UpdatePriceDefaultResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStockDecrement

> DecreaseStockResponse UpdateStockDecrement(ctx, fsId).ShopId(shopId).UpdateStockIncrementRequestInner(updateStockIncrementRequestInner).RequestId(requestId).WarehouseId(warehouseId).Execute()





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
    updateStockIncrementRequestInner := []openapiclient.UpdateStockIncrementRequestInner{*openapiclient.NewUpdateStockIncrementRequestInner(int64(123))} // []UpdateStockIncrementRequestInner | 
    requestId := true // bool | The ID of the request to be used for duplicate request validation. Only alphanumeric (case insensitive), ‘-‘, and ‘_’ characters are allowed with a maximum of 20 characters long (optional)
    warehouseId := int64(789) // int64 | Warehouse unique identifer (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductAPI.UpdateStockDecrement(context.Background(), fsId).ShopId(shopId).UpdateStockIncrementRequestInner(updateStockIncrementRequestInner).RequestId(requestId).WarehouseId(warehouseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.UpdateStockDecrement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateStockDecrement`: DecreaseStockResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductAPI.UpdateStockDecrement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsId** | **int64** | Fulfillment service unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStockDecrementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shopId** | **int64** | Shop unique identifier | 
 **updateStockIncrementRequestInner** | [**[]UpdateStockIncrementRequestInner**](UpdateStockIncrementRequestInner.md) |  | 
 **requestId** | **bool** | The ID of the request to be used for duplicate request validation. Only alphanumeric (case insensitive), ‘-‘, and ‘_’ characters are allowed with a maximum of 20 characters long | 
 **warehouseId** | **int64** | Warehouse unique identifer | 

### Return type

[**DecreaseStockResponse**](DecreaseStockResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStockIncrement

> IncreaseStockResponse UpdateStockIncrement(ctx, fsId).ShopId(shopId).UpdateStockIncrementRequestInner(updateStockIncrementRequestInner).RequestId(requestId).WarehouseId(warehouseId).Execute()





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
    updateStockIncrementRequestInner := []openapiclient.UpdateStockIncrementRequestInner{*openapiclient.NewUpdateStockIncrementRequestInner(int64(123))} // []UpdateStockIncrementRequestInner | 
    requestId := true // bool | The ID of the request to be used for duplicate request validation. Only alphanumeric (case insensitive), ‘-‘, and ‘_’ characters are allowed with a maximum of 20 characters long (optional)
    warehouseId := int64(789) // int64 | Warehouse unique identifer (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductAPI.UpdateStockIncrement(context.Background(), fsId).ShopId(shopId).UpdateStockIncrementRequestInner(updateStockIncrementRequestInner).RequestId(requestId).WarehouseId(warehouseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.UpdateStockIncrement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateStockIncrement`: IncreaseStockResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductAPI.UpdateStockIncrement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsId** | **int64** | Fulfillment service unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStockIncrementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shopId** | **int64** | Shop unique identifier | 
 **updateStockIncrementRequestInner** | [**[]UpdateStockIncrementRequestInner**](UpdateStockIncrementRequestInner.md) |  | 
 **requestId** | **bool** | The ID of the request to be used for duplicate request validation. Only alphanumeric (case insensitive), ‘-‘, and ‘_’ characters are allowed with a maximum of 20 characters long | 
 **warehouseId** | **int64** | Warehouse unique identifer | 

### Return type

[**IncreaseStockResponse**](IncreaseStockResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStockOverwrite

> UpdateStockResponse UpdateStockOverwrite(ctx, fsId).ShopId(shopId).UpdateStockInput(updateStockInput).BypassUpdateProductStatus(bypassUpdateProductStatus).WarehouseId(warehouseId).Execute()





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
    resp, r, err := apiClient.ProductAPI.UpdateStockOverwrite(context.Background(), fsId).ShopId(shopId).UpdateStockInput(updateStockInput).BypassUpdateProductStatus(bypassUpdateProductStatus).WarehouseId(warehouseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.UpdateStockOverwrite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateStockOverwrite`: UpdateStockResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductAPI.UpdateStockOverwrite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsId** | **int64** | Fulfillment service unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStockOverwriteRequest struct via the builder pattern


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

