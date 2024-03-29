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

// checks if the CreateProductV2RequestProductsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateProductV2RequestProductsInner{}

// CreateProductV2RequestProductsInner struct for CreateProductV2RequestProductsInner
type CreateProductV2RequestProductsInner struct {
	// Name of the product with length less than or equals 70 characters
	Name string `json:"Name"`
	// The condition of the product with the following available values NEW and USED
	Condition string `json:"condition"`
	// Description of the product. Maximum characters allowed is 2000
	Description *string `json:"Description,omitempty"`
	// The stock keeping unit for the product. Maximum characters allowed is 50
	Sku *string `json:"sku,omitempty"`
	// The possible value between 100 to 100.000.000. If the product variant is added, the price parameter is automatically set to the lowest price among the variant products.
	Price float64 `json:"price"`
	// Status for the product with the following available values UNLIMITED, LIMITED, and EMPTY.
	Status string `json:"status"`
	// The stock of the product. 0 indicates always available. Other than that, the possible values are from 1 to 1000. Stock should be 1 if want to add variant product
	Stock *int64 `json:"stock,omitempty"`
	// Minimum order required to purchase the product. Can only be a positive integer
	MinOrder int64 `json:"min_order"`
	// Unique identifier for the product’s category. To get available categories, please check Get All Categories. Please input the deepest category child ID
	CategoryId int64 `json:"category_id"`
	// Currency code for stated price (IDR or USD)
	PriceCurrency string `json:"price_currency"`
	// Weight of the product
	Weight float64 `json:"weight"`
	// The unit of the weight with the following available value GR (gram)
	WeightUnit string `json:"weight_unit"`
	// Determine if the product can be returned (true) by buyers or not (false)
	IsFreeReturn *bool `json:"is_free_return,omitempty"`
	// Determine if the product must be insured (true) or not (false)
	IsMustInsurance bool `json:"is_must_insurance"`
	Etalase CreateProductV2RequestProductsInnerEtalase `json:"etalase"`
	// Images information of the product. The object keys includes: file_path
	Pictures []EditProductV3RequestProductsInnerPicturesInner `json:"pictures"`
	// Wholesale price and quantity of the product. The object keys includes: min_qty and price
	Wholesale []EditProductV3RequestProductsInnerWholesaleInner `json:"wholesale,omitempty"`
	Preorder *CreateProductV2RequestProductsInnerPreorder `json:"preorder,omitempty"`
	// Video link of the product. The object keys includes: url and source. url should only contain the YouTube video id e.g. dQw4w9WgXcQ. Where the type should be youtube
	Videos []CreateProductV2RequestProductsInnerVideosInner `json:"videos,omitempty"`
	// Variant of the product. The object keys includes: variant and product_variant
	Variant []CreateProductV2RequestProductsInnerVariantInner `json:"variant,omitempty"`
}

// NewCreateProductV2RequestProductsInner instantiates a new CreateProductV2RequestProductsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateProductV2RequestProductsInner(name string, condition string, price float64, status string, minOrder int64, categoryId int64, priceCurrency string, weight float64, weightUnit string, isMustInsurance bool, etalase CreateProductV2RequestProductsInnerEtalase, pictures []EditProductV3RequestProductsInnerPicturesInner) *CreateProductV2RequestProductsInner {
	this := CreateProductV2RequestProductsInner{}
	this.Name = name
	this.Condition = condition
	this.Price = price
	this.Status = status
	this.MinOrder = minOrder
	this.CategoryId = categoryId
	this.PriceCurrency = priceCurrency
	this.Weight = weight
	this.WeightUnit = weightUnit
	this.IsMustInsurance = isMustInsurance
	this.Etalase = etalase
	this.Pictures = pictures
	return &this
}

// NewCreateProductV2RequestProductsInnerWithDefaults instantiates a new CreateProductV2RequestProductsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateProductV2RequestProductsInnerWithDefaults() *CreateProductV2RequestProductsInner {
	this := CreateProductV2RequestProductsInner{}
	var condition string = "NEW"
	this.Condition = condition
	var status string = "LIMITED"
	this.Status = status
	var priceCurrency string = "IDR"
	this.PriceCurrency = priceCurrency
	var weightUnit string = "GR"
	this.WeightUnit = weightUnit
	return &this
}

// GetName returns the Name field value
func (o *CreateProductV2RequestProductsInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateProductV2RequestProductsInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateProductV2RequestProductsInner) SetName(v string) {
	o.Name = v
}

// GetCondition returns the Condition field value
func (o *CreateProductV2RequestProductsInner) GetCondition() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Condition
}

// GetConditionOk returns a tuple with the Condition field value
// and a boolean to check if the value has been set.
func (o *CreateProductV2RequestProductsInner) GetConditionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Condition, true
}

