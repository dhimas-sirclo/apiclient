# WebhookOrderNotificationRecipientAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressFull** | Pointer to **string** | Complete Address | [optional] 
**District** | Pointer to **string** | District name | [optional] 
**City** | Pointer to **string** | City name | [optional] 
**Province** | Pointer to **string** | Province name | [optional] 
**Country** | Pointer to **string** | Country name | [optional] 
**PostalCode** | Pointer to **string** | Postal Code | [optional] 
**DistrictId** | Pointer to **int64** | District ID | [optional] 
**CityId** | Pointer to **int64** | City ID | [optional] 
**ProvinceId** | Pointer to **int64** | Province ID | [optional] 
**Geo** | Pointer to **string** | Longitude and Latitude position for instant delivery | [optional] 

## Methods

### NewWebhookOrderNotificationRecipientAddress

`func NewWebhookOrderNotificationRecipientAddress() *WebhookOrderNotificationRecipientAddress`

NewWebhookOrderNotificationRecipientAddress instantiates a new WebhookOrderNotificationRecipientAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookOrderNotificationRecipientAddressWithDefaults

`func NewWebhookOrderNotificationRecipientAddressWithDefaults() *WebhookOrderNotificationRecipientAddress`

NewWebhookOrderNotificationRecipientAddressWithDefaults instantiates a new WebhookOrderNotificationRecipientAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressFull

`func (o *WebhookOrderNotificationRecipientAddress) GetAddressFull() string`

GetAddressFull returns the AddressFull field if non-nil, zero value otherwise.

### GetAddressFullOk

`func (o *WebhookOrderNotificationRecipientAddress) GetAddressFullOk() (*string, bool)`

GetAddressFullOk returns a tuple with the AddressFull field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressFull

`func (o *WebhookOrderNotificationRecipientAddress) SetAddressFull(v string)`

SetAddressFull sets AddressFull field to given value.

### HasAddressFull

`func (o *WebhookOrderNotificationRecipientAddress) HasAddressFull() bool`

HasAddressFull returns a boolean if a field has been set.

### GetDistrict

`func (o *WebhookOrderNotificationRecipientAddress) GetDistrict() string`

GetDistrict returns the District field if non-nil, zero value otherwise.

### GetDistrictOk

`func (o *WebhookOrderNotificationRecipientAddress) GetDistrictOk() (*string, bool)`

GetDistrictOk returns a tuple with the District field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistrict

`func (o *WebhookOrderNotificationRecipientAddress) SetDistrict(v string)`

SetDistrict sets District field to given value.

### HasDistrict

`func (o *WebhookOrderNotificationRecipientAddress) HasDistrict() bool`

HasDistrict returns a boolean if a field has been set.

### GetCity

`func (o *WebhookOrderNotificationRecipientAddress) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *WebhookOrderNotificationRecipientAddress) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *WebhookOrderNotificationRecipientAddress) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *WebhookOrderNotificationRecipientAddress) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetProvince

`func (o *WebhookOrderNotificationRecipientAddress) GetProvince() string`

GetProvince returns the Province field if non-nil, zero value otherwise.

### GetProvinceOk

`func (o *WebhookOrderNotificationRecipientAddress) GetProvinceOk() (*string, bool)`

GetProvinceOk returns a tuple with the Province field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvince

`func (o *WebhookOrderNotificationRecipientAddress) SetProvince(v string)`

SetProvince sets Province field to given value.

### HasProvince

`func (o *WebhookOrderNotificationRecipientAddress) HasProvince() bool`

HasProvince returns a boolean if a field has been set.

### GetCountry

`func (o *WebhookOrderNotificationRecipientAddress) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *WebhookOrderNotificationRecipientAddress) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *WebhookOrderNotificationRecipientAddress) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *WebhookOrderNotificationRecipientAddress) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetPostalCode

`func (o *WebhookOrderNotificationRecipientAddress) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *WebhookOrderNotificationRecipientAddress) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *WebhookOrderNotificationRecipientAddress) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *WebhookOrderNotificationRecipientAddress) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetDistrictId

`func (o *WebhookOrderNotificationRecipientAddress) GetDistrictId() int64`

GetDistrictId returns the DistrictId field if non-nil, zero value otherwise.

### GetDistrictIdOk

`func (o *WebhookOrderNotificationRecipientAddress) GetDistrictIdOk() (*int64, bool)`

GetDistrictIdOk returns a tuple with the DistrictId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistrictId

`func (o *WebhookOrderNotificationRecipientAddress) SetDistrictId(v int64)`

SetDistrictId sets DistrictId field to given value.

### HasDistrictId

`func (o *WebhookOrderNotificationRecipientAddress) HasDistrictId() bool`

HasDistrictId returns a boolean if a field has been set.

### GetCityId

`func (o *WebhookOrderNotificationRecipientAddress) GetCityId() int64`

GetCityId returns the CityId field if non-nil, zero value otherwise.

### GetCityIdOk

`func (o *WebhookOrderNotificationRecipientAddress) GetCityIdOk() (*int64, bool)`

GetCityIdOk returns a tuple with the CityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityId

`func (o *WebhookOrderNotificationRecipientAddress) SetCityId(v int64)`

SetCityId sets CityId field to given value.

### HasCityId

`func (o *WebhookOrderNotificationRecipientAddress) HasCityId() bool`

HasCityId returns a boolean if a field has been set.

### GetProvinceId

`func (o *WebhookOrderNotificationRecipientAddress) GetProvinceId() int64`

GetProvinceId returns the ProvinceId field if non-nil, zero value otherwise.

### GetProvinceIdOk

`func (o *WebhookOrderNotificationRecipientAddress) GetProvinceIdOk() (*int64, bool)`

GetProvinceIdOk returns a tuple with the ProvinceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvinceId

`func (o *WebhookOrderNotificationRecipientAddress) SetProvinceId(v int64)`

SetProvinceId sets ProvinceId field to given value.

### HasProvinceId

`func (o *WebhookOrderNotificationRecipientAddress) HasProvinceId() bool`

HasProvinceId returns a boolean if a field has been set.

### GetGeo

`func (o *WebhookOrderNotificationRecipientAddress) GetGeo() string`

GetGeo returns the Geo field if non-nil, zero value otherwise.

### GetGeoOk

`func (o *WebhookOrderNotificationRecipientAddress) GetGeoOk() (*string, bool)`

GetGeoOk returns a tuple with the Geo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeo

`func (o *WebhookOrderNotificationRecipientAddress) SetGeo(v string)`

SetGeo sets Geo field to given value.

### HasGeo

`func (o *WebhookOrderNotificationRecipientAddress) HasGeo() bool`

HasGeo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


