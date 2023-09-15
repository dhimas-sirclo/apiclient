# VariantInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VariantId** | **string** |  | 
**VariantSku** | **string** |  | 
**VariantName** | **string** |  | 
**Status** | **string** |  | 
**Volume** | [**VolumeInput**](VolumeInput.md) |  | 
**Weight** | **float64** |  | 
**Attributes** | [**[]VariantAttributeInput**](VariantAttributeInput.md) |  | 
**VariantUrl** | **string** |  | 
**CurrencyCode** | **string** |  | 
**Price** | **float64** |  | 
**ImgUrls** | **[]string** |  | 

## Methods

### NewVariantInput

`func NewVariantInput(variantId string, variantSku string, variantName string, status string, volume VolumeInput, weight float64, attributes []VariantAttributeInput, variantUrl string, currencyCode string, price float64, imgUrls []string, ) *VariantInput`

NewVariantInput instantiates a new VariantInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariantInputWithDefaults

`func NewVariantInputWithDefaults() *VariantInput`

NewVariantInputWithDefaults instantiates a new VariantInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVariantId

`func (o *VariantInput) GetVariantId() string`

GetVariantId returns the VariantId field if non-nil, zero value otherwise.

### GetVariantIdOk

`func (o *VariantInput) GetVariantIdOk() (*string, bool)`

GetVariantIdOk returns a tuple with the VariantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantId

`func (o *VariantInput) SetVariantId(v string)`

SetVariantId sets VariantId field to given value.


### GetVariantSku

`func (o *VariantInput) GetVariantSku() string`

GetVariantSku returns the VariantSku field if non-nil, zero value otherwise.

### GetVariantSkuOk

`func (o *VariantInput) GetVariantSkuOk() (*string, bool)`

GetVariantSkuOk returns a tuple with the VariantSku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantSku

`func (o *VariantInput) SetVariantSku(v string)`

SetVariantSku sets VariantSku field to given value.


### GetVariantName

`func (o *VariantInput) GetVariantName() string`

GetVariantName returns the VariantName field if non-nil, zero value otherwise.

### GetVariantNameOk

`func (o *VariantInput) GetVariantNameOk() (*string, bool)`

GetVariantNameOk returns a tuple with the VariantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantName

`func (o *VariantInput) SetVariantName(v string)`

SetVariantName sets VariantName field to given value.


### GetStatus

`func (o *VariantInput) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VariantInput) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VariantInput) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetVolume

`func (o *VariantInput) GetVolume() VolumeInput`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *VariantInput) GetVolumeOk() (*VolumeInput, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *VariantInput) SetVolume(v VolumeInput)`

SetVolume sets Volume field to given value.


### GetWeight

`func (o *VariantInput) GetWeight() float64`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *VariantInput) GetWeightOk() (*float64, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *VariantInput) SetWeight(v float64)`

SetWeight sets Weight field to given value.


### GetAttributes

`func (o *VariantInput) GetAttributes() []VariantAttributeInput`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *VariantInput) GetAttributesOk() (*[]VariantAttributeInput, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *VariantInput) SetAttributes(v []VariantAttributeInput)`

SetAttributes sets Attributes field to given value.


### GetVariantUrl

`func (o *VariantInput) GetVariantUrl() string`

GetVariantUrl returns the VariantUrl field if non-nil, zero value otherwise.

### GetVariantUrlOk

`func (o *VariantInput) GetVariantUrlOk() (*string, bool)`

GetVariantUrlOk returns a tuple with the VariantUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantUrl

`func (o *VariantInput) SetVariantUrl(v string)`

SetVariantUrl sets VariantUrl field to given value.


### GetCurrencyCode

`func (o *VariantInput) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *VariantInput) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *VariantInput) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetPrice

`func (o *VariantInput) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *VariantInput) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *VariantInput) SetPrice(v float64)`

SetPrice sets Price field to given value.


### GetImgUrls

`func (o *VariantInput) GetImgUrls() []string`

GetImgUrls returns the ImgUrls field if non-nil, zero value otherwise.

### GetImgUrlsOk

`func (o *VariantInput) GetImgUrlsOk() (*[]string, bool)`

GetImgUrlsOk returns a tuple with the ImgUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImgUrls

`func (o *VariantInput) SetImgUrls(v []string)`

SetImgUrls sets ImgUrls field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


