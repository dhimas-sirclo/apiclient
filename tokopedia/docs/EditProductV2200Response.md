# EditProductV2200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Header** | Pointer to [**ResponseHeader**](ResponseHeader.md) |  | [optional] 
**Data** | Pointer to [**EditProductV2200ResponseData**](EditProductV2200ResponseData.md) |  | [optional] 

## Methods

### NewEditProductV2200Response

`func NewEditProductV2200Response() *EditProductV2200Response`

NewEditProductV2200Response instantiates a new EditProductV2200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditProductV2200ResponseWithDefaults

`func NewEditProductV2200ResponseWithDefaults() *EditProductV2200Response`

NewEditProductV2200ResponseWithDefaults instantiates a new EditProductV2200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeader

`func (o *EditProductV2200Response) GetHeader() ResponseHeader`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *EditProductV2200Response) GetHeaderOk() (*ResponseHeader, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *EditProductV2200Response) SetHeader(v ResponseHeader)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *EditProductV2200Response) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetData

`func (o *EditProductV2200Response) GetData() EditProductV2200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EditProductV2200Response) GetDataOk() (*EditProductV2200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EditProductV2200Response) SetData(v EditProductV2200ResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *EditProductV2200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


