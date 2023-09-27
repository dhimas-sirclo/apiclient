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

// checks if the CreateShowcaseRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateShowcaseRequest{}

// CreateShowcaseRequest struct for CreateShowcaseRequest
type CreateShowcaseRequest struct {
	// Showcase name to be updated
	Name string `json:"Name"`
}

// NewCreateShowcaseRequest instantiates a new CreateShowcaseRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateShowcaseRequest(name string) *CreateShowcaseRequest {
	this := CreateShowcaseRequest{}
	this.Name = name
	return &this
}

// NewCreateShowcaseRequestWithDefaults instantiates a new CreateShowcaseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateShowcaseRequestWithDefaults() *CreateShowcaseRequest {
	this := CreateShowcaseRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateShowcaseRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateShowcaseRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateShowcaseRequest) SetName(v string) {
	o.Name = v
}

func (o CreateShowcaseRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateShowcaseRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Name"] = o.Name
	return toSerialize, nil
}

type NullableCreateShowcaseRequest struct {
	value *CreateShowcaseRequest
	isSet bool
}

func (v NullableCreateShowcaseRequest) Get() *CreateShowcaseRequest {
	return v.value
}

func (v *NullableCreateShowcaseRequest) Set(val *CreateShowcaseRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateShowcaseRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateShowcaseRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateShowcaseRequest(val *CreateShowcaseRequest) *NullableCreateShowcaseRequest {
	return &NullableCreateShowcaseRequest{value: val, isSet: true}
}

func (v NullableCreateShowcaseRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateShowcaseRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


