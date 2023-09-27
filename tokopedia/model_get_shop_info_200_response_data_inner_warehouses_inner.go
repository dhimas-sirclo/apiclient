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

// checks if the GetShopInfo200ResponseDataInnerWarehousesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetShopInfo200ResponseDataInnerWarehousesInner{}

// GetShopInfo200ResponseDataInnerWarehousesInner struct for GetShopInfo200ResponseDataInnerWarehousesInner
type GetShopInfo200ResponseDataInnerWarehousesInner struct {
	WarehouseId *int64 `json:"warehouse_id,omitempty"`
	PartnerId *GetShopInfo200ResponseDataInnerWarehousesInnerPartnerId `json:"partner_id,omitempty"`
	ShopId *GetShopInfo200ResponseDataInnerWarehousesInnerPartnerId `json:"shop_id,omitempty"`
	WarehouseName *string `json:"warehouse_name,omitempty"`
	DistrictId *int64 `json:"district_id,omitempty"`
	DistrictName *string `json:"district_name,omitempty"`
	CityId *int64 `json:"city_id,omitempty"`
	CityName *string `json:"city_name,omitempty"`
	ProvinceId *int64 `json:"province_id,omitempty"`
	ProvinceName *string `json:"province_name,omitempty"`
	Status *int64 `json:"status,omitempty"`
	PostalCode *string `json:"postal_code,omitempty"`
	IsDefault *int64 `json:"is_default,omitempty"`
	Latlon *string `json:"latlon,omitempty"`
	Latitude *string `json:"latitude,omitempty"`
	Longitude *string `json:"longitude,omitempty"`
	Email *string `json:"email,omitempty"`
	AddressDetail *string `json:"address_detail,omitempty"`
	Phone *string `json:"phone,omitempty"`
	WarehoseType *string `json:"warehose_type,omitempty"`
}

// NewGetShopInfo200ResponseDataInnerWarehousesInner instantiates a new GetShopInfo200ResponseDataInnerWarehousesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetShopInfo200ResponseDataInnerWarehousesInner() *GetShopInfo200ResponseDataInnerWarehousesInner {
	this := GetShopInfo200ResponseDataInnerWarehousesInner{}
	return &this
}

// NewGetShopInfo200ResponseDataInnerWarehousesInnerWithDefaults instantiates a new GetShopInfo200ResponseDataInnerWarehousesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetShopInfo200ResponseDataInnerWarehousesInnerWithDefaults() *GetShopInfo200ResponseDataInnerWarehousesInner {
	this := GetShopInfo200ResponseDataInnerWarehousesInner{}
	return &this
}

// GetWarehouseId returns the WarehouseId field value if set, zero value otherwise.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) GetWarehouseId() int64 {
	if o == nil || IsNil(o.WarehouseId) {
		var ret int64
		return ret
	}
	return *o.WarehouseId
}

// GetWarehouseIdOk returns a tuple with the WarehouseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) GetWarehouseIdOk() (*int64, bool) {
	if o == nil || IsNil(o.WarehouseId) {
		return nil, false
	}
	return o.WarehouseId, true
}

// HasWarehouseId returns a boolean if a field has been set.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) HasWarehouseId() bool {
	if o != nil && !IsNil(o.WarehouseId) {
		return true
	}

	return false
}

// SetWarehouseId gets a reference to the given int64 and assigns it to the WarehouseId field.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) SetWarehouseId(v int64) {
	o.WarehouseId = &v
}

// GetPartnerId returns the PartnerId field value if set, zero value otherwise.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) GetPartnerId() GetShopInfo200ResponseDataInnerWarehousesInnerPartnerId {
	if o == nil || IsNil(o.PartnerId) {
		var ret GetShopInfo200ResponseDataInnerWarehousesInnerPartnerId
		return ret
	}
	return *o.PartnerId
}

// GetPartnerIdOk returns a tuple with the PartnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) GetPartnerIdOk() (*GetShopInfo200ResponseDataInnerWarehousesInnerPartnerId, bool) {
	if o == nil || IsNil(o.PartnerId) {
		return nil, false
	}
	return o.PartnerId, true
}

// HasPartnerId returns a boolean if a field has been set.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) HasPartnerId() bool {
	if o != nil && !IsNil(o.PartnerId) {
		return true
	}

	return false
}

