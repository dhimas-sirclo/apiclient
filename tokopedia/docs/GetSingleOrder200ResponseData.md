# GetSingleOrder200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | Pointer to **int64** | Order Unique Identifier | [optional] 
**BuyerId** | Pointer to **int64** | Buyer User Unique Identifier | [optional] 
**SellerId** | Pointer to **int64** | Seller User Unique Identifier | [optional] 
**PaymentId** | Pointer to **int64** | Payment Unique Identifier | [optional] 
**IsAffiliate** | Pointer to **bool** | Is affiliate? | [optional] 
**IsFulfillment** | Pointer to **bool** | Order Fulfilled by Tokocabang (true), Order Fulfilled by Shop (false) | [optional] 
**OrderWarehouse** | Pointer to [**GetSingleOrder200ResponseDataOrderWarehouse**](GetSingleOrder200ResponseDataOrderWarehouse.md) |  | [optional] 
**OrderStatus** | Pointer to **int64** | Order status code: * 0 - Seller cancel order. * 3 - Order Reject Due Empty Stock. * 5 - Order Canceled by Fraud * 6 - Order Rejected (Auto Cancel Out of Stock) * 10 - Order rejected by seller. * 15 - Instant Cancel by Buyer. * 100 - Order Created. * 103 - Wait for payment confirmation from third party. * 220 - Payment verified, order ready to process. * 221 - Waiting for partner approval. * 400 - Seller accept order. * 450 - Waiting for pickup. * 500 - Order shipment. * 501 - Status changed to waiting resi have no input. * 520 - Invalid shipment reference number (AWB). * 530 - Requested by user to correct invalid entry of shipment reference number. * 540 - Delivered to Pickup Point. * 550 - Return to Seller. * 600 - Order delivered. * 601 - Buyer open a case to finish an order. * 690 - Fraud Review * 700 - Order finished.  | [optional] 
**InvoiceNumber** | Pointer to **string** | Invoice Number | [optional] 
**InvoicePdf** | Pointer to **string** | Invoice PDF Filename | [optional] 
**InvoiceUrl** | Pointer to **string** | Invoice URL | [optional] 
**OpenAmt** | Pointer to **float64** | Total Price Amount | [optional] 
**LpAmt** | Pointer to **float64** |  | [optional] 
**CashbackAmt** | Pointer to **float64** | Cashback Amount | [optional] 
**Info** | Pointer to **string** | Info | [optional] 
**Comment** | Pointer to **string** | Order Comment | [optional] 
**ItemPrice** | Pointer to **float64** | Item Price | [optional] 
**BuyerInfo** | Pointer to [**GetSingleOrder200ResponseDataBuyerInfo**](GetSingleOrder200ResponseDataBuyerInfo.md) |  | [optional] 
**ShopInfo** | Pointer to [**GetSingleOrder200ResponseDataShopInfo**](GetSingleOrder200ResponseDataShopInfo.md) |  | [optional] 
**ShipmentFulfillment** | Pointer to [**GetSingleOrder200ResponseDataShipmentFulfillment**](GetSingleOrder200ResponseDataShipmentFulfillment.md) |  | [optional] 
**Preorder** | Pointer to [**GetSingleOrder200ResponseDataPreorder**](GetSingleOrder200ResponseDataPreorder.md) |  | [optional] 
**OrderInfo** | Pointer to [**GetSingleOrder200ResponseDataOrderInfo**](GetSingleOrder200ResponseDataOrderInfo.md) |  | [optional] 
**OriginInfo** | Pointer to [**GetSingleOrder200ResponseDataOriginInfo**](GetSingleOrder200ResponseDataOriginInfo.md) |  | [optional] 
**PaymentInfo** | Pointer to [**GetSingleOrder200ResponseDataPaymentInfo**](GetSingleOrder200ResponseDataPaymentInfo.md) |  | [optional] 
**InsuranceInfo** | Pointer to [**GetSingleOrder200ResponseDataInsuranceInfo**](GetSingleOrder200ResponseDataInsuranceInfo.md) |  | [optional] 
**HoldInfo** | Pointer to **map[string]interface{}** |  | [optional] 
**CancelRequestInfo** | Pointer to **map[string]interface{}** |  | [optional] 
**CreateTime** | Pointer to **string** | Order Creation Date (format: 2017-07-20T17:50:58.061156Z)  | [optional] 
**ShippingDate** | Pointer to **NullableString** | Order Shipping Date | [optional] 
**UpdateTime** | Pointer to **string** | Order Update Date (format: 2017-07-24T08:01:07.073696Z)  | [optional] 
**PaymentDate** | Pointer to **string** | Order Payment Date (format: 2017-07-20T17:50:58.061156Z)  | [optional] 
**DeliveredDate** | Pointer to **NullableString** | Order Delivered Date | [optional] 
**EstShippingDate** | Pointer to **NullableString** | Estimated Order Shipping Date | [optional] 
**EstDeliveryDate** | Pointer to **NullableString** | Estimated Order Delivered Date | [optional] 
**RelatedInvoices** | Pointer to **map[string]interface{}** | Related Invoices | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** | Custom Fields | [optional] 
**PromoOrderDetail** | Pointer to [**GetSingleOrder200ResponseDataPromoOrderDetail**](GetSingleOrder200ResponseDataPromoOrderDetail.md) |  | [optional] 
**PofInfo** | Pointer to [**GetSingleOrder200ResponseDataPofInfo**](GetSingleOrder200ResponseDataPofInfo.md) |  | [optional] 

