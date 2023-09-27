# GetBundleListDefaultResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Header** | Pointer to [**GetOrderWebhook200ResponseHeader**](GetOrderWebhook200ResponseHeader.md) |  | [optional] 
**Data** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewGetBundleListDefaultResponse

`func NewGetBundleListDefaultResponse() *GetBundleListDefaultResponse`

NewGetBundleListDefaultResponse instantiates a new GetBundleListDefaultResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBundleListDefaultResponseWithDefaults

`func NewGetBundleListDefaultResponseWithDefaults() *GetBundleListDefaultResponse`

NewGetBundleListDefaultResponseWithDefaults instantiates a new GetBundleListDefaultResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeader

`func (o *GetBundleListDefaultResponse) GetHeader() GetOrderWebhook200ResponseHeader`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *GetBundleListDefaultResponse) GetHeaderOk() (*GetOrderWebhook200ResponseHeader, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *GetBundleListDefaultResponse) SetHeader(v GetOrderWebhook200ResponseHeader)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *GetBundleListDefaultResponse) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetData

`func (o *GetBundleListDefaultResponse) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetBundleListDefaultResponse) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetBundleListDefaultResponse) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *GetBundleListDefaultResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *GetBundleListDefaultResponse) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *GetBundleListDefaultResponse) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


