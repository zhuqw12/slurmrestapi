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

// checks if the Dbv0039MetaSlurmVersion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0039MetaSlurmVersion{}

// Dbv0039MetaSlurmVersion struct for Dbv0039MetaSlurmVersion
type Dbv0039MetaSlurmVersion struct {
	//
	Major *int32 `json:"major,omitempty"`
	//
	Micro *int32 `json:"micro,omitempty"`
	//
	Minor *int32 `json:"minor,omitempty"`
}

// NewDbv0039MetaSlurmVersion instantiates a new Dbv0039MetaSlurmVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0039MetaSlurmVersion() *Dbv0039MetaSlurmVersion {
	this := Dbv0039MetaSlurmVersion{}
	return &this
}

// NewDbv0039MetaSlurmVersionWithDefaults instantiates a new Dbv0039MetaSlurmVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0039MetaSlurmVersionWithDefaults() *Dbv0039MetaSlurmVersion {
	this := Dbv0039MetaSlurmVersion{}
	return &this
}

// GetMajor returns the Major field value if set, zero value otherwise.
func (o *Dbv0039MetaSlurmVersion) GetMajor() int32 {
	if o == nil || IsNil(o.Major) {
		var ret int32
		return ret
	}
	return *o.Major
}

// GetMajorOk returns a tuple with the Major field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0039MetaSlurmVersion) GetMajorOk() (*int32, bool) {
	if o == nil || IsNil(o.Major) {
		return nil, false
	}
	return o.Major, true
}

// HasMajor returns a boolean if a field has been set.
func (o *Dbv0039MetaSlurmVersion) HasMajor() bool {
	if o != nil && !IsNil(o.Major) {
		return true
	}

	return false
}

// SetMajor gets a reference to the given int32 and assigns it to the Major field.
func (o *Dbv0039MetaSlurmVersion) SetMajor(v int32) {
	o.Major = &v
}

// GetMicro returns the Micro field value if set, zero value otherwise.
func (o *Dbv0039MetaSlurmVersion) GetMicro() int32 {
	if o == nil || IsNil(o.Micro) {
		var ret int32
		return ret
	}
	return *o.Micro
}

// GetMicroOk returns a tuple with the Micro field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0039MetaSlurmVersion) GetMicroOk() (*int32, bool) {
	if o == nil || IsNil(o.Micro) {
		return nil, false
	}
	return o.Micro, true
}

// HasMicro returns a boolean if a field has been set.
func (o *Dbv0039MetaSlurmVersion) HasMicro() bool {
	if o != nil && !IsNil(o.Micro) {
		return true
	}

	return false
}

// SetMicro gets a reference to the given int32 and assigns it to the Micro field.
func (o *Dbv0039MetaSlurmVersion) SetMicro(v int32) {
	o.Micro = &v
}

// GetMinor returns the Minor field value if set, zero value otherwise.
func (o *Dbv0039MetaSlurmVersion) GetMinor() int32 {
	if o == nil || IsNil(o.Minor) {
		var ret int32
		return ret
	}
	return *o.Minor
}

// GetMinorOk returns a tuple with the Minor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0039MetaSlurmVersion) GetMinorOk() (*int32, bool) {
	if o == nil || IsNil(o.Minor) {
		return nil, false
	}
	return o.Minor, true
}

// HasMinor returns a boolean if a field has been set.
func (o *Dbv0039MetaSlurmVersion) HasMinor() bool {
	if o != nil && !IsNil(o.Minor) {
		return true
	}

	return false
}

// SetMinor gets a reference to the given int32 and assigns it to the Minor field.
func (o *Dbv0039MetaSlurmVersion) SetMinor(v int32) {
	o.Minor = &v
}

func (o Dbv0039MetaSlurmVersion) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0039MetaSlurmVersion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Major) {
		toSerialize["major"] = o.Major
	}
	if !IsNil(o.Micro) {
		toSerialize["micro"] = o.Micro
	}
	if !IsNil(o.Minor) {
		toSerialize["minor"] = o.Minor
	}
	return toSerialize, nil
}

type NullableDbv0039MetaSlurmVersion struct {
	value *Dbv0039MetaSlurmVersion
	isSet bool
}

func (v NullableDbv0039MetaSlurmVersion) Get() *Dbv0039MetaSlurmVersion {
	return v.value
}

func (v *NullableDbv0039MetaSlurmVersion) Set(val *Dbv0039MetaSlurmVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0039MetaSlurmVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0039MetaSlurmVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0039MetaSlurmVersion(val *Dbv0039MetaSlurmVersion) *NullableDbv0039MetaSlurmVersion {
	return &NullableDbv0039MetaSlurmVersion{value: val, isSet: true}
}

func (v NullableDbv0039MetaSlurmVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0039MetaSlurmVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
