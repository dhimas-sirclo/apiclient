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

// checks if the GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment{}

// GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment struct for GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment
type GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment struct {
	Id *int64 `json:"id,omitempty"`
	OrderId *int64 `json:"order_id,omitempty"`
	// format: 2019-07-18T14:58:01.998063Z 
	PaymentDateTime *string `json:"payment_date_time,omitempty"`
	IsSameDay *bool `json:"is_same_day,omitempty"`
	// format: 2019-07-18T14:58:01.998063Z 
	AcceptDeadline *string `json:"accept_deadline,omitempty"`
	// format: 2019-07-18T14:58:01.998063Z 
	ConfirmShippingDeadline *string `json:"confirm_shipping_deadline,omitempty"`
	ItemDeliveredDeadline *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillmentItemDeliveredDeadline `json:"item_delivered_deadline,omitempty"`
	IsAccepted *bool `json:"is_accepted,omitempty"`
	IsConfirmShipping *bool `json:"is_confirm_shipping,omitempty"`
	IsItemDelivered *bool `json:"is_item_delivered,omitempty"`
	FulfillmentStatus *int64 `json:"fulfillment_status,omitempty"`
}

// NewGetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment instantiates a new GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment() *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment {
	this := GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment{}
	return &this
}

// NewGetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillmentWithDefaults instantiates a new GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillmentWithDefaults() *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment {
	this := GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) SetId(v int64) {
	o.Id = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) GetOrderId() int64 {
	if o == nil || IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) GetOrderIdOk() (*int64, bool) {
	if o == nil || IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) HasOrderId() bool {
	if o != nil && !IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetPaymentDateTime returns the PaymentDateTime field value if set, zero value otherwise.
func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) GetPaymentDateTime() string {
	if o == nil || IsNil(o.PaymentDateTime) {
		var ret string
		return ret
	}
	return *o.PaymentDateTime
}

// GetPaymentDateTimeOk returns a tuple with the PaymentDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) GetPaymentDateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentDateTime) {
		return nil, false
	}
	return o.PaymentDateTime, true
}

// HasPaymentDateTime returns a boolean if a field has been set.
func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) HasPaymentDateTime() bool {
	if o != nil && !IsNil(o.PaymentDateTime) {
		return true
	}

	return false
}

// SetPaymentDateTime gets a reference to the given string and assigns it to the PaymentDateTime field.
func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) SetPaymentDateTime(v string) {
	o.PaymentDateTime = &v
}

// GetIsSameDay returns the IsSameDay field value if set, zero value otherwise.
func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) GetIsSameDay() bool {
	if o == nil || IsNil(o.IsSameDay) {
		var ret bool
		return ret
	}
	return *o.IsSameDay
}

// GetIsSameDayOk returns a tuple with the IsSameDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) GetIsSameDayOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSameDay) {
		return nil, false
	}
	return o.IsSameDay, true
}

// HasIsSameDay returns a boolean if a field has been set.
func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) HasIsSameDay() bool {
	if o != nil && !IsNil(o.IsSameDay) {
		return true
	}

	return false
}

// SetIsSameDay gets a reference to the given bool and assigns it to the IsSameDay field.
func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) SetIsSameDay(v bool) {
	o.IsSameDay = &v
}

// GetAcceptDeadline returns the AcceptDeadline field value if set, zero value otherwise.
func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) GetAcceptDeadline() string {
	if o == nil || IsNil(o.AcceptDeadline) {
		var ret string
		return ret
	}
	return *o.AcceptDeadline
}

// GetAcceptDeadlineOk returns a tuple with the AcceptDeadline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) GetAcceptDeadlineOk() (*string, bool) {
	if o == nil || IsNil(o.AcceptDeadline) {
		return nil, false
	}
	return o.AcceptDeadline, true
}

// HasAcceptDeadline returns a boolean if a field has been set.
func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) HasAcceptDeadline() bool {
	if o != nil && !IsNil(o.AcceptDeadline) {
		return true
	}

	return false
}

// SetAcceptDeadline gets a reference to the given string and assigns it to the AcceptDeadline field.
func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) SetAcceptDeadline(v string) {
	o.AcceptDeadline = &v
}

// GetConfirmShippingDeadline returns the ConfirmShippingDeadline field value if set, zero value otherwise.
func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) GetConfirmShippingDeadline() string {
	if o == nil || IsNil(o.ConfirmShippingDeadline) {
		var ret string
		return ret
	}
	return *o.ConfirmShippingDeadline
}

// GetConfirmShippingDeadlineOk returns a tuple with the ConfirmShippingDeadline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) GetConfirmShippingDeadlineOk() (*string, bool) {
	if o == nil || IsNil(o.ConfirmShippingDeadline) {
		return nil, false
	}
	return o.ConfirmShippingDeadline, true
}

// HasConfirmShippingDeadline returns a boolean if a field has been set.
func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) HasConfirmShippingDeadline() bool {
	if o != nil && !IsNil(o.ConfirmShippingDeadline) {
		return true
	}

	return false
}

// SetConfirmShippingDeadline gets a reference to the given string and assigns it to the ConfirmShippingDeadline field.
func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) SetConfirmShippingDeadline(v string) {
	o.ConfirmShippingDeadline = &v
}

