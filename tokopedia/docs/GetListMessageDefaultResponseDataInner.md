# GetListMessageDefaultResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageKey** | Pointer to **string** | Message Key, Field messages_key is an identic key for related buyer and related seller that communicate at the chat. | [optional] 
**MsgId** | Pointer to **int64** | Message Unique Identifier | [optional] 
**Attributes** | Pointer to [**GetListMessageDefaultResponseDataInnerAttributes**](GetListMessageDefaultResponseDataInnerAttributes.md) |  | [optional] 

## Methods

### NewGetListMessageDefaultResponseDataInner

`func NewGetListMessageDefaultResponseDataInner() *GetListMessageDefaultResponseDataInner`

NewGetListMessageDefaultResponseDataInner instantiates a new GetListMessageDefaultResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetListMessageDefaultResponseDataInnerWithDefaults

`func NewGetListMessageDefaultResponseDataInnerWithDefaults() *GetListMessageDefaultResponseDataInner`

NewGetListMessageDefaultResponseDataInnerWithDefaults instantiates a new GetListMessageDefaultResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageKey

`func (o *GetListMessageDefaultResponseDataInner) GetMessageKey() string`

GetMessageKey returns the MessageKey field if non-nil, zero value otherwise.

### GetMessageKeyOk

`func (o *GetListMessageDefaultResponseDataInner) GetMessageKeyOk() (*string, bool)`

GetMessageKeyOk returns a tuple with the MessageKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageKey

`func (o *GetListMessageDefaultResponseDataInner) SetMessageKey(v string)`

SetMessageKey sets MessageKey field to given value.

### HasMessageKey

`func (o *GetListMessageDefaultResponseDataInner) HasMessageKey() bool`

HasMessageKey returns a boolean if a field has been set.

### GetMsgId

`func (o *GetListMessageDefaultResponseDataInner) GetMsgId() int64`

GetMsgId returns the MsgId field if non-nil, zero value otherwise.

### GetMsgIdOk

`func (o *GetListMessageDefaultResponseDataInner) GetMsgIdOk() (*int64, bool)`

GetMsgIdOk returns a tuple with the MsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgId

`func (o *GetListMessageDefaultResponseDataInner) SetMsgId(v int64)`

SetMsgId sets MsgId field to given value.

### HasMsgId

`func (o *GetListMessageDefaultResponseDataInner) HasMsgId() bool`

HasMsgId returns a boolean if a field has been set.

### GetAttributes

`func (o *GetListMessageDefaultResponseDataInner) GetAttributes() GetListMessageDefaultResponseDataInnerAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GetListMessageDefaultResponseDataInner) GetAttributesOk() (*GetListMessageDefaultResponseDataInnerAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GetListMessageDefaultResponseDataInner) SetAttributes(v GetListMessageDefaultResponseDataInnerAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *GetListMessageDefaultResponseDataInner) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


