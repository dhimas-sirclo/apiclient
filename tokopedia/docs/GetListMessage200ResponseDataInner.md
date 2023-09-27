# GetListMessage200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageKey** | Pointer to **string** | Message Key, Field messages_key is an identic key for related buyer and related seller that communicate at the chat. | [optional] 
**MsgId** | Pointer to **int64** | Message Unique Identifier | [optional] 
**Attributes** | Pointer to [**GetListMessage200ResponseDataInnerAttributes**](GetListMessage200ResponseDataInnerAttributes.md) |  | [optional] 

## Methods

### NewGetListMessage200ResponseDataInner

`func NewGetListMessage200ResponseDataInner() *GetListMessage200ResponseDataInner`

NewGetListMessage200ResponseDataInner instantiates a new GetListMessage200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetListMessage200ResponseDataInnerWithDefaults

`func NewGetListMessage200ResponseDataInnerWithDefaults() *GetListMessage200ResponseDataInner`

NewGetListMessage200ResponseDataInnerWithDefaults instantiates a new GetListMessage200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageKey

`func (o *GetListMessage200ResponseDataInner) GetMessageKey() string`

GetMessageKey returns the MessageKey field if non-nil, zero value otherwise.

### GetMessageKeyOk

`func (o *GetListMessage200ResponseDataInner) GetMessageKeyOk() (*string, bool)`

GetMessageKeyOk returns a tuple with the MessageKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageKey

`func (o *GetListMessage200ResponseDataInner) SetMessageKey(v string)`

SetMessageKey sets MessageKey field to given value.

### HasMessageKey

`func (o *GetListMessage200ResponseDataInner) HasMessageKey() bool`

HasMessageKey returns a boolean if a field has been set.

### GetMsgId

`func (o *GetListMessage200ResponseDataInner) GetMsgId() int64`

GetMsgId returns the MsgId field if non-nil, zero value otherwise.

### GetMsgIdOk

`func (o *GetListMessage200ResponseDataInner) GetMsgIdOk() (*int64, bool)`

GetMsgIdOk returns a tuple with the MsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgId

`func (o *GetListMessage200ResponseDataInner) SetMsgId(v int64)`

SetMsgId sets MsgId field to given value.

### HasMsgId

`func (o *GetListMessage200ResponseDataInner) HasMsgId() bool`

HasMsgId returns a boolean if a field has been set.

### GetAttributes

`func (o *GetListMessage200ResponseDataInner) GetAttributes() GetListMessage200ResponseDataInnerAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GetListMessage200ResponseDataInner) GetAttributesOk() (*GetListMessage200ResponseDataInnerAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GetListMessage200ResponseDataInner) SetAttributes(v GetListMessage200ResponseDataInnerAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *GetListMessage200ResponseDataInner) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


