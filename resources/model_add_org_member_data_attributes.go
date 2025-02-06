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

// checks if the AddOrgMemberDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddOrgMemberDataAttributes{}

// AddOrgMemberDataAttributes struct for AddOrgMemberDataAttributes
type AddOrgMemberDataAttributes struct {
	// organization id
	OrgId string `json:"org_id"`
	// status of member
	Status string `json:"status"`
}

type _AddOrgMemberDataAttributes AddOrgMemberDataAttributes

// NewAddOrgMemberDataAttributes instantiates a new AddOrgMemberDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddOrgMemberDataAttributes(orgId string, status string) *AddOrgMemberDataAttributes {
	this := AddOrgMemberDataAttributes{}
	this.OrgId = orgId
	this.Status = status
	return &this
}

// NewAddOrgMemberDataAttributesWithDefaults instantiates a new AddOrgMemberDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddOrgMemberDataAttributesWithDefaults() *AddOrgMemberDataAttributes {
	this := AddOrgMemberDataAttributes{}
	return &this
}

// GetOrgId returns the OrgId field value
func (o *AddOrgMemberDataAttributes) GetOrgId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *AddOrgMemberDataAttributes) GetOrgIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value
func (o *AddOrgMemberDataAttributes) SetOrgId(v string) {
	o.OrgId = v
}

// GetStatus returns the Status field value
func (o *AddOrgMemberDataAttributes) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AddOrgMemberDataAttributes) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *AddOrgMemberDataAttributes) SetStatus(v string) {
	o.Status = v
}

func (o AddOrgMemberDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddOrgMemberDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["org_id"] = o.OrgId
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *AddOrgMemberDataAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"org_id",
		"status",
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

	varAddOrgMemberDataAttributes := _AddOrgMemberDataAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddOrgMemberDataAttributes)

	if err != nil {
		return err
	}

	*o = AddOrgMemberDataAttributes(varAddOrgMemberDataAttributes)

	return err
}

type NullableAddOrgMemberDataAttributes struct {
	value *AddOrgMemberDataAttributes
	isSet bool
}

func (v NullableAddOrgMemberDataAttributes) Get() *AddOrgMemberDataAttributes {
	return v.value
}

func (v *NullableAddOrgMemberDataAttributes) Set(val *AddOrgMemberDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAddOrgMemberDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAddOrgMemberDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddOrgMemberDataAttributes(val *AddOrgMemberDataAttributes) *NullableAddOrgMemberDataAttributes {
	return &NullableAddOrgMemberDataAttributes{value: val, isSet: true}
}

func (v NullableAddOrgMemberDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddOrgMemberDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


