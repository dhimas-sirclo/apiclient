# ProductCategoryInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParentCategoryId** | Pointer to **string** |  | [optional] 
**CategoryId** | **string** |  | 
**CategoryName** | **string** |  | 
**IsLeaf** | **bool** |  | 
**Attribute** | [**[]ProductCategoryAttributeInput**](ProductCategoryAttributeInput.md) |  | 

## Methods

### NewProductCategoryInput

`func NewProductCategoryInput(categoryId string, categoryName string, isLeaf bool, attribute []ProductCategoryAttributeInput, ) *ProductCategoryInput`

NewProductCategoryInput instantiates a new ProductCategoryInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductCategoryInputWithDefaults

`func NewProductCategoryInputWithDefaults() *ProductCategoryInput`

NewProductCategoryInputWithDefaults instantiates a new ProductCategoryInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParentCategoryId

`func (o *ProductCategoryInput) GetParentCategoryId() string`

GetParentCategoryId returns the ParentCategoryId field if non-nil, zero value otherwise.

### GetParentCategoryIdOk

`func (o *ProductCategoryInput) GetParentCategoryIdOk() (*string, bool)`

GetParentCategoryIdOk returns a tuple with the ParentCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCategoryId

`func (o *ProductCategoryInput) SetParentCategoryId(v string)`

SetParentCategoryId sets ParentCategoryId field to given value.

### HasParentCategoryId

`func (o *ProductCategoryInput) HasParentCategoryId() bool`

HasParentCategoryId returns a boolean if a field has been set.

### GetCategoryId

`func (o *ProductCategoryInput) GetCategoryId() string`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *ProductCategoryInput) GetCategoryIdOk() (*string, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *ProductCategoryInput) SetCategoryId(v string)`

SetCategoryId sets CategoryId field to given value.


### GetCategoryName

`func (o *ProductCategoryInput) GetCategoryName() string`

GetCategoryName returns the CategoryName field if non-nil, zero value otherwise.

### GetCategoryNameOk

`func (o *ProductCategoryInput) GetCategoryNameOk() (*string, bool)`

GetCategoryNameOk returns a tuple with the CategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryName

`func (o *ProductCategoryInput) SetCategoryName(v string)`

SetCategoryName sets CategoryName field to given value.


### GetIsLeaf

`func (o *ProductCategoryInput) GetIsLeaf() bool`

GetIsLeaf returns the IsLeaf field if non-nil, zero value otherwise.

### GetIsLeafOk

`func (o *ProductCategoryInput) GetIsLeafOk() (*bool, bool)`

GetIsLeafOk returns a tuple with the IsLeaf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLeaf

`func (o *ProductCategoryInput) SetIsLeaf(v bool)`

SetIsLeaf sets IsLeaf field to given value.


### GetAttribute

`func (o *ProductCategoryInput) GetAttribute() []ProductCategoryAttributeInput`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *ProductCategoryInput) GetAttributeOk() (*[]ProductCategoryAttributeInput, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *ProductCategoryInput) SetAttribute(v []ProductCategoryAttributeInput)`

SetAttribute sets Attribute field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


