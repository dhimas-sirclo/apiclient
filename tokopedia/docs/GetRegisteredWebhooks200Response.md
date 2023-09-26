# GetRegisteredWebhooks200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**GetRegisteredWebhooks200ResponseData**](GetRegisteredWebhooks200ResponseData.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**ErrorMessage** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGetRegisteredWebhooks200Response

`func NewGetRegisteredWebhooks200Response() *GetRegisteredWebhooks200Response`

NewGetRegisteredWebhooks200Response instantiates a new GetRegisteredWebhooks200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRegisteredWebhooks200ResponseWithDefaults

`func NewGetRegisteredWebhooks200ResponseWithDefaults() *GetRegisteredWebhooks200Response`

NewGetRegisteredWebhooks200ResponseWithDefaults instantiates a new GetRegisteredWebhooks200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetRegisteredWebhooks200Response) GetData() GetRegisteredWebhooks200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetRegisteredWebhooks200Response) GetDataOk() (*GetRegisteredWebhooks200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetRegisteredWebhooks200Response) SetData(v GetRegisteredWebhooks200ResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *GetRegisteredWebhooks200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetStatus

`func (o *GetRegisteredWebhooks200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetRegisteredWebhooks200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetRegisteredWebhooks200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetRegisteredWebhooks200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetErrorMessage

`func (o *GetRegisteredWebhooks200Response) GetErrorMessage() []string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *GetRegisteredWebhooks200Response) GetErrorMessageOk() (*[]string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *GetRegisteredWebhooks200Response) SetErrorMessage(v []string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *GetRegisteredWebhooks200Response) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


