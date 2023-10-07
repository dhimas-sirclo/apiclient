# EditProductV3Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Products** | [**[]EditProductV3RequestProductsInner**](EditProductV3RequestProductsInner.md) |  | 

## Methods

### NewEditProductV3Request

`func NewEditProductV3Request(products []EditProductV3RequestProductsInner, ) *EditProductV3Request`

NewEditProductV3Request instantiates a new EditProductV3Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditProductV3RequestWithDefaults

`func NewEditProductV3RequestWithDefaults() *EditProductV3Request`

NewEditProductV3RequestWithDefaults instantiates a new EditProductV3Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProducts

`func (o *EditProductV3Request) GetProducts() []EditProductV3RequestProductsInner`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *EditProductV3Request) GetProductsOk() (*[]EditProductV3RequestProductsInner, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *EditProductV3Request) SetProducts(v []EditProductV3RequestProductsInner)`

SetProducts sets Products field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


