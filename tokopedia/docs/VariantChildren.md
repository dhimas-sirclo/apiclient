# VariantChildren

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**ProductId** | Pointer to **int32** |  | [optional] 
**Price** | Pointer to **int32** |  | [optional] 
**PriceFmt** | Pointer to **string** |  | [optional] 
**Stock** | Pointer to **int32** |  | [optional] 
**MainStock** | Pointer to **int32** |  | [optional] 
**ReserveStock** | Pointer to **int32** |  | [optional] 
**Sku** | Pointer to **string** |  | [optional] 
**OptionIds** | Pointer to **[]int32** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**IsBuyable** | Pointer to **bool** |  | [optional] 
**IsWishlist** | Pointer to **bool** |  | [optional] 
**Picture** | Pointer to [**VariantPicture**](VariantPicture.md) |  | [optional] 
**Campaign** | Pointer to [**VariantCampaign**](VariantCampaign.md) |  | [optional] 
**AlwaysAvailable** | Pointer to **bool** |  | [optional] 
**StockWording** | Pointer to **string** |  | [optional] 
**OtherVariantStock** | Pointer to **string** |  | [optional] 
**IsLimitedStock** | Pointer to **bool** |  | [optional] 

## Methods

### NewVariantChildren

`func NewVariantChildren() *VariantChildren`

NewVariantChildren instantiates a new VariantChildren object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariantChildrenWithDefaults

`func NewVariantChildrenWithDefaults() *VariantChildren`

NewVariantChildrenWithDefaults instantiates a new VariantChildren object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VariantChildren) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VariantChildren) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VariantChildren) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VariantChildren) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *VariantChildren) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VariantChildren) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VariantChildren) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *VariantChildren) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetProductId

`func (o *VariantChildren) GetProductId() int32`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *VariantChildren) GetProductIdOk() (*int32, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *VariantChildren) SetProductId(v int32)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *VariantChildren) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetPrice

`func (o *VariantChildren) GetPrice() int32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *VariantChildren) GetPriceOk() (*int32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *VariantChildren) SetPrice(v int32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *VariantChildren) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetPriceFmt

`func (o *VariantChildren) GetPriceFmt() string`

GetPriceFmt returns the PriceFmt field if non-nil, zero value otherwise.

### GetPriceFmtOk

`func (o *VariantChildren) GetPriceFmtOk() (*string, bool)`

GetPriceFmtOk returns a tuple with the PriceFmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceFmt

`func (o *VariantChildren) SetPriceFmt(v string)`

SetPriceFmt sets PriceFmt field to given value.

### HasPriceFmt

`func (o *VariantChildren) HasPriceFmt() bool`

HasPriceFmt returns a boolean if a field has been set.

### GetStock

`func (o *VariantChildren) GetStock() int32`

GetStock returns the Stock field if non-nil, zero value otherwise.

### GetStockOk

`func (o *VariantChildren) GetStockOk() (*int32, bool)`

GetStockOk returns a tuple with the Stock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStock

`func (o *VariantChildren) SetStock(v int32)`

SetStock sets Stock field to given value.

### HasStock

`func (o *VariantChildren) HasStock() bool`

HasStock returns a boolean if a field has been set.

### GetMainStock

`func (o *VariantChildren) GetMainStock() int32`

GetMainStock returns the MainStock field if non-nil, zero value otherwise.

### GetMainStockOk

`func (o *VariantChildren) GetMainStockOk() (*int32, bool)`

GetMainStockOk returns a tuple with the MainStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainStock

`func (o *VariantChildren) SetMainStock(v int32)`

SetMainStock sets MainStock field to given value.

### HasMainStock

`func (o *VariantChildren) HasMainStock() bool`

HasMainStock returns a boolean if a field has been set.

### GetReserveStock

`func (o *VariantChildren) GetReserveStock() int32`

GetReserveStock returns the ReserveStock field if non-nil, zero value otherwise.

### GetReserveStockOk

`func (o *VariantChildren) GetReserveStockOk() (*int32, bool)`

GetReserveStockOk returns a tuple with the ReserveStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserveStock

`func (o *VariantChildren) SetReserveStock(v int32)`

SetReserveStock sets ReserveStock field to given value.

### HasReserveStock

`func (o *VariantChildren) HasReserveStock() bool`

HasReserveStock returns a boolean if a field has been set.

### GetSku

`func (o *VariantChildren) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *VariantChildren) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *VariantChildren) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *VariantChildren) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetOptionIds

`func (o *VariantChildren) GetOptionIds() []int32`

GetOptionIds returns the OptionIds field if non-nil, zero value otherwise.

### GetOptionIdsOk

`func (o *VariantChildren) GetOptionIdsOk() (*[]int32, bool)`

GetOptionIdsOk returns a tuple with the OptionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionIds

`func (o *VariantChildren) SetOptionIds(v []int32)`

SetOptionIds sets OptionIds field to given value.

### HasOptionIds

`func (o *VariantChildren) HasOptionIds() bool`

HasOptionIds returns a boolean if a field has been set.

### GetEnabled

