# \ProductAPI

All URIs are relative to *https://api2.connexi.id.dmmy.me/v1/open/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpsertProductCategories**](ProductAPI.md#UpsertProductCategories) | **Post** /categories/upsert/{channelCode} | 
[**UpsertProducts**](ProductAPI.md#UpsertProducts) | **Post** /products/upsert/{channelCode} | 



## UpsertProductCategories

> UpsertProductCategoriesResponse UpsertProductCategories(ctx, channelCode).ProductCategoryInput(productCategoryInput).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/dhimas-sirclo/apiclient/sirclo"
)

func main() {
    channelCode := "channelCode_example" // string | Channel code
    productCategoryInput := []openapiclient.ProductCategoryInput{*openapiclient.NewProductCategoryInput("CategoryId_example", "CategoryName_example", false, []openapiclient.ProductCategoryAttributeInput{*openapiclient.NewProductCategoryAttributeInput("AttributeId_example", "FieldName_example", "InputType_example", false, false)})} // []ProductCategoryInput | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductAPI.UpsertProductCategories(context.Background(), channelCode).ProductCategoryInput(productCategoryInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.UpsertProductCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpsertProductCategories`: UpsertProductCategoriesResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductAPI.UpsertProductCategories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelCode** | **string** | Channel code | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpsertProductCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **productCategoryInput** | [**[]ProductCategoryInput**](ProductCategoryInput.md) |  | 

### Return type

[**UpsertProductCategoriesResponse**](UpsertProductCategoriesResponse.md)

### Authorization

[appKey](../README.md#appKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpsertProducts

> []Product UpsertProducts(ctx, channelCode).ProductInput(productInput).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/dhimas-sirclo/apiclient/sirclo"
)

func main() {
    channelCode := "channelCode_example" // string | Channel code
    productInput := []openapiclient.ProductInput{*openapiclient.NewProductInput("Type_example", "Name_example", "Brand_example", "ProductId_example", "ProductCode_example", "ProductDescription_example", "Condition_example", "ShopId_example", "CategoryId_example", "CategoryName_example", []openapiclient.VariantInput{*openapiclient.NewVariantInput("VariantId_example", "VariantSku_example", "VariantName_example", "Status_example", *openapiclient.NewVolumeInput(NullableFloat64(123), NullableFloat64(123), NullableFloat64(123)), float64(123), []openapiclient.VariantAttributeInput{*openapiclient.NewVariantAttributeInput("Name_example", "Value_example")}, "VariantUrl_example", "CurrencyCode_example", float64(123), []string{"ImgUrls_example"})})} // []ProductInput | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductAPI.UpsertProducts(context.Background(), channelCode).ProductInput(productInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.UpsertProducts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpsertProducts`: []Product
    fmt.Fprintf(os.Stdout, "Response from `ProductAPI.UpsertProducts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelCode** | **string** | Channel code | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpsertProductsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **productInput** | [**[]ProductInput**](ProductInput.md) |  | 

### Return type

[**[]Product**](Product.md)

### Authorization

[appKey](../README.md#appKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

