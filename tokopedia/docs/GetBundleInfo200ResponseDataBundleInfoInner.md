# GetBundleInfo200ResponseDataBundleInfoInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BundleId** | Pointer to **int64** |  | [optional] 
**GroupId** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **int64** |  | [optional] 
**ShopId** | Pointer to **int64** |  | [optional] 
**StartTimeUnix** | Pointer to **int64** |  | [optional] 
**StopTimeUnix** | Pointer to **int64** |  | [optional] 
**BundleItem** | Pointer to [**[]GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInner**](GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInner.md) |  | [optional] 
**WarehouseId** | Pointer to **int64** |  | [optional] 
**Quota** | Pointer to **int64** |  | [optional] 
**OriginalQuota** | Pointer to **int64** |  | [optional] 
**MaxOrder** | Pointer to **int64** |  | [optional] 
**Preorder** | Pointer to [**GetBundleInfo200ResponseDataBundleInfoInnerPreorder**](GetBundleInfo200ResponseDataBundleInfoInnerPreorder.md) |  | [optional] 

## Methods

### NewGetBundleInfo200ResponseDataBundleInfoInner

`func NewGetBundleInfo200ResponseDataBundleInfoInner() *GetBundleInfo200ResponseDataBundleInfoInner`

NewGetBundleInfo200ResponseDataBundleInfoInner instantiates a new GetBundleInfo200ResponseDataBundleInfoInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBundleInfo200ResponseDataBundleInfoInnerWithDefaults

`func NewGetBundleInfo200ResponseDataBundleInfoInnerWithDefaults() *GetBundleInfo200ResponseDataBundleInfoInner`

NewGetBundleInfo200ResponseDataBundleInfoInnerWithDefaults instantiates a new GetBundleInfo200ResponseDataBundleInfoInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBundleId

`func (o *GetBundleInfo200ResponseDataBundleInfoInner) GetBundleId() int64`

GetBundleId returns the BundleId field if non-nil, zero value otherwise.

### GetBundleIdOk

`func (o *GetBundleInfo200ResponseDataBundleInfoInner) GetBundleIdOk() (*int64, bool)`

GetBundleIdOk returns a tuple with the BundleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleId

`func (o *GetBundleInfo200ResponseDataBundleInfoInner) SetBundleId(v int64)`

SetBundleId sets BundleId field to given value.

### HasBundleId

`func (o *GetBundleInfo200ResponseDataBundleInfoInner) HasBundleId() bool`

HasBundleId returns a boolean if a field has been set.

### GetGroupId

`func (o *GetBundleInfo200ResponseDataBundleInfoInner) GetGroupId() int64`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *GetBundleInfo200ResponseDataBundleInfoInner) GetGroupIdOk() (*int64, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *GetBundleInfo200ResponseDataBundleInfoInner) SetGroupId(v int64)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *GetBundleInfo200ResponseDataBundleInfoInner) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetName

`func (o *GetBundleInfo200ResponseDataBundleInfoInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetBundleInfo200ResponseDataBundleInfoInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetBundleInfo200ResponseDataBundleInfoInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetBundleInfo200ResponseDataBundleInfoInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *GetBundleInfo200ResponseDataBundleInfoInner) GetType() int64`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetBundleInfo200ResponseDataBundleInfoInner) GetTypeOk() (*int64, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetBundleInfo200ResponseDataBundleInfoInner) SetType(v int64)`

SetType sets Type field to given value.

### HasType

`func (o *GetBundleInfo200ResponseDataBundleInfoInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *GetBundleInfo200ResponseDataBundleInfoInner) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetBundleInfo200ResponseDataBundleInfoInner) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetBundleInfo200ResponseDataBundleInfoInner) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetBundleInfo200ResponseDataBundleInfoInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetShopId

`func (o *GetBundleInfo200ResponseDataBundleInfoInner) GetShopId() int64`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *GetBundleInfo200ResponseDataBundleInfoInner) GetShopIdOk() (*int64, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *GetBundleInfo200ResponseDataBundleInfoInner) SetShopId(v int64)`

SetShopId sets ShopId field to given value.

### HasShopId

`func (o *GetBundleInfo200ResponseDataBundleInfoInner) HasShopId() bool`

HasShopId returns a boolean if a field has been set.

### GetStartTimeUnix

`func (o *GetBundleInfo200ResponseDataBundleInfoInner) GetStartTimeUnix() int64`

GetStartTimeUnix returns the StartTimeUnix field if non-nil, zero value otherwise.

### GetStartTimeUnixOk

`func (o *GetBundleInfo200ResponseDataBundleInfoInner) GetStartTimeUnixOk() (*int64, bool)`

GetStartTimeUnixOk returns a tuple with the StartTimeUnix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUnix

`func (o *GetBundleInfo200ResponseDataBundleInfoInner) SetStartTimeUnix(v int64)`

SetStartTimeUnix sets StartTimeUnix field to given value.

### HasStartTimeUnix

`func (o *GetBundleInfo200ResponseDataBundleInfoInner) HasStartTimeUnix() bool`

HasStartTimeUnix returns a boolean if a field has been set.

### GetStopTimeUnix

`func (o *GetBundleInfo200ResponseDataBundleInfoInner) GetStopTimeUnix() int64`

GetStopTimeUnix returns the StopTimeUnix field if non-nil, zero value otherwise.

### GetStopTimeUnixOk

`func (o *GetBundleInfo200ResponseDataBundleInfoInner) GetStopTimeUnixOk() (*int64, bool)`

GetStopTimeUnixOk returns a tuple with the StopTimeUnix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopTimeUnix

`func (o *GetBundleInfo200ResponseDataBundleInfoInner) SetStopTimeUnix(v int64)`

SetStopTimeUnix sets StopTimeUnix field to given value.

### HasStopTimeUnix

`func (o *GetBundleInfo200ResponseDataBundleInfoInner) HasStopTimeUnix() bool`

HasStopTimeUnix returns a boolean if a field has been set.

### GetBundleItem

`func (o *GetBundleInfo200ResponseDataBundleInfoInner) GetBundleItem() []GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInner`

GetBundleItem returns the BundleItem field if non-nil, zero value otherwise.

### GetBundleItemOk

`func (o *GetBundleInfo200ResponseDataBundleInfoInner) GetBundleItemOk() (*[]GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInner, bool)`

GetBundleItemOk returns a tuple with the BundleItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleItem

`func (o *GetBundleInfo200ResponseDataBundleInfoInner) SetBundleItem(v []GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInner)`

SetBundleItem sets BundleItem field to given value.

### HasBundleItem

`func (o *GetBundleInfo200ResponseDataBundleInfoInner) HasBundleItem() bool`

HasBundleItem returns a boolean if a field has been set.

### GetWarehouseId

`func (o *GetBundleInfo200ResponseDataBundleInfoInner) GetWarehouseId() int64`

GetWarehouseId returns the WarehouseId field if non-nil, zero value otherwise.

### GetWarehouseIdOk

`func (o *GetBundleInfo200ResponseDataBundleInfoInner) GetWarehouseIdOk() (*int64, bool)`

GetWarehouseIdOk returns a tuple with the WarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseId

`func (o *GetBundleInfo200ResponseDataBundleInfoInner) SetWarehouseId(v int64)`

SetWarehouseId sets WarehouseId field to given value.

### HasWarehouseId

`func (o *GetBundleInfo200ResponseDataBundleInfoInner) HasWarehouseId() bool`

HasWarehouseId returns a boolean if a field has been set.

### GetQuota

`func (o *GetBundleInfo200ResponseDataBundleInfoInner) GetQuota() int64`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *GetBundleInfo200ResponseDataBundleInfoInner) GetQuotaOk() (*int64, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *GetBundleInfo200ResponseDataBundleInfoInner) SetQuota(v int64)`

SetQuota sets Quota field to given value.

### HasQuota

`func (o *GetBundleInfo200ResponseDataBundleInfoInner) HasQuota() bool`

HasQuota returns a boolean if a field has been set.

### GetOriginalQuota

`func (o *GetBundleInfo200ResponseDataBundleInfoInner) GetOriginalQuota() int64`

GetOriginalQuota returns the OriginalQuota field if non-nil, zero value otherwise.

### GetOriginalQuotaOk

`func (o *GetBundleInfo200ResponseDataBundleInfoInner) GetOriginalQuotaOk() (*int64, bool)`

GetOriginalQuotaOk returns a tuple with the OriginalQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalQuota

`func (o *GetBundleInfo200ResponseDataBundleInfoInner) SetOriginalQuota(v int64)`

SetOriginalQuota sets OriginalQuota field to given value.

### HasOriginalQuota

`func (o *GetBundleInfo200ResponseDataBundleInfoInner) HasOriginalQuota() bool`

HasOriginalQuota returns a boolean if a field has been set.

### GetMaxOrder

`func (o *GetBundleInfo200ResponseDataBundleInfoInner) GetMaxOrder() int64`

GetMaxOrder returns the MaxOrder field if non-nil, zero value otherwise.

### GetMaxOrderOk

`func (o *GetBundleInfo200ResponseDataBundleInfoInner) GetMaxOrderOk() (*int64, bool)`

GetMaxOrderOk returns a tuple with the MaxOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxOrder

`func (o *GetBundleInfo200ResponseDataBundleInfoInner) SetMaxOrder(v int64)`

SetMaxOrder sets MaxOrder field to given value.

### HasMaxOrder

`func (o *GetBundleInfo200ResponseDataBundleInfoInner) HasMaxOrder() bool`

HasMaxOrder returns a boolean if a field has been set.

### GetPreorder

`func (o *GetBundleInfo200ResponseDataBundleInfoInner) GetPreorder() GetBundleInfo200ResponseDataBundleInfoInnerPreorder`

GetPreorder returns the Preorder field if non-nil, zero value otherwise.

### GetPreorderOk

`func (o *GetBundleInfo200ResponseDataBundleInfoInner) GetPreorderOk() (*GetBundleInfo200ResponseDataBundleInfoInnerPreorder, bool)`

GetPreorderOk returns a tuple with the Preorder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreorder

`func (o *GetBundleInfo200ResponseDataBundleInfoInner) SetPreorder(v GetBundleInfo200ResponseDataBundleInfoInnerPreorder)`

SetPreorder sets Preorder field to given value.

### HasPreorder

`func (o *GetBundleInfo200ResponseDataBundleInfoInner) HasPreorder() bool`

HasPreorder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


