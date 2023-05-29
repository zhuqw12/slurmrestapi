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

// checks if the Dbv0037ResponseUserDelete type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0037ResponseUserDelete{}

// Dbv0037ResponseUserDelete struct for Dbv0037ResponseUserDelete
type Dbv0037ResponseUserDelete struct {
	// Slurm errors
	Errors []Dbv0037Error `json:"errors,omitempty"`
}

// NewDbv0037ResponseUserDelete instantiates a new Dbv0037ResponseUserDelete object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0037ResponseUserDelete() *Dbv0037ResponseUserDelete {
	this := Dbv0037ResponseUserDelete{}
	return &this
}

// NewDbv0037ResponseUserDeleteWithDefaults instantiates a new Dbv0037ResponseUserDelete object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0037ResponseUserDeleteWithDefaults() *Dbv0037ResponseUserDelete {
	this := Dbv0037ResponseUserDelete{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *Dbv0037ResponseUserDelete) GetErrors() []Dbv0037Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Dbv0037Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037ResponseUserDelete) GetErrorsOk() ([]Dbv0037Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *Dbv0037ResponseUserDelete) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Dbv0037Error and assigns it to the Errors field.
func (o *Dbv0037ResponseUserDelete) SetErrors(v []Dbv0037Error) {
	o.Errors = v
}

func (o Dbv0037ResponseUserDelete) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0037ResponseUserDelete) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableDbv0037ResponseUserDelete struct {
	value *Dbv0037ResponseUserDelete
	isSet bool
}

func (v NullableDbv0037ResponseUserDelete) Get() *Dbv0037ResponseUserDelete {
	return v.value
}

func (v *NullableDbv0037ResponseUserDelete) Set(val *Dbv0037ResponseUserDelete) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0037ResponseUserDelete) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0037ResponseUserDelete) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0037ResponseUserDelete(val *Dbv0037ResponseUserDelete) *NullableDbv0037ResponseUserDelete {
	return &NullableDbv0037ResponseUserDelete{value: val, isSet: true}
}

func (v NullableDbv0037ResponseUserDelete) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0037ResponseUserDelete) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}