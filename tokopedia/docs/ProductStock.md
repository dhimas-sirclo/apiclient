# ProductStock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UseStock** | Pointer to **bool** |  | [optional] 
**Value** | Pointer to **int32** | Product Total Stock | [optional] 

## Methods

### NewProductStock

`func NewProductStock() *ProductStock`

NewProductStock instantiates a new ProductStock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductStockWithDefaults

`func NewProductStockWithDefaults() *ProductStock`

NewProductStockWithDefaults instantiates a new ProductStock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUseStock

`func (o *ProductStock) GetUseStock() bool`

GetUseStock returns the UseStock field if non-nil, zero value otherwise.

### GetUseStockOk

`func (o *ProductStock) GetUseStockOk() (*bool, bool)`

GetUseStockOk returns a tuple with the UseStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseStock

`func (o *ProductStock) SetUseStock(v bool)`

SetUseStock sets UseStock field to given value.

### HasUseStock

`func (o *ProductStock) HasUseStock() bool`

HasUseStock returns a boolean if a field has been set.

### GetValue

`func (o *ProductStock) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ProductStock) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ProductStock) SetValue(v int32)`

SetValue sets Value field to given value.

### HasValue

`func (o *ProductStock) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


