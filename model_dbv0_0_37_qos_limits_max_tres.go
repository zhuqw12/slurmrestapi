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

// checks if the Dbv0037QosLimitsMaxTres type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0037QosLimitsMaxTres{}

// Dbv0037QosLimitsMaxTres Limits on TRES
type Dbv0037QosLimitsMaxTres struct {
	Minutes *Dbv0037QosLimitsMaxTresMinutes `json:"minutes,omitempty"`
	Per     *Dbv0037QosLimitsMaxTresPer     `json:"per,omitempty"`
}

// NewDbv0037QosLimitsMaxTres instantiates a new Dbv0037QosLimitsMaxTres object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0037QosLimitsMaxTres() *Dbv0037QosLimitsMaxTres {
	this := Dbv0037QosLimitsMaxTres{}
	return &this
}

// NewDbv0037QosLimitsMaxTresWithDefaults instantiates a new Dbv0037QosLimitsMaxTres object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0037QosLimitsMaxTresWithDefaults() *Dbv0037QosLimitsMaxTres {
	this := Dbv0037QosLimitsMaxTres{}
	return &this
}

// GetMinutes returns the Minutes field value if set, zero value otherwise.
func (o *Dbv0037QosLimitsMaxTres) GetMinutes() Dbv0037QosLimitsMaxTresMinutes {
	if o == nil || IsNil(o.Minutes) {
		var ret Dbv0037QosLimitsMaxTresMinutes
		return ret
	}
	return *o.Minutes
}

// GetMinutesOk returns a tuple with the Minutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037QosLimitsMaxTres) GetMinutesOk() (*Dbv0037QosLimitsMaxTresMinutes, bool) {
	if o == nil || IsNil(o.Minutes) {
		return nil, false
	}
	return o.Minutes, true
}

// HasMinutes returns a boolean if a field has been set.
func (o *Dbv0037QosLimitsMaxTres) HasMinutes() bool {
	if o != nil && !IsNil(o.Minutes) {
		return true
	}

	return false
}

// SetMinutes gets a reference to the given Dbv0037QosLimitsMaxTresMinutes and assigns it to the Minutes field.
func (o *Dbv0037QosLimitsMaxTres) SetMinutes(v Dbv0037QosLimitsMaxTresMinutes) {
	o.Minutes = &v
}

// GetPer returns the Per field value if set, zero value otherwise.
func (o *Dbv0037QosLimitsMaxTres) GetPer() Dbv0037QosLimitsMaxTresPer {
	if o == nil || IsNil(o.Per) {
		var ret Dbv0037QosLimitsMaxTresPer
		return ret
	}
	return *o.Per
}

// GetPerOk returns a tuple with the Per field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037QosLimitsMaxTres) GetPerOk() (*Dbv0037QosLimitsMaxTresPer, bool) {
	if o == nil || IsNil(o.Per) {
		return nil, false
	}
	return o.Per, true
}

// HasPer returns a boolean if a field has been set.
func (o *Dbv0037QosLimitsMaxTres) HasPer() bool {
	if o != nil && !IsNil(o.Per) {
		return true
	}

	return false
}

// SetPer gets a reference to the given Dbv0037QosLimitsMaxTresPer and assigns it to the Per field.
func (o *Dbv0037QosLimitsMaxTres) SetPer(v Dbv0037QosLimitsMaxTresPer) {
	o.Per = &v
}

func (o Dbv0037QosLimitsMaxTres) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0037QosLimitsMaxTres) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Minutes) {
		toSerialize["minutes"] = o.Minutes
	}
	if !IsNil(o.Per) {
		toSerialize["per"] = o.Per
	}
	return toSerialize, nil
}

type NullableDbv0037QosLimitsMaxTres struct {
	value *Dbv0037QosLimitsMaxTres
	isSet bool
}

func (v NullableDbv0037QosLimitsMaxTres) Get() *Dbv0037QosLimitsMaxTres {
	return v.value
}

func (v *NullableDbv0037QosLimitsMaxTres) Set(val *Dbv0037QosLimitsMaxTres) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0037QosLimitsMaxTres) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0037QosLimitsMaxTres) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0037QosLimitsMaxTres(val *Dbv0037QosLimitsMaxTres) *NullableDbv0037QosLimitsMaxTres {
	return &NullableDbv0037QosLimitsMaxTres{value: val, isSet: true}
}

func (v NullableDbv0037QosLimitsMaxTres) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0037QosLimitsMaxTres) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
