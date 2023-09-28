# GetSingleOrder200ResponseDataOrderInfoOrderDetailInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderDetailId** | Pointer to **int64** | Order Detail Unique Identifier | [optional] 
**ProductId** | Pointer to **int64** | Product Unique Identifier | [optional] 
**ProductName** | Pointer to **string** | Product Name | [optional] 
**ProductDescPdp** | Pointer to **string** | Product Description PDP | [optional] 
**ProductDescAtc** | Pointer to **string** | Product Description ATC | [optional] 
**ProductPrice** | Pointer to **int64** | Product Price | [optional] 
**SubtotalPrice** | Pointer to **int64** | Product Order Total Price | [optional] 
**Weight** | Pointer to **float64** | Product Weight | [optional] 
**TotalWeight** | Pointer to **float64** | Product Order Total Weight | [optional] 
**Quantity** | Pointer to **int64** | Product Order Quantity | [optional] 
**QuantityDeliver** | Pointer to **int64** | Product Order Quantity Deliver Count | [optional] 
**QuantityReject** | Pointer to **int64** | Product Order Quantity Reject Count | [optional] 
**IsFreeReturns** | Pointer to **bool** | Is free returns? | [optional] 
**InsurancePrice** | Pointer to **int64** | Product Order Insurance Price | [optional] 
**NormalPrice** | Pointer to **int64** | Product Order Normal Price | [optional] 
**CurrencyId** | Pointer to **int64** | Currency Unique Identifier | [optional] 
**CurrencyRate** | Pointer to **int64** | Currency Rate | [optional] 
**MinOrder** | Pointer to **int64** | Product Minimum Order | [optional] 
**ChildCatId** | Pointer to **int64** | Product Child Category Unique Identifier | [optional] 
**CampaignId** | Pointer to **string** | Campaign Unique Identifier | [optional] 
**ProductPicture** | Pointer to **string** | Product Picture URL | [optional] 
**SnapshotUrl** | Pointer to **string** | Product Snapshot URL | [optional] 
**Sku** | Pointer to **string** | SKU | [optional] 

## Methods

### NewGetSingleOrder200ResponseDataOrderInfoOrderDetailInner

`func NewGetSingleOrder200ResponseDataOrderInfoOrderDetailInner() *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner`

NewGetSingleOrder200ResponseDataOrderInfoOrderDetailInner instantiates a new GetSingleOrder200ResponseDataOrderInfoOrderDetailInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSingleOrder200ResponseDataOrderInfoOrderDetailInnerWithDefaults

`func NewGetSingleOrder200ResponseDataOrderInfoOrderDetailInnerWithDefaults() *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner`

NewGetSingleOrder200ResponseDataOrderInfoOrderDetailInnerWithDefaults instantiates a new GetSingleOrder200ResponseDataOrderInfoOrderDetailInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderDetailId

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) GetOrderDetailId() int64`

GetOrderDetailId returns the OrderDetailId field if non-nil, zero value otherwise.

### GetOrderDetailIdOk

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) GetOrderDetailIdOk() (*int64, bool)`

GetOrderDetailIdOk returns a tuple with the OrderDetailId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDetailId

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) SetOrderDetailId(v int64)`

SetOrderDetailId sets OrderDetailId field to given value.

### HasOrderDetailId

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) HasOrderDetailId() bool`

HasOrderDetailId returns a boolean if a field has been set.

### GetProductId

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) GetProductId() int64`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) GetProductIdOk() (*int64, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) SetProductId(v int64)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetProductName

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetProductDescPdp

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) GetProductDescPdp() string`

GetProductDescPdp returns the ProductDescPdp field if non-nil, zero value otherwise.

### GetProductDescPdpOk

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) GetProductDescPdpOk() (*string, bool)`

GetProductDescPdpOk returns a tuple with the ProductDescPdp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductDescPdp

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) SetProductDescPdp(v string)`

SetProductDescPdp sets ProductDescPdp field to given value.

### HasProductDescPdp

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) HasProductDescPdp() bool`

HasProductDescPdp returns a boolean if a field has been set.

### GetProductDescAtc

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) GetProductDescAtc() string`

