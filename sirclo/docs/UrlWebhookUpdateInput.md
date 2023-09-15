# UrlWebhookUpdateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UrlConnect** | **string** |  | 
**UrlDisconnect** | **string** |  | 
**UrlAcceptedOrder** | **string** |  | 
**UrlPackedOrder** | **string** |  | 
**UrlShippedOrder** | Pointer to **string** |  | [optional] 
**UrlCompletedOrder** | Pointer to **string** |  | [optional] 
**UrlCancelledOrder** | **string** |  | 
**UrlUpdateStock** | **string** |  | 
**UrlBulkUpdateStock** | Pointer to **string** |  | [optional] 
**UrlUpsertProduct** | Pointer to **string** |  | [optional] 
**UrlFetchProduct** | Pointer to **string** |  | [optional] 
**UrlFetchOrder** | Pointer to **string** |  | [optional] 
**UrlConnectOauth2** | Pointer to **string** |  | [optional] 
**UrlUpdatePrice** | Pointer to **string** |  | [optional] 
**UrlBulkUpdatePrice** | Pointer to **string** |  | [optional] 
**UrlCreateDiscount** | Pointer to **string** |  | [optional] 
**UrlDeleteDiscount** | Pointer to **string** |  | [optional] 

## Methods

### NewUrlWebhookUpdateInput

`func NewUrlWebhookUpdateInput(urlConnect string, urlDisconnect string, urlAcceptedOrder string, urlPackedOrder string, urlCancelledOrder string, urlUpdateStock string, ) *UrlWebhookUpdateInput`

NewUrlWebhookUpdateInput instantiates a new UrlWebhookUpdateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUrlWebhookUpdateInputWithDefaults

`func NewUrlWebhookUpdateInputWithDefaults() *UrlWebhookUpdateInput`

NewUrlWebhookUpdateInputWithDefaults instantiates a new UrlWebhookUpdateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrlConnect

`func (o *UrlWebhookUpdateInput) GetUrlConnect() string`

GetUrlConnect returns the UrlConnect field if non-nil, zero value otherwise.

### GetUrlConnectOk

`func (o *UrlWebhookUpdateInput) GetUrlConnectOk() (*string, bool)`

GetUrlConnectOk returns a tuple with the UrlConnect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlConnect

`func (o *UrlWebhookUpdateInput) SetUrlConnect(v string)`

SetUrlConnect sets UrlConnect field to given value.


### GetUrlDisconnect

`func (o *UrlWebhookUpdateInput) GetUrlDisconnect() string`

GetUrlDisconnect returns the UrlDisconnect field if non-nil, zero value otherwise.

### GetUrlDisconnectOk

`func (o *UrlWebhookUpdateInput) GetUrlDisconnectOk() (*string, bool)`

GetUrlDisconnectOk returns a tuple with the UrlDisconnect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlDisconnect

`func (o *UrlWebhookUpdateInput) SetUrlDisconnect(v string)`

SetUrlDisconnect sets UrlDisconnect field to given value.


### GetUrlAcceptedOrder

`func (o *UrlWebhookUpdateInput) GetUrlAcceptedOrder() string`

GetUrlAcceptedOrder returns the UrlAcceptedOrder field if non-nil, zero value otherwise.

### GetUrlAcceptedOrderOk

`func (o *UrlWebhookUpdateInput) GetUrlAcceptedOrderOk() (*string, bool)`

GetUrlAcceptedOrderOk returns a tuple with the UrlAcceptedOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlAcceptedOrder

`func (o *UrlWebhookUpdateInput) SetUrlAcceptedOrder(v string)`

SetUrlAcceptedOrder sets UrlAcceptedOrder field to given value.


### GetUrlPackedOrder

`func (o *UrlWebhookUpdateInput) GetUrlPackedOrder() string`

GetUrlPackedOrder returns the UrlPackedOrder field if non-nil, zero value otherwise.

### GetUrlPackedOrderOk

`func (o *UrlWebhookUpdateInput) GetUrlPackedOrderOk() (*string, bool)`

GetUrlPackedOrderOk returns a tuple with the UrlPackedOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlPackedOrder

`func (o *UrlWebhookUpdateInput) SetUrlPackedOrder(v string)`

SetUrlPackedOrder sets UrlPackedOrder field to given value.


### GetUrlShippedOrder

`func (o *UrlWebhookUpdateInput) GetUrlShippedOrder() string`

GetUrlShippedOrder returns the UrlShippedOrder field if non-nil, zero value otherwise.

### GetUrlShippedOrderOk

`func (o *UrlWebhookUpdateInput) GetUrlShippedOrderOk() (*string, bool)`

GetUrlShippedOrderOk returns a tuple with the UrlShippedOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlShippedOrder

`func (o *UrlWebhookUpdateInput) SetUrlShippedOrder(v string)`

SetUrlShippedOrder sets UrlShippedOrder field to given value.

### HasUrlShippedOrder

