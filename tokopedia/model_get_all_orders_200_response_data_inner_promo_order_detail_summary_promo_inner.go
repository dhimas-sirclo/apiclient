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

// checks if the GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner{}

// GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner struct for GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner
type GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner struct {
	// Promo Name
	Name *string `json:"Name,omitempty"`
	// Promo Name
	IsCoupon *bool `json:"is_coupon,omitempty"`
	// Show cashback amount?
	ShowCashbackAmount *bool `json:"show_cashback_amount,omitempty"`
	// Show discount amount?
	ShowDiscountAmount *bool `json:"show_discount_amount,omitempty"`
	// Cashback Amount
	CashbackAmount *int64 `json:"cashback_amount,omitempty"`
	// Cashback Point
	CashbackPoints *int64 `json:"cashback_points,omitempty"`
	// Promo Type
	Type *string `json:"Type,omitempty"`
	// Discount Amount
	DiscountAmount *int64 `json:"discount_amount,omitempty"`
	DiscountDetails []GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInnerDiscountDetailsInner `json:"discount_details,omitempty"`
	// Discount Invoice Description
	InvoiceDesc *string `json:"invoice_desc,omitempty"`
}

// NewGetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner instantiates a new GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner() *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner {
	this := GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner{}
	return &this
}

// NewGetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInnerWithDefaults instantiates a new GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInnerWithDefaults() *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner {
	this := GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) SetName(v string) {
	o.Name = &v
}

// GetIsCoupon returns the IsCoupon field value if set, zero value otherwise.
func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) GetIsCoupon() bool {
	if o == nil || IsNil(o.IsCoupon) {
		var ret bool
		return ret
	}
	return *o.IsCoupon
}

// GetIsCouponOk returns a tuple with the IsCoupon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) GetIsCouponOk() (*bool, bool) {
	if o == nil || IsNil(o.IsCoupon) {
		return nil, false
	}
	return o.IsCoupon, true
}

// HasIsCoupon returns a boolean if a field has been set.
func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) HasIsCoupon() bool {
	if o != nil && !IsNil(o.IsCoupon) {
		return true
	}

	return false
}

// SetIsCoupon gets a reference to the given bool and assigns it to the IsCoupon field.
func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) SetIsCoupon(v bool) {
	o.IsCoupon = &v
}

// GetShowCashbackAmount returns the ShowCashbackAmount field value if set, zero value otherwise.
func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) GetShowCashbackAmount() bool {
	if o == nil || IsNil(o.ShowCashbackAmount) {
		var ret bool
		return ret
	}
	return *o.ShowCashbackAmount
}

// GetShowCashbackAmountOk returns a tuple with the ShowCashbackAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) GetShowCashbackAmountOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowCashbackAmount) {
		return nil, false
	}
	return o.ShowCashbackAmount, true
}

// HasShowCashbackAmount returns a boolean if a field has been set.
func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) HasShowCashbackAmount() bool {
	if o != nil && !IsNil(o.ShowCashbackAmount) {
		return true
	}

	return false
}

// SetShowCashbackAmount gets a reference to the given bool and assigns it to the ShowCashbackAmount field.
func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) SetShowCashbackAmount(v bool) {
	o.ShowCashbackAmount = &v
}

// GetShowDiscountAmount returns the ShowDiscountAmount field value if set, zero value otherwise.
func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) GetShowDiscountAmount() bool {
	if o == nil || IsNil(o.ShowDiscountAmount) {
		var ret bool
		return ret
	}
	return *o.ShowDiscountAmount
}

// GetShowDiscountAmountOk returns a tuple with the ShowDiscountAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) GetShowDiscountAmountOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowDiscountAmount) {
		return nil, false
	}
	return o.ShowDiscountAmount, true
}

// HasShowDiscountAmount returns a boolean if a field has been set.
func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) HasShowDiscountAmount() bool {
	if o != nil && !IsNil(o.ShowDiscountAmount) {
		return true
	}

	return false
}

// SetShowDiscountAmount gets a reference to the given bool and assigns it to the ShowDiscountAmount field.
func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) SetShowDiscountAmount(v bool) {
	o.ShowDiscountAmount = &v
}

// GetCashbackAmount returns the CashbackAmount field value if set, zero value otherwise.
func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) GetCashbackAmount() int64 {
	if o == nil || IsNil(o.CashbackAmount) {
		var ret int64
		return ret
	}
	return *o.CashbackAmount
}

// GetCashbackAmountOk returns a tuple with the CashbackAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) GetCashbackAmountOk() (*int64, bool) {
	if o == nil || IsNil(o.CashbackAmount) {
		return nil, false
	}
	return o.CashbackAmount, true
}

// HasCashbackAmount returns a boolean if a field has been set.
func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) HasCashbackAmount() bool {
	if o != nil && !IsNil(o.CashbackAmount) {
		return true
	}

	return false
}

// SetCashbackAmount gets a reference to the given int64 and assigns it to the CashbackAmount field.
func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) SetCashbackAmount(v int64) {
	o.CashbackAmount = &v
}

