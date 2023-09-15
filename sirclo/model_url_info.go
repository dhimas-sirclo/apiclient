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

// checks if the UrlInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UrlInfo{}

// UrlInfo struct for UrlInfo
type UrlInfo struct {
	UrlChannel *string `json:"url_channel,omitempty"`
	UrlHelpPage *string `json:"url_help_page,omitempty"`
	UrlLogoChannel *string `json:"url_logo_channel,omitempty"`
	UrlSquareLogoChannel *string `json:"url_square_logo_channel,omitempty"`
	UrlOauth2 *string `json:"url_oauth2,omitempty"`
}

// NewUrlInfo instantiates a new UrlInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUrlInfo() *UrlInfo {
	this := UrlInfo{}
	return &this
}

// NewUrlInfoWithDefaults instantiates a new UrlInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUrlInfoWithDefaults() *UrlInfo {
	this := UrlInfo{}
	return &this
}

// GetUrlChannel returns the UrlChannel field value if set, zero value otherwise.
func (o *UrlInfo) GetUrlChannel() string {
	if o == nil || IsNil(o.UrlChannel) {
		var ret string
		return ret
	}
	return *o.UrlChannel
}

// GetUrlChannelOk returns a tuple with the UrlChannel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UrlInfo) GetUrlChannelOk() (*string, bool) {
	if o == nil || IsNil(o.UrlChannel) {
		return nil, false
	}
	return o.UrlChannel, true
}

// HasUrlChannel returns a boolean if a field has been set.
func (o *UrlInfo) HasUrlChannel() bool {
	if o != nil && !IsNil(o.UrlChannel) {
		return true
	}

	return false
}

// SetUrlChannel gets a reference to the given string and assigns it to the UrlChannel field.
func (o *UrlInfo) SetUrlChannel(v string) {
	o.UrlChannel = &v
}

// GetUrlHelpPage returns the UrlHelpPage field value if set, zero value otherwise.
func (o *UrlInfo) GetUrlHelpPage() string {
	if o == nil || IsNil(o.UrlHelpPage) {
		var ret string
		return ret
	}
	return *o.UrlHelpPage
}

// GetUrlHelpPageOk returns a tuple with the UrlHelpPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UrlInfo) GetUrlHelpPageOk() (*string, bool) {
	if o == nil || IsNil(o.UrlHelpPage) {
		return nil, false
	}
	return o.UrlHelpPage, true
}

// HasUrlHelpPage returns a boolean if a field has been set.
func (o *UrlInfo) HasUrlHelpPage() bool {
	if o != nil && !IsNil(o.UrlHelpPage) {
		return true
	}

	return false
}

// SetUrlHelpPage gets a reference to the given string and assigns it to the UrlHelpPage field.
func (o *UrlInfo) SetUrlHelpPage(v string) {
	o.UrlHelpPage = &v
}

// GetUrlLogoChannel returns the UrlLogoChannel field value if set, zero value otherwise.
func (o *UrlInfo) GetUrlLogoChannel() string {
	if o == nil || IsNil(o.UrlLogoChannel) {
		var ret string
		return ret
	}
	return *o.UrlLogoChannel
}

// GetUrlLogoChannelOk returns a tuple with the UrlLogoChannel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UrlInfo) GetUrlLogoChannelOk() (*string, bool) {
	if o == nil || IsNil(o.UrlLogoChannel) {
		return nil, false
	}
	return o.UrlLogoChannel, true
}

// HasUrlLogoChannel returns a boolean if a field has been set.
func (o *UrlInfo) HasUrlLogoChannel() bool {
	if o != nil && !IsNil(o.UrlLogoChannel) {
		return true
	}

	return false
}

// SetUrlLogoChannel gets a reference to the given string and assigns it to the UrlLogoChannel field.
func (o *UrlInfo) SetUrlLogoChannel(v string) {
	o.UrlLogoChannel = &v
}

// GetUrlSquareLogoChannel returns the UrlSquareLogoChannel field value if set, zero value otherwise.
func (o *UrlInfo) GetUrlSquareLogoChannel() string {
	if o == nil || IsNil(o.UrlSquareLogoChannel) {
		var ret string
		return ret
	}
	return *o.UrlSquareLogoChannel
}

// GetUrlSquareLogoChannelOk returns a tuple with the UrlSquareLogoChannel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UrlInfo) GetUrlSquareLogoChannelOk() (*string, bool) {
	if o == nil || IsNil(o.UrlSquareLogoChannel) {
		return nil, false
	}
	return o.UrlSquareLogoChannel, true
}

// HasUrlSquareLogoChannel returns a boolean if a field has been set.
func (o *UrlInfo) HasUrlSquareLogoChannel() bool {
	if o != nil && !IsNil(o.UrlSquareLogoChannel) {
		return true
	}

	return false
}

// SetUrlSquareLogoChannel gets a reference to the given string and assigns it to the UrlSquareLogoChannel field.
func (o *UrlInfo) SetUrlSquareLogoChannel(v string) {
	o.UrlSquareLogoChannel = &v
}

// GetUrlOauth2 returns the UrlOauth2 field value if set, zero value otherwise.
func (o *UrlInfo) GetUrlOauth2() string {
	if o == nil || IsNil(o.UrlOauth2) {
		var ret string
		return ret
	}
	return *o.UrlOauth2
}

// GetUrlOauth2Ok returns a tuple with the UrlOauth2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UrlInfo) GetUrlOauth2Ok() (*string, bool) {
	if o == nil || IsNil(o.UrlOauth2) {
		return nil, false
	}
	return o.UrlOauth2, true
}

// HasUrlOauth2 returns a boolean if a field has been set.
func (o *UrlInfo) HasUrlOauth2() bool {
	if o != nil && !IsNil(o.UrlOauth2) {
		return true
	}

	return false
}

// SetUrlOauth2 gets a reference to the given string and assigns it to the UrlOauth2 field.
func (o *UrlInfo) SetUrlOauth2(v string) {
	o.UrlOauth2 = &v
}

func (o UrlInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UrlInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UrlChannel) {
		toSerialize["url_channel"] = o.UrlChannel
	}
	if !IsNil(o.UrlHelpPage) {
		toSerialize["url_help_page"] = o.UrlHelpPage
	}
	if !IsNil(o.UrlLogoChannel) {
		toSerialize["url_logo_channel"] = o.UrlLogoChannel
	}
	if !IsNil(o.UrlSquareLogoChannel) {
		toSerialize["url_square_logo_channel"] = o.UrlSquareLogoChannel
	}
	if !IsNil(o.UrlOauth2) {
		toSerialize["url_oauth2"] = o.UrlOauth2
	}
	return toSerialize, nil
}

type NullableUrlInfo struct {
	value *UrlInfo
	isSet bool
}

func (v NullableUrlInfo) Get() *UrlInfo {
	return v.value
}

func (v *NullableUrlInfo) Set(val *UrlInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableUrlInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableUrlInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUrlInfo(val *UrlInfo) *NullableUrlInfo {
	return &NullableUrlInfo{value: val, isSet: true}
}

func (v NullableUrlInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUrlInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


