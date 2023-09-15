# OrderInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderDate** | **string** |  | 
**OrderId** | **string** |  | 
**CustomerReference** | **string** |  | 
**CurrencyCode** | **string** |  | 
**ExchangeRate** | Pointer to **NullableFloat64** |  | [optional] 
**WarehouseId** | **string** |  | 
**PaymentMethod** | **string** |  | 
**DropshipperName** | Pointer to **string** |  | [optional] 
**DriverName** | Pointer to **string** |  | [optional] 
**DriverPhoneNumber** | Pointer to **string** |  | [optional] 
**DeliveryLatitude** | Pointer to **string** |  | [optional] 
**DeliveryLongitude** | Pointer to **string** |  | [optional] 
**DeliveryName** | **string** |  | 
**DeliveryStreetAddress** | **string** |  | 
**DeliveryStreetAddress2** | Pointer to **string** |  | [optional] 
**DeliverySuburb** | **string** |  | 
**DeliveryCity** | **string** |  | 
**DeliveryRegion** | **string** |  | 
**DeliveryPostCode** | **string** |  | 
**DeliveryCountry** | **string** |  | 
**DeliveryMethod** | **string** |  | 
**OrderStatus** | **string** |  | 
**TaxCode** | Pointer to **string** |  | [optional] 
**TaxRate** | Pointer to **NullableFloat64** |  | [optional] 
**TaxTotal** | **float64** |  | 
**ShippingTotal** | **float64** |  | 
**DiscountTotal** | **float64** |  | 
**TotalPromo** | **float64** |  | 
**TotalVoucher** | **float64** |  | 
**Subtotal** | **float64** |  | 
**Total** | **float64** |  | 
**Remarks** | Pointer to **string** |  | [optional] 
**PhoneNumber** | **string** |  | 
**AirwaybillNumber** | Pointer to **string** |  | [optional] 
**PaidAt** | Pointer to **string** |  | [optional] 
**AcceptedAt** | Pointer to **string** |  | [optional] 
**PackedAt** | Pointer to **string** |  | [optional] 
**AwbRetrievedAt** | Pointer to **string** |  | [optional] 
**CompletedAt** | Pointer to **string** |  | [optional] 
**ReceivedAt** | Pointer to **string** |  | [optional] 
**ReceivedBy** | Pointer to **string** |  | [optional] 
**CancelledAt** | Pointer to **string** |  | [optional] 
**ShopId** | **string** |  | 
**CancelReason** | Pointer to **string** |  | [optional] 
**Provider** | **string** |  | 
**Service** | **string** |  | 
**ShipmentExtras** | Pointer to **string** |  | [optional] 
**IsAwbAvailableFromChannel** | **bool** |  | 
**LineItems** | [**[]LineItemInput**](LineItemInput.md) |  | 

## Methods

### NewOrderInput

`func NewOrderInput(orderDate string, orderId string, customerReference string, currencyCode string, warehouseId string, paymentMethod string, deliveryName string, deliveryStreetAddress string, deliverySuburb string, deliveryCity string, deliveryRegion string, deliveryPostCode string, deliveryCountry string, deliveryMethod string, orderStatus string, taxTotal float64, shippingTotal float64, discountTotal float64, totalPromo float64, totalVoucher float64, subtotal float64, total float64, phoneNumber string, shopId string, provider string, service string, isAwbAvailableFromChannel bool, lineItems []LineItemInput, ) *OrderInput`

NewOrderInput instantiates a new OrderInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderInputWithDefaults

`func NewOrderInputWithDefaults() *OrderInput`

NewOrderInputWithDefaults instantiates a new OrderInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderDate

`func (o *OrderInput) GetOrderDate() string`

GetOrderDate returns the OrderDate field if non-nil, zero value otherwise.

### GetOrderDateOk

`func (o *OrderInput) GetOrderDateOk() (*string, bool)`

GetOrderDateOk returns a tuple with the OrderDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDate

