# GetResolutionTicket200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsSuccess** | Pointer to **bool** |  | [optional] 
**Startdate** | Pointer to **string** | format: YYYY-MM-DD  | [optional] 
**Enddate** | Pointer to **string** | format: YYYY-MM-DD  | [optional] 
**Shops** | Pointer to [**[]GetResolutionTicket200ResponseDataShopsInner**](GetResolutionTicket200ResponseDataShopsInner.md) |  | [optional] 

## Methods

### NewGetResolutionTicket200ResponseData

`func NewGetResolutionTicket200ResponseData() *GetResolutionTicket200ResponseData`

NewGetResolutionTicket200ResponseData instantiates a new GetResolutionTicket200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetResolutionTicket200ResponseDataWithDefaults

`func NewGetResolutionTicket200ResponseDataWithDefaults() *GetResolutionTicket200ResponseData`

NewGetResolutionTicket200ResponseDataWithDefaults instantiates a new GetResolutionTicket200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsSuccess

`func (o *GetResolutionTicket200ResponseData) GetIsSuccess() bool`

GetIsSuccess returns the IsSuccess field if non-nil, zero value otherwise.

### GetIsSuccessOk

`func (o *GetResolutionTicket200ResponseData) GetIsSuccessOk() (*bool, bool)`

GetIsSuccessOk returns a tuple with the IsSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuccess

`func (o *GetResolutionTicket200ResponseData) SetIsSuccess(v bool)`

SetIsSuccess sets IsSuccess field to given value.

### HasIsSuccess

`func (o *GetResolutionTicket200ResponseData) HasIsSuccess() bool`

HasIsSuccess returns a boolean if a field has been set.

### GetStartdate

`func (o *GetResolutionTicket200ResponseData) GetStartdate() string`

GetStartdate returns the Startdate field if non-nil, zero value otherwise.

### GetStartdateOk

`func (o *GetResolutionTicket200ResponseData) GetStartdateOk() (*string, bool)`

GetStartdateOk returns a tuple with the Startdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartdate

`func (o *GetResolutionTicket200ResponseData) SetStartdate(v string)`

SetStartdate sets Startdate field to given value.

### HasStartdate

`func (o *GetResolutionTicket200ResponseData) HasStartdate() bool`

HasStartdate returns a boolean if a field has been set.

### GetEnddate

`func (o *GetResolutionTicket200ResponseData) GetEnddate() string`

GetEnddate returns the Enddate field if non-nil, zero value otherwise.

### GetEnddateOk

`func (o *GetResolutionTicket200ResponseData) GetEnddateOk() (*string, bool)`

GetEnddateOk returns a tuple with the Enddate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnddate

`func (o *GetResolutionTicket200ResponseData) SetEnddate(v string)`

SetEnddate sets Enddate field to given value.

### HasEnddate

`func (o *GetResolutionTicket200ResponseData) HasEnddate() bool`

HasEnddate returns a boolean if a field has been set.

### GetShops

`func (o *GetResolutionTicket200ResponseData) GetShops() []GetResolutionTicket200ResponseDataShopsInner`

GetShops returns the Shops field if non-nil, zero value otherwise.

### GetShopsOk

`func (o *GetResolutionTicket200ResponseData) GetShopsOk() (*[]GetResolutionTicket200ResponseDataShopsInner, bool)`

GetShopsOk returns a tuple with the Shops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShops

`func (o *GetResolutionTicket200ResponseData) SetShops(v []GetResolutionTicket200ResponseDataShopsInner)`

SetShops sets Shops field to given value.

### HasShops

`func (o *GetResolutionTicket200ResponseData) HasShops() bool`

HasShops returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


