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

// checks if the EditProductV2RequestProductsInnerVariantProductsInnerPicturesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditProductV2RequestProductsInnerVariantProductsInnerPicturesInner{}

// EditProductV2RequestProductsInnerVariantProductsInnerPicturesInner struct for EditProductV2RequestProductsInnerVariantProductsInnerPicturesInner
type EditProductV2RequestProductsInnerVariantProductsInnerPicturesInner struct {
	FilePath string `json:"file_path"`
}

// NewEditProductV2RequestProductsInnerVariantProductsInnerPicturesInner instantiates a new EditProductV2RequestProductsInnerVariantProductsInnerPicturesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditProductV2RequestProductsInnerVariantProductsInnerPicturesInner(filePath string) *EditProductV2RequestProductsInnerVariantProductsInnerPicturesInner {
	this := EditProductV2RequestProductsInnerVariantProductsInnerPicturesInner{}
	this.FilePath = filePath
	return &this
}

// NewEditProductV2RequestProductsInnerVariantProductsInnerPicturesInnerWithDefaults instantiates a new EditProductV2RequestProductsInnerVariantProductsInnerPicturesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditProductV2RequestProductsInnerVariantProductsInnerPicturesInnerWithDefaults() *EditProductV2RequestProductsInnerVariantProductsInnerPicturesInner {
	this := EditProductV2RequestProductsInnerVariantProductsInnerPicturesInner{}
	return &this
}

// GetFilePath returns the FilePath field value
func (o *EditProductV2RequestProductsInnerVariantProductsInnerPicturesInner) GetFilePath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FilePath
}

// GetFilePathOk returns a tuple with the FilePath field value
// and a boolean to check if the value has been set.
func (o *EditProductV2RequestProductsInnerVariantProductsInnerPicturesInner) GetFilePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FilePath, true
}

// SetFilePath sets field value
func (o *EditProductV2RequestProductsInnerVariantProductsInnerPicturesInner) SetFilePath(v string) {
	o.FilePath = v
}

func (o EditProductV2RequestProductsInnerVariantProductsInnerPicturesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditProductV2RequestProductsInnerVariantProductsInnerPicturesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["file_path"] = o.FilePath
	return toSerialize, nil
}

type NullableEditProductV2RequestProductsInnerVariantProductsInnerPicturesInner struct {
	value *EditProductV2RequestProductsInnerVariantProductsInnerPicturesInner
	isSet bool
}

func (v NullableEditProductV2RequestProductsInnerVariantProductsInnerPicturesInner) Get() *EditProductV2RequestProductsInnerVariantProductsInnerPicturesInner {
	return v.value
}

func (v *NullableEditProductV2RequestProductsInnerVariantProductsInnerPicturesInner) Set(val *EditProductV2RequestProductsInnerVariantProductsInnerPicturesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableEditProductV2RequestProductsInnerVariantProductsInnerPicturesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableEditProductV2RequestProductsInnerVariantProductsInnerPicturesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditProductV2RequestProductsInnerVariantProductsInnerPicturesInner(val *EditProductV2RequestProductsInnerVariantProductsInnerPicturesInner) *NullableEditProductV2RequestProductsInnerVariantProductsInnerPicturesInner {
	return &NullableEditProductV2RequestProductsInnerVariantProductsInnerPicturesInner{value: val, isSet: true}
}

func (v NullableEditProductV2RequestProductsInnerVariantProductsInnerPicturesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditProductV2RequestProductsInnerVariantProductsInnerPicturesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


