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

// checks if the Dbv0038Cluster type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0038Cluster{}

// Dbv0038Cluster struct for Dbv0038Cluster
type Dbv0038Cluster struct {
	Controller *Dbv0038ClusterController `json:"controller,omitempty"`
	// List of properties of cluster
	Flags []string `json:"flags,omitempty"`
	// Cluster name
	Name *string `json:"name,omitempty"`
	// Assigned nodes
	Nodes *string `json:"nodes,omitempty"`
	// Configured select plugin
	SelectPlugin *string                     `json:"select_plugin,omitempty"`
	Associations *Dbv0038ClusterAssociations `json:"associations,omitempty"`
	// Number rpc version
	RpcVersion *int32 `json:"rpc_version,omitempty"`
	// TRES list of attributes
	Tres []Dbv0038TresListInner `json:"tres,omitempty"`
}

// NewDbv0038Cluster instantiates a new Dbv0038Cluster object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0038Cluster() *Dbv0038Cluster {
	this := Dbv0038Cluster{}
	return &this
}

// NewDbv0038ClusterWithDefaults instantiates a new Dbv0038Cluster object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0038ClusterWithDefaults() *Dbv0038Cluster {
	this := Dbv0038Cluster{}
	return &this
}

// GetController returns the Controller field value if set, zero value otherwise.
func (o *Dbv0038Cluster) GetController() Dbv0038ClusterController {
	if o == nil || IsNil(o.Controller) {
		var ret Dbv0038ClusterController
		return ret
	}
	return *o.Controller
}

// GetControllerOk returns a tuple with the Controller field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038Cluster) GetControllerOk() (*Dbv0038ClusterController, bool) {
	if o == nil || IsNil(o.Controller) {
		return nil, false
	}
	return o.Controller, true
}

// HasController returns a boolean if a field has been set.
func (o *Dbv0038Cluster) HasController() bool {
	if o != nil && !IsNil(o.Controller) {
		return true
	}

	return false
}

// SetController gets a reference to the given Dbv0038ClusterController and assigns it to the Controller field.
func (o *Dbv0038Cluster) SetController(v Dbv0038ClusterController) {
	o.Controller = &v
}

// GetFlags returns the Flags field value if set, zero value otherwise.
func (o *Dbv0038Cluster) GetFlags() []string {
	if o == nil || IsNil(o.Flags) {
		var ret []string
		return ret
	}
	return o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038Cluster) GetFlagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Flags) {
		return nil, false
	}
	return o.Flags, true
}

// HasFlags returns a boolean if a field has been set.
func (o *Dbv0038Cluster) HasFlags() bool {
	if o != nil && !IsNil(o.Flags) {
		return true
	}

	return false
}

// SetFlags gets a reference to the given []string and assigns it to the Flags field.
func (o *Dbv0038Cluster) SetFlags(v []string) {
	o.Flags = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Dbv0038Cluster) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038Cluster) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Dbv0038Cluster) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Dbv0038Cluster) SetName(v string) {
	o.Name = &v
}

// GetNodes returns the Nodes field value if set, zero value otherwise.
func (o *Dbv0038Cluster) GetNodes() string {
	if o == nil || IsNil(o.Nodes) {
		var ret string
		return ret
	}
	return *o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038Cluster) GetNodesOk() (*string, bool) {
	if o == nil || IsNil(o.Nodes) {
		return nil, false
	}
	return o.Nodes, true
}

// HasNodes returns a boolean if a field has been set.
func (o *Dbv0038Cluster) HasNodes() bool {
	if o != nil && !IsNil(o.Nodes) {
		return true
	}

	return false
}

// SetNodes gets a reference to the given string and assigns it to the Nodes field.
func (o *Dbv0038Cluster) SetNodes(v string) {
	o.Nodes = &v
}

// GetSelectPlugin returns the SelectPlugin field value if set, zero value otherwise.
func (o *Dbv0038Cluster) GetSelectPlugin() string {
	if o == nil || IsNil(o.SelectPlugin) {
		var ret string
		return ret
	}
	return *o.SelectPlugin
}

