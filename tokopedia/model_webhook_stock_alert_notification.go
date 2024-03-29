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

// checks if the WebhookStockAlertNotification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookStockAlertNotification{}

// WebhookStockAlertNotification Stock Alert is notification that show information alert stock into new value. This notification will be trigger when : - Edit Stock Alert At Seller Dashboard 
type WebhookStockAlertNotification struct {
	// Fulfillment service unique identifier
	FsId *int64 `json:"fs_id,omitempty"`
	// Products changes data
	ProductChangesData []WebhookStockAlertNotificationProductChangesDataInner `json:"product_changes_data,omitempty"`
}

// NewWebhookStockAlertNotification instantiates a new WebhookStockAlertNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookStockAlertNotification() *WebhookStockAlertNotification {
	this := WebhookStockAlertNotification{}
	return &this
}

// NewWebhookStockAlertNotificationWithDefaults instantiates a new WebhookStockAlertNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookStockAlertNotificationWithDefaults() *WebhookStockAlertNotification {
	this := WebhookStockAlertNotification{}
	return &this
}

// GetFsId returns the FsId field value if set, zero value otherwise.
func (o *WebhookStockAlertNotification) GetFsId() int64 {
	if o == nil || IsNil(o.FsId) {
		var ret int64
		return ret
	}
	return *o.FsId
}

// GetFsIdOk returns a tuple with the FsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookStockAlertNotification) GetFsIdOk() (*int64, bool) {
	if o == nil || IsNil(o.FsId) {
		return nil, false
	}
	return o.FsId, true
}

// HasFsId returns a boolean if a field has been set.
func (o *WebhookStockAlertNotification) HasFsId() bool {
	if o != nil && !IsNil(o.FsId) {
		return true
	}

	return false
}

// SetFsId gets a reference to the given int64 and assigns it to the FsId field.
func (o *WebhookStockAlertNotification) SetFsId(v int64) {
	o.FsId = &v
}

// GetProductChangesData returns the ProductChangesData field value if set, zero value otherwise.
func (o *WebhookStockAlertNotification) GetProductChangesData() []WebhookStockAlertNotificationProductChangesDataInner {
	if o == nil || IsNil(o.ProductChangesData) {
		var ret []WebhookStockAlertNotificationProductChangesDataInner
		return ret
	}
	return o.ProductChangesData
}

// GetProductChangesDataOk returns a tuple with the ProductChangesData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookStockAlertNotification) GetProductChangesDataOk() ([]WebhookStockAlertNotificationProductChangesDataInner, bool) {
	if o == nil || IsNil(o.ProductChangesData) {
		return nil, false
	}
	return o.ProductChangesData, true
}

// HasProductChangesData returns a boolean if a field has been set.
func (o *WebhookStockAlertNotification) HasProductChangesData() bool {
	if o != nil && !IsNil(o.ProductChangesData) {
		return true
	}

	return false
}

// SetProductChangesData gets a reference to the given []WebhookStockAlertNotificationProductChangesDataInner and assigns it to the ProductChangesData field.
func (o *WebhookStockAlertNotification) SetProductChangesData(v []WebhookStockAlertNotificationProductChangesDataInner) {
	o.ProductChangesData = v
}

func (o WebhookStockAlertNotification) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookStockAlertNotification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FsId) {
		toSerialize["fs_id"] = o.FsId
	}
	if !IsNil(o.ProductChangesData) {
		toSerialize["product_changes_data"] = o.ProductChangesData
	}
	return toSerialize, nil
}

type NullableWebhookStockAlertNotification struct {
	value *WebhookStockAlertNotification
	isSet bool
}

func (v NullableWebhookStockAlertNotification) Get() *WebhookStockAlertNotification {
	return v.value
}

func (v *NullableWebhookStockAlertNotification) Set(val *WebhookStockAlertNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookStockAlertNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookStockAlertNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookStockAlertNotification(val *WebhookStockAlertNotification) *NullableWebhookStockAlertNotification {
	return &NullableWebhookStockAlertNotification{value: val, isSet: true}
}

func (v NullableWebhookStockAlertNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookStockAlertNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


