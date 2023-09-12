# Variant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParentId** | Pointer to **int32** | Variant Parent Unique Identifier | [optional] 
**DefaultChild** | Pointer to **int32** | Variant Default Child Identifier | [optional] 
**Variant** | Pointer to [**[]VariantVariant**](VariantVariant.md) |  | [optional] 
**Children** | Pointer to [**[]VariantChildren**](VariantChildren.md) |  | [optional] 

## Methods

### NewVariant

`func NewVariant() *Variant`

NewVariant instantiates a new Variant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariantWithDefaults

`func NewVariantWithDefaults() *Variant`

NewVariantWithDefaults instantiates a new Variant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParentId

`func (o *Variant) GetParentId() int32`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *Variant) GetParentIdOk() (*int32, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *Variant) SetParentId(v int32)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *Variant) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetDefaultChild

`func (o *Variant) GetDefaultChild() int32`

GetDefaultChild returns the DefaultChild field if non-nil, zero value otherwise.

### GetDefaultChildOk

`func (o *Variant) GetDefaultChildOk() (*int32, bool)`

GetDefaultChildOk returns a tuple with the DefaultChild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultChild

`func (o *Variant) SetDefaultChild(v int32)`

SetDefaultChild sets DefaultChild field to given value.

### HasDefaultChild

`func (o *Variant) HasDefaultChild() bool`

HasDefaultChild returns a boolean if a field has been set.

### GetVariant

`func (o *Variant) GetVariant() []VariantVariant`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *Variant) GetVariantOk() (*[]VariantVariant, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *Variant) SetVariant(v []VariantVariant)`

SetVariant sets Variant field to given value.

### HasVariant

`func (o *Variant) HasVariant() bool`

HasVariant returns a boolean if a field has been set.

### GetChildren

`func (o *Variant) GetChildren() []VariantChildren`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *Variant) GetChildrenOk() (*[]VariantChildren, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *Variant) SetChildren(v []VariantChildren)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *Variant) HasChildren() bool`

HasChildren returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


