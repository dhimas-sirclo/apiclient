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

// checks if the VariantUnit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VariantUnit{}

// VariantUnit struct for VariantUnit
type VariantUnit struct {
	VariantUnitId *int64 `json:"variant_unit_id,omitempty"`
	Status *int64 `json:"status,omitempty"`
	UnitName *string `json:"unit_name,omitempty"`
	UnitShortName *string `json:"unit_short_name,omitempty"`
	UnitValues []VariantUnitValue `json:"unit_values,omitempty"`
}

// NewVariantUnit instantiates a new VariantUnit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVariantUnit() *VariantUnit {
	this := VariantUnit{}
	return &this
}

// NewVariantUnitWithDefaults instantiates a new VariantUnit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVariantUnitWithDefaults() *VariantUnit {
	this := VariantUnit{}
	return &this
}

// GetVariantUnitId returns the VariantUnitId field value if set, zero value otherwise.
func (o *VariantUnit) GetVariantUnitId() int64 {
	if o == nil || IsNil(o.VariantUnitId) {
		var ret int64
		return ret
	}
	return *o.VariantUnitId
}

// GetVariantUnitIdOk returns a tuple with the VariantUnitId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariantUnit) GetVariantUnitIdOk() (*int64, bool) {
	if o == nil || IsNil(o.VariantUnitId) {
		return nil, false
	}
	return o.VariantUnitId, true
}

// HasVariantUnitId returns a boolean if a field has been set.
func (o *VariantUnit) HasVariantUnitId() bool {
	if o != nil && !IsNil(o.VariantUnitId) {
		return true
	}

	return false
}

// SetVariantUnitId gets a reference to the given int64 and assigns it to the VariantUnitId field.
func (o *VariantUnit) SetVariantUnitId(v int64) {
	o.VariantUnitId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *VariantUnit) GetStatus() int64 {
	if o == nil || IsNil(o.Status) {
		var ret int64
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariantUnit) GetStatusOk() (*int64, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *VariantUnit) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int64 and assigns it to the Status field.
func (o *VariantUnit) SetStatus(v int64) {
	o.Status = &v
}

// GetUnitName returns the UnitName field value if set, zero value otherwise.
func (o *VariantUnit) GetUnitName() string {
	if o == nil || IsNil(o.UnitName) {
		var ret string
		return ret
	}
	return *o.UnitName
}

// GetUnitNameOk returns a tuple with the UnitName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariantUnit) GetUnitNameOk() (*string, bool) {
	if o == nil || IsNil(o.UnitName) {
		return nil, false
	}
	return o.UnitName, true
}

// HasUnitName returns a boolean if a field has been set.
func (o *VariantUnit) HasUnitName() bool {
	if o != nil && !IsNil(o.UnitName) {
		return true
	}

	return false
}

// SetUnitName gets a reference to the given string and assigns it to the UnitName field.
func (o *VariantUnit) SetUnitName(v string) {
	o.UnitName = &v
}

// GetUnitShortName returns the UnitShortName field value if set, zero value otherwise.
func (o *VariantUnit) GetUnitShortName() string {
	if o == nil || IsNil(o.UnitShortName) {
		var ret string
		return ret
	}
	return *o.UnitShortName
}

// GetUnitShortNameOk returns a tuple with the UnitShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariantUnit) GetUnitShortNameOk() (*string, bool) {
	if o == nil || IsNil(o.UnitShortName) {
		return nil, false
	}
	return o.UnitShortName, true
}

// HasUnitShortName returns a boolean if a field has been set.
func (o *VariantUnit) HasUnitShortName() bool {
	if o != nil && !IsNil(o.UnitShortName) {
		return true
	}

	return false
}

// SetUnitShortName gets a reference to the given string and assigns it to the UnitShortName field.
func (o *VariantUnit) SetUnitShortName(v string) {
	o.UnitShortName = &v
}

// GetUnitValues returns the UnitValues field value if set, zero value otherwise.
func (o *VariantUnit) GetUnitValues() []VariantUnitValue {
	if o == nil || IsNil(o.UnitValues) {
		var ret []VariantUnitValue
		return ret
	}
	return o.UnitValues
}

// GetUnitValuesOk returns a tuple with the UnitValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariantUnit) GetUnitValuesOk() ([]VariantUnitValue, bool) {
	if o == nil || IsNil(o.UnitValues) {
		return nil, false
	}
	return o.UnitValues, true
}

// HasUnitValues returns a boolean if a field has been set.
func (o *VariantUnit) HasUnitValues() bool {
	if o != nil && !IsNil(o.UnitValues) {
		return true
	}

	return false
}

// SetUnitValues gets a reference to the given []VariantUnitValue and assigns it to the UnitValues field.
func (o *VariantUnit) SetUnitValues(v []VariantUnitValue) {
	o.UnitValues = v
}

func (o VariantUnit) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VariantUnit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VariantUnitId) {
		toSerialize["variant_unit_id"] = o.VariantUnitId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.UnitName) {
		toSerialize["unit_name"] = o.UnitName
	}
	if !IsNil(o.UnitShortName) {
		toSerialize["unit_short_name"] = o.UnitShortName
	}
	if !IsNil(o.UnitValues) {
		toSerialize["unit_values"] = o.UnitValues
	}
	return toSerialize, nil
}

type NullableVariantUnit struct {
	value *VariantUnit
	isSet bool
}

func (v NullableVariantUnit) Get() *VariantUnit {
	return v.value
}

func (v *NullableVariantUnit) Set(val *VariantUnit) {
	v.value = val
	v.isSet = true
}

func (v NullableVariantUnit) IsSet() bool {
	return v.isSet
}

func (v *NullableVariantUnit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVariantUnit(val *VariantUnit) *NullableVariantUnit {
	return &NullableVariantUnit{value: val, isSet: true}
}

func (v NullableVariantUnit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVariantUnit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

