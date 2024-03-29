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

// checks if the GetBundleInfo200ResponseDataBundleInfoInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetBundleInfo200ResponseDataBundleInfoInner{}

// GetBundleInfo200ResponseDataBundleInfoInner struct for GetBundleInfo200ResponseDataBundleInfoInner
type GetBundleInfo200ResponseDataBundleInfoInner struct {
	BundleId *int64 `json:"bundle_id,omitempty"`
	GroupId *int64 `json:"group_id,omitempty"`
	Name *string `json:"Name,omitempty"`
	Type *int64 `json:"Type,omitempty"`
	Status *int64 `json:"status,omitempty"`
	ShopId *int64 `json:"shop_id,omitempty"`
	StartTimeUnix *int64 `json:"start_time_unix,omitempty"`
	StopTimeUnix *int64 `json:"stop_time_unix,omitempty"`
	BundleItem []GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInner `json:"bundle_item,omitempty"`
	WarehouseId *int64 `json:"warehouse_id,omitempty"`
	Quota *int64 `json:"quota,omitempty"`
	OriginalQuota *int64 `json:"original_quota,omitempty"`
	MaxOrder *int64 `json:"max_order,omitempty"`
	Preorder *GetBundleInfo200ResponseDataBundleInfoInnerPreorder `json:"preorder,omitempty"`
}

// NewGetBundleInfo200ResponseDataBundleInfoInner instantiates a new GetBundleInfo200ResponseDataBundleInfoInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBundleInfo200ResponseDataBundleInfoInner() *GetBundleInfo200ResponseDataBundleInfoInner {
	this := GetBundleInfo200ResponseDataBundleInfoInner{}
	return &this
}

// NewGetBundleInfo200ResponseDataBundleInfoInnerWithDefaults instantiates a new GetBundleInfo200ResponseDataBundleInfoInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBundleInfo200ResponseDataBundleInfoInnerWithDefaults() *GetBundleInfo200ResponseDataBundleInfoInner {
	this := GetBundleInfo200ResponseDataBundleInfoInner{}
	return &this
}

// GetBundleId returns the BundleId field value if set, zero value otherwise.
func (o *GetBundleInfo200ResponseDataBundleInfoInner) GetBundleId() int64 {
	if o == nil || IsNil(o.BundleId) {
		var ret int64
		return ret
	}
	return *o.BundleId
}

// GetBundleIdOk returns a tuple with the BundleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInner) GetBundleIdOk() (*int64, bool) {
	if o == nil || IsNil(o.BundleId) {
		return nil, false
	}
	return o.BundleId, true
}

// HasBundleId returns a boolean if a field has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInner) HasBundleId() bool {
	if o != nil && !IsNil(o.BundleId) {
		return true
	}

	return false
}

// SetBundleId gets a reference to the given int64 and assigns it to the BundleId field.
func (o *GetBundleInfo200ResponseDataBundleInfoInner) SetBundleId(v int64) {
	o.BundleId = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *GetBundleInfo200ResponseDataBundleInfoInner) GetGroupId() int64 {
	if o == nil || IsNil(o.GroupId) {
		var ret int64
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInner) GetGroupIdOk() (*int64, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInner) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given int64 and assigns it to the GroupId field.
func (o *GetBundleInfo200ResponseDataBundleInfoInner) SetGroupId(v int64) {
	o.GroupId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetBundleInfo200ResponseDataBundleInfoInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetBundleInfo200ResponseDataBundleInfoInner) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetBundleInfo200ResponseDataBundleInfoInner) GetType() int64 {
	if o == nil || IsNil(o.Type) {
		var ret int64
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInner) GetTypeOk() (*int64, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given int64 and assigns it to the Type field.
func (o *GetBundleInfo200ResponseDataBundleInfoInner) SetType(v int64) {
	o.Type = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetBundleInfo200ResponseDataBundleInfoInner) GetStatus() int64 {
	if o == nil || IsNil(o.Status) {
		var ret int64
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInner) GetStatusOk() (*int64, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInner) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int64 and assigns it to the Status field.
func (o *GetBundleInfo200ResponseDataBundleInfoInner) SetStatus(v int64) {
	o.Status = &v
}

// GetShopId returns the ShopId field value if set, zero value otherwise.
func (o *GetBundleInfo200ResponseDataBundleInfoInner) GetShopId() int64 {
	if o == nil || IsNil(o.ShopId) {
		var ret int64
		return ret
	}
	return *o.ShopId
}

// GetShopIdOk returns a tuple with the ShopId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInner) GetShopIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ShopId) {
		return nil, false
	}
	return o.ShopId, true
}

// HasShopId returns a boolean if a field has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInner) HasShopId() bool {
	if o != nil && !IsNil(o.ShopId) {
		return true
	}

	return false
}

