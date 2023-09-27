# GetShowcase200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Showcase** | Pointer to [**[]GetShowcase200ResponseDataShowcaseInner**](GetShowcase200ResponseDataShowcaseInner.md) |  | [optional] 
**ShowcaseGroup** | Pointer to [**[]GetShowcase200ResponseDataShowcaseGroupInner**](GetShowcase200ResponseDataShowcaseGroupInner.md) |  | [optional] 
**UseAce** | Pointer to **bool** |  | [optional] 
**ShowcaseWithoutAce** | Pointer to **[]int64** |  | [optional] 
**PrevLink** | Pointer to **string** |  | [optional] 
**NextLink** | Pointer to **string** |  | [optional] 

## Methods

### NewGetShowcase200ResponseData

`func NewGetShowcase200ResponseData() *GetShowcase200ResponseData`

NewGetShowcase200ResponseData instantiates a new GetShowcase200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetShowcase200ResponseDataWithDefaults

`func NewGetShowcase200ResponseDataWithDefaults() *GetShowcase200ResponseData`

NewGetShowcase200ResponseDataWithDefaults instantiates a new GetShowcase200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShowcase

`func (o *GetShowcase200ResponseData) GetShowcase() []GetShowcase200ResponseDataShowcaseInner`

GetShowcase returns the Showcase field if non-nil, zero value otherwise.

### GetShowcaseOk

`func (o *GetShowcase200ResponseData) GetShowcaseOk() (*[]GetShowcase200ResponseDataShowcaseInner, bool)`

GetShowcaseOk returns a tuple with the Showcase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowcase

`func (o *GetShowcase200ResponseData) SetShowcase(v []GetShowcase200ResponseDataShowcaseInner)`

SetShowcase sets Showcase field to given value.

### HasShowcase

`func (o *GetShowcase200ResponseData) HasShowcase() bool`

HasShowcase returns a boolean if a field has been set.

### GetShowcaseGroup

`func (o *GetShowcase200ResponseData) GetShowcaseGroup() []GetShowcase200ResponseDataShowcaseGroupInner`

GetShowcaseGroup returns the ShowcaseGroup field if non-nil, zero value otherwise.

### GetShowcaseGroupOk

`func (o *GetShowcase200ResponseData) GetShowcaseGroupOk() (*[]GetShowcase200ResponseDataShowcaseGroupInner, bool)`

GetShowcaseGroupOk returns a tuple with the ShowcaseGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowcaseGroup

`func (o *GetShowcase200ResponseData) SetShowcaseGroup(v []GetShowcase200ResponseDataShowcaseGroupInner)`

SetShowcaseGroup sets ShowcaseGroup field to given value.

### HasShowcaseGroup

`func (o *GetShowcase200ResponseData) HasShowcaseGroup() bool`

HasShowcaseGroup returns a boolean if a field has been set.

### GetUseAce

`func (o *GetShowcase200ResponseData) GetUseAce() bool`

GetUseAce returns the UseAce field if non-nil, zero value otherwise.

### GetUseAceOk

`func (o *GetShowcase200ResponseData) GetUseAceOk() (*bool, bool)`

GetUseAceOk returns a tuple with the UseAce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAce

`func (o *GetShowcase200ResponseData) SetUseAce(v bool)`

SetUseAce sets UseAce field to given value.

### HasUseAce

`func (o *GetShowcase200ResponseData) HasUseAce() bool`

HasUseAce returns a boolean if a field has been set.

### GetShowcaseWithoutAce

`func (o *GetShowcase200ResponseData) GetShowcaseWithoutAce() []int64`

GetShowcaseWithoutAce returns the ShowcaseWithoutAce field if non-nil, zero value otherwise.

### GetShowcaseWithoutAceOk

`func (o *GetShowcase200ResponseData) GetShowcaseWithoutAceOk() (*[]int64, bool)`

GetShowcaseWithoutAceOk returns a tuple with the ShowcaseWithoutAce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowcaseWithoutAce

`func (o *GetShowcase200ResponseData) SetShowcaseWithoutAce(v []int64)`

SetShowcaseWithoutAce sets ShowcaseWithoutAce field to given value.

### HasShowcaseWithoutAce

`func (o *GetShowcase200ResponseData) HasShowcaseWithoutAce() bool`

HasShowcaseWithoutAce returns a boolean if a field has been set.

### GetPrevLink

`func (o *GetShowcase200ResponseData) GetPrevLink() string`

GetPrevLink returns the PrevLink field if non-nil, zero value otherwise.

### GetPrevLinkOk

`func (o *GetShowcase200ResponseData) GetPrevLinkOk() (*string, bool)`

GetPrevLinkOk returns a tuple with the PrevLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevLink

`func (o *GetShowcase200ResponseData) SetPrevLink(v string)`

SetPrevLink sets PrevLink field to given value.

### HasPrevLink

`func (o *GetShowcase200ResponseData) HasPrevLink() bool`

HasPrevLink returns a boolean if a field has been set.

### GetNextLink

`func (o *GetShowcase200ResponseData) GetNextLink() string`

GetNextLink returns the NextLink field if non-nil, zero value otherwise.

### GetNextLinkOk

`func (o *GetShowcase200ResponseData) GetNextLinkOk() (*string, bool)`

GetNextLinkOk returns a tuple with the NextLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextLink

`func (o *GetShowcase200ResponseData) SetNextLink(v string)`

SetNextLink sets NextLink field to given value.

### HasNextLink

`func (o *GetShowcase200ResponseData) HasNextLink() bool`

HasNextLink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


