# UpsertProductCategoriesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailureCategory** | Pointer to [**[]ProductCategory**](ProductCategory.md) |  | [optional] 
**TotalFailureCategory** | Pointer to **int32** |  | [optional] 
**TotalSuccessCategory** | Pointer to **int32** |  | [optional] 
**TotalCategory** | Pointer to **int32** |  | [optional] 

## Methods

### NewUpsertProductCategoriesResponse

`func NewUpsertProductCategoriesResponse() *UpsertProductCategoriesResponse`

NewUpsertProductCategoriesResponse instantiates a new UpsertProductCategoriesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpsertProductCategoriesResponseWithDefaults

`func NewUpsertProductCategoriesResponseWithDefaults() *UpsertProductCategoriesResponse`

NewUpsertProductCategoriesResponseWithDefaults instantiates a new UpsertProductCategoriesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailureCategory

`func (o *UpsertProductCategoriesResponse) GetFailureCategory() []ProductCategory`

GetFailureCategory returns the FailureCategory field if non-nil, zero value otherwise.

### GetFailureCategoryOk

`func (o *UpsertProductCategoriesResponse) GetFailureCategoryOk() (*[]ProductCategory, bool)`

GetFailureCategoryOk returns a tuple with the FailureCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCategory

`func (o *UpsertProductCategoriesResponse) SetFailureCategory(v []ProductCategory)`

SetFailureCategory sets FailureCategory field to given value.

### HasFailureCategory

`func (o *UpsertProductCategoriesResponse) HasFailureCategory() bool`

HasFailureCategory returns a boolean if a field has been set.

### SetFailureCategoryNil

`func (o *UpsertProductCategoriesResponse) SetFailureCategoryNil(b bool)`

 SetFailureCategoryNil sets the value for FailureCategory to be an explicit nil

### UnsetFailureCategory
`func (o *UpsertProductCategoriesResponse) UnsetFailureCategory()`

UnsetFailureCategory ensures that no value is present for FailureCategory, not even an explicit nil
### GetTotalFailureCategory

`func (o *UpsertProductCategoriesResponse) GetTotalFailureCategory() int32`

GetTotalFailureCategory returns the TotalFailureCategory field if non-nil, zero value otherwise.

### GetTotalFailureCategoryOk

`func (o *UpsertProductCategoriesResponse) GetTotalFailureCategoryOk() (*int32, bool)`

GetTotalFailureCategoryOk returns a tuple with the TotalFailureCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFailureCategory

`func (o *UpsertProductCategoriesResponse) SetTotalFailureCategory(v int32)`

SetTotalFailureCategory sets TotalFailureCategory field to given value.

### HasTotalFailureCategory

`func (o *UpsertProductCategoriesResponse) HasTotalFailureCategory() bool`

HasTotalFailureCategory returns a boolean if a field has been set.

### GetTotalSuccessCategory

`func (o *UpsertProductCategoriesResponse) GetTotalSuccessCategory() int32`

GetTotalSuccessCategory returns the TotalSuccessCategory field if non-nil, zero value otherwise.

### GetTotalSuccessCategoryOk

`func (o *UpsertProductCategoriesResponse) GetTotalSuccessCategoryOk() (*int32, bool)`

GetTotalSuccessCategoryOk returns a tuple with the TotalSuccessCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSuccessCategory

`func (o *UpsertProductCategoriesResponse) SetTotalSuccessCategory(v int32)`

SetTotalSuccessCategory sets TotalSuccessCategory field to given value.

### HasTotalSuccessCategory

`func (o *UpsertProductCategoriesResponse) HasTotalSuccessCategory() bool`

HasTotalSuccessCategory returns a boolean if a field has been set.

### GetTotalCategory

`func (o *UpsertProductCategoriesResponse) GetTotalCategory() int32`

GetTotalCategory returns the TotalCategory field if non-nil, zero value otherwise.

### GetTotalCategoryOk

`func (o *UpsertProductCategoriesResponse) GetTotalCategoryOk() (*int32, bool)`

GetTotalCategoryOk returns a tuple with the TotalCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCategory

`func (o *UpsertProductCategoriesResponse) SetTotalCategory(v int32)`

SetTotalCategory sets TotalCategory field to given value.

### HasTotalCategory

`func (o *UpsertProductCategoriesResponse) HasTotalCategory() bool`

HasTotalCategory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