// SetShopId gets a reference to the given int64 and assigns it to the ShopId field.
func (o *GetBundleInfo200ResponseDataBundleInfoInner) SetShopId(v int64) {
	o.ShopId = &v
}

// GetStartTimeUnix returns the StartTimeUnix field value if set, zero value otherwise.
func (o *GetBundleInfo200ResponseDataBundleInfoInner) GetStartTimeUnix() int64 {
	if o == nil || IsNil(o.StartTimeUnix) {
		var ret int64
		return ret
	}
	return *o.StartTimeUnix
}

// GetStartTimeUnixOk returns a tuple with the StartTimeUnix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInner) GetStartTimeUnixOk() (*int64, bool) {
	if o == nil || IsNil(o.StartTimeUnix) {
		return nil, false
	}
	return o.StartTimeUnix, true
}

// HasStartTimeUnix returns a boolean if a field has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInner) HasStartTimeUnix() bool {
	if o != nil && !IsNil(o.StartTimeUnix) {
		return true
	}

	return false
}

// SetStartTimeUnix gets a reference to the given int64 and assigns it to the StartTimeUnix field.
func (o *GetBundleInfo200ResponseDataBundleInfoInner) SetStartTimeUnix(v int64) {
	o.StartTimeUnix = &v
}

// GetStopTimeUnix returns the StopTimeUnix field value if set, zero value otherwise.
func (o *GetBundleInfo200ResponseDataBundleInfoInner) GetStopTimeUnix() int64 {
	if o == nil || IsNil(o.StopTimeUnix) {
		var ret int64
		return ret
	}
	return *o.StopTimeUnix
}

// GetStopTimeUnixOk returns a tuple with the StopTimeUnix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInner) GetStopTimeUnixOk() (*int64, bool) {
	if o == nil || IsNil(o.StopTimeUnix) {
		return nil, false
	}
	return o.StopTimeUnix, true
}

// HasStopTimeUnix returns a boolean if a field has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInner) HasStopTimeUnix() bool {
	if o != nil && !IsNil(o.StopTimeUnix) {
		return true
	}

	return false
}

// SetStopTimeUnix gets a reference to the given int64 and assigns it to the StopTimeUnix field.
func (o *GetBundleInfo200ResponseDataBundleInfoInner) SetStopTimeUnix(v int64) {
	o.StopTimeUnix = &v
}

// GetBundleItem returns the BundleItem field value if set, zero value otherwise.
func (o *GetBundleInfo200ResponseDataBundleInfoInner) GetBundleItem() []GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInner {
	if o == nil || IsNil(o.BundleItem) {
		var ret []GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInner
		return ret
	}
	return o.BundleItem
}

// GetBundleItemOk returns a tuple with the BundleItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInner) GetBundleItemOk() ([]GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInner, bool) {
	if o == nil || IsNil(o.BundleItem) {
		return nil, false
	}
	return o.BundleItem, true
}

// HasBundleItem returns a boolean if a field has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInner) HasBundleItem() bool {
	if o != nil && !IsNil(o.BundleItem) {
		return true
	}

	return false
}

