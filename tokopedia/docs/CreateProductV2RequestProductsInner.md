# CreateProductV2RequestProductsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the product with length less than or equals 70 characters | 
**Condition** | **string** | The condition of the product with the following available values NEW and USED | [default to "NEW"]
**Description** | Pointer to **string** | Description of the product. Maximum characters allowed is 2000 | [optional] 
**Sku** | Pointer to **string** | The stock keeping unit for the product. Maximum characters allowed is 50 | [optional] 
**Price** | **float64** | The possible value between 100 to 100.000.000. If the product variant is added, the price parameter is automatically set to the lowest price among the variant products. | 
**Status** | **string** | Status for the product with the following available values UNLIMITED, LIMITED, and EMPTY. | [default to "LIMITED"]
**Stock** | Pointer to **int64** | The stock of the product. 0 indicates always available. Other than that, the possible values are from 1 to 1000. Stock should be 1 if want to add variant product | [optional] 
**MinOrder** | **int64** | Minimum order required to purchase the product. Can only be a positive integer | 
**CategoryId** | **int64** | Unique identifier for the product’s category. To get available categories, please check Get All Categories. Please input the deepest category child ID | 
**PriceCurrency** | **string** | Currency code for stated price (IDR or USD) | [default to "IDR"]
**Weight** | **float64** | Weight of the product | 
**WeightUnit** | **string** | The unit of the weight with the following available value GR (gram) | [default to "GR"]
**IsFreeReturn** | Pointer to **bool** | Determine if the product can be returned (true) by buyers or not (false) | [optional] 
**IsMustInsurance** | **bool** | Determine if the product must be insured (true) or not (false) | 
**Etalase** | [**CreateProductV2RequestProductsInnerEtalase**](CreateProductV2RequestProductsInnerEtalase.md) |  | 
**Pictures** | [**[]EditProductV3RequestProductsInnerPicturesInner**](EditProductV3RequestProductsInnerPicturesInner.md) | Images information of the product. The object keys includes: file_path | 
**Wholesale** | Pointer to [**[]CreateProductV3RequestInnerWholesaleInner**](CreateProductV3RequestInnerWholesaleInner.md) | Wholesale price and quantity of the product. The object keys includes: min_qty and price | [optional] 
**Preorder** | Pointer to [**CreateProductV2RequestProductsInnerPreorder**](CreateProductV2RequestProductsInnerPreorder.md) |  | [optional] 
**Videos** | Pointer to [**[]CreateProductV2RequestProductsInnerVideosInner**](CreateProductV2RequestProductsInnerVideosInner.md) | Video link of the product. The object keys includes: url and source. url should only contain the YouTube video id e.g. dQw4w9WgXcQ. Where the type should be youtube | [optional] 
**Variant** | Pointer to [**[]CreateProductV2RequestProductsInnerVariantInner**](CreateProductV2RequestProductsInnerVariantInner.md) | Variant of the product. The object keys includes: variant and product_variant | [optional] 

## Methods

### NewCreateProductV2RequestProductsInner

`func NewCreateProductV2RequestProductsInner(name string, condition string, price float64, status string, minOrder int64, categoryId int64, priceCurrency string, weight float64, weightUnit string, isMustInsurance bool, etalase CreateProductV2RequestProductsInnerEtalase, pictures []EditProductV3RequestProductsInnerPicturesInner, ) *CreateProductV2RequestProductsInner`

NewCreateProductV2RequestProductsInner instantiates a new CreateProductV2RequestProductsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProductV2RequestProductsInnerWithDefaults

`func NewCreateProductV2RequestProductsInnerWithDefaults() *CreateProductV2RequestProductsInner`

NewCreateProductV2RequestProductsInnerWithDefaults instantiates a new CreateProductV2RequestProductsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateProductV2RequestProductsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateProductV2RequestProductsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateProductV2RequestProductsInner) SetName(v string)`

SetName sets Name field to given value.


### GetCondition

`func (o *CreateProductV2RequestProductsInner) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *CreateProductV2RequestProductsInner) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *CreateProductV2RequestProductsInner) SetCondition(v string)`

SetCondition sets Condition field to given value.


### GetDescription

`func (o *CreateProductV2RequestProductsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateProductV2RequestProductsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateProductV2RequestProductsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateProductV2RequestProductsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSku

`func (o *CreateProductV2RequestProductsInner) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *CreateProductV2RequestProductsInner) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *CreateProductV2RequestProductsInner) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *CreateProductV2RequestProductsInner) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetPrice

`func (o *CreateProductV2RequestProductsInner) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CreateProductV2RequestProductsInner) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CreateProductV2RequestProductsInner) SetPrice(v float64)`

SetPrice sets Price field to given value.


### GetStatus

`func (o *CreateProductV2RequestProductsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateProductV2RequestProductsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateProductV2RequestProductsInner) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStock

