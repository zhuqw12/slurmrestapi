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

// checks if the Dbv0037Account type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0037Account{}

// Dbv0037Account Account description
type Dbv0037Account struct {
	// List of assigned associations
	Associations []Dbv0037AssociationShortInfo `json:"associations,omitempty"`
	// List of assigned coordinators
	Coordinators []Dbv0037CoordinatorInfo `json:"coordinators,omitempty"`
	// Description of account
	Description *string `json:"description,omitempty"`
	// Name of account
	Name *string `json:"name,omitempty"`
	// Assigned organization of account
	Organization *string `json:"organization,omitempty"`
	// List of properties of account
	Flags []string `json:"flags,omitempty"`
}

// NewDbv0037Account instantiates a new Dbv0037Account object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0037Account() *Dbv0037Account {
	this := Dbv0037Account{}
	return &this
}

// NewDbv0037AccountWithDefaults instantiates a new Dbv0037Account object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0037AccountWithDefaults() *Dbv0037Account {
	this := Dbv0037Account{}
	return &this
}

// GetAssociations returns the Associations field value if set, zero value otherwise.
func (o *Dbv0037Account) GetAssociations() []Dbv0037AssociationShortInfo {
	if o == nil || IsNil(o.Associations) {
		var ret []Dbv0037AssociationShortInfo
		return ret
	}
	return o.Associations
}

// GetAssociationsOk returns a tuple with the Associations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037Account) GetAssociationsOk() ([]Dbv0037AssociationShortInfo, bool) {
	if o == nil || IsNil(o.Associations) {
		return nil, false
	}
	return o.Associations, true
}

// HasAssociations returns a boolean if a field has been set.
func (o *Dbv0037Account) HasAssociations() bool {
	if o != nil && !IsNil(o.Associations) {
		return true
	}

	return false
}

// SetAssociations gets a reference to the given []Dbv0037AssociationShortInfo and assigns it to the Associations field.
func (o *Dbv0037Account) SetAssociations(v []Dbv0037AssociationShortInfo) {
	o.Associations = v
}

// GetCoordinators returns the Coordinators field value if set, zero value otherwise.
func (o *Dbv0037Account) GetCoordinators() []Dbv0037CoordinatorInfo {
	if o == nil || IsNil(o.Coordinators) {
		var ret []Dbv0037CoordinatorInfo
		return ret
	}
	return o.Coordinators
}

// GetCoordinatorsOk returns a tuple with the Coordinators field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037Account) GetCoordinatorsOk() ([]Dbv0037CoordinatorInfo, bool) {
	if o == nil || IsNil(o.Coordinators) {
		return nil, false
	}
	return o.Coordinators, true
}

// HasCoordinators returns a boolean if a field has been set.
func (o *Dbv0037Account) HasCoordinators() bool {
	if o != nil && !IsNil(o.Coordinators) {
		return true
	}

	return false
}

// SetCoordinators gets a reference to the given []Dbv0037CoordinatorInfo and assigns it to the Coordinators field.
func (o *Dbv0037Account) SetCoordinators(v []Dbv0037CoordinatorInfo) {
	o.Coordinators = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Dbv0037Account) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037Account) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Dbv0037Account) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Dbv0037Account) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Dbv0037Account) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037Account) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Dbv0037Account) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Dbv0037Account) SetName(v string) {
	o.Name = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *Dbv0037Account) GetOrganization() string {
	if o == nil || IsNil(o.Organization) {
		var ret string
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037Account) GetOrganizationOk() (*string, bool) {
	if o == nil || IsNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *Dbv0037Account) HasOrganization() bool {
	if o != nil && !IsNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given string and assigns it to the Organization field.
func (o *Dbv0037Account) SetOrganization(v string) {
	o.Organization = &v
}

// GetFlags returns the Flags field value if set, zero value otherwise.
func (o *Dbv0037Account) GetFlags() []string {
	if o == nil || IsNil(o.Flags) {
		var ret []string
		return ret
	}
	return o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037Account) GetFlagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Flags) {
		return nil, false
	}
	return o.Flags, true
}

// HasFlags returns a boolean if a field has been set.
func (o *Dbv0037Account) HasFlags() bool {
	if o != nil && !IsNil(o.Flags) {
		return true
	}

	return false
}

// SetFlags gets a reference to the given []string and assigns it to the Flags field.
func (o *Dbv0037Account) SetFlags(v []string) {
	o.Flags = v
}

func (o Dbv0037Account) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0037Account) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Associations) {
		toSerialize["associations"] = o.Associations
	}
	if !IsNil(o.Coordinators) {
		toSerialize["coordinators"] = o.Coordinators
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Organization) {
		toSerialize["organization"] = o.Organization
	}
	if !IsNil(o.Flags) {
		toSerialize["flags"] = o.Flags
	}
	return toSerialize, nil
}

type NullableDbv0037Account struct {
	value *Dbv0037Account
	isSet bool
}

func (v NullableDbv0037Account) Get() *Dbv0037Account {
	return v.value
}

func (v *NullableDbv0037Account) Set(val *Dbv0037Account) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0037Account) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0037Account) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0037Account(val *Dbv0037Account) *NullableDbv0037Account {
	return &NullableDbv0037Account{value: val, isSet: true}
}

func (v NullableDbv0037Account) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0037Account) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
