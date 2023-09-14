# WebhookOrderNotificationProductsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Product ID | [optional] 
**Name** | Pointer to **string** | Product name | [optional] 
**Quantity** | Pointer to **int32** | Product quantity | [optional] 
**Notes** | Pointer to **string** | Product notes | [optional] 
**Weight** | Pointer to **float64** |  | [optional] 
**TotalWeight** | Pointer to **float64** |  | [optional] 
**Price** | Pointer to **int64** | Product price | [optional] 
**TotalPrice** | Pointer to **int64** | Total price | [optional] 
**Currency** | Pointer to **string** | Currency code | [optional] 
**Sku** | Pointer to **string** | SKU code | [optional] 
**IsEligiblePof** | Pointer to **bool** | Flag which determines if the product is eligible for POF | [optional] 

## Methods

### NewWebhookOrderNotificationProductsInner

`func NewWebhookOrderNotificationProductsInner() *WebhookOrderNotificationProductsInner`

NewWebhookOrderNotificationProductsInner instantiates a new WebhookOrderNotificationProductsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookOrderNotificationProductsInnerWithDefaults

`func NewWebhookOrderNotificationProductsInnerWithDefaults() *WebhookOrderNotificationProductsInner`

NewWebhookOrderNotificationProductsInnerWithDefaults instantiates a new WebhookOrderNotificationProductsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookOrderNotificationProductsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookOrderNotificationProductsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookOrderNotificationProductsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *WebhookOrderNotificationProductsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *WebhookOrderNotificationProductsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebhookOrderNotificationProductsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebhookOrderNotificationProductsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WebhookOrderNotificationProductsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetQuantity

`func (o *WebhookOrderNotificationProductsInner) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *WebhookOrderNotificationProductsInner) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *WebhookOrderNotificationProductsInner) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *WebhookOrderNotificationProductsInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetNotes

`func (o *WebhookOrderNotificationProductsInner) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *WebhookOrderNotificationProductsInner) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *WebhookOrderNotificationProductsInner) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *WebhookOrderNotificationProductsInner) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetWeight

`func (o *WebhookOrderNotificationProductsInner) GetWeight() float64`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *WebhookOrderNotificationProductsInner) GetWeightOk() (*float64, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *WebhookOrderNotificationProductsInner) SetWeight(v float64)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *WebhookOrderNotificationProductsInner) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetTotalWeight

`func (o *WebhookOrderNotificationProductsInner) GetTotalWeight() float64`

GetTotalWeight returns the TotalWeight field if non-nil, zero value otherwise.

### GetTotalWeightOk

`func (o *WebhookOrderNotificationProductsInner) GetTotalWeightOk() (*float64, bool)`

GetTotalWeightOk returns a tuple with the TotalWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWeight

`func (o *WebhookOrderNotificationProductsInner) SetTotalWeight(v float64)`

SetTotalWeight sets TotalWeight field to given value.

### HasTotalWeight

`func (o *WebhookOrderNotificationProductsInner) HasTotalWeight() bool`

HasTotalWeight returns a boolean if a field has been set.

### GetPrice

`func (o *WebhookOrderNotificationProductsInner) GetPrice() int64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *WebhookOrderNotificationProductsInner) GetPriceOk() (*int64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *WebhookOrderNotificationProductsInner) SetPrice(v int64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *WebhookOrderNotificationProductsInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetTotalPrice

`func (o *WebhookOrderNotificationProductsInner) GetTotalPrice() int64`

GetTotalPrice returns the TotalPrice field if non-nil, zero value otherwise.

### GetTotalPriceOk

`func (o *WebhookOrderNotificationProductsInner) GetTotalPriceOk() (*int64, bool)`

GetTotalPriceOk returns a tuple with the TotalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPrice

`func (o *WebhookOrderNotificationProductsInner) SetTotalPrice(v int64)`

SetTotalPrice sets TotalPrice field to given value.

### HasTotalPrice

`func (o *WebhookOrderNotificationProductsInner) HasTotalPrice() bool`

HasTotalPrice returns a boolean if a field has been set.

### GetCurrency

`func (o *WebhookOrderNotificationProductsInner) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *WebhookOrderNotificationProductsInner) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *WebhookOrderNotificationProductsInner) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *WebhookOrderNotificationProductsInner) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetSku

`func (o *WebhookOrderNotificationProductsInner) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *WebhookOrderNotificationProductsInner) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *WebhookOrderNotificationProductsInner) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *WebhookOrderNotificationProductsInner) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetIsEligiblePof

`func (o *WebhookOrderNotificationProductsInner) GetIsEligiblePof() bool`

GetIsEligiblePof returns the IsEligiblePof field if non-nil, zero value otherwise.

### GetIsEligiblePofOk

`func (o *WebhookOrderNotificationProductsInner) GetIsEligiblePofOk() (*bool, bool)`

GetIsEligiblePofOk returns a tuple with the IsEligiblePof field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEligiblePof

`func (o *WebhookOrderNotificationProductsInner) SetIsEligiblePof(v bool)`

SetIsEligiblePof sets IsEligiblePof field to given value.

### HasIsEligiblePof

`func (o *WebhookOrderNotificationProductsInner) HasIsEligiblePof() bool`

HasIsEligiblePof returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


