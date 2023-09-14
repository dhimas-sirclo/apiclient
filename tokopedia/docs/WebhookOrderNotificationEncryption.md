# WebhookOrderNotificationEncryption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Secret** | Pointer to **string** | Contains String used to generate key for content | [optional] 
**Content** | Pointer to **string** | Contains encrypted buyer information | [optional] 

## Methods

### NewWebhookOrderNotificationEncryption

`func NewWebhookOrderNotificationEncryption() *WebhookOrderNotificationEncryption`

NewWebhookOrderNotificationEncryption instantiates a new WebhookOrderNotificationEncryption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookOrderNotificationEncryptionWithDefaults

`func NewWebhookOrderNotificationEncryptionWithDefaults() *WebhookOrderNotificationEncryption`

NewWebhookOrderNotificationEncryptionWithDefaults instantiates a new WebhookOrderNotificationEncryption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecret

`func (o *WebhookOrderNotificationEncryption) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *WebhookOrderNotificationEncryption) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *WebhookOrderNotificationEncryption) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *WebhookOrderNotificationEncryption) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetContent

`func (o *WebhookOrderNotificationEncryption) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *WebhookOrderNotificationEncryption) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *WebhookOrderNotificationEncryption) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *WebhookOrderNotificationEncryption) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