// GetItemDeliveredDeadline returns the ItemDeliveredDeadline field value if set, zero value otherwise.
func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) GetItemDeliveredDeadline() GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillmentItemDeliveredDeadline {
	if o == nil || IsNil(o.ItemDeliveredDeadline) {
		var ret GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillmentItemDeliveredDeadline
		return ret
	}
	return *o.ItemDeliveredDeadline
}

// GetItemDeliveredDeadlineOk returns a tuple with the ItemDeliveredDeadline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) GetItemDeliveredDeadlineOk() (*GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillmentItemDeliveredDeadline, bool) {
	if o == nil || IsNil(o.ItemDeliveredDeadline) {
		return nil, false
	}
	return o.ItemDeliveredDeadline, true
}

// HasItemDeliveredDeadline returns a boolean if a field has been set.
func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) HasItemDeliveredDeadline() bool {
	if o != nil && !IsNil(o.ItemDeliveredDeadline) {
		return true
	}

	return false
}

// SetItemDeliveredDeadline gets a reference to the given GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillmentItemDeliveredDeadline and assigns it to the ItemDeliveredDeadline field.
func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) SetItemDeliveredDeadline(v GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillmentItemDeliveredDeadline) {
	o.ItemDeliveredDeadline = &v
}

// GetIsAccepted returns the IsAccepted field value if set, zero value otherwise.
func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) GetIsAccepted() bool {
	if o == nil || IsNil(o.IsAccepted) {
		var ret bool
		return ret
	}
	return *o.IsAccepted
}

// GetIsAcceptedOk returns a tuple with the IsAccepted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) GetIsAcceptedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAccepted) {
		return nil, false
	}
	return o.IsAccepted, true
}

// HasIsAccepted returns a boolean if a field has been set.
func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) HasIsAccepted() bool {
	if o != nil && !IsNil(o.IsAccepted) {
		return true
	}

	return false
}

// SetIsAccepted gets a reference to the given bool and assigns it to the IsAccepted field.
func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) SetIsAccepted(v bool) {
	o.IsAccepted = &v
}

// GetIsConfirmShipping returns the IsConfirmShipping field value if set, zero value otherwise.
func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) GetIsConfirmShipping() bool {
	if o == nil || IsNil(o.IsConfirmShipping) {
		var ret bool
		return ret
	}
	return *o.IsConfirmShipping
}

// GetIsConfirmShippingOk returns a tuple with the IsConfirmShipping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) GetIsConfirmShippingOk() (*bool, bool) {
	if o == nil || IsNil(o.IsConfirmShipping) {
		return nil, false
	}
	return o.IsConfirmShipping, true
}

// HasIsConfirmShipping returns a boolean if a field has been set.
func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) HasIsConfirmShipping() bool {
	if o != nil && !IsNil(o.IsConfirmShipping) {
		return true
	}

	return false
}

// SetIsConfirmShipping gets a reference to the given bool and assigns it to the IsConfirmShipping field.
func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) SetIsConfirmShipping(v bool) {
	o.IsConfirmShipping = &v
}

// GetIsItemDelivered returns the IsItemDelivered field value if set, zero value otherwise.
func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) GetIsItemDelivered() bool {
	if o == nil || IsNil(o.IsItemDelivered) {
		var ret bool
		return ret
	}
	return *o.IsItemDelivered
}

// GetIsItemDeliveredOk returns a tuple with the IsItemDelivered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) GetIsItemDeliveredOk() (*bool, bool) {
	if o == nil || IsNil(o.IsItemDelivered) {
		return nil, false
	}
	return o.IsItemDelivered, true
}

// HasIsItemDelivered returns a boolean if a field has been set.
func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) HasIsItemDelivered() bool {
	if o != nil && !IsNil(o.IsItemDelivered) {
		return true
	}

	return false
}

// SetIsItemDelivered gets a reference to the given bool and assigns it to the IsItemDelivered field.
func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) SetIsItemDelivered(v bool) {
	o.IsItemDelivered = &v
}

// GetFulfillmentStatus returns the FulfillmentStatus field value if set, zero value otherwise.
func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) GetFulfillmentStatus() int64 {
	if o == nil || IsNil(o.FulfillmentStatus) {
		var ret int64
		return ret
	}
	return *o.FulfillmentStatus
}

// GetFulfillmentStatusOk returns a tuple with the FulfillmentStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) GetFulfillmentStatusOk() (*int64, bool) {
	if o == nil || IsNil(o.FulfillmentStatus) {
		return nil, false
	}
	return o.FulfillmentStatus, true
}

// HasFulfillmentStatus returns a boolean if a field has been set.
func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) HasFulfillmentStatus() bool {
	if o != nil && !IsNil(o.FulfillmentStatus) {
		return true
	}

	return false
}

// SetFulfillmentStatus gets a reference to the given int64 and assigns it to the FulfillmentStatus field.
func (o *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) SetFulfillmentStatus(v int64) {
	o.FulfillmentStatus = &v
}

func (o GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) ToMap() (map[string]interface{}, error) {
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

type NullableGetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment struct {
	value *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment
	isSet bool
}

func (v NullableGetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) Get() *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment {
	return v.value
}

func (v *NullableGetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) Set(val *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment(val *GetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) *NullableGetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment {
	return &NullableGetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment{value: val, isSet: true}
}

func (v NullableGetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrderCobCod200ResponseDataOrderDataInnerOrderShipmentFulfillment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