`func (o *OrderInput) SetOrderDate(v string)`

SetOrderDate sets OrderDate field to given value.


### GetOrderId

`func (o *OrderInput) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *OrderInput) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *OrderInput) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.


### GetCustomerReference

`func (o *OrderInput) GetCustomerReference() string`

GetCustomerReference returns the CustomerReference field if non-nil, zero value otherwise.

### GetCustomerReferenceOk

`func (o *OrderInput) GetCustomerReferenceOk() (*string, bool)`

GetCustomerReferenceOk returns a tuple with the CustomerReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerReference

`func (o *OrderInput) SetCustomerReference(v string)`

SetCustomerReference sets CustomerReference field to given value.


### GetCurrencyCode

`func (o *OrderInput) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *OrderInput) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *OrderInput) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetExchangeRate

`func (o *OrderInput) GetExchangeRate() float64`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *OrderInput) GetExchangeRateOk() (*float64, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *OrderInput) SetExchangeRate(v float64)`

SetExchangeRate sets ExchangeRate field to given value.

### HasExchangeRate

`func (o *OrderInput) HasExchangeRate() bool`

HasExchangeRate returns a boolean if a field has been set.

### SetExchangeRateNil

`func (o *OrderInput) SetExchangeRateNil(b bool)`

 SetExchangeRateNil sets the value for ExchangeRate to be an explicit nil

### UnsetExchangeRate
`func (o *OrderInput) UnsetExchangeRate()`

UnsetExchangeRate ensures that no value is present for ExchangeRate, not even an explicit nil
### GetWarehouseId

`func (o *OrderInput) GetWarehouseId() string`

GetWarehouseId returns the WarehouseId field if non-nil, zero value otherwise.

### GetWarehouseIdOk

`func (o *OrderInput) GetWarehouseIdOk() (*string, bool)`

GetWarehouseIdOk returns a tuple with the WarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseId

`func (o *OrderInput) SetWarehouseId(v string)`

SetWarehouseId sets WarehouseId field to given value.


### GetPaymentMethod

`func (o *OrderInput) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *OrderInput) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *OrderInput) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.


### GetDropshipperName

`func (o *OrderInput) GetDropshipperName() string`

GetDropshipperName returns the DropshipperName field if non-nil, zero value otherwise.

### GetDropshipperNameOk

`func (o *OrderInput) GetDropshipperNameOk() (*string, bool)`

GetDropshipperNameOk returns a tuple with the DropshipperName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropshipperName

`func (o *OrderInput) SetDropshipperName(v string)`

SetDropshipperName sets DropshipperName field to given value.

### HasDropshipperName

`func (o *OrderInput) HasDropshipperName() bool`

HasDropshipperName returns a boolean if a field has been set.

### GetDriverName

`func (o *OrderInput) GetDriverName() string`

GetDriverName returns the DriverName field if non-nil, zero value otherwise.

### GetDriverNameOk

`func (o *OrderInput) GetDriverNameOk() (*string, bool)`

GetDriverNameOk returns a tuple with the DriverName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverName

`func (o *OrderInput) SetDriverName(v string)`

SetDriverName sets DriverName field to given value.

### HasDriverName

`func (o *OrderInput) HasDriverName() bool`

HasDriverName returns a boolean if a field has been set.

### GetDriverPhoneNumber

`func (o *OrderInput) GetDriverPhoneNumber() string`

GetDriverPhoneNumber returns the DriverPhoneNumber field if non-nil, zero value otherwise.

### GetDriverPhoneNumberOk

`func (o *OrderInput) GetDriverPhoneNumberOk() (*string, bool)`

GetDriverPhoneNumberOk returns a tuple with the DriverPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverPhoneNumber

`func (o *OrderInput) SetDriverPhoneNumber(v string)`

SetDriverPhoneNumber sets DriverPhoneNumber field to given value.

### HasDriverPhoneNumber