`func (o *VariantChildren) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *VariantChildren) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *VariantChildren) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *VariantChildren) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetIsBuyable

`func (o *VariantChildren) GetIsBuyable() bool`

GetIsBuyable returns the IsBuyable field if non-nil, zero value otherwise.

### GetIsBuyableOk

`func (o *VariantChildren) GetIsBuyableOk() (*bool, bool)`

GetIsBuyableOk returns a tuple with the IsBuyable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBuyable

`func (o *VariantChildren) SetIsBuyable(v bool)`

SetIsBuyable sets IsBuyable field to given value.

### HasIsBuyable

`func (o *VariantChildren) HasIsBuyable() bool`

HasIsBuyable returns a boolean if a field has been set.

### GetIsWishlist

`func (o *VariantChildren) GetIsWishlist() bool`

GetIsWishlist returns the IsWishlist field if non-nil, zero value otherwise.

### GetIsWishlistOk

`func (o *VariantChildren) GetIsWishlistOk() (*bool, bool)`

GetIsWishlistOk returns a tuple with the IsWishlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWishlist

`func (o *VariantChildren) SetIsWishlist(v bool)`

SetIsWishlist sets IsWishlist field to given value.

### HasIsWishlist

`func (o *VariantChildren) HasIsWishlist() bool`

HasIsWishlist returns a boolean if a field has been set.

### GetPicture

`func (o *VariantChildren) GetPicture() VariantPicture`

GetPicture returns the Picture field if non-nil, zero value otherwise.

### GetPictureOk

`func (o *VariantChildren) GetPictureOk() (*VariantPicture, bool)`

GetPictureOk returns a tuple with the Picture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPicture

`func (o *VariantChildren) SetPicture(v VariantPicture)`

SetPicture sets Picture field to given value.

### HasPicture

`func (o *VariantChildren) HasPicture() bool`

HasPicture returns a boolean if a field has been set.

### GetCampaign

`func (o *VariantChildren) GetCampaign() VariantCampaign`

GetCampaign returns the Campaign field if non-nil, zero value otherwise.

### GetCampaignOk

`func (o *VariantChildren) GetCampaignOk() (*VariantCampaign, bool)`

GetCampaignOk returns a tuple with the Campaign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaign

`func (o *VariantChildren) SetCampaign(v VariantCampaign)`

SetCampaign sets Campaign field to given value.

### HasCampaign

`func (o *VariantChildren) HasCampaign() bool`

HasCampaign returns a boolean if a field has been set.

### GetAlwaysAvailable

`func (o *VariantChildren) GetAlwaysAvailable() bool`

GetAlwaysAvailable returns the AlwaysAvailable field if non-nil, zero value otherwise.

### GetAlwaysAvailableOk

`func (o *VariantChildren) GetAlwaysAvailableOk() (*bool, bool)`

GetAlwaysAvailableOk returns a tuple with the AlwaysAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlwaysAvailable

`func (o *VariantChildren) SetAlwaysAvailable(v bool)`

SetAlwaysAvailable sets AlwaysAvailable field to given value.

### HasAlwaysAvailable

`func (o *VariantChildren) HasAlwaysAvailable() bool`

HasAlwaysAvailable returns a boolean if a field has been set.

### GetStockWording

`func (o *VariantChildren) GetStockWording() string`

GetStockWording returns the StockWording field if non-nil, zero value otherwise.

### GetStockWordingOk

`func (o *VariantChildren) GetStockWordingOk() (*string, bool)`

GetStockWordingOk returns a tuple with the StockWording field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockWording

`func (o *VariantChildren) SetStockWording(v string)`

SetStockWording sets StockWording field to given value.

### HasStockWording

`func (o *VariantChildren) HasStockWording() bool`

HasStockWording returns a boolean if a field has been set.

### GetOtherVariantStock

`func (o *VariantChildren) GetOtherVariantStock() string`

GetOtherVariantStock returns the OtherVariantStock field if non-nil, zero value otherwise.

### GetOtherVariantStockOk

`func (o *VariantChildren) GetOtherVariantStockOk() (*string, bool)`

GetOtherVariantStockOk returns a tuple with the OtherVariantStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherVariantStock

`func (o *VariantChildren) SetOtherVariantStock(v string)`

SetOtherVariantStock sets OtherVariantStock field to given value.

### HasOtherVariantStock

`func (o *VariantChildren) HasOtherVariantStock() bool`

HasOtherVariantStock returns a boolean if a field has been set.

### GetIsLimitedStock

`func (o *VariantChildren) GetIsLimitedStock() bool`

GetIsLimitedStock returns the IsLimitedStock field if non-nil, zero value otherwise.

### GetIsLimitedStockOk

`func (o *VariantChildren) GetIsLimitedStockOk() (*bool, bool)`

GetIsLimitedStockOk returns a tuple with the IsLimitedStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLimitedStock

`func (o *VariantChildren) SetIsLimitedStock(v bool)`

SetIsLimitedStock sets IsLimitedStock field to given value.

### HasIsLimitedStock

`func (o *VariantChildren) HasIsLimitedStock() bool`

HasIsLimitedStock returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


