# \OrderAPI

All URIs are relative to *https://fs.tokopedia.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcceptOrder**](OrderAPI.md#AcceptOrder) | **Post** /v1/order/{order_id}/fs/{fs_id}/ack | 
[**ConfirmShipping**](OrderAPI.md#ConfirmShipping) | **Post** /v1/order/{order_id}/fs/{fs_id}/status | 
[**GetOrderCobCod**](OrderAPI.md#GetOrderCobCod) | **Get** /v1/fs/{fs_id}/fulfillment_order | 
[**GetOrderWebhook**](OrderAPI.md#GetOrderWebhook) | **Get** /v1/order/{order_id}/fs/{fs_id}/webhook | 
[**GetResolutionTicket**](OrderAPI.md#GetResolutionTicket) | **Get** /resolution/v1/fs/{fs_id}/ticket | 
[**GetShippingLabel**](OrderAPI.md#GetShippingLabel) | **Get** /v1/order/{order_id}/fs/{fs_id}/shipping-label | 
[**GetSingleOrder**](OrderAPI.md#GetSingleOrder) | **Get** /v2/fs/{fs_id}/order | 
[**RejectOrder**](OrderAPI.md#RejectOrder) | **Post** /v1/order/{order_id}/fs/{fs_id}/nack | 
[**RejectRequestCancel**](OrderAPI.md#RejectRequestCancel) | **Post** /v1/order/{order_id}/fs/{fs_id}/reject-cancel | 
[**RequestPartialOrderFulfillment**](OrderAPI.md#RequestPartialOrderFulfillment) | **Post** /v1/order/{order_id}/fs/{fs_id}/pof/request | 
[**RequestPickup**](OrderAPI.md#RequestPickup) | **Post** /inventory/v1/fs/{fs_id}/pick-up | 
[**TriggerWebhook**](OrderAPI.md#TriggerWebhook) | **Post** /v1/fs/{fs_id}/trigger | 



## AcceptOrder

> UpdateShopStatusDefaultResponseData AcceptOrder(ctx, orderId, fsId).WarehouseId(warehouseId).Execute()





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
    warehouseId := int64(789) // int64 | Warehouse unique identifier (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderAPI.AcceptOrder(context.Background(), orderId, fsId).WarehouseId(warehouseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.AcceptOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcceptOrder`: UpdateShopStatusDefaultResponseData
    fmt.Fprintf(os.Stdout, "Response from `OrderAPI.AcceptOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **int64** | Order service unique identifier | 
**fsId** | **int64** | Fulfillment service unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcceptOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **warehouseId** | **int64** | Warehouse unique identifier | 

### Return type

[**UpdateShopStatusDefaultResponseData**](UpdateShopStatusDefaultResponseData.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfirmShipping

> UpdateShopStatusDefaultResponseData ConfirmShipping(ctx, orderId, fsId).ConfirmShippingRequest(confirmShippingRequest).Execute()





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
    confirmShippingRequest := *openapiclient.NewConfirmShippingRequest(int64(123), "ShippingRefNum_example") // ConfirmShippingRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderAPI.ConfirmShipping(context.Background(), orderId, fsId).ConfirmShippingRequest(confirmShippingRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.ConfirmShipping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConfirmShipping`: UpdateShopStatusDefaultResponseData
    fmt.Fprintf(os.Stdout, "Response from `OrderAPI.ConfirmShipping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **int64** | Order service unique identifier | 
**fsId** | **int64** | Fulfillment service unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiConfirmShippingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **confirmShippingRequest** | [**ConfirmShippingRequest**](ConfirmShippingRequest.md) |  | 

### Return type

[**UpdateShopStatusDefaultResponseData**](UpdateShopStatusDefaultResponseData.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrderCobCod

> GetOrderCobCod200Response GetOrderCobCod(ctx, fsId).OrderId(orderId).ShopId(shopId).WarehouseId(warehouseId).PerPage(perPage).FirstOrderId(firstOrderId).NextOrderId(nextOrderId).Execute()





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
    orderId := int64(789) // int64 | Order unique identifier. Will be prioritized over other parameter. Required when there is no other parameter (optional)
    shopId := int64(789) // int64 | Shop unique identifier where the order belongs. Used when get multiple order ID (optional)
    warehouseId := int64(789) // int64 | Warehouse unique identifier where the order belongs. Used when get multiple order ID (optional)
    perPage := int64(789) // int64 | Number of order in one request. Default value is 10. Can be use together with warehouse_id or shop_id parameter (optional)
    firstOrderId := int64(789) // int64 | Order ID that can be use get data after this order ID sequential by order created time (optional)
    nextOrderId := int64(789) // int64 | Order ID that can be use get data before this order ID sequential by order created time (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderAPI.GetOrderCobCod(context.Background(), fsId).OrderId(orderId).ShopId(shopId).WarehouseId(warehouseId).PerPage(perPage).FirstOrderId(firstOrderId).NextOrderId(nextOrderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.GetOrderCobCod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrderCobCod`: GetOrderCobCod200Response
    fmt.Fprintf(os.Stdout, "Response from `OrderAPI.GetOrderCobCod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsId** | **int64** | Fulfillment service unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderCobCodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orderId** | **int64** | Order unique identifier. Will be prioritized over other parameter. Required when there is no other parameter | 
 **shopId** | **int64** | Shop unique identifier where the order belongs. Used when get multiple order ID | 
 **warehouseId** | **int64** | Warehouse unique identifier where the order belongs. Used when get multiple order ID | 
 **perPage** | **int64** | Number of order in one request. Default value is 10. Can be use together with warehouse_id or shop_id parameter | 
 **firstOrderId** | **int64** | Order ID that can be use get data after this order ID sequential by order created time | 
 **nextOrderId** | **int64** | Order ID that can be use get data before this order ID sequential by order created time | 

### Return type

[**GetOrderCobCod200Response**](GetOrderCobCod200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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


## GetResolutionTicket

> GetResolutionTicket200Response GetResolutionTicket(ctx, fsId).ShopId(shopId).FromDate(fromDate).ToDate(toDate).Execute()





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
    fromDate := "fromDate_example" // string | Shop unique identifier (format: YYYY-MM-DD) 
    toDate := "toDate_example" // string | Shop unique identifier (format: YYYY-MM-DD) 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderAPI.GetResolutionTicket(context.Background(), fsId).ShopId(shopId).FromDate(fromDate).ToDate(toDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.GetResolutionTicket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetResolutionTicket`: GetResolutionTicket200Response
    fmt.Fprintf(os.Stdout, "Response from `OrderAPI.GetResolutionTicket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsId** | **int64** | Fulfillment service unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResolutionTicketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shopId** | **int64** | Shop unique identifier | 
 **fromDate** | **string** | Shop unique identifier (format: YYYY-MM-DD)  | 
 **toDate** | **string** | Shop unique identifier (format: YYYY-MM-DD)  | 

### Return type

[**GetResolutionTicket200Response**](GetResolutionTicket200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetShippingLabel

> string GetShippingLabel(ctx, orderId, fsId).Printed(printed).Execute()





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
    printed := int64(789) // int64 | 0 or 1. Default value is 1. When flag set to 1 then the seller dashboard will show the order as already printed (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderAPI.GetShippingLabel(context.Background(), orderId, fsId).Printed(printed).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.GetShippingLabel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetShippingLabel`: string
    fmt.Fprintf(os.Stdout, "Response from `OrderAPI.GetShippingLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **int64** | Order service unique identifier | 
**fsId** | **int64** | Fulfillment service unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetShippingLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **printed** | **int64** | 0 or 1. Default value is 1. When flag set to 1 then the seller dashboard will show the order as already printed | 

### Return type

**string**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSingleOrder

> GetSingleOrder200Response GetSingleOrder(ctx, fsId).OrderId(orderId).InvoiceNum(invoiceNum).Execute()





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
    orderId := int64(789) // int64 | Order unique identifier
    invoiceNum := "invoiceNum_example" // string | Invoice Number (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderAPI.GetSingleOrder(context.Background(), fsId).OrderId(orderId).InvoiceNum(invoiceNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.GetSingleOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSingleOrder`: GetSingleOrder200Response
    fmt.Fprintf(os.Stdout, "Response from `OrderAPI.GetSingleOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsId** | **int64** | Fulfillment service unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSingleOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orderId** | **int64** | Order unique identifier | 
 **invoiceNum** | **string** | Invoice Number | 

### Return type

[**GetSingleOrder200Response**](GetSingleOrder200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RejectOrder

> UpdateShopStatusDefaultResponseData RejectOrder(ctx, orderId, fsId).RejectOrderRequest(rejectOrderRequest).Execute()





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
    rejectOrderRequest := *openapiclient.NewRejectOrderRequest() // RejectOrderRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderAPI.RejectOrder(context.Background(), orderId, fsId).RejectOrderRequest(rejectOrderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.RejectOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RejectOrder`: UpdateShopStatusDefaultResponseData
    fmt.Fprintf(os.Stdout, "Response from `OrderAPI.RejectOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **int64** | Order service unique identifier | 
**fsId** | **int64** | Fulfillment service unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiRejectOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **rejectOrderRequest** | [**RejectOrderRequest**](RejectOrderRequest.md) |  | 

### Return type

[**UpdateShopStatusDefaultResponseData**](UpdateShopStatusDefaultResponseData.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RejectRequestCancel

> RejectRequestCancel200Response RejectRequestCancel(ctx, orderId, fsId).ShopId(shopId).Execute()





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
    shopId := int64(789) // int64 | Shop unique identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderAPI.RejectRequestCancel(context.Background(), orderId, fsId).ShopId(shopId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.RejectRequestCancel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RejectRequestCancel`: RejectRequestCancel200Response
    fmt.Fprintf(os.Stdout, "Response from `OrderAPI.RejectRequestCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **int64** | Order service unique identifier | 
**fsId** | **int64** | Fulfillment service unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiRejectRequestCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **shopId** | **int64** | Shop unique identifier | 

### Return type

[**RejectRequestCancel200Response**](RejectRequestCancel200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestPartialOrderFulfillment

> RejectRequestCancel200Response RequestPartialOrderFulfillment(ctx, orderId, fsId).RequestPartialOrderFulfillmentRequest(requestPartialOrderFulfillmentRequest).Execute()





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
    requestPartialOrderFulfillmentRequest := *openapiclient.NewRequestPartialOrderFulfillmentRequest() // RequestPartialOrderFulfillmentRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderAPI.RequestPartialOrderFulfillment(context.Background(), orderId, fsId).RequestPartialOrderFulfillmentRequest(requestPartialOrderFulfillmentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.RequestPartialOrderFulfillment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RequestPartialOrderFulfillment`: RejectRequestCancel200Response
    fmt.Fprintf(os.Stdout, "Response from `OrderAPI.RequestPartialOrderFulfillment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **int64** | Order service unique identifier | 
**fsId** | **int64** | Fulfillment service unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestPartialOrderFulfillmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestPartialOrderFulfillmentRequest** | [**RequestPartialOrderFulfillmentRequest**](RequestPartialOrderFulfillmentRequest.md) |  | 

### Return type

[**RejectRequestCancel200Response**](RejectRequestCancel200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestPickup

> RequestPickup200Response RequestPickup(ctx, fsId).RequestPickupRequest(requestPickupRequest).Execute()





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
    requestPickupRequest := *openapiclient.NewRequestPickupRequest(int64(123), int64(123)) // RequestPickupRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderAPI.RequestPickup(context.Background(), fsId).RequestPickupRequest(requestPickupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.RequestPickup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RequestPickup`: RequestPickup200Response
    fmt.Fprintf(os.Stdout, "Response from `OrderAPI.RequestPickup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsId** | **int64** | Fulfillment service unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestPickupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestPickupRequest** | [**RequestPickupRequest**](RequestPickupRequest.md) |  | 

### Return type

[**RequestPickup200Response**](RequestPickup200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
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

