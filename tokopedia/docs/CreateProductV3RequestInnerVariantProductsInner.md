# CreateProductV3RequestInnerVariantProductsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsPrimary** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] [default to "LIMITED"]
**Price** | Pointer to **float64** |  | [optional] 
**Stock** | Pointer to **int64** |  | [optional] 
**Sku** | Pointer to **string** |  | [optional] 
**Combination** | Pointer to **[]int64** |  | [optional] 
**Pictures** | Pointer to [**[]EditProductV3RequestProductsInnerPicturesInner**](EditProductV3RequestProductsInnerPicturesInner.md) |  | [optional] 

## Methods

### NewCreateProductV3RequestInnerVariantProductsInner

`func NewCreateProductV3RequestInnerVariantProductsInner() *CreateProductV3RequestInnerVariantProductsInner`

NewCreateProductV3RequestInnerVariantProductsInner instantiates a new CreateProductV3RequestInnerVariantProductsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProductV3RequestInnerVariantProductsInnerWithDefaults

`func NewCreateProductV3RequestInnerVariantProductsInnerWithDefaults() *CreateProductV3RequestInnerVariantProductsInner`

NewCreateProductV3RequestInnerVariantProductsInnerWithDefaults instantiates a new CreateProductV3RequestInnerVariantProductsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsPrimary

`func (o *CreateProductV3RequestInnerVariantProductsInner) GetIsPrimary() bool`

GetIsPrimary returns the IsPrimary field if non-nil, zero value otherwise.

### GetIsPrimaryOk

`func (o *CreateProductV3RequestInnerVariantProductsInner) GetIsPrimaryOk() (*bool, bool)`

GetIsPrimaryOk returns a tuple with the IsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimary

`func (o *CreateProductV3RequestInnerVariantProductsInner) SetIsPrimary(v bool)`

SetIsPrimary sets IsPrimary field to given value.

### HasIsPrimary

`func (o *CreateProductV3RequestInnerVariantProductsInner) HasIsPrimary() bool`

HasIsPrimary returns a boolean if a field has been set.

### GetStatus

`func (o *CreateProductV3RequestInnerVariantProductsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateProductV3RequestInnerVariantProductsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateProductV3RequestInnerVariantProductsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateProductV3RequestInnerVariantProductsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPrice

`func (o *CreateProductV3RequestInnerVariantProductsInner) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CreateProductV3RequestInnerVariantProductsInner) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CreateProductV3RequestInnerVariantProductsInner) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *CreateProductV3RequestInnerVariantProductsInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetStock

`func (o *CreateProductV3RequestInnerVariantProductsInner) GetStock() int64`

GetStock returns the Stock field if non-nil, zero value otherwise.

### GetStockOk

`func (o *CreateProductV3RequestInnerVariantProductsInner) GetStockOk() (*int64, bool)`

GetStockOk returns a tuple with the Stock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStock

`func (o *CreateProductV3RequestInnerVariantProductsInner) SetStock(v int64)`

SetStock sets Stock field to given value.

### HasStock

`func (o *CreateProductV3RequestInnerVariantProductsInner) HasStock() bool`

HasStock returns a boolean if a field has been set.

### GetSku

`func (o *CreateProductV3RequestInnerVariantProductsInner) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *CreateProductV3RequestInnerVariantProductsInner) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *CreateProductV3RequestInnerVariantProductsInner) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *CreateProductV3RequestInnerVariantProductsInner) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetCombination

`func (o *CreateProductV3RequestInnerVariantProductsInner) GetCombination() []int64`

GetCombination returns the Combination field if non-nil, zero value otherwise.

### GetCombinationOk

`func (o *CreateProductV3RequestInnerVariantProductsInner) GetCombinationOk() (*[]int64, bool)`

GetCombinationOk returns a tuple with the Combination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCombination

`func (o *CreateProductV3RequestInnerVariantProductsInner) SetCombination(v []int64)`

SetCombination sets Combination field to given value.

### HasCombination

`func (o *CreateProductV3RequestInnerVariantProductsInner) HasCombination() bool`

HasCombination returns a boolean if a field has been set.

### GetPictures

`func (o *CreateProductV3RequestInnerVariantProductsInner) GetPictures() []EditProductV3RequestProductsInnerPicturesInner`

GetPictures returns the Pictures field if non-nil, zero value otherwise.

### GetPicturesOk

`func (o *CreateProductV3RequestInnerVariantProductsInner) GetPicturesOk() (*[]EditProductV3RequestProductsInnerPicturesInner, bool)`

GetPicturesOk returns a tuple with the Pictures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPictures

`func (o *CreateProductV3RequestInnerVariantProductsInner) SetPictures(v []EditProductV3RequestProductsInnerPicturesInner)`

SetPictures sets Pictures field to given value.

### HasPictures

`func (o *CreateProductV3RequestInnerVariantProductsInner) HasPictures() bool`

HasPictures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


