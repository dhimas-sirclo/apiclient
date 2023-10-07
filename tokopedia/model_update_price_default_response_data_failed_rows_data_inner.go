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

// checks if the UpdatePriceDefaultResponseDataFailedRowsDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdatePriceDefaultResponseDataFailedRowsDataInner{}

// UpdatePriceDefaultResponseDataFailedRowsDataInner struct for UpdatePriceDefaultResponseDataFailedRowsDataInner
type UpdatePriceDefaultResponseDataFailedRowsDataInner struct {
	ProductId *int64 `json:"product_id,omitempty"`
	Sku *string `json:"sku,omitempty"`
	ProductUrl *string `json:"product_url,omitempty"`
	NewPrice *int64 `json:"new_price,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewUpdatePriceDefaultResponseDataFailedRowsDataInner instantiates a new UpdatePriceDefaultResponseDataFailedRowsDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdatePriceDefaultResponseDataFailedRowsDataInner() *UpdatePriceDefaultResponseDataFailedRowsDataInner {
	this := UpdatePriceDefaultResponseDataFailedRowsDataInner{}
	return &this
}

// NewUpdatePriceDefaultResponseDataFailedRowsDataInnerWithDefaults instantiates a new UpdatePriceDefaultResponseDataFailedRowsDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdatePriceDefaultResponseDataFailedRowsDataInnerWithDefaults() *UpdatePriceDefaultResponseDataFailedRowsDataInner {
	this := UpdatePriceDefaultResponseDataFailedRowsDataInner{}
	return &this
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *UpdatePriceDefaultResponseDataFailedRowsDataInner) GetProductId() int64 {
	if o == nil || IsNil(o.ProductId) {
		var ret int64
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePriceDefaultResponseDataFailedRowsDataInner) GetProductIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ProductId) {
		return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *UpdatePriceDefaultResponseDataFailedRowsDataInner) HasProductId() bool {
	if o != nil && !IsNil(o.ProductId) {
		return true
	}

	return false
}

// SetProductId gets a reference to the given int64 and assigns it to the ProductId field.
func (o *UpdatePriceDefaultResponseDataFailedRowsDataInner) SetProductId(v int64) {
	o.ProductId = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *UpdatePriceDefaultResponseDataFailedRowsDataInner) GetSku() string {
	if o == nil || IsNil(o.Sku) {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePriceDefaultResponseDataFailedRowsDataInner) GetSkuOk() (*string, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *UpdatePriceDefaultResponseDataFailedRowsDataInner) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *UpdatePriceDefaultResponseDataFailedRowsDataInner) SetSku(v string) {
	o.Sku = &v
}

// GetProductUrl returns the ProductUrl field value if set, zero value otherwise.
func (o *UpdatePriceDefaultResponseDataFailedRowsDataInner) GetProductUrl() string {
	if o == nil || IsNil(o.ProductUrl) {
		var ret string
		return ret
	}
	return *o.ProductUrl
}

// GetProductUrlOk returns a tuple with the ProductUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePriceDefaultResponseDataFailedRowsDataInner) GetProductUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ProductUrl) {
		return nil, false
	}
	return o.ProductUrl, true
}

// HasProductUrl returns a boolean if a field has been set.
func (o *UpdatePriceDefaultResponseDataFailedRowsDataInner) HasProductUrl() bool {
	if o != nil && !IsNil(o.ProductUrl) {
		return true
	}

	return false
}

// SetProductUrl gets a reference to the given string and assigns it to the ProductUrl field.
func (o *UpdatePriceDefaultResponseDataFailedRowsDataInner) SetProductUrl(v string) {
	o.ProductUrl = &v
}

// GetNewPrice returns the NewPrice field value if set, zero value otherwise.
func (o *UpdatePriceDefaultResponseDataFailedRowsDataInner) GetNewPrice() int64 {
	if o == nil || IsNil(o.NewPrice) {
		var ret int64
		return ret
	}
	return *o.NewPrice
}

// GetNewPriceOk returns a tuple with the NewPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePriceDefaultResponseDataFailedRowsDataInner) GetNewPriceOk() (*int64, bool) {
	if o == nil || IsNil(o.NewPrice) {
		return nil, false
	}
	return o.NewPrice, true
}

// HasNewPrice returns a boolean if a field has been set.
func (o *UpdatePriceDefaultResponseDataFailedRowsDataInner) HasNewPrice() bool {
	if o != nil && !IsNil(o.NewPrice) {
		return true
	}

	return false
}

// SetNewPrice gets a reference to the given int64 and assigns it to the NewPrice field.
func (o *UpdatePriceDefaultResponseDataFailedRowsDataInner) SetNewPrice(v int64) {
	o.NewPrice = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *UpdatePriceDefaultResponseDataFailedRowsDataInner) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePriceDefaultResponseDataFailedRowsDataInner) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *UpdatePriceDefaultResponseDataFailedRowsDataInner) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *UpdatePriceDefaultResponseDataFailedRowsDataInner) SetMessage(v string) {
	o.Message = &v
}

func (o UpdatePriceDefaultResponseDataFailedRowsDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdatePriceDefaultResponseDataFailedRowsDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProductId) {
		toSerialize["product_id"] = o.ProductId
	}
	if !IsNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	if !IsNil(o.ProductUrl) {
		toSerialize["product_url"] = o.ProductUrl
	}
	if !IsNil(o.NewPrice) {
		toSerialize["new_price"] = o.NewPrice
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableUpdatePriceDefaultResponseDataFailedRowsDataInner struct {
	value *UpdatePriceDefaultResponseDataFailedRowsDataInner
	isSet bool
}

func (v NullableUpdatePriceDefaultResponseDataFailedRowsDataInner) Get() *UpdatePriceDefaultResponseDataFailedRowsDataInner {
	return v.value
}

func (v *NullableUpdatePriceDefaultResponseDataFailedRowsDataInner) Set(val *UpdatePriceDefaultResponseDataFailedRowsDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdatePriceDefaultResponseDataFailedRowsDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdatePriceDefaultResponseDataFailedRowsDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdatePriceDefaultResponseDataFailedRowsDataInner(val *UpdatePriceDefaultResponseDataFailedRowsDataInner) *NullableUpdatePriceDefaultResponseDataFailedRowsDataInner {
	return &NullableUpdatePriceDefaultResponseDataFailedRowsDataInner{value: val, isSet: true}
}

func (v NullableUpdatePriceDefaultResponseDataFailedRowsDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdatePriceDefaultResponseDataFailedRowsDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


