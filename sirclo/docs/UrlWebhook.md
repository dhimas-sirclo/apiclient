# UrlWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UrlConnect** | Pointer to **string** |  | [optional] 
**UrlDisconnect** | Pointer to **string** |  | [optional] 
**UrlAcceptedOrder** | Pointer to **string** |  | [optional] 
**UrlPackedOrder** | Pointer to **string** |  | [optional] 
**UrlShippedOrder** | Pointer to **string** |  | [optional] 
**UrlCompletedOrder** | Pointer to **string** |  | [optional] 
**UrlCancelledOrder** | Pointer to **string** |  | [optional] 
**UrlFetchOrder** | Pointer to **string** |  | [optional] 
**UrlUpdateStock** | Pointer to **string** |  | [optional] 
**UrlBulkUpdateStock** | Pointer to **string** |  | [optional] 
**UrlUpdatePrice** | Pointer to **string** |  | [optional] 
**UrlBulkUpdatePrice** | Pointer to **string** |  | [optional] 
**UrlUpsertProduct** | Pointer to **string** |  | [optional] 
**UrlFetchProduct** | Pointer to **string** |  | [optional] 
**UrlFetchAttributes** | Pointer to **string** |  | [optional] 
**UrlConnectOauth2** | Pointer to **string** |  | [optional] 
**UrlGenerateShopId** | Pointer to **string** |  | [optional] 
**UrlGetShopName** | Pointer to **string** |  | [optional] 
**UrlPassthrough** | Pointer to **string** |  | [optional] 
**UrlWebhookNewOrder** | Pointer to **string** |  | [optional] 
**UrlWebhookUpsertOrder** | Pointer to **string** |  | [optional] 
**UrlWebhookCancelOrder** | Pointer to **string** |  | [optional] 
**UrlUpdateStatusProduct** | Pointer to **string** |  | [optional] 
**UrlCreateDiscount** | Pointer to **string** |  | [optional] 
**UrlDeleteDiscount** | Pointer to **string** |  | [optional] 
**UrlPullWarehouse** | Pointer to **string** |  | [optional] 
**UrlFetchCategories** | Pointer to **string** |  | [optional] 

## Methods

### NewUrlWebhook

`func NewUrlWebhook() *UrlWebhook`

NewUrlWebhook instantiates a new UrlWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUrlWebhookWithDefaults

`func NewUrlWebhookWithDefaults() *UrlWebhook`

NewUrlWebhookWithDefaults instantiates a new UrlWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrlConnect

`func (o *UrlWebhook) GetUrlConnect() string`

GetUrlConnect returns the UrlConnect field if non-nil, zero value otherwise.

### GetUrlConnectOk

`func (o *UrlWebhook) GetUrlConnectOk() (*string, bool)`

GetUrlConnectOk returns a tuple with the UrlConnect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlConnect

`func (o *UrlWebhook) SetUrlConnect(v string)`

SetUrlConnect sets UrlConnect field to given value.

### HasUrlConnect

`func (o *UrlWebhook) HasUrlConnect() bool`

HasUrlConnect returns a boolean if a field has been set.

### GetUrlDisconnect

`func (o *UrlWebhook) GetUrlDisconnect() string`

GetUrlDisconnect returns the UrlDisconnect field if non-nil, zero value otherwise.

### GetUrlDisconnectOk

`func (o *UrlWebhook) GetUrlDisconnectOk() (*string, bool)`

GetUrlDisconnectOk returns a tuple with the UrlDisconnect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlDisconnect

`func (o *UrlWebhook) SetUrlDisconnect(v string)`

SetUrlDisconnect sets UrlDisconnect field to given value.

### HasUrlDisconnect

`func (o *UrlWebhook) HasUrlDisconnect() bool`

HasUrlDisconnect returns a boolean if a field has been set.

### GetUrlAcceptedOrder

`func (o *UrlWebhook) GetUrlAcceptedOrder() string`

