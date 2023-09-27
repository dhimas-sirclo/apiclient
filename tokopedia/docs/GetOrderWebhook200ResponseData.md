# GetOrderWebhook200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FsId** | Pointer to **int64** |  | [optional] 
**OrderStatus** | Pointer to **int64** | Order status code: * 0 - Seller cancel order. * 3 - Order Reject Due Empty Stock. * 5 - Order Canceled by Fraud * 6 - Order Rejected (Auto Cancel Out of Stock) * 10 - Order rejected by seller. * 15 - Instant Cancel by Buyer. * 100 - Order Created. * 103 - Wait for payment confirmation from third party. * 220 - Payment verified, order ready to process. * 221 - Waiting for partner approval. * 400 - Seller accept order. * 450 - Waiting for pickup. * 500 - Order shipment. * 501 - Status changed to waiting resi have no input. * 520 - Invalid shipment reference number (AWB). * 530 - Requested by user to correct invalid entry of shipment reference number. * 540 - Delivered to Pickup Point. * 550 - Return to Seller. * 600 - Order delivered. * 601 - Buyer open a case to finish an order. * 690 - Fraud Review * 700 - Order finished.  | [optional] 
**OrderId** | Pointer to **int64** |  | [optional] 
**ShopId** | Pointer to **int64** |  | [optional] 
**WarehouseId** | Pointer to **int64** |  | [optional] 
**ProductDetails** | Pointer to [**[]GetOrderWebhook200ResponseDataOneOfProductDetailsInner**](GetOrderWebhook200ResponseDataOneOfProductDetailsInner.md) |  | [optional] 
**InvoiceRefNum** | Pointer to **string** |  | [optional] 
**Products** | Pointer to [**[]GetOrderWebhook200ResponseDataOneOfProductDetailsInner**](GetOrderWebhook200ResponseDataOneOfProductDetailsInner.md) |  | [optional] 
**Customer** | Pointer to [**GetOrderWebhook200ResponseDataOneOf1Customer**](GetOrderWebhook200ResponseDataOneOf1Customer.md) |  | [optional] 
**Recipient** | Pointer to [**GetOrderWebhook200ResponseDataOneOf1Recipient**](GetOrderWebhook200ResponseDataOneOf1Recipient.md) |  | [optional] 
**ShopName** | Pointer to **string** |  | [optional] 
**PaymentId** | Pointer to **int64** |  | [optional] 
**PaymentDate** | Pointer to **string** |  | [optional] 
**Logistics** | Pointer to [**GetOrderWebhook200ResponseDataOneOf1Logistics**](GetOrderWebhook200ResponseDataOneOf1Logistics.md) |  | [optional] 
**Amt** | Pointer to [**GetOrderWebhook200ResponseDataOneOf1Amt**](GetOrderWebhook200ResponseDataOneOf1Amt.md) |  | [optional] 
**DropshipperInfo** | Pointer to [**GetOrderWebhook200ResponseDataOneOf1DropshipperInfo**](GetOrderWebhook200ResponseDataOneOf1DropshipperInfo.md) |  | [optional] 
**VoucherInfo** | Pointer to [**GetOrderWebhook200ResponseDataOneOf1VoucherInfo**](GetOrderWebhook200ResponseDataOneOf1VoucherInfo.md) |  | [optional] 
**DeviceType** | Pointer to **string** |  | [optional] 
**CreateTime** | Pointer to **string** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**AcceptPartial** | Pointer to **bool** |  | [optional] 

## Methods

### NewGetOrderWebhook200ResponseData

`func NewGetOrderWebhook200ResponseData() *GetOrderWebhook200ResponseData`

NewGetOrderWebhook200ResponseData instantiates a new GetOrderWebhook200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrderWebhook200ResponseDataWithDefaults

`func NewGetOrderWebhook200ResponseDataWithDefaults() *GetOrderWebhook200ResponseData`

NewGetOrderWebhook200ResponseDataWithDefaults instantiates a new GetOrderWebhook200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFsId

`func (o *GetOrderWebhook200ResponseData) GetFsId() int64`

GetFsId returns the FsId field if non-nil, zero value otherwise.

