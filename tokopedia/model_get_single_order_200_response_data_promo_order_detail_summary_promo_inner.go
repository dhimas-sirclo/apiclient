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

// checks if the GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner{}

// GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner struct for GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner
type GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner struct {
	// Promo Title
	Name *string `json:"name,omitempty"`
	// Is promo a coupon or not
	IsCoupon *bool `json:"is_coupon,omitempty"`
	// Flag to show cashback amount (hardcoded true)
	ShowCashbackAmount *bool `json:"show_cashback_amount,omitempty"`
	// Flag to show discount amount (hardcoded true)
	ShowDiscountAmount *bool `json:"show_discount_amount,omitempty"`
	// Cashback amount in IDR
	CashbackAmount *int64 `json:"cashback_amount,omitempty"`
	// Cashback amount in Points currency
	CashbackPoints *int64 `json:"cashback_points,omitempty"`
	// Details of cashback in 1 promo
	CashbackDetails map[string]interface{} `json:"cashback_details,omitempty"`
	// Promo benefit type. (discount, cashback)
	Type *string `json:"type,omitempty"`
	// Discount amount
	DiscountAmount *int64 `json:"discount_amount,omitempty"`
	// Details of discount in 1 promo
	DiscountDetails []GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInnerDiscountDetailsInner `json:"discount_details,omitempty"`
	InvoiceDesc *string `json:"invoice_desc,omitempty"`
}

// NewGetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner instantiates a new GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner() *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner {
	this := GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner{}
	return &this
}

// NewGetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInnerWithDefaults instantiates a new GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInnerWithDefaults() *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner {
	this := GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) SetName(v string) {
	o.Name = &v
}

// GetIsCoupon returns the IsCoupon field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) GetIsCoupon() bool {
	if o == nil || IsNil(o.IsCoupon) {
		var ret bool
		return ret
	}
	return *o.IsCoupon
}

// GetIsCouponOk returns a tuple with the IsCoupon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) GetIsCouponOk() (*bool, bool) {
	if o == nil || IsNil(o.IsCoupon) {
		return nil, false
	}
	return o.IsCoupon, true
}

// HasIsCoupon returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) HasIsCoupon() bool {
	if o != nil && !IsNil(o.IsCoupon) {
		return true
	}

	return false
}

// SetIsCoupon gets a reference to the given bool and assigns it to the IsCoupon field.
func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) SetIsCoupon(v bool) {
	o.IsCoupon = &v
}

// GetShowCashbackAmount returns the ShowCashbackAmount field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) GetShowCashbackAmount() bool {
	if o == nil || IsNil(o.ShowCashbackAmount) {
		var ret bool
		return ret
	}
	return *o.ShowCashbackAmount
}

// GetShowCashbackAmountOk returns a tuple with the ShowCashbackAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) GetShowCashbackAmountOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowCashbackAmount) {
		return nil, false
	}
	return o.ShowCashbackAmount, true
}

// HasShowCashbackAmount returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) HasShowCashbackAmount() bool {
	if o != nil && !IsNil(o.ShowCashbackAmount) {
		return true
	}

	return false
}

// SetShowCashbackAmount gets a reference to the given bool and assigns it to the ShowCashbackAmount field.
func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) SetShowCashbackAmount(v bool) {
	o.ShowCashbackAmount = &v
}

// GetShowDiscountAmount returns the ShowDiscountAmount field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) GetShowDiscountAmount() bool {
	if o == nil || IsNil(o.ShowDiscountAmount) {
		var ret bool
		return ret
	}
	return *o.ShowDiscountAmount
}

// GetShowDiscountAmountOk returns a tuple with the ShowDiscountAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) GetShowDiscountAmountOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowDiscountAmount) {
		return nil, false
	}
	return o.ShowDiscountAmount, true
}

// HasShowDiscountAmount returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) HasShowDiscountAmount() bool {
	if o != nil && !IsNil(o.ShowDiscountAmount) {
		return true
	}

	return false
}

// SetShowDiscountAmount gets a reference to the given bool and assigns it to the ShowDiscountAmount field.
func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) SetShowDiscountAmount(v bool) {
	o.ShowDiscountAmount = &v
}

// GetCashbackAmount returns the CashbackAmount field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) GetCashbackAmount() int64 {
	if o == nil || IsNil(o.CashbackAmount) {
		var ret int64
		return ret
	}
	return *o.CashbackAmount
}

// GetCashbackAmountOk returns a tuple with the CashbackAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) GetCashbackAmountOk() (*int64, bool) {
	if o == nil || IsNil(o.CashbackAmount) {
		return nil, false
	}
	return o.CashbackAmount, true
}

// HasCashbackAmount returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) HasCashbackAmount() bool {
	if o != nil && !IsNil(o.CashbackAmount) {
		return true
	}

	return false
}

