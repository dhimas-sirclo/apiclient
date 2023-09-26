# GetListMessageDefaultResponseDataInnerAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contact** | Pointer to [**GetListMessageDefaultResponseDataInnerAttributesContact**](GetListMessageDefaultResponseDataInnerAttributesContact.md) |  | [optional] 
**LastReplyMsg** | Pointer to **string** | Last Reply Message | [optional] 
**LastReplyTime** | Pointer to **int64** | Last Reply Time in UNIX | [optional] 
**ReadStatus** | Pointer to **int64** | Read Status, Field read status is to identify that status &#x3D; 1 is not read yet and status &#x3D; 2 is already read. | [optional] 
**Unreads** | Pointer to **int64** | Unreads Count, Field unreads is the number of unread messages in the related message. | [optional] 
**PinStatus** | Pointer to **int64** | Pin Status, Field pin_status is for chat pinned or not, when pinned it is 1 and for not it is 0. | [optional] 

## Methods

### NewGetListMessageDefaultResponseDataInnerAttributes

`func NewGetListMessageDefaultResponseDataInnerAttributes() *GetListMessageDefaultResponseDataInnerAttributes`

NewGetListMessageDefaultResponseDataInnerAttributes instantiates a new GetListMessageDefaultResponseDataInnerAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetListMessageDefaultResponseDataInnerAttributesWithDefaults

`func NewGetListMessageDefaultResponseDataInnerAttributesWithDefaults() *GetListMessageDefaultResponseDataInnerAttributes`

NewGetListMessageDefaultResponseDataInnerAttributesWithDefaults instantiates a new GetListMessageDefaultResponseDataInnerAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContact

`func (o *GetListMessageDefaultResponseDataInnerAttributes) GetContact() GetListMessageDefaultResponseDataInnerAttributesContact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *GetListMessageDefaultResponseDataInnerAttributes) GetContactOk() (*GetListMessageDefaultResponseDataInnerAttributesContact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *GetListMessageDefaultResponseDataInnerAttributes) SetContact(v GetListMessageDefaultResponseDataInnerAttributesContact)`

SetContact sets Contact field to given value.

### HasContact

`func (o *GetListMessageDefaultResponseDataInnerAttributes) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetLastReplyMsg

`func (o *GetListMessageDefaultResponseDataInnerAttributes) GetLastReplyMsg() string`

GetLastReplyMsg returns the LastReplyMsg field if non-nil, zero value otherwise.

### GetLastReplyMsgOk

`func (o *GetListMessageDefaultResponseDataInnerAttributes) GetLastReplyMsgOk() (*string, bool)`

GetLastReplyMsgOk returns a tuple with the LastReplyMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReplyMsg

`func (o *GetListMessageDefaultResponseDataInnerAttributes) SetLastReplyMsg(v string)`

SetLastReplyMsg sets LastReplyMsg field to given value.

### HasLastReplyMsg

`func (o *GetListMessageDefaultResponseDataInnerAttributes) HasLastReplyMsg() bool`

HasLastReplyMsg returns a boolean if a field has been set.

### GetLastReplyTime

`func (o *GetListMessageDefaultResponseDataInnerAttributes) GetLastReplyTime() int64`

GetLastReplyTime returns the LastReplyTime field if non-nil, zero value otherwise.

### GetLastReplyTimeOk

`func (o *GetListMessageDefaultResponseDataInnerAttributes) GetLastReplyTimeOk() (*int64, bool)`

GetLastReplyTimeOk returns a tuple with the LastReplyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReplyTime

`func (o *GetListMessageDefaultResponseDataInnerAttributes) SetLastReplyTime(v int64)`

SetLastReplyTime sets LastReplyTime field to given value.

### HasLastReplyTime

`func (o *GetListMessageDefaultResponseDataInnerAttributes) HasLastReplyTime() bool`

HasLastReplyTime returns a boolean if a field has been set.

### GetReadStatus

`func (o *GetListMessageDefaultResponseDataInnerAttributes) GetReadStatus() int64`

GetReadStatus returns the ReadStatus field if non-nil, zero value otherwise.

### GetReadStatusOk

`func (o *GetListMessageDefaultResponseDataInnerAttributes) GetReadStatusOk() (*int64, bool)`

GetReadStatusOk returns a tuple with the ReadStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadStatus

`func (o *GetListMessageDefaultResponseDataInnerAttributes) SetReadStatus(v int64)`

SetReadStatus sets ReadStatus field to given value.

### HasReadStatus

`func (o *GetListMessageDefaultResponseDataInnerAttributes) HasReadStatus() bool`

HasReadStatus returns a boolean if a field has been set.

### GetUnreads

`func (o *GetListMessageDefaultResponseDataInnerAttributes) GetUnreads() int64`

GetUnreads returns the Unreads field if non-nil, zero value otherwise.

### GetUnreadsOk

`func (o *GetListMessageDefaultResponseDataInnerAttributes) GetUnreadsOk() (*int64, bool)`

GetUnreadsOk returns a tuple with the Unreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnreads

`func (o *GetListMessageDefaultResponseDataInnerAttributes) SetUnreads(v int64)`

SetUnreads sets Unreads field to given value.

### HasUnreads

`func (o *GetListMessageDefaultResponseDataInnerAttributes) HasUnreads() bool`

HasUnreads returns a boolean if a field has been set.

### GetPinStatus

`func (o *GetListMessageDefaultResponseDataInnerAttributes) GetPinStatus() int64`

GetPinStatus returns the PinStatus field if non-nil, zero value otherwise.

### GetPinStatusOk

`func (o *GetListMessageDefaultResponseDataInnerAttributes) GetPinStatusOk() (*int64, bool)`

GetPinStatusOk returns a tuple with the PinStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinStatus

`func (o *GetListMessageDefaultResponseDataInnerAttributes) SetPinStatus(v int64)`

SetPinStatus sets PinStatus field to given value.

### HasPinStatus

`func (o *GetListMessageDefaultResponseDataInnerAttributes) HasPinStatus() bool`

HasPinStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


