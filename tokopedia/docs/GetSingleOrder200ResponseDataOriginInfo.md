# GetSingleOrder200ResponseDataOriginInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SenderName** | Pointer to **string** | Sender Name | [optional] 
**OriginProvince** | Pointer to **int64** | Origin Province Unique Identifier | [optional] 
**OriginProvinceName** | Pointer to **string** | Origin Province Name | [optional] 
**OriginCity** | Pointer to **int64** | Origin City Unique Identifier | [optional] 
**OriginCityName** | Pointer to **string** | Origin City Name | [optional] 
**OriginAddress** | Pointer to **string** | Origin Address | [optional] 
**OriginDistrict** | Pointer to **int64** | Origin District Unique Identifier | [optional] 
**OriginDistrictName** | Pointer to **string** | Origin District Name | [optional] 
**OriginPostalCode** | Pointer to **string** | Origin Postal Code | [optional] 
**OriginGeo** | Pointer to **string** | Origin Geolocation Coordinate | [optional] 
**ReceiverName** | Pointer to **string** | Receiver Name | [optional] 
**DestinationAddress** | Pointer to **string** | Destination Address | [optional] 
**DestinationProvince** | Pointer to **int64** | Destination Province Unique Identifier | [optional] 
**DestinationCity** | Pointer to **int64** | Destination City Unique Identifier | [optional] 
**DestinationDistrict** | Pointer to **int64** | Destination Disctrict Unique Identifier | [optional] 
**DestinationPostalCode** | Pointer to **string** | Destination Disctrict Postal Code | [optional] 
**DestinationGeo** | Pointer to **string** | Destination Geolocation Coordinate | [optional] 
**DestinationLoc** | Pointer to [**GetSingleOrder200ResponseDataOriginInfoDestinationLoc**](GetSingleOrder200ResponseDataOriginInfoDestinationLoc.md) |  | [optional] 

## Methods

### NewGetSingleOrder200ResponseDataOriginInfo

`func NewGetSingleOrder200ResponseDataOriginInfo() *GetSingleOrder200ResponseDataOriginInfo`

NewGetSingleOrder200ResponseDataOriginInfo instantiates a new GetSingleOrder200ResponseDataOriginInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSingleOrder200ResponseDataOriginInfoWithDefaults

`func NewGetSingleOrder200ResponseDataOriginInfoWithDefaults() *GetSingleOrder200ResponseDataOriginInfo`

NewGetSingleOrder200ResponseDataOriginInfoWithDefaults instantiates a new GetSingleOrder200ResponseDataOriginInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSenderName

`func (o *GetSingleOrder200ResponseDataOriginInfo) GetSenderName() string`

GetSenderName returns the SenderName field if non-nil, zero value otherwise.

### GetSenderNameOk

`func (o *GetSingleOrder200ResponseDataOriginInfo) GetSenderNameOk() (*string, bool)`

GetSenderNameOk returns a tuple with the SenderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderName

`func (o *GetSingleOrder200ResponseDataOriginInfo) SetSenderName(v string)`

SetSenderName sets SenderName field to given value.

### HasSenderName

`func (o *GetSingleOrder200ResponseDataOriginInfo) HasSenderName() bool`

HasSenderName returns a boolean if a field has been set.

### GetOriginProvince

`func (o *GetSingleOrder200ResponseDataOriginInfo) GetOriginProvince() int64`

GetOriginProvince returns the OriginProvince field if non-nil, zero value otherwise.

### GetOriginProvinceOk

`func (o *GetSingleOrder200ResponseDataOriginInfo) GetOriginProvinceOk() (*int64, bool)`

GetOriginProvinceOk returns a tuple with the OriginProvince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginProvince

`func (o *GetSingleOrder200ResponseDataOriginInfo) SetOriginProvince(v int64)`

SetOriginProvince sets OriginProvince field to given value.

### HasOriginProvince

`func (o *GetSingleOrder200ResponseDataOriginInfo) HasOriginProvince() bool`

HasOriginProvince returns a boolean if a field has been set.

### GetOriginProvinceName

`func (o *GetSingleOrder200ResponseDataOriginInfo) GetOriginProvinceName() string`

GetOriginProvinceName returns the OriginProvinceName field if non-nil, zero value otherwise.

### GetOriginProvinceNameOk

