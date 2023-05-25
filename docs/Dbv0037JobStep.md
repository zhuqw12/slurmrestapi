# Dbv0037JobStep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | Pointer to [**Dbv0038JobStepTime**](Dbv0038JobStepTime.md) |  | [optional] 
**ExitCode** | Pointer to [**Dbv0037JobExitCode**](Dbv0037JobExitCode.md) |  | [optional] 
**Nodes** | Pointer to [**Dbv0038JobStepNodes**](Dbv0038JobStepNodes.md) |  | [optional] 
**Tasks** | Pointer to [**Dbv0038JobStepTasks**](Dbv0038JobStepTasks.md) |  | [optional] 
**Pid** | Pointer to **string** | First process PID | [optional] 
**CPU** | Pointer to [**Dbv0038JobStepCPU**](Dbv0038JobStepCPU.md) |  | [optional] 
**KillRequestUser** | Pointer to **string** | User who requested job killed | [optional] 
**State** | Pointer to **string** | State of job step | [optional] 
**Statistics** | Pointer to [**Dbv0038JobStepStatistics**](Dbv0038JobStepStatistics.md) |  | [optional] 
**Step** | Pointer to [**Dbv0038JobStepStep**](Dbv0038JobStepStep.md) |  | [optional] 
**Task** | Pointer to [**Dbv0037JobStepTask**](Dbv0037JobStepTask.md) |  | [optional] 
**Tres** | Pointer to [**Dbv0037JobStepTres**](Dbv0037JobStepTres.md) |  | [optional] 

## Methods

### NewDbv0037JobStep

`func NewDbv0037JobStep() *Dbv0037JobStep`

NewDbv0037JobStep instantiates a new Dbv0037JobStep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbv0037JobStepWithDefaults

`func NewDbv0037JobStepWithDefaults() *Dbv0037JobStep`

NewDbv0037JobStepWithDefaults instantiates a new Dbv0037JobStep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *Dbv0037JobStep) GetTime() Dbv0038JobStepTime`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *Dbv0037JobStep) GetTimeOk() (*Dbv0038JobStepTime, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *Dbv0037JobStep) SetTime(v Dbv0038JobStepTime)`

SetTime sets Time field to given value.

### HasTime

`func (o *Dbv0037JobStep) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetExitCode

`func (o *Dbv0037JobStep) GetExitCode() Dbv0037JobExitCode`

GetExitCode returns the ExitCode field if non-nil, zero value otherwise.

### GetExitCodeOk

`func (o *Dbv0037JobStep) GetExitCodeOk() (*Dbv0037JobExitCode, bool)`

GetExitCodeOk returns a tuple with the ExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitCode

`func (o *Dbv0037JobStep) SetExitCode(v Dbv0037JobExitCode)`

SetExitCode sets ExitCode field to given value.

### HasExitCode

`func (o *Dbv0037JobStep) HasExitCode() bool`

HasExitCode returns a boolean if a field has been set.

### GetNodes

`func (o *Dbv0037JobStep) GetNodes() Dbv0038JobStepNodes`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *Dbv0037JobStep) GetNodesOk() (*Dbv0038JobStepNodes, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *Dbv0037JobStep) SetNodes(v Dbv0038JobStepNodes)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *Dbv0037JobStep) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetTasks

`func (o *Dbv0037JobStep) GetTasks() Dbv0038JobStepTasks`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *Dbv0037JobStep) GetTasksOk() (*Dbv0038JobStepTasks, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *Dbv0037JobStep) SetTasks(v Dbv0038JobStepTasks)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *Dbv0037JobStep) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetPid

`func (o *Dbv0037JobStep) GetPid() string`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *Dbv0037JobStep) GetPidOk() (*string, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *Dbv0037JobStep) SetPid(v string)`

SetPid sets Pid field to given value.

### HasPid

`func (o *Dbv0037JobStep) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetCPU

`func (o *Dbv0037JobStep) GetCPU() Dbv0038JobStepCPU`

GetCPU returns the CPU field if non-nil, zero value otherwise.

### GetCPUOk

`func (o *Dbv0037JobStep) GetCPUOk() (*Dbv0038JobStepCPU, bool)`

GetCPUOk returns a tuple with the CPU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCPU

`func (o *Dbv0037JobStep) SetCPU(v Dbv0038JobStepCPU)`

SetCPU sets CPU field to given value.

### HasCPU

`func (o *Dbv0037JobStep) HasCPU() bool`

HasCPU returns a boolean if a field has been set.

### GetKillRequestUser

`func (o *Dbv0037JobStep) GetKillRequestUser() string`

GetKillRequestUser returns the KillRequestUser field if non-nil, zero value otherwise.

### GetKillRequestUserOk

`func (o *Dbv0037JobStep) GetKillRequestUserOk() (*string, bool)`

GetKillRequestUserOk returns a tuple with the KillRequestUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKillRequestUser

`func (o *Dbv0037JobStep) SetKillRequestUser(v string)`

SetKillRequestUser sets KillRequestUser field to given value.

### HasKillRequestUser

`func (o *Dbv0037JobStep) HasKillRequestUser() bool`

HasKillRequestUser returns a boolean if a field has been set.

### GetState

`func (o *Dbv0037JobStep) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Dbv0037JobStep) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Dbv0037JobStep) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Dbv0037JobStep) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStatistics

`func (o *Dbv0037JobStep) GetStatistics() Dbv0038JobStepStatistics`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *Dbv0037JobStep) GetStatisticsOk() (*Dbv0038JobStepStatistics, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *Dbv0037JobStep) SetStatistics(v Dbv0038JobStepStatistics)`

SetStatistics sets Statistics field to given value.

### HasStatistics

`func (o *Dbv0037JobStep) HasStatistics() bool`

HasStatistics returns a boolean if a field has been set.

### GetStep

`func (o *Dbv0037JobStep) GetStep() Dbv0038JobStepStep`

GetStep returns the Step field if non-nil, zero value otherwise.

### GetStepOk

`func (o *Dbv0037JobStep) GetStepOk() (*Dbv0038JobStepStep, bool)`

GetStepOk returns a tuple with the Step field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStep

`func (o *Dbv0037JobStep) SetStep(v Dbv0038JobStepStep)`

SetStep sets Step field to given value.

### HasStep

`func (o *Dbv0037JobStep) HasStep() bool`

HasStep returns a boolean if a field has been set.

### GetTask

`func (o *Dbv0037JobStep) GetTask() Dbv0037JobStepTask`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *Dbv0037JobStep) GetTaskOk() (*Dbv0037JobStepTask, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *Dbv0037JobStep) SetTask(v Dbv0037JobStepTask)`

SetTask sets Task field to given value.

### HasTask

`func (o *Dbv0037JobStep) HasTask() bool`

HasTask returns a boolean if a field has been set.

### GetTres

`func (o *Dbv0037JobStep) GetTres() Dbv0037JobStepTres`

GetTres returns the Tres field if non-nil, zero value otherwise.

### GetTresOk

`func (o *Dbv0037JobStep) GetTresOk() (*Dbv0037JobStepTres, bool)`

GetTresOk returns a tuple with the Tres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTres

`func (o *Dbv0037JobStep) SetTres(v Dbv0037JobStepTres)`

SetTres sets Tres field to given value.

### HasTres

`func (o *Dbv0037JobStep) HasTres() bool`

HasTres returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


