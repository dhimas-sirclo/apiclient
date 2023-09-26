# UpdateStockInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sku** | Pointer to **string** |  | [optional] 
**ProductId** | Pointer to **int64** |  | [optional] 
**NewStock** | **int64** |  | 

## Methods

### NewUpdateStockInput

`func NewUpdateStockInput(newStock int64, ) *UpdateStockInput`

NewUpdateStockInput instantiates a new UpdateStockInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateStockInputWithDefaults

`func NewUpdateStockInputWithDefaults() *UpdateStockInput`

NewUpdateStockInputWithDefaults instantiates a new UpdateStockInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSku

`func (o *UpdateStockInput) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *UpdateStockInput) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *UpdateStockInput) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *UpdateStockInput) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetProductId

`func (o *UpdateStockInput) GetProductId() int64`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *UpdateStockInput) GetProductIdOk() (*int64, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *UpdateStockInput) SetProductId(v int64)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *UpdateStockInput) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetNewStock

`func (o *UpdateStockInput) GetNewStock() int64`

GetNewStock returns the NewStock field if non-nil, zero value otherwise.

### GetNewStockOk

`func (o *UpdateStockInput) GetNewStockOk() (*int64, bool)`

GetNewStockOk returns a tuple with the NewStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewStock

`func (o *UpdateStockInput) SetNewStock(v int64)`

SetNewStock sets NewStock field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


