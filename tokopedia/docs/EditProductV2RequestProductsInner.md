# EditProductV2RequestProductsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Product Identifier. Please choose either id or sku. | [optional] 
**Sku** | Pointer to **string** | The stock keeping unit for the product. Maximum characters allowed is 50. Please choose either id or sku. | [optional] 
**Name** | Pointer to **string** | Name of the product with length less than or equals 70 characters | [optional] 
**CategoryId** | Pointer to **int64** | Unique identifier for the product’s category. To get available categories, please check Get All Categories Please input the deepest category child ID | [optional] 
**Condition** | Pointer to **string** | The condition of the product with the following available values NEW and USED | [optional] 
**Description** | Pointer to **string** | Description of the product. Maximum characters allowed is 2000 | [optional] 
**Price** | Pointer to **float64** | The possible value between 100 to 100.000.000. If the product variant is added, the price parameter is automatically set to the lowest price among the variant products | [optional] 
**Status** | Pointer to **string** | Status for the product with the following available values UNLIMITED, LIMITED, and EMPTY | [optional] 
**Stock** | Pointer to **int64** | The stock of the product. 0 indicates always available. Other than that, the possible values are from 1 to 1000. Stock should be 1 if want to add variant product | [optional] 
**MinOrder** | Pointer to **int64** | Minimum order required to purchase the product. Can only be a positive integer | [optional] 
**PriceCurrency** | Pointer to **string** | Currency code for stated price (IDR or USD) | [optional] 
**Weight** | Pointer to **float64** | Weight of the product | [optional] 
**WeightUnit** | Pointer to **string** | The unit of the weight with the following available value GR (gram) | [optional] [default to "GR"]
**IsFreeReturn** | Pointer to **bool** | Determine if the product can be returned (true) by buyers or not (false) | [optional] 
**IsMustInsurance** | Pointer to **bool** | Determine if the product must be insured (true) or not (false) | [optional] 
**Etalase** | Pointer to [**EditProductV2RequestProductsInnerEtalase**](EditProductV2RequestProductsInnerEtalase.md) |  | [optional] 
**Pictures** | Pointer to [**[]EditProductV2RequestProductsInnerPicturesInner**](EditProductV2RequestProductsInnerPicturesInner.md) | Images information of the product. The object keys includes: file_path. Remove pictures from related product by passing empty array ([]) in this request | [optional] 
**Wholesale** | Pointer to [**[]EditProductV2RequestProductsInnerWholesaleInner**](EditProductV2RequestProductsInnerWholesaleInner.md) | Wholesale price and quantity of the product. The object keys includes: min_qty and price. Remove wholesale from related product by passing empty array ([]) in this request | [optional] 
**Preorder** | Pointer to [**EditProductV2RequestProductsInnerPreorder**](EditProductV2RequestProductsInnerPreorder.md) |  | [optional] 
**Videos** | Pointer to [**[]EditProductV2RequestProductsInnerVideosInner**](EditProductV2RequestProductsInnerVideosInner.md) | Video link of the product. The object keys includes: url and source. url should only contain the YouTube video id e.g. dQw4w9WgXcQ. Where the type type should be youtube. Remove videos from related product by passing empty array ([]) in this request | [optional] 
**Variant** | Pointer to [**EditProductV2RequestProductsInnerVariant**](EditProductV2RequestProductsInnerVariant.md) |  | [optional] 

## Methods

### NewEditProductV2RequestProductsInner

`func NewEditProductV2RequestProductsInner() *EditProductV2RequestProductsInner`

NewEditProductV2RequestProductsInner instantiates a new EditProductV2RequestProductsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditProductV2RequestProductsInnerWithDefaults

`func NewEditProductV2RequestProductsInnerWithDefaults() *EditProductV2RequestProductsInner`

NewEditProductV2RequestProductsInnerWithDefaults instantiates a new EditProductV2RequestProductsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EditProductV2RequestProductsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EditProductV2RequestProductsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EditProductV2RequestProductsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *EditProductV2RequestProductsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSku

`func (o *EditProductV2RequestProductsInner) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *EditProductV2RequestProductsInner) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *EditProductV2RequestProductsInner) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *EditProductV2RequestProductsInner) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetName

`func (o *EditProductV2RequestProductsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EditProductV2RequestProductsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EditProductV2RequestProductsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EditProductV2RequestProductsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCategoryId

`func (o *EditProductV2RequestProductsInner) GetCategoryId() int64`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *EditProductV2RequestProductsInner) GetCategoryIdOk() (*int64, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *EditProductV2RequestProductsInner) SetCategoryId(v int64)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *EditProductV2RequestProductsInner) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### GetCondition

