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

// checks if the GetProductDiscussion200ResponseDataQuestionInnerAnswerInnerAttachedProductInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetProductDiscussion200ResponseDataQuestionInnerAnswerInnerAttachedProductInner{}

// GetProductDiscussion200ResponseDataQuestionInnerAnswerInnerAttachedProductInner struct for GetProductDiscussion200ResponseDataQuestionInnerAnswerInnerAttachedProductInner
type GetProductDiscussion200ResponseDataQuestionInnerAnswerInnerAttachedProductInner struct {
	// Product Unique Identifier
	ProductID *int64 `json:"ProductID,omitempty"`
	// Product Name
	Name *string `json:"Name,omitempty"`
	// Product Price Formatted
	PriceFormatted *string `json:"PriceFormatted,omitempty"`
	// Product URL
	URL *string `json:"URL,omitempty"`
	// Product Thumbnail
	Thumbnail *string `json:"Thumbnail,omitempty"`
}

// NewGetProductDiscussion200ResponseDataQuestionInnerAnswerInnerAttachedProductInner instantiates a new GetProductDiscussion200ResponseDataQuestionInnerAnswerInnerAttachedProductInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetProductDiscussion200ResponseDataQuestionInnerAnswerInnerAttachedProductInner() *GetProductDiscussion200ResponseDataQuestionInnerAnswerInnerAttachedProductInner {
	this := GetProductDiscussion200ResponseDataQuestionInnerAnswerInnerAttachedProductInner{}
	return &this
}

// NewGetProductDiscussion200ResponseDataQuestionInnerAnswerInnerAttachedProductInnerWithDefaults instantiates a new GetProductDiscussion200ResponseDataQuestionInnerAnswerInnerAttachedProductInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetProductDiscussion200ResponseDataQuestionInnerAnswerInnerAttachedProductInnerWithDefaults() *GetProductDiscussion200ResponseDataQuestionInnerAnswerInnerAttachedProductInner {
	this := GetProductDiscussion200ResponseDataQuestionInnerAnswerInnerAttachedProductInner{}
	return &this
}

// GetProductID returns the ProductID field value if set, zero value otherwise.
func (o *GetProductDiscussion200ResponseDataQuestionInnerAnswerInnerAttachedProductInner) GetProductID() int64 {
	if o == nil || IsNil(o.ProductID) {
		var ret int64
		return ret
	}
	return *o.ProductID
}

// GetProductIDOk returns a tuple with the ProductID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductDiscussion200ResponseDataQuestionInnerAnswerInnerAttachedProductInner) GetProductIDOk() (*int64, bool) {
	if o == nil || IsNil(o.ProductID) {
		return nil, false
	}
	return o.ProductID, true
}

// HasProductID returns a boolean if a field has been set.
func (o *GetProductDiscussion200ResponseDataQuestionInnerAnswerInnerAttachedProductInner) HasProductID() bool {
	if o != nil && !IsNil(o.ProductID) {
		return true
	}

	return false
}

// SetProductID gets a reference to the given int64 and assigns it to the ProductID field.
func (o *GetProductDiscussion200ResponseDataQuestionInnerAnswerInnerAttachedProductInner) SetProductID(v int64) {
	o.ProductID = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetProductDiscussion200ResponseDataQuestionInnerAnswerInnerAttachedProductInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductDiscussion200ResponseDataQuestionInnerAnswerInnerAttachedProductInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetProductDiscussion200ResponseDataQuestionInnerAnswerInnerAttachedProductInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetProductDiscussion200ResponseDataQuestionInnerAnswerInnerAttachedProductInner) SetName(v string) {
	o.Name = &v
}

// GetPriceFormatted returns the PriceFormatted field value if set, zero value otherwise.
func (o *GetProductDiscussion200ResponseDataQuestionInnerAnswerInnerAttachedProductInner) GetPriceFormatted() string {
	if o == nil || IsNil(o.PriceFormatted) {
		var ret string
		return ret
	}
	return *o.PriceFormatted
}

// GetPriceFormattedOk returns a tuple with the PriceFormatted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductDiscussion200ResponseDataQuestionInnerAnswerInnerAttachedProductInner) GetPriceFormattedOk() (*string, bool) {
	if o == nil || IsNil(o.PriceFormatted) {
		return nil, false
	}
	return o.PriceFormatted, true
}

// HasPriceFormatted returns a boolean if a field has been set.
func (o *GetProductDiscussion200ResponseDataQuestionInnerAnswerInnerAttachedProductInner) HasPriceFormatted() bool {
	if o != nil && !IsNil(o.PriceFormatted) {
		return true
	}

	return false
}

