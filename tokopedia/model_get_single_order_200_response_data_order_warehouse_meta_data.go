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

// checks if the GetSingleOrder200ResponseDataOrderWarehouseMetaData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSingleOrder200ResponseDataOrderWarehouseMetaData{}

// GetSingleOrder200ResponseDataOrderWarehouseMetaData struct for GetSingleOrder200ResponseDataOrderWarehouseMetaData
type GetSingleOrder200ResponseDataOrderWarehouseMetaData struct {
	// Warehouse Unique Identifier
	WarehouseId *int64 `json:"warehouse_id,omitempty"`
	// Partner Unique Identifier
	PartnerId *int64 `json:"partner_id,omitempty"`
	// Shop Unique Identifier
	ShopId *int64 `json:"shop_id,omitempty"`
	// Warehouse Name
	WarehouseName *string `json:"warehouse_name,omitempty"`
	// Warehouse Distric Unique Identifier
	DistrictId *int64 `json:"district_id,omitempty"`
	// Warehouse Distric Name
	DistrictName *string `json:"district_name,omitempty"`
	// Warehouse City Unique Identifier
	CityId *int64 `json:"city_id,omitempty"`
	// Warehouse City Name
	CityName *string `json:"city_name,omitempty"`
	// Warehouse Province Unique Identifier
	ProvinceId *int64 `json:"province_id,omitempty"`
	// Warehouse Province Name
	ProvinceName *string `json:"province_name,omitempty"`
	// Warehouse Status
	Status *int64 `json:"status,omitempty"`
	// Warehouse Postal Code
	PostalCode *string `json:"postal_code,omitempty"`
	// Is default warehouse?
	IsDefault *int64 `json:"is_default,omitempty"`
	// Geolocation Coordinate
	Latlon *string `json:"latlon,omitempty"`
	// Latitude Coordinate
	Latitude *string `json:"latitude,omitempty"`
	// Longitude Coordinate
	Longitude *string `json:"longitude,omitempty"`
	// Warehouse Email
	Email *string `json:"email,omitempty"`
	// Warehouse Detail Address
	AddressDetail *string `json:"address_detail,omitempty"`
	// Warehouse Country Name
	CountryName *string `json:"country_name,omitempty"`
	// Is fulfillment warehouse?
	IsFulfillment *bool `json:"is_fulfillment,omitempty"`
}

// NewGetSingleOrder200ResponseDataOrderWarehouseMetaData instantiates a new GetSingleOrder200ResponseDataOrderWarehouseMetaData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSingleOrder200ResponseDataOrderWarehouseMetaData() *GetSingleOrder200ResponseDataOrderWarehouseMetaData {
	this := GetSingleOrder200ResponseDataOrderWarehouseMetaData{}
	return &this
}

// NewGetSingleOrder200ResponseDataOrderWarehouseMetaDataWithDefaults instantiates a new GetSingleOrder200ResponseDataOrderWarehouseMetaData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSingleOrder200ResponseDataOrderWarehouseMetaDataWithDefaults() *GetSingleOrder200ResponseDataOrderWarehouseMetaData {
	this := GetSingleOrder200ResponseDataOrderWarehouseMetaData{}
	return &this
}

// GetWarehouseId returns the WarehouseId field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) GetWarehouseId() int64 {
	if o == nil || IsNil(o.WarehouseId) {
		var ret int64
		return ret
	}
	return *o.WarehouseId
}

// GetWarehouseIdOk returns a tuple with the WarehouseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) GetWarehouseIdOk() (*int64, bool) {
	if o == nil || IsNil(o.WarehouseId) {
		return nil, false
	}
	return o.WarehouseId, true
}

// HasWarehouseId returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) HasWarehouseId() bool {
	if o != nil && !IsNil(o.WarehouseId) {
		return true
	}

	return false
}

// SetWarehouseId gets a reference to the given int64 and assigns it to the WarehouseId field.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) SetWarehouseId(v int64) {
	o.WarehouseId = &v
}

// GetPartnerId returns the PartnerId field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) GetPartnerId() int64 {
	if o == nil || IsNil(o.PartnerId) {
		var ret int64
		return ret
	}
	return *o.PartnerId
}

// GetPartnerIdOk returns a tuple with the PartnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) GetPartnerIdOk() (*int64, bool) {
	if o == nil || IsNil(o.PartnerId) {
		return nil, false
	}
	return o.PartnerId, true
}

// HasPartnerId returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) HasPartnerId() bool {
	if o != nil && !IsNil(o.PartnerId) {
		return true
	}

	return false
}

// SetPartnerId gets a reference to the given int64 and assigns it to the PartnerId field.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) SetPartnerId(v int64) {
	o.PartnerId = &v
}

