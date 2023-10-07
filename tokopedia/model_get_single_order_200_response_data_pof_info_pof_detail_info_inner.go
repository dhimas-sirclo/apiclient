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

// checks if the GetSingleOrder200ResponseDataPofInfoPofDetailInfoInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSingleOrder200ResponseDataPofInfoPofDetailInfoInner{}

// GetSingleOrder200ResponseDataPofInfoPofDetailInfoInner struct for GetSingleOrder200ResponseDataPofInfoPofDetailInfoInner
type GetSingleOrder200ResponseDataPofInfoPofDetailInfoInner struct {
	// Order Unique Identifier
	OrderId *int64 `json:"order_id,omitempty"`
	// Order Detail Unique Identifier
	OrderDtlId *int64 `json:"order_dtl_id,omitempty"`
	// Product Unique Identifier
	ProductId *int64 `json:"product_id,omitempty"`
	// Requested Product Quantity
	QuantityRequest *int64 `json:"quantity_request,omitempty"`
	// Requested Total Price
	TotalPriceRequest *int64 `json:"total_price_request,omitempty"`
}

// NewGetSingleOrder200ResponseDataPofInfoPofDetailInfoInner instantiates a new GetSingleOrder200ResponseDataPofInfoPofDetailInfoInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSingleOrder200ResponseDataPofInfoPofDetailInfoInner() *GetSingleOrder200ResponseDataPofInfoPofDetailInfoInner {
	this := GetSingleOrder200ResponseDataPofInfoPofDetailInfoInner{}
	return &this
}

// NewGetSingleOrder200ResponseDataPofInfoPofDetailInfoInnerWithDefaults instantiates a new GetSingleOrder200ResponseDataPofInfoPofDetailInfoInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSingleOrder200ResponseDataPofInfoPofDetailInfoInnerWithDefaults() *GetSingleOrder200ResponseDataPofInfoPofDetailInfoInner {
	this := GetSingleOrder200ResponseDataPofInfoPofDetailInfoInner{}
	return &this
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataPofInfoPofDetailInfoInner) GetOrderId() int64 {
	if o == nil || IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataPofInfoPofDetailInfoInner) GetOrderIdOk() (*int64, bool) {
	if o == nil || IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataPofInfoPofDetailInfoInner) HasOrderId() bool {
	if o != nil && !IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *GetSingleOrder200ResponseDataPofInfoPofDetailInfoInner) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetOrderDtlId returns the OrderDtlId field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataPofInfoPofDetailInfoInner) GetOrderDtlId() int64 {
	if o == nil || IsNil(o.OrderDtlId) {
		var ret int64
		return ret
	}
	return *o.OrderDtlId
}

// GetOrderDtlIdOk returns a tuple with the OrderDtlId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataPofInfoPofDetailInfoInner) GetOrderDtlIdOk() (*int64, bool) {
	if o == nil || IsNil(o.OrderDtlId) {
		return nil, false
	}
	return o.OrderDtlId, true
}

// HasOrderDtlId returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataPofInfoPofDetailInfoInner) HasOrderDtlId() bool {
	if o != nil && !IsNil(o.OrderDtlId) {
		return true
	}

	return false
}

// SetOrderDtlId gets a reference to the given int64 and assigns it to the OrderDtlId field.
func (o *GetSingleOrder200ResponseDataPofInfoPofDetailInfoInner) SetOrderDtlId(v int64) {
	o.OrderDtlId = &v
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataPofInfoPofDetailInfoInner) GetProductId() int64 {
	if o == nil || IsNil(o.ProductId) {
		var ret int64
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataPofInfoPofDetailInfoInner) GetProductIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ProductId) {
		return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataPofInfoPofDetailInfoInner) HasProductId() bool {
	if o != nil && !IsNil(o.ProductId) {
		return true
	}

	return false
}

// SetProductId gets a reference to the given int64 and assigns it to the ProductId field.
func (o *GetSingleOrder200ResponseDataPofInfoPofDetailInfoInner) SetProductId(v int64) {
	o.ProductId = &v
}

