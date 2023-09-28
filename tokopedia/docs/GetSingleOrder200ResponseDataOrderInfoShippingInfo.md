# GetSingleOrder200ResponseDataOrderInfoShippingInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SpId** | Pointer to **int64** | Shipper Unique Identifier | [optional] 
**ShippingId** | Pointer to **int64** | Shipping Unique Identifier | [optional] 
**LogisticName** | Pointer to **string** | Logistic Name | [optional] 
**LogisticService** | Pointer to **string** | Logistic Service Name | [optional] 
**ShippingPrice** | Pointer to **int64** | Shipping Price | [optional] 
**ShippingPriceRate** | Pointer to **int64** | Shipping Price Rate | [optional] 
**ShippingFee** | Pointer to **int64** | Shipping Fee | [optional] 
**InsurancePrice** | Pointer to **int64** | Insurance Price | [optional] 
**Fee** | Pointer to **int64** | Fee | [optional] 
**IsChangeCourier** | Pointer to **bool** | Is change courirer? | [optional] 
**SecondSpId** | Pointer to **int64** | Second Shipper Unique Identifier | [optional] 
**SecondShippingId** | Pointer to **int64** | Second Shipping Unique Identifier | [optional] 
**SecondLogisticName** | Pointer to **string** | Second Logistic Name | [optional] 
**SecondLogisticService** | Pointer to **string** | Second Logistic Service Name | [optional] 
**SecondAgencyFee** | Pointer to **int64** | Second Agency Fee | [optional] 
**SecondInsurance** | Pointer to **int64** | Second Insurance | [optional] 
**SecondRate** | Pointer to **int64** | Second Shipping Price Rate | [optional] 
**Awb** | Pointer to **string** | Airway Bill (Resi) | [optional] 
**AutoresiCashlessStatus** | Pointer to **int64** | Autoresi Cashless Status | [optional] 
**AutoresiAwb** | Pointer to **string** | Airway Bill (Auto Resi) | [optional] 
**AutoresiShippingPrice** | Pointer to **int64** | AWB Shipping Price | [optional] 
**CountAwb** | Pointer to **int64** | AWB Count | [optional] 
**IsCashless** | Pointer to **bool** | Is cashless? | [optional] 
**IsFakeDelivery** | Pointer to **bool** | Is fake delivery? | [optional] 
**RecommendedCourierInfo** | Pointer to [**[]GetSingleOrder200ResponseDataOrderInfoShippingInfoRecommendedCourierInfoInner**](GetSingleOrder200ResponseDataOrderInfoShippingInfoRecommendedCourierInfoInner.md) |  | [optional] 

## Methods

### NewGetSingleOrder200ResponseDataOrderInfoShippingInfo

`func NewGetSingleOrder200ResponseDataOrderInfoShippingInfo() *GetSingleOrder200ResponseDataOrderInfoShippingInfo`

NewGetSingleOrder200ResponseDataOrderInfoShippingInfo instantiates a new GetSingleOrder200ResponseDataOrderInfoShippingInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSingleOrder200ResponseDataOrderInfoShippingInfoWithDefaults

`func NewGetSingleOrder200ResponseDataOrderInfoShippingInfoWithDefaults() *GetSingleOrder200ResponseDataOrderInfoShippingInfo`

NewGetSingleOrder200ResponseDataOrderInfoShippingInfoWithDefaults instantiates a new GetSingleOrder200ResponseDataOrderInfoShippingInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpId

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetSpId() int64`

GetSpId returns the SpId field if non-nil, zero value otherwise.

### GetSpIdOk

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetSpIdOk() (*int64, bool)`

GetSpIdOk returns a tuple with the SpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpId

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) SetSpId(v int64)`

SetSpId sets SpId field to given value.

### HasSpId

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) HasSpId() bool`

HasSpId returns a boolean if a field has been set.

### GetShippingId

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetShippingId() int64`

GetShippingId returns the ShippingId field if non-nil, zero value otherwise.

### GetShippingIdOk

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetShippingIdOk() (*int64, bool)`

GetShippingIdOk returns a tuple with the ShippingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingId

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) SetShippingId(v int64)`

SetShippingId sets ShippingId field to given value.

### HasShippingId

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) HasShippingId() bool`

HasShippingId returns a boolean if a field has been set.

