# WebhookProductStatusChangeNotificationProductChangesDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | Action Type of Product Change Notification | [optional] [default to "status_changes"]
**ProductId** | Pointer to **int64** | Product unique identifier | [optional] 
**ShopId** | Pointer to **int64** | Shop unique identifier | [optional] 
**Value** | Pointer to **string** | New price value | [optional] 
**UpdateTime** | Pointer to **int64** | UNIX format updated time | [optional] 

## Methods

### NewWebhookProductStatusChangeNotificationProductChangesDataInner

`func NewWebhookProductStatusChangeNotificationProductChangesDataInner() *WebhookProductStatusChangeNotificationProductChangesDataInner`

NewWebhookProductStatusChangeNotificationProductChangesDataInner instantiates a new WebhookProductStatusChangeNotificationProductChangesDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookProductStatusChangeNotificationProductChangesDataInnerWithDefaults

`func NewWebhookProductStatusChangeNotificationProductChangesDataInnerWithDefaults() *WebhookProductStatusChangeNotificationProductChangesDataInner`

NewWebhookProductStatusChangeNotificationProductChangesDataInnerWithDefaults instantiates a new WebhookProductStatusChangeNotificationProductChangesDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *WebhookProductStatusChangeNotificationProductChangesDataInner) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *WebhookProductStatusChangeNotificationProductChangesDataInner) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *WebhookProductStatusChangeNotificationProductChangesDataInner) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *WebhookProductStatusChangeNotificationProductChangesDataInner) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetProductId

`func (o *WebhookProductStatusChangeNotificationProductChangesDataInner) GetProductId() int64`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *WebhookProductStatusChangeNotificationProductChangesDataInner) GetProductIdOk() (*int64, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *WebhookProductStatusChangeNotificationProductChangesDataInner) SetProductId(v int64)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *WebhookProductStatusChangeNotificationProductChangesDataInner) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetShopId

`func (o *WebhookProductStatusChangeNotificationProductChangesDataInner) GetShopId() int64`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *WebhookProductStatusChangeNotificationProductChangesDataInner) GetShopIdOk() (*int64, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *WebhookProductStatusChangeNotificationProductChangesDataInner) SetShopId(v int64)`

SetShopId sets ShopId field to given value.

### HasShopId

`func (o *WebhookProductStatusChangeNotificationProductChangesDataInner) HasShopId() bool`

HasShopId returns a boolean if a field has been set.

### GetValue

`func (o *WebhookProductStatusChangeNotificationProductChangesDataInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *WebhookProductStatusChangeNotificationProductChangesDataInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *WebhookProductStatusChangeNotificationProductChangesDataInner) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *WebhookProductStatusChangeNotificationProductChangesDataInner) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetUpdateTime

`func (o *WebhookProductStatusChangeNotificationProductChangesDataInner) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *WebhookProductStatusChangeNotificationProductChangesDataInner) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *WebhookProductStatusChangeNotificationProductChangesDataInner) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *WebhookProductStatusChangeNotificationProductChangesDataInner) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


