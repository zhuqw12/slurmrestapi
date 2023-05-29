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

// checks if the Dbv0037QosLimits type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0037QosLimits{}

// Dbv0037QosLimits Assigned limits
type Dbv0037QosLimits struct {
	// factor to apply to TRES count for associations using this QOS
	Factor *float32             `json:"factor,omitempty"`
	Max    *Dbv0037QosLimitsMax `json:"max,omitempty"`
	Min    *Dbv0037QosLimitsMin `json:"min,omitempty"`
}

// NewDbv0037QosLimits instantiates a new Dbv0037QosLimits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0037QosLimits() *Dbv0037QosLimits {
	this := Dbv0037QosLimits{}
	return &this
}

// NewDbv0037QosLimitsWithDefaults instantiates a new Dbv0037QosLimits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0037QosLimitsWithDefaults() *Dbv0037QosLimits {
	this := Dbv0037QosLimits{}
	return &this
}

// GetFactor returns the Factor field value if set, zero value otherwise.
func (o *Dbv0037QosLimits) GetFactor() float32 {
	if o == nil || IsNil(o.Factor) {
		var ret float32
		return ret
	}
	return *o.Factor
}

// GetFactorOk returns a tuple with the Factor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037QosLimits) GetFactorOk() (*float32, bool) {
	if o == nil || IsNil(o.Factor) {
		return nil, false
	}
	return o.Factor, true
}

// HasFactor returns a boolean if a field has been set.
func (o *Dbv0037QosLimits) HasFactor() bool {
	if o != nil && !IsNil(o.Factor) {
		return true
	}

	return false
}

// SetFactor gets a reference to the given float32 and assigns it to the Factor field.
func (o *Dbv0037QosLimits) SetFactor(v float32) {
	o.Factor = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *Dbv0037QosLimits) GetMax() Dbv0037QosLimitsMax {
	if o == nil || IsNil(o.Max) {
		var ret Dbv0037QosLimitsMax
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037QosLimits) GetMaxOk() (*Dbv0037QosLimitsMax, bool) {
	if o == nil || IsNil(o.Max) {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *Dbv0037QosLimits) HasMax() bool {
	if o != nil && !IsNil(o.Max) {
		return true
	}

	return false
}

// SetMax gets a reference to the given Dbv0037QosLimitsMax and assigns it to the Max field.
func (o *Dbv0037QosLimits) SetMax(v Dbv0037QosLimitsMax) {
	o.Max = &v
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *Dbv0037QosLimits) GetMin() Dbv0037QosLimitsMin {
	if o == nil || IsNil(o.Min) {
		var ret Dbv0037QosLimitsMin
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037QosLimits) GetMinOk() (*Dbv0037QosLimitsMin, bool) {
	if o == nil || IsNil(o.Min) {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *Dbv0037QosLimits) HasMin() bool {
	if o != nil && !IsNil(o.Min) {
		return true
	}

	return false
}

// SetMin gets a reference to the given Dbv0037QosLimitsMin and assigns it to the Min field.
func (o *Dbv0037QosLimits) SetMin(v Dbv0037QosLimitsMin) {
	o.Min = &v
}

func (o Dbv0037QosLimits) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0037QosLimits) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Factor) {
		toSerialize["factor"] = o.Factor
	}
	if !IsNil(o.Max) {
		toSerialize["max"] = o.Max
	}
	if !IsNil(o.Min) {
		toSerialize["min"] = o.Min
	}
	return toSerialize, nil
}

type NullableDbv0037QosLimits struct {
	value *Dbv0037QosLimits
	isSet bool
}

func (v NullableDbv0037QosLimits) Get() *Dbv0037QosLimits {
	return v.value
}

func (v *NullableDbv0037QosLimits) Set(val *Dbv0037QosLimits) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0037QosLimits) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0037QosLimits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0037QosLimits(val *Dbv0037QosLimits) *NullableDbv0037QosLimits {
	return &NullableDbv0037QosLimits{value: val, isSet: true}
}

func (v NullableDbv0037QosLimits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0037QosLimits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}