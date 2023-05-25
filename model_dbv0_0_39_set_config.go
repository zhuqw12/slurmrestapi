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

// checks if the Dbv0039SetConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0039SetConfig{}

// Dbv0039SetConfig struct for Dbv0039SetConfig
type Dbv0039SetConfig struct {
	Clusters     []V0039ClusterRec `json:"clusters,omitempty"`
	TRES         [][]V0039Tres     `json:"TRES,omitempty"`
	Accounts     []V0039Account    `json:"accounts,omitempty"`
	Users        []V0039User       `json:"users,omitempty"`
	Qos          []V0039Qos        `json:"qos,omitempty"`
	Wckeys       []V0039Wckey      `json:"wckeys,omitempty"`
	Associations []V0039Assoc      `json:"associations,omitempty"`
}

// NewDbv0039SetConfig instantiates a new Dbv0039SetConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0039SetConfig() *Dbv0039SetConfig {
	this := Dbv0039SetConfig{}
	return &this
}

// NewDbv0039SetConfigWithDefaults instantiates a new Dbv0039SetConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0039SetConfigWithDefaults() *Dbv0039SetConfig {
	this := Dbv0039SetConfig{}
	return &this
}

// GetClusters returns the Clusters field value if set, zero value otherwise.
func (o *Dbv0039SetConfig) GetClusters() []V0039ClusterRec {
	if o == nil || IsNil(o.Clusters) {
		var ret []V0039ClusterRec
		return ret
	}
	return o.Clusters
}

// GetClustersOk returns a tuple with the Clusters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0039SetConfig) GetClustersOk() ([]V0039ClusterRec, bool) {
	if o == nil || IsNil(o.Clusters) {
		return nil, false
	}
	return o.Clusters, true
}

// HasClusters returns a boolean if a field has been set.
func (o *Dbv0039SetConfig) HasClusters() bool {
	if o != nil && !IsNil(o.Clusters) {
		return true
	}

	return false
}

// SetClusters gets a reference to the given []V0039ClusterRec and assigns it to the Clusters field.
func (o *Dbv0039SetConfig) SetClusters(v []V0039ClusterRec) {
	o.Clusters = v
}

// GetTRES returns the TRES field value if set, zero value otherwise.
func (o *Dbv0039SetConfig) GetTRES() [][]V0039Tres {
	if o == nil || IsNil(o.TRES) {
		var ret [][]V0039Tres
		return ret
	}
	return o.TRES
}

// GetTRESOk returns a tuple with the TRES field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0039SetConfig) GetTRESOk() ([][]V0039Tres, bool) {
	if o == nil || IsNil(o.TRES) {
		return nil, false
	}
	return o.TRES, true
}

// HasTRES returns a boolean if a field has been set.
func (o *Dbv0039SetConfig) HasTRES() bool {
	if o != nil && !IsNil(o.TRES) {
		return true
	}

	return false
}

// SetTRES gets a reference to the given [][]V0039Tres and assigns it to the TRES field.
func (o *Dbv0039SetConfig) SetTRES(v [][]V0039Tres) {
	o.TRES = v
}

// GetAccounts returns the Accounts field value if set, zero value otherwise.
func (o *Dbv0039SetConfig) GetAccounts() []V0039Account {
	if o == nil || IsNil(o.Accounts) {
		var ret []V0039Account
		return ret
	}
	return o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0039SetConfig) GetAccountsOk() ([]V0039Account, bool) {
	if o == nil || IsNil(o.Accounts) {
		return nil, false
	}
	return o.Accounts, true
}

// HasAccounts returns a boolean if a field has been set.
func (o *Dbv0039SetConfig) HasAccounts() bool {
	if o != nil && !IsNil(o.Accounts) {
		return true
	}

	return false
}

