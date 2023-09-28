/*
Tokopedia API

Tokopedia API

API version: 1.0
Contact: dev@sirclo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tokopedia

import (
	"encoding/json"
)

// checks if the GetSingleOrder200ResponseDataShipmentFulfillment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSingleOrder200ResponseDataShipmentFulfillment{}

// GetSingleOrder200ResponseDataShipmentFulfillment struct for GetSingleOrder200ResponseDataShipmentFulfillment
type GetSingleOrder200ResponseDataShipmentFulfillment struct {
	// Shipment Fulfillment Unique Identifier
	Id *int64 `json:"id,omitempty"`
	// Order Unique Identifier
	OrderId *int64 `json:"order_id,omitempty"`
	// Payment Date Time (format: 0001-01-01T00:00:00Z) 
	PaymentDateTime *string `json:"payment_date_time,omitempty"`
	// Is same day delivery?
	IsSameDay *bool `json:"is_same_day,omitempty"`
	// Accept Deadline (format: 0001-01-01T00:00:00Z) 
	AcceptDeadline *string `json:"accept_deadline,omitempty"`
	// Accept Deadline (format: 0001-01-01T00:00:00Z) 
	ConfirmShippingDeadline *string `json:"confirm_shipping_deadline,omitempty"`
	ItemDeliveredDeadline *GetSingleOrder200ResponseDataShipmentFulfillmentItemDeliveredDeadline `json:"item_delivered_deadline,omitempty"`
	// Is order accepted?
	IsAccepted *bool `json:"is_accepted,omitempty"`
	// Is shipping confirmed?
	IsConfirmShipping *bool `json:"is_confirm_shipping,omitempty"`
	// Is item delivered?
	IsItemDelivered *bool `json:"is_item_delivered,omitempty"`
	// Fullfilment Status
	FulfillmentStatus *int64 `json:"fulfillment_status,omitempty"`
}

// NewGetSingleOrder200ResponseDataShipmentFulfillment instantiates a new GetSingleOrder200ResponseDataShipmentFulfillment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSingleOrder200ResponseDataShipmentFulfillment() *GetSingleOrder200ResponseDataShipmentFulfillment {
	this := GetSingleOrder200ResponseDataShipmentFulfillment{}
	return &this
}

// NewGetSingleOrder200ResponseDataShipmentFulfillmentWithDefaults instantiates a new GetSingleOrder200ResponseDataShipmentFulfillment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSingleOrder200ResponseDataShipmentFulfillmentWithDefaults() *GetSingleOrder200ResponseDataShipmentFulfillment {
	this := GetSingleOrder200ResponseDataShipmentFulfillment{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataShipmentFulfillment) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataShipmentFulfillment) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataShipmentFulfillment) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *GetSingleOrder200ResponseDataShipmentFulfillment) SetId(v int64) {
	o.Id = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataShipmentFulfillment) GetOrderId() int64 {
	if o == nil || IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataShipmentFulfillment) GetOrderIdOk() (*int64, bool) {
	if o == nil || IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataShipmentFulfillment) HasOrderId() bool {
	if o != nil && !IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *GetSingleOrder200ResponseDataShipmentFulfillment) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetPaymentDateTime returns the PaymentDateTime field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataShipmentFulfillment) GetPaymentDateTime() string {
	if o == nil || IsNil(o.PaymentDateTime) {
		var ret string
		return ret
	}
	return *o.PaymentDateTime
}

// GetPaymentDateTimeOk returns a tuple with the PaymentDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataShipmentFulfillment) GetPaymentDateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentDateTime) {
		return nil, false
	}
	return o.PaymentDateTime, true
}

// HasPaymentDateTime returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataShipmentFulfillment) HasPaymentDateTime() bool {
	if o != nil && !IsNil(o.PaymentDateTime) {
		return true
	}

	return false
}

// SetPaymentDateTime gets a reference to the given string and assigns it to the PaymentDateTime field.
func (o *GetSingleOrder200ResponseDataShipmentFulfillment) SetPaymentDateTime(v string) {
	o.PaymentDateTime = &v
}

// GetIsSameDay returns the IsSameDay field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataShipmentFulfillment) GetIsSameDay() bool {
	if o == nil || IsNil(o.IsSameDay) {
		var ret bool
		return ret
	}
	return *o.IsSameDay
}

// GetIsSameDayOk returns a tuple with the IsSameDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataShipmentFulfillment) GetIsSameDayOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSameDay) {
		return nil, false
	}
	return o.IsSameDay, true
}

// HasIsSameDay returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataShipmentFulfillment) HasIsSameDay() bool {
	if o != nil && !IsNil(o.IsSameDay) {
		return true
	}

	return false
}

// SetIsSameDay gets a reference to the given bool and assigns it to the IsSameDay field.
func (o *GetSingleOrder200ResponseDataShipmentFulfillment) SetIsSameDay(v bool) {
	o.IsSameDay = &v
}

// GetAcceptDeadline returns the AcceptDeadline field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataShipmentFulfillment) GetAcceptDeadline() string {
	if o == nil || IsNil(o.AcceptDeadline) {
		var ret string
		return ret
	}
	return *o.AcceptDeadline
}

// GetAcceptDeadlineOk returns a tuple with the AcceptDeadline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataShipmentFulfillment) GetAcceptDeadlineOk() (*string, bool) {
	if o == nil || IsNil(o.AcceptDeadline) {
		return nil, false
	}
	return o.AcceptDeadline, true
}

// HasAcceptDeadline returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataShipmentFulfillment) HasAcceptDeadline() bool {
	if o != nil && !IsNil(o.AcceptDeadline) {
		return true
	}

	return false
}

// SetAcceptDeadline gets a reference to the given string and assigns it to the AcceptDeadline field.
func (o *GetSingleOrder200ResponseDataShipmentFulfillment) SetAcceptDeadline(v string) {
	o.AcceptDeadline = &v
}

// GetConfirmShippingDeadline returns the ConfirmShippingDeadline field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataShipmentFulfillment) GetConfirmShippingDeadline() string {
	if o == nil || IsNil(o.ConfirmShippingDeadline) {
		var ret string
		return ret
	}
	return *o.ConfirmShippingDeadline
}

// GetConfirmShippingDeadlineOk returns a tuple with the ConfirmShippingDeadline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataShipmentFulfillment) GetConfirmShippingDeadlineOk() (*string, bool) {
	if o == nil || IsNil(o.ConfirmShippingDeadline) {
		return nil, false
	}
	return o.ConfirmShippingDeadline, true
}

// HasConfirmShippingDeadline returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataShipmentFulfillment) HasConfirmShippingDeadline() bool {
	if o != nil && !IsNil(o.ConfirmShippingDeadline) {
		return true
	}

	return false
}

// SetConfirmShippingDeadline gets a reference to the given string and assigns it to the ConfirmShippingDeadline field.
func (o *GetSingleOrder200ResponseDataShipmentFulfillment) SetConfirmShippingDeadline(v string) {
	o.ConfirmShippingDeadline = &v
}

// GetItemDeliveredDeadline returns the ItemDeliveredDeadline field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataShipmentFulfillment) GetItemDeliveredDeadline() GetSingleOrder200ResponseDataShipmentFulfillmentItemDeliveredDeadline {
	if o == nil || IsNil(o.ItemDeliveredDeadline) {
		var ret GetSingleOrder200ResponseDataShipmentFulfillmentItemDeliveredDeadline
		return ret
	}
	return *o.ItemDeliveredDeadline
}

// GetItemDeliveredDeadlineOk returns a tuple with the ItemDeliveredDeadline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataShipmentFulfillment) GetItemDeliveredDeadlineOk() (*GetSingleOrder200ResponseDataShipmentFulfillmentItemDeliveredDeadline, bool) {
	if o == nil || IsNil(o.ItemDeliveredDeadline) {
		return nil, false
	}
	return o.ItemDeliveredDeadline, true
}

// HasItemDeliveredDeadline returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataShipmentFulfillment) HasItemDeliveredDeadline() bool {
	if o != nil && !IsNil(o.ItemDeliveredDeadline) {
		return true
	}

	return false
}

// SetItemDeliveredDeadline gets a reference to the given GetSingleOrder200ResponseDataShipmentFulfillmentItemDeliveredDeadline and assigns it to the ItemDeliveredDeadline field.
func (o *GetSingleOrder200ResponseDataShipmentFulfillment) SetItemDeliveredDeadline(v GetSingleOrder200ResponseDataShipmentFulfillmentItemDeliveredDeadline) {
	o.ItemDeliveredDeadline = &v
}

// GetIsAccepted returns the IsAccepted field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataShipmentFulfillment) GetIsAccepted() bool {
	if o == nil || IsNil(o.IsAccepted) {
		var ret bool
		return ret
	}
	return *o.IsAccepted
}

// GetIsAcceptedOk returns a tuple with the IsAccepted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataShipmentFulfillment) GetIsAcceptedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAccepted) {
		return nil, false
	}
	return o.IsAccepted, true
}

// HasIsAccepted returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataShipmentFulfillment) HasIsAccepted() bool {
	if o != nil && !IsNil(o.IsAccepted) {
		return true
	}

	return false
}

// SetIsAccepted gets a reference to the given bool and assigns it to the IsAccepted field.
func (o *GetSingleOrder200ResponseDataShipmentFulfillment) SetIsAccepted(v bool) {
	o.IsAccepted = &v
}

// GetIsConfirmShipping returns the IsConfirmShipping field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataShipmentFulfillment) GetIsConfirmShipping() bool {
	if o == nil || IsNil(o.IsConfirmShipping) {
		var ret bool
		return ret
	}
	return *o.IsConfirmShipping
}

// GetIsConfirmShippingOk returns a tuple with the IsConfirmShipping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataShipmentFulfillment) GetIsConfirmShippingOk() (*bool, bool) {
	if o == nil || IsNil(o.IsConfirmShipping) {
		return nil, false
	}
	return o.IsConfirmShipping, true
}

// HasIsConfirmShipping returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataShipmentFulfillment) HasIsConfirmShipping() bool {
	if o != nil && !IsNil(o.IsConfirmShipping) {
		return true
	}

	return false
}

// SetIsConfirmShipping gets a reference to the given bool and assigns it to the IsConfirmShipping field.
func (o *GetSingleOrder200ResponseDataShipmentFulfillment) SetIsConfirmShipping(v bool) {
	o.IsConfirmShipping = &v
}

// GetIsItemDelivered returns the IsItemDelivered field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataShipmentFulfillment) GetIsItemDelivered() bool {
	if o == nil || IsNil(o.IsItemDelivered) {
		var ret bool
		return ret
	}
	return *o.IsItemDelivered
}

// GetIsItemDeliveredOk returns a tuple with the IsItemDelivered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataShipmentFulfillment) GetIsItemDeliveredOk() (*bool, bool) {
	if o == nil || IsNil(o.IsItemDelivered) {
		return nil, false
	}
	return o.IsItemDelivered, true
}

// HasIsItemDelivered returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataShipmentFulfillment) HasIsItemDelivered() bool {
	if o != nil && !IsNil(o.IsItemDelivered) {
		return true
	}

	return false
}

// SetIsItemDelivered gets a reference to the given bool and assigns it to the IsItemDelivered field.
func (o *GetSingleOrder200ResponseDataShipmentFulfillment) SetIsItemDelivered(v bool) {
	o.IsItemDelivered = &v
}

// GetFulfillmentStatus returns the FulfillmentStatus field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataShipmentFulfillment) GetFulfillmentStatus() int64 {
	if o == nil || IsNil(o.FulfillmentStatus) {
		var ret int64
		return ret
	}
	return *o.FulfillmentStatus
}

// GetFulfillmentStatusOk returns a tuple with the FulfillmentStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataShipmentFulfillment) GetFulfillmentStatusOk() (*int64, bool) {
	if o == nil || IsNil(o.FulfillmentStatus) {
		return nil, false
	}
	return o.FulfillmentStatus, true
}

// HasFulfillmentStatus returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataShipmentFulfillment) HasFulfillmentStatus() bool {
	if o != nil && !IsNil(o.FulfillmentStatus) {
		return true
	}

	return false
}

// SetFulfillmentStatus gets a reference to the given int64 and assigns it to the FulfillmentStatus field.
func (o *GetSingleOrder200ResponseDataShipmentFulfillment) SetFulfillmentStatus(v int64) {
	o.FulfillmentStatus = &v
}

func (o GetSingleOrder200ResponseDataShipmentFulfillment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSingleOrder200ResponseDataShipmentFulfillment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.OrderId) {
		toSerialize["order_id"] = o.OrderId
	}
	if !IsNil(o.PaymentDateTime) {
		toSerialize["payment_date_time"] = o.PaymentDateTime
	}
	if !IsNil(o.IsSameDay) {
		toSerialize["is_same_day"] = o.IsSameDay
	}
	if !IsNil(o.AcceptDeadline) {
		toSerialize["accept_deadline"] = o.AcceptDeadline
	}
	if !IsNil(o.ConfirmShippingDeadline) {
		toSerialize["confirm_shipping_deadline"] = o.ConfirmShippingDeadline
	}
	if !IsNil(o.ItemDeliveredDeadline) {
		toSerialize["item_delivered_deadline"] = o.ItemDeliveredDeadline
	}
	if !IsNil(o.IsAccepted) {
		toSerialize["is_accepted"] = o.IsAccepted
	}
	if !IsNil(o.IsConfirmShipping) {
		toSerialize["is_confirm_shipping"] = o.IsConfirmShipping
	}
	if !IsNil(o.IsItemDelivered) {
		toSerialize["is_item_delivered"] = o.IsItemDelivered
	}
	if !IsNil(o.FulfillmentStatus) {
		toSerialize["fulfillment_status"] = o.FulfillmentStatus
	}
	return toSerialize, nil
}

type NullableGetSingleOrder200ResponseDataShipmentFulfillment struct {
	value *GetSingleOrder200ResponseDataShipmentFulfillment
	isSet bool
}

func (v NullableGetSingleOrder200ResponseDataShipmentFulfillment) Get() *GetSingleOrder200ResponseDataShipmentFulfillment {
	return v.value
}

func (v *NullableGetSingleOrder200ResponseDataShipmentFulfillment) Set(val *GetSingleOrder200ResponseDataShipmentFulfillment) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSingleOrder200ResponseDataShipmentFulfillment) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSingleOrder200ResponseDataShipmentFulfillment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSingleOrder200ResponseDataShipmentFulfillment(val *GetSingleOrder200ResponseDataShipmentFulfillment) *NullableGetSingleOrder200ResponseDataShipmentFulfillment {
	return &NullableGetSingleOrder200ResponseDataShipmentFulfillment{value: val, isSet: true}
}

func (v NullableGetSingleOrder200ResponseDataShipmentFulfillment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSingleOrder200ResponseDataShipmentFulfillment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


