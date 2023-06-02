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

// checks if the Dbv0038AssociationMaxPer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0038AssociationMaxPer{}

// Dbv0038AssociationMaxPer Max per settings
type Dbv0038AssociationMaxPer struct {
	Account *Dbv0038AssociationMaxPerAccount `json:"account,omitempty"`
}

// NewDbv0038AssociationMaxPer instantiates a new Dbv0038AssociationMaxPer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0038AssociationMaxPer() *Dbv0038AssociationMaxPer {
	this := Dbv0038AssociationMaxPer{}
	return &this
}

// NewDbv0038AssociationMaxPerWithDefaults instantiates a new Dbv0038AssociationMaxPer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0038AssociationMaxPerWithDefaults() *Dbv0038AssociationMaxPer {
	this := Dbv0038AssociationMaxPer{}
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *Dbv0038AssociationMaxPer) GetAccount() Dbv0038AssociationMaxPerAccount {
	if o == nil || IsNil(o.Account) {
		var ret Dbv0038AssociationMaxPerAccount
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038AssociationMaxPer) GetAccountOk() (*Dbv0038AssociationMaxPerAccount, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *Dbv0038AssociationMaxPer) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given Dbv0038AssociationMaxPerAccount and assigns it to the Account field.
func (o *Dbv0038AssociationMaxPer) SetAccount(v Dbv0038AssociationMaxPerAccount) {
	o.Account = &v
}

func (o Dbv0038AssociationMaxPer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0038AssociationMaxPer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	return toSerialize, nil
}

type NullableDbv0038AssociationMaxPer struct {
	value *Dbv0038AssociationMaxPer
	isSet bool
}

func (v NullableDbv0038AssociationMaxPer) Get() *Dbv0038AssociationMaxPer {
	return v.value
}

func (v *NullableDbv0038AssociationMaxPer) Set(val *Dbv0038AssociationMaxPer) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0038AssociationMaxPer) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0038AssociationMaxPer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0038AssociationMaxPer(val *Dbv0038AssociationMaxPer) *NullableDbv0038AssociationMaxPer {
	return &NullableDbv0038AssociationMaxPer{value: val, isSet: true}
}

func (v NullableDbv0038AssociationMaxPer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0038AssociationMaxPer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
