# CreateBundleRequestBundle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BundleName** | **string** | Name of the bundle with length less than or equals 15 characters and greater than or equals 3 characters | 
**BundleType** | **int64** | Type of the bundle where 1 is PAKET DISKON and 2 is BUNDLING , refers to Seller Education | 
**BundleStatus** | **int64** | Status of the bundle where &#x60;1 is ACTIVE | 
**StartTimeUnix** | **int64** | Start time of bundle in UNIX format must be greater than the current time | 
**WarehouseId** | **[]int64** | Warehouse unique identifier | 
**MaxOrder** | Pointer to **int64** | Maximum bundle package in one order | [optional] 
**StopTimeUnix** | Pointer to **int64** | Stop time of bundle in UNIX format must be greater than start time. Set to 0 for making a bundle with no time limit / until product stock is 0 | [optional] 
**OriginalQuota** | Pointer to **int64** | Original quota for bundle | [optional] 
**BundleItems** | [**[]CreateBundleRequestBundleBundleItemsInner**](CreateBundleRequestBundleBundleItemsInner.md) | List of products in this bundle | 

## Methods

### NewCreateBundleRequestBundle

`func NewCreateBundleRequestBundle(bundleName string, bundleType int64, bundleStatus int64, startTimeUnix int64, warehouseId []int64, bundleItems []CreateBundleRequestBundleBundleItemsInner, ) *CreateBundleRequestBundle`

NewCreateBundleRequestBundle instantiates a new CreateBundleRequestBundle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBundleRequestBundleWithDefaults

`func NewCreateBundleRequestBundleWithDefaults() *CreateBundleRequestBundle`

NewCreateBundleRequestBundleWithDefaults instantiates a new CreateBundleRequestBundle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBundleName

`func (o *CreateBundleRequestBundle) GetBundleName() string`

GetBundleName returns the BundleName field if non-nil, zero value otherwise.

### GetBundleNameOk

`func (o *CreateBundleRequestBundle) GetBundleNameOk() (*string, bool)`

GetBundleNameOk returns a tuple with the BundleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleName

`func (o *CreateBundleRequestBundle) SetBundleName(v string)`

SetBundleName sets BundleName field to given value.


### GetBundleType

`func (o *CreateBundleRequestBundle) GetBundleType() int64`

GetBundleType returns the BundleType field if non-nil, zero value otherwise.

### GetBundleTypeOk

`func (o *CreateBundleRequestBundle) GetBundleTypeOk() (*int64, bool)`

GetBundleTypeOk returns a tuple with the BundleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleType

`func (o *CreateBundleRequestBundle) SetBundleType(v int64)`

SetBundleType sets BundleType field to given value.


### GetBundleStatus

`func (o *CreateBundleRequestBundle) GetBundleStatus() int64`

GetBundleStatus returns the BundleStatus field if non-nil, zero value otherwise.

### GetBundleStatusOk

`func (o *CreateBundleRequestBundle) GetBundleStatusOk() (*int64, bool)`

GetBundleStatusOk returns a tuple with the BundleStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleStatus

`func (o *CreateBundleRequestBundle) SetBundleStatus(v int64)`

SetBundleStatus sets BundleStatus field to given value.


### GetStartTimeUnix

`func (o *CreateBundleRequestBundle) GetStartTimeUnix() int64`

GetStartTimeUnix returns the StartTimeUnix field if non-nil, zero value otherwise.

### GetStartTimeUnixOk

`func (o *CreateBundleRequestBundle) GetStartTimeUnixOk() (*int64, bool)`

GetStartTimeUnixOk returns a tuple with the StartTimeUnix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUnix

`func (o *CreateBundleRequestBundle) SetStartTimeUnix(v int64)`

SetStartTimeUnix sets StartTimeUnix field to given value.


### GetWarehouseId

`func (o *CreateBundleRequestBundle) GetWarehouseId() []int64`

GetWarehouseId returns the WarehouseId field if non-nil, zero value otherwise.

### GetWarehouseIdOk

`func (o *CreateBundleRequestBundle) GetWarehouseIdOk() (*[]int64, bool)`

GetWarehouseIdOk returns a tuple with the WarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseId

`func (o *CreateBundleRequestBundle) SetWarehouseId(v []int64)`

SetWarehouseId sets WarehouseId field to given value.


### GetMaxOrder

`func (o *CreateBundleRequestBundle) GetMaxOrder() int64`

GetMaxOrder returns the MaxOrder field if non-nil, zero value otherwise.

### GetMaxOrderOk

`func (o *CreateBundleRequestBundle) GetMaxOrderOk() (*int64, bool)`

GetMaxOrderOk returns a tuple with the MaxOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxOrder

`func (o *CreateBundleRequestBundle) SetMaxOrder(v int64)`

SetMaxOrder sets MaxOrder field to given value.

### HasMaxOrder

`func (o *CreateBundleRequestBundle) HasMaxOrder() bool`

HasMaxOrder returns a boolean if a field has been set.

### GetStopTimeUnix

`func (o *CreateBundleRequestBundle) GetStopTimeUnix() int64`

GetStopTimeUnix returns the StopTimeUnix field if non-nil, zero value otherwise.

### GetStopTimeUnixOk

`func (o *CreateBundleRequestBundle) GetStopTimeUnixOk() (*int64, bool)`

GetStopTimeUnixOk returns a tuple with the StopTimeUnix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopTimeUnix

`func (o *CreateBundleRequestBundle) SetStopTimeUnix(v int64)`

SetStopTimeUnix sets StopTimeUnix field to given value.

### HasStopTimeUnix

`func (o *CreateBundleRequestBundle) HasStopTimeUnix() bool`

HasStopTimeUnix returns a boolean if a field has been set.

### GetOriginalQuota

`func (o *CreateBundleRequestBundle) GetOriginalQuota() int64`

GetOriginalQuota returns the OriginalQuota field if non-nil, zero value otherwise.

### GetOriginalQuotaOk

`func (o *CreateBundleRequestBundle) GetOriginalQuotaOk() (*int64, bool)`

GetOriginalQuotaOk returns a tuple with the OriginalQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalQuota

`func (o *CreateBundleRequestBundle) SetOriginalQuota(v int64)`

SetOriginalQuota sets OriginalQuota field to given value.

### HasOriginalQuota

`func (o *CreateBundleRequestBundle) HasOriginalQuota() bool`

HasOriginalQuota returns a boolean if a field has been set.

### GetBundleItems

`func (o *CreateBundleRequestBundle) GetBundleItems() []CreateBundleRequestBundleBundleItemsInner`

GetBundleItems returns the BundleItems field if non-nil, zero value otherwise.

### GetBundleItemsOk

`func (o *CreateBundleRequestBundle) GetBundleItemsOk() (*[]CreateBundleRequestBundleBundleItemsInner, bool)`

GetBundleItemsOk returns a tuple with the BundleItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleItems

`func (o *CreateBundleRequestBundle) SetBundleItems(v []CreateBundleRequestBundleBundleItemsInner)`

SetBundleItems sets BundleItems field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


