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

// checks if the V0039StepStatistics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0039StepStatistics{}

// V0039StepStatistics struct for V0039StepStatistics
type V0039StepStatistics struct {
	Energy *V0039StepStatisticsEnergy `json:"energy,omitempty"`
}

// NewV0039StepStatistics instantiates a new V0039StepStatistics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0039StepStatistics() *V0039StepStatistics {
	this := V0039StepStatistics{}
	return &this
}

// NewV0039StepStatisticsWithDefaults instantiates a new V0039StepStatistics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0039StepStatisticsWithDefaults() *V0039StepStatistics {
	this := V0039StepStatistics{}
	return &this
}

// GetEnergy returns the Energy field value if set, zero value otherwise.
func (o *V0039StepStatistics) GetEnergy() V0039StepStatisticsEnergy {
	if o == nil || IsNil(o.Energy) {
		var ret V0039StepStatisticsEnergy
		return ret
	}
	return *o.Energy
}

// GetEnergyOk returns a tuple with the Energy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039StepStatistics) GetEnergyOk() (*V0039StepStatisticsEnergy, bool) {
	if o == nil || IsNil(o.Energy) {
		return nil, false
	}
	return o.Energy, true
}

// HasEnergy returns a boolean if a field has been set.
func (o *V0039StepStatistics) HasEnergy() bool {
	if o != nil && !IsNil(o.Energy) {
		return true
	}

	return false
}

// SetEnergy gets a reference to the given V0039StepStatisticsEnergy and assigns it to the Energy field.
func (o *V0039StepStatistics) SetEnergy(v V0039StepStatisticsEnergy) {
	o.Energy = &v
}

func (o V0039StepStatistics) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0039StepStatistics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Energy) {
		toSerialize["energy"] = o.Energy
	}
	return toSerialize, nil
}

type NullableV0039StepStatistics struct {
	value *V0039StepStatistics
	isSet bool
}

func (v NullableV0039StepStatistics) Get() *V0039StepStatistics {
	return v.value
}

func (v *NullableV0039StepStatistics) Set(val *V0039StepStatistics) {
	v.value = val
	v.isSet = true
}

func (v NullableV0039StepStatistics) IsSet() bool {
	return v.isSet
}

func (v *NullableV0039StepStatistics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0039StepStatistics(val *V0039StepStatistics) *NullableV0039StepStatistics {
	return &NullableV0039StepStatistics{value: val, isSet: true}
}

func (v NullableV0039StepStatistics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0039StepStatistics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
