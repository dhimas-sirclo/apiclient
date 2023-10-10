# WebhookInactiveFlashSaleNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | Action Type of Active Flash Sale Notification | [optional] [default to "flash_sale_inactive"]
**FsId** | Pointer to **int64** | Fulfillment service unique identifier | [optional] 
**ProductId** | Pointer to **int64** | Product unique identifier | [optional] 
**ShopId** | Pointer to **int64** | Shop unique identifier | [optional] 
**OriginalPrice** | Pointer to **float64** | Normal Price | [optional] 
**DiscountedPrice** | Pointer to **float64** | New Price When Slash Price Already Started | [optional] 
**DiscountPercentage** | Pointer to **int64** | Discount Percentage | [optional] 
**StartDate** | Pointer to **string** | Flash Sale Start Time | [optional] 
**EndDate** | Pointer to **string** | Flash Sale End Time | [optional] 

## Methods

### NewWebhookInactiveFlashSaleNotification

`func NewWebhookInactiveFlashSaleNotification() *WebhookInactiveFlashSaleNotification`

NewWebhookInactiveFlashSaleNotification instantiates a new WebhookInactiveFlashSaleNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookInactiveFlashSaleNotificationWithDefaults

`func NewWebhookInactiveFlashSaleNotificationWithDefaults() *WebhookInactiveFlashSaleNotification`

NewWebhookInactiveFlashSaleNotificationWithDefaults instantiates a new WebhookInactiveFlashSaleNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *WebhookInactiveFlashSaleNotification) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *WebhookInactiveFlashSaleNotification) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *WebhookInactiveFlashSaleNotification) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *WebhookInactiveFlashSaleNotification) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetFsId

`func (o *WebhookInactiveFlashSaleNotification) GetFsId() int64`

GetFsId returns the FsId field if non-nil, zero value otherwise.

### GetFsIdOk

`func (o *WebhookInactiveFlashSaleNotification) GetFsIdOk() (*int64, bool)`

GetFsIdOk returns a tuple with the FsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsId

`func (o *WebhookInactiveFlashSaleNotification) SetFsId(v int64)`

SetFsId sets FsId field to given value.

### HasFsId

`func (o *WebhookInactiveFlashSaleNotification) HasFsId() bool`

HasFsId returns a boolean if a field has been set.

### GetProductId

`func (o *WebhookInactiveFlashSaleNotification) GetProductId() int64`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *WebhookInactiveFlashSaleNotification) GetProductIdOk() (*int64, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *WebhookInactiveFlashSaleNotification) SetProductId(v int64)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *WebhookInactiveFlashSaleNotification) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetShopId

`func (o *WebhookInactiveFlashSaleNotification) GetShopId() int64`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *WebhookInactiveFlashSaleNotification) GetShopIdOk() (*int64, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *WebhookInactiveFlashSaleNotification) SetShopId(v int64)`

SetShopId sets ShopId field to given value.

### HasShopId

`func (o *WebhookInactiveFlashSaleNotification) HasShopId() bool`

HasShopId returns a boolean if a field has been set.

### GetOriginalPrice

`func (o *WebhookInactiveFlashSaleNotification) GetOriginalPrice() float64`

GetOriginalPrice returns the OriginalPrice field if non-nil, zero value otherwise.

### GetOriginalPriceOk

`func (o *WebhookInactiveFlashSaleNotification) GetOriginalPriceOk() (*float64, bool)`

GetOriginalPriceOk returns a tuple with the OriginalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalPrice

`func (o *WebhookInactiveFlashSaleNotification) SetOriginalPrice(v float64)`

SetOriginalPrice sets OriginalPrice field to given value.

### HasOriginalPrice

`func (o *WebhookInactiveFlashSaleNotification) HasOriginalPrice() bool`

HasOriginalPrice returns a boolean if a field has been set.

### GetDiscountedPrice

`func (o *WebhookInactiveFlashSaleNotification) GetDiscountedPrice() float64`

GetDiscountedPrice returns the DiscountedPrice field if non-nil, zero value otherwise.

### GetDiscountedPriceOk

`func (o *WebhookInactiveFlashSaleNotification) GetDiscountedPriceOk() (*float64, bool)`

GetDiscountedPriceOk returns a tuple with the DiscountedPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountedPrice

`func (o *WebhookInactiveFlashSaleNotification) SetDiscountedPrice(v float64)`

SetDiscountedPrice sets DiscountedPrice field to given value.

### HasDiscountedPrice

`func (o *WebhookInactiveFlashSaleNotification) HasDiscountedPrice() bool`

HasDiscountedPrice returns a boolean if a field has been set.

### GetDiscountPercentage

`func (o *WebhookInactiveFlashSaleNotification) GetDiscountPercentage() int64`

GetDiscountPercentage returns the DiscountPercentage field if non-nil, zero value otherwise.

### GetDiscountPercentageOk

`func (o *WebhookInactiveFlashSaleNotification) GetDiscountPercentageOk() (*int64, bool)`

GetDiscountPercentageOk returns a tuple with the DiscountPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountPercentage

`func (o *WebhookInactiveFlashSaleNotification) SetDiscountPercentage(v int64)`

SetDiscountPercentage sets DiscountPercentage field to given value.

### HasDiscountPercentage

`func (o *WebhookInactiveFlashSaleNotification) HasDiscountPercentage() bool`

HasDiscountPercentage returns a boolean if a field has been set.

### GetStartDate

`func (o *WebhookInactiveFlashSaleNotification) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *WebhookInactiveFlashSaleNotification) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *WebhookInactiveFlashSaleNotification) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *WebhookInactiveFlashSaleNotification) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *WebhookInactiveFlashSaleNotification) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *WebhookInactiveFlashSaleNotification) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *WebhookInactiveFlashSaleNotification) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *WebhookInactiveFlashSaleNotification) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


