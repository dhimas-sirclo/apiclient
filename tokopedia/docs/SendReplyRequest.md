# SendReplyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShopId** | **int64** |  | 
**Message** | **string** |  | 
**AttachmentType** | **int64** |  | 
**Payload** | [**SendReplyRequestOneOf1Payload**](SendReplyRequestOneOf1Payload.md) |  | 

## Methods

### NewSendReplyRequest

`func NewSendReplyRequest(shopId int64, message string, attachmentType int64, payload SendReplyRequestOneOf1Payload, ) *SendReplyRequest`

NewSendReplyRequest instantiates a new SendReplyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendReplyRequestWithDefaults

`func NewSendReplyRequestWithDefaults() *SendReplyRequest`

NewSendReplyRequestWithDefaults instantiates a new SendReplyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShopId

`func (o *SendReplyRequest) GetShopId() int64`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *SendReplyRequest) GetShopIdOk() (*int64, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *SendReplyRequest) SetShopId(v int64)`

SetShopId sets ShopId field to given value.


### GetMessage

`func (o *SendReplyRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SendReplyRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SendReplyRequest) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetAttachmentType

`func (o *SendReplyRequest) GetAttachmentType() int64`

GetAttachmentType returns the AttachmentType field if non-nil, zero value otherwise.

### GetAttachmentTypeOk

`func (o *SendReplyRequest) GetAttachmentTypeOk() (*int64, bool)`

GetAttachmentTypeOk returns a tuple with the AttachmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentType

`func (o *SendReplyRequest) SetAttachmentType(v int64)`

SetAttachmentType sets AttachmentType field to given value.


### GetPayload

`func (o *SendReplyRequest) GetPayload() SendReplyRequestOneOf1Payload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *SendReplyRequest) GetPayloadOk() (*SendReplyRequestOneOf1Payload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *SendReplyRequest) SetPayload(v SendReplyRequestOneOf1Payload)`

SetPayload sets Payload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


