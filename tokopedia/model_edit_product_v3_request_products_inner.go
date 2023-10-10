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

// checks if the EditProductV3RequestProductsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditProductV3RequestProductsInner{}

// EditProductV3RequestProductsInner struct for EditProductV3RequestProductsInner
type EditProductV3RequestProductsInner struct {
	// Product Identifier
	Id int64 `json:"id"`
	// Name of the product with length less than or equals 70 characters
	Name *string `json:"Name,omitempty"`
	// The condition of the product with the following available values NEW and USED
	Condition *string `json:"condition,omitempty"`
	// Description of the product. Maximum characters allowed is 2000
	Description *string `json:"Description,omitempty"`
	// The stock keeping unit for the product. Maximum characters allowed is 50
	Sku *string `json:"sku,omitempty"`
	// The possible value between 100 to 100.000.000. If the product variant is added, the price parameter is automatically set to the lowest price among the variant products
	Price *float64 `json:"price,omitempty"`
	// Status for the product with the following available values UNLIMITED, LIMITED, and EMPTY
	Status string `json:"status"`
	// The stock of the product. 0 indicates always available. Other than that, the possible values are from 1 to 1000. Stock should be 1 if want to add variant product
	Stock *int64 `json:"stock,omitempty"`
	// Minimum order required to purchase the product. Can only be a positive integer
	MinOrder *int64 `json:"min_order,omitempty"`
	// Unique identifier for the product’s category. To get available categories, please check Get All Categories Please input the deepest category child ID
	CategoryId *int64 `json:"category_id,omitempty"`
	Dimension *EditProductV3RequestProductsInnerDimension `json:"dimension,omitempty"`
	// Custom product logistics of the product. To get the id, please check Get Active Courier. Required field to input just ShippingProductID value responses from Get Active Courier. Remove custom_product_logistics from related product by passing empty array ([]) in this request
	CustomProductLogistics []int64 `json:"custom_product_logistics,omitempty"`
	// Product Specification (anotations) By Category ID. The value is array of annotations id that can be retrieve by hit endpoint Get Product Annotation by Category ID. The location of id is at values.id. Remove annotations from related product by passing empty array (“”) in this request
	Annotations []string `json:"annotations,omitempty"`
	// Currency code for stated price (IDR or USD)
	PriceCurrency *string `json:"price_currency,omitempty"`
	// Weight of the product
	Weight *float64 `json:"weight,omitempty"`
	// The unit of the weight with the following available value GR (gram)
	WeightUnit *string `json:"weight_unit,omitempty"`
	// Determine if the product can be returned (true) by buyers or not (false)
	IsFreeReturn *bool `json:"is_free_return,omitempty"`
	// Determine if the product must be insured (true) or not (false)
	IsMustInsurance bool `json:"is_must_insurance"`
	Menu *EditProductV3RequestProductsInnerMenu `json:"menu,omitempty"`
	// Images information of the product. The object keys includes: file_path. Remove pictures from related product by passing empty array ([]) in this request 
	Pictures []EditProductV3RequestProductsInnerPicturesInner `json:"pictures,omitempty"`
	// Wholesale price and quantity of the product. The object keys includes: min_qty and price. Remove wholesale from related product by passing empty array ([]) in this request 
	Wholesale []EditProductV3RequestProductsInnerWholesaleInner `json:"wholesale,omitempty"`
	Preorder *EditProductV3RequestProductsInnerPreorder `json:"preorder,omitempty"`
	// Video link of the product. The object keys includes: url and source. url should only contain the YouTube video id e.g. dQw4w9WgXcQ. Where the type type should be youtube. Remove videos from related product by passing empty array ([]) in this request 
	Videos []EditProductV3RequestProductsInnerVideosInner `json:"videos,omitempty"`
	Etalase *EditProductV3RequestProductsInnerEtalase `json:"etalase,omitempty"`
	Variant *EditProductV3RequestProductsInnerVariant `json:"variant,omitempty"`
}

