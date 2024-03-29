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

// checks if the WebhookErrorProductEditDataFailedRowsDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookErrorProductEditDataFailedRowsDataInner{}

// WebhookErrorProductEditDataFailedRowsDataInner struct for WebhookErrorProductEditDataFailedRowsDataInner
type WebhookErrorProductEditDataFailedRowsDataInner struct {
	ProductId *int64 `json:"product_id,omitempty"`
	ProductName *string `json:"product_name,omitempty"`
	ProductPrice *int64 `json:"product_price,omitempty"`
	Sku *string `json:"sku,omitempty"`
	Error []string `json:"error,omitempty"`
}

// NewWebhookErrorProductEditDataFailedRowsDataInner instantiates a new WebhookErrorProductEditDataFailedRowsDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookErrorProductEditDataFailedRowsDataInner() *WebhookErrorProductEditDataFailedRowsDataInner {
	this := WebhookErrorProductEditDataFailedRowsDataInner{}
	return &this
}

// NewWebhookErrorProductEditDataFailedRowsDataInnerWithDefaults instantiates a new WebhookErrorProductEditDataFailedRowsDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookErrorProductEditDataFailedRowsDataInnerWithDefaults() *WebhookErrorProductEditDataFailedRowsDataInner {
	this := WebhookErrorProductEditDataFailedRowsDataInner{}
	return &this
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *WebhookErrorProductEditDataFailedRowsDataInner) GetProductId() int64 {
	if o == nil || IsNil(o.ProductId) {
		var ret int64
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookErrorProductEditDataFailedRowsDataInner) GetProductIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ProductId) {
		return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *WebhookErrorProductEditDataFailedRowsDataInner) HasProductId() bool {
	if o != nil && !IsNil(o.ProductId) {
		return true
	}

	return false
}

// SetProductId gets a reference to the given int64 and assigns it to the ProductId field.
func (o *WebhookErrorProductEditDataFailedRowsDataInner) SetProductId(v int64) {
	o.ProductId = &v
}

// GetProductName returns the ProductName field value if set, zero value otherwise.
func (o *WebhookErrorProductEditDataFailedRowsDataInner) GetProductName() string {
	if o == nil || IsNil(o.ProductName) {
		var ret string
		return ret
	}
	return *o.ProductName
}

// GetProductNameOk returns a tuple with the ProductName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookErrorProductEditDataFailedRowsDataInner) GetProductNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProductName) {
		return nil, false
	}
	return o.ProductName, true
}

// HasProductName returns a boolean if a field has been set.
func (o *WebhookErrorProductEditDataFailedRowsDataInner) HasProductName() bool {
	if o != nil && !IsNil(o.ProductName) {
		return true
	}

	return false
}

// SetProductName gets a reference to the given string and assigns it to the ProductName field.
func (o *WebhookErrorProductEditDataFailedRowsDataInner) SetProductName(v string) {
	o.ProductName = &v
}

// GetProductPrice returns the ProductPrice field value if set, zero value otherwise.
func (o *WebhookErrorProductEditDataFailedRowsDataInner) GetProductPrice() int64 {
	if o == nil || IsNil(o.ProductPrice) {
		var ret int64
		return ret
	}
	return *o.ProductPrice
}

// GetProductPriceOk returns a tuple with the ProductPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookErrorProductEditDataFailedRowsDataInner) GetProductPriceOk() (*int64, bool) {
	if o == nil || IsNil(o.ProductPrice) {
		return nil, false
	}
	return o.ProductPrice, true
}

// HasProductPrice returns a boolean if a field has been set.
func (o *WebhookErrorProductEditDataFailedRowsDataInner) HasProductPrice() bool {
	if o != nil && !IsNil(o.ProductPrice) {
		return true
	}

	return false
}

// SetProductPrice gets a reference to the given int64 and assigns it to the ProductPrice field.
func (o *WebhookErrorProductEditDataFailedRowsDataInner) SetProductPrice(v int64) {
	o.ProductPrice = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *WebhookErrorProductEditDataFailedRowsDataInner) GetSku() string {
	if o == nil || IsNil(o.Sku) {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookErrorProductEditDataFailedRowsDataInner) GetSkuOk() (*string, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *WebhookErrorProductEditDataFailedRowsDataInner) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *WebhookErrorProductEditDataFailedRowsDataInner) SetSku(v string) {
	o.Sku = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *WebhookErrorProductEditDataFailedRowsDataInner) GetError() []string {
	if o == nil || IsNil(o.Error) {
		var ret []string
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookErrorProductEditDataFailedRowsDataInner) GetErrorOk() ([]string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *WebhookErrorProductEditDataFailedRowsDataInner) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given []string and assigns it to the Error field.
func (o *WebhookErrorProductEditDataFailedRowsDataInner) SetError(v []string) {
	o.Error = v
}

func (o WebhookErrorProductEditDataFailedRowsDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookErrorProductEditDataFailedRowsDataInner) ToMap() (map[string]interface{}, error) {
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

type NullableWebhookErrorProductEditDataFailedRowsDataInner struct {
	value *WebhookErrorProductEditDataFailedRowsDataInner
	isSet bool
}

func (v NullableWebhookErrorProductEditDataFailedRowsDataInner) Get() *WebhookErrorProductEditDataFailedRowsDataInner {
	return v.value
}

func (v *NullableWebhookErrorProductEditDataFailedRowsDataInner) Set(val *WebhookErrorProductEditDataFailedRowsDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookErrorProductEditDataFailedRowsDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookErrorProductEditDataFailedRowsDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookErrorProductEditDataFailedRowsDataInner(val *WebhookErrorProductEditDataFailedRowsDataInner) *NullableWebhookErrorProductEditDataFailedRowsDataInner {
	return &NullableWebhookErrorProductEditDataFailedRowsDataInner{value: val, isSet: true}
}

func (v NullableWebhookErrorProductEditDataFailedRowsDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookErrorProductEditDataFailedRowsDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


