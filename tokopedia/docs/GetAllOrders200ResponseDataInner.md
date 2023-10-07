# GetAllOrders200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FsId** | Pointer to **string** | Fullfilment Service Unique Identifier | [optional] 
**OrderId** | Pointer to **int64** | Order Unique Identifier | [optional] 
**IsCodMitra** | Pointer to **bool** | Is COD mitra? | [optional] 
**AcceptPartial** | Pointer to **bool** | Is accept partial? | [optional] 
**InvoiceRefNum** | Pointer to **string** | Invoice Number | [optional] 
**HaveProductBundle** | Pointer to **bool** | Have product bundle? | [optional] 
**Products** | Pointer to [**[]GetAllOrders200ResponseDataInnerProductsInner**](GetAllOrders200ResponseDataInnerProductsInner.md) |  | [optional] 
**ProductsFulfilled** | Pointer to [**[]GetAllOrders200ResponseDataInnerProductsFulfilledInner**](GetAllOrders200ResponseDataInnerProductsFulfilledInner.md) |  | [optional] 
**BundleDetail** | Pointer to [**GetAllOrders200ResponseDataInnerBundleDetail**](GetAllOrders200ResponseDataInnerBundleDetail.md) |  | [optional] 
**DeviceType** | Pointer to **string** | Device Type | [optional] 
**Buyer** | Pointer to [**GetAllOrders200ResponseDataInnerBuyer**](GetAllOrders200ResponseDataInnerBuyer.md) |  | [optional] 
**ShopId** | Pointer to **int64** | Shop Unique Identifier | [optional] 
**PaymentId** | Pointer to **int64** | Payment Unique Identifier | [optional] 
**Recipient** | Pointer to [**GetAllOrders200ResponseDataInnerRecipient**](GetAllOrders200ResponseDataInnerRecipient.md) |  | [optional] 
**Logistics** | Pointer to [**GetAllOrders200ResponseDataInnerLogistics**](GetAllOrders200ResponseDataInnerLogistics.md) |  | [optional] 
**Amt** | Pointer to [**GetAllOrders200ResponseDataInnerAmt**](GetAllOrders200ResponseDataInnerAmt.md) |  | [optional] 
**DropshipperInfo** | Pointer to **map[string]interface{}** |  | [optional] 
**VoucherInfo** | Pointer to [**GetAllOrders200ResponseDataInnerVoucherInfo**](GetAllOrders200ResponseDataInnerVoucherInfo.md) |  | [optional] 
**OrderStatus** | Pointer to **int64** | Order Status Code | [optional] 
**WarehouseId** | Pointer to **int64** | Warehouse Unique Identifier | [optional] 
**FulfillBy** | Pointer to **int64** | Order Fulfilled by Tokocabang (1), Order Fulfilled by Shop (0) | [optional] 
**CreateTime** | Pointer to **int64** | Order Created Time | [optional] 
**CustomFields** | Pointer to [**GetAllOrders200ResponseDataInnerCustomFields**](GetAllOrders200ResponseDataInnerCustomFields.md) |  | [optional] 
**PromoOrderDetail** | Pointer to [**GetAllOrders200ResponseDataInnerPromoOrderDetail**](GetAllOrders200ResponseDataInnerPromoOrderDetail.md) |  | [optional] 
**Encryption** | Pointer to [**GetAllOrders200ResponseDataInnerEncryption**](GetAllOrders200ResponseDataInnerEncryption.md) |  | [optional] 
**AddonInfo** | Pointer to **map[string]interface{}** |  | [optional] 
**ShipmentFulfillment** | Pointer to [**GetAllOrders200ResponseDataInnerShipmentFulfillment**](GetAllOrders200ResponseDataInnerShipmentFulfillment.md) |  | [optional] 
**IsPlus** | Pointer to **bool** | Plus Order Flag | [optional] 
**PofInfo** | Pointer to [**GetAllOrders200ResponseDataInnerPofInfo**](GetAllOrders200ResponseDataInnerPofInfo.md) |  | [optional] 

