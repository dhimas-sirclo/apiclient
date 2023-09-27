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

// checks if the GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInner{}

// GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInner struct for GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInner
type GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInner struct {
	ProductVariantId *int64 `json:"product_variant_id,omitempty"`
	VariantId *int64 `json:"variant_id,omitempty"`
	VariantUnitId *int64 `json:"variant_unit_id,omitempty"`
	Position *int64 `json:"position,omitempty"`
	Option []GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInner `json:"option,omitempty"`
	Name *string `json:"Name,omitempty"`
	Identifier *string `json:"identifier,omitempty"`
}

// NewGetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInner instantiates a new GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInner() *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInner {
	this := GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInner{}
	return &this
}

// NewGetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerWithDefaults instantiates a new GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerWithDefaults() *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInner {
	this := GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInner{}
	return &this
}

// GetProductVariantId returns the ProductVariantId field value if set, zero value otherwise.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInner) GetProductVariantId() int64 {
	if o == nil || IsNil(o.ProductVariantId) {
		var ret int64
		return ret
	}
	return *o.ProductVariantId
}

// GetProductVariantIdOk returns a tuple with the ProductVariantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInner) GetProductVariantIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ProductVariantId) {
		return nil, false
	}
	return o.ProductVariantId, true
}

// HasProductVariantId returns a boolean if a field has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInner) HasProductVariantId() bool {
	if o != nil && !IsNil(o.ProductVariantId) {
		return true
	}

	return false
}

// SetProductVariantId gets a reference to the given int64 and assigns it to the ProductVariantId field.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInner) SetProductVariantId(v int64) {
	o.ProductVariantId = &v
}

// GetVariantId returns the VariantId field value if set, zero value otherwise.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInner) GetVariantId() int64 {
	if o == nil || IsNil(o.VariantId) {
		var ret int64
		return ret
	}
	return *o.VariantId
}

// GetVariantIdOk returns a tuple with the VariantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInner) GetVariantIdOk() (*int64, bool) {
	if o == nil || IsNil(o.VariantId) {
		return nil, false
	}
	return o.VariantId, true
}

// HasVariantId returns a boolean if a field has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInner) HasVariantId() bool {
	if o != nil && !IsNil(o.VariantId) {
		return true
	}

	return false
}

// SetVariantId gets a reference to the given int64 and assigns it to the VariantId field.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInner) SetVariantId(v int64) {
	o.VariantId = &v
}

// GetVariantUnitId returns the VariantUnitId field value if set, zero value otherwise.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInner) GetVariantUnitId() int64 {
	if o == nil || IsNil(o.VariantUnitId) {
		var ret int64
		return ret
	}
	return *o.VariantUnitId
}

// GetVariantUnitIdOk returns a tuple with the VariantUnitId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInner) GetVariantUnitIdOk() (*int64, bool) {
	if o == nil || IsNil(o.VariantUnitId) {
		return nil, false
	}
	return o.VariantUnitId, true
}

// HasVariantUnitId returns a boolean if a field has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInner) HasVariantUnitId() bool {
	if o != nil && !IsNil(o.VariantUnitId) {
		return true
	}

	return false
}

// SetVariantUnitId gets a reference to the given int64 and assigns it to the VariantUnitId field.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInner) SetVariantUnitId(v int64) {
	o.VariantUnitId = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInner) GetPosition() int64 {
	if o == nil || IsNil(o.Position) {
		var ret int64
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInner) GetPositionOk() (*int64, bool) {
	if o == nil || IsNil(o.Position) {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInner) HasPosition() bool {
	if o != nil && !IsNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given int64 and assigns it to the Position field.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInner) SetPosition(v int64) {
	o.Position = &v
}

// GetOption returns the Option field value if set, zero value otherwise.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInner) GetOption() []GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInner {
	if o == nil || IsNil(o.Option) {
		var ret []GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInner
		return ret
	}
	return o.Option
}

// GetOptionOk returns a tuple with the Option field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInner) GetOptionOk() ([]GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInner, bool) {
	if o == nil || IsNil(o.Option) {
		return nil, false
	}
	return o.Option, true
}

// HasOption returns a boolean if a field has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInner) HasOption() bool {
	if o != nil && !IsNil(o.Option) {
		return true
	}

	return false
}

// SetOption gets a reference to the given []GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInner and assigns it to the Option field.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInner) SetOption(v []GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInner) {
	o.Option = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInner) SetName(v string) {
	o.Name = &v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInner) GetIdentifier() string {
	if o == nil || IsNil(o.Identifier) {
		var ret string
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInner) GetIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.Identifier) {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInner) HasIdentifier() bool {
	if o != nil && !IsNil(o.Identifier) {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given string and assigns it to the Identifier field.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInner) SetIdentifier(v string) {
	o.Identifier = &v
}

func (o GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProductVariantId) {
		toSerialize["product_variant_id"] = o.ProductVariantId
	}
	if !IsNil(o.VariantId) {
		toSerialize["variant_id"] = o.VariantId
	}
	if !IsNil(o.VariantUnitId) {
		toSerialize["variant_unit_id"] = o.VariantUnitId
	}
	if !IsNil(o.Position) {
		toSerialize["position"] = o.Position
	}
	if !IsNil(o.Option) {
		toSerialize["option"] = o.Option
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.Identifier) {
		toSerialize["identifier"] = o.Identifier
	}
	return toSerialize, nil
}

type NullableGetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInner struct {
	value *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInner
	isSet bool
}

func (v NullableGetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInner) Get() *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInner {
	return v.value
}

func (v *NullableGetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInner) Set(val *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInner(val *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInner) *NullableGetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInner {
	return &NullableGetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInner{value: val, isSet: true}
}

func (v NullableGetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


