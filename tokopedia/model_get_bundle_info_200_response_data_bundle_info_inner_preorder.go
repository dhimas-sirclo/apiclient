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

// checks if the GetBundleInfo200ResponseDataBundleInfoInnerPreorder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetBundleInfo200ResponseDataBundleInfoInnerPreorder{}

// GetBundleInfo200ResponseDataBundleInfoInnerPreorder struct for GetBundleInfo200ResponseDataBundleInfoInnerPreorder
type GetBundleInfo200ResponseDataBundleInfoInnerPreorder struct {
	Status *string `json:"status,omitempty"`
	StatusNum *int64 `json:"status_num,omitempty"`
	ProcessType *string `json:"process_type,omitempty"`
	ProcessTypeNum *int64 `json:"process_type_num,omitempty"`
	StartTime *string `json:"start_time,omitempty"`
	EndTime *string `json:"end_time,omitempty"`
	OrderLimit *int64 `json:"order_limit,omitempty"`
	MaxOrder *int64 `json:"max_order,omitempty"`
	ProcessDay *int64 `json:"process_day,omitempty"`
	ProcessTime *int64 `json:"process_time,omitempty"`
}

// NewGetBundleInfo200ResponseDataBundleInfoInnerPreorder instantiates a new GetBundleInfo200ResponseDataBundleInfoInnerPreorder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBundleInfo200ResponseDataBundleInfoInnerPreorder() *GetBundleInfo200ResponseDataBundleInfoInnerPreorder {
	this := GetBundleInfo200ResponseDataBundleInfoInnerPreorder{}
	return &this
}

// NewGetBundleInfo200ResponseDataBundleInfoInnerPreorderWithDefaults instantiates a new GetBundleInfo200ResponseDataBundleInfoInnerPreorder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBundleInfo200ResponseDataBundleInfoInnerPreorderWithDefaults() *GetBundleInfo200ResponseDataBundleInfoInnerPreorder {
	this := GetBundleInfo200ResponseDataBundleInfoInnerPreorder{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerPreorder) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerPreorder) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerPreorder) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerPreorder) SetStatus(v string) {
	o.Status = &v
}

// GetStatusNum returns the StatusNum field value if set, zero value otherwise.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerPreorder) GetStatusNum() int64 {
	if o == nil || IsNil(o.StatusNum) {
		var ret int64
		return ret
	}
	return *o.StatusNum
}

// GetStatusNumOk returns a tuple with the StatusNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerPreorder) GetStatusNumOk() (*int64, bool) {
	if o == nil || IsNil(o.StatusNum) {
		return nil, false
	}
	return o.StatusNum, true
}

// HasStatusNum returns a boolean if a field has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerPreorder) HasStatusNum() bool {
	if o != nil && !IsNil(o.StatusNum) {
		return true
	}

	return false
}

// SetStatusNum gets a reference to the given int64 and assigns it to the StatusNum field.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerPreorder) SetStatusNum(v int64) {
	o.StatusNum = &v
}

// GetProcessType returns the ProcessType field value if set, zero value otherwise.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerPreorder) GetProcessType() string {
	if o == nil || IsNil(o.ProcessType) {
		var ret string
		return ret
	}
	return *o.ProcessType
}

// GetProcessTypeOk returns a tuple with the ProcessType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerPreorder) GetProcessTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ProcessType) {
		return nil, false
	}
	return o.ProcessType, true
}

// HasProcessType returns a boolean if a field has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerPreorder) HasProcessType() bool {
	if o != nil && !IsNil(o.ProcessType) {
		return true
	}

	return false
}

// SetProcessType gets a reference to the given string and assigns it to the ProcessType field.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerPreorder) SetProcessType(v string) {
	o.ProcessType = &v
}

// GetProcessTypeNum returns the ProcessTypeNum field value if set, zero value otherwise.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerPreorder) GetProcessTypeNum() int64 {
	if o == nil || IsNil(o.ProcessTypeNum) {
		var ret int64
		return ret
	}
	return *o.ProcessTypeNum
}

// GetProcessTypeNumOk returns a tuple with the ProcessTypeNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerPreorder) GetProcessTypeNumOk() (*int64, bool) {
	if o == nil || IsNil(o.ProcessTypeNum) {
		return nil, false
	}
	return o.ProcessTypeNum, true
}

// HasProcessTypeNum returns a boolean if a field has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerPreorder) HasProcessTypeNum() bool {
	if o != nil && !IsNil(o.ProcessTypeNum) {
		return true
	}

	return false
}

// SetProcessTypeNum gets a reference to the given int64 and assigns it to the ProcessTypeNum field.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerPreorder) SetProcessTypeNum(v int64) {
	o.ProcessTypeNum = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerPreorder) GetStartTime() string {
	if o == nil || IsNil(o.StartTime) {
		var ret string
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerPreorder) GetStartTimeOk() (*string, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerPreorder) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given string and assigns it to the StartTime field.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerPreorder) SetStartTime(v string) {
	o.StartTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerPreorder) GetEndTime() string {
	if o == nil || IsNil(o.EndTime) {
		var ret string
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerPreorder) GetEndTimeOk() (*string, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerPreorder) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given string and assigns it to the EndTime field.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerPreorder) SetEndTime(v string) {
	o.EndTime = &v
}

// GetOrderLimit returns the OrderLimit field value if set, zero value otherwise.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerPreorder) GetOrderLimit() int64 {
	if o == nil || IsNil(o.OrderLimit) {
		var ret int64
		return ret
	}
	return *o.OrderLimit
}

