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

// checks if the WebhookOrderNotificationBundleDetailNonBundleInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookOrderNotificationBundleDetailNonBundleInner{}

// WebhookOrderNotificationBundleDetailNonBundleInner struct for WebhookOrderNotificationBundleDetailNonBundleInner
type WebhookOrderNotificationBundleDetailNonBundleInner struct {
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"Name,omitempty"`
	Quantity *int64 `json:"quantity,omitempty"`
	Notes *string `json:"notes,omitempty"`
	Weight *float64 `json:"weight,omitempty"`
	TotalWeight *float64 `json:"total_weight,omitempty"`
	Price *int64 `json:"price,omitempty"`
	TotalPrice *int64 `json:"total_price,omitempty"`
	Currency *string `json:"currency,omitempty"`
	Sku *string `json:"sku,omitempty"`
}

// NewWebhookOrderNotificationBundleDetailNonBundleInner instantiates a new WebhookOrderNotificationBundleDetailNonBundleInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookOrderNotificationBundleDetailNonBundleInner() *WebhookOrderNotificationBundleDetailNonBundleInner {
	this := WebhookOrderNotificationBundleDetailNonBundleInner{}
	return &this
}

// NewWebhookOrderNotificationBundleDetailNonBundleInnerWithDefaults instantiates a new WebhookOrderNotificationBundleDetailNonBundleInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookOrderNotificationBundleDetailNonBundleInnerWithDefaults() *WebhookOrderNotificationBundleDetailNonBundleInner {
	this := WebhookOrderNotificationBundleDetailNonBundleInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WebhookOrderNotificationBundleDetailNonBundleInner) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookOrderNotificationBundleDetailNonBundleInner) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WebhookOrderNotificationBundleDetailNonBundleInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *WebhookOrderNotificationBundleDetailNonBundleInner) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WebhookOrderNotificationBundleDetailNonBundleInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookOrderNotificationBundleDetailNonBundleInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WebhookOrderNotificationBundleDetailNonBundleInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WebhookOrderNotificationBundleDetailNonBundleInner) SetName(v string) {
	o.Name = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *WebhookOrderNotificationBundleDetailNonBundleInner) GetQuantity() int64 {
	if o == nil || IsNil(o.Quantity) {
		var ret int64
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookOrderNotificationBundleDetailNonBundleInner) GetQuantityOk() (*int64, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *WebhookOrderNotificationBundleDetailNonBundleInner) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int64 and assigns it to the Quantity field.
func (o *WebhookOrderNotificationBundleDetailNonBundleInner) SetQuantity(v int64) {
	o.Quantity = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *WebhookOrderNotificationBundleDetailNonBundleInner) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookOrderNotificationBundleDetailNonBundleInner) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *WebhookOrderNotificationBundleDetailNonBundleInner) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *WebhookOrderNotificationBundleDetailNonBundleInner) SetNotes(v string) {
	o.Notes = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *WebhookOrderNotificationBundleDetailNonBundleInner) GetWeight() float64 {
	if o == nil || IsNil(o.Weight) {
		var ret float64
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookOrderNotificationBundleDetailNonBundleInner) GetWeightOk() (*float64, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *WebhookOrderNotificationBundleDetailNonBundleInner) HasWeight() bool {
	if o != nil && !IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given float64 and assigns it to the Weight field.
func (o *WebhookOrderNotificationBundleDetailNonBundleInner) SetWeight(v float64) {
	o.Weight = &v
}

// GetTotalWeight returns the TotalWeight field value if set, zero value otherwise.
func (o *WebhookOrderNotificationBundleDetailNonBundleInner) GetTotalWeight() float64 {
	if o == nil || IsNil(o.TotalWeight) {
		var ret float64
		return ret
	}
	return *o.TotalWeight
}

// GetTotalWeightOk returns a tuple with the TotalWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookOrderNotificationBundleDetailNonBundleInner) GetTotalWeightOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalWeight) {
		return nil, false
	}
	return o.TotalWeight, true
}

// HasTotalWeight returns a boolean if a field has been set.
func (o *WebhookOrderNotificationBundleDetailNonBundleInner) HasTotalWeight() bool {
	if o != nil && !IsNil(o.TotalWeight) {
		return true
	}

	return false
}

// SetTotalWeight gets a reference to the given float64 and assigns it to the TotalWeight field.
func (o *WebhookOrderNotificationBundleDetailNonBundleInner) SetTotalWeight(v float64) {
	o.TotalWeight = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *WebhookOrderNotificationBundleDetailNonBundleInner) GetPrice() int64 {
	if o == nil || IsNil(o.Price) {
		var ret int64
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookOrderNotificationBundleDetailNonBundleInner) GetPriceOk() (*int64, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *WebhookOrderNotificationBundleDetailNonBundleInner) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given int64 and assigns it to the Price field.
func (o *WebhookOrderNotificationBundleDetailNonBundleInner) SetPrice(v int64) {
	o.Price = &v
}

// GetTotalPrice returns the TotalPrice field value if set, zero value otherwise.
func (o *WebhookOrderNotificationBundleDetailNonBundleInner) GetTotalPrice() int64 {
	if o == nil || IsNil(o.TotalPrice) {
		var ret int64
		return ret
	}
	return *o.TotalPrice
}

// GetTotalPriceOk returns a tuple with the TotalPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookOrderNotificationBundleDetailNonBundleInner) GetTotalPriceOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalPrice) {
		return nil, false
	}
	return o.TotalPrice, true
}

// HasTotalPrice returns a boolean if a field has been set.
func (o *WebhookOrderNotificationBundleDetailNonBundleInner) HasTotalPrice() bool {
	if o != nil && !IsNil(o.TotalPrice) {
		return true
	}

	return false
}

// SetTotalPrice gets a reference to the given int64 and assigns it to the TotalPrice field.
func (o *WebhookOrderNotificationBundleDetailNonBundleInner) SetTotalPrice(v int64) {
	o.TotalPrice = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *WebhookOrderNotificationBundleDetailNonBundleInner) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookOrderNotificationBundleDetailNonBundleInner) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *WebhookOrderNotificationBundleDetailNonBundleInner) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *WebhookOrderNotificationBundleDetailNonBundleInner) SetCurrency(v string) {
	o.Currency = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *WebhookOrderNotificationBundleDetailNonBundleInner) GetSku() string {
	if o == nil || IsNil(o.Sku) {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookOrderNotificationBundleDetailNonBundleInner) GetSkuOk() (*string, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *WebhookOrderNotificationBundleDetailNonBundleInner) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *WebhookOrderNotificationBundleDetailNonBundleInner) SetSku(v string) {
	o.Sku = &v
}

func (o WebhookOrderNotificationBundleDetailNonBundleInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookOrderNotificationBundleDetailNonBundleInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if !IsNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	if !IsNil(o.TotalWeight) {
		toSerialize["total_weight"] = o.TotalWeight
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.TotalPrice) {
		toSerialize["total_price"] = o.TotalPrice
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	return toSerialize, nil
}

type NullableWebhookOrderNotificationBundleDetailNonBundleInner struct {
	value *WebhookOrderNotificationBundleDetailNonBundleInner
	isSet bool
}

func (v NullableWebhookOrderNotificationBundleDetailNonBundleInner) Get() *WebhookOrderNotificationBundleDetailNonBundleInner {
	return v.value
}

func (v *NullableWebhookOrderNotificationBundleDetailNonBundleInner) Set(val *WebhookOrderNotificationBundleDetailNonBundleInner) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookOrderNotificationBundleDetailNonBundleInner) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookOrderNotificationBundleDetailNonBundleInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookOrderNotificationBundleDetailNonBundleInner(val *WebhookOrderNotificationBundleDetailNonBundleInner) *NullableWebhookOrderNotificationBundleDetailNonBundleInner {
	return &NullableWebhookOrderNotificationBundleDetailNonBundleInner{value: val, isSet: true}
}

func (v NullableWebhookOrderNotificationBundleDetailNonBundleInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookOrderNotificationBundleDetailNonBundleInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