// SetPartnerId gets a reference to the given GetShopInfo200ResponseDataInnerWarehousesInnerPartnerId and assigns it to the PartnerId field.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) SetPartnerId(v GetShopInfo200ResponseDataInnerWarehousesInnerPartnerId) {
	o.PartnerId = &v
}

// GetShopId returns the ShopId field value if set, zero value otherwise.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) GetShopId() GetShopInfo200ResponseDataInnerWarehousesInnerPartnerId {
	if o == nil || IsNil(o.ShopId) {
		var ret GetShopInfo200ResponseDataInnerWarehousesInnerPartnerId
		return ret
	}
	return *o.ShopId
}

// GetShopIdOk returns a tuple with the ShopId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) GetShopIdOk() (*GetShopInfo200ResponseDataInnerWarehousesInnerPartnerId, bool) {
	if o == nil || IsNil(o.ShopId) {
		return nil, false
	}
	return o.ShopId, true
}

// HasShopId returns a boolean if a field has been set.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) HasShopId() bool {
	if o != nil && !IsNil(o.ShopId) {
		return true
	}

	return false
}

// SetShopId gets a reference to the given GetShopInfo200ResponseDataInnerWarehousesInnerPartnerId and assigns it to the ShopId field.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) SetShopId(v GetShopInfo200ResponseDataInnerWarehousesInnerPartnerId) {
	o.ShopId = &v
}

// GetWarehouseName returns the WarehouseName field value if set, zero value otherwise.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) GetWarehouseName() string {
	if o == nil || IsNil(o.WarehouseName) {
		var ret string
		return ret
	}
	return *o.WarehouseName
}

// GetWarehouseNameOk returns a tuple with the WarehouseName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) GetWarehouseNameOk() (*string, bool) {
	if o == nil || IsNil(o.WarehouseName) {
		return nil, false
	}
	return o.WarehouseName, true
}

// HasWarehouseName returns a boolean if a field has been set.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) HasWarehouseName() bool {
	if o != nil && !IsNil(o.WarehouseName) {
		return true
	}

	return false
}

// SetWarehouseName gets a reference to the given string and assigns it to the WarehouseName field.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) SetWarehouseName(v string) {
	o.WarehouseName = &v
}

// GetDistrictId returns the DistrictId field value if set, zero value otherwise.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) GetDistrictId() int64 {
	if o == nil || IsNil(o.DistrictId) {
		var ret int64
		return ret
	}
	return *o.DistrictId
}

// GetDistrictIdOk returns a tuple with the DistrictId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) GetDistrictIdOk() (*int64, bool) {
	if o == nil || IsNil(o.DistrictId) {
		return nil, false
	}
	return o.DistrictId, true
}

// HasDistrictId returns a boolean if a field has been set.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) HasDistrictId() bool {
	if o != nil && !IsNil(o.DistrictId) {
		return true
	}

	return false
}

// SetDistrictId gets a reference to the given int64 and assigns it to the DistrictId field.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) SetDistrictId(v int64) {
	o.DistrictId = &v
}

// GetDistrictName returns the DistrictName field value if set, zero value otherwise.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) GetDistrictName() string {
	if o == nil || IsNil(o.DistrictName) {
		var ret string
		return ret
	}
	return *o.DistrictName
}

// GetDistrictNameOk returns a tuple with the DistrictName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) GetDistrictNameOk() (*string, bool) {
	if o == nil || IsNil(o.DistrictName) {
		return nil, false
	}
	return o.DistrictName, true
}

// HasDistrictName returns a boolean if a field has been set.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) HasDistrictName() bool {
	if o != nil && !IsNil(o.DistrictName) {
		return true
	}

	return false
}

// SetDistrictName gets a reference to the given string and assigns it to the DistrictName field.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) SetDistrictName(v string) {
	o.DistrictName = &v
}

// GetCityId returns the CityId field value if set, zero value otherwise.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) GetCityId() int64 {
	if o == nil || IsNil(o.CityId) {
		var ret int64
		return ret
	}
	return *o.CityId
}

// GetCityIdOk returns a tuple with the CityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) GetCityIdOk() (*int64, bool) {
	if o == nil || IsNil(o.CityId) {
		return nil, false
	}
	return o.CityId, true
}

