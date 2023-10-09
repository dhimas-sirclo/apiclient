# CreateProductV3DefaultResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalData** | Pointer to **int64** |  | [optional] 
**SuccessData** | Pointer to **int64** |  | [optional] 
**FailData** | Pointer to **int64** |  | [optional] 
**FailedRowsData** | Pointer to [**[]CreateProductV3DefaultResponseDataFailedRowsDataInner**](CreateProductV3DefaultResponseDataFailedRowsDataInner.md) |  | [optional] 

## Methods

### NewCreateProductV3DefaultResponseData

`func NewCreateProductV3DefaultResponseData() *CreateProductV3DefaultResponseData`

NewCreateProductV3DefaultResponseData instantiates a new CreateProductV3DefaultResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProductV3DefaultResponseDataWithDefaults

`func NewCreateProductV3DefaultResponseDataWithDefaults() *CreateProductV3DefaultResponseData`

NewCreateProductV3DefaultResponseDataWithDefaults instantiates a new CreateProductV3DefaultResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalData

`func (o *CreateProductV3DefaultResponseData) GetTotalData() int64`

GetTotalData returns the TotalData field if non-nil, zero value otherwise.

### GetTotalDataOk

`func (o *CreateProductV3DefaultResponseData) GetTotalDataOk() (*int64, bool)`

GetTotalDataOk returns a tuple with the TotalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalData

`func (o *CreateProductV3DefaultResponseData) SetTotalData(v int64)`

SetTotalData sets TotalData field to given value.

### HasTotalData

`func (o *CreateProductV3DefaultResponseData) HasTotalData() bool`

HasTotalData returns a boolean if a field has been set.

### GetSuccessData

`func (o *CreateProductV3DefaultResponseData) GetSuccessData() int64`

GetSuccessData returns the SuccessData field if non-nil, zero value otherwise.

### GetSuccessDataOk

`func (o *CreateProductV3DefaultResponseData) GetSuccessDataOk() (*int64, bool)`

GetSuccessDataOk returns a tuple with the SuccessData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessData

`func (o *CreateProductV3DefaultResponseData) SetSuccessData(v int64)`

SetSuccessData sets SuccessData field to given value.

### HasSuccessData

`func (o *CreateProductV3DefaultResponseData) HasSuccessData() bool`

HasSuccessData returns a boolean if a field has been set.

### GetFailData

`func (o *CreateProductV3DefaultResponseData) GetFailData() int64`

GetFailData returns the FailData field if non-nil, zero value otherwise.

### GetFailDataOk

`func (o *CreateProductV3DefaultResponseData) GetFailDataOk() (*int64, bool)`

GetFailDataOk returns a tuple with the FailData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailData

`func (o *CreateProductV3DefaultResponseData) SetFailData(v int64)`

SetFailData sets FailData field to given value.

### HasFailData

`func (o *CreateProductV3DefaultResponseData) HasFailData() bool`

HasFailData returns a boolean if a field has been set.

### GetFailedRowsData

`func (o *CreateProductV3DefaultResponseData) GetFailedRowsData() []CreateProductV3DefaultResponseDataFailedRowsDataInner`

GetFailedRowsData returns the FailedRowsData field if non-nil, zero value otherwise.

### GetFailedRowsDataOk

`func (o *CreateProductV3DefaultResponseData) GetFailedRowsDataOk() (*[]CreateProductV3DefaultResponseDataFailedRowsDataInner, bool)`

GetFailedRowsDataOk returns a tuple with the FailedRowsData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedRowsData

`func (o *CreateProductV3DefaultResponseData) SetFailedRowsData(v []CreateProductV3DefaultResponseDataFailedRowsDataInner)`

SetFailedRowsData sets FailedRowsData field to given value.

### HasFailedRowsData

`func (o *CreateProductV3DefaultResponseData) HasFailedRowsData() bool`

HasFailedRowsData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


