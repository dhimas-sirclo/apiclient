# GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInnerDiscountDetailsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int64** | Discount amount per discount type | [optional] 
**Type** | Pointer to **string** | Discount type. (discount_shipping, discount_product) | [optional] 
**BudgetDetails** | Pointer to [**[]GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInnerDiscountDetailsInnerBudgetDetailsInner**](GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInnerDiscountDetailsInnerBudgetDetailsInner.md) | Details of budget in 1 promo with breakdown of benefit amount | [optional] 

## Methods

### NewGetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInnerDiscountDetailsInner

`func NewGetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInnerDiscountDetailsInner() *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInnerDiscountDetailsInner`

NewGetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInnerDiscountDetailsInner instantiates a new GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInnerDiscountDetailsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInnerDiscountDetailsInnerWithDefaults

`func NewGetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInnerDiscountDetailsInnerWithDefaults() *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInnerDiscountDetailsInner`

NewGetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInnerDiscountDetailsInnerWithDefaults instantiates a new GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInnerDiscountDetailsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInnerDiscountDetailsInner) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInnerDiscountDetailsInner) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInnerDiscountDetailsInner) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInnerDiscountDetailsInner) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetType

`func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInnerDiscountDetailsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInnerDiscountDetailsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInnerDiscountDetailsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInnerDiscountDetailsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetBudgetDetails

`func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInnerDiscountDetailsInner) GetBudgetDetails() []GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInnerDiscountDetailsInnerBudgetDetailsInner`

GetBudgetDetails returns the BudgetDetails field if non-nil, zero value otherwise.

### GetBudgetDetailsOk

`func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInnerDiscountDetailsInner) GetBudgetDetailsOk() (*[]GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInnerDiscountDetailsInnerBudgetDetailsInner, bool)`

GetBudgetDetailsOk returns a tuple with the BudgetDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetDetails

`func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInnerDiscountDetailsInner) SetBudgetDetails(v []GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInnerDiscountDetailsInnerBudgetDetailsInner)`

SetBudgetDetails sets BudgetDetails field to given value.

### HasBudgetDetails

`func (o *GetSingleOrder200ResponseDataPromoOrderDetailSummaryPromoInnerDiscountDetailsInner) HasBudgetDetails() bool`

HasBudgetDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


