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

// checks if the InitiateChat200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InitiateChat200Response{}

// InitiateChat200Response struct for InitiateChat200Response
type InitiateChat200Response struct {
	Header *GetListMessageDefaultResponseHeader `json:"header,omitempty"`
	Data NullableInitiateChat200ResponseData `json:"data,omitempty"`
}

// NewInitiateChat200Response instantiates a new InitiateChat200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInitiateChat200Response() *InitiateChat200Response {
	this := InitiateChat200Response{}
	return &this
}

// NewInitiateChat200ResponseWithDefaults instantiates a new InitiateChat200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInitiateChat200ResponseWithDefaults() *InitiateChat200Response {
	this := InitiateChat200Response{}
	return &this
}

// GetHeader returns the Header field value if set, zero value otherwise.
func (o *InitiateChat200Response) GetHeader() GetListMessageDefaultResponseHeader {
	if o == nil || IsNil(o.Header) {
		var ret GetListMessageDefaultResponseHeader
		return ret
	}
	return *o.Header
}

// GetHeaderOk returns a tuple with the Header field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InitiateChat200Response) GetHeaderOk() (*GetListMessageDefaultResponseHeader, bool) {
	if o == nil || IsNil(o.Header) {
		return nil, false
	}
	return o.Header, true
}

// HasHeader returns a boolean if a field has been set.
func (o *InitiateChat200Response) HasHeader() bool {
	if o != nil && !IsNil(o.Header) {
		return true
	}

	return false
}

// SetHeader gets a reference to the given GetListMessageDefaultResponseHeader and assigns it to the Header field.
func (o *InitiateChat200Response) SetHeader(v GetListMessageDefaultResponseHeader) {
	o.Header = &v
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InitiateChat200Response) GetData() InitiateChat200ResponseData {
	if o == nil || IsNil(o.Data.Get()) {
		var ret InitiateChat200ResponseData
		return ret
	}
	return *o.Data.Get()
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InitiateChat200Response) GetDataOk() (*InitiateChat200ResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data.Get(), o.Data.IsSet()
}

// HasData returns a boolean if a field has been set.
func (o *InitiateChat200Response) HasData() bool {
	if o != nil && o.Data.IsSet() {
		return true
	}

	return false
}

// SetData gets a reference to the given NullableInitiateChat200ResponseData and assigns it to the Data field.
func (o *InitiateChat200Response) SetData(v InitiateChat200ResponseData) {
	o.Data.Set(&v)
}
// SetDataNil sets the value for Data to be an explicit nil
func (o *InitiateChat200Response) SetDataNil() {
	o.Data.Set(nil)
}

// UnsetData ensures that no value is present for Data, not even an explicit nil
func (o *InitiateChat200Response) UnsetData() {
	o.Data.Unset()
}

func (o InitiateChat200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InitiateChat200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Header) {
		toSerialize["header"] = o.Header
	}
	if o.Data.IsSet() {
		toSerialize["data"] = o.Data.Get()
	}
	return toSerialize, nil
}

type NullableInitiateChat200Response struct {
	value *InitiateChat200Response
	isSet bool
}

func (v NullableInitiateChat200Response) Get() *InitiateChat200Response {
	return v.value
}

func (v *NullableInitiateChat200Response) Set(val *InitiateChat200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableInitiateChat200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableInitiateChat200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInitiateChat200Response(val *InitiateChat200Response) *NullableInitiateChat200Response {
	return &NullableInitiateChat200Response{value: val, isSet: true}
}

func (v NullableInitiateChat200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInitiateChat200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


