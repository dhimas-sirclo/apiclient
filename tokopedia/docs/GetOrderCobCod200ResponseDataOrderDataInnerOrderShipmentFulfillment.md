# GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**OrderId** | Pointer to **int64** |  | [optional] 
**PaymentDateTime** | Pointer to **string** | format: 2019-07-18T14:58:01.998063Z  | [optional] 
**IsSameDay** | Pointer to **bool** |  | [optional] 
**AcceptDeadline** | Pointer to **string** | format: 2019-07-18T14:58:01.998063Z  | [optional] 
**ConfirmShippingDeadline** | Pointer to **string** | format: 2019-07-18T14:58:01.998063Z  | [optional] 
**ItemDeliveredDeadline** | Pointer to [**GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillmentItemDeliveredDeadline**](GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillmentItemDeliveredDeadline.md) |  | [optional] 
**IsAccepted** | Pointer to **bool** |  | [optional] 
**IsConfirmShipping** | Pointer to **bool** |  | [optional] 
**IsItemDelivered** | Pointer to **bool** |  | [optional] 
**FulfillmentStatus** | Pointer to **int64** |  | [optional] 

## Methods

### NewGetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment

`func NewGetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment() *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment`

NewGetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment instantiates a new GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillmentWithDefaults

`func NewGetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillmentWithDefaults() *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment`

NewGetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillmentWithDefaults instantiates a new GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrderId

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetPaymentDateTime

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) GetPaymentDateTime() string`

GetPaymentDateTime returns the PaymentDateTime field if non-nil, zero value otherwise.

### GetPaymentDateTimeOk

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) GetPaymentDateTimeOk() (*string, bool)`

GetPaymentDateTimeOk returns a tuple with the PaymentDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDateTime

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) SetPaymentDateTime(v string)`

SetPaymentDateTime sets PaymentDateTime field to given value.

### HasPaymentDateTime

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) HasPaymentDateTime() bool`

HasPaymentDateTime returns a boolean if a field has been set.

### GetIsSameDay

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) GetIsSameDay() bool`

GetIsSameDay returns the IsSameDay field if non-nil, zero value otherwise.

### GetIsSameDayOk

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) GetIsSameDayOk() (*bool, bool)`

GetIsSameDayOk returns a tuple with the IsSameDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSameDay

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) SetIsSameDay(v bool)`

SetIsSameDay sets IsSameDay field to given value.

### HasIsSameDay

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) HasIsSameDay() bool`

HasIsSameDay returns a boolean if a field has been set.

### GetAcceptDeadline

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) GetAcceptDeadline() string`

GetAcceptDeadline returns the AcceptDeadline field if non-nil, zero value otherwise.

### GetAcceptDeadlineOk

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) GetAcceptDeadlineOk() (*string, bool)`

GetAcceptDeadlineOk returns a tuple with the AcceptDeadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptDeadline

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) SetAcceptDeadline(v string)`

SetAcceptDeadline sets AcceptDeadline field to given value.

### HasAcceptDeadline

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) HasAcceptDeadline() bool`

HasAcceptDeadline returns a boolean if a field has been set.

### GetConfirmShippingDeadline

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) GetConfirmShippingDeadline() string`

GetConfirmShippingDeadline returns the ConfirmShippingDeadline field if non-nil, zero value otherwise.

### GetConfirmShippingDeadlineOk

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) GetConfirmShippingDeadlineOk() (*string, bool)`

GetConfirmShippingDeadlineOk returns a tuple with the ConfirmShippingDeadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmShippingDeadline

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) SetConfirmShippingDeadline(v string)`

SetConfirmShippingDeadline sets ConfirmShippingDeadline field to given value.

### HasConfirmShippingDeadline

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) HasConfirmShippingDeadline() bool`

HasConfirmShippingDeadline returns a boolean if a field has been set.

### GetItemDeliveredDeadline

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) GetItemDeliveredDeadline() GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillmentItemDeliveredDeadline`

GetItemDeliveredDeadline returns the ItemDeliveredDeadline field if non-nil, zero value otherwise.

### GetItemDeliveredDeadlineOk

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) GetItemDeliveredDeadlineOk() (*GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillmentItemDeliveredDeadline, bool)`

GetItemDeliveredDeadlineOk returns a tuple with the ItemDeliveredDeadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemDeliveredDeadline

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) SetItemDeliveredDeadline(v GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillmentItemDeliveredDeadline)`

SetItemDeliveredDeadline sets ItemDeliveredDeadline field to given value.

### HasItemDeliveredDeadline

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) HasItemDeliveredDeadline() bool`

HasItemDeliveredDeadline returns a boolean if a field has been set.

### GetIsAccepted

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) GetIsAccepted() bool`

GetIsAccepted returns the IsAccepted field if non-nil, zero value otherwise.

### GetIsAcceptedOk

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) GetIsAcceptedOk() (*bool, bool)`

GetIsAcceptedOk returns a tuple with the IsAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAccepted

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) SetIsAccepted(v bool)`

SetIsAccepted sets IsAccepted field to given value.

### HasIsAccepted

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) HasIsAccepted() bool`

HasIsAccepted returns a boolean if a field has been set.

### GetIsConfirmShipping

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) GetIsConfirmShipping() bool`

GetIsConfirmShipping returns the IsConfirmShipping field if non-nil, zero value otherwise.

### GetIsConfirmShippingOk

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) GetIsConfirmShippingOk() (*bool, bool)`

GetIsConfirmShippingOk returns a tuple with the IsConfirmShipping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConfirmShipping

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) SetIsConfirmShipping(v bool)`

SetIsConfirmShipping sets IsConfirmShipping field to given value.

### HasIsConfirmShipping

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) HasIsConfirmShipping() bool`

HasIsConfirmShipping returns a boolean if a field has been set.

### GetIsItemDelivered

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) GetIsItemDelivered() bool`

GetIsItemDelivered returns the IsItemDelivered field if non-nil, zero value otherwise.

### GetIsItemDeliveredOk

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) GetIsItemDeliveredOk() (*bool, bool)`

GetIsItemDeliveredOk returns a tuple with the IsItemDelivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsItemDelivered

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) SetIsItemDelivered(v bool)`

SetIsItemDelivered sets IsItemDelivered field to given value.

### HasIsItemDelivered

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) HasIsItemDelivered() bool`

HasIsItemDelivered returns a boolean if a field has been set.

### GetFulfillmentStatus

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) GetFulfillmentStatus() int64`

GetFulfillmentStatus returns the FulfillmentStatus field if non-nil, zero value otherwise.

### GetFulfillmentStatusOk

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) GetFulfillmentStatusOk() (*int64, bool)`

GetFulfillmentStatusOk returns a tuple with the FulfillmentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentStatus

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) SetFulfillmentStatus(v int64)`

SetFulfillmentStatus sets FulfillmentStatus field to given value.

### HasFulfillmentStatus

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) HasFulfillmentStatus() bool`

HasFulfillmentStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