// GetShopId returns the ShopId field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) GetShopId() int64 {
	if o == nil || IsNil(o.ShopId) {
		var ret int64
		return ret
	}
	return *o.ShopId
}

// GetShopIdOk returns a tuple with the ShopId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) GetShopIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ShopId) {
		return nil, false
	}
	return o.ShopId, true
}

// HasShopId returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) HasShopId() bool {
	if o != nil && !IsNil(o.ShopId) {
		return true
	}

	return false
}

// SetShopId gets a reference to the given int64 and assigns it to the ShopId field.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) SetShopId(v int64) {
	o.ShopId = &v
}

// GetWarehouseName returns the WarehouseName field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) GetWarehouseName() string {
	if o == nil || IsNil(o.WarehouseName) {
		var ret string
		return ret
	}
	return *o.WarehouseName
}

// GetWarehouseNameOk returns a tuple with the WarehouseName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) GetWarehouseNameOk() (*string, bool) {
	if o == nil || IsNil(o.WarehouseName) {
		return nil, false
	}
	return o.WarehouseName, true
}

// HasWarehouseName returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) HasWarehouseName() bool {
	if o != nil && !IsNil(o.WarehouseName) {
		return true
	}

	return false
}

// SetWarehouseName gets a reference to the given string and assigns it to the WarehouseName field.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) SetWarehouseName(v string) {
	o.WarehouseName = &v
}

// GetDistrictId returns the DistrictId field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) GetDistrictId() int64 {
	if o == nil || IsNil(o.DistrictId) {
		var ret int64
		return ret
	}
	return *o.DistrictId
}

// GetDistrictIdOk returns a tuple with the DistrictId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) GetDistrictIdOk() (*int64, bool) {
	if o == nil || IsNil(o.DistrictId) {
		return nil, false
	}
	return o.DistrictId, true
}

// HasDistrictId returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) HasDistrictId() bool {
	if o != nil && !IsNil(o.DistrictId) {
		return true
	}

	return false
}

// SetDistrictId gets a reference to the given int64 and assigns it to the DistrictId field.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) SetDistrictId(v int64) {
	o.DistrictId = &v
}

// GetDistrictName returns the DistrictName field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) GetDistrictName() string {
	if o == nil || IsNil(o.DistrictName) {
		var ret string
		return ret
	}
	return *o.DistrictName
}

// GetDistrictNameOk returns a tuple with the DistrictName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) GetDistrictNameOk() (*string, bool) {
	if o == nil || IsNil(o.DistrictName) {
		return nil, false
	}
	return o.DistrictName, true
}

// HasDistrictName returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) HasDistrictName() bool {
	if o != nil && !IsNil(o.DistrictName) {
		return true
	}

	return false
}

// SetDistrictName gets a reference to the given string and assigns it to the DistrictName field.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) SetDistrictName(v string) {
	o.DistrictName = &v
}

// GetCityId returns the CityId field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) GetCityId() int64 {
	if o == nil || IsNil(o.CityId) {
		var ret int64
		return ret
	}
	return *o.CityId
}

// GetCityIdOk returns a tuple with the CityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) GetCityIdOk() (*int64, bool) {
	if o == nil || IsNil(o.CityId) {
		return nil, false
	}
	return o.CityId, true
}

// HasCityId returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) HasCityId() bool {
	if o != nil && !IsNil(o.CityId) {
		return true
	}

	return false
}

// SetCityId gets a reference to the given int64 and assigns it to the CityId field.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) SetCityId(v int64) {
	o.CityId = &v
}

// GetCityName returns the CityName field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) GetCityName() string {
	if o == nil || IsNil(o.CityName) {
		var ret string
		return ret
	}
	return *o.CityName
}

// GetCityNameOk returns a tuple with the CityName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) GetCityNameOk() (*string, bool) {
	if o == nil || IsNil(o.CityName) {
		return nil, false
	}
	return o.CityName, true
}

// HasCityName returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) HasCityName() bool {
	if o != nil && !IsNil(o.CityName) {
		return true
	}

	return false
}

// SetCityName gets a reference to the given string and assigns it to the CityName field.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) SetCityName(v string) {
	o.CityName = &v
}

// GetProvinceId returns the ProvinceId field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) GetProvinceId() int64 {
	if o == nil || IsNil(o.ProvinceId) {
		var ret int64
		return ret
	}
	return *o.ProvinceId
}

// GetProvinceIdOk returns a tuple with the ProvinceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) GetProvinceIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ProvinceId) {
		return nil, false
	}
	return o.ProvinceId, true
}