## Methods

### NewGetSingleOrder200ResponseData

`func NewGetSingleOrder200ResponseData() *GetSingleOrder200ResponseData`

NewGetSingleOrder200ResponseData instantiates a new GetSingleOrder200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSingleOrder200ResponseDataWithDefaults

`func NewGetSingleOrder200ResponseDataWithDefaults() *GetSingleOrder200ResponseData`

NewGetSingleOrder200ResponseDataWithDefaults instantiates a new GetSingleOrder200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *GetSingleOrder200ResponseData) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *GetSingleOrder200ResponseData) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *GetSingleOrder200ResponseData) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *GetSingleOrder200ResponseData) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetBuyerId

`func (o *GetSingleOrder200ResponseData) GetBuyerId() int64`

GetBuyerId returns the BuyerId field if non-nil, zero value otherwise.

### GetBuyerIdOk

`func (o *GetSingleOrder200ResponseData) GetBuyerIdOk() (*int64, bool)`

GetBuyerIdOk returns a tuple with the BuyerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerId

`func (o *GetSingleOrder200ResponseData) SetBuyerId(v int64)`

SetBuyerId sets BuyerId field to given value.

### HasBuyerId

`func (o *GetSingleOrder200ResponseData) HasBuyerId() bool`

HasBuyerId returns a boolean if a field has been set.

### GetSellerId

`func (o *GetSingleOrder200ResponseData) GetSellerId() int64`

GetSellerId returns the SellerId field if non-nil, zero value otherwise.

### GetSellerIdOk

`func (o *GetSingleOrder200ResponseData) GetSellerIdOk() (*int64, bool)`

GetSellerIdOk returns a tuple with the SellerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerId

`func (o *GetSingleOrder200ResponseData) SetSellerId(v int64)`

SetSellerId sets SellerId field to given value.

### HasSellerId

`func (o *GetSingleOrder200ResponseData) HasSellerId() bool`

HasSellerId returns a boolean if a field has been set.

### GetPaymentId

`func (o *GetSingleOrder200ResponseData) GetPaymentId() int64`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *GetSingleOrder200ResponseData) GetPaymentIdOk() (*int64, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *GetSingleOrder200ResponseData) SetPaymentId(v int64)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *GetSingleOrder200ResponseData) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### GetIsAffiliate

`func (o *GetSingleOrder200ResponseData) GetIsAffiliate() bool`

GetIsAffiliate returns the IsAffiliate field if non-nil, zero value otherwise.

### GetIsAffiliateOk

`func (o *GetSingleOrder200ResponseData) GetIsAffiliateOk() (*bool, bool)`

GetIsAffiliateOk returns a tuple with the IsAffiliate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAffiliate

`func (o *GetSingleOrder200ResponseData) SetIsAffiliate(v bool)`

SetIsAffiliate sets IsAffiliate field to given value.

### HasIsAffiliate

`func (o *GetSingleOrder200ResponseData) HasIsAffiliate() bool`

HasIsAffiliate returns a boolean if a field has been set.

### GetIsFulfillment

`func (o *GetSingleOrder200ResponseData) GetIsFulfillment() bool`

GetIsFulfillment returns the IsFulfillment field if non-nil, zero value otherwise.

### GetIsFulfillmentOk

