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

// checks if the GetBundleListDefaultResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetBundleListDefaultResponse{}

// GetBundleListDefaultResponse struct for GetBundleListDefaultResponse
type GetBundleListDefaultResponse struct {
	Header *GetOrderWebhook200ResponseHeader `json:"header,omitempty"`
	Data map[string]interface{} `json:"data,omitempty"`
}

// NewGetBundleListDefaultResponse instantiates a new GetBundleListDefaultResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBundleListDefaultResponse() *GetBundleListDefaultResponse {
	this := GetBundleListDefaultResponse{}
	return &this
}

// NewGetBundleListDefaultResponseWithDefaults instantiates a new GetBundleListDefaultResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBundleListDefaultResponseWithDefaults() *GetBundleListDefaultResponse {
	this := GetBundleListDefaultResponse{}
	return &this
}

// GetHeader returns the Header field value if set, zero value otherwise.
func (o *GetBundleListDefaultResponse) GetHeader() GetOrderWebhook200ResponseHeader {
	if o == nil || IsNil(o.Header) {
		var ret GetOrderWebhook200ResponseHeader
		return ret
	}
	return *o.Header
}

// GetHeaderOk returns a tuple with the Header field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBundleListDefaultResponse) GetHeaderOk() (*GetOrderWebhook200ResponseHeader, bool) {
	if o == nil || IsNil(o.Header) {
		return nil, false
	}
	return o.Header, true
}

// HasHeader returns a boolean if a field has been set.
func (o *GetBundleListDefaultResponse) HasHeader() bool {
	if o != nil && !IsNil(o.Header) {
		return true
	}

	return false
}

// SetHeader gets a reference to the given GetOrderWebhook200ResponseHeader and assigns it to the Header field.
func (o *GetBundleListDefaultResponse) SetHeader(v GetOrderWebhook200ResponseHeader) {
	o.Header = &v
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetBundleListDefaultResponse) GetData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetBundleListDefaultResponse) GetDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return map[string]interface{}{}, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetBundleListDefaultResponse) HasData() bool {
	if o != nil && IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *GetBundleListDefaultResponse) SetData(v map[string]interface{}) {
	o.Data = v
}

func (o GetBundleListDefaultResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBundleListDefaultResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Header) {
		toSerialize["header"] = o.Header
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGetBundleListDefaultResponse struct {
	value *GetBundleListDefaultResponse
	isSet bool
}

func (v NullableGetBundleListDefaultResponse) Get() *GetBundleListDefaultResponse {
	return v.value
}

func (v *NullableGetBundleListDefaultResponse) Set(val *GetBundleListDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBundleListDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBundleListDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBundleListDefaultResponse(val *GetBundleListDefaultResponse) *NullableGetBundleListDefaultResponse {
	return &NullableGetBundleListDefaultResponse{value: val, isSet: true}
}

func (v NullableGetBundleListDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBundleListDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


