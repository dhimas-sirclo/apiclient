# Header

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProcessTime** | Pointer to **float32** |  | [optional] 
**Messages** | Pointer to **string** |  | [optional] 

## Methods

### NewHeader

`func NewHeader() *Header`

NewHeader instantiates a new Header object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeaderWithDefaults

`func NewHeaderWithDefaults() *Header`

NewHeaderWithDefaults instantiates a new Header object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProcessTime

`func (o *Header) GetProcessTime() float32`

GetProcessTime returns the ProcessTime field if non-nil, zero value otherwise.

### GetProcessTimeOk

`func (o *Header) GetProcessTimeOk() (*float32, bool)`

GetProcessTimeOk returns a tuple with the ProcessTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessTime

`func (o *Header) SetProcessTime(v float32)`

SetProcessTime sets ProcessTime field to given value.

### HasProcessTime

`func (o *Header) HasProcessTime() bool`

HasProcessTime returns a boolean if a field has been set.

### GetMessages

`func (o *Header) GetMessages() string`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *Header) GetMessagesOk() (*string, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *Header) SetMessages(v string)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *Header) HasMessages() bool`

HasMessages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


