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

// checks if the ProductWarehouse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductWarehouse{}

// ProductWarehouse struct for ProductWarehouse
type ProductWarehouse struct {
	ProductID *int32 `json:"productID,omitempty"`
	WarehouseID *int32 `json:"warehouseID,omitempty"`
	Price *ProductPrice `json:"price,omitempty"`
	Stock *ProductStock `json:"stock,omitempty"`
	Bundle *ProductBundle `json:"bundle,omitempty"`
}

// NewProductWarehouse instantiates a new ProductWarehouse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductWarehouse() *ProductWarehouse {
	this := ProductWarehouse{}
	return &this
}

// NewProductWarehouseWithDefaults instantiates a new ProductWarehouse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductWarehouseWithDefaults() *ProductWarehouse {
	this := ProductWarehouse{}
	return &this
}

// GetProductID returns the ProductID field value if set, zero value otherwise.
func (o *ProductWarehouse) GetProductID() int32 {
	if o == nil || IsNil(o.ProductID) {
		var ret int32
		return ret
	}
	return *o.ProductID
}

// GetProductIDOk returns a tuple with the ProductID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductWarehouse) GetProductIDOk() (*int32, bool) {
	if o == nil || IsNil(o.ProductID) {
		return nil, false
	}
	return o.ProductID, true
}

// HasProductID returns a boolean if a field has been set.
func (o *ProductWarehouse) HasProductID() bool {
	if o != nil && !IsNil(o.ProductID) {
		return true
	}

	return false
}

// SetProductID gets a reference to the given int32 and assigns it to the ProductID field.
func (o *ProductWarehouse) SetProductID(v int32) {
	o.ProductID = &v
}

// GetWarehouseID returns the WarehouseID field value if set, zero value otherwise.
func (o *ProductWarehouse) GetWarehouseID() int32 {
	if o == nil || IsNil(o.WarehouseID) {
		var ret int32
		return ret
	}
	return *o.WarehouseID
}

// GetWarehouseIDOk returns a tuple with the WarehouseID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductWarehouse) GetWarehouseIDOk() (*int32, bool) {
	if o == nil || IsNil(o.WarehouseID) {
		return nil, false
	}
	return o.WarehouseID, true
}

// HasWarehouseID returns a boolean if a field has been set.
func (o *ProductWarehouse) HasWarehouseID() bool {
	if o != nil && !IsNil(o.WarehouseID) {
		return true
	}

	return false
}

// SetWarehouseID gets a reference to the given int32 and assigns it to the WarehouseID field.
func (o *ProductWarehouse) SetWarehouseID(v int32) {
	o.WarehouseID = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *ProductWarehouse) GetPrice() ProductPrice {
	if o == nil || IsNil(o.Price) {
		var ret ProductPrice
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductWarehouse) GetPriceOk() (*ProductPrice, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *ProductWarehouse) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given ProductPrice and assigns it to the Price field.
func (o *ProductWarehouse) SetPrice(v ProductPrice) {
	o.Price = &v
}

// GetStock returns the Stock field value if set, zero value otherwise.
func (o *ProductWarehouse) GetStock() ProductStock {
	if o == nil || IsNil(o.Stock) {
		var ret ProductStock
		return ret
	}
	return *o.Stock
}

// GetStockOk returns a tuple with the Stock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductWarehouse) GetStockOk() (*ProductStock, bool) {
	if o == nil || IsNil(o.Stock) {
		return nil, false
	}
	return o.Stock, true
}

// HasStock returns a boolean if a field has been set.
func (o *ProductWarehouse) HasStock() bool {
	if o != nil && !IsNil(o.Stock) {
		return true
	}

	return false
}

// SetStock gets a reference to the given ProductStock and assigns it to the Stock field.
func (o *ProductWarehouse) SetStock(v ProductStock) {
	o.Stock = &v
}

// GetBundle returns the Bundle field value if set, zero value otherwise.
func (o *ProductWarehouse) GetBundle() ProductBundle {
	if o == nil || IsNil(o.Bundle) {
		var ret ProductBundle
		return ret
	}
	return *o.Bundle
}

// GetBundleOk returns a tuple with the Bundle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductWarehouse) GetBundleOk() (*ProductBundle, bool) {
	if o == nil || IsNil(o.Bundle) {
		return nil, false
	}
	return o.Bundle, true
}

// HasBundle returns a boolean if a field has been set.
func (o *ProductWarehouse) HasBundle() bool {
	if o != nil && !IsNil(o.Bundle) {
		return true
	}

	return false
}

// SetBundle gets a reference to the given ProductBundle and assigns it to the Bundle field.
func (o *ProductWarehouse) SetBundle(v ProductBundle) {
	o.Bundle = &v
}

func (o ProductWarehouse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductWarehouse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProductID) {
		toSerialize["productID"] = o.ProductID
	}
	if !IsNil(o.WarehouseID) {
		toSerialize["warehouseID"] = o.WarehouseID
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.Stock) {
		toSerialize["stock"] = o.Stock
	}
	if !IsNil(o.Bundle) {
		toSerialize["bundle"] = o.Bundle
	}
	return toSerialize, nil
}

type NullableProductWarehouse struct {
	value *ProductWarehouse
	isSet bool
}

func (v NullableProductWarehouse) Get() *ProductWarehouse {
	return v.value
}

func (v *NullableProductWarehouse) Set(val *ProductWarehouse) {
	v.value = val
	v.isSet = true
}

func (v NullableProductWarehouse) IsSet() bool {
	return v.isSet
}

func (v *NullableProductWarehouse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductWarehouse(val *ProductWarehouse) *NullableProductWarehouse {
	return &NullableProductWarehouse{value: val, isSet: true}
}

func (v NullableProductWarehouse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductWarehouse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

