# InitiateChat200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contact** | Pointer to [**InitiateChat200ResponseDataContact**](InitiateChat200ResponseDataContact.md) |  | [optional] 
**IsSuccess** | Pointer to **bool** |  | [optional] 
**MsgId** | Pointer to **int64** |  | [optional] 

## Methods

### NewInitiateChat200ResponseData

`func NewInitiateChat200ResponseData() *InitiateChat200ResponseData`

NewInitiateChat200ResponseData instantiates a new InitiateChat200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInitiateChat200ResponseDataWithDefaults

`func NewInitiateChat200ResponseDataWithDefaults() *InitiateChat200ResponseData`

NewInitiateChat200ResponseDataWithDefaults instantiates a new InitiateChat200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContact

`func (o *InitiateChat200ResponseData) GetContact() InitiateChat200ResponseDataContact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *InitiateChat200ResponseData) GetContactOk() (*InitiateChat200ResponseDataContact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *InitiateChat200ResponseData) SetContact(v InitiateChat200ResponseDataContact)`

SetContact sets Contact field to given value.

### HasContact

`func (o *InitiateChat200ResponseData) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetIsSuccess

`func (o *InitiateChat200ResponseData) GetIsSuccess() bool`

GetIsSuccess returns the IsSuccess field if non-nil, zero value otherwise.

### GetIsSuccessOk

`func (o *InitiateChat200ResponseData) GetIsSuccessOk() (*bool, bool)`

GetIsSuccessOk returns a tuple with the IsSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuccess

`func (o *InitiateChat200ResponseData) SetIsSuccess(v bool)`

SetIsSuccess sets IsSuccess field to given value.

### HasIsSuccess

`func (o *InitiateChat200ResponseData) HasIsSuccess() bool`

HasIsSuccess returns a boolean if a field has been set.

### GetMsgId

`func (o *InitiateChat200ResponseData) GetMsgId() int64`

GetMsgId returns the MsgId field if non-nil, zero value otherwise.

### GetMsgIdOk

`func (o *InitiateChat200ResponseData) GetMsgIdOk() (*int64, bool)`

GetMsgIdOk returns a tuple with the MsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgId

`func (o *InitiateChat200ResponseData) SetMsgId(v int64)`

SetMsgId sets MsgId field to given value.

### HasMsgId

`func (o *InitiateChat200ResponseData) HasMsgId() bool`

HasMsgId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


