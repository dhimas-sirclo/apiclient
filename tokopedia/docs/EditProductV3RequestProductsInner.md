# EditProductV3RequestProductsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Product Identifier | 
**Name** | Pointer to **string** | Name of the product with length less than or equals 70 characters | [optional] 
**Condition** | Pointer to **string** | The condition of the product with the following available values NEW and USED | [optional] 
**Description** | Pointer to **string** | Description of the product. Maximum characters allowed is 2000 | [optional] 
**Sku** | Pointer to **string** | The stock keeping unit for the product. Maximum characters allowed is 50 | [optional] 
**Price** | Pointer to **int64** | The possible value between 100 to 100.000.000. If the product variant is added, the price parameter is automatically set to the lowest price among the variant products | [optional] 
**Status** | **string** | Status for the product with the following available values UNLIMITED, LIMITED, and EMPTY | 
**Stock** | Pointer to **int64** | The stock of the product. 0 indicates always available. Other than that, the possible values are from 1 to 1000. Stock should be 1 if want to add variant product | [optional] 
**MinOrder** | Pointer to **int64** | Minimum order required to purchase the product. Can only be a positive integer | [optional] 
**CategoryId** | Pointer to **int64** | Unique identifier for the product’s category. To get available categories, please check Get All Categories Please input the deepest category child ID | [optional] 
**Dimension** | Pointer to [**EditProductV3RequestProductsInnerDimension**](EditProductV3RequestProductsInnerDimension.md) |  | [optional] 
**CustomProductLogistics** | Pointer to **[]int64** | Custom product logistics of the product. To get the id, please check Get Active Courier. Required field to input just ShippingProductID value responses from Get Active Courier. Remove custom_product_logistics from related product by passing empty array ([]) in this request | [optional] 
**Annotations** | Pointer to **[]string** | Product Specification (anotations) By Category ID. The value is array of annotations id that can be retrieve by hit endpoint Get Product Annotation by Category ID. The location of id is at values.id. Remove annotations from related product by passing empty array (“”) in this request | [optional] 
**PriceCurrency** | Pointer to **string** | Currency code for stated price (IDR or USD) | [optional] 
**Weight** | Pointer to **float64** | Weight of the product | [optional] 
**WeightUnit** | Pointer to **string** | The unit of the weight with the following available value GR (gram) | [optional] [default to "GR"]
**IsFreeReturn** | Pointer to **bool** | Determine if the product can be returned (true) by buyers or not (false) | [optional] 
**IsMustInsurance** | **bool** | Determine if the product must be insured (true) or not (false) | 
**Menu** | Pointer to [**EditProductV3RequestProductsInnerMenu**](EditProductV3RequestProductsInnerMenu.md) |  | [optional] 
**Pictures** | Pointer to [**[]EditProductV3RequestProductsInnerPicturesInner**](EditProductV3RequestProductsInnerPicturesInner.md) | Images information of the product. The object keys includes: file_path. Remove pictures from related product by passing empty array ([]) in this request  | [optional] 
**Wholesale** | Pointer to [**[]EditProductV3RequestProductsInnerWholesaleInner**](EditProductV3RequestProductsInnerWholesaleInner.md) | Wholesale price and quantity of the product. The object keys includes: min_qty and price. Remove wholesale from related product by passing empty array ([]) in this request  | [optional] 
**Preorder** | Pointer to [**EditProductV3RequestProductsInnerPreorder**](EditProductV3RequestProductsInnerPreorder.md) |  | [optional] 
**Videos** | Pointer to [**[]EditProductV3RequestProductsInnerVideosInner**](EditProductV3RequestProductsInnerVideosInner.md) | Video link of the product. The object keys includes: url and source. url should only contain the YouTube video id e.g. dQw4w9WgXcQ. Where the type type should be youtube. Remove videos from related product by passing empty array ([]) in this request  | [optional] 
**Etalase** | Pointer to [**EditProductV3RequestProductsInnerEtalase**](EditProductV3RequestProductsInnerEtalase.md) |  | [optional] 
**Variant** | Pointer to [**EditProductV3RequestProductsInnerVariant**](EditProductV3RequestProductsInnerVariant.md) |  | [optional] 

## Methods

### NewEditProductV3RequestProductsInner

`func NewEditProductV3RequestProductsInner(id int64, status string, isMustInsurance bool, ) *EditProductV3RequestProductsInner`

