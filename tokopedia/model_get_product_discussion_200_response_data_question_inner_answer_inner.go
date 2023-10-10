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

// checks if the GetProductDiscussion200ResponseDataQuestionInnerAnswerInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetProductDiscussion200ResponseDataQuestionInnerAnswerInner{}

// GetProductDiscussion200ResponseDataQuestionInnerAnswerInner struct for GetProductDiscussion200ResponseDataQuestionInnerAnswerInner
type GetProductDiscussion200ResponseDataQuestionInnerAnswerInner struct {
	// Answer Content
	Content *string `json:"Content,omitempty"`
	// Answer Masked Content
	MaskedContent *string `json:"MaskedContent,omitempty"`
	// Answerer Username
	UserName *string `json:"UserName,omitempty"`
	// User Thumbnail URL
	UserThumbnail *string `json:"UserThumbnail,omitempty"`
	// User Unique Identifier
	UserID *int64 `json:"UserID,omitempty"`
	// Answer Timestamp (format; 2020-12-04T14:14:44Z)
	CreateTime *string `json:"CreateTime,omitempty"`
	// Total Answer Count
	CreateTimeFormatted *string `json:"CreateTimeFormatted,omitempty"`
	AttachedProduct []GetProductDiscussion200ResponseDataQuestionInnerAnswerInnerAttachedProductInner `json:"AttachedProduct,omitempty"`
	// Total Answer Count
	LikeCount *int64 `json:"LikeCount,omitempty"`
	// Total Answer Count
	AnswerID *int64 `json:"AnswerID,omitempty"`
	// Total Answer Count
	IsSeller *bool `json:"IsSeller,omitempty"`
}

// NewGetProductDiscussion200ResponseDataQuestionInnerAnswerInner instantiates a new GetProductDiscussion200ResponseDataQuestionInnerAnswerInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetProductDiscussion200ResponseDataQuestionInnerAnswerInner() *GetProductDiscussion200ResponseDataQuestionInnerAnswerInner {
	this := GetProductDiscussion200ResponseDataQuestionInnerAnswerInner{}
	return &this
}

// NewGetProductDiscussion200ResponseDataQuestionInnerAnswerInnerWithDefaults instantiates a new GetProductDiscussion200ResponseDataQuestionInnerAnswerInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetProductDiscussion200ResponseDataQuestionInnerAnswerInnerWithDefaults() *GetProductDiscussion200ResponseDataQuestionInnerAnswerInner {
	this := GetProductDiscussion200ResponseDataQuestionInnerAnswerInner{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *GetProductDiscussion200ResponseDataQuestionInnerAnswerInner) GetContent() string {
	if o == nil || IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductDiscussion200ResponseDataQuestionInnerAnswerInner) GetContentOk() (*string, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *GetProductDiscussion200ResponseDataQuestionInnerAnswerInner) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *GetProductDiscussion200ResponseDataQuestionInnerAnswerInner) SetContent(v string) {
	o.Content = &v
}

// GetMaskedContent returns the MaskedContent field value if set, zero value otherwise.
func (o *GetProductDiscussion200ResponseDataQuestionInnerAnswerInner) GetMaskedContent() string {
	if o == nil || IsNil(o.MaskedContent) {
		var ret string
		return ret
	}
	return *o.MaskedContent
}

// GetMaskedContentOk returns a tuple with the MaskedContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductDiscussion200ResponseDataQuestionInnerAnswerInner) GetMaskedContentOk() (*string, bool) {
	if o == nil || IsNil(o.MaskedContent) {
		return nil, false
	}
	return o.MaskedContent, true
}

// HasMaskedContent returns a boolean if a field has been set.
func (o *GetProductDiscussion200ResponseDataQuestionInnerAnswerInner) HasMaskedContent() bool {
	if o != nil && !IsNil(o.MaskedContent) {
		return true
	}

	return false
}

