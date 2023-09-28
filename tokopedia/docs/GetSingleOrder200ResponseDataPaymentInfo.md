# GetSingleOrder200ResponseDataPaymentInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentId** | Pointer to **int64** | Payment Unique Identifier | [optional] 
**PaymentRefNum** | Pointer to **string** | Payment Ref Number | [optional] 
**PaymentDate** | Pointer to **string** | Payment Date Timestamp (format: 2017-07-20T17:50:06Z)  | [optional] 
**PaymentMethod** | Pointer to **int64** | Payment Method | [optional] 
**PaymentStatus** | Pointer to **string** | Payment Status | [optional] 
**PaymentStatusId** | Pointer to **int64** | Payment Status | [optional] 
**CreateTime** | Pointer to **string** | Create Time Timestamp (format: 2017-07-20T17:50:06Z)  | [optional] 
**PgId** | Pointer to **int64** | Payment Gateway Identifier | [optional] 
**GatewayName** | Pointer to **string** | Gateway Name | [optional] 
**DiscountAmount** | Pointer to **int64** | Discount Amount | [optional] 
**VoucherCode** | Pointer to **string** | Voucher Code | [optional] 
**VoucherId** | Pointer to **int64** | Voucher Unique Identifier | [optional] 

## Methods

### NewGetSingleOrder200ResponseDataPaymentInfo

`func NewGetSingleOrder200ResponseDataPaymentInfo() *GetSingleOrder200ResponseDataPaymentInfo`

NewGetSingleOrder200ResponseDataPaymentInfo instantiates a new GetSingleOrder200ResponseDataPaymentInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSingleOrder200ResponseDataPaymentInfoWithDefaults

`func NewGetSingleOrder200ResponseDataPaymentInfoWithDefaults() *GetSingleOrder200ResponseDataPaymentInfo`

NewGetSingleOrder200ResponseDataPaymentInfoWithDefaults instantiates a new GetSingleOrder200ResponseDataPaymentInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentId

`func (o *GetSingleOrder200ResponseDataPaymentInfo) GetPaymentId() int64`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *GetSingleOrder200ResponseDataPaymentInfo) GetPaymentIdOk() (*int64, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *GetSingleOrder200ResponseDataPaymentInfo) SetPaymentId(v int64)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *GetSingleOrder200ResponseDataPaymentInfo) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### GetPaymentRefNum

`func (o *GetSingleOrder200ResponseDataPaymentInfo) GetPaymentRefNum() string`

GetPaymentRefNum returns the PaymentRefNum field if non-nil, zero value otherwise.

### GetPaymentRefNumOk

`func (o *GetSingleOrder200ResponseDataPaymentInfo) GetPaymentRefNumOk() (*string, bool)`

GetPaymentRefNumOk returns a tuple with the PaymentRefNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentRefNum

`func (o *GetSingleOrder200ResponseDataPaymentInfo) SetPaymentRefNum(v string)`

SetPaymentRefNum sets PaymentRefNum field to given value.

### HasPaymentRefNum

`func (o *GetSingleOrder200ResponseDataPaymentInfo) HasPaymentRefNum() bool`

HasPaymentRefNum returns a boolean if a field has been set.

### GetPaymentDate

`func (o *GetSingleOrder200ResponseDataPaymentInfo) GetPaymentDate() string`

GetPaymentDate returns the PaymentDate field if non-nil, zero value otherwise.

### GetPaymentDateOk

`func (o *GetSingleOrder200ResponseDataPaymentInfo) GetPaymentDateOk() (*string, bool)`

GetPaymentDateOk returns a tuple with the PaymentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDate

`func (o *GetSingleOrder200ResponseDataPaymentInfo) SetPaymentDate(v string)`

SetPaymentDate sets PaymentDate field to given value.

### HasPaymentDate

`func (o *GetSingleOrder200ResponseDataPaymentInfo) HasPaymentDate() bool`

