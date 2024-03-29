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

// checks if the GetAllOrders200ResponseDataInnerLogistics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAllOrders200ResponseDataInnerLogistics{}

// GetAllOrders200ResponseDataInnerLogistics struct for GetAllOrders200ResponseDataInnerLogistics
type GetAllOrders200ResponseDataInnerLogistics struct {
	// Shipping Unique Identifier
	ShippingId *int64 `json:"shipping_id,omitempty"`
	// Distric Unique Identifier
	DistrictId *int64 `json:"district_id,omitempty"`
	// City Unique Identifier
	CityId *int64 `json:"city_id,omitempty"`
	// Province Unique Identifier
	ProvinceId *int64 `json:"province_id,omitempty"`
	// Geolocation
	Geo *string `json:"geo,omitempty"`
	// Shipping Agency Name
	ShippingAgency *string `json:"shipping_agency,omitempty"`
	// Shipping Service Type
	ServiceType *string `json:"service_type,omitempty"`
}

// NewGetAllOrders200ResponseDataInnerLogistics instantiates a new GetAllOrders200ResponseDataInnerLogistics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAllOrders200ResponseDataInnerLogistics() *GetAllOrders200ResponseDataInnerLogistics {
	this := GetAllOrders200ResponseDataInnerLogistics{}
	return &this
}

// NewGetAllOrders200ResponseDataInnerLogisticsWithDefaults instantiates a new GetAllOrders200ResponseDataInnerLogistics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAllOrders200ResponseDataInnerLogisticsWithDefaults() *GetAllOrders200ResponseDataInnerLogistics {
	this := GetAllOrders200ResponseDataInnerLogistics{}
	return &this
}

// GetShippingId returns the ShippingId field value if set, zero value otherwise.
func (o *GetAllOrders200ResponseDataInnerLogistics) GetShippingId() int64 {
	if o == nil || IsNil(o.ShippingId) {
		var ret int64
		return ret
	}
	return *o.ShippingId
}

// GetShippingIdOk returns a tuple with the ShippingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllOrders200ResponseDataInnerLogistics) GetShippingIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ShippingId) {
		return nil, false
	}
	return o.ShippingId, true
}

// HasShippingId returns a boolean if a field has been set.
func (o *GetAllOrders200ResponseDataInnerLogistics) HasShippingId() bool {
	if o != nil && !IsNil(o.ShippingId) {
		return true
	}

	return false
}

// SetShippingId gets a reference to the given int64 and assigns it to the ShippingId field.
func (o *GetAllOrders200ResponseDataInnerLogistics) SetShippingId(v int64) {
	o.ShippingId = &v
}

// GetDistrictId returns the DistrictId field value if set, zero value otherwise.
func (o *GetAllOrders200ResponseDataInnerLogistics) GetDistrictId() int64 {
	if o == nil || IsNil(o.DistrictId) {
		var ret int64
		return ret
	}
	return *o.DistrictId
}

// GetDistrictIdOk returns a tuple with the DistrictId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllOrders200ResponseDataInnerLogistics) GetDistrictIdOk() (*int64, bool) {
	if o == nil || IsNil(o.DistrictId) {
		return nil, false
	}
	return o.DistrictId, true
}

// HasDistrictId returns a boolean if a field has been set.
func (o *GetAllOrders200ResponseDataInnerLogistics) HasDistrictId() bool {
	if o != nil && !IsNil(o.DistrictId) {
		return true
	}

	return false
}

// SetDistrictId gets a reference to the given int64 and assigns it to the DistrictId field.
func (o *GetAllOrders200ResponseDataInnerLogistics) SetDistrictId(v int64) {
	o.DistrictId = &v
}

// GetCityId returns the CityId field value if set, zero value otherwise.
func (o *GetAllOrders200ResponseDataInnerLogistics) GetCityId() int64 {
	if o == nil || IsNil(o.CityId) {
		var ret int64
		return ret
	}
	return *o.CityId
}

// GetCityIdOk returns a tuple with the CityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllOrders200ResponseDataInnerLogistics) GetCityIdOk() (*int64, bool) {
	if o == nil || IsNil(o.CityId) {
		return nil, false
	}
	return o.CityId, true
}

// HasCityId returns a boolean if a field has been set.
func (o *GetAllOrders200ResponseDataInnerLogistics) HasCityId() bool {
	if o != nil && !IsNil(o.CityId) {
		return true
	}

	return false
}

// SetCityId gets a reference to the given int64 and assigns it to the CityId field.
func (o *GetAllOrders200ResponseDataInnerLogistics) SetCityId(v int64) {
	o.CityId = &v
}

// GetProvinceId returns the ProvinceId field value if set, zero value otherwise.
func (o *GetAllOrders200ResponseDataInnerLogistics) GetProvinceId() int64 {
	if o == nil || IsNil(o.ProvinceId) {
		var ret int64
		return ret
	}
	return *o.ProvinceId
}

// GetProvinceIdOk returns a tuple with the ProvinceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllOrders200ResponseDataInnerLogistics) GetProvinceIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ProvinceId) {
		return nil, false
	}
	return o.ProvinceId, true
}

