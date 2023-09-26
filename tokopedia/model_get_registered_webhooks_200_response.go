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

// checks if the GetRegisteredWebhooks200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRegisteredWebhooks200Response{}

// GetRegisteredWebhooks200Response struct for GetRegisteredWebhooks200Response
type GetRegisteredWebhooks200Response struct {
	Data *GetRegisteredWebhooks200ResponseData `json:"data,omitempty"`
	Status *string `json:"status,omitempty"`
	ErrorMessage []string `json:"error_message,omitempty"`
}

// NewGetRegisteredWebhooks200Response instantiates a new GetRegisteredWebhooks200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRegisteredWebhooks200Response() *GetRegisteredWebhooks200Response {
	this := GetRegisteredWebhooks200Response{}
	return &this
}

// NewGetRegisteredWebhooks200ResponseWithDefaults instantiates a new GetRegisteredWebhooks200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRegisteredWebhooks200ResponseWithDefaults() *GetRegisteredWebhooks200Response {
	this := GetRegisteredWebhooks200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetRegisteredWebhooks200Response) GetData() GetRegisteredWebhooks200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret GetRegisteredWebhooks200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRegisteredWebhooks200Response) GetDataOk() (*GetRegisteredWebhooks200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetRegisteredWebhooks200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GetRegisteredWebhooks200ResponseData and assigns it to the Data field.
func (o *GetRegisteredWebhooks200Response) SetData(v GetRegisteredWebhooks200ResponseData) {
	o.Data = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetRegisteredWebhooks200Response) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRegisteredWebhooks200Response) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetRegisteredWebhooks200Response) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetRegisteredWebhooks200Response) SetStatus(v string) {
	o.Status = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *GetRegisteredWebhooks200Response) GetErrorMessage() []string {
	if o == nil || IsNil(o.ErrorMessage) {
		var ret []string
		return ret
	}
	return o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRegisteredWebhooks200Response) GetErrorMessageOk() ([]string, bool) {
	if o == nil || IsNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *GetRegisteredWebhooks200Response) HasErrorMessage() bool {
	if o != nil && !IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given []string and assigns it to the ErrorMessage field.
func (o *GetRegisteredWebhooks200Response) SetErrorMessage(v []string) {
	o.ErrorMessage = v
}

func (o GetRegisteredWebhooks200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRegisteredWebhooks200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.ErrorMessage) {
		toSerialize["error_message"] = o.ErrorMessage
	}
	return toSerialize, nil
}

type NullableGetRegisteredWebhooks200Response struct {
	value *GetRegisteredWebhooks200Response
	isSet bool
}

func (v NullableGetRegisteredWebhooks200Response) Get() *GetRegisteredWebhooks200Response {
	return v.value
}

func (v *NullableGetRegisteredWebhooks200Response) Set(val *GetRegisteredWebhooks200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRegisteredWebhooks200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRegisteredWebhooks200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRegisteredWebhooks200Response(val *GetRegisteredWebhooks200Response) *NullableGetRegisteredWebhooks200Response {
	return &NullableGetRegisteredWebhooks200Response{value: val, isSet: true}
}

func (v NullableGetRegisteredWebhooks200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRegisteredWebhooks200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

