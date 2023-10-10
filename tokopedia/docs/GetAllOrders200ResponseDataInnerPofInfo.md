# GetAllOrders200ResponseDataInnerPofInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | Pointer to **int64** | Order Unique Identifier | [optional] 
**PofStatus** | Pointer to **int64** | Partial Order Fulfillment Status Code | [optional] 
**ShippingPrice** | Pointer to **float64** | Shipping Price | [optional] 
**InsurancePrice** | Pointer to **float64** | Insurance Price | [optional] 
**CreateTime** | Pointer to **string** | POF Request Create Timestamp (format: 2017-07-20T17:50:58.061156Z) | [optional] 
**CreateBy** | Pointer to **int64** | POF Request Create By | [optional] 
**CreateByType** | Pointer to **int64** | POF Request Create By Type | [optional] 
**UpdateTime** | Pointer to **string** | POF Request Update Timestamp (format: 2017-07-20T17:50:58.061156Z) | [optional] 
**UpdateBy** | Pointer to **int64** | POF Request Update By | [optional] 
**UpdateByType** | Pointer to **int64** | POF Request Update By Type | [optional] 

## Methods

### NewGetAllOrders200ResponseDataInnerPofInfo

`func NewGetAllOrders200ResponseDataInnerPofInfo() *GetAllOrders200ResponseDataInnerPofInfo`

NewGetAllOrders200ResponseDataInnerPofInfo instantiates a new GetAllOrders200ResponseDataInnerPofInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllOrders200ResponseDataInnerPofInfoWithDefaults

`func NewGetAllOrders200ResponseDataInnerPofInfoWithDefaults() *GetAllOrders200ResponseDataInnerPofInfo`

NewGetAllOrders200ResponseDataInnerPofInfoWithDefaults instantiates a new GetAllOrders200ResponseDataInnerPofInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *GetAllOrders200ResponseDataInnerPofInfo) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *GetAllOrders200ResponseDataInnerPofInfo) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *GetAllOrders200ResponseDataInnerPofInfo) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *GetAllOrders200ResponseDataInnerPofInfo) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetPofStatus

`func (o *GetAllOrders200ResponseDataInnerPofInfo) GetPofStatus() int64`

GetPofStatus returns the PofStatus field if non-nil, zero value otherwise.

### GetPofStatusOk

`func (o *GetAllOrders200ResponseDataInnerPofInfo) GetPofStatusOk() (*int64, bool)`

GetPofStatusOk returns a tuple with the PofStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPofStatus

`func (o *GetAllOrders200ResponseDataInnerPofInfo) SetPofStatus(v int64)`

SetPofStatus sets PofStatus field to given value.

### HasPofStatus

`func (o *GetAllOrders200ResponseDataInnerPofInfo) HasPofStatus() bool`

HasPofStatus returns a boolean if a field has been set.

### GetShippingPrice

`func (o *GetAllOrders200ResponseDataInnerPofInfo) GetShippingPrice() float64`

GetShippingPrice returns the ShippingPrice field if non-nil, zero value otherwise.

### GetShippingPriceOk

`func (o *GetAllOrders200ResponseDataInnerPofInfo) GetShippingPriceOk() (*float64, bool)`

GetShippingPriceOk returns a tuple with the ShippingPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingPrice

`func (o *GetAllOrders200ResponseDataInnerPofInfo) SetShippingPrice(v float64)`

SetShippingPrice sets ShippingPrice field to given value.

### HasShippingPrice

`func (o *GetAllOrders200ResponseDataInnerPofInfo) HasShippingPrice() bool`

HasShippingPrice returns a boolean if a field has been set.

### GetInsurancePrice

`func (o *GetAllOrders200ResponseDataInnerPofInfo) GetInsurancePrice() float64`

GetInsurancePrice returns the InsurancePrice field if non-nil, zero value otherwise.

### GetInsurancePriceOk

`func (o *GetAllOrders200ResponseDataInnerPofInfo) GetInsurancePriceOk() (*float64, bool)`

GetInsurancePriceOk returns a tuple with the InsurancePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsurancePrice

