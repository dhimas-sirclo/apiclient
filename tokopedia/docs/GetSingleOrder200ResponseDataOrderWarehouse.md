# GetSingleOrder200ResponseDataOrderWarehouse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WarehouseId** | Pointer to **int64** | Warehouse Unique Identifier | [optional] 
**FulfillBy** | Pointer to **int64** | Order Fulfilled by Tokocabang (1), Order Fulfilled by Shop (0) | [optional] 
**MetaData** | Pointer to [**GetSingleOrder200ResponseDataOrderWarehouseMetaData**](GetSingleOrder200ResponseDataOrderWarehouseMetaData.md) |  | [optional] 

## Methods

### NewGetSingleOrder200ResponseDataOrderWarehouse

`func NewGetSingleOrder200ResponseDataOrderWarehouse() *GetSingleOrder200ResponseDataOrderWarehouse`

NewGetSingleOrder200ResponseDataOrderWarehouse instantiates a new GetSingleOrder200ResponseDataOrderWarehouse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSingleOrder200ResponseDataOrderWarehouseWithDefaults

`func NewGetSingleOrder200ResponseDataOrderWarehouseWithDefaults() *GetSingleOrder200ResponseDataOrderWarehouse`

NewGetSingleOrder200ResponseDataOrderWarehouseWithDefaults instantiates a new GetSingleOrder200ResponseDataOrderWarehouse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWarehouseId

`func (o *GetSingleOrder200ResponseDataOrderWarehouse) GetWarehouseId() int64`

GetWarehouseId returns the WarehouseId field if non-nil, zero value otherwise.

### GetWarehouseIdOk

`func (o *GetSingleOrder200ResponseDataOrderWarehouse) GetWarehouseIdOk() (*int64, bool)`

GetWarehouseIdOk returns a tuple with the WarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseId

`func (o *GetSingleOrder200ResponseDataOrderWarehouse) SetWarehouseId(v int64)`

SetWarehouseId sets WarehouseId field to given value.

### HasWarehouseId

`func (o *GetSingleOrder200ResponseDataOrderWarehouse) HasWarehouseId() bool`

HasWarehouseId returns a boolean if a field has been set.

### GetFulfillBy

`func (o *GetSingleOrder200ResponseDataOrderWarehouse) GetFulfillBy() int64`

GetFulfillBy returns the FulfillBy field if non-nil, zero value otherwise.

### GetFulfillByOk

`func (o *GetSingleOrder200ResponseDataOrderWarehouse) GetFulfillByOk() (*int64, bool)`

GetFulfillByOk returns a tuple with the FulfillBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillBy

`func (o *GetSingleOrder200ResponseDataOrderWarehouse) SetFulfillBy(v int64)`

SetFulfillBy sets FulfillBy field to given value.

### HasFulfillBy

`func (o *GetSingleOrder200ResponseDataOrderWarehouse) HasFulfillBy() bool`

HasFulfillBy returns a boolean if a field has been set.

### GetMetaData

`func (o *GetSingleOrder200ResponseDataOrderWarehouse) GetMetaData() GetSingleOrder200ResponseDataOrderWarehouseMetaData`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *GetSingleOrder200ResponseDataOrderWarehouse) GetMetaDataOk() (*GetSingleOrder200ResponseDataOrderWarehouseMetaData, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *GetSingleOrder200ResponseDataOrderWarehouse) SetMetaData(v GetSingleOrder200ResponseDataOrderWarehouseMetaData)`

SetMetaData sets MetaData field to given value.

### HasMetaData

`func (o *GetSingleOrder200ResponseDataOrderWarehouse) HasMetaData() bool`

HasMetaData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


