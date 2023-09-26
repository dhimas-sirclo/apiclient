/*
Tokopedia API

Tokopedia API

API version: 1.0
Contact: dev@sirclo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tokopedia

import (
	"encoding/json"
)

// checks if the Product type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Product{}

// Product struct for Product
type Product struct {
	Basic *ProductBasic `json:"basic,omitempty"`
	Price *ProductPrice `json:"price,omitempty"`
	Weight *ProductWeight `json:"weight,omitempty"`
	Stock *ProductStockWording `json:"stock,omitempty"`
	// Product Stock (Not Reserved)
	MainStock *int64 `json:"main_stock,omitempty"`
	// Product Stock that Reserved (ex: FlashSale)
	ReserveStock *int64 `json:"reserve_stock,omitempty"`
	Variant *ProductVariant `json:"variant,omitempty"`
	Menu *ProductMenu `json:"menu,omitempty"`
	Preorder *ProductPreOrder `json:"preorder,omitempty"`
	ExtraAttribute *ProductExtraAttribute `json:"extraAttribute,omitempty"`
	Wholesale []ProductWholesale `json:"wholesale,omitempty"`
	CategoryTree []ProductCategoryTree `json:"categoryTree,omitempty"`
	Pictures []ProductPicture `json:"pictures,omitempty"`
	GMStats *ProductGMStat `json:"GMStats,omitempty"`
	Stats *ProductStats `json:"stats,omitempty"`
	Other *ProductOther `json:"other,omitempty"`
	Campaign *ProductCampaign `json:"campaign,omitempty"`
	Volume *ProductVolume `json:"volume,omitempty"`
	Warehouses []ProductWarehouse `json:"warehouses,omitempty"`
}

// NewProduct instantiates a new Product object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProduct() *Product {
	this := Product{}
	return &this
}

// NewProductWithDefaults instantiates a new Product object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductWithDefaults() *Product {
	this := Product{}
	return &this
}

// GetBasic returns the Basic field value if set, zero value otherwise.
func (o *Product) GetBasic() ProductBasic {
	if o == nil || IsNil(o.Basic) {
		var ret ProductBasic
		return ret
	}
	return *o.Basic
}

// GetBasicOk returns a tuple with the Basic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetBasicOk() (*ProductBasic, bool) {
	if o == nil || IsNil(o.Basic) {
		return nil, false
	}
	return o.Basic, true
}

// HasBasic returns a boolean if a field has been set.
func (o *Product) HasBasic() bool {
	if o != nil && !IsNil(o.Basic) {
		return true
	}

	return false
}

// SetBasic gets a reference to the given ProductBasic and assigns it to the Basic field.
func (o *Product) SetBasic(v ProductBasic) {
	o.Basic = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *Product) GetPrice() ProductPrice {
	if o == nil || IsNil(o.Price) {
		var ret ProductPrice
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetPriceOk() (*ProductPrice, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *Product) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given ProductPrice and assigns it to the Price field.
func (o *Product) SetPrice(v ProductPrice) {
	o.Price = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *Product) GetWeight() ProductWeight {
	if o == nil || IsNil(o.Weight) {
		var ret ProductWeight
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetWeightOk() (*ProductWeight, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *Product) HasWeight() bool {
	if o != nil && !IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given ProductWeight and assigns it to the Weight field.
func (o *Product) SetWeight(v ProductWeight) {
	o.Weight = &v
}

// GetStock returns the Stock field value if set, zero value otherwise.
func (o *Product) GetStock() ProductStockWording {
	if o == nil || IsNil(o.Stock) {
		var ret ProductStockWording
		return ret
	}
	return *o.Stock
}

// GetStockOk returns a tuple with the Stock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetStockOk() (*ProductStockWording, bool) {
	if o == nil || IsNil(o.Stock) {
		return nil, false
	}
	return o.Stock, true
}

// HasStock returns a boolean if a field has been set.
func (o *Product) HasStock() bool {
	if o != nil && !IsNil(o.Stock) {
		return true
	}

	return false
}

// SetStock gets a reference to the given ProductStockWording and assigns it to the Stock field.
func (o *Product) SetStock(v ProductStockWording) {
	o.Stock = &v
}

// GetMainStock returns the MainStock field value if set, zero value otherwise.
func (o *Product) GetMainStock() int64 {
	if o == nil || IsNil(o.MainStock) {
		var ret int64
		return ret
	}
	return *o.MainStock
}

// GetMainStockOk returns a tuple with the MainStock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetMainStockOk() (*int64, bool) {
	if o == nil || IsNil(o.MainStock) {
		return nil, false
	}
	return o.MainStock, true
}

// HasMainStock returns a boolean if a field has been set.
func (o *Product) HasMainStock() bool {
	if o != nil && !IsNil(o.MainStock) {
		return true
	}

	return false
}

// SetMainStock gets a reference to the given int64 and assigns it to the MainStock field.
func (o *Product) SetMainStock(v int64) {
	o.MainStock = &v
}

// GetReserveStock returns the ReserveStock field value if set, zero value otherwise.
func (o *Product) GetReserveStock() int64 {
	if o == nil || IsNil(o.ReserveStock) {
		var ret int64
		return ret
	}
	return *o.ReserveStock
}

// GetReserveStockOk returns a tuple with the ReserveStock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetReserveStockOk() (*int64, bool) {
	if o == nil || IsNil(o.ReserveStock) {
		return nil, false
	}
	return o.ReserveStock, true
}

// HasReserveStock returns a boolean if a field has been set.
func (o *Product) HasReserveStock() bool {
	if o != nil && !IsNil(o.ReserveStock) {
		return true
	}

	return false
}

// SetReserveStock gets a reference to the given int64 and assigns it to the ReserveStock field.
func (o *Product) SetReserveStock(v int64) {
	o.ReserveStock = &v
}

// GetVariant returns the Variant field value if set, zero value otherwise.
func (o *Product) GetVariant() ProductVariant {
	if o == nil || IsNil(o.Variant) {
		var ret ProductVariant
		return ret
	}
	return *o.Variant
}

// GetVariantOk returns a tuple with the Variant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetVariantOk() (*ProductVariant, bool) {
	if o == nil || IsNil(o.Variant) {
		return nil, false
	}
	return o.Variant, true
}

// HasVariant returns a boolean if a field has been set.
func (o *Product) HasVariant() bool {
	if o != nil && !IsNil(o.Variant) {
		return true
	}

	return false
}

// SetVariant gets a reference to the given ProductVariant and assigns it to the Variant field.
func (o *Product) SetVariant(v ProductVariant) {
	o.Variant = &v
}

// GetMenu returns the Menu field value if set, zero value otherwise.
func (o *Product) GetMenu() ProductMenu {
	if o == nil || IsNil(o.Menu) {
		var ret ProductMenu
		return ret
	}
	return *o.Menu
}

// GetMenuOk returns a tuple with the Menu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetMenuOk() (*ProductMenu, bool) {
	if o == nil || IsNil(o.Menu) {
		return nil, false
	}
	return o.Menu, true
}

// HasMenu returns a boolean if a field has been set.
func (o *Product) HasMenu() bool {
	if o != nil && !IsNil(o.Menu) {
		return true
	}

	return false
}

// SetMenu gets a reference to the given ProductMenu and assigns it to the Menu field.
func (o *Product) SetMenu(v ProductMenu) {
	o.Menu = &v
}

// GetPreorder returns the Preorder field value if set, zero value otherwise.
func (o *Product) GetPreorder() ProductPreOrder {
	if o == nil || IsNil(o.Preorder) {
		var ret ProductPreOrder
		return ret
	}
	return *o.Preorder
}

// GetPreorderOk returns a tuple with the Preorder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetPreorderOk() (*ProductPreOrder, bool) {
	if o == nil || IsNil(o.Preorder) {
		return nil, false
	}
	return o.Preorder, true
}

// HasPreorder returns a boolean if a field has been set.
func (o *Product) HasPreorder() bool {
	if o != nil && !IsNil(o.Preorder) {
		return true
	}

	return false
}

// SetPreorder gets a reference to the given ProductPreOrder and assigns it to the Preorder field.
func (o *Product) SetPreorder(v ProductPreOrder) {
	o.Preorder = &v
}

// GetExtraAttribute returns the ExtraAttribute field value if set, zero value otherwise.
func (o *Product) GetExtraAttribute() ProductExtraAttribute {
	if o == nil || IsNil(o.ExtraAttribute) {
		var ret ProductExtraAttribute
		return ret
	}
	return *o.ExtraAttribute
}

// GetExtraAttributeOk returns a tuple with the ExtraAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetExtraAttributeOk() (*ProductExtraAttribute, bool) {
	if o == nil || IsNil(o.ExtraAttribute) {
		return nil, false
	}
	return o.ExtraAttribute, true
}

// HasExtraAttribute returns a boolean if a field has been set.
func (o *Product) HasExtraAttribute() bool {
	if o != nil && !IsNil(o.ExtraAttribute) {
		return true
	}

	return false
}

// SetExtraAttribute gets a reference to the given ProductExtraAttribute and assigns it to the ExtraAttribute field.
func (o *Product) SetExtraAttribute(v ProductExtraAttribute) {
	o.ExtraAttribute = &v
}

// GetWholesale returns the Wholesale field value if set, zero value otherwise.
func (o *Product) GetWholesale() []ProductWholesale {
	if o == nil || IsNil(o.Wholesale) {
		var ret []ProductWholesale
		return ret
	}
	return o.Wholesale
}

// GetWholesaleOk returns a tuple with the Wholesale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetWholesaleOk() ([]ProductWholesale, bool) {
	if o == nil || IsNil(o.Wholesale) {
		return nil, false
	}
	return o.Wholesale, true
}

// HasWholesale returns a boolean if a field has been set.
func (o *Product) HasWholesale() bool {
	if o != nil && !IsNil(o.Wholesale) {
		return true
	}

	return false
}

// SetWholesale gets a reference to the given []ProductWholesale and assigns it to the Wholesale field.
func (o *Product) SetWholesale(v []ProductWholesale) {
	o.Wholesale = v
}

// GetCategoryTree returns the CategoryTree field value if set, zero value otherwise.
func (o *Product) GetCategoryTree() []ProductCategoryTree {
	if o == nil || IsNil(o.CategoryTree) {
		var ret []ProductCategoryTree
		return ret
	}
	return o.CategoryTree
}

// GetCategoryTreeOk returns a tuple with the CategoryTree field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetCategoryTreeOk() ([]ProductCategoryTree, bool) {
	if o == nil || IsNil(o.CategoryTree) {
		return nil, false
	}
	return o.CategoryTree, true
}

// HasCategoryTree returns a boolean if a field has been set.
func (o *Product) HasCategoryTree() bool {
	if o != nil && !IsNil(o.CategoryTree) {
		return true
	}

	return false
}

// SetCategoryTree gets a reference to the given []ProductCategoryTree and assigns it to the CategoryTree field.
func (o *Product) SetCategoryTree(v []ProductCategoryTree) {
	o.CategoryTree = v
}

// GetPictures returns the Pictures field value if set, zero value otherwise.
func (o *Product) GetPictures() []ProductPicture {
	if o == nil || IsNil(o.Pictures) {
		var ret []ProductPicture
		return ret
	}
	return o.Pictures
}

// GetPicturesOk returns a tuple with the Pictures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetPicturesOk() ([]ProductPicture, bool) {
	if o == nil || IsNil(o.Pictures) {
		return nil, false
	}
	return o.Pictures, true
}

// HasPictures returns a boolean if a field has been set.
func (o *Product) HasPictures() bool {
	if o != nil && !IsNil(o.Pictures) {
		return true
	}

	return false
}

// SetPictures gets a reference to the given []ProductPicture and assigns it to the Pictures field.
func (o *Product) SetPictures(v []ProductPicture) {
	o.Pictures = v
}

// GetGMStats returns the GMStats field value if set, zero value otherwise.
func (o *Product) GetGMStats() ProductGMStat {
	if o == nil || IsNil(o.GMStats) {
		var ret ProductGMStat
		return ret
	}
	return *o.GMStats
}

// GetGMStatsOk returns a tuple with the GMStats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetGMStatsOk() (*ProductGMStat, bool) {
	if o == nil || IsNil(o.GMStats) {
		return nil, false
	}
	return o.GMStats, true
}

// HasGMStats returns a boolean if a field has been set.
func (o *Product) HasGMStats() bool {
	if o != nil && !IsNil(o.GMStats) {
		return true
	}

	return false
}

// SetGMStats gets a reference to the given ProductGMStat and assigns it to the GMStats field.
func (o *Product) SetGMStats(v ProductGMStat) {
	o.GMStats = &v
}

// GetStats returns the Stats field value if set, zero value otherwise.
func (o *Product) GetStats() ProductStats {
	if o == nil || IsNil(o.Stats) {
		var ret ProductStats
		return ret
	}
	return *o.Stats
}

// GetStatsOk returns a tuple with the Stats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetStatsOk() (*ProductStats, bool) {
	if o == nil || IsNil(o.Stats) {
		return nil, false
	}
	return o.Stats, true
}

// HasStats returns a boolean if a field has been set.
func (o *Product) HasStats() bool {
	if o != nil && !IsNil(o.Stats) {
		return true
	}

	return false
}

// SetStats gets a reference to the given ProductStats and assigns it to the Stats field.
func (o *Product) SetStats(v ProductStats) {
	o.Stats = &v
}

// GetOther returns the Other field value if set, zero value otherwise.
func (o *Product) GetOther() ProductOther {
	if o == nil || IsNil(o.Other) {
		var ret ProductOther
		return ret
	}
	return *o.Other
}

// GetOtherOk returns a tuple with the Other field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetOtherOk() (*ProductOther, bool) {
	if o == nil || IsNil(o.Other) {
		return nil, false
	}
	return o.Other, true
}

// HasOther returns a boolean if a field has been set.
func (o *Product) HasOther() bool {
	if o != nil && !IsNil(o.Other) {
		return true
	}

	return false
}

// SetOther gets a reference to the given ProductOther and assigns it to the Other field.
func (o *Product) SetOther(v ProductOther) {
	o.Other = &v
}

// GetCampaign returns the Campaign field value if set, zero value otherwise.
func (o *Product) GetCampaign() ProductCampaign {
	if o == nil || IsNil(o.Campaign) {
		var ret ProductCampaign
		return ret
	}
	return *o.Campaign
}

// GetCampaignOk returns a tuple with the Campaign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetCampaignOk() (*ProductCampaign, bool) {
	if o == nil || IsNil(o.Campaign) {
		return nil, false
	}
	return o.Campaign, true
}

// HasCampaign returns a boolean if a field has been set.
func (o *Product) HasCampaign() bool {
	if o != nil && !IsNil(o.Campaign) {
		return true
	}

	return false
}

// SetCampaign gets a reference to the given ProductCampaign and assigns it to the Campaign field.
func (o *Product) SetCampaign(v ProductCampaign) {
	o.Campaign = &v
}

// GetVolume returns the Volume field value if set, zero value otherwise.
func (o *Product) GetVolume() ProductVolume {
	if o == nil || IsNil(o.Volume) {
		var ret ProductVolume
		return ret
	}
	return *o.Volume
}

// GetVolumeOk returns a tuple with the Volume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetVolumeOk() (*ProductVolume, bool) {
	if o == nil || IsNil(o.Volume) {
		return nil, false
	}
	return o.Volume, true
}

// HasVolume returns a boolean if a field has been set.
func (o *Product) HasVolume() bool {
	if o != nil && !IsNil(o.Volume) {
		return true
	}

	return false
}

// SetVolume gets a reference to the given ProductVolume and assigns it to the Volume field.
func (o *Product) SetVolume(v ProductVolume) {
	o.Volume = &v
}

// GetWarehouses returns the Warehouses field value if set, zero value otherwise.
func (o *Product) GetWarehouses() []ProductWarehouse {
	if o == nil || IsNil(o.Warehouses) {
		var ret []ProductWarehouse
		return ret
	}
	return o.Warehouses
}

// GetWarehousesOk returns a tuple with the Warehouses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetWarehousesOk() ([]ProductWarehouse, bool) {
	if o == nil || IsNil(o.Warehouses) {
		return nil, false
	}
	return o.Warehouses, true
}

// HasWarehouses returns a boolean if a field has been set.
func (o *Product) HasWarehouses() bool {
	if o != nil && !IsNil(o.Warehouses) {
		return true
	}

	return false
}

// SetWarehouses gets a reference to the given []ProductWarehouse and assigns it to the Warehouses field.
func (o *Product) SetWarehouses(v []ProductWarehouse) {
	o.Warehouses = v
}

func (o Product) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Product) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Basic) {
		toSerialize["basic"] = o.Basic
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	if !IsNil(o.Stock) {
		toSerialize["stock"] = o.Stock
	}
	if !IsNil(o.MainStock) {
		toSerialize["main_stock"] = o.MainStock
	}
	if !IsNil(o.ReserveStock) {
		toSerialize["reserve_stock"] = o.ReserveStock
	}
	if !IsNil(o.Variant) {
		toSerialize["variant"] = o.Variant
	}
	if !IsNil(o.Menu) {
		toSerialize["menu"] = o.Menu
	}
	if !IsNil(o.Preorder) {
		toSerialize["preorder"] = o.Preorder
	}
	if !IsNil(o.ExtraAttribute) {
		toSerialize["extraAttribute"] = o.ExtraAttribute
	}
	if !IsNil(o.Wholesale) {
		toSerialize["wholesale"] = o.Wholesale
	}
	if !IsNil(o.CategoryTree) {
		toSerialize["categoryTree"] = o.CategoryTree
	}
	if !IsNil(o.Pictures) {
		toSerialize["pictures"] = o.Pictures
	}
	if !IsNil(o.GMStats) {
		toSerialize["GMStats"] = o.GMStats
	}
	if !IsNil(o.Stats) {
		toSerialize["stats"] = o.Stats
	}
	if !IsNil(o.Other) {
		toSerialize["other"] = o.Other
	}
	if !IsNil(o.Campaign) {
		toSerialize["campaign"] = o.Campaign
	}
	if !IsNil(o.Volume) {
		toSerialize["volume"] = o.Volume
	}
	if !IsNil(o.Warehouses) {
		toSerialize["warehouses"] = o.Warehouses
	}
	return toSerialize, nil
}

type NullableProduct struct {
	value *Product
	isSet bool
}

func (v NullableProduct) Get() *Product {
	return v.value
}

func (v *NullableProduct) Set(val *Product) {
	v.value = val
	v.isSet = true
}

func (v NullableProduct) IsSet() bool {
	return v.isSet
}

func (v *NullableProduct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProduct(val *Product) *NullableProduct {
	return &NullableProduct{value: val, isSet: true}
}

func (v NullableProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


