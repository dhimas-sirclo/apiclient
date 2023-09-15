# ProductCategory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryId** | Pointer to **string** |  | [optional] 
**Attribute** | Pointer to [**[]ProductCategoryAttribute**](ProductCategoryAttribute.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewProductCategory

`func NewProductCategory() *ProductCategory`

NewProductCategory instantiates a new ProductCategory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductCategoryWithDefaults

`func NewProductCategoryWithDefaults() *ProductCategory`

NewProductCategoryWithDefaults instantiates a new ProductCategory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoryId

`func (o *ProductCategory) GetCategoryId() string`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *ProductCategory) GetCategoryIdOk() (*string, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *ProductCategory) SetCategoryId(v string)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *ProductCategory) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### GetAttribute

`func (o *ProductCategory) GetAttribute() []ProductCategoryAttribute`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *ProductCategory) GetAttributeOk() (*[]ProductCategoryAttribute, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *ProductCategory) SetAttribute(v []ProductCategoryAttribute)`

SetAttribute sets Attribute field to given value.

### HasAttribute

`func (o *ProductCategory) HasAttribute() bool`

HasAttribute returns a boolean if a field has been set.

### GetMessage

`func (o *ProductCategory) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ProductCategory) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ProductCategory) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ProductCategory) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


