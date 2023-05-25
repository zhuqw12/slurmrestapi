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

// checks if the Dbv0038QosLimitsMin type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0038QosLimitsMin{}

// Dbv0038QosLimitsMin Min limit settings
type Dbv0038QosLimitsMin struct {
	// Min priority threshold
	PriorityThreshold *int32                   `json:"priority_threshold,omitempty"`
	Tres              *Dbv0038QosLimitsMinTres `json:"tres,omitempty"`
}

// NewDbv0038QosLimitsMin instantiates a new Dbv0038QosLimitsMin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0038QosLimitsMin() *Dbv0038QosLimitsMin {
	this := Dbv0038QosLimitsMin{}
	return &this
}

// NewDbv0038QosLimitsMinWithDefaults instantiates a new Dbv0038QosLimitsMin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0038QosLimitsMinWithDefaults() *Dbv0038QosLimitsMin {
	this := Dbv0038QosLimitsMin{}
	return &this
}

// GetPriorityThreshold returns the PriorityThreshold field value if set, zero value otherwise.
func (o *Dbv0038QosLimitsMin) GetPriorityThreshold() int32 {
	if o == nil || IsNil(o.PriorityThreshold) {
		var ret int32
		return ret
	}
	return *o.PriorityThreshold
}

// GetPriorityThresholdOk returns a tuple with the PriorityThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038QosLimitsMin) GetPriorityThresholdOk() (*int32, bool) {
	if o == nil || IsNil(o.PriorityThreshold) {
		return nil, false
	}
	return o.PriorityThreshold, true
}

// HasPriorityThreshold returns a boolean if a field has been set.
func (o *Dbv0038QosLimitsMin) HasPriorityThreshold() bool {
	if o != nil && !IsNil(o.PriorityThreshold) {
		return true
	}

	return false
}

// SetPriorityThreshold gets a reference to the given int32 and assigns it to the PriorityThreshold field.
func (o *Dbv0038QosLimitsMin) SetPriorityThreshold(v int32) {
	o.PriorityThreshold = &v
}

// GetTres returns the Tres field value if set, zero value otherwise.
func (o *Dbv0038QosLimitsMin) GetTres() Dbv0038QosLimitsMinTres {
	if o == nil || IsNil(o.Tres) {
		var ret Dbv0038QosLimitsMinTres
		return ret
	}
	return *o.Tres
}

// GetTresOk returns a tuple with the Tres field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038QosLimitsMin) GetTresOk() (*Dbv0038QosLimitsMinTres, bool) {
	if o == nil || IsNil(o.Tres) {
		return nil, false
	}
	return o.Tres, true
}

// HasTres returns a boolean if a field has been set.
func (o *Dbv0038QosLimitsMin) HasTres() bool {
	if o != nil && !IsNil(o.Tres) {
		return true
	}

	return false
}

// SetTres gets a reference to the given Dbv0038QosLimitsMinTres and assigns it to the Tres field.
func (o *Dbv0038QosLimitsMin) SetTres(v Dbv0038QosLimitsMinTres) {
	o.Tres = &v
}

func (o Dbv0038QosLimitsMin) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0038QosLimitsMin) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PriorityThreshold) {
		toSerialize["priority_threshold"] = o.PriorityThreshold
	}
	if !IsNil(o.Tres) {
		toSerialize["tres"] = o.Tres
	}
	return toSerialize, nil
}

type NullableDbv0038QosLimitsMin struct {
	value *Dbv0038QosLimitsMin
	isSet bool
}

func (v NullableDbv0038QosLimitsMin) Get() *Dbv0038QosLimitsMin {
	return v.value
}

func (v *NullableDbv0038QosLimitsMin) Set(val *Dbv0038QosLimitsMin) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0038QosLimitsMin) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0038QosLimitsMin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0038QosLimitsMin(val *Dbv0038QosLimitsMin) *NullableDbv0038QosLimitsMin {
	return &NullableDbv0038QosLimitsMin{value: val, isSet: true}
}

func (v NullableDbv0038QosLimitsMin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0038QosLimitsMin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
