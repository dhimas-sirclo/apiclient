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

// checks if the GetOrderWebhook200ResponseDataOneOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrderWebhook200ResponseDataOneOf{}

// GetOrderWebhook200ResponseDataOneOf struct for GetOrderWebhook200ResponseDataOneOf
type GetOrderWebhook200ResponseDataOneOf struct {
	FsId *int64 `json:"fs_id,omitempty"`
	OrderStatus *int64 `json:"order_status,omitempty"`
	OrderId *int64 `json:"order_id,omitempty"`
	ShopId *int64 `json:"shop_id,omitempty"`
	WarehouseId *int64 `json:"warehouse_id,omitempty"`
	ProductDetails []GetOrderWebhook200ResponseDataOneOfProductDetailsInner `json:"product_details,omitempty"`
}

// NewGetOrderWebhook200ResponseDataOneOf instantiates a new GetOrderWebhook200ResponseDataOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrderWebhook200ResponseDataOneOf() *GetOrderWebhook200ResponseDataOneOf {
	this := GetOrderWebhook200ResponseDataOneOf{}
	return &this
}

// NewGetOrderWebhook200ResponseDataOneOfWithDefaults instantiates a new GetOrderWebhook200ResponseDataOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrderWebhook200ResponseDataOneOfWithDefaults() *GetOrderWebhook200ResponseDataOneOf {
	this := GetOrderWebhook200ResponseDataOneOf{}
	return &this
}

// GetFsId returns the FsId field value if set, zero value otherwise.
func (o *GetOrderWebhook200ResponseDataOneOf) GetFsId() int64 {
	if o == nil || IsNil(o.FsId) {
		var ret int64
		return ret
	}
	return *o.FsId
}

// GetFsIdOk returns a tuple with the FsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderWebhook200ResponseDataOneOf) GetFsIdOk() (*int64, bool) {
	if o == nil || IsNil(o.FsId) {
		return nil, false
	}
	return o.FsId, true
}

// HasFsId returns a boolean if a field has been set.
func (o *GetOrderWebhook200ResponseDataOneOf) HasFsId() bool {
	if o != nil && !IsNil(o.FsId) {
		return true
	}

	return false
}

// SetFsId gets a reference to the given int64 and assigns it to the FsId field.
func (o *GetOrderWebhook200ResponseDataOneOf) SetFsId(v int64) {
	o.FsId = &v
}

// GetOrderStatus returns the OrderStatus field value if set, zero value otherwise.
func (o *GetOrderWebhook200ResponseDataOneOf) GetOrderStatus() int64 {
	if o == nil || IsNil(o.OrderStatus) {
		var ret int64
		return ret
	}
	return *o.OrderStatus
}

// GetOrderStatusOk returns a tuple with the OrderStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderWebhook200ResponseDataOneOf) GetOrderStatusOk() (*int64, bool) {
	if o == nil || IsNil(o.OrderStatus) {
		return nil, false
	}
	return o.OrderStatus, true
}

// HasOrderStatus returns a boolean if a field has been set.
func (o *GetOrderWebhook200ResponseDataOneOf) HasOrderStatus() bool {
	if o != nil && !IsNil(o.OrderStatus) {
		return true
	}

	return false
}

// SetOrderStatus gets a reference to the given int64 and assigns it to the OrderStatus field.
func (o *GetOrderWebhook200ResponseDataOneOf) SetOrderStatus(v int64) {
	o.OrderStatus = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *GetOrderWebhook200ResponseDataOneOf) GetOrderId() int64 {
	if o == nil || IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderWebhook200ResponseDataOneOf) GetOrderIdOk() (*int64, bool) {
	if o == nil || IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *GetOrderWebhook200ResponseDataOneOf) HasOrderId() bool {
	if o != nil && !IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *GetOrderWebhook200ResponseDataOneOf) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetShopId returns the ShopId field value if set, zero value otherwise.
func (o *GetOrderWebhook200ResponseDataOneOf) GetShopId() int64 {
	if o == nil || IsNil(o.ShopId) {
		var ret int64
		return ret
	}
	return *o.ShopId
}

// GetShopIdOk returns a tuple with the ShopId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderWebhook200ResponseDataOneOf) GetShopIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ShopId) {
		return nil, false
	}
	return o.ShopId, true
}

