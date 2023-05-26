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

// checks if the Dbv0039Diag type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0039Diag{}

// Dbv0039Diag struct for Dbv0039Diag
type Dbv0039Diag struct {
	Meta *Dbv0039Meta `json:"meta,omitempty"`
	// Slurm errors
	Errors []Dbv0039Error `json:"errors,omitempty"`
	// Slurm warnings
	Warnings   []Dbv0039Warning `json:"warnings,omitempty"`
	Statistics *V0039StatsRec   `json:"statistics,omitempty"`
}

// NewDbv0039Diag instantiates a new Dbv0039Diag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0039Diag() *Dbv0039Diag {
	this := Dbv0039Diag{}
	return &this
}

// NewDbv0039DiagWithDefaults instantiates a new Dbv0039Diag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0039DiagWithDefaults() *Dbv0039Diag {
	this := Dbv0039Diag{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *Dbv0039Diag) GetMeta() Dbv0039Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Dbv0039Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0039Diag) GetMetaOk() (*Dbv0039Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *Dbv0039Diag) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Dbv0039Meta and assigns it to the Meta field.
func (o *Dbv0039Diag) SetMeta(v Dbv0039Meta) {
	o.Meta = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *Dbv0039Diag) GetErrors() []Dbv0039Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Dbv0039Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0039Diag) GetErrorsOk() ([]Dbv0039Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *Dbv0039Diag) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Dbv0039Error and assigns it to the Errors field.
func (o *Dbv0039Diag) SetErrors(v []Dbv0039Error) {
	o.Errors = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *Dbv0039Diag) GetWarnings() []Dbv0039Warning {
	if o == nil || IsNil(o.Warnings) {
		var ret []Dbv0039Warning
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0039Diag) GetWarningsOk() ([]Dbv0039Warning, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *Dbv0039Diag) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []Dbv0039Warning and assigns it to the Warnings field.
func (o *Dbv0039Diag) SetWarnings(v []Dbv0039Warning) {
	o.Warnings = v
}

// GetStatistics returns the Statistics field value if set, zero value otherwise.
func (o *Dbv0039Diag) GetStatistics() V0039StatsRec {
	if o == nil || IsNil(o.Statistics) {
		var ret V0039StatsRec
		return ret
	}
	return *o.Statistics
}

// GetStatisticsOk returns a tuple with the Statistics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0039Diag) GetStatisticsOk() (*V0039StatsRec, bool) {
	if o == nil || IsNil(o.Statistics) {
		return nil, false
	}
	return o.Statistics, true
}

// HasStatistics returns a boolean if a field has been set.
func (o *Dbv0039Diag) HasStatistics() bool {
	if o != nil && !IsNil(o.Statistics) {
		return true
	}

	return false
}

// SetStatistics gets a reference to the given V0039StatsRec and assigns it to the Statistics field.
func (o *Dbv0039Diag) SetStatistics(v V0039StatsRec) {
	o.Statistics = &v
}

func (o Dbv0039Diag) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0039Diag) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	if !IsNil(o.Statistics) {
		toSerialize["statistics"] = o.Statistics
	}
	return toSerialize, nil
}

type NullableDbv0039Diag struct {
	value *Dbv0039Diag
	isSet bool
}

func (v NullableDbv0039Diag) Get() *Dbv0039Diag {
	return v.value
}

func (v *NullableDbv0039Diag) Set(val *Dbv0039Diag) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0039Diag) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0039Diag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0039Diag(val *Dbv0039Diag) *NullableDbv0039Diag {
	return &NullableDbv0039Diag{value: val, isSet: true}
}

func (v NullableDbv0039Diag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0039Diag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
