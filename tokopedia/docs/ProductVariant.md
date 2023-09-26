# ProductVariant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsParent** | Pointer to **bool** | Is Product Variant Parent? | [optional] 
**IsVariant** | Pointer to **bool** | Is Product Has Variants? | [optional] 
**ChildrenID** | Pointer to **[]int64** | List of Product Variant ID | [optional] 

## Methods

### NewProductVariant

`func NewProductVariant() *ProductVariant`

NewProductVariant instantiates a new ProductVariant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductVariantWithDefaults

`func NewProductVariantWithDefaults() *ProductVariant`

NewProductVariantWithDefaults instantiates a new ProductVariant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsParent

`func (o *ProductVariant) GetIsParent() bool`

GetIsParent returns the IsParent field if non-nil, zero value otherwise.

### GetIsParentOk

`func (o *ProductVariant) GetIsParentOk() (*bool, bool)`

GetIsParentOk returns a tuple with the IsParent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsParent

`func (o *ProductVariant) SetIsParent(v bool)`

SetIsParent sets IsParent field to given value.

### HasIsParent

`func (o *ProductVariant) HasIsParent() bool`

HasIsParent returns a boolean if a field has been set.

### GetIsVariant

`func (o *ProductVariant) GetIsVariant() bool`

GetIsVariant returns the IsVariant field if non-nil, zero value otherwise.

### GetIsVariantOk

`func (o *ProductVariant) GetIsVariantOk() (*bool, bool)`

GetIsVariantOk returns a tuple with the IsVariant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVariant

`func (o *ProductVariant) SetIsVariant(v bool)`

SetIsVariant sets IsVariant field to given value.

### HasIsVariant

`func (o *ProductVariant) HasIsVariant() bool`

HasIsVariant returns a boolean if a field has been set.

### GetChildrenID

`func (o *ProductVariant) GetChildrenID() []int64`

GetChildrenID returns the ChildrenID field if non-nil, zero value otherwise.

### GetChildrenIDOk

`func (o *ProductVariant) GetChildrenIDOk() (*[]int64, bool)`

GetChildrenIDOk returns a tuple with the ChildrenID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildrenID

`func (o *ProductVariant) SetChildrenID(v []int64)`

SetChildrenID sets ChildrenID field to given value.

### HasChildrenID

`func (o *ProductVariant) HasChildrenID() bool`

HasChildrenID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


