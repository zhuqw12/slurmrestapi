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

// checks if the Dbv0038Error type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0038Error{}

// Dbv0038Error struct for Dbv0038Error
type Dbv0038Error struct {
	// Slurm internal error number
	ErrorNumber *int32 `json:"error_number,omitempty"`
	// Error message
	Error *string `json:"error,omitempty"`
	// Where error occurred in the source
	Source *string `json:"source,omitempty"`
	// Explaination of cause of error
	Description *string `json:"description,omitempty"`
}

// NewDbv0038Error instantiates a new Dbv0038Error object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0038Error() *Dbv0038Error {
	this := Dbv0038Error{}
	return &this
}

// NewDbv0038ErrorWithDefaults instantiates a new Dbv0038Error object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0038ErrorWithDefaults() *Dbv0038Error {
	this := Dbv0038Error{}
	return &this
}

// GetErrorNumber returns the ErrorNumber field value if set, zero value otherwise.
func (o *Dbv0038Error) GetErrorNumber() int32 {
	if o == nil || IsNil(o.ErrorNumber) {
		var ret int32
		return ret
	}
	return *o.ErrorNumber
}

// GetErrorNumberOk returns a tuple with the ErrorNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038Error) GetErrorNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.ErrorNumber) {
		return nil, false
	}
	return o.ErrorNumber, true
}

// HasErrorNumber returns a boolean if a field has been set.
func (o *Dbv0038Error) HasErrorNumber() bool {
	if o != nil && !IsNil(o.ErrorNumber) {
		return true
	}

	return false
}

// SetErrorNumber gets a reference to the given int32 and assigns it to the ErrorNumber field.
func (o *Dbv0038Error) SetErrorNumber(v int32) {
	o.ErrorNumber = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *Dbv0038Error) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038Error) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *Dbv0038Error) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *Dbv0038Error) SetError(v string) {
	o.Error = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *Dbv0038Error) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038Error) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *Dbv0038Error) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *Dbv0038Error) SetSource(v string) {
	o.Source = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Dbv0038Error) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038Error) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Dbv0038Error) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Dbv0038Error) SetDescription(v string) {
	o.Description = &v
}

func (o Dbv0038Error) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0038Error) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ErrorNumber) {
		toSerialize["error_number"] = o.ErrorNumber
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableDbv0038Error struct {
	value *Dbv0038Error
	isSet bool
}

func (v NullableDbv0038Error) Get() *Dbv0038Error {
	return v.value
}

func (v *NullableDbv0038Error) Set(val *Dbv0038Error) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0038Error) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0038Error) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0038Error(val *Dbv0038Error) *NullableDbv0038Error {
	return &NullableDbv0038Error{value: val, isSet: true}
}

func (v NullableDbv0038Error) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0038Error) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
