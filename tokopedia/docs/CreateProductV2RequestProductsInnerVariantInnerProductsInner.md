# CreateProductV2RequestProductsInnerVariantInnerProductsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsPrimary** | **bool** |  | 
**Status** | **string** |  | [default to "LIMITED"]
**Price** | **float64** |  | 
**Stock** | **int64** |  | 
**Sku** | **string** |  | 
**Combination** | **[]int64** |  | 
**Pictures** | [**[]EditProductV2RequestProductsInnerVariantProductsInnerPicturesInner**](EditProductV2RequestProductsInnerVariantProductsInnerPicturesInner.md) |  | 

## Methods

### NewCreateProductV2RequestProductsInnerVariantInnerProductsInner

`func NewCreateProductV2RequestProductsInnerVariantInnerProductsInner(isPrimary bool, status string, price float64, stock int64, sku string, combination []int64, pictures []EditProductV2RequestProductsInnerVariantProductsInnerPicturesInner, ) *CreateProductV2RequestProductsInnerVariantInnerProductsInner`

NewCreateProductV2RequestProductsInnerVariantInnerProductsInner instantiates a new CreateProductV2RequestProductsInnerVariantInnerProductsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProductV2RequestProductsInnerVariantInnerProductsInnerWithDefaults

`func NewCreateProductV2RequestProductsInnerVariantInnerProductsInnerWithDefaults() *CreateProductV2RequestProductsInnerVariantInnerProductsInner`

NewCreateProductV2RequestProductsInnerVariantInnerProductsInnerWithDefaults instantiates a new CreateProductV2RequestProductsInnerVariantInnerProductsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsPrimary

`func (o *CreateProductV2RequestProductsInnerVariantInnerProductsInner) GetIsPrimary() bool`

GetIsPrimary returns the IsPrimary field if non-nil, zero value otherwise.

### GetIsPrimaryOk

`func (o *CreateProductV2RequestProductsInnerVariantInnerProductsInner) GetIsPrimaryOk() (*bool, bool)`

GetIsPrimaryOk returns a tuple with the IsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimary

`func (o *CreateProductV2RequestProductsInnerVariantInnerProductsInner) SetIsPrimary(v bool)`

SetIsPrimary sets IsPrimary field to given value.


### GetStatus

`func (o *CreateProductV2RequestProductsInnerVariantInnerProductsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateProductV2RequestProductsInnerVariantInnerProductsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateProductV2RequestProductsInnerVariantInnerProductsInner) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetPrice

`func (o *CreateProductV2RequestProductsInnerVariantInnerProductsInner) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CreateProductV2RequestProductsInnerVariantInnerProductsInner) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CreateProductV2RequestProductsInnerVariantInnerProductsInner) SetPrice(v float64)`

SetPrice sets Price field to given value.


### GetStock

`func (o *CreateProductV2RequestProductsInnerVariantInnerProductsInner) GetStock() int64`

GetStock returns the Stock field if non-nil, zero value otherwise.

### GetStockOk

`func (o *CreateProductV2RequestProductsInnerVariantInnerProductsInner) GetStockOk() (*int64, bool)`

GetStockOk returns a tuple with the Stock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStock

`func (o *CreateProductV2RequestProductsInnerVariantInnerProductsInner) SetStock(v int64)`

SetStock sets Stock field to given value.


### GetSku

`func (o *CreateProductV2RequestProductsInnerVariantInnerProductsInner) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *CreateProductV2RequestProductsInnerVariantInnerProductsInner) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *CreateProductV2RequestProductsInnerVariantInnerProductsInner) SetSku(v string)`

SetSku sets Sku field to given value.


### GetCombination

`func (o *CreateProductV2RequestProductsInnerVariantInnerProductsInner) GetCombination() []int64`

GetCombination returns the Combination field if non-nil, zero value otherwise.

### GetCombinationOk

`func (o *CreateProductV2RequestProductsInnerVariantInnerProductsInner) GetCombinationOk() (*[]int64, bool)`

GetCombinationOk returns a tuple with the Combination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCombination

`func (o *CreateProductV2RequestProductsInnerVariantInnerProductsInner) SetCombination(v []int64)`

SetCombination sets Combination field to given value.


### GetPictures

`func (o *CreateProductV2RequestProductsInnerVariantInnerProductsInner) GetPictures() []EditProductV2RequestProductsInnerVariantProductsInnerPicturesInner`

GetPictures returns the Pictures field if non-nil, zero value otherwise.

### GetPicturesOk

`func (o *CreateProductV2RequestProductsInnerVariantInnerProductsInner) GetPicturesOk() (*[]EditProductV2RequestProductsInnerVariantProductsInnerPicturesInner, bool)`

GetPicturesOk returns a tuple with the Pictures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPictures

`func (o *CreateProductV2RequestProductsInnerVariantInnerProductsInner) SetPictures(v []EditProductV2RequestProductsInnerVariantProductsInnerPicturesInner)`

SetPictures sets Pictures field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


