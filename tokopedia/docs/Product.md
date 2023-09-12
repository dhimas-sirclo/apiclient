# Product

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Basic** | Pointer to [**ProductBasic**](ProductBasic.md) |  | [optional] 
**Price** | Pointer to [**ProductPrice**](ProductPrice.md) |  | [optional] 
**Weight** | Pointer to [**ProductWeight**](ProductWeight.md) |  | [optional] 
**Stock** | Pointer to [**ProductStockWording**](ProductStockWording.md) |  | [optional] 
**MainStock** | Pointer to **int32** | Product Stock (Not Reserved) | [optional] 
**ReserveStock** | Pointer to **int32** | Product Stock that Reserved (ex: FlashSale) | [optional] 
**Variant** | Pointer to [**ProductVariant**](ProductVariant.md) |  | [optional] 
**Menu** | Pointer to [**ProductMenu**](ProductMenu.md) |  | [optional] 
**Preorder** | Pointer to [**ProductPreOrder**](ProductPreOrder.md) |  | [optional] 
**ExtraAttribute** | Pointer to [**ProductExtraAttribute**](ProductExtraAttribute.md) |  | [optional] 
**Wholesale** | Pointer to [**[]ProductWholesale**](ProductWholesale.md) |  | [optional] 
**CategoryTree** | Pointer to [**[]ProductCategoryTree**](ProductCategoryTree.md) |  | [optional] 
**Pictures** | Pointer to [**[]ProductPicture**](ProductPicture.md) |  | [optional] 
**GMStats** | Pointer to [**ProductGMStat**](ProductGMStat.md) |  | [optional] 
**Stats** | Pointer to [**ProductStats**](ProductStats.md) |  | [optional] 
**Other** | Pointer to [**ProductOther**](ProductOther.md) |  | [optional] 
**Campaign** | Pointer to [**ProductCampaign**](ProductCampaign.md) |  | [optional] 
**Volume** | Pointer to [**ProductVolume**](ProductVolume.md) |  | [optional] 
**Warehouses** | Pointer to [**[]ProductWarehouse**](ProductWarehouse.md) |  | [optional] 

## Methods

### NewProduct

`func NewProduct() *Product`

NewProduct instantiates a new Product object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductWithDefaults

`func NewProductWithDefaults() *Product`

NewProductWithDefaults instantiates a new Product object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBasic

`func (o *Product) GetBasic() ProductBasic`

GetBasic returns the Basic field if non-nil, zero value otherwise.

### GetBasicOk

`func (o *Product) GetBasicOk() (*ProductBasic, bool)`

GetBasicOk returns a tuple with the Basic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasic

`func (o *Product) SetBasic(v ProductBasic)`

SetBasic sets Basic field to given value.

### HasBasic

`func (o *Product) HasBasic() bool`

HasBasic returns a boolean if a field has been set.

### GetPrice

`func (o *Product) GetPrice() ProductPrice`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *Product) GetPriceOk() (*ProductPrice, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *Product) SetPrice(v ProductPrice)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *Product) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetWeight

`func (o *Product) GetWeight() ProductWeight`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *Product) GetWeightOk() (*ProductWeight, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *Product) SetWeight(v ProductWeight)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *Product) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetStock

`func (o *Product) GetStock() ProductStockWording`

GetStock returns the Stock field if non-nil, zero value otherwise.

### GetStockOk

`func (o *Product) GetStockOk() (*ProductStockWording, bool)`

GetStockOk returns a tuple with the Stock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStock

`func (o *Product) SetStock(v ProductStockWording)`

SetStock sets Stock field to given value.

### HasStock

`func (o *Product) HasStock() bool`

HasStock returns a boolean if a field has been set.

### GetMainStock

`func (o *Product) GetMainStock() int32`

GetMainStock returns the MainStock field if non-nil, zero value otherwise.

### GetMainStockOk

`func (o *Product) GetMainStockOk() (*int32, bool)`

GetMainStockOk returns a tuple with the MainStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainStock

`func (o *Product) SetMainStock(v int32)`

SetMainStock sets MainStock field to given value.

### HasMainStock

`func (o *Product) HasMainStock() bool`

HasMainStock returns a boolean if a field has been set.

### GetReserveStock

`func (o *Product) GetReserveStock() int32`

GetReserveStock returns the ReserveStock field if non-nil, zero value otherwise.

### GetReserveStockOk

`func (o *Product) GetReserveStockOk() (*int32, bool)`

GetReserveStockOk returns a tuple with the ReserveStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserveStock

`func (o *Product) SetReserveStock(v int32)`

SetReserveStock sets ReserveStock field to given value.

### HasReserveStock

`func (o *Product) HasReserveStock() bool`

HasReserveStock returns a boolean if a field has been set.

### GetVariant

