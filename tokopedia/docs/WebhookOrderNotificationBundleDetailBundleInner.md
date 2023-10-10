# WebhookOrderNotificationBundleDetailBundleInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BundleId** | Pointer to **int64** |  | [optional] 
**BundleVariantId** | Pointer to **string** |  | [optional] 
**BundleName** | Pointer to **string** |  | [optional] 
**BundlePrice** | Pointer to **float64** |  | [optional] 
**BundleQuantity** | Pointer to **int64** |  | [optional] 
**BundleSubtotalPrice** | Pointer to **float64** |  | [optional] 
**OrderDetail** | Pointer to [**[]WebhookOrderNotificationBundleDetailBundleInnerOrderDetailInner**](WebhookOrderNotificationBundleDetailBundleInnerOrderDetailInner.md) |  | [optional] 

## Methods

### NewWebhookOrderNotificationBundleDetailBundleInner

`func NewWebhookOrderNotificationBundleDetailBundleInner() *WebhookOrderNotificationBundleDetailBundleInner`

NewWebhookOrderNotificationBundleDetailBundleInner instantiates a new WebhookOrderNotificationBundleDetailBundleInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookOrderNotificationBundleDetailBundleInnerWithDefaults

`func NewWebhookOrderNotificationBundleDetailBundleInnerWithDefaults() *WebhookOrderNotificationBundleDetailBundleInner`

NewWebhookOrderNotificationBundleDetailBundleInnerWithDefaults instantiates a new WebhookOrderNotificationBundleDetailBundleInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBundleId

`func (o *WebhookOrderNotificationBundleDetailBundleInner) GetBundleId() int64`

GetBundleId returns the BundleId field if non-nil, zero value otherwise.

### GetBundleIdOk

`func (o *WebhookOrderNotificationBundleDetailBundleInner) GetBundleIdOk() (*int64, bool)`

GetBundleIdOk returns a tuple with the BundleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleId

`func (o *WebhookOrderNotificationBundleDetailBundleInner) SetBundleId(v int64)`

SetBundleId sets BundleId field to given value.

### HasBundleId

`func (o *WebhookOrderNotificationBundleDetailBundleInner) HasBundleId() bool`

HasBundleId returns a boolean if a field has been set.

### GetBundleVariantId

`func (o *WebhookOrderNotificationBundleDetailBundleInner) GetBundleVariantId() string`

GetBundleVariantId returns the BundleVariantId field if non-nil, zero value otherwise.

### GetBundleVariantIdOk

`func (o *WebhookOrderNotificationBundleDetailBundleInner) GetBundleVariantIdOk() (*string, bool)`

GetBundleVariantIdOk returns a tuple with the BundleVariantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleVariantId

`func (o *WebhookOrderNotificationBundleDetailBundleInner) SetBundleVariantId(v string)`

SetBundleVariantId sets BundleVariantId field to given value.

### HasBundleVariantId

`func (o *WebhookOrderNotificationBundleDetailBundleInner) HasBundleVariantId() bool`

HasBundleVariantId returns a boolean if a field has been set.

### GetBundleName

`func (o *WebhookOrderNotificationBundleDetailBundleInner) GetBundleName() string`

GetBundleName returns the BundleName field if non-nil, zero value otherwise.

### GetBundleNameOk

`func (o *WebhookOrderNotificationBundleDetailBundleInner) GetBundleNameOk() (*string, bool)`

GetBundleNameOk returns a tuple with the BundleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleName

`func (o *WebhookOrderNotificationBundleDetailBundleInner) SetBundleName(v string)`

SetBundleName sets BundleName field to given value.

### HasBundleName

`func (o *WebhookOrderNotificationBundleDetailBundleInner) HasBundleName() bool`

HasBundleName returns a boolean if a field has been set.

### GetBundlePrice

`func (o *WebhookOrderNotificationBundleDetailBundleInner) GetBundlePrice() float64`

GetBundlePrice returns the BundlePrice field if non-nil, zero value otherwise.

### GetBundlePriceOk

`func (o *WebhookOrderNotificationBundleDetailBundleInner) GetBundlePriceOk() (*float64, bool)`

GetBundlePriceOk returns a tuple with the BundlePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundlePrice

`func (o *WebhookOrderNotificationBundleDetailBundleInner) SetBundlePrice(v float64)`

SetBundlePrice sets BundlePrice field to given value.

### HasBundlePrice

`func (o *WebhookOrderNotificationBundleDetailBundleInner) HasBundlePrice() bool`

HasBundlePrice returns a boolean if a field has been set.

### GetBundleQuantity

`func (o *WebhookOrderNotificationBundleDetailBundleInner) GetBundleQuantity() int64`

GetBundleQuantity returns the BundleQuantity field if non-nil, zero value otherwise.

### GetBundleQuantityOk

`func (o *WebhookOrderNotificationBundleDetailBundleInner) GetBundleQuantityOk() (*int64, bool)`

GetBundleQuantityOk returns a tuple with the BundleQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleQuantity

`func (o *WebhookOrderNotificationBundleDetailBundleInner) SetBundleQuantity(v int64)`

SetBundleQuantity sets BundleQuantity field to given value.

### HasBundleQuantity

`func (o *WebhookOrderNotificationBundleDetailBundleInner) HasBundleQuantity() bool`

HasBundleQuantity returns a boolean if a field has been set.

### GetBundleSubtotalPrice

`func (o *WebhookOrderNotificationBundleDetailBundleInner) GetBundleSubtotalPrice() float64`

GetBundleSubtotalPrice returns the BundleSubtotalPrice field if non-nil, zero value otherwise.

### GetBundleSubtotalPriceOk

`func (o *WebhookOrderNotificationBundleDetailBundleInner) GetBundleSubtotalPriceOk() (*float64, bool)`

GetBundleSubtotalPriceOk returns a tuple with the BundleSubtotalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleSubtotalPrice

`func (o *WebhookOrderNotificationBundleDetailBundleInner) SetBundleSubtotalPrice(v float64)`

SetBundleSubtotalPrice sets BundleSubtotalPrice field to given value.

### HasBundleSubtotalPrice

`func (o *WebhookOrderNotificationBundleDetailBundleInner) HasBundleSubtotalPrice() bool`

HasBundleSubtotalPrice returns a boolean if a field has been set.

### GetOrderDetail

`func (o *WebhookOrderNotificationBundleDetailBundleInner) GetOrderDetail() []WebhookOrderNotificationBundleDetailBundleInnerOrderDetailInner`

GetOrderDetail returns the OrderDetail field if non-nil, zero value otherwise.

### GetOrderDetailOk

`func (o *WebhookOrderNotificationBundleDetailBundleInner) GetOrderDetailOk() (*[]WebhookOrderNotificationBundleDetailBundleInnerOrderDetailInner, bool)`

GetOrderDetailOk returns a tuple with the OrderDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDetail

`func (o *WebhookOrderNotificationBundleDetailBundleInner) SetOrderDetail(v []WebhookOrderNotificationBundleDetailBundleInnerOrderDetailInner)`

SetOrderDetail sets OrderDetail field to given value.

### HasOrderDetail

`func (o *WebhookOrderNotificationBundleDetailBundleInner) HasOrderDetail() bool`

HasOrderDetail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