`func (o *GetSingleOrder200ResponseData) GetIsFulfillmentOk() (*bool, bool)`

GetIsFulfillmentOk returns a tuple with the IsFulfillment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFulfillment

`func (o *GetSingleOrder200ResponseData) SetIsFulfillment(v bool)`

SetIsFulfillment sets IsFulfillment field to given value.

### HasIsFulfillment

`func (o *GetSingleOrder200ResponseData) HasIsFulfillment() bool`

HasIsFulfillment returns a boolean if a field has been set.

### GetOrderWarehouse

`func (o *GetSingleOrder200ResponseData) GetOrderWarehouse() GetSingleOrder200ResponseDataOrderWarehouse`

GetOrderWarehouse returns the OrderWarehouse field if non-nil, zero value otherwise.

### GetOrderWarehouseOk

`func (o *GetSingleOrder200ResponseData) GetOrderWarehouseOk() (*GetSingleOrder200ResponseDataOrderWarehouse, bool)`

GetOrderWarehouseOk returns a tuple with the OrderWarehouse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderWarehouse

`func (o *GetSingleOrder200ResponseData) SetOrderWarehouse(v GetSingleOrder200ResponseDataOrderWarehouse)`

SetOrderWarehouse sets OrderWarehouse field to given value.

### HasOrderWarehouse

`func (o *GetSingleOrder200ResponseData) HasOrderWarehouse() bool`

HasOrderWarehouse returns a boolean if a field has been set.

### GetOrderStatus

`func (o *GetSingleOrder200ResponseData) GetOrderStatus() int64`

GetOrderStatus returns the OrderStatus field if non-nil, zero value otherwise.

### GetOrderStatusOk

`func (o *GetSingleOrder200ResponseData) GetOrderStatusOk() (*int64, bool)`

GetOrderStatusOk returns a tuple with the OrderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderStatus

`func (o *GetSingleOrder200ResponseData) SetOrderStatus(v int64)`

SetOrderStatus sets OrderStatus field to given value.

### HasOrderStatus

`func (o *GetSingleOrder200ResponseData) HasOrderStatus() bool`

HasOrderStatus returns a boolean if a field has been set.

### GetInvoiceNumber

`func (o *GetSingleOrder200ResponseData) GetInvoiceNumber() string`

GetInvoiceNumber returns the InvoiceNumber field if non-nil, zero value otherwise.

### GetInvoiceNumberOk

`func (o *GetSingleOrder200ResponseData) GetInvoiceNumberOk() (*string, bool)`

GetInvoiceNumberOk returns a tuple with the InvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNumber

`func (o *GetSingleOrder200ResponseData) SetInvoiceNumber(v string)`

SetInvoiceNumber sets InvoiceNumber field to given value.

### HasInvoiceNumber

`func (o *GetSingleOrder200ResponseData) HasInvoiceNumber() bool`

HasInvoiceNumber returns a boolean if a field has been set.

### GetInvoicePdf

`func (o *GetSingleOrder200ResponseData) GetInvoicePdf() string`

GetInvoicePdf returns the InvoicePdf field if non-nil, zero value otherwise.

### GetInvoicePdfOk

`func (o *GetSingleOrder200ResponseData) GetInvoicePdfOk() (*string, bool)`

GetInvoicePdfOk returns a tuple with the InvoicePdf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicePdf

`func (o *GetSingleOrder200ResponseData) SetInvoicePdf(v string)`

SetInvoicePdf sets InvoicePdf field to given value.

### HasInvoicePdf

`func (o *GetSingleOrder200ResponseData) HasInvoicePdf() bool`

HasInvoicePdf returns a boolean if a field has been set.

### GetInvoiceUrl

`func (o *GetSingleOrder200ResponseData) GetInvoiceUrl() string`

GetInvoiceUrl returns the InvoiceUrl field if non-nil, zero value otherwise.

### GetInvoiceUrlOk

`func (o *GetSingleOrder200ResponseData) GetInvoiceUrlOk() (*string, bool)`

GetInvoiceUrlOk returns a tuple with the InvoiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceUrl

`func (o *GetSingleOrder200ResponseData) SetInvoiceUrl(v string)`

SetInvoiceUrl sets InvoiceUrl field to given value.

### HasInvoiceUrl

`func (o *GetSingleOrder200ResponseData) HasInvoiceUrl() bool`

HasInvoiceUrl returns a boolean if a field has been set.

