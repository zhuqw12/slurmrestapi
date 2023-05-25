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

// checks if the V0039QosLimits type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0039QosLimits{}

// V0039QosLimits struct for V0039QosLimits
type V0039QosLimits struct {
	Min *V0039QosLimitsMin `json:"min,omitempty"`
}

// NewV0039QosLimits instantiates a new V0039QosLimits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0039QosLimits() *V0039QosLimits {
	this := V0039QosLimits{}
	return &this
}

// NewV0039QosLimitsWithDefaults instantiates a new V0039QosLimits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0039QosLimitsWithDefaults() *V0039QosLimits {
	this := V0039QosLimits{}
	return &this
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *V0039QosLimits) GetMin() V0039QosLimitsMin {
	if o == nil || IsNil(o.Min) {
		var ret V0039QosLimitsMin
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039QosLimits) GetMinOk() (*V0039QosLimitsMin, bool) {
	if o == nil || IsNil(o.Min) {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *V0039QosLimits) HasMin() bool {
	if o != nil && !IsNil(o.Min) {
		return true
	}

	return false
}

// SetMin gets a reference to the given V0039QosLimitsMin and assigns it to the Min field.
func (o *V0039QosLimits) SetMin(v V0039QosLimitsMin) {
	o.Min = &v
}

func (o V0039QosLimits) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0039QosLimits) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Min) {
		toSerialize["min"] = o.Min
	}
	return toSerialize, nil
}

type NullableV0039QosLimits struct {
	value *V0039QosLimits
	isSet bool
}

func (v NullableV0039QosLimits) Get() *V0039QosLimits {
	return v.value
}

func (v *NullableV0039QosLimits) Set(val *V0039QosLimits) {
	v.value = val
	v.isSet = true
}

func (v NullableV0039QosLimits) IsSet() bool {
	return v.isSet
}

func (v *NullableV0039QosLimits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0039QosLimits(val *V0039QosLimits) *NullableV0039QosLimits {
	return &NullableV0039QosLimits{value: val, isSet: true}
}

func (v NullableV0039QosLimits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0039QosLimits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
