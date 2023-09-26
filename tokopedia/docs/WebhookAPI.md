# \WebhookAPI

All URIs are relative to *https://fs.tokopedia.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrderWebhook**](WebhookAPI.md#GetOrderWebhook) | **Get** /v1/order/{order_id}/fs/{fs_id}/webhook | 
[**GetRegisteredWebhooks**](WebhookAPI.md#GetRegisteredWebhooks) | **Get** /v1/fs/{fs_id} | 
[**RegisterWebhook**](WebhookAPI.md#RegisterWebhook) | **Post** /v1/fs/{fs_id}/register | 
[**TriggerWebhook**](WebhookAPI.md#TriggerWebhook) | **Post** /v1/fs/{fs_id}/trigger | 



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
    resp, r, err := apiClient.WebhookAPI.GetOrderWebhook(context.Background(), orderId, fsId).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.GetOrderWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrderWebhook`: GetOrderWebhook200Response
    fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.GetOrderWebhook`: %v\n", resp)
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


## GetRegisteredWebhooks

> GetRegisteredWebhooks200Response GetRegisteredWebhooks(ctx, fsId).Execute()





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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhookAPI.GetRegisteredWebhooks(context.Background(), fsId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.GetRegisteredWebhooks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRegisteredWebhooks`: GetRegisteredWebhooks200Response
    fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.GetRegisteredWebhooks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsId** | **int64** | Fulfillment service unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRegisteredWebhooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetRegisteredWebhooks200Response**](GetRegisteredWebhooks200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/html, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterWebhook

> RegisterWebhookDefaultResponse RegisterWebhook(ctx, fsId).RegisterWebhookRequest(registerWebhookRequest).Execute()





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
    registerWebhookRequest := *openapiclient.NewRegisterWebhookRequest(int64(123), "OrderNotificationUrl_example", "OrderCancellationUrl_example", "OrderStatusUrl_example", "OrderRequestCancellationUrl_example", "ChatNotificationUrl_example", "ProductCreationUrl_example", "ProductChangesUrl_example", "CampaignNotificationUrl_example", "WebhookSecret_example") // RegisterWebhookRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhookAPI.RegisterWebhook(context.Background(), fsId).RegisterWebhookRequest(registerWebhookRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.RegisterWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegisterWebhook`: RegisterWebhookDefaultResponse
    fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.RegisterWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsId** | **int64** | Fulfillment service unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegisterWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **registerWebhookRequest** | [**RegisterWebhookRequest**](RegisterWebhookRequest.md) |  | 

### Return type

[**RegisterWebhookDefaultResponse**](RegisterWebhookDefaultResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TriggerWebhook

> TriggerWebhookDefaultResponse TriggerWebhook(ctx, fsId).TriggerWebhookRequest(triggerWebhookRequest).Execute()





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
    resp, r, err := apiClient.WebhookAPI.TriggerWebhook(context.Background(), fsId).TriggerWebhookRequest(triggerWebhookRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.TriggerWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TriggerWebhook`: TriggerWebhookDefaultResponse
    fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.TriggerWebhook`: %v\n", resp)
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

[**TriggerWebhookDefaultResponse**](TriggerWebhookDefaultResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