// SetCondition sets field value
func (o *CreateProductV2RequestProductsInner) SetCondition(v string) {
	o.Condition = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateProductV2RequestProductsInner) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProductV2RequestProductsInner) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateProductV2RequestProductsInner) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateProductV2RequestProductsInner) SetDescription(v string) {
	o.Description = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *CreateProductV2RequestProductsInner) GetSku() string {
	if o == nil || IsNil(o.Sku) {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProductV2RequestProductsInner) GetSkuOk() (*string, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *CreateProductV2RequestProductsInner) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *CreateProductV2RequestProductsInner) SetSku(v string) {
	o.Sku = &v
}

// GetPrice returns the Price field value
func (o *CreateProductV2RequestProductsInner) GetPrice() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *CreateProductV2RequestProductsInner) GetPriceOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Price, true
}

// SetPrice sets field value
func (o *CreateProductV2RequestProductsInner) SetPrice(v float64) {
	o.Price = v
}

// GetStatus returns the Status field value
func (o *CreateProductV2RequestProductsInner) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CreateProductV2RequestProductsInner) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *CreateProductV2RequestProductsInner) SetStatus(v string) {
	o.Status = v
}

// GetStock returns the Stock field value if set, zero value otherwise.
func (o *CreateProductV2RequestProductsInner) GetStock() int64 {
	if o == nil || IsNil(o.Stock) {
		var ret int64
		return ret
	}
	return *o.Stock
}

// GetStockOk returns a tuple with the Stock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProductV2RequestProductsInner) GetStockOk() (*int64, bool) {
	if o == nil || IsNil(o.Stock) {
		return nil, false
	}
	return o.Stock, true
}

// HasStock returns a boolean if a field has been set.
func (o *CreateProductV2RequestProductsInner) HasStock() bool {
	if o != nil && !IsNil(o.Stock) {
		return true
	}

	return false
}

// SetStock gets a reference to the given int64 and assigns it to the Stock field.
func (o *CreateProductV2RequestProductsInner) SetStock(v int64) {
	o.Stock = &v
}

// GetMinOrder returns the MinOrder field value
func (o *CreateProductV2RequestProductsInner) GetMinOrder() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.MinOrder
}

// GetMinOrderOk returns a tuple with the MinOrder field value
// and a boolean to check if the value has been set.
func (o *CreateProductV2RequestProductsInner) GetMinOrderOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinOrder, true
}

// SetMinOrder sets field value
func (o *CreateProductV2RequestProductsInner) SetMinOrder(v int64) {
	o.MinOrder = v
}

// GetCategoryId returns the CategoryId field value
func (o *CreateProductV2RequestProductsInner) GetCategoryId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CategoryId
}

// GetCategoryIdOk returns a tuple with the CategoryId field value
// and a boolean to check if the value has been set.
func (o *CreateProductV2RequestProductsInner) GetCategoryIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CategoryId, true
}

// SetCategoryId sets field value
func (o *CreateProductV2RequestProductsInner) SetCategoryId(v int64) {
	o.CategoryId = v
}

// GetPriceCurrency returns the PriceCurrency field value
func (o *CreateProductV2RequestProductsInner) GetPriceCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PriceCurrency
}

// GetPriceCurrencyOk returns a tuple with the PriceCurrency field value
// and a boolean to check if the value has been set.
func (o *CreateProductV2RequestProductsInner) GetPriceCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PriceCurrency, true
}

// SetPriceCurrency sets field value
func (o *CreateProductV2RequestProductsInner) SetPriceCurrency(v string) {
	o.PriceCurrency = v
}

// GetWeight returns the Weight field value
func (o *CreateProductV2RequestProductsInner) GetWeight() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Weight
}

// GetWeightOk returns a tuple with the Weight field value
// and a boolean to check if the value has been set.
func (o *CreateProductV2RequestProductsInner) GetWeightOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Weight, true
}

// SetWeight sets field value
func (o *CreateProductV2RequestProductsInner) SetWeight(v float64) {
	o.Weight = v
}

// GetWeightUnit returns the WeightUnit field value
func (o *CreateProductV2RequestProductsInner) GetWeightUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WeightUnit
}

// GetWeightUnitOk returns a tuple with the WeightUnit field value
// and a boolean to check if the value has been set.
func (o *CreateProductV2RequestProductsInner) GetWeightUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WeightUnit, true
}

// SetWeightUnit sets field value
func (o *CreateProductV2RequestProductsInner) SetWeightUnit(v string) {
	o.WeightUnit = v
}

// GetIsFreeReturn returns the IsFreeReturn field value if set, zero value otherwise.
func (o *CreateProductV2RequestProductsInner) GetIsFreeReturn() bool {
	if o == nil || IsNil(o.IsFreeReturn) {
		var ret bool
		return ret
	}
	return *o.IsFreeReturn
}

