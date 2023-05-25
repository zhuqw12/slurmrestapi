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

// checks if the Dbv0037AssociationsInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0037AssociationsInfo{}

// Dbv0037AssociationsInfo struct for Dbv0037AssociationsInfo
type Dbv0037AssociationsInfo struct {
	// Slurm errors
	Errors []Dbv0037Error `json:"errors,omitempty"`
	// Array of associations
	Associations []Dbv0037Association `json:"associations,omitempty"`
}

// NewDbv0037AssociationsInfo instantiates a new Dbv0037AssociationsInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0037AssociationsInfo() *Dbv0037AssociationsInfo {
	this := Dbv0037AssociationsInfo{}
	return &this
}

// NewDbv0037AssociationsInfoWithDefaults instantiates a new Dbv0037AssociationsInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0037AssociationsInfoWithDefaults() *Dbv0037AssociationsInfo {
	this := Dbv0037AssociationsInfo{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *Dbv0037AssociationsInfo) GetErrors() []Dbv0037Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Dbv0037Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037AssociationsInfo) GetErrorsOk() ([]Dbv0037Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *Dbv0037AssociationsInfo) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Dbv0037Error and assigns it to the Errors field.
func (o *Dbv0037AssociationsInfo) SetErrors(v []Dbv0037Error) {
	o.Errors = v
}

// GetAssociations returns the Associations field value if set, zero value otherwise.
func (o *Dbv0037AssociationsInfo) GetAssociations() []Dbv0037Association {
	if o == nil || IsNil(o.Associations) {
		var ret []Dbv0037Association
		return ret
	}
	return o.Associations
}

// GetAssociationsOk returns a tuple with the Associations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037AssociationsInfo) GetAssociationsOk() ([]Dbv0037Association, bool) {
	if o == nil || IsNil(o.Associations) {
		return nil, false
	}
	return o.Associations, true
}

// HasAssociations returns a boolean if a field has been set.
func (o *Dbv0037AssociationsInfo) HasAssociations() bool {
	if o != nil && !IsNil(o.Associations) {
		return true
	}

	return false
}

// SetAssociations gets a reference to the given []Dbv0037Association and assigns it to the Associations field.
func (o *Dbv0037AssociationsInfo) SetAssociations(v []Dbv0037Association) {
	o.Associations = v
}

func (o Dbv0037AssociationsInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0037AssociationsInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !IsNil(o.Associations) {
		toSerialize["associations"] = o.Associations
	}
	return toSerialize, nil
}

type NullableDbv0037AssociationsInfo struct {
	value *Dbv0037AssociationsInfo
	isSet bool
}

func (v NullableDbv0037AssociationsInfo) Get() *Dbv0037AssociationsInfo {
	return v.value
}

func (v *NullableDbv0037AssociationsInfo) Set(val *Dbv0037AssociationsInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0037AssociationsInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0037AssociationsInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0037AssociationsInfo(val *Dbv0037AssociationsInfo) *NullableDbv0037AssociationsInfo {
	return &NullableDbv0037AssociationsInfo{value: val, isSet: true}
}

func (v NullableDbv0037AssociationsInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0037AssociationsInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