// SetBundleItem gets a reference to the given []GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInner and assigns it to the BundleItem field.
func (o *GetBundleInfo200ResponseDataBundleInfoInner) SetBundleItem(v []GetBundleInfo200ResponseDataBundleInfoInnerBundleItemInner) {
	o.BundleItem = v
}

// GetWarehouseId returns the WarehouseId field value if set, zero value otherwise.
func (o *GetBundleInfo200ResponseDataBundleInfoInner) GetWarehouseId() int64 {
	if o == nil || IsNil(o.WarehouseId) {
		var ret int64
		return ret
	}
	return *o.WarehouseId
}

// GetWarehouseIdOk returns a tuple with the WarehouseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInner) GetWarehouseIdOk() (*int64, bool) {
	if o == nil || IsNil(o.WarehouseId) {
		return nil, false
	}
	return o.WarehouseId, true
}

// HasWarehouseId returns a boolean if a field has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInner) HasWarehouseId() bool {
	if o != nil && !IsNil(o.WarehouseId) {
		return true
	}

	return false
}

// SetWarehouseId gets a reference to the given int64 and assigns it to the WarehouseId field.
func (o *GetBundleInfo200ResponseDataBundleInfoInner) SetWarehouseId(v int64) {
	o.WarehouseId = &v
}

// GetQuota returns the Quota field value if set, zero value otherwise.
func (o *GetBundleInfo200ResponseDataBundleInfoInner) GetQuota() int64 {
	if o == nil || IsNil(o.Quota) {
		var ret int64
		return ret
	}
	return *o.Quota
}

// GetQuotaOk returns a tuple with the Quota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInner) GetQuotaOk() (*int64, bool) {
	if o == nil || IsNil(o.Quota) {
		return nil, false
	}
	return o.Quota, true
}

// HasQuota returns a boolean if a field has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInner) HasQuota() bool {
	if o != nil && !IsNil(o.Quota) {
		return true
	}

	return false
}

// SetQuota gets a reference to the given int64 and assigns it to the Quota field.
func (o *GetBundleInfo200ResponseDataBundleInfoInner) SetQuota(v int64) {
	o.Quota = &v
}

// GetOriginalQuota returns the OriginalQuota field value if set, zero value otherwise.
func (o *GetBundleInfo200ResponseDataBundleInfoInner) GetOriginalQuota() int64 {
	if o == nil || IsNil(o.OriginalQuota) {
		var ret int64
		return ret
	}
	return *o.OriginalQuota
}

// GetOriginalQuotaOk returns a tuple with the OriginalQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInner) GetOriginalQuotaOk() (*int64, bool) {
	if o == nil || IsNil(o.OriginalQuota) {
		return nil, false
	}
	return o.OriginalQuota, true
}

// HasOriginalQuota returns a boolean if a field has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInner) HasOriginalQuota() bool {
	if o != nil && !IsNil(o.OriginalQuota) {
		return true
	}

	return false
}

// SetOriginalQuota gets a reference to the given int64 and assigns it to the OriginalQuota field.
func (o *GetBundleInfo200ResponseDataBundleInfoInner) SetOriginalQuota(v int64) {
	o.OriginalQuota = &v
}

// GetMaxOrder returns the MaxOrder field value if set, zero value otherwise.
func (o *GetBundleInfo200ResponseDataBundleInfoInner) GetMaxOrder() int64 {
	if o == nil || IsNil(o.MaxOrder) {
		var ret int64
		return ret
	}
	return *o.MaxOrder
}

// GetMaxOrderOk returns a tuple with the MaxOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInner) GetMaxOrderOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxOrder) {
		return nil, false
	}
	return o.MaxOrder, true
}

// HasMaxOrder returns a boolean if a field has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInner) HasMaxOrder() bool {
	if o != nil && !IsNil(o.MaxOrder) {
		return true
	}

	return false
}

