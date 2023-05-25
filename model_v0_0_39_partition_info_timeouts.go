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

// checks if the V0039PartitionInfoTimeouts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0039PartitionInfoTimeouts{}

// V0039PartitionInfoTimeouts struct for V0039PartitionInfoTimeouts
type V0039PartitionInfoTimeouts struct {
	Suspend *V0039Uint16NoVal `json:"suspend,omitempty"`
}

// NewV0039PartitionInfoTimeouts instantiates a new V0039PartitionInfoTimeouts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0039PartitionInfoTimeouts() *V0039PartitionInfoTimeouts {
	this := V0039PartitionInfoTimeouts{}
	return &this
}

// NewV0039PartitionInfoTimeoutsWithDefaults instantiates a new V0039PartitionInfoTimeouts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0039PartitionInfoTimeoutsWithDefaults() *V0039PartitionInfoTimeouts {
	this := V0039PartitionInfoTimeouts{}
	return &this
}

// GetSuspend returns the Suspend field value if set, zero value otherwise.
func (o *V0039PartitionInfoTimeouts) GetSuspend() V0039Uint16NoVal {
	if o == nil || IsNil(o.Suspend) {
		var ret V0039Uint16NoVal
		return ret
	}
	return *o.Suspend
}

// GetSuspendOk returns a tuple with the Suspend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039PartitionInfoTimeouts) GetSuspendOk() (*V0039Uint16NoVal, bool) {
	if o == nil || IsNil(o.Suspend) {
		return nil, false
	}
	return o.Suspend, true
}

// HasSuspend returns a boolean if a field has been set.
func (o *V0039PartitionInfoTimeouts) HasSuspend() bool {
	if o != nil && !IsNil(o.Suspend) {
		return true
	}

	return false
}

// SetSuspend gets a reference to the given V0039Uint16NoVal and assigns it to the Suspend field.
func (o *V0039PartitionInfoTimeouts) SetSuspend(v V0039Uint16NoVal) {
	o.Suspend = &v
}

func (o V0039PartitionInfoTimeouts) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0039PartitionInfoTimeouts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Suspend) {
		toSerialize["suspend"] = o.Suspend
	}
	return toSerialize, nil
}

type NullableV0039PartitionInfoTimeouts struct {
	value *V0039PartitionInfoTimeouts
	isSet bool
}

func (v NullableV0039PartitionInfoTimeouts) Get() *V0039PartitionInfoTimeouts {
	return v.value
}

func (v *NullableV0039PartitionInfoTimeouts) Set(val *V0039PartitionInfoTimeouts) {
	v.value = val
	v.isSet = true
}

func (v NullableV0039PartitionInfoTimeouts) IsSet() bool {
	return v.isSet
}

func (v *NullableV0039PartitionInfoTimeouts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0039PartitionInfoTimeouts(val *V0039PartitionInfoTimeouts) *NullableV0039PartitionInfoTimeouts {
	return &NullableV0039PartitionInfoTimeouts{value: val, isSet: true}
}

func (v NullableV0039PartitionInfoTimeouts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0039PartitionInfoTimeouts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
