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

// checks if the Dbv0037ConfigInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0037ConfigInfo{}

// Dbv0037ConfigInfo struct for Dbv0037ConfigInfo
type Dbv0037ConfigInfo struct {
	// Slurm errors
	Errors []Dbv0037Error `json:"errors,omitempty"`
	// Array of TRES
	Tres [][]Dbv0038TresListInner `json:"tres,omitempty"`
	// Array of accounts
	Accounts []Dbv0037Account `json:"accounts,omitempty"`
	// Array of associations
	Associations []Dbv0037Association `json:"associations,omitempty"`
	// Array of users
	Users []Dbv0037User `json:"users,omitempty"`
	// Array of qos
	Qos []Dbv0037Qos `json:"qos,omitempty"`
	// Array of wckeys
	Wckeys []Dbv0037Wckey `json:"wckeys,omitempty"`
}

// NewDbv0037ConfigInfo instantiates a new Dbv0037ConfigInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0037ConfigInfo() *Dbv0037ConfigInfo {
	this := Dbv0037ConfigInfo{}
	return &this
}

// NewDbv0037ConfigInfoWithDefaults instantiates a new Dbv0037ConfigInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0037ConfigInfoWithDefaults() *Dbv0037ConfigInfo {
	this := Dbv0037ConfigInfo{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *Dbv0037ConfigInfo) GetErrors() []Dbv0037Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Dbv0037Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037ConfigInfo) GetErrorsOk() ([]Dbv0037Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *Dbv0037ConfigInfo) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Dbv0037Error and assigns it to the Errors field.
func (o *Dbv0037ConfigInfo) SetErrors(v []Dbv0037Error) {
	o.Errors = v
}

// GetTres returns the Tres field value if set, zero value otherwise.
func (o *Dbv0037ConfigInfo) GetTres() [][]Dbv0038TresListInner {
	if o == nil || IsNil(o.Tres) {
		var ret [][]Dbv0038TresListInner
		return ret
	}
	return o.Tres
}

// GetTresOk returns a tuple with the Tres field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037ConfigInfo) GetTresOk() ([][]Dbv0038TresListInner, bool) {
	if o == nil || IsNil(o.Tres) {
		return nil, false
	}
	return o.Tres, true
}

// HasTres returns a boolean if a field has been set.
func (o *Dbv0037ConfigInfo) HasTres() bool {
	if o != nil && !IsNil(o.Tres) {
		return true
	}

	return false
}

// SetTres gets a reference to the given [][]Dbv0038TresListInner and assigns it to the Tres field.
func (o *Dbv0037ConfigInfo) SetTres(v [][]Dbv0038TresListInner) {
	o.Tres = v
}

// GetAccounts returns the Accounts field value if set, zero value otherwise.
func (o *Dbv0037ConfigInfo) GetAccounts() []Dbv0037Account {
	if o == nil || IsNil(o.Accounts) {
		var ret []Dbv0037Account
		return ret
	}
	return o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037ConfigInfo) GetAccountsOk() ([]Dbv0037Account, bool) {
	if o == nil || IsNil(o.Accounts) {
		return nil, false
	}
	return o.Accounts, true
}

// HasAccounts returns a boolean if a field has been set.
func (o *Dbv0037ConfigInfo) HasAccounts() bool {
	if o != nil && !IsNil(o.Accounts) {
		return true
	}

	return false
}

// SetAccounts gets a reference to the given []Dbv0037Account and assigns it to the Accounts field.
func (o *Dbv0037ConfigInfo) SetAccounts(v []Dbv0037Account) {
	o.Accounts = v
}

// GetAssociations returns the Associations field value if set, zero value otherwise.
func (o *Dbv0037ConfigInfo) GetAssociations() []Dbv0037Association {
	if o == nil || IsNil(o.Associations) {
		var ret []Dbv0037Association
		return ret
	}
	return o.Associations
}

// GetAssociationsOk returns a tuple with the Associations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037ConfigInfo) GetAssociationsOk() ([]Dbv0037Association, bool) {
	if o == nil || IsNil(o.Associations) {
		return nil, false
	}
	return o.Associations, true
}

