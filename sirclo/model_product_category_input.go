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

// checks if the ProductCategoryInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductCategoryInput{}

// ProductCategoryInput struct for ProductCategoryInput
type ProductCategoryInput struct {
	ParentCategoryId *string `json:"parent_category_id,omitempty"`
	CategoryId string `json:"category_id"`
	CategoryName string `json:"category_name"`
	IsLeaf bool `json:"is_leaf"`
	Attribute []ProductCategoryAttributeInput `json:"attribute"`
}

// NewProductCategoryInput instantiates a new ProductCategoryInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductCategoryInput(categoryId string, categoryName string, isLeaf bool, attribute []ProductCategoryAttributeInput) *ProductCategoryInput {
	this := ProductCategoryInput{}
	this.CategoryId = categoryId
	this.CategoryName = categoryName
	this.IsLeaf = isLeaf
	this.Attribute = attribute
	return &this
}

// NewProductCategoryInputWithDefaults instantiates a new ProductCategoryInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductCategoryInputWithDefaults() *ProductCategoryInput {
	this := ProductCategoryInput{}
	return &this
}

// GetParentCategoryId returns the ParentCategoryId field value if set, zero value otherwise.
func (o *ProductCategoryInput) GetParentCategoryId() string {
	if o == nil || IsNil(o.ParentCategoryId) {
		var ret string
		return ret
	}
	return *o.ParentCategoryId
}

// GetParentCategoryIdOk returns a tuple with the ParentCategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductCategoryInput) GetParentCategoryIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentCategoryId) {
		return nil, false
	}
	return o.ParentCategoryId, true
}

// HasParentCategoryId returns a boolean if a field has been set.
func (o *ProductCategoryInput) HasParentCategoryId() bool {
	if o != nil && !IsNil(o.ParentCategoryId) {
		return true
	}

	return false
}

// SetParentCategoryId gets a reference to the given string and assigns it to the ParentCategoryId field.
func (o *ProductCategoryInput) SetParentCategoryId(v string) {
	o.ParentCategoryId = &v
}

// GetCategoryId returns the CategoryId field value
func (o *ProductCategoryInput) GetCategoryId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CategoryId
}

// GetCategoryIdOk returns a tuple with the CategoryId field value
// and a boolean to check if the value has been set.
func (o *ProductCategoryInput) GetCategoryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CategoryId, true
}

// SetCategoryId sets field value
func (o *ProductCategoryInput) SetCategoryId(v string) {
	o.CategoryId = v
}

// GetCategoryName returns the CategoryName field value
func (o *ProductCategoryInput) GetCategoryName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CategoryName
}

// GetCategoryNameOk returns a tuple with the CategoryName field value
// and a boolean to check if the value has been set.
func (o *ProductCategoryInput) GetCategoryNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CategoryName, true
}

// SetCategoryName sets field value
func (o *ProductCategoryInput) SetCategoryName(v string) {
	o.CategoryName = v
}

// GetIsLeaf returns the IsLeaf field value
func (o *ProductCategoryInput) GetIsLeaf() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsLeaf
}

// GetIsLeafOk returns a tuple with the IsLeaf field value
// and a boolean to check if the value has been set.
func (o *ProductCategoryInput) GetIsLeafOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsLeaf, true
}

// SetIsLeaf sets field value
func (o *ProductCategoryInput) SetIsLeaf(v bool) {
	o.IsLeaf = v
}

// GetAttribute returns the Attribute field value
func (o *ProductCategoryInput) GetAttribute() []ProductCategoryAttributeInput {
	if o == nil {
		var ret []ProductCategoryAttributeInput
		return ret
	}

	return o.Attribute
}

// GetAttributeOk returns a tuple with the Attribute field value
// and a boolean to check if the value has been set.
func (o *ProductCategoryInput) GetAttributeOk() ([]ProductCategoryAttributeInput, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attribute, true
}

// SetAttribute sets field value
func (o *ProductCategoryInput) SetAttribute(v []ProductCategoryAttributeInput) {
	o.Attribute = v
}

func (o ProductCategoryInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductCategoryInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ParentCategoryId) {
		toSerialize["parent_category_id"] = o.ParentCategoryId
	}
	toSerialize["category_id"] = o.CategoryId
	toSerialize["category_name"] = o.CategoryName
	toSerialize["is_leaf"] = o.IsLeaf
	toSerialize["attribute"] = o.Attribute
	return toSerialize, nil
}

type NullableProductCategoryInput struct {
	value *ProductCategoryInput
	isSet bool
}

func (v NullableProductCategoryInput) Get() *ProductCategoryInput {
	return v.value
}

func (v *NullableProductCategoryInput) Set(val *ProductCategoryInput) {
	v.value = val
	v.isSet = true
}

func (v NullableProductCategoryInput) IsSet() bool {
	return v.isSet
}

func (v *NullableProductCategoryInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductCategoryInput(val *ProductCategoryInput) *NullableProductCategoryInput {
	return &NullableProductCategoryInput{value: val, isSet: true}
}

func (v NullableProductCategoryInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductCategoryInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


