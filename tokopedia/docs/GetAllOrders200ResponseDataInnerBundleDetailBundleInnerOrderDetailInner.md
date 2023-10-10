# GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderDtlId** | Pointer to **int64** | Bundle Order Detail Unique Identifier | [optional] 
**OrderId** | Pointer to **int64** | Bundle Order Unique Identifier | [optional] 
**ProductId** | Pointer to **int64** | Product Unique Identifier | [optional] 
**ProductName** | Pointer to **string** | Product Name | [optional] 
**ProductDesc** | Pointer to **string** | Product Descirption | [optional] 
**Quantity** | Pointer to **int64** | Product Order Quantity | [optional] 
**ProductPrice** | Pointer to **float64** | Product Price | [optional] 
**ProductWeight** | Pointer to **float64** | Product Weight | [optional] 
**TotalWeight** | Pointer to **float64** | Product Order Total Weight | [optional] 
**SubtotalPrice** | Pointer to **float64** | Product Order Total Price | [optional] 
**Notes** | Pointer to **string** | Product Order Notes | [optional] 
**Finsurance** | Pointer to **int64** |  | [optional] 
**Returnable** | Pointer to **int64** | Is returnable? | [optional] 
**ChildCatId** | Pointer to **int64** | Category Unique Identifier | [optional] 
**CurrencyId** | Pointer to **int64** | Currency Unique Identifier | [optional] 
**InsurancePrice** | Pointer to **float64** | Product Insurance Price | [optional] 
**NormalPrice** | Pointer to **float64** | Product Normal Price | [optional] 
**CurrencyRate** | Pointer to **float64** |  | [optional] 
**ProdPic** | Pointer to **string** |  | [optional] 
**MinOrder** | Pointer to **int64** | Minimum Product Order | [optional] 
**MustInsurance** | Pointer to **int64** | Is product must include insurance? | [optional] 
**Condition** | Pointer to **int64** | Product Condition, 1 &#x3D; New, 2 &#x3D; Used | [optional] 
**CampaignId** | Pointer to **int64** | Product Campaign Unique Identifier | [optional] 
**Sku** | Pointer to **string** | Product SKU | [optional] 
**IsSlashPrice** | Pointer to **bool** | Is product has slash price? | [optional] 
**OmsDetailData** | Pointer to **string** |  | [optional] 
**Thumbnail** | Pointer to **string** | Product Thumbnail URL | [optional] 
**BundleId** | Pointer to **int64** | Bundle Unique Identifier | [optional] 
**BundleVariantId** | Pointer to **string** | Bundle Variant Unique Identifier | [optional] 

## Methods

### NewGetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner

`func NewGetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner() *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner`

NewGetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner instantiates a new GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInnerWithDefaults

`func NewGetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInnerWithDefaults() *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner`

NewGetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInnerWithDefaults instantiates a new GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderDtlId

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) GetOrderDtlId() int64`

GetOrderDtlId returns the OrderDtlId field if non-nil, zero value otherwise.

### GetOrderDtlIdOk

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) GetOrderDtlIdOk() (*int64, bool)`

GetOrderDtlIdOk returns a tuple with the OrderDtlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDtlId

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) SetOrderDtlId(v int64)`

SetOrderDtlId sets OrderDtlId field to given value.

### HasOrderDtlId

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) HasOrderDtlId() bool`

HasOrderDtlId returns a boolean if a field has been set.

### GetOrderId

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetProductId

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) GetProductId() int64`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) GetProductIdOk() (*int64, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) SetProductId(v int64)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetProductName

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetProductDesc

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) GetProductDesc() string`

GetProductDesc returns the ProductDesc field if non-nil, zero value otherwise.

### GetProductDescOk

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) GetProductDescOk() (*string, bool)`

GetProductDescOk returns a tuple with the ProductDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductDesc

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) SetProductDesc(v string)`

SetProductDesc sets ProductDesc field to given value.

### HasProductDesc

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) HasProductDesc() bool`

HasProductDesc returns a boolean if a field has been set.

### GetQuantity

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetProductPrice

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) GetProductPrice() float64`

GetProductPrice returns the ProductPrice field if non-nil, zero value otherwise.

### GetProductPriceOk

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) GetProductPriceOk() (*float64, bool)`

GetProductPriceOk returns a tuple with the ProductPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductPrice

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) SetProductPrice(v float64)`

SetProductPrice sets ProductPrice field to given value.

### HasProductPrice

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) HasProductPrice() bool`

HasProductPrice returns a boolean if a field has been set.

### GetProductWeight

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) GetProductWeight() float64`

GetProductWeight returns the ProductWeight field if non-nil, zero value otherwise.

