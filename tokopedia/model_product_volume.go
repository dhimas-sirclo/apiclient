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

// checks if the ProductVolume type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductVolume{}

// ProductVolume struct for ProductVolume
type ProductVolume struct {
	// Product Length
	Length *int32 `json:"length,omitempty"`
	// Product Width
	Width *int32 `json:"width,omitempty"`
	// Product Height
	Height *int32 `json:"height,omitempty"`
}

// NewProductVolume instantiates a new ProductVolume object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductVolume() *ProductVolume {
	this := ProductVolume{}
	return &this
}

// NewProductVolumeWithDefaults instantiates a new ProductVolume object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductVolumeWithDefaults() *ProductVolume {
	this := ProductVolume{}
	return &this
}

// GetLength returns the Length field value if set, zero value otherwise.
func (o *ProductVolume) GetLength() int32 {
	if o == nil || IsNil(o.Length) {
		var ret int32
		return ret
	}
	return *o.Length
}

// GetLengthOk returns a tuple with the Length field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductVolume) GetLengthOk() (*int32, bool) {
	if o == nil || IsNil(o.Length) {
		return nil, false
	}
	return o.Length, true
}

// HasLength returns a boolean if a field has been set.
func (o *ProductVolume) HasLength() bool {
	if o != nil && !IsNil(o.Length) {
		return true
	}

	return false
}

// SetLength gets a reference to the given int32 and assigns it to the Length field.
func (o *ProductVolume) SetLength(v int32) {
	o.Length = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *ProductVolume) GetWidth() int32 {
	if o == nil || IsNil(o.Width) {
		var ret int32
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductVolume) GetWidthOk() (*int32, bool) {
	if o == nil || IsNil(o.Width) {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *ProductVolume) HasWidth() bool {
	if o != nil && !IsNil(o.Width) {
		return true
	}

	return false
}

// SetWidth gets a reference to the given int32 and assigns it to the Width field.
func (o *ProductVolume) SetWidth(v int32) {
	o.Width = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *ProductVolume) GetHeight() int32 {
	if o == nil || IsNil(o.Height) {
		var ret int32
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductVolume) GetHeightOk() (*int32, bool) {
	if o == nil || IsNil(o.Height) {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *ProductVolume) HasHeight() bool {
	if o != nil && !IsNil(o.Height) {
		return true
	}

	return false
}

// SetHeight gets a reference to the given int32 and assigns it to the Height field.
func (o *ProductVolume) SetHeight(v int32) {
	o.Height = &v
}

func (o ProductVolume) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductVolume) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Length) {
		toSerialize["length"] = o.Length
	}
	if !IsNil(o.Width) {
		toSerialize["width"] = o.Width
	}
	if !IsNil(o.Height) {
		toSerialize["height"] = o.Height
	}
	return toSerialize, nil
}

type NullableProductVolume struct {
	value *ProductVolume
	isSet bool
}

func (v NullableProductVolume) Get() *ProductVolume {
	return v.value
}

func (v *NullableProductVolume) Set(val *ProductVolume) {
	v.value = val
	v.isSet = true
}

func (v NullableProductVolume) IsSet() bool {
	return v.isSet
}

func (v *NullableProductVolume) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductVolume(val *ProductVolume) *NullableProductVolume {
	return &NullableProductVolume{value: val, isSet: true}
}

func (v NullableProductVolume) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductVolume) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


