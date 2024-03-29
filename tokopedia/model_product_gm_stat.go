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

// checks if the ProductGMStat type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductGMStat{}

// ProductGMStat struct for ProductGMStat
type ProductGMStat struct {
	TransactionSuccess *int64 `json:"transactionSuccess,omitempty"`
	TransactionReject *int64 `json:"transactionReject,omitempty"`
	CountSold *int64 `json:"countSold,omitempty"`
}

// NewProductGMStat instantiates a new ProductGMStat object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductGMStat() *ProductGMStat {
	this := ProductGMStat{}
	return &this
}

// NewProductGMStatWithDefaults instantiates a new ProductGMStat object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductGMStatWithDefaults() *ProductGMStat {
	this := ProductGMStat{}
	return &this
}

// GetTransactionSuccess returns the TransactionSuccess field value if set, zero value otherwise.
func (o *ProductGMStat) GetTransactionSuccess() int64 {
	if o == nil || IsNil(o.TransactionSuccess) {
		var ret int64
		return ret
	}
	return *o.TransactionSuccess
}

// GetTransactionSuccessOk returns a tuple with the TransactionSuccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductGMStat) GetTransactionSuccessOk() (*int64, bool) {
	if o == nil || IsNil(o.TransactionSuccess) {
		return nil, false
	}
	return o.TransactionSuccess, true
}

// HasTransactionSuccess returns a boolean if a field has been set.
func (o *ProductGMStat) HasTransactionSuccess() bool {
	if o != nil && !IsNil(o.TransactionSuccess) {
		return true
	}

	return false
}

// SetTransactionSuccess gets a reference to the given int64 and assigns it to the TransactionSuccess field.
func (o *ProductGMStat) SetTransactionSuccess(v int64) {
	o.TransactionSuccess = &v
}

// GetTransactionReject returns the TransactionReject field value if set, zero value otherwise.
func (o *ProductGMStat) GetTransactionReject() int64 {
	if o == nil || IsNil(o.TransactionReject) {
		var ret int64
		return ret
	}
	return *o.TransactionReject
}

// GetTransactionRejectOk returns a tuple with the TransactionReject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductGMStat) GetTransactionRejectOk() (*int64, bool) {
	if o == nil || IsNil(o.TransactionReject) {
		return nil, false
	}
	return o.TransactionReject, true
}

// HasTransactionReject returns a boolean if a field has been set.
func (o *ProductGMStat) HasTransactionReject() bool {
	if o != nil && !IsNil(o.TransactionReject) {
		return true
	}

	return false
}

// SetTransactionReject gets a reference to the given int64 and assigns it to the TransactionReject field.
func (o *ProductGMStat) SetTransactionReject(v int64) {
	o.TransactionReject = &v
}

// GetCountSold returns the CountSold field value if set, zero value otherwise.
func (o *ProductGMStat) GetCountSold() int64 {
	if o == nil || IsNil(o.CountSold) {
		var ret int64
		return ret
	}
	return *o.CountSold
}

// GetCountSoldOk returns a tuple with the CountSold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductGMStat) GetCountSoldOk() (*int64, bool) {
	if o == nil || IsNil(o.CountSold) {
		return nil, false
	}
	return o.CountSold, true
}

// HasCountSold returns a boolean if a field has been set.
func (o *ProductGMStat) HasCountSold() bool {
	if o != nil && !IsNil(o.CountSold) {
		return true
	}

	return false
}

// SetCountSold gets a reference to the given int64 and assigns it to the CountSold field.
func (o *ProductGMStat) SetCountSold(v int64) {
	o.CountSold = &v
}

func (o ProductGMStat) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductGMStat) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TransactionSuccess) {
		toSerialize["transactionSuccess"] = o.TransactionSuccess
	}
	if !IsNil(o.TransactionReject) {
		toSerialize["transactionReject"] = o.TransactionReject
	}
	if !IsNil(o.CountSold) {
		toSerialize["countSold"] = o.CountSold
	}
	return toSerialize, nil
}

type NullableProductGMStat struct {
	value *ProductGMStat
	isSet bool
}

func (v NullableProductGMStat) Get() *ProductGMStat {
	return v.value
}

func (v *NullableProductGMStat) Set(val *ProductGMStat) {
	v.value = val
	v.isSet = true
}

func (v NullableProductGMStat) IsSet() bool {
	return v.isSet
}

func (v *NullableProductGMStat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductGMStat(val *ProductGMStat) *NullableProductGMStat {
	return &NullableProductGMStat{value: val, isSet: true}
}

func (v NullableProductGMStat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductGMStat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