// NewEditProductV3RequestProductsInner instantiates a new EditProductV3RequestProductsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditProductV3RequestProductsInner(id int64, status string, isMustInsurance bool) *EditProductV3RequestProductsInner {
	this := EditProductV3RequestProductsInner{}
	this.Id = id
	this.Status = status
	var weightUnit string = "GR"
	this.WeightUnit = &weightUnit
	this.IsMustInsurance = isMustInsurance
	return &this
}

// NewEditProductV3RequestProductsInnerWithDefaults instantiates a new EditProductV3RequestProductsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditProductV3RequestProductsInnerWithDefaults() *EditProductV3RequestProductsInner {
	this := EditProductV3RequestProductsInner{}
	var weightUnit string = "GR"
	this.WeightUnit = &weightUnit
	return &this
}

// GetId returns the Id field value
func (o *EditProductV3RequestProductsInner) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EditProductV3RequestProductsInner) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EditProductV3RequestProductsInner) SetId(v int64) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EditProductV3RequestProductsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditProductV3RequestProductsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EditProductV3RequestProductsInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EditProductV3RequestProductsInner) SetName(v string) {
	o.Name = &v
}

// GetCondition returns the Condition field value if set, zero value otherwise.
func (o *EditProductV3RequestProductsInner) GetCondition() string {
	if o == nil || IsNil(o.Condition) {
		var ret string
		return ret
	}
	return *o.Condition
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditProductV3RequestProductsInner) GetConditionOk() (*string, bool) {
	if o == nil || IsNil(o.Condition) {
		return nil, false
	}
	return o.Condition, true
}

// HasCondition returns a boolean if a field has been set.
func (o *EditProductV3RequestProductsInner) HasCondition() bool {
	if o != nil && !IsNil(o.Condition) {
		return true
	}

	return false
}

// SetCondition gets a reference to the given string and assigns it to the Condition field.
func (o *EditProductV3RequestProductsInner) SetCondition(v string) {
	o.Condition = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EditProductV3RequestProductsInner) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditProductV3RequestProductsInner) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EditProductV3RequestProductsInner) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EditProductV3RequestProductsInner) SetDescription(v string) {
	o.Description = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *EditProductV3RequestProductsInner) GetSku() string {
	if o == nil || IsNil(o.Sku) {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditProductV3RequestProductsInner) GetSkuOk() (*string, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *EditProductV3RequestProductsInner) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *EditProductV3RequestProductsInner) SetSku(v string) {
	o.Sku = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *EditProductV3RequestProductsInner) GetPrice() float64 {
	if o == nil || IsNil(o.Price) {
		var ret float64
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditProductV3RequestProductsInner) GetPriceOk() (*float64, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *EditProductV3RequestProductsInner) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float64 and assigns it to the Price field.
func (o *EditProductV3RequestProductsInner) SetPrice(v float64) {
	o.Price = &v
}

// GetStatus returns the Status field value
func (o *EditProductV3RequestProductsInner) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *EditProductV3RequestProductsInner) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *EditProductV3RequestProductsInner) SetStatus(v string) {
	o.Status = v
}

// GetStock returns the Stock field value if set, zero value otherwise.
func (o *EditProductV3RequestProductsInner) GetStock() int64 {
	if o == nil || IsNil(o.Stock) {
		var ret int64
		return ret
	}
	return *o.Stock
}

// GetStockOk returns a tuple with the Stock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditProductV3RequestProductsInner) GetStockOk() (*int64, bool) {
	if o == nil || IsNil(o.Stock) {
		return nil, false
	}
	return o.Stock, true
}

// HasStock returns a boolean if a field has been set.
func (o *EditProductV3RequestProductsInner) HasStock() bool {
	if o != nil && !IsNil(o.Stock) {
		return true
	}

	return false
}

// SetStock gets a reference to the given int64 and assigns it to the Stock field.
func (o *EditProductV3RequestProductsInner) SetStock(v int64) {
	o.Stock = &v
}

