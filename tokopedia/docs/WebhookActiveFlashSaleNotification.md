# WebhookActiveFlashSaleNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | Action Type of Active Flash Sale Notification | [optional] [default to "flash_sale_active"]
**FsId** | Pointer to **int64** | Fulfillment service unique identifier | [optional] 
**ProductId** | Pointer to **int64** | Product unique identifier | [optional] 
**ShopId** | Pointer to **int64** | Shop unique identifier | [optional] 
**OriginalPrice** | Pointer to **int64** | Normal Price | [optional] 
**DiscountedPrice** | Pointer to **int64** | New Price When Slash Price Already Started | [optional] 
**DiscountPercentage** | Pointer to **int64** | Discount Percentage | [optional] 
**StartDate** | Pointer to **string** | Flash Sale Start Time | [optional] 
**EndDate** | Pointer to **string** | Flash Sale End Time | [optional] 

## Methods

### NewWebhookActiveFlashSaleNotification

`func NewWebhookActiveFlashSaleNotification() *WebhookActiveFlashSaleNotification`

NewWebhookActiveFlashSaleNotification instantiates a new WebhookActiveFlashSaleNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookActiveFlashSaleNotificationWithDefaults

`func NewWebhookActiveFlashSaleNotificationWithDefaults() *WebhookActiveFlashSaleNotification`

NewWebhookActiveFlashSaleNotificationWithDefaults instantiates a new WebhookActiveFlashSaleNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *WebhookActiveFlashSaleNotification) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *WebhookActiveFlashSaleNotification) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *WebhookActiveFlashSaleNotification) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *WebhookActiveFlashSaleNotification) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetFsId

`func (o *WebhookActiveFlashSaleNotification) GetFsId() int64`

GetFsId returns the FsId field if non-nil, zero value otherwise.

### GetFsIdOk

`func (o *WebhookActiveFlashSaleNotification) GetFsIdOk() (*int64, bool)`

GetFsIdOk returns a tuple with the FsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsId

`func (o *WebhookActiveFlashSaleNotification) SetFsId(v int64)`

SetFsId sets FsId field to given value.

### HasFsId

`func (o *WebhookActiveFlashSaleNotification) HasFsId() bool`

HasFsId returns a boolean if a field has been set.

### GetProductId

`func (o *WebhookActiveFlashSaleNotification) GetProductId() int64`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *WebhookActiveFlashSaleNotification) GetProductIdOk() (*int64, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *WebhookActiveFlashSaleNotification) SetProductId(v int64)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *WebhookActiveFlashSaleNotification) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetShopId

`func (o *WebhookActiveFlashSaleNotification) GetShopId() int64`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *WebhookActiveFlashSaleNotification) GetShopIdOk() (*int64, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *WebhookActiveFlashSaleNotification) SetShopId(v int64)`

SetShopId sets ShopId field to given value.

### HasShopId

`func (o *WebhookActiveFlashSaleNotification) HasShopId() bool`

HasShopId returns a boolean if a field has been set.

### GetOriginalPrice

`func (o *WebhookActiveFlashSaleNotification) GetOriginalPrice() int64`

GetOriginalPrice returns the OriginalPrice field if non-nil, zero value otherwise.

### GetOriginalPriceOk

`func (o *WebhookActiveFlashSaleNotification) GetOriginalPriceOk() (*int64, bool)`

GetOriginalPriceOk returns a tuple with the OriginalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalPrice

`func (o *WebhookActiveFlashSaleNotification) SetOriginalPrice(v int64)`

SetOriginalPrice sets OriginalPrice field to given value.

### HasOriginalPrice

`func (o *WebhookActiveFlashSaleNotification) HasOriginalPrice() bool`

HasOriginalPrice returns a boolean if a field has been set.

### GetDiscountedPrice

`func (o *WebhookActiveFlashSaleNotification) GetDiscountedPrice() int64`

GetDiscountedPrice returns the DiscountedPrice field if non-nil, zero value otherwise.

### GetDiscountedPriceOk

`func (o *WebhookActiveFlashSaleNotification) GetDiscountedPriceOk() (*int64, bool)`

GetDiscountedPriceOk returns a tuple with the DiscountedPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountedPrice

`func (o *WebhookActiveFlashSaleNotification) SetDiscountedPrice(v int64)`

SetDiscountedPrice sets DiscountedPrice field to given value.

### HasDiscountedPrice

`func (o *WebhookActiveFlashSaleNotification) HasDiscountedPrice() bool`

HasDiscountedPrice returns a boolean if a field has been set.

### GetDiscountPercentage

`func (o *WebhookActiveFlashSaleNotification) GetDiscountPercentage() int64`

GetDiscountPercentage returns the DiscountPercentage field if non-nil, zero value otherwise.

### GetDiscountPercentageOk

`func (o *WebhookActiveFlashSaleNotification) GetDiscountPercentageOk() (*int64, bool)`

GetDiscountPercentageOk returns a tuple with the DiscountPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountPercentage

`func (o *WebhookActiveFlashSaleNotification) SetDiscountPercentage(v int64)`

SetDiscountPercentage sets DiscountPercentage field to given value.

### HasDiscountPercentage

`func (o *WebhookActiveFlashSaleNotification) HasDiscountPercentage() bool`

HasDiscountPercentage returns a boolean if a field has been set.

### GetStartDate

`func (o *WebhookActiveFlashSaleNotification) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *WebhookActiveFlashSaleNotification) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *WebhookActiveFlashSaleNotification) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *WebhookActiveFlashSaleNotification) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *WebhookActiveFlashSaleNotification) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *WebhookActiveFlashSaleNotification) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *WebhookActiveFlashSaleNotification) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *WebhookActiveFlashSaleNotification) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


