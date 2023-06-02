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

// checks if the Dbv0038UserInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0038UserInfo{}

// Dbv0038UserInfo struct for Dbv0038UserInfo
type Dbv0038UserInfo struct {
	Meta *Dbv0038Meta `json:"meta,omitempty"`
	// Slurm errors
	Errors []Dbv0038Error `json:"errors,omitempty"`
	// Array of users
	Users []Dbv0038User `json:"users,omitempty"`
}

// NewDbv0038UserInfo instantiates a new Dbv0038UserInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0038UserInfo() *Dbv0038UserInfo {
	this := Dbv0038UserInfo{}
	return &this
}

// NewDbv0038UserInfoWithDefaults instantiates a new Dbv0038UserInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0038UserInfoWithDefaults() *Dbv0038UserInfo {
	this := Dbv0038UserInfo{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *Dbv0038UserInfo) GetMeta() Dbv0038Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Dbv0038Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038UserInfo) GetMetaOk() (*Dbv0038Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *Dbv0038UserInfo) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Dbv0038Meta and assigns it to the Meta field.
func (o *Dbv0038UserInfo) SetMeta(v Dbv0038Meta) {
	o.Meta = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *Dbv0038UserInfo) GetErrors() []Dbv0038Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Dbv0038Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038UserInfo) GetErrorsOk() ([]Dbv0038Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *Dbv0038UserInfo) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Dbv0038Error and assigns it to the Errors field.
func (o *Dbv0038UserInfo) SetErrors(v []Dbv0038Error) {
	o.Errors = v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *Dbv0038UserInfo) GetUsers() []Dbv0038User {
	if o == nil || IsNil(o.Users) {
		var ret []Dbv0038User
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038UserInfo) GetUsersOk() ([]Dbv0038User, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *Dbv0038UserInfo) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []Dbv0038User and assigns it to the Users field.
func (o *Dbv0038UserInfo) SetUsers(v []Dbv0038User) {
	o.Users = v
}

func (o Dbv0038UserInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0038UserInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}
	return toSerialize, nil
}

type NullableDbv0038UserInfo struct {
	value *Dbv0038UserInfo
	isSet bool
}

func (v NullableDbv0038UserInfo) Get() *Dbv0038UserInfo {
	return v.value
}

func (v *NullableDbv0038UserInfo) Set(val *Dbv0038UserInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0038UserInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0038UserInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0038UserInfo(val *Dbv0038UserInfo) *NullableDbv0038UserInfo {
	return &NullableDbv0038UserInfo{value: val, isSet: true}
}

func (v NullableDbv0038UserInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0038UserInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