### GetOpenAmt

`func (o *GetSingleOrder200ResponseData) GetOpenAmt() float64`

GetOpenAmt returns the OpenAmt field if non-nil, zero value otherwise.

### GetOpenAmtOk

`func (o *GetSingleOrder200ResponseData) GetOpenAmtOk() (*float64, bool)`

GetOpenAmtOk returns a tuple with the OpenAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenAmt

`func (o *GetSingleOrder200ResponseData) SetOpenAmt(v float64)`

SetOpenAmt sets OpenAmt field to given value.

### HasOpenAmt

`func (o *GetSingleOrder200ResponseData) HasOpenAmt() bool`

HasOpenAmt returns a boolean if a field has been set.

### GetLpAmt

`func (o *GetSingleOrder200ResponseData) GetLpAmt() float64`

GetLpAmt returns the LpAmt field if non-nil, zero value otherwise.

### GetLpAmtOk

`func (o *GetSingleOrder200ResponseData) GetLpAmtOk() (*float64, bool)`

GetLpAmtOk returns a tuple with the LpAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLpAmt

`func (o *GetSingleOrder200ResponseData) SetLpAmt(v float64)`

SetLpAmt sets LpAmt field to given value.

### HasLpAmt

`func (o *GetSingleOrder200ResponseData) HasLpAmt() bool`

HasLpAmt returns a boolean if a field has been set.

### GetCashbackAmt

`func (o *GetSingleOrder200ResponseData) GetCashbackAmt() float64`

GetCashbackAmt returns the CashbackAmt field if non-nil, zero value otherwise.

### GetCashbackAmtOk

`func (o *GetSingleOrder200ResponseData) GetCashbackAmtOk() (*float64, bool)`

GetCashbackAmtOk returns a tuple with the CashbackAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashbackAmt

`func (o *GetSingleOrder200ResponseData) SetCashbackAmt(v float64)`

SetCashbackAmt sets CashbackAmt field to given value.

### HasCashbackAmt

`func (o *GetSingleOrder200ResponseData) HasCashbackAmt() bool`

HasCashbackAmt returns a boolean if a field has been set.

### GetInfo

`func (o *GetSingleOrder200ResponseData) GetInfo() string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *GetSingleOrder200ResponseData) GetInfoOk() (*string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *GetSingleOrder200ResponseData) SetInfo(v string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *GetSingleOrder200ResponseData) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetComment

`func (o *GetSingleOrder200ResponseData) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetSingleOrder200ResponseData) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetSingleOrder200ResponseData) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetSingleOrder200ResponseData) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetItemPrice

`func (o *GetSingleOrder200ResponseData) GetItemPrice() float64`

GetItemPrice returns the ItemPrice field if non-nil, zero value otherwise.

### GetItemPriceOk

`func (o *GetSingleOrder200ResponseData) GetItemPriceOk() (*float64, bool)`

GetItemPriceOk returns a tuple with the ItemPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemPrice

`func (o *GetSingleOrder200ResponseData) SetItemPrice(v float64)`

SetItemPrice sets ItemPrice field to given value.

### HasItemPrice

`func (o *GetSingleOrder200ResponseData) HasItemPrice() bool`

HasItemPrice returns a boolean if a field has been set.

### GetBuyerInfo

`func (o *GetSingleOrder200ResponseData) GetBuyerInfo() GetSingleOrder200ResponseDataBuyerInfo`

GetBuyerInfo returns the BuyerInfo field if non-nil, zero value otherwise.

### GetBuyerInfoOk

`func (o *GetSingleOrder200ResponseData) GetBuyerInfoOk() (*GetSingleOrder200ResponseDataBuyerInfo, bool)`

GetBuyerInfoOk returns a tuple with the BuyerInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerInfo

`func (o *GetSingleOrder200ResponseData) SetBuyerInfo(v GetSingleOrder200ResponseDataBuyerInfo)`

SetBuyerInfo sets BuyerInfo field to given value.

### HasBuyerInfo

`func (o *GetSingleOrder200ResponseData) HasBuyerInfo() bool`

HasBuyerInfo returns a boolean if a field has been set.

### GetShopInfo

`func (o *GetSingleOrder200ResponseData) GetShopInfo() GetSingleOrder200ResponseDataShopInfo`

GetShopInfo returns the ShopInfo field if non-nil, zero value otherwise.

### GetShopInfoOk

