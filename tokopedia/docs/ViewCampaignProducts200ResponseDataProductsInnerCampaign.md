# ViewCampaignProducts200ResponseDataProductsInnerCampaign

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CampaignID** | Pointer to **int64** | Campaign Unique Identifier | [optional] 
**SlashPriceProductId** | Pointer to **int64** | Slash Price Product Unique Identifier | [optional] 
**DiscountPercentage** | Pointer to **int64** | Product Discount Percentage | [optional] 
**DiscountedPrice** | Pointer to **string** | Discounted Product Price | [optional] 
**OriginalPrice** | Pointer to **string** |  | [optional] 
**CustomStock** | Pointer to **int64** |  | [optional] 
**CampaignStatus** | Pointer to **int64** | Campaign Status | [optional] 
**StartDate** | Pointer to **string** | Campaign Start Date (format: YYYY-MM-DD HH-mm-ss)  | [optional] 
**EndDate** | Pointer to **string** | Campaign End Date (format: YYYY-MM-DD HH-mm-ss)  | [optional] 
**CampaignTypeName** | Pointer to **string** | Campaign Type Name | [optional] 
**CampaignShortName** | Pointer to **string** | Campaign Type Short Name | [optional] 
**MaxOrder** | Pointer to **int64** | Product Maximum Order | [optional] 
**DiscountedPriceFmt** | Pointer to **string** | Discounted Price Formated | [optional] 
**OriginalPriceFmt** | Pointer to **string** | Original Price Formated | [optional] 
**OriginalStock** | Pointer to **int64** | Product Original Stock | [optional] 
**MinOrder** | Pointer to **int64** | Product Minimum Order | [optional] 
**AdditionalCartInfo** | Pointer to [**ViewCampaignProducts200ResponseDataProductsInnerCampaignAdditionalCartInfo**](ViewCampaignProducts200ResponseDataProductsInnerCampaignAdditionalCartInfo.md) |  | [optional] 
**StockWording** | Pointer to **map[string]interface{}** | Stock Wording | [optional] 

## Methods

### NewViewCampaignProducts200ResponseDataProductsInnerCampaign

`func NewViewCampaignProducts200ResponseDataProductsInnerCampaign() *ViewCampaignProducts200ResponseDataProductsInnerCampaign`

NewViewCampaignProducts200ResponseDataProductsInnerCampaign instantiates a new ViewCampaignProducts200ResponseDataProductsInnerCampaign object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewCampaignProducts200ResponseDataProductsInnerCampaignWithDefaults

`func NewViewCampaignProducts200ResponseDataProductsInnerCampaignWithDefaults() *ViewCampaignProducts200ResponseDataProductsInnerCampaign`

NewViewCampaignProducts200ResponseDataProductsInnerCampaignWithDefaults instantiates a new ViewCampaignProducts200ResponseDataProductsInnerCampaign object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCampaignID

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetCampaignID() int64`

GetCampaignID returns the CampaignID field if non-nil, zero value otherwise.

### GetCampaignIDOk

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetCampaignIDOk() (*int64, bool)`

GetCampaignIDOk returns a tuple with the CampaignID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignID

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) SetCampaignID(v int64)`

SetCampaignID sets CampaignID field to given value.

### HasCampaignID

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) HasCampaignID() bool`

HasCampaignID returns a boolean if a field has been set.

### GetSlashPriceProductId

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetSlashPriceProductId() int64`

GetSlashPriceProductId returns the SlashPriceProductId field if non-nil, zero value otherwise.

### GetSlashPriceProductIdOk

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetSlashPriceProductIdOk() (*int64, bool)`

GetSlashPriceProductIdOk returns a tuple with the SlashPriceProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlashPriceProductId

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) SetSlashPriceProductId(v int64)`

SetSlashPriceProductId sets SlashPriceProductId field to given value.

### HasSlashPriceProductId

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) HasSlashPriceProductId() bool`

HasSlashPriceProductId returns a boolean if a field has been set.

