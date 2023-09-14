# WebhookOrderNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FsId** | Pointer to **int64** | Fulfillment service unique identifier | [optional] 
**OrderId** | Pointer to **int64** | Order unique identifier | [optional] 
**AcceptPartial** | Pointer to **bool** | Flag which determines if the order is partially accepted | [optional] 
**InvoiceNum** | Pointer to **string** | Invoice reference number | [optional] 
**Products** | Pointer to [**[]WebhookOrderNotificationProductsInner**](WebhookOrderNotificationProductsInner.md) | Products data | [optional] 
**Customer** | Pointer to [**WebhookOrderNotificationCustomer**](WebhookOrderNotificationCustomer.md) |  | [optional] 
**Recipient** | Pointer to [**WebhookOrderNotificationRecipient**](WebhookOrderNotificationRecipient.md) |  | [optional] 
**ShopId** | Pointer to **int64** | Shop unique identifier | [optional] 
**WarehouseId** | Pointer to **int64** | Warehouse unique identifier | [optional] 
**ShopName** | Pointer to **string** | Shop name | [optional] 
**PaymentId** | Pointer to **int64** | Payment unique identifier | [optional] 
**Logistics** | Pointer to [**WebhookOrderNotificationLogistics**](WebhookOrderNotificationLogistics.md) |  | [optional] 
**Amt** | Pointer to [**WebhookOrderNotificationAmt**](WebhookOrderNotificationAmt.md) |  | [optional] 
**DropshipperInfo** | Pointer to [**WebhookOrderNotificationDropshipperInfo**](WebhookOrderNotificationDropshipperInfo.md) |  | [optional] 
**VoucherInfo** | Pointer to [**WebhookOrderNotificationVoucherInfo**](WebhookOrderNotificationVoucherInfo.md) |  | [optional] 
**DeviceType** | Pointer to **string** | User device type | [optional] 
**CreateTime** | Pointer to **int64** | Time in UNIX timestamp | [optional] 
**OrderStatus** | Pointer to **int32** | Order status | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** | A map of String to String for custom fields for future | [optional] 
**Encryption** | Pointer to [**WebhookOrderNotificationEncryption**](WebhookOrderNotificationEncryption.md) |  | [optional] 
**BundleDetail** | Pointer to [**WebhookOrderNotificationBundleDetail**](WebhookOrderNotificationBundleDetail.md) |  | [optional] 
**IsPlus** | Pointer to **bool** | Flag which determines if the order is a Plus order | [optional] 

## Methods

### NewWebhookOrderNotification

`func NewWebhookOrderNotification() *WebhookOrderNotification`

NewWebhookOrderNotification instantiates a new WebhookOrderNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookOrderNotificationWithDefaults

`func NewWebhookOrderNotificationWithDefaults() *WebhookOrderNotification`

NewWebhookOrderNotificationWithDefaults instantiates a new WebhookOrderNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFsId

`func (o *WebhookOrderNotification) GetFsId() int64`

GetFsId returns the FsId field if non-nil, zero value otherwise.

### GetFsIdOk

`func (o *WebhookOrderNotification) GetFsIdOk() (*int64, bool)`

GetFsIdOk returns a tuple with the FsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsId

`func (o *WebhookOrderNotification) SetFsId(v int64)`

SetFsId sets FsId field to given value.

### HasFsId

`func (o *WebhookOrderNotification) HasFsId() bool`

HasFsId returns a boolean if a field has been set.

### GetOrderId

`func (o *WebhookOrderNotification) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *WebhookOrderNotification) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *WebhookOrderNotification) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *WebhookOrderNotification) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetAcceptPartial

`func (o *WebhookOrderNotification) GetAcceptPartial() bool`

GetAcceptPartial returns the AcceptPartial field if non-nil, zero value otherwise.

### GetAcceptPartialOk

`func (o *WebhookOrderNotification) GetAcceptPartialOk() (*bool, bool)`

GetAcceptPartialOk returns a tuple with the AcceptPartial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptPartial

`func (o *WebhookOrderNotification) SetAcceptPartial(v bool)`

SetAcceptPartial sets AcceptPartial field to given value.

### HasAcceptPartial

`func (o *WebhookOrderNotification) HasAcceptPartial() bool`

HasAcceptPartial returns a boolean if a field has been set.

### GetInvoiceNum

`func (o *WebhookOrderNotification) GetInvoiceNum() string`

GetInvoiceNum returns the InvoiceNum field if non-nil, zero value otherwise.

### GetInvoiceNumOk

`func (o *WebhookOrderNotification) GetInvoiceNumOk() (*string, bool)`

GetInvoiceNumOk returns a tuple with the InvoiceNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNum

`func (o *WebhookOrderNotification) SetInvoiceNum(v string)`

SetInvoiceNum sets InvoiceNum field to given value.

### HasInvoiceNum

`func (o *WebhookOrderNotification) HasInvoiceNum() bool`

HasInvoiceNum returns a boolean if a field has been set.

### GetProducts

`func (o *WebhookOrderNotification) GetProducts() []WebhookOrderNotificationProductsInner`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *WebhookOrderNotification) GetProductsOk() (*[]WebhookOrderNotificationProductsInner, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *WebhookOrderNotification) SetProducts(v []WebhookOrderNotificationProductsInner)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *WebhookOrderNotification) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetCustomer

`func (o *WebhookOrderNotification) GetCustomer() WebhookOrderNotificationCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *WebhookOrderNotification) GetCustomerOk() (*WebhookOrderNotificationCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *WebhookOrderNotification) SetCustomer(v WebhookOrderNotificationCustomer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *WebhookOrderNotification) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetRecipient

