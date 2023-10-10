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

// checks if the UpdateStockIncrementRequestInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateStockIncrementRequestInner{}

// UpdateStockIncrementRequestInner struct for UpdateStockIncrementRequestInner
type UpdateStockIncrementRequestInner struct {
	// SKU of products that will be updated. Maximum characters allowed is 50
	Sku *string `json:"sku,omitempty"`
	// Product ID to update
	ProductId *int64 `json:"product_id,omitempty"`
	// New stock to be set
	StockValue int64 `json:"stock_value"`
}

// NewUpdateStockIncrementRequestInner instantiates a new UpdateStockIncrementRequestInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateStockIncrementRequestInner(stockValue int64) *UpdateStockIncrementRequestInner {
	this := UpdateStockIncrementRequestInner{}
	this.StockValue = stockValue
	return &this
}

// NewUpdateStockIncrementRequestInnerWithDefaults instantiates a new UpdateStockIncrementRequestInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateStockIncrementRequestInnerWithDefaults() *UpdateStockIncrementRequestInner {
	this := UpdateStockIncrementRequestInner{}
	return &this
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *UpdateStockIncrementRequestInner) GetSku() string {
	if o == nil || IsNil(o.Sku) {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStockIncrementRequestInner) GetSkuOk() (*string, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *UpdateStockIncrementRequestInner) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *UpdateStockIncrementRequestInner) SetSku(v string) {
	o.Sku = &v
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *UpdateStockIncrementRequestInner) GetProductId() int64 {
	if o == nil || IsNil(o.ProductId) {
		var ret int64
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStockIncrementRequestInner) GetProductIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ProductId) {
		return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *UpdateStockIncrementRequestInner) HasProductId() bool {
	if o != nil && !IsNil(o.ProductId) {
		return true
	}

	return false
}

// SetProductId gets a reference to the given int64 and assigns it to the ProductId field.
func (o *UpdateStockIncrementRequestInner) SetProductId(v int64) {
	o.ProductId = &v
}

// GetStockValue returns the StockValue field value
func (o *UpdateStockIncrementRequestInner) GetStockValue() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.StockValue
}

// GetStockValueOk returns a tuple with the StockValue field value
// and a boolean to check if the value has been set.
func (o *UpdateStockIncrementRequestInner) GetStockValueOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StockValue, true
}

// SetStockValue sets field value
func (o *UpdateStockIncrementRequestInner) SetStockValue(v int64) {
	o.StockValue = v
}

func (o UpdateStockIncrementRequestInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateStockIncrementRequestInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	if !IsNil(o.ProductId) {
		toSerialize["product_id"] = o.ProductId
	}
	toSerialize["stock_value"] = o.StockValue
	return toSerialize, nil
}

type NullableUpdateStockIncrementRequestInner struct {
	value *UpdateStockIncrementRequestInner
	isSet bool
}

func (v NullableUpdateStockIncrementRequestInner) Get() *UpdateStockIncrementRequestInner {
	return v.value
}

func (v *NullableUpdateStockIncrementRequestInner) Set(val *UpdateStockIncrementRequestInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateStockIncrementRequestInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateStockIncrementRequestInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateStockIncrementRequestInner(val *UpdateStockIncrementRequestInner) *NullableUpdateStockIncrementRequestInner {
	return &NullableUpdateStockIncrementRequestInner{value: val, isSet: true}
}

func (v NullableUpdateStockIncrementRequestInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateStockIncrementRequestInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