## Methods

### NewGetAllOrders200ResponseDataInner

`func NewGetAllOrders200ResponseDataInner() *GetAllOrders200ResponseDataInner`

NewGetAllOrders200ResponseDataInner instantiates a new GetAllOrders200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllOrders200ResponseDataInnerWithDefaults

`func NewGetAllOrders200ResponseDataInnerWithDefaults() *GetAllOrders200ResponseDataInner`

NewGetAllOrders200ResponseDataInnerWithDefaults instantiates a new GetAllOrders200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFsId

`func (o *GetAllOrders200ResponseDataInner) GetFsId() string`

GetFsId returns the FsId field if non-nil, zero value otherwise.

### GetFsIdOk

`func (o *GetAllOrders200ResponseDataInner) GetFsIdOk() (*string, bool)`

GetFsIdOk returns a tuple with the FsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsId

`func (o *GetAllOrders200ResponseDataInner) SetFsId(v string)`

SetFsId sets FsId field to given value.

### HasFsId

`func (o *GetAllOrders200ResponseDataInner) HasFsId() bool`

HasFsId returns a boolean if a field has been set.

### GetOrderId

`func (o *GetAllOrders200ResponseDataInner) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *GetAllOrders200ResponseDataInner) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *GetAllOrders200ResponseDataInner) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *GetAllOrders200ResponseDataInner) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetIsCodMitra

`func (o *GetAllOrders200ResponseDataInner) GetIsCodMitra() bool`

GetIsCodMitra returns the IsCodMitra field if non-nil, zero value otherwise.

### GetIsCodMitraOk

`func (o *GetAllOrders200ResponseDataInner) GetIsCodMitraOk() (*bool, bool)`

GetIsCodMitraOk returns a tuple with the IsCodMitra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCodMitra

`func (o *GetAllOrders200ResponseDataInner) SetIsCodMitra(v bool)`

SetIsCodMitra sets IsCodMitra field to given value.

### HasIsCodMitra

`func (o *GetAllOrders200ResponseDataInner) HasIsCodMitra() bool`

HasIsCodMitra returns a boolean if a field has been set.

### GetAcceptPartial

`func (o *GetAllOrders200ResponseDataInner) GetAcceptPartial() bool`

GetAcceptPartial returns the AcceptPartial field if non-nil, zero value otherwise.

### GetAcceptPartialOk

`func (o *GetAllOrders200ResponseDataInner) GetAcceptPartialOk() (*bool, bool)`

GetAcceptPartialOk returns a tuple with the AcceptPartial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptPartial

`func (o *GetAllOrders200ResponseDataInner) SetAcceptPartial(v bool)`

SetAcceptPartial sets AcceptPartial field to given value.

### HasAcceptPartial

`func (o *GetAllOrders200ResponseDataInner) HasAcceptPartial() bool`

HasAcceptPartial returns a boolean if a field has been set.

### GetInvoiceRefNum

`func (o *GetAllOrders200ResponseDataInner) GetInvoiceRefNum() string`

GetInvoiceRefNum returns the InvoiceRefNum field if non-nil, zero value otherwise.

### GetInvoiceRefNumOk

`func (o *GetAllOrders200ResponseDataInner) GetInvoiceRefNumOk() (*string, bool)`

GetInvoiceRefNumOk returns a tuple with the InvoiceRefNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceRefNum

`func (o *GetAllOrders200ResponseDataInner) SetInvoiceRefNum(v string)`

SetInvoiceRefNum sets InvoiceRefNum field to given value.

### HasInvoiceRefNum

`func (o *GetAllOrders200ResponseDataInner) HasInvoiceRefNum() bool`

HasInvoiceRefNum returns a boolean if a field has been set.

### GetHaveProductBundle

`func (o *GetAllOrders200ResponseDataInner) GetHaveProductBundle() bool`

GetHaveProductBundle returns the HaveProductBundle field if non-nil, zero value otherwise.

### GetHaveProductBundleOk

`func (o *GetAllOrders200ResponseDataInner) GetHaveProductBundleOk() (*bool, bool)`

GetHaveProductBundleOk returns a tuple with the HaveProductBundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaveProductBundle

`func (o *GetAllOrders200ResponseDataInner) SetHaveProductBundle(v bool)`

SetHaveProductBundle sets HaveProductBundle field to given value.

### HasHaveProductBundle

`func (o *GetAllOrders200ResponseDataInner) HasHaveProductBundle() bool`

HasHaveProductBundle returns a boolean if a field has been set.

### GetProducts

`func (o *GetAllOrders200ResponseDataInner) GetProducts() []GetAllOrders200ResponseDataInnerProductsInner`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *GetAllOrders200ResponseDataInner) GetProductsOk() (*[]GetAllOrders200ResponseDataInnerProductsInner, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *GetAllOrders200ResponseDataInner) SetProducts(v []GetAllOrders200ResponseDataInnerProductsInner)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *GetAllOrders200ResponseDataInner) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetProductsFulfilled

`func (o *GetAllOrders200ResponseDataInner) GetProductsFulfilled() []GetAllOrders200ResponseDataInnerProductsFulfilledInner`

GetProductsFulfilled returns the ProductsFulfilled field if non-nil, zero value otherwise.

### GetProductsFulfilledOk

`func (o *GetAllOrders200ResponseDataInner) GetProductsFulfilledOk() (*[]GetAllOrders200ResponseDataInnerProductsFulfilledInner, bool)`

GetProductsFulfilledOk returns a tuple with the ProductsFulfilled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductsFulfilled

`func (o *GetAllOrders200ResponseDataInner) SetProductsFulfilled(v []GetAllOrders200ResponseDataInnerProductsFulfilledInner)`

SetProductsFulfilled sets ProductsFulfilled field to given value.

### HasProductsFulfilled

`func (o *GetAllOrders200ResponseDataInner) HasProductsFulfilled() bool`

HasProductsFulfilled returns a boolean if a field has been set.

### GetBundleDetail

`func (o *GetAllOrders200ResponseDataInner) GetBundleDetail() GetAllOrders200ResponseDataInnerBundleDetail`

GetBundleDetail returns the BundleDetail field if non-nil, zero value otherwise.

### GetBundleDetailOk

`func (o *GetAllOrders200ResponseDataInner) GetBundleDetailOk() (*GetAllOrders200ResponseDataInnerBundleDetail, bool)`

GetBundleDetailOk returns a tuple with the BundleDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleDetail

`func (o *GetAllOrders200ResponseDataInner) SetBundleDetail(v GetAllOrders200ResponseDataInnerBundleDetail)`

SetBundleDetail sets BundleDetail field to given value.

### HasBundleDetail

`func (o *GetAllOrders200ResponseDataInner) HasBundleDetail() bool`

HasBundleDetail returns a boolean if a field has been set.

### GetDeviceType

`func (o *GetAllOrders200ResponseDataInner) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *GetAllOrders200ResponseDataInner) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *GetAllOrders200ResponseDataInner) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *GetAllOrders200ResponseDataInner) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetBuyer

