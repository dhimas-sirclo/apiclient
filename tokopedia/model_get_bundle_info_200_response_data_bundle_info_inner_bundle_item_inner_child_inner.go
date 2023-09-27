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

// checks if the GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerChildInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerChildInner{}

// GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerChildInner struct for GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerChildInner
type GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerChildInner struct {
	ProductId *int64 `json:"product_id,omitempty"`
	Name *string `json:"Name,omitempty"`
	PicUrl *string `json:"pic_url,omitempty"`
	MinOrder *int64 `json:"min_order,omitempty"`
	BundlePrice *int64 `json:"bundle_price,omitempty"`
	OriginalPrice *int64 `json:"original_price,omitempty"`
	Stock *int64 `json:"stock,omitempty"`
	OptionId []int64 `json:"option_id,omitempty"`
	IsBuyable *bool `json:"is_buyable,omitempty"`
}

// NewGetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerChildInner instantiates a new GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerChildInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerChildInner() *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerChildInner {
	this := GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerChildInner{}
	return &this
}

// NewGetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerChildInnerWithDefaults instantiates a new GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerChildInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerChildInnerWithDefaults() *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerChildInner {
	this := GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerChildInner{}
	return &this
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerChildInner) GetProductId() int64 {
	if o == nil || IsNil(o.ProductId) {
		var ret int64
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerChildInner) GetProductIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ProductId) {
		return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerChildInner) HasProductId() bool {
	if o != nil && !IsNil(o.ProductId) {
		return true
	}

	return false
}

// SetProductId gets a reference to the given int64 and assigns it to the ProductId field.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerChildInner) SetProductId(v int64) {
	o.ProductId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerChildInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerChildInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerChildInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerChildInner) SetName(v string) {
	o.Name = &v
}

// GetPicUrl returns the PicUrl field value if set, zero value otherwise.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerChildInner) GetPicUrl() string {
	if o == nil || IsNil(o.PicUrl) {
		var ret string
		return ret
	}
	return *o.PicUrl
}

// GetPicUrlOk returns a tuple with the PicUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerChildInner) GetPicUrlOk() (*string, bool) {
	if o == nil || IsNil(o.PicUrl) {
		return nil, false
	}
	return o.PicUrl, true
}

// HasPicUrl returns a boolean if a field has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerChildInner) HasPicUrl() bool {
	if o != nil && !IsNil(o.PicUrl) {
		return true
	}

	return false
}

// SetPicUrl gets a reference to the given string and assigns it to the PicUrl field.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerChildInner) SetPicUrl(v string) {
	o.PicUrl = &v
}

// GetMinOrder returns the MinOrder field value if set, zero value otherwise.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerChildInner) GetMinOrder() int64 {
	if o == nil || IsNil(o.MinOrder) {
		var ret int64
		return ret
	}
	return *o.MinOrder
}

// GetMinOrderOk returns a tuple with the MinOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerChildInner) GetMinOrderOk() (*int64, bool) {
	if o == nil || IsNil(o.MinOrder) {
		return nil, false
	}
	return o.MinOrder, true
}

// HasMinOrder returns a boolean if a field has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerChildInner) HasMinOrder() bool {
	if o != nil && !IsNil(o.MinOrder) {
		return true
	}

	return false
}

// SetMinOrder gets a reference to the given int64 and assigns it to the MinOrder field.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerChildInner) SetMinOrder(v int64) {
	o.MinOrder = &v
}

// GetBundlePrice returns the BundlePrice field value if set, zero value otherwise.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerChildInner) GetBundlePrice() int64 {
	if o == nil || IsNil(o.BundlePrice) {
		var ret int64
		return ret
	}
	return *o.BundlePrice
}

// GetBundlePriceOk returns a tuple with the BundlePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerChildInner) GetBundlePriceOk() (*int64, bool) {
	if o == nil || IsNil(o.BundlePrice) {
		return nil, false
	}
	return o.BundlePrice, true
}

// HasBundlePrice returns a boolean if a field has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerChildInner) HasBundlePrice() bool {
	if o != nil && !IsNil(o.BundlePrice) {
		return true
	}

	return false
}

// SetBundlePrice gets a reference to the given int64 and assigns it to the BundlePrice field.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerChildInner) SetBundlePrice(v int64) {
	o.BundlePrice = &v
}

// GetOriginalPrice returns the OriginalPrice field value if set, zero value otherwise.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerChildInner) GetOriginalPrice() int64 {
	if o == nil || IsNil(o.OriginalPrice) {
		var ret int64
		return ret
	}
	return *o.OriginalPrice
}

