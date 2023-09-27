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

// checks if the ViewCampaignProducts200ResponseDataProductsInnerCampaign type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ViewCampaignProducts200ResponseDataProductsInnerCampaign{}

// ViewCampaignProducts200ResponseDataProductsInnerCampaign struct for ViewCampaignProducts200ResponseDataProductsInnerCampaign
type ViewCampaignProducts200ResponseDataProductsInnerCampaign struct {
	// Campaign Unique Identifier
	CampaignID *int64 `json:"campaignID,omitempty"`
	// Slash Price Product Unique Identifier
	SlashPriceProductId *int64 `json:"slash_price_product_id,omitempty"`
	// Product Discount Percentage
	DiscountPercentage *int64 `json:"discount_percentage,omitempty"`
	// Discounted Product Price
	DiscountedPrice *string `json:"discounted_price,omitempty"`
	OriginalPrice *string `json:"original_price,omitempty"`
	CustomStock *int64 `json:"custom_stock,omitempty"`
	// Campaign Status
	CampaignStatus *int64 `json:"campaign_status,omitempty"`
	// Campaign Start Date (format: YYYY-MM-DD HH-mm-ss) 
	StartDate *string `json:"start_date,omitempty"`
	// Campaign End Date (format: YYYY-MM-DD HH-mm-ss) 
	EndDate *string `json:"end_date,omitempty"`
	// Campaign Type Name
	CampaignTypeName *string `json:"campaign_type_name,omitempty"`
	// Campaign Type Short Name
	CampaignShortName *string `json:"campaign_short_name,omitempty"`
	// Product Maximum Order
	MaxOrder *int64 `json:"max_order,omitempty"`
	// Discounted Price Formated
	DiscountedPriceFmt *string `json:"discounted_price_fmt,omitempty"`
	// Original Price Formated
	OriginalPriceFmt *string `json:"original_price_fmt,omitempty"`
	// Product Original Stock
	OriginalStock *int64 `json:"original_stock,omitempty"`
	// Product Minimum Order
	MinOrder *int64 `json:"min_order,omitempty"`
	AdditionalCartInfo *ViewCampaignProducts200ResponseDataProductsInnerCampaignAdditionalCartInfo `json:"additional_cart_info,omitempty"`
	// Stock Wording
	StockWording map[string]interface{} `json:"stock_wording,omitempty"`
}

// NewViewCampaignProducts200ResponseDataProductsInnerCampaign instantiates a new ViewCampaignProducts200ResponseDataProductsInnerCampaign object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViewCampaignProducts200ResponseDataProductsInnerCampaign() *ViewCampaignProducts200ResponseDataProductsInnerCampaign {
	this := ViewCampaignProducts200ResponseDataProductsInnerCampaign{}
	return &this
}

// NewViewCampaignProducts200ResponseDataProductsInnerCampaignWithDefaults instantiates a new ViewCampaignProducts200ResponseDataProductsInnerCampaign object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViewCampaignProducts200ResponseDataProductsInnerCampaignWithDefaults() *ViewCampaignProducts200ResponseDataProductsInnerCampaign {
	this := ViewCampaignProducts200ResponseDataProductsInnerCampaign{}
	return &this
}

// GetCampaignID returns the CampaignID field value if set, zero value otherwise.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetCampaignID() int64 {
	if o == nil || IsNil(o.CampaignID) {
		var ret int64
		return ret
	}
	return *o.CampaignID
}

// GetCampaignIDOk returns a tuple with the CampaignID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetCampaignIDOk() (*int64, bool) {
	if o == nil || IsNil(o.CampaignID) {
		return nil, false
	}
	return o.CampaignID, true
}

// HasCampaignID returns a boolean if a field has been set.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) HasCampaignID() bool {
	if o != nil && !IsNil(o.CampaignID) {
		return true
	}

	return false
}

// SetCampaignID gets a reference to the given int64 and assigns it to the CampaignID field.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) SetCampaignID(v int64) {
	o.CampaignID = &v
}

// GetSlashPriceProductId returns the SlashPriceProductId field value if set, zero value otherwise.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetSlashPriceProductId() int64 {
	if o == nil || IsNil(o.SlashPriceProductId) {
		var ret int64
		return ret
	}
	return *o.SlashPriceProductId
}