### GetProductWeightOk

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) GetProductWeightOk() (*float64, bool)`

GetProductWeightOk returns a tuple with the ProductWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductWeight

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) SetProductWeight(v float64)`

SetProductWeight sets ProductWeight field to given value.

### HasProductWeight

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) HasProductWeight() bool`

HasProductWeight returns a boolean if a field has been set.

### GetTotalWeight

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) GetTotalWeight() float64`

GetTotalWeight returns the TotalWeight field if non-nil, zero value otherwise.

### GetTotalWeightOk

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) GetTotalWeightOk() (*float64, bool)`

GetTotalWeightOk returns a tuple with the TotalWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWeight

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) SetTotalWeight(v float64)`

SetTotalWeight sets TotalWeight field to given value.

### HasTotalWeight

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) HasTotalWeight() bool`

HasTotalWeight returns a boolean if a field has been set.

### GetSubtotalPrice

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) GetSubtotalPrice() float64`

GetSubtotalPrice returns the SubtotalPrice field if non-nil, zero value otherwise.

### GetSubtotalPriceOk

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) GetSubtotalPriceOk() (*float64, bool)`

GetSubtotalPriceOk returns a tuple with the SubtotalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotalPrice

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) SetSubtotalPrice(v float64)`

SetSubtotalPrice sets SubtotalPrice field to given value.

### HasSubtotalPrice

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) HasSubtotalPrice() bool`

HasSubtotalPrice returns a boolean if a field has been set.

### GetNotes

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetFinsurance

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) GetFinsurance() int64`

GetFinsurance returns the Finsurance field if non-nil, zero value otherwise.

### GetFinsuranceOk

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) GetFinsuranceOk() (*int64, bool)`

GetFinsuranceOk returns a tuple with the Finsurance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinsurance

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) SetFinsurance(v int64)`

SetFinsurance sets Finsurance field to given value.

### HasFinsurance

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) HasFinsurance() bool`

HasFinsurance returns a boolean if a field has been set.

### GetReturnable

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) GetReturnable() int64`

GetReturnable returns the Returnable field if non-nil, zero value otherwise.

### GetReturnableOk

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) GetReturnableOk() (*int64, bool)`

GetReturnableOk returns a tuple with the Returnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnable

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) SetReturnable(v int64)`

SetReturnable sets Returnable field to given value.

### HasReturnable

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) HasReturnable() bool`

HasReturnable returns a boolean if a field has been set.

### GetChildCatId

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) GetChildCatId() int64`

GetChildCatId returns the ChildCatId field if non-nil, zero value otherwise.

### GetChildCatIdOk

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) GetChildCatIdOk() (*int64, bool)`

GetChildCatIdOk returns a tuple with the ChildCatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildCatId

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) SetChildCatId(v int64)`

SetChildCatId sets ChildCatId field to given value.

### HasChildCatId

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) HasChildCatId() bool`

HasChildCatId returns a boolean if a field has been set.

### GetCurrencyId

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) GetCurrencyId() int64`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) GetCurrencyIdOk() (*int64, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) SetCurrencyId(v int64)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### GetInsurancePrice

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) GetInsurancePrice() float64`

GetInsurancePrice returns the InsurancePrice field if non-nil, zero value otherwise.

### GetInsurancePriceOk

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) GetInsurancePriceOk() (*float64, bool)`

GetInsurancePriceOk returns a tuple with the InsurancePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsurancePrice

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) SetInsurancePrice(v float64)`

SetInsurancePrice sets InsurancePrice field to given value.

### HasInsurancePrice

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) HasInsurancePrice() bool`

HasInsurancePrice returns a boolean if a field has been set.

### GetNormalPrice

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) GetNormalPrice() float64`

GetNormalPrice returns the NormalPrice field if non-nil, zero value otherwise.

### GetNormalPriceOk

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) GetNormalPriceOk() (*float64, bool)`

GetNormalPriceOk returns a tuple with the NormalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalPrice

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) SetNormalPrice(v float64)`

SetNormalPrice sets NormalPrice field to given value.

### HasNormalPrice

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) HasNormalPrice() bool`

HasNormalPrice returns a boolean if a field has been set.

### GetCurrencyRate

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) GetCurrencyRate() float64`

GetCurrencyRate returns the CurrencyRate field if non-nil, zero value otherwise.

### GetCurrencyRateOk

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) GetCurrencyRateOk() (*float64, bool)`

GetCurrencyRateOk returns a tuple with the CurrencyRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyRate

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) SetCurrencyRate(v float64)`

SetCurrencyRate sets CurrencyRate field to given value.

### HasCurrencyRate

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) HasCurrencyRate() bool`