NewEditProductV3RequestProductsInner instantiates a new EditProductV3RequestProductsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditProductV3RequestProductsInnerWithDefaults

`func NewEditProductV3RequestProductsInnerWithDefaults() *EditProductV3RequestProductsInner`

NewEditProductV3RequestProductsInnerWithDefaults instantiates a new EditProductV3RequestProductsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EditProductV3RequestProductsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EditProductV3RequestProductsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EditProductV3RequestProductsInner) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *EditProductV3RequestProductsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EditProductV3RequestProductsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EditProductV3RequestProductsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EditProductV3RequestProductsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCondition

`func (o *EditProductV3RequestProductsInner) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *EditProductV3RequestProductsInner) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *EditProductV3RequestProductsInner) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *EditProductV3RequestProductsInner) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetDescription

`func (o *EditProductV3RequestProductsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EditProductV3RequestProductsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EditProductV3RequestProductsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EditProductV3RequestProductsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSku

`func (o *EditProductV3RequestProductsInner) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *EditProductV3RequestProductsInner) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *EditProductV3RequestProductsInner) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *EditProductV3RequestProductsInner) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetPrice

`func (o *EditProductV3RequestProductsInner) GetPrice() int64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *EditProductV3RequestProductsInner) GetPriceOk() (*int64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *EditProductV3RequestProductsInner) SetPrice(v int64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *EditProductV3RequestProductsInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetStatus

`func (o *EditProductV3RequestProductsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EditProductV3RequestProductsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EditProductV3RequestProductsInner) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStock

`func (o *EditProductV3RequestProductsInner) GetStock() int64`

GetStock returns the Stock field if non-nil, zero value otherwise.

### GetStockOk

`func (o *EditProductV3RequestProductsInner) GetStockOk() (*int64, bool)`

GetStockOk returns a tuple with the Stock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStock

`func (o *EditProductV3RequestProductsInner) SetStock(v int64)`

SetStock sets Stock field to given value.

### HasStock

`func (o *EditProductV3RequestProductsInner) HasStock() bool`

HasStock returns a boolean if a field has been set.

### GetMinOrder

`func (o *EditProductV3RequestProductsInner) GetMinOrder() int64`

GetMinOrder returns the MinOrder field if non-nil, zero value otherwise.

### GetMinOrderOk

`func (o *EditProductV3RequestProductsInner) GetMinOrderOk() (*int64, bool)`

GetMinOrderOk returns a tuple with the MinOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinOrder

`func (o *EditProductV3RequestProductsInner) SetMinOrder(v int64)`

SetMinOrder sets MinOrder field to given value.

### HasMinOrder

`func (o *EditProductV3RequestProductsInner) HasMinOrder() bool`

HasMinOrder returns a boolean if a field has been set.

### GetCategoryId

`func (o *EditProductV3RequestProductsInner) GetCategoryId() int64`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *EditProductV3RequestProductsInner) GetCategoryIdOk() (*int64, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *EditProductV3RequestProductsInner) SetCategoryId(v int64)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *EditProductV3RequestProductsInner) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### GetDimension

`func (o *EditProductV3RequestProductsInner) GetDimension() EditProductV3RequestProductsInnerDimension`

GetDimension returns the Dimension field if non-nil, zero value otherwise.

### GetDimensionOk

`func (o *EditProductV3RequestProductsInner) GetDimensionOk() (*EditProductV3RequestProductsInnerDimension, bool)`

GetDimensionOk returns a tuple with the Dimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimension

`func (o *EditProductV3RequestProductsInner) SetDimension(v EditProductV3RequestProductsInnerDimension)`

SetDimension sets Dimension field to given value.

### HasDimension

`func (o *EditProductV3RequestProductsInner) HasDimension() bool`

HasDimension returns a boolean if a field has been set.

### GetCustomProductLogistics

`func (o *EditProductV3RequestProductsInner) GetCustomProductLogistics() []int64`

GetCustomProductLogistics returns the CustomProductLogistics field if non-nil, zero value otherwise.

### GetCustomProductLogisticsOk

`func (o *EditProductV3RequestProductsInner) GetCustomProductLogisticsOk() (*[]int64, bool)`

GetCustomProductLogisticsOk returns a tuple with the CustomProductLogistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProductLogistics

`func (o *EditProductV3RequestProductsInner) SetCustomProductLogistics(v []int64)`

SetCustomProductLogistics sets CustomProductLogistics field to given value.

### HasCustomProductLogistics

`func (o *EditProductV3RequestProductsInner) HasCustomProductLogistics() bool`

HasCustomProductLogistics returns a boolean if a field has been set.

### GetAnnotations

`func (o *EditProductV3RequestProductsInner) GetAnnotations() []string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *EditProductV3RequestProductsInner) GetAnnotationsOk() (*[]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *EditProductV3RequestProductsInner) SetAnnotations(v []string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *EditProductV3RequestProductsInner) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetPriceCurrency

