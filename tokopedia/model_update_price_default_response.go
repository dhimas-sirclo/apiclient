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

// checks if the UpdatePriceDefaultResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdatePriceDefaultResponse{}

// UpdatePriceDefaultResponse struct for UpdatePriceDefaultResponse
type UpdatePriceDefaultResponse struct {
	Header *ResponseHeader `json:"header,omitempty"`
	Data *UpdatePriceDefaultResponseData `json:"data,omitempty"`
}

// NewUpdatePriceDefaultResponse instantiates a new UpdatePriceDefaultResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdatePriceDefaultResponse() *UpdatePriceDefaultResponse {
	this := UpdatePriceDefaultResponse{}
	return &this
}

// NewUpdatePriceDefaultResponseWithDefaults instantiates a new UpdatePriceDefaultResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdatePriceDefaultResponseWithDefaults() *UpdatePriceDefaultResponse {
	this := UpdatePriceDefaultResponse{}
	return &this
}

// GetHeader returns the Header field value if set, zero value otherwise.
func (o *UpdatePriceDefaultResponse) GetHeader() ResponseHeader {
	if o == nil || IsNil(o.Header) {
		var ret ResponseHeader
		return ret
	}
	return *o.Header
}

// GetHeaderOk returns a tuple with the Header field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePriceDefaultResponse) GetHeaderOk() (*ResponseHeader, bool) {
	if o == nil || IsNil(o.Header) {
		return nil, false
	}
	return o.Header, true
}

// HasHeader returns a boolean if a field has been set.
func (o *UpdatePriceDefaultResponse) HasHeader() bool {
	if o != nil && !IsNil(o.Header) {
		return true
	}

	return false
}

// SetHeader gets a reference to the given ResponseHeader and assigns it to the Header field.
func (o *UpdatePriceDefaultResponse) SetHeader(v ResponseHeader) {
	o.Header = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *UpdatePriceDefaultResponse) GetData() UpdatePriceDefaultResponseData {
	if o == nil || IsNil(o.Data) {
		var ret UpdatePriceDefaultResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePriceDefaultResponse) GetDataOk() (*UpdatePriceDefaultResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *UpdatePriceDefaultResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given UpdatePriceDefaultResponseData and assigns it to the Data field.
func (o *UpdatePriceDefaultResponse) SetData(v UpdatePriceDefaultResponseData) {
	o.Data = &v
}

func (o UpdatePriceDefaultResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdatePriceDefaultResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Header) {
		toSerialize["header"] = o.Header
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableUpdatePriceDefaultResponse struct {
	value *UpdatePriceDefaultResponse
	isSet bool
}

func (v NullableUpdatePriceDefaultResponse) Get() *UpdatePriceDefaultResponse {
	return v.value
}

func (v *NullableUpdatePriceDefaultResponse) Set(val *UpdatePriceDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdatePriceDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdatePriceDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdatePriceDefaultResponse(val *UpdatePriceDefaultResponse) *NullableUpdatePriceDefaultResponse {
	return &NullableUpdatePriceDefaultResponse{value: val, isSet: true}
}

func (v NullableUpdatePriceDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdatePriceDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