`func (o *GetSingleOrder200ResponseData) GetShopInfoOk() (*GetSingleOrder200ResponseDataShopInfo, bool)`

GetShopInfoOk returns a tuple with the ShopInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopInfo

`func (o *GetSingleOrder200ResponseData) SetShopInfo(v GetSingleOrder200ResponseDataShopInfo)`

SetShopInfo sets ShopInfo field to given value.

### HasShopInfo

`func (o *GetSingleOrder200ResponseData) HasShopInfo() bool`

HasShopInfo returns a boolean if a field has been set.

### GetShipmentFulfillment

`func (o *GetSingleOrder200ResponseData) GetShipmentFulfillment() GetSingleOrder200ResponseDataShipmentFulfillment`

GetShipmentFulfillment returns the ShipmentFulfillment field if non-nil, zero value otherwise.

### GetShipmentFulfillmentOk

`func (o *GetSingleOrder200ResponseData) GetShipmentFulfillmentOk() (*GetSingleOrder200ResponseDataShipmentFulfillment, bool)`

GetShipmentFulfillmentOk returns a tuple with the ShipmentFulfillment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentFulfillment

`func (o *GetSingleOrder200ResponseData) SetShipmentFulfillment(v GetSingleOrder200ResponseDataShipmentFulfillment)`

SetShipmentFulfillment sets ShipmentFulfillment field to given value.

### HasShipmentFulfillment

`func (o *GetSingleOrder200ResponseData) HasShipmentFulfillment() bool`

HasShipmentFulfillment returns a boolean if a field has been set.

### GetPreorder

`func (o *GetSingleOrder200ResponseData) GetPreorder() GetSingleOrder200ResponseDataPreorder`

GetPreorder returns the Preorder field if non-nil, zero value otherwise.

### GetPreorderOk

`func (o *GetSingleOrder200ResponseData) GetPreorderOk() (*GetSingleOrder200ResponseDataPreorder, bool)`

GetPreorderOk returns a tuple with the Preorder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreorder

`func (o *GetSingleOrder200ResponseData) SetPreorder(v GetSingleOrder200ResponseDataPreorder)`

SetPreorder sets Preorder field to given value.

### HasPreorder

`func (o *GetSingleOrder200ResponseData) HasPreorder() bool`

HasPreorder returns a boolean if a field has been set.

### GetOrderInfo

`func (o *GetSingleOrder200ResponseData) GetOrderInfo() GetSingleOrder200ResponseDataOrderInfo`

GetOrderInfo returns the OrderInfo field if non-nil, zero value otherwise.

### GetOrderInfoOk

`func (o *GetSingleOrder200ResponseData) GetOrderInfoOk() (*GetSingleOrder200ResponseDataOrderInfo, bool)`

GetOrderInfoOk returns a tuple with the OrderInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderInfo

`func (o *GetSingleOrder200ResponseData) SetOrderInfo(v GetSingleOrder200ResponseDataOrderInfo)`

SetOrderInfo sets OrderInfo field to given value.

### HasOrderInfo

`func (o *GetSingleOrder200ResponseData) HasOrderInfo() bool`

HasOrderInfo returns a boolean if a field has been set.

### GetOriginInfo

`func (o *GetSingleOrder200ResponseData) GetOriginInfo() GetSingleOrder200ResponseDataOriginInfo`

GetOriginInfo returns the OriginInfo field if non-nil, zero value otherwise.

### GetOriginInfoOk

`func (o *GetSingleOrder200ResponseData) GetOriginInfoOk() (*GetSingleOrder200ResponseDataOriginInfo, bool)`

GetOriginInfoOk returns a tuple with the OriginInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginInfo

`func (o *GetSingleOrder200ResponseData) SetOriginInfo(v GetSingleOrder200ResponseDataOriginInfo)`

SetOriginInfo sets OriginInfo field to given value.

### HasOriginInfo

`func (o *GetSingleOrder200ResponseData) HasOriginInfo() bool`

HasOriginInfo returns a boolean if a field has been set.

### GetPaymentInfo

`func (o *GetSingleOrder200ResponseData) GetPaymentInfo() GetSingleOrder200ResponseDataPaymentInfo`

GetPaymentInfo returns the PaymentInfo field if non-nil, zero value otherwise.

### GetPaymentInfoOk

`func (o *GetSingleOrder200ResponseData) GetPaymentInfoOk() (*GetSingleOrder200ResponseDataPaymentInfo, bool)`

