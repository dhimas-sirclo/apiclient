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

// checks if the GetProductAnnotationByCategoryId200ResponseDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetProductAnnotationByCategoryId200ResponseDataInner{}

// GetProductAnnotationByCategoryId200ResponseDataInner struct for GetProductAnnotationByCategoryId200ResponseDataInner
type GetProductAnnotationByCategoryId200ResponseDataInner struct {
	// Variant Name
	Variant *string `json:"variant,omitempty"`
	SortOrder *int64 `json:"sort_order,omitempty"`
	Values []GetProductAnnotationByCategoryId200ResponseDataInnerValuesInner `json:"values,omitempty"`
}

// NewGetProductAnnotationByCategoryId200ResponseDataInner instantiates a new GetProductAnnotationByCategoryId200ResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetProductAnnotationByCategoryId200ResponseDataInner() *GetProductAnnotationByCategoryId200ResponseDataInner {
	this := GetProductAnnotationByCategoryId200ResponseDataInner{}
	return &this
}

// NewGetProductAnnotationByCategoryId200ResponseDataInnerWithDefaults instantiates a new GetProductAnnotationByCategoryId200ResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetProductAnnotationByCategoryId200ResponseDataInnerWithDefaults() *GetProductAnnotationByCategoryId200ResponseDataInner {
	this := GetProductAnnotationByCategoryId200ResponseDataInner{}
	return &this
}

// GetVariant returns the Variant field value if set, zero value otherwise.
func (o *GetProductAnnotationByCategoryId200ResponseDataInner) GetVariant() string {
	if o == nil || IsNil(o.Variant) {
		var ret string
		return ret
	}
	return *o.Variant
}

// GetVariantOk returns a tuple with the Variant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductAnnotationByCategoryId200ResponseDataInner) GetVariantOk() (*string, bool) {
	if o == nil || IsNil(o.Variant) {
		return nil, false
	}
	return o.Variant, true
}

// HasVariant returns a boolean if a field has been set.
func (o *GetProductAnnotationByCategoryId200ResponseDataInner) HasVariant() bool {
	if o != nil && !IsNil(o.Variant) {
		return true
	}

	return false
}

// SetVariant gets a reference to the given string and assigns it to the Variant field.
func (o *GetProductAnnotationByCategoryId200ResponseDataInner) SetVariant(v string) {
	o.Variant = &v
}

// GetSortOrder returns the SortOrder field value if set, zero value otherwise.
func (o *GetProductAnnotationByCategoryId200ResponseDataInner) GetSortOrder() int64 {
	if o == nil || IsNil(o.SortOrder) {
		var ret int64
		return ret
	}
	return *o.SortOrder
}

// GetSortOrderOk returns a tuple with the SortOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductAnnotationByCategoryId200ResponseDataInner) GetSortOrderOk() (*int64, bool) {
	if o == nil || IsNil(o.SortOrder) {
		return nil, false
	}
	return o.SortOrder, true
}

// HasSortOrder returns a boolean if a field has been set.
func (o *GetProductAnnotationByCategoryId200ResponseDataInner) HasSortOrder() bool {
	if o != nil && !IsNil(o.SortOrder) {
		return true
	}

	return false
}

// SetSortOrder gets a reference to the given int64 and assigns it to the SortOrder field.
func (o *GetProductAnnotationByCategoryId200ResponseDataInner) SetSortOrder(v int64) {
	o.SortOrder = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *GetProductAnnotationByCategoryId200ResponseDataInner) GetValues() []GetProductAnnotationByCategoryId200ResponseDataInnerValuesInner {
	if o == nil || IsNil(o.Values) {
		var ret []GetProductAnnotationByCategoryId200ResponseDataInnerValuesInner
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductAnnotationByCategoryId200ResponseDataInner) GetValuesOk() ([]GetProductAnnotationByCategoryId200ResponseDataInnerValuesInner, bool) {
	if o == nil || IsNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *GetProductAnnotationByCategoryId200ResponseDataInner) HasValues() bool {
	if o != nil && !IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given []GetProductAnnotationByCategoryId200ResponseDataInnerValuesInner and assigns it to the Values field.
func (o *GetProductAnnotationByCategoryId200ResponseDataInner) SetValues(v []GetProductAnnotationByCategoryId200ResponseDataInnerValuesInner) {
	o.Values = v
}

func (o GetProductAnnotationByCategoryId200ResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetProductAnnotationByCategoryId200ResponseDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Variant) {
		toSerialize["variant"] = o.Variant
	}
	if !IsNil(o.SortOrder) {
		toSerialize["sort_order"] = o.SortOrder
	}
	if !IsNil(o.Values) {
		toSerialize["values"] = o.Values
	}
	return toSerialize, nil
}

type NullableGetProductAnnotationByCategoryId200ResponseDataInner struct {
	value *GetProductAnnotationByCategoryId200ResponseDataInner
	isSet bool
}

func (v NullableGetProductAnnotationByCategoryId200ResponseDataInner) Get() *GetProductAnnotationByCategoryId200ResponseDataInner {
	return v.value
}

func (v *NullableGetProductAnnotationByCategoryId200ResponseDataInner) Set(val *GetProductAnnotationByCategoryId200ResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetProductAnnotationByCategoryId200ResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetProductAnnotationByCategoryId200ResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetProductAnnotationByCategoryId200ResponseDataInner(val *GetProductAnnotationByCategoryId200ResponseDataInner) *NullableGetProductAnnotationByCategoryId200ResponseDataInner {
	return &NullableGetProductAnnotationByCategoryId200ResponseDataInner{value: val, isSet: true}
}

func (v NullableGetProductAnnotationByCategoryId200ResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetProductAnnotationByCategoryId200ResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

