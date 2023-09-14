# WebhookStockAlertNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FsId** | Pointer to **int64** | Fulfillment service unique identifier | [optional] 
**ProductChangesData** | Pointer to [**[]WebhookStockAlertNotificationProductChangesDataInner**](WebhookStockAlertNotificationProductChangesDataInner.md) | Products changes data | [optional] 

## Methods

### NewWebhookStockAlertNotification

`func NewWebhookStockAlertNotification() *WebhookStockAlertNotification`

NewWebhookStockAlertNotification instantiates a new WebhookStockAlertNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookStockAlertNotificationWithDefaults

`func NewWebhookStockAlertNotificationWithDefaults() *WebhookStockAlertNotification`

NewWebhookStockAlertNotificationWithDefaults instantiates a new WebhookStockAlertNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFsId

`func (o *WebhookStockAlertNotification) GetFsId() int64`

GetFsId returns the FsId field if non-nil, zero value otherwise.

### GetFsIdOk

`func (o *WebhookStockAlertNotification) GetFsIdOk() (*int64, bool)`

GetFsIdOk returns a tuple with the FsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsId

`func (o *WebhookStockAlertNotification) SetFsId(v int64)`

SetFsId sets FsId field to given value.

### HasFsId

`func (o *WebhookStockAlertNotification) HasFsId() bool`

HasFsId returns a boolean if a field has been set.

### GetProductChangesData

`func (o *WebhookStockAlertNotification) GetProductChangesData() []WebhookStockAlertNotificationProductChangesDataInner`

GetProductChangesData returns the ProductChangesData field if non-nil, zero value otherwise.

### GetProductChangesDataOk

`func (o *WebhookStockAlertNotification) GetProductChangesDataOk() (*[]WebhookStockAlertNotificationProductChangesDataInner, bool)`

GetProductChangesDataOk returns a tuple with the ProductChangesData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductChangesData

`func (o *WebhookStockAlertNotification) SetProductChangesData(v []WebhookStockAlertNotificationProductChangesDataInner)`

SetProductChangesData sets ProductChangesData field to given value.

### HasProductChangesData

`func (o *WebhookStockAlertNotification) HasProductChangesData() bool`

HasProductChangesData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


