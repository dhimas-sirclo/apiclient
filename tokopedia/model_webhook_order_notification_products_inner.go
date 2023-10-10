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

// checks if the WebhookOrderNotificationProductsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookOrderNotificationProductsInner{}

// WebhookOrderNotificationProductsInner struct for WebhookOrderNotificationProductsInner
type WebhookOrderNotificationProductsInner struct {
	// Product ID
	Id *int64 `json:"id,omitempty"`
	// Product name
	Name *string `json:"Name,omitempty"`
	// Product quantity
	Quantity *int64 `json:"quantity,omitempty"`
	// Product notes
	Notes *string `json:"notes,omitempty"`
	Weight *float64 `json:"weight,omitempty"`
	TotalWeight *float64 `json:"total_weight,omitempty"`
	// Product price
	Price *float64 `json:"price,omitempty"`
	// Total price
	TotalPrice *float64 `json:"total_price,omitempty"`
	// Currency code
	Currency *string `json:"currency,omitempty"`
	// SKU code
	Sku *string `json:"sku,omitempty"`
	// Flag which determines if the product is eligible for POF
	IsEligiblePof *bool `json:"is_eligible_pof,omitempty"`
}

// NewWebhookOrderNotificationProductsInner instantiates a new WebhookOrderNotificationProductsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookOrderNotificationProductsInner() *WebhookOrderNotificationProductsInner {
	this := WebhookOrderNotificationProductsInner{}
	return &this
}

// NewWebhookOrderNotificationProductsInnerWithDefaults instantiates a new WebhookOrderNotificationProductsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookOrderNotificationProductsInnerWithDefaults() *WebhookOrderNotificationProductsInner {
	this := WebhookOrderNotificationProductsInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WebhookOrderNotificationProductsInner) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookOrderNotificationProductsInner) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WebhookOrderNotificationProductsInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *WebhookOrderNotificationProductsInner) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WebhookOrderNotificationProductsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookOrderNotificationProductsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WebhookOrderNotificationProductsInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WebhookOrderNotificationProductsInner) SetName(v string) {
	o.Name = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *WebhookOrderNotificationProductsInner) GetQuantity() int64 {
	if o == nil || IsNil(o.Quantity) {
		var ret int64
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookOrderNotificationProductsInner) GetQuantityOk() (*int64, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *WebhookOrderNotificationProductsInner) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int64 and assigns it to the Quantity field.
func (o *WebhookOrderNotificationProductsInner) SetQuantity(v int64) {
	o.Quantity = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *WebhookOrderNotificationProductsInner) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookOrderNotificationProductsInner) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *WebhookOrderNotificationProductsInner) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *WebhookOrderNotificationProductsInner) SetNotes(v string) {
	o.Notes = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *WebhookOrderNotificationProductsInner) GetWeight() float64 {
	if o == nil || IsNil(o.Weight) {
		var ret float64
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookOrderNotificationProductsInner) GetWeightOk() (*float64, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *WebhookOrderNotificationProductsInner) HasWeight() bool {
	if o != nil && !IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given float64 and assigns it to the Weight field.
func (o *WebhookOrderNotificationProductsInner) SetWeight(v float64) {
	o.Weight = &v
}

// GetTotalWeight returns the TotalWeight field value if set, zero value otherwise.
func (o *WebhookOrderNotificationProductsInner) GetTotalWeight() float64 {
	if o == nil || IsNil(o.TotalWeight) {
		var ret float64
		return ret
	}
	return *o.TotalWeight
}

// GetTotalWeightOk returns a tuple with the TotalWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookOrderNotificationProductsInner) GetTotalWeightOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalWeight) {
		return nil, false
	}
	return o.TotalWeight, true
}

// HasTotalWeight returns a boolean if a field has been set.
func (o *WebhookOrderNotificationProductsInner) HasTotalWeight() bool {
	if o != nil && !IsNil(o.TotalWeight) {
		return true
	}

	return false
}