GetPaymentInfoOk returns a tuple with the PaymentInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentInfo

`func (o *GetSingleOrder200ResponseData) SetPaymentInfo(v GetSingleOrder200ResponseDataPaymentInfo)`

SetPaymentInfo sets PaymentInfo field to given value.

### HasPaymentInfo

`func (o *GetSingleOrder200ResponseData) HasPaymentInfo() bool`

HasPaymentInfo returns a boolean if a field has been set.

### GetInsuranceInfo

`func (o *GetSingleOrder200ResponseData) GetInsuranceInfo() GetSingleOrder200ResponseDataInsuranceInfo`

GetInsuranceInfo returns the InsuranceInfo field if non-nil, zero value otherwise.

### GetInsuranceInfoOk

`func (o *GetSingleOrder200ResponseData) GetInsuranceInfoOk() (*GetSingleOrder200ResponseDataInsuranceInfo, bool)`

GetInsuranceInfoOk returns a tuple with the InsuranceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsuranceInfo

`func (o *GetSingleOrder200ResponseData) SetInsuranceInfo(v GetSingleOrder200ResponseDataInsuranceInfo)`

SetInsuranceInfo sets InsuranceInfo field to given value.

### HasInsuranceInfo

`func (o *GetSingleOrder200ResponseData) HasInsuranceInfo() bool`

HasInsuranceInfo returns a boolean if a field has been set.

### GetHoldInfo

`func (o *GetSingleOrder200ResponseData) GetHoldInfo() map[string]interface{}`

GetHoldInfo returns the HoldInfo field if non-nil, zero value otherwise.

### GetHoldInfoOk

`func (o *GetSingleOrder200ResponseData) GetHoldInfoOk() (*map[string]interface{}, bool)`

GetHoldInfoOk returns a tuple with the HoldInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldInfo

`func (o *GetSingleOrder200ResponseData) SetHoldInfo(v map[string]interface{})`

SetHoldInfo sets HoldInfo field to given value.

### HasHoldInfo

`func (o *GetSingleOrder200ResponseData) HasHoldInfo() bool`

HasHoldInfo returns a boolean if a field has been set.

### SetHoldInfoNil

`func (o *GetSingleOrder200ResponseData) SetHoldInfoNil(b bool)`

 SetHoldInfoNil sets the value for HoldInfo to be an explicit nil

### UnsetHoldInfo
`func (o *GetSingleOrder200ResponseData) UnsetHoldInfo()`

UnsetHoldInfo ensures that no value is present for HoldInfo, not even an explicit nil
### GetCancelRequestInfo

`func (o *GetSingleOrder200ResponseData) GetCancelRequestInfo() map[string]interface{}`

GetCancelRequestInfo returns the CancelRequestInfo field if non-nil, zero value otherwise.

### GetCancelRequestInfoOk

`func (o *GetSingleOrder200ResponseData) GetCancelRequestInfoOk() (*map[string]interface{}, bool)`

GetCancelRequestInfoOk returns a tuple with the CancelRequestInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelRequestInfo

`func (o *GetSingleOrder200ResponseData) SetCancelRequestInfo(v map[string]interface{})`

SetCancelRequestInfo sets CancelRequestInfo field to given value.

### HasCancelRequestInfo

`func (o *GetSingleOrder200ResponseData) HasCancelRequestInfo() bool`

HasCancelRequestInfo returns a boolean if a field has been set.

### SetCancelRequestInfoNil

`func (o *GetSingleOrder200ResponseData) SetCancelRequestInfoNil(b bool)`

 SetCancelRequestInfoNil sets the value for CancelRequestInfo to be an explicit nil

### UnsetCancelRequestInfo
`func (o *GetSingleOrder200ResponseData) UnsetCancelRequestInfo()`

UnsetCancelRequestInfo ensures that no value is present for CancelRequestInfo, not even an explicit nil
### GetCreateTime

`func (o *GetSingleOrder200ResponseData) GetCreateTime() string`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *GetSingleOrder200ResponseData) GetCreateTimeOk() (*string, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *GetSingleOrder200ResponseData) SetCreateTime(v string)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *GetSingleOrder200ResponseData) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetShippingDate

`func (o *GetSingleOrder200ResponseData) GetShippingDate() string`

GetShippingDate returns the ShippingDate field if non-nil, zero value otherwise.

