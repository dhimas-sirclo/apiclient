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

// checks if the GetSingleOrder200ResponseDataOrderInfoShippingInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSingleOrder200ResponseDataOrderInfoShippingInfo{}

// GetSingleOrder200ResponseDataOrderInfoShippingInfo struct for GetSingleOrder200ResponseDataOrderInfoShippingInfo
type GetSingleOrder200ResponseDataOrderInfoShippingInfo struct {
	// Shipper Unique Identifier
	SpId *int64 `json:"sp_id,omitempty"`
	// Shipping Unique Identifier
	ShippingId *int64 `json:"shipping_id,omitempty"`
	// Logistic Name
	LogisticName *string `json:"logistic_name,omitempty"`
	// Logistic Service Name
	LogisticService *string `json:"logistic_service,omitempty"`
	// Shipping Price
	ShippingPrice *int64 `json:"shipping_price,omitempty"`
	// Shipping Price Rate
	ShippingPriceRate *int64 `json:"shipping_price_rate,omitempty"`
	// Shipping Fee
	ShippingFee *int64 `json:"shipping_fee,omitempty"`
	// Insurance Price
	InsurancePrice *int64 `json:"insurance_price,omitempty"`
	// Fee
	Fee *int64 `json:"fee,omitempty"`
	// Is change courirer?
	IsChangeCourier *bool `json:"is_change_courier,omitempty"`
	// Second Shipper Unique Identifier
	SecondSpId *int64 `json:"second_sp_id,omitempty"`
	// Second Shipping Unique Identifier
	SecondShippingId *int64 `json:"second_shipping_id,omitempty"`
	// Second Logistic Name
	SecondLogisticName *string `json:"second_logistic_name,omitempty"`
	// Second Logistic Service Name
	SecondLogisticService *string `json:"second_logistic_service,omitempty"`
	// Second Agency Fee
	SecondAgencyFee *int64 `json:"second_agency_fee,omitempty"`
	// Second Insurance
	SecondInsurance *int64 `json:"second_insurance,omitempty"`
	// Second Shipping Price Rate
	SecondRate *int64 `json:"second_rate,omitempty"`
	// Airway Bill (Resi)
	Awb *string `json:"awb,omitempty"`
	// Autoresi Cashless Status
	AutoresiCashlessStatus *int64 `json:"autoresi_cashless_status,omitempty"`
	// Airway Bill (Auto Resi)
	AutoresiAwb *string `json:"autoresi_awb,omitempty"`
	// AWB Shipping Price
	AutoresiShippingPrice *int64 `json:"autoresi_shipping_price,omitempty"`
	// AWB Count
	CountAwb *int64 `json:"count_awb,omitempty"`
	// Is cashless?
	IsCashless *bool `json:"isCashless,omitempty"`
	// Is fake delivery?
	IsFakeDelivery *bool `json:"is_fake_delivery,omitempty"`
	RecommendedCourierInfo []GetSingleOrder200ResponseDataOrderInfoShippingInfoRecommendedCourierInfoInner `json:"recommended_courier_info,omitempty"`
}

// NewGetSingleOrder200ResponseDataOrderInfoShippingInfo instantiates a new GetSingleOrder200ResponseDataOrderInfoShippingInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSingleOrder200ResponseDataOrderInfoShippingInfo() *GetSingleOrder200ResponseDataOrderInfoShippingInfo {
	this := GetSingleOrder200ResponseDataOrderInfoShippingInfo{}
	return &this
}

// NewGetSingleOrder200ResponseDataOrderInfoShippingInfoWithDefaults instantiates a new GetSingleOrder200ResponseDataOrderInfoShippingInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSingleOrder200ResponseDataOrderInfoShippingInfoWithDefaults() *GetSingleOrder200ResponseDataOrderInfoShippingInfo {
	this := GetSingleOrder200ResponseDataOrderInfoShippingInfo{}
	return &this
}

// GetSpId returns the SpId field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetSpId() int64 {
	if o == nil || IsNil(o.SpId) {
		var ret int64
		return ret
	}
	return *o.SpId
}

