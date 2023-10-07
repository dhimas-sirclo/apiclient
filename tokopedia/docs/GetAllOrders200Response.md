# GetAllOrders200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Header** | Pointer to [**ResponseHeader**](ResponseHeader.md) |  | [optional] 
**Data** | Pointer to [**[]GetAllOrders200ResponseDataInner**](GetAllOrders200ResponseDataInner.md) |  | [optional] 

## Methods

### NewGetAllOrders200Response

`func NewGetAllOrders200Response() *GetAllOrders200Response`

NewGetAllOrders200Response instantiates a new GetAllOrders200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllOrders200ResponseWithDefaults

`func NewGetAllOrders200ResponseWithDefaults() *GetAllOrders200Response`

NewGetAllOrders200ResponseWithDefaults instantiates a new GetAllOrders200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeader

`func (o *GetAllOrders200Response) GetHeader() ResponseHeader`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *GetAllOrders200Response) GetHeaderOk() (*ResponseHeader, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *GetAllOrders200Response) SetHeader(v ResponseHeader)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *GetAllOrders200Response) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetData

`func (o *GetAllOrders200Response) GetData() []GetAllOrders200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetAllOrders200Response) GetDataOk() (*[]GetAllOrders200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetAllOrders200Response) SetData(v []GetAllOrders200ResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *GetAllOrders200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


