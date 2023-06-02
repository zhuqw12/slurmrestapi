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

// checks if the Dbv0038SetConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0038SetConfig{}

// Dbv0038SetConfig struct for Dbv0038SetConfig
type Dbv0038SetConfig struct {
	Clusters     []Dbv0038ClustersProperties `json:"clusters,omitempty"`
	TRES         [][]Dbv0038TresListInner    `json:"TRES,omitempty"`
	Accounts     []Dbv0038UpdateAccount      `json:"accounts,omitempty"`
	Users        []Dbv0038User               `json:"users,omitempty"`
	Qos          []Dbv0038Qos                `json:"qos,omitempty"`
	Wckeys       []Dbv0038Wckey              `json:"wckeys,omitempty"`
	Associations []Dbv0038Association        `json:"associations,omitempty"`
}

// NewDbv0038SetConfig instantiates a new Dbv0038SetConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0038SetConfig() *Dbv0038SetConfig {
	this := Dbv0038SetConfig{}
	return &this
}

// NewDbv0038SetConfigWithDefaults instantiates a new Dbv0038SetConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0038SetConfigWithDefaults() *Dbv0038SetConfig {
	this := Dbv0038SetConfig{}
	return &this
}

// GetClusters returns the Clusters field value if set, zero value otherwise.
func (o *Dbv0038SetConfig) GetClusters() []Dbv0038ClustersProperties {
	if o == nil || IsNil(o.Clusters) {
		var ret []Dbv0038ClustersProperties
		return ret
	}
	return o.Clusters
}

// GetClustersOk returns a tuple with the Clusters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038SetConfig) GetClustersOk() ([]Dbv0038ClustersProperties, bool) {
	if o == nil || IsNil(o.Clusters) {
		return nil, false
	}
	return o.Clusters, true
}

// HasClusters returns a boolean if a field has been set.
func (o *Dbv0038SetConfig) HasClusters() bool {
	if o != nil && !IsNil(o.Clusters) {
		return true
	}

	return false
}

// SetClusters gets a reference to the given []Dbv0038ClustersProperties and assigns it to the Clusters field.
func (o *Dbv0038SetConfig) SetClusters(v []Dbv0038ClustersProperties) {
	o.Clusters = v
}

// GetTRES returns the TRES field value if set, zero value otherwise.
func (o *Dbv0038SetConfig) GetTRES() [][]Dbv0038TresListInner {
	if o == nil || IsNil(o.TRES) {
		var ret [][]Dbv0038TresListInner
		return ret
	}
	return o.TRES
}

// GetTRESOk returns a tuple with the TRES field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038SetConfig) GetTRESOk() ([][]Dbv0038TresListInner, bool) {
	if o == nil || IsNil(o.TRES) {
		return nil, false
	}
	return o.TRES, true
}

// HasTRES returns a boolean if a field has been set.
func (o *Dbv0038SetConfig) HasTRES() bool {
	if o != nil && !IsNil(o.TRES) {
		return true
	}

	return false
}

// SetTRES gets a reference to the given [][]Dbv0038TresListInner and assigns it to the TRES field.
func (o *Dbv0038SetConfig) SetTRES(v [][]Dbv0038TresListInner) {
	o.TRES = v
}

// GetAccounts returns the Accounts field value if set, zero value otherwise.
func (o *Dbv0038SetConfig) GetAccounts() []Dbv0038UpdateAccount {
	if o == nil || IsNil(o.Accounts) {
		var ret []Dbv0038UpdateAccount
		return ret
	}
	return o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038SetConfig) GetAccountsOk() ([]Dbv0038UpdateAccount, bool) {
	if o == nil || IsNil(o.Accounts) {
		return nil, false
	}
	return o.Accounts, true
}

// HasAccounts returns a boolean if a field has been set.
func (o *Dbv0038SetConfig) HasAccounts() bool {
	if o != nil && !IsNil(o.Accounts) {
		return true
	}

	return false
}

