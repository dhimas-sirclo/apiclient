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

// checks if the GetSingleOrder200ResponseDataPreorder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSingleOrder200ResponseDataPreorder{}

// GetSingleOrder200ResponseDataPreorder struct for GetSingleOrder200ResponseDataPreorder
type GetSingleOrder200ResponseDataPreorder struct {
	// Order Unique Identifier
	OrderId *int64 `json:"order_id,omitempty"`
	// Preorder Type
	PreorderType *int64 `json:"preorder_type,omitempty"`
	// Preorder Process Time in UNIX Format
	PreorderProcessTime *int64 `json:"preorder_process_time,omitempty"`
	// Preorder Process Start Time (format: 2017-07-20T17:50:58.061156Z) 
	PreorderProcessStart *string `json:"preorder_process_start,omitempty"`
	// Preorder Deadline Time (format: 0001-01-01T00:00:00Z) 
	PreorderDeadline *string `json:"preorder_deadline,omitempty"`
	// Shop Unique Identifier
	ShopId *int64 `json:"shop_id,omitempty"`
	// Customer User Unique Identifier
	CustomerId *int64 `json:"customer_id,omitempty"`
}

// NewGetSingleOrder200ResponseDataPreorder instantiates a new GetSingleOrder200ResponseDataPreorder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSingleOrder200ResponseDataPreorder() *GetSingleOrder200ResponseDataPreorder {
	this := GetSingleOrder200ResponseDataPreorder{}
	return &this
}

// NewGetSingleOrder200ResponseDataPreorderWithDefaults instantiates a new GetSingleOrder200ResponseDataPreorder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSingleOrder200ResponseDataPreorderWithDefaults() *GetSingleOrder200ResponseDataPreorder {
	this := GetSingleOrder200ResponseDataPreorder{}
	return &this
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataPreorder) GetOrderId() int64 {
	if o == nil || IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataPreorder) GetOrderIdOk() (*int64, bool) {
	if o == nil || IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataPreorder) HasOrderId() bool {
	if o != nil && !IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *GetSingleOrder200ResponseDataPreorder) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetPreorderType returns the PreorderType field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataPreorder) GetPreorderType() int64 {
	if o == nil || IsNil(o.PreorderType) {
		var ret int64
		return ret
	}
	return *o.PreorderType
}

// GetPreorderTypeOk returns a tuple with the PreorderType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataPreorder) GetPreorderTypeOk() (*int64, bool) {
	if o == nil || IsNil(o.PreorderType) {
		return nil, false
	}
	return o.PreorderType, true
}

// HasPreorderType returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataPreorder) HasPreorderType() bool {
	if o != nil && !IsNil(o.PreorderType) {
		return true
	}

	return false
}

// SetPreorderType gets a reference to the given int64 and assigns it to the PreorderType field.
func (o *GetSingleOrder200ResponseDataPreorder) SetPreorderType(v int64) {
	o.PreorderType = &v
}

// GetPreorderProcessTime returns the PreorderProcessTime field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataPreorder) GetPreorderProcessTime() int64 {
	if o == nil || IsNil(o.PreorderProcessTime) {
		var ret int64
		return ret
	}
	return *o.PreorderProcessTime
}

// GetPreorderProcessTimeOk returns a tuple with the PreorderProcessTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataPreorder) GetPreorderProcessTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.PreorderProcessTime) {
		return nil, false
	}
	return o.PreorderProcessTime, true
}

// HasPreorderProcessTime returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataPreorder) HasPreorderProcessTime() bool {
	if o != nil && !IsNil(o.PreorderProcessTime) {
		return true
	}

	return false
}

// SetPreorderProcessTime gets a reference to the given int64 and assigns it to the PreorderProcessTime field.
func (o *GetSingleOrder200ResponseDataPreorder) SetPreorderProcessTime(v int64) {
	o.PreorderProcessTime = &v
}

// GetPreorderProcessStart returns the PreorderProcessStart field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataPreorder) GetPreorderProcessStart() string {
	if o == nil || IsNil(o.PreorderProcessStart) {
		var ret string
		return ret
	}
	return *o.PreorderProcessStart
}

// GetPreorderProcessStartOk returns a tuple with the PreorderProcessStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataPreorder) GetPreorderProcessStartOk() (*string, bool) {
	if o == nil || IsNil(o.PreorderProcessStart) {
		return nil, false
	}
	return o.PreorderProcessStart, true
}

