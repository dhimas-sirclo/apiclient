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

// checks if the CheckUploadStatus200ResponseDataUploadDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CheckUploadStatus200ResponseDataUploadDataInner{}

// CheckUploadStatus200ResponseDataUploadDataInner struct for CheckUploadStatus200ResponseDataUploadDataInner
type CheckUploadStatus200ResponseDataUploadDataInner struct {
	UploadId *int64 `json:"upload_id,omitempty"`
	Status *string `json:"status,omitempty"`
	TotalData *int64 `json:"total_data,omitempty"`
	UnprocessedRows *int64 `json:"unprocessed_rows,omitempty"`
	SuccessRows *int64 `json:"success_rows,omitempty"`
	FailedRows *int64 `json:"failed_rows,omitempty"`
	Processed *int64 `json:"processed,omitempty"`
}

// NewCheckUploadStatus200ResponseDataUploadDataInner instantiates a new CheckUploadStatus200ResponseDataUploadDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckUploadStatus200ResponseDataUploadDataInner() *CheckUploadStatus200ResponseDataUploadDataInner {
	this := CheckUploadStatus200ResponseDataUploadDataInner{}
	return &this
}

// NewCheckUploadStatus200ResponseDataUploadDataInnerWithDefaults instantiates a new CheckUploadStatus200ResponseDataUploadDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckUploadStatus200ResponseDataUploadDataInnerWithDefaults() *CheckUploadStatus200ResponseDataUploadDataInner {
	this := CheckUploadStatus200ResponseDataUploadDataInner{}
	return &this
}

// GetUploadId returns the UploadId field value if set, zero value otherwise.
func (o *CheckUploadStatus200ResponseDataUploadDataInner) GetUploadId() int64 {
	if o == nil || IsNil(o.UploadId) {
		var ret int64
		return ret
	}
	return *o.UploadId
}

// GetUploadIdOk returns a tuple with the UploadId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckUploadStatus200ResponseDataUploadDataInner) GetUploadIdOk() (*int64, bool) {
	if o == nil || IsNil(o.UploadId) {
		return nil, false
	}
	return o.UploadId, true
}

// HasUploadId returns a boolean if a field has been set.
func (o *CheckUploadStatus200ResponseDataUploadDataInner) HasUploadId() bool {
	if o != nil && !IsNil(o.UploadId) {
		return true
	}

	return false
}

// SetUploadId gets a reference to the given int64 and assigns it to the UploadId field.
func (o *CheckUploadStatus200ResponseDataUploadDataInner) SetUploadId(v int64) {
	o.UploadId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CheckUploadStatus200ResponseDataUploadDataInner) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckUploadStatus200ResponseDataUploadDataInner) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CheckUploadStatus200ResponseDataUploadDataInner) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CheckUploadStatus200ResponseDataUploadDataInner) SetStatus(v string) {
	o.Status = &v
}

// GetTotalData returns the TotalData field value if set, zero value otherwise.
func (o *CheckUploadStatus200ResponseDataUploadDataInner) GetTotalData() int64 {
	if o == nil || IsNil(o.TotalData) {
		var ret int64
		return ret
	}
	return *o.TotalData
}

// GetTotalDataOk returns a tuple with the TotalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckUploadStatus200ResponseDataUploadDataInner) GetTotalDataOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalData) {
		return nil, false
	}
	return o.TotalData, true
}

// HasTotalData returns a boolean if a field has been set.
func (o *CheckUploadStatus200ResponseDataUploadDataInner) HasTotalData() bool {
	if o != nil && !IsNil(o.TotalData) {
		return true
	}

	return false
}

// SetTotalData gets a reference to the given int64 and assigns it to the TotalData field.
func (o *CheckUploadStatus200ResponseDataUploadDataInner) SetTotalData(v int64) {
	o.TotalData = &v
}

// GetUnprocessedRows returns the UnprocessedRows field value if set, zero value otherwise.
func (o *CheckUploadStatus200ResponseDataUploadDataInner) GetUnprocessedRows() int64 {
	if o == nil || IsNil(o.UnprocessedRows) {
		var ret int64
		return ret
	}
	return *o.UnprocessedRows
}

// GetUnprocessedRowsOk returns a tuple with the UnprocessedRows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckUploadStatus200ResponseDataUploadDataInner) GetUnprocessedRowsOk() (*int64, bool) {
	if o == nil || IsNil(o.UnprocessedRows) {
		return nil, false
	}
	return o.UnprocessedRows, true
}

// HasUnprocessedRows returns a boolean if a field has been set.
func (o *CheckUploadStatus200ResponseDataUploadDataInner) HasUnprocessedRows() bool {
	if o != nil && !IsNil(o.UnprocessedRows) {
		return true
	}

	return false
}

// SetUnprocessedRows gets a reference to the given int64 and assigns it to the UnprocessedRows field.
func (o *CheckUploadStatus200ResponseDataUploadDataInner) SetUnprocessedRows(v int64) {
	o.UnprocessedRows = &v
}

