# GetSingleOrder200ResponseDataOrderInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderDetail** | Pointer to [**[]GetSingleOrder200ResponseDataOrderInfoOrderDetailInner**](GetSingleOrder200ResponseDataOrderInfoOrderDetailInner.md) |  | [optional] 
**OrderHistory** | Pointer to [**[]GetSingleOrder200ResponseDataOrderInfoOrderHistoryInner**](GetSingleOrder200ResponseDataOrderInfoOrderHistoryInner.md) |  | [optional] 
**OrderAgeDay** | Pointer to **int64** | Order Age in Day | [optional] 
**ShippingAgeDay** | Pointer to **int64** | Shipping Age in Day | [optional] 
**DeliveredAgeDay** | Pointer to **int64** | Delivered Age in Day | [optional] 
**PartialProcess** | Pointer to **bool** | Is partial process? | [optional] 
**ShippingInfo** | Pointer to [**GetSingleOrder200ResponseDataOrderInfoShippingInfo**](GetSingleOrder200ResponseDataOrderInfoShippingInfo.md) |  | [optional] 
**Destination** | Pointer to [**GetSingleOrder200ResponseDataOrderInfoDestination**](GetSingleOrder200ResponseDataOrderInfoDestination.md) |  | [optional] 
**IsReplacement** | Pointer to **bool** | Is replacement? | [optional] 
**ReplacementMultiplier** | Pointer to **int64** | Replacement Multiplier | [optional] 
**IsPlus** | Pointer to **bool** | Plus Order Flag | [optional] 

## Methods

### NewGetSingleOrder200ResponseDataOrderInfo

`func NewGetSingleOrder200ResponseDataOrderInfo() *GetSingleOrder200ResponseDataOrderInfo`

NewGetSingleOrder200ResponseDataOrderInfo instantiates a new GetSingleOrder200ResponseDataOrderInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSingleOrder200ResponseDataOrderInfoWithDefaults

`func NewGetSingleOrder200ResponseDataOrderInfoWithDefaults() *GetSingleOrder200ResponseDataOrderInfo`

NewGetSingleOrder200ResponseDataOrderInfoWithDefaults instantiates a new GetSingleOrder200ResponseDataOrderInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderDetail

`func (o *GetSingleOrder200ResponseDataOrderInfo) GetOrderDetail() []GetSingleOrder200ResponseDataOrderInfoOrderDetailInner`

GetOrderDetail returns the OrderDetail field if non-nil, zero value otherwise.

### GetOrderDetailOk

`func (o *GetSingleOrder200ResponseDataOrderInfo) GetOrderDetailOk() (*[]GetSingleOrder200ResponseDataOrderInfoOrderDetailInner, bool)`

GetOrderDetailOk returns a tuple with the OrderDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDetail

`func (o *GetSingleOrder200ResponseDataOrderInfo) SetOrderDetail(v []GetSingleOrder200ResponseDataOrderInfoOrderDetailInner)`

SetOrderDetail sets OrderDetail field to given value.

### HasOrderDetail

`func (o *GetSingleOrder200ResponseDataOrderInfo) HasOrderDetail() bool`

HasOrderDetail returns a boolean if a field has been set.

### GetOrderHistory

`func (o *GetSingleOrder200ResponseDataOrderInfo) GetOrderHistory() []GetSingleOrder200ResponseDataOrderInfoOrderHistoryInner`

GetOrderHistory returns the OrderHistory field if non-nil, zero value otherwise.

### GetOrderHistoryOk

`func (o *GetSingleOrder200ResponseDataOrderInfo) GetOrderHistoryOk() (*[]GetSingleOrder200ResponseDataOrderInfoOrderHistoryInner, bool)`

GetOrderHistoryOk returns a tuple with the OrderHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderHistory

`func (o *GetSingleOrder200ResponseDataOrderInfo) SetOrderHistory(v []GetSingleOrder200ResponseDataOrderInfoOrderHistoryInner)`

SetOrderHistory sets OrderHistory field to given value.

### HasOrderHistory

`func (o *GetSingleOrder200ResponseDataOrderInfo) HasOrderHistory() bool`

HasOrderHistory returns a boolean if a field has been set.

### GetOrderAgeDay

`func (o *GetSingleOrder200ResponseDataOrderInfo) GetOrderAgeDay() int64`

GetOrderAgeDay returns the OrderAgeDay field if non-nil, zero value otherwise.

### GetOrderAgeDayOk

`func (o *GetSingleOrder200ResponseDataOrderInfo) GetOrderAgeDayOk() (*int64, bool)`

GetOrderAgeDayOk returns a tuple with the OrderAgeDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderAgeDay

`func (o *GetSingleOrder200ResponseDataOrderInfo) SetOrderAgeDay(v int64)`

SetOrderAgeDay sets OrderAgeDay field to given value.

### HasOrderAgeDay

`func (o *GetSingleOrder200ResponseDataOrderInfo) HasOrderAgeDay() bool`

HasOrderAgeDay returns a boolean if a field has been set.

### GetShippingAgeDay

`func (o *GetSingleOrder200ResponseDataOrderInfo) GetShippingAgeDay() int64`

GetShippingAgeDay returns the ShippingAgeDay field if non-nil, zero value otherwise.

### GetShippingAgeDayOk

`func (o *GetSingleOrder200ResponseDataOrderInfo) GetShippingAgeDayOk() (*int64, bool)`

GetShippingAgeDayOk returns a tuple with the ShippingAgeDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAgeDay

`func (o *GetSingleOrder200ResponseDataOrderInfo) SetShippingAgeDay(v int64)`

SetShippingAgeDay sets ShippingAgeDay field to given value.

### HasShippingAgeDay

