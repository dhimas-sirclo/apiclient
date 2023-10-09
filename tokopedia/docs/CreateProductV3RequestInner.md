# CreateProductV3RequestInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the product with length less than or equals 70 characters | [optional] 
**Condition** | Pointer to **string** | The condition of the product with the following available values NEW and USED | [optional] 
**Description** | Pointer to **string** | Description of the product. Maximum characters allowed is 2000 | [optional] 
**Sku** | Pointer to **string** | The stock keeping unit for the product. Maximum characters allowed is 50 | [optional] 
**Price** | Pointer to **float64** | The possible value between 100 to 100.000.000. If the product variant is added, the price parameter is automatically set to the lowest price among the variant products | [optional] 
**Status** | Pointer to **string** | Status for the product with the following available values UNLIMITED, LIMITED, and EMPTY | [optional] 
**Stock** | Pointer to **int64** | The stock of the product. 0 indicates always available. Other than that, the possible values are from 1 to 1000. Stock should be 1 if want to add variant product. | [optional] 
**MinOrder** | Pointer to **int64** | Minimum order required to purchase the product. Can only be a positive integ | [optional] [default to 1]
**CategoryId** | Pointer to **int64** | Unique identifier for the product’s category. To get available categories, please check Get All Categories Please input the deepest category child ID | [optional] 
**Dimension** | Pointer to [**CreateProductV3RequestInnerDimension**](CreateProductV3RequestInnerDimension.md) |  | [optional] 
**CustomProductLogistics** | Pointer to **[]int64** | Custom product logistics of the product. To get the id, please check Get Active Courier. Required field to input just ShippingProductID value responses from Get Active Courier | [optional] 
**Annotations** | Pointer to **[]int64** | Product Specification (anotations) By Category ID. The value is array of annotations id that can be retrieve by hit endpoint Get Product Annotation by Category ID. The location of id is at values.id | [optional] 
**PriceCurrency** | Pointer to **string** | Currency code for stated price (IDR or USD) | [optional] 
**Weight** | Pointer to **float64** | Weight of the product | [optional] 
**WeightUnit** | Pointer to **string** | The unit of the weight with the following available value GR (gram) | [optional] [default to "GR"]
**IsFreeReturn** | Pointer to **bool** | Determine if the product can be returned (true) by buyers or not (false) | [optional] 
**IsMustInsurance** | Pointer to **bool** | Determine if the product must be insured (true) or not (false) | [optional] 
**Etalase** | Pointer to [**EditProductV2RequestProductsInnerEtalase**](EditProductV2RequestProductsInnerEtalase.md) |  | [optional] 
**Pictures** | Pointer to [**[]EditProductV3RequestProductsInnerPicturesInner**](EditProductV3RequestProductsInnerPicturesInner.md) | Images information of the product. The object keys includes: file_path | [optional] 
**Wholesale** | Pointer to [**[]CreateProductV3RequestInnerWholesaleInner**](CreateProductV3RequestInnerWholesaleInner.md) | Wholesale price and quantity of the product. The object keys includes: min_qty and price | [optional] 
**Preorder** | Pointer to [**CreateProductV3RequestInnerPreorder**](CreateProductV3RequestInnerPreorder.md) |  | [optional] 
**Videos** | Pointer to [**[]CreateProductV3RequestInnerVideosInner**](CreateProductV3RequestInnerVideosInner.md) | Video link of the product. The object keys includes: url and source. url should only contain the YouTube video id e.g. dQw4w9WgXcQ. Where the type type should be youtube | [optional] 
**Variant** | Pointer to [**CreateProductV3RequestInnerVariant**](CreateProductV3RequestInnerVariant.md) |  | [optional] 

## Methods

### NewCreateProductV3RequestInner

`func NewCreateProductV3RequestInner() *CreateProductV3RequestInner`

NewCreateProductV3RequestInner instantiates a new CreateProductV3RequestInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProductV3RequestInnerWithDefaults

