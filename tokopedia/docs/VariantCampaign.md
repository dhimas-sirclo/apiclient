# VariantCampaign

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsActive** | Pointer to **bool** |  | [optional] 
**DiscountedPercentage** | Pointer to **int64** |  | [optional] 
**DiscountedPrice** | Pointer to **float64** |  | [optional] 
**DiscountedPriceFmt** | Pointer to **string** |  | [optional] 
**CampaignType** | Pointer to **int64** |  | [optional] 
**CampaignTypeName** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **string** |  | [optional] 
**EndDate** | Pointer to **string** |  | [optional] 

## Methods

### NewVariantCampaign

`func NewVariantCampaign() *VariantCampaign`

NewVariantCampaign instantiates a new VariantCampaign object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariantCampaignWithDefaults

`func NewVariantCampaignWithDefaults() *VariantCampaign`

NewVariantCampaignWithDefaults instantiates a new VariantCampaign object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsActive

`func (o *VariantCampaign) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *VariantCampaign) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *VariantCampaign) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *VariantCampaign) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetDiscountedPercentage

`func (o *VariantCampaign) GetDiscountedPercentage() int64`

GetDiscountedPercentage returns the DiscountedPercentage field if non-nil, zero value otherwise.

### GetDiscountedPercentageOk

`func (o *VariantCampaign) GetDiscountedPercentageOk() (*int64, bool)`

GetDiscountedPercentageOk returns a tuple with the DiscountedPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountedPercentage

`func (o *VariantCampaign) SetDiscountedPercentage(v int64)`

SetDiscountedPercentage sets DiscountedPercentage field to given value.

### HasDiscountedPercentage

`func (o *VariantCampaign) HasDiscountedPercentage() bool`

HasDiscountedPercentage returns a boolean if a field has been set.

### GetDiscountedPrice

`func (o *VariantCampaign) GetDiscountedPrice() float64`

GetDiscountedPrice returns the DiscountedPrice field if non-nil, zero value otherwise.

### GetDiscountedPriceOk

`func (o *VariantCampaign) GetDiscountedPriceOk() (*float64, bool)`

GetDiscountedPriceOk returns a tuple with the DiscountedPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountedPrice

`func (o *VariantCampaign) SetDiscountedPrice(v float64)`

SetDiscountedPrice sets DiscountedPrice field to given value.

### HasDiscountedPrice

`func (o *VariantCampaign) HasDiscountedPrice() bool`

HasDiscountedPrice returns a boolean if a field has been set.

### GetDiscountedPriceFmt

`func (o *VariantCampaign) GetDiscountedPriceFmt() string`

GetDiscountedPriceFmt returns the DiscountedPriceFmt field if non-nil, zero value otherwise.

### GetDiscountedPriceFmtOk

`func (o *VariantCampaign) GetDiscountedPriceFmtOk() (*string, bool)`

GetDiscountedPriceFmtOk returns a tuple with the DiscountedPriceFmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountedPriceFmt

`func (o *VariantCampaign) SetDiscountedPriceFmt(v string)`

SetDiscountedPriceFmt sets DiscountedPriceFmt field to given value.

### HasDiscountedPriceFmt

`func (o *VariantCampaign) HasDiscountedPriceFmt() bool`

HasDiscountedPriceFmt returns a boolean if a field has been set.

### GetCampaignType

`func (o *VariantCampaign) GetCampaignType() int64`

GetCampaignType returns the CampaignType field if non-nil, zero value otherwise.

### GetCampaignTypeOk

`func (o *VariantCampaign) GetCampaignTypeOk() (*int64, bool)`

GetCampaignTypeOk returns a tuple with the CampaignType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignType

`func (o *VariantCampaign) SetCampaignType(v int64)`

SetCampaignType sets CampaignType field to given value.

### HasCampaignType

`func (o *VariantCampaign) HasCampaignType() bool`

HasCampaignType returns a boolean if a field has been set.

### GetCampaignTypeName

`func (o *VariantCampaign) GetCampaignTypeName() string`

GetCampaignTypeName returns the CampaignTypeName field if non-nil, zero value otherwise.

### GetCampaignTypeNameOk

`func (o *VariantCampaign) GetCampaignTypeNameOk() (*string, bool)`

GetCampaignTypeNameOk returns a tuple with the CampaignTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignTypeName

`func (o *VariantCampaign) SetCampaignTypeName(v string)`

SetCampaignTypeName sets CampaignTypeName field to given value.

### HasCampaignTypeName

`func (o *VariantCampaign) HasCampaignTypeName() bool`

HasCampaignTypeName returns a boolean if a field has been set.

### GetStartDate

`func (o *VariantCampaign) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *VariantCampaign) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *VariantCampaign) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *VariantCampaign) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *VariantCampaign) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *VariantCampaign) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *VariantCampaign) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *VariantCampaign) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


