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

// checks if the GetSingleOrder200ResponseDataPromoOrderDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSingleOrder200ResponseDataPromoOrderDetail{}

// GetSingleOrder200ResponseDataPromoOrderDetail struct for GetSingleOrder200ResponseDataPromoOrderDetail
type GetSingleOrder200ResponseDataPromoOrderDetail struct {
	OrderId *int64 `json:"order_id,omitempty"`
	TotalCashback *int64 `json:"total_cashback,omitempty"`
	TotalDiscount *int64 `json:"total_discount,omitempty"`
	TotalDiscountProduct *int64 `json:"total_discount_product,omitempty"`
	TotalDiscountShipping *int64 `json:"total_discount_shipping,omitempty"`
	TotalDiscountDetails []GetSingleOrder200ResponseDataPromoOrderDetailTotalDiscountDetailsInner `json:"total_discount_details,omitempty"`
	SummaryPromo []GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner `json:"summary_promo,omitempty"`
}

// NewGetSingleOrder200ResponseDataPromoOrderDetail instantiates a new GetSingleOrder200ResponseDataPromoOrderDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSingleOrder200ResponseDataPromoOrderDetail() *GetSingleOrder200ResponseDataPromoOrderDetail {
	this := GetSingleOrder200ResponseDataPromoOrderDetail{}
	return &this
}

// NewGetSingleOrder200ResponseDataPromoOrderDetailWithDefaults instantiates a new GetSingleOrder200ResponseDataPromoOrderDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSingleOrder200ResponseDataPromoOrderDetailWithDefaults() *GetSingleOrder200ResponseDataPromoOrderDetail {
	this := GetSingleOrder200ResponseDataPromoOrderDetail{}
	return &this
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataPromoOrderDetail) GetOrderId() int64 {
	if o == nil || IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataPromoOrderDetail) GetOrderIdOk() (*int64, bool) {
	if o == nil || IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataPromoOrderDetail) HasOrderId() bool {
	if o != nil && !IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *GetSingleOrder200ResponseDataPromoOrderDetail) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetTotalCashback returns the TotalCashback field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataPromoOrderDetail) GetTotalCashback() int64 {
	if o == nil || IsNil(o.TotalCashback) {
		var ret int64
		return ret
	}
	return *o.TotalCashback
}

// GetTotalCashbackOk returns a tuple with the TotalCashback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataPromoOrderDetail) GetTotalCashbackOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalCashback) {
		return nil, false
	}
	return o.TotalCashback, true
}

// HasTotalCashback returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataPromoOrderDetail) HasTotalCashback() bool {
	if o != nil && !IsNil(o.TotalCashback) {
		return true
	}

	return false
}

// SetTotalCashback gets a reference to the given int64 and assigns it to the TotalCashback field.
func (o *GetSingleOrder200ResponseDataPromoOrderDetail) SetTotalCashback(v int64) {
	o.TotalCashback = &v
}

// GetTotalDiscount returns the TotalDiscount field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataPromoOrderDetail) GetTotalDiscount() int64 {
	if o == nil || IsNil(o.TotalDiscount) {
		var ret int64
		return ret
	}
	return *o.TotalDiscount
}

// GetTotalDiscountOk returns a tuple with the TotalDiscount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataPromoOrderDetail) GetTotalDiscountOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalDiscount) {
		return nil, false
	}
	return o.TotalDiscount, true
}

// HasTotalDiscount returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataPromoOrderDetail) HasTotalDiscount() bool {
	if o != nil && !IsNil(o.TotalDiscount) {
		return true
	}

	return false
}

// SetTotalDiscount gets a reference to the given int64 and assigns it to the TotalDiscount field.
func (o *GetSingleOrder200ResponseDataPromoOrderDetail) SetTotalDiscount(v int64) {
	o.TotalDiscount = &v
}

// GetTotalDiscountProduct returns the TotalDiscountProduct field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataPromoOrderDetail) GetTotalDiscountProduct() int64 {
	if o == nil || IsNil(o.TotalDiscountProduct) {
		var ret int64
		return ret
	}
	return *o.TotalDiscountProduct
}

// GetTotalDiscountProductOk returns a tuple with the TotalDiscountProduct field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataPromoOrderDetail) GetTotalDiscountProductOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalDiscountProduct) {
		return nil, false
	}
	return o.TotalDiscountProduct, true
}

// HasTotalDiscountProduct returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataPromoOrderDetail) HasTotalDiscountProduct() bool {
	if o != nil && !IsNil(o.TotalDiscountProduct) {
		return true
	}

	return false
}

// SetTotalDiscountProduct gets a reference to the given int64 and assigns it to the TotalDiscountProduct field.
func (o *GetSingleOrder200ResponseDataPromoOrderDetail) SetTotalDiscountProduct(v int64) {
	o.TotalDiscountProduct = &v
}

// GetTotalDiscountShipping returns the TotalDiscountShipping field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataPromoOrderDetail) GetTotalDiscountShipping() int64 {
	if o == nil || IsNil(o.TotalDiscountShipping) {
		var ret int64
		return ret
	}
	return *o.TotalDiscountShipping
}

