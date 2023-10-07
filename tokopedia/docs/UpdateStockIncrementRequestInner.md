# UpdateStockIncrementRequestInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sku** | Pointer to **string** | SKU of products that will be updated. Maximum characters allowed is 50 | [optional] 
**ProductId** | Pointer to **int64** | Product ID to update | [optional] 
**StockValue** | **int64** | New stock to be set | 

## Methods

### NewUpdateStockIncrementRequestInner

`func NewUpdateStockIncrementRequestInner(stockValue int64, ) *UpdateStockIncrementRequestInner`

NewUpdateStockIncrementRequestInner instantiates a new UpdateStockIncrementRequestInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateStockIncrementRequestInnerWithDefaults

`func NewUpdateStockIncrementRequestInnerWithDefaults() *UpdateStockIncrementRequestInner`

NewUpdateStockIncrementRequestInnerWithDefaults instantiates a new UpdateStockIncrementRequestInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSku

`func (o *UpdateStockIncrementRequestInner) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *UpdateStockIncrementRequestInner) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *UpdateStockIncrementRequestInner) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *UpdateStockIncrementRequestInner) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetProductId

`func (o *UpdateStockIncrementRequestInner) GetProductId() int64`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *UpdateStockIncrementRequestInner) GetProductIdOk() (*int64, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *UpdateStockIncrementRequestInner) SetProductId(v int64)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *UpdateStockIncrementRequestInner) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetStockValue

`func (o *UpdateStockIncrementRequestInner) GetStockValue() int64`

GetStockValue returns the StockValue field if non-nil, zero value otherwise.

### GetStockValueOk

`func (o *UpdateStockIncrementRequestInner) GetStockValueOk() (*int64, bool)`

GetStockValueOk returns a tuple with the StockValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockValue

`func (o *UpdateStockIncrementRequestInner) SetStockValue(v int64)`

SetStockValue sets StockValue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


