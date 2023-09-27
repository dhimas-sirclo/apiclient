# ViewCampaignProducts200ResponseDataProductsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | Pointer to **int64** | Product Unique Identifier | [optional] 
**Price** | Pointer to **string** | Product Price | [optional] 
**Shop** | Pointer to **map[string]interface{}** |  | [optional] 
**MaxOrder** | Pointer to **int64** | Product Maximum Order | [optional] 
**PriceUnfmt** | Pointer to **string** | Product Price Unformated | [optional] 
**MinOrder** | Pointer to **int64** | Product Minimum Order | [optional] 
**Campaign** | Pointer to [**ViewCampaignProducts200ResponseDataProductsInnerCampaign**](ViewCampaignProducts200ResponseDataProductsInnerCampaign.md) |  | [optional] 

## Methods

### NewViewCampaignProducts200ResponseDataProductsInner

`func NewViewCampaignProducts200ResponseDataProductsInner() *ViewCampaignProducts200ResponseDataProductsInner`

NewViewCampaignProducts200ResponseDataProductsInner instantiates a new ViewCampaignProducts200ResponseDataProductsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewCampaignProducts200ResponseDataProductsInnerWithDefaults

`func NewViewCampaignProducts200ResponseDataProductsInnerWithDefaults() *ViewCampaignProducts200ResponseDataProductsInner`

NewViewCampaignProducts200ResponseDataProductsInnerWithDefaults instantiates a new ViewCampaignProducts200ResponseDataProductsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetID

`func (o *ViewCampaignProducts200ResponseDataProductsInner) GetID() int64`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *ViewCampaignProducts200ResponseDataProductsInner) GetIDOk() (*int64, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *ViewCampaignProducts200ResponseDataProductsInner) SetID(v int64)`

SetID sets ID field to given value.

### HasID

`func (o *ViewCampaignProducts200ResponseDataProductsInner) HasID() bool`

HasID returns a boolean if a field has been set.

### GetPrice

`func (o *ViewCampaignProducts200ResponseDataProductsInner) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ViewCampaignProducts200ResponseDataProductsInner) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ViewCampaignProducts200ResponseDataProductsInner) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ViewCampaignProducts200ResponseDataProductsInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetShop

`func (o *ViewCampaignProducts200ResponseDataProductsInner) GetShop() map[string]interface{}`

GetShop returns the Shop field if non-nil, zero value otherwise.

### GetShopOk

`func (o *ViewCampaignProducts200ResponseDataProductsInner) GetShopOk() (*map[string]interface{}, bool)`

GetShopOk returns a tuple with the Shop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShop

`func (o *ViewCampaignProducts200ResponseDataProductsInner) SetShop(v map[string]interface{})`

SetShop sets Shop field to given value.

### HasShop

`func (o *ViewCampaignProducts200ResponseDataProductsInner) HasShop() bool`

HasShop returns a boolean if a field has been set.

### GetMaxOrder

`func (o *ViewCampaignProducts200ResponseDataProductsInner) GetMaxOrder() int64`

GetMaxOrder returns the MaxOrder field if non-nil, zero value otherwise.

### GetMaxOrderOk

`func (o *ViewCampaignProducts200ResponseDataProductsInner) GetMaxOrderOk() (*int64, bool)`

GetMaxOrderOk returns a tuple with the MaxOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxOrder

`func (o *ViewCampaignProducts200ResponseDataProductsInner) SetMaxOrder(v int64)`

SetMaxOrder sets MaxOrder field to given value.

### HasMaxOrder

`func (o *ViewCampaignProducts200ResponseDataProductsInner) HasMaxOrder() bool`

HasMaxOrder returns a boolean if a field has been set.

### GetPriceUnfmt

`func (o *ViewCampaignProducts200ResponseDataProductsInner) GetPriceUnfmt() string`

GetPriceUnfmt returns the PriceUnfmt field if non-nil, zero value otherwise.

### GetPriceUnfmtOk

`func (o *ViewCampaignProducts200ResponseDataProductsInner) GetPriceUnfmtOk() (*string, bool)`

GetPriceUnfmtOk returns a tuple with the PriceUnfmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUnfmt

`func (o *ViewCampaignProducts200ResponseDataProductsInner) SetPriceUnfmt(v string)`

SetPriceUnfmt sets PriceUnfmt field to given value.

### HasPriceUnfmt

`func (o *ViewCampaignProducts200ResponseDataProductsInner) HasPriceUnfmt() bool`

HasPriceUnfmt returns a boolean if a field has been set.

### GetMinOrder

`func (o *ViewCampaignProducts200ResponseDataProductsInner) GetMinOrder() int64`

GetMinOrder returns the MinOrder field if non-nil, zero value otherwise.

### GetMinOrderOk

`func (o *ViewCampaignProducts200ResponseDataProductsInner) GetMinOrderOk() (*int64, bool)`

GetMinOrderOk returns a tuple with the MinOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinOrder

`func (o *ViewCampaignProducts200ResponseDataProductsInner) SetMinOrder(v int64)`

SetMinOrder sets MinOrder field to given value.

### HasMinOrder

`func (o *ViewCampaignProducts200ResponseDataProductsInner) HasMinOrder() bool`

HasMinOrder returns a boolean if a field has been set.

### GetCampaign

`func (o *ViewCampaignProducts200ResponseDataProductsInner) GetCampaign() ViewCampaignProducts200ResponseDataProductsInnerCampaign`

GetCampaign returns the Campaign field if non-nil, zero value otherwise.

### GetCampaignOk

`func (o *ViewCampaignProducts200ResponseDataProductsInner) GetCampaignOk() (*ViewCampaignProducts200ResponseDataProductsInnerCampaign, bool)`

GetCampaignOk returns a tuple with the Campaign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaign

`func (o *ViewCampaignProducts200ResponseDataProductsInner) SetCampaign(v ViewCampaignProducts200ResponseDataProductsInnerCampaign)`

SetCampaign sets Campaign field to given value.

### HasCampaign

`func (o *ViewCampaignProducts200ResponseDataProductsInner) HasCampaign() bool`

HasCampaign returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


