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

// checks if the WebhookOrderNotificationRecipientAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookOrderNotificationRecipientAddress{}

// WebhookOrderNotificationRecipientAddress Recipient Address
type WebhookOrderNotificationRecipientAddress struct {
	// Complete Address
	AddressFull *string `json:"address_full,omitempty"`
	// District name
	District *string `json:"district,omitempty"`
	// City name
	City *string `json:"city,omitempty"`
	// Province name
	Province *string `json:"province,omitempty"`
	// Country name
	Country *string `json:"country,omitempty"`
	// Postal Code
	PostalCode *string `json:"postal_code,omitempty"`
	// District ID
	DistrictId *int64 `json:"district_id,omitempty"`
	// City ID
	CityId *int64 `json:"city_id,omitempty"`
	// Province ID
	ProvinceId *int64 `json:"province_id,omitempty"`
	// Longitude and Latitude position for instant delivery
	Geo *string `json:"geo,omitempty"`
}

// NewWebhookOrderNotificationRecipientAddress instantiates a new WebhookOrderNotificationRecipientAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookOrderNotificationRecipientAddress() *WebhookOrderNotificationRecipientAddress {
	this := WebhookOrderNotificationRecipientAddress{}
	return &this
}

// NewWebhookOrderNotificationRecipientAddressWithDefaults instantiates a new WebhookOrderNotificationRecipientAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookOrderNotificationRecipientAddressWithDefaults() *WebhookOrderNotificationRecipientAddress {
	this := WebhookOrderNotificationRecipientAddress{}
	return &this
}

// GetAddressFull returns the AddressFull field value if set, zero value otherwise.
func (o *WebhookOrderNotificationRecipientAddress) GetAddressFull() string {
	if o == nil || IsNil(o.AddressFull) {
		var ret string
		return ret
	}
	return *o.AddressFull
}

// GetAddressFullOk returns a tuple with the AddressFull field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookOrderNotificationRecipientAddress) GetAddressFullOk() (*string, bool) {
	if o == nil || IsNil(o.AddressFull) {
		return nil, false
	}
	return o.AddressFull, true
}

// HasAddressFull returns a boolean if a field has been set.
func (o *WebhookOrderNotificationRecipientAddress) HasAddressFull() bool {
	if o != nil && !IsNil(o.AddressFull) {
		return true
	}

	return false
}

// SetAddressFull gets a reference to the given string and assigns it to the AddressFull field.
func (o *WebhookOrderNotificationRecipientAddress) SetAddressFull(v string) {
	o.AddressFull = &v
}

// GetDistrict returns the District field value if set, zero value otherwise.
func (o *WebhookOrderNotificationRecipientAddress) GetDistrict() string {
	if o == nil || IsNil(o.District) {
		var ret string
		return ret
	}
	return *o.District
}

// GetDistrictOk returns a tuple with the District field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookOrderNotificationRecipientAddress) GetDistrictOk() (*string, bool) {
	if o == nil || IsNil(o.District) {
		return nil, false
	}
	return o.District, true
}

// HasDistrict returns a boolean if a field has been set.
func (o *WebhookOrderNotificationRecipientAddress) HasDistrict() bool {
	if o != nil && !IsNil(o.District) {
		return true
	}

	return false
}

// SetDistrict gets a reference to the given string and assigns it to the District field.
func (o *WebhookOrderNotificationRecipientAddress) SetDistrict(v string) {
	o.District = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *WebhookOrderNotificationRecipientAddress) GetCity() string {
	if o == nil || IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookOrderNotificationRecipientAddress) GetCityOk() (*string, bool) {
	if o == nil || IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *WebhookOrderNotificationRecipientAddress) HasCity() bool {
	if o != nil && !IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *WebhookOrderNotificationRecipientAddress) SetCity(v string) {
	o.City = &v
}

// GetProvince returns the Province field value if set, zero value otherwise.
func (o *WebhookOrderNotificationRecipientAddress) GetProvince() string {
	if o == nil || IsNil(o.Province) {
		var ret string
		return ret
	}
	return *o.Province
}

// GetProvinceOk returns a tuple with the Province field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookOrderNotificationRecipientAddress) GetProvinceOk() (*string, bool) {
	if o == nil || IsNil(o.Province) {
		return nil, false
	}
	return o.Province, true
}

// HasProvince returns a boolean if a field has been set.
func (o *WebhookOrderNotificationRecipientAddress) HasProvince() bool {
	if o != nil && !IsNil(o.Province) {
		return true
	}

	return false
}

// SetProvince gets a reference to the given string and assigns it to the Province field.
func (o *WebhookOrderNotificationRecipientAddress) SetProvince(v string) {
	o.Province = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *WebhookOrderNotificationRecipientAddress) GetCountry() string {
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookOrderNotificationRecipientAddress) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *WebhookOrderNotificationRecipientAddress) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *WebhookOrderNotificationRecipientAddress) SetCountry(v string) {
	o.Country = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *WebhookOrderNotificationRecipientAddress) GetPostalCode() string {
	if o == nil || IsNil(o.PostalCode) {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookOrderNotificationRecipientAddress) GetPostalCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PostalCode) {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *WebhookOrderNotificationRecipientAddress) HasPostalCode() bool {
	if o != nil && !IsNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *WebhookOrderNotificationRecipientAddress) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetDistrictId returns the DistrictId field value if set, zero value otherwise.
