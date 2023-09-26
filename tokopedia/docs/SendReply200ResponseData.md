# SendReply200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MsgId** | Pointer to **int64** |  | [optional] 
**SenderId** | Pointer to **int64** |  | [optional] 
**Role** | Pointer to **int64** |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**ReplyTime** | Pointer to **int64** |  | [optional] 
**From** | Pointer to **string** |  | [optional] 
**Attachment** | Pointer to [**SendReply200ResponseDataAttachment**](SendReply200ResponseDataAttachment.md) |  | [optional] 
**FallbackAttachment** | Pointer to [**SendReply200ResponseDataFallbackAttachment**](SendReply200ResponseDataFallbackAttachment.md) |  | [optional] 

## Methods

### NewSendReply200ResponseData

`func NewSendReply200ResponseData() *SendReply200ResponseData`

NewSendReply200ResponseData instantiates a new SendReply200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendReply200ResponseDataWithDefaults

`func NewSendReply200ResponseDataWithDefaults() *SendReply200ResponseData`

NewSendReply200ResponseDataWithDefaults instantiates a new SendReply200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsgId

`func (o *SendReply200ResponseData) GetMsgId() int64`

GetMsgId returns the MsgId field if non-nil, zero value otherwise.

### GetMsgIdOk

`func (o *SendReply200ResponseData) GetMsgIdOk() (*int64, bool)`

GetMsgIdOk returns a tuple with the MsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgId

`func (o *SendReply200ResponseData) SetMsgId(v int64)`

SetMsgId sets MsgId field to given value.

### HasMsgId

`func (o *SendReply200ResponseData) HasMsgId() bool`

HasMsgId returns a boolean if a field has been set.

### GetSenderId

`func (o *SendReply200ResponseData) GetSenderId() int64`

GetSenderId returns the SenderId field if non-nil, zero value otherwise.

### GetSenderIdOk

`func (o *SendReply200ResponseData) GetSenderIdOk() (*int64, bool)`

GetSenderIdOk returns a tuple with the SenderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderId

`func (o *SendReply200ResponseData) SetSenderId(v int64)`

SetSenderId sets SenderId field to given value.

### HasSenderId

`func (o *SendReply200ResponseData) HasSenderId() bool`

HasSenderId returns a boolean if a field has been set.

### GetRole

`func (o *SendReply200ResponseData) GetRole() int64`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *SendReply200ResponseData) GetRoleOk() (*int64, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *SendReply200ResponseData) SetRole(v int64)`

SetRole sets Role field to given value.

### HasRole

`func (o *SendReply200ResponseData) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetMsg

`func (o *SendReply200ResponseData) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *SendReply200ResponseData) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *SendReply200ResponseData) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *SendReply200ResponseData) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetReplyTime

`func (o *SendReply200ResponseData) GetReplyTime() int64`

GetReplyTime returns the ReplyTime field if non-nil, zero value otherwise.

### GetReplyTimeOk

`func (o *SendReply200ResponseData) GetReplyTimeOk() (*int64, bool)`

GetReplyTimeOk returns a tuple with the ReplyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyTime

`func (o *SendReply200ResponseData) SetReplyTime(v int64)`

SetReplyTime sets ReplyTime field to given value.

### HasReplyTime

`func (o *SendReply200ResponseData) HasReplyTime() bool`

HasReplyTime returns a boolean if a field has been set.

### GetFrom

`func (o *SendReply200ResponseData) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *SendReply200ResponseData) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *SendReply200ResponseData) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *SendReply200ResponseData) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetAttachment

`func (o *SendReply200ResponseData) GetAttachment() SendReply200ResponseDataAttachment`

GetAttachment returns the Attachment field if non-nil, zero value otherwise.

### GetAttachmentOk

`func (o *SendReply200ResponseData) GetAttachmentOk() (*SendReply200ResponseDataAttachment, bool)`

GetAttachmentOk returns a tuple with the Attachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachment

`func (o *SendReply200ResponseData) SetAttachment(v SendReply200ResponseDataAttachment)`

SetAttachment sets Attachment field to given value.

### HasAttachment

`func (o *SendReply200ResponseData) HasAttachment() bool`

HasAttachment returns a boolean if a field has been set.

### GetFallbackAttachment

`func (o *SendReply200ResponseData) GetFallbackAttachment() SendReply200ResponseDataFallbackAttachment`

GetFallbackAttachment returns the FallbackAttachment field if non-nil, zero value otherwise.

### GetFallbackAttachmentOk

`func (o *SendReply200ResponseData) GetFallbackAttachmentOk() (*SendReply200ResponseDataFallbackAttachment, bool)`

GetFallbackAttachmentOk returns a tuple with the FallbackAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackAttachment

`func (o *SendReply200ResponseData) SetFallbackAttachment(v SendReply200ResponseDataFallbackAttachment)`

SetFallbackAttachment sets FallbackAttachment field to given value.

### HasFallbackAttachment

`func (o *SendReply200ResponseData) HasFallbackAttachment() bool`

HasFallbackAttachment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