GetUrlAcceptedOrder returns the UrlAcceptedOrder field if non-nil, zero value otherwise.

### GetUrlAcceptedOrderOk

`func (o *UrlWebhook) GetUrlAcceptedOrderOk() (*string, bool)`

GetUrlAcceptedOrderOk returns a tuple with the UrlAcceptedOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlAcceptedOrder

`func (o *UrlWebhook) SetUrlAcceptedOrder(v string)`

SetUrlAcceptedOrder sets UrlAcceptedOrder field to given value.

### HasUrlAcceptedOrder

`func (o *UrlWebhook) HasUrlAcceptedOrder() bool`

HasUrlAcceptedOrder returns a boolean if a field has been set.

### GetUrlPackedOrder

`func (o *UrlWebhook) GetUrlPackedOrder() string`

GetUrlPackedOrder returns the UrlPackedOrder field if non-nil, zero value otherwise.

### GetUrlPackedOrderOk

`func (o *UrlWebhook) GetUrlPackedOrderOk() (*string, bool)`

GetUrlPackedOrderOk returns a tuple with the UrlPackedOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlPackedOrder

`func (o *UrlWebhook) SetUrlPackedOrder(v string)`

SetUrlPackedOrder sets UrlPackedOrder field to given value.

### HasUrlPackedOrder

`func (o *UrlWebhook) HasUrlPackedOrder() bool`

HasUrlPackedOrder returns a boolean if a field has been set.

### GetUrlShippedOrder

`func (o *UrlWebhook) GetUrlShippedOrder() string`

GetUrlShippedOrder returns the UrlShippedOrder field if non-nil, zero value otherwise.

### GetUrlShippedOrderOk

`func (o *UrlWebhook) GetUrlShippedOrderOk() (*string, bool)`

GetUrlShippedOrderOk returns a tuple with the UrlShippedOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlShippedOrder

`func (o *UrlWebhook) SetUrlShippedOrder(v string)`

SetUrlShippedOrder sets UrlShippedOrder field to given value.

### HasUrlShippedOrder

`func (o *UrlWebhook) HasUrlShippedOrder() bool`

HasUrlShippedOrder returns a boolean if a field has been set.

### GetUrlCompletedOrder

`func (o *UrlWebhook) GetUrlCompletedOrder() string`

GetUrlCompletedOrder returns the UrlCompletedOrder field if non-nil, zero value otherwise.

### GetUrlCompletedOrderOk

`func (o *UrlWebhook) GetUrlCompletedOrderOk() (*string, bool)`

GetUrlCompletedOrderOk returns a tuple with the UrlCompletedOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlCompletedOrder

`func (o *UrlWebhook) SetUrlCompletedOrder(v string)`

SetUrlCompletedOrder sets UrlCompletedOrder field to given value.

### HasUrlCompletedOrder

`func (o *UrlWebhook) HasUrlCompletedOrder() bool`

HasUrlCompletedOrder returns a boolean if a field has been set.

### GetUrlCancelledOrder

`func (o *UrlWebhook) GetUrlCancelledOrder() string`

GetUrlCancelledOrder returns the UrlCancelledOrder field if non-nil, zero value otherwise.

### GetUrlCancelledOrderOk

`func (o *UrlWebhook) GetUrlCancelledOrderOk() (*string, bool)`

GetUrlCancelledOrderOk returns a tuple with the UrlCancelledOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlCancelledOrder

`func (o *UrlWebhook) SetUrlCancelledOrder(v string)`

SetUrlCancelledOrder sets UrlCancelledOrder field to given value.

### HasUrlCancelledOrder

`func (o *UrlWebhook) HasUrlCancelledOrder() bool`

HasUrlCancelledOrder returns a boolean if a field has been set.

### GetUrlFetchOrder

`func (o *UrlWebhook) GetUrlFetchOrder() string`

GetUrlFetchOrder returns the UrlFetchOrder field if non-nil, zero value otherwise.