`func (o *OrderInput) HasDriverPhoneNumber() bool`

HasDriverPhoneNumber returns a boolean if a field has been set.

### GetDeliveryLatitude

`func (o *OrderInput) GetDeliveryLatitude() string`

GetDeliveryLatitude returns the DeliveryLatitude field if non-nil, zero value otherwise.

### GetDeliveryLatitudeOk

`func (o *OrderInput) GetDeliveryLatitudeOk() (*string, bool)`

GetDeliveryLatitudeOk returns a tuple with the DeliveryLatitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryLatitude

`func (o *OrderInput) SetDeliveryLatitude(v string)`

SetDeliveryLatitude sets DeliveryLatitude field to given value.

### HasDeliveryLatitude

`func (o *OrderInput) HasDeliveryLatitude() bool`

HasDeliveryLatitude returns a boolean if a field has been set.

### GetDeliveryLongitude

`func (o *OrderInput) GetDeliveryLongitude() string`

GetDeliveryLongitude returns the DeliveryLongitude field if non-nil, zero value otherwise.

### GetDeliveryLongitudeOk

`func (o *OrderInput) GetDeliveryLongitudeOk() (*string, bool)`

GetDeliveryLongitudeOk returns a tuple with the DeliveryLongitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryLongitude

`func (o *OrderInput) SetDeliveryLongitude(v string)`

SetDeliveryLongitude sets DeliveryLongitude field to given value.

### HasDeliveryLongitude

`func (o *OrderInput) HasDeliveryLongitude() bool`

HasDeliveryLongitude returns a boolean if a field has been set.

### GetDeliveryName

`func (o *OrderInput) GetDeliveryName() string`

GetDeliveryName returns the DeliveryName field if non-nil, zero value otherwise.

### GetDeliveryNameOk

`func (o *OrderInput) GetDeliveryNameOk() (*string, bool)`

GetDeliveryNameOk returns a tuple with the DeliveryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryName

`func (o *OrderInput) SetDeliveryName(v string)`

SetDeliveryName sets DeliveryName field to given value.


### GetDeliveryStreetAddress

`func (o *OrderInput) GetDeliveryStreetAddress() string`

GetDeliveryStreetAddress returns the DeliveryStreetAddress field if non-nil, zero value otherwise.

### GetDeliveryStreetAddressOk

`func (o *OrderInput) GetDeliveryStreetAddressOk() (*string, bool)`

GetDeliveryStreetAddressOk returns a tuple with the DeliveryStreetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryStreetAddress

`func (o *OrderInput) SetDeliveryStreetAddress(v string)`

SetDeliveryStreetAddress sets DeliveryStreetAddress field to given value.


### GetDeliveryStreetAddress2

`func (o *OrderInput) GetDeliveryStreetAddress2() string`

GetDeliveryStreetAddress2 returns the DeliveryStreetAddress2 field if non-nil, zero value otherwise.

### GetDeliveryStreetAddress2Ok

`func (o *OrderInput) GetDeliveryStreetAddress2Ok() (*string, bool)`

GetDeliveryStreetAddress2Ok returns a tuple with the DeliveryStreetAddress2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryStreetAddress2

`func (o *OrderInput) SetDeliveryStreetAddress2(v string)`

SetDeliveryStreetAddress2 sets DeliveryStreetAddress2 field to given value.

### HasDeliveryStreetAddress2

`func (o *OrderInput) HasDeliveryStreetAddress2() bool`

HasDeliveryStreetAddress2 returns a boolean if a field has been set.

### GetDeliverySuburb

`func (o *OrderInput) GetDeliverySuburb() string`

GetDeliverySuburb returns the DeliverySuburb field if non-nil, zero value otherwise.

### GetDeliverySuburbOk

`func (o *OrderInput) GetDeliverySuburbOk() (*string, bool)`

GetDeliverySuburbOk returns a tuple with the DeliverySuburb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliverySuburb

`func (o *OrderInput) SetDeliverySuburb(v string)`

SetDeliverySuburb sets DeliverySuburb field to given value.


### GetDeliveryCity

`func (o *OrderInput) GetDeliveryCity() string`

GetDeliveryCity returns the DeliveryCity field if non-nil, zero value otherwise.

### GetDeliveryCityOk

`func (o *OrderInput) GetDeliveryCityOk() (*string, bool)`

GetDeliveryCityOk returns a tuple with the DeliveryCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryCity

`func (o *OrderInput) SetDeliveryCity(v string)`

SetDeliveryCity sets DeliveryCity field to given value.


### GetDeliveryRegion

`func (o *OrderInput) GetDeliveryRegion() string`

GetDeliveryRegion returns the DeliveryRegion field if non-nil, zero value otherwise.

### GetDeliveryRegionOk

`func (o *OrderInput) GetDeliveryRegionOk() (*string, bool)`

GetDeliveryRegionOk returns a tuple with the DeliveryRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryRegion

`func (o *OrderInput) SetDeliveryRegion(v string)`

SetDeliveryRegion sets DeliveryRegion field to given value.


### GetDeliveryPostCode

`func (o *OrderInput) GetDeliveryPostCode() string`

GetDeliveryPostCode returns the DeliveryPostCode field if non-nil, zero value otherwise.

### GetDeliveryPostCodeOk

`func (o *OrderInput) GetDeliveryPostCodeOk() (*string, bool)`

GetDeliveryPostCodeOk returns a tuple with the DeliveryPostCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryPostCode

`func (o *OrderInput) SetDeliveryPostCode(v string)`

SetDeliveryPostCode sets DeliveryPostCode field to given value.


### GetDeliveryCountry

`func (o *OrderInput) GetDeliveryCountry() string`

GetDeliveryCountry returns the DeliveryCountry field if non-nil, zero value otherwise.

### GetDeliveryCountryOk

`func (o *OrderInput) GetDeliveryCountryOk() (*string, bool)`

GetDeliveryCountryOk returns a tuple with the DeliveryCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryCountry

`func (o *OrderInput) SetDeliveryCountry(v string)`

SetDeliveryCountry sets DeliveryCountry field to given value.


### GetDeliveryMethod

`func (o *OrderInput) GetDeliveryMethod() string`

GetDeliveryMethod returns the DeliveryMethod field if non-nil, zero value otherwise.

### GetDeliveryMethodOk

`func (o *OrderInput) GetDeliveryMethodOk() (*string, bool)`

GetDeliveryMethodOk returns a tuple with the DeliveryMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryMethod

`func (o *OrderInput) SetDeliveryMethod(v string)`

SetDeliveryMethod sets DeliveryMethod field to given value.


### GetOrderStatus

`func (o *OrderInput) GetOrderStatus() string`

GetOrderStatus returns the OrderStatus field if non-nil, zero value otherwise.

### GetOrderStatusOk

`func (o *OrderInput) GetOrderStatusOk() (*string, bool)`

GetOrderStatusOk returns a tuple with the OrderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderStatus

`func (o *OrderInput) SetOrderStatus(v string)`

SetOrderStatus sets OrderStatus field to given value.


### GetTaxCode

