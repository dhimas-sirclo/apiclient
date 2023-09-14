# WebhookOrderRequestCancellation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | Pointer to **int64** | Requested order unique identifier | [optional] 

## Methods

### NewWebhookOrderRequestCancellation

`func NewWebhookOrderRequestCancellation() *WebhookOrderRequestCancellation`

NewWebhookOrderRequestCancellation instantiates a new WebhookOrderRequestCancellation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookOrderRequestCancellationWithDefaults

`func NewWebhookOrderRequestCancellationWithDefaults() *WebhookOrderRequestCancellation`

NewWebhookOrderRequestCancellationWithDefaults instantiates a new WebhookOrderRequestCancellation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *WebhookOrderRequestCancellation) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *WebhookOrderRequestCancellation) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *WebhookOrderRequestCancellation) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *WebhookOrderRequestCancellation) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


