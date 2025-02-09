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

// checks if the InitiativeIncluded type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InitiativeIncluded{}

// InitiativeIncluded struct for InitiativeIncluded
type InitiativeIncluded struct {
	Wallets Wallets `json:"wallets"`
	Organization []AddOrgMember `json:"organization"`
}

type _InitiativeIncluded InitiativeIncluded

// NewInitiativeIncluded instantiates a new InitiativeIncluded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInitiativeIncluded(wallets Wallets, organization []AddOrgMember) *InitiativeIncluded {
	this := InitiativeIncluded{}
	this.Wallets = wallets
	this.Organization = organization
	return &this
}

// NewInitiativeIncludedWithDefaults instantiates a new InitiativeIncluded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInitiativeIncludedWithDefaults() *InitiativeIncluded {
	this := InitiativeIncluded{}
	return &this
}

// GetWallets returns the Wallets field value
func (o *InitiativeIncluded) GetWallets() Wallets {
	if o == nil {
		var ret Wallets
		return ret
	}

	return o.Wallets
}

// GetWalletsOk returns a tuple with the Wallets field value
// and a boolean to check if the value has been set.
func (o *InitiativeIncluded) GetWalletsOk() (*Wallets, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Wallets, true
}

// SetWallets sets field value
func (o *InitiativeIncluded) SetWallets(v Wallets) {
	o.Wallets = v
}

// GetOrganization returns the Organization field value
func (o *InitiativeIncluded) GetOrganization() []AddOrgMember {
	if o == nil {
		var ret []AddOrgMember
		return ret
	}

	return o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value
// and a boolean to check if the value has been set.
func (o *InitiativeIncluded) GetOrganizationOk() ([]AddOrgMember, bool) {
	if o == nil {
		return nil, false
	}
	return o.Organization, true
}

// SetOrganization sets field value
func (o *InitiativeIncluded) SetOrganization(v []AddOrgMember) {
	o.Organization = v
}

func (o InitiativeIncluded) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InitiativeIncluded) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["wallets"] = o.Wallets
	toSerialize["organization"] = o.Organization
	return toSerialize, nil
}

func (o *InitiativeIncluded) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"wallets",
		"organization",
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

	varInitiativeIncluded := _InitiativeIncluded{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInitiativeIncluded)

	if err != nil {
		return err
	}

	*o = InitiativeIncluded(varInitiativeIncluded)

	return err
}

type NullableInitiativeIncluded struct {
	value *InitiativeIncluded
	isSet bool
}

func (v NullableInitiativeIncluded) Get() *InitiativeIncluded {
	return v.value
}

func (v *NullableInitiativeIncluded) Set(val *InitiativeIncluded) {
	v.value = val
	v.isSet = true
}

func (v NullableInitiativeIncluded) IsSet() bool {
	return v.isSet
}

func (v *NullableInitiativeIncluded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInitiativeIncluded(val *InitiativeIncluded) *NullableInitiativeIncluded {
	return &NullableInitiativeIncluded{value: val, isSet: true}
}

func (v NullableInitiativeIncluded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInitiativeIncluded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