HasCurrencyRate returns a boolean if a field has been set.

### GetProdPic

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) GetProdPic() string`

GetProdPic returns the ProdPic field if non-nil, zero value otherwise.

### GetProdPicOk

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) GetProdPicOk() (*string, bool)`

GetProdPicOk returns a tuple with the ProdPic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProdPic

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) SetProdPic(v string)`

SetProdPic sets ProdPic field to given value.

### HasProdPic

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) HasProdPic() bool`

HasProdPic returns a boolean if a field has been set.

### GetMinOrder

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) GetMinOrder() int64`

GetMinOrder returns the MinOrder field if non-nil, zero value otherwise.

### GetMinOrderOk

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) GetMinOrderOk() (*int64, bool)`

GetMinOrderOk returns a tuple with the MinOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinOrder

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) SetMinOrder(v int64)`

SetMinOrder sets MinOrder field to given value.

### HasMinOrder

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) HasMinOrder() bool`

HasMinOrder returns a boolean if a field has been set.

### GetMustInsurance

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) GetMustInsurance() int64`

GetMustInsurance returns the MustInsurance field if non-nil, zero value otherwise.

### GetMustInsuranceOk

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) GetMustInsuranceOk() (*int64, bool)`

GetMustInsuranceOk returns a tuple with the MustInsurance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMustInsurance

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) SetMustInsurance(v int64)`

SetMustInsurance sets MustInsurance field to given value.

### HasMustInsurance

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) HasMustInsurance() bool`

HasMustInsurance returns a boolean if a field has been set.

### GetCondition

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) GetCondition() int64`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) GetConditionOk() (*int64, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) SetCondition(v int64)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetCampaignId

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) GetCampaignId() int64`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) GetCampaignIdOk() (*int64, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) SetCampaignId(v int64)`

SetCampaignId sets CampaignId field to given value.

### HasCampaignId

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) HasCampaignId() bool`

HasCampaignId returns a boolean if a field has been set.

### GetSku

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetIsSlashPrice

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) GetIsSlashPrice() bool`

GetIsSlashPrice returns the IsSlashPrice field if non-nil, zero value otherwise.

### GetIsSlashPriceOk

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) GetIsSlashPriceOk() (*bool, bool)`

GetIsSlashPriceOk returns a tuple with the IsSlashPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSlashPrice

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) SetIsSlashPrice(v bool)`

SetIsSlashPrice sets IsSlashPrice field to given value.

### HasIsSlashPrice

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) HasIsSlashPrice() bool`

HasIsSlashPrice returns a boolean if a field has been set.

### GetOmsDetailData

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) GetOmsDetailData() string`

GetOmsDetailData returns the OmsDetailData field if non-nil, zero value otherwise.

### GetOmsDetailDataOk

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) GetOmsDetailDataOk() (*string, bool)`

GetOmsDetailDataOk returns a tuple with the OmsDetailData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmsDetailData

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) SetOmsDetailData(v string)`

SetOmsDetailData sets OmsDetailData field to given value.

### HasOmsDetailData

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) HasOmsDetailData() bool`

HasOmsDetailData returns a boolean if a field has been set.

### GetThumbnail

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) GetThumbnail() string`

GetThumbnail returns the Thumbnail field if non-nil, zero value otherwise.

### GetThumbnailOk

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) GetThumbnailOk() (*string, bool)`

GetThumbnailOk returns a tuple with the Thumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnail

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) SetThumbnail(v string)`

SetThumbnail sets Thumbnail field to given value.

### HasThumbnail

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) HasThumbnail() bool`

HasThumbnail returns a boolean if a field has been set.

### GetBundleId

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) GetBundleId() int64`

GetBundleId returns the BundleId field if non-nil, zero value otherwise.

### GetBundleIdOk

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) GetBundleIdOk() (*int64, bool)`

GetBundleIdOk returns a tuple with the BundleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleId

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) SetBundleId(v int64)`

SetBundleId sets BundleId field to given value.

### HasBundleId

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) HasBundleId() bool`

HasBundleId returns a boolean if a field has been set.

### GetBundleVariantId

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) GetBundleVariantId() string`

GetBundleVariantId returns the BundleVariantId field if non-nil, zero value otherwise.

### GetBundleVariantIdOk

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) GetBundleVariantIdOk() (*string, bool)`

GetBundleVariantIdOk returns a tuple with the BundleVariantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleVariantId

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) SetBundleVariantId(v string)`

SetBundleVariantId sets BundleVariantId field to given value.

### HasBundleVariantId

`func (o *GetAllOrders200ResponseDataInnerBundleDetailBundleInnerOrderDetailInner) HasBundleVariantId() bool`

HasBundleVariantId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


