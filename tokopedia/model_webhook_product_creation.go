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

// checks if the WebhookProductCreation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookProductCreation{}

// WebhookProductCreation This notification will be trigger when using Create Product V2.0.[https://developer.tokopedia.com/openapi/guide/api-reference/tokopedia/product-api/create-product]
type WebhookProductCreation struct {
	// Upload id of the product to check Notification
	UploadId *int64 `json:"upload_id,omitempty"`
	// Shop unique identifier
	ShopId *int64 `json:"shop_id,omitempty"`
	// User unique identifier
	UserId *int64 `json:"user_id,omitempty"`
	// Fulfillment service unique identifier
	FsId *int64 `json:"fs_id,omitempty"`
	Status *string `json:"status,omitempty"`
	TotalData *int64 `json:"total_data,omitempty"`
	SuccessRowData []WebhookProductCreationSuccessRowDataInner `json:"success_row_data,omitempty"`
}

// NewWebhookProductCreation instantiates a new WebhookProductCreation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookProductCreation() *WebhookProductCreation {
	this := WebhookProductCreation{}
	return &this
}

// NewWebhookProductCreationWithDefaults instantiates a new WebhookProductCreation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookProductCreationWithDefaults() *WebhookProductCreation {
	this := WebhookProductCreation{}
	return &this
}

// GetUploadId returns the UploadId field value if set, zero value otherwise.
func (o *WebhookProductCreation) GetUploadId() int64 {
	if o == nil || IsNil(o.UploadId) {
		var ret int64
		return ret
	}
	return *o.UploadId
}

// GetUploadIdOk returns a tuple with the UploadId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookProductCreation) GetUploadIdOk() (*int64, bool) {
	if o == nil || IsNil(o.UploadId) {
		return nil, false
	}
	return o.UploadId, true
}

// HasUploadId returns a boolean if a field has been set.
func (o *WebhookProductCreation) HasUploadId() bool {
	if o != nil && !IsNil(o.UploadId) {
		return true
	}

	return false
}

// SetUploadId gets a reference to the given int64 and assigns it to the UploadId field.
func (o *WebhookProductCreation) SetUploadId(v int64) {
	o.UploadId = &v
}

// GetShopId returns the ShopId field value if set, zero value otherwise.
func (o *WebhookProductCreation) GetShopId() int64 {
	if o == nil || IsNil(o.ShopId) {
		var ret int64
		return ret
	}
	return *o.ShopId
}

// GetShopIdOk returns a tuple with the ShopId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookProductCreation) GetShopIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ShopId) {
		return nil, false
	}
	return o.ShopId, true
}

// HasShopId returns a boolean if a field has been set.
func (o *WebhookProductCreation) HasShopId() bool {
	if o != nil && !IsNil(o.ShopId) {
		return true
	}

	return false
}

// SetShopId gets a reference to the given int64 and assigns it to the ShopId field.
func (o *WebhookProductCreation) SetShopId(v int64) {
	o.ShopId = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *WebhookProductCreation) GetUserId() int64 {
	if o == nil || IsNil(o.UserId) {
		var ret int64
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookProductCreation) GetUserIdOk() (*int64, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *WebhookProductCreation) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int64 and assigns it to the UserId field.
func (o *WebhookProductCreation) SetUserId(v int64) {
	o.UserId = &v
}

// GetFsId returns the FsId field value if set, zero value otherwise.
func (o *WebhookProductCreation) GetFsId() int64 {
	if o == nil || IsNil(o.FsId) {
		var ret int64
		return ret
	}
	return *o.FsId
}

// GetFsIdOk returns a tuple with the FsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookProductCreation) GetFsIdOk() (*int64, bool) {
	if o == nil || IsNil(o.FsId) {
		return nil, false
	}
	return o.FsId, true
}

// HasFsId returns a boolean if a field has been set.
func (o *WebhookProductCreation) HasFsId() bool {
	if o != nil && !IsNil(o.FsId) {
		return true
	}

	return false
}

// SetFsId gets a reference to the given int64 and assigns it to the FsId field.
func (o *WebhookProductCreation) SetFsId(v int64) {
	o.FsId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *WebhookProductCreation) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookProductCreation) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *WebhookProductCreation) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *WebhookProductCreation) SetStatus(v string) {
	o.Status = &v
}

// GetTotalData returns the TotalData field value if set, zero value otherwise.
func (o *WebhookProductCreation) GetTotalData() int64 {
	if o == nil || IsNil(o.TotalData) {
		var ret int64
		return ret
	}
	return *o.TotalData
}

// GetTotalDataOk returns a tuple with the TotalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookProductCreation) GetTotalDataOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalData) {
		return nil, false
	}
	return o.TotalData, true
}

// HasTotalData returns a boolean if a field has been set.
func (o *WebhookProductCreation) HasTotalData() bool {
	if o != nil && !IsNil(o.TotalData) {
		return true
	}

	return false
}

// SetTotalData gets a reference to the given int64 and assigns it to the TotalData field.
func (o *WebhookProductCreation) SetTotalData(v int64) {
	o.TotalData = &v
}

// GetSuccessRowData returns the SuccessRowData field value if set, zero value otherwise.
func (o *WebhookProductCreation) GetSuccessRowData() []WebhookProductCreationSuccessRowDataInner {
	if o == nil || IsNil(o.SuccessRowData) {
		var ret []WebhookProductCreationSuccessRowDataInner
		return ret
	}
	return o.SuccessRowData
}

// GetSuccessRowDataOk returns a tuple with the SuccessRowData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookProductCreation) GetSuccessRowDataOk() ([]WebhookProductCreationSuccessRowDataInner, bool) {
	if o == nil || IsNil(o.SuccessRowData) {
		return nil, false
	}
	return o.SuccessRowData, true
}

// HasSuccessRowData returns a boolean if a field has been set.
func (o *WebhookProductCreation) HasSuccessRowData() bool {
	if o != nil && !IsNil(o.SuccessRowData) {
		return true
	}

	return false
}

// SetSuccessRowData gets a reference to the given []WebhookProductCreationSuccessRowDataInner and assigns it to the SuccessRowData field.
func (o *WebhookProductCreation) SetSuccessRowData(v []WebhookProductCreationSuccessRowDataInner) {
	o.SuccessRowData = v
}

func (o WebhookProductCreation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookProductCreation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UploadId) {
		toSerialize["upload_id"] = o.UploadId
	}
	if !IsNil(o.ShopId) {
		toSerialize["shop_id"] = o.ShopId
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	if !IsNil(o.FsId) {
		toSerialize["fs_id"] = o.FsId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.TotalData) {
		toSerialize["total_data"] = o.TotalData
	}
	if !IsNil(o.SuccessRowData) {
		toSerialize["success_row_data"] = o.SuccessRowData
	}
	return toSerialize, nil
}

type NullableWebhookProductCreation struct {
	value *WebhookProductCreation
	isSet bool
}

func (v NullableWebhookProductCreation) Get() *WebhookProductCreation {
	return v.value
}

func (v *NullableWebhookProductCreation) Set(val *WebhookProductCreation) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookProductCreation) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookProductCreation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookProductCreation(val *WebhookProductCreation) *NullableWebhookProductCreation {
	return &NullableWebhookProductCreation{value: val, isSet: true}
}

func (v NullableWebhookProductCreation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookProductCreation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

