# GetResolutionTicket200ResponseDataShopsInnerTicketInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**ProblemType** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**OpenDate** | Pointer to **string** | format: YYYY-MM-DD  | [optional] 
**SlaDate** | Pointer to **string** | format: YYYY-MM-DD  | [optional] 
**CloseDate** | Pointer to **string** | format: YYYY-MM-DD  | [optional] 
**InvoiceNumber** | Pointer to **string** |  | [optional] 
**Solution** | Pointer to **string** |  | [optional] 
**ComplaintProduct** | Pointer to [**[]GetResolutionTicket200ResponseDataShopsInnerTicketInnerComplaintProductInner**](GetResolutionTicket200ResponseDataShopsInnerTicketInnerComplaintProductInner.md) |  | [optional] 
**Fault** | Pointer to **string** |  | [optional] 
**ShippingAmt** | Pointer to **int64** |  | [optional] 
**TotalIssuedFunds** | Pointer to **int64** |  | [optional] 

## Methods

### NewGetResolutionTicket200ResponseDataShopsInnerTicketInner

`func NewGetResolutionTicket200ResponseDataShopsInnerTicketInner() *GetResolutionTicket200ResponseDataShopsInnerTicketInner`

NewGetResolutionTicket200ResponseDataShopsInnerTicketInner instantiates a new GetResolutionTicket200ResponseDataShopsInnerTicketInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetResolutionTicket200ResponseDataShopsInnerTicketInnerWithDefaults

`func NewGetResolutionTicket200ResponseDataShopsInnerTicketInnerWithDefaults() *GetResolutionTicket200ResponseDataShopsInnerTicketInner`

NewGetResolutionTicket200ResponseDataShopsInnerTicketInnerWithDefaults instantiates a new GetResolutionTicket200ResponseDataShopsInnerTicketInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetResolutionTicket200ResponseDataShopsInnerTicketInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetResolutionTicket200ResponseDataShopsInnerTicketInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetResolutionTicket200ResponseDataShopsInnerTicketInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetResolutionTicket200ResponseDataShopsInnerTicketInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProblemType

`func (o *GetResolutionTicket200ResponseDataShopsInnerTicketInner) GetProblemType() string`

GetProblemType returns the ProblemType field if non-nil, zero value otherwise.

### GetProblemTypeOk

`func (o *GetResolutionTicket200ResponseDataShopsInnerTicketInner) GetProblemTypeOk() (*string, bool)`

GetProblemTypeOk returns a tuple with the ProblemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProblemType

`func (o *GetResolutionTicket200ResponseDataShopsInnerTicketInner) SetProblemType(v string)`

SetProblemType sets ProblemType field to given value.

### HasProblemType

`func (o *GetResolutionTicket200ResponseDataShopsInnerTicketInner) HasProblemType() bool`

HasProblemType returns a boolean if a field has been set.

### GetStatus