`func (o *GetSingleOrder200ResponseDataOriginInfo) GetOriginProvinceNameOk() (*string, bool)`

GetOriginProvinceNameOk returns a tuple with the OriginProvinceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginProvinceName

`func (o *GetSingleOrder200ResponseDataOriginInfo) SetOriginProvinceName(v string)`

SetOriginProvinceName sets OriginProvinceName field to given value.

### HasOriginProvinceName

`func (o *GetSingleOrder200ResponseDataOriginInfo) HasOriginProvinceName() bool`

HasOriginProvinceName returns a boolean if a field has been set.

### GetOriginCity

`func (o *GetSingleOrder200ResponseDataOriginInfo) GetOriginCity() int64`

GetOriginCity returns the OriginCity field if non-nil, zero value otherwise.

### GetOriginCityOk

`func (o *GetSingleOrder200ResponseDataOriginInfo) GetOriginCityOk() (*int64, bool)`

GetOriginCityOk returns a tuple with the OriginCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginCity

`func (o *GetSingleOrder200ResponseDataOriginInfo) SetOriginCity(v int64)`

SetOriginCity sets OriginCity field to given value.

### HasOriginCity

`func (o *GetSingleOrder200ResponseDataOriginInfo) HasOriginCity() bool`

HasOriginCity returns a boolean if a field has been set.

### GetOriginCityName

`func (o *GetSingleOrder200ResponseDataOriginInfo) GetOriginCityName() string`

GetOriginCityName returns the OriginCityName field if non-nil, zero value otherwise.

### GetOriginCityNameOk

`func (o *GetSingleOrder200ResponseDataOriginInfo) GetOriginCityNameOk() (*string, bool)`

GetOriginCityNameOk returns a tuple with the OriginCityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginCityName

`func (o *GetSingleOrder200ResponseDataOriginInfo) SetOriginCityName(v string)`

SetOriginCityName sets OriginCityName field to given value.

### HasOriginCityName

`func (o *GetSingleOrder200ResponseDataOriginInfo) HasOriginCityName() bool`

HasOriginCityName returns a boolean if a field has been set.

### GetOriginAddress

`func (o *GetSingleOrder200ResponseDataOriginInfo) GetOriginAddress() string`

GetOriginAddress returns the OriginAddress field if non-nil, zero value otherwise.

### GetOriginAddressOk

`func (o *GetSingleOrder200ResponseDataOriginInfo) GetOriginAddressOk() (*string, bool)`

GetOriginAddressOk returns a tuple with the OriginAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginAddress

`func (o *GetSingleOrder200ResponseDataOriginInfo) SetOriginAddress(v string)`

SetOriginAddress sets OriginAddress field to given value.

### HasOriginAddress

`func (o *GetSingleOrder200ResponseDataOriginInfo) HasOriginAddress() bool`

HasOriginAddress returns a boolean if a field has been set.

### GetOriginDistrict

`func (o *GetSingleOrder200ResponseDataOriginInfo) GetOriginDistrict() int64`

GetOriginDistrict returns the OriginDistrict field if non-nil, zero value otherwise.

### GetOriginDistrictOk

`func (o *GetSingleOrder200ResponseDataOriginInfo) GetOriginDistrictOk() (*int64, bool)`

GetOriginDistrictOk returns a tuple with the OriginDistrict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginDistrict

`func (o *GetSingleOrder200ResponseDataOriginInfo) SetOriginDistrict(v int64)`

SetOriginDistrict sets OriginDistrict field to given value.

### HasOriginDistrict

`func (o *GetSingleOrder200ResponseDataOriginInfo) HasOriginDistrict() bool`

HasOriginDistrict returns a boolean if a field has been set.

### GetOriginDistrictName

`func (o *GetSingleOrder200ResponseDataOriginInfo) GetOriginDistrictName() string`

GetOriginDistrictName returns the OriginDistrictName field if non-nil, zero value otherwise.

### GetOriginDistrictNameOk

`func (o *GetSingleOrder200ResponseDataOriginInfo) GetOriginDistrictNameOk() (*string, bool)`

GetOriginDistrictNameOk returns a tuple with the OriginDistrictName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginDistrictName

`func (o *GetSingleOrder200ResponseDataOriginInfo) SetOriginDistrictName(v string)`