// GetMinOrder returns the MinOrder field value if set, zero value otherwise.
func (o *EditProductV3RequestProductsInner) GetMinOrder() int64 {
	if o == nil || IsNil(o.MinOrder) {
		var ret int64
		return ret
	}
	return *o.MinOrder
}

// GetMinOrderOk returns a tuple with the MinOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditProductV3RequestProductsInner) GetMinOrderOk() (*int64, bool) {
	if o == nil || IsNil(o.MinOrder) {
		return nil, false
	}
	return o.MinOrder, true
}

// HasMinOrder returns a boolean if a field has been set.
func (o *EditProductV3RequestProductsInner) HasMinOrder() bool {
	if o != nil && !IsNil(o.MinOrder) {
		return true
	}

	return false
}

// SetMinOrder gets a reference to the given int64 and assigns it to the MinOrder field.
func (o *EditProductV3RequestProductsInner) SetMinOrder(v int64) {
	o.MinOrder = &v
}

// GetCategoryId returns the CategoryId field value if set, zero value otherwise.
func (o *EditProductV3RequestProductsInner) GetCategoryId() int64 {
	if o == nil || IsNil(o.CategoryId) {
		var ret int64
		return ret
	}
	return *o.CategoryId
}

// GetCategoryIdOk returns a tuple with the CategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditProductV3RequestProductsInner) GetCategoryIdOk() (*int64, bool) {
	if o == nil || IsNil(o.CategoryId) {
		return nil, false
	}
	return o.CategoryId, true
}

// HasCategoryId returns a boolean if a field has been set.
func (o *EditProductV3RequestProductsInner) HasCategoryId() bool {
	if o != nil && !IsNil(o.CategoryId) {
		return true
	}

	return false
}

// SetCategoryId gets a reference to the given int64 and assigns it to the CategoryId field.
func (o *EditProductV3RequestProductsInner) SetCategoryId(v int64) {
	o.CategoryId = &v
}

// GetDimension returns the Dimension field value if set, zero value otherwise.
func (o *EditProductV3RequestProductsInner) GetDimension() EditProductV3RequestProductsInnerDimension {
	if o == nil || IsNil(o.Dimension) {
		var ret EditProductV3RequestProductsInnerDimension
		return ret
	}
	return *o.Dimension
}

// GetDimensionOk returns a tuple with the Dimension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditProductV3RequestProductsInner) GetDimensionOk() (*EditProductV3RequestProductsInnerDimension, bool) {
	if o == nil || IsNil(o.Dimension) {
		return nil, false
	}
	return o.Dimension, true
}

// HasDimension returns a boolean if a field has been set.
func (o *EditProductV3RequestProductsInner) HasDimension() bool {
	if o != nil && !IsNil(o.Dimension) {
		return true
	}

	return false
}

// SetDimension gets a reference to the given EditProductV3RequestProductsInnerDimension and assigns it to the Dimension field.
func (o *EditProductV3RequestProductsInner) SetDimension(v EditProductV3RequestProductsInnerDimension) {
	o.Dimension = &v
}

// GetCustomProductLogistics returns the CustomProductLogistics field value if set, zero value otherwise.
func (o *EditProductV3RequestProductsInner) GetCustomProductLogistics() []int64 {
	if o == nil || IsNil(o.CustomProductLogistics) {
		var ret []int64
		return ret
	}
	return o.CustomProductLogistics
}

// GetCustomProductLogisticsOk returns a tuple with the CustomProductLogistics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditProductV3RequestProductsInner) GetCustomProductLogisticsOk() ([]int64, bool) {
	if o == nil || IsNil(o.CustomProductLogistics) {
		return nil, false
	}
	return o.CustomProductLogistics, true
}

// HasCustomProductLogistics returns a boolean if a field has been set.
func (o *EditProductV3RequestProductsInner) HasCustomProductLogistics() bool {
	if o != nil && !IsNil(o.CustomProductLogistics) {
		return true
	}

	return false
}