`func (o *EditProductV3RequestProductsInner) GetPriceCurrency() string`

GetPriceCurrency returns the PriceCurrency field if non-nil, zero value otherwise.

### GetPriceCurrencyOk

`func (o *EditProductV3RequestProductsInner) GetPriceCurrencyOk() (*string, bool)`

GetPriceCurrencyOk returns a tuple with the PriceCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceCurrency

`func (o *EditProductV3RequestProductsInner) SetPriceCurrency(v string)`

SetPriceCurrency sets PriceCurrency field to given value.

### HasPriceCurrency

`func (o *EditProductV3RequestProductsInner) HasPriceCurrency() bool`

HasPriceCurrency returns a boolean if a field has been set.

### GetWeight

`func (o *EditProductV3RequestProductsInner) GetWeight() float64`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *EditProductV3RequestProductsInner) GetWeightOk() (*float64, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *EditProductV3RequestProductsInner) SetWeight(v float64)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *EditProductV3RequestProductsInner) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetWeightUnit

`func (o *EditProductV3RequestProductsInner) GetWeightUnit() string`

GetWeightUnit returns the WeightUnit field if non-nil, zero value otherwise.

### GetWeightUnitOk

`func (o *EditProductV3RequestProductsInner) GetWeightUnitOk() (*string, bool)`

GetWeightUnitOk returns a tuple with the WeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightUnit

`func (o *EditProductV3RequestProductsInner) SetWeightUnit(v string)`

SetWeightUnit sets WeightUnit field to given value.

### HasWeightUnit

`func (o *EditProductV3RequestProductsInner) HasWeightUnit() bool`

HasWeightUnit returns a boolean if a field has been set.

### GetIsFreeReturn

`func (o *EditProductV3RequestProductsInner) GetIsFreeReturn() bool`

GetIsFreeReturn returns the IsFreeReturn field if non-nil, zero value otherwise.

### GetIsFreeReturnOk

`func (o *EditProductV3RequestProductsInner) GetIsFreeReturnOk() (*bool, bool)`

GetIsFreeReturnOk returns a tuple with the IsFreeReturn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFreeReturn

`func (o *EditProductV3RequestProductsInner) SetIsFreeReturn(v bool)`

SetIsFreeReturn sets IsFreeReturn field to given value.

### HasIsFreeReturn

`func (o *EditProductV3RequestProductsInner) HasIsFreeReturn() bool`

HasIsFreeReturn returns a boolean if a field has been set.

### GetIsMustInsurance

`func (o *EditProductV3RequestProductsInner) GetIsMustInsurance() bool`

GetIsMustInsurance returns the IsMustInsurance field if non-nil, zero value otherwise.

### GetIsMustInsuranceOk

`func (o *EditProductV3RequestProductsInner) GetIsMustInsuranceOk() (*bool, bool)`

GetIsMustInsuranceOk returns a tuple with the IsMustInsurance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMustInsurance

`func (o *EditProductV3RequestProductsInner) SetIsMustInsurance(v bool)`

SetIsMustInsurance sets IsMustInsurance field to given value.


### GetMenu

`func (o *EditProductV3RequestProductsInner) GetMenu() EditProductV3RequestProductsInnerMenu`

GetMenu returns the Menu field if non-nil, zero value otherwise.

### GetMenuOk

`func (o *EditProductV3RequestProductsInner) GetMenuOk() (*EditProductV3RequestProductsInnerMenu, bool)`

GetMenuOk returns a tuple with the Menu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMenu

`func (o *EditProductV3RequestProductsInner) SetMenu(v EditProductV3RequestProductsInnerMenu)`

SetMenu sets Menu field to given value.

### HasMenu

`func (o *EditProductV3RequestProductsInner) HasMenu() bool`

HasMenu returns a boolean if a field has been set.

