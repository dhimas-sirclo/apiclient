# CreateBundleRequestBundleBundleItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductId** | **int64** | Product unique identifier to include in the bundle | 
**Status** | **int64** | Status of the product where 1 is SHOWN and 2 is UNSHOWN | 
**BundlePrice** | **float64** | The product’s price in this bundle | 
**MinOrder** | Pointer to **int64** | Minimum order for this product. Used for PAKET DISKON type bundle. | [optional] 
**Children** | Pointer to [**CreateBundleRequestBundleBundleItemsInnerChildren**](CreateBundleRequestBundleBundleItemsInnerChildren.md) |  | [optional] 

## Methods

### NewCreateBundleRequestBundleBundleItemsInner

`func NewCreateBundleRequestBundleBundleItemsInner(productId int64, status int64, bundlePrice float64, ) *CreateBundleRequestBundleBundleItemsInner`

NewCreateBundleRequestBundleBundleItemsInner instantiates a new CreateBundleRequestBundleBundleItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBundleRequestBundleBundleItemsInnerWithDefaults

`func NewCreateBundleRequestBundleBundleItemsInnerWithDefaults() *CreateBundleRequestBundleBundleItemsInner`

NewCreateBundleRequestBundleBundleItemsInnerWithDefaults instantiates a new CreateBundleRequestBundleBundleItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductId

`func (o *CreateBundleRequestBundleBundleItemsInner) GetProductId() int64`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *CreateBundleRequestBundleBundleItemsInner) GetProductIdOk() (*int64, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *CreateBundleRequestBundleBundleItemsInner) SetProductId(v int64)`

SetProductId sets ProductId field to given value.


### GetStatus

`func (o *CreateBundleRequestBundleBundleItemsInner) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateBundleRequestBundleBundleItemsInner) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateBundleRequestBundleBundleItemsInner) SetStatus(v int64)`

SetStatus sets Status field to given value.


### GetBundlePrice

`func (o *CreateBundleRequestBundleBundleItemsInner) GetBundlePrice() float64`

GetBundlePrice returns the BundlePrice field if non-nil, zero value otherwise.

### GetBundlePriceOk

`func (o *CreateBundleRequestBundleBundleItemsInner) GetBundlePriceOk() (*float64, bool)`

GetBundlePriceOk returns a tuple with the BundlePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundlePrice

`func (o *CreateBundleRequestBundleBundleItemsInner) SetBundlePrice(v float64)`

SetBundlePrice sets BundlePrice field to given value.


### GetMinOrder

`func (o *CreateBundleRequestBundleBundleItemsInner) GetMinOrder() int64`

GetMinOrder returns the MinOrder field if non-nil, zero value otherwise.

### GetMinOrderOk

`func (o *CreateBundleRequestBundleBundleItemsInner) GetMinOrderOk() (*int64, bool)`

GetMinOrderOk returns a tuple with the MinOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinOrder

`func (o *CreateBundleRequestBundleBundleItemsInner) SetMinOrder(v int64)`

SetMinOrder sets MinOrder field to given value.

### HasMinOrder

`func (o *CreateBundleRequestBundleBundleItemsInner) HasMinOrder() bool`

HasMinOrder returns a boolean if a field has been set.

### GetChildren

`func (o *CreateBundleRequestBundleBundleItemsInner) GetChildren() CreateBundleRequestBundleBundleItemsInnerChildren`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *CreateBundleRequestBundleBundleItemsInner) GetChildrenOk() (*CreateBundleRequestBundleBundleItemsInnerChildren, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *CreateBundleRequestBundleBundleItemsInner) SetChildren(v CreateBundleRequestBundleBundleItemsInnerChildren)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *CreateBundleRequestBundleBundleItemsInner) HasChildren() bool`

HasChildren returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


