# WebhookOrderNotificationBundleDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalProduct** | Pointer to **int64** |  | [optional] 
**ProductBundlingIcon** | Pointer to **string** |  | [optional] 
**Bundle** | Pointer to [**[]WebhookOrderNotificationBundleDetailBundleInner**](WebhookOrderNotificationBundleDetailBundleInner.md) | Contains information about bundle items in the order | [optional] 
**NonBundle** | Pointer to [**[]WebhookOrderNotificationBundleDetailNonBundleInner**](WebhookOrderNotificationBundleDetailNonBundleInner.md) | Contains information about non-bundle items in bundled item | [optional] 

## Methods

### NewWebhookOrderNotificationBundleDetail

`func NewWebhookOrderNotificationBundleDetail() *WebhookOrderNotificationBundleDetail`

NewWebhookOrderNotificationBundleDetail instantiates a new WebhookOrderNotificationBundleDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookOrderNotificationBundleDetailWithDefaults

`func NewWebhookOrderNotificationBundleDetailWithDefaults() *WebhookOrderNotificationBundleDetail`

NewWebhookOrderNotificationBundleDetailWithDefaults instantiates a new WebhookOrderNotificationBundleDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalProduct

`func (o *WebhookOrderNotificationBundleDetail) GetTotalProduct() int64`

GetTotalProduct returns the TotalProduct field if non-nil, zero value otherwise.

### GetTotalProductOk

`func (o *WebhookOrderNotificationBundleDetail) GetTotalProductOk() (*int64, bool)`

GetTotalProductOk returns a tuple with the TotalProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalProduct

`func (o *WebhookOrderNotificationBundleDetail) SetTotalProduct(v int64)`

SetTotalProduct sets TotalProduct field to given value.

### HasTotalProduct

`func (o *WebhookOrderNotificationBundleDetail) HasTotalProduct() bool`

HasTotalProduct returns a boolean if a field has been set.

### GetProductBundlingIcon

`func (o *WebhookOrderNotificationBundleDetail) GetProductBundlingIcon() string`

GetProductBundlingIcon returns the ProductBundlingIcon field if non-nil, zero value otherwise.

### GetProductBundlingIconOk

`func (o *WebhookOrderNotificationBundleDetail) GetProductBundlingIconOk() (*string, bool)`

GetProductBundlingIconOk returns a tuple with the ProductBundlingIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductBundlingIcon

`func (o *WebhookOrderNotificationBundleDetail) SetProductBundlingIcon(v string)`

SetProductBundlingIcon sets ProductBundlingIcon field to given value.

### HasProductBundlingIcon

`func (o *WebhookOrderNotificationBundleDetail) HasProductBundlingIcon() bool`

HasProductBundlingIcon returns a boolean if a field has been set.

### GetBundle

`func (o *WebhookOrderNotificationBundleDetail) GetBundle() []WebhookOrderNotificationBundleDetailBundleInner`

GetBundle returns the Bundle field if non-nil, zero value otherwise.

### GetBundleOk

`func (o *WebhookOrderNotificationBundleDetail) GetBundleOk() (*[]WebhookOrderNotificationBundleDetailBundleInner, bool)`

GetBundleOk returns a tuple with the Bundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundle

`func (o *WebhookOrderNotificationBundleDetail) SetBundle(v []WebhookOrderNotificationBundleDetailBundleInner)`

SetBundle sets Bundle field to given value.

### HasBundle

`func (o *WebhookOrderNotificationBundleDetail) HasBundle() bool`

HasBundle returns a boolean if a field has been set.

### GetNonBundle

`func (o *WebhookOrderNotificationBundleDetail) GetNonBundle() []WebhookOrderNotificationBundleDetailNonBundleInner`

GetNonBundle returns the NonBundle field if non-nil, zero value otherwise.

### GetNonBundleOk

`func (o *WebhookOrderNotificationBundleDetail) GetNonBundleOk() (*[]WebhookOrderNotificationBundleDetailNonBundleInner, bool)`

GetNonBundleOk returns a tuple with the NonBundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonBundle

`func (o *WebhookOrderNotificationBundleDetail) SetNonBundle(v []WebhookOrderNotificationBundleDetailNonBundleInner)`

SetNonBundle sets NonBundle field to given value.

### HasNonBundle

`func (o *WebhookOrderNotificationBundleDetail) HasNonBundle() bool`

HasNonBundle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


