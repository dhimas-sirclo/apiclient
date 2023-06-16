# GetProductInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Header** | Pointer to [**Header**](Header.md) |  | [optional] 
**Data** | Pointer to [**[]Product**](Product.md) |  | [optional] 

## Methods

### NewGetProductInfoResponse

`func NewGetProductInfoResponse() *GetProductInfoResponse`

NewGetProductInfoResponse instantiates a new GetProductInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetProductInfoResponseWithDefaults

`func NewGetProductInfoResponseWithDefaults() *GetProductInfoResponse`

NewGetProductInfoResponseWithDefaults instantiates a new GetProductInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeader

`func (o *GetProductInfoResponse) GetHeader() Header`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *GetProductInfoResponse) GetHeaderOk() (*Header, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *GetProductInfoResponse) SetHeader(v Header)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *GetProductInfoResponse) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetData

`func (o *GetProductInfoResponse) GetData() []Product`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetProductInfoResponse) GetDataOk() (*[]Product, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetProductInfoResponse) SetData(v []Product)`

SetData sets Data field to given value.

### HasData

`func (o *GetProductInfoResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