// SetCustomProductLogistics gets a reference to the given []int64 and assigns it to the CustomProductLogistics field.
func (o *EditProductV3RequestProductsInner) SetCustomProductLogistics(v []int64) {
	o.CustomProductLogistics = v
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise.
func (o *EditProductV3RequestProductsInner) GetAnnotations() []string {
	if o == nil || IsNil(o.Annotations) {
		var ret []string
		return ret
	}
	return o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditProductV3RequestProductsInner) GetAnnotationsOk() ([]string, bool) {
	if o == nil || IsNil(o.Annotations) {
		return nil, false
	}
	return o.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *EditProductV3RequestProductsInner) HasAnnotations() bool {
	if o != nil && !IsNil(o.Annotations) {
		return true
	}

	return false
}

// SetAnnotations gets a reference to the given []string and assigns it to the Annotations field.
func (o *EditProductV3RequestProductsInner) SetAnnotations(v []string) {
	o.Annotations = v
}

// GetPriceCurrency returns the PriceCurrency field value if set, zero value otherwise.
func (o *EditProductV3RequestProductsInner) GetPriceCurrency() string {
	if o == nil || IsNil(o.PriceCurrency) {
		var ret string
		return ret
	}
	return *o.PriceCurrency
}

// GetPriceCurrencyOk returns a tuple with the PriceCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditProductV3RequestProductsInner) GetPriceCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.PriceCurrency) {
		return nil, false
	}
	return o.PriceCurrency, true
}

// HasPriceCurrency returns a boolean if a field has been set.
func (o *EditProductV3RequestProductsInner) HasPriceCurrency() bool {
	if o != nil && !IsNil(o.PriceCurrency) {
		return true
	}

	return false
}

// SetPriceCurrency gets a reference to the given string and assigns it to the PriceCurrency field.
func (o *EditProductV3RequestProductsInner) SetPriceCurrency(v string) {
	o.PriceCurrency = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *EditProductV3RequestProductsInner) GetWeight() float64 {
	if o == nil || IsNil(o.Weight) {
		var ret float64
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditProductV3RequestProductsInner) GetWeightOk() (*float64, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *EditProductV3RequestProductsInner) HasWeight() bool {
	if o != nil && !IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given float64 and assigns it to the Weight field.
func (o *EditProductV3RequestProductsInner) SetWeight(v float64) {
	o.Weight = &v
}

// GetWeightUnit returns the WeightUnit field value if set, zero value otherwise.
func (o *EditProductV3RequestProductsInner) GetWeightUnit() string {
	if o == nil || IsNil(o.WeightUnit) {
		var ret string
		return ret
	}
	return *o.WeightUnit
}

// GetWeightUnitOk returns a tuple with the WeightUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditProductV3RequestProductsInner) GetWeightUnitOk() (*string, bool) {
	if o == nil || IsNil(o.WeightUnit) {
		return nil, false
	}
	return o.WeightUnit, true
}

// HasWeightUnit returns a boolean if a field has been set.
func (o *EditProductV3RequestProductsInner) HasWeightUnit() bool {
	if o != nil && !IsNil(o.WeightUnit) {
		return true
	}

	return false
}

// SetWeightUnit gets a reference to the given string and assigns it to the WeightUnit field.
func (o *EditProductV3RequestProductsInner) SetWeightUnit(v string) {
	o.WeightUnit = &v
}

// GetIsFreeReturn returns the IsFreeReturn field value if set, zero value otherwise.
func (o *EditProductV3RequestProductsInner) GetIsFreeReturn() bool {
	if o == nil || IsNil(o.IsFreeReturn) {
		var ret bool
		return ret
	}
	return *o.IsFreeReturn
}

// GetIsFreeReturnOk returns a tuple with the IsFreeReturn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditProductV3RequestProductsInner) GetIsFreeReturnOk() (*bool, bool) {
	if o == nil || IsNil(o.IsFreeReturn) {
		return nil, false
	}
	return o.IsFreeReturn, true
}

