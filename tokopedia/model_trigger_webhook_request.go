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

// checks if the TriggerWebhookRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TriggerWebhookRequest{}

// TriggerWebhookRequest struct for TriggerWebhookRequest
type TriggerWebhookRequest struct {
	OrderId int64 `json:"order_id"`
	Type string `json:"type"`
	Url string `json:"url"`
	IsEncrypted bool `json:"is_encrypted"`
}

// NewTriggerWebhookRequest instantiates a new TriggerWebhookRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTriggerWebhookRequest(orderId int64, type_ string, url string, isEncrypted bool) *TriggerWebhookRequest {
	this := TriggerWebhookRequest{}
	this.OrderId = orderId
	this.Type = type_
	this.Url = url
	this.IsEncrypted = isEncrypted
	return &this
}

// NewTriggerWebhookRequestWithDefaults instantiates a new TriggerWebhookRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTriggerWebhookRequestWithDefaults() *TriggerWebhookRequest {
	this := TriggerWebhookRequest{}
	return &this
}

// GetOrderId returns the OrderId field value
func (o *TriggerWebhookRequest) GetOrderId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value
// and a boolean to check if the value has been set.
func (o *TriggerWebhookRequest) GetOrderIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderId, true
}

// SetOrderId sets field value
func (o *TriggerWebhookRequest) SetOrderId(v int64) {
	o.OrderId = v
}

// GetType returns the Type field value
func (o *TriggerWebhookRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TriggerWebhookRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TriggerWebhookRequest) SetType(v string) {
	o.Type = v
}

// GetUrl returns the Url field value
func (o *TriggerWebhookRequest) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *TriggerWebhookRequest) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *TriggerWebhookRequest) SetUrl(v string) {
	o.Url = v
}

// GetIsEncrypted returns the IsEncrypted field value
func (o *TriggerWebhookRequest) GetIsEncrypted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsEncrypted
}

// GetIsEncryptedOk returns a tuple with the IsEncrypted field value
// and a boolean to check if the value has been set.
func (o *TriggerWebhookRequest) GetIsEncryptedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsEncrypted, true
}

// SetIsEncrypted sets field value
func (o *TriggerWebhookRequest) SetIsEncrypted(v bool) {
	o.IsEncrypted = v
}

func (o TriggerWebhookRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TriggerWebhookRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["order_id"] = o.OrderId
	toSerialize["type"] = o.Type
	toSerialize["url"] = o.Url
	toSerialize["is_encrypted"] = o.IsEncrypted
	return toSerialize, nil
}

type NullableTriggerWebhookRequest struct {
	value *TriggerWebhookRequest
	isSet bool
}

func (v NullableTriggerWebhookRequest) Get() *TriggerWebhookRequest {
	return v.value
}

func (v *NullableTriggerWebhookRequest) Set(val *TriggerWebhookRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTriggerWebhookRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTriggerWebhookRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTriggerWebhookRequest(val *TriggerWebhookRequest) *NullableTriggerWebhookRequest {
	return &NullableTriggerWebhookRequest{value: val, isSet: true}
}

func (v NullableTriggerWebhookRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTriggerWebhookRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

