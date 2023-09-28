# RequestPartialOrderFulfillmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PofDetail** | Pointer to [**[]RequestPartialOrderFulfillmentRequestPofDetailInner**](RequestPartialOrderFulfillmentRequestPofDetailInner.md) | Array of POF detail info | [optional] 

## Methods

### NewRequestPartialOrderFulfillmentRequest

`func NewRequestPartialOrderFulfillmentRequest() *RequestPartialOrderFulfillmentRequest`

NewRequestPartialOrderFulfillmentRequest instantiates a new RequestPartialOrderFulfillmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestPartialOrderFulfillmentRequestWithDefaults

`func NewRequestPartialOrderFulfillmentRequestWithDefaults() *RequestPartialOrderFulfillmentRequest`

NewRequestPartialOrderFulfillmentRequestWithDefaults instantiates a new RequestPartialOrderFulfillmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPofDetail

`func (o *RequestPartialOrderFulfillmentRequest) GetPofDetail() []RequestPartialOrderFulfillmentRequestPofDetailInner`

GetPofDetail returns the PofDetail field if non-nil, zero value otherwise.

### GetPofDetailOk

`func (o *RequestPartialOrderFulfillmentRequest) GetPofDetailOk() (*[]RequestPartialOrderFulfillmentRequestPofDetailInner, bool)`

GetPofDetailOk returns a tuple with the PofDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPofDetail

`func (o *RequestPartialOrderFulfillmentRequest) SetPofDetail(v []RequestPartialOrderFulfillmentRequestPofDetailInner)`

SetPofDetail sets PofDetail field to given value.

### HasPofDetail

`func (o *RequestPartialOrderFulfillmentRequest) HasPofDetail() bool`

HasPofDetail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


