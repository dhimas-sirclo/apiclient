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

// checks if the GetSingleOrder200ResponseDataShipmentFulfillmentItemDeliveredDeadline type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSingleOrder200ResponseDataShipmentFulfillmentItemDeliveredDeadline{}

// GetSingleOrder200ResponseDataShipmentFulfillmentItemDeliveredDeadline struct for GetSingleOrder200ResponseDataShipmentFulfillmentItemDeliveredDeadline
type GetSingleOrder200ResponseDataShipmentFulfillmentItemDeliveredDeadline struct {
	// Item Delivered Deadline Timestamp (format: 0001-01-01T00:00:00Z) 
	Time *string `json:"Time,omitempty"`
	// Is valid?
	Valid *bool `json:"Valid,omitempty"`
}

// NewGetSingleOrder200ResponseDataShipmentFulfillmentItemDeliveredDeadline instantiates a new GetSingleOrder200ResponseDataShipmentFulfillmentItemDeliveredDeadline object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSingleOrder200ResponseDataShipmentFulfillmentItemDeliveredDeadline() *GetSingleOrder200ResponseDataShipmentFulfillmentItemDeliveredDeadline {
	this := GetSingleOrder200ResponseDataShipmentFulfillmentItemDeliveredDeadline{}
	return &this
}

// NewGetSingleOrder200ResponseDataShipmentFulfillmentItemDeliveredDeadlineWithDefaults instantiates a new GetSingleOrder200ResponseDataShipmentFulfillmentItemDeliveredDeadline object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSingleOrder200ResponseDataShipmentFulfillmentItemDeliveredDeadlineWithDefaults() *GetSingleOrder200ResponseDataShipmentFulfillmentItemDeliveredDeadline {
	this := GetSingleOrder200ResponseDataShipmentFulfillmentItemDeliveredDeadline{}
	return &this
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataShipmentFulfillmentItemDeliveredDeadline) GetTime() string {
	if o == nil || IsNil(o.Time) {
		var ret string
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataShipmentFulfillmentItemDeliveredDeadline) GetTimeOk() (*string, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataShipmentFulfillmentItemDeliveredDeadline) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given string and assigns it to the Time field.
func (o *GetSingleOrder200ResponseDataShipmentFulfillmentItemDeliveredDeadline) SetTime(v string) {
	o.Time = &v
}

// GetValid returns the Valid field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataShipmentFulfillmentItemDeliveredDeadline) GetValid() bool {
	if o == nil || IsNil(o.Valid) {
		var ret bool
		return ret
	}
	return *o.Valid
}

// GetValidOk returns a tuple with the Valid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataShipmentFulfillmentItemDeliveredDeadline) GetValidOk() (*bool, bool) {
	if o == nil || IsNil(o.Valid) {
		return nil, false
	}
	return o.Valid, true
}

// HasValid returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataShipmentFulfillmentItemDeliveredDeadline) HasValid() bool {
	if o != nil && !IsNil(o.Valid) {
		return true
	}

	return false
}

// SetValid gets a reference to the given bool and assigns it to the Valid field.
func (o *GetSingleOrder200ResponseDataShipmentFulfillmentItemDeliveredDeadline) SetValid(v bool) {
	o.Valid = &v
}

func (o GetSingleOrder200ResponseDataShipmentFulfillmentItemDeliveredDeadline) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSingleOrder200ResponseDataShipmentFulfillmentItemDeliveredDeadline) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Time) {
		toSerialize["Time"] = o.Time
	}
	if !IsNil(o.Valid) {
		toSerialize["Valid"] = o.Valid
	}
	return toSerialize, nil
}

type NullableGetSingleOrder200ResponseDataShipmentFulfillmentItemDeliveredDeadline struct {
	value *GetSingleOrder200ResponseDataShipmentFulfillmentItemDeliveredDeadline
	isSet bool
}

func (v NullableGetSingleOrder200ResponseDataShipmentFulfillmentItemDeliveredDeadline) Get() *GetSingleOrder200ResponseDataShipmentFulfillmentItemDeliveredDeadline {
	return v.value
}

func (v *NullableGetSingleOrder200ResponseDataShipmentFulfillmentItemDeliveredDeadline) Set(val *GetSingleOrder200ResponseDataShipmentFulfillmentItemDeliveredDeadline) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSingleOrder200ResponseDataShipmentFulfillmentItemDeliveredDeadline) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSingleOrder200ResponseDataShipmentFulfillmentItemDeliveredDeadline) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSingleOrder200ResponseDataShipmentFulfillmentItemDeliveredDeadline(val *GetSingleOrder200ResponseDataShipmentFulfillmentItemDeliveredDeadline) *NullableGetSingleOrder200ResponseDataShipmentFulfillmentItemDeliveredDeadline {
	return &NullableGetSingleOrder200ResponseDataShipmentFulfillmentItemDeliveredDeadline{value: val, isSet: true}
}

func (v NullableGetSingleOrder200ResponseDataShipmentFulfillmentItemDeliveredDeadline) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSingleOrder200ResponseDataShipmentFulfillmentItemDeliveredDeadline) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