`func (o *UrlWebhookUpdateInput) HasUrlShippedOrder() bool`

HasUrlShippedOrder returns a boolean if a field has been set.

### GetUrlCompletedOrder

`func (o *UrlWebhookUpdateInput) GetUrlCompletedOrder() string`

GetUrlCompletedOrder returns the UrlCompletedOrder field if non-nil, zero value otherwise.

### GetUrlCompletedOrderOk

`func (o *UrlWebhookUpdateInput) GetUrlCompletedOrderOk() (*string, bool)`

GetUrlCompletedOrderOk returns a tuple with the UrlCompletedOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlCompletedOrder

`func (o *UrlWebhookUpdateInput) SetUrlCompletedOrder(v string)`

SetUrlCompletedOrder sets UrlCompletedOrder field to given value.

### HasUrlCompletedOrder

`func (o *UrlWebhookUpdateInput) HasUrlCompletedOrder() bool`

HasUrlCompletedOrder returns a boolean if a field has been set.

### GetUrlCancelledOrder

`func (o *UrlWebhookUpdateInput) GetUrlCancelledOrder() string`

GetUrlCancelledOrder returns the UrlCancelledOrder field if non-nil, zero value otherwise.

### GetUrlCancelledOrderOk

`func (o *UrlWebhookUpdateInput) GetUrlCancelledOrderOk() (*string, bool)`

GetUrlCancelledOrderOk returns a tuple with the UrlCancelledOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlCancelledOrder

`func (o *UrlWebhookUpdateInput) SetUrlCancelledOrder(v string)`

SetUrlCancelledOrder sets UrlCancelledOrder field to given value.


### GetUrlUpdateStock

`func (o *UrlWebhookUpdateInput) GetUrlUpdateStock() string`

GetUrlUpdateStock returns the UrlUpdateStock field if non-nil, zero value otherwise.

### GetUrlUpdateStockOk

`func (o *UrlWebhookUpdateInput) GetUrlUpdateStockOk() (*string, bool)`

GetUrlUpdateStockOk returns a tuple with the UrlUpdateStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlUpdateStock

`func (o *UrlWebhookUpdateInput) SetUrlUpdateStock(v string)`

SetUrlUpdateStock sets UrlUpdateStock field to given value.


### GetUrlBulkUpdateStock

`func (o *UrlWebhookUpdateInput) GetUrlBulkUpdateStock() string`

GetUrlBulkUpdateStock returns the UrlBulkUpdateStock field if non-nil, zero value otherwise.

### GetUrlBulkUpdateStockOk

`func (o *UrlWebhookUpdateInput) GetUrlBulkUpdateStockOk() (*string, bool)`

GetUrlBulkUpdateStockOk returns a tuple with the UrlBulkUpdateStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlBulkUpdateStock

`func (o *UrlWebhookUpdateInput) SetUrlBulkUpdateStock(v string)`

SetUrlBulkUpdateStock sets UrlBulkUpdateStock field to given value.

### HasUrlBulkUpdateStock

`func (o *UrlWebhookUpdateInput) HasUrlBulkUpdateStock() bool`

HasUrlBulkUpdateStock returns a boolean if a field has been set.

### GetUrlUpsertProduct

`func (o *UrlWebhookUpdateInput) GetUrlUpsertProduct() string`

GetUrlUpsertProduct returns the UrlUpsertProduct field if non-nil, zero value otherwise.

### GetUrlUpsertProductOk

`func (o *UrlWebhookUpdateInput) GetUrlUpsertProductOk() (*string, bool)`

GetUrlUpsertProductOk returns a tuple with the UrlUpsertProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlUpsertProduct

`func (o *UrlWebhookUpdateInput) SetUrlUpsertProduct(v string)`

SetUrlUpsertProduct sets UrlUpsertProduct field to given value.

### HasUrlUpsertProduct

`func (o *UrlWebhookUpdateInput) HasUrlUpsertProduct() bool`

HasUrlUpsertProduct returns a boolean if a field has been set.

### GetUrlFetchProduct

`func (o *UrlWebhookUpdateInput) GetUrlFetchProduct() string`

GetUrlFetchProduct returns the UrlFetchProduct field if non-nil, zero value otherwise.

### GetUrlFetchProductOk

`func (o *UrlWebhookUpdateInput) GetUrlFetchProductOk() (*string, bool)`

GetUrlFetchProductOk returns a tuple with the UrlFetchProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlFetchProduct

`func (o *UrlWebhookUpdateInput) SetUrlFetchProduct(v string)`

SetUrlFetchProduct sets UrlFetchProduct field to given value.

### HasUrlFetchProduct

`func (o *UrlWebhookUpdateInput) HasUrlFetchProduct() bool`

HasUrlFetchProduct returns a boolean if a field has been set.

### GetUrlFetchOrder

`func (o *UrlWebhookUpdateInput) GetUrlFetchOrder() string`

GetUrlFetchOrder returns the UrlFetchOrder field if non-nil, zero value otherwise.