// SetMaskedContent gets a reference to the given string and assigns it to the MaskedContent field.
func (o *GetProductDiscussion200ResponseDataQuestionInnerAnswerInner) SetMaskedContent(v string) {
	o.MaskedContent = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *GetProductDiscussion200ResponseDataQuestionInnerAnswerInner) GetUserName() string {
	if o == nil || IsNil(o.UserName) {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductDiscussion200ResponseDataQuestionInnerAnswerInner) GetUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.UserName) {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *GetProductDiscussion200ResponseDataQuestionInnerAnswerInner) HasUserName() bool {
	if o != nil && !IsNil(o.UserName) {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *GetProductDiscussion200ResponseDataQuestionInnerAnswerInner) SetUserName(v string) {
	o.UserName = &v
}

// GetUserThumbnail returns the UserThumbnail field value if set, zero value otherwise.
func (o *GetProductDiscussion200ResponseDataQuestionInnerAnswerInner) GetUserThumbnail() string {
	if o == nil || IsNil(o.UserThumbnail) {
		var ret string
		return ret
	}
	return *o.UserThumbnail
}

// GetUserThumbnailOk returns a tuple with the UserThumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductDiscussion200ResponseDataQuestionInnerAnswerInner) GetUserThumbnailOk() (*string, bool) {
	if o == nil || IsNil(o.UserThumbnail) {
		return nil, false
	}
	return o.UserThumbnail, true
}

// HasUserThumbnail returns a boolean if a field has been set.
func (o *GetProductDiscussion200ResponseDataQuestionInnerAnswerInner) HasUserThumbnail() bool {
	if o != nil && !IsNil(o.UserThumbnail) {
		return true
	}

	return false
}

// SetUserThumbnail gets a reference to the given string and assigns it to the UserThumbnail field.
func (o *GetProductDiscussion200ResponseDataQuestionInnerAnswerInner) SetUserThumbnail(v string) {
	o.UserThumbnail = &v
}

// GetUserID returns the UserID field value if set, zero value otherwise.
func (o *GetProductDiscussion200ResponseDataQuestionInnerAnswerInner) GetUserID() int64 {
	if o == nil || IsNil(o.UserID) {
		var ret int64
		return ret
	}
	return *o.UserID
}

// GetUserIDOk returns a tuple with the UserID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductDiscussion200ResponseDataQuestionInnerAnswerInner) GetUserIDOk() (*int64, bool) {
	if o == nil || IsNil(o.UserID) {
		return nil, false
	}
	return o.UserID, true
}

// HasUserID returns a boolean if a field has been set.
func (o *GetProductDiscussion200ResponseDataQuestionInnerAnswerInner) HasUserID() bool {
	if o != nil && !IsNil(o.UserID) {
		return true
	}

	return false
}

// SetUserID gets a reference to the given int64 and assigns it to the UserID field.
func (o *GetProductDiscussion200ResponseDataQuestionInnerAnswerInner) SetUserID(v int64) {
	o.UserID = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *GetProductDiscussion200ResponseDataQuestionInnerAnswerInner) GetCreateTime() string {
	if o == nil || IsNil(o.CreateTime) {
		var ret string
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductDiscussion200ResponseDataQuestionInnerAnswerInner) GetCreateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *GetProductDiscussion200ResponseDataQuestionInnerAnswerInner) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given string and assigns it to the CreateTime field.
func (o *GetProductDiscussion200ResponseDataQuestionInnerAnswerInner) SetCreateTime(v string) {
	o.CreateTime = &v
}

// GetCreateTimeFormatted returns the CreateTimeFormatted field value if set, zero value otherwise.
func (o *GetProductDiscussion200ResponseDataQuestionInnerAnswerInner) GetCreateTimeFormatted() string {
	if o == nil || IsNil(o.CreateTimeFormatted) {
		var ret string
		return ret
	}
	return *o.CreateTimeFormatted
}

// GetCreateTimeFormattedOk returns a tuple with the CreateTimeFormatted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductDiscussion200ResponseDataQuestionInnerAnswerInner) GetCreateTimeFormattedOk() (*string, bool) {
	if o == nil || IsNil(o.CreateTimeFormatted) {
		return nil, false
	}
	return o.CreateTimeFormatted, true
}

