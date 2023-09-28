# GetSingleOrder200ResponseDataPofInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | Pointer to **int64** | Order Unique Identifier | [optional] 
**PofStatus** | Pointer to **int64** | Partial Order Fulfillment Status Code | [optional] 
**ShippingPrice** | Pointer to **int64** | Shipping Price | [optional] 
**InsurancePrice** | Pointer to **int64** | Insurance Price | [optional] 
**CreateTime** | Pointer to **string** | POF Request Create Timestamp (format: 2023-04-03T08:36:29.331837Z)  | [optional] 
**CreateBy** | Pointer to **int64** | POF Request Create By | [optional] 
**CreateByType** | Pointer to **int64** | POF Request Create By Type | [optional] 
**UpdateTime** | Pointer to **string** | POF Request Update Timestamp (format: 2023-04-03T08:36:36.482397Z)  | [optional] 
**UpdateBy** | Pointer to **int64** | POF Request Update By | [optional] 
**UpdateByType** | Pointer to **int64** | POF Request Update By Type | [optional] 
**PofDetailInfo** | Pointer to [**[]GetSingleOrder200ResponseDataPofInfoPofDetailInfoInner**](GetSingleOrder200ResponseDataPofInfoPofDetailInfoInner.md) |  | [optional] 

## Methods

### NewGetSingleOrder200ResponseDataPofInfo

`func NewGetSingleOrder200ResponseDataPofInfo() *GetSingleOrder200ResponseDataPofInfo`

NewGetSingleOrder200ResponseDataPofInfo instantiates a new GetSingleOrder200ResponseDataPofInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSingleOrder200ResponseDataPofInfoWithDefaults

`func NewGetSingleOrder200ResponseDataPofInfoWithDefaults() *GetSingleOrder200ResponseDataPofInfo`

NewGetSingleOrder200ResponseDataPofInfoWithDefaults instantiates a new GetSingleOrder200ResponseDataPofInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *GetSingleOrder200ResponseDataPofInfo) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *GetSingleOrder200ResponseDataPofInfo) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *GetSingleOrder200ResponseDataPofInfo) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *GetSingleOrder200ResponseDataPofInfo) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetPofStatus

`func (o *GetSingleOrder200ResponseDataPofInfo) GetPofStatus() int64`

GetPofStatus returns the PofStatus field if non-nil, zero value otherwise.

### GetPofStatusOk

`func (o *GetSingleOrder200ResponseDataPofInfo) GetPofStatusOk() (*int64, bool)`

GetPofStatusOk returns a tuple with the PofStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPofStatus

`func (o *GetSingleOrder200ResponseDataPofInfo) SetPofStatus(v int64)`

SetPofStatus sets PofStatus field to given value.

### HasPofStatus

`func (o *GetSingleOrder200ResponseDataPofInfo) HasPofStatus() bool`

HasPofStatus returns a boolean if a field has been set.

### GetShippingPrice

`func (o *GetSingleOrder200ResponseDataPofInfo) GetShippingPrice() int64`

GetShippingPrice returns the ShippingPrice field if non-nil, zero value otherwise.

### GetShippingPriceOk

`func (o *GetSingleOrder200ResponseDataPofInfo) GetShippingPriceOk() (*int64, bool)`

GetShippingPriceOk returns a tuple with the ShippingPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingPrice

`func (o *GetSingleOrder200ResponseDataPofInfo) SetShippingPrice(v int64)`

SetShippingPrice sets ShippingPrice field to given value.

### HasShippingPrice

`func (o *GetSingleOrder200ResponseDataPofInfo) HasShippingPrice() bool`

HasShippingPrice returns a boolean if a field has been set.

### GetInsurancePrice

`func (o *GetSingleOrder200ResponseDataPofInfo) GetInsurancePrice() int64`

GetInsurancePrice returns the InsurancePrice field if non-nil, zero value otherwise.

### GetInsurancePriceOk

`func (o *GetSingleOrder200ResponseDataPofInfo) GetInsurancePriceOk() (*int64, bool)`

GetInsurancePriceOk returns a tuple with the InsurancePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsurancePrice

`func (o *GetSingleOrder200ResponseDataPofInfo) SetInsurancePrice(v int64)`

SetInsurancePrice sets InsurancePrice field to given value.

### HasInsurancePrice

`func (o *GetSingleOrder200ResponseDataPofInfo) HasInsurancePrice() bool`

HasInsurancePrice returns a boolean if a field has been set.

### GetCreateTime