// GetSuccessRows returns the SuccessRows field value if set, zero value otherwise.
func (o *CheckUploadStatus200ResponseDataUploadDataInner) GetSuccessRows() int64 {
	if o == nil || IsNil(o.SuccessRows) {
		var ret int64
		return ret
	}
	return *o.SuccessRows
}

// GetSuccessRowsOk returns a tuple with the SuccessRows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckUploadStatus200ResponseDataUploadDataInner) GetSuccessRowsOk() (*int64, bool) {
	if o == nil || IsNil(o.SuccessRows) {
		return nil, false
	}
	return o.SuccessRows, true
}

// HasSuccessRows returns a boolean if a field has been set.
func (o *CheckUploadStatus200ResponseDataUploadDataInner) HasSuccessRows() bool {
	if o != nil && !IsNil(o.SuccessRows) {
		return true
	}

	return false
}

// SetSuccessRows gets a reference to the given int64 and assigns it to the SuccessRows field.
func (o *CheckUploadStatus200ResponseDataUploadDataInner) SetSuccessRows(v int64) {
	o.SuccessRows = &v
}

// GetFailedRows returns the FailedRows field value if set, zero value otherwise.
func (o *CheckUploadStatus200ResponseDataUploadDataInner) GetFailedRows() int64 {
	if o == nil || IsNil(o.FailedRows) {
		var ret int64
		return ret
	}
	return *o.FailedRows
}

// GetFailedRowsOk returns a tuple with the FailedRows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckUploadStatus200ResponseDataUploadDataInner) GetFailedRowsOk() (*int64, bool) {
	if o == nil || IsNil(o.FailedRows) {
		return nil, false
	}
	return o.FailedRows, true
}

// HasFailedRows returns a boolean if a field has been set.
func (o *CheckUploadStatus200ResponseDataUploadDataInner) HasFailedRows() bool {
	if o != nil && !IsNil(o.FailedRows) {
		return true
	}

	return false
}

// SetFailedRows gets a reference to the given int64 and assigns it to the FailedRows field.
func (o *CheckUploadStatus200ResponseDataUploadDataInner) SetFailedRows(v int64) {
	o.FailedRows = &v
}

// GetProcessed returns the Processed field value if set, zero value otherwise.
func (o *CheckUploadStatus200ResponseDataUploadDataInner) GetProcessed() int64 {
	if o == nil || IsNil(o.Processed) {
		var ret int64
		return ret
	}
	return *o.Processed
}

// GetProcessedOk returns a tuple with the Processed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckUploadStatus200ResponseDataUploadDataInner) GetProcessedOk() (*int64, bool) {
	if o == nil || IsNil(o.Processed) {
		return nil, false
	}
	return o.Processed, true
}

// HasProcessed returns a boolean if a field has been set.
func (o *CheckUploadStatus200ResponseDataUploadDataInner) HasProcessed() bool {
	if o != nil && !IsNil(o.Processed) {
		return true
	}

	return false
}

// SetProcessed gets a reference to the given int64 and assigns it to the Processed field.
func (o *CheckUploadStatus200ResponseDataUploadDataInner) SetProcessed(v int64) {
	o.Processed = &v
}

func (o CheckUploadStatus200ResponseDataUploadDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckUploadStatus200ResponseDataUploadDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UploadId) {
		toSerialize["upload_id"] = o.UploadId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.TotalData) {
		toSerialize["total_data"] = o.TotalData
	}
	if !IsNil(o.UnprocessedRows) {
		toSerialize["unprocessed_rows"] = o.UnprocessedRows
	}
	if !IsNil(o.SuccessRows) {
		toSerialize["success_rows"] = o.SuccessRows
	}
	if !IsNil(o.FailedRows) {
		toSerialize["failed_rows"] = o.FailedRows
	}
	if !IsNil(o.Processed) {
		toSerialize["processed"] = o.Processed
	}
	return toSerialize, nil
}

type NullableCheckUploadStatus200ResponseDataUploadDataInner struct {
	value *CheckUploadStatus200ResponseDataUploadDataInner
	isSet bool
}

func (v NullableCheckUploadStatus200ResponseDataUploadDataInner) Get() *CheckUploadStatus200ResponseDataUploadDataInner {
	return v.value
}

func (v *NullableCheckUploadStatus200ResponseDataUploadDataInner) Set(val *CheckUploadStatus200ResponseDataUploadDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckUploadStatus200ResponseDataUploadDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckUploadStatus200ResponseDataUploadDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckUploadStatus200ResponseDataUploadDataInner(val *CheckUploadStatus200ResponseDataUploadDataInner) *NullableCheckUploadStatus200ResponseDataUploadDataInner {
	return &NullableCheckUploadStatus200ResponseDataUploadDataInner{value: val, isSet: true}
}

func (v NullableCheckUploadStatus200ResponseDataUploadDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckUploadStatus200ResponseDataUploadDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