`func (o *GetAllOrders200ResponseDataInner) GetBuyer() GetAllOrders200ResponseDataInnerBuyer`

GetBuyer returns the Buyer field if non-nil, zero value otherwise.

### GetBuyerOk

`func (o *GetAllOrders200ResponseDataInner) GetBuyerOk() (*GetAllOrders200ResponseDataInnerBuyer, bool)`

GetBuyerOk returns a tuple with the Buyer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyer

`func (o *GetAllOrders200ResponseDataInner) SetBuyer(v GetAllOrders200ResponseDataInnerBuyer)`

SetBuyer sets Buyer field to given value.

### HasBuyer

`func (o *GetAllOrders200ResponseDataInner) HasBuyer() bool`

HasBuyer returns a boolean if a field has been set.

### GetShopId

`func (o *GetAllOrders200ResponseDataInner) GetShopId() int64`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *GetAllOrders200ResponseDataInner) GetShopIdOk() (*int64, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *GetAllOrders200ResponseDataInner) SetShopId(v int64)`

SetShopId sets ShopId field to given value.

### HasShopId

`func (o *GetAllOrders200ResponseDataInner) HasShopId() bool`

HasShopId returns a boolean if a field has been set.

### GetPaymentId

`func (o *GetAllOrders200ResponseDataInner) GetPaymentId() int64`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *GetAllOrders200ResponseDataInner) GetPaymentIdOk() (*int64, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *GetAllOrders200ResponseDataInner) SetPaymentId(v int64)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *GetAllOrders200ResponseDataInner) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### GetRecipient

`func (o *GetAllOrders200ResponseDataInner) GetRecipient() GetAllOrders200ResponseDataInnerRecipient`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *GetAllOrders200ResponseDataInner) GetRecipientOk() (*GetAllOrders200ResponseDataInnerRecipient, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *GetAllOrders200ResponseDataInner) SetRecipient(v GetAllOrders200ResponseDataInnerRecipient)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *GetAllOrders200ResponseDataInner) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.

### GetLogistics

`func (o *GetAllOrders200ResponseDataInner) GetLogistics() GetAllOrders200ResponseDataInnerLogistics`

GetLogistics returns the Logistics field if non-nil, zero value otherwise.

### GetLogisticsOk

`func (o *GetAllOrders200ResponseDataInner) GetLogisticsOk() (*GetAllOrders200ResponseDataInnerLogistics, bool)`

GetLogisticsOk returns a tuple with the Logistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogistics

`func (o *GetAllOrders200ResponseDataInner) SetLogistics(v GetAllOrders200ResponseDataInnerLogistics)`

SetLogistics sets Logistics field to given value.

### HasLogistics

`func (o *GetAllOrders200ResponseDataInner) HasLogistics() bool`

HasLogistics returns a boolean if a field has been set.

### GetAmt

`func (o *GetAllOrders200ResponseDataInner) GetAmt() GetAllOrders200ResponseDataInnerAmt`

GetAmt returns the Amt field if non-nil, zero value otherwise.

### GetAmtOk

`func (o *GetAllOrders200ResponseDataInner) GetAmtOk() (*GetAllOrders200ResponseDataInnerAmt, bool)`

GetAmtOk returns a tuple with the Amt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmt

`func (o *GetAllOrders200ResponseDataInner) SetAmt(v GetAllOrders200ResponseDataInnerAmt)`

SetAmt sets Amt field to given value.

### HasAmt

`func (o *GetAllOrders200ResponseDataInner) HasAmt() bool`

HasAmt returns a boolean if a field has been set.

### GetDropshipperInfo

`func (o *GetAllOrders200ResponseDataInner) GetDropshipperInfo() map[string]interface{}`

GetDropshipperInfo returns the DropshipperInfo field if non-nil, zero value otherwise.

### GetDropshipperInfoOk

`func (o *GetAllOrders200ResponseDataInner) GetDropshipperInfoOk() (*map[string]interface{}, bool)`

GetDropshipperInfoOk returns a tuple with the DropshipperInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropshipperInfo

`func (o *GetAllOrders200ResponseDataInner) SetDropshipperInfo(v map[string]interface{})`

SetDropshipperInfo sets DropshipperInfo field to given value.

### HasDropshipperInfo

`func (o *GetAllOrders200ResponseDataInner) HasDropshipperInfo() bool`

HasDropshipperInfo returns a boolean if a field has been set.

### GetVoucherInfo

`func (o *GetAllOrders200ResponseDataInner) GetVoucherInfo() GetAllOrders200ResponseDataInnerVoucherInfo`

GetVoucherInfo returns the VoucherInfo field if non-nil, zero value otherwise.

### GetVoucherInfoOk

`func (o *GetAllOrders200ResponseDataInner) GetVoucherInfoOk() (*GetAllOrders200ResponseDataInnerVoucherInfo, bool)`

GetVoucherInfoOk returns a tuple with the VoucherInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoucherInfo

`func (o *GetAllOrders200ResponseDataInner) SetVoucherInfo(v GetAllOrders200ResponseDataInnerVoucherInfo)`

SetVoucherInfo sets VoucherInfo field to given value.

### HasVoucherInfo

`func (o *GetAllOrders200ResponseDataInner) HasVoucherInfo() bool`

HasVoucherInfo returns a boolean if a field has been set.

### GetOrderStatus

`func (o *GetAllOrders200ResponseDataInner) GetOrderStatus() int64`

GetOrderStatus returns the OrderStatus field if non-nil, zero value otherwise.

### GetOrderStatusOk

`func (o *GetAllOrders200ResponseDataInner) GetOrderStatusOk() (*int64, bool)`

GetOrderStatusOk returns a tuple with the OrderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderStatus

`func (o *GetAllOrders200ResponseDataInner) SetOrderStatus(v int64)`

SetOrderStatus sets OrderStatus field to given value.

### HasOrderStatus

`func (o *GetAllOrders200ResponseDataInner) HasOrderStatus() bool`

HasOrderStatus returns a boolean if a field has been set.

### GetWarehouseId

`func (o *GetAllOrders200ResponseDataInner) GetWarehouseId() int64`

GetWarehouseId returns the WarehouseId field if non-nil, zero value otherwise.

### GetWarehouseIdOk

`func (o *GetAllOrders200ResponseDataInner) GetWarehouseIdOk() (*int64, bool)`

GetWarehouseIdOk returns a tuple with the WarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseId

`func (o *GetAllOrders200ResponseDataInner) SetWarehouseId(v int64)`

SetWarehouseId sets WarehouseId field to given value.

### HasWarehouseId

`func (o *GetAllOrders200ResponseDataInner) HasWarehouseId() bool`

HasWarehouseId returns a boolean if a field has been set.

### GetFulfillBy

`func (o *GetAllOrders200ResponseDataInner) GetFulfillBy() int64`

GetFulfillBy returns the FulfillBy field if non-nil, zero value otherwise.

### GetFulfillByOk

`func (o *GetAllOrders200ResponseDataInner) GetFulfillByOk() (*int64, bool)`

GetFulfillByOk returns a tuple with the FulfillBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillBy

`func (o *GetAllOrders200ResponseDataInner) SetFulfillBy(v int64)`

SetFulfillBy sets FulfillBy field to given value.

### HasFulfillBy

`func (o *GetAllOrders200ResponseDataInner) HasFulfillBy() bool`

HasFulfillBy returns a boolean if a field has been set.

### GetCreateTime

`func (o *GetAllOrders200ResponseDataInner) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *GetAllOrders200ResponseDataInner) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *GetAllOrders200ResponseDataInner) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *GetAllOrders200ResponseDataInner) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCustomFields

