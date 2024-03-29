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

// checks if the GetProductInfoResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetProductInfoResponse{}

// GetProductInfoResponse struct for GetProductInfoResponse
type GetProductInfoResponse struct {
	Header *ResponseHeader `json:"header,omitempty"`
	Data []Product `json:"data,omitempty"`
}

// NewGetProductInfoResponse instantiates a new GetProductInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetProductInfoResponse() *GetProductInfoResponse {
	this := GetProductInfoResponse{}
	return &this
}

// NewGetProductInfoResponseWithDefaults instantiates a new GetProductInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetProductInfoResponseWithDefaults() *GetProductInfoResponse {
	this := GetProductInfoResponse{}
	return &this
}

// GetHeader returns the Header field value if set, zero value otherwise.
func (o *GetProductInfoResponse) GetHeader() ResponseHeader {
	if o == nil || IsNil(o.Header) {
		var ret ResponseHeader
		return ret
	}
	return *o.Header
}

// GetHeaderOk returns a tuple with the Header field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductInfoResponse) GetHeaderOk() (*ResponseHeader, bool) {
	if o == nil || IsNil(o.Header) {
		return nil, false
	}
	return o.Header, true
}

// HasHeader returns a boolean if a field has been set.
func (o *GetProductInfoResponse) HasHeader() bool {
	if o != nil && !IsNil(o.Header) {
		return true
	}

	return false
}

// SetHeader gets a reference to the given ResponseHeader and assigns it to the Header field.
func (o *GetProductInfoResponse) SetHeader(v ResponseHeader) {
	o.Header = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetProductInfoResponse) GetData() []Product {
	if o == nil || IsNil(o.Data) {
		var ret []Product
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductInfoResponse) GetDataOk() ([]Product, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetProductInfoResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []Product and assigns it to the Data field.
func (o *GetProductInfoResponse) SetData(v []Product) {
	o.Data = v
}

func (o GetProductInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetProductInfoResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Header) {
		toSerialize["header"] = o.Header
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGetProductInfoResponse struct {
	value *GetProductInfoResponse
	isSet bool
}

func (v NullableGetProductInfoResponse) Get() *GetProductInfoResponse {
	return v.value
}

func (v *NullableGetProductInfoResponse) Set(val *GetProductInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetProductInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetProductInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetProductInfoResponse(val *GetProductInfoResponse) *NullableGetProductInfoResponse {
	return &NullableGetProductInfoResponse{value: val, isSet: true}
}

func (v NullableGetProductInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetProductInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


