# GetOrderCobCod200ResponseDataOrderDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | Pointer to [**GetOrderCobCod200ResponseDataOrderDataInnerOrder**](GetOrderCobCod200ResponseDataOrderDataInnerOrder.md) |  | [optional] 
**OrderHistory** | Pointer to [**[]GetOrderCobCod200ResponseDataOrderDataInnerOrderHistoryInner**](GetOrderCobCod200ResponseDataOrderDataInnerOrderHistoryInner.md) |  | [optional] 
**OrderDetail** | Pointer to [**[]GetOrderCobCod200ResponseDataOrderDataInnerOrderDetailInner**](GetOrderCobCod200ResponseDataOrderDataInnerOrderDetailInner.md) |  | [optional] 
**DropShipper** | Pointer to [**GetOrderCobCod200ResponseDataOrderDataInnerDropShipper**](GetOrderCobCod200ResponseDataOrderDataInnerDropShipper.md) |  | [optional] 
**TypeMeta** | Pointer to [**GetOrderCobCod200ResponseDataOrderDataInnerTypeMeta**](GetOrderCobCod200ResponseDataOrderDataInnerTypeMeta.md) |  | [optional] 
**OrderShipmentFulfillment** | Pointer to [**GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment**](GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment.md) |  | [optional] 
**BookingData** | Pointer to [**GetOrderCobCod200ResponseDataOrderDataInnerBookingData**](GetOrderCobCod200ResponseDataOrderDataInnerBookingData.md) |  | [optional] 

## Methods

### NewGetOrderCobCod200ResponseDataOrderDataInner

`func NewGetOrderCobCod200ResponseDataOrderDataInner() *GetOrderCobCod200ResponseDataOrderDataInner`

NewGetOrderCobCod200ResponseDataOrderDataInner instantiates a new GetOrderCobCod200ResponseDataOrderDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrderCobCod200ResponseDataOrderDataInnerWithDefaults

`func NewGetOrderCobCod200ResponseDataOrderDataInnerWithDefaults() *GetOrderCobCod200ResponseDataOrderDataInner`

NewGetOrderCobCod200ResponseDataOrderDataInnerWithDefaults instantiates a new GetOrderCobCod200ResponseDataOrderDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrder

`func (o *GetOrderCobCod200ResponseDataOrderDataInner) GetOrder() GetOrderCobCod200ResponseDataOrderDataInnerOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *GetOrderCobCod200ResponseDataOrderDataInner) GetOrderOk() (*GetOrderCobCod200ResponseDataOrderDataInnerOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *GetOrderCobCod200ResponseDataOrderDataInner) SetOrder(v GetOrderCobCod200ResponseDataOrderDataInnerOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *GetOrderCobCod200ResponseDataOrderDataInner) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetOrderHistory

`func (o *GetOrderCobCod200ResponseDataOrderDataInner) GetOrderHistory() []GetOrderCobCod200ResponseDataOrderDataInnerOrderHistoryInner`

GetOrderHistory returns the OrderHistory field if non-nil, zero value otherwise.

### GetOrderHistoryOk

`func (o *GetOrderCobCod200ResponseDataOrderDataInner) GetOrderHistoryOk() (*[]GetOrderCobCod200ResponseDataOrderDataInnerOrderHistoryInner, bool)`

GetOrderHistoryOk returns a tuple with the OrderHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderHistory

`func (o *GetOrderCobCod200ResponseDataOrderDataInner) SetOrderHistory(v []GetOrderCobCod200ResponseDataOrderDataInnerOrderHistoryInner)`

SetOrderHistory sets OrderHistory field to given value.

### HasOrderHistory

`func (o *GetOrderCobCod200ResponseDataOrderDataInner) HasOrderHistory() bool`

HasOrderHistory returns a boolean if a field has been set.

### GetOrderDetail

`func (o *GetOrderCobCod200ResponseDataOrderDataInner) GetOrderDetail() []GetOrderCobCod200ResponseDataOrderDataInnerOrderDetailInner`

GetOrderDetail returns the OrderDetail field if non-nil, zero value otherwise.

### GetOrderDetailOk

`func (o *GetOrderCobCod200ResponseDataOrderDataInner) GetOrderDetailOk() (*[]GetOrderCobCod200ResponseDataOrderDataInnerOrderDetailInner, bool)`

GetOrderDetailOk returns a tuple with the OrderDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDetail

`func (o *GetOrderCobCod200ResponseDataOrderDataInner) SetOrderDetail(v []GetOrderCobCod200ResponseDataOrderDataInnerOrderDetailInner)`

SetOrderDetail sets OrderDetail field to given value.

### HasOrderDetail

`func (o *GetOrderCobCod200ResponseDataOrderDataInner) HasOrderDetail() bool`

HasOrderDetail returns a boolean if a field has been set.

### GetDropShipper

`func (o *GetOrderCobCod200ResponseDataOrderDataInner) GetDropShipper() GetOrderCobCod200ResponseDataOrderDataInnerDropShipper`

GetDropShipper returns the DropShipper field if non-nil, zero value otherwise.

### GetDropShipperOk

`func (o *GetOrderCobCod200ResponseDataOrderDataInner) GetDropShipperOk() (*GetOrderCobCod200ResponseDataOrderDataInnerDropShipper, bool)`

GetDropShipperOk returns a tuple with the DropShipper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropShipper

`func (o *GetOrderCobCod200ResponseDataOrderDataInner) SetDropShipper(v GetOrderCobCod200ResponseDataOrderDataInnerDropShipper)`

SetDropShipper sets DropShipper field to given value.

### HasDropShipper

`func (o *GetOrderCobCod200ResponseDataOrderDataInner) HasDropShipper() bool`

HasDropShipper returns a boolean if a field has been set.

### GetTypeMeta

`func (o *GetOrderCobCod200ResponseDataOrderDataInner) GetTypeMeta() GetOrderCobCod200ResponseDataOrderDataInnerTypeMeta`

GetTypeMeta returns the TypeMeta field if non-nil, zero value otherwise.

### GetTypeMetaOk

`func (o *GetOrderCobCod200ResponseDataOrderDataInner) GetTypeMetaOk() (*GetOrderCobCod200ResponseDataOrderDataInnerTypeMeta, bool)`

GetTypeMetaOk returns a tuple with the TypeMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeMeta

`func (o *GetOrderCobCod200ResponseDataOrderDataInner) SetTypeMeta(v GetOrderCobCod200ResponseDataOrderDataInnerTypeMeta)`

SetTypeMeta sets TypeMeta field to given value.

### HasTypeMeta

`func (o *GetOrderCobCod200ResponseDataOrderDataInner) HasTypeMeta() bool`

HasTypeMeta returns a boolean if a field has been set.

### GetOrderShipmentFulfillment

`func (o *GetOrderCobCod200ResponseDataOrderDataInner) GetOrderShipmentFulfillment() GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment`

GetOrderShipmentFulfillment returns the OrderShipmentFulfillment field if non-nil, zero value otherwise.

### GetOrderShipmentFulfillmentOk

`func (o *GetOrderCobCod200ResponseDataOrderDataInner) GetOrderShipmentFulfillmentOk() (*GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment, bool)`

GetOrderShipmentFulfillmentOk returns a tuple with the OrderShipmentFulfillment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderShipmentFulfillment

`func (o *GetOrderCobCod200ResponseDataOrderDataInner) SetOrderShipmentFulfillment(v GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment)`

SetOrderShipmentFulfillment sets OrderShipmentFulfillment field to given value.

### HasOrderShipmentFulfillment

`func (o *GetOrderCobCod200ResponseDataOrderDataInner) HasOrderShipmentFulfillment() bool`

HasOrderShipmentFulfillment returns a boolean if a field has been set.

### GetBookingData

`func (o *GetOrderCobCod200ResponseDataOrderDataInner) GetBookingData() GetOrderCobCod200ResponseDataOrderDataInnerBookingData`

GetBookingData returns the BookingData field if non-nil, zero value otherwise.

### GetBookingDataOk

`func (o *GetOrderCobCod200ResponseDataOrderDataInner) GetBookingDataOk() (*GetOrderCobCod200ResponseDataOrderDataInnerBookingData, bool)`

GetBookingDataOk returns a tuple with the BookingData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookingData

`func (o *GetOrderCobCod200ResponseDataOrderDataInner) SetBookingData(v GetOrderCobCod200ResponseDataOrderDataInnerBookingData)`

SetBookingData sets BookingData field to given value.

### HasBookingData

`func (o *GetOrderCobCod200ResponseDataOrderDataInner) HasBookingData() bool`

HasBookingData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


