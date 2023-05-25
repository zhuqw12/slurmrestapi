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

// checks if the V0039StepTimeUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0039StepTimeUser{}

// V0039StepTimeUser struct for V0039StepTimeUser
type V0039StepTimeUser struct {
	Microseconds *int32 `json:"microseconds,omitempty"`
}

// NewV0039StepTimeUser instantiates a new V0039StepTimeUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0039StepTimeUser() *V0039StepTimeUser {
	this := V0039StepTimeUser{}
	return &this
}

// NewV0039StepTimeUserWithDefaults instantiates a new V0039StepTimeUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0039StepTimeUserWithDefaults() *V0039StepTimeUser {
	this := V0039StepTimeUser{}
	return &this
}

// GetMicroseconds returns the Microseconds field value if set, zero value otherwise.
func (o *V0039StepTimeUser) GetMicroseconds() int32 {
	if o == nil || IsNil(o.Microseconds) {
		var ret int32
		return ret
	}
	return *o.Microseconds
}

// GetMicrosecondsOk returns a tuple with the Microseconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039StepTimeUser) GetMicrosecondsOk() (*int32, bool) {
	if o == nil || IsNil(o.Microseconds) {
		return nil, false
	}
	return o.Microseconds, true
}

// HasMicroseconds returns a boolean if a field has been set.
func (o *V0039StepTimeUser) HasMicroseconds() bool {
	if o != nil && !IsNil(o.Microseconds) {
		return true
	}

	return false
}

// SetMicroseconds gets a reference to the given int32 and assigns it to the Microseconds field.
func (o *V0039StepTimeUser) SetMicroseconds(v int32) {
	o.Microseconds = &v
}

func (o V0039StepTimeUser) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0039StepTimeUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Microseconds) {
		toSerialize["microseconds"] = o.Microseconds
	}
	return toSerialize, nil
}

type NullableV0039StepTimeUser struct {
	value *V0039StepTimeUser
	isSet bool
}

func (v NullableV0039StepTimeUser) Get() *V0039StepTimeUser {
	return v.value
}

func (v *NullableV0039StepTimeUser) Set(val *V0039StepTimeUser) {
	v.value = val
	v.isSet = true
}

func (v NullableV0039StepTimeUser) IsSet() bool {
	return v.isSet
}

func (v *NullableV0039StepTimeUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0039StepTimeUser(val *V0039StepTimeUser) *NullableV0039StepTimeUser {
	return &NullableV0039StepTimeUser{value: val, isSet: true}
}

func (v NullableV0039StepTimeUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0039StepTimeUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
