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

// checks if the EditProductV3RequestProductsInnerPreorder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditProductV3RequestProductsInnerPreorder{}

// EditProductV3RequestProductsInnerPreorder Preorder information. The object keys includes: is_active, duration, and time_unit. Remove preorder from related product by passing empty object ({}) in this request 
type EditProductV3RequestProductsInnerPreorder struct {
	IsActive *bool `json:"is_active,omitempty"`
	Duration *int64 `json:"duration,omitempty"`
	TimeUnit *string `json:"time_unit,omitempty"`
}

// NewEditProductV3RequestProductsInnerPreorder instantiates a new EditProductV3RequestProductsInnerPreorder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditProductV3RequestProductsInnerPreorder() *EditProductV3RequestProductsInnerPreorder {
	this := EditProductV3RequestProductsInnerPreorder{}
	var timeUnit string = "DAY"
	this.TimeUnit = &timeUnit
	return &this
}

// NewEditProductV3RequestProductsInnerPreorderWithDefaults instantiates a new EditProductV3RequestProductsInnerPreorder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditProductV3RequestProductsInnerPreorderWithDefaults() *EditProductV3RequestProductsInnerPreorder {
	this := EditProductV3RequestProductsInnerPreorder{}
	var timeUnit string = "DAY"
	this.TimeUnit = &timeUnit
	return &this
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *EditProductV3RequestProductsInnerPreorder) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditProductV3RequestProductsInnerPreorder) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *EditProductV3RequestProductsInnerPreorder) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *EditProductV3RequestProductsInnerPreorder) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *EditProductV3RequestProductsInnerPreorder) GetDuration() int64 {
	if o == nil || IsNil(o.Duration) {
		var ret int64
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditProductV3RequestProductsInnerPreorder) GetDurationOk() (*int64, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *EditProductV3RequestProductsInnerPreorder) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int64 and assigns it to the Duration field.
func (o *EditProductV3RequestProductsInnerPreorder) SetDuration(v int64) {
	o.Duration = &v
}

// GetTimeUnit returns the TimeUnit field value if set, zero value otherwise.
func (o *EditProductV3RequestProductsInnerPreorder) GetTimeUnit() string {
	if o == nil || IsNil(o.TimeUnit) {
		var ret string
		return ret
	}
	return *o.TimeUnit
}

// GetTimeUnitOk returns a tuple with the TimeUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditProductV3RequestProductsInnerPreorder) GetTimeUnitOk() (*string, bool) {
	if o == nil || IsNil(o.TimeUnit) {
		return nil, false
	}
	return o.TimeUnit, true
}

// HasTimeUnit returns a boolean if a field has been set.
func (o *EditProductV3RequestProductsInnerPreorder) HasTimeUnit() bool {
	if o != nil && !IsNil(o.TimeUnit) {
		return true
	}

	return false
}

// SetTimeUnit gets a reference to the given string and assigns it to the TimeUnit field.
func (o *EditProductV3RequestProductsInnerPreorder) SetTimeUnit(v string) {
	o.TimeUnit = &v
}

func (o EditProductV3RequestProductsInnerPreorder) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditProductV3RequestProductsInnerPreorder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsActive) {
		toSerialize["is_active"] = o.IsActive
	}
	if !IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !IsNil(o.TimeUnit) {
		toSerialize["time_unit"] = o.TimeUnit
	}
	return toSerialize, nil
}

type NullableEditProductV3RequestProductsInnerPreorder struct {
	value *EditProductV3RequestProductsInnerPreorder
	isSet bool
}

func (v NullableEditProductV3RequestProductsInnerPreorder) Get() *EditProductV3RequestProductsInnerPreorder {
	return v.value
}

func (v *NullableEditProductV3RequestProductsInnerPreorder) Set(val *EditProductV3RequestProductsInnerPreorder) {
	v.value = val
	v.isSet = true
}

func (v NullableEditProductV3RequestProductsInnerPreorder) IsSet() bool {
	return v.isSet
}

func (v *NullableEditProductV3RequestProductsInnerPreorder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditProductV3RequestProductsInnerPreorder(val *EditProductV3RequestProductsInnerPreorder) *NullableEditProductV3RequestProductsInnerPreorder {
	return &NullableEditProductV3RequestProductsInnerPreorder{value: val, isSet: true}
}

func (v NullableEditProductV3RequestProductsInnerPreorder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditProductV3RequestProductsInnerPreorder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

