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

// checks if the V0038Meta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0038Meta{}

// V0038Meta struct for V0038Meta
type V0038Meta struct {
	Plugin *Dbv0038MetaPlugin `json:"plugin,omitempty"`
	Slurm  *Dbv0038MetaSlurm  `json:"Slurm,omitempty"`
}

// NewV0038Meta instantiates a new V0038Meta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0038Meta() *V0038Meta {
	this := V0038Meta{}
	return &this
}

// NewV0038MetaWithDefaults instantiates a new V0038Meta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0038MetaWithDefaults() *V0038Meta {
	this := V0038Meta{}
	return &this
}

// GetPlugin returns the Plugin field value if set, zero value otherwise.
func (o *V0038Meta) GetPlugin() Dbv0038MetaPlugin {
	if o == nil || IsNil(o.Plugin) {
		var ret Dbv0038MetaPlugin
		return ret
	}
	return *o.Plugin
}

// GetPluginOk returns a tuple with the Plugin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0038Meta) GetPluginOk() (*Dbv0038MetaPlugin, bool) {
	if o == nil || IsNil(o.Plugin) {
		return nil, false
	}
	return o.Plugin, true
}

// HasPlugin returns a boolean if a field has been set.
func (o *V0038Meta) HasPlugin() bool {
	if o != nil && !IsNil(o.Plugin) {
		return true
	}

	return false
}

// SetPlugin gets a reference to the given Dbv0038MetaPlugin and assigns it to the Plugin field.
func (o *V0038Meta) SetPlugin(v Dbv0038MetaPlugin) {
	o.Plugin = &v
}

// GetSlurm returns the Slurm field value if set, zero value otherwise.
func (o *V0038Meta) GetSlurm() Dbv0038MetaSlurm {
	if o == nil || IsNil(o.Slurm) {
		var ret Dbv0038MetaSlurm
		return ret
	}
	return *o.Slurm
}

// GetSlurmOk returns a tuple with the Slurm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0038Meta) GetSlurmOk() (*Dbv0038MetaSlurm, bool) {
	if o == nil || IsNil(o.Slurm) {
		return nil, false
	}
	return o.Slurm, true
}

// HasSlurm returns a boolean if a field has been set.
func (o *V0038Meta) HasSlurm() bool {
	if o != nil && !IsNil(o.Slurm) {
		return true
	}

	return false
}

// SetSlurm gets a reference to the given Dbv0038MetaSlurm and assigns it to the Slurm field.
func (o *V0038Meta) SetSlurm(v Dbv0038MetaSlurm) {
	o.Slurm = &v
}

func (o V0038Meta) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0038Meta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Plugin) {
		toSerialize["plugin"] = o.Plugin
	}
	if !IsNil(o.Slurm) {
		toSerialize["Slurm"] = o.Slurm
	}
	return toSerialize, nil
}

type NullableV0038Meta struct {
	value *V0038Meta
	isSet bool
}

func (v NullableV0038Meta) Get() *V0038Meta {
	return v.value
}

func (v *NullableV0038Meta) Set(val *V0038Meta) {
	v.value = val
	v.isSet = true
}

func (v NullableV0038Meta) IsSet() bool {
	return v.isSet
}

func (v *NullableV0038Meta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0038Meta(val *V0038Meta) *NullableV0038Meta {
	return &NullableV0038Meta{value: val, isSet: true}
}

func (v NullableV0038Meta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0038Meta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