// GetSpIdOk returns a tuple with the SpId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetSpIdOk() (*int64, bool) {
	if o == nil || IsNil(o.SpId) {
		return nil, false
	}
	return o.SpId, true
}

// HasSpId returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) HasSpId() bool {
	if o != nil && !IsNil(o.SpId) {
		return true
	}

	return false
}

// SetSpId gets a reference to the given int64 and assigns it to the SpId field.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) SetSpId(v int64) {
	o.SpId = &v
}

// GetShippingId returns the ShippingId field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetShippingId() int64 {
	if o == nil || IsNil(o.ShippingId) {
		var ret int64
		return ret
	}
	return *o.ShippingId
}

// GetShippingIdOk returns a tuple with the ShippingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetShippingIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ShippingId) {
		return nil, false
	}
	return o.ShippingId, true
}

// HasShippingId returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) HasShippingId() bool {
	if o != nil && !IsNil(o.ShippingId) {
		return true
	}

	return false
}

// SetShippingId gets a reference to the given int64 and assigns it to the ShippingId field.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) SetShippingId(v int64) {
	o.ShippingId = &v
}

// GetLogisticName returns the LogisticName field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetLogisticName() string {
	if o == nil || IsNil(o.LogisticName) {
		var ret string
		return ret
	}
	return *o.LogisticName
}

// GetLogisticNameOk returns a tuple with the LogisticName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetLogisticNameOk() (*string, bool) {
	if o == nil || IsNil(o.LogisticName) {
		return nil, false
	}
	return o.LogisticName, true
}

// HasLogisticName returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) HasLogisticName() bool {
	if o != nil && !IsNil(o.LogisticName) {
		return true
	}

	return false
}

// SetLogisticName gets a reference to the given string and assigns it to the LogisticName field.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) SetLogisticName(v string) {
	o.LogisticName = &v
}

// GetLogisticService returns the LogisticService field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetLogisticService() string {
	if o == nil || IsNil(o.LogisticService) {
		var ret string
		return ret
	}
	return *o.LogisticService
}

// GetLogisticServiceOk returns a tuple with the LogisticService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetLogisticServiceOk() (*string, bool) {
	if o == nil || IsNil(o.LogisticService) {
		return nil, false
	}
	return o.LogisticService, true
}

// HasLogisticService returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) HasLogisticService() bool {
	if o != nil && !IsNil(o.LogisticService) {
		return true
	}

	return false
}

// SetLogisticService gets a reference to the given string and assigns it to the LogisticService field.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) SetLogisticService(v string) {
	o.LogisticService = &v
}

// GetShippingPrice returns the ShippingPrice field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetShippingPrice() int64 {
	if o == nil || IsNil(o.ShippingPrice) {
		var ret int64
		return ret
	}
	return *o.ShippingPrice
}

// GetShippingPriceOk returns a tuple with the ShippingPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetShippingPriceOk() (*int64, bool) {
	if o == nil || IsNil(o.ShippingPrice) {
		return nil, false
	}
	return o.ShippingPrice, true
}

// HasShippingPrice returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) HasShippingPrice() bool {
	if o != nil && !IsNil(o.ShippingPrice) {
		return true
	}

	return false
}

// SetShippingPrice gets a reference to the given int64 and assigns it to the ShippingPrice field.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) SetShippingPrice(v int64) {
	o.ShippingPrice = &v
}

// GetShippingPriceRate returns the ShippingPriceRate field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetShippingPriceRate() int64 {
	if o == nil || IsNil(o.ShippingPriceRate) {
		var ret int64
		return ret
	}
	return *o.ShippingPriceRate
}

// GetShippingPriceRateOk returns a tuple with the ShippingPriceRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetShippingPriceRateOk() (*int64, bool) {
	if o == nil || IsNil(o.ShippingPriceRate) {
		return nil, false
	}
	return o.ShippingPriceRate, true
}

// HasShippingPriceRate returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) HasShippingPriceRate() bool {
	if o != nil && !IsNil(o.ShippingPriceRate) {
		return true
	}

	return false
}

// SetShippingPriceRate gets a reference to the given int64 and assigns it to the ShippingPriceRate field.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) SetShippingPriceRate(v int64) {
	o.ShippingPriceRate = &v
}

