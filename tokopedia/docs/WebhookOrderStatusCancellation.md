# WebhookOrderStatusCancellation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderStatus** | Pointer to **int64** | Order status code: * 0 - Seller cancel order. * 3 - Order Reject Due Empty Stock. * 5 - Order Canceled by Fraud * 6 - Order Rejected (Auto Cancel Out of Stock) * 10 - Order rejected by seller. * 15 - Instant Cancel by Buyer. * 100 - Order Created. * 103 - Wait for payment confirmation from third party. * 220 - Payment verified, order ready to process. * 221 - Waiting for partner approval. * 400 - Seller accept order. * 450 - Waiting for pickup. * 500 - Order shipment. * 501 - Status changed to waiting resi have no input. * 520 - Invalid shipment reference number (AWB). * 530 - Requested by user to correct invalid entry of shipment reference number. * 540 - Delivered to Pickup Point. * 550 - Return to Seller. * 600 - Order delivered. * 601 - Buyer open a case to finish an order. * 690 - Fraud Review * 700 - Order finished.  | [optional] 
**FsId** | Pointer to **int64** | Fulfillment service unique identifier | [optional] 
**ShopId** | Pointer to **int64** | Shop unique identifier | [optional] 
**OrderId** | Pointer to **int64** | Order unique identifier | [optional] 
**ProductDetails** | Pointer to [**[]GetOrderWebhook200ResponseDataOneOfProductDetailsInner**](GetOrderWebhook200ResponseDataOneOfProductDetailsInner.md) | List of products | [optional] 

## Methods

### NewWebhookOrderStatusCancellation

`func NewWebhookOrderStatusCancellation() *WebhookOrderStatusCancellation`

NewWebhookOrderStatusCancellation instantiates a new WebhookOrderStatusCancellation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookOrderStatusCancellationWithDefaults

`func NewWebhookOrderStatusCancellationWithDefaults() *WebhookOrderStatusCancellation`

NewWebhookOrderStatusCancellationWithDefaults instantiates a new WebhookOrderStatusCancellation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderStatus

`func (o *WebhookOrderStatusCancellation) GetOrderStatus() int64`

GetOrderStatus returns the OrderStatus field if non-nil, zero value otherwise.

### GetOrderStatusOk

`func (o *WebhookOrderStatusCancellation) GetOrderStatusOk() (*int64, bool)`

GetOrderStatusOk returns a tuple with the OrderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderStatus

`func (o *WebhookOrderStatusCancellation) SetOrderStatus(v int64)`

SetOrderStatus sets OrderStatus field to given value.

### HasOrderStatus

`func (o *WebhookOrderStatusCancellation) HasOrderStatus() bool`

HasOrderStatus returns a boolean if a field has been set.

### GetFsId

`func (o *WebhookOrderStatusCancellation) GetFsId() int64`

GetFsId returns the FsId field if non-nil, zero value otherwise.

### GetFsIdOk

`func (o *WebhookOrderStatusCancellation) GetFsIdOk() (*int64, bool)`

GetFsIdOk returns a tuple with the FsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsId

`func (o *WebhookOrderStatusCancellation) SetFsId(v int64)`

SetFsId sets FsId field to given value.

### HasFsId

`func (o *WebhookOrderStatusCancellation) HasFsId() bool`

HasFsId returns a boolean if a field has been set.

### GetShopId

`func (o *WebhookOrderStatusCancellation) GetShopId() int64`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *WebhookOrderStatusCancellation) GetShopIdOk() (*int64, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *WebhookOrderStatusCancellation) SetShopId(v int64)`

SetShopId sets ShopId field to given value.

### HasShopId

`func (o *WebhookOrderStatusCancellation) HasShopId() bool`

HasShopId returns a boolean if a field has been set.

### GetOrderId

`func (o *WebhookOrderStatusCancellation) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *WebhookOrderStatusCancellation) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *WebhookOrderStatusCancellation) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *WebhookOrderStatusCancellation) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetProductDetails

`func (o *WebhookOrderStatusCancellation) GetProductDetails() []GetOrderWebhook200ResponseDataOneOfProductDetailsInner`

GetProductDetails returns the ProductDetails field if non-nil, zero value otherwise.

### GetProductDetailsOk

`func (o *WebhookOrderStatusCancellation) GetProductDetailsOk() (*[]GetOrderWebhook200ResponseDataOneOfProductDetailsInner, bool)`

GetProductDetailsOk returns a tuple with the ProductDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductDetails

`func (o *WebhookOrderStatusCancellation) SetProductDetails(v []GetOrderWebhook200ResponseDataOneOfProductDetailsInner)`

SetProductDetails sets ProductDetails field to given value.

### HasProductDetails

`func (o *WebhookOrderStatusCancellation) HasProductDetails() bool`

HasProductDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