`func (o *GetAllOrders200ResponseDataInner) GetCustomFields() GetAllOrders200ResponseDataInnerCustomFields`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *GetAllOrders200ResponseDataInner) GetCustomFieldsOk() (*GetAllOrders200ResponseDataInnerCustomFields, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *GetAllOrders200ResponseDataInner) SetCustomFields(v GetAllOrders200ResponseDataInnerCustomFields)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *GetAllOrders200ResponseDataInner) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetPromoOrderDetail

`func (o *GetAllOrders200ResponseDataInner) GetPromoOrderDetail() GetAllOrders200ResponseDataInnerPromoOrderDetail`

GetPromoOrderDetail returns the PromoOrderDetail field if non-nil, zero value otherwise.

### GetPromoOrderDetailOk

`func (o *GetAllOrders200ResponseDataInner) GetPromoOrderDetailOk() (*GetAllOrders200ResponseDataInnerPromoOrderDetail, bool)`

GetPromoOrderDetailOk returns a tuple with the PromoOrderDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromoOrderDetail

`func (o *GetAllOrders200ResponseDataInner) SetPromoOrderDetail(v GetAllOrders200ResponseDataInnerPromoOrderDetail)`

SetPromoOrderDetail sets PromoOrderDetail field to given value.

### HasPromoOrderDetail

