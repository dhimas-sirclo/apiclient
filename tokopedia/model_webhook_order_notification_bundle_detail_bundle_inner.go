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

// checks if the WebhookOrderNotificationBundleDetailBundleInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookOrderNotificationBundleDetailBundleInner{}

// WebhookOrderNotificationBundleDetailBundleInner struct for WebhookOrderNotificationBundleDetailBundleInner
type WebhookOrderNotificationBundleDetailBundleInner struct {
	BundleId *int64 `json:"bundle_id,omitempty"`
	BundleVariantId *string `json:"bundle_variant_id,omitempty"`
	BundleName *string `json:"bundle_name,omitempty"`
	BundlePrice *int64 `json:"bundle_price,omitempty"`
	BundleQuantity *int64 `json:"bundle_quantity,omitempty"`
	BundleSubtotalPrice *int64 `json:"bundle_subtotal_price,omitempty"`
	OrderDetail []WebhookOrderNotificationBundleDetailBundleInnerOrderDetailInner `json:"order_detail,omitempty"`
}

// NewWebhookOrderNotificationBundleDetailBundleInner instantiates a new WebhookOrderNotificationBundleDetailBundleInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookOrderNotificationBundleDetailBundleInner() *WebhookOrderNotificationBundleDetailBundleInner {
	this := WebhookOrderNotificationBundleDetailBundleInner{}
	return &this
}

// NewWebhookOrderNotificationBundleDetailBundleInnerWithDefaults instantiates a new WebhookOrderNotificationBundleDetailBundleInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookOrderNotificationBundleDetailBundleInnerWithDefaults() *WebhookOrderNotificationBundleDetailBundleInner {
	this := WebhookOrderNotificationBundleDetailBundleInner{}
	return &this
}

// GetBundleId returns the BundleId field value if set, zero value otherwise.
func (o *WebhookOrderNotificationBundleDetailBundleInner) GetBundleId() int64 {
	if o == nil || IsNil(o.BundleId) {
		var ret int64
		return ret
	}
	return *o.BundleId
}

// GetBundleIdOk returns a tuple with the BundleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookOrderNotificationBundleDetailBundleInner) GetBundleIdOk() (*int64, bool) {
	if o == nil || IsNil(o.BundleId) {
		return nil, false
	}
	return o.BundleId, true
}

// HasBundleId returns a boolean if a field has been set.
func (o *WebhookOrderNotificationBundleDetailBundleInner) HasBundleId() bool {
	if o != nil && !IsNil(o.BundleId) {
		return true
	}

	return false
}

// SetBundleId gets a reference to the given int64 and assigns it to the BundleId field.
func (o *WebhookOrderNotificationBundleDetailBundleInner) SetBundleId(v int64) {
	o.BundleId = &v
}

// GetBundleVariantId returns the BundleVariantId field value if set, zero value otherwise.
func (o *WebhookOrderNotificationBundleDetailBundleInner) GetBundleVariantId() string {
	if o == nil || IsNil(o.BundleVariantId) {
		var ret string
		return ret
	}
	return *o.BundleVariantId
}

// GetBundleVariantIdOk returns a tuple with the BundleVariantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookOrderNotificationBundleDetailBundleInner) GetBundleVariantIdOk() (*string, bool) {
	if o == nil || IsNil(o.BundleVariantId) {
		return nil, false
	}
	return o.BundleVariantId, true
}

// HasBundleVariantId returns a boolean if a field has been set.
func (o *WebhookOrderNotificationBundleDetailBundleInner) HasBundleVariantId() bool {
	if o != nil && !IsNil(o.BundleVariantId) {
		return true
	}

	return false
}

// SetBundleVariantId gets a reference to the given string and assigns it to the BundleVariantId field.
func (o *WebhookOrderNotificationBundleDetailBundleInner) SetBundleVariantId(v string) {
	o.BundleVariantId = &v
}

