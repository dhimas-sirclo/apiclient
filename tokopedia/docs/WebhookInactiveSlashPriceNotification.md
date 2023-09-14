# WebhookInactiveSlashPriceNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | Action Type of Active Slash Price Notification | [optional] [default to "slash_price_inactive"]
**FsId** | Pointer to **int64** | Fulfillment service unique identifier | [optional] 
**ProductId** | Pointer to **int64** | Product unique identifier | [optional] 
**ShopId** | Pointer to **int64** | Shop unique identifier | [optional] 
**OriginalPrice** | Pointer to **int64** | Normal Price | [optional] 
**DiscountedPrice** | Pointer to **int64** | New Price When Slash Price Already Started | [optional] 
**DiscountPercentage** | Pointer to **int64** | Discount Percentage | [optional] 
**StartDate** | Pointer to **string** | Slash Price Start Time | [optional] 
**EndDate** | Pointer to **string** | Slash Price End Time | [optional] 

## Methods

### NewWebhookInactiveSlashPriceNotification

`func NewWebhookInactiveSlashPriceNotification() *WebhookInactiveSlashPriceNotification`

NewWebhookInactiveSlashPriceNotification instantiates a new WebhookInactiveSlashPriceNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookInactiveSlashPriceNotificationWithDefaults

`func NewWebhookInactiveSlashPriceNotificationWithDefaults() *WebhookInactiveSlashPriceNotification`

NewWebhookInactiveSlashPriceNotificationWithDefaults instantiates a new WebhookInactiveSlashPriceNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *WebhookInactiveSlashPriceNotification) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *WebhookInactiveSlashPriceNotification) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *WebhookInactiveSlashPriceNotification) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *WebhookInactiveSlashPriceNotification) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetFsId

`func (o *WebhookInactiveSlashPriceNotification) GetFsId() int64`

GetFsId returns the FsId field if non-nil, zero value otherwise.

### GetFsIdOk

`func (o *WebhookInactiveSlashPriceNotification) GetFsIdOk() (*int64, bool)`

GetFsIdOk returns a tuple with the FsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsId

`func (o *WebhookInactiveSlashPriceNotification) SetFsId(v int64)`

SetFsId sets FsId field to given value.

### HasFsId

`func (o *WebhookInactiveSlashPriceNotification) HasFsId() bool`

HasFsId returns a boolean if a field has been set.

### GetProductId

`func (o *WebhookInactiveSlashPriceNotification) GetProductId() int64`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *WebhookInactiveSlashPriceNotification) GetProductIdOk() (*int64, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *WebhookInactiveSlashPriceNotification) SetProductId(v int64)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *WebhookInactiveSlashPriceNotification) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetShopId

`func (o *WebhookInactiveSlashPriceNotification) GetShopId() int64`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *WebhookInactiveSlashPriceNotification) GetShopIdOk() (*int64, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *WebhookInactiveSlashPriceNotification) SetShopId(v int64)`

SetShopId sets ShopId field to given value.

### HasShopId

`func (o *WebhookInactiveSlashPriceNotification) HasShopId() bool`

HasShopId returns a boolean if a field has been set.

### GetOriginalPrice

`func (o *WebhookInactiveSlashPriceNotification) GetOriginalPrice() int64`

GetOriginalPrice returns the OriginalPrice field if non-nil, zero value otherwise.

### GetOriginalPriceOk

`func (o *WebhookInactiveSlashPriceNotification) GetOriginalPriceOk() (*int64, bool)`

GetOriginalPriceOk returns a tuple with the OriginalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalPrice

`func (o *WebhookInactiveSlashPriceNotification) SetOriginalPrice(v int64)`

SetOriginalPrice sets OriginalPrice field to given value.

### HasOriginalPrice

`func (o *WebhookInactiveSlashPriceNotification) HasOriginalPrice() bool`

HasOriginalPrice returns a boolean if a field has been set.

### GetDiscountedPrice

`func (o *WebhookInactiveSlashPriceNotification) GetDiscountedPrice() int64`

GetDiscountedPrice returns the DiscountedPrice field if non-nil, zero value otherwise.

### GetDiscountedPriceOk

`func (o *WebhookInactiveSlashPriceNotification) GetDiscountedPriceOk() (*int64, bool)`

GetDiscountedPriceOk returns a tuple with the DiscountedPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountedPrice

`func (o *WebhookInactiveSlashPriceNotification) SetDiscountedPrice(v int64)`

SetDiscountedPrice sets DiscountedPrice field to given value.

### HasDiscountedPrice

`func (o *WebhookInactiveSlashPriceNotification) HasDiscountedPrice() bool`

HasDiscountedPrice returns a boolean if a field has been set.

### GetDiscountPercentage

`func (o *WebhookInactiveSlashPriceNotification) GetDiscountPercentage() int64`

GetDiscountPercentage returns the DiscountPercentage field if non-nil, zero value otherwise.

### GetDiscountPercentageOk

`func (o *WebhookInactiveSlashPriceNotification) GetDiscountPercentageOk() (*int64, bool)`

GetDiscountPercentageOk returns a tuple with the DiscountPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountPercentage

`func (o *WebhookInactiveSlashPriceNotification) SetDiscountPercentage(v int64)`

SetDiscountPercentage sets DiscountPercentage field to given value.

### HasDiscountPercentage

`func (o *WebhookInactiveSlashPriceNotification) HasDiscountPercentage() bool`

HasDiscountPercentage returns a boolean if a field has been set.

### GetStartDate

`func (o *WebhookInactiveSlashPriceNotification) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *WebhookInactiveSlashPriceNotification) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *WebhookInactiveSlashPriceNotification) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *WebhookInactiveSlashPriceNotification) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *WebhookInactiveSlashPriceNotification) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *WebhookInactiveSlashPriceNotification) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *WebhookInactiveSlashPriceNotification) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *WebhookInactiveSlashPriceNotification) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


