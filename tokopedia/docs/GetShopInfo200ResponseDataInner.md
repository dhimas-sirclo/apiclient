# GetShopInfo200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShopId** | Pointer to **int64** |  | [optional] 
**UserId** | Pointer to **int64** |  | [optional] 
**ShopName** | Pointer to **string** |  | [optional] 
**Logo** | Pointer to **string** |  | [optional] 
**ShopUrl** | Pointer to **string** |  | [optional] 
**IsOpen** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **int64** |  | [optional] 
**DateShopCreated** | Pointer to **string** |  | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**AdminId** | Pointer to **[]int64** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**DistrictId** | Pointer to **int64** |  | [optional] 
**ProvinceName** | Pointer to **string** |  | [optional] 
**Warehouses** | Pointer to [**[]GetShopInfo200ResponseDataInnerWarehousesInner**](GetShopInfo200ResponseDataInnerWarehousesInner.md) |  | [optional] 
**SubscribeTokocabang** | Pointer to **bool** |  | [optional] 
**IsMitra** | Pointer to **bool** |  | [optional] 

## Methods

### NewGetShopInfo200ResponseDataInner

`func NewGetShopInfo200ResponseDataInner() *GetShopInfo200ResponseDataInner`

NewGetShopInfo200ResponseDataInner instantiates a new GetShopInfo200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetShopInfo200ResponseDataInnerWithDefaults

`func NewGetShopInfo200ResponseDataInnerWithDefaults() *GetShopInfo200ResponseDataInner`

NewGetShopInfo200ResponseDataInnerWithDefaults instantiates a new GetShopInfo200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShopId

`func (o *GetShopInfo200ResponseDataInner) GetShopId() int64`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *GetShopInfo200ResponseDataInner) GetShopIdOk() (*int64, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *GetShopInfo200ResponseDataInner) SetShopId(v int64)`

SetShopId sets ShopId field to given value.

### HasShopId

`func (o *GetShopInfo200ResponseDataInner) HasShopId() bool`

HasShopId returns a boolean if a field has been set.

### GetUserId

`func (o *GetShopInfo200ResponseDataInner) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *GetShopInfo200ResponseDataInner) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *GetShopInfo200ResponseDataInner) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *GetShopInfo200ResponseDataInner) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetShopName

`func (o *GetShopInfo200ResponseDataInner) GetShopName() string`

GetShopName returns the ShopName field if non-nil, zero value otherwise.

### GetShopNameOk

`func (o *GetShopInfo200ResponseDataInner) GetShopNameOk() (*string, bool)`

GetShopNameOk returns a tuple with the ShopName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopName

`func (o *GetShopInfo200ResponseDataInner) SetShopName(v string)`

SetShopName sets ShopName field to given value.

### HasShopName

`func (o *GetShopInfo200ResponseDataInner) HasShopName() bool`

HasShopName returns a boolean if a field has been set.

### GetLogo

`func (o *GetShopInfo200ResponseDataInner) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *GetShopInfo200ResponseDataInner) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *GetShopInfo200ResponseDataInner) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *GetShopInfo200ResponseDataInner) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetShopUrl

`func (o *GetShopInfo200ResponseDataInner) GetShopUrl() string`

GetShopUrl returns the ShopUrl field if non-nil, zero value otherwise.

### GetShopUrlOk

`func (o *GetShopInfo200ResponseDataInner) GetShopUrlOk() (*string, bool)`

GetShopUrlOk returns a tuple with the ShopUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopUrl

`func (o *GetShopInfo200ResponseDataInner) SetShopUrl(v string)`

SetShopUrl sets ShopUrl field to given value.

### HasShopUrl

`func (o *GetShopInfo200ResponseDataInner) HasShopUrl() bool`

HasShopUrl returns a boolean if a field has been set.

### GetIsOpen

`func (o *GetShopInfo200ResponseDataInner) GetIsOpen() int64`

GetIsOpen returns the IsOpen field if non-nil, zero value otherwise.

### GetIsOpenOk

`func (o *GetShopInfo200ResponseDataInner) GetIsOpenOk() (*int64, bool)`

GetIsOpenOk returns a tuple with the IsOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOpen

`func (o *GetShopInfo200ResponseDataInner) SetIsOpen(v int64)`

SetIsOpen sets IsOpen field to given value.

### HasIsOpen

`func (o *GetShopInfo200ResponseDataInner) HasIsOpen() bool`

HasIsOpen returns a boolean if a field has been set.

### GetStatus

`func (o *GetShopInfo200ResponseDataInner) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetShopInfo200ResponseDataInner) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetShopInfo200ResponseDataInner) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetShopInfo200ResponseDataInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDateShopCreated

`func (o *GetShopInfo200ResponseDataInner) GetDateShopCreated() string`

GetDateShopCreated returns the DateShopCreated field if non-nil, zero value otherwise.

### GetDateShopCreatedOk

`func (o *GetShopInfo200ResponseDataInner) GetDateShopCreatedOk() (*string, bool)`

GetDateShopCreatedOk returns a tuple with the DateShopCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateShopCreated

`func (o *GetShopInfo200ResponseDataInner) SetDateShopCreated(v string)`

