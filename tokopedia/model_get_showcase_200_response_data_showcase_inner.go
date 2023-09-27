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

// checks if the GetShowcase200ResponseDataShowcaseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetShowcase200ResponseDataShowcaseInner{}

// GetShowcase200ResponseDataShowcaseInner struct for GetShowcase200ResponseDataShowcaseInner
type GetShowcase200ResponseDataShowcaseInner struct {
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"Name,omitempty"`
	Alias *string `json:"alias,omitempty"`
	Uri *string `json:"uri,omitempty"`
	ProductCount *string `json:"product_count,omitempty"`
	IsHighlighted *bool `json:"is_highlighted,omitempty"`
	Badge *string `json:"badge,omitempty"`
	AceDefaultSort *int64 `json:"ace_default_sort,omitempty"`
}

// NewGetShowcase200ResponseDataShowcaseInner instantiates a new GetShowcase200ResponseDataShowcaseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetShowcase200ResponseDataShowcaseInner() *GetShowcase200ResponseDataShowcaseInner {
	this := GetShowcase200ResponseDataShowcaseInner{}
	return &this
}

// NewGetShowcase200ResponseDataShowcaseInnerWithDefaults instantiates a new GetShowcase200ResponseDataShowcaseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetShowcase200ResponseDataShowcaseInnerWithDefaults() *GetShowcase200ResponseDataShowcaseInner {
	this := GetShowcase200ResponseDataShowcaseInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetShowcase200ResponseDataShowcaseInner) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetShowcase200ResponseDataShowcaseInner) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetShowcase200ResponseDataShowcaseInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *GetShowcase200ResponseDataShowcaseInner) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetShowcase200ResponseDataShowcaseInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetShowcase200ResponseDataShowcaseInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetShowcase200ResponseDataShowcaseInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetShowcase200ResponseDataShowcaseInner) SetName(v string) {
	o.Name = &v
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *GetShowcase200ResponseDataShowcaseInner) GetAlias() string {
	if o == nil || IsNil(o.Alias) {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetShowcase200ResponseDataShowcaseInner) GetAliasOk() (*string, bool) {
	if o == nil || IsNil(o.Alias) {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *GetShowcase200ResponseDataShowcaseInner) HasAlias() bool {
	if o != nil && !IsNil(o.Alias) {
		return true
	}

	return false
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *GetShowcase200ResponseDataShowcaseInner) SetAlias(v string) {
	o.Alias = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *GetShowcase200ResponseDataShowcaseInner) GetUri() string {
	if o == nil || IsNil(o.Uri) {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetShowcase200ResponseDataShowcaseInner) GetUriOk() (*string, bool) {
	if o == nil || IsNil(o.Uri) {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *GetShowcase200ResponseDataShowcaseInner) HasUri() bool {
	if o != nil && !IsNil(o.Uri) {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *GetShowcase200ResponseDataShowcaseInner) SetUri(v string) {
	o.Uri = &v
}

// GetProductCount returns the ProductCount field value if set, zero value otherwise.
func (o *GetShowcase200ResponseDataShowcaseInner) GetProductCount() string {
	if o == nil || IsNil(o.ProductCount) {
		var ret string
		return ret
	}
	return *o.ProductCount
}

// GetProductCountOk returns a tuple with the ProductCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetShowcase200ResponseDataShowcaseInner) GetProductCountOk() (*string, bool) {
	if o == nil || IsNil(o.ProductCount) {
		return nil, false
	}
	return o.ProductCount, true
}

// HasProductCount returns a boolean if a field has been set.
func (o *GetShowcase200ResponseDataShowcaseInner) HasProductCount() bool {
	if o != nil && !IsNil(o.ProductCount) {
		return true
	}

	return false
}

// SetProductCount gets a reference to the given string and assigns it to the ProductCount field.
func (o *GetShowcase200ResponseDataShowcaseInner) SetProductCount(v string) {
	o.ProductCount = &v
}

// GetIsHighlighted returns the IsHighlighted field value if set, zero value otherwise.
func (o *GetShowcase200ResponseDataShowcaseInner) GetIsHighlighted() bool {
	if o == nil || IsNil(o.IsHighlighted) {
		var ret bool
		return ret
	}
	return *o.IsHighlighted
}

// GetIsHighlightedOk returns a tuple with the IsHighlighted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetShowcase200ResponseDataShowcaseInner) GetIsHighlightedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsHighlighted) {
		return nil, false
	}
	return o.IsHighlighted, true
}

// HasIsHighlighted returns a boolean if a field has been set.
func (o *GetShowcase200ResponseDataShowcaseInner) HasIsHighlighted() bool {
	if o != nil && !IsNil(o.IsHighlighted) {
		return true
	}

	return false
}

// SetIsHighlighted gets a reference to the given bool and assigns it to the IsHighlighted field.
func (o *GetShowcase200ResponseDataShowcaseInner) SetIsHighlighted(v bool) {
	o.IsHighlighted = &v
}

// GetBadge returns the Badge field value if set, zero value otherwise.
func (o *GetShowcase200ResponseDataShowcaseInner) GetBadge() string {
	if o == nil || IsNil(o.Badge) {
		var ret string
		return ret
	}
	return *o.Badge
}

// GetBadgeOk returns a tuple with the Badge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetShowcase200ResponseDataShowcaseInner) GetBadgeOk() (*string, bool) {
	if o == nil || IsNil(o.Badge) {
		return nil, false
	}
	return o.Badge, true
}

// HasBadge returns a boolean if a field has been set.
func (o *GetShowcase200ResponseDataShowcaseInner) HasBadge() bool {
	if o != nil && !IsNil(o.Badge) {
		return true
	}

	return false
}

// SetBadge gets a reference to the given string and assigns it to the Badge field.
func (o *GetShowcase200ResponseDataShowcaseInner) SetBadge(v string) {
	o.Badge = &v
}

// GetAceDefaultSort returns the AceDefaultSort field value if set, zero value otherwise.
func (o *GetShowcase200ResponseDataShowcaseInner) GetAceDefaultSort() int64 {
	if o == nil || IsNil(o.AceDefaultSort) {
		var ret int64
		return ret
	}
	return *o.AceDefaultSort
}

// GetAceDefaultSortOk returns a tuple with the AceDefaultSort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetShowcase200ResponseDataShowcaseInner) GetAceDefaultSortOk() (*int64, bool) {
	if o == nil || IsNil(o.AceDefaultSort) {
		return nil, false
	}
	return o.AceDefaultSort, true
}

// HasAceDefaultSort returns a boolean if a field has been set.
func (o *GetShowcase200ResponseDataShowcaseInner) HasAceDefaultSort() bool {
	if o != nil && !IsNil(o.AceDefaultSort) {
		return true
	}

	return false
}

// SetAceDefaultSort gets a reference to the given int64 and assigns it to the AceDefaultSort field.
func (o *GetShowcase200ResponseDataShowcaseInner) SetAceDefaultSort(v int64) {
	o.AceDefaultSort = &v
}

func (o GetShowcase200ResponseDataShowcaseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetShowcase200ResponseDataShowcaseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.Alias) {
		toSerialize["alias"] = o.Alias
	}
	if !IsNil(o.Uri) {
		toSerialize["uri"] = o.Uri
	}
	if !IsNil(o.ProductCount) {
		toSerialize["product_count"] = o.ProductCount
	}
	if !IsNil(o.IsHighlighted) {
		toSerialize["is_highlighted"] = o.IsHighlighted
	}
	if !IsNil(o.Badge) {
		toSerialize["badge"] = o.Badge
	}
	if !IsNil(o.AceDefaultSort) {
		toSerialize["ace_default_sort"] = o.AceDefaultSort
	}
	return toSerialize, nil
}

type NullableGetShowcase200ResponseDataShowcaseInner struct {
	value *GetShowcase200ResponseDataShowcaseInner
	isSet bool
}

func (v NullableGetShowcase200ResponseDataShowcaseInner) Get() *GetShowcase200ResponseDataShowcaseInner {
	return v.value
}

func (v *NullableGetShowcase200ResponseDataShowcaseInner) Set(val *GetShowcase200ResponseDataShowcaseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetShowcase200ResponseDataShowcaseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetShowcase200ResponseDataShowcaseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetShowcase200ResponseDataShowcaseInner(val *GetShowcase200ResponseDataShowcaseInner) *NullableGetShowcase200ResponseDataShowcaseInner {
	return &NullableGetShowcase200ResponseDataShowcaseInner{value: val, isSet: true}
}

func (v NullableGetShowcase200ResponseDataShowcaseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetShowcase200ResponseDataShowcaseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