// GetOrderLimitOk returns a tuple with the OrderLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerPreorder) GetOrderLimitOk() (*int64, bool) {
	if o == nil || IsNil(o.OrderLimit) {
		return nil, false
	}
	return o.OrderLimit, true
}

// HasOrderLimit returns a boolean if a field has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerPreorder) HasOrderLimit() bool {
	if o != nil && !IsNil(o.OrderLimit) {
		return true
	}

	return false
}

// SetOrderLimit gets a reference to the given int64 and assigns it to the OrderLimit field.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerPreorder) SetOrderLimit(v int64) {
	o.OrderLimit = &v
}

// GetMaxOrder returns the MaxOrder field value if set, zero value otherwise.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerPreorder) GetMaxOrder() int64 {
	if o == nil || IsNil(o.MaxOrder) {
		var ret int64
		return ret
	}
	return *o.MaxOrder
}

// GetMaxOrderOk returns a tuple with the MaxOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerPreorder) GetMaxOrderOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxOrder) {
		return nil, false
	}
	return o.MaxOrder, true
}

// HasMaxOrder returns a boolean if a field has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerPreorder) HasMaxOrder() bool {
	if o != nil && !IsNil(o.MaxOrder) {
		return true
	}

	return false
}

// SetMaxOrder gets a reference to the given int64 and assigns it to the MaxOrder field.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerPreorder) SetMaxOrder(v int64) {
	o.MaxOrder = &v
}

// GetProcessDay returns the ProcessDay field value if set, zero value otherwise.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerPreorder) GetProcessDay() int64 {
	if o == nil || IsNil(o.ProcessDay) {
		var ret int64
		return ret
	}
	return *o.ProcessDay
}

// GetProcessDayOk returns a tuple with the ProcessDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerPreorder) GetProcessDayOk() (*int64, bool) {
	if o == nil || IsNil(o.ProcessDay) {
		return nil, false
	}
	return o.ProcessDay, true
}

// HasProcessDay returns a boolean if a field has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerPreorder) HasProcessDay() bool {
	if o != nil && !IsNil(o.ProcessDay) {
		return true
	}

	return false
}

// SetProcessDay gets a reference to the given int64 and assigns it to the ProcessDay field.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerPreorder) SetProcessDay(v int64) {
	o.ProcessDay = &v
}

// GetProcessTime returns the ProcessTime field value if set, zero value otherwise.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerPreorder) GetProcessTime() int64 {
	if o == nil || IsNil(o.ProcessTime) {
		var ret int64
		return ret
	}
	return *o.ProcessTime
}

// GetProcessTimeOk returns a tuple with the ProcessTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerPreorder) GetProcessTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.ProcessTime) {
		return nil, false
	}
	return o.ProcessTime, true
}

// HasProcessTime returns a boolean if a field has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerPreorder) HasProcessTime() bool {
	if o != nil && !IsNil(o.ProcessTime) {
		return true
	}

	return false
}

// SetProcessTime gets a reference to the given int64 and assigns it to the ProcessTime field.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerPreorder) SetProcessTime(v int64) {
	o.ProcessTime = &v
}

func (o GetBundleInfo200ResponseDataBundleInfoInnerPreorder) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBundleInfo200ResponseDataBundleInfoInnerPreorder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.StatusNum) {
		toSerialize["status_num"] = o.StatusNum
	}
	if !IsNil(o.ProcessType) {
		toSerialize["process_type"] = o.ProcessType
	}
	if !IsNil(o.ProcessTypeNum) {
		toSerialize["process_type_num"] = o.ProcessTypeNum
	}
	if !IsNil(o.StartTime) {
		toSerialize["start_time"] = o.StartTime
	}
	if !IsNil(o.EndTime) {
		toSerialize["end_time"] = o.EndTime
	}
	if !IsNil(o.OrderLimit) {
		toSerialize["order_limit"] = o.OrderLimit
	}
	if !IsNil(o.MaxOrder) {
		toSerialize["max_order"] = o.MaxOrder
	}
	if !IsNil(o.ProcessDay) {
		toSerialize["process_day"] = o.ProcessDay
	}
	if !IsNil(o.ProcessTime) {
		toSerialize["process_time"] = o.ProcessTime
	}
	return toSerialize, nil
}

type NullableGetBundleInfo200ResponseDataBundleInfoInnerPreorder struct {
	value *GetBundleInfo200ResponseDataBundleInfoInnerPreorder
	isSet bool
}

func (v NullableGetBundleInfo200ResponseDataBundleInfoInnerPreorder) Get() *GetBundleInfo200ResponseDataBundleInfoInnerPreorder {
	return v.value
}

func (v *NullableGetBundleInfo200ResponseDataBundleInfoInnerPreorder) Set(val *GetBundleInfo200ResponseDataBundleInfoInnerPreorder) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBundleInfo200ResponseDataBundleInfoInnerPreorder) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBundleInfo200ResponseDataBundleInfoInnerPreorder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBundleInfo200ResponseDataBundleInfoInnerPreorder(val *GetBundleInfo200ResponseDataBundleInfoInnerPreorder) *NullableGetBundleInfo200ResponseDataBundleInfoInnerPreorder {
	return &NullableGetBundleInfo200ResponseDataBundleInfoInnerPreorder{value: val, isSet: true}
}

func (v NullableGetBundleInfo200ResponseDataBundleInfoInnerPreorder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBundleInfo200ResponseDataBundleInfoInnerPreorder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


