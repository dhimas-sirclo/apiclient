# WebhookChatReplyNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MsgId** | Pointer to **int64** | Message unique identifier | [optional] 
**BuyerId** | Pointer to **int64** | Buyer unique identifier | [optional] 
**Message** | Pointer to **string** | Message content | [optional] 
**Thumbnail** | Pointer to **string** | Chat&#39;s Thumbnail | [optional] 
**FullName** | Pointer to **string** | Opposite chat user&#39;s first name | [optional] 
**ShopId** | Pointer to **int64** | Shop unique identifier | [optional] 

## Methods

### NewWebhookChatReplyNotification

`func NewWebhookChatReplyNotification() *WebhookChatReplyNotification`

NewWebhookChatReplyNotification instantiates a new WebhookChatReplyNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookChatReplyNotificationWithDefaults

`func NewWebhookChatReplyNotificationWithDefaults() *WebhookChatReplyNotification`

NewWebhookChatReplyNotificationWithDefaults instantiates a new WebhookChatReplyNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsgId

`func (o *WebhookChatReplyNotification) GetMsgId() int64`

GetMsgId returns the MsgId field if non-nil, zero value otherwise.

### GetMsgIdOk

`func (o *WebhookChatReplyNotification) GetMsgIdOk() (*int64, bool)`

GetMsgIdOk returns a tuple with the MsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgId

`func (o *WebhookChatReplyNotification) SetMsgId(v int64)`

SetMsgId sets MsgId field to given value.

### HasMsgId

`func (o *WebhookChatReplyNotification) HasMsgId() bool`

HasMsgId returns a boolean if a field has been set.

### GetBuyerId

`func (o *WebhookChatReplyNotification) GetBuyerId() int64`

GetBuyerId returns the BuyerId field if non-nil, zero value otherwise.

### GetBuyerIdOk

`func (o *WebhookChatReplyNotification) GetBuyerIdOk() (*int64, bool)`

GetBuyerIdOk returns a tuple with the BuyerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerId

`func (o *WebhookChatReplyNotification) SetBuyerId(v int64)`

SetBuyerId sets BuyerId field to given value.

### HasBuyerId

`func (o *WebhookChatReplyNotification) HasBuyerId() bool`

HasBuyerId returns a boolean if a field has been set.

### GetMessage

`func (o *WebhookChatReplyNotification) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *WebhookChatReplyNotification) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *WebhookChatReplyNotification) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *WebhookChatReplyNotification) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetThumbnail

`func (o *WebhookChatReplyNotification) GetThumbnail() string`

GetThumbnail returns the Thumbnail field if non-nil, zero value otherwise.

### GetThumbnailOk

`func (o *WebhookChatReplyNotification) GetThumbnailOk() (*string, bool)`

GetThumbnailOk returns a tuple with the Thumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnail

`func (o *WebhookChatReplyNotification) SetThumbnail(v string)`

SetThumbnail sets Thumbnail field to given value.

### HasThumbnail

`func (o *WebhookChatReplyNotification) HasThumbnail() bool`

HasThumbnail returns a boolean if a field has been set.

### GetFullName

`func (o *WebhookChatReplyNotification) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *WebhookChatReplyNotification) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *WebhookChatReplyNotification) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *WebhookChatReplyNotification) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetShopId

`func (o *WebhookChatReplyNotification) GetShopId() int64`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *WebhookChatReplyNotification) GetShopIdOk() (*int64, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *WebhookChatReplyNotification) SetShopId(v int64)`

SetShopId sets ShopId field to given value.

### HasShopId

`func (o *WebhookChatReplyNotification) HasShopId() bool`

HasShopId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


