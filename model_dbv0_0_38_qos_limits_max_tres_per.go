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

// checks if the Dbv0038QosLimitsMaxTresPer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0038QosLimitsMaxTresPer{}

// Dbv0038QosLimitsMaxTresPer Max TRES per settings
type Dbv0038QosLimitsMaxTresPer struct {
	// TRES list of attributes
	Account []Dbv0038TresListInner `json:"account,omitempty"`
	// TRES list of attributes
	Job []Dbv0038TresListInner `json:"job,omitempty"`
	// TRES list of attributes
	Node []Dbv0038TresListInner `json:"node,omitempty"`
	// TRES list of attributes
	User []Dbv0038TresListInner `json:"user,omitempty"`
}

// NewDbv0038QosLimitsMaxTresPer instantiates a new Dbv0038QosLimitsMaxTresPer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0038QosLimitsMaxTresPer() *Dbv0038QosLimitsMaxTresPer {
	this := Dbv0038QosLimitsMaxTresPer{}
	return &this
}

// NewDbv0038QosLimitsMaxTresPerWithDefaults instantiates a new Dbv0038QosLimitsMaxTresPer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0038QosLimitsMaxTresPerWithDefaults() *Dbv0038QosLimitsMaxTresPer {
	this := Dbv0038QosLimitsMaxTresPer{}
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *Dbv0038QosLimitsMaxTresPer) GetAccount() []Dbv0038TresListInner {
	if o == nil || IsNil(o.Account) {
		var ret []Dbv0038TresListInner
		return ret
	}
	return o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038QosLimitsMaxTresPer) GetAccountOk() ([]Dbv0038TresListInner, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *Dbv0038QosLimitsMaxTresPer) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given []Dbv0038TresListInner and assigns it to the Account field.
func (o *Dbv0038QosLimitsMaxTresPer) SetAccount(v []Dbv0038TresListInner) {
	o.Account = v
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *Dbv0038QosLimitsMaxTresPer) GetJob() []Dbv0038TresListInner {
	if o == nil || IsNil(o.Job) {
		var ret []Dbv0038TresListInner
		return ret
	}
	return o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038QosLimitsMaxTresPer) GetJobOk() ([]Dbv0038TresListInner, bool) {
	if o == nil || IsNil(o.Job) {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *Dbv0038QosLimitsMaxTresPer) HasJob() bool {
	if o != nil && !IsNil(o.Job) {
		return true
	}

	return false
}

// SetJob gets a reference to the given []Dbv0038TresListInner and assigns it to the Job field.
func (o *Dbv0038QosLimitsMaxTresPer) SetJob(v []Dbv0038TresListInner) {
	o.Job = v
}

// GetNode returns the Node field value if set, zero value otherwise.
func (o *Dbv0038QosLimitsMaxTresPer) GetNode() []Dbv0038TresListInner {
	if o == nil || IsNil(o.Node) {
		var ret []Dbv0038TresListInner
		return ret
	}
	return o.Node
}

// GetNodeOk returns a tuple with the Node field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038QosLimitsMaxTresPer) GetNodeOk() ([]Dbv0038TresListInner, bool) {
	if o == nil || IsNil(o.Node) {
		return nil, false
	}
	return o.Node, true
}

// HasNode returns a boolean if a field has been set.
func (o *Dbv0038QosLimitsMaxTresPer) HasNode() bool {
	if o != nil && !IsNil(o.Node) {
		return true
	}

	return false
}

// SetNode gets a reference to the given []Dbv0038TresListInner and assigns it to the Node field.
func (o *Dbv0038QosLimitsMaxTresPer) SetNode(v []Dbv0038TresListInner) {
	o.Node = v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *Dbv0038QosLimitsMaxTresPer) GetUser() []Dbv0038TresListInner {
	if o == nil || IsNil(o.User) {
		var ret []Dbv0038TresListInner
		return ret
	}
	return o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038QosLimitsMaxTresPer) GetUserOk() ([]Dbv0038TresListInner, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *Dbv0038QosLimitsMaxTresPer) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given []Dbv0038TresListInner and assigns it to the User field.
func (o *Dbv0038QosLimitsMaxTresPer) SetUser(v []Dbv0038TresListInner) {
	o.User = v
}

func (o Dbv0038QosLimitsMaxTresPer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0038QosLimitsMaxTresPer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	if !IsNil(o.Job) {
		toSerialize["job"] = o.Job
	}
	if !IsNil(o.Node) {
		toSerialize["node"] = o.Node
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	return toSerialize, nil
}

type NullableDbv0038QosLimitsMaxTresPer struct {
	value *Dbv0038QosLimitsMaxTresPer
	isSet bool
}

func (v NullableDbv0038QosLimitsMaxTresPer) Get() *Dbv0038QosLimitsMaxTresPer {
	return v.value
}

func (v *NullableDbv0038QosLimitsMaxTresPer) Set(val *Dbv0038QosLimitsMaxTresPer) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0038QosLimitsMaxTresPer) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0038QosLimitsMaxTresPer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0038QosLimitsMaxTresPer(val *Dbv0038QosLimitsMaxTresPer) *NullableDbv0038QosLimitsMaxTresPer {
	return &NullableDbv0038QosLimitsMaxTresPer{value: val, isSet: true}
}

func (v NullableDbv0038QosLimitsMaxTresPer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0038QosLimitsMaxTresPer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