HasPaymentDate returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *GetSingleOrder200ResponseDataPaymentInfo) GetPaymentMethod() int64`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *GetSingleOrder200ResponseDataPaymentInfo) GetPaymentMethodOk() (*int64, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *GetSingleOrder200ResponseDataPaymentInfo) SetPaymentMethod(v int64)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *GetSingleOrder200ResponseDataPaymentInfo) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetPaymentStatus

`func (o *GetSingleOrder200ResponseDataPaymentInfo) GetPaymentStatus() string`

GetPaymentStatus returns the PaymentStatus field if non-nil, zero value otherwise.

### GetPaymentStatusOk

`func (o *GetSingleOrder200ResponseDataPaymentInfo) GetPaymentStatusOk() (*string, bool)`

GetPaymentStatusOk returns a tuple with the PaymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentStatus

`func (o *GetSingleOrder200ResponseDataPaymentInfo) SetPaymentStatus(v string)`

SetPaymentStatus sets PaymentStatus field to given value.

### HasPaymentStatus

`func (o *GetSingleOrder200ResponseDataPaymentInfo) HasPaymentStatus() bool`

HasPaymentStatus returns a boolean if a field has been set.

### GetPaymentStatusId

`func (o *GetSingleOrder200ResponseDataPaymentInfo) GetPaymentStatusId() int64`

GetPaymentStatusId returns the PaymentStatusId field if non-nil, zero value otherwise.

### GetPaymentStatusIdOk

`func (o *GetSingleOrder200ResponseDataPaymentInfo) GetPaymentStatusIdOk() (*int64, bool)`

GetPaymentStatusIdOk returns a tuple with the PaymentStatusId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentStatusId

`func (o *GetSingleOrder200ResponseDataPaymentInfo) SetPaymentStatusId(v int64)`

SetPaymentStatusId sets PaymentStatusId field to given value.

### HasPaymentStatusId

`func (o *GetSingleOrder200ResponseDataPaymentInfo) HasPaymentStatusId() bool`

HasPaymentStatusId returns a boolean if a field has been set.

### GetCreateTime

`func (o *GetSingleOrder200ResponseDataPaymentInfo) GetCreateTime() string`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *GetSingleOrder200ResponseDataPaymentInfo) GetCreateTimeOk() (*string, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *GetSingleOrder200ResponseDataPaymentInfo) SetCreateTime(v string)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *GetSingleOrder200ResponseDataPaymentInfo) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetPgId

`func (o *GetSingleOrder200ResponseDataPaymentInfo) GetPgId() int64`

GetPgId returns the PgId field if non-nil, zero value otherwise.

### GetPgIdOk

`func (o *GetSingleOrder200ResponseDataPaymentInfo) GetPgIdOk() (*int64, bool)`

GetPgIdOk returns a tuple with the PgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgId

`func (o *GetSingleOrder200ResponseDataPaymentInfo) SetPgId(v int64)`

SetPgId sets PgId field to given value.

### HasPgId

`func (o *GetSingleOrder200ResponseDataPaymentInfo) HasPgId() bool`

HasPgId returns a boolean if a field has been set.

### GetGatewayName

`func (o *GetSingleOrder200ResponseDataPaymentInfo) GetGatewayName() string`

GetGatewayName returns the GatewayName field if non-nil, zero value otherwise.

### GetGatewayNameOk

`func (o *GetSingleOrder200ResponseDataPaymentInfo) GetGatewayNameOk() (*string, bool)`

GetGatewayNameOk returns a tuple with the GatewayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayName

`func (o *GetSingleOrder200ResponseDataPaymentInfo) SetGatewayName(v string)`

SetGatewayName sets GatewayName field to given value.

### HasGatewayName

`func (o *GetSingleOrder200ResponseDataPaymentInfo) HasGatewayName() bool`

HasGatewayName returns a boolean if a field has been set.

### GetDiscountAmount

`func (o *GetSingleOrder200ResponseDataPaymentInfo) GetDiscountAmount() int64`

GetDiscountAmount returns the DiscountAmount field if non-nil, zero value otherwise.

### GetDiscountAmountOk

`func (o *GetSingleOrder200ResponseDataPaymentInfo) GetDiscountAmountOk() (*int64, bool)`

GetDiscountAmountOk returns a tuple with the DiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmount

`func (o *GetSingleOrder200ResponseDataPaymentInfo) SetDiscountAmount(v int64)`

SetDiscountAmount sets DiscountAmount field to given value.

### HasDiscountAmount

`func (o *GetSingleOrder200ResponseDataPaymentInfo) HasDiscountAmount() bool`

HasDiscountAmount returns a boolean if a field has been set.

### GetVoucherCode

`func (o *GetSingleOrder200ResponseDataPaymentInfo) GetVoucherCode() string`

GetVoucherCode returns the VoucherCode field if non-nil, zero value otherwise.

### GetVoucherCodeOk

`func (o *GetSingleOrder200ResponseDataPaymentInfo) GetVoucherCodeOk() (*string, bool)`

GetVoucherCodeOk returns a tuple with the VoucherCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoucherCode

`func (o *GetSingleOrder200ResponseDataPaymentInfo) SetVoucherCode(v string)`

SetVoucherCode sets VoucherCode field to given value.

### HasVoucherCode

`func (o *GetSingleOrder200ResponseDataPaymentInfo) HasVoucherCode() bool`

HasVoucherCode returns a boolean if a field has been set.

### GetVoucherId

`func (o *GetSingleOrder200ResponseDataPaymentInfo) GetVoucherId() int64`

GetVoucherId returns the VoucherId field if non-nil, zero value otherwise.

### GetVoucherIdOk

`func (o *GetSingleOrder200ResponseDataPaymentInfo) GetVoucherIdOk() (*int64, bool)`

GetVoucherIdOk returns a tuple with the VoucherId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoucherId

`func (o *GetSingleOrder200ResponseDataPaymentInfo) SetVoucherId(v int64)`

SetVoucherId sets VoucherId field to given value.

### HasVoucherId

`func (o *GetSingleOrder200ResponseDataPaymentInfo) HasVoucherId() bool`

HasVoucherId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


