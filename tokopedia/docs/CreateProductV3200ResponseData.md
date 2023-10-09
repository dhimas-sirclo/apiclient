# CreateProductV3200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalData** | Pointer to **int64** |  | [optional] 
**SuccessData** | Pointer to **int64** |  | [optional] 
**FailData** | Pointer to **int64** |  | [optional] 
**SuccessRowData** | Pointer to [**[]CreateProductV3200ResponseDataSuccessRowDataInner**](CreateProductV3200ResponseDataSuccessRowDataInner.md) |  | [optional] 

## Methods

### NewCreateProductV3200ResponseData

`func NewCreateProductV3200ResponseData() *CreateProductV3200ResponseData`

NewCreateProductV3200ResponseData instantiates a new CreateProductV3200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProductV3200ResponseDataWithDefaults

`func NewCreateProductV3200ResponseDataWithDefaults() *CreateProductV3200ResponseData`

NewCreateProductV3200ResponseDataWithDefaults instantiates a new CreateProductV3200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalData

`func (o *CreateProductV3200ResponseData) GetTotalData() int64`

GetTotalData returns the TotalData field if non-nil, zero value otherwise.

### GetTotalDataOk

`func (o *CreateProductV3200ResponseData) GetTotalDataOk() (*int64, bool)`

GetTotalDataOk returns a tuple with the TotalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalData

`func (o *CreateProductV3200ResponseData) SetTotalData(v int64)`

SetTotalData sets TotalData field to given value.

### HasTotalData

`func (o *CreateProductV3200ResponseData) HasTotalData() bool`

HasTotalData returns a boolean if a field has been set.

### GetSuccessData

`func (o *CreateProductV3200ResponseData) GetSuccessData() int64`

GetSuccessData returns the SuccessData field if non-nil, zero value otherwise.

### GetSuccessDataOk

`func (o *CreateProductV3200ResponseData) GetSuccessDataOk() (*int64, bool)`

GetSuccessDataOk returns a tuple with the SuccessData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessData

`func (o *CreateProductV3200ResponseData) SetSuccessData(v int64)`

SetSuccessData sets SuccessData field to given value.

### HasSuccessData

`func (o *CreateProductV3200ResponseData) HasSuccessData() bool`

HasSuccessData returns a boolean if a field has been set.

### GetFailData

`func (o *CreateProductV3200ResponseData) GetFailData() int64`

GetFailData returns the FailData field if non-nil, zero value otherwise.

### GetFailDataOk

`func (o *CreateProductV3200ResponseData) GetFailDataOk() (*int64, bool)`

GetFailDataOk returns a tuple with the FailData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailData

`func (o *CreateProductV3200ResponseData) SetFailData(v int64)`

SetFailData sets FailData field to given value.

### HasFailData

`func (o *CreateProductV3200ResponseData) HasFailData() bool`

HasFailData returns a boolean if a field has been set.

### GetSuccessRowData

`func (o *CreateProductV3200ResponseData) GetSuccessRowData() []CreateProductV3200ResponseDataSuccessRowDataInner`

GetSuccessRowData returns the SuccessRowData field if non-nil, zero value otherwise.

### GetSuccessRowDataOk

`func (o *CreateProductV3200ResponseData) GetSuccessRowDataOk() (*[]CreateProductV3200ResponseDataSuccessRowDataInner, bool)`

GetSuccessRowDataOk returns a tuple with the SuccessRowData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessRowData

`func (o *CreateProductV3200ResponseData) SetSuccessRowData(v []CreateProductV3200ResponseDataSuccessRowDataInner)`

SetSuccessRowData sets SuccessRowData field to given value.

### HasSuccessRowData

`func (o *CreateProductV3200ResponseData) HasSuccessRowData() bool`

HasSuccessRowData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