// GetTotalDiscountShippingOk returns a tuple with the TotalDiscountShipping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataPromoOrderDetail) GetTotalDiscountShippingOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalDiscountShipping) {
		return nil, false
	}
	return o.TotalDiscountShipping, true
}

// HasTotalDiscountShipping returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataPromoOrderDetail) HasTotalDiscountShipping() bool {
	if o != nil && !IsNil(o.TotalDiscountShipping) {
		return true
	}

	return false
}

// SetTotalDiscountShipping gets a reference to the given int64 and assigns it to the TotalDiscountShipping field.
func (o *GetSingleOrder200ResponseDataPromoOrderDetail) SetTotalDiscountShipping(v int64) {
	o.TotalDiscountShipping = &v
}

// GetTotalDiscountDetails returns the TotalDiscountDetails field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataPromoOrderDetail) GetTotalDiscountDetails() []GetSingleOrder200ResponseDataPromoOrderDetailTotalDiscountDetailsInner {
	if o == nil || IsNil(o.TotalDiscountDetails) {
		var ret []GetSingleOrder200ResponseDataPromoOrderDetailTotalDiscountDetailsInner
		return ret
	}
	return o.TotalDiscountDetails
}

// GetTotalDiscountDetailsOk returns a tuple with the TotalDiscountDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataPromoOrderDetail) GetTotalDiscountDetailsOk() ([]GetSingleOrder200ResponseDataPromoOrderDetailTotalDiscountDetailsInner, bool) {
	if o == nil || IsNil(o.TotalDiscountDetails) {
		return nil, false
	}
	return o.TotalDiscountDetails, true
}

// HasTotalDiscountDetails returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataPromoOrderDetail) HasTotalDiscountDetails() bool {
	if o != nil && !IsNil(o.TotalDiscountDetails) {
		return true
	}

	return false
}

// SetTotalDiscountDetails gets a reference to the given []GetSingleOrder200ResponseDataPromoOrderDetailTotalDiscountDetailsInner and assigns it to the TotalDiscountDetails field.
func (o *GetSingleOrder200ResponseDataPromoOrderDetail) SetTotalDiscountDetails(v []GetSingleOrder200ResponseDataPromoOrderDetailTotalDiscountDetailsInner) {
	o.TotalDiscountDetails = v
}

// GetSummaryPromo returns the SummaryPromo field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataPromoOrderDetail) GetSummaryPromo() []GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner {
	if o == nil || IsNil(o.SummaryPromo) {
		var ret []GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner
		return ret
	}
	return o.SummaryPromo
}

// GetSummaryPromoOk returns a tuple with the SummaryPromo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataPromoOrderDetail) GetSummaryPromoOk() ([]GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner, bool) {
	if o == nil || IsNil(o.SummaryPromo) {
		return nil, false
	}
	return o.SummaryPromo, true
}

// HasSummaryPromo returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataPromoOrderDetail) HasSummaryPromo() bool {
	if o != nil && !IsNil(o.SummaryPromo) {
		return true
	}

	return false
}

// SetSummaryPromo gets a reference to the given []GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner and assigns it to the SummaryPromo field.
func (o *GetSingleOrder200ResponseDataPromoOrderDetail) SetSummaryPromo(v []GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) {
	o.SummaryPromo = v
}

func (o GetSingleOrder200ResponseDataPromoOrderDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSingleOrder200ResponseDataPromoOrderDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OrderId) {
		toSerialize["order_id"] = o.OrderId
	}
	if !IsNil(o.TotalCashback) {
		toSerialize["total_cashback"] = o.TotalCashback
	}
	if !IsNil(o.TotalDiscount) {
		toSerialize["total_discount"] = o.TotalDiscount
	}
	if !IsNil(o.TotalDiscountProduct) {
		toSerialize["total_discount_product"] = o.TotalDiscountProduct
	}
	if !IsNil(o.TotalDiscountShipping) {
		toSerialize["total_discount_shipping"] = o.TotalDiscountShipping
	}
	if !IsNil(o.TotalDiscountDetails) {
		toSerialize["total_discount_details"] = o.TotalDiscountDetails
	}
	if !IsNil(o.SummaryPromo) {
		toSerialize["summary_promo"] = o.SummaryPromo
	}
	return toSerialize, nil
}

type NullableGetSingleOrder200ResponseDataPromoOrderDetail struct {
	value *GetSingleOrder200ResponseDataPromoOrderDetail
	isSet bool
}

func (v NullableGetSingleOrder200ResponseDataPromoOrderDetail) Get() *GetSingleOrder200ResponseDataPromoOrderDetail {
	return v.value
}

func (v *NullableGetSingleOrder200ResponseDataPromoOrderDetail) Set(val *GetSingleOrder200ResponseDataPromoOrderDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSingleOrder200ResponseDataPromoOrderDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSingleOrder200ResponseDataPromoOrderDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSingleOrder200ResponseDataPromoOrderDetail(val *GetSingleOrder200ResponseDataPromoOrderDetail) *NullableGetSingleOrder200ResponseDataPromoOrderDetail {
	return &NullableGetSingleOrder200ResponseDataPromoOrderDetail{value: val, isSet: true}
}

func (v NullableGetSingleOrder200ResponseDataPromoOrderDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSingleOrder200ResponseDataPromoOrderDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

