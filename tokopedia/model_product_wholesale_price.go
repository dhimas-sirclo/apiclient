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

// checks if the ProductWholesalePrice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductWholesalePrice{}

// ProductWholesalePrice struct for ProductWholesalePrice
type ProductWholesalePrice struct {
	// Product Price Value
	Value *int32 `json:"value,omitempty"`
	// Product Price Currency
	Currency *int32 `json:"currency,omitempty"`
	// Price Value
	Idr *int32 `json:"idr,omitempty"`
}

// NewProductWholesalePrice instantiates a new ProductWholesalePrice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductWholesalePrice() *ProductWholesalePrice {
	this := ProductWholesalePrice{}
	return &this
}

// NewProductWholesalePriceWithDefaults instantiates a new ProductWholesalePrice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductWholesalePriceWithDefaults() *ProductWholesalePrice {
	this := ProductWholesalePrice{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ProductWholesalePrice) GetValue() int32 {
	if o == nil || IsNil(o.Value) {
		var ret int32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductWholesalePrice) GetValueOk() (*int32, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ProductWholesalePrice) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given int32 and assigns it to the Value field.
func (o *ProductWholesalePrice) SetValue(v int32) {
	o.Value = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *ProductWholesalePrice) GetCurrency() int32 {
	if o == nil || IsNil(o.Currency) {
		var ret int32
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductWholesalePrice) GetCurrencyOk() (*int32, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *ProductWholesalePrice) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given int32 and assigns it to the Currency field.
func (o *ProductWholesalePrice) SetCurrency(v int32) {
	o.Currency = &v
}

// GetIdr returns the Idr field value if set, zero value otherwise.
func (o *ProductWholesalePrice) GetIdr() int32 {
	if o == nil || IsNil(o.Idr) {
		var ret int32
		return ret
	}
	return *o.Idr
}

// GetIdrOk returns a tuple with the Idr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductWholesalePrice) GetIdrOk() (*int32, bool) {
	if o == nil || IsNil(o.Idr) {
		return nil, false
	}
	return o.Idr, true
}

// HasIdr returns a boolean if a field has been set.
func (o *ProductWholesalePrice) HasIdr() bool {
	if o != nil && !IsNil(o.Idr) {
		return true
	}

	return false
}

// SetIdr gets a reference to the given int32 and assigns it to the Idr field.
func (o *ProductWholesalePrice) SetIdr(v int32) {
	o.Idr = &v
}

func (o ProductWholesalePrice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductWholesalePrice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.Idr) {
		toSerialize["idr"] = o.Idr
	}
	return toSerialize, nil
}

type NullableProductWholesalePrice struct {
	value *ProductWholesalePrice
	isSet bool
}

func (v NullableProductWholesalePrice) Get() *ProductWholesalePrice {
	return v.value
}

func (v *NullableProductWholesalePrice) Set(val *ProductWholesalePrice) {
	v.value = val
	v.isSet = true
}

func (v NullableProductWholesalePrice) IsSet() bool {
	return v.isSet
}

func (v *NullableProductWholesalePrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductWholesalePrice(val *ProductWholesalePrice) *NullableProductWholesalePrice {
	return &NullableProductWholesalePrice{value: val, isSet: true}
}

func (v NullableProductWholesalePrice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductWholesalePrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


