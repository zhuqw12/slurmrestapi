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

// checks if the Dbv0038JobStepCPURequestedFrequency type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0038JobStepCPURequestedFrequency{}

// Dbv0038JobStepCPURequestedFrequency CPU frequency requested
type Dbv0038JobStepCPURequestedFrequency struct {
	// Min CPU frequency
	Min *int32 `json:"min,omitempty"`
	// Max CPU frequency
	Max *int32 `json:"max,omitempty"`
}

// NewDbv0038JobStepCPURequestedFrequency instantiates a new Dbv0038JobStepCPURequestedFrequency object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0038JobStepCPURequestedFrequency() *Dbv0038JobStepCPURequestedFrequency {
	this := Dbv0038JobStepCPURequestedFrequency{}
	return &this
}

// NewDbv0038JobStepCPURequestedFrequencyWithDefaults instantiates a new Dbv0038JobStepCPURequestedFrequency object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0038JobStepCPURequestedFrequencyWithDefaults() *Dbv0038JobStepCPURequestedFrequency {
	this := Dbv0038JobStepCPURequestedFrequency{}
	return &this
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *Dbv0038JobStepCPURequestedFrequency) GetMin() int32 {
	if o == nil || IsNil(o.Min) {
		var ret int32
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038JobStepCPURequestedFrequency) GetMinOk() (*int32, bool) {
	if o == nil || IsNil(o.Min) {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *Dbv0038JobStepCPURequestedFrequency) HasMin() bool {
	if o != nil && !IsNil(o.Min) {
		return true
	}

	return false
}

// SetMin gets a reference to the given int32 and assigns it to the Min field.
func (o *Dbv0038JobStepCPURequestedFrequency) SetMin(v int32) {
	o.Min = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *Dbv0038JobStepCPURequestedFrequency) GetMax() int32 {
	if o == nil || IsNil(o.Max) {
		var ret int32
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038JobStepCPURequestedFrequency) GetMaxOk() (*int32, bool) {
	if o == nil || IsNil(o.Max) {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *Dbv0038JobStepCPURequestedFrequency) HasMax() bool {
	if o != nil && !IsNil(o.Max) {
		return true
	}

	return false
}

// SetMax gets a reference to the given int32 and assigns it to the Max field.
func (o *Dbv0038JobStepCPURequestedFrequency) SetMax(v int32) {
	o.Max = &v
}

func (o Dbv0038JobStepCPURequestedFrequency) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0038JobStepCPURequestedFrequency) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Min) {
		toSerialize["min"] = o.Min
	}
	if !IsNil(o.Max) {
		toSerialize["max"] = o.Max
	}
	return toSerialize, nil
}

type NullableDbv0038JobStepCPURequestedFrequency struct {
	value *Dbv0038JobStepCPURequestedFrequency
	isSet bool
}

func (v NullableDbv0038JobStepCPURequestedFrequency) Get() *Dbv0038JobStepCPURequestedFrequency {
	return v.value
}

func (v *NullableDbv0038JobStepCPURequestedFrequency) Set(val *Dbv0038JobStepCPURequestedFrequency) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0038JobStepCPURequestedFrequency) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0038JobStepCPURequestedFrequency) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0038JobStepCPURequestedFrequency(val *Dbv0038JobStepCPURequestedFrequency) *NullableDbv0038JobStepCPURequestedFrequency {
	return &NullableDbv0038JobStepCPURequestedFrequency{value: val, isSet: true}
}

func (v NullableDbv0038JobStepCPURequestedFrequency) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0038JobStepCPURequestedFrequency) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
