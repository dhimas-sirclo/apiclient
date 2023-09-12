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

// checks if the CategoryVariants type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CategoryVariants{}

// CategoryVariants struct for CategoryVariants
type CategoryVariants struct {
	Header *Header `json:"header,omitempty"`
	Data *CategoryVariant `json:"data,omitempty"`
}

// NewCategoryVariants instantiates a new CategoryVariants object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCategoryVariants() *CategoryVariants {
	this := CategoryVariants{}
	return &this
}

// NewCategoryVariantsWithDefaults instantiates a new CategoryVariants object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCategoryVariantsWithDefaults() *CategoryVariants {
	this := CategoryVariants{}
	return &this
}

// GetHeader returns the Header field value if set, zero value otherwise.
func (o *CategoryVariants) GetHeader() Header {
	if o == nil || IsNil(o.Header) {
		var ret Header
		return ret
	}
	return *o.Header
}

// GetHeaderOk returns a tuple with the Header field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryVariants) GetHeaderOk() (*Header, bool) {
	if o == nil || IsNil(o.Header) {
		return nil, false
	}
	return o.Header, true
}

// HasHeader returns a boolean if a field has been set.
func (o *CategoryVariants) HasHeader() bool {
	if o != nil && !IsNil(o.Header) {
		return true
	}

	return false
}

// SetHeader gets a reference to the given Header and assigns it to the Header field.
func (o *CategoryVariants) SetHeader(v Header) {
	o.Header = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CategoryVariants) GetData() CategoryVariant {
	if o == nil || IsNil(o.Data) {
		var ret CategoryVariant
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryVariants) GetDataOk() (*CategoryVariant, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CategoryVariants) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given CategoryVariant and assigns it to the Data field.
func (o *CategoryVariants) SetData(v CategoryVariant) {
	o.Data = &v
}

func (o CategoryVariants) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CategoryVariants) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Header) {
		toSerialize["header"] = o.Header
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableCategoryVariants struct {
	value *CategoryVariants
	isSet bool
}

func (v NullableCategoryVariants) Get() *CategoryVariants {
	return v.value
}

func (v *NullableCategoryVariants) Set(val *CategoryVariants) {
	v.value = val
	v.isSet = true
}

func (v NullableCategoryVariants) IsSet() bool {
	return v.isSet
}

func (v *NullableCategoryVariants) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCategoryVariants(val *CategoryVariants) *NullableCategoryVariants {
	return &NullableCategoryVariants{value: val, isSet: true}
}

func (v NullableCategoryVariants) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCategoryVariants) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