// HasIsFreeReturn returns a boolean if a field has been set.
func (o *EditProductV3RequestProductsInner) HasIsFreeReturn() bool {
	if o != nil && !IsNil(o.IsFreeReturn) {
		return true
	}

	return false
}

// SetIsFreeReturn gets a reference to the given bool and assigns it to the IsFreeReturn field.
func (o *EditProductV3RequestProductsInner) SetIsFreeReturn(v bool) {
	o.IsFreeReturn = &v
}

// GetIsMustInsurance returns the IsMustInsurance field value
func (o *EditProductV3RequestProductsInner) GetIsMustInsurance() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsMustInsurance
}

// GetIsMustInsuranceOk returns a tuple with the IsMustInsurance field value
// and a boolean to check if the value has been set.
func (o *EditProductV3RequestProductsInner) GetIsMustInsuranceOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsMustInsurance, true
}

// SetIsMustInsurance sets field value
func (o *EditProductV3RequestProductsInner) SetIsMustInsurance(v bool) {
	o.IsMustInsurance = v
}

// GetMenu returns the Menu field value if set, zero value otherwise.
func (o *EditProductV3RequestProductsInner) GetMenu() EditProductV3RequestProductsInnerMenu {
	if o == nil || IsNil(o.Menu) {
		var ret EditProductV3RequestProductsInnerMenu
		return ret
	}
	return *o.Menu
}

// GetMenuOk returns a tuple with the Menu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditProductV3RequestProductsInner) GetMenuOk() (*EditProductV3RequestProductsInnerMenu, bool) {
	if o == nil || IsNil(o.Menu) {
		return nil, false
	}
	return o.Menu, true
}

// HasMenu returns a boolean if a field has been set.
func (o *EditProductV3RequestProductsInner) HasMenu() bool {
	if o != nil && !IsNil(o.Menu) {
		return true
	}

	return false
}

// SetMenu gets a reference to the given EditProductV3RequestProductsInnerMenu and assigns it to the Menu field.
func (o *EditProductV3RequestProductsInner) SetMenu(v EditProductV3RequestProductsInnerMenu) {
	o.Menu = &v
}

// GetPictures returns the Pictures field value if set, zero value otherwise.
func (o *EditProductV3RequestProductsInner) GetPictures() []EditProductV3RequestProductsInnerPicturesInner {
	if o == nil || IsNil(o.Pictures) {
		var ret []EditProductV3RequestProductsInnerPicturesInner
		return ret
	}
	return o.Pictures
}

// GetPicturesOk returns a tuple with the Pictures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditProductV3RequestProductsInner) GetPicturesOk() ([]EditProductV3RequestProductsInnerPicturesInner, bool) {
	if o == nil || IsNil(o.Pictures) {
		return nil, false
	}
	return o.Pictures, true
}

// HasPictures returns a boolean if a field has been set.
func (o *EditProductV3RequestProductsInner) HasPictures() bool {
	if o != nil && !IsNil(o.Pictures) {
		return true
	}

	return false
}

// SetPictures gets a reference to the given []EditProductV3RequestProductsInnerPicturesInner and assigns it to the Pictures field.
func (o *EditProductV3RequestProductsInner) SetPictures(v []EditProductV3RequestProductsInnerPicturesInner) {
	o.Pictures = v
}

// GetWholesale returns the Wholesale field value if set, zero value otherwise.
func (o *EditProductV3RequestProductsInner) GetWholesale() []EditProductV3RequestProductsInnerWholesaleInner {
	if o == nil || IsNil(o.Wholesale) {
		var ret []EditProductV3RequestProductsInnerWholesaleInner
		return ret
	}
	return o.Wholesale
}

// GetWholesaleOk returns a tuple with the Wholesale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditProductV3RequestProductsInner) GetWholesaleOk() ([]EditProductV3RequestProductsInnerWholesaleInner, bool) {
	if o == nil || IsNil(o.Wholesale) {
		return nil, false
	}
	return o.Wholesale, true
}