// HasCityId returns a boolean if a field has been set.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) HasCityId() bool {
	if o != nil && !IsNil(o.CityId) {
		return true
	}

	return false
}

// SetCityId gets a reference to the given int64 and assigns it to the CityId field.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) SetCityId(v int64) {
	o.CityId = &v
}

// GetCityName returns the CityName field value if set, zero value otherwise.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) GetCityName() string {
	if o == nil || IsNil(o.CityName) {
		var ret string
		return ret
	}
	return *o.CityName
}

// GetCityNameOk returns a tuple with the CityName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) GetCityNameOk() (*string, bool) {
	if o == nil || IsNil(o.CityName) {
		return nil, false
	}
	return o.CityName, true
}

// HasCityName returns a boolean if a field has been set.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) HasCityName() bool {
	if o != nil && !IsNil(o.CityName) {
		return true
	}

	return false
}

// SetCityName gets a reference to the given string and assigns it to the CityName field.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) SetCityName(v string) {
	o.CityName = &v
}

// GetProvinceId returns the ProvinceId field value if set, zero value otherwise.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) GetProvinceId() int64 {
	if o == nil || IsNil(o.ProvinceId) {
		var ret int64
		return ret
	}
	return *o.ProvinceId
}

// GetProvinceIdOk returns a tuple with the ProvinceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) GetProvinceIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ProvinceId) {
		return nil, false
	}
	return o.ProvinceId, true
}

// HasProvinceId returns a boolean if a field has been set.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) HasProvinceId() bool {
	if o != nil && !IsNil(o.ProvinceId) {
		return true
	}

	return false
}

// SetProvinceId gets a reference to the given int64 and assigns it to the ProvinceId field.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) SetProvinceId(v int64) {
	o.ProvinceId = &v
}

// GetProvinceName returns the ProvinceName field value if set, zero value otherwise.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) GetProvinceName() string {
	if o == nil || IsNil(o.ProvinceName) {
		var ret string
		return ret
	}
	return *o.ProvinceName
}

// GetProvinceNameOk returns a tuple with the ProvinceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) GetProvinceNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProvinceName) {
		return nil, false
	}
	return o.ProvinceName, true
}

// HasProvinceName returns a boolean if a field has been set.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) HasProvinceName() bool {
	if o != nil && !IsNil(o.ProvinceName) {
		return true
	}

	return false
}

// SetProvinceName gets a reference to the given string and assigns it to the ProvinceName field.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) SetProvinceName(v string) {
	o.ProvinceName = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) GetStatus() int64 {
	if o == nil || IsNil(o.Status) {
		var ret int64
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) GetStatusOk() (*int64, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int64 and assigns it to the Status field.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) SetStatus(v int64) {
	o.Status = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) GetPostalCode() string {
	if o == nil || IsNil(o.PostalCode) {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) GetPostalCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PostalCode) {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) HasPostalCode() bool {
	if o != nil && !IsNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) GetIsDefault() int64 {
	if o == nil || IsNil(o.IsDefault) {
		var ret int64
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) GetIsDefaultOk() (*int64, bool) {
	if o == nil || IsNil(o.IsDefault) {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) HasIsDefault() bool {
	if o != nil && !IsNil(o.IsDefault) {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given int64 and assigns it to the IsDefault field.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) SetIsDefault(v int64) {
	o.IsDefault = &v
}

// GetLatlon returns the Latlon field value if set, zero value otherwise.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) GetLatlon() string {
	if o == nil || IsNil(o.Latlon) {
		var ret string
		return ret
	}
	return *o.Latlon
}

// GetLatlonOk returns a tuple with the Latlon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) GetLatlonOk() (*string, bool) {
	if o == nil || IsNil(o.Latlon) {
		return nil, false
	}
	return o.Latlon, true
}

// HasLatlon returns a boolean if a field has been set.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) HasLatlon() bool {
	if o != nil && !IsNil(o.Latlon) {
		return true
	}

	return false
}

// SetLatlon gets a reference to the given string and assigns it to the Latlon field.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) SetLatlon(v string) {
	o.Latlon = &v
}

// GetLatitude returns the Latitude field value if set, zero value otherwise.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) GetLatitude() string {
	if o == nil || IsNil(o.Latitude) {
		var ret string
		return ret
	}
	return *o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) GetLatitudeOk() (*string, bool) {
	if o == nil || IsNil(o.Latitude) {
		return nil, false
	}
	return o.Latitude, true
}

