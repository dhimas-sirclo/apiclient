# WebhookStockAlertNotificationProductChangesDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | Action Type of Product Change Notification | [optional] [default to "stock_alert_changes"]
**ProductId** | Pointer to **int64** | Product unique identifier | [optional] 
**ShopId** | Pointer to **int64** | Shop unique identifier | [optional] 
**WarehouseId** | Pointer to **int64** | Warehouse unique identifier | [optional] 
**Value** | Pointer to **string** | New decrement stock value | [optional] 
**UpdateTime** | Pointer to **int64** | UNIX format updated time | [optional] 
**IsDefaultWarehouse** | Pointer to **bool** | Indicate for default warehouse of the shop. If not implement multilocation project, pick the true response to decrement stock | [optional] 

## Methods

### NewWebhookStockAlertNotificationProductChangesDataInner

`func NewWebhookStockAlertNotificationProductChangesDataInner() *WebhookStockAlertNotificationProductChangesDataInner`

NewWebhookStockAlertNotificationProductChangesDataInner instantiates a new WebhookStockAlertNotificationProductChangesDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookStockAlertNotificationProductChangesDataInnerWithDefaults

`func NewWebhookStockAlertNotificationProductChangesDataInnerWithDefaults() *WebhookStockAlertNotificationProductChangesDataInner`

NewWebhookStockAlertNotificationProductChangesDataInnerWithDefaults instantiates a new WebhookStockAlertNotificationProductChangesDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *WebhookStockAlertNotificationProductChangesDataInner) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *WebhookStockAlertNotificationProductChangesDataInner) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *WebhookStockAlertNotificationProductChangesDataInner) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *WebhookStockAlertNotificationProductChangesDataInner) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetProductId

`func (o *WebhookStockAlertNotificationProductChangesDataInner) GetProductId() int64`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *WebhookStockAlertNotificationProductChangesDataInner) GetProductIdOk() (*int64, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *WebhookStockAlertNotificationProductChangesDataInner) SetProductId(v int64)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *WebhookStockAlertNotificationProductChangesDataInner) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetShopId

`func (o *WebhookStockAlertNotificationProductChangesDataInner) GetShopId() int64`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *WebhookStockAlertNotificationProductChangesDataInner) GetShopIdOk() (*int64, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *WebhookStockAlertNotificationProductChangesDataInner) SetShopId(v int64)`

SetShopId sets ShopId field to given value.

### HasShopId

`func (o *WebhookStockAlertNotificationProductChangesDataInner) HasShopId() bool`

HasShopId returns a boolean if a field has been set.

### GetWarehouseId

`func (o *WebhookStockAlertNotificationProductChangesDataInner) GetWarehouseId() int64`

GetWarehouseId returns the WarehouseId field if non-nil, zero value otherwise.

### GetWarehouseIdOk

`func (o *WebhookStockAlertNotificationProductChangesDataInner) GetWarehouseIdOk() (*int64, bool)`

GetWarehouseIdOk returns a tuple with the WarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseId

`func (o *WebhookStockAlertNotificationProductChangesDataInner) SetWarehouseId(v int64)`

SetWarehouseId sets WarehouseId field to given value.

### HasWarehouseId

`func (o *WebhookStockAlertNotificationProductChangesDataInner) HasWarehouseId() bool`

HasWarehouseId returns a boolean if a field has been set.

### GetValue

`func (o *WebhookStockAlertNotificationProductChangesDataInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *WebhookStockAlertNotificationProductChangesDataInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *WebhookStockAlertNotificationProductChangesDataInner) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *WebhookStockAlertNotificationProductChangesDataInner) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetUpdateTime

`func (o *WebhookStockAlertNotificationProductChangesDataInner) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *WebhookStockAlertNotificationProductChangesDataInner) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *WebhookStockAlertNotificationProductChangesDataInner) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *WebhookStockAlertNotificationProductChangesDataInner) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetIsDefaultWarehouse

`func (o *WebhookStockAlertNotificationProductChangesDataInner) GetIsDefaultWarehouse() bool`

GetIsDefaultWarehouse returns the IsDefaultWarehouse field if non-nil, zero value otherwise.

### GetIsDefaultWarehouseOk

`func (o *WebhookStockAlertNotificationProductChangesDataInner) GetIsDefaultWarehouseOk() (*bool, bool)`

GetIsDefaultWarehouseOk returns a tuple with the IsDefaultWarehouse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefaultWarehouse

`func (o *WebhookStockAlertNotificationProductChangesDataInner) SetIsDefaultWarehouse(v bool)`

SetIsDefaultWarehouse sets IsDefaultWarehouse field to given value.

### HasIsDefaultWarehouse

`func (o *WebhookStockAlertNotificationProductChangesDataInner) HasIsDefaultWarehouse() bool`

HasIsDefaultWarehouse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


