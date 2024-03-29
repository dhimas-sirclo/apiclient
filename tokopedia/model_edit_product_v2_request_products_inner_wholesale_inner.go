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

// checks if the EditProductV2RequestProductsInnerWholesaleInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditProductV2RequestProductsInnerWholesaleInner{}

// EditProductV2RequestProductsInnerWholesaleInner struct for EditProductV2RequestProductsInnerWholesaleInner
type EditProductV2RequestProductsInnerWholesaleInner struct {
	MinQty int64 `json:"min_qty"`
	Price float64 `json:"price"`
}

// NewEditProductV2RequestProductsInnerWholesaleInner instantiates a new EditProductV2RequestProductsInnerWholesaleInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditProductV2RequestProductsInnerWholesaleInner(minQty int64, price float64) *EditProductV2RequestProductsInnerWholesaleInner {
	this := EditProductV2RequestProductsInnerWholesaleInner{}
	this.MinQty = minQty
	this.Price = price
	return &this
}

// NewEditProductV2RequestProductsInnerWholesaleInnerWithDefaults instantiates a new EditProductV2RequestProductsInnerWholesaleInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditProductV2RequestProductsInnerWholesaleInnerWithDefaults() *EditProductV2RequestProductsInnerWholesaleInner {
	this := EditProductV2RequestProductsInnerWholesaleInner{}
	return &this
}

// GetMinQty returns the MinQty field value
func (o *EditProductV2RequestProductsInnerWholesaleInner) GetMinQty() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.MinQty
}

// GetMinQtyOk returns a tuple with the MinQty field value
// and a boolean to check if the value has been set.
func (o *EditProductV2RequestProductsInnerWholesaleInner) GetMinQtyOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinQty, true
}

// SetMinQty sets field value
func (o *EditProductV2RequestProductsInnerWholesaleInner) SetMinQty(v int64) {
	o.MinQty = v
}

// GetPrice returns the Price field value
func (o *EditProductV2RequestProductsInnerWholesaleInner) GetPrice() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *EditProductV2RequestProductsInnerWholesaleInner) GetPriceOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Price, true
}

// SetPrice sets field value
func (o *EditProductV2RequestProductsInnerWholesaleInner) SetPrice(v float64) {
	o.Price = v
}

func (o EditProductV2RequestProductsInnerWholesaleInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditProductV2RequestProductsInnerWholesaleInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["min_qty"] = o.MinQty
	toSerialize["price"] = o.Price
	return toSerialize, nil
}

type NullableEditProductV2RequestProductsInnerWholesaleInner struct {
	value *EditProductV2RequestProductsInnerWholesaleInner
	isSet bool
}

func (v NullableEditProductV2RequestProductsInnerWholesaleInner) Get() *EditProductV2RequestProductsInnerWholesaleInner {
	return v.value
}

func (v *NullableEditProductV2RequestProductsInnerWholesaleInner) Set(val *EditProductV2RequestProductsInnerWholesaleInner) {
	v.value = val
	v.isSet = true
}

func (v NullableEditProductV2RequestProductsInnerWholesaleInner) IsSet() bool {
	return v.isSet
}

func (v *NullableEditProductV2RequestProductsInnerWholesaleInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditProductV2RequestProductsInnerWholesaleInner(val *EditProductV2RequestProductsInnerWholesaleInner) *NullableEditProductV2RequestProductsInnerWholesaleInner {
	return &NullableEditProductV2RequestProductsInnerWholesaleInner{value: val, isSet: true}
}

func (v NullableEditProductV2RequestProductsInnerWholesaleInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditProductV2RequestProductsInnerWholesaleInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


