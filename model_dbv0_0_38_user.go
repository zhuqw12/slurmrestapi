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

// checks if the Dbv0038User type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0038User{}

// Dbv0038User User description
type Dbv0038User struct {
	// Description of administrator level
	AdministratorLevel *string `json:"administrator_level,omitempty"`
	// Assigned associations
	Associations []Dbv0038AssociationShortInfo `json:"associations,omitempty"`
	// List of assigned coordinators
	Coordinators []Dbv0038CoordinatorInfo `json:"coordinators,omitempty"`
	Default      *Dbv0038UserDefault      `json:"default,omitempty"`
	// List of properties of user
	Flags []string `json:"flags,omitempty"`
	// User name
	Name *string `json:"name,omitempty"`
}

// NewDbv0038User instantiates a new Dbv0038User object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0038User() *Dbv0038User {
	this := Dbv0038User{}
	return &this
}

// NewDbv0038UserWithDefaults instantiates a new Dbv0038User object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0038UserWithDefaults() *Dbv0038User {
	this := Dbv0038User{}
	return &this
}

// GetAdministratorLevel returns the AdministratorLevel field value if set, zero value otherwise.
func (o *Dbv0038User) GetAdministratorLevel() string {
	if o == nil || IsNil(o.AdministratorLevel) {
		var ret string
		return ret
	}
	return *o.AdministratorLevel
}

// GetAdministratorLevelOk returns a tuple with the AdministratorLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038User) GetAdministratorLevelOk() (*string, bool) {
	if o == nil || IsNil(o.AdministratorLevel) {
		return nil, false
	}
	return o.AdministratorLevel, true
}

// HasAdministratorLevel returns a boolean if a field has been set.
func (o *Dbv0038User) HasAdministratorLevel() bool {
	if o != nil && !IsNil(o.AdministratorLevel) {
		return true
	}

	return false
}

// SetAdministratorLevel gets a reference to the given string and assigns it to the AdministratorLevel field.
func (o *Dbv0038User) SetAdministratorLevel(v string) {
	o.AdministratorLevel = &v
}

// GetAssociations returns the Associations field value if set, zero value otherwise.
func (o *Dbv0038User) GetAssociations() []Dbv0038AssociationShortInfo {
	if o == nil || IsNil(o.Associations) {
		var ret []Dbv0038AssociationShortInfo
		return ret
	}
	return o.Associations
}

// GetAssociationsOk returns a tuple with the Associations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038User) GetAssociationsOk() ([]Dbv0038AssociationShortInfo, bool) {
	if o == nil || IsNil(o.Associations) {
		return nil, false
	}
	return o.Associations, true
}

// HasAssociations returns a boolean if a field has been set.
func (o *Dbv0038User) HasAssociations() bool {
	if o != nil && !IsNil(o.Associations) {
		return true
	}

	return false
}

// SetAssociations gets a reference to the given []Dbv0038AssociationShortInfo and assigns it to the Associations field.
func (o *Dbv0038User) SetAssociations(v []Dbv0038AssociationShortInfo) {
	o.Associations = v
}

// GetCoordinators returns the Coordinators field value if set, zero value otherwise.
func (o *Dbv0038User) GetCoordinators() []Dbv0038CoordinatorInfo {
	if o == nil || IsNil(o.Coordinators) {
		var ret []Dbv0038CoordinatorInfo
		return ret
	}
	return o.Coordinators
}

// GetCoordinatorsOk returns a tuple with the Coordinators field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038User) GetCoordinatorsOk() ([]Dbv0038CoordinatorInfo, bool) {
	if o == nil || IsNil(o.Coordinators) {
		return nil, false
	}
	return o.Coordinators, true
}

// HasCoordinators returns a boolean if a field has been set.
func (o *Dbv0038User) HasCoordinators() bool {
	if o != nil && !IsNil(o.Coordinators) {
		return true
	}

	return false
}

// SetCoordinators gets a reference to the given []Dbv0038CoordinatorInfo and assigns it to the Coordinators field.
func (o *Dbv0038User) SetCoordinators(v []Dbv0038CoordinatorInfo) {
	o.Coordinators = v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *Dbv0038User) GetDefault() Dbv0038UserDefault {
	if o == nil || IsNil(o.Default) {
		var ret Dbv0038UserDefault
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038User) GetDefaultOk() (*Dbv0038UserDefault, bool) {
	if o == nil || IsNil(o.Default) {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *Dbv0038User) HasDefault() bool {
	if o != nil && !IsNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given Dbv0038UserDefault and assigns it to the Default field.
func (o *Dbv0038User) SetDefault(v Dbv0038UserDefault) {
	o.Default = &v
}

// GetFlags returns the Flags field value if set, zero value otherwise.
func (o *Dbv0038User) GetFlags() []string {
	if o == nil || IsNil(o.Flags) {
		var ret []string
		return ret
	}
	return o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038User) GetFlagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Flags) {
		return nil, false
	}
	return o.Flags, true
}

// HasFlags returns a boolean if a field has been set.
func (o *Dbv0038User) HasFlags() bool {
	if o != nil && !IsNil(o.Flags) {
		return true
	}

	return false
}

// SetFlags gets a reference to the given []string and assigns it to the Flags field.
func (o *Dbv0038User) SetFlags(v []string) {
	o.Flags = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Dbv0038User) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038User) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Dbv0038User) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Dbv0038User) SetName(v string) {
	o.Name = &v
}

func (o Dbv0038User) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0038User) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdministratorLevel) {
		toSerialize["administrator_level"] = o.AdministratorLevel
	}
	if !IsNil(o.Associations) {
		toSerialize["associations"] = o.Associations
	}
	if !IsNil(o.Coordinators) {
		toSerialize["coordinators"] = o.Coordinators
	}
	if !IsNil(o.Default) {
		toSerialize["default"] = o.Default
	}
	if !IsNil(o.Flags) {
		toSerialize["flags"] = o.Flags
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableDbv0038User struct {
	value *Dbv0038User
	isSet bool
}

func (v NullableDbv0038User) Get() *Dbv0038User {
	return v.value
}

func (v *NullableDbv0038User) Set(val *Dbv0038User) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0038User) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0038User) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0038User(val *Dbv0038User) *NullableDbv0038User {
	return &NullableDbv0038User{value: val, isSet: true}
}

func (v NullableDbv0038User) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0038User) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
