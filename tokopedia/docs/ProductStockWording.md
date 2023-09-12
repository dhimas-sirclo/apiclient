# ProductStockWording

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UseStock** | Pointer to **bool** |  | [optional] 
**Value** | Pointer to **int32** | Product Total Stock | [optional] 
**StockWording** | Pointer to **string** | Product Stock Wording (Description) | [optional] 

## Methods

### NewProductStockWording

`func NewProductStockWording() *ProductStockWording`

NewProductStockWording instantiates a new ProductStockWording object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductStockWordingWithDefaults

`func NewProductStockWordingWithDefaults() *ProductStockWording`

NewProductStockWordingWithDefaults instantiates a new ProductStockWording object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUseStock

`func (o *ProductStockWording) GetUseStock() bool`

GetUseStock returns the UseStock field if non-nil, zero value otherwise.

### GetUseStockOk

`func (o *ProductStockWording) GetUseStockOk() (*bool, bool)`

GetUseStockOk returns a tuple with the UseStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseStock

`func (o *ProductStockWording) SetUseStock(v bool)`

SetUseStock sets UseStock field to given value.

### HasUseStock

`func (o *ProductStockWording) HasUseStock() bool`

HasUseStock returns a boolean if a field has been set.

### GetValue

`func (o *ProductStockWording) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ProductStockWording) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ProductStockWording) SetValue(v int32)`

SetValue sets Value field to given value.

### HasValue

`func (o *ProductStockWording) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetStockWording

`func (o *ProductStockWording) GetStockWording() string`

GetStockWording returns the StockWording field if non-nil, zero value otherwise.

### GetStockWordingOk

`func (o *ProductStockWording) GetStockWordingOk() (*string, bool)`

GetStockWordingOk returns a tuple with the StockWording field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockWording

`func (o *ProductStockWording) SetStockWording(v string)`

SetStockWording sets StockWording field to given value.

### HasStockWording

`func (o *ProductStockWording) HasStockWording() bool`

HasStockWording returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


