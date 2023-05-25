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

// checks if the V0039JobTres type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0039JobTres{}

// V0039JobTres struct for V0039JobTres
type V0039JobTres struct {
	Requested []V0039Tres `json:"requested,omitempty"`
}

// NewV0039JobTres instantiates a new V0039JobTres object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0039JobTres() *V0039JobTres {
	this := V0039JobTres{}
	return &this
}

// NewV0039JobTresWithDefaults instantiates a new V0039JobTres object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0039JobTresWithDefaults() *V0039JobTres {
	this := V0039JobTres{}
	return &this
}

// GetRequested returns the Requested field value if set, zero value otherwise.
func (o *V0039JobTres) GetRequested() []V0039Tres {
	if o == nil || IsNil(o.Requested) {
		var ret []V0039Tres
		return ret
	}
	return o.Requested
}

// GetRequestedOk returns a tuple with the Requested field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039JobTres) GetRequestedOk() ([]V0039Tres, bool) {
	if o == nil || IsNil(o.Requested) {
		return nil, false
	}
	return o.Requested, true
}

// HasRequested returns a boolean if a field has been set.
func (o *V0039JobTres) HasRequested() bool {
	if o != nil && !IsNil(o.Requested) {
		return true
	}

	return false
}

// SetRequested gets a reference to the given []V0039Tres and assigns it to the Requested field.
func (o *V0039JobTres) SetRequested(v []V0039Tres) {
	o.Requested = v
}

func (o V0039JobTres) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0039JobTres) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Requested) {
		toSerialize["requested"] = o.Requested
	}
	return toSerialize, nil
}

type NullableV0039JobTres struct {
	value *V0039JobTres
	isSet bool
}

func (v NullableV0039JobTres) Get() *V0039JobTres {
	return v.value
}

func (v *NullableV0039JobTres) Set(val *V0039JobTres) {
	v.value = val
	v.isSet = true
}

func (v NullableV0039JobTres) IsSet() bool {
	return v.isSet
}

func (v *NullableV0039JobTres) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0039JobTres(val *V0039JobTres) *NullableV0039JobTres {
	return &NullableV0039JobTres{value: val, isSet: true}
}

func (v NullableV0039JobTres) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0039JobTres) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
