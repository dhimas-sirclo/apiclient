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

// checks if the WebhookActiveSlashPriceNotification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookActiveSlashPriceNotification{}

// WebhookActiveSlashPriceNotification Active Slash Price is notification that show information of active slash prsice event. This notification will be trigger when slash price already started.
type WebhookActiveSlashPriceNotification struct {
	// Action Type of Active Slash Price Notification
	Action *string `json:"action,omitempty"`
	// Fulfillment service unique identifier
	FsId *int64 `json:"fs_id,omitempty"`
	// Product unique identifier
	ProductId *int64 `json:"product_id,omitempty"`
	// Shop unique identifier
	ShopId *int64 `json:"shop_id,omitempty"`
	// Normal Price
	OriginalPrice *int64 `json:"original_price,omitempty"`
	// New Price When Slash Price Already Started
	DiscountedPrice *float64 `json:"discounted_price,omitempty"`
	// Discount Percentage
	DiscountPercentage *int64 `json:"discount_percentage,omitempty"`
	// Slash Price Start Time
	StartDate *string `json:"start_date,omitempty"`
	// Slash Price End Time
	EndDate *string `json:"end_date,omitempty"`
}

// NewWebhookActiveSlashPriceNotification instantiates a new WebhookActiveSlashPriceNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookActiveSlashPriceNotification() *WebhookActiveSlashPriceNotification {
	this := WebhookActiveSlashPriceNotification{}
	var action string = "slash_price_active"
	this.Action = &action
	return &this
}

// NewWebhookActiveSlashPriceNotificationWithDefaults instantiates a new WebhookActiveSlashPriceNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookActiveSlashPriceNotificationWithDefaults() *WebhookActiveSlashPriceNotification {
	this := WebhookActiveSlashPriceNotification{}
	var action string = "slash_price_active"
	this.Action = &action
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *WebhookActiveSlashPriceNotification) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookActiveSlashPriceNotification) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *WebhookActiveSlashPriceNotification) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *WebhookActiveSlashPriceNotification) SetAction(v string) {
	o.Action = &v
}

// GetFsId returns the FsId field value if set, zero value otherwise.
func (o *WebhookActiveSlashPriceNotification) GetFsId() int64 {
	if o == nil || IsNil(o.FsId) {
		var ret int64
		return ret
	}
	return *o.FsId
}

// GetFsIdOk returns a tuple with the FsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookActiveSlashPriceNotification) GetFsIdOk() (*int64, bool) {
	if o == nil || IsNil(o.FsId) {
		return nil, false
	}
	return o.FsId, true
}

// HasFsId returns a boolean if a field has been set.
func (o *WebhookActiveSlashPriceNotification) HasFsId() bool {
	if o != nil && !IsNil(o.FsId) {
		return true
	}

	return false
}

// SetFsId gets a reference to the given int64 and assigns it to the FsId field.
func (o *WebhookActiveSlashPriceNotification) SetFsId(v int64) {
	o.FsId = &v
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *WebhookActiveSlashPriceNotification) GetProductId() int64 {
	if o == nil || IsNil(o.ProductId) {
		var ret int64
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookActiveSlashPriceNotification) GetProductIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ProductId) {
		return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *WebhookActiveSlashPriceNotification) HasProductId() bool {
	if o != nil && !IsNil(o.ProductId) {
		return true
	}

	return false
}

// SetProductId gets a reference to the given int64 and assigns it to the ProductId field.
func (o *WebhookActiveSlashPriceNotification) SetProductId(v int64) {
	o.ProductId = &v
}

// GetShopId returns the ShopId field value if set, zero value otherwise.
func (o *WebhookActiveSlashPriceNotification) GetShopId() int64 {
	if o == nil || IsNil(o.ShopId) {
		var ret int64
		return ret
	}
	return *o.ShopId
}

// GetShopIdOk returns a tuple with the ShopId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookActiveSlashPriceNotification) GetShopIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ShopId) {
		return nil, false
	}
	return o.ShopId, true
}

// HasShopId returns a boolean if a field has been set.
func (o *WebhookActiveSlashPriceNotification) HasShopId() bool {
	if o != nil && !IsNil(o.ShopId) {
		return true
	}

	return false
}

// SetShopId gets a reference to the given int64 and assigns it to the ShopId field.
func (o *WebhookActiveSlashPriceNotification) SetShopId(v int64) {
	o.ShopId = &v
}

// GetOriginalPrice returns the OriginalPrice field value if set, zero value otherwise.
func (o *WebhookActiveSlashPriceNotification) GetOriginalPrice() int64 {
	if o == nil || IsNil(o.OriginalPrice) {
		var ret int64
		return ret
	}
	return *o.OriginalPrice
}

// GetOriginalPriceOk returns a tuple with the OriginalPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookActiveSlashPriceNotification) GetOriginalPriceOk() (*int64, bool) {
	if o == nil || IsNil(o.OriginalPrice) {
		return nil, false
	}
	return o.OriginalPrice, true
}

