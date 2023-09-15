# \OrderAPI

All URIs are relative to *https://api2.connexi.id.dmmy.me/v1/open/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpsertOrders**](OrderAPI.md#UpsertOrders) | **Post** /orders/upsert/{channelCode} | 



## UpsertOrders

> []Order UpsertOrders(ctx, channelCode).OrderInput(orderInput).Execute()





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
    orderInput := []openapiclient.OrderInput{*openapiclient.NewOrderInput("OrderDate_example", "OrderId_example", "CustomerReference_example", "CurrencyCode_example", "WarehouseId_example", "PaymentMethod_example", "DeliveryName_example", "DeliveryStreetAddress_example", "DeliverySuburb_example", "DeliveryCity_example", "DeliveryRegion_example", "DeliveryPostCode_example", "DeliveryCountry_example", "DeliveryMethod_example", "OrderStatus_example", float64(123), float64(123), float64(123), float64(123), float64(123), float64(123), float64(123), "PhoneNumber_example", "ShopId_example", "Provider_example", "Service_example", false, []openapiclient.LineItemInput{*openapiclient.NewLineItemInput("OrderItemId_example", "Sku_example", "ProductName_example", int32(123), float64(123), float64(123), float64(123), float64(123), float64(123))})} // []OrderInput | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderAPI.UpsertOrders(context.Background(), channelCode).OrderInput(orderInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.UpsertOrders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpsertOrders`: []Order
    fmt.Fprintf(os.Stdout, "Response from `OrderAPI.UpsertOrders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelCode** | **string** | Channel code | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpsertOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orderInput** | [**[]OrderInput**](OrderInput.md) |  | 

### Return type

[**[]Order**](Order.md)

### Authorization

[appKey](../README.md#appKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

