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

// checks if the ConfirmShippingRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfirmShippingRequest{}

// ConfirmShippingRequest struct for ConfirmShippingRequest
type ConfirmShippingRequest struct {
	// Order status code: * 0 - Seller cancel order. * 3 - Order Reject Due Empty Stock. * 5 - Order Canceled by Fraud * 6 - Order Rejected (Auto Cancel Out of Stock) * 10 - Order rejected by seller. * 15 - Instant Cancel by Buyer. * 100 - Order Created. * 103 - Wait for payment confirmation from third party. * 220 - Payment verified, order ready to process. * 221 - Waiting for partner approval. * 400 - Seller accept order. * 450 - Waiting for pickup. * 500 - Order shipment. * 501 - Status changed to waiting resi have no input. * 520 - Invalid shipment reference number (AWB). * 530 - Requested by user to correct invalid entry of shipment reference number. * 540 - Delivered to Pickup Point. * 550 - Return to Seller. * 600 - Order delivered. * 601 - Buyer open a case to finish an order. * 690 - Fraud Review * 700 - Order finished. 
	OrderStatus int64 `json:"order_status"`
	ShippingRefNum string `json:"shipping_ref_num"`
}

// NewConfirmShippingRequest instantiates a new ConfirmShippingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfirmShippingRequest(orderStatus int64, shippingRefNum string) *ConfirmShippingRequest {
	this := ConfirmShippingRequest{}
	this.OrderStatus = orderStatus
	this.ShippingRefNum = shippingRefNum
	return &this
}

// NewConfirmShippingRequestWithDefaults instantiates a new ConfirmShippingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfirmShippingRequestWithDefaults() *ConfirmShippingRequest {
	this := ConfirmShippingRequest{}
	return &this
}

// GetOrderStatus returns the OrderStatus field value
func (o *ConfirmShippingRequest) GetOrderStatus() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.OrderStatus
}

// GetOrderStatusOk returns a tuple with the OrderStatus field value
// and a boolean to check if the value has been set.
func (o *ConfirmShippingRequest) GetOrderStatusOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderStatus, true
}

// SetOrderStatus sets field value
func (o *ConfirmShippingRequest) SetOrderStatus(v int64) {
	o.OrderStatus = v
}

// GetShippingRefNum returns the ShippingRefNum field value
func (o *ConfirmShippingRequest) GetShippingRefNum() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShippingRefNum
}

// GetShippingRefNumOk returns a tuple with the ShippingRefNum field value
// and a boolean to check if the value has been set.
func (o *ConfirmShippingRequest) GetShippingRefNumOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShippingRefNum, true
}

// SetShippingRefNum sets field value
func (o *ConfirmShippingRequest) SetShippingRefNum(v string) {
	o.ShippingRefNum = v
}

func (o ConfirmShippingRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfirmShippingRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["order_status"] = o.OrderStatus
	toSerialize["shipping_ref_num"] = o.ShippingRefNum
	return toSerialize, nil
}

type NullableConfirmShippingRequest struct {
	value *ConfirmShippingRequest
	isSet bool
}

func (v NullableConfirmShippingRequest) Get() *ConfirmShippingRequest {
	return v.value
}

func (v *NullableConfirmShippingRequest) Set(val *ConfirmShippingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableConfirmShippingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableConfirmShippingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfirmShippingRequest(val *ConfirmShippingRequest) *NullableConfirmShippingRequest {
	return &NullableConfirmShippingRequest{value: val, isSet: true}
}

func (v NullableConfirmShippingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfirmShippingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

