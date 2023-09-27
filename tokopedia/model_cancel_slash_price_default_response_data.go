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

// checks if the CancelSlashPriceDefaultResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CancelSlashPriceDefaultResponseData{}

// CancelSlashPriceDefaultResponseData struct for CancelSlashPriceDefaultResponseData
type CancelSlashPriceDefaultResponseData struct {
	// Total Data Count
	TotalData *int64 `json:"total_data,omitempty"`
	// Success Data Count
	SucceedRows *int64 `json:"succeed_rows,omitempty"`
	// Failed Data Count
	FailedRows *int64 `json:"failed_rows,omitempty"`
	// Total Failed Data Count
	FailedRowsData []string `json:"failed_rows_data,omitempty"`
}

// NewCancelSlashPriceDefaultResponseData instantiates a new CancelSlashPriceDefaultResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCancelSlashPriceDefaultResponseData() *CancelSlashPriceDefaultResponseData {
	this := CancelSlashPriceDefaultResponseData{}
	return &this
}

// NewCancelSlashPriceDefaultResponseDataWithDefaults instantiates a new CancelSlashPriceDefaultResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCancelSlashPriceDefaultResponseDataWithDefaults() *CancelSlashPriceDefaultResponseData {
	this := CancelSlashPriceDefaultResponseData{}
	return &this
}

// GetTotalData returns the TotalData field value if set, zero value otherwise.
func (o *CancelSlashPriceDefaultResponseData) GetTotalData() int64 {
	if o == nil || IsNil(o.TotalData) {
		var ret int64
		return ret
	}
	return *o.TotalData
}

// GetTotalDataOk returns a tuple with the TotalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelSlashPriceDefaultResponseData) GetTotalDataOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalData) {
		return nil, false
	}
	return o.TotalData, true
}

// HasTotalData returns a boolean if a field has been set.
func (o *CancelSlashPriceDefaultResponseData) HasTotalData() bool {
	if o != nil && !IsNil(o.TotalData) {
		return true
	}

	return false
}

// SetTotalData gets a reference to the given int64 and assigns it to the TotalData field.
func (o *CancelSlashPriceDefaultResponseData) SetTotalData(v int64) {
	o.TotalData = &v
}

// GetSucceedRows returns the SucceedRows field value if set, zero value otherwise.
func (o *CancelSlashPriceDefaultResponseData) GetSucceedRows() int64 {
	if o == nil || IsNil(o.SucceedRows) {
		var ret int64
		return ret
	}
	return *o.SucceedRows
}

// GetSucceedRowsOk returns a tuple with the SucceedRows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelSlashPriceDefaultResponseData) GetSucceedRowsOk() (*int64, bool) {
	if o == nil || IsNil(o.SucceedRows) {
		return nil, false
	}
	return o.SucceedRows, true
}

// HasSucceedRows returns a boolean if a field has been set.
func (o *CancelSlashPriceDefaultResponseData) HasSucceedRows() bool {
	if o != nil && !IsNil(o.SucceedRows) {
		return true
	}

	return false
}

// SetSucceedRows gets a reference to the given int64 and assigns it to the SucceedRows field.
func (o *CancelSlashPriceDefaultResponseData) SetSucceedRows(v int64) {
	o.SucceedRows = &v
}

// GetFailedRows returns the FailedRows field value if set, zero value otherwise.
func (o *CancelSlashPriceDefaultResponseData) GetFailedRows() int64 {
	if o == nil || IsNil(o.FailedRows) {
		var ret int64
		return ret
	}
	return *o.FailedRows
}

// GetFailedRowsOk returns a tuple with the FailedRows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelSlashPriceDefaultResponseData) GetFailedRowsOk() (*int64, bool) {
	if o == nil || IsNil(o.FailedRows) {
		return nil, false
	}
	return o.FailedRows, true
}

// HasFailedRows returns a boolean if a field has been set.
func (o *CancelSlashPriceDefaultResponseData) HasFailedRows() bool {
	if o != nil && !IsNil(o.FailedRows) {
		return true
	}

	return false
}

// SetFailedRows gets a reference to the given int64 and assigns it to the FailedRows field.
func (o *CancelSlashPriceDefaultResponseData) SetFailedRows(v int64) {
	o.FailedRows = &v
}

// GetFailedRowsData returns the FailedRowsData field value if set, zero value otherwise.
func (o *CancelSlashPriceDefaultResponseData) GetFailedRowsData() []string {
	if o == nil || IsNil(o.FailedRowsData) {
		var ret []string
		return ret
	}
	return o.FailedRowsData
}

// GetFailedRowsDataOk returns a tuple with the FailedRowsData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelSlashPriceDefaultResponseData) GetFailedRowsDataOk() ([]string, bool) {
	if o == nil || IsNil(o.FailedRowsData) {
		return nil, false
	}
	return o.FailedRowsData, true
}

// HasFailedRowsData returns a boolean if a field has been set.
func (o *CancelSlashPriceDefaultResponseData) HasFailedRowsData() bool {
	if o != nil && !IsNil(o.FailedRowsData) {
		return true
	}

	return false
}

// SetFailedRowsData gets a reference to the given []string and assigns it to the FailedRowsData field.
func (o *CancelSlashPriceDefaultResponseData) SetFailedRowsData(v []string) {
	o.FailedRowsData = v
}

func (o CancelSlashPriceDefaultResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CancelSlashPriceDefaultResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalData) {
		toSerialize["total_data"] = o.TotalData
	}
	if !IsNil(o.SucceedRows) {
		toSerialize["succeed_rows"] = o.SucceedRows
	}
	if !IsNil(o.FailedRows) {
		toSerialize["failed_rows"] = o.FailedRows
	}
	if !IsNil(o.FailedRowsData) {
		toSerialize["failed_rows_data"] = o.FailedRowsData
	}
	return toSerialize, nil
}

type NullableCancelSlashPriceDefaultResponseData struct {
	value *CancelSlashPriceDefaultResponseData
	isSet bool
}

func (v NullableCancelSlashPriceDefaultResponseData) Get() *CancelSlashPriceDefaultResponseData {
	return v.value
}

func (v *NullableCancelSlashPriceDefaultResponseData) Set(val *CancelSlashPriceDefaultResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableCancelSlashPriceDefaultResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableCancelSlashPriceDefaultResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCancelSlashPriceDefaultResponseData(val *CancelSlashPriceDefaultResponseData) *NullableCancelSlashPriceDefaultResponseData {
	return &NullableCancelSlashPriceDefaultResponseData{value: val, isSet: true}
}

func (v NullableCancelSlashPriceDefaultResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCancelSlashPriceDefaultResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

