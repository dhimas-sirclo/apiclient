# GetBundleList200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Header** | Pointer to [**ResponseHeader**](ResponseHeader.md) |  | [optional] 
**Data** | Pointer to [**GetBundleList200ResponseData**](GetBundleList200ResponseData.md) |  | [optional] 

## Methods

### NewGetBundleList200Response

`func NewGetBundleList200Response() *GetBundleList200Response`

NewGetBundleList200Response instantiates a new GetBundleList200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBundleList200ResponseWithDefaults

`func NewGetBundleList200ResponseWithDefaults() *GetBundleList200Response`

NewGetBundleList200ResponseWithDefaults instantiates a new GetBundleList200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeader

`func (o *GetBundleList200Response) GetHeader() ResponseHeader`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *GetBundleList200Response) GetHeaderOk() (*ResponseHeader, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *GetBundleList200Response) SetHeader(v ResponseHeader)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *GetBundleList200Response) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetData

`func (o *GetBundleList200Response) GetData() GetBundleList200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetBundleList200Response) GetDataOk() (*GetBundleList200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetBundleList200Response) SetData(v GetBundleList200ResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *GetBundleList200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


