# UpdateShopStatusRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShopId** | **int64** | Shop unique identifier | 
**Action** | **string** | Update the shop status. Value between open or close | 
**StartDate** | Pointer to **string** | Required for action : close. UNIX timestamp of date (hour, min, sec) from which the resolution ticket are  | [optional] 
**EndDate** | Pointer to **string** | Required for action : close. UNIX timestamp of date (hour, min, sec) to which the resolution ticket are  | [optional] 
**CloseWARN** | Pointer to **string** | Close shop description | [optional] 
**CloseNow** | Pointer to **bool** | Required for action : close. Value between true or false. If true, then immediately close shop &amp; ignore start_date’s date. If false it will follows start_date’s schedule to close  | [optional] 

## Methods

### NewUpdateShopStatusRequest

`func NewUpdateShopStatusRequest(shopId int64, action string, ) *UpdateShopStatusRequest`

NewUpdateShopStatusRequest instantiates a new UpdateShopStatusRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateShopStatusRequestWithDefaults

`func NewUpdateShopStatusRequestWithDefaults() *UpdateShopStatusRequest`

NewUpdateShopStatusRequestWithDefaults instantiates a new UpdateShopStatusRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShopId

`func (o *UpdateShopStatusRequest) GetShopId() int64`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *UpdateShopStatusRequest) GetShopIdOk() (*int64, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *UpdateShopStatusRequest) SetShopId(v int64)`

SetShopId sets ShopId field to given value.


### GetAction

`func (o *UpdateShopStatusRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *UpdateShopStatusRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *UpdateShopStatusRequest) SetAction(v string)`

SetAction sets Action field to given value.


### GetStartDate

`func (o *UpdateShopStatusRequest) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *UpdateShopStatusRequest) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *UpdateShopStatusRequest) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *UpdateShopStatusRequest) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *UpdateShopStatusRequest) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *UpdateShopStatusRequest) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *UpdateShopStatusRequest) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *UpdateShopStatusRequest) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetCloseWARN

`func (o *UpdateShopStatusRequest) GetCloseWARN() string`

GetCloseWARN returns the CloseWARN field if non-nil, zero value otherwise.

### GetCloseWARNOk

`func (o *UpdateShopStatusRequest) GetCloseWARNOk() (*string, bool)`

GetCloseWARNOk returns a tuple with the CloseWARN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseWARN

`func (o *UpdateShopStatusRequest) SetCloseWARN(v string)`

SetCloseWARN sets CloseWARN field to given value.

### HasCloseWARN

`func (o *UpdateShopStatusRequest) HasCloseWARN() bool`

HasCloseWARN returns a boolean if a field has been set.

### GetCloseNow

`func (o *UpdateShopStatusRequest) GetCloseNow() bool`

GetCloseNow returns the CloseNow field if non-nil, zero value otherwise.

### GetCloseNowOk

`func (o *UpdateShopStatusRequest) GetCloseNowOk() (*bool, bool)`

GetCloseNowOk returns a tuple with the CloseNow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseNow

`func (o *UpdateShopStatusRequest) SetCloseNow(v bool)`

SetCloseNow sets CloseNow field to given value.

### HasCloseNow

`func (o *UpdateShopStatusRequest) HasCloseNow() bool`

HasCloseNow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