GetProductDescAtc returns the ProductDescAtc field if non-nil, zero value otherwise.

### GetProductDescAtcOk

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) GetProductDescAtcOk() (*string, bool)`

GetProductDescAtcOk returns a tuple with the ProductDescAtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductDescAtc

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) SetProductDescAtc(v string)`

SetProductDescAtc sets ProductDescAtc field to given value.

### HasProductDescAtc

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) HasProductDescAtc() bool`

HasProductDescAtc returns a boolean if a field has been set.

### GetProductPrice

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) GetProductPrice() int64`

GetProductPrice returns the ProductPrice field if non-nil, zero value otherwise.

### GetProductPriceOk

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) GetProductPriceOk() (*int64, bool)`

GetProductPriceOk returns a tuple with the ProductPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductPrice

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) SetProductPrice(v int64)`

SetProductPrice sets ProductPrice field to given value.

### HasProductPrice

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) HasProductPrice() bool`

HasProductPrice returns a boolean if a field has been set.

### GetSubtotalPrice

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) GetSubtotalPrice() int64`

GetSubtotalPrice returns the SubtotalPrice field if non-nil, zero value otherwise.

### GetSubtotalPriceOk

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) GetSubtotalPriceOk() (*int64, bool)`

GetSubtotalPriceOk returns a tuple with the SubtotalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotalPrice

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) SetSubtotalPrice(v int64)`

SetSubtotalPrice sets SubtotalPrice field to given value.

### HasSubtotalPrice

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) HasSubtotalPrice() bool`

HasSubtotalPrice returns a boolean if a field has been set.

### GetWeight

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) GetWeight() float64`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) GetWeightOk() (*float64, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) SetWeight(v float64)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetTotalWeight

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) GetTotalWeight() float64`

GetTotalWeight returns the TotalWeight field if non-nil, zero value otherwise.

### GetTotalWeightOk

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) GetTotalWeightOk() (*float64, bool)`

GetTotalWeightOk returns a tuple with the TotalWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWeight

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) SetTotalWeight(v float64)`

SetTotalWeight sets TotalWeight field to given value.

### HasTotalWeight

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) HasTotalWeight() bool`

HasTotalWeight returns a boolean if a field has been set.

### GetQuantity

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetQuantityDeliver

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) GetQuantityDeliver() int64`

GetQuantityDeliver returns the QuantityDeliver field if non-nil, zero value otherwise.

### GetQuantityDeliverOk

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) GetQuantityDeliverOk() (*int64, bool)`

GetQuantityDeliverOk returns a tuple with the QuantityDeliver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityDeliver

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) SetQuantityDeliver(v int64)`

SetQuantityDeliver sets QuantityDeliver field to given value.

### HasQuantityDeliver

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) HasQuantityDeliver() bool`

HasQuantityDeliver returns a boolean if a field has been set.

### GetQuantityReject

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) GetQuantityReject() int64`

GetQuantityReject returns the QuantityReject field if non-nil, zero value otherwise.

### GetQuantityRejectOk

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) GetQuantityRejectOk() (*int64, bool)`

GetQuantityRejectOk returns a tuple with the QuantityReject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityReject

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) SetQuantityReject(v int64)`

SetQuantityReject sets QuantityReject field to given value.

### HasQuantityReject

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) HasQuantityReject() bool`

HasQuantityReject returns a boolean if a field has been set.

### GetIsFreeReturns

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) GetIsFreeReturns() bool`

GetIsFreeReturns returns the IsFreeReturns field if non-nil, zero value otherwise.

### GetIsFreeReturnsOk

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) GetIsFreeReturnsOk() (*bool, bool)`

GetIsFreeReturnsOk returns a tuple with the IsFreeReturns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFreeReturns

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) SetIsFreeReturns(v bool)`

SetIsFreeReturns sets IsFreeReturns field to given value.

### HasIsFreeReturns

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) HasIsFreeReturns() bool`

HasIsFreeReturns returns a boolean if a field has been set.

### GetInsurancePrice

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) GetInsurancePrice() int64`

GetInsurancePrice returns the InsurancePrice field if non-nil, zero value otherwise.

### GetInsurancePriceOk

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) GetInsurancePriceOk() (*int64, bool)`

