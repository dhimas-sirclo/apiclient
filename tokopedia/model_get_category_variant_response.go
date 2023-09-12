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

// checks if the GetCategoryVariantResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCategoryVariantResponse{}

// GetCategoryVariantResponse struct for GetCategoryVariantResponse
type GetCategoryVariantResponse struct {
	Header *Header `json:"header,omitempty"`
	Data *CategoryVariant `json:"data,omitempty"`
}

// NewGetCategoryVariantResponse instantiates a new GetCategoryVariantResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCategoryVariantResponse() *GetCategoryVariantResponse {
	this := GetCategoryVariantResponse{}
	return &this
}

// NewGetCategoryVariantResponseWithDefaults instantiates a new GetCategoryVariantResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCategoryVariantResponseWithDefaults() *GetCategoryVariantResponse {
	this := GetCategoryVariantResponse{}
	return &this
}

// GetHeader returns the Header field value if set, zero value otherwise.
func (o *GetCategoryVariantResponse) GetHeader() Header {
	if o == nil || IsNil(o.Header) {
		var ret Header
		return ret
	}
	return *o.Header
}

// GetHeaderOk returns a tuple with the Header field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCategoryVariantResponse) GetHeaderOk() (*Header, bool) {
	if o == nil || IsNil(o.Header) {
		return nil, false
	}
	return o.Header, true
}

// HasHeader returns a boolean if a field has been set.
func (o *GetCategoryVariantResponse) HasHeader() bool {
	if o != nil && !IsNil(o.Header) {
		return true
	}

	return false
}

// SetHeader gets a reference to the given Header and assigns it to the Header field.
func (o *GetCategoryVariantResponse) SetHeader(v Header) {
	o.Header = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetCategoryVariantResponse) GetData() CategoryVariant {
	if o == nil || IsNil(o.Data) {
		var ret CategoryVariant
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCategoryVariantResponse) GetDataOk() (*CategoryVariant, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetCategoryVariantResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given CategoryVariant and assigns it to the Data field.
func (o *GetCategoryVariantResponse) SetData(v CategoryVariant) {
	o.Data = &v
}

func (o GetCategoryVariantResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCategoryVariantResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Header) {
		toSerialize["header"] = o.Header
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGetCategoryVariantResponse struct {
	value *GetCategoryVariantResponse
	isSet bool
}

func (v NullableGetCategoryVariantResponse) Get() *GetCategoryVariantResponse {
	return v.value
}

func (v *NullableGetCategoryVariantResponse) Set(val *GetCategoryVariantResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCategoryVariantResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCategoryVariantResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCategoryVariantResponse(val *GetCategoryVariantResponse) *NullableGetCategoryVariantResponse {
	return &NullableGetCategoryVariantResponse{value: val, isSet: true}
}

func (v NullableGetCategoryVariantResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCategoryVariantResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

