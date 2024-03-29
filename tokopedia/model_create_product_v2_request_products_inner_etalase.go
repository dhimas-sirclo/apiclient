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

// checks if the CreateProductV2RequestProductsInnerEtalase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateProductV2RequestProductsInnerEtalase{}

// CreateProductV2RequestProductsInnerEtalase Etalase or Showcase of the product. The object contains id and name. To get available showcase, please check Get Showcase. Required field to input just id with etalase_id responses from Get Showcase
type CreateProductV2RequestProductsInnerEtalase struct {
	Id int64 `json:"id"`
}

// NewCreateProductV2RequestProductsInnerEtalase instantiates a new CreateProductV2RequestProductsInnerEtalase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateProductV2RequestProductsInnerEtalase(id int64) *CreateProductV2RequestProductsInnerEtalase {
	this := CreateProductV2RequestProductsInnerEtalase{}
	this.Id = id
	return &this
}

// NewCreateProductV2RequestProductsInnerEtalaseWithDefaults instantiates a new CreateProductV2RequestProductsInnerEtalase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateProductV2RequestProductsInnerEtalaseWithDefaults() *CreateProductV2RequestProductsInnerEtalase {
	this := CreateProductV2RequestProductsInnerEtalase{}
	return &this
}

// GetId returns the Id field value
func (o *CreateProductV2RequestProductsInnerEtalase) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CreateProductV2RequestProductsInnerEtalase) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CreateProductV2RequestProductsInnerEtalase) SetId(v int64) {
	o.Id = v
}

func (o CreateProductV2RequestProductsInnerEtalase) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateProductV2RequestProductsInnerEtalase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableCreateProductV2RequestProductsInnerEtalase struct {
	value *CreateProductV2RequestProductsInnerEtalase
	isSet bool
}

func (v NullableCreateProductV2RequestProductsInnerEtalase) Get() *CreateProductV2RequestProductsInnerEtalase {
	return v.value
}

func (v *NullableCreateProductV2RequestProductsInnerEtalase) Set(val *CreateProductV2RequestProductsInnerEtalase) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateProductV2RequestProductsInnerEtalase) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateProductV2RequestProductsInnerEtalase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateProductV2RequestProductsInnerEtalase(val *CreateProductV2RequestProductsInnerEtalase) *NullableCreateProductV2RequestProductsInnerEtalase {
	return &NullableCreateProductV2RequestProductsInnerEtalase{value: val, isSet: true}
}

func (v NullableCreateProductV2RequestProductsInnerEtalase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateProductV2RequestProductsInnerEtalase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