// GetSlashPriceProductIdOk returns a tuple with the SlashPriceProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetSlashPriceProductIdOk() (*int64, bool) {
	if o == nil || IsNil(o.SlashPriceProductId) {
		return nil, false
	}
	return o.SlashPriceProductId, true
}

// HasSlashPriceProductId returns a boolean if a field has been set.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) HasSlashPriceProductId() bool {
	if o != nil && !IsNil(o.SlashPriceProductId) {
		return true
	}

	return false
}

// SetSlashPriceProductId gets a reference to the given int64 and assigns it to the SlashPriceProductId field.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) SetSlashPriceProductId(v int64) {
	o.SlashPriceProductId = &v
}

// GetDiscountPercentage returns the DiscountPercentage field value if set, zero value otherwise.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetDiscountPercentage() int64 {
	if o == nil || IsNil(o.DiscountPercentage) {
		var ret int64
		return ret
	}
	return *o.DiscountPercentage
}

// GetDiscountPercentageOk returns a tuple with the DiscountPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetDiscountPercentageOk() (*int64, bool) {
	if o == nil || IsNil(o.DiscountPercentage) {
		return nil, false
	}
	return o.DiscountPercentage, true
}

// HasDiscountPercentage returns a boolean if a field has been set.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) HasDiscountPercentage() bool {
	if o != nil && !IsNil(o.DiscountPercentage) {
		return true
	}

	return false
}

// SetDiscountPercentage gets a reference to the given int64 and assigns it to the DiscountPercentage field.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) SetDiscountPercentage(v int64) {
	o.DiscountPercentage = &v
}

// GetDiscountedPrice returns the DiscountedPrice field value if set, zero value otherwise.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetDiscountedPrice() string {
	if o == nil || IsNil(o.DiscountedPrice) {
		var ret string
		return ret
	}
	return *o.DiscountedPrice
}

// GetDiscountedPriceOk returns a tuple with the DiscountedPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetDiscountedPriceOk() (*string, bool) {
	if o == nil || IsNil(o.DiscountedPrice) {
		return nil, false
	}
	return o.DiscountedPrice, true
}

// HasDiscountedPrice returns a boolean if a field has been set.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) HasDiscountedPrice() bool {
	if o != nil && !IsNil(o.DiscountedPrice) {
		return true
	}

	return false
}

// SetDiscountedPrice gets a reference to the given string and assigns it to the DiscountedPrice field.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) SetDiscountedPrice(v string) {
	o.DiscountedPrice = &v
}

// GetOriginalPrice returns the OriginalPrice field value if set, zero value otherwise.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetOriginalPrice() string {
	if o == nil || IsNil(o.OriginalPrice) {
		var ret string
		return ret
	}
	return *o.OriginalPrice
}

// GetOriginalPriceOk returns a tuple with the OriginalPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetOriginalPriceOk() (*string, bool) {
	if o == nil || IsNil(o.OriginalPrice) {
		return nil, false
	}
	return o.OriginalPrice, true
}

// HasOriginalPrice returns a boolean if a field has been set.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) HasOriginalPrice() bool {
	if o != nil && !IsNil(o.OriginalPrice) {
		return true
	}

	return false
}

// SetOriginalPrice gets a reference to the given string and assigns it to the OriginalPrice field.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) SetOriginalPrice(v string) {
	o.OriginalPrice = &v
}

// GetCustomStock returns the CustomStock field value if set, zero value otherwise.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetCustomStock() int64 {
	if o == nil || IsNil(o.CustomStock) {
		var ret int64
		return ret
	}
	return *o.CustomStock
}

// GetCustomStockOk returns a tuple with the CustomStock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetCustomStockOk() (*int64, bool) {
	if o == nil || IsNil(o.CustomStock) {
		return nil, false
	}
	return o.CustomStock, true
}

// HasCustomStock returns a boolean if a field has been set.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) HasCustomStock() bool {
	if o != nil && !IsNil(o.CustomStock) {
		return true
	}

	return false
}

// SetCustomStock gets a reference to the given int64 and assigns it to the CustomStock field.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) SetCustomStock(v int64) {
	o.CustomStock = &v
}

// GetCampaignStatus returns the CampaignStatus field value if set, zero value otherwise.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetCampaignStatus() int64 {
	if o == nil || IsNil(o.CampaignStatus) {
		var ret int64
		return ret
	}
	return *o.CampaignStatus
}

