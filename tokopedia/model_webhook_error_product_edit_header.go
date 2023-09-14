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

// checks if the WebhookErrorProductEditHeader type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookErrorProductEditHeader{}

// WebhookErrorProductEditHeader struct for WebhookErrorProductEditHeader
type WebhookErrorProductEditHeader struct {
	ProcessTime *int64 `json:"process_time,omitempty"`
	Messages *string `json:"messages,omitempty"`
}

// NewWebhookErrorProductEditHeader instantiates a new WebhookErrorProductEditHeader object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookErrorProductEditHeader() *WebhookErrorProductEditHeader {
	this := WebhookErrorProductEditHeader{}
	return &this
}

// NewWebhookErrorProductEditHeaderWithDefaults instantiates a new WebhookErrorProductEditHeader object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookErrorProductEditHeaderWithDefaults() *WebhookErrorProductEditHeader {
	this := WebhookErrorProductEditHeader{}
	return &this
}

// GetProcessTime returns the ProcessTime field value if set, zero value otherwise.
func (o *WebhookErrorProductEditHeader) GetProcessTime() int64 {
	if o == nil || IsNil(o.ProcessTime) {
		var ret int64
		return ret
	}
	return *o.ProcessTime
}

// GetProcessTimeOk returns a tuple with the ProcessTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookErrorProductEditHeader) GetProcessTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.ProcessTime) {
		return nil, false
	}
	return o.ProcessTime, true
}

// HasProcessTime returns a boolean if a field has been set.
func (o *WebhookErrorProductEditHeader) HasProcessTime() bool {
	if o != nil && !IsNil(o.ProcessTime) {
		return true
	}

	return false
}

// SetProcessTime gets a reference to the given int64 and assigns it to the ProcessTime field.
func (o *WebhookErrorProductEditHeader) SetProcessTime(v int64) {
	o.ProcessTime = &v
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *WebhookErrorProductEditHeader) GetMessages() string {
	if o == nil || IsNil(o.Messages) {
		var ret string
		return ret
	}
	return *o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookErrorProductEditHeader) GetMessagesOk() (*string, bool) {
	if o == nil || IsNil(o.Messages) {
		return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *WebhookErrorProductEditHeader) HasMessages() bool {
	if o != nil && !IsNil(o.Messages) {
		return true
	}

	return false
}

// SetMessages gets a reference to the given string and assigns it to the Messages field.
func (o *WebhookErrorProductEditHeader) SetMessages(v string) {
	o.Messages = &v
}

func (o WebhookErrorProductEditHeader) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookErrorProductEditHeader) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProcessTime) {
		toSerialize["process_time"] = o.ProcessTime
	}
	if !IsNil(o.Messages) {
		toSerialize["messages"] = o.Messages
	}
	return toSerialize, nil
}

type NullableWebhookErrorProductEditHeader struct {
	value *WebhookErrorProductEditHeader
	isSet bool
}

func (v NullableWebhookErrorProductEditHeader) Get() *WebhookErrorProductEditHeader {
	return v.value
}

func (v *NullableWebhookErrorProductEditHeader) Set(val *WebhookErrorProductEditHeader) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookErrorProductEditHeader) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookErrorProductEditHeader) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookErrorProductEditHeader(val *WebhookErrorProductEditHeader) *NullableWebhookErrorProductEditHeader {
	return &NullableWebhookErrorProductEditHeader{value: val, isSet: true}
}

func (v NullableWebhookErrorProductEditHeader) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookErrorProductEditHeader) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


