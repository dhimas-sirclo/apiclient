# DecreaseStockResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Header** | Pointer to [**ResponseHeader**](ResponseHeader.md) |  | [optional] 
**Data** | Pointer to [**IncreaseStockResponseData**](IncreaseStockResponseData.md) |  | [optional] 

## Methods

### NewDecreaseStockResponse

`func NewDecreaseStockResponse() *DecreaseStockResponse`

NewDecreaseStockResponse instantiates a new DecreaseStockResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDecreaseStockResponseWithDefaults

`func NewDecreaseStockResponseWithDefaults() *DecreaseStockResponse`

NewDecreaseStockResponseWithDefaults instantiates a new DecreaseStockResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeader

`func (o *DecreaseStockResponse) GetHeader() ResponseHeader`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *DecreaseStockResponse) GetHeaderOk() (*ResponseHeader, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *DecreaseStockResponse) SetHeader(v ResponseHeader)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *DecreaseStockResponse) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetData

`func (o *DecreaseStockResponse) GetData() IncreaseStockResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DecreaseStockResponse) GetDataOk() (*IncreaseStockResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DecreaseStockResponse) SetData(v IncreaseStockResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *DecreaseStockResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


