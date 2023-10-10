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

// checks if the EditProductV3DefaultResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditProductV3DefaultResponseData{}

// EditProductV3DefaultResponseData struct for EditProductV3DefaultResponseData
type EditProductV3DefaultResponseData struct {
	TotalData *int64 `json:"total_data,omitempty"`
	SuccessData *int64 `json:"success_data,omitempty"`
	FailData *int64 `json:"fail_data,omitempty"`
	FailedRowsData []EditProductV3DefaultResponseDataFailedRowsDataInner `json:"failed_rows_data,omitempty"`
}

// NewEditProductV3DefaultResponseData instantiates a new EditProductV3DefaultResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditProductV3DefaultResponseData() *EditProductV3DefaultResponseData {
	this := EditProductV3DefaultResponseData{}
	return &this
}

// NewEditProductV3DefaultResponseDataWithDefaults instantiates a new EditProductV3DefaultResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditProductV3DefaultResponseDataWithDefaults() *EditProductV3DefaultResponseData {
	this := EditProductV3DefaultResponseData{}
	return &this
}

// GetTotalData returns the TotalData field value if set, zero value otherwise.
func (o *EditProductV3DefaultResponseData) GetTotalData() int64 {
	if o == nil || IsNil(o.TotalData) {
		var ret int64
		return ret
	}
	return *o.TotalData
}

// GetTotalDataOk returns a tuple with the TotalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditProductV3DefaultResponseData) GetTotalDataOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalData) {
		return nil, false
	}
	return o.TotalData, true
}

// HasTotalData returns a boolean if a field has been set.
func (o *EditProductV3DefaultResponseData) HasTotalData() bool {
	if o != nil && !IsNil(o.TotalData) {
		return true
	}

	return false
}

// SetTotalData gets a reference to the given int64 and assigns it to the TotalData field.
func (o *EditProductV3DefaultResponseData) SetTotalData(v int64) {
	o.TotalData = &v
}

// GetSuccessData returns the SuccessData field value if set, zero value otherwise.
func (o *EditProductV3DefaultResponseData) GetSuccessData() int64 {
	if o == nil || IsNil(o.SuccessData) {
		var ret int64
		return ret
	}
	return *o.SuccessData
}

// GetSuccessDataOk returns a tuple with the SuccessData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditProductV3DefaultResponseData) GetSuccessDataOk() (*int64, bool) {
	if o == nil || IsNil(o.SuccessData) {
		return nil, false
	}
	return o.SuccessData, true
}

// HasSuccessData returns a boolean if a field has been set.
func (o *EditProductV3DefaultResponseData) HasSuccessData() bool {
	if o != nil && !IsNil(o.SuccessData) {
		return true
	}

	return false
}

// SetSuccessData gets a reference to the given int64 and assigns it to the SuccessData field.
func (o *EditProductV3DefaultResponseData) SetSuccessData(v int64) {
	o.SuccessData = &v
}

// GetFailData returns the FailData field value if set, zero value otherwise.
func (o *EditProductV3DefaultResponseData) GetFailData() int64 {
	if o == nil || IsNil(o.FailData) {
		var ret int64
		return ret
	}
	return *o.FailData
}

// GetFailDataOk returns a tuple with the FailData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditProductV3DefaultResponseData) GetFailDataOk() (*int64, bool) {
	if o == nil || IsNil(o.FailData) {
		return nil, false
	}
	return o.FailData, true
}

// HasFailData returns a boolean if a field has been set.
func (o *EditProductV3DefaultResponseData) HasFailData() bool {
	if o != nil && !IsNil(o.FailData) {
		return true
	}

	return false
}

// SetFailData gets a reference to the given int64 and assigns it to the FailData field.
func (o *EditProductV3DefaultResponseData) SetFailData(v int64) {
	o.FailData = &v
}

// GetFailedRowsData returns the FailedRowsData field value if set, zero value otherwise.
func (o *EditProductV3DefaultResponseData) GetFailedRowsData() []EditProductV3DefaultResponseDataFailedRowsDataInner {
	if o == nil || IsNil(o.FailedRowsData) {
		var ret []EditProductV3DefaultResponseDataFailedRowsDataInner
		return ret
	}
	return o.FailedRowsData
}

// GetFailedRowsDataOk returns a tuple with the FailedRowsData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditProductV3DefaultResponseData) GetFailedRowsDataOk() ([]EditProductV3DefaultResponseDataFailedRowsDataInner, bool) {
	if o == nil || IsNil(o.FailedRowsData) {
		return nil, false
	}
	return o.FailedRowsData, true
}

// HasFailedRowsData returns a boolean if a field has been set.
func (o *EditProductV3DefaultResponseData) HasFailedRowsData() bool {
	if o != nil && !IsNil(o.FailedRowsData) {
		return true
	}

	return false
}

// SetFailedRowsData gets a reference to the given []EditProductV3DefaultResponseDataFailedRowsDataInner and assigns it to the FailedRowsData field.
func (o *EditProductV3DefaultResponseData) SetFailedRowsData(v []EditProductV3DefaultResponseDataFailedRowsDataInner) {
	o.FailedRowsData = v
}

func (o EditProductV3DefaultResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditProductV3DefaultResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalData) {
		toSerialize["total_data"] = o.TotalData
	}
	if !IsNil(o.SuccessData) {
		toSerialize["success_data"] = o.SuccessData
	}
	if !IsNil(o.FailData) {
		toSerialize["fail_data"] = o.FailData
	}
	if !IsNil(o.FailedRowsData) {
		toSerialize["failed_rows_data"] = o.FailedRowsData
	}
	return toSerialize, nil
}

type NullableEditProductV3DefaultResponseData struct {
	value *EditProductV3DefaultResponseData
	isSet bool
}

func (v NullableEditProductV3DefaultResponseData) Get() *EditProductV3DefaultResponseData {
	return v.value
}

func (v *NullableEditProductV3DefaultResponseData) Set(val *EditProductV3DefaultResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableEditProductV3DefaultResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableEditProductV3DefaultResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditProductV3DefaultResponseData(val *EditProductV3DefaultResponseData) *NullableEditProductV3DefaultResponseData {
	return &NullableEditProductV3DefaultResponseData{value: val, isSet: true}
}

func (v NullableEditProductV3DefaultResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditProductV3DefaultResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

