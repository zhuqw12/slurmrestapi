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

// checks if the Dbv0037AccountResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0037AccountResponse{}

// Dbv0037AccountResponse struct for Dbv0037AccountResponse
type Dbv0037AccountResponse struct {
	// Slurm errors
	Errors []Dbv0037Error `json:"errors,omitempty"`
}

// NewDbv0037AccountResponse instantiates a new Dbv0037AccountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0037AccountResponse() *Dbv0037AccountResponse {
	this := Dbv0037AccountResponse{}
	return &this
}

// NewDbv0037AccountResponseWithDefaults instantiates a new Dbv0037AccountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0037AccountResponseWithDefaults() *Dbv0037AccountResponse {
	this := Dbv0037AccountResponse{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *Dbv0037AccountResponse) GetErrors() []Dbv0037Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Dbv0037Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037AccountResponse) GetErrorsOk() ([]Dbv0037Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *Dbv0037AccountResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Dbv0037Error and assigns it to the Errors field.
func (o *Dbv0037AccountResponse) SetErrors(v []Dbv0037Error) {
	o.Errors = v
}

func (o Dbv0037AccountResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0037AccountResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableDbv0037AccountResponse struct {
	value *Dbv0037AccountResponse
	isSet bool
}

func (v NullableDbv0037AccountResponse) Get() *Dbv0037AccountResponse {
	return v.value
}

func (v *NullableDbv0037AccountResponse) Set(val *Dbv0037AccountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0037AccountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0037AccountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0037AccountResponse(val *Dbv0037AccountResponse) *NullableDbv0037AccountResponse {
	return &NullableDbv0037AccountResponse{value: val, isSet: true}
}

func (v NullableDbv0037AccountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0037AccountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
