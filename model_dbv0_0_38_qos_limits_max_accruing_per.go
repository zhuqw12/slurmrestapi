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

// checks if the Dbv0038QosLimitsMaxAccruingPer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0038QosLimitsMaxAccruingPer{}

// Dbv0038QosLimitsMaxAccruingPer Max accuring priority per setting
type Dbv0038QosLimitsMaxAccruingPer struct {
	// Max accuring priority per account
	Account *int32 `json:"account,omitempty"`
	// Max accuring priority per user
	User *int32 `json:"user,omitempty"`
}

// NewDbv0038QosLimitsMaxAccruingPer instantiates a new Dbv0038QosLimitsMaxAccruingPer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0038QosLimitsMaxAccruingPer() *Dbv0038QosLimitsMaxAccruingPer {
	this := Dbv0038QosLimitsMaxAccruingPer{}
	return &this
}

// NewDbv0038QosLimitsMaxAccruingPerWithDefaults instantiates a new Dbv0038QosLimitsMaxAccruingPer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0038QosLimitsMaxAccruingPerWithDefaults() *Dbv0038QosLimitsMaxAccruingPer {
	this := Dbv0038QosLimitsMaxAccruingPer{}
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *Dbv0038QosLimitsMaxAccruingPer) GetAccount() int32 {
	if o == nil || IsNil(o.Account) {
		var ret int32
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038QosLimitsMaxAccruingPer) GetAccountOk() (*int32, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *Dbv0038QosLimitsMaxAccruingPer) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given int32 and assigns it to the Account field.
func (o *Dbv0038QosLimitsMaxAccruingPer) SetAccount(v int32) {
	o.Account = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *Dbv0038QosLimitsMaxAccruingPer) GetUser() int32 {
	if o == nil || IsNil(o.User) {
		var ret int32
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038QosLimitsMaxAccruingPer) GetUserOk() (*int32, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *Dbv0038QosLimitsMaxAccruingPer) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given int32 and assigns it to the User field.
func (o *Dbv0038QosLimitsMaxAccruingPer) SetUser(v int32) {
	o.User = &v
}

func (o Dbv0038QosLimitsMaxAccruingPer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0038QosLimitsMaxAccruingPer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	return toSerialize, nil
}

type NullableDbv0038QosLimitsMaxAccruingPer struct {
	value *Dbv0038QosLimitsMaxAccruingPer
	isSet bool
}

func (v NullableDbv0038QosLimitsMaxAccruingPer) Get() *Dbv0038QosLimitsMaxAccruingPer {
	return v.value
}

func (v *NullableDbv0038QosLimitsMaxAccruingPer) Set(val *Dbv0038QosLimitsMaxAccruingPer) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0038QosLimitsMaxAccruingPer) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0038QosLimitsMaxAccruingPer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0038QosLimitsMaxAccruingPer(val *Dbv0038QosLimitsMaxAccruingPer) *NullableDbv0038QosLimitsMaxAccruingPer {
	return &NullableDbv0038QosLimitsMaxAccruingPer{value: val, isSet: true}
}

func (v NullableDbv0038QosLimitsMaxAccruingPer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0038QosLimitsMaxAccruingPer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