// SetCashbackAmount gets a reference to the given int64 and assigns it to the CashbackAmount field.
func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) SetCashbackAmount(v int64) {
	o.CashbackAmount = &v
}

// GetCashbackPoints returns the CashbackPoints field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) GetCashbackPoints() int64 {
	if o == nil || IsNil(o.CashbackPoints) {
		var ret int64
		return ret
	}
	return *o.CashbackPoints
}

// GetCashbackPointsOk returns a tuple with the CashbackPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) GetCashbackPointsOk() (*int64, bool) {
	if o == nil || IsNil(o.CashbackPoints) {
		return nil, false
	}
	return o.CashbackPoints, true
}

// HasCashbackPoints returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) HasCashbackPoints() bool {
	if o != nil && !IsNil(o.CashbackPoints) {
		return true
	}

	return false
}

// SetCashbackPoints gets a reference to the given int64 and assigns it to the CashbackPoints field.
func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) SetCashbackPoints(v int64) {
	o.CashbackPoints = &v
}

// GetCashbackDetails returns the CashbackDetails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) GetCashbackDetails() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.CashbackDetails
}

// GetCashbackDetailsOk returns a tuple with the CashbackDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) GetCashbackDetailsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CashbackDetails) {
		return map[string]interface{}{}, false
	}
	return o.CashbackDetails, true
}

// HasCashbackDetails returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) HasCashbackDetails() bool {
	if o != nil && IsNil(o.CashbackDetails) {
		return true
	}

	return false
}

// SetCashbackDetails gets a reference to the given map[string]interface{} and assigns it to the CashbackDetails field.
func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) SetCashbackDetails(v map[string]interface{}) {
	o.CashbackDetails = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) SetType(v string) {
	o.Type = &v
}

// GetDiscountAmount returns the DiscountAmount field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) GetDiscountAmount() int64 {
	if o == nil || IsNil(o.DiscountAmount) {
		var ret int64
		return ret
	}
	return *o.DiscountAmount
}

// GetDiscountAmountOk returns a tuple with the DiscountAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) GetDiscountAmountOk() (*int64, bool) {
	if o == nil || IsNil(o.DiscountAmount) {
		return nil, false
	}
	return o.DiscountAmount, true
}

// HasDiscountAmount returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) HasDiscountAmount() bool {
	if o != nil && !IsNil(o.DiscountAmount) {
		return true
	}

	return false
}

// SetDiscountAmount gets a reference to the given int64 and assigns it to the DiscountAmount field.
func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) SetDiscountAmount(v int64) {
	o.DiscountAmount = &v
}

// GetDiscountDetails returns the DiscountDetails field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) GetDiscountDetails() []GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInnerDiscountDetailsInner {
	if o == nil || IsNil(o.DiscountDetails) {
		var ret []GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInnerDiscountDetailsInner
		return ret
	}
	return o.DiscountDetails
}

// GetDiscountDetailsOk returns a tuple with the DiscountDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) GetDiscountDetailsOk() ([]GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInnerDiscountDetailsInner, bool) {
	if o == nil || IsNil(o.DiscountDetails) {
		return nil, false
	}
	return o.DiscountDetails, true
}

// HasDiscountDetails returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) HasDiscountDetails() bool {
	if o != nil && !IsNil(o.DiscountDetails) {
		return true
	}

	return false
}

// SetDiscountDetails gets a reference to the given []GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInnerDiscountDetailsInner and assigns it to the DiscountDetails field.
func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) SetDiscountDetails(v []GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInnerDiscountDetailsInner) {
	o.DiscountDetails = v
}

// GetInvoiceDesc returns the InvoiceDesc field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) GetInvoiceDesc() string {
	if o == nil || IsNil(o.InvoiceDesc) {
		var ret string
		return ret
	}
	return *o.InvoiceDesc
}

// GetInvoiceDescOk returns a tuple with the InvoiceDesc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) GetInvoiceDescOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceDesc) {
		return nil, false
	}
	return o.InvoiceDesc, true
}

// HasInvoiceDesc returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) HasInvoiceDesc() bool {
	if o != nil && !IsNil(o.InvoiceDesc) {
		return true
	}

	return false
}

// SetInvoiceDesc gets a reference to the given string and assigns it to the InvoiceDesc field.
func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) SetInvoiceDesc(v string) {
	o.InvoiceDesc = &v
}

func (o GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
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
	if o.CashbackDetails != nil {
		toSerialize["cashback_details"] = o.CashbackDetails
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
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

type NullableGetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner struct {
	value *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner
	isSet bool
}

func (v NullableGetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) Get() *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner {
	return v.value
}

func (v *NullableGetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) Set(val *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner(val *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) *NullableGetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner {
	return &NullableGetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner{value: val, isSet: true}
}

func (v NullableGetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


