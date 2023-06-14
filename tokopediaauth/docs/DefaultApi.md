# \DefaultApi

All URIs are relative to *https://accounts.tokopedia.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TokenPost**](DefaultApi.md#TokenPost) | **Post** /token | 



## TokenPost

> Token TokenPost(ctx).GrantType(grantType).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/dhimas-sirclo/apiclient/tokopediaauth"
)

func main() {
    grantType := "grantType_example" // string |  (default to "client_credentials")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.TokenPost(context.Background()).GrantType(grantType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.TokenPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TokenPost`: Token
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.TokenPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTokenPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **grantType** | **string** |  | [default to &quot;client_credentials&quot;]

### Return type

[**Token**](Token.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

