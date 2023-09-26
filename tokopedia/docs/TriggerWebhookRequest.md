# TriggerWebhookRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | **int64** |  | 
**Type** | **string** |  | 
**Url** | **string** |  | 
**IsEncrypted** | **bool** |  | 

## Methods

### NewTriggerWebhookRequest

`func NewTriggerWebhookRequest(orderId int64, type_ string, url string, isEncrypted bool, ) *TriggerWebhookRequest`

NewTriggerWebhookRequest instantiates a new TriggerWebhookRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerWebhookRequestWithDefaults

`func NewTriggerWebhookRequestWithDefaults() *TriggerWebhookRequest`

NewTriggerWebhookRequestWithDefaults instantiates a new TriggerWebhookRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *TriggerWebhookRequest) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *TriggerWebhookRequest) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *TriggerWebhookRequest) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.


### GetType

`func (o *TriggerWebhookRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TriggerWebhookRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TriggerWebhookRequest) SetType(v string)`

SetType sets Type field to given value.


### GetUrl

`func (o *TriggerWebhookRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TriggerWebhookRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TriggerWebhookRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetIsEncrypted

`func (o *TriggerWebhookRequest) GetIsEncrypted() bool`

GetIsEncrypted returns the IsEncrypted field if non-nil, zero value otherwise.

### GetIsEncryptedOk

`func (o *TriggerWebhookRequest) GetIsEncryptedOk() (*bool, bool)`

GetIsEncryptedOk returns a tuple with the IsEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEncrypted

`func (o *TriggerWebhookRequest) SetIsEncrypted(v bool)`

SetIsEncrypted sets IsEncrypted field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