`func (o *EditProductV2RequestProductsInner) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *EditProductV2RequestProductsInner) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *EditProductV2RequestProductsInner) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *EditProductV2RequestProductsInner) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetDescription

`func (o *EditProductV2RequestProductsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EditProductV2RequestProductsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EditProductV2RequestProductsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EditProductV2RequestProductsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPrice

`func (o *EditProductV2RequestProductsInner) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *EditProductV2RequestProductsInner) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *EditProductV2RequestProductsInner) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *EditProductV2RequestProductsInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetStatus

`func (o *EditProductV2RequestProductsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EditProductV2RequestProductsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EditProductV2RequestProductsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EditProductV2RequestProductsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStock

`func (o *EditProductV2RequestProductsInner) GetStock() int64`

GetStock returns the Stock field if non-nil, zero value otherwise.

### GetStockOk

`func (o *EditProductV2RequestProductsInner) GetStockOk() (*int64, bool)`

GetStockOk returns a tuple with the Stock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStock

`func (o *EditProductV2RequestProductsInner) SetStock(v int64)`

SetStock sets Stock field to given value.

### HasStock

`func (o *EditProductV2RequestProductsInner) HasStock() bool`

HasStock returns a boolean if a field has been set.

### GetMinOrder

`func (o *EditProductV2RequestProductsInner) GetMinOrder() int64`

GetMinOrder returns the MinOrder field if non-nil, zero value otherwise.

### GetMinOrderOk

`func (o *EditProductV2RequestProductsInner) GetMinOrderOk() (*int64, bool)`

GetMinOrderOk returns a tuple with the MinOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinOrder

`func (o *EditProductV2RequestProductsInner) SetMinOrder(v int64)`

SetMinOrder sets MinOrder field to given value.

### HasMinOrder

`func (o *EditProductV2RequestProductsInner) HasMinOrder() bool`

HasMinOrder returns a boolean if a field has been set.

### GetPriceCurrency

`func (o *EditProductV2RequestProductsInner) GetPriceCurrency() string`

GetPriceCurrency returns the PriceCurrency field if non-nil, zero value otherwise.

### GetPriceCurrencyOk

`func (o *EditProductV2RequestProductsInner) GetPriceCurrencyOk() (*string, bool)`

GetPriceCurrencyOk returns a tuple with the PriceCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceCurrency

`func (o *EditProductV2RequestProductsInner) SetPriceCurrency(v string)`

SetPriceCurrency sets PriceCurrency field to given value.

### HasPriceCurrency

`func (o *EditProductV2RequestProductsInner) HasPriceCurrency() bool`

HasPriceCurrency returns a boolean if a field has been set.

### GetWeight

`func (o *EditProductV2RequestProductsInner) GetWeight() float64`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *EditProductV2RequestProductsInner) GetWeightOk() (*float64, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *EditProductV2RequestProductsInner) SetWeight(v float64)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *EditProductV2RequestProductsInner) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetWeightUnit

`func (o *EditProductV2RequestProductsInner) GetWeightUnit() string`

GetWeightUnit returns the WeightUnit field if non-nil, zero value otherwise.

### GetWeightUnitOk

`func (o *EditProductV2RequestProductsInner) GetWeightUnitOk() (*string, bool)`

GetWeightUnitOk returns a tuple with the WeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightUnit

`func (o *EditProductV2RequestProductsInner) SetWeightUnit(v string)`

SetWeightUnit sets WeightUnit field to given value.

### HasWeightUnit

`func (o *EditProductV2RequestProductsInner) HasWeightUnit() bool`

HasWeightUnit returns a boolean if a field has been set.

### GetIsFreeReturn

`func (o *EditProductV2RequestProductsInner) GetIsFreeReturn() bool`

GetIsFreeReturn returns the IsFreeReturn field if non-nil, zero value otherwise.

### GetIsFreeReturnOk

`func (o *EditProductV2RequestProductsInner) GetIsFreeReturnOk() (*bool, bool)`

GetIsFreeReturnOk returns a tuple with the IsFreeReturn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFreeReturn

`func (o *EditProductV2RequestProductsInner) SetIsFreeReturn(v bool)`

SetIsFreeReturn sets IsFreeReturn field to given value.

### HasIsFreeReturn

`func (o *EditProductV2RequestProductsInner) HasIsFreeReturn() bool`

HasIsFreeReturn returns a boolean if a field has been set.

### GetIsMustInsurance

`func (o *EditProductV2RequestProductsInner) GetIsMustInsurance() bool`

GetIsMustInsurance returns the IsMustInsurance field if non-nil, zero value otherwise.

### GetIsMustInsuranceOk

`func (o *EditProductV2RequestProductsInner) GetIsMustInsuranceOk() (*bool, bool)`

GetIsMustInsuranceOk returns a tuple with the IsMustInsurance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMustInsurance

`func (o *EditProductV2RequestProductsInner) SetIsMustInsurance(v bool)`

SetIsMustInsurance sets IsMustInsurance field to given value.

### HasIsMustInsurance

`func (o *EditProductV2RequestProductsInner) HasIsMustInsurance() bool`

HasIsMustInsurance returns a boolean if a field has been set.

### GetEtalase

`func (o *EditProductV2RequestProductsInner) GetEtalase() EditProductV2RequestProductsInnerEtalase`

GetEtalase returns the Etalase field if non-nil, zero value otherwise.

### GetEtalaseOk

`func (o *EditProductV2RequestProductsInner) GetEtalaseOk() (*EditProductV2RequestProductsInnerEtalase, bool)`

GetEtalaseOk returns a tuple with the Etalase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtalase

`func (o *EditProductV2RequestProductsInner) SetEtalase(v EditProductV2RequestProductsInnerEtalase)`

SetEtalase sets Etalase field to given value.

### HasEtalase

`func (o *EditProductV2RequestProductsInner) HasEtalase() bool`

HasEtalase returns a boolean if a field has been set.

### GetPictures

`func (o *EditProductV2RequestProductsInner) GetPictures() []EditProductV2RequestProductsInnerPicturesInner`

GetPictures returns the Pictures field if non-nil, zero value otherwise.

### GetPicturesOk

`func (o *EditProductV2RequestProductsInner) GetPicturesOk() (*[]EditProductV2RequestProductsInnerPicturesInner, bool)`

GetPicturesOk returns a tuple with the Pictures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPictures

`func (o *EditProductV2RequestProductsInner) SetPictures(v []EditProductV2RequestProductsInnerPicturesInner)`

SetPictures sets Pictures field to given value.

### HasPictures

`func (o *EditProductV2RequestProductsInner) HasPictures() bool`

HasPictures returns a boolean if a field has been set.

### GetWholesale

`func (o *EditProductV2RequestProductsInner) GetWholesale() []EditProductV2RequestProductsInnerWholesaleInner`

GetWholesale returns the Wholesale field if non-nil, zero value otherwise.

### GetWholesaleOk

`func (o *EditProductV2RequestProductsInner) GetWholesaleOk() (*[]EditProductV2RequestProductsInnerWholesaleInner, bool)`

GetWholesaleOk returns a tuple with the Wholesale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWholesale

`func (o *EditProductV2RequestProductsInner) SetWholesale(v []EditProductV2RequestProductsInnerWholesaleInner)`

SetWholesale sets Wholesale field to given value.

### HasWholesale

`func (o *EditProductV2RequestProductsInner) HasWholesale() bool`

HasWholesale returns a boolean if a field has been set.

### GetPreorder

`func (o *EditProductV2RequestProductsInner) GetPreorder() EditProductV2RequestProductsInnerPreorder`

GetPreorder returns the Preorder field if non-nil, zero value otherwise.

### GetPreorderOk

`func (o *EditProductV2RequestProductsInner) GetPreorderOk() (*EditProductV2RequestProductsInnerPreorder, bool)`

GetPreorderOk returns a tuple with the Preorder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreorder

`func (o *EditProductV2RequestProductsInner) SetPreorder(v EditProductV2RequestProductsInnerPreorder)`

SetPreorder sets Preorder field to given value.

### HasPreorder

`func (o *EditProductV2RequestProductsInner) HasPreorder() bool`

HasPreorder returns a boolean if a field has been set.

### GetVideos

`func (o *EditProductV2RequestProductsInner) GetVideos() []EditProductV2RequestProductsInnerVideosInner`

GetVideos returns the Videos field if non-nil, zero value otherwise.

### GetVideosOk

`func (o *EditProductV2RequestProductsInner) GetVideosOk() (*[]EditProductV2RequestProductsInnerVideosInner, bool)`

GetVideosOk returns a tuple with the Videos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideos

`func (o *EditProductV2RequestProductsInner) SetVideos(v []EditProductV2RequestProductsInnerVideosInner)`

SetVideos sets Videos field to given value.

### HasVideos

`func (o *EditProductV2RequestProductsInner) HasVideos() bool`

HasVideos returns a boolean if a field has been set.

### GetVariant

`func (o *EditProductV2RequestProductsInner) GetVariant() EditProductV2RequestProductsInnerVariant`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *EditProductV2RequestProductsInner) GetVariantOk() (*EditProductV2RequestProductsInnerVariant, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *EditProductV2RequestProductsInner) SetVariant(v EditProductV2RequestProductsInnerVariant)`

SetVariant sets Variant field to given value.

### HasVariant

`func (o *EditProductV2RequestProductsInner) HasVariant() bool`

HasVariant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


