# GetProductDiscussion200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasNext** | Pointer to **bool** |  | [optional] 
**TotalQuestion** | Pointer to **int64** | Total Question Count | [optional] 
**Question** | Pointer to [**[]GetProductDiscussion200ResponseDataQuestionInner**](GetProductDiscussion200ResponseDataQuestionInner.md) |  | [optional] 
**ProductID** | Pointer to **int64** | Product Unique Identifier | [optional] 
**ShopID** | Pointer to **int64** | Shop Unique Identifier | [optional] 
**ShopURL** | Pointer to **string** | Shop URL | [optional] 

## Methods

### NewGetProductDiscussion200ResponseData

`func NewGetProductDiscussion200ResponseData() *GetProductDiscussion200ResponseData`

NewGetProductDiscussion200ResponseData instantiates a new GetProductDiscussion200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetProductDiscussion200ResponseDataWithDefaults

`func NewGetProductDiscussion200ResponseDataWithDefaults() *GetProductDiscussion200ResponseData`

NewGetProductDiscussion200ResponseDataWithDefaults instantiates a new GetProductDiscussion200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasNext

`func (o *GetProductDiscussion200ResponseData) GetHasNext() bool`

GetHasNext returns the HasNext field if non-nil, zero value otherwise.

### GetHasNextOk

`func (o *GetProductDiscussion200ResponseData) GetHasNextOk() (*bool, bool)`

GetHasNextOk returns a tuple with the HasNext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNext

`func (o *GetProductDiscussion200ResponseData) SetHasNext(v bool)`

SetHasNext sets HasNext field to given value.

### HasHasNext

`func (o *GetProductDiscussion200ResponseData) HasHasNext() bool`

HasHasNext returns a boolean if a field has been set.

### GetTotalQuestion

`func (o *GetProductDiscussion200ResponseData) GetTotalQuestion() int64`

GetTotalQuestion returns the TotalQuestion field if non-nil, zero value otherwise.

### GetTotalQuestionOk

`func (o *GetProductDiscussion200ResponseData) GetTotalQuestionOk() (*int64, bool)`

GetTotalQuestionOk returns a tuple with the TotalQuestion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalQuestion

`func (o *GetProductDiscussion200ResponseData) SetTotalQuestion(v int64)`

SetTotalQuestion sets TotalQuestion field to given value.

### HasTotalQuestion

`func (o *GetProductDiscussion200ResponseData) HasTotalQuestion() bool`

HasTotalQuestion returns a boolean if a field has been set.

### GetQuestion

`func (o *GetProductDiscussion200ResponseData) GetQuestion() []GetProductDiscussion200ResponseDataQuestionInner`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *GetProductDiscussion200ResponseData) GetQuestionOk() (*[]GetProductDiscussion200ResponseDataQuestionInner, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *GetProductDiscussion200ResponseData) SetQuestion(v []GetProductDiscussion200ResponseDataQuestionInner)`

SetQuestion sets Question field to given value.

### HasQuestion

`func (o *GetProductDiscussion200ResponseData) HasQuestion() bool`

HasQuestion returns a boolean if a field has been set.

### GetProductID

`func (o *GetProductDiscussion200ResponseData) GetProductID() int64`

GetProductID returns the ProductID field if non-nil, zero value otherwise.

### GetProductIDOk

`func (o *GetProductDiscussion200ResponseData) GetProductIDOk() (*int64, bool)`

GetProductIDOk returns a tuple with the ProductID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductID

`func (o *GetProductDiscussion200ResponseData) SetProductID(v int64)`

SetProductID sets ProductID field to given value.

### HasProductID

`func (o *GetProductDiscussion200ResponseData) HasProductID() bool`

HasProductID returns a boolean if a field has been set.

### GetShopID

`func (o *GetProductDiscussion200ResponseData) GetShopID() int64`

GetShopID returns the ShopID field if non-nil, zero value otherwise.

### GetShopIDOk

`func (o *GetProductDiscussion200ResponseData) GetShopIDOk() (*int64, bool)`

GetShopIDOk returns a tuple with the ShopID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopID

`func (o *GetProductDiscussion200ResponseData) SetShopID(v int64)`

SetShopID sets ShopID field to given value.

### HasShopID

`func (o *GetProductDiscussion200ResponseData) HasShopID() bool`

HasShopID returns a boolean if a field has been set.

### GetShopURL

`func (o *GetProductDiscussion200ResponseData) GetShopURL() string`

GetShopURL returns the ShopURL field if non-nil, zero value otherwise.

### GetShopURLOk

`func (o *GetProductDiscussion200ResponseData) GetShopURLOk() (*string, bool)`

GetShopURLOk returns a tuple with the ShopURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopURL

`func (o *GetProductDiscussion200ResponseData) SetShopURL(v string)`

SetShopURL sets ShopURL field to given value.

### HasShopURL

`func (o *GetProductDiscussion200ResponseData) HasShopURL() bool`

HasShopURL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


