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

// checks if the Dbv0038DiagStatisticsRPCsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0038DiagStatisticsRPCsInner{}

// Dbv0038DiagStatisticsRPCsInner Statistics by RPC type
type Dbv0038DiagStatisticsRPCsInner struct {
	// RPC type
	Rpc *string `json:"rpc,omitempty"`
	// Number of RPCs
	Count *int32                              `json:"count,omitempty"`
	Time  *Dbv0038DiagStatisticsRPCsInnerTime `json:"time,omitempty"`
}

// NewDbv0038DiagStatisticsRPCsInner instantiates a new Dbv0038DiagStatisticsRPCsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0038DiagStatisticsRPCsInner() *Dbv0038DiagStatisticsRPCsInner {
	this := Dbv0038DiagStatisticsRPCsInner{}
	return &this
}

// NewDbv0038DiagStatisticsRPCsInnerWithDefaults instantiates a new Dbv0038DiagStatisticsRPCsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0038DiagStatisticsRPCsInnerWithDefaults() *Dbv0038DiagStatisticsRPCsInner {
	this := Dbv0038DiagStatisticsRPCsInner{}
	return &this
}

// GetRpc returns the Rpc field value if set, zero value otherwise.
func (o *Dbv0038DiagStatisticsRPCsInner) GetRpc() string {
	if o == nil || IsNil(o.Rpc) {
		var ret string
		return ret
	}
	return *o.Rpc
}

// GetRpcOk returns a tuple with the Rpc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038DiagStatisticsRPCsInner) GetRpcOk() (*string, bool) {
	if o == nil || IsNil(o.Rpc) {
		return nil, false
	}
	return o.Rpc, true
}

// HasRpc returns a boolean if a field has been set.
func (o *Dbv0038DiagStatisticsRPCsInner) HasRpc() bool {
	if o != nil && !IsNil(o.Rpc) {
		return true
	}

	return false
}

// SetRpc gets a reference to the given string and assigns it to the Rpc field.
func (o *Dbv0038DiagStatisticsRPCsInner) SetRpc(v string) {
	o.Rpc = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *Dbv0038DiagStatisticsRPCsInner) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038DiagStatisticsRPCsInner) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *Dbv0038DiagStatisticsRPCsInner) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *Dbv0038DiagStatisticsRPCsInner) SetCount(v int32) {
	o.Count = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *Dbv0038DiagStatisticsRPCsInner) GetTime() Dbv0038DiagStatisticsRPCsInnerTime {
	if o == nil || IsNil(o.Time) {
		var ret Dbv0038DiagStatisticsRPCsInnerTime
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038DiagStatisticsRPCsInner) GetTimeOk() (*Dbv0038DiagStatisticsRPCsInnerTime, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *Dbv0038DiagStatisticsRPCsInner) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given Dbv0038DiagStatisticsRPCsInnerTime and assigns it to the Time field.
func (o *Dbv0038DiagStatisticsRPCsInner) SetTime(v Dbv0038DiagStatisticsRPCsInnerTime) {
	o.Time = &v
}

func (o Dbv0038DiagStatisticsRPCsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0038DiagStatisticsRPCsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Rpc) {
		toSerialize["rpc"] = o.Rpc
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	return toSerialize, nil
}

type NullableDbv0038DiagStatisticsRPCsInner struct {
	value *Dbv0038DiagStatisticsRPCsInner
	isSet bool
}

func (v NullableDbv0038DiagStatisticsRPCsInner) Get() *Dbv0038DiagStatisticsRPCsInner {
	return v.value
}

func (v *NullableDbv0038DiagStatisticsRPCsInner) Set(val *Dbv0038DiagStatisticsRPCsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0038DiagStatisticsRPCsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0038DiagStatisticsRPCsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0038DiagStatisticsRPCsInner(val *Dbv0038DiagStatisticsRPCsInner) *NullableDbv0038DiagStatisticsRPCsInner {
	return &NullableDbv0038DiagStatisticsRPCsInner{value: val, isSet: true}
}

func (v NullableDbv0038DiagStatisticsRPCsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0038DiagStatisticsRPCsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
