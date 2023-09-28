# RequestPickup200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | Pointer to **int64** | Order Unique Identifier | [optional] 
**ShopId** | Pointer to **int64** | Shop Unique Identifier | [optional] 
**RequestTime** | Pointer to **string** | Request Timestamp (format: 2018-06-12 10:24:00)  | [optional] 
**Result** | Pointer to **string** | Result Message | [optional] 

## Methods

### NewRequestPickup200ResponseData

`func NewRequestPickup200ResponseData() *RequestPickup200ResponseData`

NewRequestPickup200ResponseData instantiates a new RequestPickup200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestPickup200ResponseDataWithDefaults

`func NewRequestPickup200ResponseDataWithDefaults() *RequestPickup200ResponseData`

NewRequestPickup200ResponseDataWithDefaults instantiates a new RequestPickup200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *RequestPickup200ResponseData) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *RequestPickup200ResponseData) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *RequestPickup200ResponseData) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *RequestPickup200ResponseData) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetShopId

`func (o *RequestPickup200ResponseData) GetShopId() int64`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *RequestPickup200ResponseData) GetShopIdOk() (*int64, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *RequestPickup200ResponseData) SetShopId(v int64)`

SetShopId sets ShopId field to given value.

### HasShopId

`func (o *RequestPickup200ResponseData) HasShopId() bool`

HasShopId returns a boolean if a field has been set.

### GetRequestTime

`func (o *RequestPickup200ResponseData) GetRequestTime() string`

GetRequestTime returns the RequestTime field if non-nil, zero value otherwise.

### GetRequestTimeOk

`func (o *RequestPickup200ResponseData) GetRequestTimeOk() (*string, bool)`

GetRequestTimeOk returns a tuple with the RequestTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTime

`func (o *RequestPickup200ResponseData) SetRequestTime(v string)`

SetRequestTime sets RequestTime field to given value.

### HasRequestTime

`func (o *RequestPickup200ResponseData) HasRequestTime() bool`

HasRequestTime returns a boolean if a field has been set.

### GetResult

`func (o *RequestPickup200ResponseData) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *RequestPickup200ResponseData) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *RequestPickup200ResponseData) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *RequestPickup200ResponseData) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


