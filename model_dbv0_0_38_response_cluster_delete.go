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

// checks if the Dbv0038ResponseClusterDelete type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0038ResponseClusterDelete{}

// Dbv0038ResponseClusterDelete struct for Dbv0038ResponseClusterDelete
type Dbv0038ResponseClusterDelete struct {
	Meta *Dbv0038Meta `json:"meta,omitempty"`
	// Slurm errors
	Errors []Dbv0038Error `json:"errors,omitempty"`
}

// NewDbv0038ResponseClusterDelete instantiates a new Dbv0038ResponseClusterDelete object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0038ResponseClusterDelete() *Dbv0038ResponseClusterDelete {
	this := Dbv0038ResponseClusterDelete{}
	return &this
}

// NewDbv0038ResponseClusterDeleteWithDefaults instantiates a new Dbv0038ResponseClusterDelete object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0038ResponseClusterDeleteWithDefaults() *Dbv0038ResponseClusterDelete {
	this := Dbv0038ResponseClusterDelete{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *Dbv0038ResponseClusterDelete) GetMeta() Dbv0038Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Dbv0038Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038ResponseClusterDelete) GetMetaOk() (*Dbv0038Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *Dbv0038ResponseClusterDelete) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Dbv0038Meta and assigns it to the Meta field.
func (o *Dbv0038ResponseClusterDelete) SetMeta(v Dbv0038Meta) {
	o.Meta = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *Dbv0038ResponseClusterDelete) GetErrors() []Dbv0038Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Dbv0038Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038ResponseClusterDelete) GetErrorsOk() ([]Dbv0038Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *Dbv0038ResponseClusterDelete) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Dbv0038Error and assigns it to the Errors field.
func (o *Dbv0038ResponseClusterDelete) SetErrors(v []Dbv0038Error) {
	o.Errors = v
}

func (o Dbv0038ResponseClusterDelete) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0038ResponseClusterDelete) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableDbv0038ResponseClusterDelete struct {
	value *Dbv0038ResponseClusterDelete
	isSet bool
}

func (v NullableDbv0038ResponseClusterDelete) Get() *Dbv0038ResponseClusterDelete {
	return v.value
}

func (v *NullableDbv0038ResponseClusterDelete) Set(val *Dbv0038ResponseClusterDelete) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0038ResponseClusterDelete) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0038ResponseClusterDelete) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0038ResponseClusterDelete(val *Dbv0038ResponseClusterDelete) *NullableDbv0038ResponseClusterDelete {
	return &NullableDbv0038ResponseClusterDelete{value: val, isSet: true}
}

func (v NullableDbv0038ResponseClusterDelete) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0038ResponseClusterDelete) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