SetOriginDistrictName sets OriginDistrictName field to given value.

### HasOriginDistrictName

`func (o *GetSingleOrder200ResponseDataOriginInfo) HasOriginDistrictName() bool`

HasOriginDistrictName returns a boolean if a field has been set.

### GetOriginPostalCode

`func (o *GetSingleOrder200ResponseDataOriginInfo) GetOriginPostalCode() string`

GetOriginPostalCode returns the OriginPostalCode field if non-nil, zero value otherwise.

### GetOriginPostalCodeOk

`func (o *GetSingleOrder200ResponseDataOriginInfo) GetOriginPostalCodeOk() (*string, bool)`

GetOriginPostalCodeOk returns a tuple with the OriginPostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginPostalCode

`func (o *GetSingleOrder200ResponseDataOriginInfo) SetOriginPostalCode(v string)`

SetOriginPostalCode sets OriginPostalCode field to given value.

### HasOriginPostalCode

`func (o *GetSingleOrder200ResponseDataOriginInfo) HasOriginPostalCode() bool`

HasOriginPostalCode returns a boolean if a field has been set.

### GetOriginGeo

`func (o *GetSingleOrder200ResponseDataOriginInfo) GetOriginGeo() string`

GetOriginGeo returns the OriginGeo field if non-nil, zero value otherwise.

### GetOriginGeoOk

`func (o *GetSingleOrder200ResponseDataOriginInfo) GetOriginGeoOk() (*string, bool)`

GetOriginGeoOk returns a tuple with the OriginGeo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginGeo

`func (o *GetSingleOrder200ResponseDataOriginInfo) SetOriginGeo(v string)`

SetOriginGeo sets OriginGeo field to given value.

### HasOriginGeo

`func (o *GetSingleOrder200ResponseDataOriginInfo) HasOriginGeo() bool`

HasOriginGeo returns a boolean if a field has been set.

### GetReceiverName

`func (o *GetSingleOrder200ResponseDataOriginInfo) GetReceiverName() string`

GetReceiverName returns the ReceiverName field if non-nil, zero value otherwise.

### GetReceiverNameOk

`func (o *GetSingleOrder200ResponseDataOriginInfo) GetReceiverNameOk() (*string, bool)`

GetReceiverNameOk returns a tuple with the ReceiverName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverName

`func (o *GetSingleOrder200ResponseDataOriginInfo) SetReceiverName(v string)`

SetReceiverName sets ReceiverName field to given value.

### HasReceiverName

`func (o *GetSingleOrder200ResponseDataOriginInfo) HasReceiverName() bool`

HasReceiverName returns a boolean if a field has been set.

### GetDestinationAddress

`func (o *GetSingleOrder200ResponseDataOriginInfo) GetDestinationAddress() string`

GetDestinationAddress returns the DestinationAddress field if non-nil, zero value otherwise.

### GetDestinationAddressOk

`func (o *GetSingleOrder200ResponseDataOriginInfo) GetDestinationAddressOk() (*string, bool)`

GetDestinationAddressOk returns a tuple with the DestinationAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAddress

`func (o *GetSingleOrder200ResponseDataOriginInfo) SetDestinationAddress(v string)`

SetDestinationAddress sets DestinationAddress field to given value.

### HasDestinationAddress

`func (o *GetSingleOrder200ResponseDataOriginInfo) HasDestinationAddress() bool`

HasDestinationAddress returns a boolean if a field has been set.

### GetDestinationProvince

`func (o *GetSingleOrder200ResponseDataOriginInfo) GetDestinationProvince() int64`

GetDestinationProvince returns the DestinationProvince field if non-nil, zero value otherwise.

### GetDestinationProvinceOk

`func (o *GetSingleOrder200ResponseDataOriginInfo) GetDestinationProvinceOk() (*int64, bool)`

GetDestinationProvinceOk returns a tuple with the DestinationProvince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationProvince

`func (o *GetSingleOrder200ResponseDataOriginInfo) SetDestinationProvince(v int64)`

SetDestinationProvince sets DestinationProvince field to given value.

### HasDestinationProvince

`func (o *GetSingleOrder200ResponseDataOriginInfo) HasDestinationProvince() bool`

HasDestinationProvince returns a boolean if a field has been set.

