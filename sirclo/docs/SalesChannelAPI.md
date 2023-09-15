# \SalesChannelAPI

All URIs are relative to *https://api2.connexi.id.dmmy.me/v1/open/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RegisterSalesChannel**](SalesChannelAPI.md#RegisterSalesChannel) | **Post** /saleschannel/register | 
[**UpdateSalesChannel**](SalesChannelAPI.md#UpdateSalesChannel) | **Post** /saleschannel/update/{channelCode} | 



## RegisterSalesChannel

> RegisteredSalesChannel RegisterSalesChannel(ctx).RegisterSalesChannelInput(registerSalesChannelInput).Execute()





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
    registerSalesChannelInput := *openapiclient.NewRegisterSalesChannelInput("AppName_example", "AppChannelCode_example", "Type_example", *openapiclient.NewUrlInfoInput("https://example.com", "https://example.com", "https://example.com", "https://example.com")) // RegisterSalesChannelInput | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SalesChannelAPI.RegisterSalesChannel(context.Background()).RegisterSalesChannelInput(registerSalesChannelInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SalesChannelAPI.RegisterSalesChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegisterSalesChannel`: RegisteredSalesChannel
    fmt.Fprintf(os.Stdout, "Response from `SalesChannelAPI.RegisterSalesChannel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterSalesChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **registerSalesChannelInput** | [**RegisterSalesChannelInput**](RegisterSalesChannelInput.md) |  | 

### Return type

[**RegisteredSalesChannel**](RegisteredSalesChannel.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSalesChannel

> SalesChannel UpdateSalesChannel(ctx, channelCode).UpdateSalesChannelInput(updateSalesChannelInput).Execute()





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
    updateSalesChannelInput := *openapiclient.NewUpdateSalesChannelInput("Type_example", *openapiclient.NewUrlInfoUpdateInput("https://example.com", "https://example.com/help", "https://example.com/assets/logo.png", "https://example.com/assets/square-logo.png", "https://example.com/auth"), *openapiclient.NewUrlWebhookUpdateInput("https://example.com/connect", "https://example.com/disconnect", "https://example.com/orders/accept", "https://example.com/orders/pack", "https://example.com/orders/cancel", "https://example.com/stocks")) // UpdateSalesChannelInput | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SalesChannelAPI.UpdateSalesChannel(context.Background(), channelCode).UpdateSalesChannelInput(updateSalesChannelInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SalesChannelAPI.UpdateSalesChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSalesChannel`: SalesChannel
    fmt.Fprintf(os.Stdout, "Response from `SalesChannelAPI.UpdateSalesChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelCode** | **string** | Channel code | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSalesChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateSalesChannelInput** | [**UpdateSalesChannelInput**](UpdateSalesChannelInput.md) |  | 

### Return type

[**SalesChannel**](SalesChannel.md)

### Authorization

[appKey](../README.md#appKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

