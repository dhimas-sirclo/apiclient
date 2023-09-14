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

// checks if the WebhookPriceChangeNotification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookPriceChangeNotification{}

// WebhookPriceChangeNotification Price Changes is notification that show information new price value. This notification will be trigger when : - Hit Open API Update Price - Edit Price At Seller Dashboard - Hit Open API Edit Product Endpoint by change price into new value( 
type WebhookPriceChangeNotification struct {
	// Fulfillment service unique identifier
	FsId *int64 `json:"fs_id,omitempty"`
	// Products changes data
	ProductChangesData []WebhookPriceChangeNotificationProductChangesDataInner `json:"product_changes_data,omitempty"`
}

// NewWebhookPriceChangeNotification instantiates a new WebhookPriceChangeNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookPriceChangeNotification() *WebhookPriceChangeNotification {
	this := WebhookPriceChangeNotification{}
	return &this
}

// NewWebhookPriceChangeNotificationWithDefaults instantiates a new WebhookPriceChangeNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookPriceChangeNotificationWithDefaults() *WebhookPriceChangeNotification {
	this := WebhookPriceChangeNotification{}
	return &this
}

// GetFsId returns the FsId field value if set, zero value otherwise.
func (o *WebhookPriceChangeNotification) GetFsId() int64 {
	if o == nil || IsNil(o.FsId) {
		var ret int64
		return ret
	}
	return *o.FsId
}

// GetFsIdOk returns a tuple with the FsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookPriceChangeNotification) GetFsIdOk() (*int64, bool) {
	if o == nil || IsNil(o.FsId) {
		return nil, false
	}
	return o.FsId, true
}

// HasFsId returns a boolean if a field has been set.
func (o *WebhookPriceChangeNotification) HasFsId() bool {
	if o != nil && !IsNil(o.FsId) {
		return true
	}

	return false
}

// SetFsId gets a reference to the given int64 and assigns it to the FsId field.
func (o *WebhookPriceChangeNotification) SetFsId(v int64) {
	o.FsId = &v
}

// GetProductChangesData returns the ProductChangesData field value if set, zero value otherwise.
func (o *WebhookPriceChangeNotification) GetProductChangesData() []WebhookPriceChangeNotificationProductChangesDataInner {
	if o == nil || IsNil(o.ProductChangesData) {
		var ret []WebhookPriceChangeNotificationProductChangesDataInner
		return ret
	}
	return o.ProductChangesData
}

// GetProductChangesDataOk returns a tuple with the ProductChangesData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookPriceChangeNotification) GetProductChangesDataOk() ([]WebhookPriceChangeNotificationProductChangesDataInner, bool) {
	if o == nil || IsNil(o.ProductChangesData) {
		return nil, false
	}
	return o.ProductChangesData, true
}

// HasProductChangesData returns a boolean if a field has been set.
func (o *WebhookPriceChangeNotification) HasProductChangesData() bool {
	if o != nil && !IsNil(o.ProductChangesData) {
		return true
	}

	return false
}

// SetProductChangesData gets a reference to the given []WebhookPriceChangeNotificationProductChangesDataInner and assigns it to the ProductChangesData field.
func (o *WebhookPriceChangeNotification) SetProductChangesData(v []WebhookPriceChangeNotificationProductChangesDataInner) {
	o.ProductChangesData = v
}

func (o WebhookPriceChangeNotification) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookPriceChangeNotification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FsId) {
		toSerialize["fs_id"] = o.FsId
	}
	if !IsNil(o.ProductChangesData) {
		toSerialize["product_changes_data"] = o.ProductChangesData
	}
	return toSerialize, nil
}

type NullableWebhookPriceChangeNotification struct {
	value *WebhookPriceChangeNotification
	isSet bool
}

func (v NullableWebhookPriceChangeNotification) Get() *WebhookPriceChangeNotification {
	return v.value
}

func (v *NullableWebhookPriceChangeNotification) Set(val *WebhookPriceChangeNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookPriceChangeNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookPriceChangeNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookPriceChangeNotification(val *WebhookPriceChangeNotification) *NullableWebhookPriceChangeNotification {
	return &NullableWebhookPriceChangeNotification{value: val, isSet: true}
}

func (v NullableWebhookPriceChangeNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookPriceChangeNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


