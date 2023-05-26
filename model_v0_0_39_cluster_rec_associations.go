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

// checks if the V0039ClusterRecAssociations type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0039ClusterRecAssociations{}

// V0039ClusterRecAssociations struct for V0039ClusterRecAssociations
type V0039ClusterRecAssociations struct {
	Root *V0039AssocShort `json:"root,omitempty"`
}

// NewV0039ClusterRecAssociations instantiates a new V0039ClusterRecAssociations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0039ClusterRecAssociations() *V0039ClusterRecAssociations {
	this := V0039ClusterRecAssociations{}
	return &this
}

// NewV0039ClusterRecAssociationsWithDefaults instantiates a new V0039ClusterRecAssociations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0039ClusterRecAssociationsWithDefaults() *V0039ClusterRecAssociations {
	this := V0039ClusterRecAssociations{}
	return &this
}

// GetRoot returns the Root field value if set, zero value otherwise.
func (o *V0039ClusterRecAssociations) GetRoot() V0039AssocShort {
	if o == nil || IsNil(o.Root) {
		var ret V0039AssocShort
		return ret
	}
	return *o.Root
}

// GetRootOk returns a tuple with the Root field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039ClusterRecAssociations) GetRootOk() (*V0039AssocShort, bool) {
	if o == nil || IsNil(o.Root) {
		return nil, false
	}
	return o.Root, true
}

// HasRoot returns a boolean if a field has been set.
func (o *V0039ClusterRecAssociations) HasRoot() bool {
	if o != nil && !IsNil(o.Root) {
		return true
	}

	return false
}

// SetRoot gets a reference to the given V0039AssocShort and assigns it to the Root field.
func (o *V0039ClusterRecAssociations) SetRoot(v V0039AssocShort) {
	o.Root = &v
}

func (o V0039ClusterRecAssociations) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0039ClusterRecAssociations) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Root) {
		toSerialize["root"] = o.Root
	}
	return toSerialize, nil
}

type NullableV0039ClusterRecAssociations struct {
	value *V0039ClusterRecAssociations
	isSet bool
}

func (v NullableV0039ClusterRecAssociations) Get() *V0039ClusterRecAssociations {
	return v.value
}

func (v *NullableV0039ClusterRecAssociations) Set(val *V0039ClusterRecAssociations) {
	v.value = val
	v.isSet = true
}

func (v NullableV0039ClusterRecAssociations) IsSet() bool {
	return v.isSet
}

func (v *NullableV0039ClusterRecAssociations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0039ClusterRecAssociations(val *V0039ClusterRecAssociations) *NullableV0039ClusterRecAssociations {
	return &NullableV0039ClusterRecAssociations{value: val, isSet: true}
}

func (v NullableV0039ClusterRecAssociations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0039ClusterRecAssociations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
