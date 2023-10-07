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

// checks if the RequestPartialOrderFulfillmentRequestPofDetailInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestPartialOrderFulfillmentRequestPofDetailInner{}

// RequestPartialOrderFulfillmentRequestPofDetailInner struct for RequestPartialOrderFulfillmentRequestPofDetailInner
type RequestPartialOrderFulfillmentRequestPofDetailInner struct {
	// Order detail unique identifier
	OrderDtlId *int64 `json:"order_dtl_id,omitempty"`
	// Requested quantity for POF request
	QuantityRequest *int64 `json:"quantity_request,omitempty"`
}

// NewRequestPartialOrderFulfillmentRequestPofDetailInner instantiates a new RequestPartialOrderFulfillmentRequestPofDetailInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestPartialOrderFulfillmentRequestPofDetailInner() *RequestPartialOrderFulfillmentRequestPofDetailInner {
	this := RequestPartialOrderFulfillmentRequestPofDetailInner{}
	return &this
}

// NewRequestPartialOrderFulfillmentRequestPofDetailInnerWithDefaults instantiates a new RequestPartialOrderFulfillmentRequestPofDetailInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestPartialOrderFulfillmentRequestPofDetailInnerWithDefaults() *RequestPartialOrderFulfillmentRequestPofDetailInner {
	this := RequestPartialOrderFulfillmentRequestPofDetailInner{}
	return &this
}

// GetOrderDtlId returns the OrderDtlId field value if set, zero value otherwise.
func (o *RequestPartialOrderFulfillmentRequestPofDetailInner) GetOrderDtlId() int64 {
	if o == nil || IsNil(o.OrderDtlId) {
		var ret int64
		return ret
	}
	return *o.OrderDtlId
}

// GetOrderDtlIdOk returns a tuple with the OrderDtlId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestPartialOrderFulfillmentRequestPofDetailInner) GetOrderDtlIdOk() (*int64, bool) {
	if o == nil || IsNil(o.OrderDtlId) {
		return nil, false
	}
	return o.OrderDtlId, true
}

// HasOrderDtlId returns a boolean if a field has been set.
func (o *RequestPartialOrderFulfillmentRequestPofDetailInner) HasOrderDtlId() bool {
	if o != nil && !IsNil(o.OrderDtlId) {
		return true
	}

	return false
}

// SetOrderDtlId gets a reference to the given int64 and assigns it to the OrderDtlId field.
func (o *RequestPartialOrderFulfillmentRequestPofDetailInner) SetOrderDtlId(v int64) {
	o.OrderDtlId = &v
}

// GetQuantityRequest returns the QuantityRequest field value if set, zero value otherwise.
func (o *RequestPartialOrderFulfillmentRequestPofDetailInner) GetQuantityRequest() int64 {
	if o == nil || IsNil(o.QuantityRequest) {
		var ret int64
		return ret
	}
	return *o.QuantityRequest
}

// GetQuantityRequestOk returns a tuple with the QuantityRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestPartialOrderFulfillmentRequestPofDetailInner) GetQuantityRequestOk() (*int64, bool) {
	if o == nil || IsNil(o.QuantityRequest) {
		return nil, false
	}
	return o.QuantityRequest, true
}

// HasQuantityRequest returns a boolean if a field has been set.
func (o *RequestPartialOrderFulfillmentRequestPofDetailInner) HasQuantityRequest() bool {
	if o != nil && !IsNil(o.QuantityRequest) {
		return true
	}

	return false
}

// SetQuantityRequest gets a reference to the given int64 and assigns it to the QuantityRequest field.
func (o *RequestPartialOrderFulfillmentRequestPofDetailInner) SetQuantityRequest(v int64) {
	o.QuantityRequest = &v
}

func (o RequestPartialOrderFulfillmentRequestPofDetailInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestPartialOrderFulfillmentRequestPofDetailInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OrderDtlId) {
		toSerialize["order_dtl_id"] = o.OrderDtlId
	}
	if !IsNil(o.QuantityRequest) {
		toSerialize["quantity_request"] = o.QuantityRequest
	}
	return toSerialize, nil
}

type NullableRequestPartialOrderFulfillmentRequestPofDetailInner struct {
	value *RequestPartialOrderFulfillmentRequestPofDetailInner
	isSet bool
}

func (v NullableRequestPartialOrderFulfillmentRequestPofDetailInner) Get() *RequestPartialOrderFulfillmentRequestPofDetailInner {
	return v.value
}

func (v *NullableRequestPartialOrderFulfillmentRequestPofDetailInner) Set(val *RequestPartialOrderFulfillmentRequestPofDetailInner) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestPartialOrderFulfillmentRequestPofDetailInner) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestPartialOrderFulfillmentRequestPofDetailInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestPartialOrderFulfillmentRequestPofDetailInner(val *RequestPartialOrderFulfillmentRequestPofDetailInner) *NullableRequestPartialOrderFulfillmentRequestPofDetailInner {
	return &NullableRequestPartialOrderFulfillmentRequestPofDetailInner{value: val, isSet: true}
}

func (v NullableRequestPartialOrderFulfillmentRequestPofDetailInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestPartialOrderFulfillmentRequestPofDetailInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

