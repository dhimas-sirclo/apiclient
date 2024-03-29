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

// checks if the GetListMessageDefaultResponseDataInnerAttributesContact type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetListMessageDefaultResponseDataInnerAttributesContact{}

// GetListMessageDefaultResponseDataInnerAttributesContact struct for GetListMessageDefaultResponseDataInnerAttributesContact
type GetListMessageDefaultResponseDataInnerAttributesContact struct {
	// User Unique Identifier
	Id *int64 `json:"id,omitempty"`
	// User Role, Field role has User, Shop Owner, Shop Admin, and Tokopedia Admin.
	Role *string `json:"role,omitempty"`
	Attributes *GetListMessageDefaultResponseDataInnerAttributesContactAttributes `json:"attributes,omitempty"`
}

// NewGetListMessageDefaultResponseDataInnerAttributesContact instantiates a new GetListMessageDefaultResponseDataInnerAttributesContact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetListMessageDefaultResponseDataInnerAttributesContact() *GetListMessageDefaultResponseDataInnerAttributesContact {
	this := GetListMessageDefaultResponseDataInnerAttributesContact{}
	return &this
}

// NewGetListMessageDefaultResponseDataInnerAttributesContactWithDefaults instantiates a new GetListMessageDefaultResponseDataInnerAttributesContact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetListMessageDefaultResponseDataInnerAttributesContactWithDefaults() *GetListMessageDefaultResponseDataInnerAttributesContact {
	this := GetListMessageDefaultResponseDataInnerAttributesContact{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetListMessageDefaultResponseDataInnerAttributesContact) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetListMessageDefaultResponseDataInnerAttributesContact) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetListMessageDefaultResponseDataInnerAttributesContact) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *GetListMessageDefaultResponseDataInnerAttributesContact) SetId(v int64) {
	o.Id = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *GetListMessageDefaultResponseDataInnerAttributesContact) GetRole() string {
	if o == nil || IsNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetListMessageDefaultResponseDataInnerAttributesContact) GetRoleOk() (*string, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *GetListMessageDefaultResponseDataInnerAttributesContact) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *GetListMessageDefaultResponseDataInnerAttributesContact) SetRole(v string) {
	o.Role = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *GetListMessageDefaultResponseDataInnerAttributesContact) GetAttributes() GetListMessageDefaultResponseDataInnerAttributesContactAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret GetListMessageDefaultResponseDataInnerAttributesContactAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetListMessageDefaultResponseDataInnerAttributesContact) GetAttributesOk() (*GetListMessageDefaultResponseDataInnerAttributesContactAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *GetListMessageDefaultResponseDataInnerAttributesContact) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given GetListMessageDefaultResponseDataInnerAttributesContactAttributes and assigns it to the Attributes field.
func (o *GetListMessageDefaultResponseDataInnerAttributesContact) SetAttributes(v GetListMessageDefaultResponseDataInnerAttributesContactAttributes) {
	o.Attributes = &v
}

func (o GetListMessageDefaultResponseDataInnerAttributesContact) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetListMessageDefaultResponseDataInnerAttributesContact) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableGetListMessageDefaultResponseDataInnerAttributesContact struct {
	value *GetListMessageDefaultResponseDataInnerAttributesContact
	isSet bool
}

func (v NullableGetListMessageDefaultResponseDataInnerAttributesContact) Get() *GetListMessageDefaultResponseDataInnerAttributesContact {
	return v.value
}

func (v *NullableGetListMessageDefaultResponseDataInnerAttributesContact) Set(val *GetListMessageDefaultResponseDataInnerAttributesContact) {
	v.value = val
	v.isSet = true
}

func (v NullableGetListMessageDefaultResponseDataInnerAttributesContact) IsSet() bool {
	return v.isSet
}

func (v *NullableGetListMessageDefaultResponseDataInnerAttributesContact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetListMessageDefaultResponseDataInnerAttributesContact(val *GetListMessageDefaultResponseDataInnerAttributesContact) *NullableGetListMessageDefaultResponseDataInnerAttributesContact {
	return &NullableGetListMessageDefaultResponseDataInnerAttributesContact{value: val, isSet: true}
}

func (v NullableGetListMessageDefaultResponseDataInnerAttributesContact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetListMessageDefaultResponseDataInnerAttributesContact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


