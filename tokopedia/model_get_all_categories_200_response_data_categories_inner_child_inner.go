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

// checks if the GetAllCategories200ResponseDataCategoriesInnerChildInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAllCategories200ResponseDataCategoriesInnerChildInner{}

// GetAllCategories200ResponseDataCategoriesInnerChildInner struct for GetAllCategories200ResponseDataCategoriesInnerChildInner
type GetAllCategories200ResponseDataCategoriesInnerChildInner struct {
	Name *string `json:"Name,omitempty"`
	Id *string `json:"id,omitempty"`
	Child []GetAllCategories200ResponseDataCategoriesInnerChildInnerChildInner `json:"child,omitempty"`
}

// NewGetAllCategories200ResponseDataCategoriesInnerChildInner instantiates a new GetAllCategories200ResponseDataCategoriesInnerChildInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAllCategories200ResponseDataCategoriesInnerChildInner() *GetAllCategories200ResponseDataCategoriesInnerChildInner {
	this := GetAllCategories200ResponseDataCategoriesInnerChildInner{}
	return &this
}

// NewGetAllCategories200ResponseDataCategoriesInnerChildInnerWithDefaults instantiates a new GetAllCategories200ResponseDataCategoriesInnerChildInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAllCategories200ResponseDataCategoriesInnerChildInnerWithDefaults() *GetAllCategories200ResponseDataCategoriesInnerChildInner {
	this := GetAllCategories200ResponseDataCategoriesInnerChildInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetAllCategories200ResponseDataCategoriesInnerChildInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllCategories200ResponseDataCategoriesInnerChildInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetAllCategories200ResponseDataCategoriesInnerChildInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetAllCategories200ResponseDataCategoriesInnerChildInner) SetName(v string) {
	o.Name = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetAllCategories200ResponseDataCategoriesInnerChildInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllCategories200ResponseDataCategoriesInnerChildInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetAllCategories200ResponseDataCategoriesInnerChildInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetAllCategories200ResponseDataCategoriesInnerChildInner) SetId(v string) {
	o.Id = &v
}

// GetChild returns the Child field value if set, zero value otherwise.
func (o *GetAllCategories200ResponseDataCategoriesInnerChildInner) GetChild() []GetAllCategories200ResponseDataCategoriesInnerChildInnerChildInner {
	if o == nil || IsNil(o.Child) {
		var ret []GetAllCategories200ResponseDataCategoriesInnerChildInnerChildInner
		return ret
	}
	return o.Child
}

// GetChildOk returns a tuple with the Child field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllCategories200ResponseDataCategoriesInnerChildInner) GetChildOk() ([]GetAllCategories200ResponseDataCategoriesInnerChildInnerChildInner, bool) {
	if o == nil || IsNil(o.Child) {
		return nil, false
	}
	return o.Child, true
}

// HasChild returns a boolean if a field has been set.
func (o *GetAllCategories200ResponseDataCategoriesInnerChildInner) HasChild() bool {
	if o != nil && !IsNil(o.Child) {
		return true
	}

	return false
}

// SetChild gets a reference to the given []GetAllCategories200ResponseDataCategoriesInnerChildInnerChildInner and assigns it to the Child field.
func (o *GetAllCategories200ResponseDataCategoriesInnerChildInner) SetChild(v []GetAllCategories200ResponseDataCategoriesInnerChildInnerChildInner) {
	o.Child = v
}

func (o GetAllCategories200ResponseDataCategoriesInnerChildInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAllCategories200ResponseDataCategoriesInnerChildInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Child) {
		toSerialize["child"] = o.Child
	}
	return toSerialize, nil
}

type NullableGetAllCategories200ResponseDataCategoriesInnerChildInner struct {
	value *GetAllCategories200ResponseDataCategoriesInnerChildInner
	isSet bool
}

func (v NullableGetAllCategories200ResponseDataCategoriesInnerChildInner) Get() *GetAllCategories200ResponseDataCategoriesInnerChildInner {
	return v.value
}

func (v *NullableGetAllCategories200ResponseDataCategoriesInnerChildInner) Set(val *GetAllCategories200ResponseDataCategoriesInnerChildInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAllCategories200ResponseDataCategoriesInnerChildInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAllCategories200ResponseDataCategoriesInnerChildInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAllCategories200ResponseDataCategoriesInnerChildInner(val *GetAllCategories200ResponseDataCategoriesInnerChildInner) *NullableGetAllCategories200ResponseDataCategoriesInnerChildInner {
	return &NullableGetAllCategories200ResponseDataCategoriesInnerChildInner{value: val, isSet: true}
}

func (v NullableGetAllCategories200ResponseDataCategoriesInnerChildInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAllCategories200ResponseDataCategoriesInnerChildInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

