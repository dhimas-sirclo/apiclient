# GetListReply200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Header** | Pointer to [**GetOrderWebhook200ResponseHeader**](GetOrderWebhook200ResponseHeader.md) |  | [optional] 
**Data** | Pointer to [**[]GetListReply200ResponseDataInner**](GetListReply200ResponseDataInner.md) |  | [optional] 

## Methods

### NewGetListReply200Response

`func NewGetListReply200Response() *GetListReply200Response`

NewGetListReply200Response instantiates a new GetListReply200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetListReply200ResponseWithDefaults

`func NewGetListReply200ResponseWithDefaults() *GetListReply200Response`

NewGetListReply200ResponseWithDefaults instantiates a new GetListReply200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeader

`func (o *GetListReply200Response) GetHeader() GetOrderWebhook200ResponseHeader`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *GetListReply200Response) GetHeaderOk() (*GetOrderWebhook200ResponseHeader, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *GetListReply200Response) SetHeader(v GetOrderWebhook200ResponseHeader)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *GetListReply200Response) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetData

`func (o *GetListReply200Response) GetData() []GetListReply200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetListReply200Response) GetDataOk() (*[]GetListReply200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetListReply200Response) SetData(v []GetListReply200ResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *GetListReply200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *GetListReply200Response) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *GetListReply200Response) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


