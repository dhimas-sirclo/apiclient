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

// checks if the EditProductV3DefaultResponseDataFailedRowsDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditProductV3DefaultResponseDataFailedRowsDataInner{}

// EditProductV3DefaultResponseDataFailedRowsDataInner struct for EditProductV3DefaultResponseDataFailedRowsDataInner
type EditProductV3DefaultResponseDataFailedRowsDataInner struct {
	ProductId *int64 `json:"product_id,omitempty"`
	ProductName *string `json:"product_name,omitempty"`
	ProductPrice *int64 `json:"product_price,omitempty"`
	Sku *string `json:"sku,omitempty"`
	Error []string `json:"error,omitempty"`
}

// NewEditProductV3DefaultResponseDataFailedRowsDataInner instantiates a new EditProductV3DefaultResponseDataFailedRowsDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditProductV3DefaultResponseDataFailedRowsDataInner() *EditProductV3DefaultResponseDataFailedRowsDataInner {
	this := EditProductV3DefaultResponseDataFailedRowsDataInner{}
	return &this
}

// NewEditProductV3DefaultResponseDataFailedRowsDataInnerWithDefaults instantiates a new EditProductV3DefaultResponseDataFailedRowsDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditProductV3DefaultResponseDataFailedRowsDataInnerWithDefaults() *EditProductV3DefaultResponseDataFailedRowsDataInner {
	this := EditProductV3DefaultResponseDataFailedRowsDataInner{}
	return &this
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *EditProductV3DefaultResponseDataFailedRowsDataInner) GetProductId() int64 {
	if o == nil || IsNil(o.ProductId) {
		var ret int64
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditProductV3DefaultResponseDataFailedRowsDataInner) GetProductIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ProductId) {
		return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *EditProductV3DefaultResponseDataFailedRowsDataInner) HasProductId() bool {
	if o != nil && !IsNil(o.ProductId) {
		return true
	}

	return false
}

// SetProductId gets a reference to the given int64 and assigns it to the ProductId field.
func (o *EditProductV3DefaultResponseDataFailedRowsDataInner) SetProductId(v int64) {
	o.ProductId = &v
}

// GetProductName returns the ProductName field value if set, zero value otherwise.
func (o *EditProductV3DefaultResponseDataFailedRowsDataInner) GetProductName() string {
	if o == nil || IsNil(o.ProductName) {
		var ret string
		return ret
	}
	return *o.ProductName
}

// GetProductNameOk returns a tuple with the ProductName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditProductV3DefaultResponseDataFailedRowsDataInner) GetProductNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProductName) {
		return nil, false
	}
	return o.ProductName, true
}

// HasProductName returns a boolean if a field has been set.
func (o *EditProductV3DefaultResponseDataFailedRowsDataInner) HasProductName() bool {
	if o != nil && !IsNil(o.ProductName) {
		return true
	}

	return false
}

// SetProductName gets a reference to the given string and assigns it to the ProductName field.
func (o *EditProductV3DefaultResponseDataFailedRowsDataInner) SetProductName(v string) {
	o.ProductName = &v
}

// GetProductPrice returns the ProductPrice field value if set, zero value otherwise.
func (o *EditProductV3DefaultResponseDataFailedRowsDataInner) GetProductPrice() int64 {
	if o == nil || IsNil(o.ProductPrice) {
		var ret int64
		return ret
	}
	return *o.ProductPrice
}

// GetProductPriceOk returns a tuple with the ProductPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditProductV3DefaultResponseDataFailedRowsDataInner) GetProductPriceOk() (*int64, bool) {
	if o == nil || IsNil(o.ProductPrice) {
		return nil, false
	}
	return o.ProductPrice, true
}

// HasProductPrice returns a boolean if a field has been set.
func (o *EditProductV3DefaultResponseDataFailedRowsDataInner) HasProductPrice() bool {
	if o != nil && !IsNil(o.ProductPrice) {
		return true
	}

	return false
}

// SetProductPrice gets a reference to the given int64 and assigns it to the ProductPrice field.
func (o *EditProductV3DefaultResponseDataFailedRowsDataInner) SetProductPrice(v int64) {
	o.ProductPrice = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *EditProductV3DefaultResponseDataFailedRowsDataInner) GetSku() string {
	if o == nil || IsNil(o.Sku) {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditProductV3DefaultResponseDataFailedRowsDataInner) GetSkuOk() (*string, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *EditProductV3DefaultResponseDataFailedRowsDataInner) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *EditProductV3DefaultResponseDataFailedRowsDataInner) SetSku(v string) {
	o.Sku = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *EditProductV3DefaultResponseDataFailedRowsDataInner) GetError() []string {
	if o == nil || IsNil(o.Error) {
		var ret []string
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditProductV3DefaultResponseDataFailedRowsDataInner) GetErrorOk() ([]string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *EditProductV3DefaultResponseDataFailedRowsDataInner) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given []string and assigns it to the Error field.
func (o *EditProductV3DefaultResponseDataFailedRowsDataInner) SetError(v []string) {
	o.Error = v
}

func (o EditProductV3DefaultResponseDataFailedRowsDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditProductV3DefaultResponseDataFailedRowsDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProductId) {
		toSerialize["product_id"] = o.ProductId
	}
	if !IsNil(o.ProductName) {
		toSerialize["product_name"] = o.ProductName
	}
	if !IsNil(o.ProductPrice) {
		toSerialize["product_price"] = o.ProductPrice
	}
	if !IsNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableEditProductV3DefaultResponseDataFailedRowsDataInner struct {
	value *EditProductV3DefaultResponseDataFailedRowsDataInner
	isSet bool
}

func (v NullableEditProductV3DefaultResponseDataFailedRowsDataInner) Get() *EditProductV3DefaultResponseDataFailedRowsDataInner {
	return v.value
}

func (v *NullableEditProductV3DefaultResponseDataFailedRowsDataInner) Set(val *EditProductV3DefaultResponseDataFailedRowsDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableEditProductV3DefaultResponseDataFailedRowsDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableEditProductV3DefaultResponseDataFailedRowsDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditProductV3DefaultResponseDataFailedRowsDataInner(val *EditProductV3DefaultResponseDataFailedRowsDataInner) *NullableEditProductV3DefaultResponseDataFailedRowsDataInner {
	return &NullableEditProductV3DefaultResponseDataFailedRowsDataInner{value: val, isSet: true}
}

func (v NullableEditProductV3DefaultResponseDataFailedRowsDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditProductV3DefaultResponseDataFailedRowsDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