// HasOriginalPrice returns a boolean if a field has been set.
func (o *WebhookActiveSlashPriceNotification) HasOriginalPrice() bool {
	if o != nil && !IsNil(o.OriginalPrice) {
		return true
	}

	return false
}

// SetOriginalPrice gets a reference to the given int64 and assigns it to the OriginalPrice field.
func (o *WebhookActiveSlashPriceNotification) SetOriginalPrice(v int64) {
	o.OriginalPrice = &v
}

// GetDiscountedPrice returns the DiscountedPrice field value if set, zero value otherwise.
func (o *WebhookActiveSlashPriceNotification) GetDiscountedPrice() float64 {
	if o == nil || IsNil(o.DiscountedPrice) {
		var ret float64
		return ret
	}
	return *o.DiscountedPrice
}

// GetDiscountedPriceOk returns a tuple with the DiscountedPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookActiveSlashPriceNotification) GetDiscountedPriceOk() (*float64, bool) {
	if o == nil || IsNil(o.DiscountedPrice) {
		return nil, false
	}
	return o.DiscountedPrice, true
}

// HasDiscountedPrice returns a boolean if a field has been set.
func (o *WebhookActiveSlashPriceNotification) HasDiscountedPrice() bool {
	if o != nil && !IsNil(o.DiscountedPrice) {
		return true
	}

	return false
}

// SetDiscountedPrice gets a reference to the given float64 and assigns it to the DiscountedPrice field.
func (o *WebhookActiveSlashPriceNotification) SetDiscountedPrice(v float64) {
	o.DiscountedPrice = &v
}

// GetDiscountPercentage returns the DiscountPercentage field value if set, zero value otherwise.
func (o *WebhookActiveSlashPriceNotification) GetDiscountPercentage() int64 {
	if o == nil || IsNil(o.DiscountPercentage) {
		var ret int64
		return ret
	}
	return *o.DiscountPercentage
}

// GetDiscountPercentageOk returns a tuple with the DiscountPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookActiveSlashPriceNotification) GetDiscountPercentageOk() (*int64, bool) {
	if o == nil || IsNil(o.DiscountPercentage) {
		return nil, false
	}
	return o.DiscountPercentage, true
}

// HasDiscountPercentage returns a boolean if a field has been set.
func (o *WebhookActiveSlashPriceNotification) HasDiscountPercentage() bool {
	if o != nil && !IsNil(o.DiscountPercentage) {
		return true
	}

	return false
}

// SetDiscountPercentage gets a reference to the given int64 and assigns it to the DiscountPercentage field.
func (o *WebhookActiveSlashPriceNotification) SetDiscountPercentage(v int64) {
	o.DiscountPercentage = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *WebhookActiveSlashPriceNotification) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookActiveSlashPriceNotification) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *WebhookActiveSlashPriceNotification) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *WebhookActiveSlashPriceNotification) SetStartDate(v string) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *WebhookActiveSlashPriceNotification) GetEndDate() string {
	if o == nil || IsNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookActiveSlashPriceNotification) GetEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *WebhookActiveSlashPriceNotification) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *WebhookActiveSlashPriceNotification) SetEndDate(v string) {
	o.EndDate = &v
}

func (o WebhookActiveSlashPriceNotification) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookActiveSlashPriceNotification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.FsId) {
		toSerialize["fs_id"] = o.FsId
	}
	if !IsNil(o.ProductId) {
		toSerialize["product_id"] = o.ProductId
	}
	if !IsNil(o.ShopId) {
		toSerialize["shop_id"] = o.ShopId
	}
	if !IsNil(o.OriginalPrice) {
		toSerialize["original_price"] = o.OriginalPrice
	}
	if !IsNil(o.DiscountedPrice) {
		toSerialize["discounted_price"] = o.DiscountedPrice
	}
	if !IsNil(o.DiscountPercentage) {
		toSerialize["discount_percentage"] = o.DiscountPercentage
	}
	if !IsNil(o.StartDate) {
		toSerialize["start_date"] = o.StartDate
	}
	if !IsNil(o.EndDate) {
		toSerialize["end_date"] = o.EndDate
	}
	return toSerialize, nil
}

type NullableWebhookActiveSlashPriceNotification struct {
	value *WebhookActiveSlashPriceNotification
	isSet bool
}

func (v NullableWebhookActiveSlashPriceNotification) Get() *WebhookActiveSlashPriceNotification {
	return v.value
}

func (v *NullableWebhookActiveSlashPriceNotification) Set(val *WebhookActiveSlashPriceNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookActiveSlashPriceNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookActiveSlashPriceNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookActiveSlashPriceNotification(val *WebhookActiveSlashPriceNotification) *NullableWebhookActiveSlashPriceNotification {
	return &NullableWebhookActiveSlashPriceNotification{value: val, isSet: true}
}

func (v NullableWebhookActiveSlashPriceNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookActiveSlashPriceNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