### GetLogisticName

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetLogisticName() string`

GetLogisticName returns the LogisticName field if non-nil, zero value otherwise.

### GetLogisticNameOk

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetLogisticNameOk() (*string, bool)`

GetLogisticNameOk returns a tuple with the LogisticName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogisticName

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) SetLogisticName(v string)`

SetLogisticName sets LogisticName field to given value.

### HasLogisticName

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) HasLogisticName() bool`

HasLogisticName returns a boolean if a field has been set.

### GetLogisticService

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetLogisticService() string`

GetLogisticService returns the LogisticService field if non-nil, zero value otherwise.

### GetLogisticServiceOk

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetLogisticServiceOk() (*string, bool)`

GetLogisticServiceOk returns a tuple with the LogisticService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogisticService

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) SetLogisticService(v string)`

SetLogisticService sets LogisticService field to given value.

### HasLogisticService

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) HasLogisticService() bool`

HasLogisticService returns a boolean if a field has been set.

### GetShippingPrice

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetShippingPrice() int64`

GetShippingPrice returns the ShippingPrice field if non-nil, zero value otherwise.

### GetShippingPriceOk

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetShippingPriceOk() (*int64, bool)`

GetShippingPriceOk returns a tuple with the ShippingPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingPrice

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) SetShippingPrice(v int64)`

SetShippingPrice sets ShippingPrice field to given value.

### HasShippingPrice

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) HasShippingPrice() bool`

HasShippingPrice returns a boolean if a field has been set.

### GetShippingPriceRate

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetShippingPriceRate() int64`

GetShippingPriceRate returns the ShippingPriceRate field if non-nil, zero value otherwise.

### GetShippingPriceRateOk

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetShippingPriceRateOk() (*int64, bool)`

GetShippingPriceRateOk returns a tuple with the ShippingPriceRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingPriceRate

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) SetShippingPriceRate(v int64)`

SetShippingPriceRate sets ShippingPriceRate field to given value.

### HasShippingPriceRate

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) HasShippingPriceRate() bool`

HasShippingPriceRate returns a boolean if a field has been set.

### GetShippingFee

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetShippingFee() int64`

GetShippingFee returns the ShippingFee field if non-nil, zero value otherwise.

### GetShippingFeeOk

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetShippingFeeOk() (*int64, bool)`

GetShippingFeeOk returns a tuple with the ShippingFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingFee

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) SetShippingFee(v int64)`

SetShippingFee sets ShippingFee field to given value.

### HasShippingFee

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) HasShippingFee() bool`

HasShippingFee returns a boolean if a field has been set.

### GetInsurancePrice

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetInsurancePrice() int64`

GetInsurancePrice returns the InsurancePrice field if non-nil, zero value otherwise.

### GetInsurancePriceOk

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetInsurancePriceOk() (*int64, bool)`

GetInsurancePriceOk returns a tuple with the InsurancePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsurancePrice

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) SetInsurancePrice(v int64)`

SetInsurancePrice sets InsurancePrice field to given value.

### HasInsurancePrice

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) HasInsurancePrice() bool`

HasInsurancePrice returns a boolean if a field has been set.

### GetFee

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetFee() int64`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetFeeOk() (*int64, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) SetFee(v int64)`

SetFee sets Fee field to given value.

### HasFee

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetIsChangeCourier

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetIsChangeCourier() bool`

GetIsChangeCourier returns the IsChangeCourier field if non-nil, zero value otherwise.

### GetIsChangeCourierOk

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetIsChangeCourierOk() (*bool, bool)`

GetIsChangeCourierOk returns a tuple with the IsChangeCourier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsChangeCourier

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) SetIsChangeCourier(v bool)`

SetIsChangeCourier sets IsChangeCourier field to given value.

### HasIsChangeCourier

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) HasIsChangeCourier() bool`

HasIsChangeCourier returns a boolean if a field has been set.

### GetSecondSpId

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetSecondSpId() int64`

GetSecondSpId returns the SecondSpId field if non-nil, zero value otherwise.

### GetSecondSpIdOk

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetSecondSpIdOk() (*int64, bool)`

GetSecondSpIdOk returns a tuple with the SecondSpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondSpId

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) SetSecondSpId(v int64)`

SetSecondSpId sets SecondSpId field to given value.

### HasSecondSpId

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) HasSecondSpId() bool`