`func (o *Product) GetVariant() ProductVariant`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *Product) GetVariantOk() (*ProductVariant, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *Product) SetVariant(v ProductVariant)`

SetVariant sets Variant field to given value.

### HasVariant

`func (o *Product) HasVariant() bool`

HasVariant returns a boolean if a field has been set.

### GetMenu

`func (o *Product) GetMenu() ProductMenu`

GetMenu returns the Menu field if non-nil, zero value otherwise.

### GetMenuOk

`func (o *Product) GetMenuOk() (*ProductMenu, bool)`

GetMenuOk returns a tuple with the Menu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMenu

`func (o *Product) SetMenu(v ProductMenu)`

SetMenu sets Menu field to given value.

### HasMenu

`func (o *Product) HasMenu() bool`

HasMenu returns a boolean if a field has been set.

### GetPreorder

`func (o *Product) GetPreorder() ProductPreOrder`

GetPreorder returns the Preorder field if non-nil, zero value otherwise.

### GetPreorderOk

`func (o *Product) GetPreorderOk() (*ProductPreOrder, bool)`

GetPreorderOk returns a tuple with the Preorder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreorder

`func (o *Product) SetPreorder(v ProductPreOrder)`

SetPreorder sets Preorder field to given value.

### HasPreorder

`func (o *Product) HasPreorder() bool`

HasPreorder returns a boolean if a field has been set.

### GetExtraAttribute

`func (o *Product) GetExtraAttribute() ProductExtraAttribute`

GetExtraAttribute returns the ExtraAttribute field if non-nil, zero value otherwise.

### GetExtraAttributeOk

`func (o *Product) GetExtraAttributeOk() (*ProductExtraAttribute, bool)`

GetExtraAttributeOk returns a tuple with the ExtraAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraAttribute

`func (o *Product) SetExtraAttribute(v ProductExtraAttribute)`

SetExtraAttribute sets ExtraAttribute field to given value.

### HasExtraAttribute

`func (o *Product) HasExtraAttribute() bool`

HasExtraAttribute returns a boolean if a field has been set.

### GetWholesale

`func (o *Product) GetWholesale() []ProductWholesale`

GetWholesale returns the Wholesale field if non-nil, zero value otherwise.

### GetWholesaleOk

`func (o *Product) GetWholesaleOk() (*[]ProductWholesale, bool)`

GetWholesaleOk returns a tuple with the Wholesale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWholesale

`func (o *Product) SetWholesale(v []ProductWholesale)`

SetWholesale sets Wholesale field to given value.

### HasWholesale

`func (o *Product) HasWholesale() bool`

HasWholesale returns a boolean if a field has been set.

### GetCategoryTree

`func (o *Product) GetCategoryTree() []ProductCategoryTree`

GetCategoryTree returns the CategoryTree field if non-nil, zero value otherwise.

### GetCategoryTreeOk

`func (o *Product) GetCategoryTreeOk() (*[]ProductCategoryTree, bool)`

GetCategoryTreeOk returns a tuple with the CategoryTree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryTree

`func (o *Product) SetCategoryTree(v []ProductCategoryTree)`

SetCategoryTree sets CategoryTree field to given value.

### HasCategoryTree

`func (o *Product) HasCategoryTree() bool`

HasCategoryTree returns a boolean if a field has been set.

### GetPictures

`func (o *Product) GetPictures() []ProductPicture`

GetPictures returns the Pictures field if non-nil, zero value otherwise.

### GetPicturesOk

`func (o *Product) GetPicturesOk() (*[]ProductPicture, bool)`

GetPicturesOk returns a tuple with the Pictures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPictures

`func (o *Product) SetPictures(v []ProductPicture)`

SetPictures sets Pictures field to given value.

### HasPictures

`func (o *Product) HasPictures() bool`

HasPictures returns a boolean if a field has been set.

### GetGMStats

`func (o *Product) GetGMStats() ProductGMStat`

GetGMStats returns the GMStats field if non-nil, zero value otherwise.

### GetGMStatsOk

`func (o *Product) GetGMStatsOk() (*ProductGMStat, bool)`

GetGMStatsOk returns a tuple with the GMStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGMStats

`func (o *Product) SetGMStats(v ProductGMStat)`

SetGMStats sets GMStats field to given value.

### HasGMStats

`func (o *Product) HasGMStats() bool`

HasGMStats returns a boolean if a field has been set.

### GetStats

`func (o *Product) GetStats() ProductStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *Product) GetStatsOk() (*ProductStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *Product) SetStats(v ProductStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *Product) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetOther

`func (o *Product) GetOther() ProductOther`

GetOther returns the Other field if non-nil, zero value otherwise.

### GetOtherOk

`func (o *Product) GetOtherOk() (*ProductOther, bool)`

GetOtherOk returns a tuple with the Other field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOther

`func (o *Product) SetOther(v ProductOther)`

SetOther sets Other field to given value.

### HasOther

`func (o *Product) HasOther() bool`

HasOther returns a boolean if a field has been set.

### GetCampaign

`func (o *Product) GetCampaign() ProductCampaign`

GetCampaign returns the Campaign field if non-nil, zero value otherwise.

### GetCampaignOk

`func (o *Product) GetCampaignOk() (*ProductCampaign, bool)`

GetCampaignOk returns a tuple with the Campaign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaign

`func (o *Product) SetCampaign(v ProductCampaign)`

SetCampaign sets Campaign field to given value.

### HasCampaign

`func (o *Product) HasCampaign() bool`

HasCampaign returns a boolean if a field has been set.

### GetVolume

`func (o *Product) GetVolume() ProductVolume`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *Product) GetVolumeOk() (*ProductVolume, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *Product) SetVolume(v ProductVolume)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *Product) HasVolume() bool`

HasVolume returns a boolean if a field has been set.

### GetWarehouses

`func (o *Product) GetWarehouses() []ProductWarehouse`

GetWarehouses returns the Warehouses field if non-nil, zero value otherwise.

### GetWarehousesOk

`func (o *Product) GetWarehousesOk() (*[]ProductWarehouse, bool)`

GetWarehousesOk returns a tuple with the Warehouses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouses

`func (o *Product) SetWarehouses(v []ProductWarehouse)`

SetWarehouses sets Warehouses field to given value.

### HasWarehouses

`func (o *Product) HasWarehouses() bool`

HasWarehouses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