### GetDiscountPercentage

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetDiscountPercentage() int64`

GetDiscountPercentage returns the DiscountPercentage field if non-nil, zero value otherwise.

### GetDiscountPercentageOk

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetDiscountPercentageOk() (*int64, bool)`

GetDiscountPercentageOk returns a tuple with the DiscountPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountPercentage

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) SetDiscountPercentage(v int64)`

SetDiscountPercentage sets DiscountPercentage field to given value.

### HasDiscountPercentage

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) HasDiscountPercentage() bool`

HasDiscountPercentage returns a boolean if a field has been set.

### GetDiscountedPrice

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetDiscountedPrice() string`

GetDiscountedPrice returns the DiscountedPrice field if non-nil, zero value otherwise.

### GetDiscountedPriceOk

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetDiscountedPriceOk() (*string, bool)`

GetDiscountedPriceOk returns a tuple with the DiscountedPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountedPrice

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) SetDiscountedPrice(v string)`

SetDiscountedPrice sets DiscountedPrice field to given value.

### HasDiscountedPrice

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) HasDiscountedPrice() bool`

HasDiscountedPrice returns a boolean if a field has been set.

### GetOriginalPrice

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetOriginalPrice() string`

GetOriginalPrice returns the OriginalPrice field if non-nil, zero value otherwise.

### GetOriginalPriceOk

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetOriginalPriceOk() (*string, bool)`

GetOriginalPriceOk returns a tuple with the OriginalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalPrice

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) SetOriginalPrice(v string)`

SetOriginalPrice sets OriginalPrice field to given value.

### HasOriginalPrice

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) HasOriginalPrice() bool`

HasOriginalPrice returns a boolean if a field has been set.

### GetCustomStock

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetCustomStock() int64`

GetCustomStock returns the CustomStock field if non-nil, zero value otherwise.

### GetCustomStockOk

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetCustomStockOk() (*int64, bool)`

GetCustomStockOk returns a tuple with the CustomStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomStock

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) SetCustomStock(v int64)`

SetCustomStock sets CustomStock field to given value.

### HasCustomStock

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) HasCustomStock() bool`

HasCustomStock returns a boolean if a field has been set.

### GetCampaignStatus

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetCampaignStatus() int64`

GetCampaignStatus returns the CampaignStatus field if non-nil, zero value otherwise.

### GetCampaignStatusOk

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetCampaignStatusOk() (*int64, bool)`

GetCampaignStatusOk returns a tuple with the CampaignStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignStatus

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) SetCampaignStatus(v int64)`

SetCampaignStatus sets CampaignStatus field to given value.

### HasCampaignStatus

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) HasCampaignStatus() bool`

HasCampaignStatus returns a boolean if a field has been set.

### GetStartDate

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetCampaignTypeName

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetCampaignTypeName() string`

GetCampaignTypeName returns the CampaignTypeName field if non-nil, zero value otherwise.

### GetCampaignTypeNameOk

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetCampaignTypeNameOk() (*string, bool)`

GetCampaignTypeNameOk returns a tuple with the CampaignTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignTypeName

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) SetCampaignTypeName(v string)`

SetCampaignTypeName sets CampaignTypeName field to given value.

### HasCampaignTypeName

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) HasCampaignTypeName() bool`

HasCampaignTypeName returns a boolean if a field has been set.

### GetCampaignShortName

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetCampaignShortName() string`

GetCampaignShortName returns the CampaignShortName field if non-nil, zero value otherwise.

### GetCampaignShortNameOk

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetCampaignShortNameOk() (*string, bool)`

GetCampaignShortNameOk returns a tuple with the CampaignShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignShortName

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) SetCampaignShortName(v string)`

SetCampaignShortName sets CampaignShortName field to given value.

### HasCampaignShortName

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) HasCampaignShortName() bool`

HasCampaignShortName returns a boolean if a field has been set.

### GetMaxOrder

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetMaxOrder() int64`

GetMaxOrder returns the MaxOrder field if non-nil, zero value otherwise.

### GetMaxOrderOk

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetMaxOrderOk() (*int64, bool)`

