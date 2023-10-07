# EditProductV3RequestProductsInnerVariantProductsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsPrimary** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **int64** |  | [optional] 
**Stock** | Pointer to **int64** |  | [optional] 
**Sku** | Pointer to **string** |  | [optional] 
**Combination** | Pointer to **[]int64** |  | [optional] 
**Pictures** | Pointer to [**[]EditProductV3RequestProductsInnerPicturesInner**](EditProductV3RequestProductsInnerPicturesInner.md) |  | [optional] 

## Methods

### NewEditProductV3RequestProductsInnerVariantProductsInner

`func NewEditProductV3RequestProductsInnerVariantProductsInner() *EditProductV3RequestProductsInnerVariantProductsInner`

NewEditProductV3RequestProductsInnerVariantProductsInner instantiates a new EditProductV3RequestProductsInnerVariantProductsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditProductV3RequestProductsInnerVariantProductsInnerWithDefaults

`func NewEditProductV3RequestProductsInnerVariantProductsInnerWithDefaults() *EditProductV3RequestProductsInnerVariantProductsInner`

NewEditProductV3RequestProductsInnerVariantProductsInnerWithDefaults instantiates a new EditProductV3RequestProductsInnerVariantProductsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsPrimary

`func (o *EditProductV3RequestProductsInnerVariantProductsInner) GetIsPrimary() bool`

GetIsPrimary returns the IsPrimary field if non-nil, zero value otherwise.

### GetIsPrimaryOk

`func (o *EditProductV3RequestProductsInnerVariantProductsInner) GetIsPrimaryOk() (*bool, bool)`

GetIsPrimaryOk returns a tuple with the IsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimary

`func (o *EditProductV3RequestProductsInnerVariantProductsInner) SetIsPrimary(v bool)`

SetIsPrimary sets IsPrimary field to given value.

### HasIsPrimary

`func (o *EditProductV3RequestProductsInnerVariantProductsInner) HasIsPrimary() bool`

HasIsPrimary returns a boolean if a field has been set.

### GetStatus

`func (o *EditProductV3RequestProductsInnerVariantProductsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EditProductV3RequestProductsInnerVariantProductsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EditProductV3RequestProductsInnerVariantProductsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EditProductV3RequestProductsInnerVariantProductsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPrice

`func (o *EditProductV3RequestProductsInnerVariantProductsInner) GetPrice() int64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *EditProductV3RequestProductsInnerVariantProductsInner) GetPriceOk() (*int64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *EditProductV3RequestProductsInnerVariantProductsInner) SetPrice(v int64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *EditProductV3RequestProductsInnerVariantProductsInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetStock

`func (o *EditProductV3RequestProductsInnerVariantProductsInner) GetStock() int64`

GetStock returns the Stock field if non-nil, zero value otherwise.

### GetStockOk

`func (o *EditProductV3RequestProductsInnerVariantProductsInner) GetStockOk() (*int64, bool)`

GetStockOk returns a tuple with the Stock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStock

`func (o *EditProductV3RequestProductsInnerVariantProductsInner) SetStock(v int64)`

SetStock sets Stock field to given value.

### HasStock

`func (o *EditProductV3RequestProductsInnerVariantProductsInner) HasStock() bool`

HasStock returns a boolean if a field has been set.

### GetSku

`func (o *EditProductV3RequestProductsInnerVariantProductsInner) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *EditProductV3RequestProductsInnerVariantProductsInner) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *EditProductV3RequestProductsInnerVariantProductsInner) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *EditProductV3RequestProductsInnerVariantProductsInner) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetCombination

`func (o *EditProductV3RequestProductsInnerVariantProductsInner) GetCombination() []int64`

GetCombination returns the Combination field if non-nil, zero value otherwise.

### GetCombinationOk

`func (o *EditProductV3RequestProductsInnerVariantProductsInner) GetCombinationOk() (*[]int64, bool)`

GetCombinationOk returns a tuple with the Combination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCombination

`func (o *EditProductV3RequestProductsInnerVariantProductsInner) SetCombination(v []int64)`

SetCombination sets Combination field to given value.

### HasCombination

`func (o *EditProductV3RequestProductsInnerVariantProductsInner) HasCombination() bool`

HasCombination returns a boolean if a field has been set.

### GetPictures

`func (o *EditProductV3RequestProductsInnerVariantProductsInner) GetPictures() []EditProductV3RequestProductsInnerPicturesInner`

GetPictures returns the Pictures field if non-nil, zero value otherwise.

### GetPicturesOk

`func (o *EditProductV3RequestProductsInnerVariantProductsInner) GetPicturesOk() (*[]EditProductV3RequestProductsInnerPicturesInner, bool)`

GetPicturesOk returns a tuple with the Pictures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPictures

`func (o *EditProductV3RequestProductsInnerVariantProductsInner) SetPictures(v []EditProductV3RequestProductsInnerPicturesInner)`

SetPictures sets Pictures field to given value.

### HasPictures

`func (o *EditProductV3RequestProductsInnerVariantProductsInner) HasPictures() bool`

HasPictures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


