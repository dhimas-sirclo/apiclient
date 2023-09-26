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

// checks if the GetOrderWebhook200ResponseDataOneOf1VoucherInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrderWebhook200ResponseDataOneOf1VoucherInfo{}

// GetOrderWebhook200ResponseDataOneOf1VoucherInfo struct for GetOrderWebhook200ResponseDataOneOf1VoucherInfo
type GetOrderWebhook200ResponseDataOneOf1VoucherInfo struct {
	VoucherCode *string `json:"voucher_code,omitempty"`
	VoucherType *int64 `json:"voucher_type,omitempty"`
}

// NewGetOrderWebhook200ResponseDataOneOf1VoucherInfo instantiates a new GetOrderWebhook200ResponseDataOneOf1VoucherInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrderWebhook200ResponseDataOneOf1VoucherInfo() *GetOrderWebhook200ResponseDataOneOf1VoucherInfo {
	this := GetOrderWebhook200ResponseDataOneOf1VoucherInfo{}
	return &this
}

// NewGetOrderWebhook200ResponseDataOneOf1VoucherInfoWithDefaults instantiates a new GetOrderWebhook200ResponseDataOneOf1VoucherInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrderWebhook200ResponseDataOneOf1VoucherInfoWithDefaults() *GetOrderWebhook200ResponseDataOneOf1VoucherInfo {
	this := GetOrderWebhook200ResponseDataOneOf1VoucherInfo{}
	return &this
}

// GetVoucherCode returns the VoucherCode field value if set, zero value otherwise.
func (o *GetOrderWebhook200ResponseDataOneOf1VoucherInfo) GetVoucherCode() string {
	if o == nil || IsNil(o.VoucherCode) {
		var ret string
		return ret
	}
	return *o.VoucherCode
}

// GetVoucherCodeOk returns a tuple with the VoucherCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderWebhook200ResponseDataOneOf1VoucherInfo) GetVoucherCodeOk() (*string, bool) {
	if o == nil || IsNil(o.VoucherCode) {
		return nil, false
	}
	return o.VoucherCode, true
}

// HasVoucherCode returns a boolean if a field has been set.
func (o *GetOrderWebhook200ResponseDataOneOf1VoucherInfo) HasVoucherCode() bool {
	if o != nil && !IsNil(o.VoucherCode) {
		return true
	}

	return false
}

// SetVoucherCode gets a reference to the given string and assigns it to the VoucherCode field.
func (o *GetOrderWebhook200ResponseDataOneOf1VoucherInfo) SetVoucherCode(v string) {
	o.VoucherCode = &v
}

// GetVoucherType returns the VoucherType field value if set, zero value otherwise.
func (o *GetOrderWebhook200ResponseDataOneOf1VoucherInfo) GetVoucherType() int64 {
	if o == nil || IsNil(o.VoucherType) {
		var ret int64
		return ret
	}
	return *o.VoucherType
}

// GetVoucherTypeOk returns a tuple with the VoucherType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderWebhook200ResponseDataOneOf1VoucherInfo) GetVoucherTypeOk() (*int64, bool) {
	if o == nil || IsNil(o.VoucherType) {
		return nil, false
	}
	return o.VoucherType, true
}

// HasVoucherType returns a boolean if a field has been set.
func (o *GetOrderWebhook200ResponseDataOneOf1VoucherInfo) HasVoucherType() bool {
	if o != nil && !IsNil(o.VoucherType) {
		return true
	}

	return false
}

// SetVoucherType gets a reference to the given int64 and assigns it to the VoucherType field.
func (o *GetOrderWebhook200ResponseDataOneOf1VoucherInfo) SetVoucherType(v int64) {
	o.VoucherType = &v
}

func (o GetOrderWebhook200ResponseDataOneOf1VoucherInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrderWebhook200ResponseDataOneOf1VoucherInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VoucherCode) {
		toSerialize["voucher_code"] = o.VoucherCode
	}
	if !IsNil(o.VoucherType) {
		toSerialize["voucher_type"] = o.VoucherType
	}
	return toSerialize, nil
}

type NullableGetOrderWebhook200ResponseDataOneOf1VoucherInfo struct {
	value *GetOrderWebhook200ResponseDataOneOf1VoucherInfo
	isSet bool
}

func (v NullableGetOrderWebhook200ResponseDataOneOf1VoucherInfo) Get() *GetOrderWebhook200ResponseDataOneOf1VoucherInfo {
	return v.value
}

func (v *NullableGetOrderWebhook200ResponseDataOneOf1VoucherInfo) Set(val *GetOrderWebhook200ResponseDataOneOf1VoucherInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrderWebhook200ResponseDataOneOf1VoucherInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrderWebhook200ResponseDataOneOf1VoucherInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrderWebhook200ResponseDataOneOf1VoucherInfo(val *GetOrderWebhook200ResponseDataOneOf1VoucherInfo) *NullableGetOrderWebhook200ResponseDataOneOf1VoucherInfo {
	return &NullableGetOrderWebhook200ResponseDataOneOf1VoucherInfo{value: val, isSet: true}
}

func (v NullableGetOrderWebhook200ResponseDataOneOf1VoucherInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrderWebhook200ResponseDataOneOf1VoucherInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

