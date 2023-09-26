# TriggerWebhookDefaultResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Header** | Pointer to [**GetOrderWebhook200ResponseHeader**](GetOrderWebhook200ResponseHeader.md) |  | [optional] 
**Data** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTriggerWebhookDefaultResponse

`func NewTriggerWebhookDefaultResponse() *TriggerWebhookDefaultResponse`

NewTriggerWebhookDefaultResponse instantiates a new TriggerWebhookDefaultResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerWebhookDefaultResponseWithDefaults

`func NewTriggerWebhookDefaultResponseWithDefaults() *TriggerWebhookDefaultResponse`

NewTriggerWebhookDefaultResponseWithDefaults instantiates a new TriggerWebhookDefaultResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeader

`func (o *TriggerWebhookDefaultResponse) GetHeader() GetOrderWebhook200ResponseHeader`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *TriggerWebhookDefaultResponse) GetHeaderOk() (*GetOrderWebhook200ResponseHeader, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *TriggerWebhookDefaultResponse) SetHeader(v GetOrderWebhook200ResponseHeader)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *TriggerWebhookDefaultResponse) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetData

`func (o *TriggerWebhookDefaultResponse) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TriggerWebhookDefaultResponse) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TriggerWebhookDefaultResponse) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *TriggerWebhookDefaultResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *TriggerWebhookDefaultResponse) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *TriggerWebhookDefaultResponse) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


