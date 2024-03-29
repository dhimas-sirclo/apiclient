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

// checks if the GetShowcase200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetShowcase200ResponseData{}

// GetShowcase200ResponseData struct for GetShowcase200ResponseData
type GetShowcase200ResponseData struct {
	Showcase []GetShowcase200ResponseDataShowcaseInner `json:"showcase,omitempty"`
	ShowcaseGroup []GetShowcase200ResponseDataShowcaseGroupInner `json:"showcase_group,omitempty"`
	UseAce *bool `json:"use_ace,omitempty"`
	ShowcaseWithoutAce []int64 `json:"showcase_without_ace,omitempty"`
	PrevLink *string `json:"prev_link,omitempty"`
	NextLink *string `json:"next_link,omitempty"`
}

// NewGetShowcase200ResponseData instantiates a new GetShowcase200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetShowcase200ResponseData() *GetShowcase200ResponseData {
	this := GetShowcase200ResponseData{}
	return &this
}

// NewGetShowcase200ResponseDataWithDefaults instantiates a new GetShowcase200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetShowcase200ResponseDataWithDefaults() *GetShowcase200ResponseData {
	this := GetShowcase200ResponseData{}
	return &this
}

// GetShowcase returns the Showcase field value if set, zero value otherwise.
func (o *GetShowcase200ResponseData) GetShowcase() []GetShowcase200ResponseDataShowcaseInner {
	if o == nil || IsNil(o.Showcase) {
		var ret []GetShowcase200ResponseDataShowcaseInner
		return ret
	}
	return o.Showcase
}

// GetShowcaseOk returns a tuple with the Showcase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetShowcase200ResponseData) GetShowcaseOk() ([]GetShowcase200ResponseDataShowcaseInner, bool) {
	if o == nil || IsNil(o.Showcase) {
		return nil, false
	}
	return o.Showcase, true
}

// HasShowcase returns a boolean if a field has been set.
func (o *GetShowcase200ResponseData) HasShowcase() bool {
	if o != nil && !IsNil(o.Showcase) {
		return true
	}

	return false
}

// SetShowcase gets a reference to the given []GetShowcase200ResponseDataShowcaseInner and assigns it to the Showcase field.
func (o *GetShowcase200ResponseData) SetShowcase(v []GetShowcase200ResponseDataShowcaseInner) {
	o.Showcase = v
}

// GetShowcaseGroup returns the ShowcaseGroup field value if set, zero value otherwise.
func (o *GetShowcase200ResponseData) GetShowcaseGroup() []GetShowcase200ResponseDataShowcaseGroupInner {
	if o == nil || IsNil(o.ShowcaseGroup) {
		var ret []GetShowcase200ResponseDataShowcaseGroupInner
		return ret
	}
	return o.ShowcaseGroup
}

// GetShowcaseGroupOk returns a tuple with the ShowcaseGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetShowcase200ResponseData) GetShowcaseGroupOk() ([]GetShowcase200ResponseDataShowcaseGroupInner, bool) {
	if o == nil || IsNil(o.ShowcaseGroup) {
		return nil, false
	}
	return o.ShowcaseGroup, true
}

// HasShowcaseGroup returns a boolean if a field has been set.
func (o *GetShowcase200ResponseData) HasShowcaseGroup() bool {
	if o != nil && !IsNil(o.ShowcaseGroup) {
		return true
	}

	return false
}

// SetShowcaseGroup gets a reference to the given []GetShowcase200ResponseDataShowcaseGroupInner and assigns it to the ShowcaseGroup field.
func (o *GetShowcase200ResponseData) SetShowcaseGroup(v []GetShowcase200ResponseDataShowcaseGroupInner) {
	o.ShowcaseGroup = v
}

// GetUseAce returns the UseAce field value if set, zero value otherwise.
func (o *GetShowcase200ResponseData) GetUseAce() bool {
	if o == nil || IsNil(o.UseAce) {
		var ret bool
		return ret
	}
	return *o.UseAce
}

// GetUseAceOk returns a tuple with the UseAce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetShowcase200ResponseData) GetUseAceOk() (*bool, bool) {
	if o == nil || IsNil(o.UseAce) {
		return nil, false
	}
	return o.UseAce, true
}

// HasUseAce returns a boolean if a field has been set.
func (o *GetShowcase200ResponseData) HasUseAce() bool {
	if o != nil && !IsNil(o.UseAce) {
		return true
	}

	return false
}

