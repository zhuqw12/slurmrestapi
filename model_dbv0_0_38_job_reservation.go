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

// checks if the Dbv0038JobReservation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0038JobReservation{}

// Dbv0038JobReservation Reservation usage details
type Dbv0038JobReservation struct {
	// Database id of reservation
	Id *int32 `json:"id,omitempty"`
	// Name of reservation
	Name *int32 `json:"name,omitempty"`
}

// NewDbv0038JobReservation instantiates a new Dbv0038JobReservation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0038JobReservation() *Dbv0038JobReservation {
	this := Dbv0038JobReservation{}
	return &this
}

// NewDbv0038JobReservationWithDefaults instantiates a new Dbv0038JobReservation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0038JobReservationWithDefaults() *Dbv0038JobReservation {
	this := Dbv0038JobReservation{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Dbv0038JobReservation) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038JobReservation) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Dbv0038JobReservation) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Dbv0038JobReservation) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Dbv0038JobReservation) GetName() int32 {
	if o == nil || IsNil(o.Name) {
		var ret int32
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038JobReservation) GetNameOk() (*int32, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Dbv0038JobReservation) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given int32 and assigns it to the Name field.
func (o *Dbv0038JobReservation) SetName(v int32) {
	o.Name = &v
}

func (o Dbv0038JobReservation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0038JobReservation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableDbv0038JobReservation struct {
	value *Dbv0038JobReservation
	isSet bool
}

func (v NullableDbv0038JobReservation) Get() *Dbv0038JobReservation {
	return v.value
}

func (v *NullableDbv0038JobReservation) Set(val *Dbv0038JobReservation) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0038JobReservation) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0038JobReservation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0038JobReservation(val *Dbv0038JobReservation) *NullableDbv0038JobReservation {
	return &NullableDbv0038JobReservation{value: val, isSet: true}
}

func (v NullableDbv0038JobReservation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0038JobReservation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}