`func (o *GetResolutionTicket200ResponseDataShopsInnerTicketInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetResolutionTicket200ResponseDataShopsInnerTicketInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetResolutionTicket200ResponseDataShopsInnerTicketInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetResolutionTicket200ResponseDataShopsInnerTicketInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetOpenDate

`func (o *GetResolutionTicket200ResponseDataShopsInnerTicketInner) GetOpenDate() string`

GetOpenDate returns the OpenDate field if non-nil, zero value otherwise.

### GetOpenDateOk

`func (o *GetResolutionTicket200ResponseDataShopsInnerTicketInner) GetOpenDateOk() (*string, bool)`

GetOpenDateOk returns a tuple with the OpenDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenDate

`func (o *GetResolutionTicket200ResponseDataShopsInnerTicketInner) SetOpenDate(v string)`

SetOpenDate sets OpenDate field to given value.

### HasOpenDate

`func (o *GetResolutionTicket200ResponseDataShopsInnerTicketInner) HasOpenDate() bool`

HasOpenDate returns a boolean if a field has been set.

### GetSlaDate

`func (o *GetResolutionTicket200ResponseDataShopsInnerTicketInner) GetSlaDate() string`

GetSlaDate returns the SlaDate field if non-nil, zero value otherwise.

### GetSlaDateOk

`func (o *GetResolutionTicket200ResponseDataShopsInnerTicketInner) GetSlaDateOk() (*string, bool)`

GetSlaDateOk returns a tuple with the SlaDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaDate

`func (o *GetResolutionTicket200ResponseDataShopsInnerTicketInner) SetSlaDate(v string)`

SetSlaDate sets SlaDate field to given value.

### HasSlaDate

`func (o *GetResolutionTicket200ResponseDataShopsInnerTicketInner) HasSlaDate() bool`

HasSlaDate returns a boolean if a field has been set.

### GetCloseDate

`func (o *GetResolutionTicket200ResponseDataShopsInnerTicketInner) GetCloseDate() string`

GetCloseDate returns the CloseDate field if non-nil, zero value otherwise.

### GetCloseDateOk

`func (o *GetResolutionTicket200ResponseDataShopsInnerTicketInner) GetCloseDateOk() (*string, bool)`

GetCloseDateOk returns a tuple with the CloseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseDate

`func (o *GetResolutionTicket200ResponseDataShopsInnerTicketInner) SetCloseDate(v string)`

SetCloseDate sets CloseDate field to given value.

### HasCloseDate

`func (o *GetResolutionTicket200ResponseDataShopsInnerTicketInner) HasCloseDate() bool`

HasCloseDate returns a boolean if a field has been set.

### GetInvoiceNumber

`func (o *GetResolutionTicket200ResponseDataShopsInnerTicketInner) GetInvoiceNumber() string`

GetInvoiceNumber returns the InvoiceNumber field if non-nil, zero value otherwise.

### GetInvoiceNumberOk

`func (o *GetResolutionTicket200ResponseDataShopsInnerTicketInner) GetInvoiceNumberOk() (*string, bool)`

GetInvoiceNumberOk returns a tuple with the InvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNumber

`func (o *GetResolutionTicket200ResponseDataShopsInnerTicketInner) SetInvoiceNumber(v string)`

SetInvoiceNumber sets InvoiceNumber field to given value.

### HasInvoiceNumber

`func (o *GetResolutionTicket200ResponseDataShopsInnerTicketInner) HasInvoiceNumber() bool`

HasInvoiceNumber returns a boolean if a field has been set.

### GetSolution

`func (o *GetResolutionTicket200ResponseDataShopsInnerTicketInner) GetSolution() string`

GetSolution returns the Solution field if non-nil, zero value otherwise.

### GetSolutionOk

`func (o *GetResolutionTicket200ResponseDataShopsInnerTicketInner) GetSolutionOk() (*string, bool)`

GetSolutionOk returns a tuple with the Solution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolution

`func (o *GetResolutionTicket200ResponseDataShopsInnerTicketInner) SetSolution(v string)`

SetSolution sets Solution field to given value.

### HasSolution

`func (o *GetResolutionTicket200ResponseDataShopsInnerTicketInner) HasSolution() bool`

HasSolution returns a boolean if a field has been set.

### GetComplaintProduct

`func (o *GetResolutionTicket200ResponseDataShopsInnerTicketInner) GetComplaintProduct() []GetResolutionTicket200ResponseDataShopsInnerTicketInnerComplaintProductInner`

GetComplaintProduct returns the ComplaintProduct field if non-nil, zero value otherwise.

### GetComplaintProductOk

`func (o *GetResolutionTicket200ResponseDataShopsInnerTicketInner) GetComplaintProductOk() (*[]GetResolutionTicket200ResponseDataShopsInnerTicketInnerComplaintProductInner, bool)`

GetComplaintProductOk returns a tuple with the ComplaintProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplaintProduct

`func (o *GetResolutionTicket200ResponseDataShopsInnerTicketInner) SetComplaintProduct(v []GetResolutionTicket200ResponseDataShopsInnerTicketInnerComplaintProductInner)`

SetComplaintProduct sets ComplaintProduct field to given value.

### HasComplaintProduct

`func (o *GetResolutionTicket200ResponseDataShopsInnerTicketInner) HasComplaintProduct() bool`

HasComplaintProduct returns a boolean if a field has been set.

### GetFault

`func (o *GetResolutionTicket200ResponseDataShopsInnerTicketInner) GetFault() string`

GetFault returns the Fault field if non-nil, zero value otherwise.

### GetFaultOk

`func (o *GetResolutionTicket200ResponseDataShopsInnerTicketInner) GetFaultOk() (*string, bool)`

GetFaultOk returns a tuple with the Fault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFault

`func (o *GetResolutionTicket200ResponseDataShopsInnerTicketInner) SetFault(v string)`

SetFault sets Fault field to given value.

### HasFault

`func (o *GetResolutionTicket200ResponseDataShopsInnerTicketInner) HasFault() bool`

HasFault returns a boolean if a field has been set.

### GetShippingAmt

`func (o *GetResolutionTicket200ResponseDataShopsInnerTicketInner) GetShippingAmt() int64`

GetShippingAmt returns the ShippingAmt field if non-nil, zero value otherwise.

### GetShippingAmtOk

`func (o *GetResolutionTicket200ResponseDataShopsInnerTicketInner) GetShippingAmtOk() (*int64, bool)`

GetShippingAmtOk returns a tuple with the ShippingAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAmt

`func (o *GetResolutionTicket200ResponseDataShopsInnerTicketInner) SetShippingAmt(v int64)`

SetShippingAmt sets ShippingAmt field to given value.

### HasShippingAmt

`func (o *GetResolutionTicket200ResponseDataShopsInnerTicketInner) HasShippingAmt() bool`

HasShippingAmt returns a boolean if a field has been set.

### GetTotalIssuedFunds

`func (o *GetResolutionTicket200ResponseDataShopsInnerTicketInner) GetTotalIssuedFunds() int64`

GetTotalIssuedFunds returns the TotalIssuedFunds field if non-nil, zero value otherwise.

### GetTotalIssuedFundsOk

`func (o *GetResolutionTicket200ResponseDataShopsInnerTicketInner) GetTotalIssuedFundsOk() (*int64, bool)`

GetTotalIssuedFundsOk returns a tuple with the TotalIssuedFunds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalIssuedFunds

`func (o *GetResolutionTicket200ResponseDataShopsInnerTicketInner) SetTotalIssuedFunds(v int64)`

SetTotalIssuedFunds sets TotalIssuedFunds field to given value.

### HasTotalIssuedFunds

`func (o *GetResolutionTicket200ResponseDataShopsInnerTicketInner) HasTotalIssuedFunds() bool`

HasTotalIssuedFunds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


