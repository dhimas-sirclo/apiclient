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

// checks if the CreateShowcase200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateShowcase200Response{}

// CreateShowcase200Response struct for CreateShowcase200Response
type CreateShowcase200Response struct {
	Header *ResponseHeader `json:"header,omitempty"`
	Data *CreateShowcase200ResponseData `json:"data,omitempty"`
}

// NewCreateShowcase200Response instantiates a new CreateShowcase200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateShowcase200Response() *CreateShowcase200Response {
	this := CreateShowcase200Response{}
	return &this
}

// NewCreateShowcase200ResponseWithDefaults instantiates a new CreateShowcase200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateShowcase200ResponseWithDefaults() *CreateShowcase200Response {
	this := CreateShowcase200Response{}
	return &this
}

// GetHeader returns the Header field value if set, zero value otherwise.
func (o *CreateShowcase200Response) GetHeader() ResponseHeader {
	if o == nil || IsNil(o.Header) {
		var ret ResponseHeader
		return ret
	}
	return *o.Header
}

// GetHeaderOk returns a tuple with the Header field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShowcase200Response) GetHeaderOk() (*ResponseHeader, bool) {
	if o == nil || IsNil(o.Header) {
		return nil, false
	}
	return o.Header, true
}

// HasHeader returns a boolean if a field has been set.
func (o *CreateShowcase200Response) HasHeader() bool {
	if o != nil && !IsNil(o.Header) {
		return true
	}

	return false
}

// SetHeader gets a reference to the given ResponseHeader and assigns it to the Header field.
func (o *CreateShowcase200Response) SetHeader(v ResponseHeader) {
	o.Header = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CreateShowcase200Response) GetData() CreateShowcase200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret CreateShowcase200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShowcase200Response) GetDataOk() (*CreateShowcase200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CreateShowcase200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given CreateShowcase200ResponseData and assigns it to the Data field.
func (o *CreateShowcase200Response) SetData(v CreateShowcase200ResponseData) {
	o.Data = &v
}

func (o CreateShowcase200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateShowcase200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Header) {
		toSerialize["header"] = o.Header
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableCreateShowcase200Response struct {
	value *CreateShowcase200Response
	isSet bool
}

func (v NullableCreateShowcase200Response) Get() *CreateShowcase200Response {
	return v.value
}

func (v *NullableCreateShowcase200Response) Set(val *CreateShowcase200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateShowcase200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateShowcase200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateShowcase200Response(val *CreateShowcase200Response) *NullableCreateShowcase200Response {
	return &NullableCreateShowcase200Response{value: val, isSet: true}
}

func (v NullableCreateShowcase200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateShowcase200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

