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

// checks if the GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner{}

// GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner struct for GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner
type GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner struct {
	// Bundle Order Detail Unique Identifier
	OrderDtlId *int64 `json:"order_dtl_id,omitempty"`
	// Bundle Order Unique Identifier
	OrderId *int64 `json:"order_id,omitempty"`
	// Product Unique Identifier
	ProductId *int64 `json:"product_id,omitempty"`
	// Product Name
	ProductName *string `json:"product_name,omitempty"`
	// Product Descirption
	ProductDesc *string `json:"product_desc,omitempty"`
	// Product Order Quantity
	Quantity *int64 `json:"quantity,omitempty"`
	// Product Price
	ProductPrice *float64 `json:"product_price,omitempty"`
	// Product Weight
	ProductWeight *float64 `json:"product_weight,omitempty"`
	// Product Order Total Weight
	TotalWeight *float64 `json:"total_weight,omitempty"`
	// Product Order Total Price
	SubtotalPrice *float64 `json:"subtotal_price,omitempty"`
	// Product Order Notes
	Notes *string `json:"notes,omitempty"`
	// 
	Finsurance *int64 `json:"finsurance,omitempty"`
	// Is returnable?
	Returnable *int64 `json:"returnable,omitempty"`
	// Category Unique Identifier
	ChildCatId *int64 `json:"child_cat_id,omitempty"`
	// Currency Unique Identifier
	CurrencyId *int64 `json:"currency_id,omitempty"`
	// Product Insurance Price
	InsurancePrice *float64 `json:"insurance_price,omitempty"`
	// Product Normal Price
	NormalPrice *float64 `json:"normal_price,omitempty"`
	// 
	CurrencyRate *float64 `json:"currency_rate,omitempty"`
	// 
	ProdPic *string `json:"prod_pic,omitempty"`
	// Minimum Product Order
	MinOrder *int64 `json:"min_order,omitempty"`
	// Is product must include insurance?
	MustInsurance *int64 `json:"must_insurance,omitempty"`
	// Product Condition, 1 = New, 2 = Used
	Condition *int64 `json:"condition,omitempty"`
	// Product Campaign Unique Identifier
	CampaignId *int64 `json:"campaign_id,omitempty"`
	// Product SKU
	Sku *string `json:"sku,omitempty"`
	// Is product has slash price?
	IsSlashPrice *bool `json:"is_slash_price,omitempty"`
	// 
	OmsDetailData *string `json:"oms_detail_data,omitempty"`
	// Product Thumbnail URL
	Thumbnail *string `json:"thumbnail,omitempty"`
}

// NewGetAllOrders200ResponseDataInnerBundleDetailNonBundleInner instantiates a new GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAllOrders200ResponseDataInnerBundleDetailNonBundleInner() *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner {
	this := GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner{}
	return &this
}

// NewGetAllOrders200ResponseDataInnerBundleDetailNonBundleInnerWithDefaults instantiates a new GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAllOrders200ResponseDataInnerBundleDetailNonBundleInnerWithDefaults() *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner {
	this := GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner{}
	return &this
}

// GetOrderDtlId returns the OrderDtlId field value if set, zero value otherwise.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) GetOrderDtlId() int64 {
	if o == nil || IsNil(o.OrderDtlId) {
		var ret int64
		return ret
	}
	return *o.OrderDtlId
}

// GetOrderDtlIdOk returns a tuple with the OrderDtlId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) GetOrderDtlIdOk() (*int64, bool) {
	if o == nil || IsNil(o.OrderDtlId) {
		return nil, false
	}
	return o.OrderDtlId, true
}

// HasOrderDtlId returns a boolean if a field has been set.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) HasOrderDtlId() bool {
	if o != nil && !IsNil(o.OrderDtlId) {
		return true
	}

	return false
}

// SetOrderDtlId gets a reference to the given int64 and assigns it to the OrderDtlId field.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) SetOrderDtlId(v int64) {
	o.OrderDtlId = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) GetOrderId() int64 {
	if o == nil || IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) GetOrderIdOk() (*int64, bool) {
	if o == nil || IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) HasOrderId() bool {
	if o != nil && !IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) GetProductId() int64 {
	if o == nil || IsNil(o.ProductId) {
		var ret int64
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) GetProductIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ProductId) {
		return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) HasProductId() bool {
	if o != nil && !IsNil(o.ProductId) {
		return true
	}

	return false
}