// HasWholesale returns a boolean if a field has been set.
func (o *EditProductV3RequestProductsInner) HasWholesale() bool {
	if o != nil && !IsNil(o.Wholesale) {
		return true
	}

	return false
}

// SetWholesale gets a reference to the given []EditProductV3RequestProductsInnerWholesaleInner and assigns it to the Wholesale field.
func (o *EditProductV3RequestProductsInner) SetWholesale(v []EditProductV3RequestProductsInnerWholesaleInner) {
	o.Wholesale = v
}

// GetPreorder returns the Preorder field value if set, zero value otherwise.
func (o *EditProductV3RequestProductsInner) GetPreorder() EditProductV3RequestProductsInnerPreorder {
	if o == nil || IsNil(o.Preorder) {
		var ret EditProductV3RequestProductsInnerPreorder
		return ret
	}
	return *o.Preorder
}

// GetPreorderOk returns a tuple with the Preorder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditProductV3RequestProductsInner) GetPreorderOk() (*EditProductV3RequestProductsInnerPreorder, bool) {
	if o == nil || IsNil(o.Preorder) {
		return nil, false
	}
	return o.Preorder, true
}

// HasPreorder returns a boolean if a field has been set.
func (o *EditProductV3RequestProductsInner) HasPreorder() bool {
	if o != nil && !IsNil(o.Preorder) {
		return true
	}

	return false
}

// SetPreorder gets a reference to the given EditProductV3RequestProductsInnerPreorder and assigns it to the Preorder field.
func (o *EditProductV3RequestProductsInner) SetPreorder(v EditProductV3RequestProductsInnerPreorder) {
	o.Preorder = &v
}

// GetVideos returns the Videos field value if set, zero value otherwise.
func (o *EditProductV3RequestProductsInner) GetVideos() []EditProductV3RequestProductsInnerVideosInner {
	if o == nil || IsNil(o.Videos) {
		var ret []EditProductV3RequestProductsInnerVideosInner
		return ret
	}
	return o.Videos
}

// GetVideosOk returns a tuple with the Videos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditProductV3RequestProductsInner) GetVideosOk() ([]EditProductV3RequestProductsInnerVideosInner, bool) {
	if o == nil || IsNil(o.Videos) {
		return nil, false
	}
	return o.Videos, true
}

// HasVideos returns a boolean if a field has been set.
func (o *EditProductV3RequestProductsInner) HasVideos() bool {
	if o != nil && !IsNil(o.Videos) {
		return true
	}

	return false
}

// SetVideos gets a reference to the given []EditProductV3RequestProductsInnerVideosInner and assigns it to the Videos field.
func (o *EditProductV3RequestProductsInner) SetVideos(v []EditProductV3RequestProductsInnerVideosInner) {
	o.Videos = v
}

// GetEtalase returns the Etalase field value if set, zero value otherwise.
func (o *EditProductV3RequestProductsInner) GetEtalase() EditProductV3RequestProductsInnerEtalase {
	if o == nil || IsNil(o.Etalase) {
		var ret EditProductV3RequestProductsInnerEtalase
		return ret
	}
	return *o.Etalase
}

// GetEtalaseOk returns a tuple with the Etalase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditProductV3RequestProductsInner) GetEtalaseOk() (*EditProductV3RequestProductsInnerEtalase, bool) {
	if o == nil || IsNil(o.Etalase) {
		return nil, false
	}
	return o.Etalase, true
}

// HasEtalase returns a boolean if a field has been set.
func (o *EditProductV3RequestProductsInner) HasEtalase() bool {
	if o != nil && !IsNil(o.Etalase) {
		return true
	}

	return false
}

// SetEtalase gets a reference to the given EditProductV3RequestProductsInnerEtalase and assigns it to the Etalase field.
func (o *EditProductV3RequestProductsInner) SetEtalase(v EditProductV3RequestProductsInnerEtalase) {
	o.Etalase = &v
}

