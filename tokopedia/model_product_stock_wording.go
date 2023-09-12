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

// checks if the ProductStockWording type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductStockWording{}

// ProductStockWording struct for ProductStockWording
type ProductStockWording struct {
	UseStock *bool `json:"useStock,omitempty"`
	// Product Total Stock
	Value *int32 `json:"value,omitempty"`
	// Product Stock Wording (Description)
	StockWording *string `json:"stockWording,omitempty"`
}

// NewProductStockWording instantiates a new ProductStockWording object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductStockWording() *ProductStockWording {
	this := ProductStockWording{}
	return &this
}

// NewProductStockWordingWithDefaults instantiates a new ProductStockWording object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductStockWordingWithDefaults() *ProductStockWording {
	this := ProductStockWording{}
	return &this
}

// GetUseStock returns the UseStock field value if set, zero value otherwise.
func (o *ProductStockWording) GetUseStock() bool {
	if o == nil || IsNil(o.UseStock) {
		var ret bool
		return ret
	}
	return *o.UseStock
}

// GetUseStockOk returns a tuple with the UseStock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductStockWording) GetUseStockOk() (*bool, bool) {
	if o == nil || IsNil(o.UseStock) {
		return nil, false
	}
	return o.UseStock, true
}

// HasUseStock returns a boolean if a field has been set.
func (o *ProductStockWording) HasUseStock() bool {
	if o != nil && !IsNil(o.UseStock) {
		return true
	}

	return false
}

// SetUseStock gets a reference to the given bool and assigns it to the UseStock field.
func (o *ProductStockWording) SetUseStock(v bool) {
	o.UseStock = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ProductStockWording) GetValue() int32 {
	if o == nil || IsNil(o.Value) {
		var ret int32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductStockWording) GetValueOk() (*int32, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ProductStockWording) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given int32 and assigns it to the Value field.
func (o *ProductStockWording) SetValue(v int32) {
	o.Value = &v
}

// GetStockWording returns the StockWording field value if set, zero value otherwise.
func (o *ProductStockWording) GetStockWording() string {
	if o == nil || IsNil(o.StockWording) {
		var ret string
		return ret
	}
	return *o.StockWording
}

// GetStockWordingOk returns a tuple with the StockWording field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductStockWording) GetStockWordingOk() (*string, bool) {
	if o == nil || IsNil(o.StockWording) {
		return nil, false
	}
	return o.StockWording, true
}

// HasStockWording returns a boolean if a field has been set.
func (o *ProductStockWording) HasStockWording() bool {
	if o != nil && !IsNil(o.StockWording) {
		return true
	}

	return false
}

// SetStockWording gets a reference to the given string and assigns it to the StockWording field.
func (o *ProductStockWording) SetStockWording(v string) {
	o.StockWording = &v
}

func (o ProductStockWording) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductStockWording) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UseStock) {
		toSerialize["useStock"] = o.UseStock
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.StockWording) {
		toSerialize["stockWording"] = o.StockWording
	}
	return toSerialize, nil
}

type NullableProductStockWording struct {
	value *ProductStockWording
	isSet bool
}

func (v NullableProductStockWording) Get() *ProductStockWording {
	return v.value
}

func (v *NullableProductStockWording) Set(val *ProductStockWording) {
	v.value = val
	v.isSet = true
}

func (v NullableProductStockWording) IsSet() bool {
	return v.isSet
}

func (v *NullableProductStockWording) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductStockWording(val *ProductStockWording) *NullableProductStockWording {
	return &NullableProductStockWording{value: val, isSet: true}
}

func (v NullableProductStockWording) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductStockWording) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


