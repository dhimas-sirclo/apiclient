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

// checks if the GetShowcaseAllEtalase200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetShowcaseAllEtalase200Response{}

// GetShowcaseAllEtalase200Response struct for GetShowcaseAllEtalase200Response
type GetShowcaseAllEtalase200Response struct {
	Header *ResponseHeader `json:"header,omitempty"`
	Data *GetShowcaseAllEtalase200ResponseData `json:"data,omitempty"`
}

// NewGetShowcaseAllEtalase200Response instantiates a new GetShowcaseAllEtalase200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetShowcaseAllEtalase200Response() *GetShowcaseAllEtalase200Response {
	this := GetShowcaseAllEtalase200Response{}
	return &this
}

// NewGetShowcaseAllEtalase200ResponseWithDefaults instantiates a new GetShowcaseAllEtalase200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetShowcaseAllEtalase200ResponseWithDefaults() *GetShowcaseAllEtalase200Response {
	this := GetShowcaseAllEtalase200Response{}
	return &this
}

// GetHeader returns the Header field value if set, zero value otherwise.
func (o *GetShowcaseAllEtalase200Response) GetHeader() ResponseHeader {
	if o == nil || IsNil(o.Header) {
		var ret ResponseHeader
		return ret
	}
	return *o.Header
}

// GetHeaderOk returns a tuple with the Header field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetShowcaseAllEtalase200Response) GetHeaderOk() (*ResponseHeader, bool) {
	if o == nil || IsNil(o.Header) {
		return nil, false
	}
	return o.Header, true
}

// HasHeader returns a boolean if a field has been set.
func (o *GetShowcaseAllEtalase200Response) HasHeader() bool {
	if o != nil && !IsNil(o.Header) {
		return true
	}

	return false
}

// SetHeader gets a reference to the given ResponseHeader and assigns it to the Header field.
func (o *GetShowcaseAllEtalase200Response) SetHeader(v ResponseHeader) {
	o.Header = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetShowcaseAllEtalase200Response) GetData() GetShowcaseAllEtalase200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret GetShowcaseAllEtalase200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetShowcaseAllEtalase200Response) GetDataOk() (*GetShowcaseAllEtalase200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetShowcaseAllEtalase200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GetShowcaseAllEtalase200ResponseData and assigns it to the Data field.
func (o *GetShowcaseAllEtalase200Response) SetData(v GetShowcaseAllEtalase200ResponseData) {
	o.Data = &v
}

func (o GetShowcaseAllEtalase200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetShowcaseAllEtalase200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Header) {
		toSerialize["header"] = o.Header
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGetShowcaseAllEtalase200Response struct {
	value *GetShowcaseAllEtalase200Response
	isSet bool
}

func (v NullableGetShowcaseAllEtalase200Response) Get() *GetShowcaseAllEtalase200Response {
	return v.value
}

func (v *NullableGetShowcaseAllEtalase200Response) Set(val *GetShowcaseAllEtalase200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetShowcaseAllEtalase200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetShowcaseAllEtalase200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetShowcaseAllEtalase200Response(val *GetShowcaseAllEtalase200Response) *NullableGetShowcaseAllEtalase200Response {
	return &NullableGetShowcaseAllEtalase200Response{value: val, isSet: true}
}

func (v NullableGetShowcaseAllEtalase200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetShowcaseAllEtalase200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