`func (o *OrderInput) GetTaxCode() string`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *OrderInput) GetTaxCodeOk() (*string, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *OrderInput) SetTaxCode(v string)`

SetTaxCode sets TaxCode field to given value.

### HasTaxCode

`func (o *OrderInput) HasTaxCode() bool`

HasTaxCode returns a boolean if a field has been set.

### GetTaxRate

`func (o *OrderInput) GetTaxRate() float64`

GetTaxRate returns the TaxRate field if non-nil, zero value otherwise.

### GetTaxRateOk

`func (o *OrderInput) GetTaxRateOk() (*float64, bool)`

GetTaxRateOk returns a tuple with the TaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate

`func (o *OrderInput) SetTaxRate(v float64)`

SetTaxRate sets TaxRate field to given value.

### HasTaxRate

`func (o *OrderInput) HasTaxRate() bool`

HasTaxRate returns a boolean if a field has been set.

### SetTaxRateNil

`func (o *OrderInput) SetTaxRateNil(b bool)`

 SetTaxRateNil sets the value for TaxRate to be an explicit nil

### UnsetTaxRate
`func (o *OrderInput) UnsetTaxRate()`

UnsetTaxRate ensures that no value is present for TaxRate, not even an explicit nil
### GetTaxTotal

`func (o *OrderInput) GetTaxTotal() float64`

GetTaxTotal returns the TaxTotal field if non-nil, zero value otherwise.

### GetTaxTotalOk

`func (o *OrderInput) GetTaxTotalOk() (*float64, bool)`

GetTaxTotalOk returns a tuple with the TaxTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxTotal

`func (o *OrderInput) SetTaxTotal(v float64)`

SetTaxTotal sets TaxTotal field to given value.


### GetShippingTotal

`func (o *OrderInput) GetShippingTotal() float64`

GetShippingTotal returns the ShippingTotal field if non-nil, zero value otherwise.

### GetShippingTotalOk

`func (o *OrderInput) GetShippingTotalOk() (*float64, bool)`

GetShippingTotalOk returns a tuple with the ShippingTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingTotal

`func (o *OrderInput) SetShippingTotal(v float64)`

SetShippingTotal sets ShippingTotal field to given value.


### GetDiscountTotal

`func (o *OrderInput) GetDiscountTotal() float64`

GetDiscountTotal returns the DiscountTotal field if non-nil, zero value otherwise.

### GetDiscountTotalOk

`func (o *OrderInput) GetDiscountTotalOk() (*float64, bool)`

GetDiscountTotalOk returns a tuple with the DiscountTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountTotal

`func (o *OrderInput) SetDiscountTotal(v float64)`

SetDiscountTotal sets DiscountTotal field to given value.


### GetTotalPromo

`func (o *OrderInput) GetTotalPromo() float64`

GetTotalPromo returns the TotalPromo field if non-nil, zero value otherwise.

### GetTotalPromoOk

`func (o *OrderInput) GetTotalPromoOk() (*float64, bool)`

GetTotalPromoOk returns a tuple with the TotalPromo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPromo

`func (o *OrderInput) SetTotalPromo(v float64)`

SetTotalPromo sets TotalPromo field to given value.


### GetTotalVoucher

`func (o *OrderInput) GetTotalVoucher() float64`

GetTotalVoucher returns the TotalVoucher field if non-nil, zero value otherwise.

### GetTotalVoucherOk

`func (o *OrderInput) GetTotalVoucherOk() (*float64, bool)`

GetTotalVoucherOk returns a tuple with the TotalVoucher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalVoucher

`func (o *OrderInput) SetTotalVoucher(v float64)`

SetTotalVoucher sets TotalVoucher field to given value.


### GetSubtotal

`func (o *OrderInput) GetSubtotal() float64`

GetSubtotal returns the Subtotal field if non-nil, zero value otherwise.

### GetSubtotalOk

`func (o *OrderInput) GetSubtotalOk() (*float64, bool)`

GetSubtotalOk returns a tuple with the Subtotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotal

`func (o *OrderInput) SetSubtotal(v float64)`

SetSubtotal sets Subtotal field to given value.


### GetTotal

`func (o *OrderInput) GetTotal() float64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *OrderInput) GetTotalOk() (*float64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *OrderInput) SetTotal(v float64)`

SetTotal sets Total field to given value.


### GetRemarks

`func (o *OrderInput) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *OrderInput) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *OrderInput) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *OrderInput) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *OrderInput) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *OrderInput) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *OrderInput) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.


### GetAirwaybillNumber

`func (o *OrderInput) GetAirwaybillNumber() string`