// GetBundleName returns the BundleName field value if set, zero value otherwise.
func (o *WebhookOrderNotificationBundleDetailBundleInner) GetBundleName() string {
	if o == nil || IsNil(o.BundleName) {
		var ret string
		return ret
	}
	return *o.BundleName
}

// GetBundleNameOk returns a tuple with the BundleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookOrderNotificationBundleDetailBundleInner) GetBundleNameOk() (*string, bool) {
	if o == nil || IsNil(o.BundleName) {
		return nil, false
	}
	return o.BundleName, true
}

// HasBundleName returns a boolean if a field has been set.
func (o *WebhookOrderNotificationBundleDetailBundleInner) HasBundleName() bool {
	if o != nil && !IsNil(o.BundleName) {
		return true
	}

	return false
}

// SetBundleName gets a reference to the given string and assigns it to the BundleName field.
func (o *WebhookOrderNotificationBundleDetailBundleInner) SetBundleName(v string) {
	o.BundleName = &v
}

// GetBundlePrice returns the BundlePrice field value if set, zero value otherwise.
func (o *WebhookOrderNotificationBundleDetailBundleInner) GetBundlePrice() int64 {
	if o == nil || IsNil(o.BundlePrice) {
		var ret int64
		return ret
	}
	return *o.BundlePrice
}

// GetBundlePriceOk returns a tuple with the BundlePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookOrderNotificationBundleDetailBundleInner) GetBundlePriceOk() (*int64, bool) {
	if o == nil || IsNil(o.BundlePrice) {
		return nil, false
	}
	return o.BundlePrice, true
}

// HasBundlePrice returns a boolean if a field has been set.
func (o *WebhookOrderNotificationBundleDetailBundleInner) HasBundlePrice() bool {
	if o != nil && !IsNil(o.BundlePrice) {
		return true
	}

	return false
}

// SetBundlePrice gets a reference to the given int64 and assigns it to the BundlePrice field.
func (o *WebhookOrderNotificationBundleDetailBundleInner) SetBundlePrice(v int64) {
	o.BundlePrice = &v
}

// GetBundleQuantity returns the BundleQuantity field value if set, zero value otherwise.
func (o *WebhookOrderNotificationBundleDetailBundleInner) GetBundleQuantity() int64 {
	if o == nil || IsNil(o.BundleQuantity) {
		var ret int64
		return ret
	}
	return *o.BundleQuantity
}

// GetBundleQuantityOk returns a tuple with the BundleQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookOrderNotificationBundleDetailBundleInner) GetBundleQuantityOk() (*int64, bool) {
	if o == nil || IsNil(o.BundleQuantity) {
		return nil, false
	}
	return o.BundleQuantity, true
}

// HasBundleQuantity returns a boolean if a field has been set.
func (o *WebhookOrderNotificationBundleDetailBundleInner) HasBundleQuantity() bool {
	if o != nil && !IsNil(o.BundleQuantity) {
		return true
	}

	return false
}

// SetBundleQuantity gets a reference to the given int64 and assigns it to the BundleQuantity field.
func (o *WebhookOrderNotificationBundleDetailBundleInner) SetBundleQuantity(v int64) {
	o.BundleQuantity = &v
}

// GetBundleSubtotalPrice returns the BundleSubtotalPrice field value if set, zero value otherwise.
func (o *WebhookOrderNotificationBundleDetailBundleInner) GetBundleSubtotalPrice() int64 {
	if o == nil || IsNil(o.BundleSubtotalPrice) {
		var ret int64
		return ret
	}
	return *o.BundleSubtotalPrice
}

// GetBundleSubtotalPriceOk returns a tuple with the BundleSubtotalPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookOrderNotificationBundleDetailBundleInner) GetBundleSubtotalPriceOk() (*int64, bool) {
	if o == nil || IsNil(o.BundleSubtotalPrice) {
		return nil, false
	}
	return o.BundleSubtotalPrice, true
}

