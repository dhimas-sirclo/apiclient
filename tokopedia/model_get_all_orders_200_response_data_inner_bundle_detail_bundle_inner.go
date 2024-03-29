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

// checks if the GetAllOrders200ResponseDataInnerBundleDetailBundleInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAllOrders200ResponseDataInnerBundleDetailBundleInner{}

// GetAllOrders200ResponseDataInnerBundleDetailBundleInner struct for GetAllOrders200ResponseDataInnerBundleDetailBundleInner
type GetAllOrders200ResponseDataInnerBundleDetailBundleInner struct {
	// Bundle Unique Identifier
	BundleId *int64 `json:"bundle_id,omitempty"`
	// Bundle Variant Unique Identifier
	BundleVariantId *string `json:"bundle_variant_id,omitempty"`
	// Bundle Name
	BundleName *string `json:"bundle_name,omitempty"`
	// Bundle Variant Unique Identifier
	BundlePrice *float64 `json:"bundle_price,omitempty"`
	// Bundle Order Quantity
	BundleQuantity *int64 `json:"bundle_quantity,omitempty"`
	// Bundle Order Total Price
	BundleSubtotalPrice *float64 `json:"bundle_subtotal_price,omitempty"`
	OrderDetail []GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner `json:"order_detail,omitempty"`
}

// NewGetAllOrders200ResponseDataInnerBundleDetailBundleInner instantiates a new GetAllOrders200ResponseDataInnerBundleDetailBundleInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAllOrders200ResponseDataInnerBundleDetailBundleInner() *GetAllOrders200ResponseDataInnerBundleDetailBundleInner {
	this := GetAllOrders200ResponseDataInnerBundleDetailBundleInner{}
	return &this
}

// NewGetAllOrders200ResponseDataInnerBundleDetailBundleInnerWithDefaults instantiates a new GetAllOrders200ResponseDataInnerBundleDetailBundleInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAllOrders200ResponseDataInnerBundleDetailBundleInnerWithDefaults() *GetAllOrders200ResponseDataInnerBundleDetailBundleInner {
	this := GetAllOrders200ResponseDataInnerBundleDetailBundleInner{}
	return &this
}

// GetBundleId returns the BundleId field value if set, zero value otherwise.
func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInner) GetBundleId() int64 {
	if o == nil || IsNil(o.BundleId) {
		var ret int64
		return ret
	}
	return *o.BundleId
}

// GetBundleIdOk returns a tuple with the BundleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInner) GetBundleIdOk() (*int64, bool) {
	if o == nil || IsNil(o.BundleId) {
		return nil, false
	}
	return o.BundleId, true
}

// HasBundleId returns a boolean if a field has been set.
func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInner) HasBundleId() bool {
	if o != nil && !IsNil(o.BundleId) {
		return true
	}

	return false
}

// SetBundleId gets a reference to the given int64 and assigns it to the BundleId field.
func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInner) SetBundleId(v int64) {
	o.BundleId = &v
}

// GetBundleVariantId returns the BundleVariantId field value if set, zero value otherwise.
func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInner) GetBundleVariantId() string {
	if o == nil || IsNil(o.BundleVariantId) {
		var ret string
		return ret
	}
	return *o.BundleVariantId
}

// GetBundleVariantIdOk returns a tuple with the BundleVariantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInner) GetBundleVariantIdOk() (*string, bool) {
	if o == nil || IsNil(o.BundleVariantId) {
		return nil, false
	}
	return o.BundleVariantId, true
}

// HasBundleVariantId returns a boolean if a field has been set.
func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInner) HasBundleVariantId() bool {
	if o != nil && !IsNil(o.BundleVariantId) {
		return true
	}

	return false
}

// SetBundleVariantId gets a reference to the given string and assigns it to the BundleVariantId field.
func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInner) SetBundleVariantId(v string) {
	o.BundleVariantId = &v
}

// GetBundleName returns the BundleName field value if set, zero value otherwise.
func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInner) GetBundleName() string {
	if o == nil || IsNil(o.BundleName) {
		var ret string
		return ret
	}
	return *o.BundleName
}

// GetBundleNameOk returns a tuple with the BundleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInner) GetBundleNameOk() (*string, bool) {
	if o == nil || IsNil(o.BundleName) {
		return nil, false
	}
	return o.BundleName, true
}

// HasBundleName returns a boolean if a field has been set.
func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInner) HasBundleName() bool {
	if o != nil && !IsNil(o.BundleName) {
		return true
	}

	return false
}

// SetBundleName gets a reference to the given string and assigns it to the BundleName field.
func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInner) SetBundleName(v string) {
	o.BundleName = &v
}

// GetBundlePrice returns the BundlePrice field value if set, zero value otherwise.
func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInner) GetBundlePrice() float64 {
	if o == nil || IsNil(o.BundlePrice) {
		var ret float64
		return ret
	}
	return *o.BundlePrice
}

// GetBundlePriceOk returns a tuple with the BundlePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInner) GetBundlePriceOk() (*float64, bool) {
	if o == nil || IsNil(o.BundlePrice) {
		return nil, false
	}
	return o.BundlePrice, true
}

// HasBundlePrice returns a boolean if a field has been set.
func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInner) HasBundlePrice() bool {
	if o != nil && !IsNil(o.BundlePrice) {
		return true
	}

	return false
}