### GetShippingDateOk

`func (o *GetSingleOrder200ResponseData) GetShippingDateOk() (*string, bool)`

GetShippingDateOk returns a tuple with the ShippingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingDate

`func (o *GetSingleOrder200ResponseData) SetShippingDate(v string)`

SetShippingDate sets ShippingDate field to given value.

### HasShippingDate

`func (o *GetSingleOrder200ResponseData) HasShippingDate() bool`

HasShippingDate returns a boolean if a field has been set.

### SetShippingDateNil

`func (o *GetSingleOrder200ResponseData) SetShippingDateNil(b bool)`

 SetShippingDateNil sets the value for ShippingDate to be an explicit nil

### UnsetShippingDate
`func (o *GetSingleOrder200ResponseData) UnsetShippingDate()`

UnsetShippingDate ensures that no value is present for ShippingDate, not even an explicit nil
### GetUpdateTime

`func (o *GetSingleOrder200ResponseData) GetUpdateTime() string`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *GetSingleOrder200ResponseData) GetUpdateTimeOk() (*string, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *GetSingleOrder200ResponseData) SetUpdateTime(v string)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *GetSingleOrder200ResponseData) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetPaymentDate

`func (o *GetSingleOrder200ResponseData) GetPaymentDate() string`

GetPaymentDate returns the PaymentDate field if non-nil, zero value otherwise.

### GetPaymentDateOk

`func (o *GetSingleOrder200ResponseData) GetPaymentDateOk() (*string, bool)`

GetPaymentDateOk returns a tuple with the PaymentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDate

`func (o *GetSingleOrder200ResponseData) SetPaymentDate(v string)`

SetPaymentDate sets PaymentDate field to given value.

### HasPaymentDate

`func (o *GetSingleOrder200ResponseData) HasPaymentDate() bool`

HasPaymentDate returns a boolean if a field has been set.

### GetDeliveredDate

`func (o *GetSingleOrder200ResponseData) GetDeliveredDate() string`

GetDeliveredDate returns the DeliveredDate field if non-nil, zero value otherwise.

### GetDeliveredDateOk

`func (o *GetSingleOrder200ResponseData) GetDeliveredDateOk() (*string, bool)`

GetDeliveredDateOk returns a tuple with the DeliveredDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveredDate

`func (o *GetSingleOrder200ResponseData) SetDeliveredDate(v string)`

SetDeliveredDate sets DeliveredDate field to given value.

### HasDeliveredDate

`func (o *GetSingleOrder200ResponseData) HasDeliveredDate() bool`

HasDeliveredDate returns a boolean if a field has been set.

### SetDeliveredDateNil

`func (o *GetSingleOrder200ResponseData) SetDeliveredDateNil(b bool)`

 SetDeliveredDateNil sets the value for DeliveredDate to be an explicit nil

### UnsetDeliveredDate
`func (o *GetSingleOrder200ResponseData) UnsetDeliveredDate()`

UnsetDeliveredDate ensures that no value is present for DeliveredDate, not even an explicit nil
### GetEstShippingDate

`func (o *GetSingleOrder200ResponseData) GetEstShippingDate() string`

GetEstShippingDate returns the EstShippingDate field if non-nil, zero value otherwise.

### GetEstShippingDateOk

`func (o *GetSingleOrder200ResponseData) GetEstShippingDateOk() (*string, bool)`

GetEstShippingDateOk returns a tuple with the EstShippingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstShippingDate

`func (o *GetSingleOrder200ResponseData) SetEstShippingDate(v string)`

SetEstShippingDate sets EstShippingDate field to given value.

### HasEstShippingDate

`func (o *GetSingleOrder200ResponseData) HasEstShippingDate() bool`

HasEstShippingDate returns a boolean if a field has been set.

### SetEstShippingDateNil

`func (o *GetSingleOrder200ResponseData) SetEstShippingDateNil(b bool)`

 SetEstShippingDateNil sets the value for EstShippingDate to be an explicit nil

### UnsetEstShippingDate
`func (o *GetSingleOrder200ResponseData) UnsetEstShippingDate()`

UnsetEstShippingDate ensures that no value is present for EstShippingDate, not even an explicit nil
### GetEstDeliveryDate

`func (o *GetSingleOrder200ResponseData) GetEstDeliveryDate() string`

GetEstDeliveryDate returns the EstDeliveryDate field if non-nil, zero value otherwise.

