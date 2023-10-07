# UpdatePriceDefaultResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Header** | Pointer to [**ResponseHeader**](ResponseHeader.md) |  | [optional] 
**Data** | Pointer to [**UpdatePriceDefaultResponseData**](UpdatePriceDefaultResponseData.md) |  | [optional] 

## Methods

### NewUpdatePriceDefaultResponse

`func NewUpdatePriceDefaultResponse() *UpdatePriceDefaultResponse`

NewUpdatePriceDefaultResponse instantiates a new UpdatePriceDefaultResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePriceDefaultResponseWithDefaults

`func NewUpdatePriceDefaultResponseWithDefaults() *UpdatePriceDefaultResponse`

NewUpdatePriceDefaultResponseWithDefaults instantiates a new UpdatePriceDefaultResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeader

`func (o *UpdatePriceDefaultResponse) GetHeader() ResponseHeader`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *UpdatePriceDefaultResponse) GetHeaderOk() (*ResponseHeader, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *UpdatePriceDefaultResponse) SetHeader(v ResponseHeader)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *UpdatePriceDefaultResponse) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetData

`func (o *UpdatePriceDefaultResponse) GetData() UpdatePriceDefaultResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UpdatePriceDefaultResponse) GetDataOk() (*UpdatePriceDefaultResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UpdatePriceDefaultResponse) SetData(v UpdatePriceDefaultResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *UpdatePriceDefaultResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