HasSecondSpId returns a boolean if a field has been set.

### GetSecondShippingId

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetSecondShippingId() int64`

GetSecondShippingId returns the SecondShippingId field if non-nil, zero value otherwise.

### GetSecondShippingIdOk

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetSecondShippingIdOk() (*int64, bool)`

GetSecondShippingIdOk returns a tuple with the SecondShippingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondShippingId

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) SetSecondShippingId(v int64)`

SetSecondShippingId sets SecondShippingId field to given value.

### HasSecondShippingId

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) HasSecondShippingId() bool`

HasSecondShippingId returns a boolean if a field has been set.

### GetSecondLogisticName

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetSecondLogisticName() string`

GetSecondLogisticName returns the SecondLogisticName field if non-nil, zero value otherwise.

### GetSecondLogisticNameOk

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetSecondLogisticNameOk() (*string, bool)`

GetSecondLogisticNameOk returns a tuple with the SecondLogisticName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondLogisticName

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) SetSecondLogisticName(v string)`

SetSecondLogisticName sets SecondLogisticName field to given value.

### HasSecondLogisticName

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) HasSecondLogisticName() bool`

HasSecondLogisticName returns a boolean if a field has been set.

### GetSecondLogisticService

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetSecondLogisticService() string`

GetSecondLogisticService returns the SecondLogisticService field if non-nil, zero value otherwise.

### GetSecondLogisticServiceOk

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetSecondLogisticServiceOk() (*string, bool)`

GetSecondLogisticServiceOk returns a tuple with the SecondLogisticService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondLogisticService

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) SetSecondLogisticService(v string)`

SetSecondLogisticService sets SecondLogisticService field to given value.

### HasSecondLogisticService

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) HasSecondLogisticService() bool`

HasSecondLogisticService returns a boolean if a field has been set.

### GetSecondAgencyFee

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetSecondAgencyFee() int64`

GetSecondAgencyFee returns the SecondAgencyFee field if non-nil, zero value otherwise.

### GetSecondAgencyFeeOk

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetSecondAgencyFeeOk() (*int64, bool)`

GetSecondAgencyFeeOk returns a tuple with the SecondAgencyFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondAgencyFee

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) SetSecondAgencyFee(v int64)`

SetSecondAgencyFee sets SecondAgencyFee field to given value.

### HasSecondAgencyFee

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) HasSecondAgencyFee() bool`

HasSecondAgencyFee returns a boolean if a field has been set.

### GetSecondInsurance

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetSecondInsurance() int64`

GetSecondInsurance returns the SecondInsurance field if non-nil, zero value otherwise.

### GetSecondInsuranceOk

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetSecondInsuranceOk() (*int64, bool)`

GetSecondInsuranceOk returns a tuple with the SecondInsurance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondInsurance

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) SetSecondInsurance(v int64)`

SetSecondInsurance sets SecondInsurance field to given value.

### HasSecondInsurance

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) HasSecondInsurance() bool`

HasSecondInsurance returns a boolean if a field has been set.

### GetSecondRate

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetSecondRate() int64`

GetSecondRate returns the SecondRate field if non-nil, zero value otherwise.

### GetSecondRateOk

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetSecondRateOk() (*int64, bool)`

GetSecondRateOk returns a tuple with the SecondRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondRate

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) SetSecondRate(v int64)`

SetSecondRate sets SecondRate field to given value.

### HasSecondRate

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) HasSecondRate() bool`

HasSecondRate returns a boolean if a field has been set.

### GetAwb

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetAwb() string`

GetAwb returns the Awb field if non-nil, zero value otherwise.

### GetAwbOk

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetAwbOk() (*string, bool)`

GetAwbOk returns a tuple with the Awb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwb

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) SetAwb(v string)`

SetAwb sets Awb field to given value.

### HasAwb

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) HasAwb() bool`

HasAwb returns a boolean if a field has been set.

### GetAutoresiCashlessStatus

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetAutoresiCashlessStatus() int64`

GetAutoresiCashlessStatus returns the AutoresiCashlessStatus field if non-nil, zero value otherwise.

### GetAutoresiCashlessStatusOk

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetAutoresiCashlessStatusOk() (*int64, bool)`

