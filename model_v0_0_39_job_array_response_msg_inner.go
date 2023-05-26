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

// checks if the V0039JobArrayResponseMsgInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0039JobArrayResponseMsgInner{}

// V0039JobArrayResponseMsgInner ArrayJob
type V0039JobArrayResponseMsgInner struct {
	// JobId
	JobId *int32 `json:"job_id,omitempty"`
	// numeric error code
	ErrorCode *int32 `json:"error_code,omitempty"`
	// error code description
	Error *string `json:"error,omitempty"`
	// error message
	Why *string `json:"why,omitempty"`
}

// NewV0039JobArrayResponseMsgInner instantiates a new V0039JobArrayResponseMsgInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0039JobArrayResponseMsgInner() *V0039JobArrayResponseMsgInner {
	this := V0039JobArrayResponseMsgInner{}
	return &this
}

// NewV0039JobArrayResponseMsgInnerWithDefaults instantiates a new V0039JobArrayResponseMsgInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0039JobArrayResponseMsgInnerWithDefaults() *V0039JobArrayResponseMsgInner {
	this := V0039JobArrayResponseMsgInner{}
	return &this
}

// GetJobId returns the JobId field value if set, zero value otherwise.
func (o *V0039JobArrayResponseMsgInner) GetJobId() int32 {
	if o == nil || IsNil(o.JobId) {
		var ret int32
		return ret
	}
	return *o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039JobArrayResponseMsgInner) GetJobIdOk() (*int32, bool) {
	if o == nil || IsNil(o.JobId) {
		return nil, false
	}
	return o.JobId, true
}

// HasJobId returns a boolean if a field has been set.
func (o *V0039JobArrayResponseMsgInner) HasJobId() bool {
	if o != nil && !IsNil(o.JobId) {
		return true
	}

	return false
}

// SetJobId gets a reference to the given int32 and assigns it to the JobId field.
func (o *V0039JobArrayResponseMsgInner) SetJobId(v int32) {
	o.JobId = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *V0039JobArrayResponseMsgInner) GetErrorCode() int32 {
	if o == nil || IsNil(o.ErrorCode) {
		var ret int32
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039JobArrayResponseMsgInner) GetErrorCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *V0039JobArrayResponseMsgInner) HasErrorCode() bool {
	if o != nil && !IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given int32 and assigns it to the ErrorCode field.
func (o *V0039JobArrayResponseMsgInner) SetErrorCode(v int32) {
	o.ErrorCode = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *V0039JobArrayResponseMsgInner) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039JobArrayResponseMsgInner) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *V0039JobArrayResponseMsgInner) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *V0039JobArrayResponseMsgInner) SetError(v string) {
	o.Error = &v
}

// GetWhy returns the Why field value if set, zero value otherwise.
func (o *V0039JobArrayResponseMsgInner) GetWhy() string {
	if o == nil || IsNil(o.Why) {
		var ret string
		return ret
	}
	return *o.Why
}

// GetWhyOk returns a tuple with the Why field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039JobArrayResponseMsgInner) GetWhyOk() (*string, bool) {
	if o == nil || IsNil(o.Why) {
		return nil, false
	}
	return o.Why, true
}

// HasWhy returns a boolean if a field has been set.
func (o *V0039JobArrayResponseMsgInner) HasWhy() bool {
	if o != nil && !IsNil(o.Why) {
		return true
	}

	return false
}

// SetWhy gets a reference to the given string and assigns it to the Why field.
func (o *V0039JobArrayResponseMsgInner) SetWhy(v string) {
	o.Why = &v
}

func (o V0039JobArrayResponseMsgInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0039JobArrayResponseMsgInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.JobId) {
		toSerialize["job_id"] = o.JobId
	}
	if !IsNil(o.ErrorCode) {
		toSerialize["error_code"] = o.ErrorCode
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !IsNil(o.Why) {
		toSerialize["why"] = o.Why
	}
	return toSerialize, nil
}

type NullableV0039JobArrayResponseMsgInner struct {
	value *V0039JobArrayResponseMsgInner
	isSet bool
}

func (v NullableV0039JobArrayResponseMsgInner) Get() *V0039JobArrayResponseMsgInner {
	return v.value
}

func (v *NullableV0039JobArrayResponseMsgInner) Set(val *V0039JobArrayResponseMsgInner) {
	v.value = val
	v.isSet = true
}

func (v NullableV0039JobArrayResponseMsgInner) IsSet() bool {
	return v.isSet
}

func (v *NullableV0039JobArrayResponseMsgInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0039JobArrayResponseMsgInner(val *V0039JobArrayResponseMsgInner) *NullableV0039JobArrayResponseMsgInner {
	return &NullableV0039JobArrayResponseMsgInner{value: val, isSet: true}
}

func (v NullableV0039JobArrayResponseMsgInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0039JobArrayResponseMsgInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
