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

// checks if the Dbv0038JobHet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0038JobHet{}

// Dbv0038JobHet Heterogeneous Job details (optional)
type Dbv0038JobHet struct {
	// Parent HetJob id
	JobId *int32 `json:"job_id,omitempty"`
	// Offset of this job to parent
	JobOffset *int32 `json:"job_offset,omitempty"`
}

// NewDbv0038JobHet instantiates a new Dbv0038JobHet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0038JobHet() *Dbv0038JobHet {
	this := Dbv0038JobHet{}
	return &this
}

// NewDbv0038JobHetWithDefaults instantiates a new Dbv0038JobHet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0038JobHetWithDefaults() *Dbv0038JobHet {
	this := Dbv0038JobHet{}
	return &this
}

// GetJobId returns the JobId field value if set, zero value otherwise.
func (o *Dbv0038JobHet) GetJobId() int32 {
	if o == nil || IsNil(o.JobId) {
		var ret int32
		return ret
	}
	return *o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038JobHet) GetJobIdOk() (*int32, bool) {
	if o == nil || IsNil(o.JobId) {
		return nil, false
	}
	return o.JobId, true
}

// HasJobId returns a boolean if a field has been set.
func (o *Dbv0038JobHet) HasJobId() bool {
	if o != nil && !IsNil(o.JobId) {
		return true
	}

	return false
}

// SetJobId gets a reference to the given int32 and assigns it to the JobId field.
func (o *Dbv0038JobHet) SetJobId(v int32) {
	o.JobId = &v
}

// GetJobOffset returns the JobOffset field value if set, zero value otherwise.
func (o *Dbv0038JobHet) GetJobOffset() int32 {
	if o == nil || IsNil(o.JobOffset) {
		var ret int32
		return ret
	}
	return *o.JobOffset
}

// GetJobOffsetOk returns a tuple with the JobOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038JobHet) GetJobOffsetOk() (*int32, bool) {
	if o == nil || IsNil(o.JobOffset) {
		return nil, false
	}
	return o.JobOffset, true
}

// HasJobOffset returns a boolean if a field has been set.
func (o *Dbv0038JobHet) HasJobOffset() bool {
	if o != nil && !IsNil(o.JobOffset) {
		return true
	}

	return false
}

// SetJobOffset gets a reference to the given int32 and assigns it to the JobOffset field.
func (o *Dbv0038JobHet) SetJobOffset(v int32) {
	o.JobOffset = &v
}

func (o Dbv0038JobHet) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0038JobHet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.JobId) {
		toSerialize["job_id"] = o.JobId
	}
	if !IsNil(o.JobOffset) {
		toSerialize["job_offset"] = o.JobOffset
	}
	return toSerialize, nil
}

type NullableDbv0038JobHet struct {
	value *Dbv0038JobHet
	isSet bool
}

func (v NullableDbv0038JobHet) Get() *Dbv0038JobHet {
	return v.value
}

func (v *NullableDbv0038JobHet) Set(val *Dbv0038JobHet) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0038JobHet) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0038JobHet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0038JobHet(val *Dbv0038JobHet) *NullableDbv0038JobHet {
	return &NullableDbv0038JobHet{value: val, isSet: true}
}

func (v NullableDbv0038JobHet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0038JobHet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