// SetProductId gets a reference to the given int64 and assigns it to the ProductId field.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) SetProductId(v int64) {
	o.ProductId = &v
}

// GetProductName returns the ProductName field value if set, zero value otherwise.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) GetProductName() string {
	if o == nil || IsNil(o.ProductName) {
		var ret string
		return ret
	}
	return *o.ProductName
}

// GetProductNameOk returns a tuple with the ProductName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) GetProductNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProductName) {
		return nil, false
	}
	return o.ProductName, true
}

// HasProductName returns a boolean if a field has been set.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) HasProductName() bool {
	if o != nil && !IsNil(o.ProductName) {
		return true
	}

	return false
}

// SetProductName gets a reference to the given string and assigns it to the ProductName field.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) SetProductName(v string) {
	o.ProductName = &v
}

// GetProductDesc returns the ProductDesc field value if set, zero value otherwise.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) GetProductDesc() string {
	if o == nil || IsNil(o.ProductDesc) {
		var ret string
		return ret
	}
	return *o.ProductDesc
}

// GetProductDescOk returns a tuple with the ProductDesc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) GetProductDescOk() (*string, bool) {
	if o == nil || IsNil(o.ProductDesc) {
		return nil, false
	}
	return o.ProductDesc, true
}

// HasProductDesc returns a boolean if a field has been set.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) HasProductDesc() bool {
	if o != nil && !IsNil(o.ProductDesc) {
		return true
	}

	return false
}

// SetProductDesc gets a reference to the given string and assigns it to the ProductDesc field.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) SetProductDesc(v string) {
	o.ProductDesc = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) GetQuantity() int64 {
	if o == nil || IsNil(o.Quantity) {
		var ret int64
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) GetQuantityOk() (*int64, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int64 and assigns it to the Quantity field.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) SetQuantity(v int64) {
	o.Quantity = &v
}

// GetProductPrice returns the ProductPrice field value if set, zero value otherwise.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) GetProductPrice() float64 {
	if o == nil || IsNil(o.ProductPrice) {
		var ret float64
		return ret
	}
	return *o.ProductPrice
}

// GetProductPriceOk returns a tuple with the ProductPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) GetProductPriceOk() (*float64, bool) {
	if o == nil || IsNil(o.ProductPrice) {
		return nil, false
	}
	return o.ProductPrice, true
}

// HasProductPrice returns a boolean if a field has been set.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) HasProductPrice() bool {
	if o != nil && !IsNil(o.ProductPrice) {
		return true
	}

	return false
}

// SetProductPrice gets a reference to the given float64 and assigns it to the ProductPrice field.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) SetProductPrice(v float64) {
	o.ProductPrice = &v
}

// GetProductWeight returns the ProductWeight field value if set, zero value otherwise.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) GetProductWeight() float64 {
	if o == nil || IsNil(o.ProductWeight) {
		var ret float64
		return ret
	}
	return *o.ProductWeight
}

// GetProductWeightOk returns a tuple with the ProductWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) GetProductWeightOk() (*float64, bool) {
	if o == nil || IsNil(o.ProductWeight) {
		return nil, false
	}
	return o.ProductWeight, true
}

// HasProductWeight returns a boolean if a field has been set.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) HasProductWeight() bool {
	if o != nil && !IsNil(o.ProductWeight) {
		return true
	}

	return false
}

// SetProductWeight gets a reference to the given float64 and assigns it to the ProductWeight field.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) SetProductWeight(v float64) {
	o.ProductWeight = &v
}

// GetTotalWeight returns the TotalWeight field value if set, zero value otherwise.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) GetTotalWeight() float64 {
	if o == nil || IsNil(o.TotalWeight) {
		var ret float64
		return ret
	}
	return *o.TotalWeight
}

// GetTotalWeightOk returns a tuple with the TotalWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) GetTotalWeightOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalWeight) {
		return nil, false
	}
	return o.TotalWeight, true
}

// HasTotalWeight returns a boolean if a field has been set.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) HasTotalWeight() bool {
	if o != nil && !IsNil(o.TotalWeight) {
		return true
	}

	return false
}

