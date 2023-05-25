# Dbv0037AssociationMaxJobs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **int32** | Max TRES for active total jobs | [optional] 
**Accruing** | Pointer to **int32** | Max TRES for job accruing priority | [optional] 
**Total** | Pointer to **int32** | Max TRES for job total submitted | [optional] 
**Per** | Pointer to [**Dbv0038AssociationMaxJobsPer**](Dbv0038AssociationMaxJobsPer.md) |  | [optional] 

## Methods

### NewDbv0037AssociationMaxJobs

`func NewDbv0037AssociationMaxJobs() *Dbv0037AssociationMaxJobs`

NewDbv0037AssociationMaxJobs instantiates a new Dbv0037AssociationMaxJobs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbv0037AssociationMaxJobsWithDefaults

`func NewDbv0037AssociationMaxJobsWithDefaults() *Dbv0037AssociationMaxJobs`

NewDbv0037AssociationMaxJobsWithDefaults instantiates a new Dbv0037AssociationMaxJobs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *Dbv0037AssociationMaxJobs) GetActive() int32`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Dbv0037AssociationMaxJobs) GetActiveOk() (*int32, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Dbv0037AssociationMaxJobs) SetActive(v int32)`

SetActive sets Active field to given value.

### HasActive

`func (o *Dbv0037AssociationMaxJobs) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAccruing

`func (o *Dbv0037AssociationMaxJobs) GetAccruing() int32`

GetAccruing returns the Accruing field if non-nil, zero value otherwise.

### GetAccruingOk

`func (o *Dbv0037AssociationMaxJobs) GetAccruingOk() (*int32, bool)`

GetAccruingOk returns a tuple with the Accruing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccruing

`func (o *Dbv0037AssociationMaxJobs) SetAccruing(v int32)`

SetAccruing sets Accruing field to given value.

### HasAccruing

`func (o *Dbv0037AssociationMaxJobs) HasAccruing() bool`

HasAccruing returns a boolean if a field has been set.

### GetTotal

`func (o *Dbv0037AssociationMaxJobs) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *Dbv0037AssociationMaxJobs) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *Dbv0037AssociationMaxJobs) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *Dbv0037AssociationMaxJobs) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetPer

`func (o *Dbv0037AssociationMaxJobs) GetPer() Dbv0038AssociationMaxJobsPer`

GetPer returns the Per field if non-nil, zero value otherwise.

### GetPerOk

`func (o *Dbv0037AssociationMaxJobs) GetPerOk() (*Dbv0038AssociationMaxJobsPer, bool)`

GetPerOk returns a tuple with the Per field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPer

`func (o *Dbv0037AssociationMaxJobs) SetPer(v Dbv0038AssociationMaxJobsPer)`

SetPer sets Per field to given value.

### HasPer

`func (o *Dbv0037AssociationMaxJobs) HasPer() bool`

HasPer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


