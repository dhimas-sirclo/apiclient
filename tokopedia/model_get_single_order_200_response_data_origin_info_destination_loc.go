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

// checks if the GetSingleOrder200ResponseDataOriginInfoDestinationLoc type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSingleOrder200ResponseDataOriginInfoDestinationLoc{}

// GetSingleOrder200ResponseDataOriginInfoDestinationLoc struct for GetSingleOrder200ResponseDataOriginInfoDestinationLoc
type GetSingleOrder200ResponseDataOriginInfoDestinationLoc struct {
	// Latitude Coordinate
	Lat *float64 `json:"lat,omitempty"`
	// Longitude Coordinate
	Lon *float64 `json:"lon,omitempty"`
}

// NewGetSingleOrder200ResponseDataOriginInfoDestinationLoc instantiates a new GetSingleOrder200ResponseDataOriginInfoDestinationLoc object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSingleOrder200ResponseDataOriginInfoDestinationLoc() *GetSingleOrder200ResponseDataOriginInfoDestinationLoc {
	this := GetSingleOrder200ResponseDataOriginInfoDestinationLoc{}
	return &this
}

// NewGetSingleOrder200ResponseDataOriginInfoDestinationLocWithDefaults instantiates a new GetSingleOrder200ResponseDataOriginInfoDestinationLoc object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSingleOrder200ResponseDataOriginInfoDestinationLocWithDefaults() *GetSingleOrder200ResponseDataOriginInfoDestinationLoc {
	this := GetSingleOrder200ResponseDataOriginInfoDestinationLoc{}
	return &this
}

// GetLat returns the Lat field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataOriginInfoDestinationLoc) GetLat() float64 {
	if o == nil || IsNil(o.Lat) {
		var ret float64
		return ret
	}
	return *o.Lat
}

// GetLatOk returns a tuple with the Lat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataOriginInfoDestinationLoc) GetLatOk() (*float64, bool) {
	if o == nil || IsNil(o.Lat) {
		return nil, false
	}
	return o.Lat, true
}

// HasLat returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataOriginInfoDestinationLoc) HasLat() bool {
	if o != nil && !IsNil(o.Lat) {
		return true
	}

	return false
}

// SetLat gets a reference to the given float64 and assigns it to the Lat field.
func (o *GetSingleOrder200ResponseDataOriginInfoDestinationLoc) SetLat(v float64) {
	o.Lat = &v
}

// GetLon returns the Lon field value if set, zero value otherwise.
func (o *GetSingleOrder200ResponseDataOriginInfoDestinationLoc) GetLon() float64 {
	if o == nil || IsNil(o.Lon) {
		var ret float64
		return ret
	}
	return *o.Lon
}

// GetLonOk returns a tuple with the Lon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSingleOrder200ResponseDataOriginInfoDestinationLoc) GetLonOk() (*float64, bool) {
	if o == nil || IsNil(o.Lon) {
		return nil, false
	}
	return o.Lon, true
}

// HasLon returns a boolean if a field has been set.
func (o *GetSingleOrder200ResponseDataOriginInfoDestinationLoc) HasLon() bool {
	if o != nil && !IsNil(o.Lon) {
		return true
	}

	return false
}

// SetLon gets a reference to the given float64 and assigns it to the Lon field.
func (o *GetSingleOrder200ResponseDataOriginInfoDestinationLoc) SetLon(v float64) {
	o.Lon = &v
}

func (o GetSingleOrder200ResponseDataOriginInfoDestinationLoc) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSingleOrder200ResponseDataOriginInfoDestinationLoc) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Lat) {
		toSerialize["lat"] = o.Lat
	}
	if !IsNil(o.Lon) {
		toSerialize["lon"] = o.Lon
	}
	return toSerialize, nil
}

type NullableGetSingleOrder200ResponseDataOriginInfoDestinationLoc struct {
	value *GetSingleOrder200ResponseDataOriginInfoDestinationLoc
	isSet bool
}

func (v NullableGetSingleOrder200ResponseDataOriginInfoDestinationLoc) Get() *GetSingleOrder200ResponseDataOriginInfoDestinationLoc {
	return v.value
}

func (v *NullableGetSingleOrder200ResponseDataOriginInfoDestinationLoc) Set(val *GetSingleOrder200ResponseDataOriginInfoDestinationLoc) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSingleOrder200ResponseDataOriginInfoDestinationLoc) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSingleOrder200ResponseDataOriginInfoDestinationLoc) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSingleOrder200ResponseDataOriginInfoDestinationLoc(val *GetSingleOrder200ResponseDataOriginInfoDestinationLoc) *NullableGetSingleOrder200ResponseDataOriginInfoDestinationLoc {
	return &NullableGetSingleOrder200ResponseDataOriginInfoDestinationLoc{value: val, isSet: true}
}

func (v NullableGetSingleOrder200ResponseDataOriginInfoDestinationLoc) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSingleOrder200ResponseDataOriginInfoDestinationLoc) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


