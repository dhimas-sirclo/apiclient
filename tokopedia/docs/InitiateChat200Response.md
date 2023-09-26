# InitiateChat200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Header** | Pointer to [**GetListMessageDefaultResponseHeader**](GetListMessageDefaultResponseHeader.md) |  | [optional] 
**Data** | Pointer to [**NullableInitiateChat200ResponseData**](InitiateChat200ResponseData.md) |  | [optional] 

## Methods

### NewInitiateChat200Response

`func NewInitiateChat200Response() *InitiateChat200Response`

NewInitiateChat200Response instantiates a new InitiateChat200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInitiateChat200ResponseWithDefaults

`func NewInitiateChat200ResponseWithDefaults() *InitiateChat200Response`

NewInitiateChat200ResponseWithDefaults instantiates a new InitiateChat200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeader

`func (o *InitiateChat200Response) GetHeader() GetListMessageDefaultResponseHeader`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *InitiateChat200Response) GetHeaderOk() (*GetListMessageDefaultResponseHeader, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *InitiateChat200Response) SetHeader(v GetListMessageDefaultResponseHeader)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *InitiateChat200Response) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetData

`func (o *InitiateChat200Response) GetData() InitiateChat200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InitiateChat200Response) GetDataOk() (*InitiateChat200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InitiateChat200Response) SetData(v InitiateChat200ResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *InitiateChat200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *InitiateChat200Response) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *InitiateChat200Response) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