`func (o *GetAllOrders200ResponseDataInner) HasPromoOrderDetail() bool`

HasPromoOrderDetail returns a boolean if a field has been set.

### GetEncryption

`func (o *GetAllOrders200ResponseDataInner) GetEncryption() GetAllOrders200ResponseDataInnerEncryption`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *GetAllOrders200ResponseDataInner) GetEncryptionOk() (*GetAllOrders200ResponseDataInnerEncryption, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *GetAllOrders200ResponseDataInner) SetEncryption(v GetAllOrders200ResponseDataInnerEncryption)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *GetAllOrders200ResponseDataInner) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.

### GetAddonInfo

`func (o *GetAllOrders200ResponseDataInner) GetAddonInfo() map[string]interface{}`

GetAddonInfo returns the AddonInfo field if non-nil, zero value otherwise.

### GetAddonInfoOk

`func (o *GetAllOrders200ResponseDataInner) GetAddonInfoOk() (*map[string]interface{}, bool)`

GetAddonInfoOk returns a tuple with the AddonInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddonInfo

`func (o *GetAllOrders200ResponseDataInner) SetAddonInfo(v map[string]interface{})`

SetAddonInfo sets AddonInfo field to given value.

### HasAddonInfo

`func (o *GetAllOrders200ResponseDataInner) HasAddonInfo() bool`

