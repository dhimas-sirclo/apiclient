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

// checks if the GetListMessage200ResponseDataInnerAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetListMessage200ResponseDataInnerAttributes{}

// GetListMessage200ResponseDataInnerAttributes struct for GetListMessage200ResponseDataInnerAttributes
type GetListMessage200ResponseDataInnerAttributes struct {
	Contact *GetListMessage200ResponseDataInnerAttributesContact `json:"contact,omitempty"`
	// Last Reply Message
	LastReplyMsg *string `json:"last_reply_msg,omitempty"`
	// Last Reply Time in UNIX
	LastReplyTime *int64 `json:"last_reply_time,omitempty"`
	// Read Status, Field read status is to identify that status = 1 is not read yet and status = 2 is already read.
	ReadStatus *int64 `json:"read_status,omitempty"`
	// Unreads Count, Field unreads is the number of unread messages in the related message.
	Unreads *int64 `json:"unreads,omitempty"`
	// Pin Status, Field pin_status is for chat pinned or not, when pinned it is 1 and for not it is 0.
	PinStatus *int64 `json:"pin_status,omitempty"`
}

// NewGetListMessage200ResponseDataInnerAttributes instantiates a new GetListMessage200ResponseDataInnerAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetListMessage200ResponseDataInnerAttributes() *GetListMessage200ResponseDataInnerAttributes {
	this := GetListMessage200ResponseDataInnerAttributes{}
	return &this
}

// NewGetListMessage200ResponseDataInnerAttributesWithDefaults instantiates a new GetListMessage200ResponseDataInnerAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetListMessage200ResponseDataInnerAttributesWithDefaults() *GetListMessage200ResponseDataInnerAttributes {
	this := GetListMessage200ResponseDataInnerAttributes{}
	return &this
}

// GetContact returns the Contact field value if set, zero value otherwise.
func (o *GetListMessage200ResponseDataInnerAttributes) GetContact() GetListMessage200ResponseDataInnerAttributesContact {
	if o == nil || IsNil(o.Contact) {
		var ret GetListMessage200ResponseDataInnerAttributesContact
		return ret
	}
	return *o.Contact
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetListMessage200ResponseDataInnerAttributes) GetContactOk() (*GetListMessage200ResponseDataInnerAttributesContact, bool) {
	if o == nil || IsNil(o.Contact) {
		return nil, false
	}
	return o.Contact, true
}

// HasContact returns a boolean if a field has been set.
func (o *GetListMessage200ResponseDataInnerAttributes) HasContact() bool {
	if o != nil && !IsNil(o.Contact) {
		return true
	}

	return false
}

// SetContact gets a reference to the given GetListMessage200ResponseDataInnerAttributesContact and assigns it to the Contact field.
func (o *GetListMessage200ResponseDataInnerAttributes) SetContact(v GetListMessage200ResponseDataInnerAttributesContact) {
	o.Contact = &v
}

// GetLastReplyMsg returns the LastReplyMsg field value if set, zero value otherwise.
func (o *GetListMessage200ResponseDataInnerAttributes) GetLastReplyMsg() string {
	if o == nil || IsNil(o.LastReplyMsg) {
		var ret string
		return ret
	}
	return *o.LastReplyMsg
}

// GetLastReplyMsgOk returns a tuple with the LastReplyMsg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetListMessage200ResponseDataInnerAttributes) GetLastReplyMsgOk() (*string, bool) {
	if o == nil || IsNil(o.LastReplyMsg) {
		return nil, false
	}
	return o.LastReplyMsg, true
}

// HasLastReplyMsg returns a boolean if a field has been set.
func (o *GetListMessage200ResponseDataInnerAttributes) HasLastReplyMsg() bool {
	if o != nil && !IsNil(o.LastReplyMsg) {
		return true
	}

	return false
}

// SetLastReplyMsg gets a reference to the given string and assigns it to the LastReplyMsg field.
func (o *GetListMessage200ResponseDataInnerAttributes) SetLastReplyMsg(v string) {
	o.LastReplyMsg = &v
}

// GetLastReplyTime returns the LastReplyTime field value if set, zero value otherwise.
func (o *GetListMessage200ResponseDataInnerAttributes) GetLastReplyTime() int64 {
	if o == nil || IsNil(o.LastReplyTime) {
		var ret int64
		return ret
	}
	return *o.LastReplyTime
}

// GetLastReplyTimeOk returns a tuple with the LastReplyTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetListMessage200ResponseDataInnerAttributes) GetLastReplyTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.LastReplyTime) {
		return nil, false
	}
	return o.LastReplyTime, true
}

// HasLastReplyTime returns a boolean if a field has been set.
func (o *GetListMessage200ResponseDataInnerAttributes) HasLastReplyTime() bool {
	if o != nil && !IsNil(o.LastReplyTime) {
		return true
	}

	return false
}