GetMaxOrderOk returns a tuple with the MaxOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxOrder

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) SetMaxOrder(v int64)`

SetMaxOrder sets MaxOrder field to given value.

### HasMaxOrder

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) HasMaxOrder() bool`

HasMaxOrder returns a boolean if a field has been set.

### GetDiscountedPriceFmt

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetDiscountedPriceFmt() string`

GetDiscountedPriceFmt returns the DiscountedPriceFmt field if non-nil, zero value otherwise.

### GetDiscountedPriceFmtOk

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetDiscountedPriceFmtOk() (*string, bool)`

GetDiscountedPriceFmtOk returns a tuple with the DiscountedPriceFmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountedPriceFmt

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) SetDiscountedPriceFmt(v string)`

SetDiscountedPriceFmt sets DiscountedPriceFmt field to given value.

### HasDiscountedPriceFmt

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) HasDiscountedPriceFmt() bool`

HasDiscountedPriceFmt returns a boolean if a field has been set.

### GetOriginalPriceFmt

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetOriginalPriceFmt() string`

GetOriginalPriceFmt returns the OriginalPriceFmt field if non-nil, zero value otherwise.

### GetOriginalPriceFmtOk

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetOriginalPriceFmtOk() (*string, bool)`

GetOriginalPriceFmtOk returns a tuple with the OriginalPriceFmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalPriceFmt

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) SetOriginalPriceFmt(v string)`

SetOriginalPriceFmt sets OriginalPriceFmt field to given value.

### HasOriginalPriceFmt

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) HasOriginalPriceFmt() bool`

HasOriginalPriceFmt returns a boolean if a field has been set.

### GetOriginalStock

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetOriginalStock() int64`

GetOriginalStock returns the OriginalStock field if non-nil, zero value otherwise.

### GetOriginalStockOk

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetOriginalStockOk() (*int64, bool)`

GetOriginalStockOk returns a tuple with the OriginalStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalStock

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) SetOriginalStock(v int64)`

SetOriginalStock sets OriginalStock field to given value.

### HasOriginalStock

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) HasOriginalStock() bool`

HasOriginalStock returns a boolean if a field has been set.

### GetMinOrder

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetMinOrder() int64`

GetMinOrder returns the MinOrder field if non-nil, zero value otherwise.

### GetMinOrderOk

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetMinOrderOk() (*int64, bool)`

GetMinOrderOk returns a tuple with the MinOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinOrder

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) SetMinOrder(v int64)`

SetMinOrder sets MinOrder field to given value.

### HasMinOrder

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) HasMinOrder() bool`

HasMinOrder returns a boolean if a field has been set.

### GetAdditionalCartInfo

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetAdditionalCartInfo() ViewCampaignProducts200ResponseDataProductsInnerCampaignAdditionalCartInfo`

GetAdditionalCartInfo returns the AdditionalCartInfo field if non-nil, zero value otherwise.

### GetAdditionalCartInfoOk

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetAdditionalCartInfoOk() (*ViewCampaignProducts200ResponseDataProductsInnerCampaignAdditionalCartInfo, bool)`

GetAdditionalCartInfoOk returns a tuple with the AdditionalCartInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalCartInfo

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) SetAdditionalCartInfo(v ViewCampaignProducts200ResponseDataProductsInnerCampaignAdditionalCartInfo)`

SetAdditionalCartInfo sets AdditionalCartInfo field to given value.

### HasAdditionalCartInfo

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) HasAdditionalCartInfo() bool`

HasAdditionalCartInfo returns a boolean if a field has been set.

### GetStockWording

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetStockWording() map[string]interface{}`

GetStockWording returns the StockWording field if non-nil, zero value otherwise.

### GetStockWordingOk

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetStockWordingOk() (*map[string]interface{}, bool)`

GetStockWordingOk returns a tuple with the StockWording field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockWording

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) SetStockWording(v map[string]interface{})`

SetStockWording sets StockWording field to given value.

### HasStockWording

`func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) HasStockWording() bool`

HasStockWording returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


