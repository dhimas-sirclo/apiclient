# CreateProductV3200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Header** | Pointer to [**ResponseHeader**](ResponseHeader.md) |  | [optional] 
**Data** | Pointer to [**CreateProductV3200ResponseData**](CreateProductV3200ResponseData.md) |  | [optional] 

## Methods

### NewCreateProductV3200Response

`func NewCreateProductV3200Response() *CreateProductV3200Response`

NewCreateProductV3200Response instantiates a new CreateProductV3200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProductV3200ResponseWithDefaults

`func NewCreateProductV3200ResponseWithDefaults() *CreateProductV3200Response`

NewCreateProductV3200ResponseWithDefaults instantiates a new CreateProductV3200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeader

`func (o *CreateProductV3200Response) GetHeader() ResponseHeader`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *CreateProductV3200Response) GetHeaderOk() (*ResponseHeader, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *CreateProductV3200Response) SetHeader(v ResponseHeader)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *CreateProductV3200Response) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetData

`func (o *CreateProductV3200Response) GetData() CreateProductV3200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateProductV3200Response) GetDataOk() (*CreateProductV3200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateProductV3200Response) SetData(v CreateProductV3200ResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *CreateProductV3200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


