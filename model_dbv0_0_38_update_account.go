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

// checks if the Dbv0038UpdateAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0038UpdateAccount{}

// Dbv0038UpdateAccount struct for Dbv0038UpdateAccount
type Dbv0038UpdateAccount struct {
	Accounts []Dbv0038Account `json:"accounts,omitempty"`
}

// NewDbv0038UpdateAccount instantiates a new Dbv0038UpdateAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0038UpdateAccount() *Dbv0038UpdateAccount {
	this := Dbv0038UpdateAccount{}
	return &this
}

// NewDbv0038UpdateAccountWithDefaults instantiates a new Dbv0038UpdateAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0038UpdateAccountWithDefaults() *Dbv0038UpdateAccount {
	this := Dbv0038UpdateAccount{}
	return &this
}

// GetAccounts returns the Accounts field value if set, zero value otherwise.
func (o *Dbv0038UpdateAccount) GetAccounts() []Dbv0038Account {
	if o == nil || IsNil(o.Accounts) {
		var ret []Dbv0038Account
		return ret
	}
	return o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038UpdateAccount) GetAccountsOk() ([]Dbv0038Account, bool) {
	if o == nil || IsNil(o.Accounts) {
		return nil, false
	}
	return o.Accounts, true
}

// HasAccounts returns a boolean if a field has been set.
func (o *Dbv0038UpdateAccount) HasAccounts() bool {
	if o != nil && !IsNil(o.Accounts) {
		return true
	}

	return false
}

// SetAccounts gets a reference to the given []Dbv0038Account and assigns it to the Accounts field.
func (o *Dbv0038UpdateAccount) SetAccounts(v []Dbv0038Account) {
	o.Accounts = v
}

func (o Dbv0038UpdateAccount) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0038UpdateAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Accounts) {
		toSerialize["accounts"] = o.Accounts
	}
	return toSerialize, nil
}

type NullableDbv0038UpdateAccount struct {
	value *Dbv0038UpdateAccount
	isSet bool
}

func (v NullableDbv0038UpdateAccount) Get() *Dbv0038UpdateAccount {
	return v.value
}

func (v *NullableDbv0038UpdateAccount) Set(val *Dbv0038UpdateAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0038UpdateAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0038UpdateAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0038UpdateAccount(val *Dbv0038UpdateAccount) *NullableDbv0038UpdateAccount {
	return &NullableDbv0038UpdateAccount{value: val, isSet: true}
}

func (v NullableDbv0038UpdateAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0038UpdateAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