SetDateShopCreated sets DateShopCreated field to given value.

### HasDateShopCreated

`func (o *GetShopInfo200ResponseDataInner) HasDateShopCreated() bool`

HasDateShopCreated returns a boolean if a field has been set.

### GetDomain

`func (o *GetShopInfo200ResponseDataInner) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *GetShopInfo200ResponseDataInner) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *GetShopInfo200ResponseDataInner) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *GetShopInfo200ResponseDataInner) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetAdminId

`func (o *GetShopInfo200ResponseDataInner) GetAdminId() []int64`

GetAdminId returns the AdminId field if non-nil, zero value otherwise.

### GetAdminIdOk

`func (o *GetShopInfo200ResponseDataInner) GetAdminIdOk() (*[]int64, bool)`

GetAdminIdOk returns a tuple with the AdminId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminId

`func (o *GetShopInfo200ResponseDataInner) SetAdminId(v []int64)`

SetAdminId sets AdminId field to given value.

### HasAdminId

`func (o *GetShopInfo200ResponseDataInner) HasAdminId() bool`

HasAdminId returns a boolean if a field has been set.

### GetReason

`func (o *GetShopInfo200ResponseDataInner) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *GetShopInfo200ResponseDataInner) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *GetShopInfo200ResponseDataInner) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *GetShopInfo200ResponseDataInner) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetDistrictId

`func (o *GetShopInfo200ResponseDataInner) GetDistrictId() int64`

GetDistrictId returns the DistrictId field if non-nil, zero value otherwise.

### GetDistrictIdOk

`func (o *GetShopInfo200ResponseDataInner) GetDistrictIdOk() (*int64, bool)`

GetDistrictIdOk returns a tuple with the DistrictId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistrictId

`func (o *GetShopInfo200ResponseDataInner) SetDistrictId(v int64)`

SetDistrictId sets DistrictId field to given value.

### HasDistrictId

`func (o *GetShopInfo200ResponseDataInner) HasDistrictId() bool`

HasDistrictId returns a boolean if a field has been set.

### GetProvinceName

`func (o *GetShopInfo200ResponseDataInner) GetProvinceName() string`

GetProvinceName returns the ProvinceName field if non-nil, zero value otherwise.

### GetProvinceNameOk

`func (o *GetShopInfo200ResponseDataInner) GetProvinceNameOk() (*string, bool)`

GetProvinceNameOk returns a tuple with the ProvinceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvinceName

`func (o *GetShopInfo200ResponseDataInner) SetProvinceName(v string)`

SetProvinceName sets ProvinceName field to given value.

### HasProvinceName

`func (o *GetShopInfo200ResponseDataInner) HasProvinceName() bool`

HasProvinceName returns a boolean if a field has been set.

### GetWarehouses

`func (o *GetShopInfo200ResponseDataInner) GetWarehouses() []GetShopInfo200ResponseDataInnerWarehousesInner`

GetWarehouses returns the Warehouses field if non-nil, zero value otherwise.

### GetWarehousesOk

`func (o *GetShopInfo200ResponseDataInner) GetWarehousesOk() (*[]GetShopInfo200ResponseDataInnerWarehousesInner, bool)`

GetWarehousesOk returns a tuple with the Warehouses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouses

`func (o *GetShopInfo200ResponseDataInner) SetWarehouses(v []GetShopInfo200ResponseDataInnerWarehousesInner)`

SetWarehouses sets Warehouses field to given value.

### HasWarehouses

`func (o *GetShopInfo200ResponseDataInner) HasWarehouses() bool`

HasWarehouses returns a boolean if a field has been set.

### GetSubscribeTokocabang

`func (o *GetShopInfo200ResponseDataInner) GetSubscribeTokocabang() bool`

GetSubscribeTokocabang returns the SubscribeTokocabang field if non-nil, zero value otherwise.

### GetSubscribeTokocabangOk

`func (o *GetShopInfo200ResponseDataInner) GetSubscribeTokocabangOk() (*bool, bool)`

GetSubscribeTokocabangOk returns a tuple with the SubscribeTokocabang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribeTokocabang

`func (o *GetShopInfo200ResponseDataInner) SetSubscribeTokocabang(v bool)`

SetSubscribeTokocabang sets SubscribeTokocabang field to given value.

### HasSubscribeTokocabang

`func (o *GetShopInfo200ResponseDataInner) HasSubscribeTokocabang() bool`

HasSubscribeTokocabang returns a boolean if a field has been set.

### GetIsMitra

`func (o *GetShopInfo200ResponseDataInner) GetIsMitra() bool`

GetIsMitra returns the IsMitra field if non-nil, zero value otherwise.

### GetIsMitraOk

`func (o *GetShopInfo200ResponseDataInner) GetIsMitraOk() (*bool, bool)`

GetIsMitraOk returns a tuple with the IsMitra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMitra

`func (o *GetShopInfo200ResponseDataInner) SetIsMitra(v bool)`

SetIsMitra sets IsMitra field to given value.

### HasIsMitra

`func (o *GetShopInfo200ResponseDataInner) HasIsMitra() bool`

HasIsMitra returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


