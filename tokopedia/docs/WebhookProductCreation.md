# WebhookProductCreation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UploadId** | Pointer to **int64** | Upload id of the product to check Notification | [optional] 
**ShopId** | Pointer to **int64** | Shop unique identifier | [optional] 
**UserId** | Pointer to **int64** | User unique identifier | [optional] 
**FsId** | Pointer to **int64** | Fulfillment service unique identifier | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TotalData** | Pointer to **int64** |  | [optional] 
**SuccessRowData** | Pointer to [**[]WebhookProductCreationSuccessRowDataInner**](WebhookProductCreationSuccessRowDataInner.md) |  | [optional] 

## Methods

### NewWebhookProductCreation

`func NewWebhookProductCreation() *WebhookProductCreation`

NewWebhookProductCreation instantiates a new WebhookProductCreation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookProductCreationWithDefaults

`func NewWebhookProductCreationWithDefaults() *WebhookProductCreation`

NewWebhookProductCreationWithDefaults instantiates a new WebhookProductCreation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUploadId

`func (o *WebhookProductCreation) GetUploadId() int64`

GetUploadId returns the UploadId field if non-nil, zero value otherwise.

### GetUploadIdOk

`func (o *WebhookProductCreation) GetUploadIdOk() (*int64, bool)`

GetUploadIdOk returns a tuple with the UploadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadId

`func (o *WebhookProductCreation) SetUploadId(v int64)`

SetUploadId sets UploadId field to given value.

### HasUploadId

`func (o *WebhookProductCreation) HasUploadId() bool`

HasUploadId returns a boolean if a field has been set.

### GetShopId

`func (o *WebhookProductCreation) GetShopId() int64`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *WebhookProductCreation) GetShopIdOk() (*int64, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *WebhookProductCreation) SetShopId(v int64)`

SetShopId sets ShopId field to given value.

### HasShopId

`func (o *WebhookProductCreation) HasShopId() bool`

HasShopId returns a boolean if a field has been set.

### GetUserId

`func (o *WebhookProductCreation) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *WebhookProductCreation) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *WebhookProductCreation) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *WebhookProductCreation) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetFsId

`func (o *WebhookProductCreation) GetFsId() int64`

GetFsId returns the FsId field if non-nil, zero value otherwise.

### GetFsIdOk

`func (o *WebhookProductCreation) GetFsIdOk() (*int64, bool)`

GetFsIdOk returns a tuple with the FsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsId

`func (o *WebhookProductCreation) SetFsId(v int64)`

SetFsId sets FsId field to given value.

### HasFsId

`func (o *WebhookProductCreation) HasFsId() bool`

HasFsId returns a boolean if a field has been set.

### GetStatus

`func (o *WebhookProductCreation) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WebhookProductCreation) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WebhookProductCreation) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WebhookProductCreation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTotalData

`func (o *WebhookProductCreation) GetTotalData() int64`

GetTotalData returns the TotalData field if non-nil, zero value otherwise.

### GetTotalDataOk

`func (o *WebhookProductCreation) GetTotalDataOk() (*int64, bool)`

GetTotalDataOk returns a tuple with the TotalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalData

`func (o *WebhookProductCreation) SetTotalData(v int64)`

SetTotalData sets TotalData field to given value.

### HasTotalData

`func (o *WebhookProductCreation) HasTotalData() bool`

HasTotalData returns a boolean if a field has been set.

### GetSuccessRowData

`func (o *WebhookProductCreation) GetSuccessRowData() []WebhookProductCreationSuccessRowDataInner`

GetSuccessRowData returns the SuccessRowData field if non-nil, zero value otherwise.

### GetSuccessRowDataOk

`func (o *WebhookProductCreation) GetSuccessRowDataOk() (*[]WebhookProductCreationSuccessRowDataInner, bool)`

GetSuccessRowDataOk returns a tuple with the SuccessRowData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessRowData

`func (o *WebhookProductCreation) SetSuccessRowData(v []WebhookProductCreationSuccessRowDataInner)`

SetSuccessRowData sets SuccessRowData field to given value.

### HasSuccessRowData

`func (o *WebhookProductCreation) HasSuccessRowData() bool`

HasSuccessRowData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


