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

// checks if the Dbv0038JobStepStatistics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0038JobStepStatistics{}

// Dbv0038JobStepStatistics Statistics of job step
type Dbv0038JobStepStatistics struct {
	CPU    *Dbv0038JobStepStatisticsCPU    `json:"CPU,omitempty"`
	Energy *Dbv0038JobStepStatisticsEnergy `json:"energy,omitempty"`
}

// NewDbv0038JobStepStatistics instantiates a new Dbv0038JobStepStatistics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0038JobStepStatistics() *Dbv0038JobStepStatistics {
	this := Dbv0038JobStepStatistics{}
	return &this
}

// NewDbv0038JobStepStatisticsWithDefaults instantiates a new Dbv0038JobStepStatistics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0038JobStepStatisticsWithDefaults() *Dbv0038JobStepStatistics {
	this := Dbv0038JobStepStatistics{}
	return &this
}

// GetCPU returns the CPU field value if set, zero value otherwise.
func (o *Dbv0038JobStepStatistics) GetCPU() Dbv0038JobStepStatisticsCPU {
	if o == nil || IsNil(o.CPU) {
		var ret Dbv0038JobStepStatisticsCPU
		return ret
	}
	return *o.CPU
}

// GetCPUOk returns a tuple with the CPU field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038JobStepStatistics) GetCPUOk() (*Dbv0038JobStepStatisticsCPU, bool) {
	if o == nil || IsNil(o.CPU) {
		return nil, false
	}
	return o.CPU, true
}

// HasCPU returns a boolean if a field has been set.
func (o *Dbv0038JobStepStatistics) HasCPU() bool {
	if o != nil && !IsNil(o.CPU) {
		return true
	}

	return false
}

// SetCPU gets a reference to the given Dbv0038JobStepStatisticsCPU and assigns it to the CPU field.
func (o *Dbv0038JobStepStatistics) SetCPU(v Dbv0038JobStepStatisticsCPU) {
	o.CPU = &v
}

// GetEnergy returns the Energy field value if set, zero value otherwise.
func (o *Dbv0038JobStepStatistics) GetEnergy() Dbv0038JobStepStatisticsEnergy {
	if o == nil || IsNil(o.Energy) {
		var ret Dbv0038JobStepStatisticsEnergy
		return ret
	}
	return *o.Energy
}

// GetEnergyOk returns a tuple with the Energy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038JobStepStatistics) GetEnergyOk() (*Dbv0038JobStepStatisticsEnergy, bool) {
	if o == nil || IsNil(o.Energy) {
		return nil, false
	}
	return o.Energy, true
}

// HasEnergy returns a boolean if a field has been set.
func (o *Dbv0038JobStepStatistics) HasEnergy() bool {
	if o != nil && !IsNil(o.Energy) {
		return true
	}

	return false
}

// SetEnergy gets a reference to the given Dbv0038JobStepStatisticsEnergy and assigns it to the Energy field.
func (o *Dbv0038JobStepStatistics) SetEnergy(v Dbv0038JobStepStatisticsEnergy) {
	o.Energy = &v
}

func (o Dbv0038JobStepStatistics) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0038JobStepStatistics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CPU) {
		toSerialize["CPU"] = o.CPU
	}
	if !IsNil(o.Energy) {
		toSerialize["energy"] = o.Energy
	}
	return toSerialize, nil
}

type NullableDbv0038JobStepStatistics struct {
	value *Dbv0038JobStepStatistics
	isSet bool
}

func (v NullableDbv0038JobStepStatistics) Get() *Dbv0038JobStepStatistics {
	return v.value
}

func (v *NullableDbv0038JobStepStatistics) Set(val *Dbv0038JobStepStatistics) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0038JobStepStatistics) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0038JobStepStatistics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0038JobStepStatistics(val *Dbv0038JobStepStatistics) *NullableDbv0038JobStepStatistics {
	return &NullableDbv0038JobStepStatistics{value: val, isSet: true}
}

func (v NullableDbv0038JobStepStatistics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0038JobStepStatistics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
