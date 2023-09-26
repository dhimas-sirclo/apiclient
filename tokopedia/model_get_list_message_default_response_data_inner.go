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

// checks if the GetListMessageDefaultResponseDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetListMessageDefaultResponseDataInner{}

// GetListMessageDefaultResponseDataInner struct for GetListMessageDefaultResponseDataInner
type GetListMessageDefaultResponseDataInner struct {
	// Message Key, Field messages_key is an identic key for related buyer and related seller that communicate at the chat.
	MessageKey *string `json:"message_key,omitempty"`
	// Message Unique Identifier
	MsgId *int64 `json:"msg_id,omitempty"`
	Attributes *GetListMessageDefaultResponseDataInnerAttributes `json:"attributes,omitempty"`
}

// NewGetListMessageDefaultResponseDataInner instantiates a new GetListMessageDefaultResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetListMessageDefaultResponseDataInner() *GetListMessageDefaultResponseDataInner {
	this := GetListMessageDefaultResponseDataInner{}
	return &this
}

// NewGetListMessageDefaultResponseDataInnerWithDefaults instantiates a new GetListMessageDefaultResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetListMessageDefaultResponseDataInnerWithDefaults() *GetListMessageDefaultResponseDataInner {
	this := GetListMessageDefaultResponseDataInner{}
	return &this
}

// GetMessageKey returns the MessageKey field value if set, zero value otherwise.
func (o *GetListMessageDefaultResponseDataInner) GetMessageKey() string {
	if o == nil || IsNil(o.MessageKey) {
		var ret string
		return ret
	}
	return *o.MessageKey
}

// GetMessageKeyOk returns a tuple with the MessageKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetListMessageDefaultResponseDataInner) GetMessageKeyOk() (*string, bool) {
	if o == nil || IsNil(o.MessageKey) {
		return nil, false
	}
	return o.MessageKey, true
}

// HasMessageKey returns a boolean if a field has been set.
func (o *GetListMessageDefaultResponseDataInner) HasMessageKey() bool {
	if o != nil && !IsNil(o.MessageKey) {
		return true
	}

	return false
}

// SetMessageKey gets a reference to the given string and assigns it to the MessageKey field.
func (o *GetListMessageDefaultResponseDataInner) SetMessageKey(v string) {
	o.MessageKey = &v
}

// GetMsgId returns the MsgId field value if set, zero value otherwise.
func (o *GetListMessageDefaultResponseDataInner) GetMsgId() int64 {
	if o == nil || IsNil(o.MsgId) {
		var ret int64
		return ret
	}
	return *o.MsgId
}

// GetMsgIdOk returns a tuple with the MsgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetListMessageDefaultResponseDataInner) GetMsgIdOk() (*int64, bool) {
	if o == nil || IsNil(o.MsgId) {
		return nil, false
	}
	return o.MsgId, true
}

// HasMsgId returns a boolean if a field has been set.
func (o *GetListMessageDefaultResponseDataInner) HasMsgId() bool {
	if o != nil && !IsNil(o.MsgId) {
		return true
	}

	return false
}

// SetMsgId gets a reference to the given int64 and assigns it to the MsgId field.
func (o *GetListMessageDefaultResponseDataInner) SetMsgId(v int64) {
	o.MsgId = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *GetListMessageDefaultResponseDataInner) GetAttributes() GetListMessageDefaultResponseDataInnerAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret GetListMessageDefaultResponseDataInnerAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetListMessageDefaultResponseDataInner) GetAttributesOk() (*GetListMessageDefaultResponseDataInnerAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *GetListMessageDefaultResponseDataInner) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given GetListMessageDefaultResponseDataInnerAttributes and assigns it to the Attributes field.
func (o *GetListMessageDefaultResponseDataInner) SetAttributes(v GetListMessageDefaultResponseDataInnerAttributes) {
	o.Attributes = &v
}

func (o GetListMessageDefaultResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetListMessageDefaultResponseDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MessageKey) {
		toSerialize["message_key"] = o.MessageKey
	}
	if !IsNil(o.MsgId) {
		toSerialize["msg_id"] = o.MsgId
	}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableGetListMessageDefaultResponseDataInner struct {
	value *GetListMessageDefaultResponseDataInner
	isSet bool
}

func (v NullableGetListMessageDefaultResponseDataInner) Get() *GetListMessageDefaultResponseDataInner {
	return v.value
}

func (v *NullableGetListMessageDefaultResponseDataInner) Set(val *GetListMessageDefaultResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetListMessageDefaultResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetListMessageDefaultResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetListMessageDefaultResponseDataInner(val *GetListMessageDefaultResponseDataInner) *NullableGetListMessageDefaultResponseDataInner {
	return &NullableGetListMessageDefaultResponseDataInner{value: val, isSet: true}
}

func (v NullableGetListMessageDefaultResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetListMessageDefaultResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


