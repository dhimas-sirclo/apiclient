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

// checks if the CancelSlashPrice200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CancelSlashPrice200Response{}

// CancelSlashPrice200Response struct for CancelSlashPrice200Response
type CancelSlashPrice200Response struct {
	Header *ResponseHeader `json:"header,omitempty"`
	Data *CancelSlashPrice200ResponseData `json:"data,omitempty"`
}

// NewCancelSlashPrice200Response instantiates a new CancelSlashPrice200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCancelSlashPrice200Response() *CancelSlashPrice200Response {
	this := CancelSlashPrice200Response{}
	return &this
}

// NewCancelSlashPrice200ResponseWithDefaults instantiates a new CancelSlashPrice200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCancelSlashPrice200ResponseWithDefaults() *CancelSlashPrice200Response {
	this := CancelSlashPrice200Response{}
	return &this
}

// GetHeader returns the Header field value if set, zero value otherwise.
func (o *CancelSlashPrice200Response) GetHeader() ResponseHeader {
	if o == nil || IsNil(o.Header) {
		var ret ResponseHeader
		return ret
	}
	return *o.Header
}

// GetHeaderOk returns a tuple with the Header field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelSlashPrice200Response) GetHeaderOk() (*ResponseHeader, bool) {
	if o == nil || IsNil(o.Header) {
		return nil, false
	}
	return o.Header, true
}

// HasHeader returns a boolean if a field has been set.
func (o *CancelSlashPrice200Response) HasHeader() bool {
	if o != nil && !IsNil(o.Header) {
		return true
	}

	return false
}

// SetHeader gets a reference to the given ResponseHeader and assigns it to the Header field.
func (o *CancelSlashPrice200Response) SetHeader(v ResponseHeader) {
	o.Header = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CancelSlashPrice200Response) GetData() CancelSlashPrice200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret CancelSlashPrice200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelSlashPrice200Response) GetDataOk() (*CancelSlashPrice200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CancelSlashPrice200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given CancelSlashPrice200ResponseData and assigns it to the Data field.
func (o *CancelSlashPrice200Response) SetData(v CancelSlashPrice200ResponseData) {
	o.Data = &v
}

func (o CancelSlashPrice200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CancelSlashPrice200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Header) {
		toSerialize["header"] = o.Header
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableCancelSlashPrice200Response struct {
	value *CancelSlashPrice200Response
	isSet bool
}

func (v NullableCancelSlashPrice200Response) Get() *CancelSlashPrice200Response {
	return v.value
}

func (v *NullableCancelSlashPrice200Response) Set(val *CancelSlashPrice200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCancelSlashPrice200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCancelSlashPrice200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCancelSlashPrice200Response(val *CancelSlashPrice200Response) *NullableCancelSlashPrice200Response {
	return &NullableCancelSlashPrice200Response{value: val, isSet: true}
}

func (v NullableCancelSlashPrice200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCancelSlashPrice200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


