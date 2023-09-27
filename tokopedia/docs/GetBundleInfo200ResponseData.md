# GetBundleInfo200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BundleInfo** | Pointer to [**[]GetBundleInfo200ResponseDataBundleInfoInner**](GetBundleInfo200ResponseDataBundleInfoInner.md) |  | [optional] 
**Error** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewGetBundleInfo200ResponseData

`func NewGetBundleInfo200ResponseData() *GetBundleInfo200ResponseData`

NewGetBundleInfo200ResponseData instantiates a new GetBundleInfo200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBundleInfo200ResponseDataWithDefaults

`func NewGetBundleInfo200ResponseDataWithDefaults() *GetBundleInfo200ResponseData`

NewGetBundleInfo200ResponseDataWithDefaults instantiates a new GetBundleInfo200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBundleInfo

`func (o *GetBundleInfo200ResponseData) GetBundleInfo() []GetBundleInfo200ResponseDataBundleInfoInner`

GetBundleInfo returns the BundleInfo field if non-nil, zero value otherwise.

### GetBundleInfoOk

`func (o *GetBundleInfo200ResponseData) GetBundleInfoOk() (*[]GetBundleInfo200ResponseDataBundleInfoInner, bool)`

GetBundleInfoOk returns a tuple with the BundleInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleInfo

`func (o *GetBundleInfo200ResponseData) SetBundleInfo(v []GetBundleInfo200ResponseDataBundleInfoInner)`

SetBundleInfo sets BundleInfo field to given value.

### HasBundleInfo

`func (o *GetBundleInfo200ResponseData) HasBundleInfo() bool`

HasBundleInfo returns a boolean if a field has been set.

### GetError

`func (o *GetBundleInfo200ResponseData) GetError() map[string]interface{}`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetBundleInfo200ResponseData) GetErrorOk() (*map[string]interface{}, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetBundleInfo200ResponseData) SetError(v map[string]interface{})`

SetError sets Error field to given value.

### HasError

`func (o *GetBundleInfo200ResponseData) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *GetBundleInfo200ResponseData) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *GetBundleInfo200ResponseData) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


