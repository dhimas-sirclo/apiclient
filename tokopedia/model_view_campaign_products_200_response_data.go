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

// checks if the ViewCampaignProducts200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ViewCampaignProducts200ResponseData{}

// ViewCampaignProducts200ResponseData struct for ViewCampaignProducts200ResponseData
type ViewCampaignProducts200ResponseData struct {
	Products []ViewCampaignProducts200ResponseDataProductsInner `json:"products,omitempty"`
	// Total Product Count
	TotalProduct *int64 `json:"total_product,omitempty"`
}

// NewViewCampaignProducts200ResponseData instantiates a new ViewCampaignProducts200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViewCampaignProducts200ResponseData() *ViewCampaignProducts200ResponseData {
	this := ViewCampaignProducts200ResponseData{}
	return &this
}

// NewViewCampaignProducts200ResponseDataWithDefaults instantiates a new ViewCampaignProducts200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViewCampaignProducts200ResponseDataWithDefaults() *ViewCampaignProducts200ResponseData {
	this := ViewCampaignProducts200ResponseData{}
	return &this
}

// GetProducts returns the Products field value if set, zero value otherwise.
func (o *ViewCampaignProducts200ResponseData) GetProducts() []ViewCampaignProducts200ResponseDataProductsInner {
	if o == nil || IsNil(o.Products) {
		var ret []ViewCampaignProducts200ResponseDataProductsInner
		return ret
	}
	return o.Products
}

// GetProductsOk returns a tuple with the Products field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewCampaignProducts200ResponseData) GetProductsOk() ([]ViewCampaignProducts200ResponseDataProductsInner, bool) {
	if o == nil || IsNil(o.Products) {
		return nil, false
	}
	return o.Products, true
}

// HasProducts returns a boolean if a field has been set.
func (o *ViewCampaignProducts200ResponseData) HasProducts() bool {
	if o != nil && !IsNil(o.Products) {
		return true
	}

	return false
}

// SetProducts gets a reference to the given []ViewCampaignProducts200ResponseDataProductsInner and assigns it to the Products field.
func (o *ViewCampaignProducts200ResponseData) SetProducts(v []ViewCampaignProducts200ResponseDataProductsInner) {
	o.Products = v
}

// GetTotalProduct returns the TotalProduct field value if set, zero value otherwise.
func (o *ViewCampaignProducts200ResponseData) GetTotalProduct() int64 {
	if o == nil || IsNil(o.TotalProduct) {
		var ret int64
		return ret
	}
	return *o.TotalProduct
}

// GetTotalProductOk returns a tuple with the TotalProduct field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewCampaignProducts200ResponseData) GetTotalProductOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalProduct) {
		return nil, false
	}
	return o.TotalProduct, true
}

// HasTotalProduct returns a boolean if a field has been set.
func (o *ViewCampaignProducts200ResponseData) HasTotalProduct() bool {
	if o != nil && !IsNil(o.TotalProduct) {
		return true
	}

	return false
}

// SetTotalProduct gets a reference to the given int64 and assigns it to the TotalProduct field.
func (o *ViewCampaignProducts200ResponseData) SetTotalProduct(v int64) {
	o.TotalProduct = &v
}

func (o ViewCampaignProducts200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ViewCampaignProducts200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Products) {
		toSerialize["products"] = o.Products
	}
	if !IsNil(o.TotalProduct) {
		toSerialize["total_product"] = o.TotalProduct
	}
	return toSerialize, nil
}

type NullableViewCampaignProducts200ResponseData struct {
	value *ViewCampaignProducts200ResponseData
	isSet bool
}

func (v NullableViewCampaignProducts200ResponseData) Get() *ViewCampaignProducts200ResponseData {
	return v.value
}

func (v *NullableViewCampaignProducts200ResponseData) Set(val *ViewCampaignProducts200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableViewCampaignProducts200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableViewCampaignProducts200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViewCampaignProducts200ResponseData(val *ViewCampaignProducts200ResponseData) *NullableViewCampaignProducts200ResponseData {
	return &NullableViewCampaignProducts200ResponseData{value: val, isSet: true}
}

func (v NullableViewCampaignProducts200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableViewCampaignProducts200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


