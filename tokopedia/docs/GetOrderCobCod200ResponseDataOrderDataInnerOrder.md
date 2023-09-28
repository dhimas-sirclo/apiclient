# GetOrderCobCod200ResponseDataOrderDataInnerOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | Pointer to **int64** |  | [optional] 
**BuyerId** | Pointer to **int64** |  | [optional] 
**SellerId** | Pointer to **int64** |  | [optional] 
**PaymentId** | Pointer to **int64** |  | [optional] 
**OrderStatus** | Pointer to **int64** | Order status code: * 0 - Seller cancel order. * 3 - Order Reject Due Empty Stock. * 5 - Order Canceled by Fraud * 6 - Order Rejected (Auto Cancel Out of Stock) * 10 - Order rejected by seller. * 15 - Instant Cancel by Buyer. * 100 - Order Created. * 103 - Wait for payment confirmation from third party. * 220 - Payment verified, order ready to process. * 221 - Waiting for partner approval. * 400 - Seller accept order. * 450 - Waiting for pickup. * 500 - Order shipment. * 501 - Status changed to waiting resi have no input. * 520 - Invalid shipment reference number (AWB). * 530 - Requested by user to correct invalid entry of shipment reference number. * 540 - Delivered to Pickup Point. * 550 - Return to Seller. * 600 - Order delivered. * 601 - Buyer open a case to finish an order. * 690 - Fraud Review * 700 - Order finished.  | [optional] 
**InvoiceNumber** | Pointer to **string** |  | [optional] 
**InvoicePdfLink** | Pointer to **string** |  | [optional] 
**OpenAmt** | Pointer to **int64** |  | [optional] 
**PaymentAmtCod** | Pointer to **int64** |  | [optional] 

## Methods

### NewGetOrderCobCod200ResponseDataOrderDataInnerOrder

`func NewGetOrderCobCod200ResponseDataOrderDataInnerOrder() *GetOrderCobCod200ResponseDataOrderDataInnerOrder`

NewGetOrderCobCod200ResponseDataOrderDataInnerOrder instantiates a new GetOrderCobCod200ResponseDataOrderDataInnerOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrderCobCod200ResponseDataOrderDataInnerOrderWithDefaults

`func NewGetOrderCobCod200ResponseDataOrderDataInnerOrderWithDefaults() *GetOrderCobCod200ResponseDataOrderDataInnerOrder`

NewGetOrderCobCod200ResponseDataOrderDataInnerOrderWithDefaults instantiates a new GetOrderCobCod200ResponseDataOrderDataInnerOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrder) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrder) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrder) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrder) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetBuyerId

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrder) GetBuyerId() int64`

GetBuyerId returns the BuyerId field if non-nil, zero value otherwise.

### GetBuyerIdOk

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrder) GetBuyerIdOk() (*int64, bool)`

GetBuyerIdOk returns a tuple with the BuyerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerId

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrder) SetBuyerId(v int64)`

SetBuyerId sets BuyerId field to given value.

### HasBuyerId

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrder) HasBuyerId() bool`

HasBuyerId returns a boolean if a field has been set.

### GetSellerId

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrder) GetSellerId() int64`

GetSellerId returns the SellerId field if non-nil, zero value otherwise.

### GetSellerIdOk

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrder) GetSellerIdOk() (*int64, bool)`

GetSellerIdOk returns a tuple with the SellerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerId

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrder) SetSellerId(v int64)`

SetSellerId sets SellerId field to given value.

### HasSellerId

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrder) HasSellerId() bool`

HasSellerId returns a boolean if a field has been set.

### GetPaymentId

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrder) GetPaymentId() int64`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrder) GetPaymentIdOk() (*int64, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrder) SetPaymentId(v int64)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrder) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### GetOrderStatus

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrder) GetOrderStatus() int64`

GetOrderStatus returns the OrderStatus field if non-nil, zero value otherwise.

### GetOrderStatusOk

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrder) GetOrderStatusOk() (*int64, bool)`

GetOrderStatusOk returns a tuple with the OrderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderStatus

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrder) SetOrderStatus(v int64)`

SetOrderStatus sets OrderStatus field to given value.

### HasOrderStatus

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrder) HasOrderStatus() bool`

HasOrderStatus returns a boolean if a field has been set.

### GetInvoiceNumber

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrder) GetInvoiceNumber() string`

GetInvoiceNumber returns the InvoiceNumber field if non-nil, zero value otherwise.

### GetInvoiceNumberOk

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrder) GetInvoiceNumberOk() (*string, bool)`

GetInvoiceNumberOk returns a tuple with the InvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNumber

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrder) SetInvoiceNumber(v string)`

SetInvoiceNumber sets InvoiceNumber field to given value.

### HasInvoiceNumber

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrder) HasInvoiceNumber() bool`

HasInvoiceNumber returns a boolean if a field has been set.

### GetInvoicePdfLink

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrder) GetInvoicePdfLink() string`

GetInvoicePdfLink returns the InvoicePdfLink field if non-nil, zero value otherwise.

### GetInvoicePdfLinkOk

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrder) GetInvoicePdfLinkOk() (*string, bool)`

GetInvoicePdfLinkOk returns a tuple with the InvoicePdfLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicePdfLink

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrder) SetInvoicePdfLink(v string)`

SetInvoicePdfLink sets InvoicePdfLink field to given value.

### HasInvoicePdfLink

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrder) HasInvoicePdfLink() bool`

HasInvoicePdfLink returns a boolean if a field has been set.

### GetOpenAmt

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrder) GetOpenAmt() int64`

GetOpenAmt returns the OpenAmt field if non-nil, zero value otherwise.

### GetOpenAmtOk

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrder) GetOpenAmtOk() (*int64, bool)`

GetOpenAmtOk returns a tuple with the OpenAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenAmt

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrder) SetOpenAmt(v int64)`

SetOpenAmt sets OpenAmt field to given value.

### HasOpenAmt

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrder) HasOpenAmt() bool`

HasOpenAmt returns a boolean if a field has been set.

### GetPaymentAmtCod

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrder) GetPaymentAmtCod() int64`

GetPaymentAmtCod returns the PaymentAmtCod field if non-nil, zero value otherwise.

### GetPaymentAmtCodOk

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrder) GetPaymentAmtCodOk() (*int64, bool)`

GetPaymentAmtCodOk returns a tuple with the PaymentAmtCod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAmtCod

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrder) SetPaymentAmtCod(v int64)`

SetPaymentAmtCod sets PaymentAmtCod field to given value.

### HasPaymentAmtCod

`func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrder) HasPaymentAmtCod() bool`

HasPaymentAmtCod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


