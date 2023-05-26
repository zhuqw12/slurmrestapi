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

// checks if the V0039Assoc type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0039Assoc{}

// V0039Assoc struct for V0039Assoc
type V0039Assoc struct {
	Account       *string            `json:"account,omitempty"`
	Cluster       *string            `json:"cluster,omitempty"`
	Default       *V0039AssocDefault `json:"default,omitempty"`
	Flags         []string           `json:"flags,omitempty"`
	Max           *V0039AssocMax     `json:"max,omitempty"`
	IsDefault     *bool              `json:"is_default,omitempty"`
	Min           *V0039AssocMin     `json:"min,omitempty"`
	ParentAccount *string            `json:"parent_account,omitempty"`
	Partition     *string            `json:"partition,omitempty"`
	Priority      *V0039Uint32NoVal  `json:"priority,omitempty"`
	// List of QOS names
	Qos       []string         `json:"qos,omitempty"`
	SharesRaw *int32           `json:"shares_raw,omitempty"`
	Usage     *V0039AssocUsage `json:"usage,omitempty"`
	User      string           `json:"user"`
}

// NewV0039Assoc instantiates a new V0039Assoc object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0039Assoc(user string) *V0039Assoc {
	this := V0039Assoc{}
	this.User = user
	return &this
}

// NewV0039AssocWithDefaults instantiates a new V0039Assoc object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0039AssocWithDefaults() *V0039Assoc {
	this := V0039Assoc{}
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *V0039Assoc) GetAccount() string {
	if o == nil || IsNil(o.Account) {
		var ret string
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Assoc) GetAccountOk() (*string, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *V0039Assoc) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given string and assigns it to the Account field.
func (o *V0039Assoc) SetAccount(v string) {
	o.Account = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *V0039Assoc) GetCluster() string {
	if o == nil || IsNil(o.Cluster) {
		var ret string
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Assoc) GetClusterOk() (*string, bool) {
	if o == nil || IsNil(o.Cluster) {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *V0039Assoc) HasCluster() bool {
	if o != nil && !IsNil(o.Cluster) {
		return true
	}

	return false
}

// SetCluster gets a reference to the given string and assigns it to the Cluster field.
func (o *V0039Assoc) SetCluster(v string) {
	o.Cluster = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *V0039Assoc) GetDefault() V0039AssocDefault {
	if o == nil || IsNil(o.Default) {
		var ret V0039AssocDefault
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Assoc) GetDefaultOk() (*V0039AssocDefault, bool) {
	if o == nil || IsNil(o.Default) {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *V0039Assoc) HasDefault() bool {
	if o != nil && !IsNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given V0039AssocDefault and assigns it to the Default field.
func (o *V0039Assoc) SetDefault(v V0039AssocDefault) {
	o.Default = &v
}

// GetFlags returns the Flags field value if set, zero value otherwise.
func (o *V0039Assoc) GetFlags() []string {
	if o == nil || IsNil(o.Flags) {
		var ret []string
		return ret
	}
	return o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Assoc) GetFlagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Flags) {
		return nil, false
	}
	return o.Flags, true
}

// HasFlags returns a boolean if a field has been set.
func (o *V0039Assoc) HasFlags() bool {
	if o != nil && !IsNil(o.Flags) {
		return true
	}

	return false
}

// SetFlags gets a reference to the given []string and assigns it to the Flags field.
func (o *V0039Assoc) SetFlags(v []string) {
	o.Flags = v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *V0039Assoc) GetMax() V0039AssocMax {
	if o == nil || IsNil(o.Max) {
		var ret V0039AssocMax
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Assoc) GetMaxOk() (*V0039AssocMax, bool) {
	if o == nil || IsNil(o.Max) {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *V0039Assoc) HasMax() bool {
	if o != nil && !IsNil(o.Max) {
		return true
	}

	return false
}

// SetMax gets a reference to the given V0039AssocMax and assigns it to the Max field.
func (o *V0039Assoc) SetMax(v V0039AssocMax) {
	o.Max = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *V0039Assoc) GetIsDefault() bool {
	if o == nil || IsNil(o.IsDefault) {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Assoc) GetIsDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDefault) {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *V0039Assoc) HasIsDefault() bool {
	if o != nil && !IsNil(o.IsDefault) {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *V0039Assoc) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *V0039Assoc) GetMin() V0039AssocMin {
	if o == nil || IsNil(o.Min) {
		var ret V0039AssocMin
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Assoc) GetMinOk() (*V0039AssocMin, bool) {
	if o == nil || IsNil(o.Min) {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *V0039Assoc) HasMin() bool {
	if o != nil && !IsNil(o.Min) {
		return true
	}

	return false
}

// SetMin gets a reference to the given V0039AssocMin and assigns it to the Min field.
func (o *V0039Assoc) SetMin(v V0039AssocMin) {
	o.Min = &v
}

// GetParentAccount returns the ParentAccount field value if set, zero value otherwise.
func (o *V0039Assoc) GetParentAccount() string {
	if o == nil || IsNil(o.ParentAccount) {
		var ret string
		return ret
	}
	return *o.ParentAccount
}

// GetParentAccountOk returns a tuple with the ParentAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Assoc) GetParentAccountOk() (*string, bool) {
	if o == nil || IsNil(o.ParentAccount) {
		return nil, false
	}
	return o.ParentAccount, true
}

// HasParentAccount returns a boolean if a field has been set.
func (o *V0039Assoc) HasParentAccount() bool {
	if o != nil && !IsNil(o.ParentAccount) {
		return true
	}

	return false
}

// SetParentAccount gets a reference to the given string and assigns it to the ParentAccount field.
func (o *V0039Assoc) SetParentAccount(v string) {
	o.ParentAccount = &v
}

// GetPartition returns the Partition field value if set, zero value otherwise.
func (o *V0039Assoc) GetPartition() string {
	if o == nil || IsNil(o.Partition) {
		var ret string
		return ret
	}
	return *o.Partition
}

// GetPartitionOk returns a tuple with the Partition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Assoc) GetPartitionOk() (*string, bool) {
	if o == nil || IsNil(o.Partition) {
		return nil, false
	}
	return o.Partition, true
}

// HasPartition returns a boolean if a field has been set.
func (o *V0039Assoc) HasPartition() bool {
	if o != nil && !IsNil(o.Partition) {
		return true
	}

	return false
}

// SetPartition gets a reference to the given string and assigns it to the Partition field.
func (o *V0039Assoc) SetPartition(v string) {
	o.Partition = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *V0039Assoc) GetPriority() V0039Uint32NoVal {
	if o == nil || IsNil(o.Priority) {
		var ret V0039Uint32NoVal
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Assoc) GetPriorityOk() (*V0039Uint32NoVal, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *V0039Assoc) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given V0039Uint32NoVal and assigns it to the Priority field.
func (o *V0039Assoc) SetPriority(v V0039Uint32NoVal) {
	o.Priority = &v
}

// GetQos returns the Qos field value if set, zero value otherwise.
func (o *V0039Assoc) GetQos() []string {
	if o == nil || IsNil(o.Qos) {
		var ret []string
		return ret
	}
	return o.Qos
}

// GetQosOk returns a tuple with the Qos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Assoc) GetQosOk() ([]string, bool) {
	if o == nil || IsNil(o.Qos) {
		return nil, false
	}
	return o.Qos, true
}

// HasQos returns a boolean if a field has been set.
func (o *V0039Assoc) HasQos() bool {
	if o != nil && !IsNil(o.Qos) {
		return true
	}

	return false
}

// SetQos gets a reference to the given []string and assigns it to the Qos field.
func (o *V0039Assoc) SetQos(v []string) {
	o.Qos = v
}

// GetSharesRaw returns the SharesRaw field value if set, zero value otherwise.
func (o *V0039Assoc) GetSharesRaw() int32 {
	if o == nil || IsNil(o.SharesRaw) {
		var ret int32
		return ret
	}
	return *o.SharesRaw
}

// GetSharesRawOk returns a tuple with the SharesRaw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Assoc) GetSharesRawOk() (*int32, bool) {
	if o == nil || IsNil(o.SharesRaw) {
		return nil, false
	}
	return o.SharesRaw, true
}

// HasSharesRaw returns a boolean if a field has been set.
func (o *V0039Assoc) HasSharesRaw() bool {
	if o != nil && !IsNil(o.SharesRaw) {
		return true
	}

	return false
}

// SetSharesRaw gets a reference to the given int32 and assigns it to the SharesRaw field.
func (o *V0039Assoc) SetSharesRaw(v int32) {
	o.SharesRaw = &v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *V0039Assoc) GetUsage() V0039AssocUsage {
	if o == nil || IsNil(o.Usage) {
		var ret V0039AssocUsage
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Assoc) GetUsageOk() (*V0039AssocUsage, bool) {
	if o == nil || IsNil(o.Usage) {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *V0039Assoc) HasUsage() bool {
	if o != nil && !IsNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given V0039AssocUsage and assigns it to the Usage field.
func (o *V0039Assoc) SetUsage(v V0039AssocUsage) {
	o.Usage = &v
}

// GetUser returns the User field value
func (o *V0039Assoc) GetUser() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *V0039Assoc) GetUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *V0039Assoc) SetUser(v string) {
	o.User = v
}

func (o V0039Assoc) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0039Assoc) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	if !IsNil(o.Cluster) {
		toSerialize["cluster"] = o.Cluster
	}
	if !IsNil(o.Default) {
		toSerialize["default"] = o.Default
	}
	if !IsNil(o.Flags) {
		toSerialize["flags"] = o.Flags
	}
	if !IsNil(o.Max) {
		toSerialize["max"] = o.Max
	}
	if !IsNil(o.IsDefault) {
		toSerialize["is_default"] = o.IsDefault
	}
	if !IsNil(o.Min) {
		toSerialize["min"] = o.Min
	}
	if !IsNil(o.ParentAccount) {
		toSerialize["parent_account"] = o.ParentAccount
	}
	if !IsNil(o.Partition) {
		toSerialize["partition"] = o.Partition
	}
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !IsNil(o.Qos) {
		toSerialize["qos"] = o.Qos
	}
	if !IsNil(o.SharesRaw) {
		toSerialize["shares_raw"] = o.SharesRaw
	}
	if !IsNil(o.Usage) {
		toSerialize["usage"] = o.Usage
	}
	toSerialize["user"] = o.User
	return toSerialize, nil
}

type NullableV0039Assoc struct {
	value *V0039Assoc
	isSet bool
}

func (v NullableV0039Assoc) Get() *V0039Assoc {
	return v.value
}

func (v *NullableV0039Assoc) Set(val *V0039Assoc) {
	v.value = val
	v.isSet = true
}

func (v NullableV0039Assoc) IsSet() bool {
	return v.isSet
}

func (v *NullableV0039Assoc) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0039Assoc(val *V0039Assoc) *NullableV0039Assoc {
	return &NullableV0039Assoc{value: val, isSet: true}
}

func (v NullableV0039Assoc) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0039Assoc) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