GetAirwaybillNumber returns the AirwaybillNumber field if non-nil, zero value otherwise.

### GetAirwaybillNumberOk

`func (o *OrderInput) GetAirwaybillNumberOk() (*string, bool)`

GetAirwaybillNumberOk returns a tuple with the AirwaybillNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirwaybillNumber

`func (o *OrderInput) SetAirwaybillNumber(v string)`

SetAirwaybillNumber sets AirwaybillNumber field to given value.

### HasAirwaybillNumber

`func (o *OrderInput) HasAirwaybillNumber() bool`

HasAirwaybillNumber returns a boolean if a field has been set.

### GetPaidAt

`func (o *OrderInput) GetPaidAt() string`

GetPaidAt returns the PaidAt field if non-nil, zero value otherwise.

### GetPaidAtOk

`func (o *OrderInput) GetPaidAtOk() (*string, bool)`

GetPaidAtOk returns a tuple with the PaidAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidAt

`func (o *OrderInput) SetPaidAt(v string)`

SetPaidAt sets PaidAt field to given value.

### HasPaidAt

`func (o *OrderInput) HasPaidAt() bool`

HasPaidAt returns a boolean if a field has been set.

### GetAcceptedAt

`func (o *OrderInput) GetAcceptedAt() string`

GetAcceptedAt returns the AcceptedAt field if non-nil, zero value otherwise.

### GetAcceptedAtOk

`func (o *OrderInput) GetAcceptedAtOk() (*string, bool)`

GetAcceptedAtOk returns a tuple with the AcceptedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedAt

`func (o *OrderInput) SetAcceptedAt(v string)`

SetAcceptedAt sets AcceptedAt field to given value.

### HasAcceptedAt

`func (o *OrderInput) HasAcceptedAt() bool`

HasAcceptedAt returns a boolean if a field has been set.

### GetPackedAt

`func (o *OrderInput) GetPackedAt() string`

GetPackedAt returns the PackedAt field if non-nil, zero value otherwise.

### GetPackedAtOk

`func (o *OrderInput) GetPackedAtOk() (*string, bool)`

GetPackedAtOk returns a tuple with the PackedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackedAt

`func (o *OrderInput) SetPackedAt(v string)`

SetPackedAt sets PackedAt field to given value.

### HasPackedAt

`func (o *OrderInput) HasPackedAt() bool`

HasPackedAt returns a boolean if a field has been set.

### GetAwbRetrievedAt

`func (o *OrderInput) GetAwbRetrievedAt() string`

GetAwbRetrievedAt returns the AwbRetrievedAt field if non-nil, zero value otherwise.

### GetAwbRetrievedAtOk

`func (o *OrderInput) GetAwbRetrievedAtOk() (*string, bool)`

GetAwbRetrievedAtOk returns a tuple with the AwbRetrievedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwbRetrievedAt

`func (o *OrderInput) SetAwbRetrievedAt(v string)`

SetAwbRetrievedAt sets AwbRetrievedAt field to given value.

### HasAwbRetrievedAt

`func (o *OrderInput) HasAwbRetrievedAt() bool`

HasAwbRetrievedAt returns a boolean if a field has been set.

### GetCompletedAt

`func (o *OrderInput) GetCompletedAt() string`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *OrderInput) GetCompletedAtOk() (*string, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *OrderInput) SetCompletedAt(v string)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *OrderInput) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetReceivedAt

`func (o *OrderInput) GetReceivedAt() string`

GetReceivedAt returns the ReceivedAt field if non-nil, zero value otherwise.

### GetReceivedAtOk

`func (o *OrderInput) GetReceivedAtOk() (*string, bool)`

GetReceivedAtOk returns a tuple with the ReceivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedAt

`func (o *OrderInput) SetReceivedAt(v string)`

SetReceivedAt sets ReceivedAt field to given value.

### HasReceivedAt

`func (o *OrderInput) HasReceivedAt() bool`

