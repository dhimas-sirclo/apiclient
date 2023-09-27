# GetProductVariantResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Header** | Pointer to [**ResponseHeader**](ResponseHeader.md) |  | [optional] 
**Data** | Pointer to [**Variant**](Variant.md) |  | [optional] 

## Methods

### NewGetProductVariantResponse

`func NewGetProductVariantResponse() *GetProductVariantResponse`

NewGetProductVariantResponse instantiates a new GetProductVariantResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetProductVariantResponseWithDefaults

`func NewGetProductVariantResponseWithDefaults() *GetProductVariantResponse`

NewGetProductVariantResponseWithDefaults instantiates a new GetProductVariantResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeader

`func (o *GetProductVariantResponse) GetHeader() ResponseHeader`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *GetProductVariantResponse) GetHeaderOk() (*ResponseHeader, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *GetProductVariantResponse) SetHeader(v ResponseHeader)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *GetProductVariantResponse) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetData

`func (o *GetProductVariantResponse) GetData() Variant`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetProductVariantResponse) GetDataOk() (*Variant, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetProductVariantResponse) SetData(v Variant)`

SetData sets Data field to given value.

### HasData

`func (o *GetProductVariantResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


