# UpdateStockResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailedRows** | Pointer to **int64** |  | [optional] 
**FailedRowsData** | Pointer to [**[]UpdateStockResponseDataFailedRowsDataInner**](UpdateStockResponseDataFailedRowsDataInner.md) |  | [optional] 
**SucceedRows** | Pointer to **int64** |  | [optional] 

## Methods

### NewUpdateStockResponseData

`func NewUpdateStockResponseData() *UpdateStockResponseData`

NewUpdateStockResponseData instantiates a new UpdateStockResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateStockResponseDataWithDefaults

`func NewUpdateStockResponseDataWithDefaults() *UpdateStockResponseData`

NewUpdateStockResponseDataWithDefaults instantiates a new UpdateStockResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailedRows

`func (o *UpdateStockResponseData) GetFailedRows() int64`

GetFailedRows returns the FailedRows field if non-nil, zero value otherwise.

### GetFailedRowsOk

`func (o *UpdateStockResponseData) GetFailedRowsOk() (*int64, bool)`

GetFailedRowsOk returns a tuple with the FailedRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedRows

`func (o *UpdateStockResponseData) SetFailedRows(v int64)`

SetFailedRows sets FailedRows field to given value.

### HasFailedRows

`func (o *UpdateStockResponseData) HasFailedRows() bool`

HasFailedRows returns a boolean if a field has been set.

### GetFailedRowsData

`func (o *UpdateStockResponseData) GetFailedRowsData() []UpdateStockResponseDataFailedRowsDataInner`

GetFailedRowsData returns the FailedRowsData field if non-nil, zero value otherwise.

### GetFailedRowsDataOk

`func (o *UpdateStockResponseData) GetFailedRowsDataOk() (*[]UpdateStockResponseDataFailedRowsDataInner, bool)`

GetFailedRowsDataOk returns a tuple with the FailedRowsData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedRowsData

`func (o *UpdateStockResponseData) SetFailedRowsData(v []UpdateStockResponseDataFailedRowsDataInner)`

SetFailedRowsData sets FailedRowsData field to given value.

### HasFailedRowsData

`func (o *UpdateStockResponseData) HasFailedRowsData() bool`

HasFailedRowsData returns a boolean if a field has been set.

### GetSucceedRows

`func (o *UpdateStockResponseData) GetSucceedRows() int64`

GetSucceedRows returns the SucceedRows field if non-nil, zero value otherwise.

### GetSucceedRowsOk

`func (o *UpdateStockResponseData) GetSucceedRowsOk() (*int64, bool)`

GetSucceedRowsOk returns a tuple with the SucceedRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSucceedRows

`func (o *UpdateStockResponseData) SetSucceedRows(v int64)`

SetSucceedRows sets SucceedRows field to given value.

### HasSucceedRows

`func (o *UpdateStockResponseData) HasSucceedRows() bool`

HasSucceedRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


