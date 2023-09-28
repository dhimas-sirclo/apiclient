# GetOrderCobCod200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderData** | Pointer to [**[]GetOrderCobCod200ResponseDataOrderDataInner**](GetOrderCobCod200ResponseDataOrderDataInner.md) |  | [optional] 
**NextOrderId** | Pointer to **int64** | Next Order Unique Identifier | [optional] 
**FirstOrderId** | Pointer to **int64** | FirstOrder Unique Identifier | [optional] 

## Methods

### NewGetOrderCobCod200ResponseData

`func NewGetOrderCobCod200ResponseData() *GetOrderCobCod200ResponseData`

NewGetOrderCobCod200ResponseData instantiates a new GetOrderCobCod200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrderCobCod200ResponseDataWithDefaults

`func NewGetOrderCobCod200ResponseDataWithDefaults() *GetOrderCobCod200ResponseData`

NewGetOrderCobCod200ResponseDataWithDefaults instantiates a new GetOrderCobCod200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderData

`func (o *GetOrderCobCod200ResponseData) GetOrderData() []GetOrderCobCod200ResponseDataOrderDataInner`

GetOrderData returns the OrderData field if non-nil, zero value otherwise.

### GetOrderDataOk

`func (o *GetOrderCobCod200ResponseData) GetOrderDataOk() (*[]GetOrderCobCod200ResponseDataOrderDataInner, bool)`

GetOrderDataOk returns a tuple with the OrderData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderData

`func (o *GetOrderCobCod200ResponseData) SetOrderData(v []GetOrderCobCod200ResponseDataOrderDataInner)`

SetOrderData sets OrderData field to given value.

### HasOrderData

`func (o *GetOrderCobCod200ResponseData) HasOrderData() bool`

HasOrderData returns a boolean if a field has been set.

### GetNextOrderId

`func (o *GetOrderCobCod200ResponseData) GetNextOrderId() int64`

GetNextOrderId returns the NextOrderId field if non-nil, zero value otherwise.

### GetNextOrderIdOk

`func (o *GetOrderCobCod200ResponseData) GetNextOrderIdOk() (*int64, bool)`

GetNextOrderIdOk returns a tuple with the NextOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextOrderId

`func (o *GetOrderCobCod200ResponseData) SetNextOrderId(v int64)`

SetNextOrderId sets NextOrderId field to given value.

### HasNextOrderId

`func (o *GetOrderCobCod200ResponseData) HasNextOrderId() bool`

HasNextOrderId returns a boolean if a field has been set.

### GetFirstOrderId

`func (o *GetOrderCobCod200ResponseData) GetFirstOrderId() int64`

GetFirstOrderId returns the FirstOrderId field if non-nil, zero value otherwise.

### GetFirstOrderIdOk

`func (o *GetOrderCobCod200ResponseData) GetFirstOrderIdOk() (*int64, bool)`

GetFirstOrderIdOk returns a tuple with the FirstOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstOrderId

`func (o *GetOrderCobCod200ResponseData) SetFirstOrderId(v int64)`

SetFirstOrderId sets FirstOrderId field to given value.

### HasFirstOrderId

`func (o *GetOrderCobCod200ResponseData) HasFirstOrderId() bool`

HasFirstOrderId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