### GetUrlFetchOrderOk

`func (o *UrlWebhook) GetUrlFetchOrderOk() (*string, bool)`

GetUrlFetchOrderOk returns a tuple with the UrlFetchOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlFetchOrder

`func (o *UrlWebhook) SetUrlFetchOrder(v string)`

SetUrlFetchOrder sets UrlFetchOrder field to given value.

### HasUrlFetchOrder

`func (o *UrlWebhook) HasUrlFetchOrder() bool`

HasUrlFetchOrder returns a boolean if a field has been set.

### GetUrlUpdateStock

`func (o *UrlWebhook) GetUrlUpdateStock() string`

GetUrlUpdateStock returns the UrlUpdateStock field if non-nil, zero value otherwise.

### GetUrlUpdateStockOk

`func (o *UrlWebhook) GetUrlUpdateStockOk() (*string, bool)`

GetUrlUpdateStockOk returns a tuple with the UrlUpdateStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlUpdateStock

`func (o *UrlWebhook) SetUrlUpdateStock(v string)`

SetUrlUpdateStock sets UrlUpdateStock field to given value.

### HasUrlUpdateStock

`func (o *UrlWebhook) HasUrlUpdateStock() bool`

HasUrlUpdateStock returns a boolean if a field has been set.

### GetUrlBulkUpdateStock

`func (o *UrlWebhook) GetUrlBulkUpdateStock() string`

GetUrlBulkUpdateStock returns the UrlBulkUpdateStock field if non-nil, zero value otherwise.

### GetUrlBulkUpdateStockOk

`func (o *UrlWebhook) GetUrlBulkUpdateStockOk() (*string, bool)`

GetUrlBulkUpdateStockOk returns a tuple with the UrlBulkUpdateStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlBulkUpdateStock

`func (o *UrlWebhook) SetUrlBulkUpdateStock(v string)`

SetUrlBulkUpdateStock sets UrlBulkUpdateStock field to given value.

### HasUrlBulkUpdateStock

`func (o *UrlWebhook) HasUrlBulkUpdateStock() bool`

HasUrlBulkUpdateStock returns a boolean if a field has been set.

### GetUrlUpdatePrice

`func (o *UrlWebhook) GetUrlUpdatePrice() string`

GetUrlUpdatePrice returns the UrlUpdatePrice field if non-nil, zero value otherwise.

### GetUrlUpdatePriceOk

`func (o *UrlWebhook) GetUrlUpdatePriceOk() (*string, bool)`

GetUrlUpdatePriceOk returns a tuple with the UrlUpdatePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlUpdatePrice

`func (o *UrlWebhook) SetUrlUpdatePrice(v string)`

SetUrlUpdatePrice sets UrlUpdatePrice field to given value.

### HasUrlUpdatePrice

`func (o *UrlWebhook) HasUrlUpdatePrice() bool`

HasUrlUpdatePrice returns a boolean if a field has been set.

### GetUrlBulkUpdatePrice

`func (o *UrlWebhook) GetUrlBulkUpdatePrice() string`

GetUrlBulkUpdatePrice returns the UrlBulkUpdatePrice field if non-nil, zero value otherwise.

### GetUrlBulkUpdatePriceOk

`func (o *UrlWebhook) GetUrlBulkUpdatePriceOk() (*string, bool)`

GetUrlBulkUpdatePriceOk returns a tuple with the UrlBulkUpdatePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlBulkUpdatePrice

`func (o *UrlWebhook) SetUrlBulkUpdatePrice(v string)`

SetUrlBulkUpdatePrice sets UrlBulkUpdatePrice field to given value.

### HasUrlBulkUpdatePrice

`func (o *UrlWebhook) HasUrlBulkUpdatePrice() bool`

HasUrlBulkUpdatePrice returns a boolean if a field has been set.

### GetUrlUpsertProduct

`func (o *UrlWebhook) GetUrlUpsertProduct() string`

