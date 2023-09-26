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

// checks if the GetOrderWebhook200ResponseDataOneOf1Customer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrderWebhook200ResponseDataOneOf1Customer{}

// GetOrderWebhook200ResponseDataOneOf1Customer struct for GetOrderWebhook200ResponseDataOneOf1Customer
type GetOrderWebhook200ResponseDataOneOf1Customer struct {
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"Name,omitempty"`
	Phone *string `json:"phone,omitempty"`
	Email *string `json:"email,omitempty"`
}

// NewGetOrderWebhook200ResponseDataOneOf1Customer instantiates a new GetOrderWebhook200ResponseDataOneOf1Customer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrderWebhook200ResponseDataOneOf1Customer() *GetOrderWebhook200ResponseDataOneOf1Customer {
	this := GetOrderWebhook200ResponseDataOneOf1Customer{}
	return &this
}

// NewGetOrderWebhook200ResponseDataOneOf1CustomerWithDefaults instantiates a new GetOrderWebhook200ResponseDataOneOf1Customer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrderWebhook200ResponseDataOneOf1CustomerWithDefaults() *GetOrderWebhook200ResponseDataOneOf1Customer {
	this := GetOrderWebhook200ResponseDataOneOf1Customer{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetOrderWebhook200ResponseDataOneOf1Customer) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderWebhook200ResponseDataOneOf1Customer) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetOrderWebhook200ResponseDataOneOf1Customer) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *GetOrderWebhook200ResponseDataOneOf1Customer) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetOrderWebhook200ResponseDataOneOf1Customer) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderWebhook200ResponseDataOneOf1Customer) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetOrderWebhook200ResponseDataOneOf1Customer) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetOrderWebhook200ResponseDataOneOf1Customer) SetName(v string) {
	o.Name = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *GetOrderWebhook200ResponseDataOneOf1Customer) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderWebhook200ResponseDataOneOf1Customer) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *GetOrderWebhook200ResponseDataOneOf1Customer) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *GetOrderWebhook200ResponseDataOneOf1Customer) SetPhone(v string) {
	o.Phone = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *GetOrderWebhook200ResponseDataOneOf1Customer) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderWebhook200ResponseDataOneOf1Customer) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *GetOrderWebhook200ResponseDataOneOf1Customer) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *GetOrderWebhook200ResponseDataOneOf1Customer) SetEmail(v string) {
	o.Email = &v
}

func (o GetOrderWebhook200ResponseDataOneOf1Customer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrderWebhook200ResponseDataOneOf1Customer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	return toSerialize, nil
}

type NullableGetOrderWebhook200ResponseDataOneOf1Customer struct {
	value *GetOrderWebhook200ResponseDataOneOf1Customer
	isSet bool
}

func (v NullableGetOrderWebhook200ResponseDataOneOf1Customer) Get() *GetOrderWebhook200ResponseDataOneOf1Customer {
	return v.value
}

func (v *NullableGetOrderWebhook200ResponseDataOneOf1Customer) Set(val *GetOrderWebhook200ResponseDataOneOf1Customer) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrderWebhook200ResponseDataOneOf1Customer) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrderWebhook200ResponseDataOneOf1Customer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrderWebhook200ResponseDataOneOf1Customer(val *GetOrderWebhook200ResponseDataOneOf1Customer) *NullableGetOrderWebhook200ResponseDataOneOf1Customer {
	return &NullableGetOrderWebhook200ResponseDataOneOf1Customer{value: val, isSet: true}
}

func (v NullableGetOrderWebhook200ResponseDataOneOf1Customer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrderWebhook200ResponseDataOneOf1Customer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


