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

// checks if the Dbv0038AssociationUsage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0038AssociationUsage{}

// Dbv0038AssociationUsage Association usage
type Dbv0038AssociationUsage struct {
	// Jobs accuring priority
	AccrueJobCount *int32 `json:"accrue_job_count,omitempty"`
	// Group used wallclock time (s)
	GroupUsedWallclock *float32 `json:"group_used_wallclock,omitempty"`
	// Fairshare factor
	FairshareFactor *float32 `json:"fairshare_factor,omitempty"`
	// Fairshare shares
	FairshareShares *int32 `json:"fairshare_shares,omitempty"`
	// Currently active jobs
	NormalizedPriority *int32 `json:"normalized_priority,omitempty"`
	// Normalized shares
	NormalizedShares *float32 `json:"normalized_shares,omitempty"`
	// Effective normalized usage
	EffectiveNormalizedUsage *float32 `json:"effective_normalized_usage,omitempty"`
	// Raw usage
	RawUsage *int32 `json:"raw_usage,omitempty"`
	// Total jobs submitted
	JobCount *int32 `json:"job_count,omitempty"`
	// Fairshare level
	FairshareLevel *float32 `json:"fairshare_level,omitempty"`
}

// NewDbv0038AssociationUsage instantiates a new Dbv0038AssociationUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0038AssociationUsage() *Dbv0038AssociationUsage {
	this := Dbv0038AssociationUsage{}
	return &this
}

// NewDbv0038AssociationUsageWithDefaults instantiates a new Dbv0038AssociationUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0038AssociationUsageWithDefaults() *Dbv0038AssociationUsage {
	this := Dbv0038AssociationUsage{}
	return &this
}

// GetAccrueJobCount returns the AccrueJobCount field value if set, zero value otherwise.
func (o *Dbv0038AssociationUsage) GetAccrueJobCount() int32 {
	if o == nil || IsNil(o.AccrueJobCount) {
		var ret int32
		return ret
	}
	return *o.AccrueJobCount
}

// GetAccrueJobCountOk returns a tuple with the AccrueJobCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038AssociationUsage) GetAccrueJobCountOk() (*int32, bool) {
	if o == nil || IsNil(o.AccrueJobCount) {
		return nil, false
	}
	return o.AccrueJobCount, true
}

// HasAccrueJobCount returns a boolean if a field has been set.
func (o *Dbv0038AssociationUsage) HasAccrueJobCount() bool {
	if o != nil && !IsNil(o.AccrueJobCount) {
		return true
	}

	return false
}

// SetAccrueJobCount gets a reference to the given int32 and assigns it to the AccrueJobCount field.
func (o *Dbv0038AssociationUsage) SetAccrueJobCount(v int32) {
	o.AccrueJobCount = &v
}

// GetGroupUsedWallclock returns the GroupUsedWallclock field value if set, zero value otherwise.
func (o *Dbv0038AssociationUsage) GetGroupUsedWallclock() float32 {
	if o == nil || IsNil(o.GroupUsedWallclock) {
		var ret float32
		return ret
	}
	return *o.GroupUsedWallclock
}

// GetGroupUsedWallclockOk returns a tuple with the GroupUsedWallclock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038AssociationUsage) GetGroupUsedWallclockOk() (*float32, bool) {
	if o == nil || IsNil(o.GroupUsedWallclock) {
		return nil, false
	}
	return o.GroupUsedWallclock, true
}

// HasGroupUsedWallclock returns a boolean if a field has been set.
func (o *Dbv0038AssociationUsage) HasGroupUsedWallclock() bool {
	if o != nil && !IsNil(o.GroupUsedWallclock) {
		return true
	}

	return false
}

// SetGroupUsedWallclock gets a reference to the given float32 and assigns it to the GroupUsedWallclock field.
func (o *Dbv0038AssociationUsage) SetGroupUsedWallclock(v float32) {
	o.GroupUsedWallclock = &v
}

// GetFairshareFactor returns the FairshareFactor field value if set, zero value otherwise.
func (o *Dbv0038AssociationUsage) GetFairshareFactor() float32 {
	if o == nil || IsNil(o.FairshareFactor) {
		var ret float32
		return ret
	}
	return *o.FairshareFactor
}

// GetFairshareFactorOk returns a tuple with the FairshareFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038AssociationUsage) GetFairshareFactorOk() (*float32, bool) {
	if o == nil || IsNil(o.FairshareFactor) {
		return nil, false
	}
	return o.FairshareFactor, true
}

// HasFairshareFactor returns a boolean if a field has been set.
func (o *Dbv0038AssociationUsage) HasFairshareFactor() bool {
	if o != nil && !IsNil(o.FairshareFactor) {
		return true
	}

	return false
}