HasAddonInfo returns a boolean if a field has been set.

### GetShipmentFulfillment

`func (o *GetAllOrders200ResponseDataInner) GetShipmentFulfillment() GetAllOrders200ResponseDataInnerShipmentFulfillment`

GetShipmentFulfillment returns the ShipmentFulfillment field if non-nil, zero value otherwise.

### GetShipmentFulfillmentOk

`func (o *GetAllOrders200ResponseDataInner) GetShipmentFulfillmentOk() (*GetAllOrders200ResponseDataInnerShipmentFulfillment, bool)`

GetShipmentFulfillmentOk returns a tuple with the ShipmentFulfillment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentFulfillment

`func (o *GetAllOrders200ResponseDataInner) SetShipmentFulfillment(v GetAllOrders200ResponseDataInnerShipmentFulfillment)`

SetShipmentFulfillment sets ShipmentFulfillment field to given value.

### HasShipmentFulfillment

`func (o *GetAllOrders200ResponseDataInner) HasShipmentFulfillment() bool`

HasShipmentFulfillment returns a boolean if a field has been set.

### GetIsPlus

`func (o *GetAllOrders200ResponseDataInner) GetIsPlus() bool`

GetIsPlus returns the IsPlus field if non-nil, zero value otherwise.

### GetIsPlusOk

`func (o *GetAllOrders200ResponseDataInner) GetIsPlusOk() (*bool, bool)`

GetIsPlusOk returns a tuple with the IsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPlus

`func (o *GetAllOrders200ResponseDataInner) SetIsPlus(v bool)`

SetIsPlus sets IsPlus field to given value.

### HasIsPlus

`func (o *GetAllOrders200ResponseDataInner) HasIsPlus() bool`

HasIsPlus returns a boolean if a field has been set.

### GetPofInfo

`func (o *GetAllOrders200ResponseDataInner) GetPofInfo() GetAllOrders200ResponseDataInnerPofInfo`

GetPofInfo returns the PofInfo field if non-nil, zero value otherwise.

### GetPofInfoOk

`func (o *GetAllOrders200ResponseDataInner) GetPofInfoOk() (*GetAllOrders200ResponseDataInnerPofInfo, bool)`

GetPofInfoOk returns a tuple with the PofInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPofInfo

`func (o *GetAllOrders200ResponseDataInner) SetPofInfo(v GetAllOrders200ResponseDataInnerPofInfo)`

SetPofInfo sets PofInfo field to given value.

### HasPofInfo

`func (o *GetAllOrders200ResponseDataInner) HasPofInfo() bool`

HasPofInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


