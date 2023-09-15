# UpdateSalesChannelInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**UrlInfo** | [**UrlInfoUpdateInput**](UrlInfoUpdateInput.md) |  | 
**UrlWebhook** | [**UrlWebhookUpdateInput**](UrlWebhookUpdateInput.md) |  | 

## Methods

### NewUpdateSalesChannelInput

`func NewUpdateSalesChannelInput(type_ string, urlInfo UrlInfoUpdateInput, urlWebhook UrlWebhookUpdateInput, ) *UpdateSalesChannelInput`

NewUpdateSalesChannelInput instantiates a new UpdateSalesChannelInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSalesChannelInputWithDefaults

`func NewUpdateSalesChannelInputWithDefaults() *UpdateSalesChannelInput`

NewUpdateSalesChannelInputWithDefaults instantiates a new UpdateSalesChannelInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *UpdateSalesChannelInput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateSalesChannelInput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateSalesChannelInput) SetType(v string)`

SetType sets Type field to given value.


### GetUrlInfo

`func (o *UpdateSalesChannelInput) GetUrlInfo() UrlInfoUpdateInput`

GetUrlInfo returns the UrlInfo field if non-nil, zero value otherwise.

### GetUrlInfoOk

`func (o *UpdateSalesChannelInput) GetUrlInfoOk() (*UrlInfoUpdateInput, bool)`

GetUrlInfoOk returns a tuple with the UrlInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlInfo

`func (o *UpdateSalesChannelInput) SetUrlInfo(v UrlInfoUpdateInput)`

SetUrlInfo sets UrlInfo field to given value.


### GetUrlWebhook

`func (o *UpdateSalesChannelInput) GetUrlWebhook() UrlWebhookUpdateInput`

GetUrlWebhook returns the UrlWebhook field if non-nil, zero value otherwise.

### GetUrlWebhookOk

`func (o *UpdateSalesChannelInput) GetUrlWebhookOk() (*UrlWebhookUpdateInput, bool)`

GetUrlWebhookOk returns a tuple with the UrlWebhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlWebhook

`func (o *UpdateSalesChannelInput) SetUrlWebhook(v UrlWebhookUpdateInput)`

SetUrlWebhook sets UrlWebhook field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


