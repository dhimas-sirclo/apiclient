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

// checks if the WebhookOrderStatusCancellation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookOrderStatusCancellation{}

// WebhookOrderStatusCancellation struct for WebhookOrderStatusCancellation
type WebhookOrderStatusCancellation struct {
	// Order Status
	OrderStatus *int64 `json:"order_status,omitempty"`
	// Fulfillment service unique identifier
	FsId *int64 `json:"fs_id,omitempty"`
	// Shop unique identifier
	ShopId *int64 `json:"shop_id,omitempty"`
	// Order unique identifier
	OrderId *int64 `json:"order_id,omitempty"`
	// List of products
	ProductDetails []WebhookOrderStatusCancellationProductDetailsInner `json:"product_details,omitempty"`
}

// NewWebhookOrderStatusCancellation instantiates a new WebhookOrderStatusCancellation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookOrderStatusCancellation() *WebhookOrderStatusCancellation {
	this := WebhookOrderStatusCancellation{}
	return &this
}

// NewWebhookOrderStatusCancellationWithDefaults instantiates a new WebhookOrderStatusCancellation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookOrderStatusCancellationWithDefaults() *WebhookOrderStatusCancellation {
	this := WebhookOrderStatusCancellation{}
	return &this
}

// GetOrderStatus returns the OrderStatus field value if set, zero value otherwise.
func (o *WebhookOrderStatusCancellation) GetOrderStatus() int64 {
	if o == nil || IsNil(o.OrderStatus) {
		var ret int64
		return ret
	}
	return *o.OrderStatus
}

// GetOrderStatusOk returns a tuple with the OrderStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookOrderStatusCancellation) GetOrderStatusOk() (*int64, bool) {
	if o == nil || IsNil(o.OrderStatus) {
		return nil, false
	}
	return o.OrderStatus, true
}

// HasOrderStatus returns a boolean if a field has been set.
func (o *WebhookOrderStatusCancellation) HasOrderStatus() bool {
	if o != nil && !IsNil(o.OrderStatus) {
		return true
	}

	return false
}

// SetOrderStatus gets a reference to the given int64 and assigns it to the OrderStatus field.
func (o *WebhookOrderStatusCancellation) SetOrderStatus(v int64) {
	o.OrderStatus = &v
}

// GetFsId returns the FsId field value if set, zero value otherwise.
func (o *WebhookOrderStatusCancellation) GetFsId() int64 {
	if o == nil || IsNil(o.FsId) {
		var ret int64
		return ret
	}
	return *o.FsId
}

// GetFsIdOk returns a tuple with the FsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookOrderStatusCancellation) GetFsIdOk() (*int64, bool) {
	if o == nil || IsNil(o.FsId) {
		return nil, false
	}
	return o.FsId, true
}

// HasFsId returns a boolean if a field has been set.
func (o *WebhookOrderStatusCancellation) HasFsId() bool {
	if o != nil && !IsNil(o.FsId) {
		return true
	}

	return false
}

// SetFsId gets a reference to the given int64 and assigns it to the FsId field.
func (o *WebhookOrderStatusCancellation) SetFsId(v int64) {
	o.FsId = &v
}

// GetShopId returns the ShopId field value if set, zero value otherwise.
func (o *WebhookOrderStatusCancellation) GetShopId() int64 {
	if o == nil || IsNil(o.ShopId) {
		var ret int64
		return ret
	}
	return *o.ShopId
}

// GetShopIdOk returns a tuple with the ShopId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookOrderStatusCancellation) GetShopIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ShopId) {
		return nil, false
	}
	return o.ShopId, true
}

// HasShopId returns a boolean if a field has been set.
func (o *WebhookOrderStatusCancellation) HasShopId() bool {
	if o != nil && !IsNil(o.ShopId) {
		return true
	}

	return false
}

// SetShopId gets a reference to the given int64 and assigns it to the ShopId field.
func (o *WebhookOrderStatusCancellation) SetShopId(v int64) {
	o.ShopId = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *WebhookOrderStatusCancellation) GetOrderId() int64 {
	if o == nil || IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookOrderStatusCancellation) GetOrderIdOk() (*int64, bool) {
	if o == nil || IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *WebhookOrderStatusCancellation) HasOrderId() bool {
	if o != nil && !IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *WebhookOrderStatusCancellation) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetProductDetails returns the ProductDetails field value if set, zero value otherwise.
func (o *WebhookOrderStatusCancellation) GetProductDetails() []WebhookOrderStatusCancellationProductDetailsInner {
	if o == nil || IsNil(o.ProductDetails) {
		var ret []WebhookOrderStatusCancellationProductDetailsInner
		return ret
	}
	return o.ProductDetails
}

// GetProductDetailsOk returns a tuple with the ProductDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookOrderStatusCancellation) GetProductDetailsOk() ([]WebhookOrderStatusCancellationProductDetailsInner, bool) {
	if o == nil || IsNil(o.ProductDetails) {
		return nil, false
	}
	return o.ProductDetails, true
}

// HasProductDetails returns a boolean if a field has been set.
func (o *WebhookOrderStatusCancellation) HasProductDetails() bool {
	if o != nil && !IsNil(o.ProductDetails) {
		return true
	}

	return false
}

// SetProductDetails gets a reference to the given []WebhookOrderStatusCancellationProductDetailsInner and assigns it to the ProductDetails field.
func (o *WebhookOrderStatusCancellation) SetProductDetails(v []WebhookOrderStatusCancellationProductDetailsInner) {
	o.ProductDetails = v
}

func (o WebhookOrderStatusCancellation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookOrderStatusCancellation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OrderStatus) {
		toSerialize["order_status"] = o.OrderStatus
	}
	if !IsNil(o.FsId) {
		toSerialize["fs_id"] = o.FsId
	}
	if !IsNil(o.ShopId) {
		toSerialize["shop_id"] = o.ShopId
	}
	if !IsNil(o.OrderId) {
		toSerialize["order_id"] = o.OrderId
	}
	if !IsNil(o.ProductDetails) {
		toSerialize["product_details"] = o.ProductDetails
	}
	return toSerialize, nil
}

type NullableWebhookOrderStatusCancellation struct {
	value *WebhookOrderStatusCancellation
	isSet bool
}

func (v NullableWebhookOrderStatusCancellation) Get() *WebhookOrderStatusCancellation {
	return v.value
}

func (v *NullableWebhookOrderStatusCancellation) Set(val *WebhookOrderStatusCancellation) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookOrderStatusCancellation) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookOrderStatusCancellation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookOrderStatusCancellation(val *WebhookOrderStatusCancellation) *NullableWebhookOrderStatusCancellation {
	return &NullableWebhookOrderStatusCancellation{value: val, isSet: true}
}

func (v NullableWebhookOrderStatusCancellation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookOrderStatusCancellation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


