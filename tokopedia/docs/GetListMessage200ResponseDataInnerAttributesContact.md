# GetListMessage200ResponseDataInnerAttributesContact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | User Unique Identifier | [optional] 
**Role** | Pointer to **string** | User Role, Field role has User, Shop Owner, Shop Admin, and Tokopedia Admin. | [optional] 
**Attributes** | Pointer to [**GetListMessage200ResponseDataInnerAttributesContactAttributes**](GetListMessage200ResponseDataInnerAttributesContactAttributes.md) |  | [optional] 

## Methods

### NewGetListMessage200ResponseDataInnerAttributesContact

`func NewGetListMessage200ResponseDataInnerAttributesContact() *GetListMessage200ResponseDataInnerAttributesContact`

NewGetListMessage200ResponseDataInnerAttributesContact instantiates a new GetListMessage200ResponseDataInnerAttributesContact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetListMessage200ResponseDataInnerAttributesContactWithDefaults

`func NewGetListMessage200ResponseDataInnerAttributesContactWithDefaults() *GetListMessage200ResponseDataInnerAttributesContact`

NewGetListMessage200ResponseDataInnerAttributesContactWithDefaults instantiates a new GetListMessage200ResponseDataInnerAttributesContact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetListMessage200ResponseDataInnerAttributesContact) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetListMessage200ResponseDataInnerAttributesContact) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetListMessage200ResponseDataInnerAttributesContact) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetListMessage200ResponseDataInnerAttributesContact) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRole

`func (o *GetListMessage200ResponseDataInnerAttributesContact) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *GetListMessage200ResponseDataInnerAttributesContact) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *GetListMessage200ResponseDataInnerAttributesContact) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *GetListMessage200ResponseDataInnerAttributesContact) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetAttributes

`func (o *GetListMessage200ResponseDataInnerAttributesContact) GetAttributes() GetListMessage200ResponseDataInnerAttributesContactAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GetListMessage200ResponseDataInnerAttributesContact) GetAttributesOk() (*GetListMessage200ResponseDataInnerAttributesContactAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GetListMessage200ResponseDataInnerAttributesContact) SetAttributes(v GetListMessage200ResponseDataInnerAttributesContactAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *GetListMessage200ResponseDataInnerAttributesContact) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


