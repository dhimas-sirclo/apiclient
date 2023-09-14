# WebhookStockDecrementNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FsId** | Pointer to **int64** | Fulfillment service unique identifier | [optional] 
**ProductChangesData** | Pointer to [**[]WebhookStockDecrementNotificationProductChangesDataInner**](WebhookStockDecrementNotificationProductChangesDataInner.md) | Products changes data | [optional] 

## Methods

### NewWebhookStockDecrementNotification

`func NewWebhookStockDecrementNotification() *WebhookStockDecrementNotification`

NewWebhookStockDecrementNotification instantiates a new WebhookStockDecrementNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookStockDecrementNotificationWithDefaults

`func NewWebhookStockDecrementNotificationWithDefaults() *WebhookStockDecrementNotification`

NewWebhookStockDecrementNotificationWithDefaults instantiates a new WebhookStockDecrementNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFsId

`func (o *WebhookStockDecrementNotification) GetFsId() int64`

GetFsId returns the FsId field if non-nil, zero value otherwise.

### GetFsIdOk

`func (o *WebhookStockDecrementNotification) GetFsIdOk() (*int64, bool)`

GetFsIdOk returns a tuple with the FsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsId

`func (o *WebhookStockDecrementNotification) SetFsId(v int64)`

SetFsId sets FsId field to given value.

### HasFsId

`func (o *WebhookStockDecrementNotification) HasFsId() bool`

HasFsId returns a boolean if a field has been set.

### GetProductChangesData

`func (o *WebhookStockDecrementNotification) GetProductChangesData() []WebhookStockDecrementNotificationProductChangesDataInner`

GetProductChangesData returns the ProductChangesData field if non-nil, zero value otherwise.

### GetProductChangesDataOk

`func (o *WebhookStockDecrementNotification) GetProductChangesDataOk() (*[]WebhookStockDecrementNotificationProductChangesDataInner, bool)`

GetProductChangesDataOk returns a tuple with the ProductChangesData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductChangesData

`func (o *WebhookStockDecrementNotification) SetProductChangesData(v []WebhookStockDecrementNotificationProductChangesDataInner)`

SetProductChangesData sets ProductChangesData field to given value.

### HasProductChangesData

`func (o *WebhookStockDecrementNotification) HasProductChangesData() bool`

HasProductChangesData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


