# GetListMessage200ResponseDataInnerAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contact** | Pointer to [**GetListMessage200ResponseDataInnerAttributesContact**](GetListMessage200ResponseDataInnerAttributesContact.md) |  | [optional] 
**LastReplyMsg** | Pointer to **string** | Last Reply Message | [optional] 
**LastReplyTime** | Pointer to **int64** | Last Reply Time in UNIX | [optional] 
**ReadStatus** | Pointer to **int64** | Read Status, Field read status is to identify that status &#x3D; 1 is not read yet and status &#x3D; 2 is already read. | [optional] 
**Unreads** | Pointer to **int64** | Unreads Count, Field unreads is the number of unread messages in the related message. | [optional] 
**PinStatus** | Pointer to **int64** | Pin Status, Field pin_status is for chat pinned or not, when pinned it is 1 and for not it is 0. | [optional] 

## Methods

### NewGetListMessage200ResponseDataInnerAttributes

`func NewGetListMessage200ResponseDataInnerAttributes() *GetListMessage200ResponseDataInnerAttributes`

NewGetListMessage200ResponseDataInnerAttributes instantiates a new GetListMessage200ResponseDataInnerAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetListMessage200ResponseDataInnerAttributesWithDefaults

`func NewGetListMessage200ResponseDataInnerAttributesWithDefaults() *GetListMessage200ResponseDataInnerAttributes`

NewGetListMessage200ResponseDataInnerAttributesWithDefaults instantiates a new GetListMessage200ResponseDataInnerAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContact

`func (o *GetListMessage200ResponseDataInnerAttributes) GetContact() GetListMessage200ResponseDataInnerAttributesContact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *GetListMessage200ResponseDataInnerAttributes) GetContactOk() (*GetListMessage200ResponseDataInnerAttributesContact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *GetListMessage200ResponseDataInnerAttributes) SetContact(v GetListMessage200ResponseDataInnerAttributesContact)`

SetContact sets Contact field to given value.

### HasContact

`func (o *GetListMessage200ResponseDataInnerAttributes) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetLastReplyMsg

`func (o *GetListMessage200ResponseDataInnerAttributes) GetLastReplyMsg() string`

GetLastReplyMsg returns the LastReplyMsg field if non-nil, zero value otherwise.

### GetLastReplyMsgOk

`func (o *GetListMessage200ResponseDataInnerAttributes) GetLastReplyMsgOk() (*string, bool)`

GetLastReplyMsgOk returns a tuple with the LastReplyMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReplyMsg

`func (o *GetListMessage200ResponseDataInnerAttributes) SetLastReplyMsg(v string)`

SetLastReplyMsg sets LastReplyMsg field to given value.

### HasLastReplyMsg

`func (o *GetListMessage200ResponseDataInnerAttributes) HasLastReplyMsg() bool`

HasLastReplyMsg returns a boolean if a field has been set.

### GetLastReplyTime

`func (o *GetListMessage200ResponseDataInnerAttributes) GetLastReplyTime() int64`

GetLastReplyTime returns the LastReplyTime field if non-nil, zero value otherwise.

### GetLastReplyTimeOk

`func (o *GetListMessage200ResponseDataInnerAttributes) GetLastReplyTimeOk() (*int64, bool)`

GetLastReplyTimeOk returns a tuple with the LastReplyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReplyTime

`func (o *GetListMessage200ResponseDataInnerAttributes) SetLastReplyTime(v int64)`

SetLastReplyTime sets LastReplyTime field to given value.

### HasLastReplyTime

`func (o *GetListMessage200ResponseDataInnerAttributes) HasLastReplyTime() bool`

HasLastReplyTime returns a boolean if a field has been set.

### GetReadStatus

`func (o *GetListMessage200ResponseDataInnerAttributes) GetReadStatus() int64`

GetReadStatus returns the ReadStatus field if non-nil, zero value otherwise.

### GetReadStatusOk

`func (o *GetListMessage200ResponseDataInnerAttributes) GetReadStatusOk() (*int64, bool)`

GetReadStatusOk returns a tuple with the ReadStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadStatus

`func (o *GetListMessage200ResponseDataInnerAttributes) SetReadStatus(v int64)`

SetReadStatus sets ReadStatus field to given value.

### HasReadStatus

`func (o *GetListMessage200ResponseDataInnerAttributes) HasReadStatus() bool`

HasReadStatus returns a boolean if a field has been set.

### GetUnreads

`func (o *GetListMessage200ResponseDataInnerAttributes) GetUnreads() int64`

GetUnreads returns the Unreads field if non-nil, zero value otherwise.

### GetUnreadsOk

`func (o *GetListMessage200ResponseDataInnerAttributes) GetUnreadsOk() (*int64, bool)`

GetUnreadsOk returns a tuple with the Unreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnreads

`func (o *GetListMessage200ResponseDataInnerAttributes) SetUnreads(v int64)`

SetUnreads sets Unreads field to given value.

### HasUnreads

`func (o *GetListMessage200ResponseDataInnerAttributes) HasUnreads() bool`

HasUnreads returns a boolean if a field has been set.

### GetPinStatus

`func (o *GetListMessage200ResponseDataInnerAttributes) GetPinStatus() int64`

GetPinStatus returns the PinStatus field if non-nil, zero value otherwise.

### GetPinStatusOk

`func (o *GetListMessage200ResponseDataInnerAttributes) GetPinStatusOk() (*int64, bool)`

GetPinStatusOk returns a tuple with the PinStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinStatus

`func (o *GetListMessage200ResponseDataInnerAttributes) SetPinStatus(v int64)`

SetPinStatus sets PinStatus field to given value.

### HasPinStatus

`func (o *GetListMessage200ResponseDataInnerAttributes) HasPinStatus() bool`

HasPinStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