// HasPreorderProcessStart returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataPreorder) HasPreorderProcessStart() bool {
	if o != nil && !IsNil(o.PreorderProcessStart) {
		return true
	}

	return false
}

// SetPreorderProcessStart gets a reference to the given string and assigns it to the PreorderProcessStart field.
func (o *GetSingleOrder200ResponseDataPreorder) SetPreorderProcessStart(v string) {
	o.PreorderProcessStart = &v
}

// GetPreorderDeadline returns the PreorderDeadline field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataPreorder) GetPreorderDeadline() string {
	if o == nil || IsNil(o.PreorderDeadline) {
		var ret string
		return ret
	}
	return *o.PreorderDeadline
}

// GetPreorderDeadlineOk returns a tuple with the PreorderDeadline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataPreorder) GetPreorderDeadlineOk() (*string, bool) {
	if o == nil || IsNil(o.PreorderDeadline) {
		return nil, false
	}
	return o.PreorderDeadline, true
}

// HasPreorderDeadline returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataPreorder) HasPreorderDeadline() bool {
	if o != nil && !IsNil(o.PreorderDeadline) {
		return true
	}

	return false
}

// SetPreorderDeadline gets a reference to the given string and assigns it to the PreorderDeadline field.
func (o *GetSingleOrder200ResponseDataPreorder) SetPreorderDeadline(v string) {
	o.PreorderDeadline = &v
}

// GetShopId returns the ShopId field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataPreorder) GetShopId() int64 {
	if o == nil || IsNil(o.ShopId) {
		var ret int64
		return ret
	}
	return *o.ShopId
}

// GetShopIdOk returns a tuple with the ShopId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataPreorder) GetShopIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ShopId) {
		return nil, false
	}
	return o.ShopId, true
}

// HasShopId returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataPreorder) HasShopId() bool {
	if o != nil && !IsNil(o.ShopId) {
		return true
	}

	return false
}

// SetShopId gets a reference to the given int64 and assigns it to the ShopId field.
func (o *GetSingleOrder200ResponseDataPreorder) SetShopId(v int64) {
	o.ShopId = &v
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataPreorder) GetCustomerId() int64 {
	if o == nil || IsNil(o.CustomerId) {
		var ret int64
		return ret
	}
	return *o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataPreorder) GetCustomerIdOk() (*int64, bool) {
	if o == nil || IsNil(o.CustomerId) {
		return nil, false
	}
	return o.CustomerId, true
}

// HasCustomerId returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataPreorder) HasCustomerId() bool {
	if o != nil && !IsNil(o.CustomerId) {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given int64 and assigns it to the CustomerId field.
func (o *GetSingleOrder200ResponseDataPreorder) SetCustomerId(v int64) {
	o.CustomerId = &v
}

func (o GetSingleOrder200ResponseDataPreorder) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSingleOrder200ResponseDataPreorder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OrderId) {
		toSerialize["order_id"] = o.OrderId
	}
	if !IsNil(o.PreorderType) {
		toSerialize["preorder_type"] = o.PreorderType
	}
	if !IsNil(o.PreorderProcessTime) {
		toSerialize["preorder_process_time"] = o.PreorderProcessTime
	}
	if !IsNil(o.PreorderProcessStart) {
		toSerialize["preorder_process_start"] = o.PreorderProcessStart
	}
	if !IsNil(o.PreorderDeadline) {
		toSerialize["preorder_deadline"] = o.PreorderDeadline
	}
	if !IsNil(o.ShopId) {
		toSerialize["shop_id"] = o.ShopId
	}
	if !IsNil(o.CustomerId) {
		toSerialize["customer_id"] = o.CustomerId
	}
	return toSerialize, nil
}

type NullableGetSingleOrder200ResponseDataPreorder struct {
	value *GetSingleOrder200ResponseDataPreorder
	isSet bool
}

func (v NullableGetSingleOrder200ResponseDataPreorder) Get() *GetSingleOrder200ResponseDataPreorder {
	return v.value
}

func (v *NullableGetSingleOrder200ResponseDataPreorder) Set(val *GetSingleOrder200ResponseDataPreorder) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSingleOrder200ResponseDataPreorder) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSingleOrder200ResponseDataPreorder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSingleOrder200ResponseDataPreorder(val *GetSingleOrder200ResponseDataPreorder) *NullableGetSingleOrder200ResponseDataPreorder {
	return &NullableGetSingleOrder200ResponseDataPreorder{value: val, isSet: true}
}

func (v NullableGetSingleOrder200ResponseDataPreorder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSingleOrder200ResponseDataPreorder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


