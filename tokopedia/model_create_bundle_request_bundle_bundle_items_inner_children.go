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

// checks if the CreateBundleRequestBundleBundleItemsInnerChildren type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateBundleRequestBundleBundleItemsInnerChildren{}

// CreateBundleRequestBundleBundleItemsInnerChildren The variant of the product. It has product_id and bundle_price
type CreateBundleRequestBundleBundleItemsInnerChildren struct {
	VariantId int64 `json:"variant_id"`
	VariantPrice int64 `json:"variant_price"`
}

// NewCreateBundleRequestBundleBundleItemsInnerChildren instantiates a new CreateBundleRequestBundleBundleItemsInnerChildren object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateBundleRequestBundleBundleItemsInnerChildren(variantId int64, variantPrice int64) *CreateBundleRequestBundleBundleItemsInnerChildren {
	this := CreateBundleRequestBundleBundleItemsInnerChildren{}
	this.VariantId = variantId
	this.VariantPrice = variantPrice
	return &this
}

// NewCreateBundleRequestBundleBundleItemsInnerChildrenWithDefaults instantiates a new CreateBundleRequestBundleBundleItemsInnerChildren object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateBundleRequestBundleBundleItemsInnerChildrenWithDefaults() *CreateBundleRequestBundleBundleItemsInnerChildren {
	this := CreateBundleRequestBundleBundleItemsInnerChildren{}
	return &this
}

// GetVariantId returns the VariantId field value
func (o *CreateBundleRequestBundleBundleItemsInnerChildren) GetVariantId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.VariantId
}

// GetVariantIdOk returns a tuple with the VariantId field value
// and a boolean to check if the value has been set.
func (o *CreateBundleRequestBundleBundleItemsInnerChildren) GetVariantIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VariantId, true
}

// SetVariantId sets field value
func (o *CreateBundleRequestBundleBundleItemsInnerChildren) SetVariantId(v int64) {
	o.VariantId = v
}

// GetVariantPrice returns the VariantPrice field value
func (o *CreateBundleRequestBundleBundleItemsInnerChildren) GetVariantPrice() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.VariantPrice
}

// GetVariantPriceOk returns a tuple with the VariantPrice field value
// and a boolean to check if the value has been set.
func (o *CreateBundleRequestBundleBundleItemsInnerChildren) GetVariantPriceOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VariantPrice, true
}

// SetVariantPrice sets field value
func (o *CreateBundleRequestBundleBundleItemsInnerChildren) SetVariantPrice(v int64) {
	o.VariantPrice = v
}

func (o CreateBundleRequestBundleBundleItemsInnerChildren) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateBundleRequestBundleBundleItemsInnerChildren) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["variant_id"] = o.VariantId
	toSerialize["variant_price"] = o.VariantPrice
	return toSerialize, nil
}

type NullableCreateBundleRequestBundleBundleItemsInnerChildren struct {
	value *CreateBundleRequestBundleBundleItemsInnerChildren
	isSet bool
}

func (v NullableCreateBundleRequestBundleBundleItemsInnerChildren) Get() *CreateBundleRequestBundleBundleItemsInnerChildren {
	return v.value
}

func (v *NullableCreateBundleRequestBundleBundleItemsInnerChildren) Set(val *CreateBundleRequestBundleBundleItemsInnerChildren) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateBundleRequestBundleBundleItemsInnerChildren) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateBundleRequestBundleBundleItemsInnerChildren) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateBundleRequestBundleBundleItemsInnerChildren(val *CreateBundleRequestBundleBundleItemsInnerChildren) *NullableCreateBundleRequestBundleBundleItemsInnerChildren {
	return &NullableCreateBundleRequestBundleBundleItemsInnerChildren{value: val, isSet: true}
}

func (v NullableCreateBundleRequestBundleBundleItemsInnerChildren) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateBundleRequestBundleBundleItemsInnerChildren) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

