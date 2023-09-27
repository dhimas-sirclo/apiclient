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

// checks if the GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInnerPicture type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInnerPicture{}

// GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInnerPicture struct for GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInnerPicture
type GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInnerPicture struct {
	FileName *string `json:"file_name,omitempty"`
	FilePath *string `json:"file_path,omitempty"`
	Url *string `json:"url,omitempty"`
	Url100 *string `json:"url100,omitempty"`
	Url200 *string `json:"url200,omitempty"`
}

// NewGetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInnerPicture instantiates a new GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInnerPicture object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInnerPicture() *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInnerPicture {
	this := GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInnerPicture{}
	return &this
}

// NewGetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInnerPictureWithDefaults instantiates a new GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInnerPicture object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInnerPictureWithDefaults() *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInnerPicture {
	this := GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInnerPicture{}
	return &this
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInnerPicture) GetFileName() string {
	if o == nil || IsNil(o.FileName) {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInnerPicture) GetFileNameOk() (*string, bool) {
	if o == nil || IsNil(o.FileName) {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInnerPicture) HasFileName() bool {
	if o != nil && !IsNil(o.FileName) {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInnerPicture) SetFileName(v string) {
	o.FileName = &v
}

// GetFilePath returns the FilePath field value if set, zero value otherwise.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInnerPicture) GetFilePath() string {
	if o == nil || IsNil(o.FilePath) {
		var ret string
		return ret
	}
	return *o.FilePath
}

// GetFilePathOk returns a tuple with the FilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInnerPicture) GetFilePathOk() (*string, bool) {
	if o == nil || IsNil(o.FilePath) {
		return nil, false
	}
	return o.FilePath, true
}

// HasFilePath returns a boolean if a field has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInnerPicture) HasFilePath() bool {
	if o != nil && !IsNil(o.FilePath) {
		return true
	}

	return false
}

// SetFilePath gets a reference to the given string and assigns it to the FilePath field.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInnerPicture) SetFilePath(v string) {
	o.FilePath = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInnerPicture) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInnerPicture) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInnerPicture) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInnerPicture) SetUrl(v string) {
	o.Url = &v
}

// GetUrl100 returns the Url100 field value if set, zero value otherwise.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInnerPicture) GetUrl100() string {
	if o == nil || IsNil(o.Url100) {
		var ret string
		return ret
	}
	return *o.Url100
}

// GetUrl100Ok returns a tuple with the Url100 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInnerPicture) GetUrl100Ok() (*string, bool) {
	if o == nil || IsNil(o.Url100) {
		return nil, false
	}
	return o.Url100, true
}

// HasUrl100 returns a boolean if a field has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInnerPicture) HasUrl100() bool {
	if o != nil && !IsNil(o.Url100) {
		return true
	}

	return false
}

// SetUrl100 gets a reference to the given string and assigns it to the Url100 field.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInnerPicture) SetUrl100(v string) {
	o.Url100 = &v
}

// GetUrl200 returns the Url200 field value if set, zero value otherwise.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInnerPicture) GetUrl200() string {
	if o == nil || IsNil(o.Url200) {
		var ret string
		return ret
	}
	return *o.Url200
}

// GetUrl200Ok returns a tuple with the Url200 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInnerPicture) GetUrl200Ok() (*string, bool) {
	if o == nil || IsNil(o.Url200) {
		return nil, false
	}
	return o.Url200, true
}

// HasUrl200 returns a boolean if a field has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInnerPicture) HasUrl200() bool {
	if o != nil && !IsNil(o.Url200) {
		return true
	}

	return false
}

// SetUrl200 gets a reference to the given string and assigns it to the Url200 field.
func (o *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInnerPicture) SetUrl200(v string) {
	o.Url200 = &v
}

func (o GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInnerPicture) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInnerPicture) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FileName) {
		toSerialize["file_name"] = o.FileName
	}
	if !IsNil(o.FilePath) {
		toSerialize["file_path"] = o.FilePath
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Url100) {
		toSerialize["url100"] = o.Url100
	}
	if !IsNil(o.Url200) {
		toSerialize["url200"] = o.Url200
	}
	return toSerialize, nil
}

type NullableGetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInnerPicture struct {
	value *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInnerPicture
	isSet bool
}

func (v NullableGetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInnerPicture) Get() *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInnerPicture {
	return v.value
}

func (v *NullableGetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInnerPicture) Set(val *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInnerPicture) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInnerPicture) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInnerPicture) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInnerPicture(val *GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInnerPicture) *NullableGetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInnerPicture {
	return &NullableGetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInnerPicture{value: val, isSet: true}
}

func (v NullableGetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInnerPicture) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBundleInfo200ResponseDataBundleInfoInnerBundleItemInnerSelectionInnerOptionInnerPicture) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