// SetMaxOrder gets a reference to the given int64 and assigns it to the MaxOrder field.
func (o *GetBundleInfo200ResponseDataBundleInfoInner) SetMaxOrder(v int64) {
	o.MaxOrder = &v
}

// GetPreorder returns the Preorder field value if set, zero value otherwise.
func (o *GetBundleInfo200ResponseDataBundleInfoInner) GetPreorder() GetBundleInfo200ResponseDataBundleInfoInnerPreorder {
	if o == nil || IsNil(o.Preorder) {
		var ret GetBundleInfo200ResponseDataBundleInfoInnerPreorder
		return ret
	}
	return *o.Preorder
}

// GetPreorderOk returns a tuple with the Preorder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInner) GetPreorderOk() (*GetBundleInfo200ResponseDataBundleInfoInnerPreorder, bool) {
	if o == nil || IsNil(o.Preorder) {
		return nil, false
	}
	return o.Preorder, true
}

// HasPreorder returns a boolean if a field has been set.
func (o *GetBundleInfo200ResponseDataBundleInfoInner) HasPreorder() bool {
	if o != nil && !IsNil(o.Preorder) {
		return true
	}

	return false
}

// SetPreorder gets a reference to the given GetBundleInfo200ResponseDataBundleInfoInnerPreorder and assigns it to the Preorder field.
func (o *GetBundleInfo200ResponseDataBundleInfoInner) SetPreorder(v GetBundleInfo200ResponseDataBundleInfoInnerPreorder) {
	o.Preorder = &v
}

func (o GetBundleInfo200ResponseDataBundleInfoInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBundleInfo200ResponseDataBundleInfoInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BundleId) {
		toSerialize["bundle_id"] = o.BundleId
	}
	if !IsNil(o.GroupId) {
		toSerialize["group_id"] = o.GroupId
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["Type"] = o.Type
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.ShopId) {
		toSerialize["shop_id"] = o.ShopId
	}
	if !IsNil(o.StartTimeUnix) {
		toSerialize["start_time_unix"] = o.StartTimeUnix
	}
	if !IsNil(o.StopTimeUnix) {
		toSerialize["stop_time_unix"] = o.StopTimeUnix
	}
	if !IsNil(o.BundleItem) {
		toSerialize["bundle_item"] = o.BundleItem
	}
	if !IsNil(o.WarehouseId) {
		toSerialize["warehouse_id"] = o.WarehouseId
	}
	if !IsNil(o.Quota) {
		toSerialize["quota"] = o.Quota
	}
	if !IsNil(o.OriginalQuota) {
		toSerialize["original_quota"] = o.OriginalQuota
	}
	if !IsNil(o.MaxOrder) {
		toSerialize["max_order"] = o.MaxOrder
	}
	if !IsNil(o.Preorder) {
		toSerialize["preorder"] = o.Preorder
	}
	return toSerialize, nil
}

type NullableGetBundleInfo200ResponseDataBundleInfoInner struct {
	value *GetBundleInfo200ResponseDataBundleInfoInner
	isSet bool
}

func (v NullableGetBundleInfo200ResponseDataBundleInfoInner) Get() *GetBundleInfo200ResponseDataBundleInfoInner {
	return v.value
}

func (v *NullableGetBundleInfo200ResponseDataBundleInfoInner) Set(val *GetBundleInfo200ResponseDataBundleInfoInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBundleInfo200ResponseDataBundleInfoInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBundleInfo200ResponseDataBundleInfoInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBundleInfo200ResponseDataBundleInfoInner(val *GetBundleInfo200ResponseDataBundleInfoInner) *NullableGetBundleInfo200ResponseDataBundleInfoInner {
	return &NullableGetBundleInfo200ResponseDataBundleInfoInner{value: val, isSet: true}
}

func (v NullableGetBundleInfo200ResponseDataBundleInfoInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBundleInfo200ResponseDataBundleInfoInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