GetUrlUpsertProduct returns the UrlUpsertProduct field if non-nil, zero value otherwise.

### GetUrlUpsertProductOk

`func (o *UrlWebhook) GetUrlUpsertProductOk() (*string, bool)`

GetUrlUpsertProductOk returns a tuple with the UrlUpsertProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlUpsertProduct

`func (o *UrlWebhook) SetUrlUpsertProduct(v string)`

SetUrlUpsertProduct sets UrlUpsertProduct field to given value.

### HasUrlUpsertProduct

`func (o *UrlWebhook) HasUrlUpsertProduct() bool`

HasUrlUpsertProduct returns a boolean if a field has been set.

### GetUrlFetchProduct

`func (o *UrlWebhook) GetUrlFetchProduct() string`

GetUrlFetchProduct returns the UrlFetchProduct field if non-nil, zero value otherwise.

### GetUrlFetchProductOk

`func (o *UrlWebhook) GetUrlFetchProductOk() (*string, bool)`

GetUrlFetchProductOk returns a tuple with the UrlFetchProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlFetchProduct

`func (o *UrlWebhook) SetUrlFetchProduct(v string)`

SetUrlFetchProduct sets UrlFetchProduct field to given value.

### HasUrlFetchProduct

`func (o *UrlWebhook) HasUrlFetchProduct() bool`

HasUrlFetchProduct returns a boolean if a field has been set.

### GetUrlFetchAttributes

`func (o *UrlWebhook) GetUrlFetchAttributes() string`

GetUrlFetchAttributes returns the UrlFetchAttributes field if non-nil, zero value otherwise.

### GetUrlFetchAttributesOk

`func (o *UrlWebhook) GetUrlFetchAttributesOk() (*string, bool)`

GetUrlFetchAttributesOk returns a tuple with the UrlFetchAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlFetchAttributes

`func (o *UrlWebhook) SetUrlFetchAttributes(v string)`

SetUrlFetchAttributes sets UrlFetchAttributes field to given value.

### HasUrlFetchAttributes

`func (o *UrlWebhook) HasUrlFetchAttributes() bool`

HasUrlFetchAttributes returns a boolean if a field has been set.

### GetUrlConnectOauth2

`func (o *UrlWebhook) GetUrlConnectOauth2() string`

GetUrlConnectOauth2 returns the UrlConnectOauth2 field if non-nil, zero value otherwise.

### GetUrlConnectOauth2Ok

`func (o *UrlWebhook) GetUrlConnectOauth2Ok() (*string, bool)`

GetUrlConnectOauth2Ok returns a tuple with the UrlConnectOauth2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlConnectOauth2

`func (o *UrlWebhook) SetUrlConnectOauth2(v string)`

SetUrlConnectOauth2 sets UrlConnectOauth2 field to given value.

### HasUrlConnectOauth2

`func (o *UrlWebhook) HasUrlConnectOauth2() bool`

HasUrlConnectOauth2 returns a boolean if a field has been set.

### GetUrlGenerateShopId

`func (o *UrlWebhook) GetUrlGenerateShopId() string`

GetUrlGenerateShopId returns the UrlGenerateShopId field if non-nil, zero value otherwise.

### GetUrlGenerateShopIdOk

`func (o *UrlWebhook) GetUrlGenerateShopIdOk() (*string, bool)`

GetUrlGenerateShopIdOk returns a tuple with the UrlGenerateShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlGenerateShopId

`func (o *UrlWebhook) SetUrlGenerateShopId(v string)`

SetUrlGenerateShopId sets UrlGenerateShopId field to given value.

### HasUrlGenerateShopId

`func (o *UrlWebhook) HasUrlGenerateShopId() bool`

HasUrlGenerateShopId returns a boolean if a field has been set.

### GetUrlGetShopName

`func (o *UrlWebhook) GetUrlGetShopName() string`

GetUrlGetShopName returns the UrlGetShopName field if non-nil, zero value otherwise.

