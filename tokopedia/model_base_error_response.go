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

// checks if the BaseErrorResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseErrorResponse{}

// BaseErrorResponse struct for BaseErrorResponse
type BaseErrorResponse struct {
	Header *ResponseHeader `json:"header,omitempty"`
	Data map[string]interface{} `json:"data,omitempty"`
}

// NewBaseErrorResponse instantiates a new BaseErrorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseErrorResponse() *BaseErrorResponse {
	this := BaseErrorResponse{}
	return &this
}

// NewBaseErrorResponseWithDefaults instantiates a new BaseErrorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseErrorResponseWithDefaults() *BaseErrorResponse {
	this := BaseErrorResponse{}
	return &this
}

// GetHeader returns the Header field value if set, zero value otherwise.
func (o *BaseErrorResponse) GetHeader() ResponseHeader {
	if o == nil || IsNil(o.Header) {
		var ret ResponseHeader
		return ret
	}
	return *o.Header
}

// GetHeaderOk returns a tuple with the Header field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseErrorResponse) GetHeaderOk() (*ResponseHeader, bool) {
	if o == nil || IsNil(o.Header) {
		return nil, false
	}
	return o.Header, true
}

// HasHeader returns a boolean if a field has been set.
func (o *BaseErrorResponse) HasHeader() bool {
	if o != nil && !IsNil(o.Header) {
		return true
	}

	return false
}

// SetHeader gets a reference to the given ResponseHeader and assigns it to the Header field.
func (o *BaseErrorResponse) SetHeader(v ResponseHeader) {
	o.Header = &v
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BaseErrorResponse) GetData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BaseErrorResponse) GetDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return map[string]interface{}{}, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *BaseErrorResponse) HasData() bool {
	if o != nil && IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *BaseErrorResponse) SetData(v map[string]interface{}) {
	o.Data = v
}

func (o BaseErrorResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseErrorResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Header) {
		toSerialize["header"] = o.Header
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableBaseErrorResponse struct {
	value *BaseErrorResponse
	isSet bool
}

func (v NullableBaseErrorResponse) Get() *BaseErrorResponse {
	return v.value
}

func (v *NullableBaseErrorResponse) Set(val *BaseErrorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseErrorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseErrorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseErrorResponse(val *BaseErrorResponse) *NullableBaseErrorResponse {
	return &NullableBaseErrorResponse{value: val, isSet: true}
}

func (v NullableBaseErrorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseErrorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


