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

// checks if the UpdateShipmentInfo200ResponseDataData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateShipmentInfo200ResponseDataData{}

// UpdateShipmentInfo200ResponseDataData struct for UpdateShipmentInfo200ResponseDataData
type UpdateShipmentInfo200ResponseDataData struct {
	// Is the request success?
	Success *bool `json:"success,omitempty"`
}

// NewUpdateShipmentInfo200ResponseDataData instantiates a new UpdateShipmentInfo200ResponseDataData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateShipmentInfo200ResponseDataData() *UpdateShipmentInfo200ResponseDataData {
	this := UpdateShipmentInfo200ResponseDataData{}
	return &this
}

// NewUpdateShipmentInfo200ResponseDataDataWithDefaults instantiates a new UpdateShipmentInfo200ResponseDataData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateShipmentInfo200ResponseDataDataWithDefaults() *UpdateShipmentInfo200ResponseDataData {
	this := UpdateShipmentInfo200ResponseDataData{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *UpdateShipmentInfo200ResponseDataData) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateShipmentInfo200ResponseDataData) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *UpdateShipmentInfo200ResponseDataData) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *UpdateShipmentInfo200ResponseDataData) SetSuccess(v bool) {
	o.Success = &v
}

func (o UpdateShipmentInfo200ResponseDataData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateShipmentInfo200ResponseDataData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	return toSerialize, nil
}

type NullableUpdateShipmentInfo200ResponseDataData struct {
	value *UpdateShipmentInfo200ResponseDataData
	isSet bool
}

func (v NullableUpdateShipmentInfo200ResponseDataData) Get() *UpdateShipmentInfo200ResponseDataData {
	return v.value
}

func (v *NullableUpdateShipmentInfo200ResponseDataData) Set(val *UpdateShipmentInfo200ResponseDataData) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateShipmentInfo200ResponseDataData) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateShipmentInfo200ResponseDataData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateShipmentInfo200ResponseDataData(val *UpdateShipmentInfo200ResponseDataData) *NullableUpdateShipmentInfo200ResponseDataData {
	return &NullableUpdateShipmentInfo200ResponseDataData{value: val, isSet: true}
}

func (v NullableUpdateShipmentInfo200ResponseDataData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateShipmentInfo200ResponseDataData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