// GetShippingFee returns the ShippingFee field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetShippingFee() int64 {
	if o == nil || IsNil(o.ShippingFee) {
		var ret int64
		return ret
	}
	return *o.ShippingFee
}

// GetShippingFeeOk returns a tuple with the ShippingFee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetShippingFeeOk() (*int64, bool) {
	if o == nil || IsNil(o.ShippingFee) {
		return nil, false
	}
	return o.ShippingFee, true
}

// HasShippingFee returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) HasShippingFee() bool {
	if o != nil && !IsNil(o.ShippingFee) {
		return true
	}

	return false
}

// SetShippingFee gets a reference to the given int64 and assigns it to the ShippingFee field.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) SetShippingFee(v int64) {
	o.ShippingFee = &v
}

// GetInsurancePrice returns the InsurancePrice field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetInsurancePrice() int64 {
	if o == nil || IsNil(o.InsurancePrice) {
		var ret int64
		return ret
	}
	return *o.InsurancePrice
}

// GetInsurancePriceOk returns a tuple with the InsurancePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetInsurancePriceOk() (*int64, bool) {
	if o == nil || IsNil(o.InsurancePrice) {
		return nil, false
	}
	return o.InsurancePrice, true
}

// HasInsurancePrice returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) HasInsurancePrice() bool {
	if o != nil && !IsNil(o.InsurancePrice) {
		return true
	}

	return false
}

// SetInsurancePrice gets a reference to the given int64 and assigns it to the InsurancePrice field.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) SetInsurancePrice(v int64) {
	o.InsurancePrice = &v
}

// GetFee returns the Fee field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetFee() int64 {
	if o == nil || IsNil(o.Fee) {
		var ret int64
		return ret
	}
	return *o.Fee
}

// GetFeeOk returns a tuple with the Fee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetFeeOk() (*int64, bool) {
	if o == nil || IsNil(o.Fee) {
		return nil, false
	}
	return o.Fee, true
}

// HasFee returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) HasFee() bool {
	if o != nil && !IsNil(o.Fee) {
		return true
	}

	return false
}

// SetFee gets a reference to the given int64 and assigns it to the Fee field.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) SetFee(v int64) {
	o.Fee = &v
}

// GetIsChangeCourier returns the IsChangeCourier field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetIsChangeCourier() bool {
	if o == nil || IsNil(o.IsChangeCourier) {
		var ret bool
		return ret
	}
	return *o.IsChangeCourier
}

// GetIsChangeCourierOk returns a tuple with the IsChangeCourier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetIsChangeCourierOk() (*bool, bool) {
	if o == nil || IsNil(o.IsChangeCourier) {
		return nil, false
	}
	return o.IsChangeCourier, true
}

// HasIsChangeCourier returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) HasIsChangeCourier() bool {
	if o != nil && !IsNil(o.IsChangeCourier) {
		return true
	}

	return false
}

// SetIsChangeCourier gets a reference to the given bool and assigns it to the IsChangeCourier field.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) SetIsChangeCourier(v bool) {
	o.IsChangeCourier = &v
}

// GetSecondSpId returns the SecondSpId field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetSecondSpId() int64 {
	if o == nil || IsNil(o.SecondSpId) {
		var ret int64
		return ret
	}
	return *o.SecondSpId
}

// GetSecondSpIdOk returns a tuple with the SecondSpId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetSecondSpIdOk() (*int64, bool) {
	if o == nil || IsNil(o.SecondSpId) {
		return nil, false
	}
	return o.SecondSpId, true
}

// HasSecondSpId returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) HasSecondSpId() bool {
	if o != nil && !IsNil(o.SecondSpId) {
		return true
	}

	return false
}

// SetSecondSpId gets a reference to the given int64 and assigns it to the SecondSpId field.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) SetSecondSpId(v int64) {
	o.SecondSpId = &v
}

// GetSecondShippingId returns the SecondShippingId field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetSecondShippingId() int64 {
	if o == nil || IsNil(o.SecondShippingId) {
		var ret int64
		return ret
	}
	return *o.SecondShippingId
}