### GetUrlGetShopNameOk

`func (o *UrlWebhook) GetUrlGetShopNameOk() (*string, bool)`

GetUrlGetShopNameOk returns a tuple with the UrlGetShopName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlGetShopName

`func (o *UrlWebhook) SetUrlGetShopName(v string)`

SetUrlGetShopName sets UrlGetShopName field to given value.

### HasUrlGetShopName

`func (o *UrlWebhook) HasUrlGetShopName() bool`

HasUrlGetShopName returns a boolean if a field has been set.

### GetUrlPassthrough

`func (o *UrlWebhook) GetUrlPassthrough() string`

GetUrlPassthrough returns the UrlPassthrough field if non-nil, zero value otherwise.

### GetUrlPassthroughOk

`func (o *UrlWebhook) GetUrlPassthroughOk() (*string, bool)`

GetUrlPassthroughOk returns a tuple with the UrlPassthrough field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlPassthrough

`func (o *UrlWebhook) SetUrlPassthrough(v string)`

SetUrlPassthrough sets UrlPassthrough field to given value.

### HasUrlPassthrough

`func (o *UrlWebhook) HasUrlPassthrough() bool`

HasUrlPassthrough returns a boolean if a field has been set.

### GetUrlWebhookNewOrder

`func (o *UrlWebhook) GetUrlWebhookNewOrder() string`

GetUrlWebhookNewOrder returns the UrlWebhookNewOrder field if non-nil, zero value otherwise.

### GetUrlWebhookNewOrderOk

`func (o *UrlWebhook) GetUrlWebhookNewOrderOk() (*string, bool)`

GetUrlWebhookNewOrderOk returns a tuple with the UrlWebhookNewOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlWebhookNewOrder

`func (o *UrlWebhook) SetUrlWebhookNewOrder(v string)`

SetUrlWebhookNewOrder sets UrlWebhookNewOrder field to given value.

### HasUrlWebhookNewOrder

`func (o *UrlWebhook) HasUrlWebhookNewOrder() bool`

HasUrlWebhookNewOrder returns a boolean if a field has been set.

### GetUrlWebhookUpsertOrder

`func (o *UrlWebhook) GetUrlWebhookUpsertOrder() string`

GetUrlWebhookUpsertOrder returns the UrlWebhookUpsertOrder field if non-nil, zero value otherwise.

### GetUrlWebhookUpsertOrderOk

`func (o *UrlWebhook) GetUrlWebhookUpsertOrderOk() (*string, bool)`

GetUrlWebhookUpsertOrderOk returns a tuple with the UrlWebhookUpsertOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlWebhookUpsertOrder

`func (o *UrlWebhook) SetUrlWebhookUpsertOrder(v string)`

SetUrlWebhookUpsertOrder sets UrlWebhookUpsertOrder field to given value.

### HasUrlWebhookUpsertOrder

`func (o *UrlWebhook) HasUrlWebhookUpsertOrder() bool`

HasUrlWebhookUpsertOrder returns a boolean if a field has been set.

### GetUrlWebhookCancelOrder

`func (o *UrlWebhook) GetUrlWebhookCancelOrder() string`

GetUrlWebhookCancelOrder returns the UrlWebhookCancelOrder field if non-nil, zero value otherwise.

### GetUrlWebhookCancelOrderOk

`func (o *UrlWebhook) GetUrlWebhookCancelOrderOk() (*string, bool)`

GetUrlWebhookCancelOrderOk returns a tuple with the UrlWebhookCancelOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlWebhookCancelOrder

`func (o *UrlWebhook) SetUrlWebhookCancelOrder(v string)`

SetUrlWebhookCancelOrder sets UrlWebhookCancelOrder field to given value.

### HasUrlWebhookCancelOrder

`func (o *UrlWebhook) HasUrlWebhookCancelOrder() bool`

HasUrlWebhookCancelOrder returns a boolean if a field has been set.

### GetUrlUpdateStatusProduct