// HasCreateTimeFormatted returns a boolean if a field has been set.
func (o *GetProductDiscussion200ResponseDataQuestionInnerAnswerInner) HasCreateTimeFormatted() bool {
	if o != nil && !IsNil(o.CreateTimeFormatted) {
		return true
	}

	return false
}

// SetCreateTimeFormatted gets a reference to the given string and assigns it to the CreateTimeFormatted field.
func (o *GetProductDiscussion200ResponseDataQuestionInnerAnswerInner) SetCreateTimeFormatted(v string) {
	o.CreateTimeFormatted = &v
}

// GetAttachedProduct returns the AttachedProduct field value if set, zero value otherwise.
func (o *GetProductDiscussion200ResponseDataQuestionInnerAnswerInner) GetAttachedProduct() []GetProductDiscussion200ResponseDataQuestionInnerAnswerInnerAttachedProductInner {
	if o == nil || IsNil(o.AttachedProduct) {
		var ret []GetProductDiscussion200ResponseDataQuestionInnerAnswerInnerAttachedProductInner
		return ret
	}
	return o.AttachedProduct
}

// GetAttachedProductOk returns a tuple with the AttachedProduct field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductDiscussion200ResponseDataQuestionInnerAnswerInner) GetAttachedProductOk() ([]GetProductDiscussion200ResponseDataQuestionInnerAnswerInnerAttachedProductInner, bool) {
	if o == nil || IsNil(o.AttachedProduct) {
		return nil, false
	}
	return o.AttachedProduct, true
}

// HasAttachedProduct returns a boolean if a field has been set.
func (o *GetProductDiscussion200ResponseDataQuestionInnerAnswerInner) HasAttachedProduct() bool {
	if o != nil && !IsNil(o.AttachedProduct) {
		return true
	}

	return false
}

// SetAttachedProduct gets a reference to the given []GetProductDiscussion200ResponseDataQuestionInnerAnswerInnerAttachedProductInner and assigns it to the AttachedProduct field.
func (o *GetProductDiscussion200ResponseDataQuestionInnerAnswerInner) SetAttachedProduct(v []GetProductDiscussion200ResponseDataQuestionInnerAnswerInnerAttachedProductInner) {
	o.AttachedProduct = v
}

// GetLikeCount returns the LikeCount field value if set, zero value otherwise.
func (o *GetProductDiscussion200ResponseDataQuestionInnerAnswerInner) GetLikeCount() int64 {
	if o == nil || IsNil(o.LikeCount) {
		var ret int64
		return ret
	}
	return *o.LikeCount
}

// GetLikeCountOk returns a tuple with the LikeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductDiscussion200ResponseDataQuestionInnerAnswerInner) GetLikeCountOk() (*int64, bool) {
	if o == nil || IsNil(o.LikeCount) {
		return nil, false
	}
	return o.LikeCount, true
}

// HasLikeCount returns a boolean if a field has been set.
func (o *GetProductDiscussion200ResponseDataQuestionInnerAnswerInner) HasLikeCount() bool {
	if o != nil && !IsNil(o.LikeCount) {
		return true
	}

	return false
}

// SetLikeCount gets a reference to the given int64 and assigns it to the LikeCount field.
func (o *GetProductDiscussion200ResponseDataQuestionInnerAnswerInner) SetLikeCount(v int64) {
	o.LikeCount = &v
}

// GetAnswerID returns the AnswerID field value if set, zero value otherwise.
func (o *GetProductDiscussion200ResponseDataQuestionInnerAnswerInner) GetAnswerID() int64 {
	if o == nil || IsNil(o.AnswerID) {
		var ret int64
		return ret
	}
	return *o.AnswerID
}

// GetAnswerIDOk returns a tuple with the AnswerID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductDiscussion200ResponseDataQuestionInnerAnswerInner) GetAnswerIDOk() (*int64, bool) {
	if o == nil || IsNil(o.AnswerID) {
		return nil, false
	}
	return o.AnswerID, true
}

// HasAnswerID returns a boolean if a field has been set.
func (o *GetProductDiscussion200ResponseDataQuestionInnerAnswerInner) HasAnswerID() bool {
	if o != nil && !IsNil(o.AnswerID) {
		return true
	}

	return false
}

