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

// checks if the V0039JobInfoPower type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0039JobInfoPower{}

// V0039JobInfoPower struct for V0039JobInfoPower
type V0039JobInfoPower struct {
	Flags []string `json:"flags,omitempty"`
}

// NewV0039JobInfoPower instantiates a new V0039JobInfoPower object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0039JobInfoPower() *V0039JobInfoPower {
	this := V0039JobInfoPower{}
	return &this
}

// NewV0039JobInfoPowerWithDefaults instantiates a new V0039JobInfoPower object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0039JobInfoPowerWithDefaults() *V0039JobInfoPower {
	this := V0039JobInfoPower{}
	return &this
}

// GetFlags returns the Flags field value if set, zero value otherwise.
func (o *V0039JobInfoPower) GetFlags() []string {
	if o == nil || IsNil(o.Flags) {
		var ret []string
		return ret
	}
	return o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039JobInfoPower) GetFlagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Flags) {
		return nil, false
	}
	return o.Flags, true
}

// HasFlags returns a boolean if a field has been set.
func (o *V0039JobInfoPower) HasFlags() bool {
	if o != nil && !IsNil(o.Flags) {
		return true
	}

	return false
}

// SetFlags gets a reference to the given []string and assigns it to the Flags field.
func (o *V0039JobInfoPower) SetFlags(v []string) {
	o.Flags = v
}

func (o V0039JobInfoPower) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0039JobInfoPower) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Flags) {
		toSerialize["flags"] = o.Flags
	}
	return toSerialize, nil
}

type NullableV0039JobInfoPower struct {
	value *V0039JobInfoPower
	isSet bool
}

func (v NullableV0039JobInfoPower) Get() *V0039JobInfoPower {
	return v.value
}

func (v *NullableV0039JobInfoPower) Set(val *V0039JobInfoPower) {
	v.value = val
	v.isSet = true
}

func (v NullableV0039JobInfoPower) IsSet() bool {
	return v.isSet
}

func (v *NullableV0039JobInfoPower) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0039JobInfoPower(val *V0039JobInfoPower) *NullableV0039JobInfoPower {
	return &NullableV0039JobInfoPower{value: val, isSet: true}
}

func (v NullableV0039JobInfoPower) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0039JobInfoPower) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
