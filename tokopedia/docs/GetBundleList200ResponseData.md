# GetBundleList200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BundleListInfo** | Pointer to [**[]GetBundleList200ResponseDataBundleListInfoInner**](GetBundleList200ResponseDataBundleListInfoInner.md) |  | [optional] 
**IsLastPage** | Pointer to **bool** |  | [optional] 
**LastGroupId** | Pointer to **int64** |  | [optional] 

## Methods

### NewGetBundleList200ResponseData

`func NewGetBundleList200ResponseData() *GetBundleList200ResponseData`

NewGetBundleList200ResponseData instantiates a new GetBundleList200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBundleList200ResponseDataWithDefaults

`func NewGetBundleList200ResponseDataWithDefaults() *GetBundleList200ResponseData`

NewGetBundleList200ResponseDataWithDefaults instantiates a new GetBundleList200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBundleListInfo

`func (o *GetBundleList200ResponseData) GetBundleListInfo() []GetBundleList200ResponseDataBundleListInfoInner`

GetBundleListInfo returns the BundleListInfo field if non-nil, zero value otherwise.

### GetBundleListInfoOk

`func (o *GetBundleList200ResponseData) GetBundleListInfoOk() (*[]GetBundleList200ResponseDataBundleListInfoInner, bool)`

GetBundleListInfoOk returns a tuple with the BundleListInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleListInfo

`func (o *GetBundleList200ResponseData) SetBundleListInfo(v []GetBundleList200ResponseDataBundleListInfoInner)`

SetBundleListInfo sets BundleListInfo field to given value.

### HasBundleListInfo

`func (o *GetBundleList200ResponseData) HasBundleListInfo() bool`

HasBundleListInfo returns a boolean if a field has been set.

### GetIsLastPage

`func (o *GetBundleList200ResponseData) GetIsLastPage() bool`

GetIsLastPage returns the IsLastPage field if non-nil, zero value otherwise.

### GetIsLastPageOk

`func (o *GetBundleList200ResponseData) GetIsLastPageOk() (*bool, bool)`

GetIsLastPageOk returns a tuple with the IsLastPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLastPage

`func (o *GetBundleList200ResponseData) SetIsLastPage(v bool)`

SetIsLastPage sets IsLastPage field to given value.

### HasIsLastPage

`func (o *GetBundleList200ResponseData) HasIsLastPage() bool`

HasIsLastPage returns a boolean if a field has been set.

### GetLastGroupId

`func (o *GetBundleList200ResponseData) GetLastGroupId() int64`

GetLastGroupId returns the LastGroupId field if non-nil, zero value otherwise.

### GetLastGroupIdOk

`func (o *GetBundleList200ResponseData) GetLastGroupIdOk() (*int64, bool)`

GetLastGroupIdOk returns a tuple with the LastGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastGroupId

`func (o *GetBundleList200ResponseData) SetLastGroupId(v int64)`

SetLastGroupId sets LastGroupId field to given value.

### HasLastGroupId

`func (o *GetBundleList200ResponseData) HasLastGroupId() bool`

HasLastGroupId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