// SetBundlePrice gets a reference to the given float64 and assigns it to the BundlePrice field.
func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInner) SetBundlePrice(v float64) {
	o.BundlePrice = &v
}

// GetBundleQuantity returns the BundleQuantity field value if set, zero value otherwise.
func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInner) GetBundleQuantity() int64 {
	if o == nil || IsNil(o.BundleQuantity) {
		var ret int64
		return ret
	}
	return *o.BundleQuantity
}

// GetBundleQuantityOk returns a tuple with the BundleQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInner) GetBundleQuantityOk() (*int64, bool) {
	if o == nil || IsNil(o.BundleQuantity) {
		return nil, false
	}
	return o.BundleQuantity, true
}

// HasBundleQuantity returns a boolean if a field has been set.
func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInner) HasBundleQuantity() bool {
	if o != nil && !IsNil(o.BundleQuantity) {
		return true
	}

	return false
}

// SetBundleQuantity gets a reference to the given int64 and assigns it to the BundleQuantity field.
func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInner) SetBundleQuantity(v int64) {
	o.BundleQuantity = &v
}

// GetBundleSubtotalPrice returns the BundleSubtotalPrice field value if set, zero value otherwise.
func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInner) GetBundleSubtotalPrice() float64 {
	if o == nil || IsNil(o.BundleSubtotalPrice) {
		var ret float64
		return ret
	}
	return *o.BundleSubtotalPrice
}

// GetBundleSubtotalPriceOk returns a tuple with the BundleSubtotalPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInner) GetBundleSubtotalPriceOk() (*float64, bool) {
	if o == nil || IsNil(o.BundleSubtotalPrice) {
		return nil, false
	}
	return o.BundleSubtotalPrice, true
}

// HasBundleSubtotalPrice returns a boolean if a field has been set.
func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInner) HasBundleSubtotalPrice() bool {
	if o != nil && !IsNil(o.BundleSubtotalPrice) {
		return true
	}

	return false
}

// SetBundleSubtotalPrice gets a reference to the given float64 and assigns it to the BundleSubtotalPrice field.
func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInner) SetBundleSubtotalPrice(v float64) {
	o.BundleSubtotalPrice = &v
}

// GetOrderDetail returns the OrderDetail field value if set, zero value otherwise.
func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInner) GetOrderDetail() []GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner {
	if o == nil || IsNil(o.OrderDetail) {
		var ret []GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner
		return ret
	}
	return o.OrderDetail
}

// GetOrderDetailOk returns a tuple with the OrderDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInner) GetOrderDetailOk() ([]GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner, bool) {
	if o == nil || IsNil(o.OrderDetail) {
		return nil, false
	}
	return o.OrderDetail, true
}

// HasOrderDetail returns a boolean if a field has been set.
func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInner) HasOrderDetail() bool {
	if o != nil && !IsNil(o.OrderDetail) {
		return true
	}

	return false
}

// SetOrderDetail gets a reference to the given []GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner and assigns it to the OrderDetail field.
func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInner) SetOrderDetail(v []GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) {
	o.OrderDetail = v
}

func (o GetAllOrders200ResponseDataInnerBundleDetailBundleInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAllOrders200ResponseDataInnerBundleDetailBundleInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BundleId) {
		toSerialize["bundle_id"] = o.BundleId
	}
	if !IsNil(o.BundleVariantId) {
		toSerialize["bundle_variant_id"] = o.BundleVariantId
	}
	if !IsNil(o.BundleName) {
		toSerialize["bundle_name"] = o.BundleName
	}
	if !IsNil(o.BundlePrice) {
		toSerialize["bundle_price"] = o.BundlePrice
	}
	if !IsNil(o.BundleQuantity) {
		toSerialize["bundle_quantity"] = o.BundleQuantity
	}
	if !IsNil(o.BundleSubtotalPrice) {
		toSerialize["bundle_subtotal_price"] = o.BundleSubtotalPrice
	}
	if !IsNil(o.OrderDetail) {
		toSerialize["order_detail"] = o.OrderDetail
	}
	return toSerialize, nil
}

type NullableGetAllOrders200ResponseDataInnerBundleDetailBundleInner struct {
	value *GetAllOrders200ResponseDataInnerBundleDetailBundleInner
	isSet bool
}

func (v NullableGetAllOrders200ResponseDataInnerBundleDetailBundleInner) Get() *GetAllOrders200ResponseDataInnerBundleDetailBundleInner {
	return v.value
}

func (v *NullableGetAllOrders200ResponseDataInnerBundleDetailBundleInner) Set(val *GetAllOrders200ResponseDataInnerBundleDetailBundleInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAllOrders200ResponseDataInnerBundleDetailBundleInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAllOrders200ResponseDataInnerBundleDetailBundleInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAllOrders200ResponseDataInnerBundleDetailBundleInner(val *GetAllOrders200ResponseDataInnerBundleDetailBundleInner) *NullableGetAllOrders200ResponseDataInnerBundleDetailBundleInner {
	return &NullableGetAllOrders200ResponseDataInnerBundleDetailBundleInner{value: val, isSet: true}
}

func (v NullableGetAllOrders200ResponseDataInnerBundleDetailBundleInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAllOrders200ResponseDataInnerBundleDetailBundleInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


