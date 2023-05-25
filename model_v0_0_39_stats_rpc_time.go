/*
Slurm Rest API

API to access and control Slurm DB.

API version: dbv0.0.38
Contact: sales@schedmd.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package slurmrestopenapi

import (
	"encoding/json"
)

// checks if the V0039StatsRpcTime type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0039StatsRpcTime{}

// V0039StatsRpcTime struct for V0039StatsRpcTime
type V0039StatsRpcTime struct {
	Total *int64 `json:"total,omitempty"`
}

// NewV0039StatsRpcTime instantiates a new V0039StatsRpcTime object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0039StatsRpcTime() *V0039StatsRpcTime {
	this := V0039StatsRpcTime{}
	return &this
}

// NewV0039StatsRpcTimeWithDefaults instantiates a new V0039StatsRpcTime object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0039StatsRpcTimeWithDefaults() *V0039StatsRpcTime {
	this := V0039StatsRpcTime{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *V0039StatsRpcTime) GetTotal() int64 {
	if o == nil || IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039StatsRpcTime) GetTotalOk() (*int64, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *V0039StatsRpcTime) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *V0039StatsRpcTime) SetTotal(v int64) {
	o.Total = &v
}

func (o V0039StatsRpcTime) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0039StatsRpcTime) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableV0039StatsRpcTime struct {
	value *V0039StatsRpcTime
	isSet bool
}

func (v NullableV0039StatsRpcTime) Get() *V0039StatsRpcTime {
	return v.value
}

func (v *NullableV0039StatsRpcTime) Set(val *V0039StatsRpcTime) {
	v.value = val
	v.isSet = true
}

func (v NullableV0039StatsRpcTime) IsSet() bool {
	return v.isSet
}

func (v *NullableV0039StatsRpcTime) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0039StatsRpcTime(val *V0039StatsRpcTime) *NullableV0039StatsRpcTime {
	return &NullableV0039StatsRpcTime{value: val, isSet: true}
}

func (v NullableV0039StatsRpcTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0039StatsRpcTime) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
