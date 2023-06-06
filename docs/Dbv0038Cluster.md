# Dbv0038Cluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Controller** | Pointer to [**Dbv0038ClusterController**](Dbv0038ClusterController.md) |  | [optional]
**Flags** | Pointer to **[]string** | List of properties of cluster | [optional]
**Name** | Pointer to **string** | Cluster name | [optional]
**Nodes** | Pointer to **string** | Assigned nodes | [optional]
**SelectPlugin** | Pointer to **string** | Configured select plugin | [optional]
**Associations** | Pointer to [**Dbv0038ClusterAssociations**](Dbv0038ClusterAssociations.md) |  | [optional]
**RpcVersion** | Pointer to **int32** | Number rpc version | [optional]
**Tres** | Pointer to [**[]Dbv0038TresListInner**](Dbv0038TresListInner.md) | TRES list of attributes | [optional]

## Methods

### NewDbv0038Cluster

`func NewDbv0038Cluster() *Dbv0038Cluster`

NewDbv0038Cluster instantiates a new Dbv0038Cluster object This constructor will assign default values to properties
that have it defined, and makes sure properties required by API are set, but the set of arguments will change when the
set of required properties is changed

### NewDbv0038ClusterWithDefaults

`func NewDbv0038ClusterWithDefaults() *Dbv0038Cluster`

NewDbv0038ClusterWithDefaults instantiates a new Dbv0038Cluster object This constructor will only assign default values
to properties that have it defined, but it doesn't guarantee that properties required by API are set

### GetController

`func (o *Dbv0038Cluster) GetController() Dbv0038ClusterController`

GetController returns the Controller field if non-nil, zero value otherwise.

### GetControllerOk

`func (o *Dbv0038Cluster) GetControllerOk() (*Dbv0038ClusterController, bool)`

GetControllerOk returns a tuple with the Controller field if it's non-nil, zero value otherwise and a boolean to check
if the value has been set.

### SetController

`func (o *Dbv0038Cluster) SetController(v Dbv0038ClusterController)`

SetController sets Controller field to given value.

### HasController

`func (o *Dbv0038Cluster) HasController() bool`

HasController returns a boolean if a field has been set.

### GetFlags

`func (o *Dbv0038Cluster) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *Dbv0038Cluster) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise and a boolean to check if the
value has been set.

### SetFlags

`func (o *Dbv0038Cluster) SetFlags(v []string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *Dbv0038Cluster) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetName

`func (o *Dbv0038Cluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Dbv0038Cluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise and a boolean to check if the value
has been set.

### SetName

`func (o *Dbv0038Cluster) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Dbv0038Cluster) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodes

`func (o *Dbv0038Cluster) GetNodes() string`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *Dbv0038Cluster) GetNodesOk() (*string, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise and a boolean to check if the
value has been set.

### SetNodes

`func (o *Dbv0038Cluster) SetNodes(v string)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *Dbv0038Cluster) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetSelectPlugin

`func (o *Dbv0038Cluster) GetSelectPlugin() string`

GetSelectPlugin returns the SelectPlugin field if non-nil, zero value otherwise.

### GetSelectPluginOk

`func (o *Dbv0038Cluster) GetSelectPluginOk() (*string, bool)`

GetSelectPluginOk returns a tuple with the SelectPlugin field if it's non-nil, zero value otherwise and a boolean to
check if the value has been set.

### SetSelectPlugin

`func (o *Dbv0038Cluster) SetSelectPlugin(v string)`

SetSelectPlugin sets SelectPlugin field to given value.

### HasSelectPlugin

`func (o *Dbv0038Cluster) HasSelectPlugin() bool`

HasSelectPlugin returns a boolean if a field has been set.

### GetAssociations

`func (o *Dbv0038Cluster) GetAssociations() Dbv0038ClusterAssociations`

GetAssociations returns the Associations field if non-nil, zero value otherwise.

### GetAssociationsOk

`func (o *Dbv0038Cluster) GetAssociationsOk() (*Dbv0038ClusterAssociations, bool)`

GetAssociationsOk returns a tuple with the Associations field if it's non-nil, zero value otherwise and a boolean to
check if the value has been set.

### SetAssociations

`func (o *Dbv0038Cluster) SetAssociations(v Dbv0038ClusterAssociations)`

SetAssociations sets Associations field to given value.

### HasAssociations

`func (o *Dbv0038Cluster) HasAssociations() bool`

HasAssociations returns a boolean if a field has been set.

### GetRpcVersion

`func (o *Dbv0038Cluster) GetRpcVersion() int32`

GetRpcVersion returns the RpcVersion field if non-nil, zero value otherwise.

### GetRpcVersionOk

`func (o *Dbv0038Cluster) GetRpcVersionOk() (*int32, bool)`

GetRpcVersionOk returns a tuple with the RpcVersion field if it's non-nil, zero value otherwise and a boolean to check
if the value has been set.

### SetRpcVersion

`func (o *Dbv0038Cluster) SetRpcVersion(v int32)`

SetRpcVersion sets RpcVersion field to given value.

### HasRpcVersion

`func (o *Dbv0038Cluster) HasRpcVersion() bool`

HasRpcVersion returns a boolean if a field has been set.

### GetTres

`func (o *Dbv0038Cluster) GetTres() []Dbv0038TresListInner`

GetTres returns the Tres field if non-nil, zero value otherwise.

### GetTresOk

`func (o *Dbv0038Cluster) GetTresOk() (*[]Dbv0038TresListInner, bool)`

GetTresOk returns a tuple with the Tres field if it's non-nil, zero value otherwise and a boolean to check if the value
has been set.

### SetTres

`func (o *Dbv0038Cluster) SetTres(v []Dbv0038TresListInner)`

SetTres sets Tres field to given value.

### HasTres

`func (o *Dbv0038Cluster) HasTres() bool`

HasTres returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