// SetTotalWeight gets a reference to the given float64 and assigns it to the TotalWeight field.
func (o *WebhookOrderNotificationProductsInner) SetTotalWeight(v float64) {
	o.TotalWeight = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *WebhookOrderNotificationProductsInner) GetPrice() float64 {
	if o == nil || IsNil(o.Price) {
		var ret float64
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookOrderNotificationProductsInner) GetPriceOk() (*float64, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *WebhookOrderNotificationProductsInner) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float64 and assigns it to the Price field.
func (o *WebhookOrderNotificationProductsInner) SetPrice(v float64) {
	o.Price = &v
}

// GetTotalPrice returns the TotalPrice field value if set, zero value otherwise.
func (o *WebhookOrderNotificationProductsInner) GetTotalPrice() float64 {
	if o == nil || IsNil(o.TotalPrice) {
		var ret float64
		return ret
	}
	return *o.TotalPrice
}

// GetTotalPriceOk returns a tuple with the TotalPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookOrderNotificationProductsInner) GetTotalPriceOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalPrice) {
		return nil, false
	}
	return o.TotalPrice, true
}

// HasTotalPrice returns a boolean if a field has been set.
func (o *WebhookOrderNotificationProductsInner) HasTotalPrice() bool {
	if o != nil && !IsNil(o.TotalPrice) {
		return true
	}

	return false
}

// SetTotalPrice gets a reference to the given float64 and assigns it to the TotalPrice field.
func (o *WebhookOrderNotificationProductsInner) SetTotalPrice(v float64) {
	o.TotalPrice = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *WebhookOrderNotificationProductsInner) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookOrderNotificationProductsInner) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *WebhookOrderNotificationProductsInner) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *WebhookOrderNotificationProductsInner) SetCurrency(v string) {
	o.Currency = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *WebhookOrderNotificationProductsInner) GetSku() string {
	if o == nil || IsNil(o.Sku) {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookOrderNotificationProductsInner) GetSkuOk() (*string, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *WebhookOrderNotificationProductsInner) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *WebhookOrderNotificationProductsInner) SetSku(v string) {
	o.Sku = &v
}

// GetIsEligiblePof returns the IsEligiblePof field value if set, zero value otherwise.
func (o *WebhookOrderNotificationProductsInner) GetIsEligiblePof() bool {
	if o == nil || IsNil(o.IsEligiblePof) {
		var ret bool
		return ret
	}
	return *o.IsEligiblePof
}

// GetIsEligiblePofOk returns a tuple with the IsEligiblePof field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookOrderNotificationProductsInner) GetIsEligiblePofOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEligiblePof) {
		return nil, false
	}
	return o.IsEligiblePof, true
}

// HasIsEligiblePof returns a boolean if a field has been set.
func (o *WebhookOrderNotificationProductsInner) HasIsEligiblePof() bool {
	if o != nil && !IsNil(o.IsEligiblePof) {
		return true
	}

	return false
}

// SetIsEligiblePof gets a reference to the given bool and assigns it to the IsEligiblePof field.
func (o *WebhookOrderNotificationProductsInner) SetIsEligiblePof(v bool) {
	o.IsEligiblePof = &v
}

func (o WebhookOrderNotificationProductsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookOrderNotificationProductsInner) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.IsEligiblePof) {
		toSerialize["is_eligible_pof"] = o.IsEligiblePof
	}
	return toSerialize, nil
}

type NullableWebhookOrderNotificationProductsInner struct {
	value *WebhookOrderNotificationProductsInner
	isSet bool
}

func (v NullableWebhookOrderNotificationProductsInner) Get() *WebhookOrderNotificationProductsInner {
	return v.value
}

func (v *NullableWebhookOrderNotificationProductsInner) Set(val *WebhookOrderNotificationProductsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookOrderNotificationProductsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookOrderNotificationProductsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookOrderNotificationProductsInner(val *WebhookOrderNotificationProductsInner) *NullableWebhookOrderNotificationProductsInner {
	return &NullableWebhookOrderNotificationProductsInner{value: val, isSet: true}
}

func (v NullableWebhookOrderNotificationProductsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookOrderNotificationProductsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


