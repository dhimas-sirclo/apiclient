# ProductPrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **int32** | Product Price Value | [optional] 
**Currency** | Pointer to **int32** | Product Price Currency | [optional] 
**LastUpdateUnix** | Pointer to **int32** | Product Price Last Updated timer | [optional] 
**Idr** | Pointer to **int32** | Price Value | [optional] 

## Methods

### NewProductPrice

`func NewProductPrice() *ProductPrice`

NewProductPrice instantiates a new ProductPrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductPriceWithDefaults

`func NewProductPriceWithDefaults() *ProductPrice`

NewProductPriceWithDefaults instantiates a new ProductPrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *ProductPrice) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ProductPrice) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ProductPrice) SetValue(v int32)`

SetValue sets Value field to given value.

### HasValue

`func (o *ProductPrice) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetCurrency

`func (o *ProductPrice) GetCurrency() int32`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ProductPrice) GetCurrencyOk() (*int32, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ProductPrice) SetCurrency(v int32)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ProductPrice) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetLastUpdateUnix

`func (o *ProductPrice) GetLastUpdateUnix() int32`

GetLastUpdateUnix returns the LastUpdateUnix field if non-nil, zero value otherwise.

### GetLastUpdateUnixOk

`func (o *ProductPrice) GetLastUpdateUnixOk() (*int32, bool)`

GetLastUpdateUnixOk returns a tuple with the LastUpdateUnix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateUnix

`func (o *ProductPrice) SetLastUpdateUnix(v int32)`

SetLastUpdateUnix sets LastUpdateUnix field to given value.

### HasLastUpdateUnix

`func (o *ProductPrice) HasLastUpdateUnix() bool`

HasLastUpdateUnix returns a boolean if a field has been set.

### GetIdr

`func (o *ProductPrice) GetIdr() int32`

GetIdr returns the Idr field if non-nil, zero value otherwise.

### GetIdrOk

`func (o *ProductPrice) GetIdrOk() (*int32, bool)`

GetIdrOk returns a tuple with the Idr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdr

`func (o *ProductPrice) SetIdr(v int32)`

SetIdr sets Idr field to given value.

### HasIdr

`func (o *ProductPrice) HasIdr() bool`

HasIdr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