// SetTotalWeight gets a reference to the given float64 and assigns it to the TotalWeight field.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) SetTotalWeight(v float64) {
	o.TotalWeight = &v
}

// GetSubtotalPrice returns the SubtotalPrice field value if set, zero value otherwise.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) GetSubtotalPrice() float64 {
	if o == nil || IsNil(o.SubtotalPrice) {
		var ret float64
		return ret
	}
	return *o.SubtotalPrice
}

// GetSubtotalPriceOk returns a tuple with the SubtotalPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) GetSubtotalPriceOk() (*float64, bool) {
	if o == nil || IsNil(o.SubtotalPrice) {
		return nil, false
	}
	return o.SubtotalPrice, true
}

// HasSubtotalPrice returns a boolean if a field has been set.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) HasSubtotalPrice() bool {
	if o != nil && !IsNil(o.SubtotalPrice) {
		return true
	}

	return false
}

// SetSubtotalPrice gets a reference to the given float64 and assigns it to the SubtotalPrice field.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) SetSubtotalPrice(v float64) {
	o.SubtotalPrice = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) SetNotes(v string) {
	o.Notes = &v
}

// GetFinsurance returns the Finsurance field value if set, zero value otherwise.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) GetFinsurance() int64 {
	if o == nil || IsNil(o.Finsurance) {
		var ret int64
		return ret
	}
	return *o.Finsurance
}

// GetFinsuranceOk returns a tuple with the Finsurance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) GetFinsuranceOk() (*int64, bool) {
	if o == nil || IsNil(o.Finsurance) {
		return nil, false
	}
	return o.Finsurance, true
}

// HasFinsurance returns a boolean if a field has been set.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) HasFinsurance() bool {
	if o != nil && !IsNil(o.Finsurance) {
		return true
	}

	return false
}

// SetFinsurance gets a reference to the given int64 and assigns it to the Finsurance field.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) SetFinsurance(v int64) {
	o.Finsurance = &v
}

// GetReturnable returns the Returnable field value if set, zero value otherwise.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) GetReturnable() int64 {
	if o == nil || IsNil(o.Returnable) {
		var ret int64
		return ret
	}
	return *o.Returnable
}

// GetReturnableOk returns a tuple with the Returnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) GetReturnableOk() (*int64, bool) {
	if o == nil || IsNil(o.Returnable) {
		return nil, false
	}
	return o.Returnable, true
}

// HasReturnable returns a boolean if a field has been set.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) HasReturnable() bool {
	if o != nil && !IsNil(o.Returnable) {
		return true
	}

	return false
}

// SetReturnable gets a reference to the given int64 and assigns it to the Returnable field.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) SetReturnable(v int64) {
	o.Returnable = &v
}

// GetChildCatId returns the ChildCatId field value if set, zero value otherwise.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) GetChildCatId() int64 {
	if o == nil || IsNil(o.ChildCatId) {
		var ret int64
		return ret
	}
	return *o.ChildCatId
}

// GetChildCatIdOk returns a tuple with the ChildCatId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) GetChildCatIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ChildCatId) {
		return nil, false
	}
	return o.ChildCatId, true
}

// HasChildCatId returns a boolean if a field has been set.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) HasChildCatId() bool {
	if o != nil && !IsNil(o.ChildCatId) {
		return true
	}

	return false
}

// SetChildCatId gets a reference to the given int64 and assigns it to the ChildCatId field.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) SetChildCatId(v int64) {
	o.ChildCatId = &v
}

// GetCurrencyId returns the CurrencyId field value if set, zero value otherwise.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) GetCurrencyId() int64 {
	if o == nil || IsNil(o.CurrencyId) {
		var ret int64
		return ret
	}
	return *o.CurrencyId
}

// GetCurrencyIdOk returns a tuple with the CurrencyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) GetCurrencyIdOk() (*int64, bool) {
	if o == nil || IsNil(o.CurrencyId) {
		return nil, false
	}
	return o.CurrencyId, true
}

// HasCurrencyId returns a boolean if a field has been set.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) HasCurrencyId() bool {
	if o != nil && !IsNil(o.CurrencyId) {
		return true
	}

	return false
}

