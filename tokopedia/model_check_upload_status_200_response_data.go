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

// checks if the CheckUploadStatus200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CheckUploadStatus200ResponseData{}

// CheckUploadStatus200ResponseData struct for CheckUploadStatus200ResponseData
type CheckUploadStatus200ResponseData struct {
	UploadData []CheckUploadStatus200ResponseDataUploadDataInner `json:"upload_data,omitempty"`
}

// NewCheckUploadStatus200ResponseData instantiates a new CheckUploadStatus200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckUploadStatus200ResponseData() *CheckUploadStatus200ResponseData {
	this := CheckUploadStatus200ResponseData{}
	return &this
}

// NewCheckUploadStatus200ResponseDataWithDefaults instantiates a new CheckUploadStatus200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckUploadStatus200ResponseDataWithDefaults() *CheckUploadStatus200ResponseData {
	this := CheckUploadStatus200ResponseData{}
	return &this
}

// GetUploadData returns the UploadData field value if set, zero value otherwise.
func (o *CheckUploadStatus200ResponseData) GetUploadData() []CheckUploadStatus200ResponseDataUploadDataInner {
	if o == nil || IsNil(o.UploadData) {
		var ret []CheckUploadStatus200ResponseDataUploadDataInner
		return ret
	}
	return o.UploadData
}

// GetUploadDataOk returns a tuple with the UploadData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckUploadStatus200ResponseData) GetUploadDataOk() ([]CheckUploadStatus200ResponseDataUploadDataInner, bool) {
	if o == nil || IsNil(o.UploadData) {
		return nil, false
	}
	return o.UploadData, true
}

// HasUploadData returns a boolean if a field has been set.
func (o *CheckUploadStatus200ResponseData) HasUploadData() bool {
	if o != nil && !IsNil(o.UploadData) {
		return true
	}

	return false
}

// SetUploadData gets a reference to the given []CheckUploadStatus200ResponseDataUploadDataInner and assigns it to the UploadData field.
func (o *CheckUploadStatus200ResponseData) SetUploadData(v []CheckUploadStatus200ResponseDataUploadDataInner) {
	o.UploadData = v
}

func (o CheckUploadStatus200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckUploadStatus200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UploadData) {
		toSerialize["upload_data"] = o.UploadData
	}
	return toSerialize, nil
}

type NullableCheckUploadStatus200ResponseData struct {
	value *CheckUploadStatus200ResponseData
	isSet bool
}

func (v NullableCheckUploadStatus200ResponseData) Get() *CheckUploadStatus200ResponseData {
	return v.value
}

func (v *NullableCheckUploadStatus200ResponseData) Set(val *CheckUploadStatus200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckUploadStatus200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckUploadStatus200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckUploadStatus200ResponseData(val *CheckUploadStatus200ResponseData) *NullableCheckUploadStatus200ResponseData {
	return &NullableCheckUploadStatus200ResponseData{value: val, isSet: true}
}

func (v NullableCheckUploadStatus200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckUploadStatus200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

