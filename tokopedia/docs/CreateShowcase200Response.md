# CreateShowcase200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Header** | Pointer to [**ResponseHeader**](ResponseHeader.md) |  | [optional] 
**Data** | Pointer to [**CreateShowcase200ResponseData**](CreateShowcase200ResponseData.md) |  | [optional] 

## Methods

### NewCreateShowcase200Response

`func NewCreateShowcase200Response() *CreateShowcase200Response`

NewCreateShowcase200Response instantiates a new CreateShowcase200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateShowcase200ResponseWithDefaults

`func NewCreateShowcase200ResponseWithDefaults() *CreateShowcase200Response`

NewCreateShowcase200ResponseWithDefaults instantiates a new CreateShowcase200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeader

`func (o *CreateShowcase200Response) GetHeader() ResponseHeader`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *CreateShowcase200Response) GetHeaderOk() (*ResponseHeader, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *CreateShowcase200Response) SetHeader(v ResponseHeader)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *CreateShowcase200Response) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetData

`func (o *CreateShowcase200Response) GetData() CreateShowcase200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateShowcase200Response) GetDataOk() (*CreateShowcase200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateShowcase200Response) SetData(v CreateShowcase200ResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *CreateShowcase200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


