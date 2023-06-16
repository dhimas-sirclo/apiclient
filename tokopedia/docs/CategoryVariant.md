# CategoryVariant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryId** | Pointer to **int64** |  | [optional] 
**VariantIdCombinations** | Pointer to **[][]int64** |  | [optional] 
**VariantDetails** | Pointer to [**[]CategoryVariantDetail**](CategoryVariantDetail.md) |  | [optional] 

## Methods

### NewCategoryVariant

`func NewCategoryVariant() *CategoryVariant`

NewCategoryVariant instantiates a new CategoryVariant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryVariantWithDefaults

`func NewCategoryVariantWithDefaults() *CategoryVariant`

NewCategoryVariantWithDefaults instantiates a new CategoryVariant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoryId

`func (o *CategoryVariant) GetCategoryId() int64`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *CategoryVariant) GetCategoryIdOk() (*int64, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *CategoryVariant) SetCategoryId(v int64)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *CategoryVariant) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### GetVariantIdCombinations

`func (o *CategoryVariant) GetVariantIdCombinations() [][]int64`

GetVariantIdCombinations returns the VariantIdCombinations field if non-nil, zero value otherwise.

### GetVariantIdCombinationsOk

`func (o *CategoryVariant) GetVariantIdCombinationsOk() (*[][]int64, bool)`

GetVariantIdCombinationsOk returns a tuple with the VariantIdCombinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantIdCombinations

`func (o *CategoryVariant) SetVariantIdCombinations(v [][]int64)`

SetVariantIdCombinations sets VariantIdCombinations field to given value.

### HasVariantIdCombinations

`func (o *CategoryVariant) HasVariantIdCombinations() bool`

HasVariantIdCombinations returns a boolean if a field has been set.

### GetVariantDetails

`func (o *CategoryVariant) GetVariantDetails() []CategoryVariantDetail`

GetVariantDetails returns the VariantDetails field if non-nil, zero value otherwise.

### GetVariantDetailsOk

`func (o *CategoryVariant) GetVariantDetailsOk() (*[]CategoryVariantDetail, bool)`

GetVariantDetailsOk returns a tuple with the VariantDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantDetails

`func (o *CategoryVariant) SetVariantDetails(v []CategoryVariantDetail)`

SetVariantDetails sets VariantDetails field to given value.

### HasVariantDetails

`func (o *CategoryVariant) HasVariantDetails() bool`

HasVariantDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


