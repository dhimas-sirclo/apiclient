# ProductWarehouse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductID** | Pointer to **int64** |  | [optional] 
**WarehouseID** | Pointer to **int64** |  | [optional] 
**Price** | Pointer to [**ProductPrice**](ProductPrice.md) |  | [optional] 
**Stock** | Pointer to [**ProductStock**](ProductStock.md) |  | [optional] 
**Bundle** | Pointer to [**ProductBundle**](ProductBundle.md) |  | [optional] 

## Methods

### NewProductWarehouse

`func NewProductWarehouse() *ProductWarehouse`

NewProductWarehouse instantiates a new ProductWarehouse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductWarehouseWithDefaults

`func NewProductWarehouseWithDefaults() *ProductWarehouse`

NewProductWarehouseWithDefaults instantiates a new ProductWarehouse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductID

`func (o *ProductWarehouse) GetProductID() int64`

GetProductID returns the ProductID field if non-nil, zero value otherwise.

### GetProductIDOk

`func (o *ProductWarehouse) GetProductIDOk() (*int64, bool)`

GetProductIDOk returns a tuple with the ProductID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductID

`func (o *ProductWarehouse) SetProductID(v int64)`

SetProductID sets ProductID field to given value.

### HasProductID

`func (o *ProductWarehouse) HasProductID() bool`

HasProductID returns a boolean if a field has been set.

### GetWarehouseID

`func (o *ProductWarehouse) GetWarehouseID() int64`

GetWarehouseID returns the WarehouseID field if non-nil, zero value otherwise.

### GetWarehouseIDOk

`func (o *ProductWarehouse) GetWarehouseIDOk() (*int64, bool)`

GetWarehouseIDOk returns a tuple with the WarehouseID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseID

`func (o *ProductWarehouse) SetWarehouseID(v int64)`

SetWarehouseID sets WarehouseID field to given value.

### HasWarehouseID

`func (o *ProductWarehouse) HasWarehouseID() bool`

HasWarehouseID returns a boolean if a field has been set.

### GetPrice

`func (o *ProductWarehouse) GetPrice() ProductPrice`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ProductWarehouse) GetPriceOk() (*ProductPrice, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ProductWarehouse) SetPrice(v ProductPrice)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ProductWarehouse) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetStock

`func (o *ProductWarehouse) GetStock() ProductStock`

GetStock returns the Stock field if non-nil, zero value otherwise.

### GetStockOk

`func (o *ProductWarehouse) GetStockOk() (*ProductStock, bool)`

GetStockOk returns a tuple with the Stock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStock

`func (o *ProductWarehouse) SetStock(v ProductStock)`

SetStock sets Stock field to given value.

### HasStock

`func (o *ProductWarehouse) HasStock() bool`

HasStock returns a boolean if a field has been set.

### GetBundle

`func (o *ProductWarehouse) GetBundle() ProductBundle`

GetBundle returns the Bundle field if non-nil, zero value otherwise.

### GetBundleOk

`func (o *ProductWarehouse) GetBundleOk() (*ProductBundle, bool)`

GetBundleOk returns a tuple with the Bundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundle

`func (o *ProductWarehouse) SetBundle(v ProductBundle)`

SetBundle sets Bundle field to given value.

### HasBundle

`func (o *ProductWarehouse) HasBundle() bool`

HasBundle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


