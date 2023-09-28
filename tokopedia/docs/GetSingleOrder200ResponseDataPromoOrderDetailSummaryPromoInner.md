# GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Promo Title | [optional] 
**IsCoupon** | Pointer to **bool** | Is promo a coupon or not | [optional] 
**ShowCashbackAmount** | Pointer to **bool** | Flag to show cashback amount (hardcoded true) | [optional] 
**ShowDiscountAmount** | Pointer to **bool** | Flag to show discount amount (hardcoded true) | [optional] 
**CashbackAmount** | Pointer to **int64** | Cashback amount in IDR | [optional] 
**CashbackPoints** | Pointer to **int64** | Cashback amount in Points currency | [optional] 
**CashbackDetails** | Pointer to **map[string]interface{}** | Details of cashback in 1 promo | [optional] 
**Type** | Pointer to **string** | Promo benefit type. (discount, cashback) | [optional] 
**DiscountAmount** | Pointer to **int64** | Discount amount | [optional] 
**DiscountDetails** | Pointer to [**[]GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInnerDiscountDetailsInner**](GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInnerDiscountDetailsInner.md) | Details of discount in 1 promo | [optional] 
**InvoiceDesc** | Pointer to **string** |  | [optional] 

## Methods

### NewGetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner

`func NewGetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner() *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner`

NewGetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner instantiates a new GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInnerWithDefaults

`func NewGetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInnerWithDefaults() *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner`

NewGetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInnerWithDefaults instantiates a new GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIsCoupon

`func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) GetIsCoupon() bool`

GetIsCoupon returns the IsCoupon field if non-nil, zero value otherwise.

### GetIsCouponOk

`func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) GetIsCouponOk() (*bool, bool)`

GetIsCouponOk returns a tuple with the IsCoupon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCoupon

`func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) SetIsCoupon(v bool)`

SetIsCoupon sets IsCoupon field to given value.

### HasIsCoupon

`func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) HasIsCoupon() bool`

HasIsCoupon returns a boolean if a field has been set.

### GetShowCashbackAmount

`func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) GetShowCashbackAmount() bool`

GetShowCashbackAmount returns the ShowCashbackAmount field if non-nil, zero value otherwise.

### GetShowCashbackAmountOk

`func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) GetShowCashbackAmountOk() (*bool, bool)`

GetShowCashbackAmountOk returns a tuple with the ShowCashbackAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowCashbackAmount

`func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) SetShowCashbackAmount(v bool)`

SetShowCashbackAmount sets ShowCashbackAmount field to given value.

### HasShowCashbackAmount

`func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) HasShowCashbackAmount() bool`

HasShowCashbackAmount returns a boolean if a field has been set.

### GetShowDiscountAmount

`func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) GetShowDiscountAmount() bool`

GetShowDiscountAmount returns the ShowDiscountAmount field if non-nil, zero value otherwise.

### GetShowDiscountAmountOk

`func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) GetShowDiscountAmountOk() (*bool, bool)`

GetShowDiscountAmountOk returns a tuple with the ShowDiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowDiscountAmount

`func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) SetShowDiscountAmount(v bool)`

SetShowDiscountAmount sets ShowDiscountAmount field to given value.

### HasShowDiscountAmount

`func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) HasShowDiscountAmount() bool`

HasShowDiscountAmount returns a boolean if a field has been set.

### GetCashbackAmount

`func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) GetCashbackAmount() int64`

GetCashbackAmount returns the CashbackAmount field if non-nil, zero value otherwise.

### GetCashbackAmountOk

`func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) GetCashbackAmountOk() (*int64, bool)`

GetCashbackAmountOk returns a tuple with the CashbackAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashbackAmount

`func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) SetCashbackAmount(v int64)`

SetCashbackAmount sets CashbackAmount field to given value.

### HasCashbackAmount

`func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) HasCashbackAmount() bool`

HasCashbackAmount returns a boolean if a field has been set.

### GetCashbackPoints

`func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) GetCashbackPoints() int64`

GetCashbackPoints returns the CashbackPoints field if non-nil, zero value otherwise.

### GetCashbackPointsOk

`func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) GetCashbackPointsOk() (*int64, bool)`

GetCashbackPointsOk returns a tuple with the CashbackPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashbackPoints

`func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) SetCashbackPoints(v int64)`

SetCashbackPoints sets CashbackPoints field to given value.

### HasCashbackPoints

`func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) HasCashbackPoints() bool`

HasCashbackPoints returns a boolean if a field has been set.

### GetCashbackDetails

`func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) GetCashbackDetails() map[string]interface{}`

GetCashbackDetails returns the CashbackDetails field if non-nil, zero value otherwise.

### GetCashbackDetailsOk

`func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) GetCashbackDetailsOk() (*map[string]interface{}, bool)`

GetCashbackDetailsOk returns a tuple with the CashbackDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashbackDetails

`func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) SetCashbackDetails(v map[string]interface{})`

SetCashbackDetails sets CashbackDetails field to given value.

### HasCashbackDetails

`func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) HasCashbackDetails() bool`

HasCashbackDetails returns a boolean if a field has been set.

### SetCashbackDetailsNil

`func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) SetCashbackDetailsNil(b bool)`

 SetCashbackDetailsNil sets the value for CashbackDetails to be an explicit nil

### UnsetCashbackDetails
`func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) UnsetCashbackDetails()`

UnsetCashbackDetails ensures that no value is present for CashbackDetails, not even an explicit nil
### GetType

`func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDiscountAmount

`func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) GetDiscountAmount() int64`

GetDiscountAmount returns the DiscountAmount field if non-nil, zero value otherwise.

### GetDiscountAmountOk

`func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) GetDiscountAmountOk() (*int64, bool)`

GetDiscountAmountOk returns a tuple with the DiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmount

`func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) SetDiscountAmount(v int64)`

SetDiscountAmount sets DiscountAmount field to given value.

### HasDiscountAmount

`func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) HasDiscountAmount() bool`

HasDiscountAmount returns a boolean if a field has been set.

### GetDiscountDetails

`func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) GetDiscountDetails() []GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInnerDiscountDetailsInner`

GetDiscountDetails returns the DiscountDetails field if non-nil, zero value otherwise.

### GetDiscountDetailsOk

`func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) GetDiscountDetailsOk() (*[]GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInnerDiscountDetailsInner, bool)`

GetDiscountDetailsOk returns a tuple with the DiscountDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountDetails

`func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) SetDiscountDetails(v []GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInnerDiscountDetailsInner)`

SetDiscountDetails sets DiscountDetails field to given value.

### HasDiscountDetails

`func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) HasDiscountDetails() bool`

HasDiscountDetails returns a boolean if a field has been set.

### GetInvoiceDesc

`func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) GetInvoiceDesc() string`

GetInvoiceDesc returns the InvoiceDesc field if non-nil, zero value otherwise.

### GetInvoiceDescOk

`func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) GetInvoiceDescOk() (*string, bool)`

GetInvoiceDescOk returns a tuple with the InvoiceDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceDesc

`func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) SetInvoiceDesc(v string)`

SetInvoiceDesc sets InvoiceDesc field to given value.

### HasInvoiceDesc

`func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInner) HasInvoiceDesc() bool`

HasInvoiceDesc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


