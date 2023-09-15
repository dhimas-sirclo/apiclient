/*
SIRCLO Open API

SIRCLO Open API

API version: 1.0.0
Contact: dev@sirclo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sirclo

import (
	"encoding/json"
)

// checks if the VariantInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VariantInput{}

// VariantInput struct for VariantInput
type VariantInput struct {
	VariantId string `json:"variant_id"`
	VariantSku string `json:"variant_sku"`
	VariantName string `json:"variant_name"`
	Status string `json:"status"`
	Volume VolumeInput `json:"volume"`
	Weight float64 `json:"weight"`
	Attributes []VariantAttributeInput `json:"attributes"`
	VariantUrl string `json:"variant_url"`
	CurrencyCode string `json:"currency_code"`
	Price float64 `json:"price"`
	ImgUrls []string `json:"img_urls"`
}

// NewVariantInput instantiates a new VariantInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVariantInput(variantId string, variantSku string, variantName string, status string, volume VolumeInput, weight float64, attributes []VariantAttributeInput, variantUrl string, currencyCode string, price float64, imgUrls []string) *VariantInput {
	this := VariantInput{}
	this.VariantId = variantId
	this.VariantSku = variantSku
	this.VariantName = variantName
	this.Status = status
	this.Volume = volume
	this.Weight = weight
	this.Attributes = attributes
	this.VariantUrl = variantUrl
	this.CurrencyCode = currencyCode
	this.Price = price
	this.ImgUrls = imgUrls
	return &this
}

// NewVariantInputWithDefaults instantiates a new VariantInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVariantInputWithDefaults() *VariantInput {
	this := VariantInput{}
	return &this
}

// GetVariantId returns the VariantId field value
func (o *VariantInput) GetVariantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VariantId
}

// GetVariantIdOk returns a tuple with the VariantId field value
// and a boolean to check if the value has been set.
func (o *VariantInput) GetVariantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VariantId, true
}

// SetVariantId sets field value
func (o *VariantInput) SetVariantId(v string) {
	o.VariantId = v
}

// GetVariantSku returns the VariantSku field value
func (o *VariantInput) GetVariantSku() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VariantSku
}

// GetVariantSkuOk returns a tuple with the VariantSku field value
// and a boolean to check if the value has been set.
func (o *VariantInput) GetVariantSkuOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VariantSku, true
}

// SetVariantSku sets field value
func (o *VariantInput) SetVariantSku(v string) {
	o.VariantSku = v
}

// GetVariantName returns the VariantName field value
func (o *VariantInput) GetVariantName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VariantName
}

// GetVariantNameOk returns a tuple with the VariantName field value
// and a boolean to check if the value has been set.
func (o *VariantInput) GetVariantNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VariantName, true
}

// SetVariantName sets field value
func (o *VariantInput) SetVariantName(v string) {
	o.VariantName = v
}

// GetStatus returns the Status field value
func (o *VariantInput) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *VariantInput) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *VariantInput) SetStatus(v string) {
	o.Status = v
}

// GetVolume returns the Volume field value
func (o *VariantInput) GetVolume() VolumeInput {
	if o == nil {
		var ret VolumeInput
		return ret
	}

	return o.Volume
}

// GetVolumeOk returns a tuple with the Volume field value
// and a boolean to check if the value has been set.
func (o *VariantInput) GetVolumeOk() (*VolumeInput, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Volume, true
}

// SetVolume sets field value
func (o *VariantInput) SetVolume(v VolumeInput) {
	o.Volume = v
}

// GetWeight returns the Weight field value
func (o *VariantInput) GetWeight() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Weight
}

// GetWeightOk returns a tuple with the Weight field value
// and a boolean to check if the value has been set.
func (o *VariantInput) GetWeightOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Weight, true
}

// SetWeight sets field value
func (o *VariantInput) SetWeight(v float64) {
	o.Weight = v
}

// GetAttributes returns the Attributes field value
func (o *VariantInput) GetAttributes() []VariantAttributeInput {
	if o == nil {
		var ret []VariantAttributeInput
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *VariantInput) GetAttributesOk() ([]VariantAttributeInput, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attributes, true
}

// SetAttributes sets field value
func (o *VariantInput) SetAttributes(v []VariantAttributeInput) {
	o.Attributes = v
}

// GetVariantUrl returns the VariantUrl field value
func (o *VariantInput) GetVariantUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VariantUrl
}

// GetVariantUrlOk returns a tuple with the VariantUrl field value
// and a boolean to check if the value has been set.
func (o *VariantInput) GetVariantUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VariantUrl, true
}

// SetVariantUrl sets field value
func (o *VariantInput) SetVariantUrl(v string) {
	o.VariantUrl = v
}

// GetCurrencyCode returns the CurrencyCode field value
func (o *VariantInput) GetCurrencyCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value
// and a boolean to check if the value has been set.
func (o *VariantInput) GetCurrencyCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrencyCode, true
}

// SetCurrencyCode sets field value
func (o *VariantInput) SetCurrencyCode(v string) {
	o.CurrencyCode = v
}

// GetPrice returns the Price field value
func (o *VariantInput) GetPrice() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *VariantInput) GetPriceOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Price, true
}

// SetPrice sets field value
func (o *VariantInput) SetPrice(v float64) {
	o.Price = v
}

// GetImgUrls returns the ImgUrls field value
func (o *VariantInput) GetImgUrls() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ImgUrls
}

// GetImgUrlsOk returns a tuple with the ImgUrls field value
// and a boolean to check if the value has been set.
func (o *VariantInput) GetImgUrlsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImgUrls, true
}

// SetImgUrls sets field value
func (o *VariantInput) SetImgUrls(v []string) {
	o.ImgUrls = v
}

func (o VariantInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VariantInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["variant_id"] = o.VariantId
	toSerialize["variant_sku"] = o.VariantSku
	toSerialize["variant_name"] = o.VariantName
	toSerialize["status"] = o.Status
	toSerialize["volume"] = o.Volume
	toSerialize["weight"] = o.Weight
	toSerialize["attributes"] = o.Attributes
	toSerialize["variant_url"] = o.VariantUrl
	toSerialize["currency_code"] = o.CurrencyCode
	toSerialize["price"] = o.Price
	toSerialize["img_urls"] = o.ImgUrls
	return toSerialize, nil
}

type NullableVariantInput struct {
	value *VariantInput
	isSet bool
}

func (v NullableVariantInput) Get() *VariantInput {
	return v.value
}

func (v *NullableVariantInput) Set(val *VariantInput) {
	v.value = val
	v.isSet = true
}

func (v NullableVariantInput) IsSet() bool {
	return v.isSet
}

func (v *NullableVariantInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVariantInput(val *VariantInput) *NullableVariantInput {
	return &NullableVariantInput{value: val, isSet: true}
}

func (v NullableVariantInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVariantInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