GetAutoresiCashlessStatusOk returns a tuple with the AutoresiCashlessStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoresiCashlessStatus

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) SetAutoresiCashlessStatus(v int64)`

SetAutoresiCashlessStatus sets AutoresiCashlessStatus field to given value.

### HasAutoresiCashlessStatus

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) HasAutoresiCashlessStatus() bool`

HasAutoresiCashlessStatus returns a boolean if a field has been set.

### GetAutoresiAwb

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetAutoresiAwb() string`

GetAutoresiAwb returns the AutoresiAwb field if non-nil, zero value otherwise.

### GetAutoresiAwbOk

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetAutoresiAwbOk() (*string, bool)`

GetAutoresiAwbOk returns a tuple with the AutoresiAwb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoresiAwb

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) SetAutoresiAwb(v string)`

SetAutoresiAwb sets AutoresiAwb field to given value.

### HasAutoresiAwb

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) HasAutoresiAwb() bool`

HasAutoresiAwb returns a boolean if a field has been set.

### GetAutoresiShippingPrice

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetAutoresiShippingPrice() int64`

GetAutoresiShippingPrice returns the AutoresiShippingPrice field if non-nil, zero value otherwise.

### GetAutoresiShippingPriceOk

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetAutoresiShippingPriceOk() (*int64, bool)`

GetAutoresiShippingPriceOk returns a tuple with the AutoresiShippingPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoresiShippingPrice

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) SetAutoresiShippingPrice(v int64)`

SetAutoresiShippingPrice sets AutoresiShippingPrice field to given value.

### HasAutoresiShippingPrice

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) HasAutoresiShippingPrice() bool`

HasAutoresiShippingPrice returns a boolean if a field has been set.

### GetCountAwb

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetCountAwb() int64`

GetCountAwb returns the CountAwb field if non-nil, zero value otherwise.

### GetCountAwbOk

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetCountAwbOk() (*int64, bool)`

GetCountAwbOk returns a tuple with the CountAwb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountAwb

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) SetCountAwb(v int64)`

SetCountAwb sets CountAwb field to given value.

### HasCountAwb

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) HasCountAwb() bool`

HasCountAwb returns a boolean if a field has been set.

### GetIsCashless

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetIsCashless() bool`

GetIsCashless returns the IsCashless field if non-nil, zero value otherwise.

### GetIsCashlessOk

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetIsCashlessOk() (*bool, bool)`

GetIsCashlessOk returns a tuple with the IsCashless field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCashless

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) SetIsCashless(v bool)`

SetIsCashless sets IsCashless field to given value.

### HasIsCashless

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) HasIsCashless() bool`

HasIsCashless returns a boolean if a field has been set.

### GetIsFakeDelivery

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetIsFakeDelivery() bool`

GetIsFakeDelivery returns the IsFakeDelivery field if non-nil, zero value otherwise.

### GetIsFakeDeliveryOk

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetIsFakeDeliveryOk() (*bool, bool)`

GetIsFakeDeliveryOk returns a tuple with the IsFakeDelivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFakeDelivery

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) SetIsFakeDelivery(v bool)`

SetIsFakeDelivery sets IsFakeDelivery field to given value.

### HasIsFakeDelivery

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) HasIsFakeDelivery() bool`

HasIsFakeDelivery returns a boolean if a field has been set.

### GetRecommendedCourierInfo

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetRecommendedCourierInfo() []GetSingleOrder200ResponseDataOrderInfoShippingInfoRecommendedCourierInfoInner`

GetRecommendedCourierInfo returns the RecommendedCourierInfo field if non-nil, zero value otherwise.

### GetRecommendedCourierInfoOk

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) GetRecommendedCourierInfoOk() (*[]GetSingleOrder200ResponseDataOrderInfoShippingInfoRecommendedCourierInfoInner, bool)`

GetRecommendedCourierInfoOk returns a tuple with the RecommendedCourierInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendedCourierInfo

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) SetRecommendedCourierInfo(v []GetSingleOrder200ResponseDataOrderInfoShippingInfoRecommendedCourierInfoInner)`

SetRecommendedCourierInfo sets RecommendedCourierInfo field to given value.

### HasRecommendedCourierInfo

`func (o *GetSingleOrder200ResponseDataOrderInfoShippingInfo) HasRecommendedCourierInfo() bool`

HasRecommendedCourierInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


