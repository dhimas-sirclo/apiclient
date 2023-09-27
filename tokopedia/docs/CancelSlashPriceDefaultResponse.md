# CancelSlashPriceDefaultResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Header** | Pointer to [**ResponseHeader**](ResponseHeader.md) |  | [optional] 
**Data** | Pointer to [**NullableCancelSlashPriceDefaultResponseData**](CancelSlashPriceDefaultResponseData.md) |  | [optional] 

## Methods

### NewCancelSlashPriceDefaultResponse

`func NewCancelSlashPriceDefaultResponse() *CancelSlashPriceDefaultResponse`

NewCancelSlashPriceDefaultResponse instantiates a new CancelSlashPriceDefaultResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelSlashPriceDefaultResponseWithDefaults

`func NewCancelSlashPriceDefaultResponseWithDefaults() *CancelSlashPriceDefaultResponse`

NewCancelSlashPriceDefaultResponseWithDefaults instantiates a new CancelSlashPriceDefaultResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeader

`func (o *CancelSlashPriceDefaultResponse) GetHeader() ResponseHeader`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *CancelSlashPriceDefaultResponse) GetHeaderOk() (*ResponseHeader, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *CancelSlashPriceDefaultResponse) SetHeader(v ResponseHeader)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *CancelSlashPriceDefaultResponse) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetData

`func (o *CancelSlashPriceDefaultResponse) GetData() CancelSlashPriceDefaultResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CancelSlashPriceDefaultResponse) GetDataOk() (*CancelSlashPriceDefaultResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CancelSlashPriceDefaultResponse) SetData(v CancelSlashPriceDefaultResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *CancelSlashPriceDefaultResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *CancelSlashPriceDefaultResponse) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *CancelSlashPriceDefaultResponse) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


