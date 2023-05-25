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

// checks if the V0039StatsMsgRpcsByUserInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0039StatsMsgRpcsByUserInner{}

// V0039StatsMsgRpcsByUserInner user
type V0039StatsMsgRpcsByUserInner struct {
	// user name
	User *string `json:"user,omitempty"`
	// user id (numeric)
	UserId *int32 `json:"user_id,omitempty"`
	// Number of RPCs received
	Count *int64 `json:"count,omitempty"`
	// Average time spent processing RPC in seconds
	AverageTime *int64 `json:"average_time,omitempty"`
	// Total time spent processing RPC in seconds
	TotalTime *int64 `json:"total_time,omitempty"`
}

// NewV0039StatsMsgRpcsByUserInner instantiates a new V0039StatsMsgRpcsByUserInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0039StatsMsgRpcsByUserInner() *V0039StatsMsgRpcsByUserInner {
	this := V0039StatsMsgRpcsByUserInner{}
	return &this
}

// NewV0039StatsMsgRpcsByUserInnerWithDefaults instantiates a new V0039StatsMsgRpcsByUserInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0039StatsMsgRpcsByUserInnerWithDefaults() *V0039StatsMsgRpcsByUserInner {
	this := V0039StatsMsgRpcsByUserInner{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *V0039StatsMsgRpcsByUserInner) GetUser() string {
	if o == nil || IsNil(o.User) {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039StatsMsgRpcsByUserInner) GetUserOk() (*string, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *V0039StatsMsgRpcsByUserInner) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *V0039StatsMsgRpcsByUserInner) SetUser(v string) {
	o.User = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *V0039StatsMsgRpcsByUserInner) GetUserId() int32 {
	if o == nil || IsNil(o.UserId) {
		var ret int32
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039StatsMsgRpcsByUserInner) GetUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *V0039StatsMsgRpcsByUserInner) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int32 and assigns it to the UserId field.
func (o *V0039StatsMsgRpcsByUserInner) SetUserId(v int32) {
	o.UserId = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *V0039StatsMsgRpcsByUserInner) GetCount() int64 {
	if o == nil || IsNil(o.Count) {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039StatsMsgRpcsByUserInner) GetCountOk() (*int64, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *V0039StatsMsgRpcsByUserInner) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *V0039StatsMsgRpcsByUserInner) SetCount(v int64) {
	o.Count = &v
}

// GetAverageTime returns the AverageTime field value if set, zero value otherwise.
func (o *V0039StatsMsgRpcsByUserInner) GetAverageTime() int64 {
	if o == nil || IsNil(o.AverageTime) {
		var ret int64
		return ret
	}
	return *o.AverageTime
}

// GetAverageTimeOk returns a tuple with the AverageTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039StatsMsgRpcsByUserInner) GetAverageTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.AverageTime) {
		return nil, false
	}
	return o.AverageTime, true
}

// HasAverageTime returns a boolean if a field has been set.
func (o *V0039StatsMsgRpcsByUserInner) HasAverageTime() bool {
	if o != nil && !IsNil(o.AverageTime) {
		return true
	}

	return false
}

// SetAverageTime gets a reference to the given int64 and assigns it to the AverageTime field.
func (o *V0039StatsMsgRpcsByUserInner) SetAverageTime(v int64) {
	o.AverageTime = &v
}

// GetTotalTime returns the TotalTime field value if set, zero value otherwise.
func (o *V0039StatsMsgRpcsByUserInner) GetTotalTime() int64 {
	if o == nil || IsNil(o.TotalTime) {
		var ret int64
		return ret
	}
	return *o.TotalTime
}

// GetTotalTimeOk returns a tuple with the TotalTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039StatsMsgRpcsByUserInner) GetTotalTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalTime) {
		return nil, false
	}
	return o.TotalTime, true
}

// HasTotalTime returns a boolean if a field has been set.
func (o *V0039StatsMsgRpcsByUserInner) HasTotalTime() bool {
	if o != nil && !IsNil(o.TotalTime) {
		return true
	}

	return false
}

// SetTotalTime gets a reference to the given int64 and assigns it to the TotalTime field.
func (o *V0039StatsMsgRpcsByUserInner) SetTotalTime(v int64) {
	o.TotalTime = &v
}

func (o V0039StatsMsgRpcsByUserInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0039StatsMsgRpcsByUserInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.AverageTime) {
		toSerialize["average_time"] = o.AverageTime
	}
	if !IsNil(o.TotalTime) {
		toSerialize["total_time"] = o.TotalTime
	}
	return toSerialize, nil
}

type NullableV0039StatsMsgRpcsByUserInner struct {
	value *V0039StatsMsgRpcsByUserInner
	isSet bool
}

func (v NullableV0039StatsMsgRpcsByUserInner) Get() *V0039StatsMsgRpcsByUserInner {
	return v.value
}

func (v *NullableV0039StatsMsgRpcsByUserInner) Set(val *V0039StatsMsgRpcsByUserInner) {
	v.value = val
	v.isSet = true
}

func (v NullableV0039StatsMsgRpcsByUserInner) IsSet() bool {
	return v.isSet
}

func (v *NullableV0039StatsMsgRpcsByUserInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0039StatsMsgRpcsByUserInner(val *V0039StatsMsgRpcsByUserInner) *NullableV0039StatsMsgRpcsByUserInner {
	return &NullableV0039StatsMsgRpcsByUserInner{value: val, isSet: true}
}

func (v NullableV0039StatsMsgRpcsByUserInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0039StatsMsgRpcsByUserInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
