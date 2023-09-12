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

// checks if the GetVariantResponseAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetVariantResponseAllOf{}

// GetVariantResponseAllOf struct for GetVariantResponseAllOf
type GetVariantResponseAllOf struct {
	Data *Variant `json:"data,omitempty"`
}

// NewGetVariantResponseAllOf instantiates a new GetVariantResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetVariantResponseAllOf() *GetVariantResponseAllOf {
	this := GetVariantResponseAllOf{}
	return &this
}

// NewGetVariantResponseAllOfWithDefaults instantiates a new GetVariantResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetVariantResponseAllOfWithDefaults() *GetVariantResponseAllOf {
	this := GetVariantResponseAllOf{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetVariantResponseAllOf) GetData() Variant {
	if o == nil || IsNil(o.Data) {
		var ret Variant
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVariantResponseAllOf) GetDataOk() (*Variant, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetVariantResponseAllOf) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given Variant and assigns it to the Data field.
func (o *GetVariantResponseAllOf) SetData(v Variant) {
	o.Data = &v
}

func (o GetVariantResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetVariantResponseAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGetVariantResponseAllOf struct {
	value *GetVariantResponseAllOf
	isSet bool
}

func (v NullableGetVariantResponseAllOf) Get() *GetVariantResponseAllOf {
	return v.value
}

func (v *NullableGetVariantResponseAllOf) Set(val *GetVariantResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableGetVariantResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableGetVariantResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetVariantResponseAllOf(val *GetVariantResponseAllOf) *NullableGetVariantResponseAllOf {
	return &NullableGetVariantResponseAllOf{value: val, isSet: true}
}

func (v NullableGetVariantResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetVariantResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

