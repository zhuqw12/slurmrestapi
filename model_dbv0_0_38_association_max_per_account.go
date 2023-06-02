/*
Slurm Rest API

API to access and control Slurm.

API version: v0.0.38
Contact: sales@schedmd.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package slurmrestapi

import (
	"encoding/json"
)

// checks if the Dbv0038AssociationMaxPerAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0038AssociationMaxPerAccount{}

// Dbv0038AssociationMaxPerAccount Max per accounting settings
type Dbv0038AssociationMaxPerAccount struct {
	// Max wallclock per account
	WallClock *int32 `json:"wall_clock,omitempty"`
}

// NewDbv0038AssociationMaxPerAccount instantiates a new Dbv0038AssociationMaxPerAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0038AssociationMaxPerAccount() *Dbv0038AssociationMaxPerAccount {
	this := Dbv0038AssociationMaxPerAccount{}
	return &this
}

// NewDbv0038AssociationMaxPerAccountWithDefaults instantiates a new Dbv0038AssociationMaxPerAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0038AssociationMaxPerAccountWithDefaults() *Dbv0038AssociationMaxPerAccount {
	this := Dbv0038AssociationMaxPerAccount{}
	return &this
}

// GetWallClock returns the WallClock field value if set, zero value otherwise.
func (o *Dbv0038AssociationMaxPerAccount) GetWallClock() int32 {
	if o == nil || IsNil(o.WallClock) {
		var ret int32
		return ret
	}
	return *o.WallClock
}

// GetWallClockOk returns a tuple with the WallClock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038AssociationMaxPerAccount) GetWallClockOk() (*int32, bool) {
	if o == nil || IsNil(o.WallClock) {
		return nil, false
	}
	return o.WallClock, true
}

// HasWallClock returns a boolean if a field has been set.
func (o *Dbv0038AssociationMaxPerAccount) HasWallClock() bool {
	if o != nil && !IsNil(o.WallClock) {
		return true
	}

	return false
}

// SetWallClock gets a reference to the given int32 and assigns it to the WallClock field.
func (o *Dbv0038AssociationMaxPerAccount) SetWallClock(v int32) {
	o.WallClock = &v
}

func (o Dbv0038AssociationMaxPerAccount) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0038AssociationMaxPerAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.WallClock) {
		toSerialize["wall_clock"] = o.WallClock
	}
	return toSerialize, nil
}

type NullableDbv0038AssociationMaxPerAccount struct {
	value *Dbv0038AssociationMaxPerAccount
	isSet bool
}

func (v NullableDbv0038AssociationMaxPerAccount) Get() *Dbv0038AssociationMaxPerAccount {
	return v.value
}

func (v *NullableDbv0038AssociationMaxPerAccount) Set(val *Dbv0038AssociationMaxPerAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0038AssociationMaxPerAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0038AssociationMaxPerAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0038AssociationMaxPerAccount(val *Dbv0038AssociationMaxPerAccount) *NullableDbv0038AssociationMaxPerAccount {
	return &NullableDbv0038AssociationMaxPerAccount{value: val, isSet: true}
}

func (v NullableDbv0038AssociationMaxPerAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0038AssociationMaxPerAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
