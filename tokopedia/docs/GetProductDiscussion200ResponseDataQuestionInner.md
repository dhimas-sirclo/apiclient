# GetProductDiscussion200ResponseDataQuestionInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **string** | Discussion Content | [optional] 
**MaskedContent** | Pointer to **string** | Discussion Masked Content | [optional] 
**UserName** | Pointer to **string** | User username | [optional] 
**UserID** | Pointer to **int64** | User userID | [optional] 
**CreateTime** | Pointer to **string** | Discussion Timestamp (format: 2020-12-04T13:40:19Z) | [optional] 
**CreateTimeFormatted** | Pointer to **string** | Discussion Timestamp Formatted | [optional] 
**TotalAnswer** | Pointer to **int64** | Total Answer Count | [optional] 
**Answer** | Pointer to [**[]GetProductDiscussion200ResponseDataQuestionInnerAnswerInner**](GetProductDiscussion200ResponseDataQuestionInnerAnswerInner.md) |  | [optional] 
**QuestionID** | Pointer to **int64** | Question Unique Identifier | [optional] 
**AnswererThumbnail** | Pointer to **string** | Answerer Thumbnail URL | [optional] 
**UserThumbnail** | Pointer to **string** | User Thumbnail URL | [optional] 

## Methods

### NewGetProductDiscussion200ResponseDataQuestionInner

`func NewGetProductDiscussion200ResponseDataQuestionInner() *GetProductDiscussion200ResponseDataQuestionInner`

NewGetProductDiscussion200ResponseDataQuestionInner instantiates a new GetProductDiscussion200ResponseDataQuestionInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetProductDiscussion200ResponseDataQuestionInnerWithDefaults

`func NewGetProductDiscussion200ResponseDataQuestionInnerWithDefaults() *GetProductDiscussion200ResponseDataQuestionInner`

NewGetProductDiscussion200ResponseDataQuestionInnerWithDefaults instantiates a new GetProductDiscussion200ResponseDataQuestionInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *GetProductDiscussion200ResponseDataQuestionInner) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *GetProductDiscussion200ResponseDataQuestionInner) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *GetProductDiscussion200ResponseDataQuestionInner) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *GetProductDiscussion200ResponseDataQuestionInner) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetMaskedContent

`func (o *GetProductDiscussion200ResponseDataQuestionInner) GetMaskedContent() string`

GetMaskedContent returns the MaskedContent field if non-nil, zero value otherwise.

### GetMaskedContentOk

`func (o *GetProductDiscussion200ResponseDataQuestionInner) GetMaskedContentOk() (*string, bool)`

GetMaskedContentOk returns a tuple with the MaskedContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskedContent

`func (o *GetProductDiscussion200ResponseDataQuestionInner) SetMaskedContent(v string)`

SetMaskedContent sets MaskedContent field to given value.

### HasMaskedContent

`func (o *GetProductDiscussion200ResponseDataQuestionInner) HasMaskedContent() bool`

HasMaskedContent returns a boolean if a field has been set.

### GetUserName

`func (o *GetProductDiscussion200ResponseDataQuestionInner) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *GetProductDiscussion200ResponseDataQuestionInner) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *GetProductDiscussion200ResponseDataQuestionInner) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *GetProductDiscussion200ResponseDataQuestionInner) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetUserID

`func (o *GetProductDiscussion200ResponseDataQuestionInner) GetUserID() int64`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *GetProductDiscussion200ResponseDataQuestionInner) GetUserIDOk() (*int64, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *GetProductDiscussion200ResponseDataQuestionInner) SetUserID(v int64)`

SetUserID sets UserID field to given value.

### HasUserID

`func (o *GetProductDiscussion200ResponseDataQuestionInner) HasUserID() bool`

HasUserID returns a boolean if a field has been set.

### GetCreateTime

`func (o *GetProductDiscussion200ResponseDataQuestionInner) GetCreateTime() string`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *GetProductDiscussion200ResponseDataQuestionInner) GetCreateTimeOk() (*string, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *GetProductDiscussion200ResponseDataQuestionInner) SetCreateTime(v string)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *GetProductDiscussion200ResponseDataQuestionInner) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCreateTimeFormatted

`func (o *GetProductDiscussion200ResponseDataQuestionInner) GetCreateTimeFormatted() string`

GetCreateTimeFormatted returns the CreateTimeFormatted field if non-nil, zero value otherwise.

### GetCreateTimeFormattedOk

`func (o *GetProductDiscussion200ResponseDataQuestionInner) GetCreateTimeFormattedOk() (*string, bool)`

