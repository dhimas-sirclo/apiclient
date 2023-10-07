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

// checks if the GetAllCategories200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAllCategories200ResponseData{}

// GetAllCategories200ResponseData struct for GetAllCategories200ResponseData
type GetAllCategories200ResponseData struct {
	Categories []GetAllCategories200ResponseDataCategoriesInner `json:"categories,omitempty"`
}

// NewGetAllCategories200ResponseData instantiates a new GetAllCategories200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAllCategories200ResponseData() *GetAllCategories200ResponseData {
	this := GetAllCategories200ResponseData{}
	return &this
}

// NewGetAllCategories200ResponseDataWithDefaults instantiates a new GetAllCategories200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAllCategories200ResponseDataWithDefaults() *GetAllCategories200ResponseData {
	this := GetAllCategories200ResponseData{}
	return &this
}

// GetCategories returns the Categories field value if set, zero value otherwise.
func (o *GetAllCategories200ResponseData) GetCategories() []GetAllCategories200ResponseDataCategoriesInner {
	if o == nil || IsNil(o.Categories) {
		var ret []GetAllCategories200ResponseDataCategoriesInner
		return ret
	}
	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllCategories200ResponseData) GetCategoriesOk() ([]GetAllCategories200ResponseDataCategoriesInner, bool) {
	if o == nil || IsNil(o.Categories) {
		return nil, false
	}
	return o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *GetAllCategories200ResponseData) HasCategories() bool {
	if o != nil && !IsNil(o.Categories) {
		return true
	}

	return false
}

// SetCategories gets a reference to the given []GetAllCategories200ResponseDataCategoriesInner and assigns it to the Categories field.
func (o *GetAllCategories200ResponseData) SetCategories(v []GetAllCategories200ResponseDataCategoriesInner) {
	o.Categories = v
}

func (o GetAllCategories200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAllCategories200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Categories) {
		toSerialize["categories"] = o.Categories
	}
	return toSerialize, nil
}

type NullableGetAllCategories200ResponseData struct {
	value *GetAllCategories200ResponseData
	isSet bool
}

func (v NullableGetAllCategories200ResponseData) Get() *GetAllCategories200ResponseData {
	return v.value
}

func (v *NullableGetAllCategories200ResponseData) Set(val *GetAllCategories200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAllCategories200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAllCategories200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAllCategories200ResponseData(val *GetAllCategories200ResponseData) *NullableGetAllCategories200ResponseData {
	return &NullableGetAllCategories200ResponseData{value: val, isSet: true}
}

func (v NullableGetAllCategories200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAllCategories200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