// SetCurrencyId gets a reference to the given int64 and assigns it to the CurrencyId field.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) SetCurrencyId(v int64) {
	o.CurrencyId = &v
}

// GetInsurancePrice returns the InsurancePrice field value if set, zero value otherwise.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) GetInsurancePrice() float64 {
	if o == nil || IsNil(o.InsurancePrice) {
		var ret float64
		return ret
	}
	return *o.InsurancePrice
}

// GetInsurancePriceOk returns a tuple with the InsurancePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) GetInsurancePriceOk() (*float64, bool) {
	if o == nil || IsNil(o.InsurancePrice) {
		return nil, false
	}
	return o.InsurancePrice, true
}

// HasInsurancePrice returns a boolean if a field has been set.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) HasInsurancePrice() bool {
	if o != nil && !IsNil(o.InsurancePrice) {
		return true
	}

	return false
}

// SetInsurancePrice gets a reference to the given float64 and assigns it to the InsurancePrice field.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) SetInsurancePrice(v float64) {
	o.InsurancePrice = &v
}

// GetNormalPrice returns the NormalPrice field value if set, zero value otherwise.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) GetNormalPrice() float64 {
	if o == nil || IsNil(o.NormalPrice) {
		var ret float64
		return ret
	}
	return *o.NormalPrice
}

// GetNormalPriceOk returns a tuple with the NormalPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) GetNormalPriceOk() (*float64, bool) {
	if o == nil || IsNil(o.NormalPrice) {
		return nil, false
	}
	return o.NormalPrice, true
}

// HasNormalPrice returns a boolean if a field has been set.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) HasNormalPrice() bool {
	if o != nil && !IsNil(o.NormalPrice) {
		return true
	}

	return false
}

// SetNormalPrice gets a reference to the given float64 and assigns it to the NormalPrice field.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) SetNormalPrice(v float64) {
	o.NormalPrice = &v
}

// GetCurrencyRate returns the CurrencyRate field value if set, zero value otherwise.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) GetCurrencyRate() float64 {
	if o == nil || IsNil(o.CurrencyRate) {
		var ret float64
		return ret
	}
	return *o.CurrencyRate
}

// GetCurrencyRateOk returns a tuple with the CurrencyRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) GetCurrencyRateOk() (*float64, bool) {
	if o == nil || IsNil(o.CurrencyRate) {
		return nil, false
	}
	return o.CurrencyRate, true
}

// HasCurrencyRate returns a boolean if a field has been set.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) HasCurrencyRate() bool {
	if o != nil && !IsNil(o.CurrencyRate) {
		return true
	}

	return false
}

// SetCurrencyRate gets a reference to the given float64 and assigns it to the CurrencyRate field.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) SetCurrencyRate(v float64) {
	o.CurrencyRate = &v
}

// GetProdPic returns the ProdPic field value if set, zero value otherwise.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) GetProdPic() string {
	if o == nil || IsNil(o.ProdPic) {
		var ret string
		return ret
	}
	return *o.ProdPic
}

// GetProdPicOk returns a tuple with the ProdPic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) GetProdPicOk() (*string, bool) {
	if o == nil || IsNil(o.ProdPic) {
		return nil, false
	}
	return o.ProdPic, true
}

// HasProdPic returns a boolean if a field has been set.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) HasProdPic() bool {
	if o != nil && !IsNil(o.ProdPic) {
		return true
	}

	return false
}

// SetProdPic gets a reference to the given string and assigns it to the ProdPic field.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) SetProdPic(v string) {
	o.ProdPic = &v
}

// GetMinOrder returns the MinOrder field value if set, zero value otherwise.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) GetMinOrder() int64 {
	if o == nil || IsNil(o.MinOrder) {
		var ret int64
		return ret
	}
	return *o.MinOrder
}

// GetMinOrderOk returns a tuple with the MinOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) GetMinOrderOk() (*int64, bool) {
	if o == nil || IsNil(o.MinOrder) {
		return nil, false
	}
	return o.MinOrder, true
}

// HasMinOrder returns a boolean if a field has been set.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) HasMinOrder() bool {
	if o != nil && !IsNil(o.MinOrder) {
		return true
	}

	return false
}