`func NewCreateProductV3RequestInnerWithDefaults() *CreateProductV3RequestInner`

NewCreateProductV3RequestInnerWithDefaults instantiates a new CreateProductV3RequestInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateProductV3RequestInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateProductV3RequestInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateProductV3RequestInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateProductV3RequestInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCondition

`func (o *CreateProductV3RequestInner) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *CreateProductV3RequestInner) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *CreateProductV3RequestInner) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *CreateProductV3RequestInner) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetDescription

`func (o *CreateProductV3RequestInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateProductV3RequestInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateProductV3RequestInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateProductV3RequestInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSku

`func (o *CreateProductV3RequestInner) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *CreateProductV3RequestInner) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *CreateProductV3RequestInner) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *CreateProductV3RequestInner) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetPrice

`func (o *CreateProductV3RequestInner) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CreateProductV3RequestInner) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CreateProductV3RequestInner) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *CreateProductV3RequestInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetStatus

`func (o *CreateProductV3RequestInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateProductV3RequestInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateProductV3RequestInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateProductV3RequestInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStock

`func (o *CreateProductV3RequestInner) GetStock() int64`

GetStock returns the Stock field if non-nil, zero value otherwise.

### GetStockOk

`func (o *CreateProductV3RequestInner) GetStockOk() (*int64, bool)`

GetStockOk returns a tuple with the Stock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStock

`func (o *CreateProductV3RequestInner) SetStock(v int64)`

SetStock sets Stock field to given value.

### HasStock

`func (o *CreateProductV3RequestInner) HasStock() bool`

HasStock returns a boolean if a field has been set.

### GetMinOrder

`func (o *CreateProductV3RequestInner) GetMinOrder() int64`

GetMinOrder returns the MinOrder field if non-nil, zero value otherwise.

### GetMinOrderOk

`func (o *CreateProductV3RequestInner) GetMinOrderOk() (*int64, bool)`

GetMinOrderOk returns a tuple with the MinOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinOrder

`func (o *CreateProductV3RequestInner) SetMinOrder(v int64)`

SetMinOrder sets MinOrder field to given value.

### HasMinOrder

`func (o *CreateProductV3RequestInner) HasMinOrder() bool`

HasMinOrder returns a boolean if a field has been set.

### GetCategoryId

`func (o *CreateProductV3RequestInner) GetCategoryId() int64`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *CreateProductV3RequestInner) GetCategoryIdOk() (*int64, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *CreateProductV3RequestInner) SetCategoryId(v int64)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *CreateProductV3RequestInner) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### GetDimension

`func (o *CreateProductV3RequestInner) GetDimension() CreateProductV3RequestInnerDimension`

GetDimension returns the Dimension field if non-nil, zero value otherwise.

### GetDimensionOk

`func (o *CreateProductV3RequestInner) GetDimensionOk() (*CreateProductV3RequestInnerDimension, bool)`

GetDimensionOk returns a tuple with the Dimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimension

`func (o *CreateProductV3RequestInner) SetDimension(v CreateProductV3RequestInnerDimension)`

SetDimension sets Dimension field to given value.

### HasDimension

`func (o *CreateProductV3RequestInner) HasDimension() bool`

HasDimension returns a boolean if a field has been set.

### GetCustomProductLogistics

`func (o *CreateProductV3RequestInner) GetCustomProductLogistics() []int64`

GetCustomProductLogistics returns the CustomProductLogistics field if non-nil, zero value otherwise.

### GetCustomProductLogisticsOk

`func (o *CreateProductV3RequestInner) GetCustomProductLogisticsOk() (*[]int64, bool)`

GetCustomProductLogisticsOk returns a tuple with the CustomProductLogistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProductLogistics

`func (o *CreateProductV3RequestInner) SetCustomProductLogistics(v []int64)`

SetCustomProductLogistics sets CustomProductLogistics field to given value.

### HasCustomProductLogistics

`func (o *CreateProductV3RequestInner) HasCustomProductLogistics() bool`

HasCustomProductLogistics returns a boolean if a field has been set.

### GetAnnotations

`func (o *CreateProductV3RequestInner) GetAnnotations() []int64`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *CreateProductV3RequestInner) GetAnnotationsOk() (*[]int64, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *CreateProductV3RequestInner) SetAnnotations(v []int64)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *CreateProductV3RequestInner) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetPriceCurrency

`func (o *CreateProductV3RequestInner) GetPriceCurrency() string`

GetPriceCurrency returns the PriceCurrency field if non-nil, zero value otherwise.

### GetPriceCurrencyOk

`func (o *CreateProductV3RequestInner) GetPriceCurrencyOk() (*string, bool)`

GetPriceCurrencyOk returns a tuple with the PriceCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceCurrency

`func (o *CreateProductV3RequestInner) SetPriceCurrency(v string)`

SetPriceCurrency sets PriceCurrency field to given value.

### HasPriceCurrency

`func (o *CreateProductV3RequestInner) HasPriceCurrency() bool`

HasPriceCurrency returns a boolean if a field has been set.

### GetWeight

`func (o *CreateProductV3RequestInner) GetWeight() float64`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *CreateProductV3RequestInner) GetWeightOk() (*float64, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *CreateProductV3RequestInner) SetWeight(v float64)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *CreateProductV3RequestInner) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetWeightUnit

`func (o *CreateProductV3RequestInner) GetWeightUnit() string`

GetWeightUnit returns the WeightUnit field if non-nil, zero value otherwise.

### GetWeightUnitOk

`func (o *CreateProductV3RequestInner) GetWeightUnitOk() (*string, bool)`

GetWeightUnitOk returns a tuple with the WeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightUnit

`func (o *CreateProductV3RequestInner) SetWeightUnit(v string)`

SetWeightUnit sets WeightUnit field to given value.

### HasWeightUnit

`func (o *CreateProductV3RequestInner) HasWeightUnit() bool`

HasWeightUnit returns a boolean if a field has been set.

### GetIsFreeReturn

`func (o *CreateProductV3RequestInner) GetIsFreeReturn() bool`

GetIsFreeReturn returns the IsFreeReturn field if non-nil, zero value otherwise.

### GetIsFreeReturnOk

`func (o *CreateProductV3RequestInner) GetIsFreeReturnOk() (*bool, bool)`

GetIsFreeReturnOk returns a tuple with the IsFreeReturn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFreeReturn

`func (o *CreateProductV3RequestInner) SetIsFreeReturn(v bool)`

SetIsFreeReturn sets IsFreeReturn field to given value.

### HasIsFreeReturn

`func (o *CreateProductV3RequestInner) HasIsFreeReturn() bool`

HasIsFreeReturn returns a boolean if a field has been set.

### GetIsMustInsurance

`func (o *CreateProductV3RequestInner) GetIsMustInsurance() bool`

GetIsMustInsurance returns the IsMustInsurance field if non-nil, zero value otherwise.

### GetIsMustInsuranceOk

`func (o *CreateProductV3RequestInner) GetIsMustInsuranceOk() (*bool, bool)`

GetIsMustInsuranceOk returns a tuple with the IsMustInsurance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMustInsurance

`func (o *CreateProductV3RequestInner) SetIsMustInsurance(v bool)`

SetIsMustInsurance sets IsMustInsurance field to given value.

### HasIsMustInsurance

`func (o *CreateProductV3RequestInner) HasIsMustInsurance() bool`

HasIsMustInsurance returns a boolean if a field has been set.

### GetEtalase

`func (o *CreateProductV3RequestInner) GetEtalase() EditProductV2RequestProductsInnerEtalase`

GetEtalase returns the Etalase field if non-nil, zero value otherwise.

### GetEtalaseOk

`func (o *CreateProductV3RequestInner) GetEtalaseOk() (*EditProductV2RequestProductsInnerEtalase, bool)`

GetEtalaseOk returns a tuple with the Etalase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtalase

`func (o *CreateProductV3RequestInner) SetEtalase(v EditProductV2RequestProductsInnerEtalase)`

SetEtalase sets Etalase field to given value.

### HasEtalase

`func (o *CreateProductV3RequestInner) HasEtalase() bool`

HasEtalase returns a boolean if a field has been set.

### GetPictures

`func (o *CreateProductV3RequestInner) GetPictures() []EditProductV3RequestProductsInnerPicturesInner`

GetPictures returns the Pictures field if non-nil, zero value otherwise.

### GetPicturesOk

`func (o *CreateProductV3RequestInner) GetPicturesOk() (*[]EditProductV3RequestProductsInnerPicturesInner, bool)`

GetPicturesOk returns a tuple with the Pictures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPictures

`func (o *CreateProductV3RequestInner) SetPictures(v []EditProductV3RequestProductsInnerPicturesInner)`

SetPictures sets Pictures field to given value.

### HasPictures

`func (o *CreateProductV3RequestInner) HasPictures() bool`

HasPictures returns a boolean if a field has been set.

### GetWholesale

`func (o *CreateProductV3RequestInner) GetWholesale() []CreateProductV3RequestInnerWholesaleInner`

GetWholesale returns the Wholesale field if non-nil, zero value otherwise.

### GetWholesaleOk

`func (o *CreateProductV3RequestInner) GetWholesaleOk() (*[]CreateProductV3RequestInnerWholesaleInner, bool)`

GetWholesaleOk returns a tuple with the Wholesale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWholesale

`func (o *CreateProductV3RequestInner) SetWholesale(v []CreateProductV3RequestInnerWholesaleInner)`

SetWholesale sets Wholesale field to given value.

### HasWholesale

`func (o *CreateProductV3RequestInner) HasWholesale() bool`

HasWholesale returns a boolean if a field has been set.

### GetPreorder

`func (o *CreateProductV3RequestInner) GetPreorder() CreateProductV3RequestInnerPreorder`

GetPreorder returns the Preorder field if non-nil, zero value otherwise.

### GetPreorderOk

`func (o *CreateProductV3RequestInner) GetPreorderOk() (*CreateProductV3RequestInnerPreorder, bool)`

GetPreorderOk returns a tuple with the Preorder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreorder

`func (o *CreateProductV3RequestInner) SetPreorder(v CreateProductV3RequestInnerPreorder)`

SetPreorder sets Preorder field to given value.

### HasPreorder

`func (o *CreateProductV3RequestInner) HasPreorder() bool`

HasPreorder returns a boolean if a field has been set.

### GetVideos

`func (o *CreateProductV3RequestInner) GetVideos() []CreateProductV3RequestInnerVideosInner`

GetVideos returns the Videos field if non-nil, zero value otherwise.

### GetVideosOk

`func (o *CreateProductV3RequestInner) GetVideosOk() (*[]CreateProductV3RequestInnerVideosInner, bool)`

GetVideosOk returns a tuple with the Videos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideos

`func (o *CreateProductV3RequestInner) SetVideos(v []CreateProductV3RequestInnerVideosInner)`

SetVideos sets Videos field to given value.

### HasVideos

`func (o *CreateProductV3RequestInner) HasVideos() bool`

HasVideos returns a boolean if a field has been set.

### GetVariant

`func (o *CreateProductV3RequestInner) GetVariant() CreateProductV3RequestInnerVariant`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *CreateProductV3RequestInner) GetVariantOk() (*CreateProductV3RequestInnerVariant, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *CreateProductV3RequestInner) SetVariant(v CreateProductV3RequestInnerVariant)`

SetVariant sets Variant field to given value.

### HasVariant

`func (o *CreateProductV3RequestInner) HasVariant() bool`

HasVariant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


