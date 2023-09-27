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

// checks if the GetSaldoHistory200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSaldoHistory200Response{}

// GetSaldoHistory200Response struct for GetSaldoHistory200Response
type GetSaldoHistory200Response struct {
	Header *GetOrderWebhook200ResponseHeader `json:"header,omitempty"`
	Data NullableGetSaldoHistory200ResponseData `json:"data,omitempty"`
}

// NewGetSaldoHistory200Response instantiates a new GetSaldoHistory200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSaldoHistory200Response() *GetSaldoHistory200Response {
	this := GetSaldoHistory200Response{}
	return &this
}

// NewGetSaldoHistory200ResponseWithDefaults instantiates a new GetSaldoHistory200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSaldoHistory200ResponseWithDefaults() *GetSaldoHistory200Response {
	this := GetSaldoHistory200Response{}
	return &this
}

// GetHeader returns the Header field value if set, zero value otherwise.
func (o *GetSaldoHistory200Response) GetHeader() GetOrderWebhook200ResponseHeader {
	if o == nil || IsNil(o.Header) {
		var ret GetOrderWebhook200ResponseHeader
		return ret
	}
	return *o.Header
}

// GetHeaderOk returns a tuple with the Header field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSaldoHistory200Response) GetHeaderOk() (*GetOrderWebhook200ResponseHeader, bool) {
	if o == nil || IsNil(o.Header) {
		return nil, false
	}
	return o.Header, true
}

// HasHeader returns a boolean if a field has been set.
func (o *GetSaldoHistory200Response) HasHeader() bool {
	if o != nil && !IsNil(o.Header) {
		return true
	}

	return false
}

// SetHeader gets a reference to the given GetOrderWebhook200ResponseHeader and assigns it to the Header field.
func (o *GetSaldoHistory200Response) SetHeader(v GetOrderWebhook200ResponseHeader) {
	o.Header = &v
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetSaldoHistory200Response) GetData() GetSaldoHistory200ResponseData {
	if o == nil || IsNil(o.Data.Get()) {
		var ret GetSaldoHistory200ResponseData
		return ret
	}
	return *o.Data.Get()
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetSaldoHistory200Response) GetDataOk() (*GetSaldoHistory200ResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data.Get(), o.Data.IsSet()
}

// HasData returns a boolean if a field has been set.
func (o *GetSaldoHistory200Response) HasData() bool {
	if o != nil && o.Data.IsSet() {
		return true
	}

	return false
}

// SetData gets a reference to the given NullableGetSaldoHistory200ResponseData and assigns it to the Data field.
func (o *GetSaldoHistory200Response) SetData(v GetSaldoHistory200ResponseData) {
	o.Data.Set(&v)
}
// SetDataNil sets the value for Data to be an explicit nil
func (o *GetSaldoHistory200Response) SetDataNil() {
	o.Data.Set(nil)
}

// UnsetData ensures that no value is present for Data, not even an explicit nil
func (o *GetSaldoHistory200Response) UnsetData() {
	o.Data.Unset()
}

func (o GetSaldoHistory200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSaldoHistory200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Header) {
		toSerialize["header"] = o.Header
	}
	if o.Data.IsSet() {
		toSerialize["data"] = o.Data.Get()
	}
	return toSerialize, nil
}

type NullableGetSaldoHistory200Response struct {
	value *GetSaldoHistory200Response
	isSet bool
}

func (v NullableGetSaldoHistory200Response) Get() *GetSaldoHistory200Response {
	return v.value
}

func (v *NullableGetSaldoHistory200Response) Set(val *GetSaldoHistory200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSaldoHistory200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSaldoHistory200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSaldoHistory200Response(val *GetSaldoHistory200Response) *NullableGetSaldoHistory200Response {
	return &NullableGetSaldoHistory200Response{value: val, isSet: true}
}

func (v NullableGetSaldoHistory200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSaldoHistory200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