// GetOriginalPriceOk returns a tuple with the OriginalPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerChildInner) GetOriginalPriceOk() (*int64, bool) {
	if o == nil || IsNil(o.OriginalPrice) {
		return nil, false
	}
	return o.OriginalPrice, true
}

// HasOriginalPrice returns a boolean if a field has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerChildInner) HasOriginalPrice() bool {
	if o != nil && !IsNil(o.OriginalPrice) {
		return true
	}

	return false
}

// SetOriginalPrice gets a reference to the given int64 and assigns it to the OriginalPrice field.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerChildInner) SetOriginalPrice(v int64) {
	o.OriginalPrice = &v
}

// GetStock returns the Stock field value if set, zero value otherwise.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerChildInner) GetStock() int64 {
	if o == nil || IsNil(o.Stock) {
		var ret int64
		return ret
	}
	return *o.Stock
}

// GetStockOk returns a tuple with the Stock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerChildInner) GetStockOk() (*int64, bool) {
	if o == nil || IsNil(o.Stock) {
		return nil, false
	}
	return o.Stock, true
}

// HasStock returns a boolean if a field has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerChildInner) HasStock() bool {
	if o != nil && !IsNil(o.Stock) {
		return true
	}

	return false
}

// SetStock gets a reference to the given int64 and assigns it to the Stock field.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerChildInner) SetStock(v int64) {
	o.Stock = &v
}

// GetOptionId returns the OptionId field value if set, zero value otherwise.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerChildInner) GetOptionId() []int64 {
	if o == nil || IsNil(o.OptionId) {
		var ret []int64
		return ret
	}
	return o.OptionId
}

// GetOptionIdOk returns a tuple with the OptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerChildInner) GetOptionIdOk() ([]int64, bool) {
	if o == nil || IsNil(o.OptionId) {
		return nil, false
	}
	return o.OptionId, true
}

// HasOptionId returns a boolean if a field has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerChildInner) HasOptionId() bool {
	if o != nil && !IsNil(o.OptionId) {
		return true
	}

	return false
}

// SetOptionId gets a reference to the given []int64 and assigns it to the OptionId field.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerChildInner) SetOptionId(v []int64) {
	o.OptionId = v
}

// GetIsBuyable returns the IsBuyable field value if set, zero value otherwise.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerChildInner) GetIsBuyable() bool {
	if o == nil || IsNil(o.IsBuyable) {
		var ret bool
		return ret
	}
	return *o.IsBuyable
}

// GetIsBuyableOk returns a tuple with the IsBuyable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerChildInner) GetIsBuyableOk() (*bool, bool) {
	if o == nil || IsNil(o.IsBuyable) {
		return nil, false
	}
	return o.IsBuyable, true
}

// HasIsBuyable returns a boolean if a field has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerChildInner) HasIsBuyable() bool {
	if o != nil && !IsNil(o.IsBuyable) {
		return true
	}

	return false
}

// SetIsBuyable gets a reference to the given bool and assigns it to the IsBuyable field.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerChildInner) SetIsBuyable(v bool) {
	o.IsBuyable = &v
}

func (o GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerChildInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerChildInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProductId) {
		toSerialize["product_id"] = o.ProductId
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.PicUrl) {
		toSerialize["pic_url"] = o.PicUrl
	}
	if !IsNil(o.MinOrder) {
		toSerialize["min_order"] = o.MinOrder
	}
	if !IsNil(o.BundlePrice) {
		toSerialize["bundle_price"] = o.BundlePrice
	}
	if !IsNil(o.OriginalPrice) {
		toSerialize["original_price"] = o.OriginalPrice
	}
	if !IsNil(o.Stock) {
		toSerialize["stock"] = o.Stock
	}
	if !IsNil(o.OptionId) {
		toSerialize["option_id"] = o.OptionId
	}
	if !IsNil(o.IsBuyable) {
		toSerialize["is_buyable"] = o.IsBuyable
	}
	return toSerialize, nil
}

type NullableGetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerChildInner struct {
	value *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerChildInner
	isSet bool
}

func (v NullableGetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerChildInner) Get() *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerChildInner {
	return v.value
}

func (v *NullableGetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerChildInner) Set(val *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerChildInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerChildInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerChildInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerChildInner(val *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerChildInner) *NullableGetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerChildInner {
	return &NullableGetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerChildInner{value: val, isSet: true}
}

func (v NullableGetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerChildInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerChildInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


