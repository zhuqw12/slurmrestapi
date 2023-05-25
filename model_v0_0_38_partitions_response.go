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

// checks if the V0038PartitionsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0038PartitionsResponse{}

// V0038PartitionsResponse struct for V0038PartitionsResponse
type V0038PartitionsResponse struct {
	Meta *V0038Meta `json:"meta,omitempty"`
	// slurm errors
	Errors []V0038Error `json:"errors,omitempty"`
	// partition info
	Partitions []V0038Partition `json:"partitions,omitempty"`
}

// NewV0038PartitionsResponse instantiates a new V0038PartitionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0038PartitionsResponse() *V0038PartitionsResponse {
	this := V0038PartitionsResponse{}
	return &this
}

// NewV0038PartitionsResponseWithDefaults instantiates a new V0038PartitionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0038PartitionsResponseWithDefaults() *V0038PartitionsResponse {
	this := V0038PartitionsResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *V0038PartitionsResponse) GetMeta() V0038Meta {
	if o == nil || IsNil(o.Meta) {
		var ret V0038Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0038PartitionsResponse) GetMetaOk() (*V0038Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *V0038PartitionsResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given V0038Meta and assigns it to the Meta field.
func (o *V0038PartitionsResponse) SetMeta(v V0038Meta) {
	o.Meta = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *V0038PartitionsResponse) GetErrors() []V0038Error {
	if o == nil || IsNil(o.Errors) {
		var ret []V0038Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0038PartitionsResponse) GetErrorsOk() ([]V0038Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *V0038PartitionsResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []V0038Error and assigns it to the Errors field.
func (o *V0038PartitionsResponse) SetErrors(v []V0038Error) {
	o.Errors = v
}

// GetPartitions returns the Partitions field value if set, zero value otherwise.
func (o *V0038PartitionsResponse) GetPartitions() []V0038Partition {
	if o == nil || IsNil(o.Partitions) {
		var ret []V0038Partition
		return ret
	}
	return o.Partitions
}

// GetPartitionsOk returns a tuple with the Partitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0038PartitionsResponse) GetPartitionsOk() ([]V0038Partition, bool) {
	if o == nil || IsNil(o.Partitions) {
		return nil, false
	}
	return o.Partitions, true
}

// HasPartitions returns a boolean if a field has been set.
func (o *V0038PartitionsResponse) HasPartitions() bool {
	if o != nil && !IsNil(o.Partitions) {
		return true
	}

	return false
}

// SetPartitions gets a reference to the given []V0038Partition and assigns it to the Partitions field.
func (o *V0038PartitionsResponse) SetPartitions(v []V0038Partition) {
	o.Partitions = v
}

func (o V0038PartitionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0038PartitionsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !IsNil(o.Partitions) {
		toSerialize["partitions"] = o.Partitions
	}
	return toSerialize, nil
}

type NullableV0038PartitionsResponse struct {
	value *V0038PartitionsResponse
	isSet bool
}

func (v NullableV0038PartitionsResponse) Get() *V0038PartitionsResponse {
	return v.value
}

func (v *NullableV0038PartitionsResponse) Set(val *V0038PartitionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableV0038PartitionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableV0038PartitionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0038PartitionsResponse(val *V0038PartitionsResponse) *NullableV0038PartitionsResponse {
	return &NullableV0038PartitionsResponse{value: val, isSet: true}
}

func (v NullableV0038PartitionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0038PartitionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
