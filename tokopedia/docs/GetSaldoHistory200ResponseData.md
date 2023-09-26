# GetSaldoHistory200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HaveNextPage** | Pointer to **bool** |  | [optional] 
**SaldoHistory** | Pointer to [**[]GetSaldoHistory200ResponseDataSaldoHistoryInner**](GetSaldoHistory200ResponseDataSaldoHistoryInner.md) |  | [optional] 

## Methods

### NewGetSaldoHistory200ResponseData

`func NewGetSaldoHistory200ResponseData() *GetSaldoHistory200ResponseData`

NewGetSaldoHistory200ResponseData instantiates a new GetSaldoHistory200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSaldoHistory200ResponseDataWithDefaults

`func NewGetSaldoHistory200ResponseDataWithDefaults() *GetSaldoHistory200ResponseData`

NewGetSaldoHistory200ResponseDataWithDefaults instantiates a new GetSaldoHistory200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHaveNextPage

`func (o *GetSaldoHistory200ResponseData) GetHaveNextPage() bool`

GetHaveNextPage returns the HaveNextPage field if non-nil, zero value otherwise.

### GetHaveNextPageOk

`func (o *GetSaldoHistory200ResponseData) GetHaveNextPageOk() (*bool, bool)`

GetHaveNextPageOk returns a tuple with the HaveNextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaveNextPage

`func (o *GetSaldoHistory200ResponseData) SetHaveNextPage(v bool)`

SetHaveNextPage sets HaveNextPage field to given value.

### HasHaveNextPage

`func (o *GetSaldoHistory200ResponseData) HasHaveNextPage() bool`

HasHaveNextPage returns a boolean if a field has been set.

### GetSaldoHistory

`func (o *GetSaldoHistory200ResponseData) GetSaldoHistory() []GetSaldoHistory200ResponseDataSaldoHistoryInner`

GetSaldoHistory returns the SaldoHistory field if non-nil, zero value otherwise.

### GetSaldoHistoryOk

`func (o *GetSaldoHistory200ResponseData) GetSaldoHistoryOk() (*[]GetSaldoHistory200ResponseDataSaldoHistoryInner, bool)`

GetSaldoHistoryOk returns a tuple with the SaldoHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaldoHistory

`func (o *GetSaldoHistory200ResponseData) SetSaldoHistory(v []GetSaldoHistory200ResponseDataSaldoHistoryInner)`

SetSaldoHistory sets SaldoHistory field to given value.

### HasSaldoHistory

`func (o *GetSaldoHistory200ResponseData) HasSaldoHistory() bool`

HasSaldoHistory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