### GetFsIdOk

`func (o *GetOrderWebhook200ResponseData) GetFsIdOk() (*int64, bool)`

GetFsIdOk returns a tuple with the FsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsId

`func (o *GetOrderWebhook200ResponseData) SetFsId(v int64)`

SetFsId sets FsId field to given value.

### HasFsId

`func (o *GetOrderWebhook200ResponseData) HasFsId() bool`

HasFsId returns a boolean if a field has been set.

### GetOrderStatus

`func (o *GetOrderWebhook200ResponseData) GetOrderStatus() int64`

GetOrderStatus returns the OrderStatus field if non-nil, zero value otherwise.

### GetOrderStatusOk

`func (o *GetOrderWebhook200ResponseData) GetOrderStatusOk() (*int64, bool)`

GetOrderStatusOk returns a tuple with the OrderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderStatus

`func (o *GetOrderWebhook200ResponseData) SetOrderStatus(v int64)`

SetOrderStatus sets OrderStatus field to given value.

### HasOrderStatus

`func (o *GetOrderWebhook200ResponseData) HasOrderStatus() bool`

HasOrderStatus returns a boolean if a field has been set.

### GetOrderId

`func (o *GetOrderWebhook200ResponseData) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *GetOrderWebhook200ResponseData) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *GetOrderWebhook200ResponseData) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *GetOrderWebhook200ResponseData) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetShopId

`func (o *GetOrderWebhook200ResponseData) GetShopId() int64`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *GetOrderWebhook200ResponseData) GetShopIdOk() (*int64, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *GetOrderWebhook200ResponseData) SetShopId(v int64)`

SetShopId sets ShopId field to given value.

### HasShopId

`func (o *GetOrderWebhook200ResponseData) HasShopId() bool`

HasShopId returns a boolean if a field has been set.

### GetWarehouseId

`func (o *GetOrderWebhook200ResponseData) GetWarehouseId() int64`

GetWarehouseId returns the WarehouseId field if non-nil, zero value otherwise.

### GetWarehouseIdOk

`func (o *GetOrderWebhook200ResponseData) GetWarehouseIdOk() (*int64, bool)`

GetWarehouseIdOk returns a tuple with the WarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseId

`func (o *GetOrderWebhook200ResponseData) SetWarehouseId(v int64)`

SetWarehouseId sets WarehouseId field to given value.

### HasWarehouseId

`func (o *GetOrderWebhook200ResponseData) HasWarehouseId() bool`

HasWarehouseId returns a boolean if a field has been set.

### GetProductDetails

`func (o *GetOrderWebhook200ResponseData) GetProductDetails() []GetOrderWebhook200ResponseDataOneOfProductDetailsInner`

GetProductDetails returns the ProductDetails field if non-nil, zero value otherwise.

### GetProductDetailsOk

`func (o *GetOrderWebhook200ResponseData) GetProductDetailsOk() (*[]GetOrderWebhook200ResponseDataOneOfProductDetailsInner, bool)`

GetProductDetailsOk returns a tuple with the ProductDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductDetails

`func (o *GetOrderWebhook200ResponseData) SetProductDetails(v []GetOrderWebhook200ResponseDataOneOfProductDetailsInner)`

SetProductDetails sets ProductDetails field to given value.

### HasProductDetails

`func (o *GetOrderWebhook200ResponseData) HasProductDetails() bool`

HasProductDetails returns a boolean if a field has been set.

### GetInvoiceRefNum

`func (o *GetOrderWebhook200ResponseData) GetInvoiceRefNum() string`

GetInvoiceRefNum returns the InvoiceRefNum field if non-nil, zero value otherwise.

### GetInvoiceRefNumOk

`func (o *GetOrderWebhook200ResponseData) GetInvoiceRefNumOk() (*string, bool)`

GetInvoiceRefNumOk returns a tuple with the InvoiceRefNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceRefNum

`func (o *GetOrderWebhook200ResponseData) SetInvoiceRefNum(v string)`

SetInvoiceRefNum sets InvoiceRefNum field to given value.

### HasInvoiceRefNum