func (o *WebhookOrderNotificationRecipientAddress) GetDistrictId() int64 {
	if o == nil || IsNil(o.DistrictId) {
		var ret int64
		return ret
	}
	return *o.DistrictId
}

// GetDistrictIdOk returns a tuple with the DistrictId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookOrderNotificationRecipientAddress) GetDistrictIdOk() (*int64, bool) {
	if o == nil || IsNil(o.DistrictId) {
		return nil, false
	}
	return o.DistrictId, true
}

// HasDistrictId returns a boolean if a field has been set.
func (o *WebhookOrderNotificationRecipientAddress) HasDistrictId() bool {
	if o != nil && !IsNil(o.DistrictId) {
		return true
	}

	return false
}

// SetDistrictId gets a reference to the given int64 and assigns it to the DistrictId field.
func (o *WebhookOrderNotificationRecipientAddress) SetDistrictId(v int64) {
	o.DistrictId = &v
}

// GetCityId returns the CityId field value if set, zero value otherwise.
func (o *WebhookOrderNotificationRecipientAddress) GetCityId() int64 {
	if o == nil || IsNil(o.CityId) {
		var ret int64
		return ret
	}
	return *o.CityId
}

// GetCityIdOk returns a tuple with the CityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookOrderNotificationRecipientAddress) GetCityIdOk() (*int64, bool) {
	if o == nil || IsNil(o.CityId) {
		return nil, false
	}
	return o.CityId, true
}

// HasCityId returns a boolean if a field has been set.
func (o *WebhookOrderNotificationRecipientAddress) HasCityId() bool {
	if o != nil && !IsNil(o.CityId) {
		return true
	}

	return false
}

// SetCityId gets a reference to the given int64 and assigns it to the CityId field.
func (o *WebhookOrderNotificationRecipientAddress) SetCityId(v int64) {
	o.CityId = &v
}

// GetProvinceId returns the ProvinceId field value if set, zero value otherwise.
func (o *WebhookOrderNotificationRecipientAddress) GetProvinceId() int64 {
	if o == nil || IsNil(o.ProvinceId) {
		var ret int64
		return ret
	}
	return *o.ProvinceId
}

// GetProvinceIdOk returns a tuple with the ProvinceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookOrderNotificationRecipientAddress) GetProvinceIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ProvinceId) {
		return nil, false
	}
	return o.ProvinceId, true
}

// HasProvinceId returns a boolean if a field has been set.
func (o *WebhookOrderNotificationRecipientAddress) HasProvinceId() bool {
	if o != nil && !IsNil(o.ProvinceId) {
		return true
	}

	return false
}

// SetProvinceId gets a reference to the given int64 and assigns it to the ProvinceId field.
func (o *WebhookOrderNotificationRecipientAddress) SetProvinceId(v int64) {
	o.ProvinceId = &v
}

// GetGeo returns the Geo field value if set, zero value otherwise.
func (o *WebhookOrderNotificationRecipientAddress) GetGeo() string {
	if o == nil || IsNil(o.Geo) {
		var ret string
		return ret
	}
	return *o.Geo
}

// GetGeoOk returns a tuple with the Geo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookOrderNotificationRecipientAddress) GetGeoOk() (*string, bool) {
	if o == nil || IsNil(o.Geo) {
		return nil, false
	}
	return o.Geo, true
}

// HasGeo returns a boolean if a field has been set.
func (o *WebhookOrderNotificationRecipientAddress) HasGeo() bool {
	if o != nil && !IsNil(o.Geo) {
		return true
	}

	return false
}

// SetGeo gets a reference to the given string and assigns it to the Geo field.
func (o *WebhookOrderNotificationRecipientAddress) SetGeo(v string) {
	o.Geo = &v
}

func (o WebhookOrderNotificationRecipientAddress) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookOrderNotificationRecipientAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AddressFull) {
		toSerialize["address_full"] = o.AddressFull
	}
	if !IsNil(o.District) {
		toSerialize["district"] = o.District
	}
	if !IsNil(o.City) {
		toSerialize["city"] = o.City
	}
	if !IsNil(o.Province) {
		toSerialize["province"] = o.Province
	}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !IsNil(o.PostalCode) {
		toSerialize["postal_code"] = o.PostalCode
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
	return toSerialize, nil
}

type NullableWebhookOrderNotificationRecipientAddress struct {
	value *WebhookOrderNotificationRecipientAddress
	isSet bool
}

func (v NullableWebhookOrderNotificationRecipientAddress) Get() *WebhookOrderNotificationRecipientAddress {
	return v.value
}

func (v *NullableWebhookOrderNotificationRecipientAddress) Set(val *WebhookOrderNotificationRecipientAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookOrderNotificationRecipientAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookOrderNotificationRecipientAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookOrderNotificationRecipientAddress(val *WebhookOrderNotificationRecipientAddress) *NullableWebhookOrderNotificationRecipientAddress {
	return &NullableWebhookOrderNotificationRecipientAddress{value: val, isSet: true}
}

func (v NullableWebhookOrderNotificationRecipientAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookOrderNotificationRecipientAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


