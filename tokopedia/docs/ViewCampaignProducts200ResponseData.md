# ViewCampaignProducts200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Products** | Pointer to [**[]ViewCampaignProducts200ResponseDataProductsInner**](ViewCampaignProducts200ResponseDataProductsInner.md) |  | [optional] 
**TotalProduct** | Pointer to **int64** | Total Product Count | [optional] 

## Methods

### NewViewCampaignProducts200ResponseData

`func NewViewCampaignProducts200ResponseData() *ViewCampaignProducts200ResponseData`

NewViewCampaignProducts200ResponseData instantiates a new ViewCampaignProducts200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewCampaignProducts200ResponseDataWithDefaults

`func NewViewCampaignProducts200ResponseDataWithDefaults() *ViewCampaignProducts200ResponseData`

NewViewCampaignProducts200ResponseDataWithDefaults instantiates a new ViewCampaignProducts200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProducts

`func (o *ViewCampaignProducts200ResponseData) GetProducts() []ViewCampaignProducts200ResponseDataProductsInner`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *ViewCampaignProducts200ResponseData) GetProductsOk() (*[]ViewCampaignProducts200ResponseDataProductsInner, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *ViewCampaignProducts200ResponseData) SetProducts(v []ViewCampaignProducts200ResponseDataProductsInner)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *ViewCampaignProducts200ResponseData) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetTotalProduct

`func (o *ViewCampaignProducts200ResponseData) GetTotalProduct() int64`

GetTotalProduct returns the TotalProduct field if non-nil, zero value otherwise.

### GetTotalProductOk

`func (o *ViewCampaignProducts200ResponseData) GetTotalProductOk() (*int64, bool)`

GetTotalProductOk returns a tuple with the TotalProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalProduct

`func (o *ViewCampaignProducts200ResponseData) SetTotalProduct(v int64)`

SetTotalProduct sets TotalProduct field to given value.

### HasTotalProduct

`func (o *ViewCampaignProducts200ResponseData) HasTotalProduct() bool`

HasTotalProduct returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


