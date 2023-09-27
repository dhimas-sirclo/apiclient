# \LogisticAPI

All URIs are relative to *https://fs.tokopedia.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetActiveCourier**](LogisticAPI.md#GetActiveCourier) | **Get** /v1/logistic/fs/{fs_id}/active-info | 
[**GetShipmentInfo**](LogisticAPI.md#GetShipmentInfo) | **Get** /v2/logistic/fs/{fs_id}/info | 
[**UpdateShipmentInfo**](LogisticAPI.md#UpdateShipmentInfo) | **Post** /v2/logistic/fs/{fs_id}/update | 



## GetActiveCourier

> GetActiveCourier200Response GetActiveCourier(ctx, fsId).ShopId(shopId).Execute()





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
    resp, r, err := apiClient.LogisticAPI.GetActiveCourier(context.Background(), fsId).ShopId(shopId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogisticAPI.GetActiveCourier``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetActiveCourier`: GetActiveCourier200Response
    fmt.Fprintf(os.Stdout, "Response from `LogisticAPI.GetActiveCourier`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsId** | **int64** | Fulfillment service unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetActiveCourierRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shopId** | **int64** | Shop unique identifier | 

### Return type

[**GetActiveCourier200Response**](GetActiveCourier200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetShipmentInfo

> GetShipmentInfo200Response GetShipmentInfo(ctx, fsId).ShopId(shopId).Execute()





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
    resp, r, err := apiClient.LogisticAPI.GetShipmentInfo(context.Background(), fsId).ShopId(shopId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogisticAPI.GetShipmentInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetShipmentInfo`: GetShipmentInfo200Response
    fmt.Fprintf(os.Stdout, "Response from `LogisticAPI.GetShipmentInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsId** | **int64** | Fulfillment service unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetShipmentInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shopId** | **int64** | Shop unique identifier | 

### Return type

[**GetShipmentInfo200Response**](GetShipmentInfo200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateShipmentInfo

> UpdateShipmentInfo200Response UpdateShipmentInfo(ctx, fsId).ShopId(shopId).RequestBody(requestBody).Execute()





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
    requestBody := map[string]map[string]int64{"key": map[string]int64{"key": int64(123)}} // map[string]map[string]int64 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogisticAPI.UpdateShipmentInfo(context.Background(), fsId).ShopId(shopId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogisticAPI.UpdateShipmentInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateShipmentInfo`: UpdateShipmentInfo200Response
    fmt.Fprintf(os.Stdout, "Response from `LogisticAPI.UpdateShipmentInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsId** | **int64** | Fulfillment service unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateShipmentInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shopId** | **int64** | Shop unique identifier | 
 **requestBody** | **map[string]map[string]int64** |  | 

### Return type

[**UpdateShipmentInfo200Response**](UpdateShipmentInfo200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

