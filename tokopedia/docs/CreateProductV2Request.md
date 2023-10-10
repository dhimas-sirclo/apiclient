# CreateProductV2Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Products** | Pointer to [**[]CreateProductV2RequestProductsInner**](CreateProductV2RequestProductsInner.md) | Product object array | [optional] 

## Methods

### NewCreateProductV2Request

`func NewCreateProductV2Request() *CreateProductV2Request`

NewCreateProductV2Request instantiates a new CreateProductV2Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProductV2RequestWithDefaults

`func NewCreateProductV2RequestWithDefaults() *CreateProductV2Request`

NewCreateProductV2RequestWithDefaults instantiates a new CreateProductV2Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProducts

`func (o *CreateProductV2Request) GetProducts() []CreateProductV2RequestProductsInner`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *CreateProductV2Request) GetProductsOk() (*[]CreateProductV2RequestProductsInner, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *CreateProductV2Request) SetProducts(v []CreateProductV2RequestProductsInner)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *CreateProductV2Request) HasProducts() bool`

HasProducts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


