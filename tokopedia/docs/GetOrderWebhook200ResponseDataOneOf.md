# GetOrderWebhook200ResponseDataOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FsId** | Pointer to **int64** |  | [optional] 
**OrderStatus** | Pointer to **int64** | Order status code: * 0 - Seller cancel order. * 3 - Order Reject Due Empty Stock. * 5 - Order Canceled by Fraud * 6 - Order Rejected (Auto Cancel Out of Stock) * 10 - Order rejected by seller. * 15 - Instant Cancel by Buyer. * 100 - Order Created. * 103 - Wait for payment confirmation from third party. * 220 - Payment verified, order ready to process. * 221 - Waiting for partner approval. * 400 - Seller accept order. * 450 - Waiting for pickup. * 500 - Order shipment. * 501 - Status changed to waiting resi have no input. * 520 - Invalid shipment reference number (AWB). * 530 - Requested by user to correct invalid entry of shipment reference number. * 540 - Delivered to Pickup Point. * 550 - Return to Seller. * 600 - Order delivered. * 601 - Buyer open a case to finish an order. * 690 - Fraud Review * 700 - Order finished.  | [optional] 
**OrderId** | Pointer to **int64** |  | [optional] 
**ShopId** | Pointer to **int64** |  | [optional] 
**WarehouseId** | Pointer to **int64** |  | [optional] 
**ProductDetails** | Pointer to [**[]GetOrderWebhook200ResponseDataOneOfProductDetailsInner**](GetOrderWebhook200ResponseDataOneOfProductDetailsInner.md) |  | [optional] 

## Methods

### NewGetOrderWebhook200ResponseDataOneOf

`func NewGetOrderWebhook200ResponseDataOneOf() *GetOrderWebhook200ResponseDataOneOf`

NewGetOrderWebhook200ResponseDataOneOf instantiates a new GetOrderWebhook200ResponseDataOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrderWebhook200ResponseDataOneOfWithDefaults

`func NewGetOrderWebhook200ResponseDataOneOfWithDefaults() *GetOrderWebhook200ResponseDataOneOf`

NewGetOrderWebhook200ResponseDataOneOfWithDefaults instantiates a new GetOrderWebhook200ResponseDataOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFsId

`func (o *GetOrderWebhook200ResponseDataOneOf) GetFsId() int64`

GetFsId returns the FsId field if non-nil, zero value otherwise.

### GetFsIdOk

`func (o *GetOrderWebhook200ResponseDataOneOf) GetFsIdOk() (*int64, bool)`

GetFsIdOk returns a tuple with the FsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsId

`func (o *GetOrderWebhook200ResponseDataOneOf) SetFsId(v int64)`

SetFsId sets FsId field to given value.

### HasFsId

`func (o *GetOrderWebhook200ResponseDataOneOf) HasFsId() bool`

HasFsId returns a boolean if a field has been set.

### GetOrderStatus

`func (o *GetOrderWebhook200ResponseDataOneOf) GetOrderStatus() int64`

GetOrderStatus returns the OrderStatus field if non-nil, zero value otherwise.

### GetOrderStatusOk

`func (o *GetOrderWebhook200ResponseDataOneOf) GetOrderStatusOk() (*int64, bool)`

GetOrderStatusOk returns a tuple with the OrderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderStatus

`func (o *GetOrderWebhook200ResponseDataOneOf) SetOrderStatus(v int64)`

SetOrderStatus sets OrderStatus field to given value.

### HasOrderStatus

`func (o *GetOrderWebhook200ResponseDataOneOf) HasOrderStatus() bool`

HasOrderStatus returns a boolean if a field has been set.

### GetOrderId

`func (o *GetOrderWebhook200ResponseDataOneOf) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *GetOrderWebhook200ResponseDataOneOf) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *GetOrderWebhook200ResponseDataOneOf) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *GetOrderWebhook200ResponseDataOneOf) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetShopId

`func (o *GetOrderWebhook200ResponseDataOneOf) GetShopId() int64`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *GetOrderWebhook200ResponseDataOneOf) GetShopIdOk() (*int64, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *GetOrderWebhook200ResponseDataOneOf) SetShopId(v int64)`

SetShopId sets ShopId field to given value.

### HasShopId

`func (o *GetOrderWebhook200ResponseDataOneOf) HasShopId() bool`

HasShopId returns a boolean if a field has been set.

### GetWarehouseId

`func (o *GetOrderWebhook200ResponseDataOneOf) GetWarehouseId() int64`

GetWarehouseId returns the WarehouseId field if non-nil, zero value otherwise.

### GetWarehouseIdOk

`func (o *GetOrderWebhook200ResponseDataOneOf) GetWarehouseIdOk() (*int64, bool)`

GetWarehouseIdOk returns a tuple with the WarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseId

`func (o *GetOrderWebhook200ResponseDataOneOf) SetWarehouseId(v int64)`

SetWarehouseId sets WarehouseId field to given value.

### HasWarehouseId

`func (o *GetOrderWebhook200ResponseDataOneOf) HasWarehouseId() bool`

HasWarehouseId returns a boolean if a field has been set.

### GetProductDetails

`func (o *GetOrderWebhook200ResponseDataOneOf) GetProductDetails() []GetOrderWebhook200ResponseDataOneOfProductDetailsInner`

GetProductDetails returns the ProductDetails field if non-nil, zero value otherwise.

### GetProductDetailsOk

`func (o *GetOrderWebhook200ResponseDataOneOf) GetProductDetailsOk() (*[]GetOrderWebhook200ResponseDataOneOfProductDetailsInner, bool)`

GetProductDetailsOk returns a tuple with the ProductDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductDetails

`func (o *GetOrderWebhook200ResponseDataOneOf) SetProductDetails(v []GetOrderWebhook200ResponseDataOneOfProductDetailsInner)`

SetProductDetails sets ProductDetails field to given value.

### HasProductDetails

`func (o *GetOrderWebhook200ResponseDataOneOf) HasProductDetails() bool`

HasProductDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


