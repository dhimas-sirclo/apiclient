# GetAllOrders200ResponseDataInnerPromoOrderDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | Pointer to **int64** | Order Unique Identifier | [optional] 
**TotalCashback** | Pointer to **int64** | Total Cashback Amount | [optional] 
**TotalDiscount** | Pointer to **int64** | Total Discount Amount | [optional] 
**TotalDiscountProduct** | Pointer to **int64** | Total Product Discount Amount | [optional] 
**TotalDiscountShipping** | Pointer to **int64** | Total Shipping Discount Amount | [optional] 
**TotalDiscountDetails** | Pointer to [**[]GetAllOrders200ResponseDataInnerPromoOrderDetailTotalDiscountDetailsInner**](GetAllOrders200ResponseDataInnerPromoOrderDetailTotalDiscountDetailsInner.md) |  | [optional] 
**SummaryPromo** | Pointer to [**[]GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner**](GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner.md) |  | [optional] 

## Methods

### NewGetAllOrders200ResponseDataInnerPromoOrderDetail

`func NewGetAllOrders200ResponseDataInnerPromoOrderDetail() *GetAllOrders200ResponseDataInnerPromoOrderDetail`

NewGetAllOrders200ResponseDataInnerPromoOrderDetail instantiates a new GetAllOrders200ResponseDataInnerPromoOrderDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllOrders200ResponseDataInnerPromoOrderDetailWithDefaults

`func NewGetAllOrders200ResponseDataInnerPromoOrderDetailWithDefaults() *GetAllOrders200ResponseDataInnerPromoOrderDetail`

NewGetAllOrders200ResponseDataInnerPromoOrderDetailWithDefaults instantiates a new GetAllOrders200ResponseDataInnerPromoOrderDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *GetAllOrders200ResponseDataInnerPromoOrderDetail) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *GetAllOrders200ResponseDataInnerPromoOrderDetail) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *GetAllOrders200ResponseDataInnerPromoOrderDetail) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *GetAllOrders200ResponseDataInnerPromoOrderDetail) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetTotalCashback

`func (o *GetAllOrders200ResponseDataInnerPromoOrderDetail) GetTotalCashback() int64`

GetTotalCashback returns the TotalCashback field if non-nil, zero value otherwise.

### GetTotalCashbackOk

`func (o *GetAllOrders200ResponseDataInnerPromoOrderDetail) GetTotalCashbackOk() (*int64, bool)`

GetTotalCashbackOk returns a tuple with the TotalCashback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCashback

`func (o *GetAllOrders200ResponseDataInnerPromoOrderDetail) SetTotalCashback(v int64)`

SetTotalCashback sets TotalCashback field to given value.

### HasTotalCashback

`func (o *GetAllOrders200ResponseDataInnerPromoOrderDetail) HasTotalCashback() bool`

HasTotalCashback returns a boolean if a field has been set.

### GetTotalDiscount

`func (o *GetAllOrders200ResponseDataInnerPromoOrderDetail) GetTotalDiscount() int64`

GetTotalDiscount returns the TotalDiscount field if non-nil, zero value otherwise.

### GetTotalDiscountOk

`func (o *GetAllOrders200ResponseDataInnerPromoOrderDetail) GetTotalDiscountOk() (*int64, bool)`

GetTotalDiscountOk returns a tuple with the TotalDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscount

`func (o *GetAllOrders200ResponseDataInnerPromoOrderDetail) SetTotalDiscount(v int64)`

SetTotalDiscount sets TotalDiscount field to given value.

### HasTotalDiscount

`func (o *GetAllOrders200ResponseDataInnerPromoOrderDetail) HasTotalDiscount() bool`

HasTotalDiscount returns a boolean if a field has been set.

### GetTotalDiscountProduct

`func (o *GetAllOrders200ResponseDataInnerPromoOrderDetail) GetTotalDiscountProduct() int64`

