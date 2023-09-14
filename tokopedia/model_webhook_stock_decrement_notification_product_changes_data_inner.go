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

// checks if the WebhookStockDecrementNotificationProductChangesDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookStockDecrementNotificationProductChangesDataInner{}

// WebhookStockDecrementNotificationProductChangesDataInner struct for WebhookStockDecrementNotificationProductChangesDataInner
type WebhookStockDecrementNotificationProductChangesDataInner struct {
	// Action Type of Product Change Notification
	Action *string `json:"action,omitempty"`
	// Product unique identifier
	ProductId *int64 `json:"product_id,omitempty"`
	// Shop unique identifier
	ShopId *int64 `json:"shop_id,omitempty"`
	// Warehouse unique identifier
	WarehouseId *int64 `json:"warehouse_id,omitempty"`
	// New decrement stock value
	Value *string `json:"value,omitempty"`
	// Previous main stock value
	PreviousValue *string `json:"previous_value,omitempty"`
	// UNIX format updated time
	UpdateTime *int64 `json:"update_time,omitempty"`
	// Indicate for default warehouse of the shop. If not implement multilocation project, pick the true response to decrement stock
	IsDefaultWarehouse *bool `json:"is_default_warehouse,omitempty"`
}

// NewWebhookStockDecrementNotificationProductChangesDataInner instantiates a new WebhookStockDecrementNotificationProductChangesDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookStockDecrementNotificationProductChangesDataInner() *WebhookStockDecrementNotificationProductChangesDataInner {
	this := WebhookStockDecrementNotificationProductChangesDataInner{}
	var action string = "stock_decrement"
	this.Action = &action
	return &this
}

// NewWebhookStockDecrementNotificationProductChangesDataInnerWithDefaults instantiates a new WebhookStockDecrementNotificationProductChangesDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookStockDecrementNotificationProductChangesDataInnerWithDefaults() *WebhookStockDecrementNotificationProductChangesDataInner {
	this := WebhookStockDecrementNotificationProductChangesDataInner{}
	var action string = "stock_decrement"
	this.Action = &action
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *WebhookStockDecrementNotificationProductChangesDataInner) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookStockDecrementNotificationProductChangesDataInner) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *WebhookStockDecrementNotificationProductChangesDataInner) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *WebhookStockDecrementNotificationProductChangesDataInner) SetAction(v string) {
	o.Action = &v
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *WebhookStockDecrementNotificationProductChangesDataInner) GetProductId() int64 {
	if o == nil || IsNil(o.ProductId) {
		var ret int64
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookStockDecrementNotificationProductChangesDataInner) GetProductIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ProductId) {
		return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *WebhookStockDecrementNotificationProductChangesDataInner) HasProductId() bool {
	if o != nil && !IsNil(o.ProductId) {
		return true
	}

	return false
}

// SetProductId gets a reference to the given int64 and assigns it to the ProductId field.
func (o *WebhookStockDecrementNotificationProductChangesDataInner) SetProductId(v int64) {
	o.ProductId = &v
}

// GetShopId returns the ShopId field value if set, zero value otherwise.
func (o *WebhookStockDecrementNotificationProductChangesDataInner) GetShopId() int64 {
	if o == nil || IsNil(o.ShopId) {
		var ret int64
		return ret
	}
	return *o.ShopId
}

// GetShopIdOk returns a tuple with the ShopId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookStockDecrementNotificationProductChangesDataInner) GetShopIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ShopId) {
		return nil, false
	}
	return o.ShopId, true
}

// HasShopId returns a boolean if a field has been set.
func (o *WebhookStockDecrementNotificationProductChangesDataInner) HasShopId() bool {
	if o != nil && !IsNil(o.ShopId) {
		return true
	}

	return false
}

// SetShopId gets a reference to the given int64 and assigns it to the ShopId field.
func (o *WebhookStockDecrementNotificationProductChangesDataInner) SetShopId(v int64) {
	o.ShopId = &v
}

// GetWarehouseId returns the WarehouseId field value if set, zero value otherwise.
func (o *WebhookStockDecrementNotificationProductChangesDataInner) GetWarehouseId() int64 {
	if o == nil || IsNil(o.WarehouseId) {
		var ret int64
		return ret
	}
	return *o.WarehouseId
}

// GetWarehouseIdOk returns a tuple with the WarehouseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookStockDecrementNotificationProductChangesDataInner) GetWarehouseIdOk() (*int64, bool) {
	if o == nil || IsNil(o.WarehouseId) {
		return nil, false
	}
	return o.WarehouseId, true
}

// HasWarehouseId returns a boolean if a field has been set.
func (o *WebhookStockDecrementNotificationProductChangesDataInner) HasWarehouseId() bool {
	if o != nil && !IsNil(o.WarehouseId) {
		return true
	}

	return false
}

