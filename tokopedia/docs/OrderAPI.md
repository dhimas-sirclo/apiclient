# \OrderAPI

All URIs are relative to *https://fs.tokopedia.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrderWebhook**](OrderAPI.md#GetOrderWebhook) | **Get** /v1/order/{order_id}/fs/{fs_id}/webhook | 
[**TriggerWebhook**](OrderAPI.md#TriggerWebhook) | **Post** /v1/fs/{fs_id}/trigger | 



## GetOrderWebhook

> GetOrderWebhook200Response GetOrderWebhook(ctx, orderId, fsId).Type_(type_).Execute()





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
    orderId := int64(789) // int64 | Order service unique identifier
    fsId := int64(789) // int64 | Fulfillment service unique identifier
    type_ := "type__example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderAPI.GetOrderWebhook(context.Background(), orderId, fsId).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.GetOrderWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrderWebhook`: GetOrderWebhook200Response
    fmt.Fprintf(os.Stdout, "Response from `OrderAPI.GetOrderWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **int64** | Order service unique identifier | 
**fsId** | **int64** | Fulfillment service unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **type_** | **string** |  | 

### Return type

[**GetOrderWebhook200Response**](GetOrderWebhook200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TriggerWebhook

> TriggerWebhook200Response TriggerWebhook(ctx, fsId).TriggerWebhookRequest(triggerWebhookRequest).Execute()





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
    triggerWebhookRequest := *openapiclient.NewTriggerWebhookRequest(int64(123), "Type_example", "Url_example", false) // TriggerWebhookRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderAPI.TriggerWebhook(context.Background(), fsId).TriggerWebhookRequest(triggerWebhookRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.TriggerWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TriggerWebhook`: TriggerWebhook200Response
    fmt.Fprintf(os.Stdout, "Response from `OrderAPI.TriggerWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsId** | **int64** | Fulfillment service unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiTriggerWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **triggerWebhookRequest** | [**TriggerWebhookRequest**](TriggerWebhookRequest.md) |  | 

### Return type

[**TriggerWebhook200Response**](TriggerWebhook200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

