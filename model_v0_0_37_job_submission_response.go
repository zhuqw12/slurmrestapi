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

// checks if the V0037JobSubmissionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0037JobSubmissionResponse{}

// V0037JobSubmissionResponse struct for V0037JobSubmissionResponse
type V0037JobSubmissionResponse struct {
	// slurm errors
	Errors []V0037Error `json:"errors,omitempty"`
	// new job ID
	JobId *int32 `json:"job_id,omitempty"`
	// new job step ID
	StepId *string `json:"step_id,omitempty"`
	// Message to user from job_submit plugin
	JobSubmitUserMsg *string `json:"job_submit_user_msg,omitempty"`
}

// NewV0037JobSubmissionResponse instantiates a new V0037JobSubmissionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0037JobSubmissionResponse() *V0037JobSubmissionResponse {
	this := V0037JobSubmissionResponse{}
	return &this
}

// NewV0037JobSubmissionResponseWithDefaults instantiates a new V0037JobSubmissionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0037JobSubmissionResponseWithDefaults() *V0037JobSubmissionResponse {
	this := V0037JobSubmissionResponse{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *V0037JobSubmissionResponse) GetErrors() []V0037Error {
	if o == nil || IsNil(o.Errors) {
		var ret []V0037Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0037JobSubmissionResponse) GetErrorsOk() ([]V0037Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *V0037JobSubmissionResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []V0037Error and assigns it to the Errors field.
func (o *V0037JobSubmissionResponse) SetErrors(v []V0037Error) {
	o.Errors = v
}

// GetJobId returns the JobId field value if set, zero value otherwise.
func (o *V0037JobSubmissionResponse) GetJobId() int32 {
	if o == nil || IsNil(o.JobId) {
		var ret int32
		return ret
	}
	return *o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0037JobSubmissionResponse) GetJobIdOk() (*int32, bool) {
	if o == nil || IsNil(o.JobId) {
		return nil, false
	}
	return o.JobId, true
}

// HasJobId returns a boolean if a field has been set.
func (o *V0037JobSubmissionResponse) HasJobId() bool {
	if o != nil && !IsNil(o.JobId) {
		return true
	}

	return false
}

// SetJobId gets a reference to the given int32 and assigns it to the JobId field.
func (o *V0037JobSubmissionResponse) SetJobId(v int32) {
	o.JobId = &v
}

// GetStepId returns the StepId field value if set, zero value otherwise.
func (o *V0037JobSubmissionResponse) GetStepId() string {
	if o == nil || IsNil(o.StepId) {
		var ret string
		return ret
	}
	return *o.StepId
}

// GetStepIdOk returns a tuple with the StepId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0037JobSubmissionResponse) GetStepIdOk() (*string, bool) {
	if o == nil || IsNil(o.StepId) {
		return nil, false
	}
	return o.StepId, true
}

// HasStepId returns a boolean if a field has been set.
func (o *V0037JobSubmissionResponse) HasStepId() bool {
	if o != nil && !IsNil(o.StepId) {
		return true
	}

	return false
}

// SetStepId gets a reference to the given string and assigns it to the StepId field.
func (o *V0037JobSubmissionResponse) SetStepId(v string) {
	o.StepId = &v
}

// GetJobSubmitUserMsg returns the JobSubmitUserMsg field value if set, zero value otherwise.
func (o *V0037JobSubmissionResponse) GetJobSubmitUserMsg() string {
	if o == nil || IsNil(o.JobSubmitUserMsg) {
		var ret string
		return ret
	}
	return *o.JobSubmitUserMsg
}

// GetJobSubmitUserMsgOk returns a tuple with the JobSubmitUserMsg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0037JobSubmissionResponse) GetJobSubmitUserMsgOk() (*string, bool) {
	if o == nil || IsNil(o.JobSubmitUserMsg) {
		return nil, false
	}
	return o.JobSubmitUserMsg, true
}

// HasJobSubmitUserMsg returns a boolean if a field has been set.
func (o *V0037JobSubmissionResponse) HasJobSubmitUserMsg() bool {
	if o != nil && !IsNil(o.JobSubmitUserMsg) {
		return true
	}

	return false
}

// SetJobSubmitUserMsg gets a reference to the given string and assigns it to the JobSubmitUserMsg field.
func (o *V0037JobSubmissionResponse) SetJobSubmitUserMsg(v string) {
	o.JobSubmitUserMsg = &v
}

func (o V0037JobSubmissionResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0037JobSubmissionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !IsNil(o.JobId) {
		toSerialize["job_id"] = o.JobId
	}
	if !IsNil(o.StepId) {
		toSerialize["step_id"] = o.StepId
	}
	if !IsNil(o.JobSubmitUserMsg) {
		toSerialize["job_submit_user_msg"] = o.JobSubmitUserMsg
	}
	return toSerialize, nil
}

type NullableV0037JobSubmissionResponse struct {
	value *V0037JobSubmissionResponse
	isSet bool
}

func (v NullableV0037JobSubmissionResponse) Get() *V0037JobSubmissionResponse {
	return v.value
}

func (v *NullableV0037JobSubmissionResponse) Set(val *V0037JobSubmissionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableV0037JobSubmissionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableV0037JobSubmissionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0037JobSubmissionResponse(val *V0037JobSubmissionResponse) *NullableV0037JobSubmissionResponse {
	return &NullableV0037JobSubmissionResponse{value: val, isSet: true}
}

func (v NullableV0037JobSubmissionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0037JobSubmissionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