### GetUrlFetchOrderOk

`func (o *UrlWebhookUpdateInput) GetUrlFetchOrderOk() (*string, bool)`

GetUrlFetchOrderOk returns a tuple with the UrlFetchOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlFetchOrder

`func (o *UrlWebhookUpdateInput) SetUrlFetchOrder(v string)`

SetUrlFetchOrder sets UrlFetchOrder field to given value.

### HasUrlFetchOrder

`func (o *UrlWebhookUpdateInput) HasUrlFetchOrder() bool`

HasUrlFetchOrder returns a boolean if a field has been set.

### GetUrlConnectOauth2

`func (o *UrlWebhookUpdateInput) GetUrlConnectOauth2() string`

GetUrlConnectOauth2 returns the UrlConnectOauth2 field if non-nil, zero value otherwise.

### GetUrlConnectOauth2Ok

`func (o *UrlWebhookUpdateInput) GetUrlConnectOauth2Ok() (*string, bool)`

GetUrlConnectOauth2Ok returns a tuple with the UrlConnectOauth2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlConnectOauth2

`func (o *UrlWebhookUpdateInput) SetUrlConnectOauth2(v string)`

SetUrlConnectOauth2 sets UrlConnectOauth2 field to given value.

### HasUrlConnectOauth2

`func (o *UrlWebhookUpdateInput) HasUrlConnectOauth2() bool`

HasUrlConnectOauth2 returns a boolean if a field has been set.

### GetUrlUpdatePrice

`func (o *UrlWebhookUpdateInput) GetUrlUpdatePrice() string`

GetUrlUpdatePrice returns the UrlUpdatePrice field if non-nil, zero value otherwise.

### GetUrlUpdatePriceOk

`func (o *UrlWebhookUpdateInput) GetUrlUpdatePriceOk() (*string, bool)`

GetUrlUpdatePriceOk returns a tuple with the UrlUpdatePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlUpdatePrice

`func (o *UrlWebhookUpdateInput) SetUrlUpdatePrice(v string)`

SetUrlUpdatePrice sets UrlUpdatePrice field to given value.

### HasUrlUpdatePrice

`func (o *UrlWebhookUpdateInput) HasUrlUpdatePrice() bool`

HasUrlUpdatePrice returns a boolean if a field has been set.

### GetUrlBulkUpdatePrice

`func (o *UrlWebhookUpdateInput) GetUrlBulkUpdatePrice() string`

GetUrlBulkUpdatePrice returns the UrlBulkUpdatePrice field if non-nil, zero value otherwise.

### GetUrlBulkUpdatePriceOk

`func (o *UrlWebhookUpdateInput) GetUrlBulkUpdatePriceOk() (*string, bool)`

GetUrlBulkUpdatePriceOk returns a tuple with the UrlBulkUpdatePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlBulkUpdatePrice

`func (o *UrlWebhookUpdateInput) SetUrlBulkUpdatePrice(v string)`

SetUrlBulkUpdatePrice sets UrlBulkUpdatePrice field to given value.

### HasUrlBulkUpdatePrice

`func (o *UrlWebhookUpdateInput) HasUrlBulkUpdatePrice() bool`

HasUrlBulkUpdatePrice returns a boolean if a field has been set.

### GetUrlCreateDiscount

`func (o *UrlWebhookUpdateInput) GetUrlCreateDiscount() string`

GetUrlCreateDiscount returns the UrlCreateDiscount field if non-nil, zero value otherwise.

### GetUrlCreateDiscountOk

`func (o *UrlWebhookUpdateInput) GetUrlCreateDiscountOk() (*string, bool)`

GetUrlCreateDiscountOk returns a tuple with the UrlCreateDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlCreateDiscount

`func (o *UrlWebhookUpdateInput) SetUrlCreateDiscount(v string)`

SetUrlCreateDiscount sets UrlCreateDiscount field to given value.

### HasUrlCreateDiscount

`func (o *UrlWebhookUpdateInput) HasUrlCreateDiscount() bool`

HasUrlCreateDiscount returns a boolean if a field has been set.

### GetUrlDeleteDiscount

`func (o *UrlWebhookUpdateInput) GetUrlDeleteDiscount() string`

GetUrlDeleteDiscount returns the UrlDeleteDiscount field if non-nil, zero value otherwise.

### GetUrlDeleteDiscountOk

`func (o *UrlWebhookUpdateInput) GetUrlDeleteDiscountOk() (*string, bool)`

GetUrlDeleteDiscountOk returns a tuple with the UrlDeleteDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlDeleteDiscount

`func (o *UrlWebhookUpdateInput) SetUrlDeleteDiscount(v string)`

SetUrlDeleteDiscount sets UrlDeleteDiscount field to given value.

### HasUrlDeleteDiscount

`func (o *UrlWebhookUpdateInput) HasUrlDeleteDiscount() bool`

HasUrlDeleteDiscount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