### GetEstDeliveryDateOk

`func (o *GetSingleOrder200ResponseData) GetEstDeliveryDateOk() (*string, bool)`

GetEstDeliveryDateOk returns a tuple with the EstDeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstDeliveryDate

`func (o *GetSingleOrder200ResponseData) SetEstDeliveryDate(v string)`

SetEstDeliveryDate sets EstDeliveryDate field to given value.

### HasEstDeliveryDate

`func (o *GetSingleOrder200ResponseData) HasEstDeliveryDate() bool`

HasEstDeliveryDate returns a boolean if a field has been set.

### SetEstDeliveryDateNil

`func (o *GetSingleOrder200ResponseData) SetEstDeliveryDateNil(b bool)`

 SetEstDeliveryDateNil sets the value for EstDeliveryDate to be an explicit nil

### UnsetEstDeliveryDate
`func (o *GetSingleOrder200ResponseData) UnsetEstDeliveryDate()`

UnsetEstDeliveryDate ensures that no value is present for EstDeliveryDate, not even an explicit nil
### GetRelatedInvoices

`func (o *GetSingleOrder200ResponseData) GetRelatedInvoices() map[string]interface{}`

GetRelatedInvoices returns the RelatedInvoices field if non-nil, zero value otherwise.

### GetRelatedInvoicesOk

`func (o *GetSingleOrder200ResponseData) GetRelatedInvoicesOk() (*map[string]interface{}, bool)`

GetRelatedInvoicesOk returns a tuple with the RelatedInvoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedInvoices

`func (o *GetSingleOrder200ResponseData) SetRelatedInvoices(v map[string]interface{})`

SetRelatedInvoices sets RelatedInvoices field to given value.

### HasRelatedInvoices

`func (o *GetSingleOrder200ResponseData) HasRelatedInvoices() bool`

HasRelatedInvoices returns a boolean if a field has been set.

### SetRelatedInvoicesNil

`func (o *GetSingleOrder200ResponseData) SetRelatedInvoicesNil(b bool)`

 SetRelatedInvoicesNil sets the value for RelatedInvoices to be an explicit nil

### UnsetRelatedInvoices
`func (o *GetSingleOrder200ResponseData) UnsetRelatedInvoices()`

UnsetRelatedInvoices ensures that no value is present for RelatedInvoices, not even an explicit nil
### GetCustomFields

`func (o *GetSingleOrder200ResponseData) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *GetSingleOrder200ResponseData) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *GetSingleOrder200ResponseData) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *GetSingleOrder200ResponseData) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *GetSingleOrder200ResponseData) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *GetSingleOrder200ResponseData) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil
### GetPromoOrderDetail

`func (o *GetSingleOrder200ResponseData) GetPromoOrderDetail() GetSingleOrder200ResponseDataPromoOrderDetail`

GetPromoOrderDetail returns the PromoOrderDetail field if non-nil, zero value otherwise.

### GetPromoOrderDetailOk

`func (o *GetSingleOrder200ResponseData) GetPromoOrderDetailOk() (*GetSingleOrder200ResponseDataPromoOrderDetail, bool)`

GetPromoOrderDetailOk returns a tuple with the PromoOrderDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromoOrderDetail

`func (o *GetSingleOrder200ResponseData) SetPromoOrderDetail(v GetSingleOrder200ResponseDataPromoOrderDetail)`

SetPromoOrderDetail sets PromoOrderDetail field to given value.

### HasPromoOrderDetail

`func (o *GetSingleOrder200ResponseData) HasPromoOrderDetail() bool`

HasPromoOrderDetail returns a boolean if a field has been set.

### GetPofInfo

`func (o *GetSingleOrder200ResponseData) GetPofInfo() GetSingleOrder200ResponseDataPofInfo`

GetPofInfo returns the PofInfo field if non-nil, zero value otherwise.

### GetPofInfoOk

`func (o *GetSingleOrder200ResponseData) GetPofInfoOk() (*GetSingleOrder200ResponseDataPofInfo, bool)`

GetPofInfoOk returns a tuple with the PofInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPofInfo

`func (o *GetSingleOrder200ResponseData) SetPofInfo(v GetSingleOrder200ResponseDataPofInfo)`

SetPofInfo sets PofInfo field to given value.

### HasPofInfo

`func (o *GetSingleOrder200ResponseData) HasPofInfo() bool`

HasPofInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