// SetMinOrder gets a reference to the given int64 and assigns it to the MinOrder field.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) SetMinOrder(v int64) {
	o.MinOrder = &v
}

// GetMustInsurance returns the MustInsurance field value if set, zero value otherwise.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) GetMustInsurance() int64 {
	if o == nil || IsNil(o.MustInsurance) {
		var ret int64
		return ret
	}
	return *o.MustInsurance
}

// GetMustInsuranceOk returns a tuple with the MustInsurance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) GetMustInsuranceOk() (*int64, bool) {
	if o == nil || IsNil(o.MustInsurance) {
		return nil, false
	}
	return o.MustInsurance, true
}

// HasMustInsurance returns a boolean if a field has been set.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) HasMustInsurance() bool {
	if o != nil && !IsNil(o.MustInsurance) {
		return true
	}

	return false
}

// SetMustInsurance gets a reference to the given int64 and assigns it to the MustInsurance field.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) SetMustInsurance(v int64) {
	o.MustInsurance = &v
}

// GetCondition returns the Condition field value if set, zero value otherwise.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) GetCondition() int64 {
	if o == nil || IsNil(o.Condition) {
		var ret int64
		return ret
	}
	return *o.Condition
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) GetConditionOk() (*int64, bool) {
	if o == nil || IsNil(o.Condition) {
		return nil, false
	}
	return o.Condition, true
}

// HasCondition returns a boolean if a field has been set.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) HasCondition() bool {
	if o != nil && !IsNil(o.Condition) {
		return true
	}

	return false
}

// SetCondition gets a reference to the given int64 and assigns it to the Condition field.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) SetCondition(v int64) {
	o.Condition = &v
}

// GetCampaignId returns the CampaignId field value if set, zero value otherwise.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) GetCampaignId() int64 {
	if o == nil || IsNil(o.CampaignId) {
		var ret int64
		return ret
	}
	return *o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) GetCampaignIdOk() (*int64, bool) {
	if o == nil || IsNil(o.CampaignId) {
		return nil, false
	}
	return o.CampaignId, true
}

// HasCampaignId returns a boolean if a field has been set.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) HasCampaignId() bool {
	if o != nil && !IsNil(o.CampaignId) {
		return true
	}

	return false
}

// SetCampaignId gets a reference to the given int64 and assigns it to the CampaignId field.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) SetCampaignId(v int64) {
	o.CampaignId = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) GetSku() string {
	if o == nil || IsNil(o.Sku) {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) GetSkuOk() (*string, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) SetSku(v string) {
	o.Sku = &v
}

// GetIsSlashPrice returns the IsSlashPrice field value if set, zero value otherwise.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) GetIsSlashPrice() bool {
	if o == nil || IsNil(o.IsSlashPrice) {
		var ret bool
		return ret
	}
	return *o.IsSlashPrice
}

// GetIsSlashPriceOk returns a tuple with the IsSlashPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) GetIsSlashPriceOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSlashPrice) {
		return nil, false
	}
	return o.IsSlashPrice, true
}

// HasIsSlashPrice returns a boolean if a field has been set.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) HasIsSlashPrice() bool {
	if o != nil && !IsNil(o.IsSlashPrice) {
		return true
	}

	return false
}

// SetIsSlashPrice gets a reference to the given bool and assigns it to the IsSlashPrice field.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) SetIsSlashPrice(v bool) {
	o.IsSlashPrice = &v
}

// GetOmsDetailData returns the OmsDetailData field value if set, zero value otherwise.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) GetOmsDetailData() string {
	if o == nil || IsNil(o.OmsDetailData) {
		var ret string
		return ret
	}
	return *o.OmsDetailData
}

// GetOmsDetailDataOk returns a tuple with the OmsDetailData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) GetOmsDetailDataOk() (*string, bool) {
	if o == nil || IsNil(o.OmsDetailData) {
		return nil, false
	}
	return o.OmsDetailData, true
}

// HasOmsDetailData returns a boolean if a field has been set.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) HasOmsDetailData() bool {
	if o != nil && !IsNil(o.OmsDetailData) {
		return true
	}

	return false
}

// SetOmsDetailData gets a reference to the given string and assigns it to the OmsDetailData field.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) SetOmsDetailData(v string) {
	o.OmsDetailData = &v
}

