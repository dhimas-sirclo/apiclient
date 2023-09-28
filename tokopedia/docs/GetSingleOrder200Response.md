# GetSingleOrder200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Header** | Pointer to [**ResponseHeader**](ResponseHeader.md) |  | [optional] 
**Data** | Pointer to [**GetSingleOrder200ResponseData**](GetSingleOrder200ResponseData.md) |  | [optional] 

## Methods

### NewGetSingleOrder200Response

`func NewGetSingleOrder200Response() *GetSingleOrder200Response`

NewGetSingleOrder200Response instantiates a new GetSingleOrder200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSingleOrder200ResponseWithDefaults

`func NewGetSingleOrder200ResponseWithDefaults() *GetSingleOrder200Response`

NewGetSingleOrder200ResponseWithDefaults instantiates a new GetSingleOrder200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeader

`func (o *GetSingleOrder200Response) GetHeader() ResponseHeader`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *GetSingleOrder200Response) GetHeaderOk() (*ResponseHeader, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *GetSingleOrder200Response) SetHeader(v ResponseHeader)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *GetSingleOrder200Response) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetData

`func (o *GetSingleOrder200Response) GetData() GetSingleOrder200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetSingleOrder200Response) GetDataOk() (*GetSingleOrder200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetSingleOrder200Response) SetData(v GetSingleOrder200ResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *GetSingleOrder200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


