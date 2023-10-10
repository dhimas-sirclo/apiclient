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

// checks if the EditProductV2RequestProductsInnerVariantSelectionInnerOptionsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditProductV2RequestProductsInnerVariantSelectionInnerOptionsInner{}

// EditProductV2RequestProductsInnerVariantSelectionInnerOptionsInner struct for EditProductV2RequestProductsInnerVariantSelectionInnerOptionsInner
type EditProductV2RequestProductsInnerVariantSelectionInnerOptionsInner struct {
	HexCode string `json:"hex_code"`
	UnitValueId int64 `json:"unit_value_id"`
	Value string `json:"value"`
}

// NewEditProductV2RequestProductsInnerVariantSelectionInnerOptionsInner instantiates a new EditProductV2RequestProductsInnerVariantSelectionInnerOptionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditProductV2RequestProductsInnerVariantSelectionInnerOptionsInner(hexCode string, unitValueId int64, value string) *EditProductV2RequestProductsInnerVariantSelectionInnerOptionsInner {
	this := EditProductV2RequestProductsInnerVariantSelectionInnerOptionsInner{}
	this.HexCode = hexCode
	this.UnitValueId = unitValueId
	this.Value = value
	return &this
}

// NewEditProductV2RequestProductsInnerVariantSelectionInnerOptionsInnerWithDefaults instantiates a new EditProductV2RequestProductsInnerVariantSelectionInnerOptionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditProductV2RequestProductsInnerVariantSelectionInnerOptionsInnerWithDefaults() *EditProductV2RequestProductsInnerVariantSelectionInnerOptionsInner {
	this := EditProductV2RequestProductsInnerVariantSelectionInnerOptionsInner{}
	return &this
}

// GetHexCode returns the HexCode field value
func (o *EditProductV2RequestProductsInnerVariantSelectionInnerOptionsInner) GetHexCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HexCode
}

// GetHexCodeOk returns a tuple with the HexCode field value
// and a boolean to check if the value has been set.
func (o *EditProductV2RequestProductsInnerVariantSelectionInnerOptionsInner) GetHexCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HexCode, true
}

// SetHexCode sets field value
func (o *EditProductV2RequestProductsInnerVariantSelectionInnerOptionsInner) SetHexCode(v string) {
	o.HexCode = v
}

// GetUnitValueId returns the UnitValueId field value
func (o *EditProductV2RequestProductsInnerVariantSelectionInnerOptionsInner) GetUnitValueId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UnitValueId
}

// GetUnitValueIdOk returns a tuple with the UnitValueId field value
// and a boolean to check if the value has been set.
func (o *EditProductV2RequestProductsInnerVariantSelectionInnerOptionsInner) GetUnitValueIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnitValueId, true
}

// SetUnitValueId sets field value
func (o *EditProductV2RequestProductsInnerVariantSelectionInnerOptionsInner) SetUnitValueId(v int64) {
	o.UnitValueId = v
}

// GetValue returns the Value field value
func (o *EditProductV2RequestProductsInnerVariantSelectionInnerOptionsInner) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *EditProductV2RequestProductsInnerVariantSelectionInnerOptionsInner) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *EditProductV2RequestProductsInnerVariantSelectionInnerOptionsInner) SetValue(v string) {
	o.Value = v
}

func (o EditProductV2RequestProductsInnerVariantSelectionInnerOptionsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditProductV2RequestProductsInnerVariantSelectionInnerOptionsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["hex_code"] = o.HexCode
	toSerialize["unit_value_id"] = o.UnitValueId
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

type NullableEditProductV2RequestProductsInnerVariantSelectionInnerOptionsInner struct {
	value *EditProductV2RequestProductsInnerVariantSelectionInnerOptionsInner
	isSet bool
}

func (v NullableEditProductV2RequestProductsInnerVariantSelectionInnerOptionsInner) Get() *EditProductV2RequestProductsInnerVariantSelectionInnerOptionsInner {
	return v.value
}

func (v *NullableEditProductV2RequestProductsInnerVariantSelectionInnerOptionsInner) Set(val *EditProductV2RequestProductsInnerVariantSelectionInnerOptionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableEditProductV2RequestProductsInnerVariantSelectionInnerOptionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableEditProductV2RequestProductsInnerVariantSelectionInnerOptionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditProductV2RequestProductsInnerVariantSelectionInnerOptionsInner(val *EditProductV2RequestProductsInnerVariantSelectionInnerOptionsInner) *NullableEditProductV2RequestProductsInnerVariantSelectionInnerOptionsInner {
	return &NullableEditProductV2RequestProductsInnerVariantSelectionInnerOptionsInner{value: val, isSet: true}
}

func (v NullableEditProductV2RequestProductsInnerVariantSelectionInnerOptionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditProductV2RequestProductsInnerVariantSelectionInnerOptionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

