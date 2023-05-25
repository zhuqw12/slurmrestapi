# Dbv0037JobStepTres

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Requested** | Pointer to [**Dbv0037JobStepTresRequested**](Dbv0037JobStepTresRequested.md) |  | [optional] 
**Consumed** | Pointer to [**Dbv0037JobStepTresRequested**](Dbv0037JobStepTresRequested.md) |  | [optional] 
**Allocated** | Pointer to [**[]Dbv0038TresListInner**](Dbv0038TresListInner.md) | TRES list of attributes | [optional] 

## Methods

### NewDbv0037JobStepTres

`func NewDbv0037JobStepTres() *Dbv0037JobStepTres`

NewDbv0037JobStepTres instantiates a new Dbv0037JobStepTres object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbv0037JobStepTresWithDefaults

`func NewDbv0037JobStepTresWithDefaults() *Dbv0037JobStepTres`

NewDbv0037JobStepTresWithDefaults instantiates a new Dbv0037JobStepTres object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequested

`func (o *Dbv0037JobStepTres) GetRequested() Dbv0037JobStepTresRequested`

GetRequested returns the Requested field if non-nil, zero value otherwise.

### GetRequestedOk

`func (o *Dbv0037JobStepTres) GetRequestedOk() (*Dbv0037JobStepTresRequested, bool)`

GetRequestedOk returns a tuple with the Requested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequested

`func (o *Dbv0037JobStepTres) SetRequested(v Dbv0037JobStepTresRequested)`

SetRequested sets Requested field to given value.

### HasRequested

`func (o *Dbv0037JobStepTres) HasRequested() bool`

HasRequested returns a boolean if a field has been set.

### GetConsumed

`func (o *Dbv0037JobStepTres) GetConsumed() Dbv0037JobStepTresRequested`

GetConsumed returns the Consumed field if non-nil, zero value otherwise.

### GetConsumedOk

`func (o *Dbv0037JobStepTres) GetConsumedOk() (*Dbv0037JobStepTresRequested, bool)`

GetConsumedOk returns a tuple with the Consumed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumed

`func (o *Dbv0037JobStepTres) SetConsumed(v Dbv0037JobStepTresRequested)`

SetConsumed sets Consumed field to given value.

### HasConsumed

`func (o *Dbv0037JobStepTres) HasConsumed() bool`

HasConsumed returns a boolean if a field has been set.

### GetAllocated

`func (o *Dbv0037JobStepTres) GetAllocated() []Dbv0038TresListInner`

GetAllocated returns the Allocated field if non-nil, zero value otherwise.

### GetAllocatedOk

`func (o *Dbv0037JobStepTres) GetAllocatedOk() (*[]Dbv0038TresListInner, bool)`

GetAllocatedOk returns a tuple with the Allocated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocated

`func (o *Dbv0037JobStepTres) SetAllocated(v []Dbv0038TresListInner)`

SetAllocated sets Allocated field to given value.

### HasAllocated

`func (o *Dbv0037JobStepTres) HasAllocated() bool`

HasAllocated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