// GetSecondShippingIdOk returns a tuple with the SecondShippingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetSecondShippingIdOk() (*int64, bool) {
	if o == nil || IsNil(o.SecondShippingId) {
		return nil, false
	}
	return o.SecondShippingId, true
}

// HasSecondShippingId returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) HasSecondShippingId() bool {
	if o != nil && !IsNil(o.SecondShippingId) {
		return true
	}

	return false
}

// SetSecondShippingId gets a reference to the given int64 and assigns it to the SecondShippingId field.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) SetSecondShippingId(v int64) {
	o.SecondShippingId = &v
}

// GetSecondLogisticName returns the SecondLogisticName field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetSecondLogisticName() string {
	if o == nil || IsNil(o.SecondLogisticName) {
		var ret string
		return ret
	}
	return *o.SecondLogisticName
}

// GetSecondLogisticNameOk returns a tuple with the SecondLogisticName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetSecondLogisticNameOk() (*string, bool) {
	if o == nil || IsNil(o.SecondLogisticName) {
		return nil, false
	}
	return o.SecondLogisticName, true
}

// HasSecondLogisticName returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) HasSecondLogisticName() bool {
	if o != nil && !IsNil(o.SecondLogisticName) {
		return true
	}

	return false
}

// SetSecondLogisticName gets a reference to the given string and assigns it to the SecondLogisticName field.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) SetSecondLogisticName(v string) {
	o.SecondLogisticName = &v
}

// GetSecondLogisticService returns the SecondLogisticService field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetSecondLogisticService() string {
	if o == nil || IsNil(o.SecondLogisticService) {
		var ret string
		return ret
	}
	return *o.SecondLogisticService
}

// GetSecondLogisticServiceOk returns a tuple with the SecondLogisticService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetSecondLogisticServiceOk() (*string, bool) {
	if o == nil || IsNil(o.SecondLogisticService) {
		return nil, false
	}
	return o.SecondLogisticService, true
}

// HasSecondLogisticService returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) HasSecondLogisticService() bool {
	if o != nil && !IsNil(o.SecondLogisticService) {
		return true
	}

	return false
}

// SetSecondLogisticService gets a reference to the given string and assigns it to the SecondLogisticService field.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) SetSecondLogisticService(v string) {
	o.SecondLogisticService = &v
}

// GetSecondAgencyFee returns the SecondAgencyFee field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetSecondAgencyFee() int64 {
	if o == nil || IsNil(o.SecondAgencyFee) {
		var ret int64
		return ret
	}
	return *o.SecondAgencyFee
}

// GetSecondAgencyFeeOk returns a tuple with the SecondAgencyFee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetSecondAgencyFeeOk() (*int64, bool) {
	if o == nil || IsNil(o.SecondAgencyFee) {
		return nil, false
	}
	return o.SecondAgencyFee, true
}

// HasSecondAgencyFee returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) HasSecondAgencyFee() bool {
	if o != nil && !IsNil(o.SecondAgencyFee) {
		return true
	}

	return false
}

// SetSecondAgencyFee gets a reference to the given int64 and assigns it to the SecondAgencyFee field.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) SetSecondAgencyFee(v int64) {
	o.SecondAgencyFee = &v
}

// GetSecondInsurance returns the SecondInsurance field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetSecondInsurance() int64 {
	if o == nil || IsNil(o.SecondInsurance) {
		var ret int64
		return ret
	}
	return *o.SecondInsurance
}

// GetSecondInsuranceOk returns a tuple with the SecondInsurance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetSecondInsuranceOk() (*int64, bool) {
	if o == nil || IsNil(o.SecondInsurance) {
		return nil, false
	}
	return o.SecondInsurance, true
}

// HasSecondInsurance returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) HasSecondInsurance() bool {
	if o != nil && !IsNil(o.SecondInsurance) {
		return true
	}

	return false
}

// SetSecondInsurance gets a reference to the given int64 and assigns it to the SecondInsurance field.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) SetSecondInsurance(v int64) {
	o.SecondInsurance = &v
}

