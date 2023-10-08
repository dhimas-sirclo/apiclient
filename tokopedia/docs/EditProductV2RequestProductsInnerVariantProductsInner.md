# EditProductV2RequestProductsInnerVariantProductsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsPrimary** | **bool** |  | 
**Status** | Pointer to **string** |  | [optional] 
**Price** | **float64** |  | 
**Stock** | **int64** |  | 
**Sku** | **string** |  | 
**Weight** | **float64** |  | 
**WeightUnit** | **string** |  | [default to "GR"]
**Combination** | **[]int64** |  | 
**Pictures** | [**[]EditProductV2RequestProductsInnerVariantProductsInnerPicturesInner**](EditProductV2RequestProductsInnerVariantProductsInnerPicturesInner.md) |  | 

## Methods

### NewEditProductV2RequestProductsInnerVariantProductsInner

`func NewEditProductV2RequestProductsInnerVariantProductsInner(isPrimary bool, price float64, stock int64, sku string, weight float64, weightUnit string, combination []int64, pictures []EditProductV2RequestProductsInnerVariantProductsInnerPicturesInner, ) *EditProductV2RequestProductsInnerVariantProductsInner`

NewEditProductV2RequestProductsInnerVariantProductsInner instantiates a new EditProductV2RequestProductsInnerVariantProductsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditProductV2RequestProductsInnerVariantProductsInnerWithDefaults

`func NewEditProductV2RequestProductsInnerVariantProductsInnerWithDefaults() *EditProductV2RequestProductsInnerVariantProductsInner`

NewEditProductV2RequestProductsInnerVariantProductsInnerWithDefaults instantiates a new EditProductV2RequestProductsInnerVariantProductsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsPrimary

`func (o *EditProductV2RequestProductsInnerVariantProductsInner) GetIsPrimary() bool`

GetIsPrimary returns the IsPrimary field if non-nil, zero value otherwise.

### GetIsPrimaryOk

`func (o *EditProductV2RequestProductsInnerVariantProductsInner) GetIsPrimaryOk() (*bool, bool)`

GetIsPrimaryOk returns a tuple with the IsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimary

`func (o *EditProductV2RequestProductsInnerVariantProductsInner) SetIsPrimary(v bool)`

SetIsPrimary sets IsPrimary field to given value.


### GetStatus

`func (o *EditProductV2RequestProductsInnerVariantProductsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EditProductV2RequestProductsInnerVariantProductsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EditProductV2RequestProductsInnerVariantProductsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EditProductV2RequestProductsInnerVariantProductsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPrice

`func (o *EditProductV2RequestProductsInnerVariantProductsInner) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *EditProductV2RequestProductsInnerVariantProductsInner) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *EditProductV2RequestProductsInnerVariantProductsInner) SetPrice(v float64)`

SetPrice sets Price field to given value.


### GetStock

`func (o *EditProductV2RequestProductsInnerVariantProductsInner) GetStock() int64`

GetStock returns the Stock field if non-nil, zero value otherwise.

### GetStockOk

`func (o *EditProductV2RequestProductsInnerVariantProductsInner) GetStockOk() (*int64, bool)`

GetStockOk returns a tuple with the Stock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStock

`func (o *EditProductV2RequestProductsInnerVariantProductsInner) SetStock(v int64)`

SetStock sets Stock field to given value.


### GetSku

`func (o *EditProductV2RequestProductsInnerVariantProductsInner) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *EditProductV2RequestProductsInnerVariantProductsInner) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *EditProductV2RequestProductsInnerVariantProductsInner) SetSku(v string)`

SetSku sets Sku field to given value.


### GetWeight

`func (o *EditProductV2RequestProductsInnerVariantProductsInner) GetWeight() float64`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *EditProductV2RequestProductsInnerVariantProductsInner) GetWeightOk() (*float64, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *EditProductV2RequestProductsInnerVariantProductsInner) SetWeight(v float64)`

SetWeight sets Weight field to given value.


### GetWeightUnit

`func (o *EditProductV2RequestProductsInnerVariantProductsInner) GetWeightUnit() string`

GetWeightUnit returns the WeightUnit field if non-nil, zero value otherwise.

### GetWeightUnitOk

`func (o *EditProductV2RequestProductsInnerVariantProductsInner) GetWeightUnitOk() (*string, bool)`

GetWeightUnitOk returns a tuple with the WeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightUnit

`func (o *EditProductV2RequestProductsInnerVariantProductsInner) SetWeightUnit(v string)`

SetWeightUnit sets WeightUnit field to given value.


### GetCombination

`func (o *EditProductV2RequestProductsInnerVariantProductsInner) GetCombination() []int64`

GetCombination returns the Combination field if non-nil, zero value otherwise.

### GetCombinationOk

`func (o *EditProductV2RequestProductsInnerVariantProductsInner) GetCombinationOk() (*[]int64, bool)`

GetCombinationOk returns a tuple with the Combination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCombination

`func (o *EditProductV2RequestProductsInnerVariantProductsInner) SetCombination(v []int64)`

SetCombination sets Combination field to given value.


### GetPictures

`func (o *EditProductV2RequestProductsInnerVariantProductsInner) GetPictures() []EditProductV2RequestProductsInnerVariantProductsInnerPicturesInner`

GetPictures returns the Pictures field if non-nil, zero value otherwise.

### GetPicturesOk

`func (o *EditProductV2RequestProductsInnerVariantProductsInner) GetPicturesOk() (*[]EditProductV2RequestProductsInnerVariantProductsInnerPicturesInner, bool)`

GetPicturesOk returns a tuple with the Pictures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPictures

`func (o *EditProductV2RequestProductsInnerVariantProductsInner) SetPictures(v []EditProductV2RequestProductsInnerVariantProductsInnerPicturesInner)`

SetPictures sets Pictures field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


