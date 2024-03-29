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

// checks if the EditProductV2RequestProductsInnerPreorder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditProductV2RequestProductsInnerPreorder{}

// EditProductV2RequestProductsInnerPreorder Preorder information. The object keys includes: is_active, duration, and time_unit. Remove preorder from related product by passing empty object ({}) in this request
type EditProductV2RequestProductsInnerPreorder struct {
	IsActive bool `json:"is_active"`
	Duration int64 `json:"duration"`
	TimeUnit string `json:"time_unit"`
}

// NewEditProductV2RequestProductsInnerPreorder instantiates a new EditProductV2RequestProductsInnerPreorder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditProductV2RequestProductsInnerPreorder(isActive bool, duration int64, timeUnit string) *EditProductV2RequestProductsInnerPreorder {
	this := EditProductV2RequestProductsInnerPreorder{}
	this.IsActive = isActive
	this.Duration = duration
	this.TimeUnit = timeUnit
	return &this
}

// NewEditProductV2RequestProductsInnerPreorderWithDefaults instantiates a new EditProductV2RequestProductsInnerPreorder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditProductV2RequestProductsInnerPreorderWithDefaults() *EditProductV2RequestProductsInnerPreorder {
	this := EditProductV2RequestProductsInnerPreorder{}
	return &this
}

// GetIsActive returns the IsActive field value
func (o *EditProductV2RequestProductsInnerPreorder) GetIsActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value
// and a boolean to check if the value has been set.
func (o *EditProductV2RequestProductsInnerPreorder) GetIsActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsActive, true
}

// SetIsActive sets field value
func (o *EditProductV2RequestProductsInnerPreorder) SetIsActive(v bool) {
	o.IsActive = v
}

// GetDuration returns the Duration field value
func (o *EditProductV2RequestProductsInnerPreorder) GetDuration() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *EditProductV2RequestProductsInnerPreorder) GetDurationOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *EditProductV2RequestProductsInnerPreorder) SetDuration(v int64) {
	o.Duration = v
}

// GetTimeUnit returns the TimeUnit field value
func (o *EditProductV2RequestProductsInnerPreorder) GetTimeUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TimeUnit
}

// GetTimeUnitOk returns a tuple with the TimeUnit field value
// and a boolean to check if the value has been set.
func (o *EditProductV2RequestProductsInnerPreorder) GetTimeUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeUnit, true
}

// SetTimeUnit sets field value
func (o *EditProductV2RequestProductsInnerPreorder) SetTimeUnit(v string) {
	o.TimeUnit = v
}

func (o EditProductV2RequestProductsInnerPreorder) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditProductV2RequestProductsInnerPreorder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["is_active"] = o.IsActive
	toSerialize["duration"] = o.Duration
	toSerialize["time_unit"] = o.TimeUnit
	return toSerialize, nil
}

type NullableEditProductV2RequestProductsInnerPreorder struct {
	value *EditProductV2RequestProductsInnerPreorder
	isSet bool
}

func (v NullableEditProductV2RequestProductsInnerPreorder) Get() *EditProductV2RequestProductsInnerPreorder {
	return v.value
}

func (v *NullableEditProductV2RequestProductsInnerPreorder) Set(val *EditProductV2RequestProductsInnerPreorder) {
	v.value = val
	v.isSet = true
}

func (v NullableEditProductV2RequestProductsInnerPreorder) IsSet() bool {
	return v.isSet
}

func (v *NullableEditProductV2RequestProductsInnerPreorder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditProductV2RequestProductsInnerPreorder(val *EditProductV2RequestProductsInnerPreorder) *NullableEditProductV2RequestProductsInnerPreorder {
	return &NullableEditProductV2RequestProductsInnerPreorder{value: val, isSet: true}
}

func (v NullableEditProductV2RequestProductsInnerPreorder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditProductV2RequestProductsInnerPreorder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