// HasProvinceId returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) HasProvinceId() bool {
	if o != nil && !IsNil(o.ProvinceId) {
		return true
	}

	return false
}

// SetProvinceId gets a reference to the given int64 and assigns it to the ProvinceId field.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) SetProvinceId(v int64) {
	o.ProvinceId = &v
}

// GetProvinceName returns the ProvinceName field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) GetProvinceName() string {
	if o == nil || IsNil(o.ProvinceName) {
		var ret string
		return ret
	}
	return *o.ProvinceName
}

// GetProvinceNameOk returns a tuple with the ProvinceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) GetProvinceNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProvinceName) {
		return nil, false
	}
	return o.ProvinceName, true
}

// HasProvinceName returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) HasProvinceName() bool {
	if o != nil && !IsNil(o.ProvinceName) {
		return true
	}

	return false
}

// SetProvinceName gets a reference to the given string and assigns it to the ProvinceName field.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) SetProvinceName(v string) {
	o.ProvinceName = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) GetStatus() int64 {
	if o == nil || IsNil(o.Status) {
		var ret int64
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) GetStatusOk() (*int64, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int64 and assigns it to the Status field.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) SetStatus(v int64) {
	o.Status = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) GetPostalCode() string {
	if o == nil || IsNil(o.PostalCode) {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) GetPostalCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PostalCode) {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) HasPostalCode() bool {
	if o != nil && !IsNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) GetIsDefault() int64 {
	if o == nil || IsNil(o.IsDefault) {
		var ret int64
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) GetIsDefaultOk() (*int64, bool) {
	if o == nil || IsNil(o.IsDefault) {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) HasIsDefault() bool {
	if o != nil && !IsNil(o.IsDefault) {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given int64 and assigns it to the IsDefault field.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) SetIsDefault(v int64) {
	o.IsDefault = &v
}

// GetLatlon returns the Latlon field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) GetLatlon() string {
	if o == nil || IsNil(o.Latlon) {
		var ret string
		return ret
	}
	return *o.Latlon
}

// GetLatlonOk returns a tuple with the Latlon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) GetLatlonOk() (*string, bool) {
	if o == nil || IsNil(o.Latlon) {
		return nil, false
	}
	return o.Latlon, true
}

// HasLatlon returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) HasLatlon() bool {
	if o != nil && !IsNil(o.Latlon) {
		return true
	}

	return false
}

// SetLatlon gets a reference to the given string and assigns it to the Latlon field.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) SetLatlon(v string) {
	o.Latlon = &v
}

// GetLatitude returns the Latitude field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) GetLatitude() string {
	if o == nil || IsNil(o.Latitude) {
		var ret string
		return ret
	}
	return *o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) GetLatitudeOk() (*string, bool) {
	if o == nil || IsNil(o.Latitude) {
		return nil, false
	}
	return o.Latitude, true
}

// HasLatitude returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) HasLatitude() bool {
	if o != nil && !IsNil(o.Latitude) {
		return true
	}

	return false
}

// SetLatitude gets a reference to the given string and assigns it to the Latitude field.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) SetLatitude(v string) {
	o.Latitude = &v
}

// GetLongitude returns the Longitude field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) GetLongitude() string {
	if o == nil || IsNil(o.Longitude) {
		var ret string
		return ret
	}
	return *o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) GetLongitudeOk() (*string, bool) {
	if o == nil || IsNil(o.Longitude) {
		return nil, false
	}
	return o.Longitude, true
}

// HasLongitude returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) HasLongitude() bool {
	if o != nil && !IsNil(o.Longitude) {
		return true
	}

	return false
}

// SetLongitude gets a reference to the given string and assigns it to the Longitude field.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) SetLongitude(v string) {
	o.Longitude = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) SetEmail(v string) {
	o.Email = &v
}

// GetAddressDetail returns the AddressDetail field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) GetAddressDetail() string {
	if o == nil || IsNil(o.AddressDetail) {
		var ret string
		return ret
	}
	return *o.AddressDetail
}

// GetAddressDetailOk returns a tuple with the AddressDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) GetAddressDetailOk() (*string, bool) {
	if o == nil || IsNil(o.AddressDetail) {
		return nil, false
	}
	return o.AddressDetail, true
}

// HasAddressDetail returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) HasAddressDetail() bool {
	if o != nil && !IsNil(o.AddressDetail) {
		return true
	}

	return false
}

// SetAddressDetail gets a reference to the given string and assigns it to the AddressDetail field.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) SetAddressDetail(v string) {
	o.AddressDetail = &v
}

// GetCountryName returns the CountryName field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) GetCountryName() string {
	if o == nil || IsNil(o.CountryName) {
		var ret string
		return ret
	}
	return *o.CountryName
}

