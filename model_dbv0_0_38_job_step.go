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

// checks if the Dbv0038JobStep type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0038JobStep{}

// Dbv0038JobStep struct for Dbv0038JobStep
type Dbv0038JobStep struct {
	Time     *Dbv0038JobStepTime  `json:"time,omitempty"`
	ExitCode *Dbv0038JobExitCode  `json:"exit_code,omitempty"`
	Nodes    *Dbv0038JobStepNodes `json:"nodes,omitempty"`
	Tasks    *Dbv0038JobStepTasks `json:"tasks,omitempty"`
	// First process PID
	Pid *string            `json:"pid,omitempty"`
	CPU *Dbv0038JobStepCPU `json:"CPU,omitempty"`
	// User who requested job killed
	KillRequestUser *string `json:"kill_request_user,omitempty"`
	// State of job step
	State      *string                   `json:"state,omitempty"`
	Statistics *Dbv0038JobStepStatistics `json:"statistics,omitempty"`
	Step       *Dbv0038JobStepStep       `json:"step,omitempty"`
	// Task distribution properties
	Task *string             `json:"task,omitempty"`
	Tres *Dbv0038JobStepTres `json:"tres,omitempty"`
}

// NewDbv0038JobStep instantiates a new Dbv0038JobStep object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0038JobStep() *Dbv0038JobStep {
	this := Dbv0038JobStep{}
	return &this
}

// NewDbv0038JobStepWithDefaults instantiates a new Dbv0038JobStep object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0038JobStepWithDefaults() *Dbv0038JobStep {
	this := Dbv0038JobStep{}
	return &this
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *Dbv0038JobStep) GetTime() Dbv0038JobStepTime {
	if o == nil || IsNil(o.Time) {
		var ret Dbv0038JobStepTime
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038JobStep) GetTimeOk() (*Dbv0038JobStepTime, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *Dbv0038JobStep) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given Dbv0038JobStepTime and assigns it to the Time field.
func (o *Dbv0038JobStep) SetTime(v Dbv0038JobStepTime) {
	o.Time = &v
}

// GetExitCode returns the ExitCode field value if set, zero value otherwise.
func (o *Dbv0038JobStep) GetExitCode() Dbv0038JobExitCode {
	if o == nil || IsNil(o.ExitCode) {
		var ret Dbv0038JobExitCode
		return ret
	}
	return *o.ExitCode
}

// GetExitCodeOk returns a tuple with the ExitCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038JobStep) GetExitCodeOk() (*Dbv0038JobExitCode, bool) {
	if o == nil || IsNil(o.ExitCode) {
		return nil, false
	}
	return o.ExitCode, true
}

// HasExitCode returns a boolean if a field has been set.
func (o *Dbv0038JobStep) HasExitCode() bool {
	if o != nil && !IsNil(o.ExitCode) {
		return true
	}

	return false
}

// SetExitCode gets a reference to the given Dbv0038JobExitCode and assigns it to the ExitCode field.
func (o *Dbv0038JobStep) SetExitCode(v Dbv0038JobExitCode) {
	o.ExitCode = &v
}

// GetNodes returns the Nodes field value if set, zero value otherwise.
func (o *Dbv0038JobStep) GetNodes() Dbv0038JobStepNodes {
	if o == nil || IsNil(o.Nodes) {
		var ret Dbv0038JobStepNodes
		return ret
	}
	return *o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038JobStep) GetNodesOk() (*Dbv0038JobStepNodes, bool) {
	if o == nil || IsNil(o.Nodes) {
		return nil, false
	}
	return o.Nodes, true
}

// HasNodes returns a boolean if a field has been set.
func (o *Dbv0038JobStep) HasNodes() bool {
	if o != nil && !IsNil(o.Nodes) {
		return true
	}

	return false
}

// SetNodes gets a reference to the given Dbv0038JobStepNodes and assigns it to the Nodes field.
func (o *Dbv0038JobStep) SetNodes(v Dbv0038JobStepNodes) {
	o.Nodes = &v
}

// GetTasks returns the Tasks field value if set, zero value otherwise.
func (o *Dbv0038JobStep) GetTasks() Dbv0038JobStepTasks {
	if o == nil || IsNil(o.Tasks) {
		var ret Dbv0038JobStepTasks
		return ret
	}
	return *o.Tasks
}