`func (o *GetSingleOrder200ResponseDataOrderInfo) HasShippingAgeDay() bool`

HasShippingAgeDay returns a boolean if a field has been set.

### GetDeliveredAgeDay

`func (o *GetSingleOrder200ResponseDataOrderInfo) GetDeliveredAgeDay() int64`

GetDeliveredAgeDay returns the DeliveredAgeDay field if non-nil, zero value otherwise.

### GetDeliveredAgeDayOk

`func (o *GetSingleOrder200ResponseDataOrderInfo) GetDeliveredAgeDayOk() (*int64, bool)`

GetDeliveredAgeDayOk returns a tuple with the DeliveredAgeDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveredAgeDay

`func (o *GetSingleOrder200ResponseDataOrderInfo) SetDeliveredAgeDay(v int64)`

SetDeliveredAgeDay sets DeliveredAgeDay field to given value.

### HasDeliveredAgeDay

`func (o *GetSingleOrder200ResponseDataOrderInfo) HasDeliveredAgeDay() bool`

HasDeliveredAgeDay returns a boolean if a field has been set.

### GetPartialProcess

`func (o *GetSingleOrder200ResponseDataOrderInfo) GetPartialProcess() bool`

GetPartialProcess returns the PartialProcess field if non-nil, zero value otherwise.

### GetPartialProcessOk

`func (o *GetSingleOrder200ResponseDataOrderInfo) GetPartialProcessOk() (*bool, bool)`

GetPartialProcessOk returns a tuple with the PartialProcess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartialProcess

`func (o *GetSingleOrder200ResponseDataOrderInfo) SetPartialProcess(v bool)`

SetPartialProcess sets PartialProcess field to given value.

### HasPartialProcess

`func (o *GetSingleOrder200ResponseDataOrderInfo) HasPartialProcess() bool`

HasPartialProcess returns a boolean if a field has been set.

### GetShippingInfo

`func (o *GetSingleOrder200ResponseDataOrderInfo) GetShippingInfo() GetSingleOrder200ResponseDataOrderInfoShippingInfo`

GetShippingInfo returns the ShippingInfo field if non-nil, zero value otherwise.

### GetShippingInfoOk

`func (o *GetSingleOrder200ResponseDataOrderInfo) GetShippingInfoOk() (*GetSingleOrder200ResponseDataOrderInfoShippingInfo, bool)`

GetShippingInfoOk returns a tuple with the ShippingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingInfo

`func (o *GetSingleOrder200ResponseDataOrderInfo) SetShippingInfo(v GetSingleOrder200ResponseDataOrderInfoShippingInfo)`

SetShippingInfo sets ShippingInfo field to given value.

### HasShippingInfo

`func (o *GetSingleOrder200ResponseDataOrderInfo) HasShippingInfo() bool`

HasShippingInfo returns a boolean if a field has been set.

### GetDestination

`func (o *GetSingleOrder200ResponseDataOrderInfo) GetDestination() GetSingleOrder200ResponseDataOrderInfoDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *GetSingleOrder200ResponseDataOrderInfo) GetDestinationOk() (*GetSingleOrder200ResponseDataOrderInfoDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *GetSingleOrder200ResponseDataOrderInfo) SetDestination(v GetSingleOrder200ResponseDataOrderInfoDestination)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *GetSingleOrder200ResponseDataOrderInfo) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetIsReplacement

`func (o *GetSingleOrder200ResponseDataOrderInfo) GetIsReplacement() bool`

GetIsReplacement returns the IsReplacement field if non-nil, zero value otherwise.

### GetIsReplacementOk

`func (o *GetSingleOrder200ResponseDataOrderInfo) GetIsReplacementOk() (*bool, bool)`

GetIsReplacementOk returns a tuple with the IsReplacement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReplacement

`func (o *GetSingleOrder200ResponseDataOrderInfo) SetIsReplacement(v bool)`

SetIsReplacement sets IsReplacement field to given value.

### HasIsReplacement

`func (o *GetSingleOrder200ResponseDataOrderInfo) HasIsReplacement() bool`

HasIsReplacement returns a boolean if a field has been set.

### GetReplacementMultiplier

`func (o *GetSingleOrder200ResponseDataOrderInfo) GetReplacementMultiplier() int64`

GetReplacementMultiplier returns the ReplacementMultiplier field if non-nil, zero value otherwise.

### GetReplacementMultiplierOk

`func (o *GetSingleOrder200ResponseDataOrderInfo) GetReplacementMultiplierOk() (*int64, bool)`

GetReplacementMultiplierOk returns a tuple with the ReplacementMultiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacementMultiplier

`func (o *GetSingleOrder200ResponseDataOrderInfo) SetReplacementMultiplier(v int64)`

SetReplacementMultiplier sets ReplacementMultiplier field to given value.

### HasReplacementMultiplier

`func (o *GetSingleOrder200ResponseDataOrderInfo) HasReplacementMultiplier() bool`

HasReplacementMultiplier returns a boolean if a field has been set.

### GetIsPlus

`func (o *GetSingleOrder200ResponseDataOrderInfo) GetIsPlus() bool`

GetIsPlus returns the IsPlus field if non-nil, zero value otherwise.

### GetIsPlusOk

`func (o *GetSingleOrder200ResponseDataOrderInfo) GetIsPlusOk() (*bool, bool)`

GetIsPlusOk returns a tuple with the IsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPlus

`func (o *GetSingleOrder200ResponseDataOrderInfo) SetIsPlus(v bool)`

SetIsPlus sets IsPlus field to given value.

### HasIsPlus

`func (o *GetSingleOrder200ResponseDataOrderInfo) HasIsPlus() bool`

HasIsPlus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