GetCreateTimeFormattedOk returns a tuple with the CreateTimeFormatted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTimeFormatted

`func (o *GetProductDiscussion200ResponseDataQuestionInner) SetCreateTimeFormatted(v string)`

SetCreateTimeFormatted sets CreateTimeFormatted field to given value.

### HasCreateTimeFormatted

`func (o *GetProductDiscussion200ResponseDataQuestionInner) HasCreateTimeFormatted() bool`

HasCreateTimeFormatted returns a boolean if a field has been set.

### GetTotalAnswer

`func (o *GetProductDiscussion200ResponseDataQuestionInner) GetTotalAnswer() int64`

GetTotalAnswer returns the TotalAnswer field if non-nil, zero value otherwise.

### GetTotalAnswerOk

`func (o *GetProductDiscussion200ResponseDataQuestionInner) GetTotalAnswerOk() (*int64, bool)`

GetTotalAnswerOk returns a tuple with the TotalAnswer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAnswer

`func (o *GetProductDiscussion200ResponseDataQuestionInner) SetTotalAnswer(v int64)`

SetTotalAnswer sets TotalAnswer field to given value.

### HasTotalAnswer

`func (o *GetProductDiscussion200ResponseDataQuestionInner) HasTotalAnswer() bool`

HasTotalAnswer returns a boolean if a field has been set.

### GetAnswer

`func (o *GetProductDiscussion200ResponseDataQuestionInner) GetAnswer() []GetProductDiscussion200ResponseDataQuestionInnerAnswerInner`

GetAnswer returns the Answer field if non-nil, zero value otherwise.

### GetAnswerOk

`func (o *GetProductDiscussion200ResponseDataQuestionInner) GetAnswerOk() (*[]GetProductDiscussion200ResponseDataQuestionInnerAnswerInner, bool)`

GetAnswerOk returns a tuple with the Answer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswer

`func (o *GetProductDiscussion200ResponseDataQuestionInner) SetAnswer(v []GetProductDiscussion200ResponseDataQuestionInnerAnswerInner)`

SetAnswer sets Answer field to given value.

### HasAnswer

`func (o *GetProductDiscussion200ResponseDataQuestionInner) HasAnswer() bool`

HasAnswer returns a boolean if a field has been set.

### GetQuestionID

`func (o *GetProductDiscussion200ResponseDataQuestionInner) GetQuestionID() int64`

GetQuestionID returns the QuestionID field if non-nil, zero value otherwise.

### GetQuestionIDOk

`func (o *GetProductDiscussion200ResponseDataQuestionInner) GetQuestionIDOk() (*int64, bool)`

GetQuestionIDOk returns a tuple with the QuestionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionID

`func (o *GetProductDiscussion200ResponseDataQuestionInner) SetQuestionID(v int64)`

SetQuestionID sets QuestionID field to given value.

### HasQuestionID

`func (o *GetProductDiscussion200ResponseDataQuestionInner) HasQuestionID() bool`

HasQuestionID returns a boolean if a field has been set.

### GetAnswererThumbnail

`func (o *GetProductDiscussion200ResponseDataQuestionInner) GetAnswererThumbnail() string`

GetAnswererThumbnail returns the AnswererThumbnail field if non-nil, zero value otherwise.

### GetAnswererThumbnailOk

`func (o *GetProductDiscussion200ResponseDataQuestionInner) GetAnswererThumbnailOk() (*string, bool)`

GetAnswererThumbnailOk returns a tuple with the AnswererThumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswererThumbnail

`func (o *GetProductDiscussion200ResponseDataQuestionInner) SetAnswererThumbnail(v string)`

SetAnswererThumbnail sets AnswererThumbnail field to given value.

### HasAnswererThumbnail

`func (o *GetProductDiscussion200ResponseDataQuestionInner) HasAnswererThumbnail() bool`

HasAnswererThumbnail returns a boolean if a field has been set.

### GetUserThumbnail

`func (o *GetProductDiscussion200ResponseDataQuestionInner) GetUserThumbnail() string`

GetUserThumbnail returns the UserThumbnail field if non-nil, zero value otherwise.

### GetUserThumbnailOk

`func (o *GetProductDiscussion200ResponseDataQuestionInner) GetUserThumbnailOk() (*string, bool)`

GetUserThumbnailOk returns a tuple with the UserThumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserThumbnail

`func (o *GetProductDiscussion200ResponseDataQuestionInner) SetUserThumbnail(v string)`

SetUserThumbnail sets UserThumbnail field to given value.

### HasUserThumbnail

`func (o *GetProductDiscussion200ResponseDataQuestionInner) HasUserThumbnail() bool`

HasUserThumbnail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


