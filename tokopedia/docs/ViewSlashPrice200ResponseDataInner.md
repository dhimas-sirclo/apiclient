# ViewSlashPrice200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SlashPriceProductId** | Pointer to **int64** | Slash Price Product Unique Identifier | [optional] 
**ProductId** | Pointer to **int64** | Product Unique Identifier | [optional] 
**Name** | Pointer to **string** | Product Name | [optional] 
**Price** | Pointer to [**ViewSlashPrice200ResponseDataInnerPrice**](ViewSlashPrice200ResponseDataInnerPrice.md) |  | [optional] 
**Stock** | Pointer to **int64** | Product Stock | [optional] 
**Url** | Pointer to **string** | Product URL | [optional] 
**Sku** | Pointer to **string** | Product SKU | [optional] 
**DiscountedPrice** | Pointer to **int64** | Product Discounted Price | [optional] 
**DiscountedPercentage** | Pointer to **int64** | Product Discounted Percentage | [optional] 
**MaxOrder** | Pointer to **int64** | Product Maximum Order | [optional] 
**StartDate** | Pointer to **string** | Slash Price Start Date Timestamp (format: 2021-09-27T13:25:00+07:00)  | [optional] 
**EndDate** | Pointer to **string** | Slash Price Start Date Timestamp (format: 2021-09-28T01:24:00+07:00)  | [optional] 
**Warehouses** | Pointer to [**[]ViewSlashPrice200ResponseDataInnerWarehousesInner**](ViewSlashPrice200ResponseDataInnerWarehousesInner.md) |  | [optional] 
**SlashPriceStatusId** | Pointer to **int64** | Slash Price Status | [optional] 

## Methods

### NewViewSlashPrice200ResponseDataInner

`func NewViewSlashPrice200ResponseDataInner() *ViewSlashPrice200ResponseDataInner`

NewViewSlashPrice200ResponseDataInner instantiates a new ViewSlashPrice200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewSlashPrice200ResponseDataInnerWithDefaults

`func NewViewSlashPrice200ResponseDataInnerWithDefaults() *ViewSlashPrice200ResponseDataInner`

NewViewSlashPrice200ResponseDataInnerWithDefaults instantiates a new ViewSlashPrice200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSlashPriceProductId

`func (o *ViewSlashPrice200ResponseDataInner) GetSlashPriceProductId() int64`

GetSlashPriceProductId returns the SlashPriceProductId field if non-nil, zero value otherwise.

### GetSlashPriceProductIdOk

`func (o *ViewSlashPrice200ResponseDataInner) GetSlashPriceProductIdOk() (*int64, bool)`

GetSlashPriceProductIdOk returns a tuple with the SlashPriceProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlashPriceProductId

`func (o *ViewSlashPrice200ResponseDataInner) SetSlashPriceProductId(v int64)`

SetSlashPriceProductId sets SlashPriceProductId field to given value.

### HasSlashPriceProductId

`func (o *ViewSlashPrice200ResponseDataInner) HasSlashPriceProductId() bool`

HasSlashPriceProductId returns a boolean if a field has been set.

### GetProductId

`func (o *ViewSlashPrice200ResponseDataInner) GetProductId() int64`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *ViewSlashPrice200ResponseDataInner) GetProductIdOk() (*int64, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *ViewSlashPrice200ResponseDataInner) SetProductId(v int64)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *ViewSlashPrice200ResponseDataInner) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetName

`func (o *ViewSlashPrice200ResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ViewSlashPrice200ResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ViewSlashPrice200ResponseDataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ViewSlashPrice200ResponseDataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrice

`func (o *ViewSlashPrice200ResponseDataInner) GetPrice() ViewSlashPrice200ResponseDataInnerPrice`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ViewSlashPrice200ResponseDataInner) GetPriceOk() (*ViewSlashPrice200ResponseDataInnerPrice, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ViewSlashPrice200ResponseDataInner) SetPrice(v ViewSlashPrice200ResponseDataInnerPrice)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ViewSlashPrice200ResponseDataInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetStock

`func (o *ViewSlashPrice200ResponseDataInner) GetStock() int64`

GetStock returns the Stock field if non-nil, zero value otherwise.

### GetStockOk

`func (o *ViewSlashPrice200ResponseDataInner) GetStockOk() (*int64, bool)`

GetStockOk returns a tuple with the Stock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStock

`func (o *ViewSlashPrice200ResponseDataInner) SetStock(v int64)`

SetStock sets Stock field to given value.

### HasStock

`func (o *ViewSlashPrice200ResponseDataInner) HasStock() bool`

HasStock returns a boolean if a field has been set.

### GetUrl

`func (o *ViewSlashPrice200ResponseDataInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ViewSlashPrice200ResponseDataInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ViewSlashPrice200ResponseDataInner) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ViewSlashPrice200ResponseDataInner) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetSku

`func (o *ViewSlashPrice200ResponseDataInner) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *ViewSlashPrice200ResponseDataInner) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *ViewSlashPrice200ResponseDataInner) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *ViewSlashPrice200ResponseDataInner) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetDiscountedPrice

`func (o *ViewSlashPrice200ResponseDataInner) GetDiscountedPrice() int64`

GetDiscountedPrice returns the DiscountedPrice field if non-nil, zero value otherwise.

### GetDiscountedPriceOk

`func (o *ViewSlashPrice200ResponseDataInner) GetDiscountedPriceOk() (*int64, bool)`

GetDiscountedPriceOk returns a tuple with the DiscountedPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountedPrice

`func (o *ViewSlashPrice200ResponseDataInner) SetDiscountedPrice(v int64)`

SetDiscountedPrice sets DiscountedPrice field to given value.

### HasDiscountedPrice

`func (o *ViewSlashPrice200ResponseDataInner) HasDiscountedPrice() bool`

HasDiscountedPrice returns a boolean if a field has been set.

### GetDiscountedPercentage

`func (o *ViewSlashPrice200ResponseDataInner) GetDiscountedPercentage() int64`

GetDiscountedPercentage returns the DiscountedPercentage field if non-nil, zero value otherwise.

### GetDiscountedPercentageOk

`func (o *ViewSlashPrice200ResponseDataInner) GetDiscountedPercentageOk() (*int64, bool)`

GetDiscountedPercentageOk returns a tuple with the DiscountedPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountedPercentage

`func (o *ViewSlashPrice200ResponseDataInner) SetDiscountedPercentage(v int64)`

SetDiscountedPercentage sets DiscountedPercentage field to given value.

### HasDiscountedPercentage

`func (o *ViewSlashPrice200ResponseDataInner) HasDiscountedPercentage() bool`

HasDiscountedPercentage returns a boolean if a field has been set.

### GetMaxOrder

`func (o *ViewSlashPrice200ResponseDataInner) GetMaxOrder() int64`

GetMaxOrder returns the MaxOrder field if non-nil, zero value otherwise.

### GetMaxOrderOk

`func (o *ViewSlashPrice200ResponseDataInner) GetMaxOrderOk() (*int64, bool)`

GetMaxOrderOk returns a tuple with the MaxOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxOrder

`func (o *ViewSlashPrice200ResponseDataInner) SetMaxOrder(v int64)`

SetMaxOrder sets MaxOrder field to given value.

### HasMaxOrder

`func (o *ViewSlashPrice200ResponseDataInner) HasMaxOrder() bool`

HasMaxOrder returns a boolean if a field has been set.

### GetStartDate

`func (o *ViewSlashPrice200ResponseDataInner) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ViewSlashPrice200ResponseDataInner) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ViewSlashPrice200ResponseDataInner) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ViewSlashPrice200ResponseDataInner) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *ViewSlashPrice200ResponseDataInner) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ViewSlashPrice200ResponseDataInner) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ViewSlashPrice200ResponseDataInner) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ViewSlashPrice200ResponseDataInner) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetWarehouses

`func (o *ViewSlashPrice200ResponseDataInner) GetWarehouses() []ViewSlashPrice200ResponseDataInnerWarehousesInner`

GetWarehouses returns the Warehouses field if non-nil, zero value otherwise.

### GetWarehousesOk

`func (o *ViewSlashPrice200ResponseDataInner) GetWarehousesOk() (*[]ViewSlashPrice200ResponseDataInnerWarehousesInner, bool)`

GetWarehousesOk returns a tuple with the Warehouses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouses

`func (o *ViewSlashPrice200ResponseDataInner) SetWarehouses(v []ViewSlashPrice200ResponseDataInnerWarehousesInner)`

SetWarehouses sets Warehouses field to given value.

### HasWarehouses

`func (o *ViewSlashPrice200ResponseDataInner) HasWarehouses() bool`

HasWarehouses returns a boolean if a field has been set.

### GetSlashPriceStatusId

`func (o *ViewSlashPrice200ResponseDataInner) GetSlashPriceStatusId() int64`

GetSlashPriceStatusId returns the SlashPriceStatusId field if non-nil, zero value otherwise.

### GetSlashPriceStatusIdOk

`func (o *ViewSlashPrice200ResponseDataInner) GetSlashPriceStatusIdOk() (*int64, bool)`

GetSlashPriceStatusIdOk returns a tuple with the SlashPriceStatusId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlashPriceStatusId

`func (o *ViewSlashPrice200ResponseDataInner) SetSlashPriceStatusId(v int64)`

SetSlashPriceStatusId sets SlashPriceStatusId field to given value.

### HasSlashPriceStatusId

`func (o *ViewSlashPrice200ResponseDataInner) HasSlashPriceStatusId() bool`

HasSlashPriceStatusId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


