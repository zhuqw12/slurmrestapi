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

// checks if the Dbv0038ResponseWckeyDelete type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0038ResponseWckeyDelete{}

// Dbv0038ResponseWckeyDelete struct for Dbv0038ResponseWckeyDelete
type Dbv0038ResponseWckeyDelete struct {
	Meta *Dbv0038Meta `json:"meta,omitempty"`
	// Slurm errors
	Errors []Dbv0038Error `json:"errors,omitempty"`
}

// NewDbv0038ResponseWckeyDelete instantiates a new Dbv0038ResponseWckeyDelete object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0038ResponseWckeyDelete() *Dbv0038ResponseWckeyDelete {
	this := Dbv0038ResponseWckeyDelete{}
	return &this
}

// NewDbv0038ResponseWckeyDeleteWithDefaults instantiates a new Dbv0038ResponseWckeyDelete object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0038ResponseWckeyDeleteWithDefaults() *Dbv0038ResponseWckeyDelete {
	this := Dbv0038ResponseWckeyDelete{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *Dbv0038ResponseWckeyDelete) GetMeta() Dbv0038Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Dbv0038Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038ResponseWckeyDelete) GetMetaOk() (*Dbv0038Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *Dbv0038ResponseWckeyDelete) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Dbv0038Meta and assigns it to the Meta field.
func (o *Dbv0038ResponseWckeyDelete) SetMeta(v Dbv0038Meta) {
	o.Meta = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *Dbv0038ResponseWckeyDelete) GetErrors() []Dbv0038Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Dbv0038Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038ResponseWckeyDelete) GetErrorsOk() ([]Dbv0038Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *Dbv0038ResponseWckeyDelete) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Dbv0038Error and assigns it to the Errors field.
func (o *Dbv0038ResponseWckeyDelete) SetErrors(v []Dbv0038Error) {
	o.Errors = v
}

func (o Dbv0038ResponseWckeyDelete) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0038ResponseWckeyDelete) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableDbv0038ResponseWckeyDelete struct {
	value *Dbv0038ResponseWckeyDelete
	isSet bool
}

func (v NullableDbv0038ResponseWckeyDelete) Get() *Dbv0038ResponseWckeyDelete {
	return v.value
}

func (v *NullableDbv0038ResponseWckeyDelete) Set(val *Dbv0038ResponseWckeyDelete) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0038ResponseWckeyDelete) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0038ResponseWckeyDelete) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0038ResponseWckeyDelete(val *Dbv0038ResponseWckeyDelete) *NullableDbv0038ResponseWckeyDelete {
	return &NullableDbv0038ResponseWckeyDelete{value: val, isSet: true}
}

func (v NullableDbv0038ResponseWckeyDelete) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0038ResponseWckeyDelete) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