// SetAccounts gets a reference to the given []V0039Account and assigns it to the Accounts field.
func (o *Dbv0039SetConfig) SetAccounts(v []V0039Account) {
	o.Accounts = v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *Dbv0039SetConfig) GetUsers() []V0039User {
	if o == nil || IsNil(o.Users) {
		var ret []V0039User
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0039SetConfig) GetUsersOk() ([]V0039User, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *Dbv0039SetConfig) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []V0039User and assigns it to the Users field.
func (o *Dbv0039SetConfig) SetUsers(v []V0039User) {
	o.Users = v
}

// GetQos returns the Qos field value if set, zero value otherwise.
func (o *Dbv0039SetConfig) GetQos() []V0039Qos {
	if o == nil || IsNil(o.Qos) {
		var ret []V0039Qos
		return ret
	}
	return o.Qos
}

// GetQosOk returns a tuple with the Qos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0039SetConfig) GetQosOk() ([]V0039Qos, bool) {
	if o == nil || IsNil(o.Qos) {
		return nil, false
	}
	return o.Qos, true
}

// HasQos returns a boolean if a field has been set.
func (o *Dbv0039SetConfig) HasQos() bool {
	if o != nil && !IsNil(o.Qos) {
		return true
	}

	return false
}

// SetQos gets a reference to the given []V0039Qos and assigns it to the Qos field.
func (o *Dbv0039SetConfig) SetQos(v []V0039Qos) {
	o.Qos = v
}

// GetWckeys returns the Wckeys field value if set, zero value otherwise.
func (o *Dbv0039SetConfig) GetWckeys() []V0039Wckey {
	if o == nil || IsNil(o.Wckeys) {
		var ret []V0039Wckey
		return ret
	}
	return o.Wckeys
}

// GetWckeysOk returns a tuple with the Wckeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0039SetConfig) GetWckeysOk() ([]V0039Wckey, bool) {
	if o == nil || IsNil(o.Wckeys) {
		return nil, false
	}
	return o.Wckeys, true
}

// HasWckeys returns a boolean if a field has been set.
func (o *Dbv0039SetConfig) HasWckeys() bool {
	if o != nil && !IsNil(o.Wckeys) {
		return true
	}

	return false
}

// SetWckeys gets a reference to the given []V0039Wckey and assigns it to the Wckeys field.
func (o *Dbv0039SetConfig) SetWckeys(v []V0039Wckey) {
	o.Wckeys = v
}

// GetAssociations returns the Associations field value if set, zero value otherwise.
func (o *Dbv0039SetConfig) GetAssociations() []V0039Assoc {
	if o == nil || IsNil(o.Associations) {
		var ret []V0039Assoc
		return ret
	}
	return o.Associations
}

// GetAssociationsOk returns a tuple with the Associations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0039SetConfig) GetAssociationsOk() ([]V0039Assoc, bool) {
	if o == nil || IsNil(o.Associations) {
		return nil, false
	}
	return o.Associations, true
}

// HasAssociations returns a boolean if a field has been set.
func (o *Dbv0039SetConfig) HasAssociations() bool {
	if o != nil && !IsNil(o.Associations) {
		return true
	}

	return false
}

// SetAssociations gets a reference to the given []V0039Assoc and assigns it to the Associations field.
func (o *Dbv0039SetConfig) SetAssociations(v []V0039Assoc) {
	o.Associations = v
}

func (o Dbv0039SetConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0039SetConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Clusters) {
		toSerialize["clusters"] = o.Clusters
	}
	if !IsNil(o.TRES) {
		toSerialize["TRES"] = o.TRES
	}
	if !IsNil(o.Accounts) {
		toSerialize["accounts"] = o.Accounts
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
	if !IsNil(o.Associations) {
		toSerialize["associations"] = o.Associations
	}
	return toSerialize, nil
}

type NullableDbv0039SetConfig struct {
	value *Dbv0039SetConfig
	isSet bool
}

func (v NullableDbv0039SetConfig) Get() *Dbv0039SetConfig {
	return v.value
}

func (v *NullableDbv0039SetConfig) Set(val *Dbv0039SetConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0039SetConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0039SetConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0039SetConfig(val *Dbv0039SetConfig) *NullableDbv0039SetConfig {
	return &NullableDbv0039SetConfig{value: val, isSet: true}
}

func (v NullableDbv0039SetConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0039SetConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
