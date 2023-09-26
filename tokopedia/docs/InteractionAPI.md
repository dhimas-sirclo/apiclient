# \InteractionAPI

All URIs are relative to *https://fs.tokopedia.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetListMessage**](InteractionAPI.md#GetListMessage) | **Get** /v1/chat/fs/{fs_id}/messages | 
[**GetListReply**](InteractionAPI.md#GetListReply) | **Get** /v1/chat/fs/{fs_id}/messages/{msg_id}/replies | 
[**InitiateChat**](InteractionAPI.md#InitiateChat) | **Get** /v1/chat/fs/{fs_id}/initiate | 
[**SendReply**](InteractionAPI.md#SendReply) | **Post** /v1/chat/fs/{fs_id}/messages/{msg_id}/reply | 



## GetListMessage

> GetListMessageDefaultResponse GetListMessage(ctx, fsId).ShopId(shopId).Page(page).PerPage(perPage).Filter(filter).Execute()





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
    shopId := int64(789) // int64 | Shop service unique identifier
    page := int64(789) // int64 | Determine which page the order list should start. Start from 1. You can get the next page when the previous page has have_next_page value true (optional)
    perPage := int64(789) // int64 | Determine how many orders will be shown per page. The maximum value is 15 (optional)
    filter := "filter_example" // string | Filter by message read status. Value between “all”, “read”, or “unread”identifier (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InteractionAPI.GetListMessage(context.Background(), fsId).ShopId(shopId).Page(page).PerPage(perPage).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InteractionAPI.GetListMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetListMessage`: GetListMessageDefaultResponse
    fmt.Fprintf(os.Stdout, "Response from `InteractionAPI.GetListMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsId** | **int64** | Fulfillment service unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shopId** | **int64** | Shop service unique identifier | 
 **page** | **int64** | Determine which page the order list should start. Start from 1. You can get the next page when the previous page has have_next_page value true | 
 **perPage** | **int64** | Determine how many orders will be shown per page. The maximum value is 15 | 
 **filter** | **string** | Filter by message read status. Value between “all”, “read”, or “unread”identifier | 

### Return type

[**GetListMessageDefaultResponse**](GetListMessageDefaultResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetListReply

> GetListReply200Response GetListReply(ctx, fsId, msgId).ShopId(shopId).Page(page).PerPage(perPage).Execute()





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
    msgId := int64(789) // int64 | Message service unique identifier
    shopId := int64(789) // int64 | Shop service unique identifier
    page := int64(789) // int64 | Determine which page the order list should start. Start from 1. You can get the next page when the previous page has have_next_page value true (optional)
    perPage := int64(789) // int64 | Determine how many orders will be shown per page. The maximum value is 15 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InteractionAPI.GetListReply(context.Background(), fsId, msgId).ShopId(shopId).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InteractionAPI.GetListReply``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetListReply`: GetListReply200Response
    fmt.Fprintf(os.Stdout, "Response from `InteractionAPI.GetListReply`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsId** | **int64** | Fulfillment service unique identifier | 
**msgId** | **int64** | Message service unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListReplyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **shopId** | **int64** | Shop service unique identifier | 
 **page** | **int64** | Determine which page the order list should start. Start from 1. You can get the next page when the previous page has have_next_page value true | 
 **perPage** | **int64** | Determine how many orders will be shown per page. The maximum value is 15 | 

### Return type

[**GetListReply200Response**](GetListReply200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InitiateChat

> InitiateChat200Response InitiateChat(ctx, fsId).OrderId(orderId).Execute()





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
    orderId := int64(789) // int64 | Order service unique identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InteractionAPI.InitiateChat(context.Background(), fsId).OrderId(orderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InteractionAPI.InitiateChat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InitiateChat`: InitiateChat200Response
    fmt.Fprintf(os.Stdout, "Response from `InteractionAPI.InitiateChat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsId** | **int64** | Fulfillment service unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiInitiateChatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orderId** | **int64** | Order service unique identifier | 

### Return type

[**InitiateChat200Response**](InitiateChat200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendReply

> SendReply200Response SendReply(ctx, fsId, msgId).SendReplyRequest(sendReplyRequest).Execute()





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
    msgId := int64(789) // int64 | Message service unique identifier
    sendReplyRequest := openapiclient.sendReply_request{SendReplyRequestOneOf: openapiclient.NewSendReplyRequestOneOf(int64(123), "Message_example")} // SendReplyRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InteractionAPI.SendReply(context.Background(), fsId, msgId).SendReplyRequest(sendReplyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InteractionAPI.SendReply``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendReply`: SendReply200Response
    fmt.Fprintf(os.Stdout, "Response from `InteractionAPI.SendReply`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsId** | **int64** | Fulfillment service unique identifier | 
**msgId** | **int64** | Message service unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendReplyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sendReplyRequest** | [**SendReplyRequest**](SendReplyRequest.md) |  | 

### Return type

[**SendReply200Response**](SendReply200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

