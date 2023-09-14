# WebhookStockIncrementNotificationProductChangesDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | Action Type of Product Change Notification | [optional] [default to "stock_increment"]
**ProductId** | Pointer to **int64** | Product unique identifier | [optional] 
**ShopId** | Pointer to **int64** | Shop unique identifier | [optional] 
**WarehouseId** | Pointer to **int64** | Warehouse unique identifier | [optional] 
**Value** | Pointer to **string** | New decrement stock value | [optional] 
**PreviousValue** | Pointer to **string** | Previous main stock value | [optional] 
**UpdateTime** | Pointer to **int64** | UNIX format updated time | [optional] 
**IsDefaultWarehouse** | Pointer to **bool** | Indicate for default warehouse of the shop. If not implement multilocation project, pick the true response to decrement stock | [optional] 

## Methods

### NewWebhookStockIncrementNotificationProductChangesDataInner

`func NewWebhookStockIncrementNotificationProductChangesDataInner() *WebhookStockIncrementNotificationProductChangesDataInner`

NewWebhookStockIncrementNotificationProductChangesDataInner instantiates a new WebhookStockIncrementNotificationProductChangesDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookStockIncrementNotificationProductChangesDataInnerWithDefaults

`func NewWebhookStockIncrementNotificationProductChangesDataInnerWithDefaults() *WebhookStockIncrementNotificationProductChangesDataInner`

NewWebhookStockIncrementNotificationProductChangesDataInnerWithDefaults instantiates a new WebhookStockIncrementNotificationProductChangesDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *WebhookStockIncrementNotificationProductChangesDataInner) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *WebhookStockIncrementNotificationProductChangesDataInner) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *WebhookStockIncrementNotificationProductChangesDataInner) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *WebhookStockIncrementNotificationProductChangesDataInner) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetProductId

`func (o *WebhookStockIncrementNotificationProductChangesDataInner) GetProductId() int64`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *WebhookStockIncrementNotificationProductChangesDataInner) GetProductIdOk() (*int64, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *WebhookStockIncrementNotificationProductChangesDataInner) SetProductId(v int64)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *WebhookStockIncrementNotificationProductChangesDataInner) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetShopId

`func (o *WebhookStockIncrementNotificationProductChangesDataInner) GetShopId() int64`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *WebhookStockIncrementNotificationProductChangesDataInner) GetShopIdOk() (*int64, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *WebhookStockIncrementNotificationProductChangesDataInner) SetShopId(v int64)`

SetShopId sets ShopId field to given value.

### HasShopId

`func (o *WebhookStockIncrementNotificationProductChangesDataInner) HasShopId() bool`

HasShopId returns a boolean if a field has been set.

### GetWarehouseId

`func (o *WebhookStockIncrementNotificationProductChangesDataInner) GetWarehouseId() int64`

GetWarehouseId returns the WarehouseId field if non-nil, zero value otherwise.

### GetWarehouseIdOk

`func (o *WebhookStockIncrementNotificationProductChangesDataInner) GetWarehouseIdOk() (*int64, bool)`

GetWarehouseIdOk returns a tuple with the WarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseId

`func (o *WebhookStockIncrementNotificationProductChangesDataInner) SetWarehouseId(v int64)`

SetWarehouseId sets WarehouseId field to given value.

### HasWarehouseId

`func (o *WebhookStockIncrementNotificationProductChangesDataInner) HasWarehouseId() bool`

HasWarehouseId returns a boolean if a field has been set.

### GetValue

`func (o *WebhookStockIncrementNotificationProductChangesDataInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *WebhookStockIncrementNotificationProductChangesDataInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *WebhookStockIncrementNotificationProductChangesDataInner) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *WebhookStockIncrementNotificationProductChangesDataInner) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetPreviousValue

`func (o *WebhookStockIncrementNotificationProductChangesDataInner) GetPreviousValue() string`

GetPreviousValue returns the PreviousValue field if non-nil, zero value otherwise.

### GetPreviousValueOk

`func (o *WebhookStockIncrementNotificationProductChangesDataInner) GetPreviousValueOk() (*string, bool)`

GetPreviousValueOk returns a tuple with the PreviousValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousValue

`func (o *WebhookStockIncrementNotificationProductChangesDataInner) SetPreviousValue(v string)`

SetPreviousValue sets PreviousValue field to given value.

### HasPreviousValue

`func (o *WebhookStockIncrementNotificationProductChangesDataInner) HasPreviousValue() bool`

HasPreviousValue returns a boolean if a field has been set.

### GetUpdateTime

`func (o *WebhookStockIncrementNotificationProductChangesDataInner) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *WebhookStockIncrementNotificationProductChangesDataInner) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *WebhookStockIncrementNotificationProductChangesDataInner) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *WebhookStockIncrementNotificationProductChangesDataInner) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetIsDefaultWarehouse

`func (o *WebhookStockIncrementNotificationProductChangesDataInner) GetIsDefaultWarehouse() bool`

GetIsDefaultWarehouse returns the IsDefaultWarehouse field if non-nil, zero value otherwise.

### GetIsDefaultWarehouseOk

`func (o *WebhookStockIncrementNotificationProductChangesDataInner) GetIsDefaultWarehouseOk() (*bool, bool)`

GetIsDefaultWarehouseOk returns a tuple with the IsDefaultWarehouse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefaultWarehouse

`func (o *WebhookStockIncrementNotificationProductChangesDataInner) SetIsDefaultWarehouse(v bool)`

SetIsDefaultWarehouse sets IsDefaultWarehouse field to given value.

### HasIsDefaultWarehouse

`func (o *WebhookStockIncrementNotificationProductChangesDataInner) HasIsDefaultWarehouse() bool`

HasIsDefaultWarehouse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


