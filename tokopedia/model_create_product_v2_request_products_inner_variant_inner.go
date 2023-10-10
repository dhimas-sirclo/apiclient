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

// checks if the CreateProductV2RequestProductsInnerVariantInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateProductV2RequestProductsInnerVariantInner{}

// CreateProductV2RequestProductsInnerVariantInner struct for CreateProductV2RequestProductsInnerVariantInner
type CreateProductV2RequestProductsInnerVariantInner struct {
	Products []CreateProductV2RequestProductsInnerVariantInnerProductsInner `json:"products"`
	Selection []EditProductV3RequestProductsInnerVariantSelectionInner `json:"selection"`
	Sizecharts []EditProductV2RequestProductsInnerVariantProductsInnerPicturesInner `json:"sizecharts"`
}

// NewCreateProductV2RequestProductsInnerVariantInner instantiates a new CreateProductV2RequestProductsInnerVariantInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateProductV2RequestProductsInnerVariantInner(products []CreateProductV2RequestProductsInnerVariantInnerProductsInner, selection []EditProductV3RequestProductsInnerVariantSelectionInner, sizecharts []EditProductV2RequestProductsInnerVariantProductsInnerPicturesInner) *CreateProductV2RequestProductsInnerVariantInner {
	this := CreateProductV2RequestProductsInnerVariantInner{}
	this.Products = products
	this.Selection = selection
	this.Sizecharts = sizecharts
	return &this
}

// NewCreateProductV2RequestProductsInnerVariantInnerWithDefaults instantiates a new CreateProductV2RequestProductsInnerVariantInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateProductV2RequestProductsInnerVariantInnerWithDefaults() *CreateProductV2RequestProductsInnerVariantInner {
	this := CreateProductV2RequestProductsInnerVariantInner{}
	return &this
}

// GetProducts returns the Products field value
func (o *CreateProductV2RequestProductsInnerVariantInner) GetProducts() []CreateProductV2RequestProductsInnerVariantInnerProductsInner {
	if o == nil {
		var ret []CreateProductV2RequestProductsInnerVariantInnerProductsInner
		return ret
	}

	return o.Products
}

// GetProductsOk returns a tuple with the Products field value
// and a boolean to check if the value has been set.
func (o *CreateProductV2RequestProductsInnerVariantInner) GetProductsOk() ([]CreateProductV2RequestProductsInnerVariantInnerProductsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Products, true
}

// SetProducts sets field value
func (o *CreateProductV2RequestProductsInnerVariantInner) SetProducts(v []CreateProductV2RequestProductsInnerVariantInnerProductsInner) {
	o.Products = v
}

// GetSelection returns the Selection field value
func (o *CreateProductV2RequestProductsInnerVariantInner) GetSelection() []EditProductV3RequestProductsInnerVariantSelectionInner {
	if o == nil {
		var ret []EditProductV3RequestProductsInnerVariantSelectionInner
		return ret
	}

	return o.Selection
}

// GetSelectionOk returns a tuple with the Selection field value
// and a boolean to check if the value has been set.
func (o *CreateProductV2RequestProductsInnerVariantInner) GetSelectionOk() ([]EditProductV3RequestProductsInnerVariantSelectionInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Selection, true
}

// SetSelection sets field value
func (o *CreateProductV2RequestProductsInnerVariantInner) SetSelection(v []EditProductV3RequestProductsInnerVariantSelectionInner) {
	o.Selection = v
}

// GetSizecharts returns the Sizecharts field value
func (o *CreateProductV2RequestProductsInnerVariantInner) GetSizecharts() []EditProductV2RequestProductsInnerVariantProductsInnerPicturesInner {
	if o == nil {
		var ret []EditProductV2RequestProductsInnerVariantProductsInnerPicturesInner
		return ret
	}

	return o.Sizecharts
}

// GetSizechartsOk returns a tuple with the Sizecharts field value
// and a boolean to check if the value has been set.
func (o *CreateProductV2RequestProductsInnerVariantInner) GetSizechartsOk() ([]EditProductV2RequestProductsInnerVariantProductsInnerPicturesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sizecharts, true
}

// SetSizecharts sets field value
func (o *CreateProductV2RequestProductsInnerVariantInner) SetSizecharts(v []EditProductV2RequestProductsInnerVariantProductsInnerPicturesInner) {
	o.Sizecharts = v
}

func (o CreateProductV2RequestProductsInnerVariantInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateProductV2RequestProductsInnerVariantInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["products"] = o.Products
	toSerialize["selection"] = o.Selection
	toSerialize["sizecharts"] = o.Sizecharts
	return toSerialize, nil
}

type NullableCreateProductV2RequestProductsInnerVariantInner struct {
	value *CreateProductV2RequestProductsInnerVariantInner
	isSet bool
}

func (v NullableCreateProductV2RequestProductsInnerVariantInner) Get() *CreateProductV2RequestProductsInnerVariantInner {
	return v.value
}

func (v *NullableCreateProductV2RequestProductsInnerVariantInner) Set(val *CreateProductV2RequestProductsInnerVariantInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateProductV2RequestProductsInnerVariantInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateProductV2RequestProductsInnerVariantInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateProductV2RequestProductsInnerVariantInner(val *CreateProductV2RequestProductsInnerVariantInner) *NullableCreateProductV2RequestProductsInnerVariantInner {
	return &NullableCreateProductV2RequestProductsInnerVariantInner{value: val, isSet: true}
}

func (v NullableCreateProductV2RequestProductsInnerVariantInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateProductV2RequestProductsInnerVariantInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

