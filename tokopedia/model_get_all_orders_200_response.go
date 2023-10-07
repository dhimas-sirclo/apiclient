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

// checks if the GetAllOrders200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAllOrders200Response{}

// GetAllOrders200Response struct for GetAllOrders200Response
type GetAllOrders200Response struct {
	Header *ResponseHeader `json:"header,omitempty"`
	Data []GetAllOrders200ResponseDataInner `json:"data,omitempty"`
}

// NewGetAllOrders200Response instantiates a new GetAllOrders200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAllOrders200Response() *GetAllOrders200Response {
	this := GetAllOrders200Response{}
	return &this
}

// NewGetAllOrders200ResponseWithDefaults instantiates a new GetAllOrders200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAllOrders200ResponseWithDefaults() *GetAllOrders200Response {
	this := GetAllOrders200Response{}
	return &this
}

// GetHeader returns the Header field value if set, zero value otherwise.
func (o *GetAllOrders200Response) GetHeader() ResponseHeader {
	if o == nil || IsNil(o.Header) {
		var ret ResponseHeader
		return ret
	}
	return *o.Header
}

// GetHeaderOk returns a tuple with the Header field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllOrders200Response) GetHeaderOk() (*ResponseHeader, bool) {
	if o == nil || IsNil(o.Header) {
		return nil, false
	}
	return o.Header, true
}

// HasHeader returns a boolean if a field has been set.
func (o *GetAllOrders200Response) HasHeader() bool {
	if o != nil && !IsNil(o.Header) {
		return true
	}

	return false
}

// SetHeader gets a reference to the given ResponseHeader and assigns it to the Header field.
func (o *GetAllOrders200Response) SetHeader(v ResponseHeader) {
	o.Header = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetAllOrders200Response) GetData() []GetAllOrders200ResponseDataInner {
	if o == nil || IsNil(o.Data) {
		var ret []GetAllOrders200ResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllOrders200Response) GetDataOk() ([]GetAllOrders200ResponseDataInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetAllOrders200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []GetAllOrders200ResponseDataInner and assigns it to the Data field.
func (o *GetAllOrders200Response) SetData(v []GetAllOrders200ResponseDataInner) {
	o.Data = v
}

func (o GetAllOrders200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAllOrders200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Header) {
		toSerialize["header"] = o.Header
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGetAllOrders200Response struct {
	value *GetAllOrders200Response
	isSet bool
}

func (v NullableGetAllOrders200Response) Get() *GetAllOrders200Response {
	return v.value
}

func (v *NullableGetAllOrders200Response) Set(val *GetAllOrders200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAllOrders200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAllOrders200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAllOrders200Response(val *GetAllOrders200Response) *NullableGetAllOrders200Response {
	return &NullableGetAllOrders200Response{value: val, isSet: true}
}

func (v NullableGetAllOrders200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAllOrders200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