// HasShopId returns a boolean if a field has been set.
func (o *GetOrderWebhook200ResponseDataOneOf) HasShopId() bool {
	if o != nil && !IsNil(o.ShopId) {
		return true
	}

	return false
}

// SetShopId gets a reference to the given int64 and assigns it to the ShopId field.
func (o *GetOrderWebhook200ResponseDataOneOf) SetShopId(v int64) {
	o.ShopId = &v
}

// GetWarehouseId returns the WarehouseId field value if set, zero value otherwise.
func (o *GetOrderWebhook200ResponseDataOneOf) GetWarehouseId() int64 {
	if o == nil || IsNil(o.WarehouseId) {
		var ret int64
		return ret
	}
	return *o.WarehouseId
}

// GetWarehouseIdOk returns a tuple with the WarehouseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderWebhook200ResponseDataOneOf) GetWarehouseIdOk() (*int64, bool) {
	if o == nil || IsNil(o.WarehouseId) {
		return nil, false
	}
	return o.WarehouseId, true
}

// HasWarehouseId returns a boolean if a field has been set.
func (o *GetOrderWebhook200ResponseDataOneOf) HasWarehouseId() bool {
	if o != nil && !IsNil(o.WarehouseId) {
		return true
	}

	return false
}

// SetWarehouseId gets a reference to the given int64 and assigns it to the WarehouseId field.
func (o *GetOrderWebhook200ResponseDataOneOf) SetWarehouseId(v int64) {
	o.WarehouseId = &v
}

// GetProductDetails returns the ProductDetails field value if set, zero value otherwise.
func (o *GetOrderWebhook200ResponseDataOneOf) GetProductDetails() []GetOrderWebhook200ResponseDataOneOfProductDetailsInner {
	if o == nil || IsNil(o.ProductDetails) {
		var ret []GetOrderWebhook200ResponseDataOneOfProductDetailsInner
		return ret
	}
	return o.ProductDetails
}

// GetProductDetailsOk returns a tuple with the ProductDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderWebhook200ResponseDataOneOf) GetProductDetailsOk() ([]GetOrderWebhook200ResponseDataOneOfProductDetailsInner, bool) {
	if o == nil || IsNil(o.ProductDetails) {
		return nil, false
	}
	return o.ProductDetails, true
}

// HasProductDetails returns a boolean if a field has been set.
func (o *GetOrderWebhook200ResponseDataOneOf) HasProductDetails() bool {
	if o != nil && !IsNil(o.ProductDetails) {
		return true
	}

	return false
}

// SetProductDetails gets a reference to the given []GetOrderWebhook200ResponseDataOneOfProductDetailsInner and assigns it to the ProductDetails field.
func (o *GetOrderWebhook200ResponseDataOneOf) SetProductDetails(v []GetOrderWebhook200ResponseDataOneOfProductDetailsInner) {
	o.ProductDetails = v
}

func (o GetOrderWebhook200ResponseDataOneOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrderWebhook200ResponseDataOneOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FsId) {
		toSerialize["fs_id"] = o.FsId
	}
	if !IsNil(o.OrderStatus) {
		toSerialize["order_status"] = o.OrderStatus
	}
	if !IsNil(o.OrderId) {
		toSerialize["order_id"] = o.OrderId
	}
	if !IsNil(o.ShopId) {
		toSerialize["shop_id"] = o.ShopId
	}
	if !IsNil(o.WarehouseId) {
		toSerialize["warehouse_id"] = o.WarehouseId
	}
	if !IsNil(o.ProductDetails) {
		toSerialize["product_details"] = o.ProductDetails
	}
	return toSerialize, nil
}

type NullableGetOrderWebhook200ResponseDataOneOf struct {
	value *GetOrderWebhook200ResponseDataOneOf
	isSet bool
}

func (v NullableGetOrderWebhook200ResponseDataOneOf) Get() *GetOrderWebhook200ResponseDataOneOf {
	return v.value
}

func (v *NullableGetOrderWebhook200ResponseDataOneOf) Set(val *GetOrderWebhook200ResponseDataOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrderWebhook200ResponseDataOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrderWebhook200ResponseDataOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrderWebhook200ResponseDataOneOf(val *GetOrderWebhook200ResponseDataOneOf) *NullableGetOrderWebhook200ResponseDataOneOf {
	return &NullableGetOrderWebhook200ResponseDataOneOf{value: val, isSet: true}
}

func (v NullableGetOrderWebhook200ResponseDataOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrderWebhook200ResponseDataOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