`func (o *GetAllOrders200ResponseDataInnerPofInfo) SetInsurancePrice(v float64)`

SetInsurancePrice sets InsurancePrice field to given value.

### HasInsurancePrice

`func (o *GetAllOrders200ResponseDataInnerPofInfo) HasInsurancePrice() bool`

HasInsurancePrice returns a boolean if a field has been set.

### GetCreateTime

`func (o *GetAllOrders200ResponseDataInnerPofInfo) GetCreateTime() string`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *GetAllOrders200ResponseDataInnerPofInfo) GetCreateTimeOk() (*string, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *GetAllOrders200ResponseDataInnerPofInfo) SetCreateTime(v string)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *GetAllOrders200ResponseDataInnerPofInfo) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCreateBy

`func (o *GetAllOrders200ResponseDataInnerPofInfo) GetCreateBy() int64`

GetCreateBy returns the CreateBy field if non-nil, zero value otherwise.

### GetCreateByOk

`func (o *GetAllOrders200ResponseDataInnerPofInfo) GetCreateByOk() (*int64, bool)`

GetCreateByOk returns a tuple with the CreateBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateBy

`func (o *GetAllOrders200ResponseDataInnerPofInfo) SetCreateBy(v int64)`

SetCreateBy sets CreateBy field to given value.

### HasCreateBy

`func (o *GetAllOrders200ResponseDataInnerPofInfo) HasCreateBy() bool`

HasCreateBy returns a boolean if a field has been set.

### GetCreateByType

`func (o *GetAllOrders200ResponseDataInnerPofInfo) GetCreateByType() int64`

GetCreateByType returns the CreateByType field if non-nil, zero value otherwise.

### GetCreateByTypeOk

`func (o *GetAllOrders200ResponseDataInnerPofInfo) GetCreateByTypeOk() (*int64, bool)`

GetCreateByTypeOk returns a tuple with the CreateByType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateByType

`func (o *GetAllOrders200ResponseDataInnerPofInfo) SetCreateByType(v int64)`

SetCreateByType sets CreateByType field to given value.

### HasCreateByType

`func (o *GetAllOrders200ResponseDataInnerPofInfo) HasCreateByType() bool`

HasCreateByType returns a boolean if a field has been set.

### GetUpdateTime

`func (o *GetAllOrders200ResponseDataInnerPofInfo) GetUpdateTime() string`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *GetAllOrders200ResponseDataInnerPofInfo) GetUpdateTimeOk() (*string, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *GetAllOrders200ResponseDataInnerPofInfo) SetUpdateTime(v string)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *GetAllOrders200ResponseDataInnerPofInfo) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetUpdateBy

`func (o *GetAllOrders200ResponseDataInnerPofInfo) GetUpdateBy() int64`

GetUpdateBy returns the UpdateBy field if non-nil, zero value otherwise.

### GetUpdateByOk

`func (o *GetAllOrders200ResponseDataInnerPofInfo) GetUpdateByOk() (*int64, bool)`

GetUpdateByOk returns a tuple with the UpdateBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateBy

`func (o *GetAllOrders200ResponseDataInnerPofInfo) SetUpdateBy(v int64)`

SetUpdateBy sets UpdateBy field to given value.

### HasUpdateBy

`func (o *GetAllOrders200ResponseDataInnerPofInfo) HasUpdateBy() bool`

HasUpdateBy returns a boolean if a field has been set.

### GetUpdateByType

`func (o *GetAllOrders200ResponseDataInnerPofInfo) GetUpdateByType() int64`

GetUpdateByType returns the UpdateByType field if non-nil, zero value otherwise.

### GetUpdateByTypeOk

`func (o *GetAllOrders200ResponseDataInnerPofInfo) GetUpdateByTypeOk() (*int64, bool)`

GetUpdateByTypeOk returns a tuple with the UpdateByType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateByType

`func (o *GetAllOrders200ResponseDataInnerPofInfo) SetUpdateByType(v int64)`

SetUpdateByType sets UpdateByType field to given value.

### HasUpdateByType

`func (o *GetAllOrders200ResponseDataInnerPofInfo) HasUpdateByType() bool`

HasUpdateByType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