`func (o *WebhookOrderNotification) GetRecipient() WebhookOrderNotificationRecipient`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *WebhookOrderNotification) GetRecipientOk() (*WebhookOrderNotificationRecipient, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *WebhookOrderNotification) SetRecipient(v WebhookOrderNotificationRecipient)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *WebhookOrderNotification) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.

### GetShopId

`func (o *WebhookOrderNotification) GetShopId() int64`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *WebhookOrderNotification) GetShopIdOk() (*int64, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *WebhookOrderNotification) SetShopId(v int64)`

SetShopId sets ShopId field to given value.

### HasShopId

`func (o *WebhookOrderNotification) HasShopId() bool`

HasShopId returns a boolean if a field has been set.

### GetWarehouseId

`func (o *WebhookOrderNotification) GetWarehouseId() int64`

GetWarehouseId returns the WarehouseId field if non-nil, zero value otherwise.

### GetWarehouseIdOk

`func (o *WebhookOrderNotification) GetWarehouseIdOk() (*int64, bool)`

GetWarehouseIdOk returns a tuple with the WarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseId

`func (o *WebhookOrderNotification) SetWarehouseId(v int64)`

SetWarehouseId sets WarehouseId field to given value.

### HasWarehouseId

`func (o *WebhookOrderNotification) HasWarehouseId() bool`

HasWarehouseId returns a boolean if a field has been set.

### GetShopName

`func (o *WebhookOrderNotification) GetShopName() string`

GetShopName returns the ShopName field if non-nil, zero value otherwise.

### GetShopNameOk

`func (o *WebhookOrderNotification) GetShopNameOk() (*string, bool)`

GetShopNameOk returns a tuple with the ShopName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopName

`func (o *WebhookOrderNotification) SetShopName(v string)`

SetShopName sets ShopName field to given value.

### HasShopName

`func (o *WebhookOrderNotification) HasShopName() bool`

HasShopName returns a boolean if a field has been set.

### GetPaymentId

`func (o *WebhookOrderNotification) GetPaymentId() int64`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *WebhookOrderNotification) GetPaymentIdOk() (*int64, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *WebhookOrderNotification) SetPaymentId(v int64)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *WebhookOrderNotification) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### GetLogistics

`func (o *WebhookOrderNotification) GetLogistics() WebhookOrderNotificationLogistics`

GetLogistics returns the Logistics field if non-nil, zero value otherwise.

### GetLogisticsOk

`func (o *WebhookOrderNotification) GetLogisticsOk() (*WebhookOrderNotificationLogistics, bool)`

GetLogisticsOk returns a tuple with the Logistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogistics

`func (o *WebhookOrderNotification) SetLogistics(v WebhookOrderNotificationLogistics)`

SetLogistics sets Logistics field to given value.

### HasLogistics

`func (o *WebhookOrderNotification) HasLogistics() bool`

HasLogistics returns a boolean if a field has been set.

### GetAmt

`func (o *WebhookOrderNotification) GetAmt() WebhookOrderNotificationAmt`

GetAmt returns the Amt field if non-nil, zero value otherwise.

### GetAmtOk

`func (o *WebhookOrderNotification) GetAmtOk() (*WebhookOrderNotificationAmt, bool)`

GetAmtOk returns a tuple with the Amt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmt

`func (o *WebhookOrderNotification) SetAmt(v WebhookOrderNotificationAmt)`

SetAmt sets Amt field to given value.

### HasAmt

`func (o *WebhookOrderNotification) HasAmt() bool`

HasAmt returns a boolean if a field has been set.

### GetDropshipperInfo

`func (o *WebhookOrderNotification) GetDropshipperInfo() WebhookOrderNotificationDropshipperInfo`

GetDropshipperInfo returns the DropshipperInfo field if non-nil, zero value otherwise.

### GetDropshipperInfoOk

`func (o *WebhookOrderNotification) GetDropshipperInfoOk() (*WebhookOrderNotificationDropshipperInfo, bool)`

GetDropshipperInfoOk returns a tuple with the DropshipperInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropshipperInfo

`func (o *WebhookOrderNotification) SetDropshipperInfo(v WebhookOrderNotificationDropshipperInfo)`

SetDropshipperInfo sets DropshipperInfo field to given value.

### HasDropshipperInfo

`func (o *WebhookOrderNotification) HasDropshipperInfo() bool`

HasDropshipperInfo returns a boolean if a field has been set.

### GetVoucherInfo

`func (o *WebhookOrderNotification) GetVoucherInfo() WebhookOrderNotificationVoucherInfo`

GetVoucherInfo returns the VoucherInfo field if non-nil, zero value otherwise.

### GetVoucherInfoOk

`func (o *WebhookOrderNotification) GetVoucherInfoOk() (*WebhookOrderNotificationVoucherInfo, bool)`

GetVoucherInfoOk returns a tuple with the VoucherInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoucherInfo

`func (o *WebhookOrderNotification) SetVoucherInfo(v WebhookOrderNotificationVoucherInfo)`

SetVoucherInfo sets VoucherInfo field to given value.

### HasVoucherInfo

`func (o *WebhookOrderNotification) HasVoucherInfo() bool`

HasVoucherInfo returns a boolean if a field has been set.

### GetDeviceType

`func (o *WebhookOrderNotification) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *WebhookOrderNotification) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *WebhookOrderNotification) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *WebhookOrderNotification) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetCreateTime

