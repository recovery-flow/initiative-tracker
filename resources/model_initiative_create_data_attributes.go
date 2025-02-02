/*
test

example

API version: 0.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package resources

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the InitiativeCreateDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InitiativeCreateDataAttributes{}

// InitiativeCreateDataAttributes struct for InitiativeCreateDataAttributes
type InitiativeCreateDataAttributes struct {
	// name of initiative
	Name string `json:"name"`
	// description of initiative
	Desc string `json:"desc"`
	// goal of initiative
	Goal string `json:"goal"`
	// location of initiative
	Location *string `json:"location,omitempty"`
	// types of initiative
	Type string `json:"type"`
	// status of initiative
	Status string `json:"status"`
	// final cost of initiative
	FinalCost int64 `json:"final_cost"`
	Wallets Object `json:"wallets"`
	OrgMembers []Object `json:"org_members"`
}

type _InitiativeCreateDataAttributes InitiativeCreateDataAttributes

// NewInitiativeCreateDataAttributes instantiates a new InitiativeCreateDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInitiativeCreateDataAttributes(name string, desc string, goal string, type_ string, status string, finalCost int64, wallets Object, orgMembers []Object) *InitiativeCreateDataAttributes {
	this := InitiativeCreateDataAttributes{}
	this.Name = name
	this.Desc = desc
	this.Goal = goal
	this.Type = type_
	this.Status = status
	this.FinalCost = finalCost
	this.Wallets = wallets
	this.OrgMembers = orgMembers
	return &this
}

// NewInitiativeCreateDataAttributesWithDefaults instantiates a new InitiativeCreateDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInitiativeCreateDataAttributesWithDefaults() *InitiativeCreateDataAttributes {
	this := InitiativeCreateDataAttributes{}
	return &this
}

// GetName returns the Name field value
func (o *InitiativeCreateDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InitiativeCreateDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InitiativeCreateDataAttributes) SetName(v string) {
	o.Name = v
}

// GetDesc returns the Desc field value
func (o *InitiativeCreateDataAttributes) GetDesc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Desc
}

// GetDescOk returns a tuple with the Desc field value
// and a boolean to check if the value has been set.
func (o *InitiativeCreateDataAttributes) GetDescOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Desc, true
}

// SetDesc sets field value
func (o *InitiativeCreateDataAttributes) SetDesc(v string) {
	o.Desc = v
}

// GetGoal returns the Goal field value
func (o *InitiativeCreateDataAttributes) GetGoal() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Goal
}

// GetGoalOk returns a tuple with the Goal field value
// and a boolean to check if the value has been set.
func (o *InitiativeCreateDataAttributes) GetGoalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Goal, true
}

// SetGoal sets field value
func (o *InitiativeCreateDataAttributes) SetGoal(v string) {
	o.Goal = v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *InitiativeCreateDataAttributes) GetLocation() string {
	if o == nil || IsNil(o.Location) {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InitiativeCreateDataAttributes) GetLocationOk() (*string, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *InitiativeCreateDataAttributes) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *InitiativeCreateDataAttributes) SetLocation(v string) {
	o.Location = &v
}

// GetType returns the Type field value
func (o *InitiativeCreateDataAttributes) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InitiativeCreateDataAttributes) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InitiativeCreateDataAttributes) SetType(v string) {
	o.Type = v
}

// GetStatus returns the Status field value
func (o *InitiativeCreateDataAttributes) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *InitiativeCreateDataAttributes) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *InitiativeCreateDataAttributes) SetStatus(v string) {
	o.Status = v
}

// GetFinalCost returns the FinalCost field value
func (o *InitiativeCreateDataAttributes) GetFinalCost() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.FinalCost
}

// GetFinalCostOk returns a tuple with the FinalCost field value
// and a boolean to check if the value has been set.
func (o *InitiativeCreateDataAttributes) GetFinalCostOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FinalCost, true
}

// SetFinalCost sets field value
func (o *InitiativeCreateDataAttributes) SetFinalCost(v int64) {
	o.FinalCost = v
}

// GetWallets returns the Wallets field value
func (o *InitiativeCreateDataAttributes) GetWallets() Object {
	if o == nil {
		var ret Object
		return ret
	}

	return o.Wallets
}

// GetWalletsOk returns a tuple with the Wallets field value
// and a boolean to check if the value has been set.
func (o *InitiativeCreateDataAttributes) GetWalletsOk() (*Object, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Wallets, true
}

// SetWallets sets field value
func (o *InitiativeCreateDataAttributes) SetWallets(v Object) {
	o.Wallets = v
}

// GetOrgMembers returns the OrgMembers field value
func (o *InitiativeCreateDataAttributes) GetOrgMembers() []Object {
	if o == nil {
		var ret []Object
		return ret
	}

	return o.OrgMembers
}

// GetOrgMembersOk returns a tuple with the OrgMembers field value
// and a boolean to check if the value has been set.
func (o *InitiativeCreateDataAttributes) GetOrgMembersOk() ([]Object, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrgMembers, true
}

// SetOrgMembers sets field value
func (o *InitiativeCreateDataAttributes) SetOrgMembers(v []Object) {
	o.OrgMembers = v
}

func (o InitiativeCreateDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InitiativeCreateDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["desc"] = o.Desc
	toSerialize["goal"] = o.Goal
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	toSerialize["type"] = o.Type
	toSerialize["status"] = o.Status
	toSerialize["final_cost"] = o.FinalCost
	toSerialize["wallets"] = o.Wallets
	toSerialize["org_members"] = o.OrgMembers
	return toSerialize, nil
}

func (o *InitiativeCreateDataAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"desc",
		"goal",
		"type",
		"status",
		"final_cost",
		"wallets",
		"org_members",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varInitiativeCreateDataAttributes := _InitiativeCreateDataAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInitiativeCreateDataAttributes)

	if err != nil {
		return err
	}

	*o = InitiativeCreateDataAttributes(varInitiativeCreateDataAttributes)

	return err
}

type NullableInitiativeCreateDataAttributes struct {
	value *InitiativeCreateDataAttributes
	isSet bool
}

func (v NullableInitiativeCreateDataAttributes) Get() *InitiativeCreateDataAttributes {
	return v.value
}

func (v *NullableInitiativeCreateDataAttributes) Set(val *InitiativeCreateDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableInitiativeCreateDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableInitiativeCreateDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInitiativeCreateDataAttributes(val *InitiativeCreateDataAttributes) *NullableInitiativeCreateDataAttributes {
	return &NullableInitiativeCreateDataAttributes{value: val, isSet: true}
}

func (v NullableInitiativeCreateDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInitiativeCreateDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


