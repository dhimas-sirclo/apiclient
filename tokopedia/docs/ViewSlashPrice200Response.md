# ViewSlashPrice200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Header** | Pointer to [**ResponseHeader**](ResponseHeader.md) |  | [optional] 
**Data** | Pointer to [**[]ViewSlashPrice200ResponseDataInner**](ViewSlashPrice200ResponseDataInner.md) |  | [optional] 

## Methods

### NewViewSlashPrice200Response

`func NewViewSlashPrice200Response() *ViewSlashPrice200Response`

NewViewSlashPrice200Response instantiates a new ViewSlashPrice200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewSlashPrice200ResponseWithDefaults

`func NewViewSlashPrice200ResponseWithDefaults() *ViewSlashPrice200Response`

NewViewSlashPrice200ResponseWithDefaults instantiates a new ViewSlashPrice200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeader

`func (o *ViewSlashPrice200Response) GetHeader() ResponseHeader`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *ViewSlashPrice200Response) GetHeaderOk() (*ResponseHeader, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *ViewSlashPrice200Response) SetHeader(v ResponseHeader)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *ViewSlashPrice200Response) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetData

`func (o *ViewSlashPrice200Response) GetData() []ViewSlashPrice200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ViewSlashPrice200Response) GetDataOk() (*[]ViewSlashPrice200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ViewSlashPrice200Response) SetData(v []ViewSlashPrice200ResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *ViewSlashPrice200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