// GetSecondRate returns the SecondRate field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetSecondRate() int64 {
	if o == nil || IsNil(o.SecondRate) {
		var ret int64
		return ret
	}
	return *o.SecondRate
}

// GetSecondRateOk returns a tuple with the SecondRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetSecondRateOk() (*int64, bool) {
	if o == nil || IsNil(o.SecondRate) {
		return nil, false
	}
	return o.SecondRate, true
}

// HasSecondRate returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) HasSecondRate() bool {
	if o != nil && !IsNil(o.SecondRate) {
		return true
	}

	return false
}

// SetSecondRate gets a reference to the given int64 and assigns it to the SecondRate field.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) SetSecondRate(v int64) {
	o.SecondRate = &v
}

// GetAwb returns the Awb field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetAwb() string {
	if o == nil || IsNil(o.Awb) {
		var ret string
		return ret
	}
	return *o.Awb
}

// GetAwbOk returns a tuple with the Awb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetAwbOk() (*string, bool) {
	if o == nil || IsNil(o.Awb) {
		return nil, false
	}
	return o.Awb, true
}

// HasAwb returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) HasAwb() bool {
	if o != nil && !IsNil(o.Awb) {
		return true
	}

	return false
}

// SetAwb gets a reference to the given string and assigns it to the Awb field.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) SetAwb(v string) {
	o.Awb = &v
}

// GetAutoresiCashlessStatus returns the AutoresiCashlessStatus field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetAutoresiCashlessStatus() int64 {
	if o == nil || IsNil(o.AutoresiCashlessStatus) {
		var ret int64
		return ret
	}
	return *o.AutoresiCashlessStatus
}

// GetAutoresiCashlessStatusOk returns a tuple with the AutoresiCashlessStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetAutoresiCashlessStatusOk() (*int64, bool) {
	if o == nil || IsNil(o.AutoresiCashlessStatus) {
		return nil, false
	}
	return o.AutoresiCashlessStatus, true
}

// HasAutoresiCashlessStatus returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) HasAutoresiCashlessStatus() bool {
	if o != nil && !IsNil(o.AutoresiCashlessStatus) {
		return true
	}

	return false
}

// SetAutoresiCashlessStatus gets a reference to the given int64 and assigns it to the AutoresiCashlessStatus field.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) SetAutoresiCashlessStatus(v int64) {
	o.AutoresiCashlessStatus = &v
}

// GetAutoresiAwb returns the AutoresiAwb field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetAutoresiAwb() string {
	if o == nil || IsNil(o.AutoresiAwb) {
		var ret string
		return ret
	}
	return *o.AutoresiAwb
}

// GetAutoresiAwbOk returns a tuple with the AutoresiAwb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetAutoresiAwbOk() (*string, bool) {
	if o == nil || IsNil(o.AutoresiAwb) {
		return nil, false
	}
	return o.AutoresiAwb, true
}

// HasAutoresiAwb returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) HasAutoresiAwb() bool {
	if o != nil && !IsNil(o.AutoresiAwb) {
		return true
	}

	return false
}

// SetAutoresiAwb gets a reference to the given string and assigns it to the AutoresiAwb field.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) SetAutoresiAwb(v string) {
	o.AutoresiAwb = &v
}

// GetAutoresiShippingPrice returns the AutoresiShippingPrice field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetAutoresiShippingPrice() int64 {
	if o == nil || IsNil(o.AutoresiShippingPrice) {
		var ret int64
		return ret
	}
	return *o.AutoresiShippingPrice
}

// GetAutoresiShippingPriceOk returns a tuple with the AutoresiShippingPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetAutoresiShippingPriceOk() (*int64, bool) {
	if o == nil || IsNil(o.AutoresiShippingPrice) {
		return nil, false
	}
	return o.AutoresiShippingPrice, true
}

// HasAutoresiShippingPrice returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) HasAutoresiShippingPrice() bool {
	if o != nil && !IsNil(o.AutoresiShippingPrice) {
		return true
	}

	return false
}

// SetAutoresiShippingPrice gets a reference to the given int64 and assigns it to the AutoresiShippingPrice field.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) SetAutoresiShippingPrice(v int64) {
	o.AutoresiShippingPrice = &v
}