// HasProvinceId returns a boolean if a field has been set.
func (o *GetAllOrders200ResponseDataInnerLogistics) HasProvinceId() bool {
	if o != nil && !IsNil(o.ProvinceId) {
		return true
	}

	return false
}

// SetProvinceId gets a reference to the given int64 and assigns it to the ProvinceId field.
func (o *GetAllOrders200ResponseDataInnerLogistics) SetProvinceId(v int64) {
	o.ProvinceId = &v
}

// GetGeo returns the Geo field value if set, zero value otherwise.
func (o *GetAllOrders200ResponseDataInnerLogistics) GetGeo() string {
	if o == nil || IsNil(o.Geo) {
		var ret string
		return ret
	}
	return *o.Geo
}

// GetGeoOk returns a tuple with the Geo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllOrders200ResponseDataInnerLogistics) GetGeoOk() (*string, bool) {
	if o == nil || IsNil(o.Geo) {
		return nil, false
	}
	return o.Geo, true
}

// HasGeo returns a boolean if a field has been set.
func (o *GetAllOrders200ResponseDataInnerLogistics) HasGeo() bool {
	if o != nil && !IsNil(o.Geo) {
		return true
	}

	return false
}

// SetGeo gets a reference to the given string and assigns it to the Geo field.
func (o *GetAllOrders200ResponseDataInnerLogistics) SetGeo(v string) {
	o.Geo = &v
}

// GetShippingAgency returns the ShippingAgency field value if set, zero value otherwise.
func (o *GetAllOrders200ResponseDataInnerLogistics) GetShippingAgency() string {
	if o == nil || IsNil(o.ShippingAgency) {
		var ret string
		return ret
	}
	return *o.ShippingAgency
}

// GetShippingAgencyOk returns a tuple with the ShippingAgency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllOrders200ResponseDataInnerLogistics) GetShippingAgencyOk() (*string, bool) {
	if o == nil || IsNil(o.ShippingAgency) {
		return nil, false
	}
	return o.ShippingAgency, true
}

// HasShippingAgency returns a boolean if a field has been set.
func (o *GetAllOrders200ResponseDataInnerLogistics) HasShippingAgency() bool {
	if o != nil && !IsNil(o.ShippingAgency) {
		return true
	}

	return false
}

// SetShippingAgency gets a reference to the given string and assigns it to the ShippingAgency field.
func (o *GetAllOrders200ResponseDataInnerLogistics) SetShippingAgency(v string) {
	o.ShippingAgency = &v
}

// GetServiceType returns the ServiceType field value if set, zero value otherwise.
func (o *GetAllOrders200ResponseDataInnerLogistics) GetServiceType() string {
	if o == nil || IsNil(o.ServiceType) {
		var ret string
		return ret
	}
	return *o.ServiceType
}

// GetServiceTypeOk returns a tuple with the ServiceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllOrders200ResponseDataInnerLogistics) GetServiceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceType) {
		return nil, false
	}
	return o.ServiceType, true
}

// HasServiceType returns a boolean if a field has been set.
func (o *GetAllOrders200ResponseDataInnerLogistics) HasServiceType() bool {
	if o != nil && !IsNil(o.ServiceType) {
		return true
	}

	return false
}

// SetServiceType gets a reference to the given string and assigns it to the ServiceType field.
func (o *GetAllOrders200ResponseDataInnerLogistics) SetServiceType(v string) {
	o.ServiceType = &v
}

func (o GetAllOrders200ResponseDataInnerLogistics) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAllOrders200ResponseDataInnerLogistics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ShippingId) {
		toSerialize["shipping_id"] = o.ShippingId
	}
	if !IsNil(o.DistrictId) {
		toSerialize["district_id"] = o.DistrictId
	}
	if !IsNil(o.CityId) {
		toSerialize["city_id"] = o.CityId
	}
	if !IsNil(o.ProvinceId) {
		toSerialize["province_id"] = o.ProvinceId
	}
	if !IsNil(o.Geo) {
		toSerialize["geo"] = o.Geo
	}
	if !IsNil(o.ShippingAgency) {
		toSerialize["shipping_agency"] = o.ShippingAgency
	}
	if !IsNil(o.ServiceType) {
		toSerialize["service_type"] = o.ServiceType
	}
	return toSerialize, nil
}

type NullableGetAllOrders200ResponseDataInnerLogistics struct {
	value *GetAllOrders200ResponseDataInnerLogistics
	isSet bool
}

func (v NullableGetAllOrders200ResponseDataInnerLogistics) Get() *GetAllOrders200ResponseDataInnerLogistics {
	return v.value
}

func (v *NullableGetAllOrders200ResponseDataInnerLogistics) Set(val *GetAllOrders200ResponseDataInnerLogistics) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAllOrders200ResponseDataInnerLogistics) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAllOrders200ResponseDataInnerLogistics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAllOrders200ResponseDataInnerLogistics(val *GetAllOrders200ResponseDataInnerLogistics) *NullableGetAllOrders200ResponseDataInnerLogistics {
	return &NullableGetAllOrders200ResponseDataInnerLogistics{value: val, isSet: true}
}

func (v NullableGetAllOrders200ResponseDataInnerLogistics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAllOrders200ResponseDataInnerLogistics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