GetTotalDiscountProduct returns the TotalDiscountProduct field if non-nil, zero value otherwise.

### GetTotalDiscountProductOk

`func (o *GetAllOrders200ResponseDataInnerPromoOrderDetail) GetTotalDiscountProductOk() (*int64, bool)`

GetTotalDiscountProductOk returns a tuple with the TotalDiscountProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscountProduct

`func (o *GetAllOrders200ResponseDataInnerPromoOrderDetail) SetTotalDiscountProduct(v int64)`

SetTotalDiscountProduct sets TotalDiscountProduct field to given value.

### HasTotalDiscountProduct

`func (o *GetAllOrders200ResponseDataInnerPromoOrderDetail) HasTotalDiscountProduct() bool`

HasTotalDiscountProduct returns a boolean if a field has been set.

### GetTotalDiscountShipping

`func (o *GetAllOrders200ResponseDataInnerPromoOrderDetail) GetTotalDiscountShipping() int64`

GetTotalDiscountShipping returns the TotalDiscountShipping field if non-nil, zero value otherwise.

### GetTotalDiscountShippingOk

`func (o *GetAllOrders200ResponseDataInnerPromoOrderDetail) GetTotalDiscountShippingOk() (*int64, bool)`

GetTotalDiscountShippingOk returns a tuple with the TotalDiscountShipping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscountShipping

`func (o *GetAllOrders200ResponseDataInnerPromoOrderDetail) SetTotalDiscountShipping(v int64)`

SetTotalDiscountShipping sets TotalDiscountShipping field to given value.

### HasTotalDiscountShipping

`func (o *GetAllOrders200ResponseDataInnerPromoOrderDetail) HasTotalDiscountShipping() bool`

HasTotalDiscountShipping returns a boolean if a field has been set.

### GetTotalDiscountDetails

`func (o *GetAllOrders200ResponseDataInnerPromoOrderDetail) GetTotalDiscountDetails() []GetAllOrders200ResponseDataInnerPromoOrderDetailTotalDiscountDetailsInner`

GetTotalDiscountDetails returns the TotalDiscountDetails field if non-nil, zero value otherwise.

### GetTotalDiscountDetailsOk

`func (o *GetAllOrders200ResponseDataInnerPromoOrderDetail) GetTotalDiscountDetailsOk() (*[]GetAllOrders200ResponseDataInnerPromoOrderDetailTotalDiscountDetailsInner, bool)`

GetTotalDiscountDetailsOk returns a tuple with the TotalDiscountDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscountDetails

`func (o *GetAllOrders200ResponseDataInnerPromoOrderDetail) SetTotalDiscountDetails(v []GetAllOrders200ResponseDataInnerPromoOrderDetailTotalDiscountDetailsInner)`

SetTotalDiscountDetails sets TotalDiscountDetails field to given value.

### HasTotalDiscountDetails

`func (o *GetAllOrders200ResponseDataInnerPromoOrderDetail) HasTotalDiscountDetails() bool`

HasTotalDiscountDetails returns a boolean if a field has been set.

### GetSummaryPromo

`func (o *GetAllOrders200ResponseDataInnerPromoOrderDetail) GetSummaryPromo() []GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner`

GetSummaryPromo returns the SummaryPromo field if non-nil, zero value otherwise.

### GetSummaryPromoOk

`func (o *GetAllOrders200ResponseDataInnerPromoOrderDetail) GetSummaryPromoOk() (*[]GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner, bool)`

GetSummaryPromoOk returns a tuple with the SummaryPromo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryPromo

`func (o *GetAllOrders200ResponseDataInnerPromoOrderDetail) SetSummaryPromo(v []GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner)`

SetSummaryPromo sets SummaryPromo field to given value.

### HasSummaryPromo

`func (o *GetAllOrders200ResponseDataInnerPromoOrderDetail) HasSummaryPromo() bool`

HasSummaryPromo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