// GetCountAwb returns the CountAwb field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetCountAwb() int64 {
	if o == nil || IsNil(o.CountAwb) {
		var ret int64
		return ret
	}
	return *o.CountAwb
}

// GetCountAwbOk returns a tuple with the CountAwb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetCountAwbOk() (*int64, bool) {
	if o == nil || IsNil(o.CountAwb) {
		return nil, false
	}
	return o.CountAwb, true
}

// HasCountAwb returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) HasCountAwb() bool {
	if o != nil && !IsNil(o.CountAwb) {
		return true
	}

	return false
}

// SetCountAwb gets a reference to the given int64 and assigns it to the CountAwb field.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) SetCountAwb(v int64) {
	o.CountAwb = &v
}

// GetIsCashless returns the IsCashless field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetIsCashless() bool {
	if o == nil || IsNil(o.IsCashless) {
		var ret bool
		return ret
	}
	return *o.IsCashless
}

// GetIsCashlessOk returns a tuple with the IsCashless field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetIsCashlessOk() (*bool, bool) {
	if o == nil || IsNil(o.IsCashless) {
		return nil, false
	}
	return o.IsCashless, true
}

// HasIsCashless returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) HasIsCashless() bool {
	if o != nil && !IsNil(o.IsCashless) {
		return true
	}

	return false
}

// SetIsCashless gets a reference to the given bool and assigns it to the IsCashless field.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) SetIsCashless(v bool) {
	o.IsCashless = &v
}

// GetIsFakeDelivery returns the IsFakeDelivery field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetIsFakeDelivery() bool {
	if o == nil || IsNil(o.IsFakeDelivery) {
		var ret bool
		return ret
	}
	return *o.IsFakeDelivery
}

// GetIsFakeDeliveryOk returns a tuple with the IsFakeDelivery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetIsFakeDeliveryOk() (*bool, bool) {
	if o == nil || IsNil(o.IsFakeDelivery) {
		return nil, false
	}
	return o.IsFakeDelivery, true
}

// HasIsFakeDelivery returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) HasIsFakeDelivery() bool {
	if o != nil && !IsNil(o.IsFakeDelivery) {
		return true
	}

	return false
}

// SetIsFakeDelivery gets a reference to the given bool and assigns it to the IsFakeDelivery field.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) SetIsFakeDelivery(v bool) {
	o.IsFakeDelivery = &v
}

// GetRecommendedCourierInfo returns the RecommendedCourierInfo field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetRecommendedCourierInfo() []GetSingleOrder200ResponseDataOrderInfoShippingInfoRecommendedCourierInfoInner {
	if o == nil || IsNil(o.RecommendedCourierInfo) {
		var ret []GetSingleOrder200ResponseDataOrderInfoShippingInfoRecommendedCourierInfoInner
		return ret
	}
	return o.RecommendedCourierInfo
}

// GetRecommendedCourierInfoOk returns a tuple with the RecommendedCourierInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetRecommendedCourierInfoOk() ([]GetSingleOrder200ResponseDataOrderInfoShippingInfoRecommendedCourierInfoInner, bool) {
	if o == nil || IsNil(o.RecommendedCourierInfo) {
		return nil, false
	}
	return o.RecommendedCourierInfo, true
}

// HasRecommendedCourierInfo returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) HasRecommendedCourierInfo() bool {
	if o != nil && !IsNil(o.RecommendedCourierInfo) {
		return true
	}

	return false
}

// SetRecommendedCourierInfo gets a reference to the given []GetSingleOrder200ResponseDataOrderInfoShippingInfoRecommendedCourierInfoInner and assigns it to the RecommendedCourierInfo field.
func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) SetRecommendedCourierInfo(v []GetSingleOrder200ResponseDataOrderInfoShippingInfoRecommendedCourierInfoInner) {
	o.RecommendedCourierInfo = v
}