// GetTasksOk returns a tuple with the Tasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038JobStep) GetTasksOk() (*Dbv0038JobStepTasks, bool) {
	if o == nil || IsNil(o.Tasks) {
		return nil, false
	}
	return o.Tasks, true
}

// HasTasks returns a boolean if a field has been set.
func (o *Dbv0038JobStep) HasTasks() bool {
	if o != nil && !IsNil(o.Tasks) {
		return true
	}

	return false
}

// SetTasks gets a reference to the given Dbv0038JobStepTasks and assigns it to the Tasks field.
func (o *Dbv0038JobStep) SetTasks(v Dbv0038JobStepTasks) {
	o.Tasks = &v
}

// GetPid returns the Pid field value if set, zero value otherwise.
func (o *Dbv0038JobStep) GetPid() string {
	if o == nil || IsNil(o.Pid) {
		var ret string
		return ret
	}
	return *o.Pid
}

// GetPidOk returns a tuple with the Pid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038JobStep) GetPidOk() (*string, bool) {
	if o == nil || IsNil(o.Pid) {
		return nil, false
	}
	return o.Pid, true
}

// HasPid returns a boolean if a field has been set.
func (o *Dbv0038JobStep) HasPid() bool {
	if o != nil && !IsNil(o.Pid) {
		return true
	}

	return false
}

// SetPid gets a reference to the given string and assigns it to the Pid field.
func (o *Dbv0038JobStep) SetPid(v string) {
	o.Pid = &v
}

// GetCPU returns the CPU field value if set, zero value otherwise.
func (o *Dbv0038JobStep) GetCPU() Dbv0038JobStepCPU {
	if o == nil || IsNil(o.CPU) {
		var ret Dbv0038JobStepCPU
		return ret
	}
	return *o.CPU
}

// GetCPUOk returns a tuple with the CPU field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038JobStep) GetCPUOk() (*Dbv0038JobStepCPU, bool) {
	if o == nil || IsNil(o.CPU) {
		return nil, false
	}
	return o.CPU, true
}

// HasCPU returns a boolean if a field has been set.
func (o *Dbv0038JobStep) HasCPU() bool {
	if o != nil && !IsNil(o.CPU) {
		return true
	}

	return false
}

// SetCPU gets a reference to the given Dbv0038JobStepCPU and assigns it to the CPU field.
func (o *Dbv0038JobStep) SetCPU(v Dbv0038JobStepCPU) {
	o.CPU = &v
}

// GetKillRequestUser returns the KillRequestUser field value if set, zero value otherwise.
func (o *Dbv0038JobStep) GetKillRequestUser() string {
	if o == nil || IsNil(o.KillRequestUser) {
		var ret string
		return ret
	}
	return *o.KillRequestUser
}

// GetKillRequestUserOk returns a tuple with the KillRequestUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038JobStep) GetKillRequestUserOk() (*string, bool) {
	if o == nil || IsNil(o.KillRequestUser) {
		return nil, false
	}
	return o.KillRequestUser, true
}

// HasKillRequestUser returns a boolean if a field has been set.
func (o *Dbv0038JobStep) HasKillRequestUser() bool {
	if o != nil && !IsNil(o.KillRequestUser) {
		return true
	}

	return false
}

// SetKillRequestUser gets a reference to the given string and assigns it to the KillRequestUser field.
func (o *Dbv0038JobStep) SetKillRequestUser(v string) {
	o.KillRequestUser = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Dbv0038JobStep) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038JobStep) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Dbv0038JobStep) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *Dbv0038JobStep) SetState(v string) {
	o.State = &v
}

// GetStatistics returns the Statistics field value if set, zero value otherwise.
func (o *Dbv0038JobStep) GetStatistics() Dbv0038JobStepStatistics {
	if o == nil || IsNil(o.Statistics) {
		var ret Dbv0038JobStepStatistics
		return ret
	}
	return *o.Statistics
}

// GetStatisticsOk returns a tuple with the Statistics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038JobStep) GetStatisticsOk() (*Dbv0038JobStepStatistics, bool) {
	if o == nil || IsNil(o.Statistics) {
		return nil, false
	}
	return o.Statistics, true
}

// HasStatistics returns a boolean if a field has been set.
func (o *Dbv0038JobStep) HasStatistics() bool {
	if o != nil && !IsNil(o.Statistics) {
		return true
	}

	return false
}

