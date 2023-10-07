# UpdatePriceDefaultResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SucceedRows** | Pointer to **int64** |  | [optional] 
**FailedRows** | Pointer to **int64** |  | [optional] 
**FailedRowsData** | Pointer to [**[]UpdatePriceDefaultResponseDataFailedRowsDataInner**](UpdatePriceDefaultResponseDataFailedRowsDataInner.md) |  | [optional] 

## Methods

### NewUpdatePriceDefaultResponseData

`func NewUpdatePriceDefaultResponseData() *UpdatePriceDefaultResponseData`

NewUpdatePriceDefaultResponseData instantiates a new UpdatePriceDefaultResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePriceDefaultResponseDataWithDefaults

`func NewUpdatePriceDefaultResponseDataWithDefaults() *UpdatePriceDefaultResponseData`

NewUpdatePriceDefaultResponseDataWithDefaults instantiates a new UpdatePriceDefaultResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSucceedRows

`func (o *UpdatePriceDefaultResponseData) GetSucceedRows() int64`

GetSucceedRows returns the SucceedRows field if non-nil, zero value otherwise.

### GetSucceedRowsOk

`func (o *UpdatePriceDefaultResponseData) GetSucceedRowsOk() (*int64, bool)`

GetSucceedRowsOk returns a tuple with the SucceedRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSucceedRows

`func (o *UpdatePriceDefaultResponseData) SetSucceedRows(v int64)`

SetSucceedRows sets SucceedRows field to given value.

### HasSucceedRows

`func (o *UpdatePriceDefaultResponseData) HasSucceedRows() bool`

HasSucceedRows returns a boolean if a field has been set.

### GetFailedRows

`func (o *UpdatePriceDefaultResponseData) GetFailedRows() int64`

GetFailedRows returns the FailedRows field if non-nil, zero value otherwise.

### GetFailedRowsOk

`func (o *UpdatePriceDefaultResponseData) GetFailedRowsOk() (*int64, bool)`

GetFailedRowsOk returns a tuple with the FailedRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedRows

`func (o *UpdatePriceDefaultResponseData) SetFailedRows(v int64)`

SetFailedRows sets FailedRows field to given value.

### HasFailedRows

`func (o *UpdatePriceDefaultResponseData) HasFailedRows() bool`

HasFailedRows returns a boolean if a field has been set.

### GetFailedRowsData

`func (o *UpdatePriceDefaultResponseData) GetFailedRowsData() []UpdatePriceDefaultResponseDataFailedRowsDataInner`

GetFailedRowsData returns the FailedRowsData field if non-nil, zero value otherwise.

### GetFailedRowsDataOk

`func (o *UpdatePriceDefaultResponseData) GetFailedRowsDataOk() (*[]UpdatePriceDefaultResponseDataFailedRowsDataInner, bool)`

GetFailedRowsDataOk returns a tuple with the FailedRowsData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedRowsData

`func (o *UpdatePriceDefaultResponseData) SetFailedRowsData(v []UpdatePriceDefaultResponseDataFailedRowsDataInner)`

SetFailedRowsData sets FailedRowsData field to given value.

### HasFailedRowsData

`func (o *UpdatePriceDefaultResponseData) HasFailedRowsData() bool`

HasFailedRowsData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


