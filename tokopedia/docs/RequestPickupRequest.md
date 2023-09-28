# RequestPickupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | **int64** |  | 
**ShopId** | **int64** |  | 

## Methods

### NewRequestPickupRequest

`func NewRequestPickupRequest(orderId int64, shopId int64, ) *RequestPickupRequest`

NewRequestPickupRequest instantiates a new RequestPickupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestPickupRequestWithDefaults

`func NewRequestPickupRequestWithDefaults() *RequestPickupRequest`

NewRequestPickupRequestWithDefaults instantiates a new RequestPickupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *RequestPickupRequest) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *RequestPickupRequest) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *RequestPickupRequest) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.


### GetShopId

`func (o *RequestPickupRequest) GetShopId() int64`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *RequestPickupRequest) GetShopIdOk() (*int64, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *RequestPickupRequest) SetShopId(v int64)`

SetShopId sets ShopId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