// GetThumbnail returns the Thumbnail field value if set, zero value otherwise.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) GetThumbnail() string {
	if o == nil || IsNil(o.Thumbnail) {
		var ret string
		return ret
	}
	return *o.Thumbnail
}

// GetThumbnailOk returns a tuple with the Thumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) GetThumbnailOk() (*string, bool) {
	if o == nil || IsNil(o.Thumbnail) {
		return nil, false
	}
	return o.Thumbnail, true
}

// HasThumbnail returns a boolean if a field has been set.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) HasThumbnail() bool {
	if o != nil && !IsNil(o.Thumbnail) {
		return true
	}

	return false
}

// SetThumbnail gets a reference to the given string and assigns it to the Thumbnail field.
func (o *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) SetThumbnail(v string) {
	o.Thumbnail = &v
}

func (o GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OrderDtlId) {
		toSerialize["order_dtl_id"] = o.OrderDtlId
	}
	if !IsNil(o.OrderId) {
		toSerialize["order_id"] = o.OrderId
	}
	if !IsNil(o.ProductId) {
		toSerialize["product_id"] = o.ProductId
	}
	if !IsNil(o.ProductName) {
		toSerialize["product_name"] = o.ProductName
	}
	if !IsNil(o.ProductDesc) {
		toSerialize["product_desc"] = o.ProductDesc
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.ProductPrice) {
		toSerialize["product_price"] = o.ProductPrice
	}
	if !IsNil(o.ProductWeight) {
		toSerialize["product_weight"] = o.ProductWeight
	}
	if !IsNil(o.TotalWeight) {
		toSerialize["total_weight"] = o.TotalWeight
	}
	if !IsNil(o.SubtotalPrice) {
		toSerialize["subtotal_price"] = o.SubtotalPrice
	}
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if !IsNil(o.Finsurance) {
		toSerialize["finsurance"] = o.Finsurance
	}
	if !IsNil(o.Returnable) {
		toSerialize["returnable"] = o.Returnable
	}
	if !IsNil(o.ChildCatId) {
		toSerialize["child_cat_id"] = o.ChildCatId
	}
	if !IsNil(o.CurrencyId) {
		toSerialize["currency_id"] = o.CurrencyId
	}
	if !IsNil(o.InsurancePrice) {
		toSerialize["insurance_price"] = o.InsurancePrice
	}
	if !IsNil(o.NormalPrice) {
		toSerialize["normal_price"] = o.NormalPrice
	}
	if !IsNil(o.CurrencyRate) {
		toSerialize["currency_rate"] = o.CurrencyRate
	}
	if !IsNil(o.ProdPic) {
		toSerialize["prod_pic"] = o.ProdPic
	}
	if !IsNil(o.MinOrder) {
		toSerialize["min_order"] = o.MinOrder
	}
	if !IsNil(o.MustInsurance) {
		toSerialize["must_insurance"] = o.MustInsurance
	}
	if !IsNil(o.Condition) {
		toSerialize["condition"] = o.Condition
	}
	if !IsNil(o.CampaignId) {
		toSerialize["campaign_id"] = o.CampaignId
	}
	if !IsNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	if !IsNil(o.IsSlashPrice) {
		toSerialize["is_slash_price"] = o.IsSlashPrice
	}
	if !IsNil(o.OmsDetailData) {
		toSerialize["oms_detail_data"] = o.OmsDetailData
	}
	if !IsNil(o.Thumbnail) {
		toSerialize["thumbnail"] = o.Thumbnail
	}
	return toSerialize, nil
}

type NullableGetAllOrders200ResponseDataInnerBundleDetailNonBundleInner struct {
	value *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner
	isSet bool
}

func (v NullableGetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) Get() *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner {
	return v.value
}

func (v *NullableGetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) Set(val *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAllOrders200ResponseDataInnerBundleDetailNonBundleInner(val *GetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) *NullableGetAllOrders200ResponseDataInnerBundleDetailNonBundleInner {
	return &NullableGetAllOrders200ResponseDataInnerBundleDetailNonBundleInner{value: val, isSet: true}
}

func (v NullableGetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAllOrders200ResponseDataInnerBundleDetailNonBundleInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


