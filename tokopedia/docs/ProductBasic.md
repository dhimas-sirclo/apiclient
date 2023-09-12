# ProductBasic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductID** | Pointer to **int32** | Product Unique Identifier | [optional] 
**ShopID** | Pointer to **int32** | Shop Unique Identifier | [optional] 
**Status** | Pointer to **int32** | Product Status (e.g., -2 for Banned, -1 for Pending, 0 for Deleted, 1 for Active, 2 for Best (Feature Product), 3 for Inactive (Warehouse)) | [optional] 
**Name** | Pointer to **string** | Product Name | [optional] 
**Condition** | Pointer to **int32** | Product Condition (e.g., 1 for New, 2 for Used) | [optional] 
**ChildCategoryID** | Pointer to **int32** | Product Category Unique Identifier | [optional] 
**ShortDesc** | Pointer to **string** | Product Description | [optional] 

## Methods

### NewProductBasic

`func NewProductBasic() *ProductBasic`

NewProductBasic instantiates a new ProductBasic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductBasicWithDefaults

`func NewProductBasicWithDefaults() *ProductBasic`

NewProductBasicWithDefaults instantiates a new ProductBasic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductID

`func (o *ProductBasic) GetProductID() int32`

GetProductID returns the ProductID field if non-nil, zero value otherwise.

### GetProductIDOk

`func (o *ProductBasic) GetProductIDOk() (*int32, bool)`

GetProductIDOk returns a tuple with the ProductID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductID

`func (o *ProductBasic) SetProductID(v int32)`

SetProductID sets ProductID field to given value.

### HasProductID

`func (o *ProductBasic) HasProductID() bool`

HasProductID returns a boolean if a field has been set.

### GetShopID

`func (o *ProductBasic) GetShopID() int32`

GetShopID returns the ShopID field if non-nil, zero value otherwise.

### GetShopIDOk

`func (o *ProductBasic) GetShopIDOk() (*int32, bool)`

GetShopIDOk returns a tuple with the ShopID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopID

`func (o *ProductBasic) SetShopID(v int32)`

SetShopID sets ShopID field to given value.

### HasShopID

`func (o *ProductBasic) HasShopID() bool`

HasShopID returns a boolean if a field has been set.

### GetStatus

`func (o *ProductBasic) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProductBasic) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProductBasic) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProductBasic) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetName

`func (o *ProductBasic) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductBasic) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductBasic) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProductBasic) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCondition

`func (o *ProductBasic) GetCondition() int32`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *ProductBasic) GetConditionOk() (*int32, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *ProductBasic) SetCondition(v int32)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *ProductBasic) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetChildCategoryID

`func (o *ProductBasic) GetChildCategoryID() int32`

GetChildCategoryID returns the ChildCategoryID field if non-nil, zero value otherwise.

### GetChildCategoryIDOk

`func (o *ProductBasic) GetChildCategoryIDOk() (*int32, bool)`

GetChildCategoryIDOk returns a tuple with the ChildCategoryID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildCategoryID

`func (o *ProductBasic) SetChildCategoryID(v int32)`

SetChildCategoryID sets ChildCategoryID field to given value.

### HasChildCategoryID

`func (o *ProductBasic) HasChildCategoryID() bool`

HasChildCategoryID returns a boolean if a field has been set.

### GetShortDesc

`func (o *ProductBasic) GetShortDesc() string`

GetShortDesc returns the ShortDesc field if non-nil, zero value otherwise.

### GetShortDescOk

`func (o *ProductBasic) GetShortDescOk() (*string, bool)`

GetShortDescOk returns a tuple with the ShortDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDesc

`func (o *ProductBasic) SetShortDesc(v string)`

SetShortDesc sets ShortDesc field to given value.

### HasShortDesc

`func (o *ProductBasic) HasShortDesc() bool`

HasShortDesc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