// GetSelectPluginOk returns a tuple with the SelectPlugin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038Cluster) GetSelectPluginOk() (*string, bool) {
	if o == nil || IsNil(o.SelectPlugin) {
		return nil, false
	}
	return o.SelectPlugin, true
}

// HasSelectPlugin returns a boolean if a field has been set.
func (o *Dbv0038Cluster) HasSelectPlugin() bool {
	if o != nil && !IsNil(o.SelectPlugin) {
		return true
	}

	return false
}

// SetSelectPlugin gets a reference to the given string and assigns it to the SelectPlugin field.
func (o *Dbv0038Cluster) SetSelectPlugin(v string) {
	o.SelectPlugin = &v
}

// GetAssociations returns the Associations field value if set, zero value otherwise.
func (o *Dbv0038Cluster) GetAssociations() Dbv0038ClusterAssociations {
	if o == nil || IsNil(o.Associations) {
		var ret Dbv0038ClusterAssociations
		return ret
	}
	return *o.Associations
}

// GetAssociationsOk returns a tuple with the Associations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038Cluster) GetAssociationsOk() (*Dbv0038ClusterAssociations, bool) {
	if o == nil || IsNil(o.Associations) {
		return nil, false
	}
	return o.Associations, true
}

// HasAssociations returns a boolean if a field has been set.
func (o *Dbv0038Cluster) HasAssociations() bool {
	if o != nil && !IsNil(o.Associations) {
		return true
	}

	return false
}

// SetAssociations gets a reference to the given Dbv0038ClusterAssociations and assigns it to the Associations field.
func (o *Dbv0038Cluster) SetAssociations(v Dbv0038ClusterAssociations) {
	o.Associations = &v
}

// GetRpcVersion returns the RpcVersion field value if set, zero value otherwise.
func (o *Dbv0038Cluster) GetRpcVersion() int32 {
	if o == nil || IsNil(o.RpcVersion) {
		var ret int32
		return ret
	}
	return *o.RpcVersion
}

// GetRpcVersionOk returns a tuple with the RpcVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038Cluster) GetRpcVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.RpcVersion) {
		return nil, false
	}
	return o.RpcVersion, true
}

// HasRpcVersion returns a boolean if a field has been set.
func (o *Dbv0038Cluster) HasRpcVersion() bool {
	if o != nil && !IsNil(o.RpcVersion) {
		return true
	}

	return false
}

// SetRpcVersion gets a reference to the given int32 and assigns it to the RpcVersion field.
func (o *Dbv0038Cluster) SetRpcVersion(v int32) {
	o.RpcVersion = &v
}

// GetTres returns the Tres field value if set, zero value otherwise.
func (o *Dbv0038Cluster) GetTres() []Dbv0038TresListInner {
	if o == nil || IsNil(o.Tres) {
		var ret []Dbv0038TresListInner
		return ret
	}
	return o.Tres
}

// GetTresOk returns a tuple with the Tres field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038Cluster) GetTresOk() ([]Dbv0038TresListInner, bool) {
	if o == nil || IsNil(o.Tres) {
		return nil, false
	}
	return o.Tres, true
}

// HasTres returns a boolean if a field has been set.
func (o *Dbv0038Cluster) HasTres() bool {
	if o != nil && !IsNil(o.Tres) {
		return true
	}

	return false
}

// SetTres gets a reference to the given []Dbv0038TresListInner and assigns it to the Tres field.
func (o *Dbv0038Cluster) SetTres(v []Dbv0038TresListInner) {
	o.Tres = v
}

func (o Dbv0038Cluster) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0038Cluster) ToMap() (map[string]interface{}, error) {
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

type NullableDbv0038Cluster struct {
	value *Dbv0038Cluster
	isSet bool
}

func (v NullableDbv0038Cluster) Get() *Dbv0038Cluster {
	return v.value
}

func (v *NullableDbv0038Cluster) Set(val *Dbv0038Cluster) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0038Cluster) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0038Cluster) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0038Cluster(val *Dbv0038Cluster) *NullableDbv0038Cluster {
	return &NullableDbv0038Cluster{value: val, isSet: true}
}

func (v NullableDbv0038Cluster) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0038Cluster) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}