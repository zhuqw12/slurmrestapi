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

// checks if the Dbv0039ConfigInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0039ConfigInfo{}

// Dbv0039ConfigInfo struct for Dbv0039ConfigInfo
type Dbv0039ConfigInfo struct {
	Meta *Dbv0039Meta `json:"meta,omitempty"`
	// Slurm errors
	Errors []Dbv0039Error `json:"errors,omitempty"`
	// Slurm warnings
	Warnings     []Dbv0039Warning  `json:"warnings,omitempty"`
	Tres         []V0039Tres       `json:"tres,omitempty"`
	Accounts     []V0039Account    `json:"accounts,omitempty"`
	Associations []V0039Assoc      `json:"associations,omitempty"`
	Users        []V0039User       `json:"users,omitempty"`
	Qos          []V0039Qos        `json:"qos,omitempty"`
	Wckeys       []V0039Wckey      `json:"wckeys,omitempty"`
	Clusters     []V0039ClusterRec `json:"clusters,omitempty"`
}

// NewDbv0039ConfigInfo instantiates a new Dbv0039ConfigInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0039ConfigInfo() *Dbv0039ConfigInfo {
	this := Dbv0039ConfigInfo{}
	return &this
}

// NewDbv0039ConfigInfoWithDefaults instantiates a new Dbv0039ConfigInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0039ConfigInfoWithDefaults() *Dbv0039ConfigInfo {
	this := Dbv0039ConfigInfo{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *Dbv0039ConfigInfo) GetMeta() Dbv0039Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Dbv0039Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0039ConfigInfo) GetMetaOk() (*Dbv0039Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *Dbv0039ConfigInfo) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Dbv0039Meta and assigns it to the Meta field.
func (o *Dbv0039ConfigInfo) SetMeta(v Dbv0039Meta) {
	o.Meta = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *Dbv0039ConfigInfo) GetErrors() []Dbv0039Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Dbv0039Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0039ConfigInfo) GetErrorsOk() ([]Dbv0039Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *Dbv0039ConfigInfo) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Dbv0039Error and assigns it to the Errors field.
func (o *Dbv0039ConfigInfo) SetErrors(v []Dbv0039Error) {
	o.Errors = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *Dbv0039ConfigInfo) GetWarnings() []Dbv0039Warning {
	if o == nil || IsNil(o.Warnings) {
		var ret []Dbv0039Warning
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0039ConfigInfo) GetWarningsOk() ([]Dbv0039Warning, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *Dbv0039ConfigInfo) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []Dbv0039Warning and assigns it to the Warnings field.
func (o *Dbv0039ConfigInfo) SetWarnings(v []Dbv0039Warning) {
	o.Warnings = v
}

// GetTres returns the Tres field value if set, zero value otherwise.
func (o *Dbv0039ConfigInfo) GetTres() []V0039Tres {
	if o == nil || IsNil(o.Tres) {
		var ret []V0039Tres
		return ret
	}
	return o.Tres
}

// GetTresOk returns a tuple with the Tres field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0039ConfigInfo) GetTresOk() ([]V0039Tres, bool) {
	if o == nil || IsNil(o.Tres) {
		return nil, false
	}
	return o.Tres, true
}

// HasTres returns a boolean if a field has been set.
func (o *Dbv0039ConfigInfo) HasTres() bool {
	if o != nil && !IsNil(o.Tres) {
		return true
	}

	return false
}

// SetTres gets a reference to the given []V0039Tres and assigns it to the Tres field.
func (o *Dbv0039ConfigInfo) SetTres(v []V0039Tres) {
	o.Tres = v
}

// GetAccounts returns the Accounts field value if set, zero value otherwise.
func (o *Dbv0039ConfigInfo) GetAccounts() []V0039Account {
	if o == nil || IsNil(o.Accounts) {
		var ret []V0039Account
		return ret
	}
	return o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0039ConfigInfo) GetAccountsOk() ([]V0039Account, bool) {
	if o == nil || IsNil(o.Accounts) {
		return nil, false
	}
	return o.Accounts, true
}

// HasAccounts returns a boolean if a field has been set.
func (o *Dbv0039ConfigInfo) HasAccounts() bool {
	if o != nil && !IsNil(o.Accounts) {
		return true
	}

	return false
}

// SetAccounts gets a reference to the given []V0039Account and assigns it to the Accounts field.
func (o *Dbv0039ConfigInfo) SetAccounts(v []V0039Account) {
	o.Accounts = v
}

// GetAssociations returns the Associations field value if set, zero value otherwise.
func (o *Dbv0039ConfigInfo) GetAssociations() []V0039Assoc {
	if o == nil || IsNil(o.Associations) {
		var ret []V0039Assoc
		return ret
	}
	return o.Associations
}

