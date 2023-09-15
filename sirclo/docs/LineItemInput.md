# LineItemInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderItemId** | **string** |  | 
**Sku** | **string** |  | 
**ProductName** | **string** |  | 
**OrderQuantity** | **int32** |  | 
**RawPrice** | **float64** |  | 
**DiscountAmount** | Pointer to **NullableFloat64** |  | [optional] 
**LineTax** | **float64** |  | 
**ShippingAmount** | Pointer to **NullableFloat64** |  | [optional] 
**Promo** | **float64** |  | 
**Voucher** | **float64** |  | 
**LineTotal** | **float64** |  | 
**LineComments** | Pointer to **string** |  | [optional] 

## Methods

### NewLineItemInput

`func NewLineItemInput(orderItemId string, sku string, productName string, orderQuantity int32, rawPrice float64, lineTax float64, promo float64, voucher float64, lineTotal float64, ) *LineItemInput`

NewLineItemInput instantiates a new LineItemInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLineItemInputWithDefaults

`func NewLineItemInputWithDefaults() *LineItemInput`

NewLineItemInputWithDefaults instantiates a new LineItemInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderItemId

`func (o *LineItemInput) GetOrderItemId() string`

GetOrderItemId returns the OrderItemId field if non-nil, zero value otherwise.

### GetOrderItemIdOk

`func (o *LineItemInput) GetOrderItemIdOk() (*string, bool)`

GetOrderItemIdOk returns a tuple with the OrderItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderItemId

`func (o *LineItemInput) SetOrderItemId(v string)`

SetOrderItemId sets OrderItemId field to given value.


### GetSku

`func (o *LineItemInput) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *LineItemInput) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *LineItemInput) SetSku(v string)`

SetSku sets Sku field to given value.


### GetProductName

`func (o *LineItemInput) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *LineItemInput) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *LineItemInput) SetProductName(v string)`

SetProductName sets ProductName field to given value.


### GetOrderQuantity

`func (o *LineItemInput) GetOrderQuantity() int32`

GetOrderQuantity returns the OrderQuantity field if non-nil, zero value otherwise.

### GetOrderQuantityOk

`func (o *LineItemInput) GetOrderQuantityOk() (*int32, bool)`

GetOrderQuantityOk returns a tuple with the OrderQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderQuantity

`func (o *LineItemInput) SetOrderQuantity(v int32)`

SetOrderQuantity sets OrderQuantity field to given value.


### GetRawPrice

`func (o *LineItemInput) GetRawPrice() float64`

GetRawPrice returns the RawPrice field if non-nil, zero value otherwise.

### GetRawPriceOk

`func (o *LineItemInput) GetRawPriceOk() (*float64, bool)`

GetRawPriceOk returns a tuple with the RawPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawPrice

`func (o *LineItemInput) SetRawPrice(v float64)`

SetRawPrice sets RawPrice field to given value.


### GetDiscountAmount

`func (o *LineItemInput) GetDiscountAmount() float64`

GetDiscountAmount returns the DiscountAmount field if non-nil, zero value otherwise.

### GetDiscountAmountOk

`func (o *LineItemInput) GetDiscountAmountOk() (*float64, bool)`

GetDiscountAmountOk returns a tuple with the DiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmount

`func (o *LineItemInput) SetDiscountAmount(v float64)`

SetDiscountAmount sets DiscountAmount field to given value.

### HasDiscountAmount

`func (o *LineItemInput) HasDiscountAmount() bool`

HasDiscountAmount returns a boolean if a field has been set.

### SetDiscountAmountNil

`func (o *LineItemInput) SetDiscountAmountNil(b bool)`

 SetDiscountAmountNil sets the value for DiscountAmount to be an explicit nil

### UnsetDiscountAmount
`func (o *LineItemInput) UnsetDiscountAmount()`

UnsetDiscountAmount ensures that no value is present for DiscountAmount, not even an explicit nil
### GetLineTax

`func (o *LineItemInput) GetLineTax() float64`

GetLineTax returns the LineTax field if non-nil, zero value otherwise.

### GetLineTaxOk

`func (o *LineItemInput) GetLineTaxOk() (*float64, bool)`

GetLineTaxOk returns a tuple with the LineTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineTax

`func (o *LineItemInput) SetLineTax(v float64)`

SetLineTax sets LineTax field to given value.


### GetShippingAmount

`func (o *LineItemInput) GetShippingAmount() float64`

GetShippingAmount returns the ShippingAmount field if non-nil, zero value otherwise.

### GetShippingAmountOk

`func (o *LineItemInput) GetShippingAmountOk() (*float64, bool)`

GetShippingAmountOk returns a tuple with the ShippingAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAmount

`func (o *LineItemInput) SetShippingAmount(v float64)`

SetShippingAmount sets ShippingAmount field to given value.

### HasShippingAmount

`func (o *LineItemInput) HasShippingAmount() bool`

HasShippingAmount returns a boolean if a field has been set.

### SetShippingAmountNil

`func (o *LineItemInput) SetShippingAmountNil(b bool)`

 SetShippingAmountNil sets the value for ShippingAmount to be an explicit nil

### UnsetShippingAmount
`func (o *LineItemInput) UnsetShippingAmount()`

UnsetShippingAmount ensures that no value is present for ShippingAmount, not even an explicit nil
### GetPromo

`func (o *LineItemInput) GetPromo() float64`

GetPromo returns the Promo field if non-nil, zero value otherwise.

### GetPromoOk

`func (o *LineItemInput) GetPromoOk() (*float64, bool)`

GetPromoOk returns a tuple with the Promo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromo

`func (o *LineItemInput) SetPromo(v float64)`

SetPromo sets Promo field to given value.


### GetVoucher

`func (o *LineItemInput) GetVoucher() float64`

GetVoucher returns the Voucher field if non-nil, zero value otherwise.

### GetVoucherOk

`func (o *LineItemInput) GetVoucherOk() (*float64, bool)`

GetVoucherOk returns a tuple with the Voucher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoucher

`func (o *LineItemInput) SetVoucher(v float64)`

SetVoucher sets Voucher field to given value.


### GetLineTotal

`func (o *LineItemInput) GetLineTotal() float64`

GetLineTotal returns the LineTotal field if non-nil, zero value otherwise.

### GetLineTotalOk

`func (o *LineItemInput) GetLineTotalOk() (*float64, bool)`

GetLineTotalOk returns a tuple with the LineTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineTotal

`func (o *LineItemInput) SetLineTotal(v float64)`

SetLineTotal sets LineTotal field to given value.


### GetLineComments

`func (o *LineItemInput) GetLineComments() string`

GetLineComments returns the LineComments field if non-nil, zero value otherwise.

### GetLineCommentsOk

`func (o *LineItemInput) GetLineCommentsOk() (*string, bool)`

GetLineCommentsOk returns a tuple with the LineComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineComments

`func (o *LineItemInput) SetLineComments(v string)`

SetLineComments sets LineComments field to given value.

### HasLineComments

`func (o *LineItemInput) HasLineComments() bool`

HasLineComments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


