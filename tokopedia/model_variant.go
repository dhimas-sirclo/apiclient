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

// checks if the Variant type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Variant{}

// Variant struct for Variant
type Variant struct {
	// Variant Parent Unique Identifier
	ParentId *int32 `json:"parent_id,omitempty"`
	// Variant Default Child Identifier
	DefaultChild *int32 `json:"default_child,omitempty"`
	Variant []VariantVariant `json:"variant,omitempty"`
	Children []VariantChildren `json:"children,omitempty"`
}

// NewVariant instantiates a new Variant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVariant() *Variant {
	this := Variant{}
	return &this
}

// NewVariantWithDefaults instantiates a new Variant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVariantWithDefaults() *Variant {
	this := Variant{}
	return &this
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *Variant) GetParentId() int32 {
	if o == nil || IsNil(o.ParentId) {
		var ret int32
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Variant) GetParentIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ParentId) {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *Variant) HasParentId() bool {
	if o != nil && !IsNil(o.ParentId) {
		return true
	}

	return false
}

// SetParentId gets a reference to the given int32 and assigns it to the ParentId field.
func (o *Variant) SetParentId(v int32) {
	o.ParentId = &v
}

// GetDefaultChild returns the DefaultChild field value if set, zero value otherwise.
func (o *Variant) GetDefaultChild() int32 {
	if o == nil || IsNil(o.DefaultChild) {
		var ret int32
		return ret
	}
	return *o.DefaultChild
}

// GetDefaultChildOk returns a tuple with the DefaultChild field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Variant) GetDefaultChildOk() (*int32, bool) {
	if o == nil || IsNil(o.DefaultChild) {
		return nil, false
	}
	return o.DefaultChild, true
}

// HasDefaultChild returns a boolean if a field has been set.
func (o *Variant) HasDefaultChild() bool {
	if o != nil && !IsNil(o.DefaultChild) {
		return true
	}

	return false
}

// SetDefaultChild gets a reference to the given int32 and assigns it to the DefaultChild field.
func (o *Variant) SetDefaultChild(v int32) {
	o.DefaultChild = &v
}

// GetVariant returns the Variant field value if set, zero value otherwise.
func (o *Variant) GetVariant() []VariantVariant {
	if o == nil || IsNil(o.Variant) {
		var ret []VariantVariant
		return ret
	}
	return o.Variant
}

// GetVariantOk returns a tuple with the Variant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Variant) GetVariantOk() ([]VariantVariant, bool) {
	if o == nil || IsNil(o.Variant) {
		return nil, false
	}
	return o.Variant, true
}

// HasVariant returns a boolean if a field has been set.
func (o *Variant) HasVariant() bool {
	if o != nil && !IsNil(o.Variant) {
		return true
	}

	return false
}

// SetVariant gets a reference to the given []VariantVariant and assigns it to the Variant field.
func (o *Variant) SetVariant(v []VariantVariant) {
	o.Variant = v
}

// GetChildren returns the Children field value if set, zero value otherwise.
func (o *Variant) GetChildren() []VariantChildren {
	if o == nil || IsNil(o.Children) {
		var ret []VariantChildren
		return ret
	}
	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Variant) GetChildrenOk() ([]VariantChildren, bool) {
	if o == nil || IsNil(o.Children) {
		return nil, false
	}
	return o.Children, true
}

// HasChildren returns a boolean if a field has been set.
func (o *Variant) HasChildren() bool {
	if o != nil && !IsNil(o.Children) {
		return true
	}

	return false
}

// SetChildren gets a reference to the given []VariantChildren and assigns it to the Children field.
func (o *Variant) SetChildren(v []VariantChildren) {
	o.Children = v
}

func (o Variant) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Variant) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ParentId) {
		toSerialize["parent_id"] = o.ParentId
	}
	if !IsNil(o.DefaultChild) {
		toSerialize["default_child"] = o.DefaultChild
	}
	if !IsNil(o.Variant) {
		toSerialize["variant"] = o.Variant
	}
	if !IsNil(o.Children) {
		toSerialize["children"] = o.Children
	}
	return toSerialize, nil
}

type NullableVariant struct {
	value *Variant
	isSet bool
}

func (v NullableVariant) Get() *Variant {
	return v.value
}

func (v *NullableVariant) Set(val *Variant) {
	v.value = val
	v.isSet = true
}

func (v NullableVariant) IsSet() bool {
	return v.isSet
}

func (v *NullableVariant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVariant(val *Variant) *NullableVariant {
	return &NullableVariant{value: val, isSet: true}
}

func (v NullableVariant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVariant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

