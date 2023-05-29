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

// checks if the Dbv0038Diag type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0038Diag{}

// Dbv0038Diag struct for Dbv0038Diag
type Dbv0038Diag struct {
	Meta *Dbv0038Meta `json:"meta,omitempty"`
	// Slurm errors
	Errors     []Dbv0038Error         `json:"errors,omitempty"`
	Statistics *Dbv0038DiagStatistics `json:"statistics,omitempty"`
}

// NewDbv0038Diag instantiates a new Dbv0038Diag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0038Diag() *Dbv0038Diag {
	this := Dbv0038Diag{}
	return &this
}

// NewDbv0038DiagWithDefaults instantiates a new Dbv0038Diag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0038DiagWithDefaults() *Dbv0038Diag {
	this := Dbv0038Diag{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *Dbv0038Diag) GetMeta() Dbv0038Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Dbv0038Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038Diag) GetMetaOk() (*Dbv0038Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *Dbv0038Diag) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Dbv0038Meta and assigns it to the Meta field.
func (o *Dbv0038Diag) SetMeta(v Dbv0038Meta) {
	o.Meta = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *Dbv0038Diag) GetErrors() []Dbv0038Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Dbv0038Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038Diag) GetErrorsOk() ([]Dbv0038Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *Dbv0038Diag) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Dbv0038Error and assigns it to the Errors field.
func (o *Dbv0038Diag) SetErrors(v []Dbv0038Error) {
	o.Errors = v
}

// GetStatistics returns the Statistics field value if set, zero value otherwise.
func (o *Dbv0038Diag) GetStatistics() Dbv0038DiagStatistics {
	if o == nil || IsNil(o.Statistics) {
		var ret Dbv0038DiagStatistics
		return ret
	}
	return *o.Statistics
}

// GetStatisticsOk returns a tuple with the Statistics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038Diag) GetStatisticsOk() (*Dbv0038DiagStatistics, bool) {
	if o == nil || IsNil(o.Statistics) {
		return nil, false
	}
	return o.Statistics, true
}

// HasStatistics returns a boolean if a field has been set.
func (o *Dbv0038Diag) HasStatistics() bool {
	if o != nil && !IsNil(o.Statistics) {
		return true
	}

	return false
}

// SetStatistics gets a reference to the given Dbv0038DiagStatistics and assigns it to the Statistics field.
func (o *Dbv0038Diag) SetStatistics(v Dbv0038DiagStatistics) {
	o.Statistics = &v
}

func (o Dbv0038Diag) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0038Diag) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !IsNil(o.Statistics) {
		toSerialize["statistics"] = o.Statistics
	}
	return toSerialize, nil
}

type NullableDbv0038Diag struct {
	value *Dbv0038Diag
	isSet bool
}

func (v NullableDbv0038Diag) Get() *Dbv0038Diag {
	return v.value
}

func (v *NullableDbv0038Diag) Set(val *Dbv0038Diag) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0038Diag) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0038Diag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0038Diag(val *Dbv0038Diag) *NullableDbv0038Diag {
	return &NullableDbv0038Diag{value: val, isSet: true}
}

func (v NullableDbv0038Diag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0038Diag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}