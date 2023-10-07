# WebhookErrorProductEditData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**TotalData** | Pointer to **int64** |  | [optional] 
**UnprocessedRows** | Pointer to **int64** |  | [optional] 
**SuccessRows** | Pointer to **int64** |  | [optional] 
**FailedRows** | Pointer to **int64** |  | [optional] 
**FailedRowsData** | Pointer to [**[]EditProductV3DefaultResponseDataFailedRowsDataInner**](EditProductV3DefaultResponseDataFailedRowsDataInner.md) |  | [optional] 
**ProcessedRows** | Pointer to **int64** |  | [optional] 

## Methods

### NewWebhookErrorProductEditData

`func NewWebhookErrorProductEditData() *WebhookErrorProductEditData`

NewWebhookErrorProductEditData instantiates a new WebhookErrorProductEditData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookErrorProductEditDataWithDefaults

`func NewWebhookErrorProductEditDataWithDefaults() *WebhookErrorProductEditData`

NewWebhookErrorProductEditDataWithDefaults instantiates a new WebhookErrorProductEditData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *WebhookErrorProductEditData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WebhookErrorProductEditData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WebhookErrorProductEditData) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WebhookErrorProductEditData) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTotalData

`func (o *WebhookErrorProductEditData) GetTotalData() int64`

GetTotalData returns the TotalData field if non-nil, zero value otherwise.

### GetTotalDataOk

`func (o *WebhookErrorProductEditData) GetTotalDataOk() (*int64, bool)`

GetTotalDataOk returns a tuple with the TotalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalData

`func (o *WebhookErrorProductEditData) SetTotalData(v int64)`

SetTotalData sets TotalData field to given value.

### HasTotalData

`func (o *WebhookErrorProductEditData) HasTotalData() bool`

HasTotalData returns a boolean if a field has been set.

### GetUnprocessedRows

`func (o *WebhookErrorProductEditData) GetUnprocessedRows() int64`

GetUnprocessedRows returns the UnprocessedRows field if non-nil, zero value otherwise.

### GetUnprocessedRowsOk

`func (o *WebhookErrorProductEditData) GetUnprocessedRowsOk() (*int64, bool)`

GetUnprocessedRowsOk returns a tuple with the UnprocessedRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnprocessedRows

`func (o *WebhookErrorProductEditData) SetUnprocessedRows(v int64)`

SetUnprocessedRows sets UnprocessedRows field to given value.

### HasUnprocessedRows

`func (o *WebhookErrorProductEditData) HasUnprocessedRows() bool`

HasUnprocessedRows returns a boolean if a field has been set.

### GetSuccessRows

`func (o *WebhookErrorProductEditData) GetSuccessRows() int64`

GetSuccessRows returns the SuccessRows field if non-nil, zero value otherwise.

### GetSuccessRowsOk

`func (o *WebhookErrorProductEditData) GetSuccessRowsOk() (*int64, bool)`

GetSuccessRowsOk returns a tuple with the SuccessRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessRows

`func (o *WebhookErrorProductEditData) SetSuccessRows(v int64)`

SetSuccessRows sets SuccessRows field to given value.

### HasSuccessRows

`func (o *WebhookErrorProductEditData) HasSuccessRows() bool`

HasSuccessRows returns a boolean if a field has been set.

### GetFailedRows

`func (o *WebhookErrorProductEditData) GetFailedRows() int64`

GetFailedRows returns the FailedRows field if non-nil, zero value otherwise.

### GetFailedRowsOk

`func (o *WebhookErrorProductEditData) GetFailedRowsOk() (*int64, bool)`

GetFailedRowsOk returns a tuple with the FailedRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedRows

`func (o *WebhookErrorProductEditData) SetFailedRows(v int64)`

SetFailedRows sets FailedRows field to given value.

### HasFailedRows

`func (o *WebhookErrorProductEditData) HasFailedRows() bool`

HasFailedRows returns a boolean if a field has been set.

### GetFailedRowsData

`func (o *WebhookErrorProductEditData) GetFailedRowsData() []EditProductV3DefaultResponseDataFailedRowsDataInner`

GetFailedRowsData returns the FailedRowsData field if non-nil, zero value otherwise.

### GetFailedRowsDataOk

`func (o *WebhookErrorProductEditData) GetFailedRowsDataOk() (*[]EditProductV3DefaultResponseDataFailedRowsDataInner, bool)`

GetFailedRowsDataOk returns a tuple with the FailedRowsData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedRowsData

`func (o *WebhookErrorProductEditData) SetFailedRowsData(v []EditProductV3DefaultResponseDataFailedRowsDataInner)`

SetFailedRowsData sets FailedRowsData field to given value.

### HasFailedRowsData

`func (o *WebhookErrorProductEditData) HasFailedRowsData() bool`

HasFailedRowsData returns a boolean if a field has been set.

### GetProcessedRows

`func (o *WebhookErrorProductEditData) GetProcessedRows() int64`

GetProcessedRows returns the ProcessedRows field if non-nil, zero value otherwise.

### GetProcessedRowsOk

`func (o *WebhookErrorProductEditData) GetProcessedRowsOk() (*int64, bool)`

GetProcessedRowsOk returns a tuple with the ProcessedRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedRows

`func (o *WebhookErrorProductEditData) SetProcessedRows(v int64)`

SetProcessedRows sets ProcessedRows field to given value.

### HasProcessedRows

`func (o *WebhookErrorProductEditData) HasProcessedRows() bool`

HasProcessedRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


