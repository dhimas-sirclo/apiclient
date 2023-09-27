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

// checks if the GetBundleInfo200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetBundleInfo200Response{}

// GetBundleInfo200Response struct for GetBundleInfo200Response
type GetBundleInfo200Response struct {
	Header *ResponseHeader `json:"header,omitempty"`
	Data *GetBundleInfo200ResponseData `json:"data,omitempty"`
}

// NewGetBundleInfo200Response instantiates a new GetBundleInfo200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBundleInfo200Response() *GetBundleInfo200Response {
	this := GetBundleInfo200Response{}
	return &this
}

// NewGetBundleInfo200ResponseWithDefaults instantiates a new GetBundleInfo200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBundleInfo200ResponseWithDefaults() *GetBundleInfo200Response {
	this := GetBundleInfo200Response{}
	return &this
}

// GetHeader returns the Header field value if set, zero value otherwise.
func (o *GetBundleInfo200Response) GetHeader() ResponseHeader {
	if o == nil || IsNil(o.Header) {
		var ret ResponseHeader
		return ret
	}
	return *o.Header
}

// GetHeaderOk returns a tuple with the Header field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBundleInfo200Response) GetHeaderOk() (*ResponseHeader, bool) {
	if o == nil || IsNil(o.Header) {
		return nil, false
	}
	return o.Header, true
}

// HasHeader returns a boolean if a field has been set.
func (o *GetBundleInfo200Response) HasHeader() bool {
	if o != nil && !IsNil(o.Header) {
		return true
	}

	return false
}

// SetHeader gets a reference to the given ResponseHeader and assigns it to the Header field.
func (o *GetBundleInfo200Response) SetHeader(v ResponseHeader) {
	o.Header = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetBundleInfo200Response) GetData() GetBundleInfo200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret GetBundleInfo200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBundleInfo200Response) GetDataOk() (*GetBundleInfo200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetBundleInfo200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GetBundleInfo200ResponseData and assigns it to the Data field.
func (o *GetBundleInfo200Response) SetData(v GetBundleInfo200ResponseData) {
	o.Data = &v
}

func (o GetBundleInfo200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBundleInfo200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Header) {
		toSerialize["header"] = o.Header
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGetBundleInfo200Response struct {
	value *GetBundleInfo200Response
	isSet bool
}

func (v NullableGetBundleInfo200Response) Get() *GetBundleInfo200Response {
	return v.value
}

func (v *NullableGetBundleInfo200Response) Set(val *GetBundleInfo200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBundleInfo200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBundleInfo200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBundleInfo200Response(val *GetBundleInfo200Response) *NullableGetBundleInfo200Response {
	return &NullableGetBundleInfo200Response{value: val, isSet: true}
}

func (v NullableGetBundleInfo200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBundleInfo200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