// SetStatistics gets a reference to the given Dbv0038JobStepStatistics and assigns it to the Statistics field.
func (o *Dbv0038JobStep) SetStatistics(v Dbv0038JobStepStatistics) {
	o.Statistics = &v
}

// GetStep returns the Step field value if set, zero value otherwise.
func (o *Dbv0038JobStep) GetStep() Dbv0038JobStepStep {
	if o == nil || IsNil(o.Step) {
		var ret Dbv0038JobStepStep
		return ret
	}
	return *o.Step
}

// GetStepOk returns a tuple with the Step field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038JobStep) GetStepOk() (*Dbv0038JobStepStep, bool) {
	if o == nil || IsNil(o.Step) {
		return nil, false
	}
	return o.Step, true
}

// HasStep returns a boolean if a field has been set.
func (o *Dbv0038JobStep) HasStep() bool {
	if o != nil && !IsNil(o.Step) {
		return true
	}

	return false
}

// SetStep gets a reference to the given Dbv0038JobStepStep and assigns it to the Step field.
func (o *Dbv0038JobStep) SetStep(v Dbv0038JobStepStep) {
	o.Step = &v
}

// GetTask returns the Task field value if set, zero value otherwise.
func (o *Dbv0038JobStep) GetTask() string {
	if o == nil || IsNil(o.Task) {
		var ret string
		return ret
	}
	return *o.Task
}

// GetTaskOk returns a tuple with the Task field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038JobStep) GetTaskOk() (*string, bool) {
	if o == nil || IsNil(o.Task) {
		return nil, false
	}
	return o.Task, true
}

// HasTask returns a boolean if a field has been set.
func (o *Dbv0038JobStep) HasTask() bool {
	if o != nil && !IsNil(o.Task) {
		return true
	}

	return false
}

// SetTask gets a reference to the given string and assigns it to the Task field.
func (o *Dbv0038JobStep) SetTask(v string) {
	o.Task = &v
}

// GetTres returns the Tres field value if set, zero value otherwise.
func (o *Dbv0038JobStep) GetTres() Dbv0038JobStepTres {
	if o == nil || IsNil(o.Tres) {
		var ret Dbv0038JobStepTres
		return ret
	}
	return *o.Tres
}

// GetTresOk returns a tuple with the Tres field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038JobStep) GetTresOk() (*Dbv0038JobStepTres, bool) {
	if o == nil || IsNil(o.Tres) {
		return nil, false
	}
	return o.Tres, true
}

// HasTres returns a boolean if a field has been set.
func (o *Dbv0038JobStep) HasTres() bool {
	if o != nil && !IsNil(o.Tres) {
		return true
	}

	return false
}

// SetTres gets a reference to the given Dbv0038JobStepTres and assigns it to the Tres field.
func (o *Dbv0038JobStep) SetTres(v Dbv0038JobStepTres) {
	o.Tres = &v
}

func (o Dbv0038JobStep) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0038JobStep) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !IsNil(o.ExitCode) {
		toSerialize["exit_code"] = o.ExitCode
	}
	if !IsNil(o.Nodes) {
		toSerialize["nodes"] = o.Nodes
	}
	if !IsNil(o.Tasks) {
		toSerialize["tasks"] = o.Tasks
	}
	if !IsNil(o.Pid) {
		toSerialize["pid"] = o.Pid
	}
	if !IsNil(o.CPU) {
		toSerialize["CPU"] = o.CPU
	}
	if !IsNil(o.KillRequestUser) {
		toSerialize["kill_request_user"] = o.KillRequestUser
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Statistics) {
		toSerialize["statistics"] = o.Statistics
	}
	if !IsNil(o.Step) {
		toSerialize["step"] = o.Step
	}
	if !IsNil(o.Task) {
		toSerialize["task"] = o.Task
	}
	if !IsNil(o.Tres) {
		toSerialize["tres"] = o.Tres
	}
	return toSerialize, nil
}

type NullableDbv0038JobStep struct {
	value *Dbv0038JobStep
	isSet bool
}

func (v NullableDbv0038JobStep) Get() *Dbv0038JobStep {
	return v.value
}

func (v *NullableDbv0038JobStep) Set(val *Dbv0038JobStep) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0038JobStep) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0038JobStep) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0038JobStep(val *Dbv0038JobStep) *NullableDbv0038JobStep {
	return &NullableDbv0038JobStep{value: val, isSet: true}
}

func (v NullableDbv0038JobStep) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0038JobStep) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
