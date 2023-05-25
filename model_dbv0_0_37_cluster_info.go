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

// checks if the Dbv0037ClusterInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0037ClusterInfo{}

// Dbv0037ClusterInfo struct for Dbv0037ClusterInfo
type Dbv0037ClusterInfo struct {
	Controller *Dbv0038ClusterInfoController `json:"controller,omitempty"`
	// List of properties of cluster
	Flags []string `json:"flags,omitempty"`
	// Cluster name
	Name *string `json:"name,omitempty"`
	// Assigned nodes
	Nodes *string `json:"nodes,omitempty"`
	// Configured select plugin
	SelectPlugin *string                         `json:"select_plugin,omitempty"`
	Associations *Dbv0037ClusterInfoAssociations `json:"associations,omitempty"`
	// Number rpc version
	RpcVersion *int32 `json:"rpc_version,omitempty"`
	// List of TRES in cluster
	Tres []Dbv0037ResponseTres `json:"tres,omitempty"`
}

// NewDbv0037ClusterInfo instantiates a new Dbv0037ClusterInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0037ClusterInfo() *Dbv0037ClusterInfo {
	this := Dbv0037ClusterInfo{}
	return &this
}

// NewDbv0037ClusterInfoWithDefaults instantiates a new Dbv0037ClusterInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0037ClusterInfoWithDefaults() *Dbv0037ClusterInfo {
	this := Dbv0037ClusterInfo{}
	return &this
}

// GetController returns the Controller field value if set, zero value otherwise.
func (o *Dbv0037ClusterInfo) GetController() Dbv0038ClusterInfoController {
	if o == nil || IsNil(o.Controller) {
		var ret Dbv0038ClusterInfoController
		return ret
	}
	return *o.Controller
}

// GetControllerOk returns a tuple with the Controller field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037ClusterInfo) GetControllerOk() (*Dbv0038ClusterInfoController, bool) {
	if o == nil || IsNil(o.Controller) {
		return nil, false
	}
	return o.Controller, true
}

// HasController returns a boolean if a field has been set.
func (o *Dbv0037ClusterInfo) HasController() bool {
	if o != nil && !IsNil(o.Controller) {
		return true
	}

	return false
}

// SetController gets a reference to the given Dbv0038ClusterInfoController and assigns it to the Controller field.
func (o *Dbv0037ClusterInfo) SetController(v Dbv0038ClusterInfoController) {
	o.Controller = &v
}

// GetFlags returns the Flags field value if set, zero value otherwise.
func (o *Dbv0037ClusterInfo) GetFlags() []string {
	if o == nil || IsNil(o.Flags) {
		var ret []string
		return ret
	}
	return o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037ClusterInfo) GetFlagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Flags) {
		return nil, false
	}
	return o.Flags, true
}

// HasFlags returns a boolean if a field has been set.
func (o *Dbv0037ClusterInfo) HasFlags() bool {
	if o != nil && !IsNil(o.Flags) {
		return true
	}

	return false
}

// SetFlags gets a reference to the given []string and assigns it to the Flags field.
func (o *Dbv0037ClusterInfo) SetFlags(v []string) {
	o.Flags = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Dbv0037ClusterInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037ClusterInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Dbv0037ClusterInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Dbv0037ClusterInfo) SetName(v string) {
	o.Name = &v
}

// GetNodes returns the Nodes field value if set, zero value otherwise.
func (o *Dbv0037ClusterInfo) GetNodes() string {
	if o == nil || IsNil(o.Nodes) {
		var ret string
		return ret
	}
	return *o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037ClusterInfo) GetNodesOk() (*string, bool) {
	if o == nil || IsNil(o.Nodes) {
		return nil, false
	}
	return o.Nodes, true
}

// HasNodes returns a boolean if a field has been set.
func (o *Dbv0037ClusterInfo) HasNodes() bool {
	if o != nil && !IsNil(o.Nodes) {
		return true
	}

	return false
}

// SetNodes gets a reference to the given string and assigns it to the Nodes field.
func (o *Dbv0037ClusterInfo) SetNodes(v string) {
	o.Nodes = &v
}

// GetSelectPlugin returns the SelectPlugin field value if set, zero value otherwise.
func (o *Dbv0037ClusterInfo) GetSelectPlugin() string {
	if o == nil || IsNil(o.SelectPlugin) {
		var ret string
		return ret
	}
	return *o.SelectPlugin
}

// GetSelectPluginOk returns a tuple with the SelectPlugin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037ClusterInfo) GetSelectPluginOk() (*string, bool) {
	if o == nil || IsNil(o.SelectPlugin) {
		return nil, false
	}
	return o.SelectPlugin, true
}

