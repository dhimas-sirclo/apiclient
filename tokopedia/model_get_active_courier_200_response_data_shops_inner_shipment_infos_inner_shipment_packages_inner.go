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

// checks if the GetActiveCourier200ResponseDataShopsInnerShipmentInfosInnerShipmentPackagesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetActiveCourier200ResponseDataShopsInnerShipmentInfosInnerShipmentPackagesInner{}

// GetActiveCourier200ResponseDataShopsInnerShipmentInfosInnerShipmentPackagesInner struct for GetActiveCourier200ResponseDataShopsInnerShipmentInfosInnerShipmentPackagesInner
type GetActiveCourier200ResponseDataShopsInnerShipmentInfosInnerShipmentPackagesInner struct {
	// Is Shipment Available? 1 = true, 0 = false
	IsAvailable *int64 `json:"IsAvailable,omitempty"`
	// Shipment Product Name
	ProductName *string `json:"ProductName,omitempty"`
	// Shipping Product Unique Identifier
	ShippingProductID *int64 `json:"ShippingProductID,omitempty"`
}

// NewGetActiveCourier200ResponseDataShopsInnerShipmentInfosInnerShipmentPackagesInner instantiates a new GetActiveCourier200ResponseDataShopsInnerShipmentInfosInnerShipmentPackagesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetActiveCourier200ResponseDataShopsInnerShipmentInfosInnerShipmentPackagesInner() *GetActiveCourier200ResponseDataShopsInnerShipmentInfosInnerShipmentPackagesInner {
	this := GetActiveCourier200ResponseDataShopsInnerShipmentInfosInnerShipmentPackagesInner{}
	return &this
}

// NewGetActiveCourier200ResponseDataShopsInnerShipmentInfosInnerShipmentPackagesInnerWithDefaults instantiates a new GetActiveCourier200ResponseDataShopsInnerShipmentInfosInnerShipmentPackagesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetActiveCourier200ResponseDataShopsInnerShipmentInfosInnerShipmentPackagesInnerWithDefaults() *GetActiveCourier200ResponseDataShopsInnerShipmentInfosInnerShipmentPackagesInner {
	this := GetActiveCourier200ResponseDataShopsInnerShipmentInfosInnerShipmentPackagesInner{}
	return &this
}

// GetIsAvailable returns the IsAvailable field value if set, zero value otherwise.
func (o *GetActiveCourier200ResponseDataShopsInnerShipmentInfosInnerShipmentPackagesInner) GetIsAvailable() int64 {
	if o == nil || IsNil(o.IsAvailable) {
		var ret int64
		return ret
	}
	return *o.IsAvailable
}

// GetIsAvailableOk returns a tuple with the IsAvailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetActiveCourier200ResponseDataShopsInnerShipmentInfosInnerShipmentPackagesInner) GetIsAvailableOk() (*int64, bool) {
	if o == nil || IsNil(o.IsAvailable) {
		return nil, false
	}
	return o.IsAvailable, true
}

// HasIsAvailable returns a boolean if a field has been set.
func (o *GetActiveCourier200ResponseDataShopsInnerShipmentInfosInnerShipmentPackagesInner) HasIsAvailable() bool {
	if o != nil && !IsNil(o.IsAvailable) {
		return true
	}

	return false
}

// SetIsAvailable gets a reference to the given int64 and assigns it to the IsAvailable field.
func (o *GetActiveCourier200ResponseDataShopsInnerShipmentInfosInnerShipmentPackagesInner) SetIsAvailable(v int64) {
	o.IsAvailable = &v
}

// GetProductName returns the ProductName field value if set, zero value otherwise.
func (o *GetActiveCourier200ResponseDataShopsInnerShipmentInfosInnerShipmentPackagesInner) GetProductName() string {
	if o == nil || IsNil(o.ProductName) {
		var ret string
		return ret
	}
	return *o.ProductName
}

// GetProductNameOk returns a tuple with the ProductName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetActiveCourier200ResponseDataShopsInnerShipmentInfosInnerShipmentPackagesInner) GetProductNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProductName) {
		return nil, false
	}
	return o.ProductName, true
}

