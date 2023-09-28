# GetSingleOrder200ResponseDataPromoOrderDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | Pointer to **int64** |  | [optional] 
**TotalCashback** | Pointer to **int64** |  | [optional] 
**TotalDiscount** | Pointer to **int64** |  | [optional] 
**TotalDiscountProduct** | Pointer to **int64** |  | [optional] 
**TotalDiscountShipping** | Pointer to **int64** |  | [optional] 
**TotalDiscountDetails** | Pointer to [**[]GetSingleOrder200ResponseDataPromoOrderDetailTotalDiscountDetailsInner**](GetSingleOrder200ResponseDataPromoOrderDetailTotalDiscountDetailsInner.md) |  | [optional] 
**SummaryPromo** | Pointer to [**[]GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner**](GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner.md) |  | [optional] 

## Methods

### NewGetSingleOrder200ResponseDataPromoOrderDetail

`func NewGetSingleOrder200ResponseDataPromoOrderDetail() *GetSingleOrder200ResponseDataPromoOrderDetail`

NewGetSingleOrder200ResponseDataPromoOrderDetail instantiates a new GetSingleOrder200ResponseDataPromoOrderDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSingleOrder200ResponseDataPromoOrderDetailWithDefaults

`func NewGetSingleOrder200ResponseDataPromoOrderDetailWithDefaults() *GetSingleOrder200ResponseDataPromoOrderDetail`

NewGetSingleOrder200ResponseDataPromoOrderDetailWithDefaults instantiates a new GetSingleOrder200ResponseDataPromoOrderDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *GetSingleOrder200ResponseDataPromoOrderDetail) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *GetSingleOrder200ResponseDataPromoOrderDetail) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *GetSingleOrder200ResponseDataPromoOrderDetail) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *GetSingleOrder200ResponseDataPromoOrderDetail) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetTotalCashback

`func (o *GetSingleOrder200ResponseDataPromoOrderDetail) GetTotalCashback() int64`

GetTotalCashback returns the TotalCashback field if non-nil, zero value otherwise.

### GetTotalCashbackOk

`func (o *GetSingleOrder200ResponseDataPromoOrderDetail) GetTotalCashbackOk() (*int64, bool)`

GetTotalCashbackOk returns a tuple with the TotalCashback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCashback

`func (o *GetSingleOrder200ResponseDataPromoOrderDetail) SetTotalCashback(v int64)`

SetTotalCashback sets TotalCashback field to given value.

### HasTotalCashback

`func (o *GetSingleOrder200ResponseDataPromoOrderDetail) HasTotalCashback() bool`

HasTotalCashback returns a boolean if a field has been set.

### GetTotalDiscount

`func (o *GetSingleOrder200ResponseDataPromoOrderDetail) GetTotalDiscount() int64`

GetTotalDiscount returns the TotalDiscount field if non-nil, zero value otherwise.

### GetTotalDiscountOk

`func (o *GetSingleOrder200ResponseDataPromoOrderDetail) GetTotalDiscountOk() (*int64, bool)`

GetTotalDiscountOk returns a tuple with the TotalDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscount

`func (o *GetSingleOrder200ResponseDataPromoOrderDetail) SetTotalDiscount(v int64)`

SetTotalDiscount sets TotalDiscount field to given value.

### HasTotalDiscount

`func (o *GetSingleOrder200ResponseDataPromoOrderDetail) HasTotalDiscount() bool`

HasTotalDiscount returns a boolean if a field has been set.

### GetTotalDiscountProduct

`func (o *GetSingleOrder200ResponseDataPromoOrderDetail) GetTotalDiscountProduct() int64`

GetTotalDiscountProduct returns the TotalDiscountProduct field if non-nil, zero value otherwise.

### GetTotalDiscountProductOk

`func (o *GetSingleOrder200ResponseDataPromoOrderDetail) GetTotalDiscountProductOk() (*int64, bool)`

GetTotalDiscountProductOk returns a tuple with the TotalDiscountProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscountProduct

`func (o *GetSingleOrder200ResponseDataPromoOrderDetail) SetTotalDiscountProduct(v int64)`

SetTotalDiscountProduct sets TotalDiscountProduct field to given value.

### HasTotalDiscountProduct

`func (o *GetSingleOrder200ResponseDataPromoOrderDetail) HasTotalDiscountProduct() bool`

HasTotalDiscountProduct returns a boolean if a field has been set.

### GetTotalDiscountShipping

`func (o *GetSingleOrder200ResponseDataPromoOrderDetail) GetTotalDiscountShipping() int64`

GetTotalDiscountShipping returns the TotalDiscountShipping field if non-nil, zero value otherwise.

### GetTotalDiscountShippingOk

`func (o *GetSingleOrder200ResponseDataPromoOrderDetail) GetTotalDiscountShippingOk() (*int64, bool)`

GetTotalDiscountShippingOk returns a tuple with the TotalDiscountShipping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscountShipping

`func (o *GetSingleOrder200ResponseDataPromoOrderDetail) SetTotalDiscountShipping(v int64)`

SetTotalDiscountShipping sets TotalDiscountShipping field to given value.

### HasTotalDiscountShipping

`func (o *GetSingleOrder200ResponseDataPromoOrderDetail) HasTotalDiscountShipping() bool`

HasTotalDiscountShipping returns a boolean if a field has been set.

### GetTotalDiscountDetails

`func (o *GetSingleOrder200ResponseDataPromoOrderDetail) GetTotalDiscountDetails() []GetSingleOrder200ResponseDataPromoOrderDetailTotalDiscountDetailsInner`

GetTotalDiscountDetails returns the TotalDiscountDetails field if non-nil, zero value otherwise.

### GetTotalDiscountDetailsOk

`func (o *GetSingleOrder200ResponseDataPromoOrderDetail) GetTotalDiscountDetailsOk() (*[]GetSingleOrder200ResponseDataPromoOrderDetailTotalDiscountDetailsInner, bool)`

GetTotalDiscountDetailsOk returns a tuple with the TotalDiscountDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscountDetails

`func (o *GetSingleOrder200ResponseDataPromoOrderDetail) SetTotalDiscountDetails(v []GetSingleOrder200ResponseDataPromoOrderDetailTotalDiscountDetailsInner)`

SetTotalDiscountDetails sets TotalDiscountDetails field to given value.

### HasTotalDiscountDetails

`func (o *GetSingleOrder200ResponseDataPromoOrderDetail) HasTotalDiscountDetails() bool`

HasTotalDiscountDetails returns a boolean if a field has been set.

### GetSummaryPromo

`func (o *GetSingleOrder200ResponseDataPromoOrderDetail) GetSummaryPromo() []GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner`

GetSummaryPromo returns the SummaryPromo field if non-nil, zero value otherwise.

### GetSummaryPromoOk

`func (o *GetSingleOrder200ResponseDataPromoOrderDetail) GetSummaryPromoOk() (*[]GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner, bool)`

GetSummaryPromoOk returns a tuple with the SummaryPromo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryPromo

`func (o *GetSingleOrder200ResponseDataPromoOrderDetail) SetSummaryPromo(v []GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner)`

SetSummaryPromo sets SummaryPromo field to given value.

### HasSummaryPromo

`func (o *GetSingleOrder200ResponseDataPromoOrderDetail) HasSummaryPromo() bool`

HasSummaryPromo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


