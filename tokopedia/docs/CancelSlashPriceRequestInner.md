# CancelSlashPriceRequestInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SlashPriceProductId** | **int64** | Slash price product ID. To retrieve this value, please refer to View Slash Price Products or View Campaign Products | 
**ProductId** | **int64** | Product ID that associated with slash_price_product_id, please refer to View Slash Price Products to view product_id that associated with slash_price_product_id | 

## Methods

### NewCancelSlashPriceRequestInner

`func NewCancelSlashPriceRequestInner(slashPriceProductId int64, productId int64, ) *CancelSlashPriceRequestInner`

NewCancelSlashPriceRequestInner instantiates a new CancelSlashPriceRequestInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelSlashPriceRequestInnerWithDefaults

`func NewCancelSlashPriceRequestInnerWithDefaults() *CancelSlashPriceRequestInner`

NewCancelSlashPriceRequestInnerWithDefaults instantiates a new CancelSlashPriceRequestInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSlashPriceProductId

`func (o *CancelSlashPriceRequestInner) GetSlashPriceProductId() int64`

GetSlashPriceProductId returns the SlashPriceProductId field if non-nil, zero value otherwise.

### GetSlashPriceProductIdOk

`func (o *CancelSlashPriceRequestInner) GetSlashPriceProductIdOk() (*int64, bool)`

GetSlashPriceProductIdOk returns a tuple with the SlashPriceProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlashPriceProductId

`func (o *CancelSlashPriceRequestInner) SetSlashPriceProductId(v int64)`

SetSlashPriceProductId sets SlashPriceProductId field to given value.


### GetProductId

`func (o *CancelSlashPriceRequestInner) GetProductId() int64`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *CancelSlashPriceRequestInner) GetProductIdOk() (*int64, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *CancelSlashPriceRequestInner) SetProductId(v int64)`

SetProductId sets ProductId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