func (o GetSingleOrder200ResponseDataOrderInfoShippingInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSingleOrder200ResponseDataOrderInfoShippingInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SpId) {
		toSerialize["sp_id"] = o.SpId
	}
	if !IsNil(o.ShippingId) {
		toSerialize["shipping_id"] = o.ShippingId
	}
	if !IsNil(o.LogisticName) {
		toSerialize["logistic_name"] = o.LogisticName
	}
	if !IsNil(o.LogisticService) {
		toSerialize["logistic_service"] = o.LogisticService
	}
	if !IsNil(o.ShippingPrice) {
		toSerialize["shipping_price"] = o.ShippingPrice
	}
	if !IsNil(o.ShippingPriceRate) {
		toSerialize["shipping_price_rate"] = o.ShippingPriceRate
	}
	if !IsNil(o.ShippingFee) {
		toSerialize["shipping_fee"] = o.ShippingFee
	}
	if !IsNil(o.InsurancePrice) {
		toSerialize["insurance_price"] = o.InsurancePrice
	}
	if !IsNil(o.Fee) {
		toSerialize["fee"] = o.Fee
	}
	if !IsNil(o.IsChangeCourier) {
		toSerialize["is_change_courier"] = o.IsChangeCourier
	}
	if !IsNil(o.SecondSpId) {
		toSerialize["second_sp_id"] = o.SecondSpId
	}
	if !IsNil(o.SecondShippingId) {
		toSerialize["second_shipping_id"] = o.SecondShippingId
	}
	if !IsNil(o.SecondLogisticName) {
		toSerialize["second_logistic_name"] = o.SecondLogisticName
	}
	if !IsNil(o.SecondLogisticService) {
		toSerialize["second_logistic_service"] = o.SecondLogisticService
	}
	if !IsNil(o.SecondAgencyFee) {
		toSerialize["second_agency_fee"] = o.SecondAgencyFee
	}
	if !IsNil(o.SecondInsurance) {
		toSerialize["second_insurance"] = o.SecondInsurance
	}
	if !IsNil(o.SecondRate) {
		toSerialize["second_rate"] = o.SecondRate
	}
	if !IsNil(o.Awb) {
		toSerialize["awb"] = o.Awb
	}
	if !IsNil(o.AutoresiCashlessStatus) {
		toSerialize["autoresi_cashless_status"] = o.AutoresiCashlessStatus
	}
	if !IsNil(o.AutoresiAwb) {
		toSerialize["autoresi_awb"] = o.AutoresiAwb
	}
	if !IsNil(o.AutoresiShippingPrice) {
		toSerialize["autoresi_shipping_price"] = o.AutoresiShippingPrice
	}
	if !IsNil(o.CountAwb) {
		toSerialize["count_awb"] = o.CountAwb
	}
	if !IsNil(o.IsCashless) {
		toSerialize["isCashless"] = o.IsCashless
	}
	if !IsNil(o.IsFakeDelivery) {
		toSerialize["is_fake_delivery"] = o.IsFakeDelivery
	}
	if !IsNil(o.RecommendedCourierInfo) {
		toSerialize["recommended_courier_info"] = o.RecommendedCourierInfo
	}
	return toSerialize, nil
}

type NullableGetSingleOrder200ResponseDataOrderInfoShippingInfo struct {
	value *GetSingleOrder200ResponseDataOrderInfoShippingInfo
	isSet bool
}

func (v NullableGetSingleOrder200ResponseDataOrderInfoShippingInfo) Get() *GetSingleOrder200ResponseDataOrderInfoShippingInfo {
	return v.value
}

func (v *NullableGetSingleOrder200ResponseDataOrderInfoShippingInfo) Set(val *GetSingleOrder200ResponseDataOrderInfoShippingInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSingleOrder200ResponseDataOrderInfoShippingInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSingleOrder200ResponseDataOrderInfoShippingInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSingleOrder200ResponseDataOrderInfoShippingInfo(val *GetSingleOrder200ResponseDataOrderInfoShippingInfo) *NullableGetSingleOrder200ResponseDataOrderInfoShippingInfo {
	return &NullableGetSingleOrder200ResponseDataOrderInfoShippingInfo{value: val, isSet: true}
}

func (v NullableGetSingleOrder200ResponseDataOrderInfoShippingInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSingleOrder200ResponseDataOrderInfoShippingInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


