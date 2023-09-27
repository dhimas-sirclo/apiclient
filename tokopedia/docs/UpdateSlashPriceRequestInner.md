# UpdateSlashPriceRequestInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductId** | **int64** | Product unique identifier | 
**DiscountedPrice** | Pointer to **int64** | Price that will be set into slash price campaign (please use either discounted_price or discount_percentage) | [optional] 
**DiscountPercentage** | Pointer to **int64** | Total discount percentage that will be set into slash price campaign (please use either discounted_price or discount_percentage) | [optional] 
**StartTimeUnix** | **int64** | Start time campaign in UNIX format(use GMT +7) | 
**EndTimeUnix** | **int64** | End time campaign in UNIX format (use GMT +7) | 
**MaxOrder** | **int64** | Maximum order per one transaction | 
**IsCustomWarehouse** | Pointer to **bool** | body Flag to add slash price into specific warehouse | [optional] 
**WarehouseId** | Pointer to **int64** | Warehouse unique identifier | [optional] 
**SlashPriceWarehouses** | Pointer to [**map[string]UpdateSlashPriceRequestInnerSlashPriceWarehousesValue**](UpdateSlashPriceRequestInnerSlashPriceWarehousesValue.md) |  | [optional] 

## Methods

### NewUpdateSlashPriceRequestInner

`func NewUpdateSlashPriceRequestInner(productId int64, startTimeUnix int64, endTimeUnix int64, maxOrder int64, ) *UpdateSlashPriceRequestInner`

NewUpdateSlashPriceRequestInner instantiates a new UpdateSlashPriceRequestInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSlashPriceRequestInnerWithDefaults

`func NewUpdateSlashPriceRequestInnerWithDefaults() *UpdateSlashPriceRequestInner`

NewUpdateSlashPriceRequestInnerWithDefaults instantiates a new UpdateSlashPriceRequestInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductId

`func (o *UpdateSlashPriceRequestInner) GetProductId() int64`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *UpdateSlashPriceRequestInner) GetProductIdOk() (*int64, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *UpdateSlashPriceRequestInner) SetProductId(v int64)`

SetProductId sets ProductId field to given value.


### GetDiscountedPrice

`func (o *UpdateSlashPriceRequestInner) GetDiscountedPrice() int64`

GetDiscountedPrice returns the DiscountedPrice field if non-nil, zero value otherwise.

### GetDiscountedPriceOk

`func (o *UpdateSlashPriceRequestInner) GetDiscountedPriceOk() (*int64, bool)`

GetDiscountedPriceOk returns a tuple with the DiscountedPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountedPrice

`func (o *UpdateSlashPriceRequestInner) SetDiscountedPrice(v int64)`

SetDiscountedPrice sets DiscountedPrice field to given value.

### HasDiscountedPrice

`func (o *UpdateSlashPriceRequestInner) HasDiscountedPrice() bool`

HasDiscountedPrice returns a boolean if a field has been set.

### GetDiscountPercentage

`func (o *UpdateSlashPriceRequestInner) GetDiscountPercentage() int64`

GetDiscountPercentage returns the DiscountPercentage field if non-nil, zero value otherwise.

### GetDiscountPercentageOk

`func (o *UpdateSlashPriceRequestInner) GetDiscountPercentageOk() (*int64, bool)`

GetDiscountPercentageOk returns a tuple with the DiscountPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountPercentage

`func (o *UpdateSlashPriceRequestInner) SetDiscountPercentage(v int64)`

SetDiscountPercentage sets DiscountPercentage field to given value.

### HasDiscountPercentage

`func (o *UpdateSlashPriceRequestInner) HasDiscountPercentage() bool`

HasDiscountPercentage returns a boolean if a field has been set.

### GetStartTimeUnix

`func (o *UpdateSlashPriceRequestInner) GetStartTimeUnix() int64`

GetStartTimeUnix returns the StartTimeUnix field if non-nil, zero value otherwise.

### GetStartTimeUnixOk

`func (o *UpdateSlashPriceRequestInner) GetStartTimeUnixOk() (*int64, bool)`

GetStartTimeUnixOk returns a tuple with the StartTimeUnix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUnix

`func (o *UpdateSlashPriceRequestInner) SetStartTimeUnix(v int64)`

SetStartTimeUnix sets StartTimeUnix field to given value.


### GetEndTimeUnix

`func (o *UpdateSlashPriceRequestInner) GetEndTimeUnix() int64`

GetEndTimeUnix returns the EndTimeUnix field if non-nil, zero value otherwise.

### GetEndTimeUnixOk

`func (o *UpdateSlashPriceRequestInner) GetEndTimeUnixOk() (*int64, bool)`

GetEndTimeUnixOk returns a tuple with the EndTimeUnix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUnix

`func (o *UpdateSlashPriceRequestInner) SetEndTimeUnix(v int64)`

SetEndTimeUnix sets EndTimeUnix field to given value.


### GetMaxOrder

`func (o *UpdateSlashPriceRequestInner) GetMaxOrder() int64`

GetMaxOrder returns the MaxOrder field if non-nil, zero value otherwise.

### GetMaxOrderOk

`func (o *UpdateSlashPriceRequestInner) GetMaxOrderOk() (*int64, bool)`

GetMaxOrderOk returns a tuple with the MaxOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxOrder

`func (o *UpdateSlashPriceRequestInner) SetMaxOrder(v int64)`

SetMaxOrder sets MaxOrder field to given value.


### GetIsCustomWarehouse

`func (o *UpdateSlashPriceRequestInner) GetIsCustomWarehouse() bool`

GetIsCustomWarehouse returns the IsCustomWarehouse field if non-nil, zero value otherwise.

### GetIsCustomWarehouseOk

`func (o *UpdateSlashPriceRequestInner) GetIsCustomWarehouseOk() (*bool, bool)`

GetIsCustomWarehouseOk returns a tuple with the IsCustomWarehouse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCustomWarehouse

`func (o *UpdateSlashPriceRequestInner) SetIsCustomWarehouse(v bool)`

SetIsCustomWarehouse sets IsCustomWarehouse field to given value.

### HasIsCustomWarehouse

`func (o *UpdateSlashPriceRequestInner) HasIsCustomWarehouse() bool`

HasIsCustomWarehouse returns a boolean if a field has been set.

### GetWarehouseId

`func (o *UpdateSlashPriceRequestInner) GetWarehouseId() int64`

GetWarehouseId returns the WarehouseId field if non-nil, zero value otherwise.

### GetWarehouseIdOk

`func (o *UpdateSlashPriceRequestInner) GetWarehouseIdOk() (*int64, bool)`

GetWarehouseIdOk returns a tuple with the WarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseId

`func (o *UpdateSlashPriceRequestInner) SetWarehouseId(v int64)`

SetWarehouseId sets WarehouseId field to given value.

### HasWarehouseId

`func (o *UpdateSlashPriceRequestInner) HasWarehouseId() bool`

HasWarehouseId returns a boolean if a field has been set.

### GetSlashPriceWarehouses

`func (o *UpdateSlashPriceRequestInner) GetSlashPriceWarehouses() map[string]UpdateSlashPriceRequestInnerSlashPriceWarehousesValue`

GetSlashPriceWarehouses returns the SlashPriceWarehouses field if non-nil, zero value otherwise.

### GetSlashPriceWarehousesOk

`func (o *UpdateSlashPriceRequestInner) GetSlashPriceWarehousesOk() (*map[string]UpdateSlashPriceRequestInnerSlashPriceWarehousesValue, bool)`

GetSlashPriceWarehousesOk returns a tuple with the SlashPriceWarehouses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlashPriceWarehouses

`func (o *UpdateSlashPriceRequestInner) SetSlashPriceWarehouses(v map[string]UpdateSlashPriceRequestInnerSlashPriceWarehousesValue)`

SetSlashPriceWarehouses sets SlashPriceWarehouses field to given value.

### HasSlashPriceWarehouses

`func (o *UpdateSlashPriceRequestInner) HasSlashPriceWarehouses() bool`

HasSlashPriceWarehouses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


