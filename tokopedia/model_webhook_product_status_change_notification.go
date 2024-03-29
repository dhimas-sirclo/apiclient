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

// checks if the WebhookProductStatusChangeNotification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookProductStatusChangeNotification{}

// WebhookProductStatusChangeNotification Status Changes is notification that show information new product status value. This notification will be trigger when : - Hit Open API Update Stock - Hit Open API Set Active Product - Hit Open API Set Inactive Product - Hit Open API Set Delete Product - Edit Status At Seller Dashboard - Hit Open API Edit Product Endpoint by change status into new value 
type WebhookProductStatusChangeNotification struct {
	// Fulfillment service unique identifier
	FsId *int64 `json:"fs_id,omitempty"`
	// Products changes data
	ProductChangesData []WebhookProductStatusChangeNotificationProductChangesDataInner `json:"product_changes_data,omitempty"`
}

// NewWebhookProductStatusChangeNotification instantiates a new WebhookProductStatusChangeNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookProductStatusChangeNotification() *WebhookProductStatusChangeNotification {
	this := WebhookProductStatusChangeNotification{}
	return &this
}

// NewWebhookProductStatusChangeNotificationWithDefaults instantiates a new WebhookProductStatusChangeNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookProductStatusChangeNotificationWithDefaults() *WebhookProductStatusChangeNotification {
	this := WebhookProductStatusChangeNotification{}
	return &this
}

// GetFsId returns the FsId field value if set, zero value otherwise.
func (o *WebhookProductStatusChangeNotification) GetFsId() int64 {
	if o == nil || IsNil(o.FsId) {
		var ret int64
		return ret
	}
	return *o.FsId
}

// GetFsIdOk returns a tuple with the FsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookProductStatusChangeNotification) GetFsIdOk() (*int64, bool) {
	if o == nil || IsNil(o.FsId) {
		return nil, false
	}
	return o.FsId, true
}

// HasFsId returns a boolean if a field has been set.
func (o *WebhookProductStatusChangeNotification) HasFsId() bool {
	if o != nil && !IsNil(o.FsId) {
		return true
	}

	return false
}

// SetFsId gets a reference to the given int64 and assigns it to the FsId field.
func (o *WebhookProductStatusChangeNotification) SetFsId(v int64) {
	o.FsId = &v
}

// GetProductChangesData returns the ProductChangesData field value if set, zero value otherwise.
func (o *WebhookProductStatusChangeNotification) GetProductChangesData() []WebhookProductStatusChangeNotificationProductChangesDataInner {
	if o == nil || IsNil(o.ProductChangesData) {
		var ret []WebhookProductStatusChangeNotificationProductChangesDataInner
		return ret
	}
	return o.ProductChangesData
}

// GetProductChangesDataOk returns a tuple with the ProductChangesData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookProductStatusChangeNotification) GetProductChangesDataOk() ([]WebhookProductStatusChangeNotificationProductChangesDataInner, bool) {
	if o == nil || IsNil(o.ProductChangesData) {
		return nil, false
	}
	return o.ProductChangesData, true
}

// HasProductChangesData returns a boolean if a field has been set.
func (o *WebhookProductStatusChangeNotification) HasProductChangesData() bool {
	if o != nil && !IsNil(o.ProductChangesData) {
		return true
	}

	return false
}

// SetProductChangesData gets a reference to the given []WebhookProductStatusChangeNotificationProductChangesDataInner and assigns it to the ProductChangesData field.
func (o *WebhookProductStatusChangeNotification) SetProductChangesData(v []WebhookProductStatusChangeNotificationProductChangesDataInner) {
	o.ProductChangesData = v
}

func (o WebhookProductStatusChangeNotification) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookProductStatusChangeNotification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FsId) {
		toSerialize["fs_id"] = o.FsId
	}
	if !IsNil(o.ProductChangesData) {
		toSerialize["product_changes_data"] = o.ProductChangesData
	}
	return toSerialize, nil
}

type NullableWebhookProductStatusChangeNotification struct {
	value *WebhookProductStatusChangeNotification
	isSet bool
}

func (v NullableWebhookProductStatusChangeNotification) Get() *WebhookProductStatusChangeNotification {
	return v.value
}

func (v *NullableWebhookProductStatusChangeNotification) Set(val *WebhookProductStatusChangeNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookProductStatusChangeNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookProductStatusChangeNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookProductStatusChangeNotification(val *WebhookProductStatusChangeNotification) *NullableWebhookProductStatusChangeNotification {
	return &NullableWebhookProductStatusChangeNotification{value: val, isSet: true}
}

func (v NullableWebhookProductStatusChangeNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookProductStatusChangeNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