// GetCashbackPoints returns the CashbackPoints field value if set, zero value otherwise.
func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) GetCashbackPoints() int64 {
	if o == nil || IsNil(o.CashbackPoints) {
		var ret int64
		return ret
	}
	return *o.CashbackPoints
}

// GetCashbackPointsOk returns a tuple with the CashbackPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) GetCashbackPointsOk() (*int64, bool) {
	if o == nil || IsNil(o.CashbackPoints) {
		return nil, false
	}
	return o.CashbackPoints, true
}

// HasCashbackPoints returns a boolean if a field has been set.
func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) HasCashbackPoints() bool {
	if o != nil && !IsNil(o.CashbackPoints) {
		return true
	}

	return false
}

// SetCashbackPoints gets a reference to the given int64 and assigns it to the CashbackPoints field.
func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) SetCashbackPoints(v int64) {
	o.CashbackPoints = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) SetType(v string) {
	o.Type = &v
}

// GetDiscountAmount returns the DiscountAmount field value if set, zero value otherwise.
func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) GetDiscountAmount() int64 {
	if o == nil || IsNil(o.DiscountAmount) {
		var ret int64
		return ret
	}
	return *o.DiscountAmount
}

// GetDiscountAmountOk returns a tuple with the DiscountAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) GetDiscountAmountOk() (*int64, bool) {
	if o == nil || IsNil(o.DiscountAmount) {
		return nil, false
	}
	return o.DiscountAmount, true
}

// HasDiscountAmount returns a boolean if a field has been set.
func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) HasDiscountAmount() bool {
	if o != nil && !IsNil(o.DiscountAmount) {
		return true
	}

	return false
}

// SetDiscountAmount gets a reference to the given int64 and assigns it to the DiscountAmount field.
func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) SetDiscountAmount(v int64) {
	o.DiscountAmount = &v
}

// GetDiscountDetails returns the DiscountDetails field value if set, zero value otherwise.
func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) GetDiscountDetails() []GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInnerDiscountDetailsInner {
	if o == nil || IsNil(o.DiscountDetails) {
		var ret []GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInnerDiscountDetailsInner
		return ret
	}
	return o.DiscountDetails
}

// GetDiscountDetailsOk returns a tuple with the DiscountDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) GetDiscountDetailsOk() ([]GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInnerDiscountDetailsInner, bool) {
	if o == nil || IsNil(o.DiscountDetails) {
		return nil, false
	}
	return o.DiscountDetails, true
}

// HasDiscountDetails returns a boolean if a field has been set.
func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) HasDiscountDetails() bool {
	if o != nil && !IsNil(o.DiscountDetails) {
		return true
	}

	return false
}

// SetDiscountDetails gets a reference to the given []GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInnerDiscountDetailsInner and assigns it to the DiscountDetails field.
func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) SetDiscountDetails(v []GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInnerDiscountDetailsInner) {
	o.DiscountDetails = v
}

// GetInvoiceDesc returns the InvoiceDesc field value if set, zero value otherwise.
func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) GetInvoiceDesc() string {
	if o == nil || IsNil(o.InvoiceDesc) {
		var ret string
		return ret
	}
	return *o.InvoiceDesc
}

// GetInvoiceDescOk returns a tuple with the InvoiceDesc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) GetInvoiceDescOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceDesc) {
		return nil, false
	}
	return o.InvoiceDesc, true
}

// HasInvoiceDesc returns a boolean if a field has been set.
func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) HasInvoiceDesc() bool {
	if o != nil && !IsNil(o.InvoiceDesc) {
		return true
	}

	return false
}

// SetInvoiceDesc gets a reference to the given string and assigns it to the InvoiceDesc field.
func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) SetInvoiceDesc(v string) {
	o.InvoiceDesc = &v
}

func (o GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.IsCoupon) {
		toSerialize["is_coupon"] = o.IsCoupon
	}
	if !IsNil(o.ShowCashbackAmount) {
		toSerialize["show_cashback_amount"] = o.ShowCashbackAmount
	}
	if !IsNil(o.ShowDiscountAmount) {
		toSerialize["show_discount_amount"] = o.ShowDiscountAmount
	}
	if !IsNil(o.CashbackAmount) {
		toSerialize["cashback_amount"] = o.CashbackAmount
	}
	if !IsNil(o.CashbackPoints) {
		toSerialize["cashback_points"] = o.CashbackPoints
	}
	if !IsNil(o.Type) {
		toSerialize["Type"] = o.Type
	}
	if !IsNil(o.DiscountAmount) {
		toSerialize["discount_amount"] = o.DiscountAmount
	}
	if !IsNil(o.DiscountDetails) {
		toSerialize["discount_details"] = o.DiscountDetails
	}
	if !IsNil(o.InvoiceDesc) {
		toSerialize["invoice_desc"] = o.InvoiceDesc
	}
	return toSerialize, nil
}

type NullableGetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner struct {
	value *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner
	isSet bool
}

func (v NullableGetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) Get() *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner {
	return v.value
}

func (v *NullableGetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) Set(val *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner(val *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) *NullableGetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner {
	return &NullableGetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner{value: val, isSet: true}
}

func (v NullableGetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


