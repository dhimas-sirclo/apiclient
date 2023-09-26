# RegisterWebhookRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FsId** | **int64** |  | 
**OrderNotificationUrl** | **string** |  | 
**OrderCancellationUrl** | **string** |  | 
**OrderStatusUrl** | **string** |  | 
**OrderRequestCancellationUrl** | **string** |  | 
**ChatNotificationUrl** | **string** |  | 
**ProductCreationUrl** | **string** |  | 
**ProductChangesUrl** | **string** |  | 
**CampaignNotificationUrl** | **string** |  | 
**WebhookSecret** | **string** |  | 

## Methods

### NewRegisterWebhookRequest

`func NewRegisterWebhookRequest(fsId int64, orderNotificationUrl string, orderCancellationUrl string, orderStatusUrl string, orderRequestCancellationUrl string, chatNotificationUrl string, productCreationUrl string, productChangesUrl string, campaignNotificationUrl string, webhookSecret string, ) *RegisterWebhookRequest`

NewRegisterWebhookRequest instantiates a new RegisterWebhookRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterWebhookRequestWithDefaults

`func NewRegisterWebhookRequestWithDefaults() *RegisterWebhookRequest`

NewRegisterWebhookRequestWithDefaults instantiates a new RegisterWebhookRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFsId

`func (o *RegisterWebhookRequest) GetFsId() int64`

GetFsId returns the FsId field if non-nil, zero value otherwise.

### GetFsIdOk

`func (o *RegisterWebhookRequest) GetFsIdOk() (*int64, bool)`

GetFsIdOk returns a tuple with the FsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsId

`func (o *RegisterWebhookRequest) SetFsId(v int64)`

SetFsId sets FsId field to given value.


### GetOrderNotificationUrl

`func (o *RegisterWebhookRequest) GetOrderNotificationUrl() string`

GetOrderNotificationUrl returns the OrderNotificationUrl field if non-nil, zero value otherwise.

### GetOrderNotificationUrlOk

`func (o *RegisterWebhookRequest) GetOrderNotificationUrlOk() (*string, bool)`

GetOrderNotificationUrlOk returns a tuple with the OrderNotificationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderNotificationUrl

`func (o *RegisterWebhookRequest) SetOrderNotificationUrl(v string)`

SetOrderNotificationUrl sets OrderNotificationUrl field to given value.


### GetOrderCancellationUrl

`func (o *RegisterWebhookRequest) GetOrderCancellationUrl() string`

GetOrderCancellationUrl returns the OrderCancellationUrl field if non-nil, zero value otherwise.

### GetOrderCancellationUrlOk

`func (o *RegisterWebhookRequest) GetOrderCancellationUrlOk() (*string, bool)`

GetOrderCancellationUrlOk returns a tuple with the OrderCancellationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderCancellationUrl

`func (o *RegisterWebhookRequest) SetOrderCancellationUrl(v string)`

SetOrderCancellationUrl sets OrderCancellationUrl field to given value.


### GetOrderStatusUrl

`func (o *RegisterWebhookRequest) GetOrderStatusUrl() string`

GetOrderStatusUrl returns the OrderStatusUrl field if non-nil, zero value otherwise.

### GetOrderStatusUrlOk

`func (o *RegisterWebhookRequest) GetOrderStatusUrlOk() (*string, bool)`

GetOrderStatusUrlOk returns a tuple with the OrderStatusUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderStatusUrl

`func (o *RegisterWebhookRequest) SetOrderStatusUrl(v string)`

SetOrderStatusUrl sets OrderStatusUrl field to given value.


### GetOrderRequestCancellationUrl

`func (o *RegisterWebhookRequest) GetOrderRequestCancellationUrl() string`

GetOrderRequestCancellationUrl returns the OrderRequestCancellationUrl field if non-nil, zero value otherwise.

### GetOrderRequestCancellationUrlOk

`func (o *RegisterWebhookRequest) GetOrderRequestCancellationUrlOk() (*string, bool)`

GetOrderRequestCancellationUrlOk returns a tuple with the OrderRequestCancellationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderRequestCancellationUrl

`func (o *RegisterWebhookRequest) SetOrderRequestCancellationUrl(v string)`

SetOrderRequestCancellationUrl sets OrderRequestCancellationUrl field to given value.


### GetChatNotificationUrl

`func (o *RegisterWebhookRequest) GetChatNotificationUrl() string`

GetChatNotificationUrl returns the ChatNotificationUrl field if non-nil, zero value otherwise.

### GetChatNotificationUrlOk

`func (o *RegisterWebhookRequest) GetChatNotificationUrlOk() (*string, bool)`

GetChatNotificationUrlOk returns a tuple with the ChatNotificationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatNotificationUrl

`func (o *RegisterWebhookRequest) SetChatNotificationUrl(v string)`

SetChatNotificationUrl sets ChatNotificationUrl field to given value.


### GetProductCreationUrl

`func (o *RegisterWebhookRequest) GetProductCreationUrl() string`

GetProductCreationUrl returns the ProductCreationUrl field if non-nil, zero value otherwise.

### GetProductCreationUrlOk

`func (o *RegisterWebhookRequest) GetProductCreationUrlOk() (*string, bool)`

GetProductCreationUrlOk returns a tuple with the ProductCreationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCreationUrl

`func (o *RegisterWebhookRequest) SetProductCreationUrl(v string)`

SetProductCreationUrl sets ProductCreationUrl field to given value.


### GetProductChangesUrl

`func (o *RegisterWebhookRequest) GetProductChangesUrl() string`

GetProductChangesUrl returns the ProductChangesUrl field if non-nil, zero value otherwise.

### GetProductChangesUrlOk

`func (o *RegisterWebhookRequest) GetProductChangesUrlOk() (*string, bool)`

GetProductChangesUrlOk returns a tuple with the ProductChangesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductChangesUrl

`func (o *RegisterWebhookRequest) SetProductChangesUrl(v string)`

SetProductChangesUrl sets ProductChangesUrl field to given value.


### GetCampaignNotificationUrl

`func (o *RegisterWebhookRequest) GetCampaignNotificationUrl() string`

GetCampaignNotificationUrl returns the CampaignNotificationUrl field if non-nil, zero value otherwise.

### GetCampaignNotificationUrlOk

`func (o *RegisterWebhookRequest) GetCampaignNotificationUrlOk() (*string, bool)`

GetCampaignNotificationUrlOk returns a tuple with the CampaignNotificationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignNotificationUrl

`func (o *RegisterWebhookRequest) SetCampaignNotificationUrl(v string)`

SetCampaignNotificationUrl sets CampaignNotificationUrl field to given value.


### GetWebhookSecret

`func (o *RegisterWebhookRequest) GetWebhookSecret() string`

GetWebhookSecret returns the WebhookSecret field if non-nil, zero value otherwise.

### GetWebhookSecretOk

`func (o *RegisterWebhookRequest) GetWebhookSecretOk() (*string, bool)`

GetWebhookSecretOk returns a tuple with the WebhookSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookSecret

`func (o *RegisterWebhookRequest) SetWebhookSecret(v string)`

SetWebhookSecret sets WebhookSecret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