`func (o *CreateProductV2RequestProductsInner) GetStock() int64`

GetStock returns the Stock field if non-nil, zero value otherwise.

### GetStockOk

`func (o *CreateProductV2RequestProductsInner) GetStockOk() (*int64, bool)`

GetStockOk returns a tuple with the Stock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStock

`func (o *CreateProductV2RequestProductsInner) SetStock(v int64)`

SetStock sets Stock field to given value.

### HasStock

`func (o *CreateProductV2RequestProductsInner) HasStock() bool`

HasStock returns a boolean if a field has been set.

### GetMinOrder

`func (o *CreateProductV2RequestProductsInner) GetMinOrder() int64`

GetMinOrder returns the MinOrder field if non-nil, zero value otherwise.

### GetMinOrderOk

`func (o *CreateProductV2RequestProductsInner) GetMinOrderOk() (*int64, bool)`

GetMinOrderOk returns a tuple with the MinOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinOrder

`func (o *CreateProductV2RequestProductsInner) SetMinOrder(v int64)`

SetMinOrder sets MinOrder field to given value.


### GetCategoryId

`func (o *CreateProductV2RequestProductsInner) GetCategoryId() int64`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *CreateProductV2RequestProductsInner) GetCategoryIdOk() (*int64, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *CreateProductV2RequestProductsInner) SetCategoryId(v int64)`

SetCategoryId sets CategoryId field to given value.


### GetPriceCurrency

`func (o *CreateProductV2RequestProductsInner) GetPriceCurrency() string`

GetPriceCurrency returns the PriceCurrency field if non-nil, zero value otherwise.

### GetPriceCurrencyOk

`func (o *CreateProductV2RequestProductsInner) GetPriceCurrencyOk() (*string, bool)`

GetPriceCurrencyOk returns a tuple with the PriceCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceCurrency

`func (o *CreateProductV2RequestProductsInner) SetPriceCurrency(v string)`

SetPriceCurrency sets PriceCurrency field to given value.


### GetWeight

`func (o *CreateProductV2RequestProductsInner) GetWeight() float64`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *CreateProductV2RequestProductsInner) GetWeightOk() (*float64, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *CreateProductV2RequestProductsInner) SetWeight(v float64)`

SetWeight sets Weight field to given value.


### GetWeightUnit

`func (o *CreateProductV2RequestProductsInner) GetWeightUnit() string`

GetWeightUnit returns the WeightUnit field if non-nil, zero value otherwise.

### GetWeightUnitOk

`func (o *CreateProductV2RequestProductsInner) GetWeightUnitOk() (*string, bool)`

GetWeightUnitOk returns a tuple with the WeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightUnit

`func (o *CreateProductV2RequestProductsInner) SetWeightUnit(v string)`

SetWeightUnit sets WeightUnit field to given value.


### GetIsFreeReturn

`func (o *CreateProductV2RequestProductsInner) GetIsFreeReturn() bool`

GetIsFreeReturn returns the IsFreeReturn field if non-nil, zero value otherwise.

### GetIsFreeReturnOk

`func (o *CreateProductV2RequestProductsInner) GetIsFreeReturnOk() (*bool, bool)`

GetIsFreeReturnOk returns a tuple with the IsFreeReturn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFreeReturn

`func (o *CreateProductV2RequestProductsInner) SetIsFreeReturn(v bool)`

SetIsFreeReturn sets IsFreeReturn field to given value.

### HasIsFreeReturn

`func (o *CreateProductV2RequestProductsInner) HasIsFreeReturn() bool`

HasIsFreeReturn returns a boolean if a field has been set.

### GetIsMustInsurance

`func (o *CreateProductV2RequestProductsInner) GetIsMustInsurance() bool`

GetIsMustInsurance returns the IsMustInsurance field if non-nil, zero value otherwise.

### GetIsMustInsuranceOk

`func (o *CreateProductV2RequestProductsInner) GetIsMustInsuranceOk() (*bool, bool)`

