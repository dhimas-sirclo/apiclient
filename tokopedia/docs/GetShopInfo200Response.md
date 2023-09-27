# GetShopInfo200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Header** | Pointer to [**ResponseHeader**](ResponseHeader.md) |  | [optional] 
**Data** | Pointer to [**[]GetShopInfo200ResponseDataInner**](GetShopInfo200ResponseDataInner.md) |  | [optional] 

## Methods

### NewGetShopInfo200Response

`func NewGetShopInfo200Response() *GetShopInfo200Response`

NewGetShopInfo200Response instantiates a new GetShopInfo200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetShopInfo200ResponseWithDefaults

`func NewGetShopInfo200ResponseWithDefaults() *GetShopInfo200Response`

NewGetShopInfo200ResponseWithDefaults instantiates a new GetShopInfo200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeader

`func (o *GetShopInfo200Response) GetHeader() ResponseHeader`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *GetShopInfo200Response) GetHeaderOk() (*ResponseHeader, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *GetShopInfo200Response) SetHeader(v ResponseHeader)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *GetShopInfo200Response) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetData

`func (o *GetShopInfo200Response) GetData() []GetShopInfo200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetShopInfo200Response) GetDataOk() (*[]GetShopInfo200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetShopInfo200Response) SetData(v []GetShopInfo200ResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *GetShopInfo200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


