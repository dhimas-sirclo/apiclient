# GetOrderCobCod200ResponseDataOrderDataInnerOrderHistoryInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderHistId** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **int64** | Order status code: * 0 - Seller cancel order. * 3 - Order Reject Due Empty Stock. * 5 - Order Canceled by Fraud * 6 - Order Rejected (Auto Cancel Out of Stock) * 10 - Order rejected by seller. * 15 - Instant Cancel by Buyer. * 100 - Order Created. * 103 - Wait for payment confirmation from third party. * 220 - Payment verified, order ready to process. * 221 - Waiting for partner approval. * 400 - Seller accept order. * 450 - Waiting for pickup. * 500 - Order shipment. * 501 - Status changed to waiting resi have no input. * 520 - Invalid shipment reference number (AWB). * 530 - Requested by user to correct invalid entry of shipment reference number. * 540 - Delivered to Pickup Point. * 550 - Return to Seller. * 600 - Order delivered. * 601 - Buyer open a case to finish an order. * 690 - Fraud Review * 700 - Order finished.  | [optional] 
**ShippingDate** | Pointer to **int64** |  | [optional] 
**CreateBy** | Pointer to **int64** |  | [optional] 

## Methods

### NewGetOrderCobCod200ResponseDataOrderDataInnerOrderHistoryInner

`func NewGetOrderCobCod200ResponseDataOrderDataInnerOrderHistoryInner() *GetOrderCobCod200ResponseDataOrderDataInnerOrderHistoryInner`

NewGetOrderCobCod200ResponseDataOrderDataInnerOrderHistoryInner instantiates a new GetOrderCobCod200ResponseDataOrderDataInnerOrderHistoryInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrderCobCod200ResponseDataOrderDataInnerOrderHistoryInnerWithDefaults

`func NewGetOrderCobCod200ResponseDataOrderDataInnerOrderHistoryInnerWithDefaults() *GetOrderCobCod200ResponseDataOrderDataInnerOrderHistoryInner`

NewGetOrderCobCod200ResponseDataOrderDataInnerOrderHistoryInnerWithDefaults instantiates a new GetOrderCobCod200ResponseDataOrderDataInnerOrderHistoryInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderHistId

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderHistoryInner) GetOrderHistId() int64`

GetOrderHistId returns the OrderHistId field if non-nil, zero value otherwise.

### GetOrderHistIdOk

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderHistoryInner) GetOrderHistIdOk() (*int64, bool)`

GetOrderHistIdOk returns a tuple with the OrderHistId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderHistId

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderHistoryInner) SetOrderHistId(v int64)`

SetOrderHistId sets OrderHistId field to given value.

### HasOrderHistId

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderHistoryInner) HasOrderHistId() bool`

HasOrderHistId returns a boolean if a field has been set.

### GetStatus

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderHistoryInner) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderHistoryInner) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderHistoryInner) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderHistoryInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetShippingDate

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderHistoryInner) GetShippingDate() int64`

GetShippingDate returns the ShippingDate field if non-nil, zero value otherwise.

### GetShippingDateOk

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderHistoryInner) GetShippingDateOk() (*int64, bool)`

GetShippingDateOk returns a tuple with the ShippingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingDate

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderHistoryInner) SetShippingDate(v int64)`

SetShippingDate sets ShippingDate field to given value.

### HasShippingDate

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderHistoryInner) HasShippingDate() bool`

HasShippingDate returns a boolean if a field has been set.

### GetCreateBy

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderHistoryInner) GetCreateBy() int64`

GetCreateBy returns the CreateBy field if non-nil, zero value otherwise.

### GetCreateByOk

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderHistoryInner) GetCreateByOk() (*int64, bool)`

GetCreateByOk returns a tuple with the CreateBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateBy

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderHistoryInner) SetCreateBy(v int64)`

SetCreateBy sets CreateBy field to given value.

### HasCreateBy

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderHistoryInner) HasCreateBy() bool`

HasCreateBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


