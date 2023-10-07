# GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Promo Name | [optional] 
**IsCoupon** | Pointer to **bool** | Promo Name | [optional] 
**ShowCashbackAmount** | Pointer to **bool** | Show cashback amount? | [optional] 
**ShowDiscountAmount** | Pointer to **bool** | Show discount amount? | [optional] 
**CashbackAmount** | Pointer to **int64** | Cashback Amount | [optional] 
**CashbackPoints** | Pointer to **int64** | Cashback Point | [optional] 
**Type** | Pointer to **string** | Promo Type | [optional] 
**DiscountAmount** | Pointer to **int64** | Discount Amount | [optional] 
**DiscountDetails** | Pointer to [**[]GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInnerDiscountDetailsInner**](GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInnerDiscountDetailsInner.md) |  | [optional] 
**InvoiceDesc** | Pointer to **string** | Discount Invoice Description | [optional] 

## Methods

### NewGetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner

`func NewGetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner() *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner`

NewGetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner instantiates a new GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInnerWithDefaults

`func NewGetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInnerWithDefaults() *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner`

NewGetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInnerWithDefaults instantiates a new GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIsCoupon

`func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) GetIsCoupon() bool`

GetIsCoupon returns the IsCoupon field if non-nil, zero value otherwise.

### GetIsCouponOk

`func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) GetIsCouponOk() (*bool, bool)`

GetIsCouponOk returns a tuple with the IsCoupon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCoupon

`func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) SetIsCoupon(v bool)`

SetIsCoupon sets IsCoupon field to given value.

### HasIsCoupon

`func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) HasIsCoupon() bool`

HasIsCoupon returns a boolean if a field has been set.

### GetShowCashbackAmount

`func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) GetShowCashbackAmount() bool`

GetShowCashbackAmount returns the ShowCashbackAmount field if non-nil, zero value otherwise.

### GetShowCashbackAmountOk

`func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) GetShowCashbackAmountOk() (*bool, bool)`

GetShowCashbackAmountOk returns a tuple with the ShowCashbackAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowCashbackAmount

`func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) SetShowCashbackAmount(v bool)`

SetShowCashbackAmount sets ShowCashbackAmount field to given value.

### HasShowCashbackAmount

`func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) HasShowCashbackAmount() bool`

HasShowCashbackAmount returns a boolean if a field has been set.

### GetShowDiscountAmount

`func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) GetShowDiscountAmount() bool`

GetShowDiscountAmount returns the ShowDiscountAmount field if non-nil, zero value otherwise.

### GetShowDiscountAmountOk

`func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) GetShowDiscountAmountOk() (*bool, bool)`

GetShowDiscountAmountOk returns a tuple with the ShowDiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowDiscountAmount

`func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) SetShowDiscountAmount(v bool)`

SetShowDiscountAmount sets ShowDiscountAmount field to given value.

### HasShowDiscountAmount

`func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) HasShowDiscountAmount() bool`

HasShowDiscountAmount returns a boolean if a field has been set.

### GetCashbackAmount

`func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) GetCashbackAmount() int64`

GetCashbackAmount returns the CashbackAmount field if non-nil, zero value otherwise.

### GetCashbackAmountOk

`func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) GetCashbackAmountOk() (*int64, bool)`

GetCashbackAmountOk returns a tuple with the CashbackAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashbackAmount

`func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) SetCashbackAmount(v int64)`

SetCashbackAmount sets CashbackAmount field to given value.

### HasCashbackAmount

`func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) HasCashbackAmount() bool`

HasCashbackAmount returns a boolean if a field has been set.

### GetCashbackPoints

`func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) GetCashbackPoints() int64`

GetCashbackPoints returns the CashbackPoints field if non-nil, zero value otherwise.

### GetCashbackPointsOk

`func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) GetCashbackPointsOk() (*int64, bool)`

GetCashbackPointsOk returns a tuple with the CashbackPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashbackPoints

`func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) SetCashbackPoints(v int64)`

SetCashbackPoints sets CashbackPoints field to given value.

### HasCashbackPoints

`func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) HasCashbackPoints() bool`

HasCashbackPoints returns a boolean if a field has been set.

### GetType

`func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDiscountAmount

`func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) GetDiscountAmount() int64`

GetDiscountAmount returns the DiscountAmount field if non-nil, zero value otherwise.

### GetDiscountAmountOk

`func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) GetDiscountAmountOk() (*int64, bool)`

GetDiscountAmountOk returns a tuple with the DiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmount

`func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) SetDiscountAmount(v int64)`

SetDiscountAmount sets DiscountAmount field to given value.

### HasDiscountAmount

`func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) HasDiscountAmount() bool`

HasDiscountAmount returns a boolean if a field has been set.

### GetDiscountDetails

`func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) GetDiscountDetails() []GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInnerDiscountDetailsInner`

GetDiscountDetails returns the DiscountDetails field if non-nil, zero value otherwise.

### GetDiscountDetailsOk

`func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) GetDiscountDetailsOk() (*[]GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInnerDiscountDetailsInner, bool)`

GetDiscountDetailsOk returns a tuple with the DiscountDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountDetails

`func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) SetDiscountDetails(v []GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInnerDiscountDetailsInner)`

SetDiscountDetails sets DiscountDetails field to given value.

### HasDiscountDetails

`func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) HasDiscountDetails() bool`

HasDiscountDetails returns a boolean if a field has been set.

### GetInvoiceDesc

`func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) GetInvoiceDesc() string`

GetInvoiceDesc returns the InvoiceDesc field if non-nil, zero value otherwise.

### GetInvoiceDescOk

`func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) GetInvoiceDescOk() (*string, bool)`

GetInvoiceDescOk returns a tuple with the InvoiceDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceDesc

`func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) SetInvoiceDesc(v string)`

SetInvoiceDesc sets InvoiceDesc field to given value.

### HasInvoiceDesc

`func (o *GetAllOrders200ResponseDataInnerPromoOrderDetailSummaryPromoInner) HasInvoiceDesc() bool`

HasInvoiceDesc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


