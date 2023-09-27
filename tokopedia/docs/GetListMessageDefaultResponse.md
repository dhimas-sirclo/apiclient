# GetListMessageDefaultResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Header** | Pointer to [**GetOrderWebhook200ResponseHeader**](GetOrderWebhook200ResponseHeader.md) |  | [optional] 
**Data** | Pointer to [**[]GetListMessageDefaultResponseDataInner**](GetListMessageDefaultResponseDataInner.md) |  | [optional] 

## Methods

### NewGetListMessageDefaultResponse

`func NewGetListMessageDefaultResponse() *GetListMessageDefaultResponse`

NewGetListMessageDefaultResponse instantiates a new GetListMessageDefaultResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetListMessageDefaultResponseWithDefaults

`func NewGetListMessageDefaultResponseWithDefaults() *GetListMessageDefaultResponse`

NewGetListMessageDefaultResponseWithDefaults instantiates a new GetListMessageDefaultResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeader

`func (o *GetListMessageDefaultResponse) GetHeader() GetOrderWebhook200ResponseHeader`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *GetListMessageDefaultResponse) GetHeaderOk() (*GetOrderWebhook200ResponseHeader, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *GetListMessageDefaultResponse) SetHeader(v GetOrderWebhook200ResponseHeader)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *GetListMessageDefaultResponse) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetData

`func (o *GetListMessageDefaultResponse) GetData() []GetListMessageDefaultResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetListMessageDefaultResponse) GetDataOk() (*[]GetListMessageDefaultResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetListMessageDefaultResponse) SetData(v []GetListMessageDefaultResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *GetListMessageDefaultResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *GetListMessageDefaultResponse) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *GetListMessageDefaultResponse) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