// GetAssociationsOk returns a tuple with the Associations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0039ConfigInfo) GetAssociationsOk() ([]V0039Assoc, bool) {
	if o == nil || IsNil(o.Associations) {
		return nil, false
	}
	return o.Associations, true
}

// HasAssociations returns a boolean if a field has been set.
func (o *Dbv0039ConfigInfo) HasAssociations() bool {
	if o != nil && !IsNil(o.Associations) {
		return true
	}

	return false
}

// SetAssociations gets a reference to the given []V0039Assoc and assigns it to the Associations field.
func (o *Dbv0039ConfigInfo) SetAssociations(v []V0039Assoc) {
	o.Associations = v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *Dbv0039ConfigInfo) GetUsers() []V0039User {
	if o == nil || IsNil(o.Users) {
		var ret []V0039User
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0039ConfigInfo) GetUsersOk() ([]V0039User, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *Dbv0039ConfigInfo) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []V0039User and assigns it to the Users field.
func (o *Dbv0039ConfigInfo) SetUsers(v []V0039User) {
	o.Users = v
}

// GetQos returns the Qos field value if set, zero value otherwise.
func (o *Dbv0039ConfigInfo) GetQos() []V0039Qos {
	if o == nil || IsNil(o.Qos) {
		var ret []V0039Qos
		return ret
	}
	return o.Qos
}

// GetQosOk returns a tuple with the Qos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0039ConfigInfo) GetQosOk() ([]V0039Qos, bool) {
	if o == nil || IsNil(o.Qos) {
		return nil, false
	}
	return o.Qos, true
}

// HasQos returns a boolean if a field has been set.
func (o *Dbv0039ConfigInfo) HasQos() bool {
	if o != nil && !IsNil(o.Qos) {
		return true
	}

	return false
}

// SetQos gets a reference to the given []V0039Qos and assigns it to the Qos field.
func (o *Dbv0039ConfigInfo) SetQos(v []V0039Qos) {
	o.Qos = v
}

// GetWckeys returns the Wckeys field value if set, zero value otherwise.
func (o *Dbv0039ConfigInfo) GetWckeys() []V0039Wckey {
	if o == nil || IsNil(o.Wckeys) {
		var ret []V0039Wckey
		return ret
	}
	return o.Wckeys
}

// GetWckeysOk returns a tuple with the Wckeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0039ConfigInfo) GetWckeysOk() ([]V0039Wckey, bool) {
	if o == nil || IsNil(o.Wckeys) {
		return nil, false
	}
	return o.Wckeys, true
}

// HasWckeys returns a boolean if a field has been set.
func (o *Dbv0039ConfigInfo) HasWckeys() bool {
	if o != nil && !IsNil(o.Wckeys) {
		return true
	}

	return false
}

// SetWckeys gets a reference to the given []V0039Wckey and assigns it to the Wckeys field.
func (o *Dbv0039ConfigInfo) SetWckeys(v []V0039Wckey) {
	o.Wckeys = v
}

// GetClusters returns the Clusters field value if set, zero value otherwise.
func (o *Dbv0039ConfigInfo) GetClusters() []V0039ClusterRec {
	if o == nil || IsNil(o.Clusters) {
		var ret []V0039ClusterRec
		return ret
	}
	return o.Clusters
}

// GetClustersOk returns a tuple with the Clusters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0039ConfigInfo) GetClustersOk() ([]V0039ClusterRec, bool) {
	if o == nil || IsNil(o.Clusters) {
		return nil, false
	}
	return o.Clusters, true
}

// HasClusters returns a boolean if a field has been set.
func (o *Dbv0039ConfigInfo) HasClusters() bool {
	if o != nil && !IsNil(o.Clusters) {
		return true
	}

	return false
}

// SetClusters gets a reference to the given []V0039ClusterRec and assigns it to the Clusters field.
func (o *Dbv0039ConfigInfo) SetClusters(v []V0039ClusterRec) {
	o.Clusters = v
}

func (o Dbv0039ConfigInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0039ConfigInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
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
	if !IsNil(o.Clusters) {
		toSerialize["clusters"] = o.Clusters
	}
	return toSerialize, nil
}

type NullableDbv0039ConfigInfo struct {
	value *Dbv0039ConfigInfo
	isSet bool
}

func (v NullableDbv0039ConfigInfo) Get() *Dbv0039ConfigInfo {
	return v.value
}

func (v *NullableDbv0039ConfigInfo) Set(val *Dbv0039ConfigInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0039ConfigInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0039ConfigInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0039ConfigInfo(val *Dbv0039ConfigInfo) *NullableDbv0039ConfigInfo {
	return &NullableDbv0039ConfigInfo{value: val, isSet: true}
}

func (v NullableDbv0039ConfigInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0039ConfigInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
