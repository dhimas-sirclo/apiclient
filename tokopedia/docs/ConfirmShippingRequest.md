# ConfirmShippingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderStatus** | **int64** | Order status code: * 0 - Seller cancel order. * 3 - Order Reject Due Empty Stock. * 5 - Order Canceled by Fraud * 6 - Order Rejected (Auto Cancel Out of Stock) * 10 - Order rejected by seller. * 15 - Instant Cancel by Buyer. * 100 - Order Created. * 103 - Wait for payment confirmation from third party. * 220 - Payment verified, order ready to process. * 221 - Waiting for partner approval. * 400 - Seller accept order. * 450 - Waiting for pickup. * 500 - Order shipment. * 501 - Status changed to waiting resi have no input. * 520 - Invalid shipment reference number (AWB). * 530 - Requested by user to correct invalid entry of shipment reference number. * 540 - Delivered to Pickup Point. * 550 - Return to Seller. * 600 - Order delivered. * 601 - Buyer open a case to finish an order. * 690 - Fraud Review * 700 - Order finished.  | 
**ShippingRefNum** | **string** |  | 

## Methods

### NewConfirmShippingRequest

`func NewConfirmShippingRequest(orderStatus int64, shippingRefNum string, ) *ConfirmShippingRequest`

NewConfirmShippingRequest instantiates a new ConfirmShippingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfirmShippingRequestWithDefaults

`func NewConfirmShippingRequestWithDefaults() *ConfirmShippingRequest`

NewConfirmShippingRequestWithDefaults instantiates a new ConfirmShippingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderStatus

`func (o *ConfirmShippingRequest) GetOrderStatus() int64`

GetOrderStatus returns the OrderStatus field if non-nil, zero value otherwise.

### GetOrderStatusOk

`func (o *ConfirmShippingRequest) GetOrderStatusOk() (*int64, bool)`

GetOrderStatusOk returns a tuple with the OrderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderStatus

`func (o *ConfirmShippingRequest) SetOrderStatus(v int64)`

SetOrderStatus sets OrderStatus field to given value.


### GetShippingRefNum

`func (o *ConfirmShippingRequest) GetShippingRefNum() string`

GetShippingRefNum returns the ShippingRefNum field if non-nil, zero value otherwise.

### GetShippingRefNumOk

`func (o *ConfirmShippingRequest) GetShippingRefNumOk() (*string, bool)`

GetShippingRefNumOk returns a tuple with the ShippingRefNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingRefNum

`func (o *ConfirmShippingRequest) SetShippingRefNum(v string)`

SetShippingRefNum sets ShippingRefNum field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


