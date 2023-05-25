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

// checks if the Dbv0038UserDefault type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0038UserDefault{}

// Dbv0038UserDefault Default settings
type Dbv0038UserDefault struct {
	// Default account name
	Account *string `json:"account,omitempty"`
	// Default wckey
	Wckey *string `json:"wckey,omitempty"`
}

// NewDbv0038UserDefault instantiates a new Dbv0038UserDefault object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0038UserDefault() *Dbv0038UserDefault {
	this := Dbv0038UserDefault{}
	return &this
}

// NewDbv0038UserDefaultWithDefaults instantiates a new Dbv0038UserDefault object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0038UserDefaultWithDefaults() *Dbv0038UserDefault {
	this := Dbv0038UserDefault{}
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *Dbv0038UserDefault) GetAccount() string {
	if o == nil || IsNil(o.Account) {
		var ret string
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038UserDefault) GetAccountOk() (*string, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *Dbv0038UserDefault) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given string and assigns it to the Account field.
func (o *Dbv0038UserDefault) SetAccount(v string) {
	o.Account = &v
}

// GetWckey returns the Wckey field value if set, zero value otherwise.
func (o *Dbv0038UserDefault) GetWckey() string {
	if o == nil || IsNil(o.Wckey) {
		var ret string
		return ret
	}
	return *o.Wckey
}

// GetWckeyOk returns a tuple with the Wckey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038UserDefault) GetWckeyOk() (*string, bool) {
	if o == nil || IsNil(o.Wckey) {
		return nil, false
	}
	return o.Wckey, true
}

// HasWckey returns a boolean if a field has been set.
func (o *Dbv0038UserDefault) HasWckey() bool {
	if o != nil && !IsNil(o.Wckey) {
		return true
	}

	return false
}

// SetWckey gets a reference to the given string and assigns it to the Wckey field.
func (o *Dbv0038UserDefault) SetWckey(v string) {
	o.Wckey = &v
}

func (o Dbv0038UserDefault) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0038UserDefault) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	if !IsNil(o.Wckey) {
		toSerialize["wckey"] = o.Wckey
	}
	return toSerialize, nil
}

type NullableDbv0038UserDefault struct {
	value *Dbv0038UserDefault
	isSet bool
}

func (v NullableDbv0038UserDefault) Get() *Dbv0038UserDefault {
	return v.value
}

func (v *NullableDbv0038UserDefault) Set(val *Dbv0038UserDefault) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0038UserDefault) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0038UserDefault) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0038UserDefault(val *Dbv0038UserDefault) *NullableDbv0038UserDefault {
	return &NullableDbv0038UserDefault{value: val, isSet: true}
}

func (v NullableDbv0038UserDefault) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0038UserDefault) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