// SetWarehouseId gets a reference to the given int64 and assigns it to the WarehouseId field.
func (o *WebhookStockDecrementNotificationProductChangesDataInner) SetWarehouseId(v int64) {
	o.WarehouseId = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *WebhookStockDecrementNotificationProductChangesDataInner) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookStockDecrementNotificationProductChangesDataInner) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *WebhookStockDecrementNotificationProductChangesDataInner) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *WebhookStockDecrementNotificationProductChangesDataInner) SetValue(v string) {
	o.Value = &v
}

// GetPreviousValue returns the PreviousValue field value if set, zero value otherwise.
func (o *WebhookStockDecrementNotificationProductChangesDataInner) GetPreviousValue() string {
	if o == nil || IsNil(o.PreviousValue) {
		var ret string
		return ret
	}
	return *o.PreviousValue
}

// GetPreviousValueOk returns a tuple with the PreviousValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookStockDecrementNotificationProductChangesDataInner) GetPreviousValueOk() (*string, bool) {
	if o == nil || IsNil(o.PreviousValue) {
		return nil, false
	}
	return o.PreviousValue, true
}

// HasPreviousValue returns a boolean if a field has been set.
func (o *WebhookStockDecrementNotificationProductChangesDataInner) HasPreviousValue() bool {
	if o != nil && !IsNil(o.PreviousValue) {
		return true
	}

	return false
}

// SetPreviousValue gets a reference to the given string and assigns it to the PreviousValue field.
func (o *WebhookStockDecrementNotificationProductChangesDataInner) SetPreviousValue(v string) {
	o.PreviousValue = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *WebhookStockDecrementNotificationProductChangesDataInner) GetUpdateTime() int64 {
	if o == nil || IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookStockDecrementNotificationProductChangesDataInner) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *WebhookStockDecrementNotificationProductChangesDataInner) HasUpdateTime() bool {
	if o != nil && !IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *WebhookStockDecrementNotificationProductChangesDataInner) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

// GetIsDefaultWarehouse returns the IsDefaultWarehouse field value if set, zero value otherwise.
func (o *WebhookStockDecrementNotificationProductChangesDataInner) GetIsDefaultWarehouse() bool {
	if o == nil || IsNil(o.IsDefaultWarehouse) {
		var ret bool
		return ret
	}
	return *o.IsDefaultWarehouse
}

// GetIsDefaultWarehouseOk returns a tuple with the IsDefaultWarehouse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookStockDecrementNotificationProductChangesDataInner) GetIsDefaultWarehouseOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDefaultWarehouse) {
		return nil, false
	}
	return o.IsDefaultWarehouse, true
}

// HasIsDefaultWarehouse returns a boolean if a field has been set.
func (o *WebhookStockDecrementNotificationProductChangesDataInner) HasIsDefaultWarehouse() bool {
	if o != nil && !IsNil(o.IsDefaultWarehouse) {
		return true
	}

	return false
}

// SetIsDefaultWarehouse gets a reference to the given bool and assigns it to the IsDefaultWarehouse field.
func (o *WebhookStockDecrementNotificationProductChangesDataInner) SetIsDefaultWarehouse(v bool) {
	o.IsDefaultWarehouse = &v
}

func (o WebhookStockDecrementNotificationProductChangesDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookStockDecrementNotificationProductChangesDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.ProductId) {
		toSerialize["product_id"] = o.ProductId
	}
	if !IsNil(o.ShopId) {
		toSerialize["shop_id"] = o.ShopId
	}
	if !IsNil(o.WarehouseId) {
		toSerialize["warehouse_id"] = o.WarehouseId
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.PreviousValue) {
		toSerialize["previous_value"] = o.PreviousValue
	}
	if !IsNil(o.UpdateTime) {
		toSerialize["update_time"] = o.UpdateTime
	}
	if !IsNil(o.IsDefaultWarehouse) {
		toSerialize["is_default_warehouse"] = o.IsDefaultWarehouse
	}
	return toSerialize, nil
}

type NullableWebhookStockDecrementNotificationProductChangesDataInner struct {
	value *WebhookStockDecrementNotificationProductChangesDataInner
	isSet bool
}

func (v NullableWebhookStockDecrementNotificationProductChangesDataInner) Get() *WebhookStockDecrementNotificationProductChangesDataInner {
	return v.value
}

func (v *NullableWebhookStockDecrementNotificationProductChangesDataInner) Set(val *WebhookStockDecrementNotificationProductChangesDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookStockDecrementNotificationProductChangesDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookStockDecrementNotificationProductChangesDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookStockDecrementNotificationProductChangesDataInner(val *WebhookStockDecrementNotificationProductChangesDataInner) *NullableWebhookStockDecrementNotificationProductChangesDataInner {
	return &NullableWebhookStockDecrementNotificationProductChangesDataInner{value: val, isSet: true}
}

func (v NullableWebhookStockDecrementNotificationProductChangesDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookStockDecrementNotificationProductChangesDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