// SetFairshareFactor gets a reference to the given float32 and assigns it to the FairshareFactor field.
func (o *Dbv0038AssociationUsage) SetFairshareFactor(v float32) {
	o.FairshareFactor = &v
}

// GetFairshareShares returns the FairshareShares field value if set, zero value otherwise.
func (o *Dbv0038AssociationUsage) GetFairshareShares() int32 {
	if o == nil || IsNil(o.FairshareShares) {
		var ret int32
		return ret
	}
	return *o.FairshareShares
}

// GetFairshareSharesOk returns a tuple with the FairshareShares field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038AssociationUsage) GetFairshareSharesOk() (*int32, bool) {
	if o == nil || IsNil(o.FairshareShares) {
		return nil, false
	}
	return o.FairshareShares, true
}

// HasFairshareShares returns a boolean if a field has been set.
func (o *Dbv0038AssociationUsage) HasFairshareShares() bool {
	if o != nil && !IsNil(o.FairshareShares) {
		return true
	}

	return false
}

// SetFairshareShares gets a reference to the given int32 and assigns it to the FairshareShares field.
func (o *Dbv0038AssociationUsage) SetFairshareShares(v int32) {
	o.FairshareShares = &v
}

// GetNormalizedPriority returns the NormalizedPriority field value if set, zero value otherwise.
func (o *Dbv0038AssociationUsage) GetNormalizedPriority() int32 {
	if o == nil || IsNil(o.NormalizedPriority) {
		var ret int32
		return ret
	}
	return *o.NormalizedPriority
}

// GetNormalizedPriorityOk returns a tuple with the NormalizedPriority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038AssociationUsage) GetNormalizedPriorityOk() (*int32, bool) {
	if o == nil || IsNil(o.NormalizedPriority) {
		return nil, false
	}
	return o.NormalizedPriority, true
}

// HasNormalizedPriority returns a boolean if a field has been set.
func (o *Dbv0038AssociationUsage) HasNormalizedPriority() bool {
	if o != nil && !IsNil(o.NormalizedPriority) {
		return true
	}

	return false
}

// SetNormalizedPriority gets a reference to the given int32 and assigns it to the NormalizedPriority field.
func (o *Dbv0038AssociationUsage) SetNormalizedPriority(v int32) {
	o.NormalizedPriority = &v
}

// GetNormalizedShares returns the NormalizedShares field value if set, zero value otherwise.
func (o *Dbv0038AssociationUsage) GetNormalizedShares() float32 {
	if o == nil || IsNil(o.NormalizedShares) {
		var ret float32
		return ret
	}
	return *o.NormalizedShares
}

// GetNormalizedSharesOk returns a tuple with the NormalizedShares field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038AssociationUsage) GetNormalizedSharesOk() (*float32, bool) {
	if o == nil || IsNil(o.NormalizedShares) {
		return nil, false
	}
	return o.NormalizedShares, true
}

// HasNormalizedShares returns a boolean if a field has been set.
func (o *Dbv0038AssociationUsage) HasNormalizedShares() bool {
	if o != nil && !IsNil(o.NormalizedShares) {
		return true
	}

	return false
}

// SetNormalizedShares gets a reference to the given float32 and assigns it to the NormalizedShares field.
func (o *Dbv0038AssociationUsage) SetNormalizedShares(v float32) {
	o.NormalizedShares = &v
}

// GetEffectiveNormalizedUsage returns the EffectiveNormalizedUsage field value if set, zero value otherwise.
func (o *Dbv0038AssociationUsage) GetEffectiveNormalizedUsage() float32 {
	if o == nil || IsNil(o.EffectiveNormalizedUsage) {
		var ret float32
		return ret
	}
	return *o.EffectiveNormalizedUsage
}

// GetEffectiveNormalizedUsageOk returns a tuple with the EffectiveNormalizedUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038AssociationUsage) GetEffectiveNormalizedUsageOk() (*float32, bool) {
	if o == nil || IsNil(o.EffectiveNormalizedUsage) {
		return nil, false
	}
	return o.EffectiveNormalizedUsage, true
}

// HasEffectiveNormalizedUsage returns a boolean if a field has been set.
func (o *Dbv0038AssociationUsage) HasEffectiveNormalizedUsage() bool {
	if o != nil && !IsNil(o.EffectiveNormalizedUsage) {
		return true
	}

	return false
}

// SetEffectiveNormalizedUsage gets a reference to the given float32 and assigns it to the EffectiveNormalizedUsage field.
func (o *Dbv0038AssociationUsage) SetEffectiveNormalizedUsage(v float32) {
	o.EffectiveNormalizedUsage = &v
}

