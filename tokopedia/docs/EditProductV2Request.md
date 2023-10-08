# EditProductV2Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Products** | [**[]EditProductV2RequestProductsInner**](EditProductV2RequestProductsInner.md) |  | 

## Methods

### NewEditProductV2Request

`func NewEditProductV2Request(products []EditProductV2RequestProductsInner, ) *EditProductV2Request`

NewEditProductV2Request instantiates a new EditProductV2Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditProductV2RequestWithDefaults

`func NewEditProductV2RequestWithDefaults() *EditProductV2Request`

NewEditProductV2RequestWithDefaults instantiates a new EditProductV2Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProducts

`func (o *EditProductV2Request) GetProducts() []EditProductV2RequestProductsInner`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *EditProductV2Request) GetProductsOk() (*[]EditProductV2RequestProductsInner, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *EditProductV2Request) SetProducts(v []EditProductV2RequestProductsInner)`

SetProducts sets Products field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