// HasBundleSubtotalPrice returns a boolean if a field has been set.
func (o *WebhookOrderNotificationBundleDetailBundleInner) HasBundleSubtotalPrice() bool {
	if o != nil && !IsNil(o.BundleSubtotalPrice) {
		return true
	}

	return false
}

// SetBundleSubtotalPrice gets a reference to the given int64 and assigns it to the BundleSubtotalPrice field.
func (o *WebhookOrderNotificationBundleDetailBundleInner) SetBundleSubtotalPrice(v int64) {
	o.BundleSubtotalPrice = &v
}

// GetOrderDetail returns the OrderDetail field value if set, zero value otherwise.
func (o *WebhookOrderNotificationBundleDetailBundleInner) GetOrderDetail() []WebhookOrderNotificationBundleDetailBundleInnerOrderDetailInner {
	if o == nil || IsNil(o.OrderDetail) {
		var ret []WebhookOrderNotificationBundleDetailBundleInnerOrderDetailInner
		return ret
	}
	return o.OrderDetail
}

// GetOrderDetailOk returns a tuple with the OrderDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookOrderNotificationBundleDetailBundleInner) GetOrderDetailOk() ([]WebhookOrderNotificationBundleDetailBundleInnerOrderDetailInner, bool) {
	if o == nil || IsNil(o.OrderDetail) {
		return nil, false
	}
	return o.OrderDetail, true
}

// HasOrderDetail returns a boolean if a field has been set.
func (o *WebhookOrderNotificationBundleDetailBundleInner) HasOrderDetail() bool {
	if o != nil && !IsNil(o.OrderDetail) {
		return true
	}

	return false
}

// SetOrderDetail gets a reference to the given []WebhookOrderNotificationBundleDetailBundleInnerOrderDetailInner and assigns it to the OrderDetail field.
func (o *WebhookOrderNotificationBundleDetailBundleInner) SetOrderDetail(v []WebhookOrderNotificationBundleDetailBundleInnerOrderDetailInner) {
	o.OrderDetail = v
}

func (o WebhookOrderNotificationBundleDetailBundleInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookOrderNotificationBundleDetailBundleInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BundleId) {
		toSerialize["bundle_id"] = o.BundleId
	}
	if !IsNil(o.BundleVariantId) {
		toSerialize["bundle_variant_id"] = o.BundleVariantId
	}
	if !IsNil(o.BundleName) {
		toSerialize["bundle_name"] = o.BundleName
	}
	if !IsNil(o.BundlePrice) {
		toSerialize["bundle_price"] = o.BundlePrice
	}
	if !IsNil(o.BundleQuantity) {
		toSerialize["bundle_quantity"] = o.BundleQuantity
	}
	if !IsNil(o.BundleSubtotalPrice) {
		toSerialize["bundle_subtotal_price"] = o.BundleSubtotalPrice
	}
	if !IsNil(o.OrderDetail) {
		toSerialize["order_detail"] = o.OrderDetail
	}
	return toSerialize, nil
}

type NullableWebhookOrderNotificationBundleDetailBundleInner struct {
	value *WebhookOrderNotificationBundleDetailBundleInner
	isSet bool
}

func (v NullableWebhookOrderNotificationBundleDetailBundleInner) Get() *WebhookOrderNotificationBundleDetailBundleInner {
	return v.value
}

func (v *NullableWebhookOrderNotificationBundleDetailBundleInner) Set(val *WebhookOrderNotificationBundleDetailBundleInner) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookOrderNotificationBundleDetailBundleInner) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookOrderNotificationBundleDetailBundleInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookOrderNotificationBundleDetailBundleInner(val *WebhookOrderNotificationBundleDetailBundleInner) *NullableWebhookOrderNotificationBundleDetailBundleInner {
	return &NullableWebhookOrderNotificationBundleDetailBundleInner{value: val, isSet: true}
}

func (v NullableWebhookOrderNotificationBundleDetailBundleInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookOrderNotificationBundleDetailBundleInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