// GetVariant returns the Variant field value if set, zero value otherwise.
func (o *EditProductV3RequestProductsInner) GetVariant() EditProductV3RequestProductsInnerVariant {
	if o == nil || IsNil(o.Variant) {
		var ret EditProductV3RequestProductsInnerVariant
		return ret
	}
	return *o.Variant
}

// GetVariantOk returns a tuple with the Variant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditProductV3RequestProductsInner) GetVariantOk() (*EditProductV3RequestProductsInnerVariant, bool) {
	if o == nil || IsNil(o.Variant) {
		return nil, false
	}
	return o.Variant, true
}

// HasVariant returns a boolean if a field has been set.
func (o *EditProductV3RequestProductsInner) HasVariant() bool {
	if o != nil && !IsNil(o.Variant) {
		return true
	}

	return false
}

// SetVariant gets a reference to the given EditProductV3RequestProductsInnerVariant and assigns it to the Variant field.
func (o *EditProductV3RequestProductsInner) SetVariant(v EditProductV3RequestProductsInnerVariant) {
	o.Variant = &v
}

func (o EditProductV3RequestProductsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditProductV3RequestProductsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.Condition) {
		toSerialize["condition"] = o.Condition
	}
	if !IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if !IsNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	toSerialize["status"] = o.Status
	if !IsNil(o.Stock) {
		toSerialize["stock"] = o.Stock
	}
	if !IsNil(o.MinOrder) {
		toSerialize["min_order"] = o.MinOrder
	}
	if !IsNil(o.CategoryId) {
		toSerialize["category_id"] = o.CategoryId
	}
	if !IsNil(o.Dimension) {
		toSerialize["dimension"] = o.Dimension
	}
	if !IsNil(o.CustomProductLogistics) {
		toSerialize["custom_product_logistics"] = o.CustomProductLogistics
	}
	if !IsNil(o.Annotations) {
		toSerialize["annotations"] = o.Annotations
	}
	if !IsNil(o.PriceCurrency) {
		toSerialize["price_currency"] = o.PriceCurrency
	}
	if !IsNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	if !IsNil(o.WeightUnit) {
		toSerialize["weight_unit"] = o.WeightUnit
	}
	if !IsNil(o.IsFreeReturn) {
		toSerialize["is_free_return"] = o.IsFreeReturn
	}
	toSerialize["is_must_insurance"] = o.IsMustInsurance
	if !IsNil(o.Menu) {
		toSerialize["menu"] = o.Menu
	}
	if !IsNil(o.Pictures) {
		toSerialize["pictures"] = o.Pictures
	}
	if !IsNil(o.Wholesale) {
		toSerialize["wholesale"] = o.Wholesale
	}
	if !IsNil(o.Preorder) {
		toSerialize["preorder"] = o.Preorder
	}
	if !IsNil(o.Videos) {
		toSerialize["videos"] = o.Videos
	}
	if !IsNil(o.Etalase) {
		toSerialize["etalase"] = o.Etalase
	}
	if !IsNil(o.Variant) {
		toSerialize["variant"] = o.Variant
	}
	return toSerialize, nil
}

type NullableEditProductV3RequestProductsInner struct {
	value *EditProductV3RequestProductsInner
	isSet bool
}

func (v NullableEditProductV3RequestProductsInner) Get() *EditProductV3RequestProductsInner {
	return v.value
}

func (v *NullableEditProductV3RequestProductsInner) Set(val *EditProductV3RequestProductsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableEditProductV3RequestProductsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableEditProductV3RequestProductsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditProductV3RequestProductsInner(val *EditProductV3RequestProductsInner) *NullableEditProductV3RequestProductsInner {
	return &NullableEditProductV3RequestProductsInner{value: val, isSet: true}
}

func (v NullableEditProductV3RequestProductsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditProductV3RequestProductsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