HasReceivedAt returns a boolean if a field has been set.

### GetReceivedBy

`func (o *OrderInput) GetReceivedBy() string`

GetReceivedBy returns the ReceivedBy field if non-nil, zero value otherwise.

### GetReceivedByOk

`func (o *OrderInput) GetReceivedByOk() (*string, bool)`

GetReceivedByOk returns a tuple with the ReceivedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedBy

`func (o *OrderInput) SetReceivedBy(v string)`

SetReceivedBy sets ReceivedBy field to given value.

### HasReceivedBy

`func (o *OrderInput) HasReceivedBy() bool`

HasReceivedBy returns a boolean if a field has been set.

### GetCancelledAt

`func (o *OrderInput) GetCancelledAt() string`

GetCancelledAt returns the CancelledAt field if non-nil, zero value otherwise.

### GetCancelledAtOk

`func (o *OrderInput) GetCancelledAtOk() (*string, bool)`

GetCancelledAtOk returns a tuple with the CancelledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelledAt

`func (o *OrderInput) SetCancelledAt(v string)`

SetCancelledAt sets CancelledAt field to given value.

### HasCancelledAt

`func (o *OrderInput) HasCancelledAt() bool`

HasCancelledAt returns a boolean if a field has been set.

### GetShopId

`func (o *OrderInput) GetShopId() string`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *OrderInput) GetShopIdOk() (*string, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *OrderInput) SetShopId(v string)`

SetShopId sets ShopId field to given value.


### GetCancelReason

`func (o *OrderInput) GetCancelReason() string`

GetCancelReason returns the CancelReason field if non-nil, zero value otherwise.

### GetCancelReasonOk

`func (o *OrderInput) GetCancelReasonOk() (*string, bool)`

GetCancelReasonOk returns a tuple with the CancelReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelReason

`func (o *OrderInput) SetCancelReason(v string)`

SetCancelReason sets CancelReason field to given value.

### HasCancelReason

`func (o *OrderInput) HasCancelReason() bool`

HasCancelReason returns a boolean if a field has been set.

### GetProvider

`func (o *OrderInput) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *OrderInput) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *OrderInput) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetService

`func (o *OrderInput) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *OrderInput) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *OrderInput) SetService(v string)`

SetService sets Service field to given value.


### GetShipmentExtras

`func (o *OrderInput) GetShipmentExtras() string`

GetShipmentExtras returns the ShipmentExtras field if non-nil, zero value otherwise.

### GetShipmentExtrasOk

`func (o *OrderInput) GetShipmentExtrasOk() (*string, bool)`

GetShipmentExtrasOk returns a tuple with the ShipmentExtras field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentExtras

`func (o *OrderInput) SetShipmentExtras(v string)`

SetShipmentExtras sets ShipmentExtras field to given value.

### HasShipmentExtras

`func (o *OrderInput) HasShipmentExtras() bool`

HasShipmentExtras returns a boolean if a field has been set.

### GetIsAwbAvailableFromChannel

`func (o *OrderInput) GetIsAwbAvailableFromChannel() bool`

GetIsAwbAvailableFromChannel returns the IsAwbAvailableFromChannel field if non-nil, zero value otherwise.

### GetIsAwbAvailableFromChannelOk

`func (o *OrderInput) GetIsAwbAvailableFromChannelOk() (*bool, bool)`

GetIsAwbAvailableFromChannelOk returns a tuple with the IsAwbAvailableFromChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAwbAvailableFromChannel

`func (o *OrderInput) SetIsAwbAvailableFromChannel(v bool)`

SetIsAwbAvailableFromChannel sets IsAwbAvailableFromChannel field to given value.


### GetLineItems

`func (o *OrderInput) GetLineItems() []LineItemInput`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *OrderInput) GetLineItemsOk() (*[]LineItemInput, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *OrderInput) SetLineItems(v []LineItemInput)`

SetLineItems sets LineItems field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