// HasAssociations returns a boolean if a field has been set.
func (o *Dbv0037ConfigInfo) HasAssociations() bool {
	if o != nil && !IsNil(o.Associations) {
		return true
	}

	return false
}

// SetAssociations gets a reference to the given []Dbv0037Association and assigns it to the Associations field.
func (o *Dbv0037ConfigInfo) SetAssociations(v []Dbv0037Association) {
	o.Associations = v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *Dbv0037ConfigInfo) GetUsers() []Dbv0037User {
	if o == nil || IsNil(o.Users) {
		var ret []Dbv0037User
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037ConfigInfo) GetUsersOk() ([]Dbv0037User, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *Dbv0037ConfigInfo) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []Dbv0037User and assigns it to the Users field.
func (o *Dbv0037ConfigInfo) SetUsers(v []Dbv0037User) {
	o.Users = v
}

// GetQos returns the Qos field value if set, zero value otherwise.
func (o *Dbv0037ConfigInfo) GetQos() []Dbv0037Qos {
	if o == nil || IsNil(o.Qos) {
		var ret []Dbv0037Qos
		return ret
	}
	return o.Qos
}

// GetQosOk returns a tuple with the Qos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037ConfigInfo) GetQosOk() ([]Dbv0037Qos, bool) {
	if o == nil || IsNil(o.Qos) {
		return nil, false
	}
	return o.Qos, true
}

// HasQos returns a boolean if a field has been set.
func (o *Dbv0037ConfigInfo) HasQos() bool {
	if o != nil && !IsNil(o.Qos) {
		return true
	}

	return false
}

// SetQos gets a reference to the given []Dbv0037Qos and assigns it to the Qos field.
func (o *Dbv0037ConfigInfo) SetQos(v []Dbv0037Qos) {
	o.Qos = v
}

// GetWckeys returns the Wckeys field value if set, zero value otherwise.
func (o *Dbv0037ConfigInfo) GetWckeys() []Dbv0037Wckey {
	if o == nil || IsNil(o.Wckeys) {
		var ret []Dbv0037Wckey
		return ret
	}
	return o.Wckeys
}

// GetWckeysOk returns a tuple with the Wckeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037ConfigInfo) GetWckeysOk() ([]Dbv0037Wckey, bool) {
	if o == nil || IsNil(o.Wckeys) {
		return nil, false
	}
	return o.Wckeys, true
}

// HasWckeys returns a boolean if a field has been set.
func (o *Dbv0037ConfigInfo) HasWckeys() bool {
	if o != nil && !IsNil(o.Wckeys) {
		return true
	}

	return false
}

// SetWckeys gets a reference to the given []Dbv0037Wckey and assigns it to the Wckeys field.
func (o *Dbv0037ConfigInfo) SetWckeys(v []Dbv0037Wckey) {
	o.Wckeys = v
}

func (o Dbv0037ConfigInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0037ConfigInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !IsNil(o.Tres) {
		toSerialize["tres"] = o.Tres
	}
	if !IsNil(o.Accounts) {
		toSerialize["accounts"] = o.Accounts
	}
	if !IsNil(o.Associations) {
		toSerialize["associations"] = o.Associations
	}
	if !IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}
	if !IsNil(o.Qos) {
		toSerialize["qos"] = o.Qos
	}
	if !IsNil(o.Wckeys) {
		toSerialize["wckeys"] = o.Wckeys
	}
	return toSerialize, nil
}

type NullableDbv0037ConfigInfo struct {
	value *Dbv0037ConfigInfo
	isSet bool
}

func (v NullableDbv0037ConfigInfo) Get() *Dbv0037ConfigInfo {
	return v.value
}

func (v *NullableDbv0037ConfigInfo) Set(val *Dbv0037ConfigInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0037ConfigInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0037ConfigInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0037ConfigInfo(val *Dbv0037ConfigInfo) *NullableDbv0037ConfigInfo {
	return &NullableDbv0037ConfigInfo{value: val, isSet: true}
}

func (v NullableDbv0037ConfigInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0037ConfigInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
