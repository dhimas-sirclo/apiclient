/*
Tokopedia

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
Contact: dev@sirclo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tokopedia

import (
	"encoding/json"
)

// checks if the GetVariantResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetVariantResponse{}

// GetVariantResponse struct for GetVariantResponse
type GetVariantResponse struct {
	Header *Header `json:"header,omitempty"`
	Data *Variant `json:"data,omitempty"`
}

// NewGetVariantResponse instantiates a new GetVariantResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetVariantResponse() *GetVariantResponse {
	this := GetVariantResponse{}
	return &this
}

// NewGetVariantResponseWithDefaults instantiates a new GetVariantResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetVariantResponseWithDefaults() *GetVariantResponse {
	this := GetVariantResponse{}
	return &this
}

// GetHeader returns the Header field value if set, zero value otherwise.
func (o *GetVariantResponse) GetHeader() Header {
	if o == nil || IsNil(o.Header) {
		var ret Header
		return ret
	}
	return *o.Header
}

// GetHeaderOk returns a tuple with the Header field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVariantResponse) GetHeaderOk() (*Header, bool) {
	if o == nil || IsNil(o.Header) {
		return nil, false
	}
	return o.Header, true
}

// HasHeader returns a boolean if a field has been set.
func (o *GetVariantResponse) HasHeader() bool {
	if o != nil && !IsNil(o.Header) {
		return true
	}

	return false
}

// SetHeader gets a reference to the given Header and assigns it to the Header field.
func (o *GetVariantResponse) SetHeader(v Header) {
	o.Header = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetVariantResponse) GetData() Variant {
	if o == nil || IsNil(o.Data) {
		var ret Variant
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVariantResponse) GetDataOk() (*Variant, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetVariantResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given Variant and assigns it to the Data field.
func (o *GetVariantResponse) SetData(v Variant) {
	o.Data = &v
}

func (o GetVariantResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetVariantResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Header) {
		toSerialize["header"] = o.Header
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGetVariantResponse struct {
	value *GetVariantResponse
	isSet bool
}

func (v NullableGetVariantResponse) Get() *GetVariantResponse {
	return v.value
}

func (v *NullableGetVariantResponse) Set(val *GetVariantResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetVariantResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetVariantResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetVariantResponse(val *GetVariantResponse) *NullableGetVariantResponse {
	return &NullableGetVariantResponse{value: val, isSet: true}
}

func (v NullableGetVariantResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetVariantResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