// GetIsFreeReturnOk returns a tuple with the IsFreeReturn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProductV2RequestProductsInner) GetIsFreeReturnOk() (*bool, bool) {
	if o == nil || IsNil(o.IsFreeReturn) {
		return nil, false
	}
	return o.IsFreeReturn, true
}

// HasIsFreeReturn returns a boolean if a field has been set.
func (o *CreateProductV2RequestProductsInner) HasIsFreeReturn() bool {
	if o != nil && !IsNil(o.IsFreeReturn) {
		return true
	}

	return false
}

// SetIsFreeReturn gets a reference to the given bool and assigns it to the IsFreeReturn field.
func (o *CreateProductV2RequestProductsInner) SetIsFreeReturn(v bool) {
	o.IsFreeReturn = &v
}

// GetIsMustInsurance returns the IsMustInsurance field value
func (o *CreateProductV2RequestProductsInner) GetIsMustInsurance() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsMustInsurance
}

// GetIsMustInsuranceOk returns a tuple with the IsMustInsurance field value
// and a boolean to check if the value has been set.
func (o *CreateProductV2RequestProductsInner) GetIsMustInsuranceOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsMustInsurance, true
}

// SetIsMustInsurance sets field value
func (o *CreateProductV2RequestProductsInner) SetIsMustInsurance(v bool) {
	o.IsMustInsurance = v
}

// GetEtalase returns the Etalase field value
func (o *CreateProductV2RequestProductsInner) GetEtalase() CreateProductV2RequestProductsInnerEtalase {
	if o == nil {
		var ret CreateProductV2RequestProductsInnerEtalase
		return ret
	}

	return o.Etalase
}

// GetEtalaseOk returns a tuple with the Etalase field value
// and a boolean to check if the value has been set.
func (o *CreateProductV2RequestProductsInner) GetEtalaseOk() (*CreateProductV2RequestProductsInnerEtalase, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Etalase, true
}

// SetEtalase sets field value
func (o *CreateProductV2RequestProductsInner) SetEtalase(v CreateProductV2RequestProductsInnerEtalase) {
	o.Etalase = v
}

// GetPictures returns the Pictures field value
func (o *CreateProductV2RequestProductsInner) GetPictures() []EditProductV3RequestProductsInnerPicturesInner {
	if o == nil {
		var ret []EditProductV3RequestProductsInnerPicturesInner
		return ret
	}

	return o.Pictures
}

// GetPicturesOk returns a tuple with the Pictures field value
// and a boolean to check if the value has been set.
func (o *CreateProductV2RequestProductsInner) GetPicturesOk() ([]EditProductV3RequestProductsInnerPicturesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Pictures, true
}

// SetPictures sets field value
func (o *CreateProductV2RequestProductsInner) SetPictures(v []EditProductV3RequestProductsInnerPicturesInner) {
	o.Pictures = v
}

// GetWholesale returns the Wholesale field value if set, zero value otherwise.
func (o *CreateProductV2RequestProductsInner) GetWholesale() []EditProductV3RequestProductsInnerWholesaleInner {
	if o == nil || IsNil(o.Wholesale) {
		var ret []EditProductV3RequestProductsInnerWholesaleInner
		return ret
	}
	return o.Wholesale
}

// GetWholesaleOk returns a tuple with the Wholesale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProductV2RequestProductsInner) GetWholesaleOk() ([]EditProductV3RequestProductsInnerWholesaleInner, bool) {
	if o == nil || IsNil(o.Wholesale) {
		return nil, false
	}
	return o.Wholesale, true
}

// HasWholesale returns a boolean if a field has been set.
func (o *CreateProductV2RequestProductsInner) HasWholesale() bool {
	if o != nil && !IsNil(o.Wholesale) {
		return true
	}

	return false
}

// SetWholesale gets a reference to the given []EditProductV3RequestProductsInnerWholesaleInner and assigns it to the Wholesale field.
func (o *CreateProductV2RequestProductsInner) SetWholesale(v []EditProductV3RequestProductsInnerWholesaleInner) {
	o.Wholesale = v
}

// GetPreorder returns the Preorder field value if set, zero value otherwise.
func (o *CreateProductV2RequestProductsInner) GetPreorder() CreateProductV2RequestProductsInnerPreorder {
	if o == nil || IsNil(o.Preorder) {
		var ret CreateProductV2RequestProductsInnerPreorder
		return ret
	}
	return *o.Preorder
}

// GetPreorderOk returns a tuple with the Preorder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProductV2RequestProductsInner) GetPreorderOk() (*CreateProductV2RequestProductsInnerPreorder, bool) {
	if o == nil || IsNil(o.Preorder) {
		return nil, false
	}
	return o.Preorder, true
}

