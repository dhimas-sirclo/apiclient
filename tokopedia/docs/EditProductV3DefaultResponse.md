# EditProductV3DefaultResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Header** | Pointer to [**ResponseHeader**](ResponseHeader.md) |  | [optional] 
**Data** | Pointer to [**EditProductV3DefaultResponseData**](EditProductV3DefaultResponseData.md) |  | [optional] 

## Methods

### NewEditProductV3DefaultResponse

`func NewEditProductV3DefaultResponse() *EditProductV3DefaultResponse`

NewEditProductV3DefaultResponse instantiates a new EditProductV3DefaultResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditProductV3DefaultResponseWithDefaults

`func NewEditProductV3DefaultResponseWithDefaults() *EditProductV3DefaultResponse`

NewEditProductV3DefaultResponseWithDefaults instantiates a new EditProductV3DefaultResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeader

`func (o *EditProductV3DefaultResponse) GetHeader() ResponseHeader`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *EditProductV3DefaultResponse) GetHeaderOk() (*ResponseHeader, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *EditProductV3DefaultResponse) SetHeader(v ResponseHeader)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *EditProductV3DefaultResponse) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetData

`func (o *EditProductV3DefaultResponse) GetData() EditProductV3DefaultResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EditProductV3DefaultResponse) GetDataOk() (*EditProductV3DefaultResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EditProductV3DefaultResponse) SetData(v EditProductV3DefaultResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *EditProductV3DefaultResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