GetInsurancePriceOk returns a tuple with the InsurancePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsurancePrice

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) SetInsurancePrice(v int64)`

SetInsurancePrice sets InsurancePrice field to given value.

### HasInsurancePrice

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) HasInsurancePrice() bool`

HasInsurancePrice returns a boolean if a field has been set.

### GetNormalPrice

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) GetNormalPrice() int64`

GetNormalPrice returns the NormalPrice field if non-nil, zero value otherwise.

### GetNormalPriceOk

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) GetNormalPriceOk() (*int64, bool)`

GetNormalPriceOk returns a tuple with the NormalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalPrice

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) SetNormalPrice(v int64)`

SetNormalPrice sets NormalPrice field to given value.

### HasNormalPrice

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) HasNormalPrice() bool`

HasNormalPrice returns a boolean if a field has been set.

### GetCurrencyId

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) GetCurrencyId() int64`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) GetCurrencyIdOk() (*int64, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) SetCurrencyId(v int64)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### GetCurrencyRate

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) GetCurrencyRate() int64`

GetCurrencyRate returns the CurrencyRate field if non-nil, zero value otherwise.

### GetCurrencyRateOk

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) GetCurrencyRateOk() (*int64, bool)`

GetCurrencyRateOk returns a tuple with the CurrencyRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyRate

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) SetCurrencyRate(v int64)`

SetCurrencyRate sets CurrencyRate field to given value.

### HasCurrencyRate

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) HasCurrencyRate() bool`

HasCurrencyRate returns a boolean if a field has been set.

### GetMinOrder

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) GetMinOrder() int64`

GetMinOrder returns the MinOrder field if non-nil, zero value otherwise.

### GetMinOrderOk

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) GetMinOrderOk() (*int64, bool)`

GetMinOrderOk returns a tuple with the MinOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinOrder

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) SetMinOrder(v int64)`

SetMinOrder sets MinOrder field to given value.

### HasMinOrder

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) HasMinOrder() bool`

HasMinOrder returns a boolean if a field has been set.

### GetChildCatId

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) GetChildCatId() int64`

GetChildCatId returns the ChildCatId field if non-nil, zero value otherwise.

### GetChildCatIdOk

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) GetChildCatIdOk() (*int64, bool)`

GetChildCatIdOk returns a tuple with the ChildCatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildCatId

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) SetChildCatId(v int64)`

SetChildCatId sets ChildCatId field to given value.

### HasChildCatId

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) HasChildCatId() bool`

HasChildCatId returns a boolean if a field has been set.

### GetCampaignId

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) GetCampaignId() string`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) GetCampaignIdOk() (*string, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) SetCampaignId(v string)`

SetCampaignId sets CampaignId field to given value.

### HasCampaignId

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) HasCampaignId() bool`

HasCampaignId returns a boolean if a field has been set.

### GetProductPicture

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) GetProductPicture() string`

GetProductPicture returns the ProductPicture field if non-nil, zero value otherwise.

### GetProductPictureOk

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) GetProductPictureOk() (*string, bool)`

GetProductPictureOk returns a tuple with the ProductPicture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductPicture

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) SetProductPicture(v string)`

SetProductPicture sets ProductPicture field to given value.

### HasProductPicture

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) HasProductPicture() bool`

HasProductPicture returns a boolean if a field has been set.

### GetSnapshotUrl

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) GetSnapshotUrl() string`

GetSnapshotUrl returns the SnapshotUrl field if non-nil, zero value otherwise.

### GetSnapshotUrlOk

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) GetSnapshotUrlOk() (*string, bool)`

GetSnapshotUrlOk returns a tuple with the SnapshotUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotUrl

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) SetSnapshotUrl(v string)`

SetSnapshotUrl sets SnapshotUrl field to given value.

### HasSnapshotUrl

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) HasSnapshotUrl() bool`

HasSnapshotUrl returns a boolean if a field has been set.

### GetSku

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *GetSingleOrder200ResponseDataOrderInfoOrderDetailInner) HasSku() bool`

HasSku returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


