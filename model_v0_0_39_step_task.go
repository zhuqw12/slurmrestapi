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

// checks if the V0039StepTask type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0039StepTask{}

// V0039StepTask struct for V0039StepTask
type V0039StepTask struct {
	Distribution *string `json:"distribution,omitempty"`
}

// NewV0039StepTask instantiates a new V0039StepTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0039StepTask() *V0039StepTask {
	this := V0039StepTask{}
	return &this
}

// NewV0039StepTaskWithDefaults instantiates a new V0039StepTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0039StepTaskWithDefaults() *V0039StepTask {
	this := V0039StepTask{}
	return &this
}

// GetDistribution returns the Distribution field value if set, zero value otherwise.
func (o *V0039StepTask) GetDistribution() string {
	if o == nil || IsNil(o.Distribution) {
		var ret string
		return ret
	}
	return *o.Distribution
}

// GetDistributionOk returns a tuple with the Distribution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039StepTask) GetDistributionOk() (*string, bool) {
	if o == nil || IsNil(o.Distribution) {
		return nil, false
	}
	return o.Distribution, true
}

// HasDistribution returns a boolean if a field has been set.
func (o *V0039StepTask) HasDistribution() bool {
	if o != nil && !IsNil(o.Distribution) {
		return true
	}

	return false
}

// SetDistribution gets a reference to the given string and assigns it to the Distribution field.
func (o *V0039StepTask) SetDistribution(v string) {
	o.Distribution = &v
}

func (o V0039StepTask) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0039StepTask) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Distribution) {
		toSerialize["distribution"] = o.Distribution
	}
	return toSerialize, nil
}

type NullableV0039StepTask struct {
	value *V0039StepTask
	isSet bool
}

func (v NullableV0039StepTask) Get() *V0039StepTask {
	return v.value
}

func (v *NullableV0039StepTask) Set(val *V0039StepTask) {
	v.value = val
	v.isSet = true
}

func (v NullableV0039StepTask) IsSet() bool {
	return v.isSet
}

func (v *NullableV0039StepTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0039StepTask(val *V0039StepTask) *NullableV0039StepTask {
	return &NullableV0039StepTask{value: val, isSet: true}
}

func (v NullableV0039StepTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0039StepTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