`func (o *UrlWebhook) GetUrlUpdateStatusProduct() string`

GetUrlUpdateStatusProduct returns the UrlUpdateStatusProduct field if non-nil, zero value otherwise.

### GetUrlUpdateStatusProductOk

`func (o *UrlWebhook) GetUrlUpdateStatusProductOk() (*string, bool)`

GetUrlUpdateStatusProductOk returns a tuple with the UrlUpdateStatusProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlUpdateStatusProduct

`func (o *UrlWebhook) SetUrlUpdateStatusProduct(v string)`

SetUrlUpdateStatusProduct sets UrlUpdateStatusProduct field to given value.

### HasUrlUpdateStatusProduct

`func (o *UrlWebhook) HasUrlUpdateStatusProduct() bool`

HasUrlUpdateStatusProduct returns a boolean if a field has been set.

### GetUrlCreateDiscount

`func (o *UrlWebhook) GetUrlCreateDiscount() string`

GetUrlCreateDiscount returns the UrlCreateDiscount field if non-nil, zero value otherwise.

### GetUrlCreateDiscountOk

`func (o *UrlWebhook) GetUrlCreateDiscountOk() (*string, bool)`

GetUrlCreateDiscountOk returns a tuple with the UrlCreateDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlCreateDiscount

`func (o *UrlWebhook) SetUrlCreateDiscount(v string)`

SetUrlCreateDiscount sets UrlCreateDiscount field to given value.

### HasUrlCreateDiscount

`func (o *UrlWebhook) HasUrlCreateDiscount() bool`

HasUrlCreateDiscount returns a boolean if a field has been set.

### GetUrlDeleteDiscount

`func (o *UrlWebhook) GetUrlDeleteDiscount() string`

GetUrlDeleteDiscount returns the UrlDeleteDiscount field if non-nil, zero value otherwise.

### GetUrlDeleteDiscountOk

`func (o *UrlWebhook) GetUrlDeleteDiscountOk() (*string, bool)`

GetUrlDeleteDiscountOk returns a tuple with the UrlDeleteDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlDeleteDiscount

`func (o *UrlWebhook) SetUrlDeleteDiscount(v string)`

SetUrlDeleteDiscount sets UrlDeleteDiscount field to given value.

### HasUrlDeleteDiscount

`func (o *UrlWebhook) HasUrlDeleteDiscount() bool`

HasUrlDeleteDiscount returns a boolean if a field has been set.

### GetUrlPullWarehouse

`func (o *UrlWebhook) GetUrlPullWarehouse() string`

GetUrlPullWarehouse returns the UrlPullWarehouse field if non-nil, zero value otherwise.

### GetUrlPullWarehouseOk

`func (o *UrlWebhook) GetUrlPullWarehouseOk() (*string, bool)`

GetUrlPullWarehouseOk returns a tuple with the UrlPullWarehouse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlPullWarehouse

`func (o *UrlWebhook) SetUrlPullWarehouse(v string)`

SetUrlPullWarehouse sets UrlPullWarehouse field to given value.

### HasUrlPullWarehouse

`func (o *UrlWebhook) HasUrlPullWarehouse() bool`

HasUrlPullWarehouse returns a boolean if a field has been set.

### GetUrlFetchCategories

`func (o *UrlWebhook) GetUrlFetchCategories() string`

GetUrlFetchCategories returns the UrlFetchCategories field if non-nil, zero value otherwise.

### GetUrlFetchCategoriesOk

`func (o *UrlWebhook) GetUrlFetchCategoriesOk() (*string, bool)`

GetUrlFetchCategoriesOk returns a tuple with the UrlFetchCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlFetchCategories

`func (o *UrlWebhook) SetUrlFetchCategories(v string)`

SetUrlFetchCategories sets UrlFetchCategories field to given value.

### HasUrlFetchCategories

`func (o *UrlWebhook) HasUrlFetchCategories() bool`

HasUrlFetchCategories returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


