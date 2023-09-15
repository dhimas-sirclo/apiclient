# ProductInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Name** | **string** |  | 
**Brand** | **string** |  | 
**ProductId** | **string** |  | 
**ProductCode** | **string** |  | 
**ProductDescription** | **string** |  | 
**Condition** | **string** |  | 
**ShopId** | **string** |  | 
**CategoryId** | **string** |  | 
**CategoryName** | **string** |  | 
**Variants** | [**[]VariantInput**](VariantInput.md) |  | 

## Methods

### NewProductInput

`func NewProductInput(type_ string, name string, brand string, productId string, productCode string, productDescription string, condition string, shopId string, categoryId string, categoryName string, variants []VariantInput, ) *ProductInput`

NewProductInput instantiates a new ProductInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductInputWithDefaults

`func NewProductInputWithDefaults() *ProductInput`

NewProductInputWithDefaults instantiates a new ProductInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ProductInput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProductInput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProductInput) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *ProductInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductInput) SetName(v string)`

SetName sets Name field to given value.


### GetBrand

`func (o *ProductInput) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *ProductInput) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *ProductInput) SetBrand(v string)`

SetBrand sets Brand field to given value.


### GetProductId

`func (o *ProductInput) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *ProductInput) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *ProductInput) SetProductId(v string)`

SetProductId sets ProductId field to given value.


### GetProductCode

`func (o *ProductInput) GetProductCode() string`

GetProductCode returns the ProductCode field if non-nil, zero value otherwise.

### GetProductCodeOk

`func (o *ProductInput) GetProductCodeOk() (*string, bool)`

GetProductCodeOk returns a tuple with the ProductCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCode

`func (o *ProductInput) SetProductCode(v string)`

SetProductCode sets ProductCode field to given value.


### GetProductDescription

`func (o *ProductInput) GetProductDescription() string`

GetProductDescription returns the ProductDescription field if non-nil, zero value otherwise.

### GetProductDescriptionOk

`func (o *ProductInput) GetProductDescriptionOk() (*string, bool)`

GetProductDescriptionOk returns a tuple with the ProductDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductDescription

`func (o *ProductInput) SetProductDescription(v string)`

SetProductDescription sets ProductDescription field to given value.


### GetCondition

`func (o *ProductInput) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *ProductInput) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *ProductInput) SetCondition(v string)`

SetCondition sets Condition field to given value.


### GetShopId

`func (o *ProductInput) GetShopId() string`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *ProductInput) GetShopIdOk() (*string, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *ProductInput) SetShopId(v string)`

SetShopId sets ShopId field to given value.


### GetCategoryId

`func (o *ProductInput) GetCategoryId() string`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *ProductInput) GetCategoryIdOk() (*string, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *ProductInput) SetCategoryId(v string)`

SetCategoryId sets CategoryId field to given value.


### GetCategoryName

`func (o *ProductInput) GetCategoryName() string`

GetCategoryName returns the CategoryName field if non-nil, zero value otherwise.

### GetCategoryNameOk

`func (o *ProductInput) GetCategoryNameOk() (*string, bool)`

GetCategoryNameOk returns a tuple with the CategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryName

`func (o *ProductInput) SetCategoryName(v string)`

SetCategoryName sets CategoryName field to given value.


### GetVariants

`func (o *ProductInput) GetVariants() []VariantInput`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *ProductInput) GetVariantsOk() (*[]VariantInput, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *ProductInput) SetVariants(v []VariantInput)`

SetVariants sets Variants field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


