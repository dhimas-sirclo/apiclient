# RejectOrderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReasonCode** | Pointer to **int64** | Order reject reason: * empty/0 - Only for reject shipping case - reason required with max length 200 characters * 1 - Product(s) out of stock * 2 - Product variant unavailable * 3 - Wrong price or weight * 4 - Shop closed. close_end and closed_note required * 5 - Others * 7 - Courier problem - reason required with max length 200 characters * 8 - Buyer’s request - reason required with max length 200 characters  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**ShopCloseEndDate** | Pointer to **string** | format: 17/05/2017  | [optional] 
**ShopCloseNote** | Pointer to **string** |  | [optional] 
**EmptyProducts** | Pointer to **[]int64** |  | [optional] 

## Methods

### NewRejectOrderRequest

`func NewRejectOrderRequest() *RejectOrderRequest`

NewRejectOrderRequest instantiates a new RejectOrderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRejectOrderRequestWithDefaults

`func NewRejectOrderRequestWithDefaults() *RejectOrderRequest`

NewRejectOrderRequestWithDefaults instantiates a new RejectOrderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReasonCode

`func (o *RejectOrderRequest) GetReasonCode() int64`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *RejectOrderRequest) GetReasonCodeOk() (*int64, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *RejectOrderRequest) SetReasonCode(v int64)`

SetReasonCode sets ReasonCode field to given value.

### HasReasonCode

`func (o *RejectOrderRequest) HasReasonCode() bool`

HasReasonCode returns a boolean if a field has been set.

### GetReason

`func (o *RejectOrderRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *RejectOrderRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *RejectOrderRequest) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *RejectOrderRequest) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetShopCloseEndDate

`func (o *RejectOrderRequest) GetShopCloseEndDate() string`

GetShopCloseEndDate returns the ShopCloseEndDate field if non-nil, zero value otherwise.

### GetShopCloseEndDateOk

`func (o *RejectOrderRequest) GetShopCloseEndDateOk() (*string, bool)`

GetShopCloseEndDateOk returns a tuple with the ShopCloseEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopCloseEndDate

`func (o *RejectOrderRequest) SetShopCloseEndDate(v string)`

SetShopCloseEndDate sets ShopCloseEndDate field to given value.

### HasShopCloseEndDate

`func (o *RejectOrderRequest) HasShopCloseEndDate() bool`

HasShopCloseEndDate returns a boolean if a field has been set.

### GetShopCloseNote

`func (o *RejectOrderRequest) GetShopCloseNote() string`

GetShopCloseNote returns the ShopCloseNote field if non-nil, zero value otherwise.

### GetShopCloseNoteOk

`func (o *RejectOrderRequest) GetShopCloseNoteOk() (*string, bool)`

GetShopCloseNoteOk returns a tuple with the ShopCloseNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopCloseNote

`func (o *RejectOrderRequest) SetShopCloseNote(v string)`

SetShopCloseNote sets ShopCloseNote field to given value.

### HasShopCloseNote

`func (o *RejectOrderRequest) HasShopCloseNote() bool`

HasShopCloseNote returns a boolean if a field has been set.

### GetEmptyProducts

`func (o *RejectOrderRequest) GetEmptyProducts() []int64`

GetEmptyProducts returns the EmptyProducts field if non-nil, zero value otherwise.

### GetEmptyProductsOk

`func (o *RejectOrderRequest) GetEmptyProductsOk() (*[]int64, bool)`

GetEmptyProductsOk returns a tuple with the EmptyProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmptyProducts

`func (o *RejectOrderRequest) SetEmptyProducts(v []int64)`

SetEmptyProducts sets EmptyProducts field to given value.

### HasEmptyProducts

`func (o *RejectOrderRequest) HasEmptyProducts() bool`

HasEmptyProducts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


