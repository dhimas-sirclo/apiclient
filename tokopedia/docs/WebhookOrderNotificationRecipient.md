# WebhookOrderNotificationRecipient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Recipient Name | [optional] 
**Phone** | Pointer to **string** | Recipient Phone Number | [optional] 
**Address** | Pointer to [**WebhookOrderNotificationRecipientAddress**](WebhookOrderNotificationRecipientAddress.md) |  | [optional] 

## Methods

### NewWebhookOrderNotificationRecipient

`func NewWebhookOrderNotificationRecipient() *WebhookOrderNotificationRecipient`

NewWebhookOrderNotificationRecipient instantiates a new WebhookOrderNotificationRecipient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookOrderNotificationRecipientWithDefaults

`func NewWebhookOrderNotificationRecipientWithDefaults() *WebhookOrderNotificationRecipient`

NewWebhookOrderNotificationRecipientWithDefaults instantiates a new WebhookOrderNotificationRecipient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WebhookOrderNotificationRecipient) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebhookOrderNotificationRecipient) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebhookOrderNotificationRecipient) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WebhookOrderNotificationRecipient) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPhone

`func (o *WebhookOrderNotificationRecipient) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *WebhookOrderNotificationRecipient) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *WebhookOrderNotificationRecipient) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *WebhookOrderNotificationRecipient) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetAddress

`func (o *WebhookOrderNotificationRecipient) GetAddress() WebhookOrderNotificationRecipientAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *WebhookOrderNotificationRecipient) GetAddressOk() (*WebhookOrderNotificationRecipientAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *WebhookOrderNotificationRecipient) SetAddress(v WebhookOrderNotificationRecipientAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *WebhookOrderNotificationRecipient) HasAddress() bool`

HasAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