`func (o *GetOrderWebhook200ResponseData) HasInvoiceRefNum() bool`

HasInvoiceRefNum returns a boolean if a field has been set.

### GetProducts

`func (o *GetOrderWebhook200ResponseData) GetProducts() []GetOrderWebhook200ResponseDataOneOfProductDetailsInner`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *GetOrderWebhook200ResponseData) GetProductsOk() (*[]GetOrderWebhook200ResponseDataOneOfProductDetailsInner, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *GetOrderWebhook200ResponseData) SetProducts(v []GetOrderWebhook200ResponseDataOneOfProductDetailsInner)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *GetOrderWebhook200ResponseData) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetCustomer

`func (o *GetOrderWebhook200ResponseData) GetCustomer() GetOrderWebhook200ResponseDataOneOf1Customer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *GetOrderWebhook200ResponseData) GetCustomerOk() (*GetOrderWebhook200ResponseDataOneOf1Customer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *GetOrderWebhook200ResponseData) SetCustomer(v GetOrderWebhook200ResponseDataOneOf1Customer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *GetOrderWebhook200ResponseData) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetRecipient

`func (o *GetOrderWebhook200ResponseData) GetRecipient() GetOrderWebhook200ResponseDataOneOf1Recipient`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *GetOrderWebhook200ResponseData) GetRecipientOk() (*GetOrderWebhook200ResponseDataOneOf1Recipient, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *GetOrderWebhook200ResponseData) SetRecipient(v GetOrderWebhook200ResponseDataOneOf1Recipient)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *GetOrderWebhook200ResponseData) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.

### GetShopName

`func (o *GetOrderWebhook200ResponseData) GetShopName() string`

GetShopName returns the ShopName field if non-nil, zero value otherwise.

### GetShopNameOk

`func (o *GetOrderWebhook200ResponseData) GetShopNameOk() (*string, bool)`

GetShopNameOk returns a tuple with the ShopName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopName

`func (o *GetOrderWebhook200ResponseData) SetShopName(v string)`

SetShopName sets ShopName field to given value.

### HasShopName

`func (o *GetOrderWebhook200ResponseData) HasShopName() bool`

HasShopName returns a boolean if a field has been set.

### GetPaymentId

`func (o *GetOrderWebhook200ResponseData) GetPaymentId() int64`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *GetOrderWebhook200ResponseData) GetPaymentIdOk() (*int64, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *GetOrderWebhook200ResponseData) SetPaymentId(v int64)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *GetOrderWebhook200ResponseData) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### GetPaymentDate

`func (o *GetOrderWebhook200ResponseData) GetPaymentDate() string`

GetPaymentDate returns the PaymentDate field if non-nil, zero value otherwise.

### GetPaymentDateOk

`func (o *GetOrderWebhook200ResponseData) GetPaymentDateOk() (*string, bool)`

GetPaymentDateOk returns a tuple with the PaymentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDate

`func (o *GetOrderWebhook200ResponseData) SetPaymentDate(v string)`

SetPaymentDate sets PaymentDate field to given value.

### HasPaymentDate

`func (o *GetOrderWebhook200ResponseData) HasPaymentDate() bool`

HasPaymentDate returns a boolean if a field has been set.

### GetLogistics

`func (o *GetOrderWebhook200ResponseData) GetLogistics() GetOrderWebhook200ResponseDataOneOf1Logistics`

GetLogistics returns the Logistics field if non-nil, zero value otherwise.

### GetLogisticsOk

`func (o *GetOrderWebhook200ResponseData) GetLogisticsOk() (*GetOrderWebhook200ResponseDataOneOf1Logistics, bool)`

GetLogisticsOk returns a tuple with the Logistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogistics

`func (o *GetOrderWebhook200ResponseData) SetLogistics(v GetOrderWebhook200ResponseDataOneOf1Logistics)`

SetLogistics sets Logistics field to given value.

### HasLogistics

`func (o *GetOrderWebhook200ResponseData) HasLogistics() bool`

HasLogistics returns a boolean if a field has been set.

### GetAmt

`func (o *GetOrderWebhook200ResponseData) GetAmt() GetOrderWebhook200ResponseDataOneOf1Amt`

GetAmt returns the Amt field if non-nil, zero value otherwise.

### GetAmtOk

`func (o *GetOrderWebhook200ResponseData) GetAmtOk() (*GetOrderWebhook200ResponseDataOneOf1Amt, bool)`

GetAmtOk returns a tuple with the Amt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmt

`func (o *GetOrderWebhook200ResponseData) SetAmt(v GetOrderWebhook200ResponseDataOneOf1Amt)`

SetAmt sets Amt field to given value.

### HasAmt

`func (o *GetOrderWebhook200ResponseData) HasAmt() bool`

HasAmt returns a boolean if a field has been set.

### GetDropshipperInfo

`func (o *GetOrderWebhook200ResponseData) GetDropshipperInfo() GetOrderWebhook200ResponseDataOneOf1DropshipperInfo`

GetDropshipperInfo returns the DropshipperInfo field if non-nil, zero value otherwise.

### GetDropshipperInfoOk

`func (o *GetOrderWebhook200ResponseData) GetDropshipperInfoOk() (*GetOrderWebhook200ResponseDataOneOf1DropshipperInfo, bool)`

GetDropshipperInfoOk returns a tuple with the DropshipperInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropshipperInfo

`func (o *GetOrderWebhook200ResponseData) SetDropshipperInfo(v GetOrderWebhook200ResponseDataOneOf1DropshipperInfo)`

SetDropshipperInfo sets DropshipperInfo field to given value.

### HasDropshipperInfo

`func (o *GetOrderWebhook200ResponseData) HasDropshipperInfo() bool`

HasDropshipperInfo returns a boolean if a field has been set.

### GetVoucherInfo

`func (o *GetOrderWebhook200ResponseData) GetVoucherInfo() GetOrderWebhook200ResponseDataOneOf1VoucherInfo`

GetVoucherInfo returns the VoucherInfo field if non-nil, zero value otherwise.

### GetVoucherInfoOk

`func (o *GetOrderWebhook200ResponseData) GetVoucherInfoOk() (*GetOrderWebhook200ResponseDataOneOf1VoucherInfo, bool)`

GetVoucherInfoOk returns a tuple with the VoucherInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoucherInfo

`func (o *GetOrderWebhook200ResponseData) SetVoucherInfo(v GetOrderWebhook200ResponseDataOneOf1VoucherInfo)`

SetVoucherInfo sets VoucherInfo field to given value.

### HasVoucherInfo

`func (o *GetOrderWebhook200ResponseData) HasVoucherInfo() bool`

HasVoucherInfo returns a boolean if a field has been set.

### GetDeviceType

`func (o *GetOrderWebhook200ResponseData) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *GetOrderWebhook200ResponseData) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *GetOrderWebhook200ResponseData) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *GetOrderWebhook200ResponseData) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetCreateTime

`func (o *GetOrderWebhook200ResponseData) GetCreateTime() string`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *GetOrderWebhook200ResponseData) GetCreateTimeOk() (*string, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *GetOrderWebhook200ResponseData) SetCreateTime(v string)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *GetOrderWebhook200ResponseData) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCustomFields

`func (o *GetOrderWebhook200ResponseData) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *GetOrderWebhook200ResponseData) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *GetOrderWebhook200ResponseData) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *GetOrderWebhook200ResponseData) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetAcceptPartial

`func (o *GetOrderWebhook200ResponseData) GetAcceptPartial() bool`

GetAcceptPartial returns the AcceptPartial field if non-nil, zero value otherwise.

### GetAcceptPartialOk

`func (o *GetOrderWebhook200ResponseData) GetAcceptPartialOk() (*bool, bool)`

GetAcceptPartialOk returns a tuple with the AcceptPartial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptPartial

`func (o *GetOrderWebhook200ResponseData) SetAcceptPartial(v bool)`

SetAcceptPartial sets AcceptPartial field to given value.

### HasAcceptPartial

`func (o *GetOrderWebhook200ResponseData) HasAcceptPartial() bool`

HasAcceptPartial returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


