# UpdateShipmentInfo200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Response Status | [optional] 
**ServerProcessTime** | Pointer to **float64** | Server Process Time | [optional] 
**MessageStatus** | Pointer to **[]string** | Shipper Logo URL | [optional] 
**Data** | Pointer to [**UpdateShipmentInfo200ResponseDataData**](UpdateShipmentInfo200ResponseDataData.md) |  | [optional] 

## Methods

### NewUpdateShipmentInfo200ResponseData

`func NewUpdateShipmentInfo200ResponseData() *UpdateShipmentInfo200ResponseData`

NewUpdateShipmentInfo200ResponseData instantiates a new UpdateShipmentInfo200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateShipmentInfo200ResponseDataWithDefaults

`func NewUpdateShipmentInfo200ResponseDataWithDefaults() *UpdateShipmentInfo200ResponseData`

NewUpdateShipmentInfo200ResponseDataWithDefaults instantiates a new UpdateShipmentInfo200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *UpdateShipmentInfo200ResponseData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateShipmentInfo200ResponseData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateShipmentInfo200ResponseData) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateShipmentInfo200ResponseData) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetServerProcessTime

`func (o *UpdateShipmentInfo200ResponseData) GetServerProcessTime() float64`

GetServerProcessTime returns the ServerProcessTime field if non-nil, zero value otherwise.

### GetServerProcessTimeOk

`func (o *UpdateShipmentInfo200ResponseData) GetServerProcessTimeOk() (*float64, bool)`

GetServerProcessTimeOk returns a tuple with the ServerProcessTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerProcessTime

`func (o *UpdateShipmentInfo200ResponseData) SetServerProcessTime(v float64)`

SetServerProcessTime sets ServerProcessTime field to given value.

### HasServerProcessTime

`func (o *UpdateShipmentInfo200ResponseData) HasServerProcessTime() bool`

HasServerProcessTime returns a boolean if a field has been set.

### GetMessageStatus

`func (o *UpdateShipmentInfo200ResponseData) GetMessageStatus() []string`

GetMessageStatus returns the MessageStatus field if non-nil, zero value otherwise.

### GetMessageStatusOk

`func (o *UpdateShipmentInfo200ResponseData) GetMessageStatusOk() (*[]string, bool)`

GetMessageStatusOk returns a tuple with the MessageStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageStatus

`func (o *UpdateShipmentInfo200ResponseData) SetMessageStatus(v []string)`

SetMessageStatus sets MessageStatus field to given value.

### HasMessageStatus

`func (o *UpdateShipmentInfo200ResponseData) HasMessageStatus() bool`

HasMessageStatus returns a boolean if a field has been set.

### GetData

`func (o *UpdateShipmentInfo200ResponseData) GetData() UpdateShipmentInfo200ResponseDataData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UpdateShipmentInfo200ResponseData) GetDataOk() (*UpdateShipmentInfo200ResponseDataData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UpdateShipmentInfo200ResponseData) SetData(v UpdateShipmentInfo200ResponseDataData)`

SetData sets Data field to given value.

### HasData

`func (o *UpdateShipmentInfo200ResponseData) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


