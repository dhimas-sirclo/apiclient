# ProductWeight

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **int32** | Product Weight Value | [optional] 
**Unit** | Pointer to **int32** | Weight (e.g., 1 for Gram, 2 for Kilogram) | [optional] 

## Methods

### NewProductWeight

`func NewProductWeight() *ProductWeight`

NewProductWeight instantiates a new ProductWeight object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductWeightWithDefaults

`func NewProductWeightWithDefaults() *ProductWeight`

NewProductWeightWithDefaults instantiates a new ProductWeight object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *ProductWeight) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ProductWeight) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ProductWeight) SetValue(v int32)`

SetValue sets Value field to given value.

### HasValue

`func (o *ProductWeight) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetUnit

`func (o *ProductWeight) GetUnit() int32`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *ProductWeight) GetUnitOk() (*int32, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *ProductWeight) SetUnit(v int32)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *ProductWeight) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


