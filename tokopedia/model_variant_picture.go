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

// checks if the VariantPicture type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VariantPicture{}

// VariantPicture struct for VariantPicture
type VariantPicture struct {
	Original *string `json:"original,omitempty"`
	Thumbnail *string `json:"thumbnail,omitempty"`
}

// NewVariantPicture instantiates a new VariantPicture object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVariantPicture() *VariantPicture {
	this := VariantPicture{}
	return &this
}

// NewVariantPictureWithDefaults instantiates a new VariantPicture object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVariantPictureWithDefaults() *VariantPicture {
	this := VariantPicture{}
	return &this
}

// GetOriginal returns the Original field value if set, zero value otherwise.
func (o *VariantPicture) GetOriginal() string {
	if o == nil || IsNil(o.Original) {
		var ret string
		return ret
	}
	return *o.Original
}

// GetOriginalOk returns a tuple with the Original field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariantPicture) GetOriginalOk() (*string, bool) {
	if o == nil || IsNil(o.Original) {
		return nil, false
	}
	return o.Original, true
}

// HasOriginal returns a boolean if a field has been set.
func (o *VariantPicture) HasOriginal() bool {
	if o != nil && !IsNil(o.Original) {
		return true
	}

	return false
}

// SetOriginal gets a reference to the given string and assigns it to the Original field.
func (o *VariantPicture) SetOriginal(v string) {
	o.Original = &v
}

// GetThumbnail returns the Thumbnail field value if set, zero value otherwise.
func (o *VariantPicture) GetThumbnail() string {
	if o == nil || IsNil(o.Thumbnail) {
		var ret string
		return ret
	}
	return *o.Thumbnail
}

// GetThumbnailOk returns a tuple with the Thumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariantPicture) GetThumbnailOk() (*string, bool) {
	if o == nil || IsNil(o.Thumbnail) {
		return nil, false
	}
	return o.Thumbnail, true
}

// HasThumbnail returns a boolean if a field has been set.
func (o *VariantPicture) HasThumbnail() bool {
	if o != nil && !IsNil(o.Thumbnail) {
		return true
	}

	return false
}

// SetThumbnail gets a reference to the given string and assigns it to the Thumbnail field.
func (o *VariantPicture) SetThumbnail(v string) {
	o.Thumbnail = &v
}

func (o VariantPicture) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VariantPicture) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Original) {
		toSerialize["original"] = o.Original
	}
	if !IsNil(o.Thumbnail) {
		toSerialize["thumbnail"] = o.Thumbnail
	}
	return toSerialize, nil
}

type NullableVariantPicture struct {
	value *VariantPicture
	isSet bool
}

func (v NullableVariantPicture) Get() *VariantPicture {
	return v.value
}

func (v *NullableVariantPicture) Set(val *VariantPicture) {
	v.value = val
	v.isSet = true
}

func (v NullableVariantPicture) IsSet() bool {
	return v.isSet
}

func (v *NullableVariantPicture) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVariantPicture(val *VariantPicture) *NullableVariantPicture {
	return &NullableVariantPicture{value: val, isSet: true}
}

func (v NullableVariantPicture) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVariantPicture) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


