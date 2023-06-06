# Dbv0038ClusterInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**Dbv0038Meta**](Dbv0038Meta.md) |  | [optional]
**Errors** | Pointer to [**[]Dbv0038Error**](Dbv0038Error.md) | Slurm errors | [optional]
**Clusters** | Pointer to [**[]Dbv0038Cluster**](Dbv0038Cluster.md) | Array of jobs | [optional]

## Methods

### NewDbv0038ClusterInfo

`func NewDbv0038ClusterInfo() *Dbv0038ClusterInfo`

NewDbv0038ClusterInfo instantiates a new Dbv0038ClusterInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbv0038ClusterInfoWithDefaults

`func NewDbv0038ClusterInfoWithDefaults() *Dbv0038ClusterInfo`

NewDbv0038ClusterInfoWithDefaults instantiates a new Dbv0038ClusterInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *Dbv0038ClusterInfo) GetMeta() Dbv0038Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *Dbv0038ClusterInfo) GetMetaOk() (*Dbv0038Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise and a boolean to check if the value
has been set.

### SetMeta

`func (o *Dbv0038ClusterInfo) SetMeta(v Dbv0038Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *Dbv0038ClusterInfo) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *Dbv0038ClusterInfo) GetErrors() []Dbv0038Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *Dbv0038ClusterInfo) GetErrorsOk() (*[]Dbv0038Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise and a boolean to check if the
value has been set.

### SetErrors

`func (o *Dbv0038ClusterInfo) SetErrors(v []Dbv0038Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *Dbv0038ClusterInfo) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetClusters

`func (o *Dbv0038ClusterInfo) GetClusters() []Dbv0038Cluster`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *Dbv0038ClusterInfo) GetClustersOk() (*[]Dbv0038Cluster, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise and a boolean to check if
the value has been set.

### SetClusters

`func (o *Dbv0038ClusterInfo) SetClusters(v []Dbv0038Cluster)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *Dbv0038ClusterInfo) HasClusters() bool`

HasClusters returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