// HasPreorder returns a boolean if a field has been set.
func (o *CreateProductV2RequestProductsInner) HasPreorder() bool {
	if o != nil && !IsNil(o.Preorder) {
		return true
	}

	return false
}

// SetPreorder gets a reference to the given CreateProductV2RequestProductsInnerPreorder and assigns it to the Preorder field.
func (o *CreateProductV2RequestProductsInner) SetPreorder(v CreateProductV2RequestProductsInnerPreorder) {
	o.Preorder = &v
}

// GetVideos returns the Videos field value if set, zero value otherwise.
func (o *CreateProductV2RequestProductsInner) GetVideos() []CreateProductV2RequestProductsInnerVideosInner {
	if o == nil || IsNil(o.Videos) {
		var ret []CreateProductV2RequestProductsInnerVideosInner
		return ret
	}
	return o.Videos
}

// GetVideosOk returns a tuple with the Videos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProductV2RequestProductsInner) GetVideosOk() ([]CreateProductV2RequestProductsInnerVideosInner, bool) {
	if o == nil || IsNil(o.Videos) {
		return nil, false
	}
	return o.Videos, true
}

// HasVideos returns a boolean if a field has been set.
func (o *CreateProductV2RequestProductsInner) HasVideos() bool {
	if o != nil && !IsNil(o.Videos) {
		return true
	}

	return false
}

// SetVideos gets a reference to the given []CreateProductV2RequestProductsInnerVideosInner and assigns it to the Videos field.
func (o *CreateProductV2RequestProductsInner) SetVideos(v []CreateProductV2RequestProductsInnerVideosInner) {
	o.Videos = v
}

// GetVariant returns the Variant field value if set, zero value otherwise.
func (o *CreateProductV2RequestProductsInner) GetVariant() []CreateProductV2RequestProductsInnerVariantInner {
	if o == nil || IsNil(o.Variant) {
		var ret []CreateProductV2RequestProductsInnerVariantInner
		return ret
	}
	return o.Variant
}

// GetVariantOk returns a tuple with the Variant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProductV2RequestProductsInner) GetVariantOk() ([]CreateProductV2RequestProductsInnerVariantInner, bool) {
	if o == nil || IsNil(o.Variant) {
		return nil, false
	}
	return o.Variant, true
}

// HasVariant returns a boolean if a field has been set.
func (o *CreateProductV2RequestProductsInner) HasVariant() bool {
	if o != nil && !IsNil(o.Variant) {
		return true
	}

	return false
}

// SetVariant gets a reference to the given []CreateProductV2RequestProductsInnerVariantInner and assigns it to the Variant field.
func (o *CreateProductV2RequestProductsInner) SetVariant(v []CreateProductV2RequestProductsInnerVariantInner) {
	o.Variant = v
}

func (o CreateProductV2RequestProductsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateProductV2RequestProductsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Name"] = o.Name
	toSerialize["condition"] = o.Condition
	if !IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if !IsNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	toSerialize["price"] = o.Price
	toSerialize["status"] = o.Status
	if !IsNil(o.Stock) {
		toSerialize["stock"] = o.Stock
	}
	toSerialize["min_order"] = o.MinOrder
	toSerialize["category_id"] = o.CategoryId
	toSerialize["price_currency"] = o.PriceCurrency
	toSerialize["weight"] = o.Weight
	toSerialize["weight_unit"] = o.WeightUnit
	if !IsNil(o.IsFreeReturn) {
		toSerialize["is_free_return"] = o.IsFreeReturn
	}
	toSerialize["is_must_insurance"] = o.IsMustInsurance
	toSerialize["etalase"] = o.Etalase
	toSerialize["pictures"] = o.Pictures
	if !IsNil(o.Wholesale) {
		toSerialize["wholesale"] = o.Wholesale
	}
	if !IsNil(o.Preorder) {
		toSerialize["preorder"] = o.Preorder
	}
	if !IsNil(o.Videos) {
		toSerialize["videos"] = o.Videos
	}
	if !IsNil(o.Variant) {
		toSerialize["variant"] = o.Variant
	}
	return toSerialize, nil
}

type NullableCreateProductV2RequestProductsInner struct {
	value *CreateProductV2RequestProductsInner
	isSet bool
}

func (v NullableCreateProductV2RequestProductsInner) Get() *CreateProductV2RequestProductsInner {
	return v.value
}

func (v *NullableCreateProductV2RequestProductsInner) Set(val *CreateProductV2RequestProductsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateProductV2RequestProductsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateProductV2RequestProductsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateProductV2RequestProductsInner(val *CreateProductV2RequestProductsInner) *NullableCreateProductV2RequestProductsInner {
	return &NullableCreateProductV2RequestProductsInner{value: val, isSet: true}
}

func (v NullableCreateProductV2RequestProductsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateProductV2RequestProductsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


