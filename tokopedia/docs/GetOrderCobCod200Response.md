# GetOrderCobCod200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Header** | Pointer to [**ResponseHeader**](ResponseHeader.md) |  | [optional] 
**Data** | Pointer to [**GetOrderCobCod200ResponseData**](GetOrderCobCod200ResponseData.md) |  | [optional] 
**Status** | Pointer to **int64** | Response Status | [optional] 

## Methods

### NewGetOrderCobCod200Response

`func NewGetOrderCobCod200Response() *GetOrderCobCod200Response`

NewGetOrderCobCod200Response instantiates a new GetOrderCobCod200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrderCobCod200ResponseWithDefaults

`func NewGetOrderCobCod200ResponseWithDefaults() *GetOrderCobCod200Response`

NewGetOrderCobCod200ResponseWithDefaults instantiates a new GetOrderCobCod200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeader

`func (o *GetOrderCobCod200Response) GetHeader() ResponseHeader`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *GetOrderCobCod200Response) GetHeaderOk() (*ResponseHeader, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *GetOrderCobCod200Response) SetHeader(v ResponseHeader)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *GetOrderCobCod200Response) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetData

`func (o *GetOrderCobCod200Response) GetData() GetOrderCobCod200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetOrderCobCod200Response) GetDataOk() (*GetOrderCobCod200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetOrderCobCod200Response) SetData(v GetOrderCobCod200ResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *GetOrderCobCod200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetStatus

`func (o *GetOrderCobCod200Response) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetOrderCobCod200Response) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetOrderCobCod200Response) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetOrderCobCod200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