// GetRawUsage returns the RawUsage field value if set, zero value otherwise.
func (o *Dbv0038AssociationUsage) GetRawUsage() int32 {
	if o == nil || IsNil(o.RawUsage) {
		var ret int32
		return ret
	}
	return *o.RawUsage
}

// GetRawUsageOk returns a tuple with the RawUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038AssociationUsage) GetRawUsageOk() (*int32, bool) {
	if o == nil || IsNil(o.RawUsage) {
		return nil, false
	}
	return o.RawUsage, true
}

// HasRawUsage returns a boolean if a field has been set.
func (o *Dbv0038AssociationUsage) HasRawUsage() bool {
	if o != nil && !IsNil(o.RawUsage) {
		return true
	}

	return false
}

// SetRawUsage gets a reference to the given int32 and assigns it to the RawUsage field.
func (o *Dbv0038AssociationUsage) SetRawUsage(v int32) {
	o.RawUsage = &v
}

// GetJobCount returns the JobCount field value if set, zero value otherwise.
func (o *Dbv0038AssociationUsage) GetJobCount() int32 {
	if o == nil || IsNil(o.JobCount) {
		var ret int32
		return ret
	}
	return *o.JobCount
}

// GetJobCountOk returns a tuple with the JobCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038AssociationUsage) GetJobCountOk() (*int32, bool) {
	if o == nil || IsNil(o.JobCount) {
		return nil, false
	}
	return o.JobCount, true
}

// HasJobCount returns a boolean if a field has been set.
func (o *Dbv0038AssociationUsage) HasJobCount() bool {
	if o != nil && !IsNil(o.JobCount) {
		return true
	}

	return false
}

// SetJobCount gets a reference to the given int32 and assigns it to the JobCount field.
func (o *Dbv0038AssociationUsage) SetJobCount(v int32) {
	o.JobCount = &v
}

// GetFairshareLevel returns the FairshareLevel field value if set, zero value otherwise.
func (o *Dbv0038AssociationUsage) GetFairshareLevel() float32 {
	if o == nil || IsNil(o.FairshareLevel) {
		var ret float32
		return ret
	}
	return *o.FairshareLevel
}

// GetFairshareLevelOk returns a tuple with the FairshareLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038AssociationUsage) GetFairshareLevelOk() (*float32, bool) {
	if o == nil || IsNil(o.FairshareLevel) {
		return nil, false
	}
	return o.FairshareLevel, true
}

// HasFairshareLevel returns a boolean if a field has been set.
func (o *Dbv0038AssociationUsage) HasFairshareLevel() bool {
	if o != nil && !IsNil(o.FairshareLevel) {
		return true
	}

	return false
}

// SetFairshareLevel gets a reference to the given float32 and assigns it to the FairshareLevel field.
func (o *Dbv0038AssociationUsage) SetFairshareLevel(v float32) {
	o.FairshareLevel = &v
}

func (o Dbv0038AssociationUsage) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0038AssociationUsage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccrueJobCount) {
		toSerialize["accrue_job_count"] = o.AccrueJobCount
	}
	if !IsNil(o.GroupUsedWallclock) {
		toSerialize["group_used_wallclock"] = o.GroupUsedWallclock
	}
	if !IsNil(o.FairshareFactor) {
		toSerialize["fairshare_factor"] = o.FairshareFactor
	}
	if !IsNil(o.FairshareShares) {
		toSerialize["fairshare_shares"] = o.FairshareShares
	}
	if !IsNil(o.NormalizedPriority) {
		toSerialize["normalized_priority"] = o.NormalizedPriority
	}
	if !IsNil(o.NormalizedShares) {
		toSerialize["normalized_shares"] = o.NormalizedShares
	}
	if !IsNil(o.EffectiveNormalizedUsage) {
		toSerialize["effective_normalized_usage"] = o.EffectiveNormalizedUsage
	}
	if !IsNil(o.RawUsage) {
		toSerialize["raw_usage"] = o.RawUsage
	}
	if !IsNil(o.JobCount) {
		toSerialize["job_count"] = o.JobCount
	}
	if !IsNil(o.FairshareLevel) {
		toSerialize["fairshare_level"] = o.FairshareLevel
	}
	return toSerialize, nil
}

type NullableDbv0038AssociationUsage struct {
	value *Dbv0038AssociationUsage
	isSet bool
}

func (v NullableDbv0038AssociationUsage) Get() *Dbv0038AssociationUsage {
	return v.value
}

func (v *NullableDbv0038AssociationUsage) Set(val *Dbv0038AssociationUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0038AssociationUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0038AssociationUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0038AssociationUsage(val *Dbv0038AssociationUsage) *NullableDbv0038AssociationUsage {
	return &NullableDbv0038AssociationUsage{value: val, isSet: true}
}

func (v NullableDbv0038AssociationUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0038AssociationUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