`func (o *WebhookOrderNotification) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *WebhookOrderNotification) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *WebhookOrderNotification) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *WebhookOrderNotification) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetOrderStatus

`func (o *WebhookOrderNotification) GetOrderStatus() int32`

GetOrderStatus returns the OrderStatus field if non-nil, zero value otherwise.

### GetOrderStatusOk

`func (o *WebhookOrderNotification) GetOrderStatusOk() (*int32, bool)`

GetOrderStatusOk returns a tuple with the OrderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderStatus

`func (o *WebhookOrderNotification) SetOrderStatus(v int32)`

SetOrderStatus sets OrderStatus field to given value.

### HasOrderStatus

`func (o *WebhookOrderNotification) HasOrderStatus() bool`

HasOrderStatus returns a boolean if a field has been set.

### GetCustomFields

`func (o *WebhookOrderNotification) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *WebhookOrderNotification) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *WebhookOrderNotification) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *WebhookOrderNotification) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetEncryption

`func (o *WebhookOrderNotification) GetEncryption() WebhookOrderNotificationEncryption`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *WebhookOrderNotification) GetEncryptionOk() (*WebhookOrderNotificationEncryption, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *WebhookOrderNotification) SetEncryption(v WebhookOrderNotificationEncryption)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *WebhookOrderNotification) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.

### GetBundleDetail

`func (o *WebhookOrderNotification) GetBundleDetail() WebhookOrderNotificationBundleDetail`

GetBundleDetail returns the BundleDetail field if non-nil, zero value otherwise.

### GetBundleDetailOk

`func (o *WebhookOrderNotification) GetBundleDetailOk() (*WebhookOrderNotificationBundleDetail, bool)`

GetBundleDetailOk returns a tuple with the BundleDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleDetail

`func (o *WebhookOrderNotification) SetBundleDetail(v WebhookOrderNotificationBundleDetail)`

SetBundleDetail sets BundleDetail field to given value.

### HasBundleDetail

`func (o *WebhookOrderNotification) HasBundleDetail() bool`

HasBundleDetail returns a boolean if a field has been set.

### GetIsPlus

`func (o *WebhookOrderNotification) GetIsPlus() bool`

GetIsPlus returns the IsPlus field if non-nil, zero value otherwise.

### GetIsPlusOk

`func (o *WebhookOrderNotification) GetIsPlusOk() (*bool, bool)`

GetIsPlusOk returns a tuple with the IsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPlus

`func (o *WebhookOrderNotification) SetIsPlus(v bool)`

SetIsPlus sets IsPlus field to given value.

### HasIsPlus

`func (o *WebhookOrderNotification) HasIsPlus() bool`

HasIsPlus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


