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

// checks if the EditProductV2200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditProductV2200ResponseData{}

// EditProductV2200ResponseData struct for EditProductV2200ResponseData
type EditProductV2200ResponseData struct {
	UploadId *int64 `json:"upload_id,omitempty"`
}

// NewEditProductV2200ResponseData instantiates a new EditProductV2200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditProductV2200ResponseData() *EditProductV2200ResponseData {
	this := EditProductV2200ResponseData{}
	return &this
}

// NewEditProductV2200ResponseDataWithDefaults instantiates a new EditProductV2200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditProductV2200ResponseDataWithDefaults() *EditProductV2200ResponseData {
	this := EditProductV2200ResponseData{}
	return &this
}

// GetUploadId returns the UploadId field value if set, zero value otherwise.
func (o *EditProductV2200ResponseData) GetUploadId() int64 {
	if o == nil || IsNil(o.UploadId) {
		var ret int64
		return ret
	}
	return *o.UploadId
}

// GetUploadIdOk returns a tuple with the UploadId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditProductV2200ResponseData) GetUploadIdOk() (*int64, bool) {
	if o == nil || IsNil(o.UploadId) {
		return nil, false
	}
	return o.UploadId, true
}

// HasUploadId returns a boolean if a field has been set.
func (o *EditProductV2200ResponseData) HasUploadId() bool {
	if o != nil && !IsNil(o.UploadId) {
		return true
	}

	return false
}

// SetUploadId gets a reference to the given int64 and assigns it to the UploadId field.
func (o *EditProductV2200ResponseData) SetUploadId(v int64) {
	o.UploadId = &v
}

func (o EditProductV2200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditProductV2200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UploadId) {
		toSerialize["upload_id"] = o.UploadId
	}
	return toSerialize, nil
}

type NullableEditProductV2200ResponseData struct {
	value *EditProductV2200ResponseData
	isSet bool
}

func (v NullableEditProductV2200ResponseData) Get() *EditProductV2200ResponseData {
	return v.value
}

func (v *NullableEditProductV2200ResponseData) Set(val *EditProductV2200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableEditProductV2200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableEditProductV2200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditProductV2200ResponseData(val *EditProductV2200ResponseData) *NullableEditProductV2200ResponseData {
	return &NullableEditProductV2200ResponseData{value: val, isSet: true}
}

func (v NullableEditProductV2200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditProductV2200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