// SetLastReplyTime gets a reference to the given int64 and assigns it to the LastReplyTime field.
func (o *GetListMessage200ResponseDataInnerAttributes) SetLastReplyTime(v int64) {
	o.LastReplyTime = &v
}

// GetReadStatus returns the ReadStatus field value if set, zero value otherwise.
func (o *GetListMessage200ResponseDataInnerAttributes) GetReadStatus() int64 {
	if o == nil || IsNil(o.ReadStatus) {
		var ret int64
		return ret
	}
	return *o.ReadStatus
}

// GetReadStatusOk returns a tuple with the ReadStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetListMessage200ResponseDataInnerAttributes) GetReadStatusOk() (*int64, bool) {
	if o == nil || IsNil(o.ReadStatus) {
		return nil, false
	}
	return o.ReadStatus, true
}

// HasReadStatus returns a boolean if a field has been set.
func (o *GetListMessage200ResponseDataInnerAttributes) HasReadStatus() bool {
	if o != nil && !IsNil(o.ReadStatus) {
		return true
	}

	return false
}

// SetReadStatus gets a reference to the given int64 and assigns it to the ReadStatus field.
func (o *GetListMessage200ResponseDataInnerAttributes) SetReadStatus(v int64) {
	o.ReadStatus = &v
}

// GetUnreads returns the Unreads field value if set, zero value otherwise.
func (o *GetListMessage200ResponseDataInnerAttributes) GetUnreads() int64 {
	if o == nil || IsNil(o.Unreads) {
		var ret int64
		return ret
	}
	return *o.Unreads
}

// GetUnreadsOk returns a tuple with the Unreads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetListMessage200ResponseDataInnerAttributes) GetUnreadsOk() (*int64, bool) {
	if o == nil || IsNil(o.Unreads) {
		return nil, false
	}
	return o.Unreads, true
}

// HasUnreads returns a boolean if a field has been set.
func (o *GetListMessage200ResponseDataInnerAttributes) HasUnreads() bool {
	if o != nil && !IsNil(o.Unreads) {
		return true
	}

	return false
}

// SetUnreads gets a reference to the given int64 and assigns it to the Unreads field.
func (o *GetListMessage200ResponseDataInnerAttributes) SetUnreads(v int64) {
	o.Unreads = &v
}

// GetPinStatus returns the PinStatus field value if set, zero value otherwise.
func (o *GetListMessage200ResponseDataInnerAttributes) GetPinStatus() int64 {
	if o == nil || IsNil(o.PinStatus) {
		var ret int64
		return ret
	}
	return *o.PinStatus
}

// GetPinStatusOk returns a tuple with the PinStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetListMessage200ResponseDataInnerAttributes) GetPinStatusOk() (*int64, bool) {
	if o == nil || IsNil(o.PinStatus) {
		return nil, false
	}
	return o.PinStatus, true
}

// HasPinStatus returns a boolean if a field has been set.
func (o *GetListMessage200ResponseDataInnerAttributes) HasPinStatus() bool {
	if o != nil && !IsNil(o.PinStatus) {
		return true
	}

	return false
}

// SetPinStatus gets a reference to the given int64 and assigns it to the PinStatus field.
func (o *GetListMessage200ResponseDataInnerAttributes) SetPinStatus(v int64) {
	o.PinStatus = &v
}

func (o GetListMessage200ResponseDataInnerAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetListMessage200ResponseDataInnerAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Contact) {
		toSerialize["contact"] = o.Contact
	}
	if !IsNil(o.LastReplyMsg) {
		toSerialize["last_reply_msg"] = o.LastReplyMsg
	}
	if !IsNil(o.LastReplyTime) {
		toSerialize["last_reply_time"] = o.LastReplyTime
	}
	if !IsNil(o.ReadStatus) {
		toSerialize["read_status"] = o.ReadStatus
	}
	if !IsNil(o.Unreads) {
		toSerialize["unreads"] = o.Unreads
	}
	if !IsNil(o.PinStatus) {
		toSerialize["pin_status"] = o.PinStatus
	}
	return toSerialize, nil
}

type NullableGetListMessage200ResponseDataInnerAttributes struct {
	value *GetListMessage200ResponseDataInnerAttributes
	isSet bool
}

func (v NullableGetListMessage200ResponseDataInnerAttributes) Get() *GetListMessage200ResponseDataInnerAttributes {
	return v.value
}

func (v *NullableGetListMessage200ResponseDataInnerAttributes) Set(val *GetListMessage200ResponseDataInnerAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGetListMessage200ResponseDataInnerAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGetListMessage200ResponseDataInnerAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetListMessage200ResponseDataInnerAttributes(val *GetListMessage200ResponseDataInnerAttributes) *NullableGetListMessage200ResponseDataInnerAttributes {
	return &NullableGetListMessage200ResponseDataInnerAttributes{value: val, isSet: true}
}

func (v NullableGetListMessage200ResponseDataInnerAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetListMessage200ResponseDataInnerAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