// GetCampaignStatusOk returns a tuple with the CampaignStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetCampaignStatusOk() (*int64, bool) {
	if o == nil || IsNil(o.CampaignStatus) {
		return nil, false
	}
	return o.CampaignStatus, true
}

// HasCampaignStatus returns a boolean if a field has been set.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) HasCampaignStatus() bool {
	if o != nil && !IsNil(o.CampaignStatus) {
		return true
	}

	return false
}

// SetCampaignStatus gets a reference to the given int64 and assigns it to the CampaignStatus field.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) SetCampaignStatus(v int64) {
	o.CampaignStatus = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) SetStartDate(v string) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetEndDate() string {
	if o == nil || IsNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) SetEndDate(v string) {
	o.EndDate = &v
}

// GetCampaignTypeName returns the CampaignTypeName field value if set, zero value otherwise.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetCampaignTypeName() string {
	if o == nil || IsNil(o.CampaignTypeName) {
		var ret string
		return ret
	}
	return *o.CampaignTypeName
}

// GetCampaignTypeNameOk returns a tuple with the CampaignTypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetCampaignTypeNameOk() (*string, bool) {
	if o == nil || IsNil(o.CampaignTypeName) {
		return nil, false
	}
	return o.CampaignTypeName, true
}

// HasCampaignTypeName returns a boolean if a field has been set.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) HasCampaignTypeName() bool {
	if o != nil && !IsNil(o.CampaignTypeName) {
		return true
	}

	return false
}

// SetCampaignTypeName gets a reference to the given string and assigns it to the CampaignTypeName field.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) SetCampaignTypeName(v string) {
	o.CampaignTypeName = &v
}

// GetCampaignShortName returns the CampaignShortName field value if set, zero value otherwise.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetCampaignShortName() string {
	if o == nil || IsNil(o.CampaignShortName) {
		var ret string
		return ret
	}
	return *o.CampaignShortName
}

// GetCampaignShortNameOk returns a tuple with the CampaignShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetCampaignShortNameOk() (*string, bool) {
	if o == nil || IsNil(o.CampaignShortName) {
		return nil, false
	}
	return o.CampaignShortName, true
}

// HasCampaignShortName returns a boolean if a field has been set.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) HasCampaignShortName() bool {
	if o != nil && !IsNil(o.CampaignShortName) {
		return true
	}

	return false
}

// SetCampaignShortName gets a reference to the given string and assigns it to the CampaignShortName field.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) SetCampaignShortName(v string) {
	o.CampaignShortName = &v
}

// GetMaxOrder returns the MaxOrder field value if set, zero value otherwise.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetMaxOrder() int64 {
	if o == nil || IsNil(o.MaxOrder) {
		var ret int64
		return ret
	}
	return *o.MaxOrder
}

// GetMaxOrderOk returns a tuple with the MaxOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetMaxOrderOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxOrder) {
		return nil, false
	}
	return o.MaxOrder, true
}

// HasMaxOrder returns a boolean if a field has been set.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) HasMaxOrder() bool {
	if o != nil && !IsNil(o.MaxOrder) {
		return true
	}

	return false
}

// SetMaxOrder gets a reference to the given int64 and assigns it to the MaxOrder field.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) SetMaxOrder(v int64) {
	o.MaxOrder = &v
}

// GetDiscountedPriceFmt returns the DiscountedPriceFmt field value if set, zero value otherwise.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetDiscountedPriceFmt() string {
	if o == nil || IsNil(o.DiscountedPriceFmt) {
		var ret string
		return ret
	}
	return *o.DiscountedPriceFmt
}

// GetDiscountedPriceFmtOk returns a tuple with the DiscountedPriceFmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetDiscountedPriceFmtOk() (*string, bool) {
	if o == nil || IsNil(o.DiscountedPriceFmt) {
		return nil, false
	}
	return o.DiscountedPriceFmt, true
}

// HasDiscountedPriceFmt returns a boolean if a field has been set.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) HasDiscountedPriceFmt() bool {
	if o != nil && !IsNil(o.DiscountedPriceFmt) {
		return true
	}

	return false
}

// SetDiscountedPriceFmt gets a reference to the given string and assigns it to the DiscountedPriceFmt field.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) SetDiscountedPriceFmt(v string) {
	o.DiscountedPriceFmt = &v
}