GetIsMustInsuranceOk returns a tuple with the IsMustInsurance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMustInsurance

`func (o *CreateProductV2RequestProductsInner) SetIsMustInsurance(v bool)`

SetIsMustInsurance sets IsMustInsurance field to given value.


### GetEtalase

`func (o *CreateProductV2RequestProductsInner) GetEtalase() CreateProductV2RequestProductsInnerEtalase`

GetEtalase returns the Etalase field if non-nil, zero value otherwise.

### GetEtalaseOk

`func (o *CreateProductV2RequestProductsInner) GetEtalaseOk() (*CreateProductV2RequestProductsInnerEtalase, bool)`

GetEtalaseOk returns a tuple with the Etalase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtalase

`func (o *CreateProductV2RequestProductsInner) SetEtalase(v CreateProductV2RequestProductsInnerEtalase)`

SetEtalase sets Etalase field to given value.


### GetPictures

`func (o *CreateProductV2RequestProductsInner) GetPictures() []EditProductV3RequestProductsInnerPicturesInner`

GetPictures returns the Pictures field if non-nil, zero value otherwise.

### GetPicturesOk

`func (o *CreateProductV2RequestProductsInner) GetPicturesOk() (*[]EditProductV3RequestProductsInnerPicturesInner, bool)`

GetPicturesOk returns a tuple with the Pictures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPictures

`func (o *CreateProductV2RequestProductsInner) SetPictures(v []EditProductV3RequestProductsInnerPicturesInner)`

SetPictures sets Pictures field to given value.


### GetWholesale

`func (o *CreateProductV2RequestProductsInner) GetWholesale() []CreateProductV3RequestInnerWholesaleInner`

GetWholesale returns the Wholesale field if non-nil, zero value otherwise.

### GetWholesaleOk

`func (o *CreateProductV2RequestProductsInner) GetWholesaleOk() (*[]CreateProductV3RequestInnerWholesaleInner, bool)`

GetWholesaleOk returns a tuple with the Wholesale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWholesale

`func (o *CreateProductV2RequestProductsInner) SetWholesale(v []CreateProductV3RequestInnerWholesaleInner)`

SetWholesale sets Wholesale field to given value.

### HasWholesale

`func (o *CreateProductV2RequestProductsInner) HasWholesale() bool`

HasWholesale returns a boolean if a field has been set.

### GetPreorder

`func (o *CreateProductV2RequestProductsInner) GetPreorder() CreateProductV2RequestProductsInnerPreorder`

GetPreorder returns the Preorder field if non-nil, zero value otherwise.

### GetPreorderOk

`func (o *CreateProductV2RequestProductsInner) GetPreorderOk() (*CreateProductV2RequestProductsInnerPreorder, bool)`

GetPreorderOk returns a tuple with the Preorder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreorder

`func (o *CreateProductV2RequestProductsInner) SetPreorder(v CreateProductV2RequestProductsInnerPreorder)`

SetPreorder sets Preorder field to given value.

### HasPreorder

`func (o *CreateProductV2RequestProductsInner) HasPreorder() bool`

HasPreorder returns a boolean if a field has been set.

### GetVideos

`func (o *CreateProductV2RequestProductsInner) GetVideos() []CreateProductV2RequestProductsInnerVideosInner`

GetVideos returns the Videos field if non-nil, zero value otherwise.

### GetVideosOk

`func (o *CreateProductV2RequestProductsInner) GetVideosOk() (*[]CreateProductV2RequestProductsInnerVideosInner, bool)`

GetVideosOk returns a tuple with the Videos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideos

`func (o *CreateProductV2RequestProductsInner) SetVideos(v []CreateProductV2RequestProductsInnerVideosInner)`

SetVideos sets Videos field to given value.

### HasVideos

`func (o *CreateProductV2RequestProductsInner) HasVideos() bool`

HasVideos returns a boolean if a field has been set.

### GetVariant

`func (o *CreateProductV2RequestProductsInner) GetVariant() []CreateProductV2RequestProductsInnerVariantInner`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *CreateProductV2RequestProductsInner) GetVariantOk() (*[]CreateProductV2RequestProductsInnerVariantInner, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *CreateProductV2RequestProductsInner) SetVariant(v []CreateProductV2RequestProductsInnerVariantInner)`

SetVariant sets Variant field to given value.

### HasVariant

`func (o *CreateProductV2RequestProductsInner) HasVariant() bool`

HasVariant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