// SetPriceFormatted gets a reference to the given string and assigns it to the PriceFormatted field.
func (o *GetProductDiscussion200ResponseDataQuestionInnerAnswerInnerAttachedProductInner) SetPriceFormatted(v string) {
	o.PriceFormatted = &v
}

// GetURL returns the URL field value if set, zero value otherwise.
func (o *GetProductDiscussion200ResponseDataQuestionInnerAnswerInnerAttachedProductInner) GetURL() string {
	if o == nil || IsNil(o.URL) {
		var ret string
		return ret
	}
	return *o.URL
}

// GetURLOk returns a tuple with the URL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductDiscussion200ResponseDataQuestionInnerAnswerInnerAttachedProductInner) GetURLOk() (*string, bool) {
	if o == nil || IsNil(o.URL) {
		return nil, false
	}
	return o.URL, true
}

// HasURL returns a boolean if a field has been set.
func (o *GetProductDiscussion200ResponseDataQuestionInnerAnswerInnerAttachedProductInner) HasURL() bool {
	if o != nil && !IsNil(o.URL) {
		return true
	}

	return false
}

// SetURL gets a reference to the given string and assigns it to the URL field.
func (o *GetProductDiscussion200ResponseDataQuestionInnerAnswerInnerAttachedProductInner) SetURL(v string) {
	o.URL = &v
}

// GetThumbnail returns the Thumbnail field value if set, zero value otherwise.
func (o *GetProductDiscussion200ResponseDataQuestionInnerAnswerInnerAttachedProductInner) GetThumbnail() string {
	if o == nil || IsNil(o.Thumbnail) {
		var ret string
		return ret
	}
	return *o.Thumbnail
}

// GetThumbnailOk returns a tuple with the Thumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductDiscussion200ResponseDataQuestionInnerAnswerInnerAttachedProductInner) GetThumbnailOk() (*string, bool) {
	if o == nil || IsNil(o.Thumbnail) {
		return nil, false
	}
	return o.Thumbnail, true
}

// HasThumbnail returns a boolean if a field has been set.
func (o *GetProductDiscussion200ResponseDataQuestionInnerAnswerInnerAttachedProductInner) HasThumbnail() bool {
	if o != nil && !IsNil(o.Thumbnail) {
		return true
	}

	return false
}

// SetThumbnail gets a reference to the given string and assigns it to the Thumbnail field.
func (o *GetProductDiscussion200ResponseDataQuestionInnerAnswerInnerAttachedProductInner) SetThumbnail(v string) {
	o.Thumbnail = &v
}

func (o GetProductDiscussion200ResponseDataQuestionInnerAnswerInnerAttachedProductInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetProductDiscussion200ResponseDataQuestionInnerAnswerInnerAttachedProductInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProductID) {
		toSerialize["ProductID"] = o.ProductID
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.PriceFormatted) {
		toSerialize["PriceFormatted"] = o.PriceFormatted
	}
	if !IsNil(o.URL) {
		toSerialize["URL"] = o.URL
	}
	if !IsNil(o.Thumbnail) {
		toSerialize["Thumbnail"] = o.Thumbnail
	}
	return toSerialize, nil
}

type NullableGetProductDiscussion200ResponseDataQuestionInnerAnswerInnerAttachedProductInner struct {
	value *GetProductDiscussion200ResponseDataQuestionInnerAnswerInnerAttachedProductInner
	isSet bool
}

func (v NullableGetProductDiscussion200ResponseDataQuestionInnerAnswerInnerAttachedProductInner) Get() *GetProductDiscussion200ResponseDataQuestionInnerAnswerInnerAttachedProductInner {
	return v.value
}

func (v *NullableGetProductDiscussion200ResponseDataQuestionInnerAnswerInnerAttachedProductInner) Set(val *GetProductDiscussion200ResponseDataQuestionInnerAnswerInnerAttachedProductInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetProductDiscussion200ResponseDataQuestionInnerAnswerInnerAttachedProductInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetProductDiscussion200ResponseDataQuestionInnerAnswerInnerAttachedProductInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetProductDiscussion200ResponseDataQuestionInnerAnswerInnerAttachedProductInner(val *GetProductDiscussion200ResponseDataQuestionInnerAnswerInnerAttachedProductInner) *NullableGetProductDiscussion200ResponseDataQuestionInnerAnswerInnerAttachedProductInner {
	return &NullableGetProductDiscussion200ResponseDataQuestionInnerAnswerInnerAttachedProductInner{value: val, isSet: true}
}

func (v NullableGetProductDiscussion200ResponseDataQuestionInnerAnswerInnerAttachedProductInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetProductDiscussion200ResponseDataQuestionInnerAnswerInnerAttachedProductInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