// HasProductName returns a boolean if a field has been set.
func (o *GetActiveCourier200ResponseDataShopsInnerShipmentInfosInnerShipmentPackagesInner) HasProductName() bool {
	if o != nil && !IsNil(o.ProductName) {
		return true
	}

	return false
}

// SetProductName gets a reference to the given string and assigns it to the ProductName field.
func (o *GetActiveCourier200ResponseDataShopsInnerShipmentInfosInnerShipmentPackagesInner) SetProductName(v string) {
	o.ProductName = &v
}

// GetShippingProductID returns the ShippingProductID field value if set, zero value otherwise.
func (o *GetActiveCourier200ResponseDataShopsInnerShipmentInfosInnerShipmentPackagesInner) GetShippingProductID() int64 {
	if o == nil || IsNil(o.ShippingProductID) {
		var ret int64
		return ret
	}
	return *o.ShippingProductID
}

// GetShippingProductIDOk returns a tuple with the ShippingProductID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetActiveCourier200ResponseDataShopsInnerShipmentInfosInnerShipmentPackagesInner) GetShippingProductIDOk() (*int64, bool) {
	if o == nil || IsNil(o.ShippingProductID) {
		return nil, false
	}
	return o.ShippingProductID, true
}

// HasShippingProductID returns a boolean if a field has been set.
func (o *GetActiveCourier200ResponseDataShopsInnerShipmentInfosInnerShipmentPackagesInner) HasShippingProductID() bool {
	if o != nil && !IsNil(o.ShippingProductID) {
		return true
	}

	return false
}

// SetShippingProductID gets a reference to the given int64 and assigns it to the ShippingProductID field.
func (o *GetActiveCourier200ResponseDataShopsInnerShipmentInfosInnerShipmentPackagesInner) SetShippingProductID(v int64) {
	o.ShippingProductID = &v
}

func (o GetActiveCourier200ResponseDataShopsInnerShipmentInfosInnerShipmentPackagesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetActiveCourier200ResponseDataShopsInnerShipmentInfosInnerShipmentPackagesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsAvailable) {
		toSerialize["IsAvailable"] = o.IsAvailable
	}
	if !IsNil(o.ProductName) {
		toSerialize["ProductName"] = o.ProductName
	}
	if !IsNil(o.ShippingProductID) {
		toSerialize["ShippingProductID"] = o.ShippingProductID
	}
	return toSerialize, nil
}

type NullableGetActiveCourier200ResponseDataShopsInnerShipmentInfosInnerShipmentPackagesInner struct {
	value *GetActiveCourier200ResponseDataShopsInnerShipmentInfosInnerShipmentPackagesInner
	isSet bool
}

func (v NullableGetActiveCourier200ResponseDataShopsInnerShipmentInfosInnerShipmentPackagesInner) Get() *GetActiveCourier200ResponseDataShopsInnerShipmentInfosInnerShipmentPackagesInner {
	return v.value
}

func (v *NullableGetActiveCourier200ResponseDataShopsInnerShipmentInfosInnerShipmentPackagesInner) Set(val *GetActiveCourier200ResponseDataShopsInnerShipmentInfosInnerShipmentPackagesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetActiveCourier200ResponseDataShopsInnerShipmentInfosInnerShipmentPackagesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetActiveCourier200ResponseDataShopsInnerShipmentInfosInnerShipmentPackagesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetActiveCourier200ResponseDataShopsInnerShipmentInfosInnerShipmentPackagesInner(val *GetActiveCourier200ResponseDataShopsInnerShipmentInfosInnerShipmentPackagesInner) *NullableGetActiveCourier200ResponseDataShopsInnerShipmentInfosInnerShipmentPackagesInner {
	return &NullableGetActiveCourier200ResponseDataShopsInnerShipmentInfosInnerShipmentPackagesInner{value: val, isSet: true}
}

func (v NullableGetActiveCourier200ResponseDataShopsInnerShipmentInfosInnerShipmentPackagesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetActiveCourier200ResponseDataShopsInnerShipmentInfosInnerShipmentPackagesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


