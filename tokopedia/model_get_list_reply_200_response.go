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

// checks if the GetListReply200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetListReply200Response{}

// GetListReply200Response struct for GetListReply200Response
type GetListReply200Response struct {
	Header *GetOrderWebhook200ResponseHeader `json:"header,omitempty"`
	Data []GetListReply200ResponseDataInner `json:"data,omitempty"`
}

// NewGetListReply200Response instantiates a new GetListReply200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetListReply200Response() *GetListReply200Response {
	this := GetListReply200Response{}
	return &this
}

// NewGetListReply200ResponseWithDefaults instantiates a new GetListReply200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetListReply200ResponseWithDefaults() *GetListReply200Response {
	this := GetListReply200Response{}
	return &this
}

// GetHeader returns the Header field value if set, zero value otherwise.
func (o *GetListReply200Response) GetHeader() GetOrderWebhook200ResponseHeader {
	if o == nil || IsNil(o.Header) {
		var ret GetOrderWebhook200ResponseHeader
		return ret
	}
	return *o.Header
}

// GetHeaderOk returns a tuple with the Header field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetListReply200Response) GetHeaderOk() (*GetOrderWebhook200ResponseHeader, bool) {
	if o == nil || IsNil(o.Header) {
		return nil, false
	}
	return o.Header, true
}

// HasHeader returns a boolean if a field has been set.
func (o *GetListReply200Response) HasHeader() bool {
	if o != nil && !IsNil(o.Header) {
		return true
	}

	return false
}

// SetHeader gets a reference to the given GetOrderWebhook200ResponseHeader and assigns it to the Header field.
func (o *GetListReply200Response) SetHeader(v GetOrderWebhook200ResponseHeader) {
	o.Header = &v
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetListReply200Response) GetData() []GetListReply200ResponseDataInner {
	if o == nil {
		var ret []GetListReply200ResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetListReply200Response) GetDataOk() ([]GetListReply200ResponseDataInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetListReply200Response) HasData() bool {
	if o != nil && IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []GetListReply200ResponseDataInner and assigns it to the Data field.
func (o *GetListReply200Response) SetData(v []GetListReply200ResponseDataInner) {
	o.Data = v
}

func (o GetListReply200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetListReply200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Header) {
		toSerialize["header"] = o.Header
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGetListReply200Response struct {
	value *GetListReply200Response
	isSet bool
}

func (v NullableGetListReply200Response) Get() *GetListReply200Response {
	return v.value
}

func (v *NullableGetListReply200Response) Set(val *GetListReply200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetListReply200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetListReply200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetListReply200Response(val *GetListReply200Response) *NullableGetListReply200Response {
	return &NullableGetListReply200Response{value: val, isSet: true}
}

func (v NullableGetListReply200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetListReply200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


