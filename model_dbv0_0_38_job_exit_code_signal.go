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

// checks if the Dbv0038JobExitCodeSignal type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0038JobExitCodeSignal{}

// Dbv0038JobExitCodeSignal Signal details (if signaled)
type Dbv0038JobExitCodeSignal struct {
	// Signal number process received
	SignalId *int32 `json:"signal_id,omitempty"`
	// Name of signal received
	Name *string `json:"name,omitempty"`
}

// NewDbv0038JobExitCodeSignal instantiates a new Dbv0038JobExitCodeSignal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0038JobExitCodeSignal() *Dbv0038JobExitCodeSignal {
	this := Dbv0038JobExitCodeSignal{}
	return &this
}

// NewDbv0038JobExitCodeSignalWithDefaults instantiates a new Dbv0038JobExitCodeSignal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0038JobExitCodeSignalWithDefaults() *Dbv0038JobExitCodeSignal {
	this := Dbv0038JobExitCodeSignal{}
	return &this
}

// GetSignalId returns the SignalId field value if set, zero value otherwise.
func (o *Dbv0038JobExitCodeSignal) GetSignalId() int32 {
	if o == nil || IsNil(o.SignalId) {
		var ret int32
		return ret
	}
	return *o.SignalId
}

// GetSignalIdOk returns a tuple with the SignalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038JobExitCodeSignal) GetSignalIdOk() (*int32, bool) {
	if o == nil || IsNil(o.SignalId) {
		return nil, false
	}
	return o.SignalId, true
}

// HasSignalId returns a boolean if a field has been set.
func (o *Dbv0038JobExitCodeSignal) HasSignalId() bool {
	if o != nil && !IsNil(o.SignalId) {
		return true
	}

	return false
}

// SetSignalId gets a reference to the given int32 and assigns it to the SignalId field.
func (o *Dbv0038JobExitCodeSignal) SetSignalId(v int32) {
	o.SignalId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Dbv0038JobExitCodeSignal) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038JobExitCodeSignal) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Dbv0038JobExitCodeSignal) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Dbv0038JobExitCodeSignal) SetName(v string) {
	o.Name = &v
}

func (o Dbv0038JobExitCodeSignal) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0038JobExitCodeSignal) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SignalId) {
		toSerialize["signal_id"] = o.SignalId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableDbv0038JobExitCodeSignal struct {
	value *Dbv0038JobExitCodeSignal
	isSet bool
}

func (v NullableDbv0038JobExitCodeSignal) Get() *Dbv0038JobExitCodeSignal {
	return v.value
}

func (v *NullableDbv0038JobExitCodeSignal) Set(val *Dbv0038JobExitCodeSignal) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0038JobExitCodeSignal) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0038JobExitCodeSignal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0038JobExitCodeSignal(val *Dbv0038JobExitCodeSignal) *NullableDbv0038JobExitCodeSignal {
	return &NullableDbv0038JobExitCodeSignal{value: val, isSet: true}
}

func (v NullableDbv0038JobExitCodeSignal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0038JobExitCodeSignal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
