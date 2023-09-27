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

// checks if the GetBundleList200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetBundleList200ResponseData{}

// GetBundleList200ResponseData struct for GetBundleList200ResponseData
type GetBundleList200ResponseData struct {
	BundleListInfo []GetBundleList200ResponseDataBundleListInfoInner `json:"bundle_list_info,omitempty"`
	IsLastPage *bool `json:"is_last_page,omitempty"`
	LastGroupId *int64 `json:"last_group_id,omitempty"`
}

// NewGetBundleList200ResponseData instantiates a new GetBundleList200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBundleList200ResponseData() *GetBundleList200ResponseData {
	this := GetBundleList200ResponseData{}
	return &this
}

// NewGetBundleList200ResponseDataWithDefaults instantiates a new GetBundleList200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBundleList200ResponseDataWithDefaults() *GetBundleList200ResponseData {
	this := GetBundleList200ResponseData{}
	return &this
}

// GetBundleListInfo returns the BundleListInfo field value if set, zero value otherwise.
func (o *GetBundleList200ResponseData) GetBundleListInfo() []GetBundleList200ResponseDataBundleListInfoInner {
	if o == nil || IsNil(o.BundleListInfo) {
		var ret []GetBundleList200ResponseDataBundleListInfoInner
		return ret
	}
	return o.BundleListInfo
}

// GetBundleListInfoOk returns a tuple with the BundleListInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBundleList200ResponseData) GetBundleListInfoOk() ([]GetBundleList200ResponseDataBundleListInfoInner, bool) {
	if o == nil || IsNil(o.BundleListInfo) {
		return nil, false
	}
	return o.BundleListInfo, true
}

// HasBundleListInfo returns a boolean if a field has been set.
func (o *GetBundleList200ResponseData) HasBundleListInfo() bool {
	if o != nil && !IsNil(o.BundleListInfo) {
		return true
	}

	return false
}

// SetBundleListInfo gets a reference to the given []GetBundleList200ResponseDataBundleListInfoInner and assigns it to the BundleListInfo field.
func (o *GetBundleList200ResponseData) SetBundleListInfo(v []GetBundleList200ResponseDataBundleListInfoInner) {
	o.BundleListInfo = v
}

// GetIsLastPage returns the IsLastPage field value if set, zero value otherwise.
func (o *GetBundleList200ResponseData) GetIsLastPage() bool {
	if o == nil || IsNil(o.IsLastPage) {
		var ret bool
		return ret
	}
	return *o.IsLastPage
}

// GetIsLastPageOk returns a tuple with the IsLastPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBundleList200ResponseData) GetIsLastPageOk() (*bool, bool) {
	if o == nil || IsNil(o.IsLastPage) {
		return nil, false
	}
	return o.IsLastPage, true
}

// HasIsLastPage returns a boolean if a field has been set.
func (o *GetBundleList200ResponseData) HasIsLastPage() bool {
	if o != nil && !IsNil(o.IsLastPage) {
		return true
	}

	return false
}

// SetIsLastPage gets a reference to the given bool and assigns it to the IsLastPage field.
func (o *GetBundleList200ResponseData) SetIsLastPage(v bool) {
	o.IsLastPage = &v
}

// GetLastGroupId returns the LastGroupId field value if set, zero value otherwise.
func (o *GetBundleList200ResponseData) GetLastGroupId() int64 {
	if o == nil || IsNil(o.LastGroupId) {
		var ret int64
		return ret
	}
	return *o.LastGroupId
}

// GetLastGroupIdOk returns a tuple with the LastGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBundleList200ResponseData) GetLastGroupIdOk() (*int64, bool) {
	if o == nil || IsNil(o.LastGroupId) {
		return nil, false
	}
	return o.LastGroupId, true
}

// HasLastGroupId returns a boolean if a field has been set.
func (o *GetBundleList200ResponseData) HasLastGroupId() bool {
	if o != nil && !IsNil(o.LastGroupId) {
		return true
	}

	return false
}

// SetLastGroupId gets a reference to the given int64 and assigns it to the LastGroupId field.
func (o *GetBundleList200ResponseData) SetLastGroupId(v int64) {
	o.LastGroupId = &v
}

func (o GetBundleList200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBundleList200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BundleListInfo) {
		toSerialize["bundle_list_info"] = o.BundleListInfo
	}
	if !IsNil(o.IsLastPage) {
		toSerialize["is_last_page"] = o.IsLastPage
	}
	if !IsNil(o.LastGroupId) {
		toSerialize["last_group_id"] = o.LastGroupId
	}
	return toSerialize, nil
}

type NullableGetBundleList200ResponseData struct {
	value *GetBundleList200ResponseData
	isSet bool
}

func (v NullableGetBundleList200ResponseData) Get() *GetBundleList200ResponseData {
	return v.value
}

func (v *NullableGetBundleList200ResponseData) Set(val *GetBundleList200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBundleList200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBundleList200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBundleList200ResponseData(val *GetBundleList200ResponseData) *NullableGetBundleList200ResponseData {
	return &NullableGetBundleList200ResponseData{value: val, isSet: true}
}

func (v NullableGetBundleList200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBundleList200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


