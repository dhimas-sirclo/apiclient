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

// checks if the WebhookProductDiscussionNewQuestionDiscussionData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookProductDiscussionNewQuestionDiscussionData{}

// WebhookProductDiscussionNewQuestionDiscussionData struct for WebhookProductDiscussionNewQuestionDiscussionData
type WebhookProductDiscussionNewQuestionDiscussionData struct {
	Id *string `json:"id,omitempty"`
	ProductId *int64 `json:"product_id,omitempty"`
	ShopId *int64 `json:"shop_id,omitempty"`
	UserId *int64 `json:"user_id,omitempty"`
	Message *string `json:"message,omitempty"`
	CreateTime *WebhookProductDiscussionNewQuestionDiscussionDataCreateTime `json:"create_time,omitempty"`
}

// NewWebhookProductDiscussionNewQuestionDiscussionData instantiates a new WebhookProductDiscussionNewQuestionDiscussionData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookProductDiscussionNewQuestionDiscussionData() *WebhookProductDiscussionNewQuestionDiscussionData {
	this := WebhookProductDiscussionNewQuestionDiscussionData{}
	return &this
}

// NewWebhookProductDiscussionNewQuestionDiscussionDataWithDefaults instantiates a new WebhookProductDiscussionNewQuestionDiscussionData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookProductDiscussionNewQuestionDiscussionDataWithDefaults() *WebhookProductDiscussionNewQuestionDiscussionData {
	this := WebhookProductDiscussionNewQuestionDiscussionData{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WebhookProductDiscussionNewQuestionDiscussionData) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookProductDiscussionNewQuestionDiscussionData) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WebhookProductDiscussionNewQuestionDiscussionData) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WebhookProductDiscussionNewQuestionDiscussionData) SetId(v string) {
	o.Id = &v
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *WebhookProductDiscussionNewQuestionDiscussionData) GetProductId() int64 {
	if o == nil || IsNil(o.ProductId) {
		var ret int64
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookProductDiscussionNewQuestionDiscussionData) GetProductIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ProductId) {
		return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *WebhookProductDiscussionNewQuestionDiscussionData) HasProductId() bool {
	if o != nil && !IsNil(o.ProductId) {
		return true
	}

	return false
}

// SetProductId gets a reference to the given int64 and assigns it to the ProductId field.
func (o *WebhookProductDiscussionNewQuestionDiscussionData) SetProductId(v int64) {
	o.ProductId = &v
}

// GetShopId returns the ShopId field value if set, zero value otherwise.
func (o *WebhookProductDiscussionNewQuestionDiscussionData) GetShopId() int64 {
	if o == nil || IsNil(o.ShopId) {
		var ret int64
		return ret
	}
	return *o.ShopId
}

// GetShopIdOk returns a tuple with the ShopId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookProductDiscussionNewQuestionDiscussionData) GetShopIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ShopId) {
		return nil, false
	}
	return o.ShopId, true
}

// HasShopId returns a boolean if a field has been set.
func (o *WebhookProductDiscussionNewQuestionDiscussionData) HasShopId() bool {
	if o != nil && !IsNil(o.ShopId) {
		return true
	}

	return false
}

// SetShopId gets a reference to the given int64 and assigns it to the ShopId field.
func (o *WebhookProductDiscussionNewQuestionDiscussionData) SetShopId(v int64) {
	o.ShopId = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *WebhookProductDiscussionNewQuestionDiscussionData) GetUserId() int64 {
	if o == nil || IsNil(o.UserId) {
		var ret int64
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookProductDiscussionNewQuestionDiscussionData) GetUserIdOk() (*int64, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *WebhookProductDiscussionNewQuestionDiscussionData) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int64 and assigns it to the UserId field.
func (o *WebhookProductDiscussionNewQuestionDiscussionData) SetUserId(v int64) {
	o.UserId = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *WebhookProductDiscussionNewQuestionDiscussionData) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookProductDiscussionNewQuestionDiscussionData) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *WebhookProductDiscussionNewQuestionDiscussionData) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *WebhookProductDiscussionNewQuestionDiscussionData) SetMessage(v string) {
	o.Message = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *WebhookProductDiscussionNewQuestionDiscussionData) GetCreateTime() WebhookProductDiscussionNewQuestionDiscussionDataCreateTime {
	if o == nil || IsNil(o.CreateTime) {
		var ret WebhookProductDiscussionNewQuestionDiscussionDataCreateTime
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookProductDiscussionNewQuestionDiscussionData) GetCreateTimeOk() (*WebhookProductDiscussionNewQuestionDiscussionDataCreateTime, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *WebhookProductDiscussionNewQuestionDiscussionData) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given WebhookProductDiscussionNewQuestionDiscussionDataCreateTime and assigns it to the CreateTime field.
func (o *WebhookProductDiscussionNewQuestionDiscussionData) SetCreateTime(v WebhookProductDiscussionNewQuestionDiscussionDataCreateTime) {
	o.CreateTime = &v
}

func (o WebhookProductDiscussionNewQuestionDiscussionData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookProductDiscussionNewQuestionDiscussionData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ProductId) {
		toSerialize["product_id"] = o.ProductId
	}
	if !IsNil(o.ShopId) {
		toSerialize["shop_id"] = o.ShopId
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.CreateTime) {
		toSerialize["create_time"] = o.CreateTime
	}
	return toSerialize, nil
}

type NullableWebhookProductDiscussionNewQuestionDiscussionData struct {
	value *WebhookProductDiscussionNewQuestionDiscussionData
	isSet bool
}

func (v NullableWebhookProductDiscussionNewQuestionDiscussionData) Get() *WebhookProductDiscussionNewQuestionDiscussionData {
	return v.value
}

func (v *NullableWebhookProductDiscussionNewQuestionDiscussionData) Set(val *WebhookProductDiscussionNewQuestionDiscussionData) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookProductDiscussionNewQuestionDiscussionData) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookProductDiscussionNewQuestionDiscussionData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookProductDiscussionNewQuestionDiscussionData(val *WebhookProductDiscussionNewQuestionDiscussionData) *NullableWebhookProductDiscussionNewQuestionDiscussionData {
	return &NullableWebhookProductDiscussionNewQuestionDiscussionData{value: val, isSet: true}
}

func (v NullableWebhookProductDiscussionNewQuestionDiscussionData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookProductDiscussionNewQuestionDiscussionData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