// GetCountryNameOk returns a tuple with the CountryName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) GetCountryNameOk() (*string, bool) {
	if o == nil || IsNil(o.CountryName) {
		return nil, false
	}
	return o.CountryName, true
}

// HasCountryName returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) HasCountryName() bool {
	if o != nil && !IsNil(o.CountryName) {
		return true
	}

	return false
}

// SetCountryName gets a reference to the given string and assigns it to the CountryName field.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) SetCountryName(v string) {
	o.CountryName = &v
}

// GetIsFulfillment returns the IsFulfillment field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) GetIsFulfillment() bool {
	if o == nil || IsNil(o.IsFulfillment) {
		var ret bool
		return ret
	}
	return *o.IsFulfillment
}

// GetIsFulfillmentOk returns a tuple with the IsFulfillment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) GetIsFulfillmentOk() (*bool, bool) {
	if o == nil || IsNil(o.IsFulfillment) {
		return nil, false
	}
	return o.IsFulfillment, true
}

// HasIsFulfillment returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) HasIsFulfillment() bool {
	if o != nil && !IsNil(o.IsFulfillment) {
		return true
	}

	return false
}

// SetIsFulfillment gets a reference to the given bool and assigns it to the IsFulfillment field.
func (o *GetSingleOrder200ResponseDataOrderWarehouseMetaData) SetIsFulfillment(v bool) {
	o.IsFulfillment = &v
}

func (o GetSingleOrder200ResponseDataOrderWarehouseMetaData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSingleOrder200ResponseDataOrderWarehouseMetaData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.WarehouseId) {
		toSerialize["warehouse_id"] = o.WarehouseId
	}
	if !IsNil(o.PartnerId) {
		toSerialize["partner_id"] = o.PartnerId
	}
	if !IsNil(o.ShopId) {
		toSerialize["shop_id"] = o.ShopId
	}
	if !IsNil(o.WarehouseName) {
		toSerialize["warehouse_name"] = o.WarehouseName
	}
	if !IsNil(o.DistrictId) {
		toSerialize["district_id"] = o.DistrictId
	}
	if !IsNil(o.DistrictName) {
		toSerialize["district_name"] = o.DistrictName
	}
	if !IsNil(o.CityId) {
		toSerialize["city_id"] = o.CityId
	}
	if !IsNil(o.CityName) {
		toSerialize["city_name"] = o.CityName
	}
	if !IsNil(o.ProvinceId) {
		toSerialize["province_id"] = o.ProvinceId
	}
	if !IsNil(o.ProvinceName) {
		toSerialize["province_name"] = o.ProvinceName
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.PostalCode) {
		toSerialize["postal_code"] = o.PostalCode
	}
	if !IsNil(o.IsDefault) {
		toSerialize["is_default"] = o.IsDefault
	}
	if !IsNil(o.Latlon) {
		toSerialize["latlon"] = o.Latlon
	}
	if !IsNil(o.Latitude) {
		toSerialize["latitude"] = o.Latitude
	}
	if !IsNil(o.Longitude) {
		toSerialize["longitude"] = o.Longitude
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.AddressDetail) {
		toSerialize["address_detail"] = o.AddressDetail
	}
	if !IsNil(o.CountryName) {
		toSerialize["country_name"] = o.CountryName
	}
	if !IsNil(o.IsFulfillment) {
		toSerialize["is_fulfillment"] = o.IsFulfillment
	}
	return toSerialize, nil
}

type NullableGetSingleOrder200ResponseDataOrderWarehouseMetaData struct {
	value *GetSingleOrder200ResponseDataOrderWarehouseMetaData
	isSet bool
}

func (v NullableGetSingleOrder200ResponseDataOrderWarehouseMetaData) Get() *GetSingleOrder200ResponseDataOrderWarehouseMetaData {
	return v.value
}

func (v *NullableGetSingleOrder200ResponseDataOrderWarehouseMetaData) Set(val *GetSingleOrder200ResponseDataOrderWarehouseMetaData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSingleOrder200ResponseDataOrderWarehouseMetaData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSingleOrder200ResponseDataOrderWarehouseMetaData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSingleOrder200ResponseDataOrderWarehouseMetaData(val *GetSingleOrder200ResponseDataOrderWarehouseMetaData) *NullableGetSingleOrder200ResponseDataOrderWarehouseMetaData {
	return &NullableGetSingleOrder200ResponseDataOrderWarehouseMetaData{value: val, isSet: true}
}

func (v NullableGetSingleOrder200ResponseDataOrderWarehouseMetaData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSingleOrder200ResponseDataOrderWarehouseMetaData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


