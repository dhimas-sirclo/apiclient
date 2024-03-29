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

// checks if the WebhookOrderNotificationCustomer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookOrderNotificationCustomer{}

// WebhookOrderNotificationCustomer Customer data
type WebhookOrderNotificationCustomer struct {
	// Customer ID
	Id *int64 `json:"id,omitempty"`
	// Customer name
	Name *string `json:"Name,omitempty"`
	// Customer phone (Masked)
	Phone *string `json:"phone,omitempty"`
	Email *string `json:"email,omitempty"`
}

// NewWebhookOrderNotificationCustomer instantiates a new WebhookOrderNotificationCustomer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookOrderNotificationCustomer() *WebhookOrderNotificationCustomer {
	this := WebhookOrderNotificationCustomer{}
	return &this
}

// NewWebhookOrderNotificationCustomerWithDefaults instantiates a new WebhookOrderNotificationCustomer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookOrderNotificationCustomerWithDefaults() *WebhookOrderNotificationCustomer {
	this := WebhookOrderNotificationCustomer{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WebhookOrderNotificationCustomer) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookOrderNotificationCustomer) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WebhookOrderNotificationCustomer) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *WebhookOrderNotificationCustomer) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WebhookOrderNotificationCustomer) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookOrderNotificationCustomer) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WebhookOrderNotificationCustomer) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WebhookOrderNotificationCustomer) SetName(v string) {
	o.Name = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *WebhookOrderNotificationCustomer) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookOrderNotificationCustomer) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *WebhookOrderNotificationCustomer) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *WebhookOrderNotificationCustomer) SetPhone(v string) {
	o.Phone = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *WebhookOrderNotificationCustomer) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookOrderNotificationCustomer) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *WebhookOrderNotificationCustomer) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *WebhookOrderNotificationCustomer) SetEmail(v string) {
	o.Email = &v
}

func (o WebhookOrderNotificationCustomer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookOrderNotificationCustomer) ToMap() (map[string]interface{}, error) {
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

type NullableWebhookOrderNotificationCustomer struct {
	value *WebhookOrderNotificationCustomer
	isSet bool
}

func (v NullableWebhookOrderNotificationCustomer) Get() *WebhookOrderNotificationCustomer {
	return v.value
}

func (v *NullableWebhookOrderNotificationCustomer) Set(val *WebhookOrderNotificationCustomer) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookOrderNotificationCustomer) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookOrderNotificationCustomer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookOrderNotificationCustomer(val *WebhookOrderNotificationCustomer) *NullableWebhookOrderNotificationCustomer {
	return &NullableWebhookOrderNotificationCustomer{value: val, isSet: true}
}

func (v NullableWebhookOrderNotificationCustomer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookOrderNotificationCustomer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


