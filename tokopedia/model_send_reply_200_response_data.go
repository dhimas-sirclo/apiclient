/*
Tokopedia API

Tokopedia API

API version: 1.0
Contact: dev@sirclo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tokopedia

import (
	"encoding/json"
)

// checks if the SendReply200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SendReply200ResponseData{}

// SendReply200ResponseData struct for SendReply200ResponseData
type SendReply200ResponseData struct {
	MsgId *int64 `json:"msg_id,omitempty"`
	SenderId *int64 `json:"sender_id,omitempty"`
	Role *int64 `json:"role,omitempty"`
	Msg *string `json:"msg,omitempty"`
	ReplyTime *int64 `json:"reply_time,omitempty"`
	From *string `json:"from,omitempty"`
	Attachment *SendReply200ResponseDataAttachment `json:"attachment,omitempty"`
	FallbackAttachment *SendReply200ResponseDataFallbackAttachment `json:"fallback_attachment,omitempty"`
}

// NewSendReply200ResponseData instantiates a new SendReply200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSendReply200ResponseData() *SendReply200ResponseData {
	this := SendReply200ResponseData{}
	return &this
}

// NewSendReply200ResponseDataWithDefaults instantiates a new SendReply200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSendReply200ResponseDataWithDefaults() *SendReply200ResponseData {
	this := SendReply200ResponseData{}
	return &this
}

// GetMsgId returns the MsgId field value if set, zero value otherwise.
func (o *SendReply200ResponseData) GetMsgId() int64 {
	if o == nil || IsNil(o.MsgId) {
		var ret int64
		return ret
	}
	return *o.MsgId
}

// GetMsgIdOk returns a tuple with the MsgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendReply200ResponseData) GetMsgIdOk() (*int64, bool) {
	if o == nil || IsNil(o.MsgId) {
		return nil, false
	}
	return o.MsgId, true
}

// HasMsgId returns a boolean if a field has been set.
func (o *SendReply200ResponseData) HasMsgId() bool {
	if o != nil && !IsNil(o.MsgId) {
		return true
	}

	return false
}

// SetMsgId gets a reference to the given int64 and assigns it to the MsgId field.
func (o *SendReply200ResponseData) SetMsgId(v int64) {
	o.MsgId = &v
}

// GetSenderId returns the SenderId field value if set, zero value otherwise.
func (o *SendReply200ResponseData) GetSenderId() int64 {
	if o == nil || IsNil(o.SenderId) {
		var ret int64
		return ret
	}
	return *o.SenderId
}

// GetSenderIdOk returns a tuple with the SenderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendReply200ResponseData) GetSenderIdOk() (*int64, bool) {
	if o == nil || IsNil(o.SenderId) {
		return nil, false
	}
	return o.SenderId, true
}

// HasSenderId returns a boolean if a field has been set.
func (o *SendReply200ResponseData) HasSenderId() bool {
	if o != nil && !IsNil(o.SenderId) {
		return true
	}

	return false
}

// SetSenderId gets a reference to the given int64 and assigns it to the SenderId field.
func (o *SendReply200ResponseData) SetSenderId(v int64) {
	o.SenderId = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *SendReply200ResponseData) GetRole() int64 {
	if o == nil || IsNil(o.Role) {
		var ret int64
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendReply200ResponseData) GetRoleOk() (*int64, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *SendReply200ResponseData) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given int64 and assigns it to the Role field.
func (o *SendReply200ResponseData) SetRole(v int64) {
	o.Role = &v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *SendReply200ResponseData) GetMsg() string {
	if o == nil || IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendReply200ResponseData) GetMsgOk() (*string, bool) {
	if o == nil || IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *SendReply200ResponseData) HasMsg() bool {
	if o != nil && !IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *SendReply200ResponseData) SetMsg(v string) {
	o.Msg = &v
}

// GetReplyTime returns the ReplyTime field value if set, zero value otherwise.
func (o *SendReply200ResponseData) GetReplyTime() int64 {
	if o == nil || IsNil(o.ReplyTime) {
		var ret int64
		return ret
	}
	return *o.ReplyTime
}

// GetReplyTimeOk returns a tuple with the ReplyTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendReply200ResponseData) GetReplyTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.ReplyTime) {
		return nil, false
	}
	return o.ReplyTime, true
}

// HasReplyTime returns a boolean if a field has been set.
func (o *SendReply200ResponseData) HasReplyTime() bool {
	if o != nil && !IsNil(o.ReplyTime) {
		return true
	}

	return false
}

// SetReplyTime gets a reference to the given int64 and assigns it to the ReplyTime field.
func (o *SendReply200ResponseData) SetReplyTime(v int64) {
	o.ReplyTime = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *SendReply200ResponseData) GetFrom() string {
	if o == nil || IsNil(o.From) {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendReply200ResponseData) GetFromOk() (*string, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *SendReply200ResponseData) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *SendReply200ResponseData) SetFrom(v string) {
	o.From = &v
}

// GetAttachment returns the Attachment field value if set, zero value otherwise.
func (o *SendReply200ResponseData) GetAttachment() SendReply200ResponseDataAttachment {
	if o == nil || IsNil(o.Attachment) {
		var ret SendReply200ResponseDataAttachment
		return ret
	}
	return *o.Attachment
}

// GetAttachmentOk returns a tuple with the Attachment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendReply200ResponseData) GetAttachmentOk() (*SendReply200ResponseDataAttachment, bool) {
	if o == nil || IsNil(o.Attachment) {
		return nil, false
	}
	return o.Attachment, true
}

// HasAttachment returns a boolean if a field has been set.
func (o *SendReply200ResponseData) HasAttachment() bool {
	if o != nil && !IsNil(o.Attachment) {
		return true
	}

	return false
}

// SetAttachment gets a reference to the given SendReply200ResponseDataAttachment and assigns it to the Attachment field.
func (o *SendReply200ResponseData) SetAttachment(v SendReply200ResponseDataAttachment) {
	o.Attachment = &v
}

// GetFallbackAttachment returns the FallbackAttachment field value if set, zero value otherwise.
func (o *SendReply200ResponseData) GetFallbackAttachment() SendReply200ResponseDataFallbackAttachment {
	if o == nil || IsNil(o.FallbackAttachment) {
		var ret SendReply200ResponseDataFallbackAttachment
		return ret
	}
	return *o.FallbackAttachment
}

// GetFallbackAttachmentOk returns a tuple with the FallbackAttachment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendReply200ResponseData) GetFallbackAttachmentOk() (*SendReply200ResponseDataFallbackAttachment, bool) {
	if o == nil || IsNil(o.FallbackAttachment) {
		return nil, false
	}
	return o.FallbackAttachment, true
}

// HasFallbackAttachment returns a boolean if a field has been set.
func (o *SendReply200ResponseData) HasFallbackAttachment() bool {
	if o != nil && !IsNil(o.FallbackAttachment) {
		return true
	}

	return false
}

// SetFallbackAttachment gets a reference to the given SendReply200ResponseDataFallbackAttachment and assigns it to the FallbackAttachment field.
func (o *SendReply200ResponseData) SetFallbackAttachment(v SendReply200ResponseDataFallbackAttachment) {
	o.FallbackAttachment = &v
}

func (o SendReply200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SendReply200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MsgId) {
		toSerialize["msg_id"] = o.MsgId
	}
	if !IsNil(o.SenderId) {
		toSerialize["sender_id"] = o.SenderId
	}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if !IsNil(o.Msg) {
		toSerialize["msg"] = o.Msg
	}
	if !IsNil(o.ReplyTime) {
		toSerialize["reply_time"] = o.ReplyTime
	}
	if !IsNil(o.From) {
		toSerialize["from"] = o.From
	}
	if !IsNil(o.Attachment) {
		toSerialize["attachment"] = o.Attachment
	}
	if !IsNil(o.FallbackAttachment) {
		toSerialize["fallback_attachment"] = o.FallbackAttachment
	}
	return toSerialize, nil
}

type NullableSendReply200ResponseData struct {
	value *SendReply200ResponseData
	isSet bool
}

func (v NullableSendReply200ResponseData) Get() *SendReply200ResponseData {
	return v.value
}

func (v *NullableSendReply200ResponseData) Set(val *SendReply200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableSendReply200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableSendReply200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSendReply200ResponseData(val *SendReply200ResponseData) *NullableSendReply200ResponseData {
	return &NullableSendReply200ResponseData{value: val, isSet: true}
}

func (v NullableSendReply200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSendReply200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