// GetOriginalPriceFmt returns the OriginalPriceFmt field value if set, zero value otherwise.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetOriginalPriceFmt() string {
	if o == nil || IsNil(o.OriginalPriceFmt) {
		var ret string
		return ret
	}
	return *o.OriginalPriceFmt
}

// GetOriginalPriceFmtOk returns a tuple with the OriginalPriceFmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetOriginalPriceFmtOk() (*string, bool) {
	if o == nil || IsNil(o.OriginalPriceFmt) {
		return nil, false
	}
	return o.OriginalPriceFmt, true
}

// HasOriginalPriceFmt returns a boolean if a field has been set.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) HasOriginalPriceFmt() bool {
	if o != nil && !IsNil(o.OriginalPriceFmt) {
		return true
	}

	return false
}

// SetOriginalPriceFmt gets a reference to the given string and assigns it to the OriginalPriceFmt field.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) SetOriginalPriceFmt(v string) {
	o.OriginalPriceFmt = &v
}

// GetOriginalStock returns the OriginalStock field value if set, zero value otherwise.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetOriginalStock() int64 {
	if o == nil || IsNil(o.OriginalStock) {
		var ret int64
		return ret
	}
	return *o.OriginalStock
}

// GetOriginalStockOk returns a tuple with the OriginalStock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetOriginalStockOk() (*int64, bool) {
	if o == nil || IsNil(o.OriginalStock) {
		return nil, false
	}
	return o.OriginalStock, true
}

// HasOriginalStock returns a boolean if a field has been set.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) HasOriginalStock() bool {
	if o != nil && !IsNil(o.OriginalStock) {
		return true
	}

	return false
}

// SetOriginalStock gets a reference to the given int64 and assigns it to the OriginalStock field.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) SetOriginalStock(v int64) {
	o.OriginalStock = &v
}

// GetMinOrder returns the MinOrder field value if set, zero value otherwise.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetMinOrder() int64 {
	if o == nil || IsNil(o.MinOrder) {
		var ret int64
		return ret
	}
	return *o.MinOrder
}

// GetMinOrderOk returns a tuple with the MinOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetMinOrderOk() (*int64, bool) {
	if o == nil || IsNil(o.MinOrder) {
		return nil, false
	}
	return o.MinOrder, true
}

// HasMinOrder returns a boolean if a field has been set.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) HasMinOrder() bool {
	if o != nil && !IsNil(o.MinOrder) {
		return true
	}

	return false
}

// SetMinOrder gets a reference to the given int64 and assigns it to the MinOrder field.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) SetMinOrder(v int64) {
	o.MinOrder = &v
}

// GetAdditionalCartInfo returns the AdditionalCartInfo field value if set, zero value otherwise.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetAdditionalCartInfo() ViewCampaignProducts200ResponseDataProductsInnerCampaignAdditionalCartInfo {
	if o == nil || IsNil(o.AdditionalCartInfo) {
		var ret ViewCampaignProducts200ResponseDataProductsInnerCampaignAdditionalCartInfo
		return ret
	}
	return *o.AdditionalCartInfo
}

// GetAdditionalCartInfoOk returns a tuple with the AdditionalCartInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetAdditionalCartInfoOk() (*ViewCampaignProducts200ResponseDataProductsInnerCampaignAdditionalCartInfo, bool) {
	if o == nil || IsNil(o.AdditionalCartInfo) {
		return nil, false
	}
	return o.AdditionalCartInfo, true
}

// HasAdditionalCartInfo returns a boolean if a field has been set.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) HasAdditionalCartInfo() bool {
	if o != nil && !IsNil(o.AdditionalCartInfo) {
		return true
	}

	return false
}

// SetAdditionalCartInfo gets a reference to the given ViewCampaignProducts200ResponseDataProductsInnerCampaignAdditionalCartInfo and assigns it to the AdditionalCartInfo field.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) SetAdditionalCartInfo(v ViewCampaignProducts200ResponseDataProductsInnerCampaignAdditionalCartInfo) {
	o.AdditionalCartInfo = &v
}

// GetStockWording returns the StockWording field value if set, zero value otherwise.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetStockWording() map[string]interface{} {
	if o == nil || IsNil(o.StockWording) {
		var ret map[string]interface{}
		return ret
	}
	return o.StockWording
}

