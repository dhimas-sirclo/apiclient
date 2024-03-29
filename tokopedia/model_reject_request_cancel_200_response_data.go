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

// checks if the RejectRequestCancel200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RejectRequestCancel200ResponseData{}

// RejectRequestCancel200ResponseData struct for RejectRequestCancel200ResponseData
type RejectRequestCancel200ResponseData struct {
	Success *bool `json:"success,omitempty"`
}

// NewRejectRequestCancel200ResponseData instantiates a new RejectRequestCancel200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRejectRequestCancel200ResponseData() *RejectRequestCancel200ResponseData {
	this := RejectRequestCancel200ResponseData{}
	return &this
}

// NewRejectRequestCancel200ResponseDataWithDefaults instantiates a new RejectRequestCancel200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRejectRequestCancel200ResponseDataWithDefaults() *RejectRequestCancel200ResponseData {
	this := RejectRequestCancel200ResponseData{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *RejectRequestCancel200ResponseData) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RejectRequestCancel200ResponseData) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *RejectRequestCancel200ResponseData) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *RejectRequestCancel200ResponseData) SetSuccess(v bool) {
	o.Success = &v
}

func (o RejectRequestCancel200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RejectRequestCancel200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	return toSerialize, nil
}

type NullableRejectRequestCancel200ResponseData struct {
	value *RejectRequestCancel200ResponseData
	isSet bool
}

func (v NullableRejectRequestCancel200ResponseData) Get() *RejectRequestCancel200ResponseData {
	return v.value
}

func (v *NullableRejectRequestCancel200ResponseData) Set(val *RejectRequestCancel200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableRejectRequestCancel200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableRejectRequestCancel200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRejectRequestCancel200ResponseData(val *RejectRequestCancel200ResponseData) *NullableRejectRequestCancel200ResponseData {
	return &NullableRejectRequestCancel200ResponseData{value: val, isSet: true}
}

func (v NullableRejectRequestCancel200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRejectRequestCancel200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


