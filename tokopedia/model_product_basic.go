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

// checks if the ProductBasic type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductBasic{}

// ProductBasic struct for ProductBasic
type ProductBasic struct {
	// Product Unique Identifier
	ProductID *int32 `json:"productID,omitempty"`
	// Shop Unique Identifier
	ShopID *int32 `json:"shopID,omitempty"`
	// Product Status (e.g., -2 for Banned, -1 for Pending, 0 for Deleted, 1 for Active, 2 for Best (Feature Product), 3 for Inactive (Warehouse))
	Status *int32 `json:"status,omitempty"`
	// Product Name
	Name *string `json:"Name,omitempty"`
	// Product Condition (e.g., 1 for New, 2 for Used)
	Condition *int32 `json:"condition,omitempty"`
	// Product Category Unique Identifier
	ChildCategoryID *int32 `json:"childCategoryID,omitempty"`
	// Product Description
	ShortDesc *string `json:"shortDesc,omitempty"`
}

// NewProductBasic instantiates a new ProductBasic object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductBasic() *ProductBasic {
	this := ProductBasic{}
	return &this
}

// NewProductBasicWithDefaults instantiates a new ProductBasic object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductBasicWithDefaults() *ProductBasic {
	this := ProductBasic{}
	return &this
}

// GetProductID returns the ProductID field value if set, zero value otherwise.
func (o *ProductBasic) GetProductID() int32 {
	if o == nil || IsNil(o.ProductID) {
		var ret int32
		return ret
	}
	return *o.ProductID
}

// GetProductIDOk returns a tuple with the ProductID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductBasic) GetProductIDOk() (*int32, bool) {
	if o == nil || IsNil(o.ProductID) {
		return nil, false
	}
	return o.ProductID, true
}

// HasProductID returns a boolean if a field has been set.
func (o *ProductBasic) HasProductID() bool {
	if o != nil && !IsNil(o.ProductID) {
		return true
	}

	return false
}

// SetProductID gets a reference to the given int32 and assigns it to the ProductID field.
func (o *ProductBasic) SetProductID(v int32) {
	o.ProductID = &v
}

// GetShopID returns the ShopID field value if set, zero value otherwise.
func (o *ProductBasic) GetShopID() int32 {
	if o == nil || IsNil(o.ShopID) {
		var ret int32
		return ret
	}
	return *o.ShopID
}

// GetShopIDOk returns a tuple with the ShopID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductBasic) GetShopIDOk() (*int32, bool) {
	if o == nil || IsNil(o.ShopID) {
		return nil, false
	}
	return o.ShopID, true
}

// HasShopID returns a boolean if a field has been set.
func (o *ProductBasic) HasShopID() bool {
	if o != nil && !IsNil(o.ShopID) {
		return true
	}

	return false
}

// SetShopID gets a reference to the given int32 and assigns it to the ShopID field.
func (o *ProductBasic) SetShopID(v int32) {
	o.ShopID = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ProductBasic) GetStatus() int32 {
	if o == nil || IsNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductBasic) GetStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ProductBasic) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *ProductBasic) SetStatus(v int32) {
	o.Status = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProductBasic) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductBasic) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProductBasic) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProductBasic) SetName(v string) {
	o.Name = &v
}

// GetCondition returns the Condition field value if set, zero value otherwise.
func (o *ProductBasic) GetCondition() int32 {
	if o == nil || IsNil(o.Condition) {
		var ret int32
		return ret
	}
	return *o.Condition
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductBasic) GetConditionOk() (*int32, bool) {
	if o == nil || IsNil(o.Condition) {
		return nil, false
	}
	return o.Condition, true
}

// HasCondition returns a boolean if a field has been set.
func (o *ProductBasic) HasCondition() bool {
	if o != nil && !IsNil(o.Condition) {
		return true
	}

	return false
}

// SetCondition gets a reference to the given int32 and assigns it to the Condition field.
func (o *ProductBasic) SetCondition(v int32) {
	o.Condition = &v
}

// GetChildCategoryID returns the ChildCategoryID field value if set, zero value otherwise.
func (o *ProductBasic) GetChildCategoryID() int32 {
	if o == nil || IsNil(o.ChildCategoryID) {
		var ret int32
		return ret
	}
	return *o.ChildCategoryID
}

// GetChildCategoryIDOk returns a tuple with the ChildCategoryID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductBasic) GetChildCategoryIDOk() (*int32, bool) {
	if o == nil || IsNil(o.ChildCategoryID) {
		return nil, false
	}
	return o.ChildCategoryID, true
}

// HasChildCategoryID returns a boolean if a field has been set.
func (o *ProductBasic) HasChildCategoryID() bool {
	if o != nil && !IsNil(o.ChildCategoryID) {
		return true
	}

	return false
}

// SetChildCategoryID gets a reference to the given int32 and assigns it to the ChildCategoryID field.
func (o *ProductBasic) SetChildCategoryID(v int32) {
	o.ChildCategoryID = &v
}

// GetShortDesc returns the ShortDesc field value if set, zero value otherwise.
func (o *ProductBasic) GetShortDesc() string {
	if o == nil || IsNil(o.ShortDesc) {
		var ret string
		return ret
	}
	return *o.ShortDesc
}

// GetShortDescOk returns a tuple with the ShortDesc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductBasic) GetShortDescOk() (*string, bool) {
	if o == nil || IsNil(o.ShortDesc) {
		return nil, false
	}
	return o.ShortDesc, true
}

// HasShortDesc returns a boolean if a field has been set.
func (o *ProductBasic) HasShortDesc() bool {
	if o != nil && !IsNil(o.ShortDesc) {
		return true
	}

	return false
}

// SetShortDesc gets a reference to the given string and assigns it to the ShortDesc field.
func (o *ProductBasic) SetShortDesc(v string) {
	o.ShortDesc = &v
}

func (o ProductBasic) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductBasic) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProductID) {
		toSerialize["productID"] = o.ProductID
	}
	if !IsNil(o.ShopID) {
		toSerialize["shopID"] = o.ShopID
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.Condition) {
		toSerialize["condition"] = o.Condition
	}
	if !IsNil(o.ChildCategoryID) {
		toSerialize["childCategoryID"] = o.ChildCategoryID
	}
	if !IsNil(o.ShortDesc) {
		toSerialize["shortDesc"] = o.ShortDesc
	}
	return toSerialize, nil
}

type NullableProductBasic struct {
	value *ProductBasic
	isSet bool
}

func (v NullableProductBasic) Get() *ProductBasic {
	return v.value
}

func (v *NullableProductBasic) Set(val *ProductBasic) {
	v.value = val
	v.isSet = true
}

func (v NullableProductBasic) IsSet() bool {
	return v.isSet
}

func (v *NullableProductBasic) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductBasic(val *ProductBasic) *NullableProductBasic {
	return &NullableProductBasic{value: val, isSet: true}
}

func (v NullableProductBasic) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductBasic) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

