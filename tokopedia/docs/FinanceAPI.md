# \FinanceAPI

All URIs are relative to *https://fs.tokopedia.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSaldoHistory**](FinanceAPI.md#GetSaldoHistory) | **Get** /v1/fs/{fs_id}/shop/{shop_id}/saldo-history | 



## GetSaldoHistory

> GetSaldoHistory200Response GetSaldoHistory(ctx, fsId, shopId).FromDate(fromDate).ToDate(toDate).Page(page).PerPage(perPage).Export(export).Execute()





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
    fromDate := "fromDate_example" // string | UNIX timestamp of date (hour, min, sec) from which the order details are requested. Max Range per request 31 Days
    toDate := "toDate_example" // string | UNIX timestamp of date (hour, min, sec) to which the order details are requested. Max Range per request 31 Days
    page := int64(789) // int64 | Determine which page the order list should start. Start from 1. You can get the next page when the previous page has have_next_page value true
    perPage := int64(789) // int64 | Determine how many orders will be shown per page. The maximum value is 500
    export := int64(789) // int64 | When the value is 1 then data will be send as files with .xls extension (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FinanceAPI.GetSaldoHistory(context.Background(), fsId, shopId).FromDate(fromDate).ToDate(toDate).Page(page).PerPage(perPage).Export(export).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FinanceAPI.GetSaldoHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSaldoHistory`: GetSaldoHistory200Response
    fmt.Fprintf(os.Stdout, "Response from `FinanceAPI.GetSaldoHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsId** | **int64** | Fulfillment service unique identifier | 
**shopId** | **int64** | Shop service unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSaldoHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fromDate** | **string** | UNIX timestamp of date (hour, min, sec) from which the order details are requested. Max Range per request 31 Days | 
 **toDate** | **string** | UNIX timestamp of date (hour, min, sec) to which the order details are requested. Max Range per request 31 Days | 
 **page** | **int64** | Determine which page the order list should start. Start from 1. You can get the next page when the previous page has have_next_page value true | 
 **perPage** | **int64** | Determine how many orders will be shown per page. The maximum value is 500 | 
 **export** | **int64** | When the value is 1 then data will be send as files with .xls extension | 

### Return type

[**GetSaldoHistory200Response**](GetSaldoHistory200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

