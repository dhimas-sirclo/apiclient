# CreateProductV3DefaultResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Header** | Pointer to [**ResponseHeader**](ResponseHeader.md) |  | [optional] 
**Data** | Pointer to [**CreateProductV3DefaultResponseData**](CreateProductV3DefaultResponseData.md) |  | [optional] 

## Methods

### NewCreateProductV3DefaultResponse

`func NewCreateProductV3DefaultResponse() *CreateProductV3DefaultResponse`

NewCreateProductV3DefaultResponse instantiates a new CreateProductV3DefaultResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProductV3DefaultResponseWithDefaults

`func NewCreateProductV3DefaultResponseWithDefaults() *CreateProductV3DefaultResponse`

NewCreateProductV3DefaultResponseWithDefaults instantiates a new CreateProductV3DefaultResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeader

`func (o *CreateProductV3DefaultResponse) GetHeader() ResponseHeader`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *CreateProductV3DefaultResponse) GetHeaderOk() (*ResponseHeader, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *CreateProductV3DefaultResponse) SetHeader(v ResponseHeader)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *CreateProductV3DefaultResponse) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetData

`func (o *CreateProductV3DefaultResponse) GetData() CreateProductV3DefaultResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateProductV3DefaultResponse) GetDataOk() (*CreateProductV3DefaultResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateProductV3DefaultResponse) SetData(v CreateProductV3DefaultResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *CreateProductV3DefaultResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