### GetPictures

`func (o *EditProductV3RequestProductsInner) GetPictures() []EditProductV3RequestProductsInnerPicturesInner`

GetPictures returns the Pictures field if non-nil, zero value otherwise.

### GetPicturesOk

`func (o *EditProductV3RequestProductsInner) GetPicturesOk() (*[]EditProductV3RequestProductsInnerPicturesInner, bool)`

GetPicturesOk returns a tuple with the Pictures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPictures

`func (o *EditProductV3RequestProductsInner) SetPictures(v []EditProductV3RequestProductsInnerPicturesInner)`

SetPictures sets Pictures field to given value.

### HasPictures

`func (o *EditProductV3RequestProductsInner) HasPictures() bool`

HasPictures returns a boolean if a field has been set.

### GetWholesale

`func (o *EditProductV3RequestProductsInner) GetWholesale() []EditProductV3RequestProductsInnerWholesaleInner`

GetWholesale returns the Wholesale field if non-nil, zero value otherwise.

### GetWholesaleOk

`func (o *EditProductV3RequestProductsInner) GetWholesaleOk() (*[]EditProductV3RequestProductsInnerWholesaleInner, bool)`

GetWholesaleOk returns a tuple with the Wholesale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWholesale

`func (o *EditProductV3RequestProductsInner) SetWholesale(v []EditProductV3RequestProductsInnerWholesaleInner)`

SetWholesale sets Wholesale field to given value.

### HasWholesale

`func (o *EditProductV3RequestProductsInner) HasWholesale() bool`

HasWholesale returns a boolean if a field has been set.

### GetPreorder

`func (o *EditProductV3RequestProductsInner) GetPreorder() EditProductV3RequestProductsInnerPreorder`

GetPreorder returns the Preorder field if non-nil, zero value otherwise.

### GetPreorderOk

`func (o *EditProductV3RequestProductsInner) GetPreorderOk() (*EditProductV3RequestProductsInnerPreorder, bool)`

GetPreorderOk returns a tuple with the Preorder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreorder

`func (o *EditProductV3RequestProductsInner) SetPreorder(v EditProductV3RequestProductsInnerPreorder)`

SetPreorder sets Preorder field to given value.

### HasPreorder

`func (o *EditProductV3RequestProductsInner) HasPreorder() bool`

HasPreorder returns a boolean if a field has been set.

### GetVideos

`func (o *EditProductV3RequestProductsInner) GetVideos() []EditProductV3RequestProductsInnerVideosInner`

GetVideos returns the Videos field if non-nil, zero value otherwise.

### GetVideosOk

`func (o *EditProductV3RequestProductsInner) GetVideosOk() (*[]EditProductV3RequestProductsInnerVideosInner, bool)`

GetVideosOk returns a tuple with the Videos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideos

`func (o *EditProductV3RequestProductsInner) SetVideos(v []EditProductV3RequestProductsInnerVideosInner)`

SetVideos sets Videos field to given value.

### HasVideos

`func (o *EditProductV3RequestProductsInner) HasVideos() bool`

HasVideos returns a boolean if a field has been set.

### GetEtalase

`func (o *EditProductV3RequestProductsInner) GetEtalase() EditProductV3RequestProductsInnerEtalase`

GetEtalase returns the Etalase field if non-nil, zero value otherwise.

### GetEtalaseOk

`func (o *EditProductV3RequestProductsInner) GetEtalaseOk() (*EditProductV3RequestProductsInnerEtalase, bool)`

GetEtalaseOk returns a tuple with the Etalase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtalase

`func (o *EditProductV3RequestProductsInner) SetEtalase(v EditProductV3RequestProductsInnerEtalase)`

SetEtalase sets Etalase field to given value.

### HasEtalase

`func (o *EditProductV3RequestProductsInner) HasEtalase() bool`

HasEtalase returns a boolean if a field has been set.

### GetVariant

`func (o *EditProductV3RequestProductsInner) GetVariant() EditProductV3RequestProductsInnerVariant`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *EditProductV3RequestProductsInner) GetVariantOk() (*EditProductV3RequestProductsInnerVariant, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *EditProductV3RequestProductsInner) SetVariant(v EditProductV3RequestProductsInnerVariant)`

SetVariant sets Variant field to given value.

### HasVariant

`func (o *EditProductV3RequestProductsInner) HasVariant() bool`

HasVariant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


