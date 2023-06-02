/*
Slurm Rest API

API to access and control Slurm.

API version: v0.0.38
Contact: sales@schedmd.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package slurmrestapi

import (
	"encoding/json"
)

// checks if the Dbv0038QosLimitsMaxWallClock type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0038QosLimitsMaxWallClock{}

// Dbv0038QosLimitsMaxWallClock Limit on wallclock settings
type Dbv0038QosLimitsMaxWallClock struct {
	Per *Dbv0038QosLimitsMaxWallClockPer `json:"per,omitempty"`
}

// NewDbv0038QosLimitsMaxWallClock instantiates a new Dbv0038QosLimitsMaxWallClock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0038QosLimitsMaxWallClock() *Dbv0038QosLimitsMaxWallClock {
	this := Dbv0038QosLimitsMaxWallClock{}
	return &this
}

// NewDbv0038QosLimitsMaxWallClockWithDefaults instantiates a new Dbv0038QosLimitsMaxWallClock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0038QosLimitsMaxWallClockWithDefaults() *Dbv0038QosLimitsMaxWallClock {
	this := Dbv0038QosLimitsMaxWallClock{}
	return &this
}

// GetPer returns the Per field value if set, zero value otherwise.
func (o *Dbv0038QosLimitsMaxWallClock) GetPer() Dbv0038QosLimitsMaxWallClockPer {
	if o == nil || IsNil(o.Per) {
		var ret Dbv0038QosLimitsMaxWallClockPer
		return ret
	}
	return *o.Per
}

// GetPerOk returns a tuple with the Per field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038QosLimitsMaxWallClock) GetPerOk() (*Dbv0038QosLimitsMaxWallClockPer, bool) {
	if o == nil || IsNil(o.Per) {
		return nil, false
	}
	return o.Per, true
}

// HasPer returns a boolean if a field has been set.
func (o *Dbv0038QosLimitsMaxWallClock) HasPer() bool {
	if o != nil && !IsNil(o.Per) {
		return true
	}

	return false
}

// SetPer gets a reference to the given Dbv0038QosLimitsMaxWallClockPer and assigns it to the Per field.
func (o *Dbv0038QosLimitsMaxWallClock) SetPer(v Dbv0038QosLimitsMaxWallClockPer) {
	o.Per = &v
}

func (o Dbv0038QosLimitsMaxWallClock) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0038QosLimitsMaxWallClock) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Per) {
		toSerialize["per"] = o.Per
	}
	return toSerialize, nil
}

type NullableDbv0038QosLimitsMaxWallClock struct {
	value *Dbv0038QosLimitsMaxWallClock
	isSet bool
}

func (v NullableDbv0038QosLimitsMaxWallClock) Get() *Dbv0038QosLimitsMaxWallClock {
	return v.value
}

func (v *NullableDbv0038QosLimitsMaxWallClock) Set(val *Dbv0038QosLimitsMaxWallClock) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0038QosLimitsMaxWallClock) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0038QosLimitsMaxWallClock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0038QosLimitsMaxWallClock(val *Dbv0038QosLimitsMaxWallClock) *NullableDbv0038QosLimitsMaxWallClock {
	return &NullableDbv0038QosLimitsMaxWallClock{value: val, isSet: true}
}

func (v NullableDbv0038QosLimitsMaxWallClock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0038QosLimitsMaxWallClock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
