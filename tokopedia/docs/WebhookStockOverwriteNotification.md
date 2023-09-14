# WebhookStockOverwriteNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FsId** | Pointer to **int64** |  | [optional] 
**ProductChangesData** | Pointer to [**[]WebhookStockOverwriteNotificationProductChangesDataInner**](WebhookStockOverwriteNotificationProductChangesDataInner.md) |  | [optional] 

## Methods

### NewWebhookStockOverwriteNotification

`func NewWebhookStockOverwriteNotification() *WebhookStockOverwriteNotification`

NewWebhookStockOverwriteNotification instantiates a new WebhookStockOverwriteNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookStockOverwriteNotificationWithDefaults

`func NewWebhookStockOverwriteNotificationWithDefaults() *WebhookStockOverwriteNotification`

NewWebhookStockOverwriteNotificationWithDefaults instantiates a new WebhookStockOverwriteNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFsId

`func (o *WebhookStockOverwriteNotification) GetFsId() int64`

GetFsId returns the FsId field if non-nil, zero value otherwise.

### GetFsIdOk

`func (o *WebhookStockOverwriteNotification) GetFsIdOk() (*int64, bool)`

GetFsIdOk returns a tuple with the FsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsId

`func (o *WebhookStockOverwriteNotification) SetFsId(v int64)`

SetFsId sets FsId field to given value.

### HasFsId

`func (o *WebhookStockOverwriteNotification) HasFsId() bool`

HasFsId returns a boolean if a field has been set.

### GetProductChangesData

`func (o *WebhookStockOverwriteNotification) GetProductChangesData() []WebhookStockOverwriteNotificationProductChangesDataInner`

GetProductChangesData returns the ProductChangesData field if non-nil, zero value otherwise.

### GetProductChangesDataOk

`func (o *WebhookStockOverwriteNotification) GetProductChangesDataOk() (*[]WebhookStockOverwriteNotificationProductChangesDataInner, bool)`

GetProductChangesDataOk returns a tuple with the ProductChangesData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductChangesData

`func (o *WebhookStockOverwriteNotification) SetProductChangesData(v []WebhookStockOverwriteNotificationProductChangesDataInner)`

SetProductChangesData sets ProductChangesData field to given value.

### HasProductChangesData

`func (o *WebhookStockOverwriteNotification) HasProductChangesData() bool`

HasProductChangesData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