// SetUseAce gets a reference to the given bool and assigns it to the UseAce field.
func (o *GetShowcase200ResponseData) SetUseAce(v bool) {
	o.UseAce = &v
}

// GetShowcaseWithoutAce returns the ShowcaseWithoutAce field value if set, zero value otherwise.
func (o *GetShowcase200ResponseData) GetShowcaseWithoutAce() []int64 {
	if o == nil || IsNil(o.ShowcaseWithoutAce) {
		var ret []int64
		return ret
	}
	return o.ShowcaseWithoutAce
}

// GetShowcaseWithoutAceOk returns a tuple with the ShowcaseWithoutAce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetShowcase200ResponseData) GetShowcaseWithoutAceOk() ([]int64, bool) {
	if o == nil || IsNil(o.ShowcaseWithoutAce) {
		return nil, false
	}
	return o.ShowcaseWithoutAce, true
}

// HasShowcaseWithoutAce returns a boolean if a field has been set.
func (o *GetShowcase200ResponseData) HasShowcaseWithoutAce() bool {
	if o != nil && !IsNil(o.ShowcaseWithoutAce) {
		return true
	}

	return false
}

// SetShowcaseWithoutAce gets a reference to the given []int64 and assigns it to the ShowcaseWithoutAce field.
func (o *GetShowcase200ResponseData) SetShowcaseWithoutAce(v []int64) {
	o.ShowcaseWithoutAce = v
}

// GetPrevLink returns the PrevLink field value if set, zero value otherwise.
func (o *GetShowcase200ResponseData) GetPrevLink() string {
	if o == nil || IsNil(o.PrevLink) {
		var ret string
		return ret
	}
	return *o.PrevLink
}

// GetPrevLinkOk returns a tuple with the PrevLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetShowcase200ResponseData) GetPrevLinkOk() (*string, bool) {
	if o == nil || IsNil(o.PrevLink) {
		return nil, false
	}
	return o.PrevLink, true
}

// HasPrevLink returns a boolean if a field has been set.
func (o *GetShowcase200ResponseData) HasPrevLink() bool {
	if o != nil && !IsNil(o.PrevLink) {
		return true
	}

	return false
}

// SetPrevLink gets a reference to the given string and assigns it to the PrevLink field.
func (o *GetShowcase200ResponseData) SetPrevLink(v string) {
	o.PrevLink = &v
}

// GetNextLink returns the NextLink field value if set, zero value otherwise.
func (o *GetShowcase200ResponseData) GetNextLink() string {
	if o == nil || IsNil(o.NextLink) {
		var ret string
		return ret
	}
	return *o.NextLink
}

// GetNextLinkOk returns a tuple with the NextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetShowcase200ResponseData) GetNextLinkOk() (*string, bool) {
	if o == nil || IsNil(o.NextLink) {
		return nil, false
	}
	return o.NextLink, true
}

// HasNextLink returns a boolean if a field has been set.
func (o *GetShowcase200ResponseData) HasNextLink() bool {
	if o != nil && !IsNil(o.NextLink) {
		return true
	}

	return false
}

// SetNextLink gets a reference to the given string and assigns it to the NextLink field.
func (o *GetShowcase200ResponseData) SetNextLink(v string) {
	o.NextLink = &v
}

func (o GetShowcase200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetShowcase200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Showcase) {
		toSerialize["showcase"] = o.Showcase
	}
	if !IsNil(o.ShowcaseGroup) {
		toSerialize["showcase_group"] = o.ShowcaseGroup
	}
	if !IsNil(o.UseAce) {
		toSerialize["use_ace"] = o.UseAce
	}
	if !IsNil(o.ShowcaseWithoutAce) {
		toSerialize["showcase_without_ace"] = o.ShowcaseWithoutAce
	}
	if !IsNil(o.PrevLink) {
		toSerialize["prev_link"] = o.PrevLink
	}
	if !IsNil(o.NextLink) {
		toSerialize["next_link"] = o.NextLink
	}
	return toSerialize, nil
}

type NullableGetShowcase200ResponseData struct {
	value *GetShowcase200ResponseData
	isSet bool
}

func (v NullableGetShowcase200ResponseData) Get() *GetShowcase200ResponseData {
	return v.value
}

func (v *NullableGetShowcase200ResponseData) Set(val *GetShowcase200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetShowcase200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetShowcase200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetShowcase200ResponseData(val *GetShowcase200ResponseData) *NullableGetShowcase200ResponseData {
	return &NullableGetShowcase200ResponseData{value: val, isSet: true}
}

func (v NullableGetShowcase200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetShowcase200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