// HasLatitude returns a boolean if a field has been set.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) HasLatitude() bool {
	if o != nil && !IsNil(o.Latitude) {
		return true
	}

	return false
}

// SetLatitude gets a reference to the given string and assigns it to the Latitude field.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) SetLatitude(v string) {
	o.Latitude = &v
}

// GetLongitude returns the Longitude field value if set, zero value otherwise.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) GetLongitude() string {
	if o == nil || IsNil(o.Longitude) {
		var ret string
		return ret
	}
	return *o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) GetLongitudeOk() (*string, bool) {
	if o == nil || IsNil(o.Longitude) {
		return nil, false
	}
	return o.Longitude, true
}

// HasLongitude returns a boolean if a field has been set.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) HasLongitude() bool {
	if o != nil && !IsNil(o.Longitude) {
		return true
	}

	return false
}

// SetLongitude gets a reference to the given string and assigns it to the Longitude field.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) SetLongitude(v string) {
	o.Longitude = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) SetEmail(v string) {
	o.Email = &v
}

// GetAddressDetail returns the AddressDetail field value if set, zero value otherwise.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) GetAddressDetail() string {
	if o == nil || IsNil(o.AddressDetail) {
		var ret string
		return ret
	}
	return *o.AddressDetail
}

// GetAddressDetailOk returns a tuple with the AddressDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) GetAddressDetailOk() (*string, bool) {
	if o == nil || IsNil(o.AddressDetail) {
		return nil, false
	}
	return o.AddressDetail, true
}

// HasAddressDetail returns a boolean if a field has been set.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) HasAddressDetail() bool {
	if o != nil && !IsNil(o.AddressDetail) {
		return true
	}

	return false
}

// SetAddressDetail gets a reference to the given string and assigns it to the AddressDetail field.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) SetAddressDetail(v string) {
	o.AddressDetail = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) SetPhone(v string) {
	o.Phone = &v
}

// GetWarehoseType returns the WarehoseType field value if set, zero value otherwise.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) GetWarehoseType() string {
	if o == nil || IsNil(o.WarehoseType) {
		var ret string
		return ret
	}
	return *o.WarehoseType
}

// GetWarehoseTypeOk returns a tuple with the WarehoseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) GetWarehoseTypeOk() (*string, bool) {
	if o == nil || IsNil(o.WarehoseType) {
		return nil, false
	}
	return o.WarehoseType, true
}

// HasWarehoseType returns a boolean if a field has been set.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) HasWarehoseType() bool {
	if o != nil && !IsNil(o.WarehoseType) {
		return true
	}

	return false
}

// SetWarehoseType gets a reference to the given string and assigns it to the WarehoseType field.
func (o *GetShopInfo200ResponseDataInnerWarehousesInner) SetWarehoseType(v string) {
	o.WarehoseType = &v
}

func (o GetShopInfo200ResponseDataInnerWarehousesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetShopInfo200ResponseDataInnerWarehousesInner) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.WarehoseType) {
		toSerialize["warehose_type"] = o.WarehoseType
	}
	return toSerialize, nil
}

type NullableGetShopInfo200ResponseDataInnerWarehousesInner struct {
	value *GetShopInfo200ResponseDataInnerWarehousesInner
	isSet bool
}

func (v NullableGetShopInfo200ResponseDataInnerWarehousesInner) Get() *GetShopInfo200ResponseDataInnerWarehousesInner {
	return v.value
}

func (v *NullableGetShopInfo200ResponseDataInnerWarehousesInner) Set(val *GetShopInfo200ResponseDataInnerWarehousesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetShopInfo200ResponseDataInnerWarehousesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetShopInfo200ResponseDataInnerWarehousesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetShopInfo200ResponseDataInnerWarehousesInner(val *GetShopInfo200ResponseDataInnerWarehousesInner) *NullableGetShopInfo200ResponseDataInnerWarehousesInner {
	return &NullableGetShopInfo200ResponseDataInnerWarehousesInner{value: val, isSet: true}
}

func (v NullableGetShopInfo200ResponseDataInnerWarehousesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetShopInfo200ResponseDataInnerWarehousesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