// SetAnswerID gets a reference to the given int64 and assigns it to the AnswerID field.
func (o *GetProductDiscussion200ResponseDataQuestionInnerAnswerInner) SetAnswerID(v int64) {
	o.AnswerID = &v
}

// GetIsSeller returns the IsSeller field value if set, zero value otherwise.
func (o *GetProductDiscussion200ResponseDataQuestionInnerAnswerInner) GetIsSeller() bool {
	if o == nil || IsNil(o.IsSeller) {
		var ret bool
		return ret
	}
	return *o.IsSeller
}

// GetIsSellerOk returns a tuple with the IsSeller field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductDiscussion200ResponseDataQuestionInnerAnswerInner) GetIsSellerOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSeller) {
		return nil, false
	}
	return o.IsSeller, true
}

// HasIsSeller returns a boolean if a field has been set.
func (o *GetProductDiscussion200ResponseDataQuestionInnerAnswerInner) HasIsSeller() bool {
	if o != nil && !IsNil(o.IsSeller) {
		return true
	}

	return false
}

// SetIsSeller gets a reference to the given bool and assigns it to the IsSeller field.
func (o *GetProductDiscussion200ResponseDataQuestionInnerAnswerInner) SetIsSeller(v bool) {
	o.IsSeller = &v
}

func (o GetProductDiscussion200ResponseDataQuestionInnerAnswerInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetProductDiscussion200ResponseDataQuestionInnerAnswerInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Content) {
		toSerialize["Content"] = o.Content
	}
	if !IsNil(o.MaskedContent) {
		toSerialize["MaskedContent"] = o.MaskedContent
	}
	if !IsNil(o.UserName) {
		toSerialize["UserName"] = o.UserName
	}
	if !IsNil(o.UserThumbnail) {
		toSerialize["UserThumbnail"] = o.UserThumbnail
	}
	if !IsNil(o.UserID) {
		toSerialize["UserID"] = o.UserID
	}
	if !IsNil(o.CreateTime) {
		toSerialize["CreateTime"] = o.CreateTime
	}
	if !IsNil(o.CreateTimeFormatted) {
		toSerialize["CreateTimeFormatted"] = o.CreateTimeFormatted
	}
	if !IsNil(o.AttachedProduct) {
		toSerialize["AttachedProduct"] = o.AttachedProduct
	}
	if !IsNil(o.LikeCount) {
		toSerialize["LikeCount"] = o.LikeCount
	}
	if !IsNil(o.AnswerID) {
		toSerialize["AnswerID"] = o.AnswerID
	}
	if !IsNil(o.IsSeller) {
		toSerialize["IsSeller"] = o.IsSeller
	}
	return toSerialize, nil
}

type NullableGetProductDiscussion200ResponseDataQuestionInnerAnswerInner struct {
	value *GetProductDiscussion200ResponseDataQuestionInnerAnswerInner
	isSet bool
}

func (v NullableGetProductDiscussion200ResponseDataQuestionInnerAnswerInner) Get() *GetProductDiscussion200ResponseDataQuestionInnerAnswerInner {
	return v.value
}

func (v *NullableGetProductDiscussion200ResponseDataQuestionInnerAnswerInner) Set(val *GetProductDiscussion200ResponseDataQuestionInnerAnswerInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetProductDiscussion200ResponseDataQuestionInnerAnswerInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetProductDiscussion200ResponseDataQuestionInnerAnswerInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetProductDiscussion200ResponseDataQuestionInnerAnswerInner(val *GetProductDiscussion200ResponseDataQuestionInnerAnswerInner) *NullableGetProductDiscussion200ResponseDataQuestionInnerAnswerInner {
	return &NullableGetProductDiscussion200ResponseDataQuestionInnerAnswerInner{value: val, isSet: true}
}

func (v NullableGetProductDiscussion200ResponseDataQuestionInnerAnswerInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetProductDiscussion200ResponseDataQuestionInnerAnswerInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

