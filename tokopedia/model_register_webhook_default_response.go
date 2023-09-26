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

// checks if the RegisterWebhookDefaultResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegisterWebhookDefaultResponse{}

// RegisterWebhookDefaultResponse struct for RegisterWebhookDefaultResponse
type RegisterWebhookDefaultResponse struct {
	Data *string `json:"data,omitempty"`
	Status *string `json:"status,omitempty"`
	ErrorMessage []string `json:"error_message,omitempty"`
}

// NewRegisterWebhookDefaultResponse instantiates a new RegisterWebhookDefaultResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegisterWebhookDefaultResponse() *RegisterWebhookDefaultResponse {
	this := RegisterWebhookDefaultResponse{}
	return &this
}

// NewRegisterWebhookDefaultResponseWithDefaults instantiates a new RegisterWebhookDefaultResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegisterWebhookDefaultResponseWithDefaults() *RegisterWebhookDefaultResponse {
	this := RegisterWebhookDefaultResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *RegisterWebhookDefaultResponse) GetData() string {
	if o == nil || IsNil(o.Data) {
		var ret string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterWebhookDefaultResponse) GetDataOk() (*string, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *RegisterWebhookDefaultResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given string and assigns it to the Data field.
func (o *RegisterWebhookDefaultResponse) SetData(v string) {
	o.Data = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *RegisterWebhookDefaultResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterWebhookDefaultResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *RegisterWebhookDefaultResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *RegisterWebhookDefaultResponse) SetStatus(v string) {
	o.Status = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *RegisterWebhookDefaultResponse) GetErrorMessage() []string {
	if o == nil || IsNil(o.ErrorMessage) {
		var ret []string
		return ret
	}
	return o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterWebhookDefaultResponse) GetErrorMessageOk() ([]string, bool) {
	if o == nil || IsNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *RegisterWebhookDefaultResponse) HasErrorMessage() bool {
	if o != nil && !IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given []string and assigns it to the ErrorMessage field.
func (o *RegisterWebhookDefaultResponse) SetErrorMessage(v []string) {
	o.ErrorMessage = v
}

func (o RegisterWebhookDefaultResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegisterWebhookDefaultResponse) ToMap() (map[string]interface{}, error) {
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

type NullableRegisterWebhookDefaultResponse struct {
	value *RegisterWebhookDefaultResponse
	isSet bool
}

func (v NullableRegisterWebhookDefaultResponse) Get() *RegisterWebhookDefaultResponse {
	return v.value
}

func (v *NullableRegisterWebhookDefaultResponse) Set(val *RegisterWebhookDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRegisterWebhookDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRegisterWebhookDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegisterWebhookDefaultResponse(val *RegisterWebhookDefaultResponse) *NullableRegisterWebhookDefaultResponse {
	return &NullableRegisterWebhookDefaultResponse{value: val, isSet: true}
}

func (v NullableRegisterWebhookDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegisterWebhookDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