### GetDestinationCity

`func (o *GetSingleOrder200ResponseDataOriginInfo) GetDestinationCity() int64`

GetDestinationCity returns the DestinationCity field if non-nil, zero value otherwise.

### GetDestinationCityOk

`func (o *GetSingleOrder200ResponseDataOriginInfo) GetDestinationCityOk() (*int64, bool)`

GetDestinationCityOk returns a tuple with the DestinationCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationCity

`func (o *GetSingleOrder200ResponseDataOriginInfo) SetDestinationCity(v int64)`

SetDestinationCity sets DestinationCity field to given value.

### HasDestinationCity

`func (o *GetSingleOrder200ResponseDataOriginInfo) HasDestinationCity() bool`

HasDestinationCity returns a boolean if a field has been set.

### GetDestinationDistrict

`func (o *GetSingleOrder200ResponseDataOriginInfo) GetDestinationDistrict() int64`

GetDestinationDistrict returns the DestinationDistrict field if non-nil, zero value otherwise.

### GetDestinationDistrictOk

`func (o *GetSingleOrder200ResponseDataOriginInfo) GetDestinationDistrictOk() (*int64, bool)`

GetDestinationDistrictOk returns a tuple with the DestinationDistrict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationDistrict

`func (o *GetSingleOrder200ResponseDataOriginInfo) SetDestinationDistrict(v int64)`

SetDestinationDistrict sets DestinationDistrict field to given value.

### HasDestinationDistrict

`func (o *GetSingleOrder200ResponseDataOriginInfo) HasDestinationDistrict() bool`

HasDestinationDistrict returns a boolean if a field has been set.

### GetDestinationPostalCode

`func (o *GetSingleOrder200ResponseDataOriginInfo) GetDestinationPostalCode() string`

GetDestinationPostalCode returns the DestinationPostalCode field if non-nil, zero value otherwise.

### GetDestinationPostalCodeOk

`func (o *GetSingleOrder200ResponseDataOriginInfo) GetDestinationPostalCodeOk() (*string, bool)`

GetDestinationPostalCodeOk returns a tuple with the DestinationPostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationPostalCode

`func (o *GetSingleOrder200ResponseDataOriginInfo) SetDestinationPostalCode(v string)`

SetDestinationPostalCode sets DestinationPostalCode field to given value.

### HasDestinationPostalCode

`func (o *GetSingleOrder200ResponseDataOriginInfo) HasDestinationPostalCode() bool`

HasDestinationPostalCode returns a boolean if a field has been set.

### GetDestinationGeo

`func (o *GetSingleOrder200ResponseDataOriginInfo) GetDestinationGeo() string`

GetDestinationGeo returns the DestinationGeo field if non-nil, zero value otherwise.

### GetDestinationGeoOk

`func (o *GetSingleOrder200ResponseDataOriginInfo) GetDestinationGeoOk() (*string, bool)`

GetDestinationGeoOk returns a tuple with the DestinationGeo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationGeo

`func (o *GetSingleOrder200ResponseDataOriginInfo) SetDestinationGeo(v string)`

SetDestinationGeo sets DestinationGeo field to given value.

### HasDestinationGeo

`func (o *GetSingleOrder200ResponseDataOriginInfo) HasDestinationGeo() bool`

HasDestinationGeo returns a boolean if a field has been set.

### GetDestinationLoc

`func (o *GetSingleOrder200ResponseDataOriginInfo) GetDestinationLoc() GetSingleOrder200ResponseDataOriginInfoDestinationLoc`

GetDestinationLoc returns the DestinationLoc field if non-nil, zero value otherwise.

### GetDestinationLocOk

`func (o *GetSingleOrder200ResponseDataOriginInfo) GetDestinationLocOk() (*GetSingleOrder200ResponseDataOriginInfoDestinationLoc, bool)`

GetDestinationLocOk returns a tuple with the DestinationLoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationLoc

`func (o *GetSingleOrder200ResponseDataOriginInfo) SetDestinationLoc(v GetSingleOrder200ResponseDataOriginInfoDestinationLoc)`

SetDestinationLoc sets DestinationLoc field to given value.

### HasDestinationLoc

`func (o *GetSingleOrder200ResponseDataOriginInfo) HasDestinationLoc() bool`

HasDestinationLoc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


