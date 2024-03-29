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

// checks if the GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInner{}

// GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInner struct for GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInner
type GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInner struct {
	ProductVariantOptionId *int64 `json:"product_variant_option_id,omitempty"`
	UnitValueId *int64 `json:"unit_value_id,omitempty"`
	Value *string `json:"value,omitempty"`
	Hex *string `json:"hex,omitempty"`
	Picture *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInnerPicture `json:"picture,omitempty"`
}

// NewGetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInner instantiates a new GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInner() *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInner {
	this := GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInner{}
	return &this
}

// NewGetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInnerWithDefaults instantiates a new GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInnerWithDefaults() *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInner {
	this := GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInner{}
	return &this
}

// GetProductVariantOptionId returns the ProductVariantOptionId field value if set, zero value otherwise.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInner) GetProductVariantOptionId() int64 {
	if o == nil || IsNil(o.ProductVariantOptionId) {
		var ret int64
		return ret
	}
	return *o.ProductVariantOptionId
}

// GetProductVariantOptionIdOk returns a tuple with the ProductVariantOptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInner) GetProductVariantOptionIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ProductVariantOptionId) {
		return nil, false
	}
	return o.ProductVariantOptionId, true
}

// HasProductVariantOptionId returns a boolean if a field has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInner) HasProductVariantOptionId() bool {
	if o != nil && !IsNil(o.ProductVariantOptionId) {
		return true
	}

	return false
}

// SetProductVariantOptionId gets a reference to the given int64 and assigns it to the ProductVariantOptionId field.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInner) SetProductVariantOptionId(v int64) {
	o.ProductVariantOptionId = &v
}

// GetUnitValueId returns the UnitValueId field value if set, zero value otherwise.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInner) GetUnitValueId() int64 {
	if o == nil || IsNil(o.UnitValueId) {
		var ret int64
		return ret
	}
	return *o.UnitValueId
}

// GetUnitValueIdOk returns a tuple with the UnitValueId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInner) GetUnitValueIdOk() (*int64, bool) {
	if o == nil || IsNil(o.UnitValueId) {
		return nil, false
	}
	return o.UnitValueId, true
}

// HasUnitValueId returns a boolean if a field has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInner) HasUnitValueId() bool {
	if o != nil && !IsNil(o.UnitValueId) {
		return true
	}

	return false
}

// SetUnitValueId gets a reference to the given int64 and assigns it to the UnitValueId field.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInner) SetUnitValueId(v int64) {
	o.UnitValueId = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInner) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInner) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInner) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInner) SetValue(v string) {
	o.Value = &v
}

// GetHex returns the Hex field value if set, zero value otherwise.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInner) GetHex() string {
	if o == nil || IsNil(o.Hex) {
		var ret string
		return ret
	}
	return *o.Hex
}

// GetHexOk returns a tuple with the Hex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInner) GetHexOk() (*string, bool) {
	if o == nil || IsNil(o.Hex) {
		return nil, false
	}
	return o.Hex, true
}

// HasHex returns a boolean if a field has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInner) HasHex() bool {
	if o != nil && !IsNil(o.Hex) {
		return true
	}

	return false
}

// SetHex gets a reference to the given string and assigns it to the Hex field.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInner) SetHex(v string) {
	o.Hex = &v
}

// GetPicture returns the Picture field value if set, zero value otherwise.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInner) GetPicture() GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInnerPicture {
	if o == nil || IsNil(o.Picture) {
		var ret GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInnerPicture
		return ret
	}
	return *o.Picture
}

// GetPictureOk returns a tuple with the Picture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInner) GetPictureOk() (*GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInnerPicture, bool) {
	if o == nil || IsNil(o.Picture) {
		return nil, false
	}
	return o.Picture, true
}

// HasPicture returns a boolean if a field has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInner) HasPicture() bool {
	if o != nil && !IsNil(o.Picture) {
		return true
	}

	return false
}

// SetPicture gets a reference to the given GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInnerPicture and assigns it to the Picture field.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInner) SetPicture(v GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInnerPicture) {
	o.Picture = &v
}

func (o GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProductVariantOptionId) {
		toSerialize["product_variant_option_id"] = o.ProductVariantOptionId
	}
	if !IsNil(o.UnitValueId) {
		toSerialize["unit_value_id"] = o.UnitValueId
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Hex) {
		toSerialize["hex"] = o.Hex
	}
	if !IsNil(o.Picture) {
		toSerialize["picture"] = o.Picture
	}
	return toSerialize, nil
}

type NullableGetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInner struct {
	value *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInner
	isSet bool
}

func (v NullableGetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInner) Get() *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInner {
	return v.value
}

func (v *NullableGetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInner) Set(val *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInner(val *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInner) *NullableGetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInner {
	return &NullableGetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInner{value: val, isSet: true}
}

func (v NullableGetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


