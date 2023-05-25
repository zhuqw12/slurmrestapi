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
	"fmt"
)

// V0038Signal POSIX signal name
type V0038Signal string

// List of v0.0.38_signal
/*const (
	HUP V0038Signal = "HUP"
	INT V0038Signal = "INT"
	QUIT V0038Signal = "QUIT"
	ABRT V0038Signal = "ABRT"
	KILL V0038Signal = "KILL"
	ALRM V0038Signal = "ALRM"
	TERM V0038Signal = "TERM"
	USR1 V0038Signal = "USR1"
	USR2 V0038Signal = "USR2"
	URG V0038Signal = "URG"
	CONT V0038Signal = "CONT"
	STOP V0038Signal = "STOP"
	TSTP V0038Signal = "TSTP"
	TTIN V0038Signal = "TTIN"
	TTOU V0038Signal = "TTOU"
)*/

// All allowed values of V0038Signal enum
var AllowedV0038SignalEnumValues = []V0038Signal{
	"HUP",
	"INT",
	"QUIT",
	"ABRT",
	"KILL",
	"ALRM",
	"TERM",
	"USR1",
	"USR2",
	"URG",
	"CONT",
	"STOP",
	"TSTP",
	"TTIN",
	"TTOU",
}

func (v *V0038Signal) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := V0038Signal(value)
	for _, existing := range AllowedV0038SignalEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid V0038Signal", value)
}

// NewV0038SignalFromValue returns a pointer to a valid V0038Signal
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewV0038SignalFromValue(v string) (*V0038Signal, error) {
	ev := V0038Signal(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for V0038Signal: valid values are %v", v, AllowedV0038SignalEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v V0038Signal) IsValid() bool {
	for _, existing := range AllowedV0038SignalEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to v0.0.38_signal value
func (v V0038Signal) Ptr() *V0038Signal {
	return &v
}

type NullableV0038Signal struct {
	value *V0038Signal
	isSet bool
}

func (v NullableV0038Signal) Get() *V0038Signal {
	return v.value
}

func (v *NullableV0038Signal) Set(val *V0038Signal) {
	v.value = val
	v.isSet = true
}

func (v NullableV0038Signal) IsSet() bool {
	return v.isSet
}

func (v *NullableV0038Signal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0038Signal(val *V0038Signal) *NullableV0038Signal {
	return &NullableV0038Signal{value: val, isSet: true}
}

func (v NullableV0038Signal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0038Signal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