`func (o *GetSingleOrder200ResponseDataPofInfo) GetCreateTime() string`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *GetSingleOrder200ResponseDataPofInfo) GetCreateTimeOk() (*string, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *GetSingleOrder200ResponseDataPofInfo) SetCreateTime(v string)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *GetSingleOrder200ResponseDataPofInfo) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCreateBy

`func (o *GetSingleOrder200ResponseDataPofInfo) GetCreateBy() int64`

GetCreateBy returns the CreateBy field if non-nil, zero value otherwise.

### GetCreateByOk

`func (o *GetSingleOrder200ResponseDataPofInfo) GetCreateByOk() (*int64, bool)`

GetCreateByOk returns a tuple with the CreateBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateBy

`func (o *GetSingleOrder200ResponseDataPofInfo) SetCreateBy(v int64)`

SetCreateBy sets CreateBy field to given value.

### HasCreateBy

`func (o *GetSingleOrder200ResponseDataPofInfo) HasCreateBy() bool`

HasCreateBy returns a boolean if a field has been set.

### GetCreateByType

`func (o *GetSingleOrder200ResponseDataPofInfo) GetCreateByType() int64`

GetCreateByType returns the CreateByType field if non-nil, zero value otherwise.

### GetCreateByTypeOk

`func (o *GetSingleOrder200ResponseDataPofInfo) GetCreateByTypeOk() (*int64, bool)`

GetCreateByTypeOk returns a tuple with the CreateByType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateByType

`func (o *GetSingleOrder200ResponseDataPofInfo) SetCreateByType(v int64)`

SetCreateByType sets CreateByType field to given value.

### HasCreateByType

`func (o *GetSingleOrder200ResponseDataPofInfo) HasCreateByType() bool`

HasCreateByType returns a boolean if a field has been set.

### GetUpdateTime

`func (o *GetSingleOrder200ResponseDataPofInfo) GetUpdateTime() string`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *GetSingleOrder200ResponseDataPofInfo) GetUpdateTimeOk() (*string, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *GetSingleOrder200ResponseDataPofInfo) SetUpdateTime(v string)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *GetSingleOrder200ResponseDataPofInfo) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetUpdateBy

`func (o *GetSingleOrder200ResponseDataPofInfo) GetUpdateBy() int64`

GetUpdateBy returns the UpdateBy field if non-nil, zero value otherwise.

### GetUpdateByOk

`func (o *GetSingleOrder200ResponseDataPofInfo) GetUpdateByOk() (*int64, bool)`

GetUpdateByOk returns a tuple with the UpdateBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateBy

`func (o *GetSingleOrder200ResponseDataPofInfo) SetUpdateBy(v int64)`

SetUpdateBy sets UpdateBy field to given value.

### HasUpdateBy

`func (o *GetSingleOrder200ResponseDataPofInfo) HasUpdateBy() bool`

HasUpdateBy returns a boolean if a field has been set.

### GetUpdateByType

`func (o *GetSingleOrder200ResponseDataPofInfo) GetUpdateByType() int64`

GetUpdateByType returns the UpdateByType field if non-nil, zero value otherwise.

### GetUpdateByTypeOk

`func (o *GetSingleOrder200ResponseDataPofInfo) GetUpdateByTypeOk() (*int64, bool)`

GetUpdateByTypeOk returns a tuple with the UpdateByType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateByType

`func (o *GetSingleOrder200ResponseDataPofInfo) SetUpdateByType(v int64)`

SetUpdateByType sets UpdateByType field to given value.

### HasUpdateByType

`func (o *GetSingleOrder200ResponseDataPofInfo) HasUpdateByType() bool`

HasUpdateByType returns a boolean if a field has been set.

### GetPofDetailInfo

`func (o *GetSingleOrder200ResponseDataPofInfo) GetPofDetailInfo() []GetSingleOrder200ResponseDataPofInfoPofDetailInfoInner`

GetPofDetailInfo returns the PofDetailInfo field if non-nil, zero value otherwise.

### GetPofDetailInfoOk

`func (o *GetSingleOrder200ResponseDataPofInfo) GetPofDetailInfoOk() (*[]GetSingleOrder200ResponseDataPofInfoPofDetailInfoInner, bool)`

GetPofDetailInfoOk returns a tuple with the PofDetailInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPofDetailInfo

`func (o *GetSingleOrder200ResponseDataPofInfo) SetPofDetailInfo(v []GetSingleOrder200ResponseDataPofInfoPofDetailInfoInner)`

SetPofDetailInfo sets PofDetailInfo field to given value.

### HasPofDetailInfo

`func (o *GetSingleOrder200ResponseDataPofInfo) HasPofDetailInfo() bool`

HasPofDetailInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


