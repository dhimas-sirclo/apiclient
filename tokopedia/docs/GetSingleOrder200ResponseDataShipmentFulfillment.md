# GetSingleOrder200ResponseDataShipmentFulfillment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Shipment Fulfillment Unique Identifier | [optional] 
**OrderId** | Pointer to **int64** | Order Unique Identifier | [optional] 
**PaymentDateTime** | Pointer to **string** | Payment Date Time (format: 0001-01-01T00:00:00Z)  | [optional] 
**IsSameDay** | Pointer to **bool** | Is same day delivery? | [optional] 
**AcceptDeadline** | Pointer to **string** | Accept Deadline (format: 0001-01-01T00:00:00Z)  | [optional] 
**ConfirmShippingDeadline** | Pointer to **string** | Accept Deadline (format: 0001-01-01T00:00:00Z)  | [optional] 
**ItemDeliveredDeadline** | Pointer to [**GetSingleOrder200ResponseDataShipmentFulfillmentItemDeliveredDeadline**](GetSingleOrder200ResponseDataShipmentFulfillmentItemDeliveredDeadline.md) |  | [optional] 
**IsAccepted** | Pointer to **bool** | Is order accepted? | [optional] 
**IsConfirmShipping** | Pointer to **bool** | Is shipping confirmed? | [optional] 
**IsItemDelivered** | Pointer to **bool** | Is item delivered? | [optional] 
**FulfillmentStatus** | Pointer to **int64** | Fullfilment Status | [optional] 

## Methods

### NewGetSingleOrder200ResponseDataShipmentFulfillment

`func NewGetSingleOrder200ResponseDataShipmentFulfillment() *GetSingleOrder200ResponseDataShipmentFulfillment`

NewGetSingleOrder200ResponseDataShipmentFulfillment instantiates a new GetSingleOrder200ResponseDataShipmentFulfillment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSingleOrder200ResponseDataShipmentFulfillmentWithDefaults

`func NewGetSingleOrder200ResponseDataShipmentFulfillmentWithDefaults() *GetSingleOrder200ResponseDataShipmentFulfillment`

NewGetSingleOrder200ResponseDataShipmentFulfillmentWithDefaults instantiates a new GetSingleOrder200ResponseDataShipmentFulfillment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetSingleOrder200ResponseDataShipmentFulfillment) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetSingleOrder200ResponseDataShipmentFulfillment) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetSingleOrder200ResponseDataShipmentFulfillment) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetSingleOrder200ResponseDataShipmentFulfillment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrderId

`func (o *GetSingleOrder200ResponseDataShipmentFulfillment) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *GetSingleOrder200ResponseDataShipmentFulfillment) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *GetSingleOrder200ResponseDataShipmentFulfillment) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *GetSingleOrder200ResponseDataShipmentFulfillment) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetPaymentDateTime

`func (o *GetSingleOrder200ResponseDataShipmentFulfillment) GetPaymentDateTime() string`

GetPaymentDateTime returns the PaymentDateTime field if non-nil, zero value otherwise.

### GetPaymentDateTimeOk

`func (o *GetSingleOrder200ResponseDataShipmentFulfillment) GetPaymentDateTimeOk() (*string, bool)`

GetPaymentDateTimeOk returns a tuple with the PaymentDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDateTime

`func (o *GetSingleOrder200ResponseDataShipmentFulfillment) SetPaymentDateTime(v string)`

SetPaymentDateTime sets PaymentDateTime field to given value.

### HasPaymentDateTime

`func (o *GetSingleOrder200ResponseDataShipmentFulfillment) HasPaymentDateTime() bool`

HasPaymentDateTime returns a boolean if a field has been set.

### GetIsSameDay

`func (o *GetSingleOrder200ResponseDataShipmentFulfillment) GetIsSameDay() bool`

GetIsSameDay returns the IsSameDay field if non-nil, zero value otherwise.

### GetIsSameDayOk

`func (o *GetSingleOrder200ResponseDataShipmentFulfillment) GetIsSameDayOk() (*bool, bool)`

GetIsSameDayOk returns a tuple with the IsSameDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSameDay

`func (o *GetSingleOrder200ResponseDataShipmentFulfillment) SetIsSameDay(v bool)`

SetIsSameDay sets IsSameDay field to given value.

### HasIsSameDay

`func (o *GetSingleOrder200ResponseDataShipmentFulfillment) HasIsSameDay() bool`

HasIsSameDay returns a boolean if a field has been set.

### GetAcceptDeadline

`func (o *GetSingleOrder200ResponseDataShipmentFulfillment) GetAcceptDeadline() string`

GetAcceptDeadline returns the AcceptDeadline field if non-nil, zero value otherwise.

### GetAcceptDeadlineOk

`func (o *GetSingleOrder200ResponseDataShipmentFulfillment) GetAcceptDeadlineOk() (*string, bool)`

GetAcceptDeadlineOk returns a tuple with the AcceptDeadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptDeadline

`func (o *GetSingleOrder200ResponseDataShipmentFulfillment) SetAcceptDeadline(v string)`

SetAcceptDeadline sets AcceptDeadline field to given value.

### HasAcceptDeadline

`func (o *GetSingleOrder200ResponseDataShipmentFulfillment) HasAcceptDeadline() bool`

HasAcceptDeadline returns a boolean if a field has been set.

### GetConfirmShippingDeadline

`func (o *GetSingleOrder200ResponseDataShipmentFulfillment) GetConfirmShippingDeadline() string`

GetConfirmShippingDeadline returns the ConfirmShippingDeadline field if non-nil, zero value otherwise.

### GetConfirmShippingDeadlineOk

`func (o *GetSingleOrder200ResponseDataShipmentFulfillment) GetConfirmShippingDeadlineOk() (*string, bool)`

GetConfirmShippingDeadlineOk returns a tuple with the ConfirmShippingDeadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmShippingDeadline

`func (o *GetSingleOrder200ResponseDataShipmentFulfillment) SetConfirmShippingDeadline(v string)`

SetConfirmShippingDeadline sets ConfirmShippingDeadline field to given value.

### HasConfirmShippingDeadline

`func (o *GetSingleOrder200ResponseDataShipmentFulfillment) HasConfirmShippingDeadline() bool`

HasConfirmShippingDeadline returns a boolean if a field has been set.

### GetItemDeliveredDeadline

`func (o *GetSingleOrder200ResponseDataShipmentFulfillment) GetItemDeliveredDeadline() GetSingleOrder200ResponseDataShipmentFulfillmentItemDeliveredDeadline`

GetItemDeliveredDeadline returns the ItemDeliveredDeadline field if non-nil, zero value otherwise.

### GetItemDeliveredDeadlineOk

`func (o *GetSingleOrder200ResponseDataShipmentFulfillment) GetItemDeliveredDeadlineOk() (*GetSingleOrder200ResponseDataShipmentFulfillmentItemDeliveredDeadline, bool)`

GetItemDeliveredDeadlineOk returns a tuple with the ItemDeliveredDeadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemDeliveredDeadline

`func (o *GetSingleOrder200ResponseDataShipmentFulfillment) SetItemDeliveredDeadline(v GetSingleOrder200ResponseDataShipmentFulfillmentItemDeliveredDeadline)`

SetItemDeliveredDeadline sets ItemDeliveredDeadline field to given value.

### HasItemDeliveredDeadline

`func (o *GetSingleOrder200ResponseDataShipmentFulfillment) HasItemDeliveredDeadline() bool`

HasItemDeliveredDeadline returns a boolean if a field has been set.

### GetIsAccepted

`func (o *GetSingleOrder200ResponseDataShipmentFulfillment) GetIsAccepted() bool`

GetIsAccepted returns the IsAccepted field if non-nil, zero value otherwise.

### GetIsAcceptedOk

`func (o *GetSingleOrder200ResponseDataShipmentFulfillment) GetIsAcceptedOk() (*bool, bool)`

GetIsAcceptedOk returns a tuple with the IsAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAccepted

`func (o *GetSingleOrder200ResponseDataShipmentFulfillment) SetIsAccepted(v bool)`

SetIsAccepted sets IsAccepted field to given value.

### HasIsAccepted

`func (o *GetSingleOrder200ResponseDataShipmentFulfillment) HasIsAccepted() bool`

HasIsAccepted returns a boolean if a field has been set.

### GetIsConfirmShipping

`func (o *GetSingleOrder200ResponseDataShipmentFulfillment) GetIsConfirmShipping() bool`

GetIsConfirmShipping returns the IsConfirmShipping field if non-nil, zero value otherwise.

### GetIsConfirmShippingOk

`func (o *GetSingleOrder200ResponseDataShipmentFulfillment) GetIsConfirmShippingOk() (*bool, bool)`

GetIsConfirmShippingOk returns a tuple with the IsConfirmShipping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConfirmShipping

`func (o *GetSingleOrder200ResponseDataShipmentFulfillment) SetIsConfirmShipping(v bool)`

SetIsConfirmShipping sets IsConfirmShipping field to given value.

### HasIsConfirmShipping

`func (o *GetSingleOrder200ResponseDataShipmentFulfillment) HasIsConfirmShipping() bool`

HasIsConfirmShipping returns a boolean if a field has been set.

### GetIsItemDelivered

`func (o *GetSingleOrder200ResponseDataShipmentFulfillment) GetIsItemDelivered() bool`

GetIsItemDelivered returns the IsItemDelivered field if non-nil, zero value otherwise.

### GetIsItemDeliveredOk

`func (o *GetSingleOrder200ResponseDataShipmentFulfillment) GetIsItemDeliveredOk() (*bool, bool)`

GetIsItemDeliveredOk returns a tuple with the IsItemDelivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsItemDelivered

`func (o *GetSingleOrder200ResponseDataShipmentFulfillment) SetIsItemDelivered(v bool)`

SetIsItemDelivered sets IsItemDelivered field to given value.

### HasIsItemDelivered

`func (o *GetSingleOrder200ResponseDataShipmentFulfillment) HasIsItemDelivered() bool`

HasIsItemDelivered returns a boolean if a field has been set.

### GetFulfillmentStatus

`func (o *GetSingleOrder200ResponseDataShipmentFulfillment) GetFulfillmentStatus() int64`

GetFulfillmentStatus returns the FulfillmentStatus field if non-nil, zero value otherwise.

### GetFulfillmentStatusOk

`func (o *GetSingleOrder200ResponseDataShipmentFulfillment) GetFulfillmentStatusOk() (*int64, bool)`

GetFulfillmentStatusOk returns a tuple with the FulfillmentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentStatus

`func (o *GetSingleOrder200ResponseDataShipmentFulfillment) SetFulfillmentStatus(v int64)`

SetFulfillmentStatus sets FulfillmentStatus field to given value.

### HasFulfillmentStatus

`func (o *GetSingleOrder200ResponseDataShipmentFulfillment) HasFulfillmentStatus() bool`

HasFulfillmentStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