// HasSelectPlugin returns a boolean if a field has been set.
func (o *Dbv0037ClusterInfo) HasSelectPlugin() bool {
	if o != nil && !IsNil(o.SelectPlugin) {
		return true
	}

	return false
}

// SetSelectPlugin gets a reference to the given string and assigns it to the SelectPlugin field.
func (o *Dbv0037ClusterInfo) SetSelectPlugin(v string) {
	o.SelectPlugin = &v
}

// GetAssociations returns the Associations field value if set, zero value otherwise.
func (o *Dbv0037ClusterInfo) GetAssociations() Dbv0037ClusterInfoAssociations {
	if o == nil || IsNil(o.Associations) {
		var ret Dbv0037ClusterInfoAssociations
		return ret
	}
	return *o.Associations
}

// GetAssociationsOk returns a tuple with the Associations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037ClusterInfo) GetAssociationsOk() (*Dbv0037ClusterInfoAssociations, bool) {
	if o == nil || IsNil(o.Associations) {
		return nil, false
	}
	return o.Associations, true
}

// HasAssociations returns a boolean if a field has been set.
func (o *Dbv0037ClusterInfo) HasAssociations() bool {
	if o != nil && !IsNil(o.Associations) {
		return true
	}

	return false
}

// SetAssociations gets a reference to the given Dbv0037ClusterInfoAssociations and assigns it to the Associations field.
func (o *Dbv0037ClusterInfo) SetAssociations(v Dbv0037ClusterInfoAssociations) {
	o.Associations = &v
}

// GetRpcVersion returns the RpcVersion field value if set, zero value otherwise.
func (o *Dbv0037ClusterInfo) GetRpcVersion() int32 {
	if o == nil || IsNil(o.RpcVersion) {
		var ret int32
		return ret
	}
	return *o.RpcVersion
}

// GetRpcVersionOk returns a tuple with the RpcVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037ClusterInfo) GetRpcVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.RpcVersion) {
		return nil, false
	}
	return o.RpcVersion, true
}

// HasRpcVersion returns a boolean if a field has been set.
func (o *Dbv0037ClusterInfo) HasRpcVersion() bool {
	if o != nil && !IsNil(o.RpcVersion) {
		return true
	}

	return false
}

// SetRpcVersion gets a reference to the given int32 and assigns it to the RpcVersion field.
func (o *Dbv0037ClusterInfo) SetRpcVersion(v int32) {
	o.RpcVersion = &v
}

// GetTres returns the Tres field value if set, zero value otherwise.
func (o *Dbv0037ClusterInfo) GetTres() []Dbv0037ResponseTres {
	if o == nil || IsNil(o.Tres) {
		var ret []Dbv0037ResponseTres
		return ret
	}
	return o.Tres
}

// GetTresOk returns a tuple with the Tres field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037ClusterInfo) GetTresOk() ([]Dbv0037ResponseTres, bool) {
	if o == nil || IsNil(o.Tres) {
		return nil, false
	}
	return o.Tres, true
}

// HasTres returns a boolean if a field has been set.
func (o *Dbv0037ClusterInfo) HasTres() bool {
	if o != nil && !IsNil(o.Tres) {
		return true
	}

	return false
}

// SetTres gets a reference to the given []Dbv0037ResponseTres and assigns it to the Tres field.
func (o *Dbv0037ClusterInfo) SetTres(v []Dbv0037ResponseTres) {
	o.Tres = v
}

func (o Dbv0037ClusterInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0037ClusterInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Controller) {
		toSerialize["controller"] = o.Controller
	}
	if !IsNil(o.Flags) {
		toSerialize["flags"] = o.Flags
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Nodes) {
		toSerialize["nodes"] = o.Nodes
	}
	if !IsNil(o.SelectPlugin) {
		toSerialize["select_plugin"] = o.SelectPlugin
	}
	if !IsNil(o.Associations) {
		toSerialize["associations"] = o.Associations
	}
	if !IsNil(o.RpcVersion) {
		toSerialize["rpc_version"] = o.RpcVersion
	}
	if !IsNil(o.Tres) {
		toSerialize["tres"] = o.Tres
	}
	return toSerialize, nil
}

type NullableDbv0037ClusterInfo struct {
	value *Dbv0037ClusterInfo
	isSet bool
}

func (v NullableDbv0037ClusterInfo) Get() *Dbv0037ClusterInfo {
	return v.value
}

func (v *NullableDbv0037ClusterInfo) Set(val *Dbv0037ClusterInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0037ClusterInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0037ClusterInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0037ClusterInfo(val *Dbv0037ClusterInfo) *NullableDbv0037ClusterInfo {
	return &NullableDbv0037ClusterInfo{value: val, isSet: true}
}

func (v NullableDbv0037ClusterInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0037ClusterInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
