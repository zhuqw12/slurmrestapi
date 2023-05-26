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

// checks if the V0039Tres type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0039Tres{}

// V0039Tres struct for V0039Tres
type V0039Tres struct {
	Type  string  `json:"type"`
	Name  *string `json:"name,omitempty"`
	Id    *int32  `json:"id,omitempty"`
	Count *int64  `json:"count,omitempty"`
}

// NewV0039Tres instantiates a new V0039Tres object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0039Tres(type_ string) *V0039Tres {
	this := V0039Tres{}
	this.Type = type_
	return &this
}

// NewV0039TresWithDefaults instantiates a new V0039Tres object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0039TresWithDefaults() *V0039Tres {
	this := V0039Tres{}
	return &this
}

// GetType returns the Type field value
func (o *V0039Tres) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *V0039Tres) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *V0039Tres) SetType(v string) {
	o.Type = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V0039Tres) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Tres) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V0039Tres) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V0039Tres) SetName(v string) {
	o.Name = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *V0039Tres) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Tres) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *V0039Tres) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *V0039Tres) SetId(v int32) {
	o.Id = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *V0039Tres) GetCount() int64 {
	if o == nil || IsNil(o.Count) {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Tres) GetCountOk() (*int64, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *V0039Tres) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *V0039Tres) SetCount(v int64) {
	o.Count = &v
}

func (o V0039Tres) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0039Tres) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	return toSerialize, nil
}

type NullableV0039Tres struct {
	value *V0039Tres
	isSet bool
}

func (v NullableV0039Tres) Get() *V0039Tres {
	return v.value
}

func (v *NullableV0039Tres) Set(val *V0039Tres) {
	v.value = val
	v.isSet = true
}

func (v NullableV0039Tres) IsSet() bool {
	return v.isSet
}

func (v *NullableV0039Tres) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0039Tres(val *V0039Tres) *NullableV0039Tres {
	return &NullableV0039Tres{value: val, isSet: true}
}

func (v NullableV0039Tres) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0039Tres) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