// GetQuantityRequest returns the QuantityRequest field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataPofInfoPofDetailInfoInner) GetQuantityRequest() int64 {
	if o == nil || IsNil(o.QuantityRequest) {
		var ret int64
		return ret
	}
	return *o.QuantityRequest
}

// GetQuantityRequestOk returns a tuple with the QuantityRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataPofInfoPofDetailInfoInner) GetQuantityRequestOk() (*int64, bool) {
	if o == nil || IsNil(o.QuantityRequest) {
		return nil, false
	}
	return o.QuantityRequest, true
}

// HasQuantityRequest returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataPofInfoPofDetailInfoInner) HasQuantityRequest() bool {
	if o != nil && !IsNil(o.QuantityRequest) {
		return true
	}

	return false
}

// SetQuantityRequest gets a reference to the given int64 and assigns it to the QuantityRequest field.
func (o *GetSingleOrder200ResponseDataPofInfoPofDetailInfoInner) SetQuantityRequest(v int64) {
	o.QuantityRequest = &v
}

// GetTotalPriceRequest returns the TotalPriceRequest field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataPofInfoPofDetailInfoInner) GetTotalPriceRequest() int64 {
	if o == nil || IsNil(o.TotalPriceRequest) {
		var ret int64
		return ret
	}
	return *o.TotalPriceRequest
}

// GetTotalPriceRequestOk returns a tuple with the TotalPriceRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataPofInfoPofDetailInfoInner) GetTotalPriceRequestOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalPriceRequest) {
		return nil, false
	}
	return o.TotalPriceRequest, true
}

// HasTotalPriceRequest returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataPofInfoPofDetailInfoInner) HasTotalPriceRequest() bool {
	if o != nil && !IsNil(o.TotalPriceRequest) {
		return true
	}

	return false
}

// SetTotalPriceRequest gets a reference to the given int64 and assigns it to the TotalPriceRequest field.
func (o *GetSingleOrder200ResponseDataPofInfoPofDetailInfoInner) SetTotalPriceRequest(v int64) {
	o.TotalPriceRequest = &v
}

func (o GetSingleOrder200ResponseDataPofInfoPofDetailInfoInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSingleOrder200ResponseDataPofInfoPofDetailInfoInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OrderId) {
		toSerialize["order_id"] = o.OrderId
	}
	if !IsNil(o.OrderDtlId) {
		toSerialize["order_dtl_id"] = o.OrderDtlId
	}
	if !IsNil(o.ProductId) {
		toSerialize["product_id"] = o.ProductId
	}
	if !IsNil(o.QuantityRequest) {
		toSerialize["quantity_request"] = o.QuantityRequest
	}
	if !IsNil(o.TotalPriceRequest) {
		toSerialize["total_price_request"] = o.TotalPriceRequest
	}
	return toSerialize, nil
}

type NullableGetSingleOrder200ResponseDataPofInfoPofDetailInfoInner struct {
	value *GetSingleOrder200ResponseDataPofInfoPofDetailInfoInner
	isSet bool
}

func (v NullableGetSingleOrder200ResponseDataPofInfoPofDetailInfoInner) Get() *GetSingleOrder200ResponseDataPofInfoPofDetailInfoInner {
	return v.value
}

func (v *NullableGetSingleOrder200ResponseDataPofInfoPofDetailInfoInner) Set(val *GetSingleOrder200ResponseDataPofInfoPofDetailInfoInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSingleOrder200ResponseDataPofInfoPofDetailInfoInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSingleOrder200ResponseDataPofInfoPofDetailInfoInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSingleOrder200ResponseDataPofInfoPofDetailInfoInner(val *GetSingleOrder200ResponseDataPofInfoPofDetailInfoInner) *NullableGetSingleOrder200ResponseDataPofInfoPofDetailInfoInner {
	return &NullableGetSingleOrder200ResponseDataPofInfoPofDetailInfoInner{value: val, isSet: true}
}

func (v NullableGetSingleOrder200ResponseDataPofInfoPofDetailInfoInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSingleOrder200ResponseDataPofInfoPofDetailInfoInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

