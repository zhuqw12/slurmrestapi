/*
Slurm Rest API

API to access and control Slurm.

API version: v0.0.39
Contact: sales@schedmd.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package slurmrestapi

import (
	"encoding/json"
)

// checks if the V0039QosLimitsMinTres type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0039QosLimitsMinTres{}

// V0039QosLimitsMinTres struct for V0039QosLimitsMinTres
type V0039QosLimitsMinTres struct {
	Per *V0039QosLimitsMinTresPer `json:"per,omitempty"`
}

// NewV0039QosLimitsMinTres instantiates a new V0039QosLimitsMinTres object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0039QosLimitsMinTres() *V0039QosLimitsMinTres {
	this := V0039QosLimitsMinTres{}
	return &this
}

// NewV0039QosLimitsMinTresWithDefaults instantiates a new V0039QosLimitsMinTres object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0039QosLimitsMinTresWithDefaults() *V0039QosLimitsMinTres {
	this := V0039QosLimitsMinTres{}
	return &this
}

// GetPer returns the Per field value if set, zero value otherwise.
func (o *V0039QosLimitsMinTres) GetPer() V0039QosLimitsMinTresPer {
	if o == nil || IsNil(o.Per) {
		var ret V0039QosLimitsMinTresPer
		return ret
	}
	return *o.Per
}

// GetPerOk returns a tuple with the Per field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039QosLimitsMinTres) GetPerOk() (*V0039QosLimitsMinTresPer, bool) {
	if o == nil || IsNil(o.Per) {
		return nil, false
	}
	return o.Per, true
}

// HasPer returns a boolean if a field has been set.
func (o *V0039QosLimitsMinTres) HasPer() bool {
	if o != nil && !IsNil(o.Per) {
		return true
	}

	return false
}

// SetPer gets a reference to the given V0039QosLimitsMinTresPer and assigns it to the Per field.
func (o *V0039QosLimitsMinTres) SetPer(v V0039QosLimitsMinTresPer) {
	o.Per = &v
}

func (o V0039QosLimitsMinTres) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0039QosLimitsMinTres) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Per) {
		toSerialize["per"] = o.Per
	}
	return toSerialize, nil
}

type NullableV0039QosLimitsMinTres struct {
	value *V0039QosLimitsMinTres
	isSet bool
}

func (v NullableV0039QosLimitsMinTres) Get() *V0039QosLimitsMinTres {
	return v.value
}

func (v *NullableV0039QosLimitsMinTres) Set(val *V0039QosLimitsMinTres) {
	v.value = val
	v.isSet = true
}

func (v NullableV0039QosLimitsMinTres) IsSet() bool {
	return v.isSet
}

func (v *NullableV0039QosLimitsMinTres) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0039QosLimitsMinTres(val *V0039QosLimitsMinTres) *NullableV0039QosLimitsMinTres {
	return &NullableV0039QosLimitsMinTres{value: val, isSet: true}
}

func (v NullableV0039QosLimitsMinTres) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0039QosLimitsMinTres) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
