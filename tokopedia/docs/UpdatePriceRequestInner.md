# UpdatePriceRequestInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sku** | Pointer to **string** | SKU of products that will be updated. Maximum characters allowed is 50 | [optional] 
**ProductId** | Pointer to **int64** | Product ID to update | [optional] 
**NewPrice** | Pointer to **float64** | New price to be set | [optional] 

## Methods

### NewUpdatePriceRequestInner

`func NewUpdatePriceRequestInner() *UpdatePriceRequestInner`

NewUpdatePriceRequestInner instantiates a new UpdatePriceRequestInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePriceRequestInnerWithDefaults

`func NewUpdatePriceRequestInnerWithDefaults() *UpdatePriceRequestInner`

NewUpdatePriceRequestInnerWithDefaults instantiates a new UpdatePriceRequestInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSku

`func (o *UpdatePriceRequestInner) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *UpdatePriceRequestInner) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *UpdatePriceRequestInner) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *UpdatePriceRequestInner) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetProductId

`func (o *UpdatePriceRequestInner) GetProductId() int64`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *UpdatePriceRequestInner) GetProductIdOk() (*int64, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *UpdatePriceRequestInner) SetProductId(v int64)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *UpdatePriceRequestInner) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetNewPrice

`func (o *UpdatePriceRequestInner) GetNewPrice() float64`

GetNewPrice returns the NewPrice field if non-nil, zero value otherwise.

### GetNewPriceOk

`func (o *UpdatePriceRequestInner) GetNewPriceOk() (*float64, bool)`

GetNewPriceOk returns a tuple with the NewPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPrice

`func (o *UpdatePriceRequestInner) SetNewPrice(v float64)`

SetNewPrice sets NewPrice field to given value.

### HasNewPrice

`func (o *UpdatePriceRequestInner) HasNewPrice() bool`

HasNewPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


