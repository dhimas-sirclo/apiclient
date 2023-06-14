/*
Tokopedia

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
Contact: dev@sirclo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tokopedia

import (
	"encoding/json"
)

// checks if the ProductPreOrder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductPreOrder{}

// ProductPreOrder struct for ProductPreOrder
type ProductPreOrder struct {
	Duration *bool `json:"duration,omitempty"`
	Unit *int64 `json:"unit,omitempty"`
	Day *int64 `json:"day,omitempty"`
}

// NewProductPreOrder instantiates a new ProductPreOrder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductPreOrder() *ProductPreOrder {
	this := ProductPreOrder{}
	return &this
}

// NewProductPreOrderWithDefaults instantiates a new ProductPreOrder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductPreOrderWithDefaults() *ProductPreOrder {
	this := ProductPreOrder{}
	return &this
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *ProductPreOrder) GetDuration() bool {
	if o == nil || IsNil(o.Duration) {
		var ret bool
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductPreOrder) GetDurationOk() (*bool, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *ProductPreOrder) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given bool and assigns it to the Duration field.
func (o *ProductPreOrder) SetDuration(v bool) {
	o.Duration = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *ProductPreOrder) GetUnit() int64 {
	if o == nil || IsNil(o.Unit) {
		var ret int64
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductPreOrder) GetUnitOk() (*int64, bool) {
	if o == nil || IsNil(o.Unit) {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *ProductPreOrder) HasUnit() bool {
	if o != nil && !IsNil(o.Unit) {
		return true
	}

	return false
}

// SetUnit gets a reference to the given int64 and assigns it to the Unit field.
func (o *ProductPreOrder) SetUnit(v int64) {
	o.Unit = &v
}

// GetDay returns the Day field value if set, zero value otherwise.
func (o *ProductPreOrder) GetDay() int64 {
	if o == nil || IsNil(o.Day) {
		var ret int64
		return ret
	}
	return *o.Day
}

// GetDayOk returns a tuple with the Day field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductPreOrder) GetDayOk() (*int64, bool) {
	if o == nil || IsNil(o.Day) {
		return nil, false
	}
	return o.Day, true
}

// HasDay returns a boolean if a field has been set.
func (o *ProductPreOrder) HasDay() bool {
	if o != nil && !IsNil(o.Day) {
		return true
	}

	return false
}

// SetDay gets a reference to the given int64 and assigns it to the Day field.
func (o *ProductPreOrder) SetDay(v int64) {
	o.Day = &v
}

func (o ProductPreOrder) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductPreOrder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !IsNil(o.Unit) {
		toSerialize["unit"] = o.Unit
	}
	if !IsNil(o.Day) {
		toSerialize["day"] = o.Day
	}
	return toSerialize, nil
}

type NullableProductPreOrder struct {
	value *ProductPreOrder
	isSet bool
}

func (v NullableProductPreOrder) Get() *ProductPreOrder {
	return v.value
}

func (v *NullableProductPreOrder) Set(val *ProductPreOrder) {
	v.value = val
	v.isSet = true
}

func (v NullableProductPreOrder) IsSet() bool {
	return v.isSet
}

func (v *NullableProductPreOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductPreOrder(val *ProductPreOrder) *NullableProductPreOrder {
	return &NullableProductPreOrder{value: val, isSet: true}
}

func (v NullableProductPreOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductPreOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