// GetStockWordingOk returns a tuple with the StockWording field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) GetStockWordingOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.StockWording) {
		return map[string]interface{}{}, false
	}
	return o.StockWording, true
}

// HasStockWording returns a boolean if a field has been set.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) HasStockWording() bool {
	if o != nil && !IsNil(o.StockWording) {
		return true
	}

	return false
}

// SetStockWording gets a reference to the given map[string]interface{} and assigns it to the StockWording field.
func (o *ViewCampaignProducts200ResponseDataProductsInnerCampaign) SetStockWording(v map[string]interface{}) {
	o.StockWording = v
}

func (o ViewCampaignProducts200ResponseDataProductsInnerCampaign) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ViewCampaignProducts200ResponseDataProductsInnerCampaign) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CampaignID) {
		toSerialize["campaignID"] = o.CampaignID
	}
	if !IsNil(o.SlashPriceProductId) {
		toSerialize["slash_price_product_id"] = o.SlashPriceProductId
	}
	if !IsNil(o.DiscountPercentage) {
		toSerialize["discount_percentage"] = o.DiscountPercentage
	}
	if !IsNil(o.DiscountedPrice) {
		toSerialize["discounted_price"] = o.DiscountedPrice
	}
	if !IsNil(o.OriginalPrice) {
		toSerialize["original_price"] = o.OriginalPrice
	}
	if !IsNil(o.CustomStock) {
		toSerialize["custom_stock"] = o.CustomStock
	}
	if !IsNil(o.CampaignStatus) {
		toSerialize["campaign_status"] = o.CampaignStatus
	}
	if !IsNil(o.StartDate) {
		toSerialize["start_date"] = o.StartDate
	}
	if !IsNil(o.EndDate) {
		toSerialize["end_date"] = o.EndDate
	}
	if !IsNil(o.CampaignTypeName) {
		toSerialize["campaign_type_name"] = o.CampaignTypeName
	}
	if !IsNil(o.CampaignShortName) {
		toSerialize["campaign_short_name"] = o.CampaignShortName
	}
	if !IsNil(o.MaxOrder) {
		toSerialize["max_order"] = o.MaxOrder
	}
	if !IsNil(o.DiscountedPriceFmt) {
		toSerialize["discounted_price_fmt"] = o.DiscountedPriceFmt
	}
	if !IsNil(o.OriginalPriceFmt) {
		toSerialize["original_price_fmt"] = o.OriginalPriceFmt
	}
	if !IsNil(o.OriginalStock) {
		toSerialize["original_stock"] = o.OriginalStock
	}
	if !IsNil(o.MinOrder) {
		toSerialize["min_order"] = o.MinOrder
	}
	if !IsNil(o.AdditionalCartInfo) {
		toSerialize["additional_cart_info"] = o.AdditionalCartInfo
	}
	if !IsNil(o.StockWording) {
		toSerialize["stock_wording"] = o.StockWording
	}
	return toSerialize, nil
}

type NullableViewCampaignProducts200ResponseDataProductsInnerCampaign struct {
	value *ViewCampaignProducts200ResponseDataProductsInnerCampaign
	isSet bool
}

func (v NullableViewCampaignProducts200ResponseDataProductsInnerCampaign) Get() *ViewCampaignProducts200ResponseDataProductsInnerCampaign {
	return v.value
}

func (v *NullableViewCampaignProducts200ResponseDataProductsInnerCampaign) Set(val *ViewCampaignProducts200ResponseDataProductsInnerCampaign) {
	v.value = val
	v.isSet = true
}

func (v NullableViewCampaignProducts200ResponseDataProductsInnerCampaign) IsSet() bool {
	return v.isSet
}

func (v *NullableViewCampaignProducts200ResponseDataProductsInnerCampaign) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViewCampaignProducts200ResponseDataProductsInnerCampaign(val *ViewCampaignProducts200ResponseDataProductsInnerCampaign) *NullableViewCampaignProducts200ResponseDataProductsInnerCampaign {
	return &NullableViewCampaignProducts200ResponseDataProductsInnerCampaign{value: val, isSet: true}
}

func (v NullableViewCampaignProducts200ResponseDataProductsInnerCampaign) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableViewCampaignProducts200ResponseDataProductsInnerCampaign) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

