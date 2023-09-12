# ErrorHeader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProcessTime** | Pointer to **float64** |  | [optional] 
**Messages** | Pointer to **string** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**ErrorCode** | Pointer to **string** |  | [optional] 

## Methods

### NewErrorHeader

`func NewErrorHeader() *ErrorHeader`

NewErrorHeader instantiates a new ErrorHeader object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorHeaderWithDefaults

`func NewErrorHeaderWithDefaults() *ErrorHeader`

NewErrorHeaderWithDefaults instantiates a new ErrorHeader object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProcessTime

`func (o *ErrorHeader) GetProcessTime() float64`

GetProcessTime returns the ProcessTime field if non-nil, zero value otherwise.

### GetProcessTimeOk

`func (o *ErrorHeader) GetProcessTimeOk() (*float64, bool)`

GetProcessTimeOk returns a tuple with the ProcessTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessTime

`func (o *ErrorHeader) SetProcessTime(v float64)`

SetProcessTime sets ProcessTime field to given value.

### HasProcessTime

`func (o *ErrorHeader) HasProcessTime() bool`

HasProcessTime returns a boolean if a field has been set.

### GetMessages

`func (o *ErrorHeader) GetMessages() string`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *ErrorHeader) GetMessagesOk() (*string, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *ErrorHeader) SetMessages(v string)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *ErrorHeader) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetReason

`func (o *ErrorHeader) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ErrorHeader) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ErrorHeader) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ErrorHeader) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetErrorCode

`func (o *ErrorHeader) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *ErrorHeader) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *ErrorHeader) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *ErrorHeader) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