// SetAccounts gets a reference to the given []Dbv0038UpdateAccount and assigns it to the Accounts field.
func (o *Dbv0038SetConfig) SetAccounts(v []Dbv0038UpdateAccount) {
	o.Accounts = v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *Dbv0038SetConfig) GetUsers() []Dbv0038User {
	if o == nil || IsNil(o.Users) {
		var ret []Dbv0038User
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038SetConfig) GetUsersOk() ([]Dbv0038User, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *Dbv0038SetConfig) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []Dbv0038User and assigns it to the Users field.
func (o *Dbv0038SetConfig) SetUsers(v []Dbv0038User) {
	o.Users = v
}

// GetQos returns the Qos field value if set, zero value otherwise.
func (o *Dbv0038SetConfig) GetQos() []Dbv0038Qos {
	if o == nil || IsNil(o.Qos) {
		var ret []Dbv0038Qos
		return ret
	}
	return o.Qos
}

// GetQosOk returns a tuple with the Qos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038SetConfig) GetQosOk() ([]Dbv0038Qos, bool) {
	if o == nil || IsNil(o.Qos) {
		return nil, false
	}
	return o.Qos, true
}

// HasQos returns a boolean if a field has been set.
func (o *Dbv0038SetConfig) HasQos() bool {
	if o != nil && !IsNil(o.Qos) {
		return true
	}

	return false
}

// SetQos gets a reference to the given []Dbv0038Qos and assigns it to the Qos field.
func (o *Dbv0038SetConfig) SetQos(v []Dbv0038Qos) {
	o.Qos = v
}

// GetWckeys returns the Wckeys field value if set, zero value otherwise.
func (o *Dbv0038SetConfig) GetWckeys() []Dbv0038Wckey {
	if o == nil || IsNil(o.Wckeys) {
		var ret []Dbv0038Wckey
		return ret
	}
	return o.Wckeys
}

// GetWckeysOk returns a tuple with the Wckeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038SetConfig) GetWckeysOk() ([]Dbv0038Wckey, bool) {
	if o == nil || IsNil(o.Wckeys) {
		return nil, false
	}
	return o.Wckeys, true
}

// HasWckeys returns a boolean if a field has been set.
func (o *Dbv0038SetConfig) HasWckeys() bool {
	if o != nil && !IsNil(o.Wckeys) {
		return true
	}

	return false
}

// SetWckeys gets a reference to the given []Dbv0038Wckey and assigns it to the Wckeys field.
func (o *Dbv0038SetConfig) SetWckeys(v []Dbv0038Wckey) {
	o.Wckeys = v
}

// GetAssociations returns the Associations field value if set, zero value otherwise.
func (o *Dbv0038SetConfig) GetAssociations() []Dbv0038Association {
	if o == nil || IsNil(o.Associations) {
		var ret []Dbv0038Association
		return ret
	}
	return o.Associations
}

// GetAssociationsOk returns a tuple with the Associations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038SetConfig) GetAssociationsOk() ([]Dbv0038Association, bool) {
	if o == nil || IsNil(o.Associations) {
		return nil, false
	}
	return o.Associations, true
}

// HasAssociations returns a boolean if a field has been set.
func (o *Dbv0038SetConfig) HasAssociations() bool {
	if o != nil && !IsNil(o.Associations) {
		return true
	}

	return false
}

// SetAssociations gets a reference to the given []Dbv0038Association and assigns it to the Associations field.
func (o *Dbv0038SetConfig) SetAssociations(v []Dbv0038Association) {
	o.Associations = v
}

func (o Dbv0038SetConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0038SetConfig) ToMap() (map[string]interface{}, error) {
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

type NullableDbv0038SetConfig struct {
	value *Dbv0038SetConfig
	isSet bool
}

func (v NullableDbv0038SetConfig) Get() *Dbv0038SetConfig {
	return v.value
}

func (v *NullableDbv0038SetConfig) Set(val *Dbv0038SetConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0038SetConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0038SetConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0038SetConfig(val *Dbv0038SetConfig) *NullableDbv0038SetConfig {
	return &NullableDbv0038SetConfig{value: val, isSet: true}
}

func (v NullableDbv0038SetConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0038SetConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}