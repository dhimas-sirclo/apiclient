# RegisterSalesChannelInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppName** | **string** |  | 
**AppChannelCode** | **string** |  | 
**Type** | **string** |  | 
**UrlInfo** | [**UrlInfoInput**](UrlInfoInput.md) |  | 
**UrlWebhook** | Pointer to [**UrlWebhookInput**](UrlWebhookInput.md) |  | [optional] 

## Methods

### NewRegisterSalesChannelInput

`func NewRegisterSalesChannelInput(appName string, appChannelCode string, type_ string, urlInfo UrlInfoInput, ) *RegisterSalesChannelInput`

NewRegisterSalesChannelInput instantiates a new RegisterSalesChannelInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterSalesChannelInputWithDefaults

`func NewRegisterSalesChannelInputWithDefaults() *RegisterSalesChannelInput`

NewRegisterSalesChannelInputWithDefaults instantiates a new RegisterSalesChannelInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppName

`func (o *RegisterSalesChannelInput) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *RegisterSalesChannelInput) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *RegisterSalesChannelInput) SetAppName(v string)`

SetAppName sets AppName field to given value.


### GetAppChannelCode

`func (o *RegisterSalesChannelInput) GetAppChannelCode() string`

GetAppChannelCode returns the AppChannelCode field if non-nil, zero value otherwise.

### GetAppChannelCodeOk

`func (o *RegisterSalesChannelInput) GetAppChannelCodeOk() (*string, bool)`

GetAppChannelCodeOk returns a tuple with the AppChannelCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppChannelCode

`func (o *RegisterSalesChannelInput) SetAppChannelCode(v string)`

SetAppChannelCode sets AppChannelCode field to given value.


### GetType

`func (o *RegisterSalesChannelInput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RegisterSalesChannelInput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RegisterSalesChannelInput) SetType(v string)`

SetType sets Type field to given value.


### GetUrlInfo

`func (o *RegisterSalesChannelInput) GetUrlInfo() UrlInfoInput`

GetUrlInfo returns the UrlInfo field if non-nil, zero value otherwise.

### GetUrlInfoOk

`func (o *RegisterSalesChannelInput) GetUrlInfoOk() (*UrlInfoInput, bool)`

GetUrlInfoOk returns a tuple with the UrlInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlInfo

`func (o *RegisterSalesChannelInput) SetUrlInfo(v UrlInfoInput)`

SetUrlInfo sets UrlInfo field to given value.


### GetUrlWebhook

`func (o *RegisterSalesChannelInput) GetUrlWebhook() UrlWebhookInput`

GetUrlWebhook returns the UrlWebhook field if non-nil, zero value otherwise.

### GetUrlWebhookOk

`func (o *RegisterSalesChannelInput) GetUrlWebhookOk() (*UrlWebhookInput, bool)`

GetUrlWebhookOk returns a tuple with the UrlWebhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlWebhook

`func (o *RegisterSalesChannelInput) SetUrlWebhook(v UrlWebhookInput)`

SetUrlWebhook sets UrlWebhook field to given value.

### HasUrlWebhook

`func (o *RegisterSalesChannelInput) HasUrlWebhook() bool`

HasUrlWebhook returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


