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

// checks if the ViewSlashPrice200ResponseDataInnerPrice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ViewSlashPrice200ResponseDataInnerPrice{}

// ViewSlashPrice200ResponseDataInnerPrice struct for ViewSlashPrice200ResponseDataInnerPrice
type ViewSlashPrice200ResponseDataInnerPrice struct {
	Min *float64 `json:"min,omitempty"`
	// Formated
	MinFormated *string `json:"min_formated,omitempty"`
	Max *float64 `json:"max,omitempty"`
	// Formated
	MaxFormated *string `json:"max_formated,omitempty"`
}

// NewViewSlashPrice200ResponseDataInnerPrice instantiates a new ViewSlashPrice200ResponseDataInnerPrice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViewSlashPrice200ResponseDataInnerPrice() *ViewSlashPrice200ResponseDataInnerPrice {
	this := ViewSlashPrice200ResponseDataInnerPrice{}
	return &this
}

// NewViewSlashPrice200ResponseDataInnerPriceWithDefaults instantiates a new ViewSlashPrice200ResponseDataInnerPrice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViewSlashPrice200ResponseDataInnerPriceWithDefaults() *ViewSlashPrice200ResponseDataInnerPrice {
	this := ViewSlashPrice200ResponseDataInnerPrice{}
	return &this
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *ViewSlashPrice200ResponseDataInnerPrice) GetMin() float64 {
	if o == nil || IsNil(o.Min) {
		var ret float64
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewSlashPrice200ResponseDataInnerPrice) GetMinOk() (*float64, bool) {
	if o == nil || IsNil(o.Min) {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *ViewSlashPrice200ResponseDataInnerPrice) HasMin() bool {
	if o != nil && !IsNil(o.Min) {
		return true
	}

	return false
}

// SetMin gets a reference to the given float64 and assigns it to the Min field.
func (o *ViewSlashPrice200ResponseDataInnerPrice) SetMin(v float64) {
	o.Min = &v
}

// GetMinFormated returns the MinFormated field value if set, zero value otherwise.
func (o *ViewSlashPrice200ResponseDataInnerPrice) GetMinFormated() string {
	if o == nil || IsNil(o.MinFormated) {
		var ret string
		return ret
	}
	return *o.MinFormated
}

// GetMinFormatedOk returns a tuple with the MinFormated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewSlashPrice200ResponseDataInnerPrice) GetMinFormatedOk() (*string, bool) {
	if o == nil || IsNil(o.MinFormated) {
		return nil, false
	}
	return o.MinFormated, true
}

// HasMinFormated returns a boolean if a field has been set.
func (o *ViewSlashPrice200ResponseDataInnerPrice) HasMinFormated() bool {
	if o != nil && !IsNil(o.MinFormated) {
		return true
	}

	return false
}

// SetMinFormated gets a reference to the given string and assigns it to the MinFormated field.
func (o *ViewSlashPrice200ResponseDataInnerPrice) SetMinFormated(v string) {
	o.MinFormated = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *ViewSlashPrice200ResponseDataInnerPrice) GetMax() float64 {
	if o == nil || IsNil(o.Max) {
		var ret float64
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewSlashPrice200ResponseDataInnerPrice) GetMaxOk() (*float64, bool) {
	if o == nil || IsNil(o.Max) {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *ViewSlashPrice200ResponseDataInnerPrice) HasMax() bool {
	if o != nil && !IsNil(o.Max) {
		return true
	}

	return false
}

// SetMax gets a reference to the given float64 and assigns it to the Max field.
func (o *ViewSlashPrice200ResponseDataInnerPrice) SetMax(v float64) {
	o.Max = &v
}

// GetMaxFormated returns the MaxFormated field value if set, zero value otherwise.
func (o *ViewSlashPrice200ResponseDataInnerPrice) GetMaxFormated() string {
	if o == nil || IsNil(o.MaxFormated) {
		var ret string
		return ret
	}
	return *o.MaxFormated
}

// GetMaxFormatedOk returns a tuple with the MaxFormated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewSlashPrice200ResponseDataInnerPrice) GetMaxFormatedOk() (*string, bool) {
	if o == nil || IsNil(o.MaxFormated) {
		return nil, false
	}
	return o.MaxFormated, true
}

// HasMaxFormated returns a boolean if a field has been set.
func (o *ViewSlashPrice200ResponseDataInnerPrice) HasMaxFormated() bool {
	if o != nil && !IsNil(o.MaxFormated) {
		return true
	}

	return false
}

// SetMaxFormated gets a reference to the given string and assigns it to the MaxFormated field.
func (o *ViewSlashPrice200ResponseDataInnerPrice) SetMaxFormated(v string) {
	o.MaxFormated = &v
}

func (o ViewSlashPrice200ResponseDataInnerPrice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ViewSlashPrice200ResponseDataInnerPrice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Min) {
		toSerialize["min"] = o.Min
	}
	if !IsNil(o.MinFormated) {
		toSerialize["min_formated"] = o.MinFormated
	}
	if !IsNil(o.Max) {
		toSerialize["max"] = o.Max
	}
	if !IsNil(o.MaxFormated) {
		toSerialize["max_formated"] = o.MaxFormated
	}
	return toSerialize, nil
}

type NullableViewSlashPrice200ResponseDataInnerPrice struct {
	value *ViewSlashPrice200ResponseDataInnerPrice
	isSet bool
}

func (v NullableViewSlashPrice200ResponseDataInnerPrice) Get() *ViewSlashPrice200ResponseDataInnerPrice {
	return v.value
}

func (v *NullableViewSlashPrice200ResponseDataInnerPrice) Set(val *ViewSlashPrice200ResponseDataInnerPrice) {
	v.value = val
	v.isSet = true
}

func (v NullableViewSlashPrice200ResponseDataInnerPrice) IsSet() bool {
	return v.isSet
}

func (v *NullableViewSlashPrice200ResponseDataInnerPrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViewSlashPrice200ResponseDataInnerPrice(val *ViewSlashPrice200ResponseDataInnerPrice) *NullableViewSlashPrice200ResponseDataInnerPrice {
	return &NullableViewSlashPrice200ResponseDataInnerPrice{value: val, isSet: true}
}

func (v NullableViewSlashPrice200ResponseDataInnerPrice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableViewSlashPrice200ResponseDataInnerPrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


