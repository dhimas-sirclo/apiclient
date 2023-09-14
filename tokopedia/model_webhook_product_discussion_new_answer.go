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

// checks if the WebhookProductDiscussionNewAnswer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookProductDiscussionNewAnswer{}

// WebhookProductDiscussionNewAnswer struct for WebhookProductDiscussionNewAnswer
type WebhookProductDiscussionNewAnswer struct {
	FsId *int64 `json:"fs_id,omitempty"`
	DiscussionType *string `json:"discussion_type,omitempty"`
	DiscussionData *WebhookProductDiscussionNewAnswerDiscussionData `json:"discussion_data,omitempty"`
}

// NewWebhookProductDiscussionNewAnswer instantiates a new WebhookProductDiscussionNewAnswer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookProductDiscussionNewAnswer() *WebhookProductDiscussionNewAnswer {
	this := WebhookProductDiscussionNewAnswer{}
	var discussionType string = "DiscussionAnswer"
	this.DiscussionType = &discussionType
	return &this
}

// NewWebhookProductDiscussionNewAnswerWithDefaults instantiates a new WebhookProductDiscussionNewAnswer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookProductDiscussionNewAnswerWithDefaults() *WebhookProductDiscussionNewAnswer {
	this := WebhookProductDiscussionNewAnswer{}
	var discussionType string = "DiscussionAnswer"
	this.DiscussionType = &discussionType
	return &this
}

// GetFsId returns the FsId field value if set, zero value otherwise.
func (o *WebhookProductDiscussionNewAnswer) GetFsId() int64 {
	if o == nil || IsNil(o.FsId) {
		var ret int64
		return ret
	}
	return *o.FsId
}

// GetFsIdOk returns a tuple with the FsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookProductDiscussionNewAnswer) GetFsIdOk() (*int64, bool) {
	if o == nil || IsNil(o.FsId) {
		return nil, false
	}
	return o.FsId, true
}

// HasFsId returns a boolean if a field has been set.
func (o *WebhookProductDiscussionNewAnswer) HasFsId() bool {
	if o != nil && !IsNil(o.FsId) {
		return true
	}

	return false
}

// SetFsId gets a reference to the given int64 and assigns it to the FsId field.
func (o *WebhookProductDiscussionNewAnswer) SetFsId(v int64) {
	o.FsId = &v
}

// GetDiscussionType returns the DiscussionType field value if set, zero value otherwise.
func (o *WebhookProductDiscussionNewAnswer) GetDiscussionType() string {
	if o == nil || IsNil(o.DiscussionType) {
		var ret string
		return ret
	}
	return *o.DiscussionType
}

// GetDiscussionTypeOk returns a tuple with the DiscussionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookProductDiscussionNewAnswer) GetDiscussionTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DiscussionType) {
		return nil, false
	}
	return o.DiscussionType, true
}

// HasDiscussionType returns a boolean if a field has been set.
func (o *WebhookProductDiscussionNewAnswer) HasDiscussionType() bool {
	if o != nil && !IsNil(o.DiscussionType) {
		return true
	}

	return false
}

// SetDiscussionType gets a reference to the given string and assigns it to the DiscussionType field.
func (o *WebhookProductDiscussionNewAnswer) SetDiscussionType(v string) {
	o.DiscussionType = &v
}

// GetDiscussionData returns the DiscussionData field value if set, zero value otherwise.
func (o *WebhookProductDiscussionNewAnswer) GetDiscussionData() WebhookProductDiscussionNewAnswerDiscussionData {
	if o == nil || IsNil(o.DiscussionData) {
		var ret WebhookProductDiscussionNewAnswerDiscussionData
		return ret
	}
	return *o.DiscussionData
}

// GetDiscussionDataOk returns a tuple with the DiscussionData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookProductDiscussionNewAnswer) GetDiscussionDataOk() (*WebhookProductDiscussionNewAnswerDiscussionData, bool) {
	if o == nil || IsNil(o.DiscussionData) {
		return nil, false
	}
	return o.DiscussionData, true
}

// HasDiscussionData returns a boolean if a field has been set.
func (o *WebhookProductDiscussionNewAnswer) HasDiscussionData() bool {
	if o != nil && !IsNil(o.DiscussionData) {
		return true
	}

	return false
}

// SetDiscussionData gets a reference to the given WebhookProductDiscussionNewAnswerDiscussionData and assigns it to the DiscussionData field.
func (o *WebhookProductDiscussionNewAnswer) SetDiscussionData(v WebhookProductDiscussionNewAnswerDiscussionData) {
	o.DiscussionData = &v
}

func (o WebhookProductDiscussionNewAnswer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookProductDiscussionNewAnswer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FsId) {
		toSerialize["fs_id"] = o.FsId
	}
	if !IsNil(o.DiscussionType) {
		toSerialize["discussion_type"] = o.DiscussionType
	}
	if !IsNil(o.DiscussionData) {
		toSerialize["discussion_data"] = o.DiscussionData
	}
	return toSerialize, nil
}

type NullableWebhookProductDiscussionNewAnswer struct {
	value *WebhookProductDiscussionNewAnswer
	isSet bool
}

func (v NullableWebhookProductDiscussionNewAnswer) Get() *WebhookProductDiscussionNewAnswer {
	return v.value
}

func (v *NullableWebhookProductDiscussionNewAnswer) Set(val *WebhookProductDiscussionNewAnswer) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookProductDiscussionNewAnswer) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookProductDiscussionNewAnswer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookProductDiscussionNewAnswer(val *WebhookProductDiscussionNewAnswer) *NullableWebhookProductDiscussionNewAnswer {
	return &NullableWebhookProductDiscussionNewAnswer{value: val, isSet: true}
}

func (v NullableWebhookProductDiscussionNewAnswer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookProductDiscussionNewAnswer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


