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

// checks if the EditProductV2RequestProductsInnerPicturesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditProductV2RequestProductsInnerPicturesInner{}

// EditProductV2RequestProductsInnerPicturesInner struct for EditProductV2RequestProductsInnerPicturesInner
type EditProductV2RequestProductsInnerPicturesInner struct {
	FilePath *string `json:"file_path,omitempty"`
}

// NewEditProductV2RequestProductsInnerPicturesInner instantiates a new EditProductV2RequestProductsInnerPicturesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditProductV2RequestProductsInnerPicturesInner() *EditProductV2RequestProductsInnerPicturesInner {
	this := EditProductV2RequestProductsInnerPicturesInner{}
	return &this
}

// NewEditProductV2RequestProductsInnerPicturesInnerWithDefaults instantiates a new EditProductV2RequestProductsInnerPicturesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditProductV2RequestProductsInnerPicturesInnerWithDefaults() *EditProductV2RequestProductsInnerPicturesInner {
	this := EditProductV2RequestProductsInnerPicturesInner{}
	return &this
}

// GetFilePath returns the FilePath field value if set, zero value otherwise.
func (o *EditProductV2RequestProductsInnerPicturesInner) GetFilePath() string {
	if o == nil || IsNil(o.FilePath) {
		var ret string
		return ret
	}
	return *o.FilePath
}

// GetFilePathOk returns a tuple with the FilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditProductV2RequestProductsInnerPicturesInner) GetFilePathOk() (*string, bool) {
	if o == nil || IsNil(o.FilePath) {
		return nil, false
	}
	return o.FilePath, true
}

// HasFilePath returns a boolean if a field has been set.
func (o *EditProductV2RequestProductsInnerPicturesInner) HasFilePath() bool {
	if o != nil && !IsNil(o.FilePath) {
		return true
	}

	return false
}

// SetFilePath gets a reference to the given string and assigns it to the FilePath field.
func (o *EditProductV2RequestProductsInnerPicturesInner) SetFilePath(v string) {
	o.FilePath = &v
}

func (o EditProductV2RequestProductsInnerPicturesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditProductV2RequestProductsInnerPicturesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FilePath) {
		toSerialize["file_path"] = o.FilePath
	}
	return toSerialize, nil
}

type NullableEditProductV2RequestProductsInnerPicturesInner struct {
	value *EditProductV2RequestProductsInnerPicturesInner
	isSet bool
}

func (v NullableEditProductV2RequestProductsInnerPicturesInner) Get() *EditProductV2RequestProductsInnerPicturesInner {
	return v.value
}

func (v *NullableEditProductV2RequestProductsInnerPicturesInner) Set(val *EditProductV2RequestProductsInnerPicturesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableEditProductV2RequestProductsInnerPicturesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableEditProductV2RequestProductsInnerPicturesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditProductV2RequestProductsInnerPicturesInner(val *EditProductV2RequestProductsInnerPicturesInner) *NullableEditProductV2RequestProductsInnerPicturesInner {
	return &NullableEditProductV2RequestProductsInnerPicturesInner{value: val, isSet: true}
}

func (v NullableEditProductV2RequestProductsInnerPicturesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditProductV2RequestProductsInnerPicturesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


